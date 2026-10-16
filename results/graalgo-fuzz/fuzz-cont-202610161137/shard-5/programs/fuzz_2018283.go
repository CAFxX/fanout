package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
func genFn0(p0 uint, p1 int, p2 int) bool {
    var v6 uint16 = 65535
    v7 := []int{p2}
    if (!false) {
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        var v8 int8 = -36
        if len(v7) > 0 {
            v7[0] = -9223372036854775808
        }
        _ = v8
    }
    var v9 uint16 = 32767
    p0 -= (^(p0))
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v9
    return true
}

func main() {
    var v1 float64 = 1e10
    var v2 int = 474
    v3 := make(map[uint32]int)
    var v4 int = 128
    var v5 int8 = 42
    switch v4 {
    case 4294967295:
        var v10 uint = 4294967295
        _ = v10
    case -870:
        if (!false) {
            v5 = v5
            if (18446744073709551615 <= 2) {
                v5 = (v5 ^ 28)
            }
            v1 += (float64(0.0) + v1)
        }
    case -9223372036854775808:
        _ = len(v3)
    }
    fmt.Println(v4, v2, v2)
    for fi11 := 0; fi11 < 16; fi11++ {
        for fi12 := 0; fi12 < 19; fi12++ {
            v4 = (fi11 & fi11)
            _ = fi12
        }
        v2 += 2147483647
        fmt.Println(fi11)
        _ = fi11
    }
    fmt.Println(v2, v5)
    v13 := []int{(v2 % ((1) | 1))}
    v4 = 42
    if true {
        if true {
            var v14 int32 = -2147483648
            for fi15 := 0; fi15 < 12; fi15++ {
                var v16 int16 = 404
                fmt.Println(v2, v2)
                _ = fi15
                _ = v16
            }
            v1 -= v1
            _ = v14
        } else {
            switch "hello" {
            case "foo bar":
                var v17 string = "a"
                _ = v17
            default:
                v4 += int(v5)
            }
        }
    } else {
        if ((0.1 - float32(0.1)) >= ((float32((v1 * (1.5 - v1))) / -2.25) - 3.14159)) {
            v18 := GenS1{F0: 842, F1: true}
            switch (v4 & (v4 ^ (v4 & (v4 >> ((255) & 63))))) {
            case 359:
                var v19 int32 = 2147483647
                _ = v19
            default:
                for fi20 := 0; fi20 < 12; fi20++ {
                    var v21 uint8 = 42
                    for fi22 := 0; fi22 < 6; fi22++ {
                        v1 = float64(((0.1 * (0.1 * (((float32(((v1 - v1) + v1)) + 0.0) * (0.0 - 1.5)) + 1e10))) + 0.0))
                        var v23 int16 = -859
                        var v24 int32 = 1
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                    v13 = append(v13, v2)
                    _ = fi20
                    _ = v21
                }
            }
            _ = v18
        }
        fmt.Println(v1, v1, v5)
        fmt.Println(v1, v2)
    }
    var v25 int64 = 1
    v4 = (v4 + (v2 & 1))
    v4 = v2
    if (-2.25 > 1e10) {
        v26 := GenS1{F0: 32767, F1: false}
        _ = v26
    } else {
        var v27 uint64 = 18446744073709551615
        v28 := GenS1{F0: 42, F1: (!(!((v5 == (v5 - 0)) && (4294967295 <= 65535))))}
        _ = v27
        _ = v28
    }
    var v29 uint = 1
    var v30 int = (v4 ^ v2)
    delete(v3, 2147483648)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
