package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint16; F1 int32 }
type GenS6 struct { F0 bool; F1 int16; F2 uint64 }
type GenS7 struct { F0 string; F1 int64; F2 int }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int32; F1 uint64; F2 int32 }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 uint8; F1 int8; F2 int64 }
type GenS12 struct { F0 uint16 }
func genFn0(p0 string) uint16 {
    switch p0 {
    case "\n\t":
        for fi5 := 0; fi5 < 18; fi5++ {
            switch (fi5 << ((0) & 63)) {
            case -2147483648:
                var v6 float32 = 0.1
                _ = v6
            case 128:
                if ((2 != 682) && (2147483647 < 65535)) {
                    var v7 string = (p0 + (p0 + "\n\t"))
                    var v8 uint32 = 2
                    for fi9 := 0; fi9 < 5; fi9++ {
                        var v10 bool = (p0 != ("\n\t" + "hello"))
                        _ = fi9
                        _ = v10
                    }
                    _ = v7
                    _ = v8
                } else {
                    for fi11 := 0; fi11 < 20; fi11++ {
                        fmt.Println(fi11)
                        _ = fi11
                    }
                    v12 := GenS1{F0: 216, F1: 1}
                    _ = v12
                }
            default:
                v13 := GenS1{F0: 128, F1: 42}
                _ = v13
            }
            _ = fi5
        }
    case "foo bar":
        for fi14 := 0; fi14 < 4; fi14++ {
            var v15 float32 = 0.0
            var v16 float64 = (0.0 / (((1e10 - 0.0) / float64((float32(0.0) + float32(1.5)))) / 0.1))
            _ = fi14
            _ = v15
            _ = v16
        }
    default:
        var v17 uint8 = 114
        _ = v17
    }
    switch -2147483648 {
    case 32768:
        var v18 bool = (42 >= 42)
        _ = v18
    }
    v19 := []int{32767, 423, 0, 581}
    var v20 int = 32767
    if (v20 != (v20 % ((-533) | 1))) {
        v21 := GenS2{F0: 42}
        if (true || true) {
            var v22 uint32 = 278
            for fi23 := 0; fi23 < 6; fi23++ {
                if len(v19) >= 2 {
                    v19 = v19[1:len(v19)]
                }
                fmt.Println(v22)
                if false {
                    if (1 <= 0) {
                        var v24 int64 = 25
                        v25 := GenS3{F0: 0}
                        if true {
                            for fi26 := 0; fi26 < 10; fi26++ {
                                p0 = p0
                                v27 := GenS2{F0: 1000}
                                _ = fi26
                                _ = v27
                            }
                            var v28 uint32 = 837
                            _ = v28
                        } else {
                            v29 := GenS2{F0: -1}
                            if true {
                                v24 += 9223372036854775807
                            }
                            fmt.Println(v24, p0)
                            _ = v29
                        }
                        _ = v24
                        _ = v25
                    }
                    v22 = (v22 - (v22 << ((42) & 63)))
                } else {
                    for fi30 := 0; fi30 < 8; fi30++ {
                        v31 := GenS2{F0: 204}
                        if true {
                            if true {
                                if (((-2147483648 <= -1) && (!true)) && false) {
                                    var v32 int32 = -2147483648
                                    _ = v32
                                }
                                var v33 uint8 = 191
                                _ = v33
                            }
                            v34 := GenS4{F0: 600}
                            v35 := GenS5{F0: 2, F1: -154}
                            _ = v34
                            _ = v35
                        }
                        _ = fi30
                        _ = v31
                    }
                }
                _ = fi23
            }
            if len(v19) > 0 {
                v19[0] = -47
            }
            _ = v22
        } else {
            v19 = append(v19, (v20 / ((128) | 1)))
            if len(v19) > 0 {
                v19[0] = (v20 + v20)
            }
            if true {
                v19 = append(v19, 1)
                v20 -= (v20 / (((v20 | v20)) | 1))
            } else {
                var v36 int8 = -1
                _ = v36
            }
        }
        _ = v21
    } else {
        var v37 int32 = 1
        for fi38 := 0; fi38 < 18; fi38++ {
            for fi39 := 0; fi39 < 4; fi39++ {
                if len(v19) >= 2 {
                    v19 = v19[1:len(v19)]
                }
                fmt.Println(p0, fi38)
                _ = fi39
            }
            fmt.Println(v20, v37)
            _ = fi38
        }
        _ = v37
    }
    _ = p0
    _ = v19
    _ = v20
    return 1
}
func genFn1(p0 string, p1 int64) string {
    p1 = (p1 / ((p1) | 1))
    for fi40 := 0; fi40 < 6; fi40++ {
        for fi41 := 0; fi41 < 1; fi41++ {
            var v42 uint = 2
            _ = fi41
            _ = v42
        }
        _ = fi40
    }
    for fi43 := 0; fi43 < 4; fi43++ {
        var v44 bool = (-2.25 != 0.0)
        var v45 uint16 = 549
        _ = fi43
        _ = v44
        _ = v45
    }
    if (0.1 <= (1e10 / 0.1)) {
        p1 += (^(p1))
        fmt.Println(p0, p0)
        if (true || true) {
            if true {
                fmt.Println(p0)
                if ((0.0 + (1.5 / ((1.5 * (float32(1.5) + float32(1e10))) - float32(3.14159)))) == 1.5) {
                    v46 := []int{127, 1, -780}
                    _ = v46
                }
            }
            if (true && ((!false) || ((p0 == p0) && ((0.0 / 1e10) >= 0.1)))) {
                for fi47 := 0; fi47 < 6; fi47++ {
                    switch (fi47 | (-(fi47))) {
                    case -1:
                        p1 = (p1 - p1)
                    case -158:
                        for fi48 := 0; fi48 < 7; fi48++ {
                            p0 = (("foo bar" + p0) + "\n\t")
                            if false {
                                var v49 int = fi48
                                var v50 uint32 = 0
                                _ = v49
                                _ = v50
                            }
                            _ = fi48
                        }
                    }
                    v51 := []int{fi47}
                    _ = fi47
                    _ = v51
                }
            } else {
                var v52 int64 = (p1 >> ((4294967295) & 63))
                for fi53 := 0; fi53 < 5; fi53++ {
                    fmt.Println(p1, v52, p0)
                    var v54 int32 = -2147483648
                    p1 += v52
                    _ = fi53
                    _ = v54
                }
                _ = v52
            }
        } else {
            var v55 float64 = 1.5
            _ = v55
        }
    } else {
        v56 := GenS3{F0: 255}
        _ = v56
    }
    _ = p0
    _ = p1
    return "\n\t"
}

