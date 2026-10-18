package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8; F1 uint }
type GenS4 struct { F0 uint16; F1 int8 }
type GenS5 struct { F0 uint; F1 int32; F2 uint8 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint16; F1 uint16 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 int64; F1 int; F2 int16 }
type GenS10 struct { F0 int8; F1 string; F2 uint }
type GenS11 struct { F0 uint64; F1 int8 }
type GenS12 struct { F0 string }
func genFn0() uint8 {
    v7 := GenS1{F0: 42, F1: 42}
    if (42 >= 192) {
        var v8 int16 = 42
        if true {
            var v9 uint16 = 942
            v10 := GenS2{F0: 2}
            _ = v10
            _ = v9
        }
        var v11 uint16 = 0
        _ = v11
        _ = v8
    } else {
        v12 := []int{}
        v12 = append(v12, -958)
        _ = v12
    }
    v13 := GenS1{F0: -128, F1: 2147483647}
    v14 := GenS2{F0: 0}
    switch 4294967295 {
    case -1:
        var v15 int64 = 42
        _ = v15
    }
    _ = v13
    _ = v14
    _ = v7
    return 235
}
func genFn1(p0 uint32, p1 int64) bool {
    if (42 == int32(p1)) {
        v16 := GenS2{F0: 85}
        p1 -= 9223372036854775807
        p1 -= 0
        _ = v16
    } else {
        v17 := GenS3{F0: 86, F1: 2}
        for fi18 := 0; fi18 < 13; fi18++ {
            v19 := []int{(fi18 << ((268) & 63))}
            _ = fi18
            _ = v19
        }
        var v20 uint8 = genFn0()
        _ = v17
        _ = v20
    }
    v21 := GenS4{F0: 285, F1: 85}
    var v22 int = 606
    switch "a" {
    case "foo bar":
        var v23 uint32 = 2147483648
        _ = v23
    case "a":
        switch 0 {
        case 127:
            if true {
                for fi24 := 0; fi24 < 11; fi24++ {
                    for fi25 := 0; fi25 < 13; fi25++ {
                        var v26 uint16 = 65535
                        _ = fi25
                        _ = v26
                    }
                    var v27 uint8 = 1
                    _ = fi24
                    _ = v27
                }
                var v28 int16 = -1
                var v29 uint16 = 255
                _ = v28
                _ = v29
            } else {
                switch v22 {
                case 32767:
                    v30 := GenS5{F0: 1, F1: 1, F2: 64}
                    _ = v30
                case -9223372036854775808:
                    var v31 int8 = 86
                    _ = v31
                }
                v32 := []int{v22, (v22 * v22)}
                v22 = v22
                _ = v32
            }
        default:
            if (104 <= p0) {
                if (2147483647 <= -1) {
                    var v33 uint64 = 2
                    switch (v22 + (v22 | v22)) {
                    case -686:
                        for fi34 := 0; fi34 < 4; fi34++ {
                            var v35 uint = 636
                            v36 := []int{}
                            switch fi34 {
                            case 0:
                                v36 = append(v36, (fi34 - (v22 - v22)))
                            default:
                                v37 := GenS6{F0: 18446744073709551615}
                                _ = v37
                            }
                            _ = fi34
                            _ = v35
                            _ = v36
                        }
                    }
                    var v38 string = "a"
                    _ = v33
                    _ = v38
                }
                v39 := []int{v22, 128}
                if ((!false) || (381 >= 0)) {
                    var v40 bool = false
                    _ = v40
                }
                _ = v39
            } else {
                v41 := GenS5{F0: 42, F1: 0, F2: 1}
                _ = v41
            }
        }
    case "\n\t":
        var v42 int16 = 712
        _ = v42
    }
    _ = p0
    _ = p1
    _ = v21
    _ = v22
    return false
}

