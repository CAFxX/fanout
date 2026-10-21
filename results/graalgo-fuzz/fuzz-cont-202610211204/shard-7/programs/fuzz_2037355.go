package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
func genFn0(p0 int64, p1 uint32, p2 int32) int16 {
    p2 -= p2
    p2 += p2
    var v4 uint64 = 93
    for fi5 := 0; fi5 < 7; fi5++ {
        fmt.Println(fi5, p2)
        _ = fi5
    }
    for fi6 := 0; fi6 < 14; fi6++ {
        var v7 int32 = p2
        for fi8 := 0; fi8 < 16; fi8++ {
            var v9 bool = (((("foo bar" + "a") + "") <= "a") || (true || false))
            if (!v9) {
                if true {
                    v10 := []int{(fi6 * fi6)}
                    _ = v10
                } else {
                    var v11 uint16 = 42
                    if v9 {
                        var v12 uint64 = 1
                        _ = v12
                    }
                    fmt.Println(v11, p0, fi6)
                    v4 += 0
                    _ = v11
                }
            }
            _ = fi8
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return -244
}
func genFn1() uint16 {
    for fi13 := 0; fi13 < 1; fi13++ {
        var v14 uint16 = 65535
        v14 = v14
        fmt.Println(fi13)
        _ = fi13
        _ = v14
    }
    var v15 uint = 178
    switch -9223372036854775808 {
    case 578:
        var v16 uint = v15
        _ = v16
    case -1:
        var v17 uint16 = 1
        _ = v17
    case 4294967295:
        v15 += 65535
    default:
        if (true && (v15 == (v15 % (((v15 / (((^(v15))) | 1))) | 1)))) {
            v18 := []int{2147483647, 32768, -2147483648}
            switch "\n\t" {
            case "a":
                var v19 int16 = 42
                _ = v19
            default:
                fmt.Println(v15)
            }
            fmt.Println(v15)
            _ = v18
        } else {
            v20 := []int{127, 0, 2147483647, 2147483647}
            _ = v20
        }
    }
    _ = v15
    return 851
}

func main() {
    v1 := make(map[string]string)
    var v2 uint8 = 42
    var v3 uint64 = 347
    v21 := v1[("foo bar" + "a")]
    var v22 float32 = -2.25
    var v23 uint = 2
    v24 := GenS1{F0: (v2 * 252), F1: 128}
    v25 := []int{1, 9223372036854775807, 735, -1}
    var v26 int64 = -9223372036854775808
    var v27 int16 = -56
    if false {
        v22 = 0.0
    } else {
        v27 -= (v27 & (v27 / ((-489) | 1)))
    }
    if false {
        v3 = 2
    } else {
        delete(v1, "\n\t")
        v22 = (0.1 - (v22 + 1.5))
        var v28 uint64 = (v3 & 60)
        _ = v28
    }
    v2 += (v2 - v2)
    if len(v25) >= 2 {
        v25 = v25[1:len(v25)]
    }
    v26 = -919
    var v29 uint32 = 363
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v29
    _ = v3
}
