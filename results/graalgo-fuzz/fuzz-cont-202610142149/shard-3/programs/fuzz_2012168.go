package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    var v1 bool = (32767 >= -32768)
    v2 := make(map[string]bool)
    var v3 uint16 = 255
    v4 := make(map[string]string)
    var v5 float32 = (0.1 / 1.5)
    v6 := []string{("" + ("hello" + ("hello" + "a"))), "\n\t"}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v1 = (v1 || v1)
    v5 = v5
    v1 = false
    v7 := GenS1{F0: 42}
    v1 = v1
    var v8 uint = 4294967295
    v9 := GenS1{F0: -32768}
    if len(v6) > 0 {
        v6[0] = "hello"
    }
    v10 := GenS1{F0: -32768}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
