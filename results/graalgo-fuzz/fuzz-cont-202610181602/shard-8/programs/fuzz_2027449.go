package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 int16; F2 int32 }
type GenS3 struct { F0 int8; F1 int64; F2 uint64 }
type GenS4 struct { F0 string; F1 int16 }
type GenS5 struct { F0 uint32; F1 int32 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int; F1 int; F2 string }
func genFn0(p0 string, p1 bool, p2 int64) int64 {
    v4 := GenS1{F0: 3}
    var v5 uint8 = 188
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 4
}
func genFn1() string {
    var v6 uint64 = 4294967295
    fmt.Println(v6)
    v6 += 1
    v6 = 18446744073709551615
    _ = v6
    return "\n\t"
}

func main() {
    v1 := make(map[int]string)
    var v2 uint = 880
    var v3 string = "\n\t"
    v2 = 611
    v3 = v3
    v3 = v3
    var v7 int = 73
    switch v7 {
    case 128:
        fmt.Println(v2, v3)
    case -115:
        for fi8 := 0; fi8 < 11; fi8++ {
            if (false || (2 >= 4294967295)) {
                for fi9 := 0; fi9 < 18; fi9++ {
                    v10 := GenS1{F0: 42}
                    _ = fi9
                    _ = v10
                }
                v11 := GenS2{F0: genFn0(("a" + v3), (128 == 128), 254), F1: 600, F2: 0}
                var v12 int = v7
                _ = v11
                _ = v12
            } else {
                v7 -= fi8
                var v13 int16 = 42
                _ = v13
            }
            v14 := []int{}
            _ = fi8
            _ = v14
        }
    default:
        var v15 int16 = -76
        _ = v15
    }
    switch "" {
    case "hello":
        switch v7 {
        case 9223372036854775807:
            v16 := []int{v7, v7}
            _ = v16
        case 0:
            for fi17 := 0; fi17 < 8; fi17++ {
                v18 := GenS3{F0: 120, F1: -9223372036854775808, F2: 412}
                _ = fi17
                _ = v18
            }
        default:
            if true {
                v3 = ((v3 + "foo bar") + "")
            }
        }
    case "\n\t":
        var v19 float64 = ((1e10 - 1.5) - 3.14159)
        _ = v19
    default:
        var v20 int32 = -1
        _ = v20
    }
    if true {
        for fi21 := 0; fi21 < 4; fi21++ {
            fmt.Println(v3, v2)
            v2 -= (v2 | (v2 * (^(v2))))
            var v22 uint = v2
            _ = fi21
            _ = v22
        }
        var v23 uint64 = 0
        fmt.Println(v3)
        _ = v23
    } else {
        v24 := GenS1{F0: 255}
        v25 := GenS2{F0: 1, F1: 720, F2: 0}
        var v26 uint32 = 141
        _ = v24
        _ = v25
        _ = v26
    }
    for fi27 := 0; fi27 < 15; fi27++ {
        var v28 int8 = 0
        if (!(1.5 != 0.0)) {
            for fi29 := 0; fi29 < 18; fi29++ {
                v30 := v1[v7]
                _ = fi29
                _ = v30
            }
            v31 := GenS2{F0: 1, F1: 1000, F2: -126}
            switch 128 {
            case 238:
                switch 128 {
                case -1:
                    v32 := GenS2{F0: 0, F1: 32767, F2: 0}
                    _ = v32
                case 0:
                    v2 -= v2
                default:
                    if (42 > 0) {
                        fmt.Println(v7, v2)
                        v33 := GenS4{F0: (v3 + ((v3 + v3) + ((v3 + "hello") + v3))), F1: -285}
                        v28 -= (v28 * (v28 % ((v28) | 1)))
                        _ = v33
                    }
                }
            default:
                v34 := GenS2{F0: 9223372036854775807, F1: int16(fi27), F2: 1}
                _ = v34
            }
            _ = v31
        } else {
            var v35 string = v3
            _ = v35
        }
        var v36 int32 = 0
        _ = fi27
        _ = v28
        _ = v36
    }
    for fi37 := 0; fi37 < 6; fi37++ {
        var v38 int32 = -277
        fmt.Println(v7, v2, v3)
        _ = fi37
        _ = v38
    }
    var v39 string = (v3 + v3)
    if (float32(0.0) != 0.0) {
        if false {
            fmt.Println(v2, v39, v3)
            var v40 uint32 = 0
            var v41 uint = 1
            _ = v40
            _ = v41
        } else {
            fmt.Println(v7)
            v42 := GenS3{F0: 1, F1: -9223372036854775808, F2: 2}
            v2 += (v2 - (v2 | (^(v2))))
            _ = v42
        }
        if false {
            fmt.Println(v2, v7)
            fmt.Println(v7, v7)
            v43 := []int{}
            _ = v43
        }
        v44 := GenS2{F0: genFn0((v3 + "a"), true, 0), F1: 1, F2: -1}
        _ = v44
    } else {
        var v45 int = (v7 / (((v7 << ((v2) & 63))) | 1))
        var v46 int8 = -54
        var v47 int64 = 968
        _ = v45
        _ = v46
        _ = v47
    }
    if true {
        v48 := []int{v7, (v7 - v7), v7, v7}
        if false {
            var v49 int32 = -2
            for fi50 := 0; fi50 < 9; fi50++ {
                for fi51 := 0; fi51 < 19; fi51++ {
                    v49 -= v49
                    switch (^(v7)) {
                    case -1:
                        v39 = (v39 + "\n\t")
                        var v52 float64 = float64((-2.25 + 1e10))
                        if true {
                            delete(v1, -2147483648)
                        }
                        _ = v52
                    }
                    _ = fi51
                }
                _ = fi50
            }
            _ = v49
        } else {
            if len(v48) > 0 {
                v48[0] = 1
            }
            for fi53 := 0; fi53 < 2; fi53++ {
                for fi54 := 0; fi54 < 4; fi54++ {
                    v39 = v39
                    fmt.Println(v2, fi53)
                    if true {
                        v7 = fi54
                        v3 = (((((v3 + v3) + "hello") + (v3 + "a")) + v39) + ("\n\t" + ""))
                    } else {
                        var v55 int = v7
                        var v56 uint32 = 2147483647
                        if (true || false) {
                            var v57 uint32 = v56
                            if (127 < 2) {
                                switch 785 {
                                case -2147483648:
                                    switch v55 {
                                    case 1:
                                        switch v39 {
                                        case "foo bar":
                                            switch (v3 + ("\n\t" + v39)) {
                                            case "hello":
                                                var v58 uint16 = 228
                                                _ = v58
                                            case "\n\t":
                                                if ((255 < 0) && false) {
                                                    var v59 int16 = 1000
                                                    var v60 uint32 = (v57 << (((v2 + 481)) & 63))
                                                    _ = v59
                                                    _ = v60
                                                } else {
                                                    v61 := GenS5{F0: 2, F1: 42}
                                                    _ = v61
                                                }
                                                v62 := v1[(v7 & 128)]
                                                _ = v62
                                            }
                                        }
                                    case 0:
                                        if false {
                                            v39 = v39
                                            v3 = v3
                                        }
                                    case 2147483647:
                                        var v63 uint64 = 994
                                        _ = v63
                                    default:
                                        fmt.Println(fi54, v39)
                                    }
                                default:
                                    var v64 uint8 = 42
                                    _ = v64
                                }
                                switch v3 {
                                case "foo bar":
                                    v48 = append(v48, 4294967295)
                                case "a":
                                    var v65 float32 = -2.25
                                    _ = v65
                                }
                                v66 := GenS6{F0: "hello"}
                                _ = v66
                            } else {
                                fmt.Println(v2)
                            }
                            var v67 uint8 = 1
                            _ = v57
                            _ = v67
                        }
                        _ = v55
                        _ = v56
                    }
                    switch fi54 {
                    case 9223372036854775807:
                        v68 := GenS4{F0: v3, F1: -1}
                        _ = v68
                    case 1:
                        v1[9223372036854775807] = v3
                    default:
                        v69 := GenS7{F0: (v3 + v3)}
                        _ = v69
                    }
                    switch v3 {
                    case "":
                        v48 = append(v48, (fi54 & -1))
                    case "a":
                        v39 = v3
                    case "\n\t":
                        v3 = v3
                        switch v39 {
                        case "":
                            _ = len(v48)
                        case "a":
                            v70 := GenS5{F0: uint32(v2), F1: 42}
                            _ = v70
                        default:
                            if (!true) {
                                if true {
                                    fmt.Println(fi54, v3)
                                    v39 = v39
                                } else {
                                    fmt.Println(fi54, fi53)
                                }
                                var v71 uint64 = 2
                                v7 = fi53
                                _ = v71
                            } else {
                                v1[(^(v7))] = ("" + v39)
                                var v72 uint16 = 0
                                _ = v72
                            }
                        }
                        v73 := GenS8{F0: -9223372036854775808, F1: -644, F2: v39}
                        _ = v73
                    default:
                        var v74 bool = false
                        _ = v74
                    }
                    _ = fi54
                }
                v39 = "a"
                v2 -= 0
                _ = fi53
            }
        }
        var v75 string = v3
        _ = v48
        _ = v75
    }
    if (false || (float32(float64(0.0)) >= 1e10)) {
        switch v7 {
        case 0:
            var v76 float64 = 0.0
            _ = v76
        case 4294967295:
            v77 := GenS8{F0: v7, F1: (v7 + v7), F2: "hello"}
            _ = v77
        }
        var v78 string = (v3 + v39)
        _ = len(v1)
        _ = v78
    } else {
        v3 = v3
        if (0 != -1) {
            switch v7 {
            case 9223372036854775807:
                if (2147483647 > 2147483647) {
                    var v79 int16 = 32767
                    v3 = (v39 + v39)
                    _ = v79
                } else {
                    var v80 float64 = (((float64(1.5) + 0.1) / 3.14159) - float64(float32(0.1)))
                    v81 := []int{v7}
                    _ = v80
                    _ = v81
                }
            }
            for fi82 := 0; fi82 < 18; fi82++ {
                var v83 uint16 = 893
                fmt.Println(v39, fi82)
                fmt.Println(v83, v39)
                _ = fi82
                _ = v83
            }
            v84 := []int{(v7 >> (((^(v2))) & 63))}
            _ = v84
        }
    }
    v85 := GenS7{F0: v3}
    if true {
        var v86 float64 = 0.0
        _ = v86
    } else {
        for fi87 := 0; fi87 < 17; fi87++ {
            v2 -= 0
            v7 += v7
            if ((3.14159 + ((1.5 - float64(((float32(1.5) / 0.1) / 1.0))) - 1e10)) <= 3.14159) {
                var v88 bool = true
                var v89 bool = (878 != 2)
                _ = v88
                _ = v89
            } else {
                v3 = "foo bar"
                for fi90 := 0; fi90 < 16; fi90++ {
                    v3 = v39
                    var v91 int = fi87
                    var v92 string = (("" + "foo bar") + v39)
                    _ = fi90
                    _ = v91
                    _ = v92
                }
            }
            _ = fi87
        }
        v2 += (v2 / (((v2 * 749)) | 1))
    }
    v39 = ("a" + v39)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v7
    _ = v85
}
