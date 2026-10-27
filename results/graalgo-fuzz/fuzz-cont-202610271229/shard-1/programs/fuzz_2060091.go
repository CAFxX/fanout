package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 int64; F1 uint64 }
type GenS3 struct { F0 bool; F1 uint; F2 string }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint; F1 uint; F2 uint32 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 string; F1 uint64 }
func genFn0(p0 int16, p1 int) int16 {
    var v7 uint64 = 18446744073709551615
    v8 := GenS1{F0: -65, F1: 824}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return -32768
}
func genFn1(p0 uint32, p1 uint8) int16 {
    p1 = p1
    for fi9 := 0; fi9 < 19; fi9++ {
        v10 := GenS2{F0: 0, F1: 0}
        if true {
            var v11 float64 = 3.14159
            _ = v11
        } else {
            v12 := GenS3{F0: ("\n\t" >= "foo bar"), F1: 1, F2: "foo bar"}
            for fi13 := 0; fi13 < 9; fi13++ {
                if false {
                    var v14 string = "a"
                    switch v14 {
                    case "a":
                        var v15 string = "hello"
                        _ = v15
                    default:
                        v16 := GenS3{F0: true, F1: 0, F2: v14}
                        _ = v16
                    }
                    var v17 string = ""
                    _ = v14
                    _ = v17
                } else {
                    if true {
                        var v18 int8 = -20
                        _ = v18
                    }
                }
                p1 = 6
                var v19 int32 = -28
                _ = fi13
                _ = v19
            }
            if false {
                fmt.Println(fi9, p0)
            } else {
                v20 := GenS1{F0: 97, F1: 32767}
                for fi21 := 0; fi21 < 5; fi21++ {
                    fmt.Println(p0, p1)
                    p1 -= (p1 << ((18446744073709551615) & 63))
                    if false {
                        switch (fi21 & fi21) {
                        case -2147483648:
                            var v22 uint16 = 1
                            _ = v22
                        case -1:
                            var v23 int = 128
                            _ = v23
                        case 0:
                            var v24 int = (-(fi9))
                            _ = v24
                        default:
                            if (515 >= p0) {
                                var v25 int32 = -2147483648
                                _ = v25
                            }
                        }
                    } else {
                        v26 := GenS4{F0: (fi21 | fi21)}
                        if (!true) {
                            switch "foo bar" {
                            case "":
                                if (true && false) {
                                    v27 := []int{(fi21 * fi9), 1}
                                    var v28 int8 = 1
                                    var v29 int32 = 2147483647
                                    _ = v27
                                    _ = v28
                                    _ = v29
                                }
                            case "a":
                                p1 = p1
                            default:
                                var v30 uint16 = 591
                                _ = v30
                            }
                            var v31 int16 = -1
                            if true {
                                if true {
                                    v32 := GenS4{F0: 0}
                                    _ = v32
                                } else {
                                    var v33 uint = 481
                                    _ = v33
                                }
                            }
                            _ = v31
                        } else {
                            switch "\n\t" {
                            case "a":
                                var v34 uint64 = 42
                                _ = v34
                            case "":
                                fmt.Println(p1, fi21)
                            }
                            p0 += 68
                            v35 := GenS5{F0: 65535}
                            if (false && true) {
                                fmt.Println(fi9, p0)
                                fmt.Println(fi21, fi21, p0)
                                v36 := GenS6{F0: 887, F1: 255, F2: p0}
                                fmt.Println(fi21, fi9)
                                if ((1 >= 32767) && ((p0 ^ (p0 % ((0) | 1))) != p0)) {
                                    if false {
                                        switch "\n\t" {
                                        case "hello":
                                            if (!("hello" < (("hello" + ("foo bar" + "\n\t")) + ("hello" + "")))) {
                                                var v37 uint64 = 353
                                                v38 := []int{}
                                                p1 = p1
                                                v39 := GenS1{F0: -2147483648, F1: 629}
                                                v40 := GenS5{F0: 255}
                                                _ = len(v38)
                                                _ = v37
                                                _ = v38
                                                _ = v39
                                                _ = v40
                                            } else {
                                                v41 := []int{(fi9 + (fi9 | fi21)), (fi21 % (((fi21 - (fi21 / ((fi9) | 1)))) | 1)), -1, 127}
                                                _ = v41
                                            }
                                        case "":
                                            fmt.Println(p0, p0)
                                        case "foo bar":
                                            var v42 int64 = -1
                                            _ = v42
                                        }
                                        v43 := []int{}
                                        _ = v43
                                    } else {
                                        var v44 string = "\n\t"
                                        _ = v44
                                    }
                                    p1 += p1
                                }
                                p1 += 42
                                switch fi9 {
                                case 121:
                                    switch (fi21 << ((42) & 63)) {
                                    case 32767:
                                        if (2147483647 > 2147483647) {
                                            var v45 uint64 = 18446744073709551615
                                            if (859 != 0) {
                                                var v46 bool = (fi9 <= 1)
                                                p0 -= (^(p0))
                                                var v47 uint8 = (^(p1))
                                                _ = v46
                                                _ = v47
                                            } else {
                                                v48 := GenS7{F0: 2147483648}
                                                _ = v48
                                            }
                                            _ = v45
                                        } else {
                                            p1 = uint8(p0)
                                            v49 := GenS4{F0: fi21}
                                            _ = v49
                                        }
                                    case -1:
                                        p1 = p1
                                        var v50 int8 = 100
                                        p0 -= p0
                                        _ = v50
                                    default:
                                        fmt.Println(p0, p0)
                                    }
                                case 32767:
                                    var v51 uint64 = 18446744073709551615
                                    _ = v51
                                case 0:
                                    v52 := []int{fi9}
                                    _ = v52
                                }
                                _ = v36
                            }
                            _ = v35
                        }
                        _ = v26
                    }
                    _ = fi21
                }
                _ = v20
            }
            _ = v12
        }
        for fi53 := 0; fi53 < 13; fi53++ {
            p1 += 128
            _ = fi53
        }
        _ = fi9
        _ = v10
    }
    p1 = 128
    for fi54 := 0; fi54 < 14; fi54++ {
        if true {
            if false {
                var v55 uint64 = 42
                fmt.Println(fi54, v55)
                _ = v55
            } else {
                var v56 uint16 = 0
                for fi57 := 0; fi57 < 15; fi57++ {
                    switch fi54 {
                    case -9223372036854775808:
                        switch "\n\t" {
                        case "hello":
                            v56 = (v56 & v56)
                        }
                    case -824:
                        v58 := GenS3{F0: (!false), F1: 390, F2: "hello"}
                        _ = v58
                    }
                    _ = fi57
                }
                p1 += 128
                _ = v56
            }
        } else {
            var v59 uint8 = p1
            var v60 int32 = -1
            var v61 float64 = 3.14159
            _ = v59
            _ = v60
            _ = v61
        }
        v62 := []int{(-(fi54)), (fi54 - fi54), fi54, fi54}
        p0 += 2147483648
        _ = fi54
        _ = v62
    }
    v63 := GenS8{F0: "\n\t", F1: 18446744073709551615}
    _ = p0
    _ = p1
    _ = v63
    return -1
}

