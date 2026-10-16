package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }

func main() {
    var v1 uint64 = 2
    var v2 int = 127
    v3 := make(map[uint32]bool)
    v4 := v3[342]
    if (!v4) {
        v5 := v3[943]
        _ = v5
    }
    delete(v3, 0)
    var v6 int8 = 1
    fmt.Println(v4, v2, v1)
    v1 += v1
    var v7 int8 = v6
    switch ("\n\t" + "a") {
    case "\n\t":
        if (uint16(v1) > 42) {
            _ = len(v3)
            fmt.Println(v2, v4, v4)
        }
    case "a":
        for fi8 := 0; fi8 < 14; fi8++ {
            if (!(42 > 18446744073709551615)) {
                fmt.Println(v4)
            } else {
                v7 += v6
                delete(v3, 2)
            }
            _ = fi8
        }
    case "foo bar":
        for fi9 := 0; fi9 < 13; fi9++ {
            v4 = false
            for fi10 := 0; fi10 < 17; fi10++ {
                v1 = (v1 + 2)
                _ = fi10
            }
            _ = fi9
        }
    default:
        v11 := GenS1{F0: 32767, F1: v7}
        _ = v11
    }
    if (!((0 > int16(v2)) || v4)) {
        v12 := GenS1{F0: -32768, F1: 100}
        _ = v12
    }
    v6 += (v7 | v6)
    var v13 uint = 65535
    var v14 uint8 = 128
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
