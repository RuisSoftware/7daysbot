# 7 Days to Die Discord bot

7daysbot is a Discord bot for 7 Days to Die written in Go that integrates in-game chat with Discord. The bot also allows for various commands to be run against the 7 Days to Die server.

This bot is a stand-alone application, no game mods are required.

This bot is currently written for Alpha 19 builds, future game updates may cause issues.

# Commands

The following commands are supported by the bot:

`!info`
`!time`
`!players`

# Installation
1. Download the appropriate release for your operating system
2. Extract somewhere on your server's system
3. Using config-sample.json as a template, create a config.json file.
4. Edit the config.json file setting appropriate discord, telnet and game settings. See 'How to set up bot account' section for info on how to generate discord keys.
5. Start the bot using either the executable or startup.sh on linux systems.
    - startup.sh executes the linux binary under the 'screen' terminal multiplexer. Use the command `screen -r 7bot' to resume the session. Ctrl+A then D to detach from session or Ctrl+C to exit the bot.  

# Configuration
- consolelogging : Log telnet data to console. NOTE: This will expose your telnet password to the console.
- Game
    - bloodmoonfrequency : Number of days between blood moons. This should match the BloodMoonFrequency setting in your server configuration.
    - DayLightLength : Should match setting of same name in server config.ini. How many hours of daylight per day.
- Discord
    - token : Discord Bot token
    - channel : Discord channel ID for bot to be linked to (Right click on channel in discord and copy ID to get the correct ID)
    - prefix : Prefix for bot commands in discord. Defaults to '!' for commands such as !time
- Telnet
    - ip : IP address of 7 Days to Die server. If running bot on same machine set to "localhost"
    - port : Port number for 7 Days to Die server telnet. Needs to be in string format (surround in double quotes). Default port is "8081"
    - password : Telnet password. Check 7 Days to Die server configuration.

## How to set up bot account
1. Log in to the [Discord Developer Portal](https://discordapp.com/developers/applications/) in a browser and click "New Application". 
2. Give the bot a name and click Create. 
3. Store the client ID as you will need it in later steps.
4. In the left hand menu, click "Bot". Now click the "Add Bot" button to create your bot. Set an avatar for your bot if desired.
5. After the bot is created, find the toggle labeled "Public bot". Make sure that this is **OFF**. Click "Save Changes" to confirm.
6. Click the link that says "Click to Reveal Token", this token is the password for the bot discord account. This token should *never* be shared with anyone. Store this token somewhere as it will be needed later and in your config.json file.
    - If the token is accidentally leaked, click the "Regenerate" button as soon as possible. This will revoke the existing token and generate a new one. You will need the new token in config.json.
7. Click 'Save Changes' if presented.
8. On the left hand menu, click "OAuth2"
9. Select the "bot" checkbox under "Scopes"
10. Copy and paste the resulting URL into a browser window. Choose a server to invite the bot to, and click "Authorize". Your bot should now be added to your Discord server.

# Currently known issues
- Bot does not handle game server restarts gracefully. If the server is restarted or crashes, bot needs to be restarted.
- Bot may not handle Discord outages gracefully in certain conditions. Restart of bot required to resolve.