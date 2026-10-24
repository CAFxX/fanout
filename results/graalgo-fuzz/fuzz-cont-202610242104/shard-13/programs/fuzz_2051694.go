package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 bool }
type GenS2 struct { F0 uint8; F1 uint16 }
type GenS3 struct { F0 uint64; F1 uint16 }
type GenS4 struct { F0 string; F1 int8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint8; F1 uint16; F2 int64 }
func genFn0(p0 uint32, p1 int8) int64 {
    for fi7 := 0; fi7 < 12; fi7++ {
        for fi8 := 0; fi8 < 11; fi8++ {
            fmt.Println(p1)
            var v9 int32 = 2147483647
            v10 := GenS1{F0: false, F1: 42, F2: false}
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi7
    }
    p0 -= p0
    _ = p0
    _ = p1
    return 672
}
func genFn1(p0 string) uint8 {
    v11 := GenS1{F0: (!true), F1: 279, F2: (false && true)}
    for fi12 := 0; fi12 < 15; fi12++ {
        var v13 int16 = 0
        fmt.Println(v13, v13, p0)
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = v11
    return 128
}

func main() {
    var v1 bool = false
    v2 := make(map[int]int)
    v3 := make(map[uint32]int)
    v4 := make(map[string]string)
    var v5 int16 = 1
    var v6 int64 = 9223372036854775807
    v14 := []int{}
    var v15 uint8 = 255
    for fi16 := 0; fi16 < 16; fi16++ {
        v6 = v6
        v17 := GenS1{F0: (true && (((("" + "") + "foo bar") != "") && (942 < 1))), F1: 2, F2: true}
        var v18 float64 = 0.0
        _ = fi16
        _ = v17
        _ = v18
    }
    v5 = v5
    var v19 string = "foo bar"
    switch -1 {
    case 128:
        v20 := GenS1{F0: v1, F1: 980, F2: ((0.1 + 1e10) <= (float64(float32(3.14159)) + float64((-2.25 * 1e10))))}
        _ = v20
    case 4294967295:
        for fi21 := 0; fi21 < 4; fi21++ {
            for fi22 := 0; fi22 < 2; fi22++ {
                var v23 float64 = 1.5
                _ = fi22
                _ = v23
            }
            switch (v19 + v19) {
            case "a":
                var v24 float64 = 0.1
                _ = v24
            }
            _ = fi21
        }
    case -2147483648:
        _ = len(v14)
    }
    var v25 uint64 = 442
    for fi26 := 0; fi26 < 12; fi26++ {
        var v27 int16 = (v5 ^ v5)
        v28 := GenS1{F0: false, F1: 939, F2: (-2.25 <= 0.0)}
        if len(v14) > 0 {
            v14[0] = (fi26 + fi26)
        }
        _ = fi26
        _ = v27
        _ = v28
    }
    var v29 float64 = -2.25
    if (false || false) {
        if (false && ((v6 ^ (v6 | (v6 << ((18446744073709551615) & 63)))) > v6)) {
            fmt.Println(v1, v5)
            for fi30 := 0; fi30 < 4; fi30++ {
                var v31 float64 = (0.1 + v29)
                v25 = (v25 & 18446744073709551615)
                _ = fi30
                _ = v31
            }
        }
        switch 794 {
        case 32768:
            v32 := v2[-1]
            _ = v32
        case -1:
            if len(v14) >= 2 {
                v14 = v14[1:len(v14)]
            }
        default:
            fmt.Println(v25, v15)
        }
        v19 = ((v19 + v19) + v19)
    }
    v33 := GenS1{F0: ((false && v1) && false), F1: 255, F2: v1}
    switch (v19 + "foo bar") {
    case "":
        switch -2147483648 {
        case 330:
            var v34 string = "\n\t"
            _ = v34
        case -901:
            for fi35 := 0; fi35 < 8; fi35++ {
                if (-9223372036854775808 < (fi35 * (fi35 % (((-(fi35))) | 1)))) {
                    v36 := GenS2{F0: v15, F1: 65535}
                    if v1 {
                        v6 -= v6
                    }
                    _ = v36
                }
                v37 := GenS3{F0: (v25 / ((4294967295) | 1)), F1: 405}
                _ = fi35
                _ = v37
            }
        case 2147483647:
            if ((v6 >> ((234) & 63)) > (^(v6))) {
                switch v19 {
                case "":
                    v15 += (^(v15))
                case "\n\t":
                    var v38 int32 = 1
                    _ = v38
                case "a":
                    v39 := GenS4{F0: (v19 + v19), F1: -128}
                    _ = v39
                }
                if (int16(v6) <= (v5 & v5)) {
                    v40 := GenS3{F0: v25, F1: 121}
                    v41 := GenS5{F0: 1}
                    _ = v40
                    _ = v41
                } else {
                    v1 = true
                }
                fmt.Println(v5, v15)
            } else {
                for fi42 := 0; fi42 < 9; fi42++ {
                    if true {
                        if len(v14) >= 2 {
                            v14 = v14[1:len(v14)]
                        }
                        delete(v3, 4294967295)
                    }
                    _ = fi42
                }
                var v43 int16 = v5
                v15 = (^(v15))
                _ = v43
            }
        }
    case "hello":
        if (537 != 2147483647) {
            for fi44 := 0; fi44 < 12; fi44++ {
                fmt.Println(v6, fi44)
                _ = fi44
            }
            for fi45 := 0; fi45 < 18; fi45++ {
                v46 := GenS6{F0: 127}
                if ((!(v25 != 466)) || (v1 && true)) {
                    var v47 uint16 = 699
                    _ = v47
                } else {
                    _ = len(v4)
                    for fi48 := 0; fi48 < 15; fi48++ {
                        var v49 uint32 = 2147483648
                        var v50 float32 = 0.1
                        v50 -= (v50 / v50)
                        _ = fi48
                        _ = v49
                        _ = v50
                    }
                }
                for fi51 := 0; fi51 < 15; fi51++ {
                    if v1 {
                        v1 = (v19 >= v19)
                    } else {
                        v1 = (!(!v1))
                    }
                    var v52 bool = v1
                    _ = fi51
                    _ = v52
                }
                _ = fi45
                _ = v46
            }
            v53 := GenS2{F0: (v15 * v15), F1: 793}
            _ = v53
        } else {
            v25 += v25
            for fi54 := 0; fi54 < 17; fi54++ {
                fmt.Println(v5, v29)
                for fi55 := 0; fi55 < 5; fi55++ {
                    if (v1 && true) {
                        var v56 int = int(v6)
                        _ = v56
                    }
                    _ = fi55
                }
                v6 += (v6 | v6)
                _ = fi54
            }
            if len(v14) >= 2 {
                v14 = v14[1:len(v14)]
            }
        }
    default:
        var v57 uint32 = 878
        _ = v57
    }
    v25 += (v25 & v25)
    if v1 {
        var v58 bool = v1
        for fi59 := 0; fi59 < 6; fi59++ {
            for fi60 := 0; fi60 < 9; fi60++ {
                v15 = (v15 & v15)
                v25 = 18446744073709551615
                v1 = v58
                _ = fi60
            }
            for fi61 := 0; fi61 < 4; fi61++ {
                var v62 string = "foo bar"
                _ = fi61
                _ = v62
            }
            v63 := GenS6{F0: 127}
            _ = fi59
            _ = v63
        }
        if (207 > -2147483648) {
            if v58 {
                var v64 uint = 65535
                _ = v64
            } else {
                v65 := GenS4{F0: (v19 + v19), F1: 42}
                _ = v65
            }
        } else {
            fmt.Println(v19, v5, v15)
        }
        _ = v58
    } else {
        fmt.Println(v1, v1, v19)
    }
    if (int(v6) < -118) {
        v66 := GenS7{F0: uint8(v25), F1: 2, F2: v6}
        if len(v14) > 0 {
            v14[0] = 44
        }
        switch 1 {
        case 782:
            if (v5 < (v5 % ((v5) | 1))) {
                v5 = v5
                v14 = append(v14, 38)
            } else {
                var v67 int64 = 9223372036854775807
                v6 = (v67 << ((135) & 63))
                _ = v67
            }
        }
        _ = v66
    } else {
        v68 := GenS3{F0: (v25 + v25), F1: 228}
        _ = v68
    }
    if len(v14) > 0 {
        v14[0] = 4294967295
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
