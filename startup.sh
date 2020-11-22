#!/bin/sh
# Auto-restart
{
  echo "Launching... Use \"screen -r 7bot\" to access the terminal and ctrl+a+d to exit the terminal."
  screen -dmS 7bot ./7daysbot
} || {
  read -p "Failed to launch, see log above. Please make sure that \"screen\" is installed on your system. Press enter to exit."
}