package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 int; F1 int32; F2 uint16 }
func genFn0(p0 int64, p1 uint8) bool {
    v7 := GenS1{F0: 0, F1: 85}
    p1 += (p1 * 255)
    _ = p0
    _ = p1
    _ = v7
    return (1 == 65535)
}
func genFn1(p0 int, p1 int) uint8 {
    v8 := []int{p1}
    if (float64(0.0) < 0.1) {
        var v9 uint16 = 630
        _ = v9
    }
    p1 = (-(p1))
    for fi10 := 0; fi10 < 5; fi10++ {
        _ = len(v8)
        for fi11 := 0; fi11 < 8; fi11++ {
            var v12 int8 = 1
            var v13 uint = 1
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = v8
    return 127
}

func main() {
    v1 := []string{"hello"}
    v2 := make(map[string]bool)
    v3 := []string{}
    var v4 uint16 = 77
    v5 := []uint32{65535, 0}
    var v6 float32 = -2.25
    v6 = 1.5
    if false {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        fmt.Println(v4)
    } else {
        var v14 bool = (true && true)
        _ = v14
    }
    var v15 uint64 = 4294967295
    var v16 float64 = 0.1
    fmt.Println(v15)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println(v16, v4, v6)
    v17 := GenS1{F0: 8, F1: 0}
    switch 127 {
    case -964:
        if (v16 <= 0.0) {
            v18 := GenS2{F0: -877, F1: 2147483647, F2: (v4 & (v4 - (^(v4))))}
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v2["foo bar"] = true
            _ = v18
        } else {
            var v19 uint8 = 42
            var v20 int8 = 1
            fmt.Println(v19, v16)
            _ = v19
            _ = v20
        }
    default:
        delete(v2, ("" + ((("hello" + ((("" + (("foo bar" + ("a" + "")) + "foo bar")) + "a") + "\n\t")) + "foo bar") + "\n\t")))
    }
    var v21 float64 = v16
    v21 -= ((v21 + v16) - (v21 - 0.1))
    var v22 int64 = -271
    var v23 uint32 = 291
    var v24 float32 = ((0.1 * (float32((v16 + float64(v6))) / ((v6 * (v6 * v6)) + v6))) * (float32(v21) / 0.1))
    if (!true) {
        var v25 int = 128
        v26 := GenS2{F0: (-(v25)), F1: int32(v4), F2: (v4 >> ((797) & 63))}
        _ = v25
        _ = v26
    }
    delete(v2, (("hello" + "hello") + (("foo bar" + "a") + "")))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
