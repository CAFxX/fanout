package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint32; F1 int32 }
type GenS5 struct { F0 uint; F1 bool }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int64; F1 uint; F2 int8 }
type GenS8 struct { F0 uint16 }
func genFn0(p0 uint16, p1 uint32, p2 bool) int {
    for fi5 := 0; fi5 < 11; fi5++ {
        p2 = p2
        if p2 {
            var v6 int8 = 100
            _ = v6
        } else {
            v7 := GenS1{F0: 2, F1: 517}
            _ = v7
        }
        _ = fi5
    }
    v8 := []int{32768, 272, 32768, 895}
    var v9 int32 = 0
    var v10 int64 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v8
    _ = v9
    return 1
}
func genFn1(p0 uint8, p1 uint16, p2 int64) string {
    for fi11 := 0; fi11 < 15; fi11++ {
        fmt.Println(p2)
        _ = fi11
    }
    if (42 != -638) {
        for fi12 := 0; fi12 < 5; fi12++ {
            if (true && (true || false)) {
                v13 := GenS1{F0: 478, F1: p1}
                _ = v13
            } else {
                switch ("foo bar" + (("a" + "foo bar") + "hello")) {
                case "hello":
                    v14 := GenS2{F0: 20}
                    _ = v14
                case "\n\t":
                    var v15 uint16 = 0
                    _ = v15
                case "":
                    v16 := GenS2{F0: 1}
                    _ = v16
                default:
                    p2 -= (p2 % (((p2 * p2)) | 1))
                }
                var v17 float32 = (-2.25 - 1e10)
                _ = v17
            }
            _ = fi12
        }
        fmt.Println(p0, p2)
    } else {
        v18 := GenS1{F0: 4294967295, F1: p1}
        v19 := GenS3{F0: 18446744073709551615}
        _ = v18
        _ = v19
    }
    v20 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    return ""
}

