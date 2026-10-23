package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint; F1 int32 }
type GenS5 struct { F0 uint; F1 bool }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int64; F1 uint; F2 string }
type GenS8 struct { F0 int8; F1 string; F2 uint64 }
type GenS9 struct { F0 uint; F1 int }
type GenS10 struct { F0 uint16; F1 uint }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 bool; F1 int32 }
func genFn0(p0 uint16, p1 uint64, p2 int) int {
    for fi5 := 0; fi5 < 1; fi5++ {
        v6 := []int{(-(p2)), fi5, int(p0)}
        switch "" {
        case "\n\t":
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        }
        _ = fi5
        _ = v6
    }
    p1 = 18446744073709551615
    v7 := []int{}
    v8 := GenS1{F0: 157}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return p2
}
func genFn1() int32 {
    var v9 uint64 = 4294967295
    var v10 bool = true
    _ = v10
    _ = v9
    return -1
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]uint32)
    v3 := make(map[string]string)
    v4 := make(map[string]bool)
    var v11 int8 = -128
    if true {
        if true {
            v11 = v11
        }
    } else {
        v11 -= 127
        delete(v1, 42)
        v12 := GenS2{F0: 42}
        _ = v12
    }
    v13 := GenS2{F0: 798}
    fmt.Println(v11)
    var v14 uint16 = 722
    var v15 string = "foo bar"
    v15 = v15
    var v16 uint8 = 255
    for fi17 := 0; fi17 < 6; fi17++ {
        var v18 uint16 = v14
        for fi19 := 0; fi19 < 8; fi19++ {
            var v20 int16 = 1000
            switch genFn0((v14 % (((v18 & v18)) | 1)), 4294967295, fi19) {
            case 0:
                switch genFn0((v14 - (v14 * (v18 % ((v18) | 1)))), 1, (fi17 % (((fi19 % ((genFn0(v14, 0, fi19)) | 1))) | 1))) {
                case -189:
                    if (2147483647 > 505) {
                        v15 = v15
                        switch (fi19 * fi19) {
                        case 281:
                            switch "" {
                            case "foo bar":
                                var v21 uint32 = 65535
                                _ = v21
                            case "\n\t":
                                fmt.Println(fi19)
                            }
                        case 0:
                            v18 += v18
                            v22 := GenS2{F0: fi19}
                            v23 := GenS3{F0: (v18 << ((417) & 63))}
                            _ = v22
                            _ = v23
                        default:
                            v11 = v11
                        }
                        if (((0.1 * (0.1 - ((0.1 - float64(0.0)) - (-2.25 - 3.14159)))) + 1e10) > 0.1) {
                            switch "a" {
                            case "hello":
                                v16 += uint8(v18)
                            case "":
                                if (!((!(("" + ((v15 + "\n\t") + (v15 + (v15 + "a")))) != ("a" + ("foo bar" + v15)))) || false)) {
                                    v24 := []int{(fi17 | (fi17 % (((fi19 - genFn0((v14 / (((v18 % ((65535) | 1))) | 1)), 1, (fi19 - (fi17 << ((65535) & 63)))))) | 1)))}
                                    v14 -= v18
                                    v25 := GenS2{F0: int(v11)}
                                    _ = v24
                                    _ = v25
                                }
                            }
                            var v26 int = genFn0(v14, 4294967295, (fi17 << ((255) & 63)))
                            if false {
                                var v27 bool = (!(!((genFn1() > 0) || (-1 == (v20 + v20)))))
                                _ = v27
                            }
                            _ = v26
                        } else {
                            switch "\n\t" {
                            case "foo bar":
                                v15 = v15
                            case "a":
                                v18 += (v18 >> ((706) & 63))
                            case "hello":
                                if true {
                                    v3[v15] = "a"
                                    v28 := GenS4{F0: 42, F1: genFn1()}
                                    _ = len(v4)
                                    var v29 string = "\n\t"
                                    _ = v28
                                    _ = v29
                                }
                            default:
                                if false {
                                    v30 := GenS2{F0: genFn0(v18, 0, (fi17 << ((0) & 63)))}
                                    _ = len(v4)
                                    if ((v20 / ((v20) | 1)) >= v20) {
                                        var v31 float64 = (float64((-2.25 - ((-2.25 / -2.25) / -2.25))) - 1e10)
                                        v32 := v3[v15]
                                        _ = v31
                                        _ = v32
                                    } else {
                                        var v33 uint64 = 0
                                        if (v11 >= (v11 << ((4294967295) & 63))) {
                                            var v34 uint = 0
                                            if (!(genFn0(2, (v33 - 2), int(v34)) != (fi17 + genFn0(v14, v33, genFn0((^(v14)), (v33 - v33), (^(fi19))))))) {
                                                v35 := GenS5{F0: 18446744073709551615, F1: (!true)}
                                                _ = v35
                                            } else {
                                                v36 := v1[1]
                                                _ = v36
                                            }
                                            _ = v34
                                        } else {
                                            if false {
                                                if (v20 != (-(v20))) {
                                                    v11 -= v11
                                                } else {
                                                    var v37 int = 32767
                                                    _ = v37
                                                }
                                                v38 := GenS3{F0: (^(v18))}
                                                _ = v38
                                            } else {
                                                v14 = (v18 ^ 32767)
                                                v39 := []int{(fi17 | fi17), -630}
                                                if true {
                                                    var v40 float32 = (1.5 * 3.14159)
                                                    _ = v40
                                                } else {
                                                    v41 := GenS2{F0: int(v16)}
                                                    switch 9223372036854775807 {
                                                    case 1:
                                                        v16 = 0
                                                    }
                                                    if (1 <= 9223372036854775807) {
                                                        var v42 float32 = (-2.25 * 0.1)
                                                        _ = v42
                                                    }
                                                    _ = v41
                                                }
                                                _ = v39
                                            }
                                            switch "a" {
                                            case "\n\t":
                                                fmt.Println(fi19, v33, fi19)
                                            case "a":
                                                switch -105 {
                                                case -2147483648:
                                                    v11 -= -1
                                                case 42:
                                                    if (((v16 << ((65) & 63)) > v16) && (!false)) {
                                                        fmt.Println(v20, v11, v18)
                                                        v43 := GenS6{F0: 65535}
                                                        var v44 uint = 537
                                                        _ = v43
                                                        _ = v44
                                                    } else {
                                                        v18 = (v14 + v18)
                                                        v45 := []int{fi19, (-(fi19)), fi17, 2147483647}
                                                        switch -9223372036854775808 {
                                                        case 1:
                                                            v11 = v11
                                                        case 127:
                                                            var v46 float32 = 1e10
                                                            _ = v46
                                                        default:
                                                            fmt.Println(v33, v11)
                                                        }
                                                        _ = v45
                                                    }
                                                case -9223372036854775808:
                                                    v47 := GenS4{F0: 1, F1: -2147483648}
                                                    _ = v47
                                                }
                                            case "foo bar":
                                                switch "hello" {
                                                case "":
                                                    v14 = v14
                                                case "hello":
                                                    var v48 uint32 = 1
                                                    _ = v48
                                                case "\n\t":
                                                    var v49 uint32 = uint32(v16)
                                                    _ = v49
                                                default:
                                                    if false {
                                                        var v50 string = v15
                                                        _ = v50
                                                    }
                                                    v2[v15] = 4294967295
                                                }
                                            default:
                                                fmt.Println(v33, v15, fi17)
                                            }
                                            v51 := []int{genFn0(v18, 905, fi17), int(v18), (-(fi19)), int(v11)}
                                            _ = v51
                                        }
                                        _ = v33
                                    }
                                    _ = v30
                                } else {
                                    v20 += 950
                                    var v52 int32 = -646
                                    v53 := []int{fi17, (fi17 * (-(fi19))), (-(fi17))}
                                    _ = v52
                                    _ = v53
                                }
                            }
                            if ((1.5 - 1e10) == -2.25) {
                                v11 = v11
                            } else {
                                if ((!(867 > 0)) && ((v11 < v11) && false)) {
                                    v54 := GenS2{F0: (fi17 % ((21) | 1))}
                                    var v55 int64 = 42
                                    var v56 string = (v15 + "hello")
                                    _ = v54
                                    _ = v55
                                    _ = v56
                                }
                                v57 := GenS6{F0: 2}
                                v15 = v15
                                _ = v57
                            }
                            v15 = v15
                        }
                        if (false && ((!(1.5 > 1e10)) || (!true))) {
                            if true {
                                if true {
                                    var v58 bool = false
                                    _ = v58
                                } else {
                                    if true {
                                        if true {
                                            var v59 uint32 = 984
                                            _ = v59
                                        } else {
                                            var v60 uint16 = v14
                                            v15 = "a"
                                            v61 := GenS5{F0: 42, F1: false}
                                            fmt.Println(v11)
                                            var v62 int32 = 1
                                            v14 += 42
                                            v15 = (v15 + v15)
                                            _ = v60
                                            _ = v61
                                            _ = v62
                                        }
                                        v63 := GenS7{F0: 268, F1: 65535, F2: v15}
                                        switch v15 {
                                        case "hello":
                                            v64 := GenS8{F0: 100, F1: v15, F2: 42}
                                            _ = v64
                                        case "":
                                            v11 -= 0
                                            var v65 int16 = (v20 & (v20 - (v20 - (^(v20)))))
                                            _ = v65
                                        case "a":
                                            v66 := GenS9{F0: 1, F1: 127}
                                            _ = v66
                                        }
                                        _ = v63
                                    } else {
                                        var v67 uint8 = (^(v16))
                                        v68 := GenS10{F0: (v14 << ((1) & 63)), F1: 65535}
                                        switch fi19 {
                                        case -408:
                                            v69 := []int{int(v16), -973, (-(fi19)), fi17}
                                            _ = v69
                                        default:
                                            var v70 bool = (true && false)
                                            _ = v70
                                        }
                                        if true {
                                            if ((v20 % ((42) | 1)) > v20) {
                                                var v71 int64 = -755
                                                fmt.Println(v67, v20)
                                                v16 -= v16
                                                _ = v71
                                            }
                                        } else {
                                            if true {
                                                v72 := GenS11{F0: (v16 - v16)}
                                                v73 := []int{}
                                                var v74 uint32 = 4294967295
                                                _ = v72
                                                _ = v73
                                                _ = v74
                                            }
                                            v75 := GenS1{F0: v16}
                                            var v76 uint = 0
                                            _ = v75
                                            _ = v76
                                        }
                                        if false {
                                            var v77 uint8 = (v67 & (v67 & v67))
                                            _ = v77
                                        } else {
                                            if false {
                                                v67 -= (v67 / ((v16) | 1))
                                                delete(v1, 926)
                                            }
                                        }
                                        switch "foo bar" {
                                        case "hello":
                                            var v78 int = 127
                                            fmt.Println(v18, v14)
                                            var v79 int32 = 2147483647
                                            _ = v78
                                            _ = v79
                                        case "":
                                            fmt.Println(fi19)
                                            v67 = (v67 % (((v16 * 2)) | 1))
                                        case "foo bar":
                                            var v80 bool = false
                                            _ = v80
                                        }
                                        var v81 uint32 = 4294967295
                                        _ = v67
                                        _ = v68
                                        _ = v81
                                    }
                                    switch (fi17 + fi17) {
                                    case 2147483647:
                                        var v82 uint = 42
                                        _ = v82
                                    case 945:
                                        if ((true && (true || true)) || (9223372036854775807 <= int64(fi19))) {
                                            v4[("hello" + ((v15 + ("\n\t" + "\n\t")) + v15))] = true
                                            v14 = v18
                                        } else {
                                            v83 := GenS8{F0: 127, F1: v15, F2: 0}
                                            _ = v83
                                        }
                                    case 140:
                                        v16 = 128
                                    default:
                                        v84 := GenS12{F0: 517}
                                        _ = v84
                                    }
                                }
                                if false {
                                    fmt.Println(fi17)
                                    var v85 uint8 = (v16 - (^(v16)))
                                    v86 := GenS13{F0: true, F1: -1}
                                    _ = v85
                                    _ = v86
                                } else {
                                    delete(v4, v15)
                                    v18 -= (v14 * (v14 % (((v18 % ((v18) | 1))) | 1)))
                                    v15 = v15
                                }
                                fmt.Println(fi19, v15)
                            }
                            v11 = 14
                        }
                        if ((v11 / ((int8(v18)) | 1)) <= v11) {
                            var v87 uint32 = 65535
                            _ = v87
                        } else {
                            var v88 float32 = 3.14159
                            _ = v88
                        }
                    }
                case 42:
                    var v89 uint64 = 18446744073709551615
                    v89 -= v89
                    _ = v89
                default:
                    var v90 int16 = (v20 * v20)
                    _ = v90
                }
            case 32767:
                var v91 uint64 = 4294967295
                _ = v91
            default:
                _ = len(v2)
            }
            _ = fi19
            _ = v20
        }
        var v92 int8 = v11
        _ = fi17
        _ = v18
        _ = v92
    }
    for fi93 := 0; fi93 < 4; fi93++ {
        v11 += (^(v11))
        fmt.Println(v11, v11, v11)
        fmt.Println(v14, v11, fi93)
        _ = fi93
    }
    var v94 int32 = genFn1()
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v94:", v94)
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v94
}
