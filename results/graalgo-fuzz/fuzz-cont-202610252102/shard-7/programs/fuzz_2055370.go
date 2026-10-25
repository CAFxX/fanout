package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }

func main() {
    var v1 uint16 = 65535
    v2 := []string{("hello" + "hello"), "foo bar", "hello", ("a" + "foo bar")}
    v3 := []string{"foo bar", "", ("hello" + ("\n\t" + ("" + ("foo bar" + "foo bar")))), "hello"}
    var v4 int16 = 1000
    var v5 uint8 = 153
    v5 += (^(v5))
    if false {
        fmt.Println(v5, v1, v5)
    }
    var v6 int32 = -1
    var v7 uint16 = 65535
    v6 -= (v6 + (-(v6)))
    v8 := GenS1{F0: (42 >= 9223372036854775807), F1: ("" + "hello")}
    var v9 int64 = 1
    v1 = v1
    var v10 string = "foo bar"
    if (137 <= uint64(v1)) {
        v5 += v5
    }
    var v11 uint32 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
