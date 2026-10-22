package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int; F1 uint8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int64; F1 string; F2 int16 }
type GenS5 struct { F0 uint8; F1 uint }
type GenS6 struct { F0 uint; F1 int8; F2 uint64 }
type GenS7 struct { F0 bool; F1 int32; F2 uint16 }
type GenS8 struct { F0 int; F1 uint32; F2 int16 }
type GenS9 struct { F0 uint32; F1 string; F2 uint32 }
func genFn0(p0 int8, p1 uint64) uint32 {
    p0 = p0
    var v6 float64 = (1.5 + 0.1)
    p0 += p0
    if (1 == -32768) {
        switch 566 {
        case 127:
            v7 := []int{}
            _ = v7
        default:
            v8 := GenS1{F0: 0}
            _ = v8
        }
    }
    _ = p0
    _ = p1
    _ = v6
    return 2147483648
}
func genFn1(p0 bool, p1 int64, p2 string) int16 {
    var v9 uint8 = 17
    switch ((p2 + "") + p2) {
    case "a":
        var v10 uint16 = 1
        _ = v10
    case "\n\t":
        var v11 int16 = -715
        _ = v11
    }
    var v12 uint8 = 17
    var v13 string = "\n\t"
    switch (p2 + "hello") {
    case "hello":
        if (!(2 <= 0)) {
            v12 -= (v12 + 8)
            var v14 uint64 = 0
            fmt.Println(p2)
            _ = v14
        } else {
            for fi15 := 0; fi15 < 5; fi15++ {
                p2 = v13
                p2 = ""
                var v16 bool = (36 < -86)
                _ = fi15
                _ = v16
            }
        }
    case "\n\t":
        var v17 uint32 = 65535
        _ = v17
    default:
        var v18 uint16 = 840
        _ = v18
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    _ = v9
    return 42
}

func main() {
    v1 := []uint32{157, 2147483647}
    v2 := make(map[uint32]string)
    var v3 int = 205
    var v4 float64 = -2.25
    var v5 string = (("\n\t" + "a") + "a")
    for fi19 := 0; fi19 < 6; fi19++ {
        v5 = "\n\t"
        var v20 int64 = 393
        _ = fi19
        _ = v20
    }
    v3 -= (v3 | 561)
    if (2 > 790) {
        if false {
            var v21 uint64 = 482
            v22 := GenS2{F0: 4294967295, F1: 252}
            _ = v21
            _ = v22
        } else {
            var v23 int = (v3 ^ v3)
            _ = len(v1)
            v24 := GenS3{F0: genFn1(false, -1, v5)}
            _ = v23
            _ = v24
        }
        var v25 int16 = int16(v3)
        for fi26 := 0; fi26 < 12; fi26++ {
            var v27 float64 = v4
            _ = fi26
            _ = v27
        }
        _ = v25
    }
    if ((0 < 1) || false) {
        v5 = ""
        for fi28 := 0; fi28 < 10; fi28++ {
            var v29 uint8 = 151
            if false {
                var v30 int8 = 100
                for fi31 := 0; fi31 < 20; fi31++ {
                    if (true && ((v29 | 0) == v29)) {
                        var v32 int = (-(v3))
                        _ = v32
                    }
                    if true {
                        var v33 int8 = 127
                        v29 = 56
                        _ = v33
                    }
                    var v34 int32 = -1
                    if false {
                        var v35 int64 = -1
                        _ = v35
                    } else {
                        if (567 >= 4294967295) {
                            if true {
                                v36 := GenS1{F0: v29}
                                _ = v36
                            } else {
                                v37 := GenS1{F0: (v29 & (v29 % ((42) | 1)))}
                                var v38 int16 = 0
                                _ = v37
                                _ = v38
                            }
                            if (42 != -805) {
                                fmt.Println(v30, v3)
                            } else {
                                if (9 > 4294967295) {
                                    v39 := GenS1{F0: (v29 << ((1) & 63))}
                                    _ = v39
                                } else {
                                    switch (fi31 & (-(fi28))) {
                                    case 141:
                                        var v40 uint16 = 50
                                        _ = v40
                                    case -502:
                                        v4 += v4
                                    case 1:
                                        v41 := GenS4{F0: 9223372036854775807, F1: v5, F2: -1}
                                        _ = v41
                                    default:
                                        var v42 bool = false
                                        _ = v42
                                    }
                                }
                                fmt.Println(fi28)
                                if (654 == 385) {
                                    var v43 string = v5
                                    if (v30 >= v30) {
                                        v5 = v43
                                        var v44 float64 = 3.14159
                                        v45 := GenS1{F0: 2}
                                        fmt.Println(fi28)
                                        v4 -= 1e10
                                        v46 := GenS2{F0: (fi31 << ((4294967295) & 63)), F1: v29}
                                        if false {
                                            v47 := GenS5{F0: (v29 | v29), F1: 945}
                                            _ = v47
                                        } else {
                                            var v48 int32 = (v34 / ((v34) | 1))
                                            _ = v48
                                        }
                                        v2[2147483647] = "foo bar"
                                        _ = v44
                                        _ = v45
                                        _ = v46
                                    }
                                    v43 = v5
                                    _ = v43
                                }
                            }
                        } else {
                            var v49 uint32 = genFn0(127, 0)
                            _ = v49
                        }
                    }
                    _ = len(v2)
                    _ = fi31
                    _ = v34
                }
                _ = v30
            } else {
                v29 = 250
                if false {
                    var v50 uint64 = 2
                    _ = v50
                } else {
                    switch (fi28 ^ fi28) {
                    case -384:
                        switch v5 {
                        case "hello":
                            v3 = v3
                        case "a":
                            for fi51 := 0; fi51 < 18; fi51++ {
                                switch -9223372036854775808 {
                                case -492:
                                    v4 += (0.1 / ((v4 + v4) * v4))
                                case 2147483647:
                                    v52 := v2[0]
                                    _ = v52
                                }
                                _ = fi51
                            }
                        default:
                            if false {
                                switch -568 {
                                case 872:
                                    v53 := GenS6{F0: 4294967295, F1: 93, F2: 0}
                                    _ = v53
                                case 715:
                                    if true {
                                        if (!false) {
                                            var v54 uint64 = 641
                                            _ = v54
                                        } else {
                                            v55 := GenS7{F0: false, F1: -285, F2: 2}
                                            for fi56 := 0; fi56 < 16; fi56++ {
                                                var v57 string = (v5 + v5)
                                                v4 = 0.0
                                                _ = fi56
                                                _ = v57
                                            }
                                            switch fi28 {
                                            case 0:
                                                switch "" {
                                                case "a":
                                                    v29 = 255
                                                case "foo bar":
                                                    for fi58 := 0; fi58 < 5; fi58++ {
                                                        var v59 uint = 0
                                                        if (521 <= 183) {
                                                            fmt.Println(fi58)
                                                            var v60 bool = true
                                                            if (!v60) {
                                                                v29 = (v29 << (((v59 << (((^(v59))) & 63))) & 63))
                                                                var v61 int8 = 101
                                                                fmt.Println(v60, v60)
                                                                _ = v61
                                                            } else {
                                                                v62 := GenS3{F0: -28}
                                                                switch ("foo bar" + "hello") {
                                                                case "hello":
                                                                    var v63 int = v3
                                                                    _ = v63
                                                                }
                                                                _ = v62
                                                            }
                                                            v5 = "foo bar"
                                                            _ = v60
                                                        } else {
                                                            if false {
                                                                v64 := GenS4{F0: 42, F1: ("\n\t" + (v5 + ("" + "a"))), F2: 292}
                                                                _ = v64
                                                            } else {
                                                                if (657 <= -690) {
                                                                    v65 := GenS1{F0: (v29 & v29)}
                                                                    fmt.Println(v5, fi28)
                                                                    _ = v65
                                                                } else {
                                                                    v66 := v2[42]
                                                                    var v67 int8 = 9
                                                                    _ = v66
                                                                    _ = v67
                                                                }
                                                            }
                                                            v68 := GenS2{F0: (fi58 - fi58), F1: v29}
                                                            if false {
                                                                v59 -= v59
                                                                v3 = (fi58 << ((v59) & 63))
                                                                v69 := GenS4{F0: -9223372036854775808, F1: ((v5 + (v5 + v5)) + "a"), F2: 0}
                                                                _ = v69
                                                            }
                                                            _ = v68
                                                        }
                                                        _ = fi58
                                                        _ = v59
                                                    }
                                                case "\n\t":
                                                    fmt.Println(v3, fi28, v29)
                                                default:
                                                    fmt.Println(v29, v29)
                                                }
                                            default:
                                                v70 := GenS5{F0: v29, F1: 42}
                                                _ = v70
                                            }
                                            _ = v55
                                        }
                                        delete(v2, 0)
                                        v4 = v4
                                    }
                                case -688:
                                    var v71 string = (v5 + v5)
                                    _ = v71
                                default:
                                    delete(v2, 65535)
                                }
                                if (true || (-1 != -128)) {
                                    for fi72 := 0; fi72 < 19; fi72++ {
                                        var v73 uint32 = 2147483648
                                        _ = fi72
                                        _ = v73
                                    }
                                }
                            } else {
                                if true {
                                    for fi74 := 0; fi74 < 18; fi74++ {
                                        if len(v1) > 0 {
                                            v1[0] = 2147483647
                                        }
                                        _ = fi74
                                    }
                                } else {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    var v75 int32 = 0
                                    v75 -= (v75 - (^(v75)))
                                    _ = v75
                                }
                                switch v3 {
                                case 575:
                                    switch ((v5 + v5) + v5) {
                                    case "foo bar":
                                        for fi76 := 0; fi76 < 7; fi76++ {
                                            switch fi76 {
                                            case 134:
                                                switch v5 {
                                                case "hello":
                                                    v2[753] = (v5 + "hello")
                                                case "foo bar":
                                                    v4 = 0.1
                                                case "\n\t":
                                                    v3 = (-(v3))
                                                default:
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    var v77 string = (("" + "a") + v5)
                                                    _ = v77
                                                }
                                            default:
                                                if true {
                                                    if true {
                                                        v78 := GenS8{F0: 32767, F1: 0, F2: genFn1(false, 9223372036854775807, v5)}
                                                        var v79 int = v3
                                                        _ = v78
                                                        _ = v79
                                                    } else {
                                                        v80 := GenS5{F0: (v29 / ((128) | 1)), F1: 65535}
                                                        if (2 != 757) {
                                                            v81 := v2[2147483648]
                                                            switch fi28 {
                                                            case 2147483647:
                                                                var v82 float32 = 3.14159
                                                                _ = v82
                                                            }
                                                            var v83 float64 = 1e10
                                                            _ = v81
                                                            _ = v83
                                                        }
                                                        if len(v1) > 0 {
                                                            v1[0] = 162
                                                        }
                                                        fmt.Println(fi76)
                                                        _ = v80
                                                    }
                                                    var v84 int32 = 1
                                                    var v85 int16 = 0
                                                    fmt.Println(v29)
                                                    _ = v84
                                                    _ = v85
                                                } else {
                                                    switch fi76 {
                                                    case -176:
                                                        v86 := GenS9{F0: genFn0(127, 42), F1: "a", F2: 0}
                                                        _ = v86
                                                    case 378:
                                                        var v87 int64 = -1
                                                        _ = v87
                                                    case 42:
                                                        var v88 uint64 = 0
                                                        v88 = (^(v88))
                                                        v4 += 3.14159
                                                        _ = v88
                                                    }
                                                    var v89 uint64 = 1
                                                    var v90 float32 = float32((0.0 + (0.1 + v4)))
                                                    v3 += 32768
                                                    var v91 float64 = (float64(v90) - 3.14159)
                                                    _ = v89
                                                    _ = v90
                                                    _ = v91
                                                }
                                            }
                                            fmt.Println(fi28, fi28, fi28)
                                            _ = fi76
                                        }
                                    case "hello":
                                        v5 = "hello"
                                    default:
                                        v1 = append(v1, 44)
                                    }
                                case 1:
                                    if true {
                                        v29 += (^(v29))
                                    }
                                case 127:
                                    _ = len(v1)
                                }
                                fmt.Println(v3)
                            }
                        }
                    }
                }
                if false {
                    fmt.Println(v5, v29)
                } else {
                    var v92 float64 = 3.14159
                    v93 := GenS4{F0: 925, F1: v5, F2: -1}
                    _ = v92
                    _ = v93
                }
            }
            _ = fi28
            _ = v29
        }
    }
    var v94 int8 = 100
    v95 := GenS3{F0: int16(v94)}
    if len(v1) > 0 {
        v1[0] = 1
    }
    _ = len(v1)
    var v96 uint64 = 4294967295
    var v97 int64 = 633
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v94:", v94)
    fmt.Println("v96:", v96)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v94
    _ = v95
    _ = v96
    _ = v97
}
