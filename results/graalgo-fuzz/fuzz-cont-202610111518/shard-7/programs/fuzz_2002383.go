package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 uint8; F1 int }
type GenS3 struct { F0 int32; F1 uint8; F2 int }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int16; F1 uint64; F2 bool }
type GenS7 struct { F0 uint16; F1 bool }
type GenS8 struct { F0 int; F1 uint; F2 uint16 }
type GenS9 struct { F0 uint8 }
func genFn0(p0 bool, p1 int, p2 bool) int {
    v6 := []int{}
    if true {
        v7 := GenS1{F0: 1, F1: 1}
        var v8 uint = 0
        _ = v7
        _ = v8
    } else {
        p1 = p1
    }
    for fi9 := 0; fi9 < 6; fi9++ {
        for fi10 := 0; fi10 < 3; fi10++ {
            v11 := GenS2{F0: 42, F1: (p1 >> ((0) & 63))}
            _ = fi10
            _ = v11
        }
        for fi12 := 0; fi12 < 19; fi12++ {
            if (p0 || false) {
                _ = len(v6)
                fmt.Println(fi9)
                if p2 {
                    var v13 int16 = 42
                    _ = v13
                }
                p2 = p0
                p0 = p0
                if p2 {
                    p2 = p2
                    v14 := GenS3{F0: 955, F1: 142, F2: 2147483647}
                    p1 -= 0
                    _ = v14
                } else {
                    v15 := GenS4{F0: 255}
                    _ = v15
                }
                var v16 string = "\n\t"
                if (0 == 964) {
                    v16 = (v16 + v16)
                } else {
                    v17 := GenS1{F0: 2, F1: 114}
                    _ = v17
                }
                p1 = fi12
                _ = v16
            }
            _ = fi12
        }
        for fi18 := 0; fi18 < 5; fi18++ {
            v6 = append(v6, 2147483647)
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            p0 = (false && false)
            _ = fi18
        }
        _ = fi9
    }
    p1 = (p1 << ((1) & 63))
    v19 := GenS1{F0: 238, F1: 253}
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    _ = v6
    return 0
}
func genFn1(p0 int64, p1 bool) int64 {
    v20 := GenS5{F0: 0}
    for fi21 := 0; fi21 < 14; fi21++ {
        var v22 uint32 = 1
        _ = fi21
        _ = v22
    }
    _ = p0
    _ = p1
    _ = v20
    return (p0 % (((p0 + p0)) | 1))
}

