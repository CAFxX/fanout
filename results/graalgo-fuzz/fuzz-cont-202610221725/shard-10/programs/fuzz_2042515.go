package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 int32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8; F1 int16; F2 uint }
type GenS4 struct { F0 string; F1 uint64 }
type GenS5 struct { F0 int; F1 uint32; F2 int8 }
type GenS6 struct { F0 bool; F1 string }
func genFn0(p0 uint8, p1 int8, p2 int8) uint32 {
    var v5 uint32 = 1
    fmt.Println(p0, p1, p1)
    var v6 uint32 = 259
    p1 = -64
    var v7 string = "a"
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 42
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]bool)
    var v3 int32 = 1
    var v4 uint8 = 1
    for fi8 := 0; fi8 < 15; fi8++ {
        for fi9 := 0; fi9 < 15; fi9++ {
            v10 := GenS1{F0: 42, F1: -463, F2: v3}
            _ = fi9
            _ = v10
        }
        if true {
            var v11 int = (fi8 >> ((18446744073709551615) & 63))
            for fi12 := 0; fi12 < 6; fi12++ {
                var v13 float32 = 3.14159
                if true {
                    v14 := GenS1{F0: -1, F1: 1, F2: 938}
                    switch ("a" + "hello") {
                    case "hello":
                        var v15 int64 = 1
                        _ = v15
                    }
                    var v16 string = "\n\t"
                    _ = v14
                    _ = v16
                }
                _ = fi12
                _ = v13
            }
            for fi17 := 0; fi17 < 10; fi17++ {
                v4 = (v4 >> ((4294967295) & 63))
                _ = fi17
            }
            _ = v11
        }
        _ = fi8
    }
    if true {
        v18 := v2[genFn0((v4 + (v4 << ((0) & 63))), -13, -8)]
        var v19 int16 = 1
        var v20 uint16 = 0
        _ = v18
        _ = v19
        _ = v20
    } else {
        switch "" {
        case "":
            for fi21 := 0; fi21 < 9; fi21++ {
                for fi22 := 0; fi22 < 18; fi22++ {
                    v23 := GenS2{F0: 269}
                    _ = fi22
                    _ = v23
                }
                var v24 uint16 = 255
                v3 = (v3 - v3)
                _ = fi21
                _ = v24
            }
        case "foo bar":
            switch ("" + "") {
            case "hello":
                var v25 uint64 = 555
                _ = v25
            case "foo bar":
                fmt.Println(v4)
            case "":
                for fi26 := 0; fi26 < 6; fi26++ {
                    var v27 string = "foo bar"
                    _ = fi26
                    _ = v27
                }
            default:
                switch "" {
                case "hello":
                    v4 = 4
                }
            }
        default:
            v28 := GenS2{F0: 354}
            _ = v28
        }
        switch 42 {
        case 32767:
            for fi29 := 0; fi29 < 1; fi29++ {
                var v30 uint32 = 1
                fmt.Println(v30, v3, v4)
                _ = fi29
                _ = v30
            }
        case -591:
            v3 += (v3 ^ (v3 % ((v3) | 1)))
        }
        var v31 uint8 = (^(v4))
        _ = v31
    }
    fmt.Println(v4, v4)
    switch -1 {
    case 848:
        if (((1.5 - 0.1) / 3.14159) == float64(1e10)) {
            v2[2] = true
            var v32 float64 = 1e10
            _ = v32
        } else {
            for fi33 := 0; fi33 < 12; fi33++ {
                switch fi33 {
                case -9223372036854775808:
                    _ = len(v2)
                default:
                    fmt.Println(fi33)
                }
                _ = fi33
            }
            v34 := GenS3{F0: 1, F1: 0, F2: 468}
            if false {
                var v35 int8 = -69
                _ = v35
            } else {
                for fi36 := 0; fi36 < 8; fi36++ {
                    var v37 float32 = (3.14159 / 0.1)
                    if (!(false && false)) {
                        fmt.Println(v37, v4, fi36)
                        var v38 bool = false
                        _ = v38
                    } else {
                        v39 := v1[fi36]
                        for fi40 := 0; fi40 < 7; fi40++ {
                            v41 := GenS4{F0: (("a" + "a") + "hello"), F1: 122}
                            var v42 int = (fi36 + (fi40 / ((fi36) | 1)))
                            var v43 int = v42
                            _ = fi40
                            _ = v41
                            _ = v42
                            _ = v43
                        }
                        for fi44 := 0; fi44 < 12; fi44++ {
                            var v45 bool = true
                            var v46 float64 = 3.14159
                            v47 := []int{}
                            v3 = -623
                            _ = fi44
                            _ = v45
                            _ = v46
                            _ = v47
                        }
                        _ = v39
                    }
                    for fi48 := 0; fi48 < 19; fi48++ {
                        if (true || false) {
                            var v49 int32 = (v3 | (v3 & v3))
                            _ = v49
                        }
                        fmt.Println(v3, fi36, v4)
                        v50 := GenS1{F0: -1, F1: -470, F2: 1}
                        _ = fi48
                        _ = v50
                    }
                    _ = fi36
                    _ = v37
                }
                var v51 int8 = 100
                for fi52 := 0; fi52 < 7; fi52++ {
                    switch fi52 {
                    case -2147483648:
                        v53 := GenS1{F0: 0, F1: 1000, F2: (v3 * (v3 << ((65535) & 63)))}
                        _ = v53
                    case -828:
                        for fi54 := 0; fi54 < 17; fi54++ {
                            v4 -= (^(v4))
                            var v55 int8 = (v51 % ((v51) | 1))
                            _ = fi54
                            _ = v55
                        }
                    }
                    _ = fi52
                }
                _ = v51
            }
            _ = v34
        }
    case 32768:
        v3 += (v3 << ((610) & 63))
    case 545:
        v4 = v4
    }
    switch 128 {
    case -463:
        var v56 int64 = 1
        _ = v56
    case 872:
        fmt.Println(v3, v3)
    }
    var v57 uint8 = v4
    v58 := GenS5{F0: -607, F1: 0, F2: 1}
    for fi59 := 0; fi59 < 4; fi59++ {
        if (!true) {
            for fi60 := 0; fi60 < 2; fi60++ {
                var v61 string = "\n\t"
                fmt.Println(v57, v61)
                _ = fi60
                _ = v61
            }
        } else {
            v62 := GenS5{F0: 32767, F1: 2147483647, F2: -35}
            v3 = 42
            _ = v62
        }
        for fi63 := 0; fi63 < 7; fi63++ {
            v64 := GenS1{F0: int64(v3), F1: -319, F2: (-(v3))}
            if (18446744073709551615 != 255) {
                v57 = (v57 + v4)
                var v65 int64 = -1
                if true {
                    var v66 uint16 = 527
                    _ = v66
                } else {
                    var v67 bool = false
                    var v68 string = (("" + "foo bar") + "a")
                    _ = v67
                    _ = v68
                }
                var v69 uint = 0
                v70 := GenS3{F0: 1, F1: 32767, F2: 18446744073709551615}
                _ = v65
                _ = v69
                _ = v70
            }
            _ = fi63
            _ = v64
        }
        _ = fi59
    }
    var v71 uint16 = 66
    var v72 uint8 = v4
    fmt.Println(v72, v72, v72)
    for fi73 := 0; fi73 < 12; fi73++ {
        for fi74 := 0; fi74 < 9; fi74++ {
            v4 = (v72 + (^(v72)))
            _ = fi74
        }
        _ = fi73
    }
    var v75 uint32 = genFn0(117, -1, 100)
    v76 := GenS6{F0: false, F1: ("" + (("" + "foo bar") + "foo bar"))}
    if true {
        for fi77 := 0; fi77 < 10; fi77++ {
            var v78 uint = 0
            if (994 == 42) {
                for fi79 := 0; fi79 < 6; fi79++ {
                    v72 -= (v72 >> ((2) & 63))
                    fmt.Println(v3, v57)
                    v2[224] = true
                    _ = fi79
                }
                v4 -= 255
            } else {
                fmt.Println(v3, v4)
                v75 = 562
                v71 -= (v71 * 42)
            }
            _ = fi77
            _ = v78
        }
        v4 = (v4 >> ((107) & 63))
        v80 := GenS1{F0: int64(v75), F1: 0, F2: v3}
        _ = v80
    }
    var v81 int64 = 802
    for fi82 := 0; fi82 < 6; fi82++ {
        v57 -= (v72 ^ (v4 + (^(v72))))
        _ = fi82
    }
    switch -997 {
    case 9223372036854775807:
        var v83 int32 = 221
        _ = v83
    case 127:
        var v84 uint64 = 2
        _ = v84
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v57:", v57)
    fmt.Println("v71:", v71)
    fmt.Println("v72:", v72)
    fmt.Println("v75:", v75)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v57
    _ = v58
    _ = v71
    _ = v72
    _ = v75
    _ = v76
    _ = v81
}
