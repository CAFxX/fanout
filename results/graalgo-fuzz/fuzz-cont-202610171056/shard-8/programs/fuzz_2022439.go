package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 bool }
func genFn0(p0 uint32, p1 int8, p2 int8) int {
    if (("foo bar" + ("foo bar" + ("" + ""))) >= "hello") {
        switch "a" {
        case "\n\t":
            v6 := GenS1{F0: p0, F1: -755, F2: (!(p1 >= p2))}
            _ = v6
        case "a":
            if true {
                if (227 > 0) {
                    fmt.Println(p1)
                } else {
                    fmt.Println(p1, p2)
                    for fi7 := 0; fi7 < 14; fi7++ {
                        var v8 uint = 18446744073709551615
                        if ((v8 ^ (v8 % ((v8) | 1))) > v8) {
                            switch (fi7 + 32767) {
                            case 42:
                                v9 := GenS1{F0: p0, F1: 754, F2: true}
                                _ = v9
                            case 32768:
                                p2 = (p2 << (((v8 / ((2) | 1))) & 63))
                            }
                        } else {
                            v10 := GenS1{F0: p0, F1: (fi7 / ((fi7) | 1)), F2: (!true)}
                            p2 -= -128
                            for fi11 := 0; fi11 < 12; fi11++ {
                                var v12 uint = v8
                                var v13 int16 = 1000
                                if (1.5 >= ((float64(-2.25) - (-2.25 + 0.1)) - 1e10)) {
                                    v14 := GenS1{F0: p0, F1: fi7, F2: (p1 <= (-(p1)))}
                                    fmt.Println(p2)
                                    _ = v14
                                }
                                _ = fi11
                                _ = v12
                                _ = v13
                            }
                            _ = v10
                        }
                        v15 := GenS1{F0: 278, F1: 32768, F2: (0.1 > 0.0)}
                        _ = fi7
                        _ = v15
                        _ = v8
                    }
                    var v16 uint = 2
                    _ = v16
                }
                v17 := []int{}
                p2 = 0
                _ = v17
            }
        case "foo bar":
            p1 = 95
        default:
            p1 += -88
        }
    }
    p0 += (p0 ^ 490)
    v18 := []int{2147483647}
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    return -1
}
func genFn1() uint8 {
    var v19 int = 32768
    var v20 uint = 679
    _ = v19
    _ = v20
    return 2
}

func main() {
    v1 := []uint32{42}
    var v2 uint32 = 0
    v3 := []uint32{v2, (v2 + v2), v2, v2}
    var v4 int8 = 42
    var v5 int32 = 204
    fmt.Println(v4)
    if len(v1) > 0 {
        v1[0] = 42
    }
    for fi21 := 0; fi21 < 13; fi21++ {
        var v22 bool = true
        _ = fi21
        _ = v22
    }
    var v23 float64 = (-2.25 * 1.5)
    var v24 uint64 = 0
    var v25 uint16 = 65535
    var v26 bool = true
    v27 := GenS1{F0: (v2 ^ v2), F1: 2147483647, F2: (128 <= 42)}
    var v28 int64 = -860
    if (((3.14159 * (1e10 * 0.1)) != 0.1) || true) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    } else {
        if len(v3) > 0 {
            v3[0] = v2
        }
        var v29 int32 = v5
        _ = v29
    }
    v28 = v28
    var v30 int8 = 24
    v4 += 100
    var v31 int8 = (v4 & (v4 | (v4 / (((v4 ^ v30)) | 1))))
    switch "a" {
    case "foo bar":
        fmt.Println(v26)
    default:
        var v32 float32 = 3.14159
        _ = v32
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
