package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16; F1 uint }
type GenS4 struct { F0 int; F1 int16 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int8; F1 int16; F2 uint16 }
type GenS8 struct { F0 uint16; F1 uint64 }
func genFn0(p0 uint) uint8 {
    p0 -= p0
    if (((float64(((float32(float64((3.14159 * 1.5))) * -2.25) + (3.14159 / 1e10))) + 0.1) + 0.1) < -2.25) {
        p0 = 827
        v4 := GenS1{F0: 100}
        var v5 int64 = 0
        _ = v4
        _ = v5
    } else {
        for fi6 := 0; fi6 < 6; fi6++ {
            p0 += (^(p0))
            if (202 <= 9) {
                v7 := GenS2{F0: 1000}
                v8 := []int{}
                _ = v7
                _ = v8
            } else {
                var v9 uint64 = 0
                _ = v9
            }
            for fi10 := 0; fi10 < 14; fi10++ {
                p0 -= (p0 << (((p0 % (((^(p0))) | 1))) & 63))
                p0 = p0
                if (true || (false || false)) {
                    if true {
                        p0 -= (p0 - 0)
                    } else {
                        var v11 bool = true
                        _ = v11
                    }
                    v12 := []int{fi10, (fi6 + (^(fi6))), -769}
                    _ = len(v12)
                    _ = v12
                }
                fmt.Println(fi6)
                _ = fi10
            }
            _ = fi6
        }
        if (!(9223372036854775807 == -570)) {
            v13 := []int{127, 907, -9223372036854775808}
            fmt.Println(p0)
            _ = v13
        }
    }
    _ = p0
    return 255
}
func genFn1(p0 uint, p1 bool, p2 uint32) int32 {
    var v14 int = 127
    for fi15 := 0; fi15 < 18; fi15++ {
        v16 := []int{491, fi15, 881}
        v17 := GenS2{F0: 42}
        if len(v16) > 0 {
            v16[0] = fi15
        }
        _ = fi15
        _ = v16
        _ = v17
    }
    v18 := GenS2{F0: -736}
    var v19 uint = (p0 << (((p0 + p0)) & 63))
    for fi20 := 0; fi20 < 12; fi20++ {
        v21 := []int{0, (fi20 >> ((255) & 63))}
        switch ("foo bar" + ("hello" + "a")) {
        case "":
            if true {
                var v22 int8 = 0
                v19 += p0
                _ = v22
            } else {
                if p1 {
                    var v23 bool = p1
                    _ = v23
                } else {
                    fmt.Println(p2, v19, fi20)
                    for fi24 := 0; fi24 < 2; fi24++ {
                        v25 := GenS3{F0: 0, F1: (^(p0))}
                        _ = fi24
                        _ = v25
                    }
                }
                for fi26 := 0; fi26 < 10; fi26++ {
                    switch (-(fi20)) {
                    case -802:
                        var v27 uint64 = 58
                        _ = v27
                    case 9223372036854775807:
                        if p1 {
                            if p1 {
                                v28 := GenS2{F0: 0}
                                _ = v28
                            }
                        } else {
                            fmt.Println(fi20, v14, v14)
                            var v29 uint8 = uint8(p2)
                            _ = v29
                        }
                        v21 = append(v21, -9223372036854775808)
                        v30 := GenS1{F0: -42}
                        p0 = p0
                        _ = v30
                    default:
                        v31 := GenS4{F0: (fi26 >> ((p0) & 63)), F1: 0}
                        _ = v31
                    }
                    _ = fi26
                }
                if (1 > -1) {
                    switch "hello" {
                    case "":
                        if ((-511 < 0) || true) {
                            var v32 int64 = 42
                            _ = v32
                        }
                    case "foo bar":
                        if len(v21) >= 2 {
                            v21 = v21[1:len(v21)]
                        }
                    case "a":
                        var v33 bool = p1
                        _ = v33
                    }
                    v34 := GenS2{F0: -1}
                    if (p1 && (false && p1)) {
                        switch "\n\t" {
                        case "\n\t":
                            var v35 uint8 = 42
                            _ = v35
                        case "foo bar":
                            _ = len(v21)
                        case "":
                            var v36 float32 = (((float32(1.5) + float32((float64(1e10) + 0.1))) + 1.5) * -2.25)
                            _ = v36
                        default:
                            var v37 float64 = 0.0
                            _ = v37
                        }
                        var v38 int = 4294967295
                        fmt.Println(p1)
                        _ = v38
                    } else {
                        var v39 int16 = -1
                        fmt.Println(p2, v19, v14)
                        _ = v39
                    }
                    _ = v34
                }
            }
        case "foo bar":
            var v40 int8 = 85
            _ = v40
        case "a":
            p2 = 42
        }
        v41 := GenS5{F0: 0}
        _ = fi20
        _ = v21
        _ = v41
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v18
    _ = v19
    return -107
}

func main() {
    var v1 string = (("a" + "") + "\n\t")
    var v2 uint64 = 996
    var v3 int16 = 0
    fmt.Println(v1, v1)
    v1 = (v1 + v1)
    if false {
        v42 := GenS6{F0: 0}
        var v43 int16 = v3
        _ = v42
        _ = v43
    } else {
        switch v1 {
        case "a":
            v2 += (v2 | (v2 | (v2 | (v2 >> ((509) & 63)))))
        case "hello":
            v44 := []int{-720, 127, 387, 222}
            _ = v44
        case "foo bar":
            v45 := []int{}
            _ = v45
        }
    }
    v2 = (^(v2))
    if false {
        v1 = (v1 + (v1 + "foo bar"))
    }
    if true {
        v46 := GenS7{F0: -73, F1: v3, F2: 33}
        for fi47 := 0; fi47 < 19; fi47++ {
            fmt.Println(v1, v3, v3)
            _ = fi47
        }
        _ = v46
    }
    for fi48 := 0; fi48 < 20; fi48++ {
        v49 := GenS4{F0: (fi48 % (((fi48 - fi48)) | 1)), F1: v3}
        var v50 float32 = 0.1
        _ = fi48
        _ = v49
        _ = v50
    }
    v51 := []int{42, -9223372036854775808, 127, -2147483648}
    for fi52 := 0; fi52 < 1; fi52++ {
        if false {
            var v53 uint = 4294967295
            _ = v53
        } else {
            var v54 uint64 = (v2 % ((v2) | 1))
            _ = v54
        }
        v55 := GenS5{F0: 562}
        for fi56 := 0; fi56 < 13; fi56++ {
            if false {
                if (127 <= 90) {
                    if len(v51) > 0 {
                        v51[0] = fi52
                    }
                    var v57 float64 = (1e10 * 0.1)
                    if len(v51) > 0 {
                        v51[0] = (fi56 ^ (fi52 / ((-2147483648) | 1)))
                    }
                    fmt.Println(v3)
                    _ = v57
                }
                v58 := GenS8{F0: 971, F1: (v2 + (v2 / (((v2 / ((v2) | 1))) | 1)))}
                _ = v58
            }
            v2 -= v2
            _ = fi56
        }
        _ = fi52
        _ = v55
    }
    v3 = v3
    switch "foo bar" {
    case "":
        switch (v1 + (((((v1 + v1) + "") + ((v1 + (v1 + v1)) + ("hello" + "hello"))) + "hello") + v1)) {
        case "hello":
            v59 := GenS7{F0: 1, F1: (v3 / (((^(v3))) | 1)), F2: 65535}
            _ = v59
        case "foo bar":
            fmt.Println(v2)
        default:
            for fi60 := 0; fi60 < 13; fi60++ {
                for fi61 := 0; fi61 < 15; fi61++ {
                    if false {
                        v62 := GenS8{F0: 65535, F1: 18446744073709551615}
                        if (430 < 9223372036854775807) {
                            v3 -= (-(v3))
                        }
                        v2 += 545
                        _ = v62
                    }
                    _ = fi61
                }
                fmt.Println(v1)
                fmt.Println(v2, v3)
                _ = fi60
            }
        }
    case "\n\t":
        fmt.Println(v3, v2)
    }
    v3 = v3
    v51 = append(v51, 799)
    v1 = ((v1 + (v1 + v1)) + v1)
    for fi63 := 0; fi63 < 16; fi63++ {
        _ = len(v51)
        _ = fi63
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v51:", len(v51), v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v51
}
