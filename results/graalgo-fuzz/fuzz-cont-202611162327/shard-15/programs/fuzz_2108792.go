package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 int32; F2 uint }
type GenS3 struct { F0 string; F1 int32 }

func main() {
    v1 := []string{"foo bar", "\n\t"}
    var v2 uint = 4294967295
    v3 := make(map[int]uint32)
    v4 := GenS1{F0: 32767}
    for fi5 := 0; fi5 < 19; fi5++ {
        if false {
            switch "\n\t" {
            case "a":
                var v6 uint32 = 2147483648
                _ = v6
            case "\n\t":
                v2 = v2
            case "foo bar":
                fmt.Println(v2, v2)
            default:
                v2 = (v2 - (^(v2)))
            }
        } else {
            switch "a" {
            case "foo bar":
                if len(v1) > 0 {
                    v1[0] = ((("" + "foo bar") + "hello") + "")
                }
            default:
                v2 = 176
            }
            v7 := GenS2{F0: 2, F1: 0, F2: (v2 & (v2 & v2))}
            var v8 bool = ((v2 - v2) > 4294967295)
            _ = v7
            _ = v8
        }
        _ = fi5
    }
    var v9 float64 = 0.0
    v2 = v2
    for fi10 := 0; fi10 < 11; fi10++ {
        if true {
            v11 := GenS3{F0: "\n\t", F1: 1}
            _ = v11
        }
        _ = fi10
    }
    var v12 uint16 = 65535
    switch "\n\t" {
    case "foo bar":
        v9 += v9
    case "\n\t":
        v2 += 1
    case "hello":
        if (true || true) {
            _ = len(v1)
        }
    }
    var v13 int8 = -39
    if (true && (1 == 42)) {
        v14 := GenS2{F0: 108, F1: -119, F2: (v2 | v2)}
        switch "\n\t" {
        case "":
            v1 = append(v1, "")
        case "foo bar":
            fmt.Println(v2, v9)
        case "hello":
            v1 = append(v1, "a")
        default:
            var v15 int32 = 2147483647
            _ = v15
        }
        _ = v14
    } else {
        fmt.Println(v13)
        var v16 uint8 = 42
        if true {
            var v17 int64 = -1
            if (-926 == 127) {
                var v18 int64 = (v17 * v17)
                v13 -= v13
                _ = v18
            } else {
                v9 = (v9 - v9)
                if (true || false) {
                    var v19 int = 32768
                    v2 += 0
                    _ = v19
                }
            }
            v20 := GenS2{F0: (v16 + v16), F1: 2147483647, F2: v2}
            _ = v17
            _ = v20
        } else {
            v3[-2147483648] = 530
        }
        _ = v16
    }
    v21 := v3[-2147483648]
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v9
}
