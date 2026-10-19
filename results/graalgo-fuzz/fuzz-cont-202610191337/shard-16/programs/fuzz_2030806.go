package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 string }
type GenS2 struct { F0 int32; F1 uint32; F2 string }
type GenS3 struct { F0 int; F1 uint16 }
type GenS4 struct { F0 uint16 }

func main() {
    var v1 uint64 = 2
    v2 := make(map[int]uint32)
    var v3 int = 1
    v4 := []string{(((("\n\t" + ("a" + (("hello" + "hello") + "hello"))) + "\n\t") + ((("\n\t" + ("\n\t" + "foo bar")) + "hello") + "hello")) + "a"), "hello"}
    v1 = 42
    v5 := GenS1{F0: -9223372036854775808, F1: 0, F2: ((("hello" + "foo bar") + (((("hello" + ((((("\n\t" + "hello") + "hello") + "a") + "hello") + "\n\t")) + "") + "") + "foo bar")) + "\n\t")}
    var v6 uint = 42
    delete(v2, v3)
    var v7 float64 = (((float64((1e10 - float32(-2.25))) * (0.1 + -2.25)) / 3.14159) / (0.1 - 1e10))
    v1 = v1
    if ((0.0 + 0.1) != (((3.14159 / ((0.1 * 0.1) / 3.14159)) * (1e10 * float32(-2.25))) + (1e10 - (-2.25 / -2.25)))) {
        var v8 int16 = 138
        v7 += v7
        for fi9 := 0; fi9 < 15; fi9++ {
            var v10 int32 = -1
            _ = fi9
            _ = v10
        }
        _ = v8
    } else {
        if false {
            v1 -= v1
        }
    }
    var v11 string = ("a" + "hello")
    fmt.Println(v6)
    switch v3 {
    case 88:
        v12 := GenS2{F0: -476, F1: 42, F2: (v11 + v11)}
        _ = v12
    default:
        if len(v4) > 0 {
            v4[0] = (v11 + "\n\t")
        }
    }
    if ((v3 * (v3 / ((v3) | 1))) != v3) {
        _ = len(v2)
    } else {
        v3 += v3
        _ = len(v4)
        v6 = 0
    }
    v7 -= v7
    switch (v3 << ((v6) & 63)) {
    case -9223372036854775808:
        var v13 float64 = -2.25
        _ = v13
    case 127:
        v1 = (v1 << (((v6 | (v6 % ((1) | 1)))) & 63))
    case 1:
        v6 -= v6
    default:
        if (4294967295 > 1) {
            v14 := GenS1{F0: 431, F1: 2147483647, F2: v11}
            for fi15 := 0; fi15 < 2; fi15++ {
                v16 := GenS3{F0: (v3 / (((v3 / ((128) | 1))) | 1)), F1: uint16(v6)}
                _ = fi15
                _ = v16
            }
            var v17 bool = true
            _ = v14
            _ = v17
        } else {
            v18 := GenS3{F0: (v3 - v3), F1: 0}
            _ = v18
        }
    }
    switch v11 {
    case "a":
        var v19 int64 = -1
        _ = v19
    case "hello":
        switch v11 {
        case "a":
            var v20 int32 = 1
            _ = v20
        }
    case "foo bar":
        v21 := GenS4{F0: 32767}
        _ = v21
    }
    v11 = v11
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
