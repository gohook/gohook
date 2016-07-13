# Gohook

Run anything with a webhook

## What is it?

Gohook is a command line tool for executing commands, programs, and scripts when a webhook is hit.

## CLI commands

- `start` - Runs the client to connect and listen for webhooks from the server
- `ls [<id>]` - List all webhook commands
- `add [--output] [--params] <command>` - Creates a new webhook command
- `remove <id>` - Removes an existing webhook command

## Examples

Create a basic webhook script

```
$ gohook add ./my-script
Webhook created:

'------------'-------------------------------------'---------------------------------'
| ID         | Webhook                             | Command                         |
'------------'-------------------------------------'---------------------------------'
| 8a1ab948c2 | https://gohook.io/begizi/8a1ab948c2 | /Users/begizi/scripts/my-script |
'------------'-------------------------------------'---------------------------------'
```

Create a basic command line webhook

```
$ gohook add "date > ~/dates.txt"
Webhook created:

'------------'-------------------------------------'--------------------'
| ID         | Webhook                             | Command            |
'------------'-------------------------------------'--------------------'
| 3b7ca81a29 | https://gohook.io/begizi/3b7ca81a29 | date > ~/dates.txt |
'------------'-------------------------------------'--------------------'
```

Create a webhook with stdin params

```
$ gohook add --params=name,greeting --output="~/output.txt" echo "Hello, $name. $greeting"
Webhook created:

'------------'-------------------------------------'---------------------------------'--------------'----------------'
| ID         | Webhook                             | Command                         | Output       | Params         |
'------------'-------------------------------------'---------------------------------|--------------'----------------'
| 877a238a7f | https://gohook.io/begizi/877a238a7f | echo "Hello, $name. $greeting"  | ~/output.txt | name, greeting |
'------------'-------------------------------------'---------------------------------'--------------'----------------'
```

Remove a webhook

```
$ gohook remove 877a238a7f
Webhook removed.
```

List all webhooks

```
$ gohook ls
Webhooks:

'------------'-------------------------------------'---------------------------------'--------------'--------------------------------'
| ID         | Webhook                             | Command                         | Output       | Params                         |
'------------'-------------------------------------'---------------------------------|--------------'--------------------------------'
| 877a238a7f | https://gohook.io/begizi/877a238a7f | echo "Hello, $name. $greeting"  | ~/output.txt | name, greeting                 |
'------------'-------------------------------------'---------------------------------'--------------|--------------------------------'
| 3b7ca81a29 | https://gohook.io/begizi/3b7ca81a29 | date > ~/dates.txt              |              |                                |
'------------'-------------------------------------'---------------------------------'--------------|--------------------------------'
| 8a1ab948c2 | https://gohook.io/begizi/8a1ab948c2 | /Users/begizi/scripts/my-script |              |                                |
'------------'-------------------------------------'---------------------------------'--------------'--------------------------------'
```

Start the client

```
$ gohook start
Listening for webhooks..
```

Start the client in the background

```
$ gohook start &
```
