package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 int8 }
type GenS2 struct { F0 uint; F1 bool; F2 int64 }
type GenS3 struct { F0 uint16; F1 uint }
type GenS4 struct { F0 int; F1 uint }
type GenS5 struct { F0 bool; F1 string }
type GenS6 struct { F0 uint32; F1 uint32; F2 uint64 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint8; F1 int64; F2 int64 }
func genFn0(p0 bool, p1 int16, p2 int) uint16 {
    for fi6 := 0; fi6 < 6; fi6++ {
        v7 := GenS1{F0: (("" + "\n\t") + ""), F1: 149, F2: 42}
        for fi8 := 0; fi8 < 1; fi8++ {
            p0 = true
            _ = fi8
        }
        p0 = p0
        _ = fi6
        _ = v7
    }
    fmt.Println(p0, p1)
    v9 := []int{}
    var v10 uint8 = 218
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 2
}

func main() {
    v1 := make(map[string]bool)
    var v2 int64 = 42
    var v3 uint32 = 2
    var v4 uint64 = 1
    v5 := make(map[int]uint32)
    for fi11 := 0; fi11 < 17; fi11++ {
        if (v3 > 2147483648) {
            var v12 float64 = 0.0
            if ((v4 | v4) != 1) {
                var v13 bool = false
                for fi14 := 0; fi14 < 20; fi14++ {
                    switch (fi14 - fi11) {
                    case 32767:
                        _ = len(v5)
                    }
                    v15 := []int{fi14, 4294967295, fi14, (fi14 | fi11)}
                    _ = fi14
                    _ = v15
                }
                var v16 int16 = int16(v3)
                _ = v13
                _ = v16
            } else {
                switch (fi11 / ((fi11) | 1)) {
                case 42:
                    for fi17 := 0; fi17 < 12; fi17++ {
                        _ = len(v5)
                        if (false && (false && false)) {
                            v12 = float64(3.14159)
                            v2 += int64(fi11)
                            var v18 int16 = 1
                            v4 = (v4 % ((v4) | 1))
                            var v19 bool = ((fi17 - (fi17 & fi11)) <= (fi17 / ((fi11) | 1)))
                            v20 := []int{(^(fi11)), fi17}
                            _ = v18
                            _ = v19
                            _ = v20
                        } else {
                            if (108 != 0) {
                                v21 := GenS1{F0: "a", F1: (v2 << ((18446744073709551615) & 63)), F2: -127}
                                var v22 string = (("hello" + "hello") + "\n\t")
                                _ = v21
                                _ = v22
                            } else {
                                v4 = 2
                                fmt.Println(fi11, fi17, fi17)
                                if (fi11 < fi11) {
                                    var v23 uint32 = 42
                                    v24 := GenS2{F0: 2, F1: (false || true), F2: (v2 & v2)}
                                    _ = v23
                                    _ = v24
                                }
                            }
                            v5[127] = v3
                        }
                        _ = fi17
                    }
                default:
                    switch "foo bar" {
                    case "":
                        v12 += 3.14159
                    case "\n\t":
                        v25 := GenS2{F0: 42, F1: ((221 != 0) && (!((2 < 928) && (-1 >= -270)))), F2: (v2 / (((v2 + 42)) | 1))}
                        _ = v25
                    }
                }
                if (v2 != v2) {
                    v5[fi11] = 42
                    for fi26 := 0; fi26 < 8; fi26++ {
                        v2 -= v2
                        if false {
                            v3 -= v3
                            fmt.Println(v4)
                            v3 = 81
                        } else {
                            if false {
                                v27 := []int{fi11, (fi26 * (fi26 + (fi26 % ((-100) | 1)))), fi11}
                                v28 := v5[fi26]
                                _ = v27
                                _ = v28
                            }
                            delete(v5, (fi26 * fi26))
                        }
                        if false {
                            var v29 int8 = 87
                            _ = v29
                        }
                        var v30 int64 = (v2 % ((9223372036854775807) | 1))
                        _ = fi26
                        _ = v30
                    }
                } else {
                    var v31 uint16 = 32767
                    if ((75 < 42) || false) {
                        if ((v12 / ((-2.25 - v12) + (((1.5 / (0.1 / (v12 * float64(3.14159)))) + 0.0) + v12))) > (((v12 + v12) / (float64(1e10) - (v12 * -2.25))) + (v12 - float64(0.1)))) {
                            fmt.Println(v2, v31)
                            for fi32 := 0; fi32 < 12; fi32++ {
                                _ = len(v1)
                                _ = fi32
                            }
                            var v33 int = 802
                            _ = v33
                        }
                        v3 += (^(v3))
                        fmt.Println(v12, v12)
                    }
                    v31 -= (v31 - genFn0(((true || false) || true), int16(v31), 128))
                    _ = v31
                }
            }
            if ((1 > 255) || ("hello" != ("" + ("hello" + "\n\t")))) {
                var v34 uint8 = 1
                _ = v34
            } else {
                for fi35 := 0; fi35 < 8; fi35++ {
                    switch (fi11 | (fi11 / ((fi11) | 1))) {
                    case 2147483647:
                        v4 -= v4
                    case -2147483648:
                        var v36 uint = 42
                        _ = v36
                    }
                    _ = fi35
                }
                fmt.Println(v3, fi11, v4)
                var v37 float64 = (float64(((3.14159 + 3.14159) + 0.0)) * v12)
                _ = v37
            }
            _ = v12
        } else {
            v38 := []int{-1}
            _ = v38
        }
        if (!(v2 != (v2 | (-(v2))))) {
            if false {
                v2 -= v2
                v3 = (v3 | (v3 & 42))
            }
        }
        switch fi11 {
        case 1:
            v39 := []int{fi11, fi11, (fi11 << ((30) & 63))}
            _ = v39
        case 560:
            fmt.Println(fi11, v3)
        }
        _ = fi11
    }
    v40 := v5[9223372036854775807]
    if (!((((172 > 83) || (true && false)) && (("foo bar" > "foo bar") && false)) && true)) {
        for fi41 := 0; fi41 < 5; fi41++ {
            if true {
                for fi42 := 0; fi42 < 1; fi42++ {
                    v2 -= v2
                    v4 = v4
                    fmt.Println(v4, fi41)
                    _ = fi42
                }
                fmt.Println(v2, v3)
                if false {
                    v43 := GenS3{F0: genFn0((false && (2147483647 == int32(fi41))), -554, (^(fi41))), F1: 279}
                    _ = v43
                }
            } else {
                v2 = (-(v2))
                v44 := GenS4{F0: 417, F1: 18446744073709551615}
                _ = v44
            }
            v2 -= 826
            v40 = 0
            _ = fi41
        }
        v40 = 0
        var v45 uint64 = v4
        _ = v45
    }
    v2 = (^(v2))
    var v46 uint16 = 32767
    v40 = 65535
    switch -476 {
    case 959:
        var v47 string = "foo bar"
        _ = v47
    case -335:
        var v48 int32 = 209
        _ = v48
    default:
        v46 -= (v46 / ((genFn0(true, -32768, -1)) | 1))
    }
    var v49 float32 = ((1.5 + (0.0 / ((1e10 + 0.0) + 0.0))) + -2.25)
    v3 += (^(v3))
    switch int(v40) {
    case -2147483648:
        if true {
            switch -1 {
            case 127:
                v4 = 18446744073709551615
            }
        }
    }
    var v50 uint32 = 473
    v46 = (v46 | (v46 >> ((18446744073709551615) & 63)))
    v51 := GenS5{F0: false, F1: "a"}
    if (96 > 127) {
        if (true || ((!false) || true)) {
            v52 := GenS6{F0: v50, F1: (v50 - (v40 >> ((42) & 63))), F2: v4}
            v2 = v2
            switch "a" {
            case "\n\t":
                var v53 int8 = -1
                _ = v53
            case "foo bar":
                if false {
                    var v54 bool = (100 >= -36)
                    if v54 {
                        v55 := GenS7{F0: v2}
                        if (v46 != (v46 % (((v46 % ((v46) | 1))) | 1))) {
                            var v56 uint = 18446744073709551615
                            _ = v56
                        } else {
                            if (0 != uint8(v4)) {
                                if (v46 != v46) {
                                    var v57 float32 = (v49 / v49)
                                    _ = v57
                                } else {
                                    v58 := v1["foo bar"]
                                    v40 = (v3 + (v50 & (v40 >> ((0) & 63))))
                                    var v59 string = ("\n\t" + "hello")
                                    _ = v58
                                    _ = v59
                                }
                                v60 := v5[-341]
                                _ = v60
                            } else {
                                v61 := GenS7{F0: (-(v2))}
                                v62 := v5[-741]
                                if v54 {
                                    if true {
                                        v4 = v4
                                        fmt.Println(v3, v4)
                                    }
                                    if true {
                                        v63 := GenS8{F0: 128, F1: v2, F2: v2}
                                        _ = len(v1)
                                        var v64 float64 = 1.5
                                        _ = v63
                                        _ = v64
                                    } else {
                                        v65 := []int{32767, -9223372036854775808, 32767}
                                        _ = len(v5)
                                        _ = v65
                                    }
                                }
                                _ = v61
                                _ = v62
                            }
                            v3 += (v50 / ((v3) | 1))
                        }
                        v46 = 32767
                        _ = v55
                    } else {
                        v54 = ((false && ((1.5 + ((0.1 + 1e10) / 1e10)) < 0.0)) && v54)
                        v50 = v40
                    }
                    var v66 int64 = 0
                    _ = v54
                    _ = v66
                } else {
                    fmt.Println(v40, v3)
                    if ((v46 >> ((255) & 63)) == v46) {
                        if (!true) {
                            var v67 int8 = 0
                            if (-1 != -563) {
                                var v68 int32 = -732
                                v46 = 1
                                switch -2147483648 {
                                case -2147483648:
                                    var v69 int32 = (v68 + (v68 / ((v68) | 1)))
                                    _ = v69
                                case 127:
                                    v70 := GenS2{F0: 188, F1: true, F2: (^(v2))}
                                    _ = v70
                                case 42:
                                    v67 = (v67 >> ((0) & 63))
                                default:
                                    var v71 bool = false
                                    _ = v71
                                }
                                _ = v68
                            }
                            v3 = (v50 & (v3 + v50))
                            _ = v67
                        } else {
                            var v72 bool = (((v49 + (0.0 - 0.1)) / float32((0.0 - 3.14159))) < float32(0.0))
                            switch 666 {
                            case -9223372036854775808:
                                v46 = (v46 | (v46 / ((v46) | 1)))
                            case -2147483648:
                                v73 := GenS7{F0: -39}
                                _ = v73
                            default:
                                v74 := []int{-9223372036854775808, -9223372036854775808, -794, -38}
                                _ = v74
                            }
                            _ = len(v1)
                            _ = v72
                        }
                        fmt.Println(v2, v3, v49)
                        if ((((float64((v49 - float32(3.14159))) - (0.1 - ((1e10 - 3.14159) + (3.14159 + 1.5)))) == 1e10) && true) && true) {
                            v75 := v1[("" + ("\n\t" + ""))]
                            v3 += (v50 ^ v40)
                            _ = v75
                        }
                    } else {
                        var v76 float64 = (0.0 + 0.0)
                        for fi77 := 0; fi77 < 12; fi77++ {
                            v3 -= (v40 << ((uint(v3)) & 63))
                            _ = fi77
                        }
                        fmt.Println(v50, v4, v76)
                        _ = v76
                    }
                }
            }
            _ = v52
        }
    } else {
        var v78 uint8 = 63
        _ = v78
    }
    v50 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v46:", v46)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v46
    _ = v49
    _ = v5
    _ = v50
    _ = v51
}
