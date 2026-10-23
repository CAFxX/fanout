package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 bool }
type GenS2 struct { F0 int16; F1 uint64; F2 bool }
type GenS3 struct { F0 bool; F1 uint; F2 uint16 }
type GenS4 struct { F0 int32; F1 uint; F2 uint32 }
type GenS5 struct { F0 bool; F1 uint64 }
func genFn0() uint32 {
    var v5 int = 804
    v5 = (^(v5))
    _ = v5
    return 65535
}
func genFn1(p0 uint, p1 uint32) uint32 {
    fmt.Println(p1, p1)
    switch "" {
    case "foo bar":
        var v6 int64 = 9223372036854775807
        _ = v6
    default:
        var v7 int64 = -9223372036854775808
        _ = v7
    }
    var v8 string = ("a" + "a")
    switch 2147483647 {
    case 32768:
        v8 = v8
    case 42:
        for fi9 := 0; fi9 < 16; fi9++ {
            var v10 uint32 = 0
            p1 = (v10 | (p1 | v10))
            _ = fi9
            _ = v10
        }
    case 177:
        var v11 uint = p0
        _ = v11
    }
    switch v8 {
    case "hello":
        var v12 int16 = 1000
        _ = v12
    case "a":
        for fi13 := 0; fi13 < 14; fi13++ {
            var v14 float32 = 0.1
            _ = fi13
            _ = v14
        }
    case "foo bar":
        switch 9223372036854775807 {
        case 127:
            fmt.Println(v8)
        case -9223372036854775808:
            v15 := GenS1{F0: 720, F1: p1, F2: ((!true) && true)}
            _ = v15
        default:
            if false {
                v16 := []int{-1, 2147483647, 32768, -9223372036854775808}
                for fi17 := 0; fi17 < 9; fi17++ {
                    p1 = genFn0()
                    var v18 int = (fi17 | (fi17 >> ((p0) & 63)))
                    if len(v16) >= 2 {
                        v16 = v16[1:len(v16)]
                    }
                    _ = fi17
                    _ = v18
                }
                _ = v16
            }
        }
    default:
        fmt.Println(v8)
    }
    _ = p0
    _ = p1
    _ = v8
    return p1
}