func main() {
    var v1 int16 = 860
    var v2 float64 = ((float64(0.0) - 0.1) / -2.25)
    var v3 float64 = v2
    var v4 int = -1
    v21 := GenS2{F0: -480}
    fmt.Println(v1, v1, v1)
    v22 := GenS3{F0: 53}
    if true {
        v23 := []int{genFn0(0, 42, true), (^(v4)), v4}
        v24 := GenS1{F0: 42, F1: 65535}
        if (!(false && false)) {
            if len(v23) > 0 {
                v23[0] = 2147483647
            }
            fmt.Println(v2)
            if (("a" == ((("\n\t" + "") + ("" + ("\n\t" + "\n\t"))) + ("foo bar" + "hello"))) || (128 != 128)) {
                for fi25 := 0; fi25 < 10; fi25++ {
                    var v26 int = fi25
                    for fi27 := 0; fi27 < 17; fi27++ {
                        var v28 uint16 = 2
                        v29 := GenS4{F0: 682, F1: 2147483647}
                        v26 = 0
                        _ = fi27
                        _ = v28
                        _ = v29
                    }
                    _ = fi25
                    _ = v26
                }
                v4 = v4
                for fi30 := 0; fi30 < 1; fi30++ {
                    if len(v23) >= 2 {
                        v23 = v23[1:len(v23)]
                    }
                    _ = fi30
                }
            } else {
                v31 := GenS4{F0: 2147483647, F1: 0}
                var v32 int8 = 127
                if (!false) {
                    var v33 float64 = v2
                    if false {
                        v32 += (v32 | (-(v32)))
                        for fi34 := 0; fi34 < 1; fi34++ {
                            switch (-(v4)) {
                            case 316:
                                var v35 uint8 = 127
                                _ = v35
                            case 320:
                                switch "a" {
                                case "foo bar":
                                    if (v2 == v3) {
                                        if true {
                                            v23 = append(v23, 159)
                                            fmt.Println(v3, v33)
                                        }
                                    } else {
                                        var v36 uint16 = 65535
                                        for fi37 := 0; fi37 < 15; fi37++ {
                                            var v38 uint16 = (v36 % ((v36) | 1))
                                            _ = fi37
                                            _ = v38
                                        }
                                        if (false && true) {
                                            var v39 float32 = (1e10 + 3.14159)
                                            var v40 int32 = -1
                                            var v41 int16 = (v1 >> ((42) & 63))
                                            _ = v39
                                            _ = v40
                                            _ = v41
                                        }
                                        _ = v36
                                    }
                                default:
                                    if (42 < 658) {
                                        if len(v23) > 0 {
                                            v23[0] = fi34
                                        }
                                    }
                                }
                            }
                            _ = fi34
                        }
                    } else {
                        switch "" {
                        case "\n\t":
                            var v42 uint32 = 65535
                            _ = v42
                        case "hello":
                            switch "a" {
                            case "hello":
                                if len(v23) > 0 {
                                    v23[0] = v4
                                }
                            case "\n\t":
                                var v43 int16 = 1000
                                _ = v43
                            case "a":
                                fmt.Println(v33, v33, v32)
                            default:
                                v23 = append(v23, v4)
                            }
                        }
                        switch v4 {
                        case 0:
                            var v44 uint16 = 608
                            _ = v44
                        case 1:
                            var v45 float64 = v33
                            _ = v45
                        default:
                            switch "" {
                            case "hello":
                                v46 := GenS4{F0: 2147483647, F1: -1}
                                _ = v46
                            }
                        }
                        v47 := GenS3{F0: 42}
                        _ = v47
                    }
                    v1 = (v1 << ((2) & 63))
                    _ = v33
                }
                _ = v31
                _ = v32
            }
        } else {
            var v48 int = (^(v4))
            for fi49 := 0; fi49 < 12; fi49++ {
                v48 = 2147483647
                fmt.Println(v1, v3, v2)
                fmt.Println(fi49, v3, v1)
                _ = fi49
            }
            if len(v23) >= 2 {
                v23 = v23[1:len(v23)]
            }
            _ = v48
        }
        _ = v23
        _ = v24
    }
    v1 -= v1
    var v50 int16 = v1
    if true {
        if ((v2 + v2) >= v3) {
            if (!false) {
                if (787 == 444) {
                    var v51 float64 = v3
                    fmt.Println(v50, v3, v1)
                    _ = v51
                } else {
                    v52 := []int{(^(v4))}
                    _ = v52
                }
            } else {
                for fi53 := 0; fi53 < 5; fi53++ {
                    if ((("hello" + (("" + "\n\t") + "\n\t")) + ("foo bar" + "hello")) <= (("a" + "\n\t") + "foo bar")) {
                        for fi54 := 0; fi54 < 8; fi54++ {
                            v2 += (v2 - (((v3 - v3) + 3.14159) + v2))
                            _ = fi54
                        }
                        v3 = 3.14159
                    }
                    if true {
                        switch fi53 {
                        case 128:
                            v55 := []int{}
                            _ = v55
                        case 977:
                            var v56 float64 = v3
                            _ = v56
                        default:
                            v50 = -32768
                        }
                        var v57 uint32 = 2147483648
                        v58 := []int{32767}
                        _ = v57
                        _ = v58
                    } else {
                        v3 = v3
                    }
                    _ = fi53
                }
            }
            fmt.Println(v50, v1, v4)
        }
        var v59 float64 = (1.5 - v2)
        _ = v59
    }
    fmt.Println(v4)
    var v60 uint16 = 1
    v61 := []int{v4, (^(v4)), (v4 / ((42) | 1)), 9223372036854775807}
    if true {
        if (101 < 916) {
            var v62 int8 = 42
            for fi63 := 0; fi63 < 8; fi63++ {
                var v64 string = "a"
                if true {
                    fmt.Println(v4)
                }
                switch (("foo bar" + v64) + v64) {
                case "hello":
                    if true {
                        v4 -= v4
                    } else {
                        fmt.Println(v50, v1)
                    }
                case "\n\t":
                    v65 := GenS2{F0: 0}
                    _ = v65
                }
                _ = fi63
                _ = v64
            }
            _ = v62
        }
        var v66 uint = 962
        if (true || ((v60 / (((v60 | v60)) | 1)) > (v60 + v60))) {
            if len(v61) >= 2 {
                v61 = v61[1:len(v61)]
            }
            for fi67 := 0; fi67 < 2; fi67++ {
                switch ("a" + "") {
                case "foo bar":
                    switch "\n\t" {
                    case "\n\t":
                        for fi68 := 0; fi68 < 4; fi68++ {
                            switch 9223372036854775807 {
                            case -552:
                                v69 := GenS4{F0: 2147483647, F1: -1}
                                _ = v69
                            case -113:
                                fmt.Println(v3)
                            case -221:
                                if false {
                                    if true {
                                        fmt.Println(v4, v66)
                                    } else {
                                        var v70 uint16 = v60
                                        _ = v70
                                    }
                                    fmt.Println(v1, fi67, v60)
                                } else {
                                    if len(v61) > 0 {
                                        v61[0] = genFn0((^(v60)), 77, true)
                                    }
                                    v71 := GenS3{F0: 1}
                                    v1 = (v50 | (^(v1)))
                                    _ = v71
                                }
                                switch v4 {
                                case 42:
                                    v2 = v3
                                case 32768:
                                    v72 := GenS5{F0: v66, F1: true}
                                    _ = v72
                                case 210:
                                    v4 += (fi67 & genFn0(uint16(v66), 65535, true))
                                }
                                var v73 string = "\n\t"
                                _ = v73
                            }
                            _ = fi68
                        }
                    case "":
                        fmt.Println(v50, fi67)
                    case "hello":
                        for fi74 := 0; fi74 < 2; fi74++ {
                            v75 := GenS2{F0: 2147483647}
                            _ = fi74
                            _ = v75
                        }
                    default:
                        fmt.Println(v66, v50)
                    }
                case "\n\t":
                    v2 -= v3
                case "hello":
                    v1 = (v1 - (v50 + v1))
                default:
                    v1 -= 1000
                }
                v4 += (v4 | fi67)
                _ = fi67
            }
            var v76 int8 = 127
            _ = v76
        }
        _ = v66
    }
    for fi77 := 0; fi77 < 11; fi77++ {
        fmt.Println(v60, v2, fi77)
        for fi78 := 0; fi78 < 10; fi78++ {
            switch (v4 - v4) {
            case -963:
                v3 -= (1e10 / (float64(1e10) + ((v3 - float64(1.5)) + 3.14159)))
            case 688:
                v60 -= v60
            case 1:
                if true {
                    if len(v61) >= 2 {
                        v61 = v61[1:len(v61)]
                    }
                    var v79 uint8 = 42
                    _ = v79
                } else {
                    v80 := GenS6{F0: "a"}
                    v61 = append(v61, v4)
                    v3 = v2
                    _ = v80
                }
            }
            if false {
                v2 -= (-2.25 - (v2 - (1e10 + v2)))
                switch "" {
                case "":
                    if (true && true) {
                        var v81 int16 = v50
                        v82 := GenS5{F0: 4294967295, F1: false}
                        _ = v81
                        _ = v82
                    } else {
                        if len(v61) > 0 {
                            v61[0] = fi78
                        }
                    }
                case "foo bar":
                    v1 = (v1 ^ (-(v1)))
                default:
                    var v83 uint8 = 2
                    _ = v83
                }
            } else {
                if true {
                    v84 := GenS7{F0: -136, F1: 4294967295, F2: 24}
                    fmt.Println(v50)
                    v2 = 3.14159
                    if ((-1 != -1) && false) {
                        v85 := GenS8{F0: (v60 & (v60 % ((v60) | 1)))}
                        _ = len(v61)
                        _ = v85
                    }
                    v60 = v60
                    if len(v61) > 0 {
                        v61[0] = fi78
                    }
                    v50 -= -1
                    v86 := GenS6{F0: "hello"}
                    _ = v84
                    _ = v86
                }
            }
            _ = fi78
        }
        _ = fi77
    }
    fmt.Println(v60, v50, v50)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v50:", v50)
    fmt.Println("v60:", v60)
    fmt.Println("slicelen v61:", len(v61), v61)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v50
    _ = v60
    _ = v61
}
