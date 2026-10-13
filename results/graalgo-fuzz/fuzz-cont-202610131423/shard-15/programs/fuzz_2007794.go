package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 string }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 int8; F1 uint8 }
type GenS4 struct { F0 int64; F1 int64; F2 int }

func main() {
    var v1 uint = 455
    var v2 uint = 42
    v3 := make(map[string]string)
    var v4 uint32 = 2147483648
    if true {
        v2 -= (v1 + 65535)
        var v5 int16 = -32768
        _ = v5
    }
    v6 := GenS1{F0: 831, F1: 0, F2: ("a" + (("hello" + "\n\t") + "foo bar"))}
    var v7 int64 = -1
    for fi8 := 0; fi8 < 11; fi8++ {
        v7 += (v7 - (v7 * 127))
        if false {
            v4 = v4
            fmt.Println(fi8)
        }
        for fi9 := 0; fi9 < 13; fi9++ {
            if (42 == 4294967295) {
                var v10 int64 = v7
                if ("" <= "hello") {
                    switch ("foo bar" + "foo bar") {
                    case "a":
                        v11 := GenS2{F0: "hello", F1: 0}
                        _ = v11
                    case "\n\t":
                        fmt.Println(v1, v10)
                    }
                }
                if (true || false) {
                    var v12 float32 = (-2.25 * 0.0)
                    _ = v12
                }
                _ = v10
            } else {
                v4 += (v4 - v4)
                _ = len(v3)
                v13 := GenS1{F0: (fi8 | fi9), F1: -128, F2: "hello"}
                v4 = v4
                _ = v13
            }
            v14 := GenS3{F0: int8(fi8), F1: 42}
            _ = fi9
            _ = v14
        }
        _ = fi8
    }
    v15 := GenS4{F0: 0, F1: (v7 ^ (v7 >> (((v2 & v2)) & 63))), F2: 4294967295}
    var v16 int32 = int32(v2)
    v17 := []int{-1, 127}
    v16 += v16
    if false {
        var v18 uint8 = 27
        _ = v18
    }
    if true {
        if (((0.0 + (0.0 + (3.14159 + (0.0 + 1e10)))) / -2.25) >= (1e10 + 1e10)) {
            var v19 uint16 = 42
            var v20 uint32 = (v4 - v4)
            switch "hello" {
            case "hello":
                v21 := GenS2{F0: "a", F1: 0}
                _ = v21
            case "foo bar":
                if len(v17) >= 2 {
                    v17 = v17[1:len(v17)]
                }
            case "":
                v7 += (v7 >> (((v1 / ((65535) | 1))) & 63))
            default:
                var v22 int = 42
                _ = v22
            }
            _ = v19
            _ = v20
        } else {
            var v23 int32 = v16
            _ = v23
        }
    }
    var v24 int64 = (v7 % (((v7 / ((0) | 1))) | 1))
    if (!((!true) && (!false))) {
        v7 = -1
        v16 = 590
    }
    for fi25 := 0; fi25 < 13; fi25++ {
        var v26 int8 = 1
        var v27 uint32 = (v4 * (v4 ^ 2))
        _ = fi25
        _ = v26
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
