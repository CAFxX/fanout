package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint32; F1 int32 }
type GenS5 struct { F0 uint8; F1 uint16 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint64; F1 int8 }
type GenS8 struct { F0 uint; F1 int64; F2 uint32 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 bool }
func genFn0(p0 uint8, p1 uint16) uint {
    var v4 uint64 = 904
    var v5 bool = true
    for fi6 := 0; fi6 < 13; fi6++ {
        v7 := GenS1{F0: (822 < (v4 * (v4 + v4)))}
        switch ("" + "a") {
        case "a":
            switch fi6 {
            case 2147483647:
                p1 += (p1 >> ((18446744073709551615) & 63))
            case -9223372036854775808:
                for fi8 := 0; fi8 < 10; fi8++ {
                    v5 = false
                    _ = fi8
                }
            case 42:
                v9 := []int{(fi6 + fi6), fi6}
                _ = v9
            default:
                v5 = ((v4 * 1) < (v4 % (((v4 << ((4294967295) & 63))) | 1)))
            }
        case "\n\t":
            v5 = v5
        case "hello":
            v10 := GenS1{F0: v5}
            _ = v10
        default:
            for fi11 := 0; fi11 < 8; fi11++ {
                v12 := GenS2{F0: v4, F1: v5}
                if (100 > 42) {
                    var v13 uint32 = 78
                    fmt.Println(p0, fi11)
                    v4 = v4
                    v14 := []int{(fi11 - fi11), (^(fi11)), fi11, (-(fi6))}
                    if len(v14) > 0 {
                        v14[0] = -9223372036854775808
                    }
                    v14 = append(v14, (fi11 / ((fi6) | 1)))
                    if len(v14) > 0 {
                        v14[0] = (fi6 % ((fi11) | 1))
                    }
                    _ = v13
                    _ = v14
                }
                if v5 {
                    p0 = (^(p0))
                    p0 = 127
                    var v15 uint8 = p0
                    _ = v15
                }
                _ = fi11
                _ = v12
            }
        }
        for fi16 := 0; fi16 < 14; fi16++ {
            var v17 float32 = float32((3.14159 / (3.14159 * -2.25)))
            v4 += 1
            v18 := GenS1{F0: false}
            _ = fi16
            _ = v17
            _ = v18
        }
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 1
}
func genFn1() uint8 {
    switch 32767 {
    case 1:
        v19 := GenS3{F0: 540}
        _ = v19
    default:
        for fi20 := 0; fi20 < 6; fi20++ {
            switch "\n\t" {
            case "hello":
                var v21 int = (-(fi20))
                _ = v21
            case "":
                switch ("hello" + "") {
                case "hello":
                    if true {
                        switch 914 {
                        case -9223372036854775808:
                            if true {
                                if false {
                                    v22 := []int{-254}
                                    if len(v22) > 0 {
                                        v22[0] = (fi20 ^ fi20)
                                    }
                                    _ = v22
                                }
                                if false {
                                    var v23 float32 = (1.5 * 3.14159)
                                    _ = v23
                                }
                                var v24 uint16 = 255
                                _ = v24
                            }
                        case 127:
                            var v25 uint64 = 1
                            _ = v25
                        }
                        var v26 uint8 = 22
                        if false {
                            v27 := GenS1{F0: ((true && true) || false)}
                            v28 := []int{}
                            for fi29 := 0; fi29 < 1; fi29++ {
                                v30 := GenS4{F0: 42, F1: 2147483647}
                                fmt.Println(fi20, v26, fi29)
                                v26 = v26
                                fmt.Println(fi29, fi29, fi20)
                                _ = fi29
                                _ = v30
                            }
                            _ = v27
                            _ = v28
                        }
                        _ = v26
                    } else {
                        var v31 uint16 = 149
                        v32 := []int{}
                        for fi33 := 0; fi33 < 14; fi33++ {
                            v31 += 65535
                            v34 := GenS5{F0: 128, F1: (v31 ^ 18)}
                            var v35 uint64 = uint64(v31)
                            _ = fi33
                            _ = v34
                            _ = v35
                        }
                        _ = v31
                        _ = v32
                    }
                case "a":
                    var v36 int16 = 42
                    _ = v36
                case "foo bar":
                    fmt.Println(fi20)
                }
            case "foo bar":
                if (42 < 349) {
                    if ((true && true) || false) {
                        var v37 int8 = -128
                        v37 = (v37 / ((v37) | 1))
                        _ = v37
                    } else {
                        for fi38 := 0; fi38 < 16; fi38++ {
                            var v39 uint8 = 1
                            var v40 int = (-(fi20))
                            _ = fi38
                            _ = v39
                            _ = v40
                        }
                        for fi41 := 0; fi41 < 2; fi41++ {
                            var v42 string = ""
                            _ = fi41
                            _ = v42
                        }
                        for fi43 := 0; fi43 < 18; fi43++ {
                            var v44 uint64 = 954
                            _ = fi43
                            _ = v44
                        }
                    }
                    var v45 float32 = 1.5
                    _ = v45
                }
            }
            _ = fi20
        }
    }
    v46 := GenS2{F0: 551, F1: true}
    _ = v46
    return 0
}

func main() {
    var v1 uint64 = 130
    var v2 int32 = -1
    var v3 float32 = 3.14159
    v2 = (v2 / (((v2 % ((-1) | 1))) | 1))
    for fi47 := 0; fi47 < 7; fi47++ {
        if false {
            fmt.Println(fi47, v1, v3)
        } else {
            for fi48 := 0; fi48 < 10; fi48++ {
                v49 := []int{(fi47 >> ((561) & 63)), (fi48 >> ((102) & 63)), int(v1)}
                switch fi47 {
                case -272:
                    v50 := GenS6{F0: ("" + "\n\t")}
                    _ = v50
                case 4294967295:
                    v51 := GenS7{F0: 18446744073709551615, F1: 42}
                    _ = v51
                default:
                    v1 += (v1 * v1)
                }
                _ = fi48
                _ = v49
            }
            if true {
                for fi52 := 0; fi52 < 3; fi52++ {
                    v3 = (-2.25 / (v3 + ((v3 * 0.0) + 0.0)))
                    var v53 float64 = 0.1
                    _ = fi52
                    _ = v53
                }
            }
        }
        v54 := GenS8{F0: 139, F1: 9223372036854775807, F2: 457}
        if (234 < 42) {
            if (674 == 648) {
                var v55 uint8 = 255
                fmt.Println(v2)
                var v56 bool = true
                _ = v55
                _ = v56
            } else {
                if false {
                    switch (fi47 | fi47) {
                    case 128:
                        fmt.Println(v2)
                    case 234:
                        var v57 int = fi47
                        _ = v57
                    case 4294967295:
                        v58 := []int{fi47}
                        _ = v58
                    }
                }
                v3 += 1e10
            }
            var v59 uint16 = 1
            _ = v59
        }
        _ = fi47
        _ = v54
    }
    v60 := []int{42}
    if len(v60) > 0 {
        v60[0] = 788
    }
    v61 := GenS5{F0: 143, F1: 255}
    v62 := GenS9{F0: (1 == 42)}
    for fi63 := 0; fi63 < 13; fi63++ {
        if true {
            var v64 uint64 = v1
            v3 -= (-2.25 - 3.14159)
            _ = v64
        } else {
            if true {
                v1 = 513
            }
            v3 -= ((v3 - (v3 - (float32((0.1 + 0.0)) * v3))) * v3)
        }
        var v65 string = ""
        for fi66 := 0; fi66 < 12; fi66++ {
            var v67 bool = true
            var v68 int32 = -2147483648
            _ = fi66
            _ = v67
            _ = v68
        }
        _ = fi63
        _ = v65
    }
    var v69 float32 = ((v3 + -2.25) / 3.14159)
    for fi70 := 0; fi70 < 5; fi70++ {
        var v71 uint64 = (v1 ^ (v1 | (v1 ^ (^(v1)))))
        if false {
            for fi72 := 0; fi72 < 3; fi72++ {
                var v73 int32 = v2
                fmt.Println(v69, v2, fi70)
                if len(v60) > 0 {
                    v60[0] = (fi70 * (fi72 / ((fi70) | 1)))
                }
                _ = len(v60)
                _ = fi72
                _ = v73
            }
        } else {
            v3 = (v3 * 0.1)
        }
        for fi74 := 0; fi74 < 8; fi74++ {
            v60 = append(v60, (fi70 >> ((426) & 63)))
            _ = len(v60)
            v75 := GenS1{F0: false}
            _ = fi74
            _ = v75
        }
        _ = fi70
        _ = v71
    }
    v3 = (0.0 - 0.1)
    v76 := GenS8{F0: 0, F1: 924, F2: 2}
    v1 = v1
    if true {
        var v77 int = 32768
        var v78 int16 = -32768
        v79 := GenS10{F0: (!false)}
        _ = v77
        _ = v78
        _ = v79
    }
    v3 -= (0.0 - (v69 + (1.5 * (v69 / v69))))
    v80 := GenS4{F0: 0, F1: v2}
    v69 = (3.14159 - 1.5)
    var v81 uint8 = 1
    v1 = (^(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v60:", len(v60), v60)
    fmt.Println("v69:", v69)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v3
    _ = v60
    _ = v61
    _ = v62
    _ = v69
    _ = v76
    _ = v80
    _ = v81
}
