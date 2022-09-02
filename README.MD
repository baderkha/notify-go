# Notify-Go

 <p align="center"><img src="./assets/logo/logo_new_1.png?raw=true" width="256px" align="center"/> </p>


Notify Go is a **free** command line / API that allows you to easily send (bot style) messages to a recipient . This features a base implementation for all your favourite messaging platforms 

like :
-  telegram 
-  discord 
-  slack 



**Note**: 
*this is not a 2 way messaging protocol . This project only publishes messages*

Also , I would like to mention ,  Slack and Discord both use `webhooks` . So if you just want to use those and have simple needs . **NO NEED TO USE THIS SOLUTION** it's overkill. Just Send an http request to the webhook route with the body. ie curl it. 




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
- [Installation](#-installation-)
    - [API](#a-api-mode-installation)
    - [Command Line](#b-command-line-installation)
- [Authorization Setup](#-authorization-setup-) -- you cannot skip this step
- [API Usage](#-api-usage-)
- [Cli Usage](#-cli-usage-)


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

<h2 id="authorization-setup" align="center"> Authorization Setup </h2>

<p style="color:red"> *Note you must do this in order to use the cli/api*</p>


This section will cover how to setup authorization for each of the message senders . Note that 

- [Slack](https://api.slack.com/messaging/webhooks)
    - you need to create an app
    - give it permissions 
    - create a webhook 
    - use that webhook with this cli / api
- [Discord](https://support.discord.com/hc/en-us/articles/228383668-Intro-to-Webhooks)
    - go to your server settings
    - under integrations (create a webhook)
    - call it a cool name
    - use that webhook with this cli / api
- [Telegram]()
    - // todo




---

 <h2 align="center"> API Usage </h2> 
 This section will cover how to use the api.

 *Note You Need Go **v1.18+*** Since The API uses generics

## MessageSender
`notify.MessageSender` is an interface which has implementations for slack,discord,and telegram. If your usage is simple and you only want to message Discord for example . Then attaching this interface is perfect for you.





### Slack Instructions  

1) init
``` go
slackSender := notify.NewSlackSender()
```

2) send a message to a webhook

``` go
err := slackSender.Send("https://slack.channel.com/your/other/channel/webhook",[]byte("some_message"))
if err != nil {
    log.Fatal(err)
}

```



### Discord Instructions
Discord is very similar to slack since it also uses webhooks to send messages to a chat . So the init logic is the same .

1) init
``` go
discSender := notify.NewDiscordSender()
```

2) send a message to a different webhook

``` go
err := discSender.Send("https://slack.channel.com/your/other/channel/webhook",[]byte("somemessage"))
if err != nil {
    log.Fatal(err)
}

```
<br>
<br>

## Manager

`notify.Manager` enables you to orchestrate sending messages to multiple social platforms . This should be used only if you expect your application to use more than 1 notification platform . 

[Otherwise see](#MessageSender)

1) init

``` go
// default contains all the clients
msgMgr := notify.Default()
```

2) send a message to a specific client (discord , slack , telegram)

``` go
// send to a specific client like discord
err := msgMgr.SendToSpecificType(
    notify.DiscordSenderType,
    "https://www.some-webhook.url.com",
    []byte("hi mom"),
)
```

3) broadcast same message to all clients

``` go
// 1 - create a mapping
ralias := notify.NewEmptyRecieverAlias()
// returns an error 
_ = ralias.Add(notfy.DiscordSenderType,"https://www.google.com")
_ = ralias.Add(notify.SlackSenderType,"https://www.google.com")

// 2 - send msg
msgMgr.SendAll(ralias,[]byte("hi mom "))

```
<br>
<br>


## Implement your own sender

To implement your own Message publisher client you have to implement the `notify.Sender` Method and then add your sender to the manager


``` go
Send(reciever string, bodyContent []byte) error
```

Example : 

1) Implement the interface
```go 
// equivalent of implements keyword
var _ notify.Sender = &WhatsappSender{}

type WhatsappSender struct {}

func (w *WhatsappSender) Send(reciver string , bodyContent[]byte) error {
    return nil
}
```

2) [Optional] Enroll it to the manager if you're using it in tandum with other services
``` go
// your main go
func main() {
    mgr := notify.Default()
    mgr.AddSender("whatsapp",new(WhatsappSender))

    _ = msgMgr.SendToSpecificType(
        "whatsapp",
        "chat_id",
        []byte("hi mom"),
    )

}
```

----

<h2 align="center"> Cli Usage </h2>

### Help

```bash
notify-go --help
```

<br>

### Managing Contacts

Contacts allow you to map channels / entities to different social profiles . IE you can map your 1 channel to discord , slack , telegram

This is **powerful** , because you can leverage the broadcast functionality of the cli to send to all those channel with 1 alias

Example : 

- Add Contact Entry
    ```bash
    notify-go newcon crypto_channel
    ```

- Add Contact Social Mapping

    this will make it easy for you to reference by name
    rather than have to repaste the webhook everytime

    ```bash
    notify-go apcon crypto_channel discord https://www.google.com
    notify-go apcon crypto_channel slack https://www.google.com
    ```

- List Contacts

    Allows you to list all / 1 contact

    - all
        ```bash
        notify-go cons 
        ```
    - specific
        ```bash
        notify-go cons crypto_channel
        ```
<br>
<br>


### Sending Messages

- with webhook
    ```bash
    notify-go msg discord https://webhook.com "hi mom"
    ```
- through contact alias (see above section for how to create contacts)
    ```bash
    notify-go msg discord crypto_channel "hi mom"
    ```
- to all social platforms via contact mapping*
    in this scenario we setup mapping for crypto_channel to discord
    and slack . So this will send to **both** channels concurrently
    ```bash
    notify-go msgcon crypto_channel "hi mom"
    ```


### Broadcasting messages

***Note** this will message everyone in the contact list and all platforms. use this with caution :)*

```bash
notify-go msgbrod "hi mom"
```



---

[^1]: By Ahmad Baderkhan

[^2]: License *Apache-V2*
