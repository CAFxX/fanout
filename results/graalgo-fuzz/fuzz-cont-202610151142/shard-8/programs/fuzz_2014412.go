package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 int8 }
func genFn0(p0 int, p1 bool) uint8 {
    var v4 int16 = 1000
    var v5 uint64 = 992
    var v6 float32 = (1e10 + -2.25)
    v6 += v6
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 42
}
func genFn1(p0 uint16) uint16 {
    var v7 uint8 = 42
    switch -1 {
    case 32767:
        var v8 float64 = ((-2.25 / 1e10) + 1e10)
        _ = v8
    case -9223372036854775808:
        fmt.Println(v7, v7)
    case -1:
        p0 -= (p0 + p0)
    default:
        v9 := GenS1{F0: (v7 << ((0) & 63)), F1: 0, F2: -128}
        _ = v9
    }
    var v10 uint32 = 0
    _ = p0
    _ = v10
    _ = v7
    return uint16(v10)
}

func main() {
    var v1 int64 = 110
    v2 := make(map[string]int)
    var v3 int32 = 42
    for fi11 := 0; fi11 < 11; fi11++ {
        var v12 float64 = (0.1 / (1e10 - ((-2.25 / -2.25) - 0.0)))
        _ = fi11
        _ = v12
    }
    v3 = v3
    fmt.Println(v3, v3)
    v13 := []int{9223372036854775807}
    var v14 uint = 255
    var v15 float32 = (1e10 - 3.14159)
    v2[("" + ("foo bar" + (("foo bar" + "") + "\n\t")))] = 2147483647
    v15 = 3.14159
    v14 -= 255
    v2[""] = 0
    if (1 <= genFn1(1)) {
        for fi16 := 0; fi16 < 11; fi16++ {
            if (109 < 255) {
                v15 += 1e10
            } else {
                v14 = v14
                var v17 uint64 = 666
                _ = v17
            }
            _ = fi16
        }
    } else {
        for fi18 := 0; fi18 < 8; fi18++ {
            var v19 uint8 = 2
            _ = fi18
            _ = v19
        }
        if (1 <= 103) {
            var v20 float64 = -2.25
            _ = v20
        } else {
            var v21 float32 = v15
            fmt.Println(v1, v14, v21)
            _ = v21
        }
    }
    v13 = append(v13, 550)
    switch ("\n\t" + (("hello" + "\n\t") + "hello")) {
    case "":
        v14 -= 980
    case "hello":
        v22 := GenS1{F0: 1, F1: 424, F2: 127}
        _ = v22
    default:
        fmt.Println(v15, v3, v14)
    }
    var v23 int8 = 0
    v14 += (v14 * 685)
    v23 = 0
    for fi24 := 0; fi24 < 13; fi24++ {
        if len(v13) >= 2 {
            v13 = v13[1:len(v13)]
        }
        v25 := GenS1{F0: 255, F1: uint64(v14), F2: v23}
        _ = fi24
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v23
    _ = v3
}
