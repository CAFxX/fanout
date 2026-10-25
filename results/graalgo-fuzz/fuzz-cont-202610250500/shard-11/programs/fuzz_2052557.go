package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 uint32 }
func genFn0(p0 int64, p1 uint, p2 uint16) bool {
    p2 -= p2
    v7 := []int{-620}
    fmt.Println(p0, p1)
    for fi8 := 0; fi8 < 13; fi8++ {
        v9 := GenS1{F0: 229, F1: true, F2: 42}
        var v10 string = ("" + "a")
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return false
}
func genFn1(p0 int16, p1 int16, p2 bool) uint8 {
    p2 = (p2 || true)
    fmt.Println(p0, p0)
    var v11 int = -305
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 0
}

func main() {
    var v1 int = -9223372036854775808
    var v2 int64 = -148
    v3 := make(map[int]uint32)
    var v4 float64 = 0.1
    var v5 int64 = -336
    v6 := make(map[uint32]int)
    v1 -= (v1 + v1)
    v5 = (v2 | v2)
    if true {
        if false {
            if true {
                for fi12 := 0; fi12 < 15; fi12++ {
                    for fi13 := 0; fi13 < 18; fi13++ {
                        var v14 int16 = 1000
                        _ = fi13
                        _ = v14
                    }
                    var v15 int32 = -2147483648
                    _ = fi12
                    _ = v15
                }
            }
            fmt.Println(v4, v5)
        }
        v16 := []int{(v1 - -1), v1, v1, v1}
        switch (v1 << ((768) & 63)) {
        case 127:
            var v17 string = "hello"
            _ = v17
        case -1:
            v1 -= v1
        default:
            v18 := GenS1{F0: 1, F1: (!false), F2: 2147483648}
            _ = v18
        }
        _ = v16
    } else {
        v19 := GenS1{F0: 126, F1: true, F2: 587}
        for fi20 := 0; fi20 < 16; fi20++ {
            v2 += (-(v5))
            switch (v1 << ((2) & 63)) {
            case 274:
                var v21 uint16 = 255
                _ = v21
            default:
                v22 := GenS1{F0: 0, F1: false, F2: 65535}
                _ = v22
            }
            for fi23 := 0; fi23 < 18; fi23++ {
                var v24 int8 = 0
                _ = fi23
                _ = v24
            }
            _ = fi20
        }
        _ = v19
    }
    var v25 float32 = 1.5
    var v26 uint32 = 898
    var v27 int = (^(v1))
    fmt.Println(v26, v2)
    v2 += v5
    switch "foo bar" {
    case "foo bar":
        fmt.Println(v5, v1)
    }
    var v28 uint8 = 0
    v29 := GenS1{F0: genFn1(1, -1, ((v26 >= (v26 - 65535)) && (v2 <= 1))), F1: true, F2: v26}
    v4 = (1.5 + v4)
    if ((v2 - (v5 / (((v5 << ((253) & 63))) | 1))) == v5) {
        v30 := []int{(v27 - v27), (-(v27))}
        _ = v30
    }
    var v31 uint = 42
    for fi32 := 0; fi32 < 15; fi32++ {
        v33 := []int{(v1 - (fi32 - 2147483647)), (v1 - fi32), 128, v27}
        _ = fi32
        _ = v33
    }
    v34 := []int{}
    v2 = (-(v2))
    _ = len(v34)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v31
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
