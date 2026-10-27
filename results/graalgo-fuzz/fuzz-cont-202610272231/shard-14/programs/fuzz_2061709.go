package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 uint8 }
type GenS2 struct { F0 uint64; F1 string; F2 int }
type GenS3 struct { F0 uint8 }
func genFn0(p0 uint32, p1 uint16, p2 int32) uint {
    var v5 float64 = 1e10
    fmt.Println(p2, v5)
    p2 = (p2 << ((478) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 289
}

func main() {
    var v1 uint16 = 255
    var v2 int = -344
    var v3 float64 = (((0.0 - 1.5) + (-2.25 / -2.25)) / (1e10 / 1.5))
    v4 := []int{v2, v2}
    if (!(!false)) {
        _ = len(v4)
        var v6 int = v2
        _ = v6
    } else {
        for fi7 := 0; fi7 < 18; fi7++ {
            var v8 int32 = 0
            _ = fi7
            _ = v8
        }
    }
    for fi9 := 0; fi9 < 12; fi9++ {
        var v10 int8 = -128
        _ = fi9
        _ = v10
    }
    v2 = v2
    switch ("hello" + "a") {
    case "\n\t":
        if ((1.5 < ((((0.1 * -2.25) - 0.0) + -2.25) * 1e10)) || true) {
            if false {
                v3 += v3
                v4 = append(v4, v2)
                v2 = v2
            } else {
                v3 = v3
                for fi11 := 0; fi11 < 8; fi11++ {
                    v12 := GenS1{F0: 100, F1: 334, F2: 2}
                    for fi13 := 0; fi13 < 8; fi13++ {
                        fmt.Println(v2, v1, fi13)
                        v2 = (fi11 ^ fi13)
                        var v14 float64 = (float64(1e10) - (v3 - (v3 / 1.0)))
                        _ = fi13
                        _ = v14
                    }
                    v4 = append(v4, (v2 | v2))
                    _ = fi11
                    _ = v12
                }
            }
        }
    }
    switch "" {
    case "":
        if true {
            if false {
                v2 += v2
            } else {
                var v15 uint16 = 0
                _ = v15
            }
            var v16 bool = true
            v2 -= v2
            _ = v16
        }
    case "foo bar":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    default:
        for fi17 := 0; fi17 < 19; fi17++ {
            v1 += (v1 % (((v1 - (v1 >> ((65535) & 63)))) | 1))
            _ = fi17
        }
    }
    _ = len(v4)
    fmt.Println(v1, v2, v2)
    var v18 string = ("a" + "\n\t")
    if false {
        if (true || (true && true)) {
            v3 += 3.14159
        }
    } else {
        fmt.Println(v1, v3, v18)
        v19 := GenS2{F0: 18446744073709551615, F1: "", F2: -2147483648}
        _ = v19
    }
    switch v18 {
    case "foo bar":
        switch 32768 {
        case 128:
            for fi20 := 0; fi20 < 12; fi20++ {
                v4 = append(v4, 128)
                var v21 uint8 = 27
                var v22 uint = 4294967295
                _ = fi20
                _ = v21
                _ = v22
            }
        case -2147483648:
            switch v2 {
            case -724:
                if (90 < 42) {
                    if len(v4) > 0 {
                        v4[0] = (-(v2))
                    }
                    var v23 int16 = 142
                    _ = v23
                }
            default:
                if false {
                    switch (v2 - v2) {
                    case -1:
                        _ = len(v4)
                    default:
                        var v24 int16 = -1
                        _ = v24
                    }
                } else {
                    var v25 uint16 = (v1 / (((v1 | (v1 * v1))) | 1))
                    if (-9223372036854775808 >= -33) {
                        var v26 int32 = 1
                        v27 := GenS3{F0: 1}
                        _ = v26
                        _ = v27
                    }
                    _ = v25
                }
            }
        default:
            v18 = (v18 + "foo bar")
        }
    default:
        v1 = (v1 * 488)
    }
    var v28 int64 = 1
    if ((!(0 >= 0)) || (false && true)) {
        var v29 uint = 65535
        var v30 int = v2
        _ = v29
        _ = v30
    }
    if (false && false) {
        var v31 bool = true
        v32 := GenS2{F0: 946, F1: v18, F2: (-(v2))}
        _ = v31
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v28
    _ = v3
    _ = v4
}
