# Style Logging Package

The `style` package provides a set of functions to log messages with various statuses and customizable prefixes, such as "Notice", "Documentation", and "Running". This package allows for formatted logging to the console and log files, supporting both simple and formatted messages, as well as conditional logging.

## Purpose

The purpose of the `style` package is to enhance logging by offering structured and formatted logging mechanisms, allowing developers to log messages with consistent styling and categorization. This is particularly useful in large applications where different kinds of messages (e.g., informational, errors, actions, and successes) need to be clearly distinguished for readability and debugging.

## Features

- Log messages with predefined statuses like `Notice`, `Documentation`, `Running`, `ThumbsUp`, `Failure`, and `Deleted`.
- Log formatted and simple strings with variadic arguments.
- Conditional logging based on boolean flags.
- Console printing with status-based prefixes and colors.
- Flexibility in message customization with format strings.

## Installation

To install the `style` package, you can simply include it in your project by copying the source code or adding it to your Go module.

```bash
go get [https://github.com/DanyelMorales/style](https://github.com/DanyelMorales/style)


## Usage

Below are some common examples of how to use the style package for logging and printing formatted messages.

### Basic Logging
``` go
package main

import "style"

func main() {
    style.OnReport("Starting the application")
    style.OnNotice("This is a notice")
    style.SuccessfulAction("Operation completed successfully")
    style.FailedAction("Operation failed due to an error")
}
```

### Formatted Logging

``` go
package main

import "style"

func main() {
    style.OnReportF("Task %d of %d is in progress", 2, 5)
    style.OnNoticeF("Notice for user ID: %s", "12345")
    style.SuccessfulActionF("User %s created successfully", "john_doe")
}
```

### Conditional Logging

``` go
package main

import "style"

func main() {
    // Conditionally log if the print flag is true
    printLog := true
    style.ActLogCF(printLog, "Action is running")
    style.OkLogCF(printLog, "Action completed successfully")
    style.ErrorLogCF(printLog, "Error occurred during execution")
}
```

Functions Overview

Simple Logging Functions

	OnReport(format string): Logs a simple message with the Documentation status.
	OnNotice(format string): Logs a simple message with the Notice status.
	OnAction(format string): Logs a simple message with the Running status.
	SuccessfulAction(format string): Logs a simple message indicating a successful action.
	FailedAction(format string): Logs a simple message indicating a failed action.
	ExitAction(format string): Logs a simple message indicating an action has ended.

Formatted Logging Functions

	OnReportF(format string, a ...interface{}): Logs a formatted message with the Documentation status.
	OnNoticeF(format string, a ...interface{}): Logs a formatted message with the Notice status.
	OnActionF(format string, a ...interface{}): Logs a formatted message with the Running status.
	SuccessfulActionF(format string, a ...interface{}): Logs a formatted message indicating a successful action.
	FailedActionF(format string, a ...interface{}): Logs a formatted message indicating a failed action.
	ExitActionF(format string, a ...interface{}): Logs a formatted message indicating an action has ended.

Conditional Logging Functions

	ActLogCF(print bool, format string, a ...interface{}): Conditionally logs a formatted message with the Running status.
	OkLogCF(print bool, format string, a ...interface{}): Conditionally logs a formatted message with the ThumbsUp status.
	ErrorLogCF(print bool, format string, a ...interface{}): Conditionally logs a formatted message with the Failure status.

Console Output Functions

	PrintF(st Enum, format string, a ...interface{}): Prints a formatted message to the console.
	Print(st Enum, format string): Prints a simple message to the console.

Core Logging Functions

	LogF(st Enum, format string, a ...interface{}): Logs a formatted message to the log.
	Log(st Enum, format string): Logs a simple message to the log.
	String(st Enum, format string): Returns a formatted string with the given status.
	StringF(st Enum, format string, a ...interface{}): Returns a formatted string with variadic arguments.

Customization

The style package can be customized by modifying the Config map to change the prefix and color of each status. This allows for flexible and consistent logging output.


### License
This package is provided as-is with no specific licensing. Feel free to modify and adapt it to your needs.
