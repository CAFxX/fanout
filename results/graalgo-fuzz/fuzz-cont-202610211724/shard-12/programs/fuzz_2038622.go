package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 string }
func genFn0(p0 bool, p1 string) int16 {
    fmt.Println(p0, p0)
    if p0 {
        v7 := GenS1{F0: 1, F1: 25}
        _ = v7
    }
    if (!true) {
        if (1.5 != (((float32((float64((-2.25 / -2.25)) - 3.14159)) * 0.0) / -2.25) * 1e10)) {
            if (!p0) {
                var v8 float32 = 1e10
                p0 = (!true)
                _ = v8
            }
        } else {
            fmt.Println(p0, p1)
            p0 = (401 > 42)
        }
    } else {
        if (9223372036854775807 < 0) {
            p0 = (p0 || true)
            var v9 int8 = -128
            _ = v9
        } else {
            if (!(4294967295 != 65535)) {
                v10 := []int{-681, 32767, -1, -939}
                fmt.Println(p1, p0)
                _ = v10
            } else {
                var v11 uint = 0
                _ = v11
            }
            p0 = p0
        }
        fmt.Println(p1, p0)
        v12 := GenS1{F0: 2147483648, F1: 65535}
        _ = v12
    }
    p1 = p1
    _ = p0
    _ = p1
    return -1
}

func main() {
    var v1 float64 = (1.5 / 0.1)
    v2 := make(map[int]bool)
    var v3 uint64 = 0
    v4 := make(map[uint32]uint32)
    v5 := make(map[string]int)
    var v6 uint32 = 412
    if (false || true) {
        v3 = 1
    }
    v2[127] = false
    switch 1 {
    case -2147483648:
        v13 := GenS1{F0: 2147483647, F1: 919}
        _ = v13
    case -9223372036854775808:
        var v14 uint8 = 179
        _ = v14
    case 2147483647:
        var v15 uint64 = v3
        _ = v15
    default:
        v16 := []int{-1, -36}
        _ = v16
    }
    v6 += v6
    fmt.Println(v1)
    var v17 int = 128
    fmt.Println(v3, v3)
    if true {
        v18 := GenS1{F0: (v6 + v6), F1: 958}
        v3 -= v3
        fmt.Println(v1, v17, v3)
        _ = v18
    }
    for fi19 := 0; fi19 < 19; fi19++ {
        v6 -= (v6 * v6)
        _ = fi19
    }
    v20 := []int{(v17 << ((255) & 63))}
    if len(v20) > 0 {
        v20[0] = 240
    }
    var v21 bool = false
    v22 := GenS1{F0: (v6 ^ 4294967295), F1: 4294967295}
    if v21 {
        switch "\n\t" {
        case "\n\t":
            switch ("" + "a") {
            case "\n\t":
                switch v17 {
                case 32767:
                    var v23 uint = 255
                    _ = v23
                case 32768:
                    v17 = v17
                }
            case "foo bar":
                delete(v4, v6)
            case "":
                v24 := GenS2{F0: "\n\t"}
                _ = v24
            default:
                if v21 {
                    v17 = v17
                    var v25 int16 = -32768
                    v4[v6] = v6
                    _ = v25
                }
            }
        default:
            if true {
                v26 := GenS2{F0: ((((("" + ("" + ("a" + "\n\t"))) + "hello") + "a") + "a") + "")}
                _ = v26
            } else {
                if len(v20) > 0 {
                    v20[0] = -921
                }
            }
        }
        fmt.Println(v6, v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