func main() {
    var v1 int32 = -761
    v2 := []int{127, 32767, 310}
    v3 := make(map[int]int)
    var v4 int32 = -2147483648
    var v5 uint8 = 42
    v2 = append(v2, 127)
    var v23 int = 362
    if true {
        for fi24 := 0; fi24 < 10; fi24++ {
            for fi25 := 0; fi25 < 20; fi25++ {
                fmt.Println(v1)
                switch "a" {
                case "":
                    v4 -= v1
                case "\n\t":
                    switch 112 {
                    case 845:
                        v4 += (v4 & v1)
                    case -2147483648:
                        if false {
                            _ = len(v3)
                        } else {
                            if (4294967295 > 2147483648) {
                                v2 = append(v2, -1)
                            }
                            delete(v3, (v23 % (((fi24 << ((2) & 63))) | 1)))
                        }
                    case 32768:
                        var v26 int16 = 490
                        _ = v26
                    default:
                        fmt.Println(fi24, fi25)
                    }
                case "a":
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                default:
                    v4 -= 854
                }
                if len(v2) > 0 {
                    v2[0] = genFn0((125 <= 100), 42, ((v4 ^ (v4 - v4)) > (v1 | -1)))
                }
                if true {
                    var v27 uint = 18446744073709551615
                    switch 134 {
                    case -2147483648:
                        v23 = v23
                    case -209:
                        if (976 != 65535) {
                            var v28 int16 = 32767
                            if false {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                if false {
                                    var v29 int32 = -544
                                    fmt.Println(v29, v28, v5)
                                    v1 += (-(v1))
                                    v23 = 1
                                    var v30 int = 32768
                                    _ = v29
                                    _ = v30
                                } else {
                                    var v31 uint32 = 0
                                    v31 = 65535
                                    var v32 int16 = (v28 % ((v28) | 1))
                                    if false {
                                        if false {
                                            v32 = v28
                                            fmt.Println(fi24, v4)
                                            if (!true) {
                                                var v33 int8 = -113
                                                var v34 uint16 = 42
                                                _ = v33
                                                _ = v34
                                            }
                                        } else {
                                            v28 = (v32 * (v28 / ((v32) | 1)))
                                            v3[(v23 % ((int(v31)) | 1))] = (fi25 | (fi24 ^ (^(fi25))))
                                        }
                                        if true {
                                            v35 := GenS6{F0: int16(v23), F1: 0, F2: (!(42 >= genFn1(-9223372036854775808, (true || (true && (v31 >= 0))))))}
                                            v28 = -364
                                            var v36 uint32 = (^(v31))
                                            _ = v35
                                            _ = v36
                                        }
                                    }
                                    var v37 bool = false
                                    switch -9223372036854775808 {
                                    case 0:
                                        var v38 uint8 = (v5 & (v5 * (v5 * (^(v5)))))
                                        _ = v38
                                    default:
                                        var v39 int32 = (v4 | (v1 / ((v1) | 1)))
                                        _ = v39
                                    }
                                    var v40 float64 = 0.0
                                    _ = v31
                                    _ = v32
                                    _ = v37
                                    _ = v40
                                }
                                fmt.Println(fi24)
                            } else {
                                v4 += (v1 << ((v27) & 63))
                                var v41 uint32 = 855
                                var v42 uint = v27
                                v2 = append(v2, int(v4))
                                _ = v41
                                _ = v42
                            }
                            _ = v28
                        } else {
                            if true {
                                if ((v27 - 1) == v27) {
                                    v43 := GenS7{F0: 594, F1: false}
                                    var v44 int32 = (v1 << ((uint(v5)) & 63))
                                    _ = v43
                                    _ = v44
                                }
                            } else {
                                v5 = 88
                                if ((v1 >> (((v27 >> (((v27 + (v27 % ((42) | 1)))) & 63))) & 63)) == (v4 >> (((^(v27))) & 63))) {
                                    fmt.Println(fi24, v1, fi24)
                                    v45 := GenS1{F0: 0, F1: v5}
                                    fmt.Println(v23, fi24, v4)
                                    _ = v45
                                } else {
                                    var v46 string = "\n\t"
                                    v47 := GenS2{F0: (v5 & (^(v5))), F1: (fi25 << ((v27) & 63))}
                                    switch (-(fi24)) {
                                    case 32767:
                                        var v48 int = (v23 >> ((65535) & 63))
                                        _ = v48
                                    case -2147483648:
                                        v1 -= 793
                                    case 626:
                                        v2 = append(v2, fi25)
                                    }
                                    _ = v46
                                    _ = v47
                                }
                                var v49 int16 = -32768
                                _ = v49
                            }
                            v3[fi24] = v23
                        }
                    case 334:
                        switch (v23 + genFn0((true && true), (v23 ^ 127), false)) {
                        case 32768:
                            var v50 float64 = 0.1
                            _ = v50
                        default:
                            v27 = 2
                        }
                        var v51 float64 = (((float64(3.14159) * -2.25) * (float64(3.14159) + float64(1e10))) / float64((1.5 / 1.5)))
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = v51
                    }
                    var v52 bool = (!((!((-(fi25)) == v23)) && ((v5 >> (((v27 - 255)) & 63)) < 42)))
                    _ = v27
                    _ = v52
                } else {
                    var v53 bool = (((fi25 >> ((65535) & 63)) < fi24) || (int64(v5) <= 1))
                    v4 += (^(v1))
                    fmt.Println(v1, v5, v23)
                    _ = v53
                }
                fmt.Println(v1, fi24)
                _ = fi25
            }
            if true {
                for fi54 := 0; fi54 < 19; fi54++ {
                    v5 += (^(v5))
                    var v55 int16 = -238
                    _ = fi54
                    _ = v55
                }
            }
            fmt.Println(fi24)
            _ = fi24
        }
    }
    v2 = append(v2, v23)
    v23 += 32768
    for fi56 := 0; fi56 < 11; fi56++ {
        for fi57 := 0; fi57 < 19; fi57++ {
            if ("\n\t" != "") {
                var v58 float32 = 3.14159
                fmt.Println(v4, v4, v4)
                _ = v58
            } else {
                if (!(0.1 < (3.14159 / 1.5))) {
                    switch (fi56 - (v23 + v23)) {
                    case 42:
                        var v59 float32 = (0.0 + (float32((0.1 - 1e10)) - -2.25))
                        _ = v59
                    }
                    v2 = append(v2, (v23 % (((fi56 - fi56)) | 1)))
                    v23 = (fi57 % (((v23 << ((255) & 63))) | 1))
                } else {
                    if true {
                        v4 = (v1 * v1)
                    }
                    v23 = (fi57 | (v23 ^ (fi57 << ((4294967295) & 63))))
                    v60 := GenS8{F0: (fi57 | (fi57 - 2147483647)), F1: 4294967295, F2: 255}
                    _ = v60
                }
            }
            v4 = (v4 << ((207) & 63))
            _ = fi57
        }
        var v61 float32 = -2.25
        if (100 != 100) {
            v5 += v5
            v2 = append(v2, v23)
        } else {
            switch (fi56 >> ((18446744073709551615) & 63)) {
            case 32767:
                for fi62 := 0; fi62 < 10; fi62++ {
                    v63 := GenS6{F0: 42, F1: 2, F2: (false || (817 < 2))}
                    var v64 float32 = v61
                    v3[(v23 + genFn0(false, 128, (!(0 > 9223372036854775807))))] = fi56
                    _ = fi62
                    _ = v63
                    _ = v64
                }
            case 0:
                if len(v2) > 0 {
                    v2[0] = (-(fi56))
                }
            case -1:
                fmt.Println(v23, fi56, v4)
            default:
                v65 := GenS3{F0: (v1 % ((v1) | 1)), F1: 155, F2: -928}
                _ = v65
            }
        }
        _ = fi56
        _ = v61
    }
    if len(v2) > 0 {
        v2[0] = -846
    }
    v1 -= (v1 % (((v1 + v4)) | 1))
    if (!false) {
        fmt.Println(v4, v5)
    } else {
        var v66 float32 = 1.5
        v67 := v3[-2147483648]
        fmt.Println(v5, v1)
        _ = v66
        _ = v67
    }
    var v68 int = v23
    v5 = v5
    v69 := GenS9{F0: (v5 / ((v5) | 1))}
    var v70 string = ""
    for fi71 := 0; fi71 < 17; fi71++ {
        for fi72 := 0; fi72 < 7; fi72++ {
            switch v70 {
            case "a":
                var v73 uint64 = 92
                _ = v73
            }
            _ = fi72
        }
        _ = fi71
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v68:", v68)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v68
    _ = v69
    _ = v70
}
