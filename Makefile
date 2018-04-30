# Target
TARGET := veth-peer

# Source files
SRCS := veth-peer.go


.PHONY: $(TARGET)
$(TARGET): 
	go build -o $(TARGET)


.PHONY: install
install: $(TARGET)
	install -c -m 555 -o root -g root veth-peer /usr/bin
	install -c -m 444 -o root -g root veth-peer.man /usr/share/man/man1

.PHONY: clean
clean: 
	@rm -f $(TARGET) *~
