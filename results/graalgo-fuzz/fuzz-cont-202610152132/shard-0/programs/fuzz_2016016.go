package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 int32 }
func genFn0(p0 int16, p1 string, p2 uint32) int16 {
    switch 9223372036854775807 {
    case 0:
        switch p1 {
        case "\n\t":
            for fi6 := 0; fi6 < 4; fi6++ {
                for fi7 := 0; fi7 < 9; fi7++ {
                    v8 := GenS1{F0: (p1 + "hello"), F1: p2, F2: int32(fi6)}
                    _ = fi7
                    _ = v8
                }
                v9 := []int{(fi6 & fi6), 4294967295, (fi6 | (fi6 % ((fi6) | 1)))}
                _ = fi6
                _ = v9
            }
        }
    default:
        v10 := []int{-9223372036854775808, 42}
        _ = v10
    }
    var v11 int64 = 0
    var v12 uint64 = 0
    if true {
        var v13 int = -2147483648
        v14 := []int{}
        p2 = p2
        _ = v13
        _ = v14
    } else {
        fmt.Println(v12, p2, p0)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return (p0 << ((114) & 63))
}
func genFn1(p0 uint32, p1 uint, p2 uint32) int16 {
    if false {
        switch -755 {
        case 32768:
            var v15 float64 = 3.14159
            _ = v15
        case -524:
            for fi16 := 0; fi16 < 17; fi16++ {
                for fi17 := 0; fi17 < 11; fi17++ {
                    p0 += 490
                    _ = fi17
                }
                _ = fi16
            }
        }
        p0 -= (p2 + 827)
        var v18 uint8 = 31
        _ = v18
    } else {
        fmt.Println(p2, p2, p0)
        v19 := []int{}
        _ = v19
    }
    if (false || (!false)) {
        p2 = p0
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    var v1 int = 42
    var v2 uint = 981
    var v3 bool = true
    v4 := make(map[int]string)
    var v5 string = ""
    v5 = v5
    for fi20 := 0; fi20 < 19; fi20++ {
        v3 = (!(true || v3))
        v21 := GenS1{F0: ("" + "hello"), F1: 399, F2: 644}
        _ = fi20
        _ = v21
    }
    v22 := v4[v1]
    v3 = true
    var v23 float64 = 1e10
    v24 := GenS1{F0: (v22 + v22), F1: 2147483648, F2: 957}
    for fi25 := 0; fi25 < 14; fi25++ {
        for fi26 := 0; fi26 < 7; fi26++ {
            if v3 {
                switch (v1 % ((fi26) | 1)) {
                case -847:
                    var v27 int64 = 9223372036854775807
                    _ = v27
                }
            }
            _ = fi26
        }
        var v28 uint64 = 0
        var v29 uint = 311
        _ = fi25
        _ = v28
        _ = v29
    }
    v30 := GenS1{F0: v5, F1: 65535, F2: 2147483647}
    v31 := GenS1{F0: (v22 + v5), F1: 2, F2: -1}
    v32 := []int{v1, v1, v1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
