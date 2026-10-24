package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 uint16; F1 int64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint; F1 int }
type GenS5 struct { F0 uint64; F1 int32; F2 int64 }
type GenS6 struct { F0 int32; F1 int64 }
type GenS7 struct { F0 string; F1 string; F2 uint32 }
type GenS8 struct { F0 int8; F1 uint; F2 int }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 string; F1 uint64; F2 uint64 }
func genFn0(p0 int, p1 int64) uint16 {
    p1 -= p1
    if ((65535 <= 2147483648) || true) {
        var v7 bool = true
        _ = v7
    } else {
        for fi8 := 0; fi8 < 16; fi8++ {
            fmt.Println(fi8, p1)
            _ = fi8
        }
        for fi9 := 0; fi9 < 17; fi9++ {
            var v10 int16 = 0
            switch "hello" {
            case "a":
                v11 := GenS1{F0: 42, F1: 2}
                _ = v11
            case "foo bar":
                for fi12 := 0; fi12 < 15; fi12++ {
                    var v13 string = "hello"
                    fmt.Println(p0, v10)
                    switch (^(fi9)) {
                    case -1:
                        v14 := GenS1{F0: 100, F1: 292}
                        _ = v14
                    case 1:
                        v15 := []int{(fi9 & (p0 << ((2) & 63))), (fi9 >> ((685) & 63))}
                        _ = v15
                    case 2:
                        if true {
                            var v16 float32 = float32(0.1)
                            v17 := []int{p0, 878, (-(fi9)), 604}
                            v10 -= (v10 + 32767)
                            _ = v16
                            _ = v17
                        }
                    }
                    _ = fi12
                    _ = v13
                }
            case "\n\t":
                var v18 uint64 = 276
                _ = v18
            default:
                p1 += p1
            }
            _ = fi9
            _ = v10
        }
        if true {
            var v19 uint64 = 2
            _ = v19
        }
    }
    if (-608 != 0) {
        if ((1000 < 1) && false) {
            v20 := []int{p0, 4294967295}
            p0 -= p0
            switch p0 {
            case 42:
                if len(v20) >= 2 {
                    v20 = v20[1:len(v20)]
                }
            default:
                if (3.14159 >= 1.5) {
                    var v21 int = p0
                    for fi22 := 0; fi22 < 6; fi22++ {
                        p1 = p1
                        var v23 float32 = 1e10
                        var v24 uint64 = 2
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                    _ = v21
                }
            }
            _ = v20
        } else {
            v25 := GenS1{F0: 55, F1: 519}
            var v26 float64 = -2.25
            _ = v25
            _ = v26
        }
        p0 += 296
    } else {
        v27 := []int{p0, p0}
        p0 += (^(p0))
        _ = v27
    }
    v28 := GenS2{F0: 361, F1: p1}
    _ = p0
    _ = p1
    _ = v28
    return 255
}
func genFn1(p0 int64, p1 uint) bool {
    p1 = p1
    var v29 int32 = -1
    for fi30 := 0; fi30 < 2; fi30++ {
        p1 = p1
        fmt.Println(p1)
        _ = fi30
    }
    _ = p0
    _ = p1
    _ = v29
    return false
}

func main() {
    var v1 uint = 255
    var v2 float64 = (0.0 - 1.5)
    var v3 uint64 = 832
    v4 := make(map[string]bool)
    var v5 int64 = -9223372036854775808
    var v6 uint8 = 127
    fmt.Println(v6, v2, v1)
    v31 := GenS2{F0: 313, F1: 1}
    var v32 int32 = -56
    for fi33 := 0; fi33 < 16; fi33++ {
        v1 = uint(v3)
        _ = fi33
    }
    var v34 uint32 = 42
    var v35 int16 = 32767
    switch "a" {
    case "hello":
        v36 := v4[("a" + "")]
        _ = v36
    case "foo bar":
        fmt.Println(v3, v1)
    }
    var v37 int64 = (v5 | v5)
    v38 := GenS1{F0: 100, F1: 2147483648}
    if false {
        var v39 int8 = -128
        v1 = 42
        for fi40 := 0; fi40 < 15; fi40++ {
            for fi41 := 0; fi41 < 11; fi41++ {
                v42 := []int{(fi41 - fi40)}
                v3 = v3
                var v43 bool = true
                v44 := GenS3{F0: v34}
                var v45 int64 = (v37 << (((v1 ^ uint(v34))) & 63))
                _ = fi41
                _ = v42
                _ = v43
                _ = v44
                _ = v45
            }
            _ = fi40
        }
        _ = v39
    } else {
        if false {
            for fi46 := 0; fi46 < 7; fi46++ {
                if true {
                    v47 := GenS4{F0: v1, F1: (-(fi46))}
                    var v48 int64 = v37
                    _ = v47
                    _ = v48
                } else {
                    v49 := GenS5{F0: v3, F1: v32, F2: 1}
                    _ = v49
                }
                _ = fi46
            }
            v50 := GenS6{F0: (v32 ^ -1), F1: (v37 * v5)}
            _ = v50
        }
        var v51 float64 = float64(0.1)
        if true {
            v52 := GenS7{F0: "hello", F1: "foo bar", F2: 42}
            var v53 uint8 = v6
            _ = v52
            _ = v53
        }
        _ = v51
    }
    v35 -= (v35 * 1000)
    for fi54 := 0; fi54 < 6; fi54++ {
        v34 -= v34
        var v55 string = ("hello" + "foo bar")
        var v56 int32 = (v32 + (v32 >> ((uint(v34)) & 63)))
        _ = fi54
        _ = v55
        _ = v56
    }
    var v57 uint64 = v3
    var v58 int64 = -1
    if true {
        if false {
            v59 := []int{238, 537}
            _ = v59
        } else {
            if false {
                switch -9223372036854775808 {
                case -823:
                    if (uint16(v1) == 939) {
                        v37 += 42
                    }
                case 42:
                    v60 := GenS7{F0: ("" + ("" + "a")), F1: ("foo bar" + ""), F2: (^(v34))}
                    _ = v60
                case -384:
                    fmt.Println(v1, v1)
                default:
                    v6 += (v6 * v6)
                }
                v61 := v4["hello"]
                if v61 {
                    var v62 int8 = -128
                    for fi63 := 0; fi63 < 3; fi63++ {
                        var v64 int32 = -523
                        fmt.Println(v3)
                        switch ("" + "foo bar") {
                        case "a":
                            v65 := GenS2{F0: 459, F1: 9223372036854775807}
                            _ = v65
                        case "\n\t":
                            if v61 {
                                var v66 uint64 = (v3 ^ v57)
                                v67 := GenS8{F0: (^(v62)), F1: 255, F2: fi63}
                                _ = v66
                                _ = v67
                            } else {
                                v68 := GenS5{F0: 911, F1: v64, F2: v5}
                                _ = v68
                            }
                        case "foo bar":
                            switch (fi63 & (fi63 | -1)) {
                            case 9223372036854775807:
                                v69 := []int{}
                                _ = v69
                            case 1:
                                if (fi63 <= int(v57)) {
                                    var v70 int16 = (v35 / (((v35 >> ((v1) & 63))) | 1))
                                    _ = v70
                                } else {
                                    var v71 uint16 = 32767
                                    _ = v71
                                }
                            }
                        default:
                            var v72 int = fi63
                            _ = v72
                        }
                        _ = fi63
                        _ = v64
                    }
                    _ = v62
                }
                _ = v61
            }
            v32 = v32
            var v73 float64 = 1.5
            _ = v73
        }
        v74 := GenS9{F0: v35}
        _ = v74
    }
    for fi75 := 0; fi75 < 14; fi75++ {
        var v76 uint32 = (^(v34))
        fmt.Println(v76, v3)
        for fi77 := 0; fi77 < 7; fi77++ {
            var v78 uint = v1
            _ = fi77
            _ = v78
        }
        _ = fi75
        _ = v76
    }
    switch int(v34) {
    case -9223372036854775808:
        if (!(1.5 < (float32(v2) + float32(v2)))) {
            for fi79 := 0; fi79 < 14; fi79++ {
                v3 += (v57 << ((4294967295) & 63))
                fmt.Println(v1)
                switch "a" {
                case "a":
                    var v80 float64 = float64((-2.25 + ((0.1 * float32(v2)) * -2.25)))
                    _ = v80
                case "":
                    var v81 bool = (true || false)
                    _ = v81
                }
                _ = fi79
            }
            v82 := GenS10{F0: "\n\t", F1: uint64(v6), F2: (v57 * v57)}
            _ = v82
        } else {
            for fi83 := 0; fi83 < 14; fi83++ {
                v84 := GenS4{F0: uint(v34), F1: fi83}
                _ = fi83
                _ = v84
            }
        }
    case 449:
        if ("" <= "hello") {
            v85 := GenS6{F0: v32, F1: (v5 << (((v1 << ((v1) & 63))) & 63))}
            var v86 int64 = -9223372036854775808
            _ = v85
            _ = v86
        } else {
            if true {
                for fi87 := 0; fi87 < 15; fi87++ {
                    v35 += (v35 / (((-(v35))) | 1))
                    _ = fi87
                }
            } else {
                v34 = v34
                v35 += v35
                v3 = v57
            }
        }
    case -2147483648:
        for fi88 := 0; fi88 < 20; fi88++ {
            var v89 uint = v1
            v35 -= 545
            _ = fi88
            _ = v89
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v34
    _ = v35
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v57
    _ = v58
    _ = v6
}
