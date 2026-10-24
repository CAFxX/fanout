package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 uint16 }
func genFn0(p0 string, p1 uint16) uint8 {
    fmt.Println(p0)
    v5 := []int{128, -1}
    var v6 float32 = float32((1e10 / -2.25))
    switch -1 {
    case 127:
        v7 := GenS1{F0: 1, F1: p0, F2: p1}
        _ = v7
    case -595:
        for fi8 := 0; fi8 < 9; fi8++ {
            v9 := GenS1{F0: 953, F1: (p0 + "hello"), F2: 42}
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            var v10 bool = (((v6 - (v6 / v6)) * 1e10) > 0.0)
            _ = fi8
            _ = v10
            _ = v9
        }
    case 9223372036854775807:
        var v11 int = 32767
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 2
}
func genFn1(p0 int64, p1 bool) int8 {
    var v12 int32 = 1
    if p1 {
        var v13 int64 = -507
        _ = v13
    } else {
        v12 += (v12 / ((v12) | 1))
        var v14 float32 = 0.1
        if (p1 || p1) {
            p1 = (1e10 != v14)
        }
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v12
    return 127
}

func main() {
    var v1 int64 = -1
    v2 := make(map[uint32]bool)
    v3 := []string{"a", "a", "\n\t"}
    var v4 string = "foo bar"
    var v15 uint = 273
    v16 := GenS1{F0: 4294967295, F1: "hello", F2: 13}
    v1 = (v1 % (((v1 / (((^(v1))) | 1))) | 1))
    var v17 uint8 = 48
    if true {
        switch "a" {
        case "a":
            v18 := GenS1{F0: 1, F1: "a", F2: 1}
            _ = v18
        case "":
            v19 := GenS1{F0: 2, F1: (v4 + (v4 + v4)), F2: 32767}
            _ = v19
        case "\n\t":
            if (0.1 <= 0.0) {
                var v20 int = 32767
                for fi21 := 0; fi21 < 12; fi21++ {
                    v22 := GenS1{F0: 4294967295, F1: v4, F2: 65535}
                    _ = fi21
                    _ = v22
                }
                for fi23 := 0; fi23 < 10; fi23++ {
                    delete(v2, 2147483647)
                    _ = len(v3)
                    _ = fi23
                }
                _ = v20
            }
        default:
            var v24 float32 = 0.0
            _ = v24
        }
    }
    var v25 uint8 = genFn0("\n\t", 2)
    for fi26 := 0; fi26 < 3; fi26++ {
        v15 = v15
        _ = len(v3)
        _ = fi26
    }
    if false {
        v15 = 255
        if (false || false) {
            for fi27 := 0; fi27 < 17; fi27++ {
                var v28 int = fi27
                fmt.Println(v15, v28)
                _ = fi27
                _ = v28
            }
            var v29 string = ("" + (v4 + v4))
            _ = v29
        }
    }
    v30 := GenS1{F0: 538, F1: v4, F2: 255}
    for fi31 := 0; fi31 < 4; fi31++ {
        if len(v3) > 0 {
            v3[0] = v4
        }
        v17 = v25
        var v32 string = (v4 + v4)
        _ = fi31
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v25
    _ = v3
    _ = v30
    _ = v4
}
