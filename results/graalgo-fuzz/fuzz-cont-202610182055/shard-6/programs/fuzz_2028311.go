package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }

func main() {
    var v1 uint64 = 0
    v2 := []uint32{65535, 1}
    var v3 int64 = -9223372036854775808
    v4 := GenS1{F0: 1, F1: 2}
    for fi5 := 0; fi5 < 12; fi5++ {
        switch "hello" {
        case "a":
            var v6 int16 = 585
            _ = v6
        case "hello":
            fmt.Println(v3, v3)
        }
        _ = fi5
    }
    switch ("foo bar" + "\n\t") {
    case "a":
        v1 = (v1 + v1)
    case "":
        var v7 bool = true
        _ = v7
    case "hello":
        switch "a" {
        case "\n\t":
            v1 += (^(v1))
        case "foo bar":
            var v8 uint16 = 638
            _ = v8
        default:
            if true {
                v2 = append(v2, 398)
            } else {
                v3 = (v3 >> ((1) & 63))
                v3 = 9223372036854775807
            }
        }
    }
    v1 -= (v1 ^ (^(v1)))
    var v9 int16 = 801
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi10 := 0; fi10 < 10; fi10++ {
        var v11 float32 = (1.5 + 1.5)
        _ = fi10
        _ = v11
    }
    fmt.Println(v3, v3)
    fmt.Println(v3, v9, v1)
    if true {
        for fi12 := 0; fi12 < 10; fi12++ {
            var v13 float32 = 0.1
            if (2 > 3) {
                v13 -= float32(0.0)
            }
            if true {
                v3 = v3
                fmt.Println(v3, fi12)
            }
            _ = fi12
            _ = v13
        }
    }
    v3 = -1
    switch 32768 {
    case -9223372036854775808:
        if true {
            fmt.Println(v9)
            v1 = 164
        } else {
            v9 = 421
            fmt.Println(v1, v1, v1)
            v1 = (v1 >> ((42) & 63))
        }
    }
    for fi14 := 0; fi14 < 2; fi14++ {
        v3 = (^(v3))
        var v15 float32 = 1.5
        _ = fi14
        _ = v15
    }
    v9 = 0
    fmt.Println(v9, v3)
    _ = len(v2)
    var v16 int8 = 42
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
