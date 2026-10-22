package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 int16; F1 int8 }
type GenS4 struct { F0 string; F1 string }
type GenS5 struct { F0 string }
type GenS6 struct { F0 string; F1 uint8; F2 uint32 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint8; F1 int8 }
type GenS9 struct { F0 int32; F1 int; F2 uint64 }
func genFn0(p0 uint16, p1 int32) bool {
    for fi6 := 0; fi6 < 8; fi6++ {
        v7 := GenS1{F0: 0}
        if (!("" >= "\n\t")) {
            v8 := []int{-2147483648, (fi6 % ((fi6) | 1)), fi6}
            var v9 int = (^(fi6))
            for fi10 := 0; fi10 < 15; fi10++ {
                p0 += p0
                _ = fi10
            }
            _ = v8
            _ = v9
        } else {
            var v11 string = ""
            _ = v11
        }
        _ = fi6
        _ = v7
    }
    switch "hello" {
    case "a":
        p1 = p1
    default:
        var v12 int16 = 821
        _ = v12
    }
    for fi13 := 0; fi13 < 17; fi13++ {
        var v14 bool = (816 != 4294967295)
        _ = fi13
        _ = v14
    }
    _ = p0
    _ = p1
    return (42 >= 42)
}
func genFn1(p0 int64, p1 uint16) uint32 {
    v15 := GenS2{F0: 254, F1: 1000}
    for fi16 := 0; fi16 < 5; fi16++ {
        for fi17 := 0; fi17 < 7; fi17++ {
            p1 -= 479
            switch ((("\n\t" + "foo bar") + "hello") + ("a" + ("\n\t" + "\n\t"))) {
            case "foo bar":
                fmt.Println(p0, p0)
            case "":
                if false {
                    var v18 bool = (42 == 42)
                    if v18 {
                        var v19 int = -1
                        var v20 float64 = 0.1
                        fmt.Println(v19, v19, fi16)
                        _ = v19
                        _ = v20
                    }
                    var v21 int16 = -1
                    _ = v18
                    _ = v21
                } else {
                    switch -9223372036854775808 {
                    case 2147483647:
                        p0 += (p0 % (((^(p0))) | 1))
                        v22 := GenS3{F0: -41, F1: 0}
                        _ = v22
                    }
                }
            case "\n\t":
                var v23 float32 = (1.5 - ((0.0 - (1.5 / 0.1)) - ((3.14159 + (3.14159 - (0.1 - 1e10))) * float32(1.5))))
                _ = v23
            default:
                var v24 int = fi17
                _ = v24
            }
            if (float32(1.5) >= float32(((3.14159 / (1e10 * float64((-2.25 * 3.14159)))) + 1.5))) {
                fmt.Println(p1)
                v25 := []int{9223372036854775807, 1, (fi17 >> ((2) & 63))}
                v26 := GenS2{F0: 128, F1: 0}
                _ = v25
                _ = v26
            } else {
                var v27 bool = false
                _ = v27
            }
            _ = fi17
        }
        if (849 > 4294967295) {
            p0 = -250
        } else {
            v28 := []int{714, fi16, (fi16 ^ fi16)}
            switch "foo bar" {
            case "\n\t":
                if (p0 >= p0) {
                    var v29 uint = 866
                    _ = v29
                }
            }
            _ = v28
        }
        v30 := GenS4{F0: "a", F1: ""}
        _ = fi16
        _ = v30
    }
    switch 1 {
    case -1:
        switch 128 {
        case 2147483647:
            var v31 bool = false
            _ = v31
        default:
            for fi32 := 0; fi32 < 6; fi32++ {
                if true {
                    p1 -= 2
                }
                p0 -= p0
                _ = fi32
            }
        }
    case -9223372036854775808:
        var v33 float64 = (3.14159 / 1e10)
        _ = v33
    case 32768:
        var v34 int8 = 1
        _ = v34
    }
    var v35 string = "\n\t"
    if false {
        for fi36 := 0; fi36 < 6; fi36++ {
            if (!(42 < 2147483647)) {
                v37 := GenS1{F0: p0}
                _ = v37
            }
            fmt.Println(v35)
            var v38 float64 = float64(0.0)
            _ = fi36
            _ = v38
        }
    }
    _ = p0
    _ = p1
    _ = v15
    _ = v35
    return 677
}

