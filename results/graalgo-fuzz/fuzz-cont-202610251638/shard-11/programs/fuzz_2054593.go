package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 uint16; F1 uint8; F2 uint }
func genFn0(p0 int32, p1 int16) int32 {
    var v5 int8 = -106
    v6 := []int{290, 1, -163}
    p1 -= p1
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return (p0 * 0)
}
func genFn1(p0 bool) uint16 {
    p0 = p0
    v7 := GenS1{F0: 42, F1: -1}
    _ = p0
    _ = v7
    return 42
}

func main() {
    var v1 float32 = 0.1
    v2 := make(map[int]bool)
    v3 := make(map[string]uint32)
    v4 := make(map[int]bool)
    v1 += v1
    fmt.Println(v1)
    v1 -= v1
    for fi8 := 0; fi8 < 14; fi8++ {
        var v9 uint16 = 32767
        _ = fi8
        _ = v9
    }
    switch 101 {
    case -1:
        v10 := GenS1{F0: 782, F1: -1}
        _ = v10
    case -2147483648:
        v1 = 0.1
    default:
        fmt.Println(v1)
    }
    v1 -= ((v1 + float32(3.14159)) + -2.25)
    switch ("a" + ("a" + ("hello" + "foo bar"))) {
    case "\n\t":
        v11 := []int{570, 42}
        _ = v11
    case "a":
        v1 = v1
    default:
        var v12 int16 = 428
        _ = v12
    }
    var v13 int32 = 1
    if false {
        fmt.Println(v1, v1)
        switch "foo bar" {
        case "foo bar":
            v1 = v1
        case "hello":
            var v14 uint64 = 988
            _ = v14
        }
    } else {
        switch 2147483647 {
        case 32767:
            if false {
                v15 := v2[-98]
                _ = v15
            }
        case -689:
            switch 32768 {
            case 9223372036854775807:
                var v16 int8 = -1
                _ = v16
            case 128:
                fmt.Println(v13, v1)
            case 32768:
                fmt.Println(v13)
            default:
                v17 := GenS2{F0: 0, F1: 128, F2: 0}
                _ = v17
            }
        default:
            for fi18 := 0; fi18 < 16; fi18++ {
                var v19 uint64 = 77
                v20 := []int{(fi18 / (((fi18 << ((487) & 63))) | 1)), fi18, fi18, 362}
                var v21 int8 = 100
                _ = fi18
                _ = v19
                _ = v20
                _ = v21
            }
        }
        v1 = v1
        v1 = float32(1e10)
    }
    fmt.Println(v13)
    var v22 uint8 = 126
    for fi23 := 0; fi23 < 1; fi23++ {
        var v24 uint8 = (v22 / ((88) | 1))
        _ = fi23
        _ = v24
    }
    if true {
        if true {
            fmt.Println(v22, v22)
            if true {
                v25 := GenS2{F0: 65535, F1: 127, F2: 65535}
                _ = v25
            } else {
                var v26 float64 = 1e10
                _ = v26
            }
        } else {
            fmt.Println(v13)
            var v27 int64 = -610
            _ = v27
        }
        delete(v4, 32767)
        v1 = 1e10
    }
    var v28 uint = 288
    v29 := v3["\n\t"]
    var v30 uint16 = genFn1(false)
    v29 -= 2
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v2
    _ = v22
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
