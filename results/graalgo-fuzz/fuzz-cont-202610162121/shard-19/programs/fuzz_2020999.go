package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 uint8; F1 int64; F2 uint16 }
func genFn0() int16 {
    v7 := GenS1{F0: "hello", F1: true}
    switch ("foo bar" + "a") {
    case "foo bar":
        var v8 uint32 = 2147483647
        _ = v8
    default:
        switch "foo bar" {
        case "hello":
            if false {
                var v9 int64 = -9223372036854775808
                v10 := GenS2{F0: 42, F1: v9, F2: 65535}
                _ = v10
                _ = v9
            }
        }
    }
    _ = v7
    return 334
}

func main() {
    var v1 uint32 = 42
    var v2 uint32 = v1
    var v3 string = ""
    var v4 int16 = 42
    var v5 int64 = -1
    v6 := make(map[uint32]string)
    for fi11 := 0; fi11 < 12; fi11++ {
        fmt.Println(v5, fi11, v1)
        _ = fi11
    }
    v12 := GenS1{F0: v3, F1: false}
    if (490 != 0) {
        var v13 uint16 = 255
        var v14 uint16 = v13
        _ = v13
        _ = v14
    }
    v6[v2] = ("\n\t" + (v3 + ((("hello" + v3) + ("" + ((v3 + ((("hello" + v3) + v3) + ("\n\t" + (("hello" + "foo bar") + "foo bar")))) + "foo bar"))) + (v3 + v3))))
    switch 34 {
    case 128:
        var v15 uint32 = (^(v2))
        _ = v15
    case -585:
        if false {
            fmt.Println(v5)
        }
    default:
        switch "a" {
        case "hello":
            var v16 int8 = 1
            _ = v16
        case "foo bar":
            v1 -= (v1 & 2147483648)
        case "a":
            v17 := v6[1]
            _ = v17
        default:
            v2 += (^(v1))
        }
    }
    v1 += 655
    v5 += 586
    var v18 float32 = float32((0.0 - -2.25))
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
