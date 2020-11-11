CC = clang
CFLAGS += -g -fno-inline
LDFLAGS += -lSDL2

RUNTIME_OBJS = \
	runtime/runtime.o

RUNTIME_ARCHIVE = runtime/runtime.a

$(RUNTIME_ARCHIVE): $(RUNTIME_OBJS)
	$(AR) r $@ $^

%.bc: %.ch8
	go run chairlift $^

%.s: %.bc
	llc $^

%.lifted: %.s $(RUNTIME_ARCHIVE)
	$(CC) $(LDFLAGS) $^ -o $@

clean:
	$(RM) *.bc $(RUNTIME_ARCHIVE) $(RUNTIME_OBJS)


