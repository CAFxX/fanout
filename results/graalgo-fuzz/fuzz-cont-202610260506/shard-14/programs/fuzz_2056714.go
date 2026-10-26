package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint; F1 int }
type GenS3 struct { F0 uint8; F1 int16 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int64; F1 string }
type GenS6 struct { F0 uint; F1 int16 }
type GenS7 struct { F0 bool; F1 bool }
type GenS8 struct { F0 uint32; F1 uint32; F2 uint8 }
type GenS9 struct { F0 int16; F1 int8 }
type GenS10 struct { F0 uint; F1 uint8 }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 int16; F1 bool }
type GenS13 struct { F0 int; F1 uint8 }
type GenS14 struct { F0 int32; F1 uint32; F2 int16 }
type GenS15 struct { F0 int16 }
type GenS16 struct { F0 uint16 }
type GenS17 struct { F0 uint16 }
type GenS18 struct { F0 int32; F1 uint8 }
type GenS19 struct { F0 uint8; F1 int }
type GenS20 struct { F0 int8 }
type GenS21 struct { F0 uint64 }
type GenS22 struct { F0 int16; F1 uint16 }
type GenS23 struct { F0 uint; F1 int8; F2 uint64 }
func genFn0(p0 uint32, p1 uint16) uint32 {
    v4 := GenS1{F0: -2147483648}
    for fi5 := 0; fi5 < 11; fi5++ {
        v6 := GenS1{F0: 1}
        switch "a" {
        case "":
            switch (("\n\t" + "foo bar") + ("a" + "foo bar")) {
            case "foo bar":
                if true {
                    for fi7 := 0; fi7 < 17; fi7++ {
                        p0 -= p0
                        p0 = p0
                        _ = fi7
                    }
                } else {
                    v8 := GenS1{F0: -2147483648}
                    _ = v8
                }
            default:
                v9 := GenS2{F0: 691, F1: (fi5 | -595)}
                _ = v9
            }
        }
        for fi10 := 0; fi10 < 11; fi10++ {
            v11 := GenS3{F0: 32, F1: 1}
            _ = fi10
            _ = v11
        }
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = p1
    _ = v4
    return (p0 - (^(p0)))
}

func main() {
    var v1 int = -2147483648
    v2 := make(map[uint32]string)
    var v3 int16 = -350
    fmt.Println(v3, v1)
    for fi12 := 0; fi12 < 13; fi12++ {
        v13 := GenS3{F0: 2, F1: v3}
        _ = fi12
        _ = v13
    }
    var v14 int8 = 5
    v15 := v2[0]
    v16 := GenS1{F0: 941}
    var v17 float32 = 1.5
    fmt.Println(v15, v1, v17)
    v18 := []int{(v1 & (-(v1))), -58, v1}
    for fi19 := 0; fi19 < 1; fi19++ {
        v15 = v15
        for fi20 := 0; fi20 < 9; fi20++ {
            if ((3.14159 / 1.0) >= 1e10) {
                v21 := GenS4{F0: 0}
                switch 1 {
                case 2147483647:
                    v15 = v15
                case 0:
                    var v22 float32 = v17
                    _ = v22
                case -2147483648:
                    var v23 int16 = 79
                    _ = v23
                default:
                    v15 = ((("" + "foo bar") + "foo bar") + v15)
                    if (!true) {
                        if (v15 != v15) {
                            var v24 float64 = float64(v17)
                            _ = v24
                        } else {
                            v18 = append(v18, v1)
                        }
                        if false {
                            var v25 uint32 = genFn0(42, 2)
                            v15 = ("" + v15)
                            if false {
                                switch ("hello" + v15) {
                                case "\n\t":
                                    var v26 int16 = (v3 ^ -862)
                                    _ = v26
                                case "":
                                    if (1 > (fi20 | 32767)) {
                                        var v27 uint32 = v25
                                        _ = v27
                                    }
                                default:
                                    v28 := GenS2{F0: 950, F1: (fi19 | fi20)}
                                    v14 += 1
                                    _ = v28
                                }
                            } else {
                                fmt.Println(fi19, v1, fi19)
                                if (!false) {
                                    v29 := GenS5{F0: -1, F1: "a"}
                                    if true {
                                        v15 = "a"
                                        v17 = (v17 * (v17 / (float32(0.1) / (v17 + ((float32((((((((1.5 - 1.5) / -2.25) * float64(0.0)) * 0.0) / 0.1) * 0.0) * float64(((v17 * (v17 + (3.14159 * v17))) * ((((v17 - (((v17 * v17) * ((((v17 * (float32((1e10 / (1.5 / ((0.0 - 3.14159) * 0.1)))) * ((v17 / (v17 - (-2.25 / 1.5))) * (((float32(0.1) + 0.0) - v17) + v17)))) / 1.0) * 0.0) / (0.0 - -2.25))) + v17)) - 1.5) + 0.0) * v17))))) - ((v17 - v17) - 0.0)) + -2.25)))))
                                    }
                                    _ = v29
                                } else {
                                    v30 := GenS2{F0: 1, F1: (v1 & (fi19 + -983))}
                                    v1 = (fi20 ^ v1)
                                    if true {
                                        v14 -= -1
                                        _ = len(v18)
                                    } else {
                                        v15 = ((v15 + "foo bar") + "")
                                        switch v15 {
                                        case "":
                                            v31 := GenS5{F0: 268, F1: ("foo bar" + (v15 + v15))}
                                            v1 = (-(fi20))
                                            if (!(false && false)) {
                                                v32 := GenS1{F0: 681}
                                                v15 = (v15 + (v15 + ("hello" + (v15 + (v15 + v15)))))
                                                if false {
                                                    v15 = v15
                                                }
                                                _ = v32
                                            } else {
                                                if false {
                                                    var v33 int16 = 1000
                                                    v33 = v33
                                                    _ = len(v18)
                                                    _ = v33
                                                } else {
                                                    v3 += 42
                                                }
                                                var v34 int = fi19
                                                if (1e10 >= (0.0 / 3.14159)) {
                                                    v1 = v34
                                                } else {
                                                    v3 = (v3 >> ((0) & 63))
                                                }
                                                if true {
                                                    v35 := GenS6{F0: 266, F1: int16(v34)}
                                                    v36 := GenS5{F0: 279, F1: (v15 + v15)}
                                                    var v37 int16 = v3
                                                    _ = v35
                                                    _ = v36
                                                    _ = v37
                                                } else {
                                                    v34 = 32768
                                                    v17 = 0.0
                                                    if (-9223372036854775808 > 4294967295) {
                                                        switch (v34 * fi19) {
                                                        case 0:
                                                            switch (fi19 - fi19) {
                                                            case 127:
                                                                v15 = v15
                                                            }
                                                        default:
                                                            switch 32768 {
                                                            case -2147483648:
                                                                if true {
                                                                    switch (v15 + (v15 + v15)) {
                                                                    case "hello":
                                                                        v38 := GenS6{F0: 1, F1: (v3 * v3)}
                                                                        _ = v38
                                                                    case "":
                                                                        v18 = append(v18, -850)
                                                                    case "foo bar":
                                                                        v39 := GenS5{F0: -102, F1: ((("hello" + v15) + v15) + ((v15 + ("\n\t" + "")) + "hello"))}
                                                                        _ = v39
                                                                    default:
                                                                        switch int(v14) {
                                                                        case -2147483648:
                                                                            if true {
                                                                                var v40 uint16 = 653
                                                                                _ = v40
                                                                            } else {
                                                                                v34 -= (v34 << ((42) & 63))
                                                                                v41 := GenS7{F0: ((v1 >> ((516) & 63)) == (-(fi19))), F1: (4294967295 == 784)}
                                                                                _ = v41
                                                                            }
                                                                        case 0:
                                                                            v34 = v34
                                                                        case 209:
                                                                            switch "\n\t" {
                                                                            case "foo bar":
                                                                                v17 += (3.14159 * v17)
                                                                            case "a":
                                                                                var v42 string = "foo bar"
                                                                                _ = v42
                                                                            default:
                                                                                switch v1 {
                                                                                case -2147483648:
                                                                                    var v43 int16 = (v3 ^ v3)
                                                                                    _ = v43
                                                                                case 32768:
                                                                                    if ("\n\t" > ("hello" + ("" + "\n\t"))) {
                                                                                        if (false || (1.5 != (1e10 * ((-2.25 / (0.1 / ((float64(((1e10 * v17) + (float32(-2.25) * v17))) - (3.14159 - (1.5 * 3.14159))) * 1e10))) + 1.5)))) {
                                                                                            var v44 string = (v15 + "hello")
                                                                                            fmt.Println(v17, v3, v25)
                                                                                            _ = v44
                                                                                        }
                                                                                    } else {
                                                                                        var v45 int = (^(fi20))
                                                                                        if ((!false) || (v14 > v14)) {
                                                                                            _ = len(v2)
                                                                                            v2[650] = (v15 + "\n\t")
                                                                                            if true {
                                                                                                if false {
                                                                                                    var v46 int64 = 9223372036854775807
                                                                                                    v47 := GenS8{F0: v25, F1: 4294967295, F2: 2}
                                                                                                    _ = v46
                                                                                                    _ = v47
                                                                                                } else {
                                                                                                    switch v15 {
                                                                                                    case "\n\t":
                                                                                                        if ((!(!((866 > 1) && (36 == (v14 ^ (^(v14))))))) || true) {
                                                                                                            if ((12 > 664) && true) {
                                                                                                                if false {
                                                                                                                    fmt.Println(v14, v3, fi19)
                                                                                                                    if (!false) {
                                                                                                                        if (v25 == (v25 / ((v25) | 1))) {
                                                                                                                            v48 := GenS7{F0: (v14 >= v14), F1: false}
                                                                                                                            switch 9223372036854775807 {
                                                                                                                            case 32768:
                                                                                                                                v15 = ((("" + v15) + ((v15 + v15) + "hello")) + v15)
                                                                                                                            default:
                                                                                                                                var v49 uint8 = 0
                                                                                                                                _ = v49
                                                                                                                            }
                                                                                                                            switch v15 {
                                                                                                                            case "":
                                                                                                                                fmt.Println(v34)
                                                                                                                            case "foo bar":
                                                                                                                                if false {
                                                                                                                                    v14 = -1
                                                                                                                                    var v50 bool = true
                                                                                                                                    _ = v50
                                                                                                                                } else {
                                                                                                                                    v45 = -620
                                                                                                                                }
                                                                                                                            default:
                                                                                                                                switch v34 {
                                                                                                                                case -9223372036854775808:
                                                                                                                                    if true {
                                                                                                                                        if true {
                                                                                                                                            var v51 float32 = ((v17 + 0.1) / (v17 + v17))
                                                                                                                                            if false {
                                                                                                                                                if (!true) {
                                                                                                                                                    if true {
                                                                                                                                                        fmt.Println(fi20, v34)
                                                                                                                                                    } else {
                                                                                                                                                        switch (v34 ^ (v34 / ((v34) | 1))) {
                                                                                                                                                        case 203:
                                                                                                                                                            v34 = (fi19 >> ((65535) & 63))
                                                                                                                                                        default:
                                                                                                                                                            v2[v25] = v15
                                                                                                                                                        }
                                                                                                                                                        v17 = float32(-2.25)
                                                                                                                                                    }
                                                                                                                                                    if (v25 > v25) {
                                                                                                                                                        var v52 int = (v34 * (fi19 << ((42) & 63)))
                                                                                                                                                        _ = v52
                                                                                                                                                    }
                                                                                                                                                } else {
                                                                                                                                                    var v53 int = v45
                                                                                                                                                    switch "a" {
                                                                                                                                                    case "foo bar":
                                                                                                                                                        switch "a" {
                                                                                                                                                        case "":
                                                                                                                                                            fmt.Println(v14, v17, v45)
                                                                                                                                                            v54 := GenS9{F0: v3, F1: (v14 + (v14 - v14))}
                                                                                                                                                            _ = v54
                                                                                                                                                        case "foo bar":
                                                                                                                                                            fmt.Println(v53)
                                                                                                                                                        }
                                                                                                                                                    case "":
                                                                                                                                                        if false {
                                                                                                                                                            switch fi20 {
                                                                                                                                                            case 2147483647:
                                                                                                                                                                var v55 int = -845
                                                                                                                                                                _ = v55
                                                                                                                                                            case 42:
                                                                                                                                                                v34 = v34
                                                                                                                                                            case -747:
                                                                                                                                                                if (false && (true || (255 != 9))) {
                                                                                                                                                                    v18 = append(v18, fi20)
                                                                                                                                                                }
                                                                                                                                                                v1 += 42
                                                                                                                                                            default:
                                                                                                                                                                var v56 uint32 = v25
                                                                                                                                                                _ = v56
                                                                                                                                                            }
                                                                                                                                                            v53 = (^(v45))
                                                                                                                                                            var v57 int8 = v14
                                                                                                                                                            v58 := v2[(^(v25))]
                                                                                                                                                            _ = v57
                                                                                                                                                            _ = v58
                                                                                                                                                        }
                                                                                                                                                    default:
                                                                                                                                                        v17 += 1e10
                                                                                                                                                    }
                                                                                                                                                    switch v15 {
                                                                                                                                                    case "a":
                                                                                                                                                        var v59 int16 = v3
                                                                                                                                                        _ = v59
                                                                                                                                                    case "":
                                                                                                                                                        var v60 uint8 = 0
                                                                                                                                                        _ = v60
                                                                                                                                                    case "foo bar":
                                                                                                                                                        fmt.Println(v45, v17, v45)
                                                                                                                                                    default:
                                                                                                                                                        if false {
                                                                                                                                                            if true {
                                                                                                                                                                v61 := GenS4{F0: 724}
                                                                                                                                                                if true {
                                                                                                                                                                    v62 := GenS10{F0: 252, F1: 0}
                                                                                                                                                                    if ((v25 ^ genFn0(genFn0(42, 32767), 255)) <= (^(v25))) {
                                                                                                                                                                        var v63 int = 127
                                                                                                                                                                        v25 += (v25 << ((255) & 63))
                                                                                                                                                                        _ = v63
                                                                                                                                                                    }
                                                                                                                                                                    v45 = 2147483647
                                                                                                                                                                    _ = v62
                                                                                                                                                                }
                                                                                                                                                                v64 := GenS8{F0: (v25 * (v25 * (^(v25)))), F1: (v25 % (((^(v25))) | 1)), F2: 85}
                                                                                                                                                                _ = v61
                                                                                                                                                                _ = v64
                                                                                                                                                            } else {
                                                                                                                                                                if (true && false) {
                                                                                                                                                                    v15 = "hello"
                                                                                                                                                                }
                                                                                                                                                            }
                                                                                                                                                        }
                                                                                                                                                    }
                                                                                                                                                    _ = v53
                                                                                                                                                }
                                                                                                                                            }
                                                                                                                                            var v65 int32 = 42
                                                                                                                                            _ = v51
                                                                                                                                            _ = v65
                                                                                                                                        } else {
                                                                                                                                            switch ("\n\t" + ((v15 + (v15 + (v15 + v15))) + v15)) {
                                                                                                                                            case "\n\t":
                                                                                                                                                if true {
                                                                                                                                                    v66 := GenS10{F0: 1, F1: 119}
                                                                                                                                                    var v67 uint16 = 2
                                                                                                                                                    var v68 float32 = v17
                                                                                                                                                    _ = v66
                                                                                                                                                    _ = v67
                                                                                                                                                    _ = v68
                                                                                                                                                } else {
                                                                                                                                                    var v69 bool = (128 > 2)
                                                                                                                                                    if v69 {
                                                                                                                                                        v17 -= -2.25
                                                                                                                                                        var v70 int = v34
                                                                                                                                                        v25 = 0
                                                                                                                                                        _ = v70
                                                                                                                                                    } else {
                                                                                                                                                        v71 := GenS7{F0: (false && v69), F1: (int64(v1) <= 756)}
                                                                                                                                                        switch (fi20 & (-(v45))) {
                                                                                                                                                        case 128:
                                                                                                                                                            v34 = fi20
                                                                                                                                                        case 487:
                                                                                                                                                            v3 -= v3
                                                                                                                                                        default:
                                                                                                                                                            if (!((828 >= 0) && v69)) {
                                                                                                                                                                var v72 int = 128
                                                                                                                                                                _ = v72
                                                                                                                                                            } else {
                                                                                                                                                                fmt.Println(v14, fi19, v14)
                                                                                                                                                                var v73 float32 = -2.25
                                                                                                                                                                switch (v34 >> ((120) & 63)) {
                                                                                                                                                                case -915:
                                                                                                                                                                    if false {
                                                                                                                                                                        var v74 uint = 2
                                                                                                                                                                        _ = v74
                                                                                                                                                                    }
                                                                                                                                                                    fmt.Println(v14, v1)
                                                                                                                                                                    v75 := GenS11{F0: (v25 << ((139) & 63))}
                                                                                                                                                                    _ = v75
                                                                                                                                                                case 265:
                                                                                                                                                                    var v76 bool = true
                                                                                                                                                                    _ = v76
                                                                                                                                                                }
                                                                                                                                                                _ = v73
                                                                                                                                                            }
                                                                                                                                                        }
                                                                                                                                                        fmt.Println(fi20)
                                                                                                                                                        _ = v71
                                                                                                                                                    }
                                                                                                                                                    _ = v69
                                                                                                                                                }
                                                                                                                                            case "":
                                                                                                                                                if ((true || ((v1 & (fi20 - (^(v34)))) > fi20)) || ((((v17 / v17) * 1.5) * (v17 * (1e10 * 3.14159))) <= v17)) {
                                                                                                                                                    var v77 int64 = -1
                                                                                                                                                    if false {
                                                                                                                                                        var v78 int16 = 451
                                                                                                                                                        v79 := GenS12{F0: int16(fi20), F1: true}
                                                                                                                                                        _ = v78
                                                                                                                                                        _ = v79
                                                                                                                                                    }
                                                                                                                                                    var v80 uint8 = 42
                                                                                                                                                    _ = v77
                                                                                                                                                    _ = v80
                                                                                                                                                } else {
                                                                                                                                                    v81 := GenS5{F0: 1, F1: v15}
                                                                                                                                                    switch "a" {
                                                                                                                                                    case "":
                                                                                                                                                        if len(v18) > 0 {
                                                                                                                                                            v18[0] = (fi20 / ((2147483647) | 1))
                                                                                                                                                        }
                                                                                                                                                    }
                                                                                                                                                    _ = v81
                                                                                                                                                }
                                                                                                                                            case "a":
                                                                                                                                                v25 = v25
                                                                                                                                            }
                                                                                                                                            v82 := v2[v25]
                                                                                                                                            if (1.5 > 0.0) {
                                                                                                                                                v83 := GenS13{F0: fi19, F1: 123}
                                                                                                                                                switch "\n\t" {
                                                                                                                                                case "hello":
                                                                                                                                                    if len(v18) >= 2 {
                                                                                                                                                        v18 = v18[1:len(v18)]
                                                                                                                                                    }
                                                                                                                                                case "\n\t":
                                                                                                                                                    v45 += v1
                                                                                                                                                case "":
                                                                                                                                                    var v84 uint16 = 975
                                                                                                                                                    _ = v84
                                                                                                                                                }
                                                                                                                                                if len(v18) >= 2 {
                                                                                                                                                    v18 = v18[1:len(v18)]
                                                                                                                                                }
                                                                                                                                                _ = v83
                                                                                                                                            } else {
                                                                                                                                                v85 := GenS10{F0: 255, F1: 36}
                                                                                                                                                fmt.Println(fi19, v1)
                                                                                                                                                _ = v85
                                                                                                                                            }
                                                                                                                                            _ = v82
                                                                                                                                        }
                                                                                                                                    } else {
                                                                                                                                        fmt.Println(v1)
                                                                                                                                        v34 = 2147483647
                                                                                                                                        v86 := v2[(v25 ^ v25)]
                                                                                                                                        _ = v86
                                                                                                                                    }
                                                                                                                                case 1:
                                                                                                                                    v18 = append(v18, fi20)
                                                                                                                                case 783:
                                                                                                                                    v34 -= v1
                                                                                                                                default:
                                                                                                                                    v2[(v25 / (((v25 / ((2147483648) | 1))) | 1))] = "a"
                                                                                                                                }
                                                                                                                            }
                                                                                                                            _ = v48
                                                                                                                        } else {
                                                                                                                            v87 := GenS14{F0: int32(v14), F1: (v25 ^ (v25 * v25)), F2: v3}
                                                                                                                            _ = v87
                                                                                                                        }
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    v88 := GenS9{F0: (v3 * -957), F1: 42}
                                                                                                                    v3 = v3
                                                                                                                    _ = v88
                                                                                                                }
                                                                                                                var v89 uint = 4294967295
                                                                                                                _ = v89
                                                                                                            }
                                                                                                        }
                                                                                                    case "hello":
                                                                                                        v17 = v17
                                                                                                    case "a":
                                                                                                        var v90 int = (-(v1))
                                                                                                        _ = v90
                                                                                                    default:
                                                                                                        var v91 uint = 255
                                                                                                        _ = v91
                                                                                                    }
                                                                                                }
                                                                                                v15 = (v15 + v15)
                                                                                                if (!(42 < 427)) {
                                                                                                    fmt.Println(v45, v15)
                                                                                                    var v92 uint32 = v25
                                                                                                    _ = v92
                                                                                                } else {
                                                                                                    if (((((v17 + float32(1.5)) + ((((0.0 / 1.0) * v17) - float32((3.14159 / (1e10 - 3.14159)))) * v17)) - 0.0) + 1e10) > ((float32((-2.25 + 0.0)) / (((1.5 - 0.1) + float32(-2.25)) - v17)) + v17)) {
                                                                                                        v34 = 580
                                                                                                    } else {
                                                                                                        var v93 uint16 = 42
                                                                                                        v93 -= (v93 + (v93 + v93))
                                                                                                        if true {
                                                                                                            v45 = (v1 + 32768)
                                                                                                            if true {
                                                                                                                var v94 bool = false
                                                                                                                var v95 int = (fi20 << ((42) & 63))
                                                                                                                if v94 {
                                                                                                                    var v96 uint = 1
                                                                                                                    v3 = v3
                                                                                                                    _ = v96
                                                                                                                }
                                                                                                                _ = v94
                                                                                                                _ = v95
                                                                                                            }
                                                                                                        } else {
                                                                                                            v25 = 65535
                                                                                                        }
                                                                                                        v97 := GenS15{F0: v3}
                                                                                                        _ = v93
                                                                                                        _ = v97
                                                                                                    }
                                                                                                    v98 := GenS16{F0: 0}
                                                                                                    _ = v98
                                                                                                }
                                                                                                v3 -= v3
                                                                                                var v99 float32 = v17
                                                                                                _ = v99
                                                                                            } else {
                                                                                                if false {
                                                                                                    var v100 float32 = (float32(((3.14159 / 1.5) / 1.5)) - (v17 / (1e10 / v17)))
                                                                                                    _ = v100
                                                                                                }
                                                                                                var v101 float32 = v17
                                                                                                if (!(255 >= 714)) {
                                                                                                    if true {
                                                                                                        v102 := GenS17{F0: 398}
                                                                                                        v103 := v2[v25]
                                                                                                        _ = len(v2)
                                                                                                        _ = v102
                                                                                                        _ = v103
                                                                                                    }
                                                                                                } else {
                                                                                                    var v104 bool = (!(v25 != v25))
                                                                                                    switch (v15 + ("hello" + v15)) {
                                                                                                    case "\n\t":
                                                                                                        v14 += (v14 & v14)
                                                                                                    }
                                                                                                    fmt.Println(v3, fi20)
                                                                                                    _ = v104
                                                                                                }
                                                                                                _ = v101
                                                                                            }
                                                                                            var v105 int32 = 42
                                                                                            v18 = append(v18, v1)
                                                                                            switch v15 {
                                                                                            case "a":
                                                                                                var v106 int64 = 1
                                                                                                _ = v106
                                                                                            default:
                                                                                                v34 += (-(v1))
                                                                                            }
                                                                                            if (2 < 32767) {
                                                                                                fmt.Println(v15, v105, v25)
                                                                                            }
                                                                                            _ = v105
                                                                                        }
                                                                                        _ = v45
                                                                                    }
                                                                                default:
                                                                                    v2[v25] = "\n\t"
                                                                                    v107 := v2[(v25 & v25)]
                                                                                    _ = len(v2)
                                                                                    _ = v107
                                                                                }
                                                                            }
                                                                        default:
                                                                            var v108 string = (v15 + (v15 + v15))
                                                                            _ = v108
                                                                        }
                                                                    }
                                                                    v18 = append(v18, (fi20 << ((2) & 63)))
                                                                    var v109 int16 = -1
                                                                    _ = v109
                                                                } else {
                                                                    v110 := v2[0]
                                                                    _ = v110
                                                                }
                                                                if false {
                                                                    v34 = (fi20 - v1)
                                                                } else {
                                                                    if ((!(0 > 65535)) || true) {
                                                                        var v111 uint16 = 2
                                                                        delete(v2, v25)
                                                                        var v112 string = "hello"
                                                                        var v113 uint = 18446744073709551615
                                                                        _ = v111
                                                                        _ = v112
                                                                        _ = v113
                                                                    }
                                                                    switch "hello" {
                                                                    case "a":
                                                                        switch ("foo bar" + "foo bar") {
                                                                        case "\n\t":
                                                                            if (!false) {
                                                                                if false {
                                                                                    v114 := GenS17{F0: 214}
                                                                                    var v115 float32 = v17
                                                                                    v18 = append(v18, fi19)
                                                                                    _ = v114
                                                                                    _ = v115
                                                                                } else {
                                                                                    v34 = (v34 - v34)
                                                                                    switch int(v25) {
                                                                                    case 127:
                                                                                        switch (v1 + 0) {
                                                                                        case 9223372036854775807:
                                                                                            v15 = ""
                                                                                        case 4294967295:
                                                                                            var v116 uint16 = 0
                                                                                            _ = v116
                                                                                        case -9223372036854775808:
                                                                                            v14 = v14
                                                                                        }
                                                                                    case -877:
                                                                                        var v117 int8 = (v14 >> ((2) & 63))
                                                                                        _ = v117
                                                                                    }
                                                                                }
                                                                                v118 := GenS13{F0: (-(fi20)), F1: 15}
                                                                                _ = v118
                                                                            } else {
                                                                                fmt.Println(v25)
                                                                                if (!(9223372036854775807 > 1)) {
                                                                                    var v119 int8 = 100
                                                                                    v14 = v119
                                                                                    _ = v119
                                                                                } else {
                                                                                    var v120 int16 = (v3 | v3)
                                                                                    _ = v120
                                                                                }
                                                                            }
                                                                        default:
                                                                            var v121 float32 = v17
                                                                            _ = v121
                                                                        }
                                                                    case "\n\t":
                                                                        switch "hello" {
                                                                        case "hello":
                                                                            if true {
                                                                                v25 = v25
                                                                                v122 := v2[(^(v25))]
                                                                                _ = v122
                                                                            } else {
                                                                                v1 = fi20
                                                                                var v123 int32 = -4
                                                                                fmt.Println(v17, fi19, v25)
                                                                                _ = v123
                                                                            }
                                                                            var v124 int32 = 1
                                                                            v125 := GenS11{F0: 4294967295}
                                                                            _ = v124
                                                                            _ = v125
                                                                        case "foo bar":
                                                                            if false {
                                                                                v3 = v3
                                                                                var v126 int = (v1 >> ((18446744073709551615) & 63))
                                                                                _ = v126
                                                                            } else {
                                                                                v25 = 65535
                                                                                fmt.Println(v3, v25)
                                                                            }
                                                                        }
                                                                    }
                                                                    delete(v2, (v25 ^ (^(v25))))
                                                                }
                                                                v127 := GenS18{F0: 1, F1: 128}
                                                                _ = v127
                                                            case 9223372036854775807:
                                                                _ = len(v18)
                                                            }
                                                        }
                                                        switch (fi19 & -452) {
                                                        case 9223372036854775807:
                                                            v128 := GenS19{F0: 2, F1: (fi20 | fi20)}
                                                            _ = v128
                                                        case 1:
                                                            v129 := GenS6{F0: 163, F1: v3}
                                                            _ = v129
                                                        default:
                                                            var v130 float64 = float64(((0.0 * v17) / v17))
                                                            _ = v130
                                                        }
                                                        var v131 float32 = float32(float64((float32(1.5) * -2.25)))
                                                        _ = v131
                                                    } else {
                                                        var v132 int16 = v3
                                                        v34 = v34
                                                        if (v17 == (v17 - (((1e10 / (v17 / 3.14159)) / (0.0 * v17)) * float32((((0.1 - float64((v17 - v17))) / -2.25) * (-2.25 - 3.14159)))))) {
                                                            v17 = v17
                                                            v1 += v1
                                                        }
                                                        _ = v132
                                                    }
                                                }
                                                _ = v34
                                            }
                                            _ = v31
                                        case "a":
                                            switch fi19 {
                                            case 128:
                                                if (1 == 4294967295) {
                                                    v1 -= 127
                                                    fmt.Println(v1, v15)
                                                }
                                            case 9223372036854775807:
                                                switch "hello" {
                                                case "foo bar":
                                                    v25 -= v25
                                                case "\n\t":
                                                    v17 = v17
                                                }
                                            case -1:
                                                v2[2147483648] = "hello"
                                            }
                                        case "\n\t":
                                            switch v15 {
                                            case "a":
                                                var v133 uint = 255
                                                _ = v133
                                            case "\n\t":
                                                switch (v1 & v1) {
                                                case 9223372036854775807:
                                                    if (1.5 != 3.14159) {
                                                        var v134 uint32 = v25
                                                        v135 := GenS15{F0: v3}
                                                        v136 := GenS19{F0: 168, F1: -422}
                                                        v17 -= float32(1.5)
                                                        v15 = (v15 + v15)
                                                        var v137 int8 = (v14 & (-(v14)))
                                                        _ = v134
                                                        _ = v135
                                                        _ = v136
                                                        _ = v137
                                                    }
                                                case -9223372036854775808:
                                                    v1 = (fi19 & (v1 >> ((18446744073709551615) & 63)))
                                                case 1:
                                                    var v138 int = 42
                                                    _ = v138
                                                default:
                                                    switch v15 {
                                                    case "a":
                                                        v139 := GenS4{F0: 42}
                                                        _ = v139
                                                    case "\n\t":
                                                        var v140 uint64 = 135
                                                        _ = v140
                                                    case "":
                                                        var v141 uint64 = 4294967295
                                                        _ = v141
                                                    default:
                                                        v18 = append(v18, v1)
                                                    }
                                                    v142 := GenS10{F0: 0, F1: 141}
                                                    _ = v142
                                                }
                                            case "":
                                                var v143 bool = false
                                                _ = v143
                                            default:
                                                if true {
                                                    switch "foo bar" {
                                                    case "":
                                                        _ = len(v2)
                                                    }
                                                } else {
                                                    var v144 int = (fi19 << ((18446744073709551615) & 63))
                                                    _ = v144
                                                }
                                            }
                                        default:
                                            var v145 int8 = v14
                                            _ = v145
                                        }
                                        fmt.Println(v14, v3)
                                        if false {
                                            if (-1 < -1) {
                                                var v146 float32 = float32(float64((v17 * ((((((((v17 / v17) - (((((v17 * (((3.14159 * 0.0) + (v17 - (0.0 * (v17 * (v17 * 1.5))))) + (v17 + (v17 / v17)))) - v17) * 0.0) / v17) - float32(float64((0.0 * v17))))) + v17) + v17) * (((float32((1.5 - float64(((v17 * (float32(0.1) + (v17 - 1.5))) / 1.5)))) - 3.14159) / float32(1e10)) / (float32(-2.25) - v17))) / 1.0) * (3.14159 * 1e10)) - v17))))
                                                _ = v146
                                            } else {
                                                var v147 uint32 = genFn0(0, 65535)
                                                if ((v14 ^ 48) != v14) {
                                                    fmt.Println(v14, v15)
                                                    v148 := GenS20{F0: -78}
                                                    _ = v148
                                                } else {
                                                    fmt.Println(fi20, v1, fi20)
                                                    switch (fi20 & 20) {
                                                    case 0:
                                                        if (((3.14159 - (1.5 + 3.14159)) == 3.14159) || ((v147 + v147) == (v147 % (((v147 + v25)) | 1)))) {
                                                            switch -250 {
                                                            case 529:
                                                                var v149 float32 = ((3.14159 / (1.5 / (v17 - (v17 * v17)))) - 0.0)
                                                                _ = v149
                                                            default:
                                                                fmt.Println(v25)
                                                            }
                                                        } else {
                                                            if (-120 >= (fi19 % ((fi19) | 1))) {
                                                                fmt.Println(v14, v17)
                                                                if ((!false) && false) {
                                                                    var v150 uint = 0
                                                                    _ = v150
                                                                } else {
                                                                    if (42 < 370) {
                                                                        v3 = (v3 * (v3 - v3))
                                                                    }
                                                                    v151 := GenS16{F0: 0}
                                                                    var v152 uint = 18446744073709551615
                                                                    _ = v151
                                                                    _ = v152
                                                                }
                                                                fmt.Println(v17, v25)
                                                            } else {
                                                                v14 += v14
                                                                v1 -= v1
                                                            }
                                                            v1 = fi19
                                                            v3 = v3
                                                        }
                                                    }
                                                    fmt.Println(v3, fi19)
                                                }
                                                _ = len(v18)
                                                _ = v147
                                            }
                                            if (-516 <= -798) {
                                                if (v15 < "") {
                                                    v25 = v25
                                                    v3 += v3
                                                } else {
                                                    v15 = (v15 + v15)
                                                    var v153 int8 = (v14 << ((544) & 63))
                                                    fmt.Println(v14, fi19)
                                                    v1 += (v1 + (fi20 + v1))
                                                    _ = v153
                                                }
                                            } else {
                                                if true {
                                                    var v154 uint32 = genFn0((v25 % ((genFn0((v25 / ((v25) | 1)), 1)) | 1)), 652)
                                                    v154 += v25
                                                    _ = v154
                                                } else {
                                                    v155 := v2[genFn0((^(v25)), 2)]
                                                    _ = v155
                                                }
                                            }
                                        } else {
                                            var v156 int32 = -2147483648
                                            v17 -= ((3.14159 / v17) * v17)
                                            if len(v18) >= 2 {
                                                v18 = v18[1:len(v18)]
                                            }
                                            v156 = v156
                                            v17 -= (((float32(-2.25) / (-2.25 - v17)) + 1.5) + v17)
                                            _ = v156
                                        }
                                        v25 = v25
                                    }
                                    _ = v30
                                }
                                if true {
                                    var v157 uint = 208
                                    _ = v157
                                }
                            }
                            _ = v25
                        } else {
                            if len(v18) > 0 {
                                v18[0] = (fi20 >> ((864) & 63))
                            }
                            var v158 string = v15
                            v1 = 4294967295
                            _ = v158
                        }
                        switch "hello" {
                        case "":
                            v14 = (^(v14))
                        case "\n\t":
                            var v159 int64 = 9223372036854775807
                            _ = v159
                        case "foo bar":
                            v14 += 127
                        }
                    } else {
                        if (2 <= 65535) {
                            var v160 bool = false
                            if v160 {
                                v3 -= -32768
                                if len(v18) >= 2 {
                                    v18 = v18[1:len(v18)]
                                }
                                fmt.Println(fi20)
                                switch -650 {
                                case 127:
                                    if (true || (!v160)) {
                                        var v161 float32 = (((v17 - v17) + (v17 + v17)) * v17)
                                        var v162 uint64 = 317
                                        _ = v161
                                        _ = v162
                                    }
                                case 1:
                                    v15 = ""
                                }
                            }
                            fmt.Println(v160, v17, fi19)
                            var v163 uint32 = 4294967295
                            if v160 {
                                if v160 {
                                    v3 = (v3 ^ v3)
                                    switch fi20 {
                                    case 663:
                                        var v164 bool = (2 == 46)
                                        _ = v164
                                    }
                                    if v160 {
                                        if v160 {
                                            if v160 {
                                                v1 += 32768
                                                var v165 bool = v160
                                                v14 = v14
                                                var v166 float64 = 1e10
                                                var v167 int32 = -137
                                                var v168 bool = (!v165)
                                                if v168 {
                                                    v169 := GenS1{F0: (v167 + -2147483648)}
                                                    _ = v169
                                                } else {
                                                    v18 = append(v18, 4294967295)
                                                }
                                                if (false && (!((false || v165) || false))) {
                                                    if v168 {
                                                        var v170 uint32 = (v163 + v163)
                                                        v163 -= 2147483648
                                                        _ = len(v2)
                                                        _ = v170
                                                    } else {
                                                        v1 -= -9223372036854775808
                                                        var v171 bool = false
                                                        _ = v171
                                                    }
                                                    if false {
                                                        v17 += (1.5 - v17)
                                                        v3 -= (v3 & v3)
                                                    }
                                                    _ = len(v2)
                                                }
                                                var v172 uint = 23
                                                switch "hello" {
                                                case "hello":
                                                    v14 += v14
                                                case "":
                                                    if len(v18) > 0 {
                                                        v18[0] = (fi20 & fi19)
                                                    }
                                                case "foo bar":
                                                    switch v15 {
                                                    case "":
                                                        var v173 uint = 2
                                                        _ = v173
                                                    default:
                                                        v3 = (v3 + v3)
                                                    }
                                                default:
                                                    if (!(!((v160 || true) || false))) {
                                                        switch fi19 {
                                                        case 2147483647:
                                                            if v168 {
                                                                fmt.Println(v15)
                                                                var v174 int8 = (v14 - (^(v14)))
                                                                var v175 uint64 = 2
                                                                _ = v174
                                                                _ = v175
                                                            } else {
                                                                v172 = v172
                                                            }
                                                        default:
                                                            var v176 uint64 = 1
                                                            _ = v176
                                                        }
                                                        v17 += float32(((v166 - ((v166 / (v166 * v166)) * v166)) * v166))
                                                        switch ((("hello" + ((v15 + v15) + v15)) + v15) + v15) {
                                                        case "\n\t":
                                                            var v177 int = (fi19 >> ((0) & 63))
                                                            _ = v177
                                                        case "":
                                                            if len(v18) >= 2 {
                                                                v18 = v18[1:len(v18)]
                                                            }
                                                        case "foo bar":
                                                            fmt.Println(v165, v17)
                                                        }
                                                        switch v15 {
                                                        case "\n\t":
                                                            v1 += 4294967295
                                                        default:
                                                            fmt.Println(v14, v1, v167)
                                                        }
                                                    } else {
                                                        var v178 uint8 = 255
                                                        var v179 uint8 = v178
                                                        _ = v178
                                                        _ = v179
                                                    }
                                                }
                                                _ = v165
                                                _ = v166
                                                _ = v167
                                                _ = v168
                                                _ = v172
                                            } else {
                                                var v180 int = fi20
                                                var v181 int8 = (v14 % (((-(v14))) | 1))
                                                _ = v180
                                                _ = v181
                                            }
                                        }
                                        if (4294967295 <= 162) {
                                            v182 := GenS3{F0: 128, F1: 0}
                                            _ = v182
                                        }
                                    } else {
                                        switch "\n\t" {
                                        case "a":
                                            var v183 uint32 = v163
                                            _ = v183
                                        case "foo bar":
                                            if false {
                                                if (394 <= 1) {
                                                    var v184 int16 = (v3 >> ((108) & 63))
                                                    v18 = append(v18, 912)
                                                    var v185 float64 = (((float64(((v17 + 1.5) + 1.5)) / 0.1) - 1e10) * -2.25)
                                                    switch (fi20 % (((fi19 >> ((566) & 63))) | 1)) {
                                                    case -990:
                                                        _ = len(v2)
                                                    }
                                                    _ = v184
                                                    _ = v185
                                                }
                                                if len(v18) >= 2 {
                                                    v18 = v18[1:len(v18)]
                                                }
                                                if v160 {
                                                    fmt.Println(v17, v17)
                                                    fmt.Println(v17, v163)
                                                } else {
                                                    v186 := GenS19{F0: 217, F1: int(v163)}
                                                    v160 = v160
                                                    fmt.Println(v163)
                                                    _ = v186
                                                }
                                            }
                                        case "":
                                            switch 9223372036854775807 {
                                            case 32768:
                                                fmt.Println(v1, fi20)
                                            case -9223372036854775808:
                                                v17 = v17
                                            default:
                                                v15 = v15
                                            }
                                        }
                                    }
                                } else {
                                    v14 += v14
                                    _ = len(v18)
                                }
                                fmt.Println(v1)
                                if v160 {
                                    fmt.Println(v1)
                                    if v160 {
                                        if (!(1 <= (v3 + 0))) {
                                            v187 := GenS21{F0: 2}
                                            v14 = v14
                                            _ = v187
                                        }
                                    } else {
                                        v160 = (v15 < v15)
                                        if (0 > 140) {
                                            switch (fi20 * fi19) {
                                            case 0:
                                                fmt.Println(v3)
                                            case 9223372036854775807:
                                                fmt.Println(v14, fi19, v14)
                                            }
                                        }
                                    }
                                }
                                fmt.Println(fi20, fi19)
                            }
                            _ = v160
                            _ = v163
                        }
                        if false {
                            if false {
                                if (-1 >= (^(v14))) {
                                    v14 = (v14 >> ((981) & 63))
                                    v188 := GenS22{F0: 1, F1: 1}
                                    if (!(false || true)) {
                                        var v189 uint32 = 2
                                        switch v1 {
                                        case 504:
                                            if (1 >= int64(v189)) {
                                                if false {
                                                    var v190 int16 = (v3 / (((v3 + (v3 ^ 1000))) | 1))
                                                    var v191 int8 = v14
                                                    v192 := GenS22{F0: (v3 / ((-32768) | 1)), F1: 42}
                                                    _ = v190
                                                    _ = v191
                                                    _ = v192
                                                }
                                                var v193 int16 = v3
                                                fmt.Println(v1, v15, v1)
                                                _ = v193
                                            } else {
                                                var v194 int64 = 0
                                                _ = v194
                                            }
                                        case 2147483647:
                                            var v195 float32 = ((v17 / 1.0) - v17)
                                            _ = v195
                                        case -515:
                                            v14 = (v14 / ((v14) | 1))
                                        default:
                                            v15 = v15
                                        }
                                        v3 = (v3 ^ (-(v3)))
                                        if false {
                                            v17 -= v17
                                            switch v15 {
                                            case "\n\t":
                                                switch v15 {
                                                case "hello":
                                                    var v196 uint = 2
                                                    _ = v196
                                                case "":
                                                    var v197 int16 = (v3 << ((65535) & 63))
                                                    var v198 int32 = 42
                                                    v199 := GenS1{F0: -2147483648}
                                                    _ = v197
                                                    _ = v198
                                                    _ = v199
                                                default:
                                                    switch (-(fi19)) {
                                                    case 1:
                                                        if true {
                                                            if (1 > int64(v14)) {
                                                                if len(v18) > 0 {
                                                                    v18[0] = 42
                                                                }
                                                            } else {
                                                                var v200 uint64 = 97
                                                                if false {
                                                                    v3 = v3
                                                                    v3 -= (v3 >> ((1) & 63))
                                                                    v201 := GenS23{F0: 1, F1: 127, F2: 0}
                                                                    if ((-376 >= 272) && (100 == 1)) {
                                                                        if len(v18) > 0 {
                                                                            v18[0] = fi19
                                                                        }
                                                                    }
                                                                    _ = v201
                                                                } else {
                                                                    if true {
                                                                        fmt.Println(v3, fi20)
                                                                        fmt.Println(v3, v15)
                                                                    } else {
                                                                        v17 -= ((v17 + (v17 - v17)) / v17)
                                                                        var v202 float64 = (-2.25 + float64(float32(float64(0.1))))
                                                                        _ = v202
                                                                    }
                                                                }
                                                                _ = v200
                                                            }
                                                        }
                                                    case 128:
                                                        v203 := GenS13{F0: (fi20 ^ fi20), F1: 0}
                                                        _ = v203
                                                    default:
                                                        fmt.Println(v15, v14)
                                                    }
                                                }
                                            case "foo bar":
                                                if (true && true) {
                                                    v17 -= 3.14159
                                                }
                                            case "a":
                                                _ = len(v2)
                                            default:
                                                if true {
                                                    if true {
                                                        switch ("\n\t" + (v15 + (((("hello" + v15) + v15) + v15) + v15))) {
                                                        case "a":
                                                            switch (v15 + ((((("foo bar" + v15) + v15) + v15) + v15) + v15)) {
                                                            case "hello":
                                                                fmt.Println(v3, v14, v189)
                                                            case "":
                                                                var v204 uint32 = (v189 >> ((4294967295) & 63))
                                                                var v205 int = (^(fi20))
                                                                _ = v204
                                                                _ = v205
                                                            case "foo bar":
                                                                var v206 uint8 = 156
                                                                _ = v206
                                                            default:
                                                                fmt.Println(v3, fi19)
                                                            }
                                                        default:
                                                            var v207 int32 = 661
                                                            _ = v207
                                                        }
                                                    }
                                                    if len(v18) > 0 {
                                                        v18[0] = (fi20 | fi19)
                                                    }
                                                    var v208 uint64 = 76
                                                    _ = v208
                                                }
                                            }
                                        }
                                        if true {
                                            var v209 float64 = -2.25
                                            _ = v209
                                        } else {
                                            if true {
                                                var v210 int = (fi20 + fi20)
                                                var v211 uint8 = 0
                                                _ = v210
                                                _ = v211
                                            } else {
                                                if len(v18) > 0 {
                                                    v18[0] = fi20
                                                }
                                                var v212 int16 = int16(v189)
                                                _ = v212
                                            }
                                            fmt.Println(v1, v15)
                                            v213 := v2[v189]
                                            var v214 float32 = v17
                                            switch -90 {
                                            case -2147483648:
                                                var v215 int8 = (v14 & (v14 >> ((255) & 63)))
                                                _ = v215
                                            case 42:
                                                _ = len(v2)
                                            case 32767:
                                                var v216 float32 = (float32((0.0 * 1.5)) * 0.1)
                                                _ = v216
                                            }
                                            _ = v213
                                            _ = v214
                                        }
                                        v217 := GenS12{F0: (v3 % (((v3 << ((1) & 63))) | 1)), F1: (!true)}
                                        _ = v189
                                        _ = v217
                                    }
                                    _ = v188
                                } else {
                                    v218 := GenS23{F0: 2, F1: (^(v14)), F2: 18446744073709551615}
                                    _ = len(v2)
                                    _ = v218
                                }
                                v219 := v2[122]
                                _ = v219
                            } else {
                                v17 = (((float32(1.5) / -2.25) * v17) - v17)
                            }
                            var v220 uint16 = 2
                            v15 = v15
                            _ = v220
                        }
                    }
                    _ = len(v18)
                }
                if len(v18) >= 2 {
                    v18 = v18[1:len(v18)]
                }
                _ = v21
            }
            v17 -= (v17 * (v17 * (v17 + v17)))
            _ = fi20
        }
        v1 -= 621
        _ = fi19
    }
    if false {
        delete(v2, 65535)
        fmt.Println(v3, v17, v17)
    } else {
        _ = len(v18)
        for fi221 := 0; fi221 < 1; fi221++ {
            var v222 uint32 = 4294967295
            if true {
                var v223 int8 = 1
                _ = v223
            } else {
                for fi224 := 0; fi224 < 6; fi224++ {
                    var v225 uint8 = 42
                    var v226 string = "foo bar"
                    var v227 uint8 = 0
                    _ = fi224
                    _ = v225
                    _ = v226
                    _ = v227
                }
                v228 := GenS21{F0: 18446744073709551615}
                _ = v228
            }
            _ = fi221
            _ = v222
        }
    }
    var v229 uint64 = 2
    fmt.Println(v17, v1, v229)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v229:", v229)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v229
    _ = v3
}
