package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 int8 = 69
    v2 := make(map[int]bool)
    v3 := make(map[uint32]uint32)
    v4 := []int{-9223372036854775808, 32768, 2147483647}
    if false {
        var v5 float32 = 0.0
        _ = v5
    }
    v1 += (-(v1))
    if true {
        var v6 int = -313
        var v7 uint = 65535
        _ = v6
        _ = v7
    }
    v1 += v1
    if (0 <= 65535) {
        v1 += (v1 + v1)
        switch ("a" + ("a" + ("foo bar" + ""))) {
        case "":
            fmt.Println(v1)
        case "hello":
            v1 = (v1 & v1)
        }
    } else {
        fmt.Println(v1)
        fmt.Println(v1)
    }
    switch 127 {
    case -1:
        v1 += (v1 - v1)
    case 2147483647:
        if true {
            if (32767 == 843) {
                switch "hello" {
                case "hello":
                    switch "\n\t" {
                    case "foo bar":
                        v1 -= (v1 << ((2) & 63))
                    case "":
                        var v8 uint64 = 0
                        _ = v8
                    case "a":
                        switch (("a" + "") + "foo bar") {
                        case "hello":
                            if (!false) {
                                v1 = -128
                            }
                        case "a":
                            var v9 uint32 = 2147483647
                            _ = v9
                        case "":
                            var v10 int = 890
                            _ = v10
                        }
                    }
                }
            }
        }
    default:
        var v11 uint = 190
        _ = v11
    }
    v12 := GenS1{F0: 18446744073709551615}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
