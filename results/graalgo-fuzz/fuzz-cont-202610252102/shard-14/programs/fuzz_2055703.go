package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 int8 }
type GenS2 struct { F0 int; F1 int; F2 uint16 }
func genFn0(p0 int16, p1 int16, p2 bool) string {
    p1 = p1
    v5 := GenS1{F0: p1, F1: 127, F2: 1}
    if false {
        var v6 int8 = 0
        v7 := []int{-1}
        _ = v6
        _ = v7
    } else {
        v8 := []int{}
        v8 = append(v8, -980)
        _ = v8
    }
    for fi9 := 0; fi9 < 17; fi9++ {
        p0 = p1
        if p2 {
            if (!p2) {
                var v10 int8 = 1
                _ = v10
            }
            if (p2 && p2) {
                var v11 int16 = (p0 / ((p0) | 1))
                var v12 bool = ((v11 << ((772) & 63)) > (p1 >> ((382) & 63)))
                var v13 float32 = (1e10 - 1.5)
                _ = v11
                _ = v12
                _ = v13
            }
            var v14 uint = 174
            _ = v14
        } else {
            if (!p2) {
                var v15 float32 = ((3.14159 + 1e10) * (1e10 + ((0.1 / 3.14159) - -2.25)))
                v16 := GenS1{F0: 1, F1: 127, F2: 24}
                var v17 float32 = v15
                _ = v15
                _ = v16
                _ = v17
            }
            switch "a" {
            case "\n\t":
                var v18 uint8 = 0
                _ = v18
            }
        }
        p1 += 32767
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return "\n\t"
}
func genFn1(p0 uint) uint {
    var v19 uint8 = 1
    v20 := GenS2{F0: 32768, F1: -373, F2: 42}
    v21 := []int{9223372036854775807, 42, -186}
    fmt.Println(v19, p0)
    _ = p0
    _ = v19
    _ = v20
    _ = v21
    return (^(p0))
}

func main() {
    var v1 bool = ((!true) && (766 < 32767))
    var v2 float32 = (-2.25 - float32(0.0))
    var v3 int8 = -43
    v4 := make(map[int]int)
    var v22 uint64 = 42
    if (((1 >= 128) && true) || v1) {
        v22 -= (v22 << ((genFn1(18446744073709551615)) & 63))
        var v23 uint = 2
        var v24 uint = v23
        _ = v23
        _ = v24
    } else {
        v25 := GenS2{F0: 620, F1: 2147483647, F2: 1}
        var v26 int32 = 0
        if ((!v1) || ((v3 & (-(v3))) == (-(v3)))) {
            var v27 float64 = (1e10 / 1.0)
            v2 = v2
            fmt.Println(v3, v3, v2)
            _ = v27
        }
        _ = v25
        _ = v26
    }
    var v28 bool = (!(!v1))
    fmt.Println(v1, v2, v28)
    v29 := v4[720]
    for fi30 := 0; fi30 < 11; fi30++ {
        for fi31 := 0; fi31 < 6; fi31++ {
            v32 := []int{(v29 << ((228) & 63)), (-(fi31))}
            _ = fi31
            _ = v32
        }
        var v33 string = "\n\t"
        _ = fi30
        _ = v33
    }
    v28 = v28
    var v34 float64 = ((3.14159 - (1e10 / 1.0)) / 1.0)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v22
    _ = v28
    _ = v29
    _ = v3
    _ = v34
    _ = v4
}
