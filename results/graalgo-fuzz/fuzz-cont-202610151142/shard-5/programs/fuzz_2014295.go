package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint; F1 uint8; F2 int64 }
func genFn0(p0 uint8, p1 int8) int32 {
    fmt.Println(p1, p0)
    if (128 < -819) {
        p1 -= (p1 * (p1 >> ((255) & 63)))
        for fi7 := 0; fi7 < 17; fi7++ {
            v8 := []int{}
            _ = fi7
            _ = v8
        }
        v9 := []int{32768, 127, 127}
        _ = v9
    }
    var v10 int64 = -1
    for fi11 := 0; fi11 < 20; fi11++ {
        v12 := GenS1{F0: (p0 << ((4294967295) & 63))}
        if ((p0 >= p0) && false) {
            var v13 uint8 = p0
            var v14 int16 = 635
            var v15 int64 = (v10 | v10)
            _ = v13
            _ = v14
            _ = v15
        }
        for fi16 := 0; fi16 < 9; fi16++ {
            switch 127 {
            case 127:
                p0 += (p0 % ((128) | 1))
            case 1:
                var v17 uint32 = 4294967295
                _ = v17
            case -2147483648:
                var v18 int64 = v10
                _ = v18
            }
            var v19 bool = (299 >= 42)
            _ = fi16
            _ = v19
        }
        _ = fi11
        _ = v12
    }
    p1 = p1
    _ = p0
    _ = p1
    _ = v10
    return 70
}
func genFn1(p0 int16) int16 {
    for fi20 := 0; fi20 < 6; fi20++ {
        v21 := []int{fi20}
        _ = len(v21)
        p0 -= -32768
        _ = fi20
        _ = v21
    }
    v22 := GenS1{F0: 7}
    _ = p0
    _ = v22
    return p0
}

func main() {
    var v1 float32 = float32(0.1)
    var v2 uint32 = 0
    v3 := []int{852, 127, 127}
    var v4 uint16 = 2
    v5 := []uint32{(^(v2)), 616}
    var v6 uint16 = v4
    var v23 int32 = 42
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println(v4, v1, v23)
    v2 = v2
    if len(v3) > 0 {
        v3[0] = 1
    }
    v1 = 0.0
    fmt.Println(v2)
    v4 += (v6 ^ (^(v6)))
    var v24 uint = 550
    if (1 == -32768) {
        if ((v4 << ((uint(v2)) & 63)) <= v4) {
            if len(v5) > 0 {
                v5[0] = 742
            }
            fmt.Println(v6)
            fmt.Println(v4, v23)
        }
        v25 := GenS2{F0: (v4 <= 875), F1: 1}
        switch 4294967295 {
        case 4294967295:
            var v26 uint32 = v2
            _ = v26
        case 2147483647:
            var v27 uint32 = v2
            _ = v27
        default:
            v28 := GenS2{F0: false, F1: 255}
            _ = v28
        }
        _ = v25
    } else {
        v1 = v1
        v5 = append(v5, (^(v2)))
    }
    var v29 string = "\n\t"
    v3 = append(v3, 215)
    if (false || (v29 < v29)) {
        if (!false) {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        } else {
            var v30 int64 = -9223372036854775808
            v31 := GenS1{F0: uint8(v2)}
            v32 := GenS3{F0: 32767}
            _ = v30
            _ = v31
            _ = v32
        }
        var v33 float64 = 0.1
        v2 -= 347
        _ = v33
    }
    var v34 uint16 = v4
    for fi35 := 0; fi35 < 16; fi35++ {
        var v36 float64 = ((float64(0.0) / 0.1) * ((-2.25 + ((1.5 * (float64(3.14159) - 1.5)) * (1.5 + (float64(v1) / 1.0)))) - 3.14159))
        for fi37 := 0; fi37 < 5; fi37++ {
            if true {
                var v38 int64 = -1
                switch ("hello" + "") {
                case "hello":
                    fmt.Println(v38, v23, v34)
                case "\n\t":
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                case "":
                    if false {
                        v29 = v29
                        var v39 int64 = v38
                        v40 := GenS4{F0: uint(v4), F1: 12, F2: v38}
                        _ = v39
                        _ = v40
                    }
                default:
                    if len(v3) > 0 {
                        v3[0] = fi35
                    }
                }
                _ = v38
            }
            v1 = (v1 / v1)
            _ = fi37
        }
        for fi41 := 0; fi41 < 1; fi41++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi41
        }
        _ = fi35
        _ = v36
    }
    var v42 int32 = -332
    v34 += v34
    fmt.Println(v34, v42, v24)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v29
    _ = v3
    _ = v34
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
