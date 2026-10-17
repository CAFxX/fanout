package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 int16 }
type GenS2 struct { F0 uint }
func genFn0(p0 uint32) int8 {
    var v6 int16 = -1
    for fi7 := 0; fi7 < 4; fi7++ {
        var v8 uint16 = 42
        v8 = uint16(p0)
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 1; fi9++ {
        var v10 uint64 = 1
        var v11 uint32 = 2147483648
        _ = fi9
        _ = v10
        _ = v11
    }
    if ("foo bar" <= "a") {
        p0 = p0
    } else {
        for fi12 := 0; fi12 < 7; fi12++ {
            v13 := []int{}
            _ = fi12
            _ = v13
        }
    }
    _ = p0
    _ = v6
    return 42
}
func genFn1() uint32 {
    switch 4294967295 {
    case 127:
        var v14 float32 = (-2.25 / float32((0.1 / (float64(((1.5 * 0.0) / 1.0)) - 0.1))))
        _ = v14
    case 128:
        if false {
            if (18446744073709551615 >= 4294967295) {
                v15 := GenS1{F0: -9223372036854775808, F1: 9223372036854775807, F2: -32768}
                v16 := []int{1}
                var v17 uint = 255
                _ = v15
                _ = v16
                _ = v17
            }
            var v18 float32 = 1e10
            _ = v18
        } else {
            var v19 int8 = 42
            var v20 float32 = 0.1
            _ = v19
            _ = v20
        }
    case 1:
        if true {
            var v21 float32 = 0.0
            v22 := GenS1{F0: 9223372036854775807, F1: 42, F2: 0}
            _ = v21
            _ = v22
        }
    }
    var v23 uint64 = 0
    fmt.Println(v23)
    v24 := GenS1{F0: 127, F1: -136, F2: 0}
    _ = v23
    _ = v24
    return 773
}

func main() {
    var v1 uint8 = 127
    v2 := []uint32{524, 2147483647}
    var v3 uint = 989
    v4 := []uint32{2147483648, 4294967295}
    var v5 bool = false
    var v25 int8 = -1
    v25 += int8(v3)
    v5 = v5
    _ = len(v2)
    for fi26 := 0; fi26 < 8; fi26++ {
        v4 = append(v4, genFn1())
        v27 := GenS2{F0: (v3 * 908)}
        _ = fi26
        _ = v27
    }
    v28 := GenS1{F0: -9223372036854775808, F1: 42, F2: 593}
    var v29 int64 = -268
    var v30 uint8 = 255
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
