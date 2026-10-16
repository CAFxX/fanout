package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int16; F1 uint64; F2 uint16 }
type GenS5 struct { F0 uint; F1 bool; F2 uint32 }
type GenS6 struct { F0 bool; F1 int64 }
type GenS7 struct { F0 uint; F1 bool }
func genFn0(p0 int32, p1 int, p2 int) uint64 {
    v7 := GenS1{F0: 184}
    if (!false) {
        var v8 int32 = (p0 - p0)
        v8 = p0
        _ = v8
    } else {
        fmt.Println(p0)
    }
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 4294967295
}
func genFn1() uint {
    switch -973 {
    case -9223372036854775808:
        var v9 bool = false
        _ = v9
    case 1:
        var v10 uint8 = 187
        _ = v10
    default:
        var v11 uint16 = 895
        _ = v11
    }
    v12 := []int{128, 2147483647, 32767}
    if false {
        var v13 uint = 249
        v14 := GenS1{F0: 32767}
        if (1 != -9223372036854775808) {
            var v15 bool = (("hello" + ("foo bar" + "a")) > "a")
            v13 = 0
            _ = v15
        } else {
            var v16 int32 = 42
            v17 := GenS1{F0: 42}
            _ = v16
            _ = v17
        }
        _ = v13
        _ = v14
    } else {
        var v18 string = "\n\t"
        v18 = v18
        for fi19 := 0; fi19 < 15; fi19++ {
            if (-815 > (fi19 - 32768)) {
                v20 := GenS1{F0: -32768}
                v18 = (v18 + "hello")
                _ = v20
            } else {
                if (0.1 >= (((float32(1e10) * float32((1.5 - ((3.14159 - 0.1) - 1.5)))) - 1e10) + float32(3.14159))) {
                    v18 = ((("foo bar" + v18) + v18) + ((v18 + "") + v18))
                    var v21 float32 = ((0.1 - 0.1) * (0.0 - 3.14159))
                    v22 := GenS2{F0: -1}
                    _ = v21
                    _ = v22
                } else {
                    switch fi19 {
                    case 32768:
                        switch (v18 + v18) {
                        case "":
                            switch ("" + v18) {
                            case "\n\t":
                                fmt.Println(fi19, v18)
                            }
                        case "\n\t":
                            switch "hello" {
                            case "\n\t":
                                var v23 string = "foo bar"
                                _ = v23
                            }
                        }
                    case 1:
                        v18 = ("" + v18)
                    case -926:
                        v18 = v18
                    }
                    var v24 bool = (!(!false))
                    fmt.Println(v24)
                    _ = v24
                }
            }
            _ = fi19
        }
        _ = v18
    }
    _ = v12
    return 704
}

