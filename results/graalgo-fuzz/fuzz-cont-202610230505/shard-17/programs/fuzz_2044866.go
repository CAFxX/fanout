package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int32 }

func main() {
    v1 := []uint32{2147483647, 130, 2147483648}
    var v2 uint64 = 857
    var v3 uint64 = 4294967295
    switch -75 {
    case 32767:
        var v4 int = 9223372036854775807
        _ = v4
    case -2147483648:
        v3 -= (v3 % (((^(v2))) | 1))
    case 2147483647:
        v5 := GenS1{F0: "foo bar"}
        _ = v5
    default:
        var v6 uint16 = 32767
        _ = v6
    }
    v3 -= (v2 / (((v2 / ((0) | 1))) | 1))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if len(v1) > 0 {
        v1[0] = 4294967295
    }
    fmt.Println(v3)
    var v7 float32 = 0.1
    v2 = (v3 - 424)
    var v8 int64 = -1
    for fi9 := 0; fi9 < 12; fi9++ {
        v1 = append(v1, 411)
        for fi10 := 0; fi10 < 8; fi10++ {
            if ("a" != "hello") {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                fmt.Println(v8, fi9)
                switch fi9 {
                case 32768:
                    if true {
                        v1 = append(v1, 935)
                        v11 := GenS2{F0: uint16(v3), F1: 65535}
                        _ = v11
                    }
                default:
                    var v12 int = (fi9 + fi10)
                    _ = v12
                }
                v7 = 0.1
                _ = len(v1)
                v2 = (v2 * v2)
            } else {
                var v13 int32 = int32(v8)
                _ = v13
            }
            fmt.Println(v7, v2)
            _ = fi10
        }
        v14 := GenS3{F0: "a"}
        _ = fi9
        _ = v14
    }
    fmt.Println(v3, v2)
    for fi15 := 0; fi15 < 15; fi15++ {
        v8 -= v8
        v1 = append(v1, 112)
        v2 += 42
        _ = fi15
    }
    if ((v8 / ((v8) | 1)) < -154) {
        if (!(-2147483648 < 42)) {
            var v16 uint16 = 42
            _ = v16
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v17 uint64 = v2
        _ = v17
    } else {
        var v18 bool = true
        var v19 uint8 = 194
        _ = v18
        _ = v19
    }
    v2 = v3
    fmt.Println(v3)
    switch ("" + "hello") {
    case "a":
        for fi20 := 0; fi20 < 17; fi20++ {
            v21 := GenS4{F0: -1}
            _ = fi20
            _ = v21
        }
    case "\n\t":
        fmt.Println(v7)
    case "hello":
        v1 = append(v1, 2147483648)
    }
    _ = len(v1)
    v8 = (v8 >> ((0) & 63))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
