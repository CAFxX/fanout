package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 bool; F1 uint; F2 uint16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint; F1 bool }
type GenS7 struct { F0 uint }
func genFn0(p0 int64, p1 int8, p2 uint) int {
    var v7 int64 = p0
    var v8 int8 = -1
    v9 := GenS1{F0: v8, F1: p0}
    var v10 bool = false
    p2 += p2
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 1
}

func main() {
    var v1 bool = false
    var v2 int8 = 0
    v3 := make(map[int]uint32)
    v4 := make(map[int]uint32)
    var v5 int8 = v2
    var v6 uint32 = 42
    v11 := GenS2{F0: v1, F1: 0, F2: 65535}
    v12 := GenS1{F0: v2, F1: 1}
    v13 := []int{-914}
    for fi14 := 0; fi14 < 2; fi14++ {
        v6 = v6
        var v15 int16 = -1
        v15 -= (v15 ^ v15)
        _ = fi14
        _ = v15
    }
    var v16 uint32 = v6
    v6 += (v6 * (v6 | v16))
    fmt.Println(v16)
    if v1 {
        switch ("hello" + "a") {
        case "foo bar":
            v17 := GenS1{F0: 2, F1: -9223372036854775808}
            _ = v17
        case "":
            v3[-130] = v6
        case "a":
            if (false || (!(85 > 4294967295))) {
                v16 += (^(v16))
                v18 := GenS3{F0: (v2 >> ((18446744073709551615) & 63))}
                v2 += (v2 ^ (v5 / (((^(v2))) | 1)))
                _ = v18
            } else {
                fmt.Println(v1, v6, v2)
                v1 = v1
            }
        }
        switch 819 {
        case -1:
            switch 32768 {
            case 4294967295:
                v2 = v2
            default:
                fmt.Println(v2)
            }
        case -2147483648:
            fmt.Println(v6, v1)
        case 0:
            v19 := GenS4{F0: 996}
            _ = v19
        }
    } else {
        if len(v13) >= 2 {
            v13 = v13[1:len(v13)]
        }
    }
    var v20 bool = (false || v1)
    v21 := GenS3{F0: v5}
    switch "foo bar" {
    case "":
        v22 := GenS5{F0: uint8(v6)}
        _ = v22
    case "hello":
        var v23 float64 = -2.25
        _ = v23
    case "a":
        fmt.Println(v20)
    }
    v4[42] = 875
    fmt.Println(v20)
    for fi24 := 0; fi24 < 13; fi24++ {
        for fi25 := 0; fi25 < 6; fi25++ {
            if true {
                v5 = 42
                v2 -= 100
            }
            if (v16 <= 4294967295) {
                var v26 int = (fi25 % ((fi24) | 1))
                v27 := GenS6{F0: 644, F1: (!false)}
                _ = v26
                _ = v27
            } else {
                var v28 float64 = 1.5
                v29 := v3[fi24]
                _ = v28
                _ = v29
            }
            _ = fi25
        }
        v6 -= (v6 + (v16 | v16))
        _ = fi24
    }
    var v30 uint32 = (v16 - (v6 * 65535))
    v31 := GenS7{F0: 379}
    v32 := GenS1{F0: v5, F1: 625}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
