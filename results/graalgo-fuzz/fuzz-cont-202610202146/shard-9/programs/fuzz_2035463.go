package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 string }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint8; F1 int32; F2 uint32 }
type GenS6 struct { F0 int32; F1 uint64 }
type GenS7 struct { F0 int; F1 uint8 }
func genFn0() uint {
    for fi5 := 0; fi5 < 13; fi5++ {
        fmt.Println(fi5)
        _ = fi5
    }
    var v6 uint32 = 4294967295
    _ = v6
    return 65535
}

func main() {
    var v1 bool = ((!false) || ((-9223372036854775808 <= -924) || true))
    var v2 uint8 = 2
    var v3 uint64 = 929
    var v4 bool = (v1 && (v1 && (v3 <= v3)))
    for fi7 := 0; fi7 < 18; fi7++ {
        v8 := []int{}
        v2 += (v2 | v2)
        v9 := GenS1{F0: 4294967295, F1: 405}
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := GenS2{F0: 1}
    v11 := GenS1{F0: 2, F1: v3}
    fmt.Println(v2, v4)
    var v12 bool = v4
    v4 = true
    switch 461 {
    case 621:
        v13 := []int{}
        _ = v13
    case -9223372036854775808:
        v14 := GenS3{F0: "a"}
        _ = v14
    default:
        if v12 {
            var v15 uint8 = v2
            v2 -= 42
            _ = v15
        } else {
            v1 = (-1 != -1)
        }
    }
    if false {
        if v4 {
            for fi16 := 0; fi16 < 15; fi16++ {
                var v17 string = "foo bar"
                _ = fi16
                _ = v17
            }
            v18 := []int{-567, -920, -9223372036854775808}
            _ = v18
        }
        v19 := []int{-1, 362, 42}
        _ = v19
    }
    for fi20 := 0; fi20 < 16; fi20++ {
        for fi21 := 0; fi21 < 14; fi21++ {
            v1 = v12
            v22 := GenS2{F0: uint(v3)}
            v23 := GenS2{F0: 42}
            fmt.Println(fi20)
            if true {
                var v24 float32 = 3.14159
                v4 = true
                v25 := GenS4{F0: (538 != 42)}
                _ = v24
                _ = v25
            } else {
                var v26 int16 = 1
                var v27 uint64 = v3
                _ = v26
                _ = v27
            }
            v28 := GenS1{F0: 255, F1: v3}
            _ = fi21
            _ = v22
            _ = v23
            _ = v28
        }
        if v4 {
            for fi29 := 0; fi29 < 18; fi29++ {
                switch (^(fi20)) {
                case -2147483648:
                    switch (^(fi29)) {
                    case 0:
                        v12 = v1
                    case 9223372036854775807:
                        var v30 uint64 = (v3 ^ 989)
                        _ = v30
                    }
                default:
                    v31 := []int{fi29, (fi20 & fi20), fi29}
                    _ = v31
                }
                var v32 int16 = -32768
                _ = fi29
                _ = v32
            }
            v4 = v12
            v2 = v2
        } else {
            v4 = (155 == 2147483647)
            v33 := []int{fi20}
            if (!false) {
                if (-2.25 > (((float64(0.0) + (3.14159 * -2.25)) * 1e10) * (((1e10 + (0.0 / -2.25)) / 3.14159) - 1.5))) {
                    var v34 uint32 = 0
                    var v35 int64 = 0
                    _ = v34
                    _ = v35
                } else {
                    for fi36 := 0; fi36 < 14; fi36++ {
                        if len(v33) > 0 {
                            v33[0] = -9223372036854775808
                        }
                        _ = fi36
                    }
                }
            } else {
                if len(v33) > 0 {
                    v33[0] = fi20
                }
            }
            _ = v33
        }
        for fi37 := 0; fi37 < 14; fi37++ {
            var v38 float64 = -2.25
            v39 := []int{}
            if (255 > 42) {
                if ((^(v3)) >= v3) {
                    var v40 uint32 = 252
                    _ = len(v39)
                    _ = v40
                }
                switch "" {
                case "a":
                    if ((1e10 > 1e10) && (42 == v3)) {
                        if len(v39) >= 2 {
                            v39 = v39[1:len(v39)]
                        }
                        v2 = (v2 / (((v2 * (v2 / ((v2) | 1)))) | 1))
                        v1 = (42 != 2147483647)
                        if len(v39) >= 2 {
                            v39 = v39[1:len(v39)]
                        }
                    }
                case "foo bar":
                    if v12 {
                        v41 := GenS5{F0: v2, F1: 42, F2: 4294967295}
                        var v42 int64 = 976
                        v39 = append(v39, (fi37 * (fi37 / ((fi20) | 1))))
                        _ = v41
                        _ = v42
                    }
                default:
                    fmt.Println(fi20, v3, v4)
                    switch 32767 {
                    case -9223372036854775808:
                        if len(v39) >= 2 {
                            v39 = v39[1:len(v39)]
                        }
                    default:
                        v39 = append(v39, (fi20 >> ((42) & 63)))
                    }
                }
            }
            v39 = append(v39, (fi37 ^ fi20))
            _ = fi37
            _ = v38
            _ = v39
        }
        _ = fi20
    }
    for fi43 := 0; fi43 < 5; fi43++ {
        v44 := []int{(fi43 ^ 1), fi43, fi43}
        for fi45 := 0; fi45 < 18; fi45++ {
            v46 := GenS4{F0: false}
            var v47 bool = ((!(0 >= 2147483648)) && ((^(v2)) < (v2 & v2)))
            v48 := GenS5{F0: (v2 % (((v2 ^ v2)) | 1)), F1: 1, F2: 65535}
            if (2147483647 != -319) {
                _ = len(v44)
                var v49 uint = 4294967295
                v50 := GenS3{F0: "a"}
                _ = v49
                _ = v50
            }
            v3 = uint64(v2)
            _ = fi45
            _ = v46
            _ = v47
            _ = v48
        }
        fmt.Println(v3)
        _ = fi43
        _ = v44
    }
    var v51 bool = v12
    var v52 uint8 = v2
    if (!v12) {
        var v53 int8 = 42
        _ = v53
    }
    for fi54 := 0; fi54 < 19; fi54++ {
        v55 := []int{(fi54 * 569)}
        if (-2.25 < (1.5 - ((3.14159 - 1e10) + (((1.5 * float64((3.14159 + -2.25))) / ((0.1 + (3.14159 * 0.0)) - 1e10)) / (1.5 - (float64(((((1.5 + (1e10 / 1e10)) - (-2.25 + 1.5)) / 0.1) * 1e10)) * 1.5)))))) {
            for fi56 := 0; fi56 < 2; fi56++ {
                fmt.Println(v4, fi54)
                _ = len(v55)
                var v57 bool = (((0.1 + float32(1e10)) >= 1.5) || ((1 < 32) && v12))
                _ = fi56
                _ = v57
            }
        } else {
            var v58 string = ("" + "hello")
            _ = v58
        }
        _ = fi54
        _ = v55
    }
    if (true || ((950 >= 255) && false)) {
        if v51 {
            if true {
                v59 := []int{128, 0}
                var v60 int32 = int32(v3)
                _ = v59
                _ = v60
            }
        } else {
            for fi61 := 0; fi61 < 10; fi61++ {
                if (v1 || (1.5 == ((((3.14159 + 0.0) - ((1e10 - ((0.1 - 0.1) / 1.5)) * -2.25)) - 1.5) * 1.5))) {
                    v62 := []int{(fi61 ^ 9223372036854775807)}
                    var v63 int = fi61
                    v12 = (65535 > 0)
                    _ = v62
                    _ = v63
                } else {
                    var v64 int = (fi61 << ((genFn0()) & 63))
                    _ = v64
                }
                _ = fi61
            }
        }
    } else {
        for fi65 := 0; fi65 < 12; fi65++ {
            v66 := []int{}
            v67 := GenS3{F0: "hello"}
            _ = fi65
            _ = v66
            _ = v67
        }
        switch "" {
        case "\n\t":
            var v68 int64 = 42
            _ = v68
        default:
            v69 := []int{-831, 986}
            _ = v69
        }
    }
    for fi70 := 0; fi70 < 6; fi70++ {
        v52 -= 127
        switch (fi70 | (-(fi70))) {
        case 4294967295:
            for fi71 := 0; fi71 < 4; fi71++ {
                switch (fi71 / ((-908) | 1)) {
                case 364:
                    v72 := []int{-1, (fi71 - (^(fi71))), (fi70 ^ fi70), 42}
                    _ = v72
                case 42:
                    v73 := []int{fi70, fi70, 127, (fi71 & fi70)}
                    _ = v73
                }
                if true {
                    if (!(65535 <= 32767)) {
                        fmt.Println(v52, v1)
                        var v74 uint32 = 4294967295
                        v75 := GenS6{F0: 1, F1: (v3 + (v3 % ((v3) | 1)))}
                        v76 := GenS7{F0: (fi71 - fi71), F1: v52}
                        v1 = true
                        _ = v74
                        _ = v75
                        _ = v76
                    }
                    v77 := GenS1{F0: 18446744073709551615, F1: v3}
                    _ = v77
                } else {
                    v78 := GenS2{F0: 603}
                    v3 -= (v3 & (v3 / (((v3 ^ 1)) | 1)))
                    var v79 float64 = 0.1
                    _ = v78
                    _ = v79
                }
                v12 = v51
                if false {
                    v12 = v12
                }
                _ = fi71
            }
        }
        v80 := GenS3{F0: "foo bar"}
        _ = fi70
        _ = v80
    }
    var v81 uint = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v51
    _ = v52
    _ = v81
}
