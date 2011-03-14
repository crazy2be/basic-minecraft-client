include $(GOROOT)/src/Make.inc

TARG=minecraft-macro
GOFILES=\
types.go\
handlers.go\
senders.go\
main.go\
data.go\
chunks.go\
loggers.go\
player.go


include $(GOROOT)/src/Make.cmd