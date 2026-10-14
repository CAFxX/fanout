package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }

func main() {
    var v1 int32 = 42
    var v2 bool = ("" < ("hello" + ("foo bar" + "\n\t")))
    v3 := make(map[int]int)
    if (88 >= 2) {
        v4 := v3[654]
        _ = v4
    } else {
        fmt.Println(v2)
    }
    switch "foo bar" {
    case "\n\t":
        fmt.Println(v2, v1)
    default:
        var v5 uint = 475
        _ = v5
    }
    _ = len(v3)
    var v6 uint16 = 32767
    var v7 int64 = -1
    switch -2147483648 {
    case -620:
        _ = len(v3)
    case 9223372036854775807:
        var v8 uint64 = 2
        _ = v8
    case 259:
        v9 := []int{-1, 691}
        _ = v9
    default:
        v2 = v2
    }
    var v10 uint8 = 1
    switch 32767 {
    case 420:
        fmt.Println(v2)
    default:
        _ = len(v3)
    }
    v2 = (3.14159 >= 1e10)
    for fi11 := 0; fi11 < 16; fi11++ {
        switch "a" {
        case "a":
            if ((v2 || (-32768 != -32768)) || false) {
                v6 += (v6 << ((255) & 63))
                fmt.Println(v2, v1)
                switch (("a" + "foo bar") + "hello") {
                case "a":
                    var v12 int = -737
                    _ = v12
                case "\n\t":
                    v13 := []int{fi11, (fi11 % (((fi11 ^ fi11)) | 1))}
                    _ = v13
                case "foo bar":
                    v6 -= 1
                }
            } else {
                v1 += (v1 & 2147483647)
                v14 := GenS1{F0: fi11, F1: 0}
                _ = v14
            }
        case "\n\t":
            v15 := GenS1{F0: -1, F1: 0}
            _ = v15
        case "foo bar":
            v6 = (v6 << ((358) & 63))
        }
        var v16 uint = 42
        _ = fi11
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
