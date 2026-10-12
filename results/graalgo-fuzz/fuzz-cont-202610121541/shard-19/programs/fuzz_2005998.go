package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 bool = (1 < 18446744073709551615)
    var v2 uint8 = 221
    var v3 uint = 0
    v4 := make(map[uint32]int)
    fmt.Println(v2, v3, v3)
    for fi5 := 0; fi5 < 10; fi5++ {
        var v6 string = "foo bar"
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: -1}
    v1 = v1
    v2 = 127
    fmt.Println(v3)
    var v8 uint16 = 65535
    v1 = (v1 && false)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
