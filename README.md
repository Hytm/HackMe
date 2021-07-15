# HackMe

## Purpose

This application lets you add message to a list, like a Guiestbook for example.

There is a dedicated endpoint where you can find a way to run arbitrary command line on the host.

This is an easily portable application you can run on any host to demonstrate different scenario of attack, and run them in a customer environment.

## How to use it

You can use the Docker version with the dockerfile. It'll build the app and create the container after that. The file is compatible with GCP Cloud Run.

Otherwise you can run the linux binary directly (hackme file).

You have 2 endpoints available: 
* http://{yourhost}:{port}/ is the app by itself
* http://{yourhost}:{port}/admin is where you can start hacking

## Stories

You can use the stories from examples folder, or build you own.

## TODOs

1. I'm working on a deployable hydra package that will fit. Just run a wget, unzip, and run the script inside the zip.
