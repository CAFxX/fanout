package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32; F1 string; F2 bool }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int8; F1 bool }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint64; F1 uint16; F2 int16 }
func genFn0() uint8 {
    var v6 int32 = 42
    v7 := []int{-564, 926, 128, -794}
    switch 2147483647 {
    case 42:
        v8 := GenS1{F0: 42}
        _ = v8
    case 128:
        var v9 string = ""
        _ = v9
    case -797:
        v6 -= v6
    default:
        v10 := GenS1{F0: 353}
        _ = v10
    }
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    var v11 uint32 = 65535
    _ = v11
    _ = v6
    _ = v7
    return 128
}
func genFn1(p0 int64) uint8 {
    for fi12 := 0; fi12 < 20; fi12++ {
        v13 := GenS2{F0: 2, F1: ("hello" + ("a" + "foo bar")), F2: true}
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 11; fi14++ {
        var v15 uint32 = 2147483647
        _ = fi14
        _ = v15
    }
    var v16 uint = 42
    if (!(995 >= 0)) {
        for fi17 := 0; fi17 < 18; fi17++ {
            if true {
                for fi18 := 0; fi18 < 16; fi18++ {
                    if true {
                        p0 -= 176
                        if true {
                            v19 := GenS1{F0: 2}
                            _ = v19
                        } else {
                            var v20 uint = 18446744073709551615
                            _ = v20
                        }
                    }
                    var v21 string = "a"
                    _ = fi18
                    _ = v21
                }
            } else {
                switch ("" + "hello") {
                case "\n\t":
                    fmt.Println(fi17, v16)
                case "foo bar":
                    if (!false) {
                        var v22 int8 = 0
                        switch 128 {
                        case 505:
                            v23 := GenS1{F0: 822}
                            _ = v23
                        case -9223372036854775808:
                            if (-2.25 >= (1.5 / (float64(0.1) / -2.25))) {
                                fmt.Println(v22)
                                v24 := GenS3{F0: false}
                                _ = v24
                            } else {
                                v25 := GenS1{F0: 255}
                                v16 -= 4294967295
                                _ = v25
                            }
                        }
                        _ = v22
                    } else {
                        p0 = 1
                    }
                default:
                    v26 := GenS2{F0: 42, F1: "a", F2: ("foo bar" > ("\n\t" + "a"))}
                    _ = v26
                }
                for fi27 := 0; fi27 < 19; fi27++ {
                    p0 = (p0 ^ (p0 * (p0 / ((p0) | 1))))
                    v28 := []int{fi27, (fi27 >> ((v16) & 63)), (fi27 ^ 127), -9223372036854775808}
                    _ = fi27
                    _ = v28
                }
                for fi29 := 0; fi29 < 1; fi29++ {
                    v30 := []int{(fi17 ^ (fi29 * fi17)), 32768, fi29}
                    _ = fi29
                    _ = v30
                }
            }
            p0 -= (p0 | p0)
            _ = fi17
        }
    } else {
        var v31 int8 = 127
        _ = v31
    }
    _ = p0
    _ = v16
    return 2
}

func main() {
    var v1 int8 = 29
    var v2 uint = 0
    var v3 float32 = 0.0
    var v4 uint32 = 171
    var v5 uint64 = 42
    var v32 int32 = 2147483647
    var v33 uint16 = 32767
    if false {
        if false {
            v4 = 4294967295
        }
        switch ("\n\t" + ("foo bar" + "")) {
        case "hello":
            v34 := []int{127, 342, -2147483648}
            _ = v34
        case "":
            var v35 int = 32767
            _ = v35
        case "a":
            if true {
                if (!true) {
                    if (42 > (v1 ^ (v1 >> ((v2) & 63)))) {
                        var v36 float32 = v3
                        _ = v36
                    } else {
                        v5 = 0
                        var v37 float64 = 3.14159
                        switch ("foo bar" + "hello") {
                        case "foo bar":
                            var v38 uint8 = uint8(v5)
                            _ = v38
                        case "":
                            var v39 int = -997
                            _ = v39
                        case "a":
                            var v40 string = "foo bar"
                            _ = v40
                        }
                        _ = v37
                    }
                }
                v41 := []int{9223372036854775807, -372}
                v2 = 255
                _ = v41
            } else {
                if (v33 > v33) {
                    if false {
                        var v42 uint8 = 0
                        _ = v42
                    } else {
                        var v43 float32 = v3
                        v33 -= 0
                        for fi44 := 0; fi44 < 12; fi44++ {
                            var v45 string = "\n\t"
                            v4 = v4
                            _ = fi44
                            _ = v45
                        }
                        _ = v43
                    }
                    if true {
                        if (!(true && (!(false || true)))) {
                            fmt.Println(v2)
                            if true {
                                switch ("foo bar" + "\n\t") {
                                case "\n\t":
                                    v46 := []int{32768, 127, 128, 32768}
                                    _ = v46
                                case "":
                                    v5 = (v5 & 18446744073709551615)
                                default:
                                    v4 = (^(v4))
                                }
                                if true {
                                    for fi47 := 0; fi47 < 10; fi47++ {
                                        var v48 uint32 = 42
                                        var v49 uint64 = v5
                                        v50 := GenS1{F0: 32767}
                                        _ = fi47
                                        _ = v48
                                        _ = v49
                                        _ = v50
                                    }
                                    for fi51 := 0; fi51 < 20; fi51++ {
                                        var v52 uint32 = (^(v4))
                                        fmt.Println(fi51, v33)
                                        fmt.Println(v52)
                                        _ = fi51
                                        _ = v52
                                    }
                                } else {
                                    var v53 bool = true
                                    var v54 string = "hello"
                                    switch -1 {
                                    case 1:
                                        var v55 float64 = float64(1e10)
                                        _ = v55
                                    default:
                                        v54 = v54
                                    }
                                    _ = v53
                                    _ = v54
                                }
                            }
                            for fi56 := 0; fi56 < 11; fi56++ {
                                v57 := []int{2147483647, (fi56 >> ((65535) & 63)), int(v5), fi56}
                                var v58 uint32 = (v4 | v4)
                                v1 += v1
                                _ = fi56
                                _ = v57
                                _ = v58
                            }
                        } else {
                            var v59 float32 = 3.14159
                            v60 := GenS1{F0: 335}
                            _ = v59
                            _ = v60
                        }
                    }
                    fmt.Println(v3, v33, v3)
                }
                if false {
                    for fi61 := 0; fi61 < 14; fi61++ {
                        v4 += 1
                        fmt.Println(v3, v2)
                        _ = fi61
                    }
                }
            }
        default:
            v62 := []int{}
            _ = v62
        }
    }
    switch int(v1) {
    case -947:
        v63 := GenS4{F0: int8(v33), F1: true}
        _ = v63
    case -2147483648:
        v64 := GenS1{F0: v33}
        _ = v64
    default:
        if ((1.5 + 0.0) > -2.25) {
            v4 = 42
            switch "" {
            case "\n\t":
                v4 -= 660
            case "hello":
                for fi65 := 0; fi65 < 19; fi65++ {
                    v66 := GenS4{F0: int8(v5), F1: (((v33 - (v33 << ((0) & 63))) != v33) && false)}
                    _ = fi65
                    _ = v66
                }
            }
        }
    }
    v2 -= (v2 + v2)
    var v67 uint8 = 35
    switch "\n\t" {
    case "":
        var v68 bool = false
        _ = v68
    case "foo bar":
        var v69 uint = uint(v33)
        _ = v69
    }
    v70 := []int{746, -2147483648, -9223372036854775808, int(v4)}
    if false {
        if (((v32 >> (((v2 & uint(v4))) & 63)) >= (v32 / (((-(v32))) | 1))) || (!true)) {
            if (((1 < -868) || false) || (false && (true || (false && false)))) {
                var v71 bool = (false && true)
                for fi72 := 0; fi72 < 15; fi72++ {
                    v33 = 126
                    _ = fi72
                }
                for fi73 := 0; fi73 < 20; fi73++ {
                    if len(v70) > 0 {
                        v70[0] = fi73
                    }
                    var v74 int8 = v1
                    var v75 uint8 = genFn0()
                    _ = fi73
                    _ = v74
                    _ = v75
                }
                _ = v71
            } else {
                v2 -= (v2 ^ 337)
                switch 139 {
                case 1:
                    fmt.Println(v5, v67)
                case 128:
                    for fi76 := 0; fi76 < 7; fi76++ {
                        if len(v70) >= 2 {
                            v70 = v70[1:len(v70)]
                        }
                        switch fi76 {
                        case -2147483648:
                            var v77 int = fi76
                            _ = v77
                        default:
                            v78 := GenS2{F0: uint32(v67), F1: "a", F2: (-786 < -32768)}
                            _ = v78
                        }
                        var v79 uint = (v2 | (v2 / ((v2) | 1)))
                        _ = fi76
                        _ = v79
                    }
                case 815:
                    switch "foo bar" {
                    case "foo bar":
                        for fi80 := 0; fi80 < 19; fi80++ {
                            switch (fi80 | int(v1)) {
                            case -324:
                                v81 := GenS3{F0: (v32 > 2147483647)}
                                _ = v81
                            case 8:
                                v70 = append(v70, fi80)
                            }
                            v82 := GenS2{F0: (v4 << ((uint(v67)) & 63)), F1: ("foo bar" + "foo bar"), F2: false}
                            if true {
                                var v83 uint8 = (^(v67))
                                fmt.Println(fi80, v4)
                                _ = v83
                            } else {
                                fmt.Println(v1)
                            }
                            _ = fi80
                            _ = v82
                        }
                    case "":
                        v84 := GenS4{F0: (-(v1)), F1: (((^(v5)) >= v5) && false)}
                        _ = v84
                    }
                default:
                    for fi85 := 0; fi85 < 13; fi85++ {
                        for fi86 := 0; fi86 < 19; fi86++ {
                            if ((1.5 * 0.1) <= (1.5 + (3.14159 + 0.1))) {
                                if true {
                                    if len(v70) >= 2 {
                                        v70 = v70[1:len(v70)]
                                    }
                                } else {
                                    if (747 != uint32(v2)) {
                                        var v87 int8 = (v1 / (((v1 - (v1 ^ (v1 ^ -108)))) | 1))
                                        var v88 uint64 = (^(v5))
                                        _ = v87
                                        _ = v88
                                    } else {
                                        var v89 int32 = int32(fi85)
                                        v5 = v5
                                        switch fi86 {
                                        case 32768:
                                            v5 -= 18446744073709551615
                                        }
                                        _ = v89
                                    }
                                }
                                var v90 uint32 = (^(v4))
                                _ = v90
                            } else {
                                v91 := GenS4{F0: int8(v5), F1: (-2.25 > 0.1)}
                                _ = v91
                            }
                            v70 = append(v70, int(v2))
                            if true {
                                v33 += 32767
                                v92 := GenS3{F0: false}
                                var v93 int = (-(fi85))
                                _ = v92
                                _ = v93
                            } else {
                                if len(v70) >= 2 {
                                    v70 = v70[1:len(v70)]
                                }
                                v1 = (v1 / ((127) | 1))
                                v2 = 18446744073709551615
                            }
                            _ = fi86
                        }
                        v94 := GenS5{F0: int64(fi85)}
                        _ = fi85
                        _ = v94
                    }
                }
                switch "\n\t" {
                case "\n\t":
                    v2 += v2
                case "":
                    switch 842 {
                    case 32768:
                        v32 -= v32
                    case 2147483647:
                        for fi95 := 0; fi95 < 17; fi95++ {
                            if ((v67 << (((v2 ^ (^(v2)))) & 63)) != (^(v67))) {
                                for fi96 := 0; fi96 < 10; fi96++ {
                                    if len(v70) > 0 {
                                        v70[0] = (fi96 | fi96)
                                    }
                                    _ = fi96
                                }
                            }
                            for fi97 := 0; fi97 < 4; fi97++ {
                                v67 = (v67 / (((v67 | (v67 - 1))) | 1))
                                _ = fi97
                            }
                            _ = fi95
                        }
                    default:
                        for fi98 := 0; fi98 < 20; fi98++ {
                            if len(v70) > 0 {
                                v70[0] = (fi98 - (fi98 >> ((v2) & 63)))
                            }
                            fmt.Println(fi98, v4, v32)
                            _ = fi98
                        }
                    }
                default:
                    switch -604 {
                    case 128:
                        v70 = append(v70, -397)
                    default:
                        v99 := GenS5{F0: -459}
                        _ = v99
                    }
                }
            }
        } else {
            v32 -= v32
        }
    } else {
        if false {
            v2 = (^(v2))
            v100 := GenS2{F0: (v4 / ((v4) | 1)), F1: "a", F2: (!((4294967295 != v5) && (v32 == (v32 + (-(v32))))))}
            v101 := GenS3{F0: (false && (false || false))}
            _ = v100
            _ = v101
        }
    }
    v102 := GenS6{F0: v5, F1: v33, F2: 765}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v67:", v67)
    fmt.Println("slicelen v70:", len(v70), v70)
    _ = v1
    _ = v102
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v67
    _ = v70
}
