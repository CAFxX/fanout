package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 int64; F1 int8 }
type GenS4 struct { F0 uint16; F1 uint64; F2 int64 }
func genFn0(p0 int16, p1 string, p2 bool) int {
    v7 := GenS1{F0: 4294967295, F1: -955}
    if true {
        switch 128 {
        case -676:
            var v8 uint8 = 64
            _ = v8
        case 127:
            v9 := GenS1{F0: 50, F1: 990}
            _ = v9
        default:
            v10 := GenS2{F0: p0, F1: 100}
            _ = v10
        }
        switch 9223372036854775807 {
        case 4294967295:
            v11 := GenS1{F0: 619, F1: 2147483647}
            _ = v11
        case 1:
            var v12 uint8 = 42
            _ = v12
        default:
            v13 := GenS3{F0: 1, F1: -11}
            _ = v13
        }
    }
    var v14 uint64 = 998
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v7
    return 1
}
func genFn1() int64 {
    v15 := []int{-1, -9223372036854775808, 128}
    switch 32767 {
    case -1:
        var v16 float64 = ((1e10 - (3.14159 / (0.1 / 1.5))) - (1e10 / 1.0))
        _ = v16
    case -2147483648:
        var v17 string = ("a" + "")
        _ = v17
    }
    var v18 int = 4294967295
    v18 -= v18
    _ = v15
    _ = v18
    return 42
}

func main() {
    v1 := []int{1, 128, 1, 1}
    var v2 int = 42
    v3 := make(map[int]string)
    var v4 float32 = float32((1e10 / 3.14159))
    var v5 int64 = -1
    var v6 int = (v2 - (v2 % (((-(v2))) | 1)))
    var v19 uint32 = 2
    switch (("a" + "") + "a") {
    case "\n\t":
        v2 = v6
    }
    v6 -= (v2 / (((v6 / ((-2147483648) | 1))) | 1))
    for fi20 := 0; fi20 < 4; fi20++ {
        var v21 uint16 = 529
        _ = fi20
        _ = v21
    }
    fmt.Println(v5)
    v3[(v6 | (v6 % (((v6 >> ((505) & 63))) | 1)))] = "a"
    v22 := v3[v6]
    v23 := GenS3{F0: v5, F1: 127}
    var v24 int16 = 1
    for fi25 := 0; fi25 < 11; fi25++ {
        for fi26 := 0; fi26 < 16; fi26++ {
            var v27 bool = false
            v28 := GenS1{F0: 42, F1: 212}
            _ = fi26
            _ = v27
            _ = v28
        }
        _ = fi25
    }
    var v29 string = "hello"
    v19 = 306
    if true {
        var v30 bool = (1 == 1)
        var v31 uint = 0
        if v30 {
            v4 -= (((v4 + float32(-2.25)) + float32(0.0)) - v4)
        } else {
            v32 := v3[(-(v6))]
            _ = v32
        }
        _ = v30
        _ = v31
    } else {
        var v33 int32 = -2147483648
        _ = v33
    }
    v34 := GenS4{F0: 255, F1: 42, F2: (^(v5))}
    switch v29 {
    case "hello":
        var v35 int = 2147483647
        _ = v35
    default:
        v36 := GenS2{F0: (^(v24)), F1: 1}
        _ = v36
    }
    switch v29 {
    case "foo bar":
        var v37 uint16 = 1
        _ = v37
    case "a":
        var v38 uint = 56
        _ = v38
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v29
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
