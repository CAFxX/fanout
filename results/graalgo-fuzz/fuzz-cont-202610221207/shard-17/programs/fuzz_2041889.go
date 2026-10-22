package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int; F1 int }
type GenS4 struct { F0 int16; F1 uint16 }
type GenS5 struct { F0 int32; F1 uint }
type GenS6 struct { F0 uint64; F1 int64; F2 uint16 }
type GenS7 struct { F0 string; F1 int32 }
type GenS8 struct { F0 uint32; F1 uint; F2 uint8 }
type GenS9 struct { F0 bool; F1 int64; F2 int64 }
func genFn0(p0 string) int32 {
    v7 := GenS1{F0: -437, F1: "foo bar"}
    v8 := GenS2{F0: 255}
    _ = p0
    _ = v7
    _ = v8
    return -2147483648
}
func genFn1() uint32 {
    var v9 bool = true
    if (42 <= 0) {
        v10 := GenS2{F0: 65535}
        if v9 {
            v9 = v9
            for fi11 := 0; fi11 < 4; fi11++ {
                v12 := []int{(fi11 / ((fi11) | 1)), fi11, fi11, fi11}
                switch fi11 {
                case 309:
                    if v9 {
                        if v9 {
                            v9 = false
                            for fi13 := 0; fi13 < 3; fi13++ {
                                v14 := GenS3{F0: -9223372036854775808, F1: fi11}
                                _ = fi13
                                _ = v14
                            }
                        } else {
                            v9 = ((2147483647 > 42) || false)
                            for fi15 := 0; fi15 < 17; fi15++ {
                                fmt.Println(fi11, v9)
                                fmt.Println(fi15, v9, fi15)
                                _ = fi15
                            }
                            v16 := GenS4{F0: 32767, F1: 32767}
                            _ = v16
                        }
                        if len(v12) > 0 {
                            v12[0] = fi11
                        }
                        v17 := GenS3{F0: fi11, F1: (fi11 / ((fi11) | 1))}
                        _ = v17
                    } else {
                        v9 = (false || false)
                        v18 := GenS5{F0: 1, F1: 0}
                        _ = v18
                    }
                default:
                    fmt.Println(v9, v9)
                }
                for fi19 := 0; fi19 < 12; fi19++ {
                    if ((v9 && v9) && v9) {
                        v9 = true
                        v20 := GenS2{F0: 977}
                        _ = v20
                    }
                    _ = fi19
                }
                _ = fi11
                _ = v12
            }
            if v9 {
                var v21 uint8 = 0
                for fi22 := 0; fi22 < 17; fi22++ {
                    for fi23 := 0; fi23 < 8; fi23++ {
                        v24 := GenS6{F0: 1, F1: 0, F2: 2}
                        v21 -= v21
                        _ = fi23
                        _ = v24
                    }
                    _ = fi22
                }
                _ = v21
            } else {
                var v25 int64 = -1
                fmt.Println(v25)
                _ = v25
            }
        } else {
            for fi26 := 0; fi26 < 16; fi26++ {
                switch "" {
                case "\n\t":
                    if (!(-128 == 27)) {
                        v9 = v9
                        v27 := []int{(^(fi26))}
                        _ = v27
                    } else {
                        for fi28 := 0; fi28 < 6; fi28++ {
                            if false {
                                switch (fi26 - (fi28 / (((fi28 << ((482) & 63))) | 1))) {
                                case -546:
                                    v29 := []int{(fi28 ^ fi28), (fi26 ^ fi26), -9223372036854775808}
                                    _ = v29
                                case 2147483647:
                                    v30 := []int{(fi26 % ((42) | 1)), fi28, fi26, (fi26 + -1)}
                                    _ = v30
                                default:
                                    v9 = v9
                                }
                                if true {
                                    fmt.Println(v9, fi26)
                                    var v31 bool = (v9 || false)
                                    v9 = (32767 > 1)
                                    _ = v31
                                } else {
                                    v32 := GenS7{F0: "foo bar", F1: -2147483648}
                                    v33 := []int{32768, (fi28 ^ 32767), (fi28 << ((18446744073709551615) & 63))}
                                    v9 = (v9 && true)
                                    _ = v32
                                    _ = v33
                                }
                                v9 = false
                            } else {
                                var v34 int32 = 780
                                _ = v34
                            }
                            _ = fi28
                        }
                        for fi35 := 0; fi35 < 7; fi35++ {
                            v9 = ((3.14159 >= (-2.25 + ((-2.25 + 1.5) + 0.0))) || (42 > 371))
                            _ = fi35
                        }
                    }
                case "":
                    var v36 float32 = float32(((1.5 / 3.14159) - 1.5))
                    _ = v36
                case "hello":
                    var v37 int8 = 40
                    _ = v37
                default:
                    fmt.Println(fi26, v9)
                }
                for fi38 := 0; fi38 < 7; fi38++ {
                    if true {
                        v39 := []int{}
                        v9 = v9
                        _ = v39
                    } else {
                        var v40 uint32 = 65535
                        v40 -= v40
                        var v41 int16 = 1
                        _ = v40
                        _ = v41
                    }
                    _ = fi38
                }
                v42 := []int{-1, (fi26 >> ((4294967295) & 63)), 159, (-(fi26))}
                _ = fi26
                _ = v42
            }
        }
        _ = v10
    } else {
        for fi43 := 0; fi43 < 6; fi43++ {
            for fi44 := 0; fi44 < 16; fi44++ {
                v45 := GenS5{F0: 0, F1: 4294967295}
                _ = fi44
                _ = v45
            }
            v46 := GenS8{F0: 561, F1: 922, F2: 62}
            if (959 > fi43) {
                for fi47 := 0; fi47 < 7; fi47++ {
                    var v48 float32 = 3.14159
                    var v49 int = (fi43 >> ((255) & 63))
                    fmt.Println(fi43, fi47)
                    _ = fi47
                    _ = v48
                    _ = v49
                }
                v50 := []int{}
                v9 = false
                _ = v50
            }
            _ = fi43
            _ = v46
        }
        for fi51 := 0; fi51 < 12; fi51++ {
            if v9 {
                v52 := GenS9{F0: (307 > -1), F1: 739, F2: 770}
                _ = v52
            } else {
                v53 := GenS3{F0: -811, F1: fi51}
                _ = v53
            }
            _ = fi51
        }
    }
    v54 := []int{}
    var v55 float64 = float64((0.0 + 1.5))
    v55 += float64((1.5 * 3.14159))
    _ = v54
    _ = v55
    _ = v9
    return 952
}

