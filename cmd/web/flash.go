package main

type Flash struct {
	Type    string // "error", "info", "success", "warning"
	Message string
}

const FlashKey = "flash_messages"
