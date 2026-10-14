package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 uint16; F1 int8 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int; F1 int; F2 uint64 }
type GenS7 struct { F0 uint8; F1 uint16; F2 int64 }
type GenS8 struct { F0 uint64; F1 int32; F2 int32 }
func genFn0(p0 uint) int16 {
    if (true && (false || true)) {
        v4 := []int{-77, 380}
        if len(v4) > 0 {
            v4[0] = 1
        }
        _ = v4
    } else {
        if true {
            v5 := []int{-9223372036854775808, 32767}
            var v6 uint32 = 668
            for fi7 := 0; fi7 < 8; fi7++ {
                v5 = append(v5, -585)
                var v8 int16 = 42
                _ = fi7
                _ = v8
            }
            _ = v5
            _ = v6
        } else {
            v9 := []int{32767, -415}
            _ = v9
        }
        if false {
            v10 := GenS1{F0: -2147483648, F1: 4294967295}
            var v11 uint8 = 128
            _ = v10
            _ = v11
        }
        p0 = p0
    }
    if (((("foo bar" + ("foo bar" + "foo bar")) + "foo bar") + "") >= "") {
        if (((!false) || true) || false) {
            v12 := []int{310, 128, -2147483648, 9223372036854775807}
            switch "" {
            case "hello":
                if ((p0 & 65535) != p0) {
                    p0 += 955
                    for fi13 := 0; fi13 < 4; fi13++ {
                        var v14 bool = true
                        v15 := GenS2{F0: (!(v14 && (!("\n\t" == "")))), F1: 42}
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                } else {
                    var v16 uint16 = 2
                    for fi17 := 0; fi17 < 1; fi17++ {
                        v18 := GenS3{F0: (v16 ^ v16), F1: 100}
                        _ = fi17
                        _ = v18
                    }
                    if (-1 <= -275) {
                        if true {
                            if true {
                                if true {
                                    var v19 int32 = 251
                                    p0 = p0
                                    _ = v19
                                }
                                if (!(!(!true))) {
                                    var v20 float32 = ((1e10 + ((1e10 + 0.0) - ((((0.1 - 1.5) / 3.14159) + 3.14159) / float32(0.1)))) / (0.1 * (1e10 / -2.25)))
                                    var v21 float64 = float64(v20)
                                    v12 = append(v12, -78)
                                    _ = v20
                                    _ = v21
                                }
                            }
                            v16 -= v16
                        }
                        var v22 uint16 = (v16 - v16)
                        _ = v22
                    } else {
                        for fi23 := 0; fi23 < 17; fi23++ {
                            if false {
                                v24 := GenS1{F0: fi23, F1: 2}
                                _ = v24
                            } else {
                                fmt.Println(p0, p0)
                                if len(v12) > 0 {
                                    v12[0] = (fi23 % (((fi23 / (((fi23 / ((4294967295) | 1))) | 1))) | 1))
                                }
                                fmt.Println(fi23, fi23)
                            }
                            if len(v12) > 0 {
                                v12[0] = (fi23 & -32)
                            }
                            v25 := GenS1{F0: 39, F1: 0}
                            _ = fi23
                            _ = v25
                        }
                        var v26 uint16 = 963
                        v26 = (v16 - (v16 ^ (v16 % ((255) | 1))))
                        _ = v26
                    }
                    _ = v16
                }
            case "foo bar":
                _ = len(v12)
            default:
                fmt.Println(p0)
            }
            _ = v12
        } else {
            v27 := []int{127, 394, 128}
            _ = v27
        }
        var v28 int64 = 42
        fmt.Println(p0)
        _ = v28
    }
    v29 := []int{0, -509, 127}
    _ = p0
    _ = v29
    return 982
}
func genFn1(p0 int, p1 uint8) uint64 {
    var v30 string = ("a" + "hello")
    for fi31 := 0; fi31 < 7; fi31++ {
        fmt.Println(p0, fi31, fi31)
        p0 = fi31
        _ = fi31
    }
    _ = p0
    _ = p1
    _ = v30
    return 403
}

func main() {
    var v1 int = 937
    var v2 float32 = 1e10
    var v3 uint32 = 2147483647
    v32 := GenS2{F0: (false || (1 <= (v3 ^ v3))), F1: -1}
    for fi33 := 0; fi33 < 3; fi33++ {
        v2 -= (0.1 + float32(-2.25))
        _ = fi33
    }
    v3 = (^(v3))
    v34 := []int{v1}
    switch (v1 << ((4294967295) & 63)) {
    case 4294967295:
        v35 := GenS4{F0: 4294967295}
        _ = v35
    case 128:
        switch "a" {
        case "\n\t":
            fmt.Println(v2, v2)
        case "a":
            v1 = 0
        case "":
            v36 := GenS5{F0: 434}
            _ = v36
        }
    case -307:
        fmt.Println(v3)
    }
    for fi37 := 0; fi37 < 17; fi37++ {
        if ("hello" == "") {
            for fi38 := 0; fi38 < 6; fi38++ {
                if false {
                    fmt.Println(fi37)
                } else {
                    var v39 int8 = 100
                    v1 -= fi38
                    fmt.Println(fi38, v3)
                    _ = v39
                }
                _ = fi38
            }
        }
        if len(v34) > 0 {
            v34[0] = (v1 << ((255) & 63))
        }
        for fi40 := 0; fi40 < 5; fi40++ {
            if false {
                if (((127 <= 16) || (!(true || (((3.14159 - (0.0 - (0.0 / (1e10 / 0.1)))) * 0.0) > -2.25)))) || (!((^(v3)) <= (v3 - (v3 | v3))))) {
                    v41 := GenS2{F0: ((v2 + v2) > (v2 / 1e10)), F1: -1}
                    v3 = 2147483648
                    _ = v41
                } else {
                    v42 := GenS1{F0: (fi40 | fi37), F1: 207}
                    v34 = append(v34, 42)
                    if ((v3 + 1) > v3) {
                        if true {
                            var v43 bool = (((false || true) && false) || (65535 > 2))
                            var v44 int16 = 661
                            _ = v43
                            _ = v44
                        }
                        v34 = append(v34, fi37)
                    }
                    _ = v42
                }
            }
            var v45 int16 = genFn0(255)
            switch ("a" + "hello") {
            case "\n\t":
                v1 += v1
            }
            v46 := GenS2{F0: true, F1: 42}
            if true {
                v2 = 0.0
                var v47 int16 = v45
                _ = v47
            }
            v48 := GenS6{F0: 2147483647, F1: (fi37 - fi40), F2: 4294967295}
            _ = fi40
            _ = v45
            _ = v46
            _ = v48
        }
        _ = fi37
    }
    switch ("a" + ("foo bar" + ("foo bar" + ""))) {
    case "\n\t":
        fmt.Println(v3, v2, v3)
    case "":
        v49 := GenS5{F0: (v3 ^ (v3 / (((v3 % ((v3) | 1))) | 1)))}
        _ = v49
    case "foo bar":
        if (2147483647 != (v3 * 4294967295)) {
            _ = len(v34)
            v50 := GenS2{F0: (false && (-90 > 42)), F1: 42}
            _ = v50
        } else {
            v1 = (v1 - 42)
        }
    default:
        var v51 int8 = -127
        _ = v51
    }
    switch (v1 | v1) {
    case 0:
        for fi52 := 0; fi52 < 16; fi52++ {
            for fi53 := 0; fi53 < 11; fi53++ {
                var v54 int8 = 42
                v34 = append(v34, 32768)
                _ = fi53
                _ = v54
            }
            _ = fi52
        }
    }
    v55 := GenS7{F0: 127, F1: 0, F2: 42}
    if (2147483647 < 0) {
        v2 = 3.14159
        switch 189 {
        case -2147483648:
            if len(v34) >= 2 {
                v34 = v34[1:len(v34)]
            }
        case 127:
            v1 += (-(v1))
        case 2147483647:
            if ((v3 * v3) <= v3) {
                v3 -= (v3 ^ v3)
            } else {
                v2 = 3.14159
                v56 := GenS8{F0: 106, F1: 2147483647, F2: 251}
                fmt.Println(v2, v1, v1)
                _ = v56
            }
        default:
            v2 = v2
        }
        if len(v34) >= 2 {
            v34 = v34[1:len(v34)]
        }
    }
    var v57 float32 = ((((0.1 + v2) / v2) + float32((0.0 / -2.25))) / (v2 + v2))
    for fi58 := 0; fi58 < 15; fi58++ {
        v59 := GenS8{F0: uint64(v3), F1: 0, F2: -1}
        _ = fi58
        _ = v59
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v34
    _ = v55
    _ = v57
}