func main() {
    v1 := []string{"", "foo bar"}
    v2 := make(map[int]uint32)
    v3 := []int{42}
    v4 := make(map[int]uint32)
    var v5 int64 = 293
    v6 := []int{0, 400, 4294967295}
    v5 += v5
    fmt.Println(v5)
    _ = len(v2)
    delete(v4, -2147483648)
    var v56 bool = (1.5 <= (-2.25 + (1.5 * 1.5)))
    var v57 float64 = (((((1.5 / 1e10) / ((0.0 / 1.0) + 0.1)) - 0.0) / ((float64(0.1) + (0.0 / 3.14159)) + (1.5 + (1e10 * ((((((1.5 + (1e10 - ((0.0 / float64(-2.25)) / ((3.14159 / 1e10) - ((1e10 * 0.1) - (0.0 + (0.0 * (float64((0.0 * 1.5)) - 0.0)))))))) - ((0.1 * (float64(-2.25) * float64(1.5))) - 3.14159)) - 0.0) + 1e10) / -2.25) + (((0.1 * (-2.25 * 1.5)) + 0.1) - 3.14159)))))) + 0.0)
    v58 := GenS2{F0: 42}
    var v59 int8 = 0
    var v60 uint = 153
    v57 = v57
    var v61 float32 = 0.1
    if (!(true || true)) {
        var v62 int16 = -700
        if v56 {
            for fi63 := 0; fi63 < 14; fi63++ {
                v61 = v61
                var v64 uint = (v60 * (v60 & v60))
                v65 := GenS4{F0: (v62 << ((0) & 63)), F1: 342}
                _ = fi63
                _ = v64
                _ = v65
            }
        }
        _ = v62
    }
    for fi66 := 0; fi66 < 16; fi66++ {
        if v56 {
            for fi67 := 0; fi67 < 12; fi67++ {
                _ = len(v3)
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                if ((v59 % ((v59) | 1)) != (v59 & -128)) {
                    if false {
                        switch (fi67 << ((v60) & 63)) {
                        case 42:
                            switch ("" + "hello") {
                            case "a":
                                if v56 {
                                    v68 := GenS1{F0: v5, F1: "foo bar"}
                                    v61 = v61
                                    _ = v68
                                } else {
                                    var v69 string = (((("hello" + "foo bar") + "hello") + ("foo bar" + (("" + "hello") + ("foo bar" + ("hello" + ("a" + ("a" + "\n\t"))))))) + "a")
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    _ = len(v6)
                                    _ = v69
                                }
                            }
                        case 707:
                            v70 := GenS3{F0: fi66, F1: fi67}
                            _ = v70
                        case -961:
                            var v71 int32 = 1
                            v60 = (v60 % (((v60 >> ((v60) & 63))) | 1))
                            _ = v71
                        default:
                            fmt.Println(v5, fi67, v60)
                        }
                    }
                    switch fi66 {
                    case 1:
                        var v72 uint64 = 1
                        _ = v72
                    case -2147483648:
                        v73 := GenS5{F0: -2147483648, F1: 2}
                        _ = v73
                    default:
                        v61 -= ((v61 / v61) / 1e10)
                    }
                    var v74 string = "\n\t"
                    v75 := GenS1{F0: 544, F1: v74}
                    switch (v74 + v74) {
                    case "":
                        var v76 int64 = -220
                        var v77 uint16 = 65535
                        v76 = 0
                        _ = v76
                        _ = v77
                    case "hello":
                        if v56 {
                            fmt.Println(v57, v56, v60)
                            if v56 {
                                v6 = append(v6, -2147483648)
                                if (!(883 < genFn1())) {
                                    _ = len(v1)
                                } else {
                                    var v78 uint16 = 151
                                    delete(v2, (fi66 >> ((18446744073709551615) & 63)))
                                    _ = v78
                                }
                            } else {
                                var v79 int16 = -1
                                _ = v79
                            }
                        } else {
                            v61 += v61
                            v80 := GenS1{F0: int64(v60), F1: ("foo bar" + (v74 + (v74 + "foo bar")))}
                            _ = v80
                        }
                    default:
                        fmt.Println(v57, fi66)
                    }
                    v74 = "\n\t"
                    _ = v74
                    _ = v75
                } else {
                    v60 += (v60 % ((v60) | 1))
                }
                _ = fi67
            }
            v81 := GenS2{F0: 32767}
            switch fi66 {
            case -40:
                _ = len(v3)
            }
            _ = v81
        }
        var v82 int8 = v59
        _ = fi66
        _ = v82
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v59:", v59)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v57
    _ = v58
    _ = v59
    _ = v6
    _ = v60
    _ = v61
}
