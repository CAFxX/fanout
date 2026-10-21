package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint32; F1 string }
type GenS4 struct { F0 uint; F1 int64; F2 uint64 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint32; F1 int16 }
func genFn0() int16 {
    var v6 uint32 = 2
    switch (("" + "hello") + "\n\t") {
    case "\n\t":
        for fi7 := 0; fi7 < 9; fi7++ {
            v8 := []int{}
            var v9 uint = 4294967295
            for fi10 := 0; fi10 < 13; fi10++ {
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                fmt.Println(v6, fi7, v9)
                switch (fi10 - 0) {
                case 127:
                    v6 += (^(v6))
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                case -1:
                    v9 = v9
                case -2147483648:
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                default:
                    v6 = (v6 % ((v6) | 1))
                }
                _ = fi10
            }
            _ = fi7
            _ = v8
            _ = v9
        }
    case "":
        for fi11 := 0; fi11 < 8; fi11++ {
            for fi12 := 0; fi12 < 7; fi12++ {
                var v13 int = (-(fi11))
                var v14 uint64 = 4294967295
                if ("hello" < ("\n\t" + ("hello" + ((("\n\t" + "\n\t") + ("hello" + (("" + "a") + "a"))) + "\n\t")))) {
                    fmt.Println(v6, fi11, v13)
                }
                var v15 uint32 = (v6 / (((v6 & v6)) | 1))
                v16 := GenS1{F0: 32767}
                _ = fi12
                _ = v13
                _ = v14
                _ = v15
                _ = v16
            }
            v6 -= v6
            var v17 string = "foo bar"
            _ = fi11
            _ = v17
        }
    case "hello":
        for fi18 := 0; fi18 < 15; fi18++ {
            for fi19 := 0; fi19 < 19; fi19++ {
                v20 := []int{fi19, (-(fi19)), (fi19 % (((^(fi19))) | 1))}
                if (v6 > 1) {
                    v6 += v6
                    var v21 int64 = 42
                    _ = v21
                }
                var v22 int8 = -111
                _ = fi19
                _ = v20
                _ = v22
            }
            _ = fi18
        }
    }
    switch 952 {
    case 32768:
        switch -2147483648 {
        case 32768:
            v23 := GenS2{F0: 2}
            _ = v23
        case 51:
            v6 -= (v6 & (v6 / ((v6) | 1)))
        default:
            v24 := []int{1, 1, 32768, 718}
            _ = v24
        }
    case -344:
        switch "hello" {
        case "\n\t":
            var v25 uint = 18446744073709551615
            _ = v25
        case "hello":
            for fi26 := 0; fi26 < 2; fi26++ {
                var v27 float32 = float32((-2.25 + 1e10))
                _ = fi26
                _ = v27
            }
        case "":
            v6 += (^(v6))
        }
    case -636:
        for fi28 := 0; fi28 < 11; fi28++ {
            v29 := []int{1, (fi28 << ((436) & 63)), 32767}
            for fi30 := 0; fi30 < 9; fi30++ {
                fmt.Println(fi28, v6)
                if false {
                    v6 -= v6
                } else {
                    v6 = (v6 << ((42) & 63))
                }
                _ = fi30
            }
            _ = fi28
            _ = v29
        }
    default:
        v31 := GenS1{F0: 1}
        _ = v31
    }
    _ = v6
    return 1000
}
func genFn1(p0 uint8, p1 int64, p2 uint8) int8 {
    var v32 int = -9223372036854775808
    p1 = (p1 | (p1 ^ (p1 >> ((65535) & 63))))
    for fi33 := 0; fi33 < 7; fi33++ {
        if ((217 > 479) || (593 != 2)) {
            fmt.Println(fi33, v32, v32)
            v34 := []int{}
            _ = v34
        }
        if (!(true || (1 != p1))) {
            for fi35 := 0; fi35 < 2; fi35++ {
                v36 := []int{fi35}
                fmt.Println(p0)
                switch ("a" + ("" + "a")) {
                case "a":
                    v32 = fi33
                case "":
                    fmt.Println(fi35, p2, fi33)
                case "\n\t":
                    _ = len(v36)
                default:
                    fmt.Println(v32, p2, v32)
                    v37 := GenS3{F0: 2, F1: ("a" + "\n\t")}
                    _ = v37
                }
                _ = fi35
                _ = v36
            }
        } else {
            var v38 int32 = 1
            var v39 uint64 = 679
            v40 := []int{fi33, v32}
            _ = v38
            _ = v39
            _ = v40
        }
        if true {
            v41 := []int{(-(fi33)), (v32 & (v32 | 4294967295)), (fi33 >> ((51) & 63)), (v32 ^ fi33)}
            _ = v41
        }
        _ = fi33
    }
    p2 -= (^(p0))
    _ = p0
    _ = p1
    _ = p2
    _ = v32
    return -46
}

