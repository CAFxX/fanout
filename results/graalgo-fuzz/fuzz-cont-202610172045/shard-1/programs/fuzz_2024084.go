package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 uint16; F2 uint8 }
type GenS3 struct { F0 int8; F1 int }
type GenS4 struct { F0 int8 }
func genFn0(p0 uint32, p1 uint) int {
    for fi6 := 0; fi6 < 4; fi6++ {
        v7 := GenS1{F0: 200}
        if (101 < 443) {
            v8 := GenS2{F0: 9223372036854775807, F1: 1, F2: 255}
            _ = v8
        } else {
            v9 := GenS1{F0: 2}
            if (2 <= 187) {
                v10 := GenS3{F0: 0, F1: (fi6 / ((fi6) | 1))}
                v11 := GenS2{F0: 0, F1: 65535, F2: uint8(p0)}
                _ = v10
                _ = v11
            }
            v12 := GenS2{F0: -198, F1: 2, F2: 0}
            _ = v12
            _ = v9
        }
        if (fi6 == -105) {
            var v13 uint16 = 65535
            var v14 int16 = -803
            _ = v13
            _ = v14
        } else {
            v15 := []int{(fi6 << (((p1 >> ((4294967295) & 63))) & 63))}
            _ = v15
        }
        _ = fi6
        _ = v7
    }
    p0 = 42
    if false {
        v16 := GenS1{F0: 200}
        v17 := []int{481, 0, -9223372036854775808, 4294967295}
        _ = v16
        _ = v17
    } else {
        switch 1 {
        case 81:
            var v18 int16 = -198
            _ = v18
        case -598:
            switch "hello" {
            case "foo bar":
                if false {
                    var v19 int32 = -2147483648
                    for fi20 := 0; fi20 < 5; fi20++ {
                        fmt.Println(v19, fi20)
                        switch "hello" {
                        case "foo bar":
                            var v21 uint = 786
                            _ = v21
                        case "a":
                            for fi22 := 0; fi22 < 6; fi22++ {
                                v23 := GenS3{F0: -5, F1: fi20}
                                v24 := GenS2{F0: 1, F1: 141, F2: 255}
                                fmt.Println(p0)
                                _ = fi22
                                _ = v23
                                _ = v24
                            }
                        case "\n\t":
                            var v25 uint8 = 180
                            _ = v25
                        default:
                            var v26 int64 = 479
                            _ = v26
                        }
                        _ = fi20
                    }
                    _ = v19
                }
            }
        }
    }
    var v27 uint16 = 864
    _ = p0
    _ = p1
    _ = v27
    return 1
}

func main() {
    var v1 uint8 = 0
    v2 := make(map[uint32]string)
    v3 := []int{1}
    v4 := make(map[int]int)
    v5 := []uint32{116}
    v28 := GenS2{F0: 42, F1: 65535, F2: (^(v1))}
    switch "\n\t" {
    case "foo bar":
        var v29 float32 = -2.25
        _ = v29
    case "\n\t":
        if len(v3) > 0 {
            v3[0] = -1
        }
    default:
        v1 = 128
    }
    v1 = 43
    for fi30 := 0; fi30 < 8; fi30++ {
        v3 = append(v3, fi30)
        for fi31 := 0; fi31 < 9; fi31++ {
            if true {
                var v32 int32 = -1
                v32 = v32
                switch fi30 {
                case -2147483648:
                    var v33 uint = 227
                    _ = v33
                case 0:
                    var v34 uint32 = 2147483648
                    _ = v34
                default:
                    if ((fi31 & 32767) < (-(fi30))) {
                        v35 := v4[fi30]
                        v3 = append(v3, (^(fi30)))
                        _ = v35
                    }
                }
                _ = len(v2)
                _ = v32
            } else {
                v36 := GenS1{F0: (^(v1))}
                _ = v36
            }
            if (false && (1e10 == ((0.1 + 0.0) + 3.14159))) {
                v3 = append(v3, fi31)
                v1 = v1
            } else {
                var v37 uint16 = 466
                if true {
                    _ = len(v2)
                    v1 = v1
                    var v38 float32 = 3.14159
                    _ = v38
                } else {
                    if true {
                        v39 := GenS3{F0: int8(v37), F1: genFn0(150, 284)}
                        _ = v39
                    }
                    var v40 int16 = 1000
                    if false {
                        var v41 uint16 = 483
                        v42 := GenS4{F0: 104}
                        _ = v41
                        _ = v42
                    } else {
                        fmt.Println(fi31, fi30, v1)
                        if len(v3) > 0 {
                            v3[0] = (fi30 & 127)
                        }
                    }
                    _ = v40
                }
                var v43 uint32 = 2147483648
                _ = v37
                _ = v43
            }
            v44 := v2[2147483648]
            v44 = ((v44 + ((v44 + (((v44 + (v44 + (v44 + v44))) + "hello") + (v44 + "\n\t"))) + (v44 + (v44 + v44)))) + ((v44 + ((((v44 + ("a" + v44)) + "a") + "") + (v44 + "a"))) + "hello"))
            v1 = 41
            _ = fi31
            _ = v44
        }
        var v45 uint32 = 751
        _ = fi30
        _ = v45
    }
    var v46 bool = (("\n\t" + ("hello" + ("hello" + "a"))) != (("" + ("" + "hello")) + (("a" + ("foo bar" + ("\n\t" + ""))) + "foo bar")))
    for fi47 := 0; fi47 < 8; fi47++ {
        var v48 bool = ("foo bar" <= "\n\t")
        var v49 float32 = 3.14159
        _ = fi47
        _ = v48
        _ = v49
    }
    fmt.Println(v1, v46)
    v50 := GenS3{F0: -1, F1: 807}
    v46 = (("foo bar" != ("hello" + ("foo bar" + "\n\t"))) && v46)
    if true {
        var v51 uint32 = 2
        var v52 float64 = float64(-2.25)
        v51 -= v51
        _ = v51
        _ = v52
    } else {
        fmt.Println(v1)
    }
    for fi53 := 0; fi53 < 3; fi53++ {
        for fi54 := 0; fi54 < 2; fi54++ {
            v55 := v4[(fi54 ^ (^(fi53)))]
            _ = fi54
            _ = v55
        }
        _ = fi53
    }
    for fi56 := 0; fi56 < 5; fi56++ {
        if false {
            v3 = append(v3, (fi56 << ((262) & 63)))
            fmt.Println(v46)
        }
        v46 = (32767 >= 42)
        _ = fi56
    }
    v57 := GenS3{F0: -128, F1: 2147483647}
    fmt.Println(v1, v1)
    v1 = (v1 & (^(v1)))
    for fi58 := 0; fi58 < 19; fi58++ {
        v59 := v4[fi58]
        if v46 {
            v60 := GenS2{F0: 9223372036854775807, F1: 703, F2: v1}
            v61 := v2[42]
            var v62 int8 = 0
            _ = v60
            _ = v61
            _ = v62
        } else {
            if (!(v46 || (0 == 37))) {
                var v63 uint32 = 2147483647
                if (!(42 == -374)) {
                    v59 += (v59 % ((v59) | 1))
                    fmt.Println(fi58, v1)
                }
                fmt.Println(v59, fi58)
                _ = v63
            }
            v5 = append(v5, 4294967295)
            var v64 uint16 = 255
            _ = v64
        }
        _ = fi58
        _ = v59
    }
    fmt.Println(v46)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v50
    _ = v57
}
