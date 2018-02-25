#!/usr/bin/python3

IPs = open("src/IPs.log").readlines()
uniqueIPs = set(IPs)
open("src/unique.IPs.log", "w").writelines(uniqueIPs)

users = open("src/users.log").readlines()
uniqueUsers = set(users)
open("src/unique.users.log", "w").writelines(uniqueUsers)