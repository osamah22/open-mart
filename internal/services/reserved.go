package services

// a list of some usernames that are not allowed for normal users
var reservedUsernames = map[string]struct{}{
	"admin": {}, "administrator": {}, "root": {}, "system": {}, "support": {},
	"help": {}, "contact": {}, "info": {}, "staff": {}, "mod": {}, "moderator": {},
	"owner": {}, "superuser": {}, "security": {}, "user": {}, "users": {},
	"api": {}, "docs": {}, "static": {}, "assets": {}, "public": {}, "private": {},
	"internal": {}, "server": {}, "status": {}, "config": {}, "settings": {},
	"about": {}, "terms": {}, "privacy": {}, "login": {}, "logout": {},
	"signup": {}, "register": {}, "auth": {}, "oauth": {}, "session": {},
	"sessions": {}, "profile": {}, "dashboard": {}, "home": {}, "index": {},
	"search": {}, "notifications": {}, "messages": {}, "inbox": {}, "outbox": {},

	// impersonation blockers
	"mail": {}, "email": {}, "noreply": {}, "no-reply": {}, "billing": {},
	"payments": {}, "payment": {}, "stripe": {}, "paypal": {}, "cloudflare": {},
	"github": {}, "google": {}, "meta": {}, "facebook": {}, "twitter": {},
	"discord": {}, "bot": {}, "assistant": {}, "helpdesk": {}, "support-team": {},

	// technical
	"null": {}, "nil": {}, "undefined": {}, "unknown": {}, "test": {}, "testing": {},
	"example": {}, "demo": {}, "staging": {}, "production": {}, "dev": {},
	"developer": {}, "development": {}, "release": {}, "build": {}, "version": {},
	"new": {}, "old": {}, "backup": {}, "archive": {}, "temp": {}, "tmp": {},
}

func isReserved(username string) bool {
	_, exists := reservedUsernames[username]
	return exists
}
