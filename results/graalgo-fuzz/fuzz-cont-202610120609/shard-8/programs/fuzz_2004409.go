package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool; F1 int; F2 int }
type GenS3 struct { F0 int; F1 int8; F2 uint }
func genFn0(p0 int32) uint32 {
    for fi4 := 0; fi4 < 8; fi4++ {
        v5 := GenS1{F0: 2147483648}
        p0 += (p0 ^ p0)
        _ = fi4
        _ = v5
    }
    switch "\n\t" {
    case "hello":
        p0 += p0
    case "\n\t":
        var v6 uint = 4294967295
        _ = v6
    default:
        if (true || (!true)) {
            if (false && ((456 == 1) || false)) {
                fmt.Println(p0)
            }
            if (true || true) {
                v7 := GenS1{F0: 263}
                switch 4294967295 {
                case 315:
                    if (("" + "\n\t") > "\n\t") {
                        var v8 float32 = 1e10
                        _ = v8
                    } else {
                        fmt.Println(p0)
                    }
                }
                v9 := []int{-1, 9223372036854775807, 32767, 519}
                _ = v7
                _ = v9
            } else {
                var v10 uint16 = 65535
                _ = v10
            }
        }
    }
    _ = p0
    return 0
}

func main() {
    v1 := make(map[string]int)
    var v2 uint64 = 2
    v3 := make(map[string]uint32)
    v11 := GenS1{F0: genFn0(-313)}
    var v12 int64 = -1
    var v13 float32 = 1.5
    if true {
        var v14 int32 = 0
        v2 = v2
        fmt.Println(v14)
        _ = v14
    }
    v15 := v3[""]
    v13 -= float32(float64(float32(3.14159)))
    v16 := GenS2{F0: (true && true), F1: 42, F2: 0}
    v15 -= v15
    var v17 float64 = 0.1
    for fi18 := 0; fi18 < 8; fi18++ {
        for fi19 := 0; fi19 < 13; fi19++ {
            if false {
                v12 = (-(v12))
                v20 := GenS1{F0: v15}
                v12 -= (v12 | v12)
                _ = v20
            } else {
                switch ("\n\t" + "foo bar") {
                case "\n\t":
                    var v21 uint8 = uint8(v15)
                    _ = v21
                case "foo bar":
                    var v22 int32 = 471
                    _ = v22
                default:
                    var v23 int64 = -9223372036854775808
                    _ = v23
                }
                var v24 int8 = 100
                v25 := GenS3{F0: 613, F1: (v24 & 0), F2: 18446744073709551615}
                _ = v24
                _ = v25
            }
            _ = fi19
        }
        delete(v1, "")
        _ = fi18
    }
    var v26 uint8 = 42
    v27 := []int{}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v26
    _ = v27
    _ = v3
}
