package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int32; F1 int; F2 string }

func main() {
    v1 := make(map[int]uint32)
    var v2 int32 = 42
    v3 := make(map[int]bool)
    var v4 uint16 = 893
    var v5 float64 = (0.0 * 1e10)
    v6 := GenS1{F0: 42, F1: 2147483647, F2: (^(v4))}
    v2 += v2
    var v7 int16 = 1
    switch "\n\t" {
    case "":
        v4 = (v4 / ((v4) | 1))
    default:
        var v8 int32 = -1
        _ = v8
    }
    for fi9 := 0; fi9 < 9; fi9++ {
        v7 -= -1
        fmt.Println(v2, fi9)
        _ = fi9
    }
    if true {
        switch "hello" {
        case "":
            v5 += (v5 / v5)
        case "a":
            if (((v5 / v5) * (0.1 + v5)) == v5) {
                v5 = (-2.25 * float64(0.0))
                var v10 uint64 = 0
                var v11 uint16 = v4
                _ = v10
                _ = v11
            } else {
                fmt.Println(v7)
                var v12 string = ""
                _ = v12
            }
        case "foo bar":
            v7 += (v7 + (v7 << ((0) & 63)))
        default:
            var v13 int32 = v2
            _ = v13
        }
    }
    v14 := GenS2{F0: ("a" + ("\n\t" + (("\n\t" + "hello") + "\n\t")))}
    var v15 int = 32767
    v2 = (v2 / ((2147483647) | 1))
    v2 = (v2 & 1)
    fmt.Println(v2)
    for fi16 := 0; fi16 < 8; fi16++ {
        v15 -= (fi16 ^ fi16)
        _ = fi16
    }
    v17 := GenS3{F0: 2}
    v18 := GenS4{F0: v2, F1: v15, F2: "hello"}
    var v19 int = 2147483647
    for fi20 := 0; fi20 < 18; fi20++ {
        for fi21 := 0; fi21 < 7; fi21++ {
            var v22 int = (v19 << ((255) & 63))
            v23 := []int{}
            v23 = append(v23, 9223372036854775807)
            var v24 uint = 46
            fmt.Println(v15)
            _ = fi21
            _ = v22
            _ = v23
            _ = v24
        }
        for fi25 := 0; fi25 < 3; fi25++ {
            v4 = v4
            v26 := []int{}
            _ = fi25
            _ = v26
        }
        _ = fi20
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
