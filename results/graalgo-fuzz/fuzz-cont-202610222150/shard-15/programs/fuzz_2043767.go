package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 uint16 }
type GenS2 struct { F0 int64; F1 bool; F2 int16 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 uint16) uint16 {
    for fi6 := 0; fi6 < 18; fi6++ {
        if (!false) {
            p0 = p0
        } else {
            switch 169 {
            case -515:
                p0 = (p0 >> ((0) & 63))
            case 4294967295:
                for fi7 := 0; fi7 < 18; fi7++ {
                    if (87 >= 33) {
                        var v8 float64 = 1.5
                        _ = v8
                    }
                    switch "\n\t" {
                    case "foo bar":
                        switch (fi7 << ((4294967295) & 63)) {
                        case 91:
                            v9 := GenS1{F0: (!false), F1: int8(fi6), F2: 2}
                            _ = v9
                        case 42:
                            v10 := []int{fi7}
                            _ = v10
                        case -2147483648:
                            p0 = 0
                        default:
                            v11 := []int{(fi7 << ((65535) & 63)), fi6}
                            _ = v11
                        }
                    default:
                        if false {
                            v12 := GenS1{F0: true, F1: -128, F2: p0}
                            var v13 int16 = 32767
                            _ = v12
                            _ = v13
                        } else {
                            var v14 uint = 0
                            var v15 string = ((("\n\t" + "foo bar") + (("hello" + ("foo bar" + ("hello" + "\n\t"))) + "")) + "\n\t")
                            _ = v14
                            _ = v15
                        }
                    }
                    switch 1 {
                    case 127:
                        v16 := GenS1{F0: true, F1: 42, F2: (p0 + p0)}
                        _ = v16
                    }
                    _ = fi7
                }
            default:
                v17 := GenS1{F0: false, F1: -77, F2: (p0 << ((65535) & 63))}
                _ = v17
            }
            p0 = p0
            v18 := GenS2{F0: 118, F1: true, F2: 1000}
            _ = v18
        }
        var v19 float32 = 0.0
        _ = fi6
        _ = v19
    }
    v20 := GenS2{F0: 1, F1: false, F2: -658}
    fmt.Println(p0)
    fmt.Println(p0)
    p0 = (^(p0))
    _ = p0
    _ = v20
    return p0
}
func genFn1(p0 int16) uint {
    v21 := []int{127, 4294967295, 2147483647}
    var v22 int16 = p0
    _ = p0
    _ = v21
    _ = v22
    return 1
}

func main() {
    v1 := []uint32{0, 0, 65535, 552}
    var v2 int = 32768
    var v3 float64 = 0.0
    var v4 int = (v2 * v2)
    var v5 float64 = float64(1e10)
    for fi23 := 0; fi23 < 14; fi23++ {
        for fi24 := 0; fi24 < 12; fi24++ {
            if ((true && (!(!("hello" > "")))) || (32767 > 536)) {
                var v25 string = "foo bar"
                _ = v25
            }
            fmt.Println(fi23, v2, v4)
            _ = fi24
        }
        _ = fi23
    }
    v5 = v5
    var v26 int32 = -1
    v2 = v2
    if (901 > 657) {
        for fi27 := 0; fi27 < 3; fi27++ {
            var v28 uint64 = 0
            if len(v1) > 0 {
                v1[0] = 491
            }
            _ = fi27
            _ = v28
        }
    }
    fmt.Println(v3, v2)
    var v29 uint32 = 2
    v30 := GenS3{F0: 4294967295}
    v4 = v2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
