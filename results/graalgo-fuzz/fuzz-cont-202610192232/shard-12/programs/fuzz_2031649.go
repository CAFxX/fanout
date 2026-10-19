package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 uint64; F1 int }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint8; F1 string }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 bool; F1 bool; F2 int32 }
type GenS7 struct { F0 int; F1 uint32; F2 int8 }
func genFn0(p0 int8, p1 uint16, p2 string) uint32 {
    var v7 uint16 = 0
    v8 := []int{127, 127, 1}
    var v9 uint16 = 1
    v8 = append(v8, 1)
    switch p2 {
    case "a":
        if (true || false) {
            if true {
                var v10 uint32 = 2147483647
                _ = v10
            }
            if (2 >= 0) {
                var v11 string = ("a" + ("" + p2))
                v12 := GenS1{F0: 2147483647, F1: 1}
                _ = v11
                _ = v12
            }
            if len(v8) > 0 {
                v8[0] = 32768
            }
        }
    case "":
        p0 -= (-(p0))
    case "\n\t":
        if (!false) {
            var v13 uint64 = 0
            _ = len(v8)
            if (1 != int(v7)) {
                var v14 string = p2
                var v15 int8 = p0
                if true {
                    switch v14 {
                    case "":
                        v16 := GenS1{F0: -1, F1: 1000}
                        _ = v16
                    case "\n\t":
                        p2 = ("a" + ((v14 + (v14 + "a")) + "foo bar"))
                    case "a":
                        v17 := GenS1{F0: -1, F1: 0}
                        _ = v17
                    default:
                        if len(v8) >= 2 {
                            v8 = v8[1:len(v8)]
                        }
                    }
                    v15 = p0
                }
                _ = v14
                _ = v15
            } else {
                if len(v8) > 0 {
                    v8[0] = -1
                }
            }
            _ = v13
        }
    default:
        p2 = "a"
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 2147483648
}

func main() {
    var v1 int16 = 1
    v2 := make(map[uint32]int)
    var v3 uint16 = 42
    var v4 uint64 = 906
    var v5 uint32 = 534
    var v6 uint32 = 4294967295
    if (!false) {
        delete(v2, genFn0(42, v3, (("" + "hello") + "\n\t")))
    } else {
        var v18 int32 = 0
        var v19 float64 = (1e10 + ((3.14159 / 1e10) + 1e10))
        _ = v18
        _ = v19
    }
    v20 := GenS2{F0: v4, F1: 1}
    fmt.Println(v5, v6)
    for fi21 := 0; fi21 < 12; fi21++ {
        for fi22 := 0; fi22 < 7; fi22++ {
            v23 := v2[v5]
            _ = fi22
            _ = v23
        }
        if (false && true) {
            fmt.Println(v1, v1)
            v24 := []int{379}
            _ = v24
        }
        var v25 uint8 = 0
        _ = fi21
        _ = v25
    }
    if (true && (!true)) {
        for fi26 := 0; fi26 < 13; fi26++ {
            switch "hello" {
            case "a":
                for fi27 := 0; fi27 < 5; fi27++ {
                    var v28 bool = false
                    var v29 uint64 = 978
                    _ = fi27
                    _ = v28
                    _ = v29
                }
            case "":
                if true {
                    switch "a" {
                    case "\n\t":
                        v30 := GenS2{F0: 921, F1: fi26}
                        _ = v30
                    case "hello":
                        switch ("a" + "foo bar") {
                        case "\n\t":
                            v5 = v6
                        case "a":
                            v5 += 42
                        case "hello":
                            fmt.Println(v6, v5, v4)
                        default:
                            v4 = (v4 * 0)
                        }
                    }
                    v6 -= genFn0(70, 255, ("" + "hello"))
                } else {
                    for fi31 := 0; fi31 < 20; fi31++ {
                        fmt.Println(fi26, v5, v1)
                        var v32 float64 = 0.1
                        _ = fi31
                        _ = v32
                    }
                }
            case "\n\t":
                v33 := GenS3{F0: (^(v6))}
                _ = v33
            }
            _ = fi26
        }
    } else {
        var v34 int8 = -98
        var v35 int8 = v34
        delete(v2, v5)
        _ = v34
        _ = v35
    }
    var v36 int = 2147483647
    for fi37 := 0; fi37 < 4; fi37++ {
        var v38 int8 = -124
        if true {
            _ = len(v2)
            var v39 float64 = 1.5
            var v40 bool = false
            _ = v39
            _ = v40
        }
        _ = fi37
        _ = v38
    }
    v3 -= (v3 << ((42) & 63))
    fmt.Println(v4, v36, v4)
    if (true || (202 != 42)) {
        v1 = 42
        switch "\n\t" {
        case "\n\t":
            for fi41 := 0; fi41 < 20; fi41++ {
                v36 += 1
                v4 = v4
                _ = fi41
            }
        }
        var v42 int32 = -335
        _ = v42
    } else {
        fmt.Println(v1, v5, v1)
    }
    switch (v36 ^ int(v4)) {
    case 42:
        var v43 int64 = -9223372036854775808
        _ = v43
    case 756:
        var v44 float32 = 1e10
        _ = v44
    default:
        for fi45 := 0; fi45 < 13; fi45++ {
            var v46 uint8 = 80
            v36 -= (^(v36))
            var v47 uint64 = (v4 >> ((65535) & 63))
            _ = fi45
            _ = v46
            _ = v47
        }
    }
    for fi48 := 0; fi48 < 14; fi48++ {
        v4 -= (v4 % ((0) | 1))
        var v49 int16 = (v1 - -898)
        v3 += v3
        _ = fi48
        _ = v49
    }
    for fi50 := 0; fi50 < 5; fi50++ {
        var v51 uint32 = 2
        _ = fi50
        _ = v51
    }
    switch 32768 {
    case 893:
        if (!false) {
            if false {
                fmt.Println(v1)
                fmt.Println(v3, v4, v4)
            }
            v52 := GenS3{F0: (^(v6))}
            _ = v52
        } else {
            var v53 int32 = int32(v1)
            _ = v53
        }
    case 567:
        for fi54 := 0; fi54 < 20; fi54++ {
            var v55 uint32 = (^(v6))
            for fi56 := 0; fi56 < 1; fi56++ {
                fmt.Println(v3)
                v57 := GenS3{F0: (v55 ^ 42)}
                if (904 < 812) {
                    v3 = v3
                    v6 = (v55 ^ 4294967295)
                } else {
                    if (!false) {
                        if false {
                            v4 = v4
                            var v58 bool = ((!true) && true)
                            var v59 uint8 = 62
                            var v60 uint8 = (v59 | v59)
                            _ = v58
                            _ = v59
                            _ = v60
                        }
                        v5 = 65535
                    } else {
                        switch "" {
                        case "":
                            var v61 uint32 = 4294967295
                            _ = v61
                        }
                    }
                    v62 := GenS4{F0: 235, F1: ""}
                    _ = v62
                }
                switch fi56 {
                case -659:
                    v55 -= v55
                case 0:
                    switch (^(v36)) {
                    case 2147483647:
                        v63 := GenS4{F0: 0, F1: ("hello" + ("\n\t" + "a"))}
                        _ = v63
                    case 9223372036854775807:
                        if (int(v3) == v36) {
                            v55 = (v55 + v6)
                        } else {
                            switch "hello" {
                            case "hello":
                                switch "foo bar" {
                                case "\n\t":
                                    v5 = v5
                                default:
                                    v64 := []int{}
                                    _ = v64
                                }
                            }
                        }
                    }
                default:
                    v6 += (v5 - 2147483647)
                }
                switch -859 {
                case 163:
                    switch v36 {
                    case 2147483647:
                        if (!true) {
                            fmt.Println(fi54)
                            var v65 uint32 = v6
                            _ = v65
                        } else {
                            v36 = (fi54 & 128)
                            var v66 float64 = (0.1 - 0.1)
                            switch "hello" {
                            case "hello":
                                if (!(-128 == 42)) {
                                    v67 := GenS5{F0: 1}
                                    if (v1 < -1) {
                                        var v68 float64 = (0.0 / float64(1e10))
                                        v69 := []int{(fi56 >> ((18446744073709551615) & 63)), 128, v36}
                                        fmt.Println(v4, v55)
                                        fmt.Println(v68, v68, fi54)
                                        _ = v68
                                        _ = v69
                                    }
                                    _ = v67
                                } else {
                                    var v70 int8 = -1
                                    switch fi54 {
                                    case 4294967295:
                                        var v71 uint64 = 2
                                        _ = v71
                                    case 127:
                                        v5 += uint32(v3)
                                    case -701:
                                        if (!((true || (428 == 65535)) && false)) {
                                            v70 = v70
                                            if false {
                                                v72 := []int{int(v4), v36, (v36 | -869)}
                                                var v73 int = fi56
                                                _ = v72
                                                _ = v73
                                            } else {
                                                delete(v2, v5)
                                                fmt.Println(fi54)
                                            }
                                        }
                                    default:
                                        var v74 float64 = 0.1
                                        _ = v74
                                    }
                                    _ = v70
                                }
                            case "a":
                                var v75 uint64 = v4
                                _ = v75
                            case "":
                                if (fi56 >= (fi56 | fi54)) {
                                    var v76 uint32 = (v5 & v5)
                                    switch "foo bar" {
                                    case "":
                                        fmt.Println(v66)
                                    case "foo bar":
                                        if false {
                                            v77 := GenS4{F0: 32, F1: "a"}
                                            v78 := GenS4{F0: 255, F1: "\n\t"}
                                            fmt.Println(fi54)
                                            _ = v77
                                            _ = v78
                                        } else {
                                            v79 := []int{fi54, 128, fi54, -2147483648}
                                            v80 := GenS6{F0: (42 == 1), F1: false, F2: 0}
                                            v6 -= genFn0(-3, (v3 | 976), "a")
                                            _ = v79
                                            _ = v80
                                        }
                                    case "a":
                                        var v81 int16 = (v1 / ((v1) | 1))
                                        _ = v81
                                    }
                                    _ = v76
                                } else {
                                    fmt.Println(fi54)
                                }
                            }
                            _ = v66
                        }
                    case -1:
                        var v82 uint8 = 2
                        _ = v82
                    }
                case 32768:
                    v83 := GenS7{F0: 1, F1: genFn0(93, v3, ("foo bar" + "a")), F2: 13}
                    _ = v83
                default:
                    v36 = fi56
                }
                var v84 uint16 = (v3 & v3)
                _ = fi56
                _ = v57
                _ = v84
            }
            _ = fi54
            _ = v55
        }
    }
    if ((!((true && (-153 != -569)) && true)) || ((-2.25 + 3.14159) < 1.5)) {
        var v85 int8 = 1
        v5 = 403
        switch v36 {
        case 434:
            v2[genFn0((v85 + 127), v3, ("hello" + ("hello" + "a")))] = (v36 | (v36 - int(v1)))
        }
        _ = v85
    } else {
        v6 -= (v6 ^ v5)
        var v86 float32 = 0.0
        fmt.Println(v4, v4)
        _ = v86
    }
    v5 = (v5 * (v6 / ((v6) | 1)))
    for fi87 := 0; fi87 < 1; fi87++ {
        v88 := []int{(v36 ^ (fi87 | v36)), (fi87 * (fi87 | 741)), v36}
        _ = fi87
        _ = v88
    }
    fmt.Println(v5, v5)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
