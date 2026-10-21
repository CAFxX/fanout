package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 uint16 }
type GenS2 struct { F0 int16; F1 int8 }
func genFn0(p0 uint16) int {
    for fi4 := 0; fi4 < 16; fi4++ {
        p0 = 2
        var v5 int16 = 979
        _ = fi4
        _ = v5
    }
    for fi6 := 0; fi6 < 14; fi6++ {
        if true {
            p0 = p0
            var v7 uint64 = 42
            p0 = 32767
            _ = v7
        }
        _ = fi6
    }
    _ = p0
    return 32767
}
func genFn1() uint64 {
    for fi8 := 0; fi8 < 20; fi8++ {
        if true {
            var v9 float64 = 0.1
            _ = v9
        } else {
            var v10 int64 = 1
            var v11 int = fi8
            _ = v10
            _ = v11
        }
        _ = fi8
    }
    switch 2147483647 {
    case 9223372036854775807:
        if false {
            switch ("\n\t" + "") {
            case "foo bar":
                v12 := GenS1{F0: 4294967295, F1: (true && false), F2: 127}
                _ = v12
            default:
                var v13 uint64 = 42
                _ = v13
            }
            var v14 float64 = 1.5
            _ = v14
        }
    }
    for fi15 := 0; fi15 < 12; fi15++ {
        var v16 int8 = -1
        fmt.Println(fi15, fi15)
        _ = fi15
        _ = v16
    }
    var v17 int8 = 100
    v18 := []int{-1, 32767, -1, 127}
    _ = v17
    _ = v18
    return 2
}

func main() {
    var v1 uint64 = 1
    var v2 float64 = 0.0
    var v3 int16 = 42
    v19 := []int{32768, 4294967295, 128}
    var v20 uint64 = v1
    var v21 uint64 = (v1 << ((uint(v1)) & 63))
    fmt.Println(v1, v21)
    if (true || false) {
        fmt.Println(v20, v1, v1)
        v22 := GenS2{F0: v3, F1: -128}
        _ = v22
    } else {
        if false {
            v3 += v3
            var v23 uint16 = 667
            var v24 uint = 2
            _ = v23
            _ = v24
        }
        v20 = 810
        v25 := GenS1{F0: 2147483647, F1: false, F2: 65535}
        _ = v25
    }
    v1 += v21
    v3 = (v3 * 42)
    v20 += v21
    v26 := GenS2{F0: (v3 ^ v3), F1: 100}
    var v27 uint8 = 0
    v20 -= (^(v21))
    var v28 uint32 = 65535
    v19 = append(v19, -870)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v26
    _ = v27
    _ = v28
    _ = v3
}
