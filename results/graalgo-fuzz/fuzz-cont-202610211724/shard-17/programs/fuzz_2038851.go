package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 int; F2 uint64 }
type GenS3 struct { F0 uint; F1 int8 }
func genFn0(p0 uint16, p1 int32, p2 uint8) int16 {
    v5 := GenS1{F0: -1}
    p1 = p1
    if true {
        var v6 int8 = 49
        _ = v6
    } else {
        v7 := GenS2{F0: -558, F1: 1, F2: 709}
        p2 -= (p2 & (p2 | (^(p2))))
        _ = v7
    }
    p0 -= p0
    switch (("hello" + "a") + ("\n\t" + "a")) {
    case "foo bar":
        v8 := GenS1{F0: (p1 >> ((42) & 63))}
        _ = v8
    case "\n\t":
        var v9 int8 = -128
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 0
}
func genFn1(p0 uint64, p1 int64, p2 int8) int {
    for fi10 := 0; fi10 < 9; fi10++ {
        var v11 int16 = 42
        _ = fi10
        _ = v11
    }
    var v12 float64 = (-2.25 * (float64(float32(1e10)) * 0.1))
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 374
}

func main() {
    var v1 uint16 = 0
    var v2 float64 = (1.5 + (0.1 + 1e10))
    var v3 int8 = -5
    var v4 float64 = v2
    v4 -= v2
    for fi13 := 0; fi13 < 5; fi13++ {
        v1 = v1
        v4 = v4
        if false {
            var v14 int16 = -32768
            v3 = (v3 - 127)
            _ = v14
        }
        _ = fi13
    }
    for fi15 := 0; fi15 < 4; fi15++ {
        fmt.Println(v1, v1)
        var v16 float32 = 0.0
        v17 := []int{fi15, genFn1(138, -9223372036854775808, (v3 - v3))}
        _ = fi15
        _ = v16
        _ = v17
    }
    if (703 >= 912) {
        for fi18 := 0; fi18 < 7; fi18++ {
            var v19 uint8 = 130
            v2 += float64(3.14159)
            _ = fi18
            _ = v19
        }
        v20 := GenS1{F0: -1}
        _ = v20
    }
    var v21 bool = false
    v1 = v1
    var v22 int32 = -2147483648
    var v23 uint64 = 2
    v21 = v21
    v24 := []int{-2147483648}
    if (-2.25 < v4) {
        for fi25 := 0; fi25 < 4; fi25++ {
            v26 := GenS3{F0: 4294967295, F1: v3}
            switch "a" {
            case "hello":
                var v27 float64 = float64(((1.5 / float32((v2 + v2))) + float32((((v2 + -2.25) - v2) - 0.1))))
                _ = v27
            case "\n\t":
                v24 = append(v24, 0)
            }
            _ = fi25
            _ = v26
        }
        var v28 uint = 255
        _ = v28
    } else {
        v29 := GenS2{F0: 42, F1: 1, F2: v23}
        v23 -= 2
        _ = v29
    }
    if v21 {
        fmt.Println(v1, v4)
    }
    v24 = append(v24, -2147483648)
    var v30 int8 = -1
    var v31 int = -9223372036854775808
    v24 = append(v24, (v31 << ((951) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