func main() {
    var v1 int16 = -32768
    var v2 int32 = 2147483647
    v3 := []uint32{42, 645, 2, 2}
    var v4 float32 = (0.1 - ((1e10 * (1e10 - (3.14159 - 0.1))) - 1e10))
    v3 = append(v3, genFn1(4294967295, 2147483647))
    if false {
        var v19 uint16 = 65535
        v4 += v4
        _ = v19
    }
    if (!false) {
        var v20 int64 = -9223372036854775808
        _ = v20
    } else {
        v1 -= 1
        if true {
            v2 -= (^(v2))
            v4 = float32(((1e10 - 1.5) + 3.14159))
        } else {
            v21 := GenS2{F0: (v1 | v1), F1: 2, F2: (128 != 42)}
            v4 -= 3.14159
            _ = v21
        }
    }
    switch 4294967295 {
    case 0:
        _ = len(v3)
    case -9223372036854775808:
        var v22 string = "foo bar"
        _ = v22
    default:
        var v23 string = (((("\n\t" + "") + "a") + (("" + "\n\t") + "a")) + (("foo bar" + "") + "hello"))
        _ = v23
    }
    v2 = v2
    var v24 int16 = 1
    _ = len(v3)
    _ = len(v3)
    fmt.Println(v1)
    var v25 float64 = (3.14159 - -2.25)
    var v26 uint32 = 544
    var v27 uint16 = 789
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v28 int = 271
    if (!(896 > v28)) {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        switch ("a" + "a") {
        case "\n\t":
            switch (v28 << ((202) & 63)) {
            case 699:
                if len(v3) > 0 {
                    v3[0] = (v26 + (v26 << ((2) & 63)))
                }
            case 32768:
                if false {
                    for fi29 := 0; fi29 < 16; fi29++ {
                        v25 = (v25 + v25)
                        fmt.Println(v28)
                        for fi30 := 0; fi30 < 6; fi30++ {
                            var v31 int8 = -14
                            _ = fi30
                            _ = v31
                        }
                        _ = fi29
                    }
                    var v32 int8 = -54
                    _ = v32
                } else {
                    if (!false) {
                        var v33 uint64 = 18446744073709551615
                        v34 := GenS2{F0: 42, F1: v33, F2: (!true)}
                        _ = v33
                        _ = v34
                    }
                    if true {
                        switch "\n\t" {
                        case "":
                            for fi35 := 0; fi35 < 20; fi35++ {
                                v24 = v1
                                switch fi35 {
                                case 4294967295:
                                    v3 = append(v3, (v26 ^ 2))
                                case -1:
                                    if (!true) {
                                        var v36 uint8 = 42
                                        _ = v36
                                    }
                                case -832:
                                    switch 32767 {
                                    case 42:
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    case 2147483647:
                                        v2 += v2
                                    default:
                                        for fi37 := 0; fi37 < 13; fi37++ {
                                            v1 -= -540
                                            v3 = append(v3, genFn1(255, 693))
                                            _ = fi37
                                        }
                                    }
                                default:
                                    v24 = v24
                                }
                                _ = fi35
                            }
                        case "\n\t":
                            v3 = append(v3, genFn0())
                        default:
                            if (v27 == (v27 >> ((4294967295) & 63))) {
                                if ((v26 ^ v26) != 2) {
                                    fmt.Println(v1, v1, v27)
                                    fmt.Println(v24, v1, v26)
                                } else {
                                    v1 += (v1 ^ v24)
                                    v25 = (((v25 / 3.14159) + v25) - (v25 / -2.25))
                                    var v38 uint16 = v27
                                    _ = v38
                                }
                                if (!(false || (!(v27 < (v27 << ((36) & 63)))))) {
                                    fmt.Println(v4, v1, v25)
                                } else {
                                    if true {
                                        v1 += (v1 - (v1 * (v1 ^ int16(v28))))
                                        v39 := GenS1{F0: 18446744073709551615, F1: v26, F2: (42 <= 17)}
                                        _ = v39
                                    }
                                    v3 = append(v3, v26)
                                    _ = len(v3)
                                }
                                fmt.Println(v24, v2, v27)
                            } else {
                                for fi40 := 0; fi40 < 13; fi40++ {
                                    fmt.Println(fi40, v1)
                                    if true {
                                        v3 = append(v3, v26)
                                    } else {
                                        for fi41 := 0; fi41 < 16; fi41++ {
                                            if false {
                                                var v42 float32 = float32(v25)
                                                switch "\n\t" {
                                                case "\n\t":
                                                    if (false || false) {
                                                        var v43 int = int(v2)
                                                        v28 += (^(fi41))
                                                        var v44 float32 = (float32((v25 * (((1.5 - 1e10) - v25) * v25))) - v4)
                                                        _ = v43
                                                        _ = v44
                                                    }
                                                case "foo bar":
                                                    v45 := GenS3{F0: (844 == 132), F1: 768, F2: v27}
                                                    _ = v45
                                                case "hello":
                                                    var v46 int8 = 0
                                                    _ = v46
                                                }
                                                _ = len(v3)
                                                _ = v42
                                            }
                                            if (true && false) {
                                                v25 += v25
                                                v47 := GenS4{F0: v2, F1: 106, F2: v26}
                                                v2 -= v2
                                                _ = v47
                                            }
                                            _ = fi41
                                        }
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        v24 += v24
                                    }
                                    fmt.Println(v26)
                                    _ = fi40
                                }
                                _ = len(v3)
                            }
                        }
                        v48 := GenS5{F0: false, F1: 42}
                        _ = v48
                    }
                    var v49 string = ""
                    _ = v49
                }
            default:
                if (true || false) {
                    _ = len(v3)
                }
            }
        case "a":
            fmt.Println(v27)
        case "hello":
            v50 := GenS2{F0: (v24 + 1000), F1: 886, F2: true}
            _ = v50
        }
        switch v28 {
        case 4294967295:
            v51 := GenS5{F0: true, F1: 4294967295}
            _ = v51
        case 0:
            v25 = 0.1
        case -9223372036854775808:
            for fi52 := 0; fi52 < 9; fi52++ {
                fmt.Println(v26, v24, v25)
                _ = fi52
            }
        }
    }
    var v53 string = "hello"
    v27 += 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v53
}
