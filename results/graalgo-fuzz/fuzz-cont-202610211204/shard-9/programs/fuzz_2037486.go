package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 uint8 }

func main() {
    v1 := make(map[string]string)
    var v2 bool = (-9223372036854775808 != 4294967295)
    v3 := make(map[uint32]string)
    var v4 int64 = 0
    if (v2 || true) {
        fmt.Println(v4, v2)
        var v5 int = -1
        fmt.Println(v5)
        _ = v5
    }
    switch ("a" + "\n\t") {
    case "foo bar":
        switch "hello" {
        case "hello":
            switch "" {
            case "foo bar":
                switch -730 {
                case 216:
                    var v6 int16 = 271
                    _ = v6
                default:
                    if (128 <= 42) {
                        var v7 int64 = 196
                        _ = v7
                    }
                }
            case "hello":
                fmt.Println(v2, v4)
            }
        case "a":
            var v8 uint64 = 42
            _ = v8
        case "":
            v9 := GenS1{F0: 1, F1: v2}
            _ = v9
        default:
            v4 += 552
        }
    case "\n\t":
        if (-59 != 1000) {
            v10 := GenS1{F0: 255, F1: (v2 && true)}
            for fi11 := 0; fi11 < 18; fi11++ {
                v2 = true
                _ = fi11
            }
            v12 := GenS1{F0: 4294967295, F1: (1 < 4294967295)}
            _ = v10
            _ = v12
        }
    case "":
        v2 = (!v2)
    default:
        var v13 uint64 = 2
        _ = v13
    }
    if ((v4 * v4) <= 428) {
        for fi14 := 0; fi14 < 14; fi14++ {
            v2 = (v2 && (!v2))
            _ = fi14
        }
        v2 = v2
    } else {
        v2 = v2
        v4 += v4
        v4 = 1
    }
    fmt.Println(v2, v2)
    v15 := GenS1{F0: 65535, F1: (true && v2)}
    var v16 int = 128
    switch "hello" {
    case "":
        v4 = 0
    case "hello":
        if true {
            v17 := GenS2{F0: 128}
            _ = v17
        } else {
            v16 += (v16 / (((v16 + (v16 ^ (-(v16))))) | 1))
        }
    case "\n\t":
        v2 = false
    }
    var v18 uint16 = 42
    for fi19 := 0; fi19 < 6; fi19++ {
        v18 = v18
        _ = fi19
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
