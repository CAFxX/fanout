package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 uint16 }
type GenS2 struct { F0 int32 }
func genFn0() int8 {
    var v7 string = ((("foo bar" + "hello") + "foo bar") + "a")
    var v8 uint = 255
    for fi9 := 0; fi9 < 18; fi9++ {
        v10 := []int{(fi9 / ((-444) | 1))}
        if true {
            fmt.Println(v7, v8, fi9)
        }
        _ = fi9
        _ = v10
    }
    _ = v7
    _ = v8
    return -63
}
func genFn1(p0 string, p1 int64, p2 int64) uint64 {
    p1 = (p1 + p1)
    if (false || true) {
        v11 := []int{}
        var v12 int64 = p2
        v12 += p2
        _ = v11
        _ = v12
    }
    v13 := GenS1{F0: true, F1: 4294967295, F2: 958}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 1
}

func main() {
    var v1 float32 = float32(3.14159)
    var v2 uint32 = 484
    var v3 uint16 = 0
    var v4 int32 = 1
    v5 := make(map[string]uint32)
    var v6 uint64 = 414
    var v14 float64 = 1.5
    v3 -= v3
    if true {
        switch ("foo bar" + "") {
        case "a":
            if (!false) {
                fmt.Println(v14, v2)
                fmt.Println(v2, v4)
                var v15 int64 = 1
                _ = v15
            } else {
                var v16 string = "\n\t"
                _ = v16
            }
        default:
            v17 := []int{int(v3)}
            _ = v17
        }
        if false {
            v2 -= (v2 * v2)
        }
    } else {
        var v18 int64 = 196
        var v19 string = ("foo bar" + "")
        v6 = (v6 + genFn1(("hello" + v19), (v18 / (((v18 << ((2) & 63))) | 1)), v18))
        _ = v18
        _ = v19
    }
    if (!(v3 <= (v3 * (^(v3))))) {
        var v20 int32 = v4
        _ = v20
    } else {
        fmt.Println(v14, v1, v2)
    }
    v21 := []int{int(v2), -25, -366, -616}
    var v22 uint32 = (v2 >> ((4294967295) & 63))
    _ = len(v5)
    var v23 float64 = float64(-2.25)
    fmt.Println(v4, v3, v4)
    v6 -= (^(v6))
    v14 += (v23 - v14)
    fmt.Println(v3, v6)
    v5[("" + "hello")] = (v2 % ((v2) | 1))
    var v24 uint32 = (v22 << ((1) & 63))
    var v25 string = "foo bar"
    for fi26 := 0; fi26 < 3; fi26++ {
        if len(v21) >= 2 {
            v21 = v21[1:len(v21)]
        }
        v27 := GenS1{F0: false, F1: (v6 & uint64(v22)), F2: v3}
        _ = fi26
        _ = v27
    }
    if true {
        var v28 bool = ((18446744073709551615 > 556) && false)
        _ = v28
    }
    v29 := GenS2{F0: (v4 >> ((0) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
