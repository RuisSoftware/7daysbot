package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/reiver/go-oi"
	"github.com/reiver/go-telnet"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	BotCaller telnet.Caller = internalBotCaller{}
	writer    telnet.Writer
	reader    telnet.Reader
	ctx       telnet.Context
)

type internalBotCaller struct{}

func (caller internalBotCaller) CallTELNET(ctxt telnet.Context, w telnet.Writer, r telnet.Reader) {
	writer = w
	reader = r
	ctx = ctxt
	botCallerCallTELNET(os.Stdin, os.Stdout, os.Stderr)
}

func handleOutput(line string) {

	if strings.Contains(line, "Please enter password:") {
		sendTelnet(config.Telnet.Password)
		return
	}

	if strings.HasPrefix(line, "Day ") {
		currentDayString := strings.Split(line, ",")[0][4:]
		currentDay, _ := strconv.Atoi(currentDayString)
		timeToBloodMoon := config.Game.BloodMoonFrequency - (currentDay % config.Game.BloodMoonFrequency)
		sendDiscordMessage(line)
		bloodMoonMessage := strconv.FormatInt(int64(timeToBloodMoon), 10) + " day(s) until blood moon."
		sendDiscordMessage(bloodMoonMessage)
		return
	}

	// Handle player list output
	r, _ := regexp.Compile("^\\d+. id=")
	if r.MatchString(line) {
		substr := strings.Split(line, ",")
		message := substr[1] + " (" + substr[10] + ")"
		sendDiscordMessage(message)
		return
	}

	// Handle no players in game
	matched, err := regexp.MatchString(` Total of 0 in the game`, line)
	if nil != err {
		return
	}
	if matched {
		message := " No players currently in game."
		sendDiscordMessage(message)
		return
	}

	// Handle Log in
	matched, err = regexp.MatchString(` INF PlayerSpawnedInWorld \(`, line)
	if nil != err {
		return
	}
	if matched {
		substr := strings.Split(line, ",")[6]
		pName := strings.TrimSuffix(strings.TrimSpace(substr[13:]), "'")
		message := pName + " has logged on"
		sendDiscordMessage(message)
		return
	}

	// Handle log out
	matched, err = regexp.MatchString(` INF Player disconnected:`, line)
	if nil != err {
		return
	}
	if matched {
		substr := strings.Split(line, ",")[3]
		pName := strings.TrimSuffix(strings.TrimSpace(substr[13:]), "'")
		message := pName + " has logged off"
		sendDiscordMessage(message)
		return
	}

	// Handle chat
	matched, err = regexp.MatchString(` INF Chat \(`, line)
	if nil != err {
		return
	}
	if matched {
		substr := strings.SplitAfterN(line[30:], ":", 2)[1]
		substr = strings.TrimLeft(substr, " ")

		// Prevent loopback chat
		if strings.HasPrefix(substr, "'Server'") {
			return
		}

		// remove quotes from playername
		substr = strings.Replace(substr, "'", "", 2)

		sendDiscordMessage(substr)
		return
	}

}

func sendTelnet(text string) {
	text += "\r\n"

	if config.Logging {
		fmt.Println(text)
	}

	byteText := []byte(text)
	p := byteText
	n, err := oi.LongWrite(writer, p)
	if nil != err {
		return
	}
	if expected, actual := int64(len(p)), n; expected != actual {
		err := fmt.Errorf("transmission problem: tried sending %d bytes, but actually only sent %d bytes", expected, actual)
		_, _ = fmt.Fprint(os.Stderr, err.Error())
		return
	}
}

func botCallerCallTELNET(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser) {
	go func(writer io.Writer, reader io.Reader) {
		var line bytes.Buffer
		linebreak := []byte("\n")

		var buffer [1]byte // Seems like the length of the buffer needs to be small, otherwise will have to wait for buffer to fill up.
		p := buffer[:]

		for {
			// Read 1 byte.
			n, err := reader.Read(p)
			if n <= 0 && nil == err {
				continue
			} else if n <= 0 && nil != err {
				break
			}

			if bytes.Equal(p, linebreak) {
				handleOutput(line.String())

				if config.Logging {
					fmt.Println(line.String())
				}

				// Clear buffer
				line.Reset()
			} else {
				line.Write(p)
			}
		}
	}(stdout, reader)

	var buffer bytes.Buffer
	var p []byte

	var crlfBuffer = [2]byte{'\r', '\n'}
	crlf := crlfBuffer[:]

	scanner := bufio.NewScanner(stdin)
	scanner.Split(scannerSplitFunc)

	for scanner.Scan() {
		buffer.Write(scanner.Bytes())
		buffer.Write(crlf)

		p = buffer.Bytes()

		n, err := oi.LongWrite(writer, p)
		if nil != err {
			break
		}
		if expected, actual := int64(len(p)), n; expected != actual {
			err := fmt.Errorf("transmission problem: tried sending %d bytes, but actually only sent %d bytes", expected, actual)
			_, _ = fmt.Fprint(stderr, err.Error())
			return
		}

		buffer.Reset()
	}

	// Wait a bit to receive data from the server (that we would send to io.Stdout).
	time.Sleep(3 * time.Millisecond)
}

func scannerSplitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF {
		return 0, nil, nil
	}

	return bufio.ScanLines(data, atEOF)
}

func Connect() {
	var caller = BotCaller

	err := telnet.DialToAndCall(config.Telnet.Ip+":"+config.Telnet.Port, caller)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
