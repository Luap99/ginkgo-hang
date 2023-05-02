# Ginkgo v2.9.2 run hang reproducer

Compile the small c daemon example that is executed in the test.
```
$ gcc -o daemon mydaemon/main.c
```

Execute ginkgo test:
```
$ ginkgo run ./
```

Now try the same with the parallel mode:
```
$ ginkgo run -p ./
```

Here it looks ok from the output but ginkgo hangs and does not return.
Running this with a debugger it looks like it hangs here:
https://github.com/onsi/ginkgo/blob/8b925abe6639e4bfd62b9defc35c5da41da08d6a/ginkgo/internal/run.go#L225

The problem is that wait() tries to finish copying all IO.
ginkgo seems to dup stdout/err to higher fds. The fds are then leaked into the daemon process
that we launch in out test. Thus the fds are kept permanently open because the child process of
the daemon outlives the tests. This means the read call on that pipe hangs forever.

To fix this we should make sure the fds are not leaked into child processes, the easiest way IMO
seems to be adding the O_CLOEXEC option to the duplicated fds.
