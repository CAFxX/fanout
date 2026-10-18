package main

import (
    "fmt"
)


func main() {
    v1 := []int{451, 2147483647}
    v2 := []string{(("" + "foo bar") + "hello"), ("hello" + ""), (("" + "foo bar") + ""), "\n\t"}
    v3 := []string{"\n\t", "hello", ("a" + ""), (((("" + "") + (("" + "") + ((("\n\t" + "hello") + "hello") + "foo bar"))) + "\n\t") + ("a" + "a"))}
    var v4 uint8 = 127
    var v5 uint16 = 32767
    var v6 uint32 = 446
    fmt.Println(v6, v6)
    v6 += (v6 | (v6 + (v6 % ((1) | 1))))
    var v7 int16 = 1000
    v7 = 31
    var v8 int = -1
    v8 = 32767
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
