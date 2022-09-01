# Notify-Go

 <p align="center"><img src="./assets/logo/logo_new_1.png?raw=true" width="256px" align="center"/> </p>


Notify Go is a **free** command line / API that allows you to easily send (bot style) messages to a recipient . This features a base implementation for all your favourite messaging platforms 

like :
-  telegram 
-  discord 
-  slack 

**Note**: *this is not a 2 way messaging protocol . This project only publishes messages*

It's aimed to be super minimal and simple to configure.

This solution is great for the following scenarios :
- Command Line Notifications
- System Notifications
- IOT Devices
- Home Automation
- Anything that can run linux .

----

<h2 align="center"> Contents </h2>

Table of contents
- [Usage](#Notify-go)
- [Installation](#Installation)
    - [API](#a-api-mode-installation)
    - [Command Line](#b-command-line-installation)
- [API Usage](#api)
- [Cli Usage](#cli-usage)
- Setting Up Messaging Auth
    - something


---

<h2 align="center"> Installation </h2>

## A) API Mode Installation

To use the API in your existing go code simply run the following
```bash 
go get -u github.com/baderkha/notify-go
```

## B) Command Line Installation

// todo


---

 <h2 align="center"> API Usage </h2> 
 This section will cover how to use the api.

## Manager

`notify.Manager` enables you to orchestrate sending messages to multiple social platforms . This should be used only if you expect your application to use more than 1 notification platform . 

[Otherwise see](#MessageSender)


## MessageSender:
`notify.MessageSender` is an interface which has implementations for slack,discord,and telegram. If your usage is simple and you only want to message Discord for example . Then attaching this interface is perfect for you.





### Slack Instructions  


1) init
``` go
slackSender := notify.NewSlackSender(&notify.SlackConfig{
    DefaultChannelWebhook : "https://slack.channel.com/somelongauthtoken/embedded"
})
```

2) send a message to default webhook

``` go
// this will send to 
// https://slack.channel.com/somelongauthtoken/embedded
err := slackSender.SendToDefaultReciever()
if err != nil {
    log.Fatal(err)
}
```

3) send a message to a different webhook

``` go
err := slackSender.SendToReciever("https://slack.channel.com/your/other/channel/webhook")
if err != nil {
    log.Fatal(err)
}

```



### Discord
Discord is very similar to slack since it also uses webhooks to send messages to a chat . So the init logic is the same .

1) init
``` go
discSender := notify.NewDiscordSender(&notify.DiscordConfig{
    DefaultChannelWebhook : "https://slack.channel.com/somelongauthtoken/embedded"
})
```

2) send a message to default webhook

``` go
// this will send to 
// https://slack.channel.com/somelongauthtoken/embedded
err := discSender.SendToDefaultReciever()
if err != nil {
    log.Fatal(err)
}
```

3) send a message to a different webhook

``` go
err := discSender.SendToReciever("https://slack.channel.com/your/other/channel/webhook")
if err != nil {
    log.Fatal(err)
}

```


----

<h2 align="center"> Cli Usage </h2>

// todo

---

<h2 align="center"> Setting Up Messaging Auth </h2>


This section will cover how to setup authorization for each of the message senders . Note that 




---

By *Ahmad Baderkhan*

License *Apache-V2*
