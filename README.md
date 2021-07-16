# HackMe

## Purpose

This application lets you add message to a list, like a Guiestbook for example.

There is a dedicated endpoint where you can find a way to run arbitrary command line on the host.

This is an easily portable application you can run on any host to demonstrate different scenario of attack, and run them in a customer environment.

## How to use it

You can use the Docker version with the dockerfile. This is a 2 steps docker file, so it'll build the app in a container and create the container for the app after that. This is compatible with GCP Cloud Run (you can attach it to continuous deployment in Cloud Run).

Otherwise you can run the linux binary directly (hackme file).

You have 2 endpoints available: 
* http://{yourhost}:{port}/ is the app by itself
* http://{yourhost}:{port}/admin is where you can start hacking

## Stories

You can use the stories from examples folder, or build your own.

## TODOs

1. I'm working on a deployable hydra package that will fit. Just run a wget, unzip, and run the script inside the zip.
