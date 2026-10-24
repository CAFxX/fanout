package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 int64; F1 uint16; F2 int32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint16 }
func genFn0() uint32 {
    v7 := []int{-9223372036854775808, 1, 969}
    var v8 int64 = 1
    if ((false && false) || (false || true)) {
        v9 := GenS1{F0: 1, F1: 4294967295}
        _ = v9
    } else {
        for fi10 := 0; fi10 < 6; fi10++ {
            switch fi10 {
            case 577:
                v11 := GenS1{F0: -437, F1: 2}
                _ = v11
            case 42:
                var v12 int32 = 585
                _ = v12
            case 4294967295:
                if len(v7) > 0 {
                    v7[0] = (fi10 ^ fi10)
                }
            }
            v8 = v8
            _ = fi10
        }
        if false {
            switch -46 {
            case -316:
                var v13 int64 = -9223372036854775808
                _ = v13
            default:
                switch 2147483647 {
                case 530:
                    v8 += (v8 % ((v8) | 1))
                }
            }
            var v14 float32 = float32(1.5)
            fmt.Println(v8, v14)
            _ = v14
        }
    }
    var v15 uint64 = 801
    _ = v15
    _ = v7
    _ = v8
    return 681
}
func genFn1() uint32 {
    if false {
        if (894 >= 65535) {
            switch 4294967295 {
            case 406:
                switch "foo bar" {
                case "a":
                    var v16 uint8 = 128
                    _ = v16
                case "hello":
                    var v17 int8 = 100
                    _ = v17
                default:
                    var v18 int8 = 42
                    _ = v18
                }
            case 759:
                v19 := GenS1{F0: 1000, F1: 0}
                _ = v19
            default:
                for fi20 := 0; fi20 < 5; fi20++ {
                    if (false || false) {
                        if true {
                            var v21 uint16 = 42
                            _ = v21
                        }
                        v22 := []int{fi20, fi20, (-(fi20)), fi20}
                        switch "foo bar" {
                        case "foo bar":
                            for fi23 := 0; fi23 < 16; fi23++ {
                                var v24 uint8 = 2
                                v24 = (v24 / ((127) | 1))
                                v24 = 128
                                var v25 uint16 = 207
                                switch ("\n\t" + "") {
                                case "":
                                    fmt.Println(v25, v25, fi23)
                                    var v26 string = "\n\t"
                                    var v27 string = ("foo bar" + ((v26 + v26) + v26))
                                    v25 = 92
                                    v24 = v24
                                    _ = v26
                                    _ = v27
                                }
                                _ = fi23
                                _ = v24
                                _ = v25
                            }
                        case "":
                            switch (-(fi20)) {
                            case 954:
                                if false {
                                    var v28 int32 = -2147483648
                                    _ = v28
                                }
                            case 42:
                                if false {
                                    switch "a" {
                                    case "hello":
                                        fmt.Println(fi20)
                                    case "\n\t":
                                        if len(v22) >= 2 {
                                            v22 = v22[1:len(v22)]
                                        }
                                    case "foo bar":
                                        fmt.Println(fi20)
                                    default:
                                        v29 := GenS1{F0: 1, F1: 18446744073709551615}
                                        _ = v29
                                    }
                                    if len(v22) > 0 {
                                        v22[0] = (fi20 >> ((65535) & 63))
                                    }
                                    v30 := GenS1{F0: -234, F1: 907}
                                    _ = v30
                                }
                            }
                        case "hello":
                            var v31 int32 = 0
                            _ = v31
                        default:
                            for fi32 := 0; fi32 < 20; fi32++ {
                                var v33 uint32 = 475
                                fmt.Println(fi32)
                                _ = fi32
                                _ = v33
                            }
                        }
                        _ = v22
                    }
                    var v34 int64 = 9223372036854775807
                    for fi35 := 0; fi35 < 4; fi35++ {
                        fmt.Println(v34)
                        _ = fi35
                    }
                    _ = fi20
                    _ = v34
                }
            }
            var v36 uint8 = 1
            var v37 float32 = (float32(-2.25) * float32(1.5))
            _ = v36
            _ = v37
        }
        var v38 string = "a"
        _ = v38
    } else {
        var v39 float64 = 0.0
        _ = v39
    }
    var v40 uint = 1
    v41 := GenS2{F0: 1, F1: 32767, F2: -590}
    if (false && (!(!false))) {
        v42 := []int{2147483647, 32768, 856, 827}
        var v43 float32 = (((-2.25 - -2.25) - 0.0) * 0.0)
        if (-1 >= -1) {
            if len(v42) > 0 {
                v42[0] = -1
            }
            if len(v42) > 0 {
                v42[0] = 1
            }
        } else {
            var v44 float64 = ((0.1 - 3.14159) / ((1.5 + (0.1 + 0.0)) / 1.0))
            _ = v44
        }
        _ = v42
        _ = v43
    } else {
        v45 := GenS2{F0: 1, F1: 42, F2: 0}
        _ = v45
    }
    v46 := []int{1, 0, -2147483648}
    _ = v40
    _ = v41
    _ = v46
    return 65535
}

