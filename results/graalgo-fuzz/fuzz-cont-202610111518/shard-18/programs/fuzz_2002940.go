package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int64; F1 uint64 }
type GenS4 struct { F0 int16; F1 int64; F2 uint16 }
type GenS5 struct { F0 int8; F1 int8 }
type GenS6 struct { F0 string; F1 uint16 }
type GenS7 struct { F0 uint; F1 uint16; F2 uint8 }
type GenS8 struct { F0 int16; F1 int8; F2 bool }
func genFn0(p0 int, p1 int16) int16 {
    v6 := []int{(p0 | (^(p0)))}
    p1 += p1
    p0 += p0
    v7 := GenS1{F0: -1}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return (p1 << ((255) & 63))
}

func main() {
    v1 := make(map[int]string)
    var v2 int64 = -979
    var v3 uint32 = 2147483647
    v4 := make(map[int]int)
    v5 := make(map[uint32]int)
    v3 = v3
    v8 := v1[-9223372036854775808]
    v9 := GenS2{F0: 175}
    delete(v1, 16)
    v10 := GenS1{F0: -1}
    var v11 uint64 = 18446744073709551615
    fmt.Println(v8)
    for fi12 := 0; fi12 < 19; fi12++ {
        v13 := GenS2{F0: 2}
        _ = fi12
        _ = v13
    }
    if false {
        var v14 uint8 = 127
        _ = v14
    }
    if (65535 > 464) {
        if true {
            if true {
                v15 := []int{32768, 230, -242, 42}
                v16 := GenS3{F0: 42, F1: 518}
                var v17 int16 = 42
                _ = v15
                _ = v16
                _ = v17
            }
            for fi18 := 0; fi18 < 13; fi18++ {
                v11 = (v11 & (v11 - (v11 % ((v11) | 1))))
                var v19 uint64 = (^(v11))
                if (!((v2 ^ v2) <= -690)) {
                    for fi20 := 0; fi20 < 8; fi20++ {
                        v21 := GenS4{F0: genFn0(42, 32767), F1: 1, F2: 255}
                        _ = fi20
                        _ = v21
                    }
                } else {
                    for fi22 := 0; fi22 < 13; fi22++ {
                        v23 := GenS3{F0: (v2 - v2), F1: v11}
                        v24 := GenS2{F0: 4294967295}
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                }
                _ = fi18
                _ = v19
            }
        }
        v2 = (v2 + (v2 - v2))
    } else {
        var v25 int8 = -1
        for fi26 := 0; fi26 < 15; fi26++ {
            if true {
                for fi27 := 0; fi27 < 9; fi27++ {
                    var v28 bool = false
                    fmt.Println(v11)
                    var v29 uint16 = 1
                    _ = fi27
                    _ = v28
                    _ = v29
                }
            } else {
                v2 = int64(v25)
            }
            if (v2 != (v2 << ((0) & 63))) {
                for fi30 := 0; fi30 < 12; fi30++ {
                    v31 := GenS3{F0: (-(v2)), F1: (v11 / (((v11 * 1)) | 1))}
                    v2 += (v2 + 1)
                    switch (-(fi26)) {
                    case 32768:
                        v3 -= v3
                    default:
                        switch v8 {
                        case "a":
                            v25 += (v25 * -80)
                        case "foo bar":
                            var v32 int = fi26
                            _ = v32
                        case "\n\t":
                            v2 -= (v2 & v2)
                        default:
                            var v33 uint64 = (^(v11))
                            _ = v33
                        }
                    }
                    _ = fi30
                    _ = v31
                }
            } else {
                var v34 float32 = 1e10
                switch ("" + v8) {
                case "a":
                    v35 := []int{(fi26 ^ (fi26 / ((fi26) | 1)))}
                    _ = v35
                case "\n\t":
                    fmt.Println(v8, fi26, v3)
                default:
                    var v36 string = (v8 + v8)
                    _ = v36
                }
                if false {
                    if (v8 <= "") {
                        var v37 uint = 1
                        _ = v37
                    }
                    fmt.Println(v2)
                    v38 := GenS2{F0: 255}
                    _ = v38
                } else {
                    fmt.Println(fi26, v25)
                    v3 += 2147483647
                    var v39 int16 = 335
                    _ = v39
                }
                _ = v34
            }
            var v40 uint16 = 750
            _ = fi26
            _ = v40
        }
        v41 := GenS3{F0: v2, F1: (^(v11))}
        _ = v25
        _ = v41
    }
    if true {
        var v42 bool = true
        v42 = true
        _ = v42
    } else {
        v8 = v8
        v8 = "\n\t"
        fmt.Println(v8, v8, v11)
    }
    var v43 float32 = float32((1e10 / (-2.25 - 0.1)))
    switch 954 {
    case 127:
        for fi44 := 0; fi44 < 6; fi44++ {
            for fi45 := 0; fi45 < 6; fi45++ {
                _ = len(v5)
                switch "\n\t" {
                case "":
                    if (!((true && false) || (712 > 0))) {
                        v46 := GenS5{F0: 0, F1: 1}
                        var v47 uint8 = 127
                        var v48 int16 = -554
                        v49 := GenS5{F0: -128, F1: 1}
                        v50 := v5[v3]
                        _ = v46
                        _ = v47
                        _ = v48
                        _ = v49
                        _ = v50
                    }
                case "foo bar":
                    var v51 int64 = v2
                    _ = v51
                case "\n\t":
                    var v52 int = -685
                    _ = v52
                default:
                    fmt.Println(v43, v3, fi44)
                }
                switch fi44 {
                case 0:
                    _ = len(v5)
                    delete(v5, v3)
                case 128:
                    v11 += (^(v11))
                case -2147483648:
                    var v53 float32 = v43
                    _ = v53
                }
                _ = fi45
            }
            for fi54 := 0; fi54 < 18; fi54++ {
                fmt.Println(v2, v3, fi44)
                v8 = "a"
                v43 = 0.0
                _ = fi54
            }
            if (((-39 > 0) && false) && true) {
                v55 := GenS2{F0: 42}
                var v56 string = v8
                _ = v55
                _ = v56
            } else {
                v43 -= v43
            }
            _ = fi44
        }
    case 32768:
        for fi57 := 0; fi57 < 7; fi57++ {
            v2 = (v2 % ((v2) | 1))
            for fi58 := 0; fi58 < 17; fi58++ {
                var v59 string = v8
                switch int(v2) {
                case 545:
                    var v60 uint16 = 55
                    _ = v60
                case -2147483648:
                    var v61 string = ""
                    v62 := GenS1{F0: -1}
                    v63 := GenS2{F0: 0}
                    _ = v61
                    _ = v62
                    _ = v63
                case 32767:
                    if (2147483647 < (v3 + (v3 & v3))) {
                        fmt.Println(v8)
                    }
                }
                var v64 bool = (0.1 == 1.5)
                _ = fi58
                _ = v59
                _ = v64
            }
            _ = fi57
        }
    case 9223372036854775807:
        fmt.Println(v2, v11, v2)
    default:
        if true {
            v2 = 9223372036854775807
            var v65 string = v8
            for fi66 := 0; fi66 < 2; fi66++ {
                for fi67 := 0; fi67 < 7; fi67++ {
                    fmt.Println(v11, v2)
                    v2 -= v2
                    if (32767 == -32768) {
                        v68 := GenS6{F0: ("hello" + ("" + "a")), F1: 65535}
                        v2 += v2
                        fmt.Println(fi66, v3, v3)
                        _ = v68
                    }
                    _ = fi67
                }
                for fi69 := 0; fi69 < 18; fi69++ {
                    v70 := GenS7{F0: 570, F1: 42, F2: 100}
                    if false {
                        v43 = 0.1
                    } else {
                        switch (fi66 | fi69) {
                        case 128:
                            v71 := GenS8{F0: 42, F1: 1, F2: ("\n\t" < v65)}
                            _ = v71
                        case 4294967295:
                            if (false && false) {
                                var v72 int32 = 735
                                _ = v72
                            } else {
                                fmt.Println(v2, fi66)
                            }
                        case -1:
                            if true {
                                var v73 int8 = 0
                                _ = v73
                            } else {
                                var v74 float64 = 3.14159
                                _ = v74
                            }
                        }
                        v43 -= 3.14159
                    }
                    v43 += 0.1
                    _ = fi69
                    _ = v70
                }
                fmt.Println(v11, v8)
                _ = fi66
            }
            _ = v65
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v8
    _ = v9
}
