package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32; F1 string; F2 string }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint16; F1 uint; F2 bool }
type GenS7 struct { F0 int64; F1 uint16 }
type GenS8 struct { F0 uint32; F1 uint }
func genFn0() string {
    v5 := GenS1{F0: 42}
    v6 := GenS2{F0: 2}
    for fi7 := 0; fi7 < 6; fi7++ {
        switch "a" {
        case "\n\t":
            var v8 int32 = -1
            _ = v8
        default:
            v9 := GenS3{F0: 42, F1: "\n\t", F2: "\n\t"}
            _ = v9
        }
        _ = fi7
    }
    switch "foo bar" {
    case "foo bar":
        var v10 uint8 = 42
        _ = v10
    case "\n\t":
        var v11 uint64 = 0
        _ = v11
    case "a":
        var v12 uint16 = 1
        _ = v12
    default:
        for fi13 := 0; fi13 < 18; fi13++ {
            var v14 int16 = 558
            _ = fi13
            _ = v14
        }
    }
    v15 := []int{-9223372036854775808, 128, 42, 128}
    _ = v15
    _ = v5
    _ = v6
    return "a"
}

func main() {
    var v1 bool = (3.14159 <= -2.25)
    v2 := make(map[int]bool)
    var v3 int = 4294967295
    v4 := make(map[string]bool)
    fmt.Println(v1, v3)
    if ((0.1 * (1e10 * 0.1)) >= ((1.5 - -2.25) / -2.25)) {
        var v16 int16 = -1
        _ = v16
    } else {
        v1 = true
        v17 := GenS2{F0: 42}
        _ = v17
    }
    v1 = true
    v18 := []int{-9223372036854775808, v3, (v3 % ((-328) | 1)), -2147483648}
    fmt.Println(v3, v1)
    v19 := GenS1{F0: 0}
    v20 := GenS4{F0: true}
    v21 := GenS5{F0: 1000}
    v22 := GenS1{F0: 1}
    switch v3 {
    case -1:
        for fi23 := 0; fi23 < 14; fi23++ {
            var v24 bool = true
            var v25 uint16 = 2
            _ = fi23
            _ = v24
            _ = v25
        }
    case 1:
        for fi26 := 0; fi26 < 20; fi26++ {
            _ = len(v2)
            if v1 {
                fmt.Println(v1, v3)
            } else {
                if true {
                    v27 := GenS6{F0: 65535, F1: 0, F2: (false || v1)}
                    fmt.Println(fi26)
                    v18 = append(v18, v3)
                    _ = v27
                } else {
                    v28 := GenS4{F0: v1}
                    var v29 int16 = 1000
                    _ = v28
                    _ = v29
                }
            }
            v30 := v4[""]
            _ = fi26
            _ = v30
        }
    default:
        var v31 int32 = 27
        _ = v31
    }
    if v1 {
        v32 := GenS7{F0: 1, F1: 322}
        v1 = v1
        var v33 string = "a"
        _ = v32
        _ = v33
    }
    v3 -= (v3 | (v3 / ((405) | 1)))
    var v34 bool = ("hello" <= "")
    v35 := v4["\n\t"]
    v36 := GenS8{F0: 148, F1: 0}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
}
