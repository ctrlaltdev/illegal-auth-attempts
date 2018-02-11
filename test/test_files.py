#!/usr/bin/python3

from pathlib import Path

def fileExists(source):
  file = Path(source)
  return file.is_file()

def test_IPs():
  assert fileExists("src/IPs.log") == True

def test_Users():
  assert fileExists("src/users.log") == True