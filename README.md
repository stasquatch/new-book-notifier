# new-book-notifier

The new book notifier works by using a cron job weekly. It uses the Google Books API to check a list of authors' recent work. If something's been published in the last week (since it previously ran), it'll output a string that the cronjob will use to email yourself. Hacky, but simple and all local.

## Updating author list

Edit the slice `authorArr` in `main.go`.

## Setting up the cron job

Edit cron jobs:

```sh
crontab -e
```

Note: you'll probably need to change the file path and [cron scheduling](https://crontab.guru/)

```sh
*/10 10 * * 0 cd ~/Code/new-book-notifier && /usr/local/bin/go run main.go | mail -s "New book alert" stacyharrison
```

### Setting up your email forwarding

Run `cat ~/.forward` - if the content of the file is your email address, you're done setting up email forwarding. Otherwise, proceed.

```sh
touch ~/.forward
echo "YOUR EMAIL GOES HERE" >> ~/.forward
```

The email generated from the cron job will likely end up in spam the first time. You can send a test email like with `echo "hello world" | mail -s testing <YOUR MAC USER NAME>`. If you're using gmail, create a filter for your email to mark it as "Never send to spam".