func main() {
    var v1 bool = true
    var v2 int = 2147483647
    v3 := []int{}
    v4 := []int{(v2 % ((v2) | 1))}
    var v5 uint16 = 688
    var v6 uint32 = 310
    v2 = 2147483647
    v64 := GenS7{F0: v6}
    var v65 uint64 = 90
    v2 = (v2 >> ((0) & 63))
    for fi66 := 0; fi66 < 13; fi66++ {
        var v67 uint = 65535
        _ = fi66
        _ = v67
    }
    v3 = append(v3, (v2 - (v2 >> ((18446744073709551615) & 63))))
    switch "" {
    case "foo bar":
        for fi68 := 0; fi68 < 18; fi68++ {
            if v1 {
                v65 -= v65
                var v69 uint8 = 255
                _ = v69
            } else {
                fmt.Println(v2)
            }
            if v1 {
                fmt.Println(v65, v1)
                var v70 uint16 = (v5 | uint16(v6))
                _ = v70
            } else {
                v1 = (!v1)
                fmt.Println(v6)
            }
            _ = fi68
        }
    default:
        switch (v2 ^ (v2 + v2)) {
        case 9223372036854775807:
            var v71 string = ("\n\t" + "\n\t")
            _ = v71
        case 32767:
            v65 = v65
        case 402:
            var v72 bool = (32768 > 2147483647)
            _ = v72
        default:
            var v73 uint = 899
            _ = v73
        }
    }
    var v74 int32 = int32(v5)
    var v75 uint32 = (v6 ^ v6)
    fmt.Println(v1)
    if len(v3) > 0 {
        v3[0] = (v2 + v2)
    }
    v75 -= (v75 + v6)
    if false {
        if len(v3) > 0 {
            v3[0] = v2
        }
    } else {
        if true {
            var v76 float32 = 1e10
            if (!true) {
                v74 = 283
                v77 := GenS7{F0: v75}
                _ = v77
            }
            _ = len(v4)
            _ = v76
        } else {
            v65 = (v65 ^ v65)
            if true {
                var v78 int8 = -91
                for fi79 := 0; fi79 < 4; fi79++ {
                    fmt.Println(v2, v1, v65)
                    var v80 uint16 = (v5 * v5)
                    var v81 int32 = (v74 << ((1) & 63))
                    _ = fi79
                    _ = v80
                    _ = v81
                }
                _ = v78
            }
            if ("foo bar" == "\n\t") {
                v74 = -1
                var v82 uint8 = 0
                if len(v3) > 0 {
                    v3[0] = (v2 ^ v2)
                }
                _ = v82
            } else {
                switch "\n\t" {
                case "":
                    for fi83 := 0; fi83 < 4; fi83++ {
                        v1 = ((((float32(1e10) + -2.25) - 0.1) < 0.0) && (1 >= v74))
                        _ = fi83
                    }
                default:
                    switch "\n\t" {
                    case "foo bar":
                        var v84 uint64 = v65
                        _ = v84
                    case "\n\t":
                        var v85 float64 = 1.5
                        _ = v85
                    }
                }
                fmt.Println(v6, v74, v65)
            }
        }
    }
    v6 -= (^(v75))
    if false {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        for fi86 := 0; fi86 < 4; fi86++ {
            v5 = (v5 / ((v5) | 1))
            _ = fi86
        }
    } else {
        v65 = 1
    }
    for fi87 := 0; fi87 < 3; fi87++ {
        fmt.Println(v75, v1)
        _ = fi87
    }
    var v88 uint16 = 255
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v65:", v65)
    fmt.Println("v74:", v74)
    fmt.Println("v75:", v75)
    fmt.Println("v88:", v88)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v64
    _ = v65
    _ = v74
    _ = v75
    _ = v88
}
