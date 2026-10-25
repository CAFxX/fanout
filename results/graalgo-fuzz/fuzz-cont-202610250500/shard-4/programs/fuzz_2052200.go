package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 string; F1 int32 }
type GenS4 struct { F0 bool; F1 uint8; F2 string }

func main() {
    var v1 bool = (!false)
    v2 := []string{}
    var v3 uint = 226
    var v4 uint64 = 995
    fmt.Println(v3, v4, v3)
    var v5 int32 = -326
    if ("\n\t" >= "") {
        if (0 <= -128) {
            switch 127 {
            case 1:
                v6 := GenS1{F0: 2}
                _ = v6
            case 4294967295:
                v7 := GenS1{F0: 2}
                _ = v7
            case 9223372036854775807:
                var v8 bool = false
                _ = v8
            default:
                v9 := GenS2{F0: v4, F1: 255}
                _ = v9
            }
            v1 = (true || true)
            v10 := GenS1{F0: 32767}
            _ = v10
        }
    }
    v5 -= v5
    switch 32768 {
    case 2147483647:
        var v11 uint8 = 91
        _ = v11
    case -2147483648:
        v3 -= (v3 ^ (v3 << (((^(v3))) & 63)))
    default:
        v12 := GenS3{F0: "hello", F1: 1}
        _ = v12
    }
    v4 += v4
    if (718 >= 1000) {
        var v13 int32 = 1
        fmt.Println(v5, v3, v5)
        _ = v13
    }
    var v14 int16 = -32768
    v4 = v4
    if (!v1) {
        if v1 {
            v15 := GenS4{F0: false, F1: 118, F2: ("hello" + "a")}
            var v16 int8 = -102
            var v17 uint16 = 498
            _ = v15
            _ = v16
            _ = v17
        }
        for fi18 := 0; fi18 < 1; fi18++ {
            v1 = false
            var v19 string = "foo bar"
            v14 += (v14 >> ((v3) & 63))
            _ = fi18
            _ = v19
        }
    }
    if v1 {
        v20 := GenS1{F0: 42}
        v3 = 510
        _ = v20
    }
    var v21 float64 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
