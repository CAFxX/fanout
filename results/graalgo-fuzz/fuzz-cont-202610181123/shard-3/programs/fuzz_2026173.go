package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 int16, p1 int8, p2 uint16) uint16 {
    var v7 uint16 = (^(p2))
    if (4294967295 < 524) {
        v8 := []int{4294967295, -1}
        _ = len(v8)
        if false {
            if (!false) {
                fmt.Println(p1, v7, v7)
                v9 := GenS1{F0: 9223372036854775807}
                for fi10 := 0; fi10 < 5; fi10++ {
                    var v11 uint16 = (p2 + 1)
                    var v12 float64 = ((0.1 + float64((float32(0.1) * 1e10))) - (((1e10 / 3.14159) / (1.5 * 0.1)) + float64(3.14159)))
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                _ = v9
            }
            fmt.Println(p0, v7)
        } else {
            p2 += (p2 / ((98) | 1))
            v13 := GenS1{F0: -1}
            fmt.Println(p1, p2, p2)
            _ = v13
        }
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return v7
}
func genFn1(p0 uint64) int8 {
    switch "a" {
    case "foo bar":
        for fi14 := 0; fi14 < 13; fi14++ {
            var v15 uint16 = genFn0(-32768, 127, 278)
            _ = fi14
            _ = v15
        }
    case "\n\t":
        if true {
            v16 := []int{}
            _ = v16
        }
    case "a":
        var v17 uint8 = 1
        _ = v17
    }
    for fi18 := 0; fi18 < 20; fi18++ {
        var v19 float64 = 0.1
        _ = fi18
        _ = v19
    }
    _ = p0
    return -1
}

func main() {
    v1 := make(map[int]string)
    var v2 uint8 = 42
    v3 := make(map[uint32]string)
    var v4 float64 = 1e10
    v5 := []string{"a", "hello"}
    v6 := []uint32{840, 1, 4294967295}
    var v20 int8 = -98
    v4 = 1e10
    fmt.Println(v4)
    switch -9223372036854775808 {
    case 32767:
        switch 852 {
        case -1:
            var v21 int64 = 42
            _ = v21
        case 996:
            v4 -= (v4 - v4)
        case 1:
            if true {
                if ((float32(1e10) + 3.14159) <= float32(v4)) {
                    for fi22 := 0; fi22 < 1; fi22++ {
                        var v23 uint64 = 665
                        _ = fi22
                        _ = v23
                    }
                }
                fmt.Println(v2, v20)
                if (669 >= 411) {
                    var v24 uint64 = 334
                    _ = v24
                } else {
                    v20 = genFn1(0)
                }
            }
        default:
            switch 9223372036854775807 {
            case 9223372036854775807:
                v25 := GenS1{F0: 580}
                _ = v25
            default:
                v4 = v4
            }
        }
    }
    if (("a" < "a") && (true && (0 != 131))) {
        fmt.Println(v2)
    }
    v20 = v20
    v20 = v20
    switch "" {
    case "\n\t":
        _ = len(v5)
    case "hello":
        v4 -= -2.25
    }
    if len(v6) > 0 {
        v6[0] = 42
    }
    if true {
        v2 = (^(v2))
        if len(v6) > 0 {
            v6[0] = 1
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