func main() {
    var v1 int8 = 1
    var v2 string = "foo bar"
    var v3 bool = ((42 > 32767) && false)
    var v4 uint64 = 4294967295
    var v5 float32 = (0.0 - ((1e10 / -2.25) - (1.5 + 0.0)))
    var v6 uint = 42
    if (0 > 4294967295) {
        v2 = v2
        v47 := []int{127, 2147483647, 740}
        _ = v47
    }
    var v48 uint32 = 4294967295
    fmt.Println(v48)
    var v49 int = -421
    v50 := GenS3{F0: (-(v1))}
    for fi51 := 0; fi51 < 2; fi51++ {
        var v52 float64 = 1e10
        if v3 {
            var v53 float64 = v52
            v54 := []int{(-(v49)), (v49 ^ v49)}
            v55 := GenS1{F0: 1, F1: v4}
            _ = v53
            _ = v54
            _ = v55
        } else {
            if (true && (42 >= 788)) {
                v56 := GenS4{F0: 489}
                v57 := []int{(fi51 & (-(fi51))), (fi51 * v49), -9223372036854775808, (v49 - int(v1))}
                if v3 {
                    if v3 {
                        v4 = v4
                    } else {
                        var v58 uint8 = 71
                        var v59 uint8 = 0
                        _ = v58
                        _ = v59
                    }
                }
                _ = v56
                _ = v57
            }
        }
        _ = fi51
        _ = v52
    }
    var v60 uint64 = uint64(v6)
    v61 := GenS1{F0: 1000, F1: v4}
    v62 := GenS2{F0: 9223372036854775807, F1: 949, F2: -2147483648}
    v5 = (v5 / v5)
    v5 += (float32(-2.25) * v5)
    for fi63 := 0; fi63 < 17; fi63++ {
        v64 := GenS5{F0: 2}
        _ = fi63
        _ = v64
    }
    if v3 {
        if (!((v3 && v3) || (1 == v4))) {
            for fi65 := 0; fi65 < 7; fi65++ {
                v48 = v48
                v3 = ((v5 - (v5 / v5)) > (float32(0.0) / float32(((0.0 * -2.25) - 3.14159))))
                if ((!((v6 ^ v6) > v6)) || v3) {
                    fmt.Println(v60, v48, v3)
                } else {
                    v66 := []int{(^(fi65))}
                    var v67 int = fi65
                    for fi68 := 0; fi68 < 15; fi68++ {
                        v66 = append(v66, -1)
                        _ = len(v66)
                        var v69 float64 = float64(v5)
                        _ = fi68
                        _ = v69
                    }
                    _ = v66
                    _ = v67
                }
                _ = fi65
            }
            v70 := GenS3{F0: int8(v4)}
            _ = v70
        }
        v71 := []int{v49, v49}
        _ = v71
    } else {
        for fi72 := 0; fi72 < 6; fi72++ {
            v1 = (v1 & (v1 + v1))
            _ = fi72
        }
        for fi73 := 0; fi73 < 14; fi73++ {
            var v74 string = v2
            if true {
                v60 = 42
                var v75 string = (v74 + (v2 + (v74 + v74)))
                _ = v75
            } else {
                v60 -= (v60 / ((v4) | 1))
                if (1e10 < 1e10) {
                    v76 := GenS6{F0: 18446744073709551615}
                    if v3 {
                        switch v74 {
                        case "hello":
                            if (((float32(3.14159) + v5) <= v5) && (-9223372036854775808 < (fi73 >> ((v6) & 63)))) {
                                if v3 {
                                    v77 := GenS4{F0: v48}
                                    var v78 bool = v3
                                    var v79 uint32 = v48
                                    _ = v77
                                    _ = v78
                                    _ = v79
                                } else {
                                    switch "hello" {
                                    case "hello":
                                        switch (^(v49)) {
                                        case 0:
                                            v80 := []int{(fi73 + 597), (v49 % ((127) | 1)), v49, (v49 - fi73)}
                                            _ = v80
                                        }
                                    case "\n\t":
                                        v81 := []int{v49, (v49 >> (((v6 % ((v6) | 1))) & 63)), (v49 % ((fi73) | 1)), v49}
                                        _ = v81
                                    case "":
                                        switch v74 {
                                        case "a":
                                            var v82 bool = (v60 <= v60)
                                            _ = v82
                                        default:
                                            v2 = (v74 + (v74 + v2))
                                        }
                                    }
                                    v83 := GenS7{F0: 882}
                                    _ = v83
                                }
                                v84 := []int{128, (fi73 & (-(fi73))), 32767, (v49 & (fi73 - (v49 & int(v1))))}
                                var v85 int16 = -32768
                                _ = v84
                                _ = v85
                            }
                        }
                        v86 := []int{}
                        v2 = "\n\t"
                        _ = v86
                    }
                    var v87 int8 = v1
                    _ = v76
                    _ = v87
                } else {
                    var v88 int32 = 42
                    v49 += 127
                    v89 := []int{(-(fi73)), (fi73 / ((v49) | 1)), (v49 * v49), (v49 % ((v49) | 1))}
                    _ = v88
                    _ = v89
                }
            }
            for fi90 := 0; fi90 < 5; fi90++ {
                v60 += v4
                var v91 uint32 = (v48 / ((2147483647) | 1))
                _ = fi90
                _ = v91
            }
            _ = fi73
            _ = v74
        }
        fmt.Println(v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v6
    _ = v60
    _ = v61
    _ = v62
}
