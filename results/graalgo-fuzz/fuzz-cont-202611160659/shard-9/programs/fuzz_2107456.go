package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 int64; F1 int64; F2 uint }
type GenS4 struct { F0 int8; F1 int64; F2 uint8 }
type GenS5 struct { F0 uint16; F1 bool }
type GenS6 struct { F0 uint; F1 uint }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 string; F1 int16 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 bool; F1 bool }
type GenS11 struct { F0 int8; F1 bool }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 uint32; F1 bool }
type GenS14 struct { F0 int64; F1 bool; F2 bool }
type GenS15 struct { F0 string }
func genFn0(p0 uint64, p1 int8) int {
    p0 -= (^(p0))
    switch 127 {
    case -9223372036854775808:
        var v7 float64 = ((3.14159 + 0.1) / 0.1)
        _ = v7
    default:
        if (-2147483648 > 9223372036854775807) {
            v8 := GenS1{F0: p1}
            _ = v8
        } else {
            var v9 uint16 = 0
            var v10 uint32 = 65535
            if (false || (-348 < -1)) {
                for fi11 := 0; fi11 < 12; fi11++ {
                    for fi12 := 0; fi12 < 1; fi12++ {
                        if false {
                            switch "foo bar" {
                            case "a":
                                v13 := []int{fi11, (^(fi11))}
                                v14 := GenS2{F0: -1, F1: 278}
                                if len(v13) >= 2 {
                                    v13 = v13[1:len(v13)]
                                }
                                switch fi12 {
                                case -702:
                                    var v15 float64 = (-2.25 * -2.25)
                                    _ = v15
                                default:
                                    var v16 int16 = -32768
                                    _ = v16
                                }
                                var v17 int8 = (^(p1))
                                _ = v13
                                _ = v14
                                _ = v17
                            case "hello":
                                v18 := GenS3{F0: 440, F1: -1, F2: 471}
                                _ = v18
                            default:
                                if true {
                                    var v19 int16 = 42
                                    _ = v19
                                } else {
                                    fmt.Println(fi12, v10)
                                    var v20 int64 = 9223372036854775807
                                    _ = v20
                                }
                            }
                            v21 := GenS3{F0: -334, F1: -86, F2: 521}
                            _ = v21
                        } else {
                            p0 += 0
                            if true {
                                v22 := GenS4{F0: 42, F1: -9223372036854775808, F2: 255}
                                _ = v22
                            }
                        }
                        _ = fi12
                    }
                    p1 = -128
                    _ = fi11
                }
                var v23 int32 = 0
                _ = v23
            }
            _ = v10
            _ = v9
        }
    }
    var v24 string = "foo bar"
    var v25 int16 = 32767
    switch 1 {
    case 128:
        p1 -= p1
    case -9223372036854775808:
        v26 := GenS1{F0: p1}
        _ = v26
    case -445:
        v27 := GenS1{F0: (p1 << ((0) & 63))}
        _ = v27
    default:
        if false {
            for fi28 := 0; fi28 < 17; fi28++ {
                p0 = p0
                _ = fi28
            }
            for fi29 := 0; fi29 < 14; fi29++ {
                fmt.Println(v24, fi29)
                switch fi29 {
                case 1:
                    v30 := GenS5{F0: 65535, F1: true}
                    _ = v30
                }
                for fi31 := 0; fi31 < 2; fi31++ {
                    v24 = (("a" + "a") + ("\n\t" + (v24 + ("foo bar" + v24))))
                    v32 := GenS6{F0: 4294967295, F1: 2}
                    p0 = 42
                    v33 := []int{(fi29 % ((fi31) | 1)), (^(fi29)), (fi31 - int(v25))}
                    fmt.Println(fi29, p1, fi31)
                    _ = fi31
                    _ = v32
                    _ = v33
                }
                _ = fi29
            }
            var v34 uint8 = 67
            _ = v34
        }
    }
    _ = p0
    _ = p1
    _ = v24
    _ = v25
    return 9223372036854775807
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]bool)
    var v3 int = 2147483647
    var v4 uint64 = 1
    var v5 float32 = 1e10
    v6 := make(map[string]int)
    if true {
        switch "" {
        case "":
            var v35 uint64 = v4
            _ = v35
        case "hello":
            v4 = v4
        }
        if (1 <= 32767) {
            if (false && ((-2.25 / 1e10) != 0.0)) {
                if (70 <= 90) {
                    for fi36 := 0; fi36 < 4; fi36++ {
                        if (860 > -101) {
                            v3 += (v3 - (v3 % ((1) | 1)))
                            v37 := GenS7{F0: -753}
                            _ = v37
                        } else {
                            var v38 uint = 18446744073709551615
                            var v39 int32 = -584
                            v39 = v39
                            _ = v38
                            _ = v39
                        }
                        fmt.Println(fi36, v4, fi36)
                        if (76 != 42) {
                            var v40 uint64 = 823
                            for fi41 := 0; fi41 < 16; fi41++ {
                                switch (v3 + genFn0(9, 100)) {
                                case 840:
                                    v42 := v1[1]
                                    _ = v42
                                }
                                v43 := GenS5{F0: 0, F1: (255 < 18446744073709551615)}
                                _ = fi41
                                _ = v43
                            }
                            for fi44 := 0; fi44 < 19; fi44++ {
                                v3 = (^(fi44))
                                _ = fi44
                            }
                            _ = v40
                        } else {
                            v45 := []int{(v3 * genFn0(v4, 1)), genFn0(v4, 127), v3}
                            var v46 int64 = 308
                            _ = v45
                            _ = v46
                        }
                        _ = fi36
                    }
                    v47 := GenS8{F0: "a", F1: 832}
                    _ = v47
                }
                var v48 float32 = (1.5 / v5)
                v48 += v5
                _ = v48
            } else {
                fmt.Println(v3, v4, v5)
                v49 := []int{}
                fmt.Println(v4, v3)
                _ = v49
            }
            v3 = (v3 >> ((349) & 63))
        } else {
            if (v3 <= (v3 % (((-(v3))) | 1))) {
                if ((false || (true || false)) && true) {
                    v50 := GenS9{F0: ""}
                    _ = v50
                }
                var v51 int8 = 1
                _ = v51
            } else {
                v3 += genFn0(91, 42)
                var v52 int32 = -2147483648
                _ = v52
            }
        }
        v3 += 0
    }
    v5 -= 0.1
    if ("a" <= "hello") {
        v53 := GenS9{F0: ("" + "foo bar")}
        if false {
            v54 := GenS4{F0: -48, F1: 122, F2: 0}
            _ = v54
        } else {
            v55 := []int{}
            v56 := GenS10{F0: false, F1: ((!(1000 == 42)) && ((^(v4)) > 0))}
            v57 := GenS11{F0: 127, F1: (v3 <= v3)}
            _ = v55
            _ = v56
            _ = v57
        }
        _ = v53
    } else {
        var v58 int32 = 2147483647
        v59 := []int{(v3 >> ((255) & 63)), genFn0(v4, 1)}
        v5 -= (v5 * 0.1)
        _ = v58
        _ = v59
    }
    v60 := GenS12{F0: (^(v4))}
    if (false || true) {
        fmt.Println(v4)
        var v61 bool = (true && true)
        var v62 string = (("" + "hello") + "\n\t")
        _ = v61
        _ = v62
    } else {
        if (100 != 71) {
            v63 := []int{4294967295, genFn0(v4, 42), (v3 * (-(v3))), (v3 * -860)}
            for fi64 := 0; fi64 < 7; fi64++ {
                delete(v6, "a")
                fmt.Println(v5, v4, v3)
                var v65 int64 = -1
                _ = fi64
                _ = v65
            }
            _ = v63
        } else {
            switch (v3 | (v3 << ((18446744073709551615) & 63))) {
            case 1:
                switch "" {
                case "":
                    if (!((0 <= -1) && true)) {
                        switch genFn0(v4, -1) {
                        case -17:
                            fmt.Println(v5, v3)
                        case 922:
                            v6[("foo bar" + ("a" + "foo bar"))] = (v3 * (v3 ^ (v3 | genFn0((v4 - (v4 << ((1) & 63))), 100))))
                        default:
                            v66 := GenS11{F0: 3, F1: ((714 > 18446744073709551615) || (!(-2147483648 == 0)))}
                            _ = v66
                        }
                    }
                case "\n\t":
                    var v67 uint = 187
                    _ = v67
                default:
                    switch (v3 - v3) {
                    case -808:
                        v4 = v4
                    case 2147483647:
                        if (42 > 42) {
                            var v68 uint32 = 1
                            for fi69 := 0; fi69 < 7; fi69++ {
                                fmt.Println(v3, v68, v4)
                                for fi70 := 0; fi70 < 16; fi70++ {
                                    if true {
                                        v71 := []int{}
                                        if true {
                                            var v72 int32 = -1
                                            var v73 int64 = 0
                                            _ = v72
                                            _ = v73
                                        } else {
                                            v4 += (v4 ^ (v4 - v4))
                                        }
                                        v68 -= v68
                                        _ = v71
                                    }
                                    _ = fi70
                                }
                                _ = fi69
                            }
                            var v74 int64 = 1
                            _ = v68
                            _ = v74
                        } else {
                            for fi75 := 0; fi75 < 2; fi75++ {
                                v3 -= v3
                                _ = fi75
                            }
                            for fi76 := 0; fi76 < 18; fi76++ {
                                switch "" {
                                case "hello":
                                    var v77 int = genFn0(2, 100)
                                    _ = v77
                                case "\n\t":
                                    switch "foo bar" {
                                    case "\n\t":
                                        var v78 uint8 = 42
                                        _ = v78
                                    case "":
                                        for fi79 := 0; fi79 < 2; fi79++ {
                                            v80 := GenS13{F0: 42, F1: true}
                                            fmt.Println(v3, v3, fi79)
                                            _ = fi79
                                            _ = v80
                                        }
                                    }
                                case "a":
                                    v4 = (v4 * v4)
                                }
                                v81 := GenS14{F0: 1, F1: false, F2: (-511 < 1)}
                                switch "a" {
                                case "foo bar":
                                    fmt.Println(v4, fi76)
                                }
                                _ = fi76
                                _ = v81
                            }
                        }
                    }
                }
            case -9223372036854775808:
                v5 = v5
            default:
                v3 += genFn0(v4, -128)
            }
            for fi82 := 0; fi82 < 13; fi82++ {
                v83 := GenS5{F0: 2, F1: false}
                v4 = (v4 << ((4294967295) & 63))
                for fi84 := 0; fi84 < 16; fi84++ {
                    v85 := []int{-72, 0, fi82}
                    _ = fi84
                    _ = v85
                }
                _ = fi82
                _ = v83
            }
            var v86 float64 = 3.14159
            _ = v86
        }
    }
    for fi87 := 0; fi87 < 5; fi87++ {
        switch "foo bar" {
        case "\n\t":
            fmt.Println(v5)
        case "a":
            var v88 uint8 = 255
            _ = v88
        }
        if false {
            _ = len(v1)
            if false {
                v5 = 0.0
            } else {
                v89 := GenS12{F0: (v4 / ((759) | 1))}
                var v90 uint = 1
                var v91 uint32 = 42
                _ = v89
                _ = v90
                _ = v91
            }
            for fi92 := 0; fi92 < 9; fi92++ {
                switch "a" {
                case "foo bar":
                    var v93 int = 42
                    _ = v93
                case "\n\t":
                    v3 += (fi87 ^ fi92)
                case "hello":
                    if (false && (v4 <= (v4 * 899))) {
                        v5 = (v5 / (v5 * float32(3.14159)))
                        if (!(255 >= 435)) {
                            fmt.Println(v5)
                        }
                        var v94 uint32 = 864
                        _ = v94
                    } else {
                        v95 := GenS7{F0: 9223372036854775807}
                        _ = v95
                    }
                default:
                    var v96 int64 = 1
                    _ = v96
                }
                v3 = 128
                _ = fi92
            }
        }
        _ = fi87
    }
    fmt.Println(v5, v5, v3)
    switch (v3 << ((2) & 63)) {
    case 32768:
        var v97 string = "a"
        _ = v97
    default:
        v98 := GenS15{F0: ("foo bar" + "foo bar")}
        _ = v98
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v60
}