func main() {
    var v1 bool = false
    var v2 float32 = (0.1 / float32(-2.25))
    v3 := make(map[uint32]string)
    var v4 string = "foo bar"
    var v5 uint = 861
    for fi39 := 0; fi39 < 17; fi39++ {
        var v40 string = "hello"
        for fi41 := 0; fi41 < 5; fi41++ {
            v4 = v40
            var v42 int = (^(fi41))
            var v43 float32 = 1.5
            _ = fi41
            _ = v42
            _ = v43
        }
        if (v1 && true) {
            var v44 uint64 = 984
            _ = v44
        } else {
            v1 = (v2 != (v2 - ((float32(0.0) + 1.5) - 1.5)))
            v1 = (0 <= (fi39 << ((v5) & 63)))
        }
        _ = fi39
        _ = v40
    }
    v45 := GenS4{F0: (v4 + v4), F1: v4}
    fmt.Println(v5, v2, v2)
    v1 = v1
    for fi46 := 0; fi46 < 9; fi46++ {
        if (v1 && true) {
            v47 := GenS5{F0: "a"}
            v48 := GenS4{F0: (("a" + "") + v4), F1: ("a" + v4)}
            _ = v47
            _ = v48
        } else {
            var v49 int32 = 2147483647
            for fi50 := 0; fi50 < 18; fi50++ {
                var v51 int64 = 0
                fmt.Println(v4, fi50)
                v52 := GenS2{F0: 42, F1: 0}
                v53 := v3[4294967295]
                _ = fi50
                _ = v51
                _ = v52
                _ = v53
            }
            v54 := []int{(fi46 << ((v5) & 63)), 2147483647}
            _ = v49
            _ = v54
        }
        v5 = (v5 / (((^(v5))) | 1))
        switch fi46 {
        case 127:
            var v55 uint = v5
            _ = v55
        case -2147483648:
            v4 = v4
        }
        _ = fi46
    }
    for fi56 := 0; fi56 < 3; fi56++ {
        v57 := GenS3{F0: 0, F1: -128}
        for fi58 := 0; fi58 < 17; fi58++ {
            v5 = (v5 & v5)
            fmt.Println(fi56)
            switch (fi56 % (((fi56 - fi58)) | 1)) {
            case -638:
                var v59 uint32 = 2147483647
                _ = v59
            case 9223372036854775807:
                v60 := GenS6{F0: (("foo bar" + v4) + "foo bar"), F1: 97, F2: 1}
                _ = v60
            default:
                v5 = (v5 ^ 18446744073709551615)
            }
            _ = fi58
        }
        v4 = (("foo bar" + (((v4 + v4) + "") + ("\n\t" + (v4 + ("hello" + ""))))) + (v4 + "foo bar"))
        _ = fi56
        _ = v57
    }
    if false {
        var v61 uint64 = 0
        v62 := []int{}
        _ = v61
        _ = v62
    } else {
        v63 := GenS6{F0: v4, F1: 75, F2: 1}
        _ = v63
    }
    var v64 uint16 = 414
    for fi65 := 0; fi65 < 9; fi65++ {
        var v66 uint32 = 0
        var v67 float32 = ((v2 - v2) * 3.14159)
        if ((2 != uint64(v5)) || (!false)) {
            if (v1 && false) {
                switch "" {
                case "a":
                    for fi68 := 0; fi68 < 8; fi68++ {
                        v64 -= (v64 * (^(v64)))
                        v69 := v3[uint32(v5)]
                        _ = fi68
                        _ = v69
                    }
                }
                v4 = (v4 + v4)
            } else {
                delete(v3, v66)
                var v70 string = "\n\t"
                for fi71 := 0; fi71 < 5; fi71++ {
                    if ((-2.25 * 3.14159) >= v67) {
                        if (!(true || v1)) {
                            switch 32768 {
                            case -66:
                                if (-172 < 1) {
                                    if v1 {
                                        var v72 uint16 = (v64 ^ 667)
                                        var v73 uint64 = 4294967295
                                        v2 += (((float32(1e10) * v2) + ((v67 * v67) * 0.1)) - 0.0)
                                        _ = v72
                                        _ = v73
                                    }
                                    v74 := GenS7{F0: 691}
                                    _ = v74
                                } else {
                                    var v75 bool = (v66 == v66)
                                    if (v75 || false) {
                                        v76 := GenS8{F0: 128, F1: 0}
                                        v77 := GenS2{F0: 127, F1: 828}
                                        var v78 int32 = 2147483647
                                        _ = v76
                                        _ = v77
                                        _ = v78
                                    }
                                    fmt.Println(v1, v67, v75)
                                    _ = v75
                                }
                            case 0:
                                v79 := GenS7{F0: 18446744073709551615}
                                _ = v79
                            default:
                                v64 = v64
                            }
                            if (!(!(!v1))) {
                                v66 -= (v66 % ((v66) | 1))
                            } else {
                                if v1 {
                                    v67 -= (v2 * v2)
                                    delete(v3, (v66 * (v66 - (v66 + (v66 - v66)))))
                                    v80 := []int{(fi71 / ((-384) | 1)), 9223372036854775807}
                                    _ = v80
                                } else {
                                    fmt.Println(v5)
                                    var v81 uint64 = 18446744073709551615
                                    v1 = true
                                    _ = v81
                                }
                                v4 = v4
                                var v82 int32 = 1
                                _ = v82
                            }
                        }
                        fmt.Println(v67, v1)
                        fmt.Println(v67)
                    }
                    v70 = ((((v4 + (v70 + v4)) + v4) + v4) + v70)
                    var v83 uint8 = 129
                    _ = fi71
                    _ = v83
                }
                _ = v70
            }
            var v84 string = v4
            var v85 uint = (v5 << (((v5 % (((v5 % (((v5 + v5)) | 1))) | 1))) & 63))
            _ = v84
            _ = v85
        }
        _ = fi65
        _ = v66
        _ = v67
    }
    v86 := GenS4{F0: v4, F1: v4}
    if ((-1 == 2147483647) && false) {
        v87 := GenS2{F0: 181, F1: 1000}
        v88 := v3[42]
        var v89 uint64 = 424
        _ = v87
        _ = v88
        _ = v89
    } else {
        v90 := GenS4{F0: (((v4 + v4) + (v4 + ((((v4 + v4) + "a") + (("hello" + ((v4 + "\n\t") + v4)) + v4)) + ((("hello" + (v4 + v4)) + (((v4 + ("hello" + ("foo bar" + v4))) + "foo bar") + v4)) + v4)))) + v4), F1: ""}
        switch -516 {
        case 9223372036854775807:
            v4 = ""
        case 0:
            v91 := GenS8{F0: 42, F1: int8(v64)}
            _ = v91
        case 773:
            v92 := GenS3{F0: 1, F1: -85}
            _ = v92
        default:
            v93 := GenS9{F0: 2147483647, F1: 32767, F2: 18446744073709551615}
            _ = v93
        }
        _ = v90
    }
    var v94 int64 = 131
    var v95 float64 = 3.14159
    var v96 uint64 = 0
    fmt.Println(v2)
    for fi97 := 0; fi97 < 3; fi97++ {
        var v98 bool = (true && v1)
        _ = fi97
        _ = v98
    }
    v4 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v64:", v64)
    fmt.Println("v94:", v94)
    fmt.Println("v95:", v95)
    fmt.Println("v96:", v96)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v64
    _ = v86
    _ = v94
    _ = v95
    _ = v96
}
