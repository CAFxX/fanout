package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 int16 }
type GenS2 struct { F0 int64 }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]bool)
    v3 := []string{(("hello" + ((("a" + ("\n\t" + ("hello" + "hello"))) + ("foo bar" + "foo bar")) + "hello")) + ("hello" + (("" + "foo bar") + "hello"))), "", (("hello" + "hello") + (("" + "foo bar") + "hello"))}
    switch "hello" {
    case "\n\t":
        var v4 float32 = 0.1
        _ = v4
    default:
        switch "\n\t" {
        case "":
            switch "hello" {
            case "foo bar":
                var v5 int16 = -32768
                _ = v5
            case "":
                var v6 uint16 = 0
                _ = v6
            case "a":
                var v7 float32 = 1.5
                _ = v7
            }
        case "\n\t":
            var v8 uint = 255
            _ = v8
        case "foo bar":
            _ = len(v1)
        }
    }
    var v9 uint64 = 0
    var v10 uint16 = 980
    var v11 uint8 = 140
    switch 32768 {
    case 4294967295:
        var v12 int8 = int8(v10)
        _ = v12
    case 9223372036854775807:
        v13 := GenS1{F0: v9, F1: 4294967295, F2: 0}
        _ = v13
    case -9223372036854775808:
        v9 -= (v9 - v9)
    }
    for fi14 := 0; fi14 < 6; fi14++ {
        for fi15 := 0; fi15 < 16; fi15++ {
            v10 -= v10
            _ = fi15
        }
        if true {
            switch fi14 {
            case 15:
                v11 += v11
            case 127:
                v16 := GenS1{F0: v9, F1: 42, F2: 1}
                _ = v16
            case 1:
                var v17 float32 = 0.1
                _ = v17
            default:
                v10 -= (v10 & 255)
            }
            var v18 string = (("hello" + ("a" + "a")) + "\n\t")
            fmt.Println(v18, v11, v10)
            _ = v18
        } else {
            var v19 int64 = int64(v10)
            _ = v19
        }
        _ = fi14
    }
    v20 := GenS2{F0: -1}
    _ = len(v3)
    fmt.Println(v11, v9, v9)
    v10 -= 65535
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v20
    _ = v3
    _ = v9
}
