package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 int }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 uint32; F1 bool; F2 uint8 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int8; F1 int8 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint32; F1 int64 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 uint; F1 int64 }
func genFn0() int16 {
    if false {
        for fi7 := 0; fi7 < 15; fi7++ {
            v8 := GenS1{F0: -32768, F1: true, F2: (fi7 + fi7)}
            _ = fi7
            _ = v8
        }
        if false {
            v9 := []int{-271}
            var v10 int = 127
            if len(v9) > 0 {
                v9[0] = (v10 % (((v10 / ((668) | 1))) | 1))
            }
            _ = v10
            _ = v9
        }
        for fi11 := 0; fi11 < 3; fi11++ {
            var v12 float32 = 3.14159
            for fi13 := 0; fi13 < 10; fi13++ {
                v12 += v12
                v12 += v12
                v14 := []int{}
                _ = fi13
                _ = v14
            }
            _ = fi11
            _ = v12
        }
    } else {
        switch ("foo bar" + "a") {
        case "hello":
            var v15 uint16 = 698
            _ = v15
        case "":
            switch "\n\t" {
            case "hello":
                var v16 float64 = 3.14159
                _ = v16
            }
        }
        for fi17 := 0; fi17 < 17; fi17++ {
            var v18 float64 = 0.1
            _ = fi17
            _ = v18
        }
        var v19 uint8 = 221
        _ = v19
    }
    var v20 uint16 = 1
    var v21 int64 = 0
    _ = v20
    _ = v21
    return 0
}
func genFn1(p0 int8, p1 bool, p2 uint64) bool {
    p2 += (p2 & (p2 * 23))
    var v22 uint = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    return (1 != 2147483648)
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]string)
    v3 := []uint32{4294967295, 895, 12, 65535}
    var v4 string = "hello"
    var v5 uint16 = 65535
    var v6 uint64 = 467
    v4 = ("hello" + "hello")
    switch ((v4 + v4) + (v4 + v4)) {
    case "":
        v4 = v4
    case "foo bar":
        for fi23 := 0; fi23 < 8; fi23++ {
            switch v4 {
            case "":
                if true {
                    var v24 int8 = 127
                    _ = v24
                } else {
                    for fi25 := 0; fi25 < 14; fi25++ {
                        switch -410 {
                        case 591:
                            v26 := GenS2{F0: ((v4 + "hello") + "foo bar"), F1: -1}
                            _ = v26
                        case -2147483648:
                            fmt.Println(fi25, fi25, fi23)
                        case 628:
                            var v27 int64 = 9223372036854775807
                            _ = v27
                        default:
                            var v28 int8 = 102
                            _ = v28
                        }
                        var v29 float32 = 3.14159
                        if (127 <= uint8(v5)) {
                            v2[(v4 + (v4 + "hello"))] = "a"
                            if (true && ((v5 * (v5 * 980)) >= v5)) {
                                v30 := GenS2{F0: "foo bar", F1: 42}
                                var v31 string = "a"
                                _ = v30
                                _ = v31
                            } else {
                                v32 := GenS3{F0: 4294967295, F1: ((!(v5 != v5)) || (true || (2 > 126))), F2: 128}
                                _ = len(v1)
                                _ = v32
                            }
                            v4 = ((v4 + "hello") + ("" + v4))
                            delete(v1, 0)
                            if false {
                                v29 += float32((1e10 / 1.0))
                            } else {
                                v33 := GenS2{F0: (v4 + "foo bar"), F1: -1}
                                v34 := GenS4{F0: (v5 << ((784) & 63))}
                                switch "hello" {
                                case "hello":
                                    var v35 uint64 = (v6 | v6)
                                    _ = v35
                                }
                                _ = v33
                                _ = v34
                            }
                        } else {
                            v6 = (v6 | 4294967295)
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                        _ = fi25
                        _ = v29
                    }
                }
            }
            v36 := v2["a"]
            _ = fi23
            _ = v36
        }
    case "\n\t":
        var v37 int = 2147483647
        _ = v37
    }
    v4 = v4
    v38 := GenS4{F0: (v5 - v5)}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi39 := 0; fi39 < 14; fi39++ {
        for fi40 := 0; fi40 < 12; fi40++ {
            switch (fi40 << ((2) & 63)) {
            case 32768:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            case -1:
                v2["\n\t"] = v4
            }
            _ = fi40
        }
        for fi41 := 0; fi41 < 20; fi41++ {
            v42 := GenS2{F0: (v4 + ((("\n\t" + ("foo bar" + (((((v4 + v4) + ("\n\t" + "\n\t")) + "") + (v4 + (v4 + (v4 + ("" + v4))))) + (v4 + v4)))) + (v4 + v4)) + v4)), F1: -128}
            var v43 uint = 1
            _ = fi41
            _ = v42
            _ = v43
        }
        switch (-(fi39)) {
        case -419:
            v5 -= v5
        case 700:
            fmt.Println(v6, v6, v5)
        case -811:
            v3 = append(v3, 2)
        default:
            fmt.Println(v4)
        }
        _ = fi39
    }
    var v44 string = ("hello" + (((v4 + ((v4 + (v4 + (v4 + "hello"))) + "\n\t")) + v4) + v4))
    v6 += (^(v6))
    v44 = v44
    if (-1 != -199) {
        v5 = (v5 << ((752) & 63))
    } else {
        if true {
            var v45 bool = false
            if v45 {
                v46 := GenS1{F0: 1000, F1: (v5 >= 2), F2: -9223372036854775808}
                v45 = v45
                switch -1 {
                case 32768:
                    if v45 {
                        v4 = "a"
                        switch (v4 + v44) {
                        case "a":
                            if false {
                                _ = len(v3)
                                _ = len(v3)
                            } else {
                                v4 = (v44 + v44)
                            }
                        case "\n\t":
                            var v47 float32 = -2.25
                            _ = v47
                        }
                        v1[1] = (0 < 2147483648)
                    }
                default:
                    if v45 {
                        fmt.Println(v44, v44, v4)
                        v48 := GenS3{F0: 0, F1: (-32768 < -32768), F2: 1}
                        _ = v48
                    } else {
                        v5 -= v5
                        for fi49 := 0; fi49 < 19; fi49++ {
                            var v50 uint64 = 318
                            var v51 uint16 = v5
                            _ = fi49
                            _ = v50
                            _ = v51
                        }
                        fmt.Println(v6)
                    }
                }
                _ = v46
            } else {
                var v52 uint32 = 276
                switch "foo bar" {
                case "":
                    var v53 float32 = 0.1
                    _ = v53
                case "hello":
                    _ = len(v3)
                case "\n\t":
                    var v54 uint8 = 1
                    _ = v54
                }
                _ = v52
            }
            _ = v45
        } else {
            var v55 uint8 = 128
            _ = v55
        }
        v56 := v2[v4]
        delete(v2, v4)
        _ = v56
    }
    switch v44 {
    case "":
        if false {
            fmt.Println(v5, v4)
            switch v44 {
            case "hello":
                if ((0.0 == float64(1.5)) && false) {
                    v6 = v6
                }
            case "":
                v57 := GenS3{F0: 2147483648, F1: true, F2: 1}
                _ = v57
            case "a":
                if (678 != 2) {
                    if (!(0 < 370)) {
                        if (65535 < 521) {
                            v44 = v4
                            v6 = 4294967295
                            if true {
                                var v58 uint16 = (^(v5))
                                _ = v58
                            } else {
                                var v59 float64 = float64(1.5)
                                var v60 int = 42
                                v59 += (v59 + (3.14159 / 1.0))
                                _ = v59
                                _ = v60
                            }
                        }
                        v61 := GenS5{F0: -87, F1: -128}
                        var v62 float32 = 0.0
                        _ = v61
                        _ = v62
                    } else {
                        if true {
                            if (!false) {
                                for fi63 := 0; fi63 < 1; fi63++ {
                                    v64 := GenS6{F0: 869}
                                    fmt.Println(v4, v4)
                                    _ = fi63
                                    _ = v64
                                }
                                var v65 int32 = 452
                                v66 := GenS7{F0: 42, F1: -1}
                                _ = v65
                                _ = v66
                            } else {
                                v44 = "foo bar"
                            }
                        }
                        v67 := GenS8{F0: -100}
                        var v68 bool = false
                        _ = v67
                        _ = v68
                    }
                    if false {
                        v69 := GenS9{F0: (v6 % ((v6) | 1))}
                        v70 := v2[v44]
                        _ = v69
                        _ = v70
                    }
                    v71 := v2[v4]
                    _ = v71
                }
            }
            if len(v3) > 0 {
                v3[0] = 2
            }
        } else {
            var v72 float64 = 1.5
            v73 := GenS3{F0: 165, F1: (0.0 <= -2.25), F2: 127}
            _ = v72
            _ = v73
        }
    case "a":
        var v74 uint16 = v5
        _ = v74
    case "\n\t":
        if (v44 != v4) {
            if false {
                switch v44 {
                case "foo bar":
                    if (!(42 >= -40)) {
                        var v75 bool = true
                        v76 := GenS10{F0: uint(v5), F1: 0}
                        var v77 int = 754
                        _ = v75
                        _ = v76
                        _ = v77
                    }
                default:
                    v44 = ("hello" + v4)
                }
                var v78 bool = false
                var v79 int16 = -1
                _ = v78
                _ = v79
            } else {
                v6 += v6
            }
            v4 = (v44 + v4)
            v5 = 65535
        } else {
            fmt.Println(v6)
            v5 += 255
            v3 = append(v3, 703)
        }
    default:
        for fi80 := 0; fi80 < 20; fi80++ {
            v81 := GenS9{F0: (^(v6))}
            if len(v3) > 0 {
                v3[0] = 2147483647
            }
            _ = fi80
            _ = v81
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v44
    _ = v5
    _ = v6
}
