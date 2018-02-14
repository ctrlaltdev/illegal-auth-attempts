# Illegal Auth Attempts

This set of scripts aims to extract from auth attempts or device scanning IPs and users used for those auth attempts.

## IPS

The top IPs are:

| IP              | Count |
| --------------- | -----:|
| 123.183.209.139 | 10079 |
| 61.177.172.64   | 5505  |
| 59.63.166.104   | 5385  |
| 18.217.140.251  | 5064  |
| 59.63.188.32    | 4788  |
| 59.63.188.36    | 4466  |
| 59.63.166.105   | 4027  |
| 123.183.209.140 | 3964  |
| 42.7.26.49      | 3649  |
| 103.99.0.188    | 3180  |

## Users

The top 10 users are:

| User    | Count |
| ------- | -----:|
| admin   | 2024  |
| user    | 1322  |
| test    | 1270  |
| guest   | 1189  |
| pi      | 278   |
| ftpuser | 199   |
| support | 180   |
| ubnt    | 146   |
| ftp     | 126   |
| 1234    | 76    |

## How and what

### Fetching

The `fetch.sh` script get from /var/log/auth.log the IPs and users of the previous day lines of the log. Hence it has to be run only once a day to get everything and to not duplicate data.

Moreover, for it to work, the cron has to be able to read /var/log/auth.log and write in /var/log/

### Counting and sorting

Once enough data gathered, and the IPs.log and users.log moved or copied to src/, the `order.sh` will create unique IPs and users lists, as well as lists with count of their occurences in the original logs, sorted descendingly.

### Prerequisites

- An /var/log/auth.log (or the fetch script will have to be adaptated to your auth logging)
- Python 3

### Contributing

You can run this script on your public facing devices to collect the IPs and users too, and if you want you can append your IPs log to src/IPs.log and your users.log to src/users.log, run `./order.sh` and `./archive.sh` and create a pull request.

Not that you will need git lfs for src/ and archives/