func main() {
    var v1 uint32 = 42
    var v2 uint16 = 32767
    v3 := make(map[string]string)
    var v4 uint = 4294967295
    var v5 int64 = 42
    v1 += 2147483648
    switch "hello" {
    case "\n\t":
        v5 += (v5 >> (((v4 | v4)) & 63))
    default:
        var v42 string = ""
        _ = v42
    }
    for fi43 := 0; fi43 < 10; fi43++ {
        for fi44 := 0; fi44 < 16; fi44++ {
            v4 -= (v4 >> ((255) & 63))
            var v45 int16 = -32768
            switch "foo bar" {
            case "hello":
                var v46 uint = 42
                _ = v46
            case "":
                var v47 string = "foo bar"
                _ = v47
            }
            _ = fi44
            _ = v45
        }
        v48 := v3[("\n\t" + "")]
        _ = fi43
        _ = v48
    }
    if false {
        v5 -= v5
        v2 -= v2
        for fi49 := 0; fi49 < 13; fi49++ {
            if true {
                _ = len(v3)
                for fi50 := 0; fi50 < 20; fi50++ {
                    v51 := []int{fi50}
                    if (false && false) {
                        var v52 uint8 = 128
                        v2 = v2
                        var v53 int8 = 42
                        _ = v52
                        _ = v53
                    } else {
                        switch ("hello" + "a") {
                        case "hello":
                            v5 = v5
                        default:
                            v54 := GenS4{F0: (v4 / ((65535) | 1)), F1: 850, F2: 717}
                            _ = v54
                        }
                    }
                    var v55 uint8 = uint8(v1)
                    var v56 int = fi50
                    if true {
                        var v57 int32 = 2147483647
                        v1 -= 221
                        _ = v57
                    } else {
                        v56 = v56
                        var v58 uint32 = v1
                        _ = v58
                    }
                    var v59 uint8 = (v55 | (^(v55)))
                    var v60 bool = true
                    if v60 {
                        if len(v51) >= 2 {
                            v51 = v51[1:len(v51)]
                        }
                        if true {
                            var v61 bool = v60
                            _ = v61
                        } else {
                            v55 = v55
                            delete(v3, "a")
                            var v62 int64 = v5
                            switch fi50 {
                            case 32768:
                                if (((127 < -98) || v60) && (("\n\t" > ("" + "hello")) || (!v60))) {
                                    var v63 float32 = 3.14159
                                    v64 := GenS2{F0: (v4 >> (((^(v4))) & 63))}
                                    var v65 int32 = 0
                                    _ = v63
                                    _ = v64
                                    _ = v65
                                }
                            case -1:
                                v5 -= (v5 * (v5 * v62))
                            case -2147483648:
                                if (!true) {
                                    var v66 float64 = (0.1 * 0.1)
                                    var v67 int32 = 721
                                    v60 = v60
                                    _ = v66
                                    _ = v67
                                } else {
                                    if ((v1 << ((v4) & 63)) >= v1) {
                                        v56 = v56
                                        _ = len(v3)
                                    } else {
                                        var v68 uint32 = (^(v1))
                                        v56 = -9223372036854775808
                                        _ = v68
                                    }
                                    if v60 {
                                        if v60 {
                                            if true {
                                                if v60 {
                                                    delete(v3, ((("\n\t" + "\n\t") + ((("foo bar" + "") + "hello") + "\n\t")) + "hello"))
                                                }
                                            }
                                            v69 := GenS3{F0: 2147483648, F1: "\n\t"}
                                            var v70 int = (fi49 % (((-(fi49))) | 1))
                                            _ = v69
                                            _ = v70
                                        }
                                    } else {
                                        fmt.Println(fi49, v5)
                                        var v71 uint16 = (v2 | (v2 + v2))
                                        var v72 string = ("" + "hello")
                                        _ = v71
                                        _ = v72
                                    }
                                }
                            default:
                                v59 = (v59 / (((^(v59))) | 1))
                            }
                            _ = v62
                        }
                        if ("a" != "hello") {
                            if v60 {
                                v59 += (^(v55))
                                v4 = 2
                                var v73 uint64 = 42
                                _ = v73
                            }
                        }
                        if false {
                            v5 = (v5 ^ (v5 >> ((uint(v2)) & 63)))
                        } else {
                            v74 := GenS3{F0: (v1 / ((v1) | 1)), F1: "\n\t"}
                            _ = v74
                        }
                        if v60 {
                            var v75 uint64 = 621
                            var v76 bool = (!(genFn0() < 1000))
                            _ = v75
                            _ = v76
                        } else {
                            v55 += 15
                            var v77 uint16 = 0
                            v51 = append(v51, int(v5))
                            _ = v77
                        }
                    } else {
                        var v78 bool = (-1 != (v5 | (v5 / ((9223372036854775807) | 1))))
                        var v79 int16 = -1
                        _ = v78
                        _ = v79
                    }
                    switch -2147483648 {
                    case 4294967295:
                        if len(v51) > 0 {
                            v51[0] = (fi49 * (fi50 / (((fi49 << ((uint(v2)) & 63))) | 1)))
                        }
                    case -435:
                        v80 := GenS5{F0: 866}
                        _ = v80
                    case -812:
                        fmt.Println(fi49, v55)
                    }
                    v56 = (v56 << (((^(v4))) & 63))
                    v81 := GenS6{F0: v55}
                    _ = fi50
                    _ = v51
                    _ = v55
                    _ = v56
                    _ = v59
                    _ = v60
                    _ = v81
                }
                var v82 int = fi49
                _ = v82
            } else {
                if (("\n\t" <= "a") || false) {
                    _ = len(v3)
                } else {
                    v83 := []int{}
                    _ = v83
                }
                v5 -= (v5 % (((v5 * (v5 * (v5 | v5)))) | 1))
            }
            v84 := []int{(fi49 + fi49), int(v4), (fi49 - 1), fi49}
            _ = fi49
            _ = v84
        }
    } else {
        v85 := []int{-357, 9223372036854775807}
        v86 := GenS1{F0: -152}
        _ = v85
        _ = v86
    }
    v1 += (v1 ^ v1)
    v5 = 0
    var v87 float32 = float32(1.5)
    v88 := v3["\n\t"]
    switch v88 {
    case "foo bar":
        var v89 uint8 = 255
        _ = v89
    case "\n\t":
        v4 -= v4
    default:
        if (true || (-996 <= 127)) {
            v88 = (v88 + v88)
            v90 := []int{128}
            switch v88 {
            case "a":
                v88 = v88
            case "":
                var v91 int8 = -1
                _ = v91
            default:
                v92 := GenS4{F0: (v4 & v4), F1: v5, F2: 824}
                _ = v92
            }
            _ = v90
        } else {
            v2 -= (^(v2))
            switch 128 {
            case 128:
                for fi93 := 0; fi93 < 5; fi93++ {
                    fmt.Println(v4, v1)
                    _ = fi93
                }
            case 0:
                v94 := GenS3{F0: (v1 + (v1 >> ((v4) & 63))), F1: v88}
                _ = v94
            case 4294967295:
                for fi95 := 0; fi95 < 3; fi95++ {
                    v96 := GenS7{F0: 2, F1: 626}
                    _ = fi95
                    _ = v96
                }
            default:
                v2 -= (v2 << ((v4) & 63))
            }
        }
    }
    v97 := GenS4{F0: v4, F1: v5, F2: 42}
    if true {
        v2 += uint16(v4)
        fmt.Println(v2, v5, v87)
    }
    v98 := GenS3{F0: (v1 & (^(v1))), F1: "foo bar"}
    v99 := []int{-847, 820}
    var v100 float64 = (3.14159 - float64(((((v87 - -2.25) - (1.5 * (-2.25 - v87))) * v87) / ((3.14159 * ((((((v87 * v87) + v87) + 1.5) + v87) + v87) + v87)) - v87))))
    fmt.Println(v4, v87, v1)
    v4 = (^(v4))
    fmt.Println("v1:", v1)
    fmt.Println("v100:", v100)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v87:", v87)
    fmt.Println("v88:", v88)
    fmt.Println("slicelen v99:", len(v99), v99)
    _ = v1
    _ = v100
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v87
    _ = v88
    _ = v97
    _ = v98
    _ = v99
}
