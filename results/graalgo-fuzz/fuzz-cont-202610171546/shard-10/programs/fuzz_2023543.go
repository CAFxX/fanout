package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32; F1 bool }
func genFn0(p0 int32, p1 uint8) int64 {
    p0 = (p0 / ((p0) | 1))
    v5 := GenS1{F0: -2147483648}
    _ = p0
    _ = p1
    _ = v5
    return 42
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 float64 = 1e10
    var v3 int16 = 546
    v4 := make(map[uint32]uint32)
    v3 = v3
    fmt.Println(v2)
    var v6 string = ("foo bar" + "a")
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 uint64 = 613
        _ = fi7
        _ = v8
    }
    var v9 uint8 = 2
    v3 -= 1000
    v10 := []int{128, 635}
    v6 = ("a" + ("\n\t" + (v6 + "\n\t")))
    switch -893 {
    case -327:
        delete(v4, 2147483648)
    }
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    v11 := GenS2{F0: -1, F1: (false || (v6 == v6))}
    fmt.Println(v2)
    if ((!(true && (1 >= 32767))) || true) {
        v6 = "foo bar"
    } else {
        for fi12 := 0; fi12 < 17; fi12++ {
            v13 := GenS2{F0: 0, F1: (!false)}
            fmt.Println(fi12, v9)
            _ = fi12
            _ = v13
        }
    }
    var v14 float32 = ((1e10 + 0.1) - 1.5)
    if (((-505 < 677) && false) && (3.14159 == v14)) {
        v15 := GenS2{F0: 2147483647, F1: true}
        var v16 uint64 = 18446744073709551615
        _ = v15
        _ = v16
    } else {
        for fi17 := 0; fi17 < 11; fi17++ {
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
            v18 := GenS2{F0: int32(v3), F1: (v9 != v9)}
            _ = fi17
            _ = v18
        }
    }
    var v19 int32 = 1
    var v20 int16 = (v3 | 884)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v6
    _ = v9
}
