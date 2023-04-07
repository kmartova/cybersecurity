# Passive
Introduction

Information gathering is one of the most important steps during a pentest, it can be considered the longest step.
Objective

The goal of this project is for you to become more comfortable with open source investigative methods
Advice

Before asking help, ask yourself if you have really thought about all the possibilities.
https://en.kali.tools/all/?category=recon
https://github.com/topics/osint-tools
https://en.wikipedia.org/wiki/Open-source_intelligence
https://en.wikipedia.org/wiki/Doxing
Guidelines

You are going here to create your first passive recognition tool, you have the choice of language, however your program will have to recognize the information entered (FULL NAME, IP, @login).

For the case of the full name, it will have to recognize the entry: Last name First Name, then look in the directories for the telephone number and the address.

If it is the ip address, your tool should display at least the city and the name of the internet service provider.

If it is an username, your tool will have to check if this username is used in at least 5 known social networks.

The result should be stored in a result.txt file (result2.txt if the file already exists)

## Usage
go run cmd/main.go [options] + FullName || IP || @login

