# Illegal Auth Attempts

This set of scripts aims to extract from auth attempts or device scanning IPs and users used for those auth attempts.

## IPS

The top 10 IPs are:

| IP              | Count |
| --------------- | -----:|
| 61.144.96.98    | 21704 |
| 1.51.198.195    | 21704 |
| 123.183.209.139 | 10079 |
| 61.177.172.64   | 5505  |
| 59.63.166.104   | 5385  |
| 18.217.140.251  | 5064  |
| 36.189.242.19   | 4854  |
| 59.63.188.32    | 4788  |
| 59.63.188.36    | 4466  |
| 59.63.166.105   | 4027  |

## Users

The top 10 users are:

| User    | Count |
| ------- | -----:|
| admin   | 2719  |
| test    | 1728  |
| user    | 1562  |
| guest   | 1339  |
| pi      | 418   |
| support | 306   |
| oracle  | 252   |
| ftpuser | 246   |
| ubnt    | 237   |
| nagios  | 225   |

## Files

The uncompressed files are available in src/ and the .tar.gz are in archives/

## How and what

### Fetching

The `fetch.sh` script get from /var/log/auth.log the IPs and users of the previous day lines of the log. Hence it has to be run only once a day to get everything and to not duplicate data.

Moreover, for it to work, the cron has to be able to read /var/log/auth.log and write in /var/log/

### Importing

If you're fetching IPs and Users on several devices and want to centralize everything on one, you can put your auth.log or secure files in import/sources.
Then `cp import/sources/filetoimport import/import` then `cd import && ./import.sh`
This will append the IPs and Users to src/IPs.log and src/users.log

### Counting and sorting

Once enough data gathered, and the IPs.log and users.log moved or copied to src/, the `order.sh` will create unique IPs and users lists, as well as lists with count of their occurences in the original logs, sorted descendingly.

### Prerequisites

- An /var/log/auth.log (or the fetch script will have to be adaptated to your auth logging)
- Python 3

### Contributing

You can run this script on your public facing devices to collect the IPs and users too, and if you want to contribute, please refer to Import section.
Once you're done, run `./order.sh` and `./archive.sh` and create a pull request.

Not that you will need git lfs for src/ and archives/
