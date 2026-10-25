package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 uint16; F1 uint16; F2 uint16 }
type GenS4 struct { F0 int8; F1 uint32 }
type GenS5 struct { F0 uint64; F1 int8 }
type GenS6 struct { F0 bool; F1 int32; F2 uint }
type GenS7 struct { F0 uint8; F1 int8; F2 uint32 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 string; F1 int64 }
type GenS10 struct { F0 int8; F1 uint32; F2 uint16 }
type GenS11 struct { F0 uint; F1 int64; F2 string }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 bool; F1 uint16; F2 int64 }
func genFn0(p0 int64) int8 {
    switch "a" {
    case "hello":
        p0 += -9223372036854775808
    case "\n\t":
        var v4 int32 = -2147483648
        _ = v4
    case "foo bar":
        p0 = (p0 ^ (p0 / (((p0 << ((18446744073709551615) & 63))) | 1)))
    default:
        fmt.Println(p0)
    }
    for fi5 := 0; fi5 < 16; fi5++ {
        var v6 string = ("foo bar" + ("\n\t" + ""))
        _ = fi5
        _ = v6
    }
    _ = p0
    return -1
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]bool)
    var v3 uint32 = 2147483648
    v7 := []int{32767, 42, -2147483648, -1}
    fmt.Println(v3)
    for fi8 := 0; fi8 < 4; fi8++ {
        fmt.Println(v3, fi8)
        fmt.Println(v3)
        _ = fi8
    }
    if false {
        var v9 float64 = 3.14159
        for fi10 := 0; fi10 < 12; fi10++ {
            v11 := GenS1{F0: 128}
            _ = fi10
            _ = v11
        }
        _ = v9
    } else {
        v3 -= v3
        fmt.Println(v3)
    }
    if (487 <= 0) {
        v3 = v3
        for fi12 := 0; fi12 < 2; fi12++ {
            var v13 int64 = -1
            v14 := GenS2{F0: -32768, F1: (!(!((56 != 255) && false)))}
            if (132 >= 0) {
                var v15 bool = ((-(fi12)) < fi12)
                var v16 bool = v15
                switch "a" {
                case "":
                    for fi17 := 0; fi17 < 7; fi17++ {
                        v7 = append(v7, -465)
                        v18 := GenS2{F0: 100, F1: (!((1.5 > float32(-2.25)) && true))}
                        _ = fi17
                        _ = v18
                    }
                }
                _ = v15
                _ = v16
            }
            _ = fi12
            _ = v13
            _ = v14
        }
        delete(v2, "hello")
    }
    for fi19 := 0; fi19 < 18; fi19++ {
        if true {
            for fi20 := 0; fi20 < 9; fi20++ {
                var v21 bool = false
                fmt.Println(v3, v3)
                var v22 uint = 562
                _ = fi20
                _ = v21
                _ = v22
            }
            if (2 <= 18446744073709551615) {
                if true {
                    v3 = v3
                    fmt.Println(v3)
                    var v23 uint32 = (v3 - 0)
                    _ = v23
                } else {
                    v7 = append(v7, (^(fi19)))
                    var v24 int16 = 32767
                    _ = v24
                }
                v3 += v3
                switch "\n\t" {
                case "hello":
                    if true {
                        if (!false) {
                            v25 := GenS3{F0: 42, F1: 1, F2: 42}
                            _ = v25
                        }
                        for fi26 := 0; fi26 < 9; fi26++ {
                            v3 = (v3 << ((509) & 63))
                            fmt.Println(fi26)
                            _ = fi26
                        }
                    } else {
                        v27 := GenS2{F0: 32767, F1: (2 < 2)}
                        if (!true) {
                            if true {
                                v28 := GenS4{F0: -1, F1: (v3 / (((v3 ^ 1)) | 1))}
                                for fi29 := 0; fi29 < 10; fi29++ {
                                    switch 128 {
                                    case -781:
                                        switch -1 {
                                        case 935:
                                            var v30 uint64 = 678
                                            _ = v30
                                        case 941:
                                            v3 = (v3 ^ v3)
                                        default:
                                            v3 -= (v3 | (v3 + v3))
                                        }
                                    }
                                    fmt.Println(fi19, fi29, fi29)
                                    _ = fi29
                                }
                                _ = v28
                            }
                        } else {
                            v31 := GenS5{F0: 42, F1: 0}
                            fmt.Println(fi19, v3)
                            v32 := GenS6{F0: ((!(-128 >= genFn0(-9223372036854775808))) || true), F1: -155, F2: 957}
                            _ = v31
                            _ = v32
                        }
                        var v33 int8 = 124
                        _ = v27
                        _ = v33
                    }
                default:
                    var v34 int8 = genFn0(int64(v3))
                    _ = v34
                }
            }
            _ = len(v1)
        }
        var v35 int16 = 42
        v35 -= (v35 & v35)
        _ = fi19
        _ = v35
    }
    if (genFn0(42) > -104) {
        switch 165 {
        case -208:
            if true {
                var v36 uint32 = v3
                v37 := GenS7{F0: 255, F1: 31, F2: (v36 % (((v36 * (v3 ^ (v3 | v3)))) | 1))}
                _ = v36
                _ = v37
            }
        case 212:
            for fi38 := 0; fi38 < 10; fi38++ {
                v39 := GenS6{F0: (uint64(v3) == 4294967295), F1: -1, F2: 42}
                v3 -= (v3 + (v3 | (v3 << ((981) & 63))))
                switch (fi38 >> ((4294967295) & 63)) {
                case 0:
                    v3 = 4294967295
                }
                _ = fi38
                _ = v39
            }
        default:
            v3 -= 1
        }
        var v40 string = "foo bar"
        v40 = ""
        _ = v40
    } else {
        for fi41 := 0; fi41 < 2; fi41++ {
            v3 += v3
            var v42 int32 = 42
            for fi43 := 0; fi43 < 20; fi43++ {
                v42 = 0
                fmt.Println(fi41)
                var v44 string = "a"
                _ = fi43
                _ = v44
            }
            _ = fi41
            _ = v42
        }
    }
    for fi45 := 0; fi45 < 17; fi45++ {
        fmt.Println(v3, fi45)
        _ = fi45
    }
    switch 128 {
    case -2147483648:
        for fi46 := 0; fi46 < 12; fi46++ {
            if false {
                var v47 uint8 = 128
                for fi48 := 0; fi48 < 1; fi48++ {
                    if (2147483647 == 2147483647) {
                        if (416 != 957) {
                            var v49 string = "a"
                            v49 = "foo bar"
                            var v50 uint = 18446744073709551615
                            v51 := GenS8{F0: 56}
                            if (1e10 != 1.5) {
                                v52 := GenS9{F0: (v49 + v49), F1: 0}
                                _ = v52
                            } else {
                                var v53 float64 = (1.5 / 1e10)
                                var v54 uint16 = 32767
                                if false {
                                    v54 = 32767
                                    switch (fi46 ^ 428) {
                                    case 0:
                                        var v55 uint32 = (v3 / ((470) | 1))
                                        _ = v55
                                    case 127:
                                        if (true && (-128 > -1)) {
                                            v2[("hello" + ((((v49 + v49) + v49) + v49) + (v49 + "hello")))] = (float32((v53 * v53)) == 3.14159)
                                        } else {
                                            _ = len(v1)
                                            switch "\n\t" {
                                            case "a":
                                                var v56 uint8 = 127
                                                _ = v56
                                            case "\n\t":
                                                fmt.Println(v3, v49, v3)
                                            case "hello":
                                                v50 = 0
                                            }
                                        }
                                    case 128:
                                        v57 := GenS10{F0: 1, F1: v3, F2: (v54 & v54)}
                                        _ = v57
                                    }
                                    v58 := GenS11{F0: v50, F1: int64(v54), F2: v49}
                                    var v59 uint32 = 2
                                    v60 := GenS3{F0: v54, F1: v54, F2: (v54 / (((^(v54))) | 1))}
                                    var v61 int = fi48
                                    if (46 > 0) {
                                        v61 = (fi46 | (fi48 >> ((v50) & 63)))
                                        var v62 bool = true
                                        _ = v62
                                    }
                                    v59 = (^(v59))
                                    v63 := v1[(v3 << (((v50 & (v50 & (v50 + v50)))) & 63))]
                                    _ = v58
                                    _ = v59
                                    _ = v60
                                    _ = v61
                                    _ = v63
                                } else {
                                    if (!false) {
                                        fmt.Println(v49)
                                        var v64 int8 = 0
                                        v47 += v47
                                        _ = v64
                                    } else {
                                        if true {
                                            fmt.Println(fi46)
                                            var v65 int32 = -959
                                            v2["\n\t"] = (v49 > "\n\t")
                                            _ = v65
                                        } else {
                                            fmt.Println(v54)
                                            v47 = 229
                                            fmt.Println(fi46)
                                        }
                                        v66 := v2[v49]
                                        _ = v66
                                    }
                                    v54 += (^(v54))
                                    v67 := GenS12{F0: v3}
                                    _ = v67
                                }
                                _ = v53
                                _ = v54
                            }
                            _ = v49
                            _ = v50
                            _ = v51
                        }
                        fmt.Println(fi46)
                        v47 = (v47 & (v47 ^ v47))
                    }
                    v68 := GenS1{F0: fi48}
                    _ = fi48
                    _ = v68
                }
                var v69 uint32 = (v3 % ((2147483648) | 1))
                _ = v47
                _ = v69
            }
            fmt.Println(v3)
            switch fi46 {
            case 0:
                fmt.Println(v3, v3)
            case -9223372036854775808:
                fmt.Println(fi46, v3)
            case 2147483647:
                v70 := GenS3{F0: 42, F1: 32767, F2: 836}
                _ = v70
            default:
                var v71 uint = 42
                _ = v71
            }
            _ = fi46
        }
    case -1:
        v72 := GenS5{F0: 0, F1: 127}
        _ = v72
    case -9223372036854775808:
        v73 := GenS5{F0: 932, F1: 13}
        _ = v73
    default:
        v74 := GenS10{F0: 100, F1: (v3 * (v3 << ((771) & 63))), F2: 2}
        _ = v74
    }
    switch 32768 {
    case -1:
        var v75 int16 = -118
        _ = v75
    default:
        v76 := GenS13{F0: false, F1: 32767, F2: 1}
        _ = v76
    }
    v77 := GenS11{F0: 1, F1: 9223372036854775807, F2: "hello"}
    var v78 int = -9223372036854775808
    var v79 uint32 = 47
    var v80 float64 = 0.0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v78:", v78)
    fmt.Println("v79:", v79)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v2
    _ = v3
    _ = v7
    _ = v77
    _ = v78
    _ = v79
    _ = v80
}
