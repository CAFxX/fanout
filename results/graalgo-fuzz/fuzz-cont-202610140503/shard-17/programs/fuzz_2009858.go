package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    var v1 uint8 = 2
    var v2 int = 38
    var v3 uint32 = 890
    var v4 bool = false
    fmt.Println(v1)
    for fi5 := 0; fi5 < 14; fi5++ {
        if (false || v4) {
            v3 = 460
            var v6 string = "hello"
            if false {
                var v7 uint = 13
                _ = v7
            }
            _ = v6
        }
        v3 = 2
        _ = fi5
    }
    v8 := []int{32768, (v2 >> ((uint(v1)) & 63))}
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    var v9 uint32 = 665
    switch "foo bar" {
    case "\n\t":
        if len(v8) > 0 {
            v8[0] = 128
        }
    case "foo bar":
        if v4 {
            v9 += (v9 << ((241) & 63))
            v10 := GenS1{F0: 65535}
            _ = v10
        }
    case "hello":
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
    }
    v2 -= (v2 ^ 705)
    v2 -= -477
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
