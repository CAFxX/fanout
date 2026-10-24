package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 int32; F1 uint64; F2 int16 }
func genFn0(p0 int32, p1 uint) bool {
    if ((false || true) || (!(false || false))) {
        if (!false) {
            p0 += p0
            var v5 int = -9223372036854775808
            _ = v5
        }
        v6 := GenS1{F0: 2, F1: true}
        v7 := GenS1{F0: 562, F1: (!false)}
        _ = v6
        _ = v7
    }
    var v8 float32 = 1.5
    v9 := []int{32767}
    switch int(p0) {
    case 127:
        p1 -= p1
    case 2147483647:
        var v10 uint64 = 1
        _ = v10
    }
    if false {
        if len(v9) > 0 {
            v9[0] = 831
        }
        if (1000 >= -1) {
            var v11 uint16 = 255
            var v12 int = 127
            v11 = (v11 % ((255) | 1))
            _ = v11
            _ = v12
        } else {
            var v13 uint = p1
            for fi14 := 0; fi14 < 1; fi14++ {
                var v15 uint16 = 65535
                for fi16 := 0; fi16 < 6; fi16++ {
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                    _ = fi16
                }
                if (-9223372036854775808 > -775) {
                    if true {
                        v13 += (p1 / (((^(v13))) | 1))
                        var v17 uint32 = 707
                        switch "" {
                        case "hello":
                            fmt.Println(p1, fi14, fi14)
                        case "":
                            var v18 uint16 = v15
                            _ = v18
                        }
                        _ = v17
                    }
                }
                _ = fi14
                _ = v15
            }
            _ = len(v9)
            _ = v13
        }
    }
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return (137 < 2)
}

func main() {
    var v1 float32 = float32(float64(1.5))
    var v2 string = ""
    v3 := []int{-715, 1, 113}
    v4 := []int{9223372036854775807, -9223372036854775808}
    v1 -= v1
    v4 = append(v4, -504)
    var v19 bool = (!(0 <= 2))
    v20 := GenS2{F0: 42, F1: 4294967295, F2: 832}
    if (v2 > v2) {
        var v21 string = (v2 + "foo bar")
        switch 603 {
        case 485:
            var v22 uint64 = 114
            _ = v22
        case 2147483647:
            fmt.Println(v19, v2, v2)
        default:
            fmt.Println(v2, v2)
        }
        var v23 uint16 = 693
        _ = v21
        _ = v23
    }
    v1 = (v1 + v1)
    var v24 float64 = -2.25
    for fi25 := 0; fi25 < 8; fi25++ {
        for fi26 := 0; fi26 < 2; fi26++ {
            switch fi25 {
            case -9223372036854775808:
                var v27 int16 = -1
                _ = v27
            case -377:
                if true {
                    v4 = append(v4, 9223372036854775807)
                }
            case 1:
                if len(v4) > 0 {
                    v4[0] = 32767
                }
            default:
                fmt.Println(fi25)
            }
            _ = fi26
        }
        v28 := GenS2{F0: 1, F1: 0, F2: 0}
        _ = fi25
        _ = v28
    }
    var v29 uint32 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v29
    _ = v3
    _ = v4
}
