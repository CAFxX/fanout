package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 int16 }
type GenS3 struct { F0 int64; F1 uint64; F2 uint16 }
type GenS4 struct { F0 int32; F1 uint }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int; F1 uint }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint64; F1 int64 }
func genFn0(p0 bool, p1 int, p2 uint64) int8 {
    v5 := GenS1{F0: 42}
    for fi6 := 0; fi6 < 19; fi6++ {
        switch "a" {
        case "":
            for fi7 := 0; fi7 < 1; fi7++ {
                switch "hello" {
                case "foo bar":
                    v8 := []int{(-(p1)), (fi7 + (-(fi7))), p1, (p1 % ((p1) | 1))}
                    _ = v8
                }
                v9 := []int{p1, p1}
                _ = fi7
                _ = v9
            }
        }
        _ = fi6
    }
    p2 = 0
    v10 := []int{p1}
    fmt.Println(p1, p2, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v5
    return 100
}
func genFn1(p0 int16, p1 uint16, p2 int64) int8 {
    var v11 bool = ((false && ((p0 + -32768) != p0)) && false)
    v12 := GenS2{F0: (p0 >> ((0) & 63)), F1: (p0 & p0)}
    v11 = true
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return -1
}

func main() {
    v1 := []int{42, -2147483648, -197}
    var v2 uint16 = 0
    v3 := make(map[int]int)
    var v4 int64 = 42
    v2 += v2
    for fi13 := 0; fi13 < 5; fi13++ {
        var v14 float32 = 3.14159
        delete(v3, (fi13 ^ 127))
        _ = fi13
        _ = v14
    }
    var v15 uint32 = 404
    for fi16 := 0; fi16 < 7; fi16++ {
        if true {
            for fi17 := 0; fi17 < 16; fi17++ {
                if (!(100 >= -1)) {
                    _ = len(v3)
                    var v18 int = -474
                    _ = v18
                }
                v2 -= (v2 * 0)
                _ = fi17
            }
            for fi19 := 0; fi19 < 20; fi19++ {
                fmt.Println(fi16, fi16, fi19)
                var v20 int8 = -1
                delete(v3, fi16)
                _ = fi19
                _ = v20
            }
        } else {
            var v21 uint64 = 42
            var v22 int16 = 1000
            _ = v21
            _ = v22
        }
        if (127 <= 1) {
            v4 = (v4 % ((161) | 1))
            if false {
                fmt.Println(fi16, fi16)
                v23 := GenS3{F0: -9223372036854775808, F1: 271, F2: (v2 << ((150) & 63))}
                _ = v23
            }
        }
        _ = fi16
    }
    v24 := GenS2{F0: -32768, F1: int16(v4)}
    if false {
        switch 127 {
        case -1:
            for fi25 := 0; fi25 < 19; fi25++ {
                for fi26 := 0; fi26 < 19; fi26++ {
                    v4 = 468
                    _ = fi26
                }
                _ = fi25
            }
        case -2147483648:
            for fi27 := 0; fi27 < 3; fi27++ {
                v3[(fi27 << ((255) & 63))] = (fi27 >> ((466) & 63))
                _ = fi27
            }
        case 4294967295:
            v4 += (v4 & 0)
        }
    }
    var v28 int16 = 0
    var v29 int8 = -21
    v30 := GenS1{F0: -2147483648}
    v31 := GenS2{F0: 364, F1: (v28 | v28)}
    var v32 int8 = 127
    for fi33 := 0; fi33 < 12; fi33++ {
        if ((false || false) || true) {
            for fi34 := 0; fi34 < 19; fi34++ {
                v35 := GenS3{F0: (-(v4)), F1: 136, F2: v2}
                var v36 string = ""
                var v37 int = fi34
                v37 = (v37 | (fi33 & (fi33 - v37)))
                _ = fi34
                _ = v35
                _ = v36
                _ = v37
            }
            if true {
                fmt.Println(v4)
                if ((uint64(v2) == 233) || true) {
                    v38 := GenS4{F0: 0, F1: 544}
                    _ = v38
                } else {
                    var v39 uint = 735
                    v40 := GenS5{F0: (-1 >= int32(v2))}
                    _ = v39
                    _ = v40
                }
                if ((-466 <= 359) || ("a" <= ("\n\t" + "hello"))) {
                    var v41 uint = 339
                    v3[(fi33 >> (((v41 & (^(v41)))) & 63))] = (fi33 - fi33)
                    _ = v41
                } else {
                    fmt.Println(fi33, v29)
                    for fi42 := 0; fi42 < 19; fi42++ {
                        v28 = v28
                        if (1 < 101) {
                            var v43 string = ""
                            if false {
                                fmt.Println(fi33, v28, v29)
                                switch v43 {
                                case "hello":
                                    _ = len(v1)
                                case "":
                                    v15 = v15
                                }
                            } else {
                                v2 = (v2 ^ 65535)
                                var v44 int32 = 314
                                v15 += 816
                                _ = v44
                            }
                            if false {
                                v15 = v15
                            }
                            switch v43 {
                            case "foo bar":
                                fmt.Println(fi42, v43, v43)
                                switch 42 {
                                case 32768:
                                    _ = len(v1)
                                case 1:
                                    if false {
                                        switch v43 {
                                        case "foo bar":
                                            switch ("a" + v43) {
                                            case "hello":
                                                var v45 float64 = -2.25
                                                _ = v45
                                            case "":
                                                if (!true) {
                                                    switch "" {
                                                    case "a":
                                                        fmt.Println(fi33)
                                                    case "":
                                                        var v46 uint16 = 42
                                                        _ = v46
                                                    case "\n\t":
                                                        _ = len(v1)
                                                        if false {
                                                            _ = len(v3)
                                                            if false {
                                                                v4 -= 0
                                                            } else {
                                                                v47 := v3[(fi33 / ((fi42) | 1))]
                                                                if len(v1) > 0 {
                                                                    v1[0] = v47
                                                                }
                                                                _ = v47
                                                            }
                                                            v29 = (v29 | genFn1(v28, 42, -525))
                                                        } else {
                                                            if true {
                                                                var v48 uint16 = 65535
                                                                _ = v48
                                                            } else {
                                                                v29 = (v29 / ((genFn0(("hello" < "a"), 253, 2)) | 1))
                                                                _ = len(v1)
                                                            }
                                                            var v49 int64 = v4
                                                            switch fi33 {
                                                            case 32768:
                                                                v49 -= -9223372036854775808
                                                            default:
                                                                fmt.Println(v4)
                                                            }
                                                            _ = v49
                                                        }
                                                    }
                                                    if (0.0 > (1.5 * 1.5)) {
                                                        var v50 int32 = -166
                                                        _ = v50
                                                    }
                                                    delete(v3, (fi42 & fi42))
                                                }
                                            }
                                        case "":
                                            var v51 int64 = 9223372036854775807
                                            _ = v51
                                        }
                                    } else {
                                        var v52 bool = (!(!false))
                                        v53 := GenS2{F0: (v28 | -32768), F1: v28}
                                        if len(v1) > 0 {
                                            v1[0] = fi33
                                        }
                                        v54 := GenS6{F0: fi33, F1: 71}
                                        fmt.Println(v28, v2, v29)
                                        _ = v52
                                        _ = v53
                                        _ = v54
                                    }
                                }
                                switch (fi33 + -1) {
                                case -897:
                                    if false {
                                        v3[fi42] = (^(fi42))
                                    }
                                    var v55 int32 = 0
                                    _ = v55
                                case -710:
                                    v28 = v28
                                case -2147483648:
                                    var v56 float32 = 0.1
                                    _ = v56
                                }
                                v2 = 42
                                if (!true) {
                                    v15 = (^(v15))
                                }
                            case "":
                                v57 := GenS1{F0: -468}
                                _ = v57
                            case "\n\t":
                                switch "" {
                                case "a":
                                    var v58 uint16 = (v2 % (((^(v2))) | 1))
                                    _ = v58
                                case "hello":
                                    v32 = -1
                                default:
                                    fmt.Println(v15, v29)
                                }
                            default:
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                if (!((255 != 56) && false)) {
                                    fmt.Println(fi42, v43, v32)
                                    v43 = (v43 + "\n\t")
                                }
                                var v59 uint16 = (v2 / ((1) | 1))
                                _ = v59
                            }
                            var v60 bool = true
                            var v61 float32 = (3.14159 / 1e10)
                            var v62 int64 = v4
                            switch v43 {
                            case "a":
                                v28 = v28
                            case "hello":
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                            var v63 uint32 = v15
                            _ = v43
                            _ = v60
                            _ = v61
                            _ = v62
                            _ = v63
                        } else {
                            v64 := GenS7{F0: v2}
                            v15 = v15
                            v65 := GenS7{F0: (^(v2))}
                            v2 = uint16(v15)
                            _ = v64
                            _ = v65
                        }
                        _ = fi42
                    }
                    if false {
                        fmt.Println(fi33)
                    } else {
                        for fi66 := 0; fi66 < 6; fi66++ {
                            v3[(-(fi66))] = (fi33 - fi33)
                            v28 = int16(v29)
                            _ = fi66
                        }
                        var v67 float64 = 1e10
                        v68 := v3[0]
                        _ = v67
                        _ = v68
                    }
                }
            }
        }
        _ = fi33
    }
    _ = len(v1)
    v69 := GenS8{F0: 4294967295, F1: 9223372036854775807}
    v70 := v3[-111]
    v32 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v15
    _ = v2
    _ = v24
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v69
    _ = v70
}