func main() {
    var v1 int = 838
    v2 := make(map[string]string)
    var v3 string = ("foo bar" + "")
    v4 := make(map[int]bool)
    var v57 int = (v1 * (v1 / (((v1 * v1)) | 1)))
    switch (v57 * 0) {
    case 32768:
        v57 += 42
    }
    v57 -= v1
    v58 := []int{(v1 + 9223372036854775807)}
    fmt.Println(v3, v3, v3)
    fmt.Println(v57)
    for fi59 := 0; fi59 < 3; fi59++ {
        var v60 uint = 424
        v57 += fi59
        if (v3 != v3) {
            v57 = (v57 / ((fi59) | 1))
            if false {
                fmt.Println(v60, v1)
                v61 := GenS6{F0: true, F1: 954, F2: 18446744073709551615}
                _ = v61
            } else {
                switch (v1 - 0) {
                case 71:
                    v62 := GenS7{F0: "a", F1: 895, F2: fi59}
                    _ = v62
                case 4294967295:
                    if (v3 <= ("a" + v3)) {
                        v4[-772] = true
                        v60 = 4294967295
                    } else {
                        v63 := GenS2{F0: 1}
                        if false {
                            var v64 float32 = 3.14159
                            _ = v64
                        }
                        _ = v63
                    }
                default:
                    v65 := GenS6{F0: false, F1: -656, F2: 0}
                    _ = v65
                }
                for fi66 := 0; fi66 < 3; fi66++ {
                    if len(v58) >= 2 {
                        v58 = v58[1:len(v58)]
                    }
                    v3 = "\n\t"
                    v3 = "foo bar"
                    _ = fi66
                }
                for fi67 := 0; fi67 < 9; fi67++ {
                    v60 = (v60 ^ 255)
                    _ = fi67
                }
            }
            if ((v3 + (v3 + "")) == v3) {
                v3 = v3
            }
        } else {
            for fi68 := 0; fi68 < 10; fi68++ {
                if false {
                    var v69 uint = (^(v60))
                    v58 = append(v58, (fi68 - (-(fi68))))
                    v70 := GenS8{F0: v3}
                    v57 += 127
                    if ((731 > 854) || (true && false)) {
                        var v71 int8 = -128
                        _ = v71
                    }
                    if true {
                        v72 := GenS9{F0: -2147483648, F1: 42, F2: -1}
                        v1 = (fi68 / ((v1) | 1))
                        _ = len(v4)
                        _ = v72
                    }
                    _ = v69
                    _ = v70
                } else {
                    v3 = "\n\t"
                }
                _ = fi68
            }
        }
        _ = fi59
        _ = v60
    }
    switch 42 {
    case -2147483648:
        v73 := GenS9{F0: 42, F1: 860, F2: -103}
        _ = v73
    case -974:
        if true {
            v2[(((v3 + v3) + (v3 + "")) + v3)] = (v3 + v3)
            if true {
                v58 = append(v58, v57)
            } else {
                fmt.Println(v57, v1, v3)
                var v74 int8 = -128
                fmt.Println(v1, v1)
                _ = v74
            }
        }
    case -531:
        _ = len(v4)
    default:
        v1 -= v1
    }
    if true {
        switch (v3 + v3) {
        case "a":
            switch ("foo bar" + v3) {
            case "":
                var v75 int64 = -928
                _ = v75
            case "foo bar":
                fmt.Println(v57, v1)
            case "hello":
                if (1.5 > 3.14159) {
                    v3 = "a"
                    v57 -= v1
                    v57 = (v57 & (v1 | v57))
                } else {
                    v76 := v4[(v57 / ((v1) | 1))]
                    _ = v76
                }
            }
        case "\n\t":
            v57 += v1
        case "":
            delete(v2, (v3 + "hello"))
        }
    } else {
        delete(v2, v3)
        if true {
            v3 = v3
            v77 := GenS10{F0: 0}
            v78 := GenS11{F0: 2, F1: -128, F2: -9223372036854775808}
            _ = v77
            _ = v78
        }
    }
    for fi79 := 0; fi79 < 9; fi79++ {
        for fi80 := 0; fi80 < 16; fi80++ {
            v58 = append(v58, fi79)
            v81 := GenS12{F0: 32767}
            v4[-1] = (2 <= 708)
            _ = fi80
            _ = v81
        }
        v82 := GenS6{F0: false, F1: 355, F2: 428}
        _ = fi79
        _ = v82
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v57:", v57)
    fmt.Println("slicelen v58:", len(v58), v58)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v57
    _ = v58
}