func main() {
    var v1 uint64 = 180
    v2 := make(map[int]bool)
    var v3 int64 = -1
    var v4 int64 = v3
    var v5 int32 = -360
    var v6 uint64 = v1
    v43 := GenS7{F0: 2, F1: 1}
    v1 = 167
    for fi44 := 0; fi44 < 15; fi44++ {
        var v45 int64 = v4
        v46 := GenS8{F0: 42}
        _ = fi44
        _ = v45
        _ = v46
    }
    switch (("" + "foo bar") + "foo bar") {
    case "\n\t":
        v5 -= v5
    default:
        v47 := []int{4294967295, -395, 32767, -2147483648}
        _ = v47
    }
    v48 := GenS9{F0: v3, F1: 83, F2: 32767}
    if ("" > "\n\t") {
        v3 = (v3 / (((^(v3))) | 1))
    }
    if (true || true) {
        v6 += (v1 << ((880) & 63))
        v1 = (v6 % ((v1) | 1))
    }
    v49 := GenS10{F0: int8(v5), F1: "hello", F2: 427}
    v3 = (v4 << ((719) & 63))
    switch 127 {
    case 4294967295:
        var v50 uint32 = 4294967295
        _ = v50
    case 603:
        v51 := GenS8{F0: 65535}
        _ = v51
    case -721:
        v6 += v6
    default:
        var v52 int16 = 1000
        _ = v52
    }
    if false {
        var v53 uint16 = 65535
        v54 := GenS11{F0: v1, F1: -1}
        _ = v53
        _ = v54
    }
    var v55 uint64 = v1
    switch "hello" {
    case "\n\t":
        v56 := []int{-901}
        _ = v56
    default:
        for fi57 := 0; fi57 < 18; fi57++ {
            var v58 string = "a"
            if true {
                for fi59 := 0; fi59 < 7; fi59++ {
                    var v60 float32 = (1.5 / 1e10)
                    v61 := []int{int(v5)}
                    v55 -= v1
                    _ = fi59
                    _ = v60
                    _ = v61
                }
            } else {
                var v62 int64 = 1
                v1 = v55
                v58 = v58
                _ = v62
            }
            _ = fi57
            _ = v58
        }
    }
    var v63 uint64 = v1
    var v64 int = 4294967295
    if ((((1.5 + float64(0.0)) - 0.0) + 0.1) <= (1.5 / -2.25)) {
        for fi65 := 0; fi65 < 18; fi65++ {
            switch (fi65 & (fi65 ^ -9223372036854775808)) {
            case -2147483648:
                fmt.Println(v5, v63)
            case 32767:
                v66 := GenS7{F0: 65535, F1: 42}
                _ = v66
            default:
                switch "" {
                case "":
                    v67 := GenS3{F0: -1, F1: uint(v63)}
                    _ = v67
                case "hello":
                    fmt.Println(v64, v4, v3)
                }
            }
            _ = fi65
        }
        var v68 uint64 = v55
        if ((((("" + "") + "") + "") + ("foo bar" + "\n\t")) > ("hello" + "a")) {
            var v69 int64 = v4
            for fi70 := 0; fi70 < 2; fi70++ {
                v71 := GenS5{F0: 42, F1: (^(v5)), F2: uint8(v6)}
                var v72 bool = (42 > 151)
                if v72 {
                    for fi73 := 0; fi73 < 16; fi73++ {
                        var v74 int8 = 0
                        _ = fi73
                        _ = v74
                    }
                    if false {
                        for fi75 := 0; fi75 < 5; fi75++ {
                            var v76 bool = false
                            v77 := GenS12{F0: "\n\t"}
                            if v76 {
                                var v78 int64 = (v4 + (v4 >> ((65535) & 63)))
                                v4 += (v4 >> ((255) & 63))
                                _ = v78
                            }
                            _ = fi75
                            _ = v76
                            _ = v77
                        }
                        if (((^(v55)) <= v55) && (!((v5 % ((v5) | 1)) == (v5 / ((716) | 1))))) {
                            v72 = v72
                            v64 = fi70
                            v64 = v64
                        } else {
                            for fi79 := 0; fi79 < 7; fi79++ {
                                var v80 uint16 = 74
                                _ = len(v2)
                                v4 += (v69 % ((v69) | 1))
                                _ = fi79
                                _ = v80
                            }
                        }
                    }
                    for fi81 := 0; fi81 < 14; fi81++ {
                        v63 = (v1 << ((65535) & 63))
                        switch 127 {
                        case 4294967295:
                            if (1.5 >= 0.0) {
                                v82 := GenS1{F0: 0, F1: (v5 & (-(v5)))}
                                v83 := v2[(^(fi70))]
                                _ = v82
                                _ = v83
                            }
                        case 2147483647:
                            var v84 int16 = -1
                            _ = v84
                        case -131:
                            v85 := []int{fi81}
                            _ = v85
                        default:
                            fmt.Println(fi70, v3)
                        }
                        _ = fi81
                    }
                } else {
                    v86 := v2[-468]
                    _ = v86
                }
                _ = fi70
                _ = v71
                _ = v72
            }
            _ = v69
        }
        _ = v68
    } else {
        fmt.Println(v6)
    }
    v1 += v55
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("v6:", v6)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v48
    _ = v49
    _ = v5
    _ = v55
    _ = v6
    _ = v63
    _ = v64
}
