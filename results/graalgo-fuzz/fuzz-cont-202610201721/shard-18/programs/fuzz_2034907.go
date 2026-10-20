package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32; F1 int; F2 uint }
type GenS4 struct { F0 uint16; F1 bool }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int }
func genFn0(p0 string) uint16 {
    p0 = (p0 + (p0 + ("\n\t" + ("" + "\n\t"))))
    v7 := GenS1{F0: 389, F1: -1, F2: 2}
    p0 = ""
    var v8 uint8 = 0
    v9 := GenS2{F0: -2147483648}
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 202
}
func genFn1() uint {
    var v10 int = -580
    var v11 uint8 = 1
    var v12 uint = uint(v11)
    switch ("" + "") {
    case "hello":
        v13 := GenS3{F0: 4294967295, F1: (v10 ^ (v10 * (v10 >> (((^(v12))) & 63)))), F2: (v12 | v12)}
        _ = v13
    case "":
        fmt.Println(v11)
    }
    _ = v10
    _ = v11
    _ = v12
    return (^(v12))
}

func main() {
    var v1 uint32 = 506
    var v2 int8 = -75
    v3 := []uint32{(v1 - 0), v1}
    v4 := make(map[int]uint32)
    v5 := make(map[uint32]bool)
    var v6 int32 = 2147483647
    v14 := v4[0]
    for fi15 := 0; fi15 < 20; fi15++ {
        if true {
            v1 = v1
            var v16 int64 = 9223372036854775807
            _ = v16
        }
        var v17 string = "hello"
        _ = fi15
        _ = v17
    }
    switch 0 {
    case 127:
        switch 271 {
        case 9223372036854775807:
            v14 = (v14 - 42)
        case 128:
            fmt.Println(v6, v6)
        default:
            v18 := GenS4{F0: 42, F1: (!true)}
            _ = v18
        }
    }
    if len(v3) > 0 {
        v3[0] = 2147483648
    }
    if (!(-2.25 != 0.1)) {
        if (9223372036854775807 != -728) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        } else {
            for fi19 := 0; fi19 < 12; fi19++ {
                var v20 int32 = (-(v6))
                _ = fi19
                _ = v20
            }
        }
        v2 += 66
    } else {
        v21 := GenS5{F0: 1}
        fmt.Println(v14, v14, v1)
        var v22 int32 = (v6 & v6)
        _ = v21
        _ = v22
    }
    var v23 uint16 = 42
    switch "a" {
    case "a":
        var v24 float64 = 0.0
        _ = v24
    case "":
        if true {
            if ((uint16(v1) <= uint16(v1)) || (false && (!true))) {
                v1 = (v1 | v1)
                fmt.Println(v1)
            } else {
                v14 = v14
            }
            v23 -= v23
        } else {
            switch -483 {
            case -2147483648:
                v23 = uint16(v14)
            case 0:
                for fi25 := 0; fi25 < 20; fi25++ {
                    v23 += (v23 % (((v23 | 1)) | 1))
                    fmt.Println(v1, v23, v1)
                    _ = fi25
                }
            case -74:
                v1 = 0
            default:
                v3 = append(v3, (v14 / ((961) | 1)))
            }
        }
    default:
        for fi26 := 0; fi26 < 18; fi26++ {
            for fi27 := 0; fi27 < 17; fi27++ {
                switch "\n\t" {
                case "hello":
                    if true {
                        v28 := GenS6{F0: fi26}
                        fmt.Println(fi26, v14)
                        delete(v4, 32767)
                        _ = v28
                    }
                default:
                    var v29 uint8 = 127
                    switch 725 {
                    case -2147483648:
                        var v30 int8 = (v2 >> ((0) & 63))
                        _ = v30
                    case 1:
                        v1 -= (v14 & v14)
                    case 176:
                        if (true || false) {
                            v31 := GenS5{F0: 42}
                            var v32 int32 = 42
                            v14 += (v1 * v14)
                            _ = v31
                            _ = v32
                        } else {
                            v14 = 627
                        }
                    }
                    _ = v29
                }
                _ = fi27
            }
            _ = fi26
        }
    }
    if (false || (!false)) {
        var v33 bool = (genFn1() != 0)
        v34 := v5[v1]
        delete(v4, -584)
        _ = v33
        _ = v34
    }
    var v35 uint8 = 2
    switch "\n\t" {
    case "\n\t":
        if ((v6 >> ((255) & 63)) >= (v6 | -2147483648)) {
            v14 = v1
            v2 = 100
            fmt.Println(v2, v1)
        } else {
            for fi36 := 0; fi36 < 8; fi36++ {
                var v37 bool = ("hello" >= ("a" + ""))
                var v38 int = (-(fi36))
                var v39 bool = (false || v37)
                _ = fi36
                _ = v37
                _ = v38
                _ = v39
            }
            v3 = append(v3, uint32(v23))
        }
    }
    var v40 float32 = (0.1 - ((0.1 - (0.1 + 0.0)) / ((float32(float64(0.1)) + (1e10 / -2.25)) + (3.14159 + 1.5))))
    var v41 int32 = int32(v14)
    var v42 float32 = v40
    var v43 int64 = -209
    v14 -= v1
    v6 = (v41 / ((v6) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v23
    _ = v3
    _ = v35
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
