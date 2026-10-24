package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 uint32; F1 uint8; F2 uint8 }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]bool)
    var v3 int64 = 175
    v3 -= v3
    v2[65535] = false
    v3 = (v3 >> ((2) & 63))
    v4 := v2[0]
    if (-2147483648 <= int(v3)) {
        if (v4 && false) {
            v5 := GenS1{F0: ("\n\t" > "hello")}
            _ = v5
        } else {
            v6 := GenS1{F0: (0 >= 32767)}
            _ = v6
        }
        var v7 int64 = 0
        _ = v7
    }
    v2[42] = true
    v8 := GenS1{F0: ((1e10 + 0.0) < (1e10 + float32(-2.25)))}
    v9 := GenS1{F0: false}
    v3 = v3
    if v4 {
        v10 := GenS1{F0: (v4 && v4)}
        _ = v10
    } else {
        if (127 != 127) {
            var v11 float64 = float64(1.5)
            v12 := []int{-374, 127, 42, -489}
            _ = v11
            _ = v12
        }
    }
    if (false || (false || v4)) {
        v3 = 642
        v13 := GenS1{F0: false}
        v3 -= v3
        _ = v13
    } else {
        switch -314 {
        case 32767:
            var v14 int32 = 12
            _ = v14
        case 883:
            if (0 > 100) {
                v15 := GenS2{F0: v3, F1: false}
                for fi16 := 0; fi16 < 11; fi16++ {
                    for fi17 := 0; fi17 < 12; fi17++ {
                        fmt.Println(fi16, v3, fi16)
                        _ = fi17
                    }
                    var v18 uint64 = 4294967295
                    _ = fi16
                    _ = v18
                }
                _ = v15
            } else {
                var v19 uint16 = 939
                _ = v19
            }
        case 9223372036854775807:
            fmt.Println(v4, v3)
        default:
            delete(v2, 65535)
        }
        if v4 {
            v3 -= v3
            fmt.Println(v3)
        }
    }
    v20 := GenS2{F0: (v3 - 42), F1: v4}
    v21 := []int{32767}
    v2[4294967295] = v4
    var v22 bool = v4
    if ((0 < 0) || v22) {
        for fi23 := 0; fi23 < 2; fi23++ {
            var v24 uint16 = 57
            _ = fi23
            _ = v24
        }
    } else {
        switch "\n\t" {
        case "a":
            fmt.Println(v3, v22)
        case "hello":
            v25 := GenS2{F0: 1, F1: (9223372036854775807 >= v3)}
            _ = v25
        case "\n\t":
            fmt.Println(v22)
        default:
            v26 := GenS3{F0: 65535, F1: 2, F2: 42}
            _ = v26
        }
    }
    delete(v1, 42)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
