package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
func genFn0(p0 bool, p1 int64, p2 int32) string {
    for fi5 := 0; fi5 < 6; fi5++ {
        if (!p0) {
            var v6 float64 = 0.0
            fmt.Println(p0, p0)
            for fi7 := 0; fi7 < 16; fi7++ {
                v8 := []int{32768, (fi5 + fi7)}
                v9 := GenS1{F0: 2, F1: 790}
                _ = fi7
                _ = v8
                _ = v9
            }
            _ = v6
        }
        var v10 int64 = -1
        _ = fi5
        _ = v10
    }
    fmt.Println(p1)
    v11 := []int{1}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return "a"
}
func genFn1(p0 int16, p1 uint64) int16 {
    v12 := []int{4294967295, 32768, 32767, 42}
    v13 := GenS1{F0: (^(p1)), F1: 1000}
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    return p0
}

func main() {
    var v1 uint8 = 127
    var v2 uint16 = 2
    var v3 int64 = 42
    v4 := make(map[int]string)
    v14 := GenS1{F0: 18446744073709551615, F1: -247}
    v3 = v3
    for fi15 := 0; fi15 < 4; fi15++ {
        v16 := GenS1{F0: 42, F1: 1000}
        _ = fi15
        _ = v16
    }
    _ = len(v4)
    switch "foo bar" {
    case "a":
        switch 2147483647 {
        case 32767:
            v2 = (v2 ^ v2)
        case 88:
            if false {
                switch "foo bar" {
                case "a":
                    var v17 uint64 = 0
                    _ = v17
                case "":
                    if (1.5 > -2.25) {
                        var v18 float64 = ((0.1 + 0.0) / 0.1)
                        v19 := []int{}
                        _ = v18
                        _ = v19
                    }
                }
                v20 := GenS1{F0: 630, F1: 572}
                _ = v20
            }
        }
    default:
        var v21 uint32 = 65535
        _ = v21
    }
    for fi22 := 0; fi22 < 9; fi22++ {
        v23 := GenS1{F0: 358, F1: 1}
        var v24 bool = true
        _ = fi22
        _ = v23
        _ = v24
    }
    fmt.Println(v1)
    v25 := GenS1{F0: uint64(v2), F1: 644}
    var v26 float64 = 0.1
    var v27 uint64 = 949
    var v28 uint8 = (v1 << ((557) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
