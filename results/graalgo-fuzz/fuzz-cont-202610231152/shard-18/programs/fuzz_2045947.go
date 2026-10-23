package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 int }
func genFn0(p0 int16, p1 uint8, p2 uint8) int64 {
    v6 := GenS1{F0: 0, F1: p0}
    if false {
        for fi7 := 0; fi7 < 13; fi7++ {
            v8 := []int{fi7, 4294967295, -301, fi7}
            for fi9 := 0; fi9 < 10; fi9++ {
                if (!false) {
                    p1 += (p2 / ((128) | 1))
                    var v10 uint = 997
                    if ((true || true) || false) {
                        switch fi9 {
                        case 128:
                            var v11 float64 = (float64((1.5 + (0.1 / 1.0))) - 3.14159)
                            _ = v11
                        case 2147483647:
                            var v12 string = "a"
                            _ = v12
                        case 4294967295:
                            _ = len(v8)
                        default:
                            fmt.Println(p2)
                        }
                    } else {
                        v10 = 396
                    }
                    _ = v10
                }
                var v13 string = (("" + ("a" + "a")) + "hello")
                var v14 uint8 = (p2 & p1)
                _ = fi9
                _ = v13
                _ = v14
            }
            _ = fi7
            _ = v8
        }
    }
    var v15 uint16 = 255
    if true {
        v15 = (v15 * v15)
        v15 = (v15 ^ v15)
    } else {
        if true {
            var v16 int = 2147483647
            v17 := []int{(^(v16)), v16, 705}
            _ = v16
            _ = v17
        } else {
            p0 = p0
        }
        for fi18 := 0; fi18 < 20; fi18++ {
            v19 := GenS2{F0: -579}
            switch fi18 {
            case 2147483647:
                var v20 float64 = (((((0.1 + 1.5) - ((1.5 / 1e10) - -2.25)) * (1e10 + -2.25)) + (float64(-2.25) + 0.0)) / 0.1)
                _ = v20
            case 42:
                v21 := []int{354}
                _ = v21
            }
            _ = fi18
            _ = v19
        }
    }
    var v22 uint16 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v22
    _ = v6
    return 68
}
func genFn1() uint64 {
    var v23 int32 = -504
    v24 := GenS1{F0: 80, F1: 42}
    v23 = v23
    v23 = (v23 << ((255) & 63))
    _ = v23
    _ = v24
    return 18446744073709551615
}

func main() {
    var v1 uint8 = 127
    var v2 uint16 = 446
    var v3 bool = (v2 < (v2 | (v2 ^ v2)))
    var v4 string = "foo bar"
    v5 := []uint32{}
    v25 := GenS1{F0: 128, F1: -656}
    for fi26 := 0; fi26 < 18; fi26++ {
        v3 = v3
        if v3 {
            for fi27 := 0; fi27 < 3; fi27++ {
                v1 = v1
                if v3 {
                    var v28 uint64 = 4294967295
                    v1 = 104
                    _ = v28
                }
                v5 = append(v5, 0)
                _ = fi27
            }
        }
        _ = fi26
    }
    var v29 float64 = -2.25
    v29 += v29
    switch 4294967295 {
    case 2147483647:
        for fi30 := 0; fi30 < 7; fi30++ {
            if (((0.1 + (v29 / ((((v29 - 0.0) + v29) * v29) - v29))) <= v29) || v3) {
                if len(v5) > 0 {
                    v5[0] = 376
                }
                fmt.Println(v29, v2, v4)
            }
            if len(v5) > 0 {
                v5[0] = 2147483648
            }
            _ = fi30
        }
    case 32768:
        v3 = (v1 >= (v1 << ((2) & 63)))
    case 4294967295:
        var v31 uint16 = (^(v2))
        _ = v31
    }
    for fi32 := 0; fi32 < 2; fi32++ {
        fmt.Println(v3, v3)
        _ = fi32
    }
    switch (v4 + v4) {
    case "\n\t":
        var v33 uint32 = 2147483647
        _ = v33
    case "hello":
        v34 := GenS2{F0: 473}
        _ = v34
    }
    v29 = (v29 / 1e10)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
