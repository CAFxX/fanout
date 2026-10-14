package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 uint }
type GenS2 struct { F0 uint; F1 int16; F2 uint }
type GenS3 struct { F0 int32; F1 int8 }
type GenS4 struct { F0 int32; F1 int8; F2 int64 }
type GenS5 struct { F0 bool }
func genFn0(p0 int64, p1 string) uint8 {
    var v7 int8 = 100
    for fi8 := 0; fi8 < 8; fi8++ {
        if false {
            v9 := []int{(fi8 ^ fi8), fi8, 32768}
            for fi10 := 0; fi10 < 4; fi10++ {
                if false {
                    fmt.Println(fi10, p1, p1)
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                } else {
                    var v11 uint8 = 42
                    _ = v11
                }
                v9 = append(v9, (fi8 - (fi10 * fi8)))
                var v12 uint16 = 0
                _ = fi10
                _ = v12
            }
            _ = v9
        } else {
            if true {
                v13 := []int{fi8}
                if (!false) {
                    if false {
                        if false {
                            var v14 int8 = (v7 - v7)
                            switch fi8 {
                            case 32768:
                                var v15 int32 = 2147483647
                                _ = v15
                            case 707:
                                var v16 float64 = 1.5
                                _ = v16
                            case 1:
                                if len(v13) >= 2 {
                                    v13 = v13[1:len(v13)]
                                }
                            default:
                                for fi17 := 0; fi17 < 12; fi17++ {
                                    v14 = v7
                                    var v18 int64 = (p0 * p0)
                                    _ = fi17
                                    _ = v18
                                }
                            }
                            _ = v14
                        }
                        v7 -= v7
                        p0 += p0
                    } else {
                        var v19 uint16 = 55
                        p0 -= (p0 & int64(fi8))
                        _ = v19
                    }
                } else {
                    for fi20 := 0; fi20 < 11; fi20++ {
                        var v21 uint8 = 127
                        v21 -= (v21 ^ (v21 ^ (v21 + v21)))
                        _ = fi20
                        _ = v21
                    }
                }
                v22 := GenS1{F0: 255, F1: false, F2: 331}
                _ = v13
                _ = v22
            } else {
                var v23 bool = false
                fmt.Println(fi8, fi8, p1)
                _ = v23
            }
        }
        _ = fi8
    }
    if false {
        var v24 int8 = v7
        v7 += v7
        _ = v24
    } else {
        v25 := GenS1{F0: 32767, F1: ((v7 + v7) < (v7 | v7)), F2: 444}
        p0 -= p0
        _ = v25
    }
    var v26 uint8 = 1
    _ = p0
    _ = p1
    _ = v26
    _ = v7
    return 253
}
func genFn1(p0 int64, p1 string) int32 {
    var v27 int64 = p0
    v28 := GenS2{F0: 408, F1: 32767, F2: 2}
    v29 := []int{127}
    for fi30 := 0; fi30 < 17; fi30++ {
        switch 114 {
        case -546:
            var v31 int8 = 26
            _ = v31
        case 423:
            if len(v29) >= 2 {
                v29 = v29[1:len(v29)]
            }
        default:
            var v32 uint64 = 18446744073709551615
            _ = v32
        }
        _ = fi30
    }
    if (1 < 82) {
        v29 = append(v29, -1)
    }
    _ = p0
    _ = p1
    _ = v27
    _ = v28
    _ = v29
    return 42
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint16 = 32767
    var v3 string = "foo bar"
    v4 := make(map[int]string)
    var v5 uint = 42
    var v6 int16 = -1
    fmt.Println(v3, v6)
    v3 = v3
    switch 32767 {
    case 127:
        v3 = (v3 + v3)
    case 879:
        var v33 int64 = 42
        _ = v33
    case 32767:
        var v34 int32 = -1
        _ = v34
    }
    if false {
        for fi35 := 0; fi35 < 17; fi35++ {
            switch v3 {
            case "hello":
                switch "" {
                case "":
                    v36 := GenS1{F0: (v2 / (((v2 * (^(v2)))) | 1)), F1: false, F2: v5}
                    _ = v36
                case "a":
                    for fi37 := 0; fi37 < 20; fi37++ {
                        v38 := GenS2{F0: (v5 & (v5 % (((v5 / ((v5) | 1))) | 1))), F1: (^(v6)), F2: 42}
                        v3 = v3
                        _ = fi37
                        _ = v38
                    }
                case "foo bar":
                    var v39 uint = v5
                    _ = v39
                }
            default:
                var v40 int = fi35
                _ = v40
            }
            switch fi35 {
            case -1:
                for fi41 := 0; fi41 < 19; fi41++ {
                    var v42 bool = (false && false)
                    switch "" {
                    case "\n\t":
                        v43 := v4[(fi35 | 1)]
                        _ = v43
                    case "foo bar":
                        v4[(fi41 | (fi35 << ((18446744073709551615) & 63)))] = "\n\t"
                    case "hello":
                        v2 -= v2
                        v3 = ("a" + (v3 + "hello"))
                    }
                    switch "\n\t" {
                    case "hello":
                        v3 = ("hello" + ((((v3 + v3) + "") + ("" + "")) + ""))
                    case "\n\t":
                        v2 = (v2 / ((v2) | 1))
                    }
                    _ = fi41
                    _ = v42
                }
            default:
                if (971 < 2147483647) {
                    fmt.Println(v6)
                    var v44 uint32 = 757
                    _ = v44
                }
            }
            _ = fi35
        }
        if true {
            v45 := GenS1{F0: v2, F1: (!true), F2: v5}
            for fi46 := 0; fi46 < 9; fi46++ {
                fmt.Println(v2, v2, v2)
                _ = fi46
            }
            _ = v45
        } else {
            switch 127 {
            case -2147483648:
                if true {
                    var v47 uint = 42
                    _ = v47
                } else {
                    var v48 int16 = 1000
                    if (!(3.14159 == -2.25)) {
                        var v49 int = -1
                        if true {
                            switch v3 {
                            case "hello":
                                v6 = (v48 << ((v5) & 63))
                            }
                        } else {
                            v50 := v4[v49]
                            if true {
                                if false {
                                    var v51 uint32 = 161
                                    _ = v51
                                }
                            } else {
                                for fi52 := 0; fi52 < 5; fi52++ {
                                    var v53 uint32 = 2
                                    _ = fi52
                                    _ = v53
                                }
                                if (uint64(v2) < 2) {
                                    for fi54 := 0; fi54 < 6; fi54++ {
                                        var v55 string = (v3 + "\n\t")
                                        _ = fi54
                                        _ = v55
                                    }
                                    fmt.Println(v49)
                                    switch 622 {
                                    case 2147483647:
                                        if false {
                                            v5 -= v5
                                            v50 = "a"
                                        } else {
                                            v6 -= -32768
                                            v56 := GenS2{F0: 255, F1: v48, F2: 0}
                                            _ = v56
                                        }
                                    case 4294967295:
                                        fmt.Println(v50)
                                    }
                                }
                            }
                            var v57 uint16 = v2
                            _ = v50
                            _ = v57
                        }
                        v5 -= v5
                        _ = v49
                    }
                    if false {
                        v2 = 32767
                    }
                    _ = v48
                }
            case -1:
                v58 := []int{-354, 9223372036854775807, -740, 4294967295}
                _ = v58
            case 32768:
                v59 := []int{128, -1}
                _ = v59
            }
            v6 -= (v6 % ((-1) | 1))
        }
    } else {
        v60 := GenS3{F0: genFn1(0, ""), F1: 1}
        _ = v60
    }
    v61 := GenS4{F0: 161, F1: 0, F2: -9223372036854775808}
    v62 := GenS5{F0: true}
    v2 = 32767
    for fi63 := 0; fi63 < 9; fi63++ {
        for fi64 := 0; fi64 < 7; fi64++ {
            fmt.Println(v3)
            _ = fi64
        }
        var v65 uint32 = 42
        var v66 int8 = 0
        _ = fi63
        _ = v65
        _ = v66
    }
    var v67 uint16 = v2
    fmt.Println(v5, v6)
    switch -1 {
    case 9223372036854775807:
        for fi68 := 0; fi68 < 7; fi68++ {
            v5 -= v5
            _ = fi68
        }
    }
    v69 := GenS3{F0: 42, F1: 0}
    v70 := []int{0, 127, -2147483648, 4294967295}
    v71 := v4[9223372036854775807]
    v72 := GenS2{F0: 631, F1: (v6 >> ((v5) & 63)), F2: (v5 ^ (v5 << ((v5) & 63)))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v67:", v67)
    fmt.Println("slicelen v70:", len(v70), v70)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v61
    _ = v62
    _ = v67
    _ = v69
    _ = v70
    _ = v71
    _ = v72
}
