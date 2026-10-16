package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 uint8 }
type GenS2 struct { F0 bool; F1 bool; F2 string }
type GenS3 struct { F0 uint16; F1 uint32; F2 int }
type GenS4 struct { F0 string; F1 uint16; F2 uint64 }

func main() {
    v1 := []string{"foo bar", "", "foo bar", ""}
    v2 := make(map[int]uint32)
    var v3 int16 = -388
    var v4 int16 = 1000
    v4 = 1
    switch "\n\t" {
    case "\n\t":
        v4 = v3
    default:
        var v5 int8 = -128
        _ = v5
    }
    for fi6 := 0; fi6 < 1; fi6++ {
        v1 = append(v1, (("foo bar" + ("hello" + "foo bar")) + "a"))
        v4 = 0
        _ = fi6
    }
    var v7 uint64 = 42
    if (v7 != (^(v7))) {
        fmt.Println(v4)
        var v8 uint16 = 540
        v3 = (v3 >> ((42) & 63))
        _ = v8
    }
    v9 := GenS1{F0: 255, F1: (false && true), F2: 255}
    switch -1 {
    case -909:
        v10 := GenS2{F0: (!(42 >= -618)), F1: false, F2: ("hello" + ("a" + "\n\t"))}
        _ = v10
    default:
        switch "foo bar" {
        case "foo bar":
            v7 -= v7
        case "a":
            var v11 float32 = 3.14159
            _ = v11
        default:
            for fi12 := 0; fi12 < 14; fi12++ {
                v4 += v4
                _ = fi12
            }
        }
    }
    var v13 int32 = 2147483647
    v14 := GenS3{F0: 521, F1: 935, F2: 1}
    if false {
        for fi15 := 0; fi15 < 19; fi15++ {
            v2[997] = 841
            v16 := GenS4{F0: "", F1: 0, F2: 18446744073709551615}
            fmt.Println(v7)
            _ = fi15
            _ = v16
        }
        fmt.Println(v13)
        v17 := GenS3{F0: 65535, F1: 988, F2: 32768}
        _ = v17
    }
    v4 += (v3 + (v4 & v4))
    if false {
        var v18 uint = 331
        _ = v18
    } else {
        v19 := GenS1{F0: 65535, F1: false, F2: 0}
        _ = v19
    }
    for fi20 := 0; fi20 < 3; fi20++ {
        if false {
            v21 := v2[fi20]
            _ = v21
        }
        _ = fi20
    }
    var v22 float32 = 0.0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v7
    _ = v9
}
