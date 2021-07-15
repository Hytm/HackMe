# HackMe

## Purpose

This application lets you add message to a list, like a Guiestbook for example.
It is intended to simulate the use of the [eval command on Redis](https://redis.io/commands/eval).

There is a dedicated endpoint where you can find a way to run arbitrary command line on the host.

This is an easily portable application you can run on any host to demonstrate different scenario of attack, and run them in a customer environment.

## Story

This application lets you write messages, whatever they are (it is certainly possible to hack the message form by itself).
This is entirely managed in-memory, so you decided to add endpoints where you can check your log file and decide to load messages directly from the log in case of a restart.

Hitting the http://{yourhost}:{port}/ will let you add messages. 

You also have a dedicated endpoint, without any link for security reason, to access your log: http://{yourhost}:{port}/admin .

By checking common endpoints, an attacker can easily access this "hidden" endpoint. He'll then discover how the application is running command (like the eval command) on the host.

From there, you can imagine whatever you want to have an alert in Lacework.

## TODOs

1. I'm working on a deployable hydra package that will fit. Just run a wget, unzip, and run the script inside the zip.
2. Bundle html files in the go package for usability
