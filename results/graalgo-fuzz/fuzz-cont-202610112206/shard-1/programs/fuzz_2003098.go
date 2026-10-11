package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 uint8; F2 uint64 }
func genFn0(p0 string, p1 uint16) int32 {
    fmt.Println(p0, p0)
    var v4 uint32 = 1
    fmt.Println(v4, v4)
    if true {
        var v5 int16 = 240
        if (!(65535 == 785)) {
            v4 += v4
            var v6 uint64 = 42
            _ = v6
        }
        _ = v5
    }
    _ = p0
    _ = p1
    _ = v4
    return 2147483647
}
func genFn1(p0 int32, p1 string, p2 uint64) int64 {
    var v7 uint64 = 18446744073709551615
    switch 4294967295 {
    case -792:
        p1 = "\n\t"
    case 249:
        if (false || false) {
            v8 := GenS1{F0: 2}
            _ = v8
        } else {
            v7 += v7
            if true {
                switch p1 {
                case "a":
                    var v9 int64 = 1
                    _ = v9
                default:
                    for fi10 := 0; fi10 < 5; fi10++ {
                        var v11 int8 = 42
                        var v12 uint32 = 2
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                }
                p2 = (p2 - v7)
            }
            p1 = p1
        }
    case 9223372036854775807:
        var v13 int64 = -9
        _ = v13
    }
    if ((!false) && false) {
        switch (p1 + p1) {
        case "foo bar":
            var v14 uint32 = 112
            _ = v14
        case "a":
            for fi15 := 0; fi15 < 20; fi15++ {
                var v16 int16 = 42
                if ((255 <= 0) || false) {
                    if (false || false) {
                        var v17 int32 = 42
                        var v18 int16 = 1000
                        var v19 uint16 = 1
                        _ = v17
                        _ = v18
                        _ = v19
                    }
                    p0 -= (p0 - p0)
                    v20 := GenS1{F0: uint8(p2)}
                    _ = v20
                }
                _ = fi15
                _ = v16
            }
        }
        var v21 int32 = (p0 ^ (p0 * p0))
        _ = v21
    }
    p2 -= (v7 % ((1) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 556
}

func main() {
    var v1 int16 = 1
    var v2 uint32 = 2147483648
    var v3 uint64 = 4294967295
    for fi22 := 0; fi22 < 7; fi22++ {
        fmt.Println(v1, v2, fi22)
        for fi23 := 0; fi23 < 2; fi23++ {
            v2 -= v2
            _ = fi23
        }
        _ = fi22
    }
    v24 := GenS1{F0: 42}
    fmt.Println(v1)
    v25 := GenS2{F0: (v2 * 2147483647), F1: 2, F2: v3}
    var v26 bool = ((v1 ^ 1000) == (v1 + (-(v1))))
    if (("\n\t" > "hello") && (2 <= 255)) {
        v2 -= v2
        for fi27 := 0; fi27 < 5; fi27++ {
            var v28 string = ""
            _ = fi27
            _ = v28
        }
    } else {
        var v29 bool = false
        _ = v29
    }
    fmt.Println(v1)
    for fi30 := 0; fi30 < 1; fi30++ {
        var v31 float64 = float64(1e10)
        _ = fi30
        _ = v31
    }
    var v32 uint64 = v3
    if (-304 >= genFn0((("a" + "hello") + "foo bar"), 2)) {
        v26 = (!(!(!v26)))
    }
    v33 := []int{}
    v3 = 0
    fmt.Println(v3, v26)
    var v34 float64 = float64(0.1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v32
    _ = v33
    _ = v34
}
