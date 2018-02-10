#!/usr/bin/python3

from operator import itemgetter

IPs = open("src/unique.IPs.count.log").readlines()

IPcounts = []
for IP in IPs:
  ip, count = IP.split(":")
  IPcounts.append((ip, int(count)))

with open("src/unique.IPs.count.log", "w") as IPsfile:
  for ip, count in sorted(IPcounts, key=itemgetter(1), reverse=True):
    IPsfile.write("{}:{}\n".format(ip, count))

users = open("src/unique.users.count.log").readlines()

Usercounts = []
for user in users:
  u, count = user.split(":")
  Usercounts.append((u, int(count)))

with open("src/unique.users.count.log", "w") as Usersfile:
  for u, count in sorted(Usercounts, key=itemgetter(1), reverse=True):
    Usersfile.write("{}:{}\n".format(u, count))