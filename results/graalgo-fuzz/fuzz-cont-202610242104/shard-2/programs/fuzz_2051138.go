package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string; F1 uint8; F2 int8 }
type GenS4 struct { F0 uint32; F1 int32 }
func genFn0(p0 uint64, p1 int16, p2 uint8) int16 {
    p1 = p1
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    return -121
}

func main() {
    v1 := make(map[string]string)
    var v2 uint16 = 214
    var v3 uint32 = 2147483648
    var v4 uint64 = 1
    if (42 <= -2147483648) {
        for fi5 := 0; fi5 < 19; fi5++ {
            v6 := GenS1{F0: 1, F1: ((((false || ((!((!true) && ((!false) && ((v2 - v2) >= 435)))) && true)) && true) || (!(2 >= 42))) && true)}
            _ = fi5
            _ = v6
        }
    } else {
        var v7 uint8 = 1
        _ = v7
    }
    v2 = (v2 & (v2 | v2))
    var v8 float32 = float32((float64(1.5) + -2.25))
    v3 += v3
    switch "foo bar" {
    case "a":
        fmt.Println(v4)
    case "":
        v9 := GenS2{F0: 255}
        _ = v9
    default:
        for fi10 := 0; fi10 < 2; fi10++ {
            v11 := v1[""]
            v4 += (v4 % (((v4 % (((v4 + (v4 >> ((837) & 63)))) | 1))) | 1))
            _ = fi10
            _ = v11
        }
    }
    fmt.Println(v2, v8)
    switch "hello" {
    case "hello":
        v8 += 1e10
    default:
        if false {
            if false {
                var v12 float32 = 1.5
                v4 -= (v4 & v4)
                _ = v12
            } else {
                var v13 uint64 = (v4 + (v4 + 1))
                var v14 int8 = 0
                v3 += 4294967295
                _ = v13
                _ = v14
            }
            v15 := GenS3{F0: "a", F1: 128, F2: -80}
            _ = v15
        } else {
            var v16 int64 = 42
            _ = v16
        }
    }
    v17 := GenS4{F0: v3, F1: -356}
    switch "" {
    case "\n\t":
        switch "\n\t" {
        case "\n\t":
            var v18 int32 = -799
            _ = v18
        case "a":
            var v19 uint8 = 127
            _ = v19
        }
    }
    fmt.Println(v8, v4, v2)
    for fi20 := 0; fi20 < 5; fi20++ {
        var v21 int8 = 100
        switch fi20 {
        case 42:
            v4 += (^(v4))
        case 2147483647:
            fmt.Println(fi20)
        }
        v22 := []int{9223372036854775807, 128}
        _ = fi20
        _ = v21
        _ = v22
    }
    for fi23 := 0; fi23 < 13; fi23++ {
        var v24 float32 = v8
        _ = fi23
        _ = v24
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