func main() {
    v1 := []string{("foo bar" + ""), "hello"}
    v2 := make(map[int]uint32)
    var v3 int16 = -32768
    v4 := []int{127, -160}
    v5 := make(map[uint32]uint32)
    var v6 uint8 = 255
    if false {
        v25 := v2[-44]
        var v26 bool = (222 == v6)
        _ = v25
        _ = v26
    } else {
        fmt.Println(v3, v6)
    }
    var v27 uint16 = 2
    v4 = append(v4, -724)
    v3 -= v3
    fmt.Println(v6, v3)
    var v28 uint8 = v6
    if false {
        v27 += v27
    }
    var v29 uint32 = 2147483648
    v30 := GenS2{F0: (v3 << ((708) & 63))}
    v3 += (v3 | (v3 % ((v3) | 1)))
    for fi31 := 0; fi31 < 13; fi31++ {
        for fi32 := 0; fi32 < 3; fi32++ {
            if ((true || (v27 <= (v27 << ((genFn1()) & 63)))) && false) {
                v6 = v28
                if (!(false && (!true))) {
                    if (194 == -1) {
                        _ = len(v4)
                    } else {
                        v6 -= (v28 + v6)
                    }
                }
                v1 = append(v1, "a")
            }
            switch ("foo bar" + "") {
            case "a":
                fmt.Println(v3, fi32, v28)
            case "":
                var v33 int32 = 0
                _ = v33
            }
            switch fi32 {
            case -981:
                v34 := GenS3{F0: 4294967295}
                _ = v34
            case 4294967295:
                if true {
                    v27 -= (^(v27))
                    if (false && false) {
                        fmt.Println(v29, v28)
                        if false {
                            var v35 int16 = (-(v3))
                            if ((!true) && (uint(v6) > 255)) {
                                switch (fi31 & fi32) {
                                case 127:
                                    if (2147483647 >= -354) {
                                        v29 = v29
                                        v3 += v3
                                        var v36 int32 = -608
                                        _ = v36
                                    } else {
                                        var v37 int32 = -2147483648
                                        switch fi31 {
                                        case 127:
                                            var v38 float32 = 0.1
                                            _ = v38
                                        case 2147483647:
                                            v28 += 128
                                        }
                                        fmt.Println(fi31, v29, v35)
                                        _ = v37
                                    }
                                    v28 = v28
                                case -155:
                                    var v39 uint16 = (v27 ^ (v27 % (((v27 / (((^(v27))) | 1))) | 1)))
                                    _ = v39
                                case 42:
                                    v40 := GenS2{F0: v3}
                                    _ = v40
                                default:
                                    v29 = (v29 * v29)
                                }
                                var v41 int32 = 128
                                v42 := GenS3{F0: 0}
                                fmt.Println(v28)
                                v43 := v5[(v29 / ((65535) | 1))]
                                _ = v41
                                _ = v42
                                _ = v43
                            }
                            _ = v35
                        }
                    } else {
                        v27 += (^(v27))
                        switch "hello" {
                        case "\n\t":
                            if false {
                                v44 := GenS2{F0: 1000}
                                if true {
                                    var v45 float32 = (-2.25 - 3.14159)
                                    v46 := GenS4{F0: v3, F1: 25, F2: uint16(v28)}
                                    if (-543 != -578) {
                                        v27 += 255
                                        var v47 int64 = 9223372036854775807
                                        _ = v47
                                    }
                                    _ = v45
                                    _ = v46
                                } else {
                                    v27 += v27
                                    if (!(100 != 1)) {
                                        switch "\n\t" {
                                        case "hello":
                                            if (!true) {
                                                if false {
                                                    var v48 int64 = 9223372036854775807
                                                    _ = v48
                                                }
                                            } else {
                                                v27 += v27
                                                if false {
                                                    var v49 uint32 = (v29 & v29)
                                                    _ = v49
                                                }
                                            }
                                        case "\n\t":
                                            v50 := GenS1{F0: v3}
                                            _ = v50
                                        default:
                                            var v51 float32 = ((1.5 - 0.1) + 3.14159)
                                            _ = v51
                                        }
                                        v4 = append(v4, fi32)
                                        var v52 string = ""
                                        _ = v52
                                    } else {
                                        if (uint32(v27) != v29) {
                                            var v53 float64 = (3.14159 / (0.1 + 0.0))
                                            if len(v4) > 0 {
                                                v4[0] = (fi31 >> ((0) & 63))
                                            }
                                            if (!(false || false)) {
                                                v27 = (v27 | (^(v27)))
                                                v27 -= v27
                                                v28 -= 246
                                            } else {
                                                if len(v4) > 0 {
                                                    v4[0] = (fi32 & (-(fi31)))
                                                }
                                                var v54 uint64 = 4294967295
                                                var v55 uint8 = v6
                                                _ = v54
                                                _ = v55
                                            }
                                            _ = v53
                                        } else {
                                            var v56 bool = (("a" != "foo bar") && (1.5 == 1.5))
                                            _ = v56
                                        }
                                    }
                                }
                                v29 = 1
                                _ = v44
                            } else {
                                v27 = 0
                            }
                        case "foo bar":
                            v2[fi32] = 42
                        default:
                            v4 = append(v4, int(v29))
                        }
                        if ("" < "hello") {
                            v57 := GenS5{F0: 1, F1: true, F2: v29}
                            var v58 int32 = 0
                            var v59 string = "hello"
                            _ = v57
                            _ = v58
                            _ = v59
                        }
                    }
                    v3 = v3
                }
            }
            v4 = append(v4, (fi32 ^ fi31))
            _ = fi32
        }
        switch 4294967295 {
        case 0:
            var v60 int64 = int64(fi31)
            _ = v60
        default:
            v1 = append(v1, (("hello" + "foo bar") + ("a" + "foo bar")))
        }
        var v61 bool = false
        _ = fi31
        _ = v61
    }
    if true {
        if false {
            var v62 float64 = ((0.1 * (3.14159 * 1.5)) * ((((1e10 - (0.0 / 1.0)) * 0.0) + 3.14159) - (-2.25 / 1.5)))
            _ = v62
        }
        switch "\n\t" {
        case "a":
            v27 += v27
        case "":
            for fi63 := 0; fi63 < 10; fi63++ {
                if true {
                    var v64 uint = 600
                    var v65 uint32 = 65535
                    _ = v64
                    _ = v65
                } else {
                    var v66 uint32 = v29
                    _ = v66
                }
                v67 := GenS6{F0: (false && (v3 == -32768)), F1: -1}
                if (381 != (v3 * (-(v3)))) {
                    v6 = (v6 - (v28 / ((v6) | 1)))
                    if true {
                        var v68 uint8 = 42
                        v6 -= (v68 | v6)
                        _ = v68
                    } else {
                        var v69 uint = 264
                        v70 := GenS7{F0: (v69 % ((v69) | 1)), F1: false}
                        _ = len(v4)
                        _ = v69
                        _ = v70
                    }
                }
                _ = fi63
                _ = v67
            }
        }
    } else {
        v71 := GenS7{F0: 65535, F1: ((v6 >> ((4294967295) & 63)) >= 0)}
        fmt.Println(v6, v6, v3)
        var v72 int32 = 2147483647
        _ = v71
        _ = v72
    }
    v3 -= (v3 & (v3 ^ v3))
    delete(v5, v29)
    var v73 float64 = float64(1.5)
    var v74 int64 = -9223372036854775808
    if true {
        v75 := GenS3{F0: (v29 & v29)}
        for fi76 := 0; fi76 < 18; fi76++ {
            v28 += (v28 << ((998) & 63))
            switch fi76 {
            case -2147483648:
                v28 = v6
            case 329:
                for fi77 := 0; fi77 < 20; fi77++ {
                    var v78 bool = false
                    _ = fi77
                    _ = v78
                }
            case 0:
                var v79 bool = true
                _ = v79
            default:
                v3 += v3
            }
            delete(v2, fi76)
            _ = fi76
        }
        _ = v75
    }
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v73:", v73)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v73
    _ = v74
}
