package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 uint16; F1 int32; F2 string }
type GenS3 struct { F0 uint64; F1 int; F2 bool }
type GenS4 struct { F0 int8; F1 uint; F2 uint32 }
type GenS5 struct { F0 int64; F1 string; F2 uint64 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint16; F1 int64 }

func main() {
    var v1 int16 = 0
    var v2 string = "\n\t"
    v3 := make(map[string]uint32)
    var v4 uint64 = 869
    var v5 uint16 = 32767
    v1 = v1
    var v6 int64 = 1
    v5 -= v5
    v7 := v3["foo bar"]
    var v8 uint32 = v7
    if false {
        v9 := []int{236, 4294967295}
        var v10 uint32 = uint32(v5)
        _ = v10
        _ = v9
    } else {
        v11 := GenS1{F0: 128, F1: -1}
        v12 := GenS2{F0: v5, F1: int32(v4), F2: v2}
        if ((!((1 > 0) || ((v1 + 32767) <= v1))) && true) {
            v13 := v3[v2]
            for fi14 := 0; fi14 < 2; fi14++ {
                v15 := GenS3{F0: (^(v4)), F1: fi14, F2: false}
                _ = fi14
                _ = v15
            }
            _ = v13
        } else {
            if (!(false || true)) {
                var v16 int32 = -1
                _ = v16
            } else {
                if false {
                    v6 -= (v6 | v6)
                    if (!false) {
                        if (0 >= 460) {
                            v17 := GenS4{F0: -1, F1: 1, F2: (^(v8))}
                            for fi18 := 0; fi18 < 15; fi18++ {
                                switch v2 {
                                case "\n\t":
                                    if true {
                                        v19 := GenS4{F0: -1, F1: 819, F2: v7}
                                        switch (fi18 + 191) {
                                        case 994:
                                            var v20 int8 = -53
                                            _ = v20
                                        case 518:
                                            var v21 float64 = (-2.25 * float64(float32((1e10 + -2.25))))
                                            _ = v21
                                        default:
                                            var v22 float64 = (0.0 + float64((((-2.25 * 1e10) / 3.14159) + 3.14159)))
                                            _ = v22
                                        }
                                        _ = v19
                                    }
                                }
                                _ = fi18
                            }
                            if true {
                                var v23 int32 = 1
                                _ = v23
                            }
                            _ = v17
                        } else {
                            fmt.Println(v2, v5)
                            var v24 uint16 = (v5 & v5)
                            v25 := GenS4{F0: -1, F1: 18446744073709551615, F2: v7}
                            _ = v24
                            _ = v25
                        }
                    }
                } else {
                    for fi26 := 0; fi26 < 1; fi26++ {
                        if false {
                            v3[v2] = 0
                        }
                        var v27 uint8 = 10
                        v6 = v6
                        _ = fi26
                        _ = v27
                    }
                    var v28 int8 = -1
                    _ = v28
                }
                switch 4294967295 {
                case -1:
                    if true {
                        if (v2 < (((v2 + "a") + v2) + "foo bar")) {
                            switch "foo bar" {
                            case "hello":
                                if false {
                                    var v29 int32 = -191
                                    _ = v29
                                } else {
                                    var v30 float64 = (1e10 + float64(3.14159))
                                    switch v2 {
                                    case "a":
                                        for fi31 := 0; fi31 < 13; fi31++ {
                                            v32 := []int{fi31, 4294967295, fi31, (^(fi31))}
                                            if false {
                                                v33 := GenS5{F0: (v6 / ((v6) | 1)), F1: v2, F2: (v4 - 262)}
                                                _ = v33
                                            } else {
                                                v7 = (v7 - v7)
                                                if ((v1 & (v1 | (v1 & (v1 + v1)))) < v1) {
                                                    var v34 uint16 = (v5 / ((255) | 1))
                                                    var v35 int = (^(fi31))
                                                    v36 := GenS6{F0: 130}
                                                    _ = v34
                                                    _ = v35
                                                    _ = v36
                                                }
                                            }
                                            _ = fi31
                                            _ = v32
                                        }
                                    case "\n\t":
                                        var v37 int32 = 804
                                        _ = v37
                                    case "hello":
                                        var v38 int16 = -317
                                        _ = v38
                                    default:
                                        v8 = v8
                                    }
                                    _ = v30
                                }
                            default:
                                v39 := []int{2147483647, int(v6), 451}
                                _ = v39
                            }
                            var v40 uint8 = 255
                            _ = v40
                        } else {
                            var v41 uint32 = v8
                            if (false && true) {
                                v3[v2] = 614
                            } else {
                                v42 := GenS2{F0: (v5 & (v5 / (((v5 * 891)) | 1))), F1: -1, F2: "hello"}
                                fmt.Println(v4, v6, v4)
                                _ = v42
                            }
                            _ = v41
                        }
                    }
                case -286:
                    fmt.Println(v7)
                }
            }
            if (v8 > (v8 / (((v8 & 2)) | 1))) {
                switch 591 {
                case -2147483648:
                    v1 = (v1 - (v1 + 923))
                default:
                    var v43 bool = (((!(((1 == -2147483648) && false) && false)) && (0.1 != (0.1 / 1.5))) || (true || false))
                    _ = v43
                }
            } else {
                if (false || false) {
                    v5 = 0
                    fmt.Println(v6)
                    fmt.Println(v5, v1, v6)
                }
                delete(v3, (v2 + v2))
                if true {
                    var v44 uint32 = (v8 << ((65535) & 63))
                    _ = v44
                }
            }
            v45 := v3["foo bar"]
            _ = v45
        }
        _ = v11
        _ = v12
    }
    v1 += v1
    for fi46 := 0; fi46 < 19; fi46++ {
        var v47 float64 = float64(0.1)
        _ = fi46
        _ = v47
    }
    if (false || false) {
        v48 := GenS6{F0: 81}
        _ = v48
    }
    var v49 uint = 42
    var v50 uint8 = 161
    v4 += (v4 % ((v4) | 1))
    var v51 uint32 = (^(v7))
    switch 32768 {
    case 32767:
        for fi52 := 0; fi52 < 5; fi52++ {
            if (fi52 < fi52) {
                var v53 uint32 = (v8 | (v7 >> (((v49 * v49)) & 63)))
                for fi54 := 0; fi54 < 1; fi54++ {
                    v55 := GenS7{F0: 65535, F1: -819}
                    var v56 uint = (v49 * v49)
                    var v57 int8 = 42
                    _ = fi54
                    _ = v55
                    _ = v56
                    _ = v57
                }
                _ = v53
            }
            _ = fi52
        }
    case 9223372036854775807:
        for fi58 := 0; fi58 < 9; fi58++ {
            v59 := []int{}
            v51 += v51
            if (!((^(v5)) >= v5)) {
                var v60 int8 = 0
                var v61 uint8 = v50
                if (v49 >= v49) {
                    v62 := GenS4{F0: 0, F1: uint(v4), F2: v7}
                    _ = v62
                } else {
                    var v63 string = v2
                    if len(v59) > 0 {
                        v59[0] = fi58
                    }
                    if (false && ((v51 & v51) != 906)) {
                        var v64 int16 = (v1 & (v1 * 42))
                        _ = v64
                    } else {
                        v65 := GenS7{F0: (v5 % ((v5) | 1)), F1: v6}
                        if true {
                            v5 += v5
                            v63 = v63
                        } else {
                            fmt.Println(v61)
                            v61 = uint8(v4)
                        }
                        _ = v65
                    }
                    _ = v63
                }
                _ = v60
                _ = v61
            }
            _ = fi58
            _ = v59
        }
    default:
        switch 127 {
        case -1:
            var v66 float64 = float64((0.1 + (-2.25 / (-2.25 / 0.1))))
            _ = v66
        case 127:
            for fi67 := 0; fi67 < 13; fi67++ {
                var v68 float64 = 0.0
                _ = fi67
                _ = v68
            }
        case -494:
            for fi69 := 0; fi69 < 3; fi69++ {
                switch "hello" {
                case "\n\t":
                    if false {
                        for fi70 := 0; fi70 < 3; fi70++ {
                            var v71 uint8 = (v50 << ((uint(v4)) & 63))
                            var v72 uint8 = v50
                            if false {
                                var v73 uint = 0
                                v72 = 190
                                v4 += 824
                                _ = v73
                            }
                            _ = fi70
                            _ = v71
                            _ = v72
                        }
                    } else {
                        fmt.Println(v6, v49, v1)
                        var v74 float32 = 1e10
                        switch (^(fi69)) {
                        case 4294967295:
                            _ = len(v3)
                        case 32767:
                            v75 := GenS2{F0: v5, F1: 0, F2: v2}
                            _ = v75
                        case 128:
                            v6 -= (-(v6))
                        default:
                            v76 := []int{}
                            _ = v76
                        }
                        _ = v74
                    }
                case "a":
                    var v77 float64 = 1.5
                    _ = v77
                default:
                    if ("\n\t" != (v2 + "hello")) {
                        _ = len(v3)
                        v78 := v3[v2]
                        v79 := []int{(fi69 - 606), (fi69 + -485)}
                        _ = v78
                        _ = v79
                    }
                }
                _ = fi69
            }
        default:
            var v80 int8 = 127
            _ = v80
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v6
    _ = v7
    _ = v8
}
