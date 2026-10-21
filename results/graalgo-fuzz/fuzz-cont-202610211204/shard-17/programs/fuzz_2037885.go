package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int; F1 uint8; F2 int64 }
type GenS3 struct { F0 string; F1 int64 }
type GenS4 struct { F0 int16; F1 uint16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 string; F1 uint }
type GenS7 struct { F0 int32; F1 int32; F2 int16 }
func genFn0(p0 uint8, p1 int64, p2 int32) uint16 {
    v7 := GenS1{F0: (^(p0))}
    var v8 uint = 42
    for fi9 := 0; fi9 < 16; fi9++ {
        v10 := GenS1{F0: (p0 | p0)}
        if (!(p1 >= (p1 % (((p1 - 1)) | 1)))) {
            for fi11 := 0; fi11 < 11; fi11++ {
                v12 := GenS2{F0: (fi11 - (^(fi11))), F1: 2, F2: p1}
                var v13 int8 = -1
                v14 := GenS3{F0: "a", F1: (-(p1))}
                fmt.Println(fi11, fi11, fi11)
                _ = fi11
                _ = v12
                _ = v13
                _ = v14
            }
        }
        var v15 float64 = 1.5
        _ = fi9
        _ = v10
        _ = v15
    }
    v8 -= v8
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 436
}

func main() {
    v1 := make(map[int]int)
    var v2 float64 = float64((0.1 - float32((float64((float32(float64(float32((0.1 * (float64((1e10 - ((0.0 - -2.25) / 1.0))) - 0.1))))) + 0.1)) * 1.5))))
    v3 := []int{}
    v4 := []string{}
    v5 := []uint32{751, 2147483647, 4294967295, 1}
    var v6 bool = false
    if v6 {
        var v16 bool = (42 >= 2147483648)
        if v6 {
            v2 = v2
        } else {
            for fi17 := 0; fi17 < 4; fi17++ {
                var v18 uint32 = 0
                for fi19 := 0; fi19 < 3; fi19++ {
                    switch "" {
                    case "":
                        var v20 bool = true
                        _ = v20
                    case "a":
                        if (v16 || v16) {
                            v18 += (v18 + 153)
                            v21 := GenS4{F0: 0, F1: 1}
                            var v22 int8 = 104
                            _ = v21
                            _ = v22
                        } else {
                            var v23 uint8 = 255
                            if ((42 != 0) || v16) {
                                var v24 uint64 = 42
                                v6 = (0.0 != ((-2.25 - 3.14159) + 1e10))
                                v25 := GenS2{F0: fi17, F1: v23, F2: 1}
                                _ = v24
                                _ = v25
                            } else {
                                if v6 {
                                    var v26 uint8 = v23
                                    _ = v26
                                } else {
                                    fmt.Println(v2, v18)
                                    v16 = v16
                                }
                                delete(v1, fi19)
                            }
                            _ = v23
                        }
                    case "\n\t":
                        v16 = false
                    }
                    if v16 {
                        v6 = (v16 && (-2.25 <= 0.1))
                        v27 := GenS3{F0: "foo bar", F1: -1}
                        _ = v27
                    } else {
                        var v28 int8 = 42
                        _ = v28
                    }
                    _ = fi19
                }
                switch fi17 {
                case 0:
                    for fi29 := 0; fi29 < 20; fi29++ {
                        switch -9223372036854775808 {
                        case 820:
                            v6 = (v16 && ((158 < 200) || (!(v2 < ((((((1e10 * v2) * (v2 + 1.5)) / 1.0) + (0.1 * (v2 / (-2.25 + (v2 / v2))))) - 0.1) - v2)))))
                            fmt.Println(fi29, fi17)
                        case 1:
                            if true {
                                if false {
                                    v6 = v6
                                    v18 -= (v18 * 4294967295)
                                    var v30 int32 = 597
                                    _ = v30
                                } else {
                                    fmt.Println(v18, fi17)
                                    if (127 >= 1) {
                                        var v31 uint8 = uint8(v18)
                                        v32 := GenS5{F0: (("hello" + "\n\t") + "")}
                                        _ = v31
                                        _ = v32
                                    }
                                }
                                var v33 uint8 = 128
                                _ = v33
                            } else {
                                v34 := v1[(fi29 | (-(fi17)))]
                                v34 = fi29
                                if (!((v6 && (1 == 1)) || v16)) {
                                    if (65535 < 0) {
                                        if false {
                                            var v35 uint8 = 128
                                            fmt.Println(v2, v6, fi17)
                                            _ = v35
                                        } else {
                                            fmt.Println(v18, v34)
                                        }
                                        fmt.Println(fi29, v2, v6)
                                    } else {
                                        if true {
                                            switch 32767 {
                                            case 32768:
                                                v18 = (v18 ^ 971)
                                            default:
                                                var v36 int8 = -117
                                                _ = v36
                                            }
                                            v37 := GenS6{F0: "a", F1: 42}
                                            var v38 int = fi29
                                            _ = v37
                                            _ = v38
                                        } else {
                                            var v39 int8 = 100
                                            _ = v39
                                        }
                                    }
                                }
                                _ = len(v4)
                                _ = v34
                            }
                        case 9223372036854775807:
                            fmt.Println(fi17, v16, v6)
                        default:
                            v16 = v6
                        }
                        _ = fi29
                    }
                case -1:
                    var v40 bool = v16
                    _ = v40
                case -144:
                    for fi41 := 0; fi41 < 20; fi41++ {
                        v2 = v2
                        v42 := GenS2{F0: (fi17 - 0), F1: 100, F2: 395}
                        switch ("a" + "a") {
                        case "a":
                            delete(v1, fi41)
                        }
                        _ = fi41
                        _ = v42
                    }
                default:
                    for fi43 := 0; fi43 < 15; fi43++ {
                        v5 = append(v5, v18)
                        _ = fi43
                    }
                }
                _ = fi17
                _ = v18
            }
            v6 = (false && false)
            var v44 float64 = 0.0
            _ = v44
        }
        fmt.Println(v16, v16, v6)
        _ = v16
    } else {
        v6 = v6
    }
    if (1 < 9223372036854775807) {
        delete(v1, -2147483648)
    } else {
        v45 := GenS1{F0: 144}
        v6 = (!("\n\t" > "\n\t"))
        var v46 bool = v6
        _ = v45
        _ = v46
    }
    var v47 uint16 = 687
    for fi48 := 0; fi48 < 15; fi48++ {
        var v49 string = (("foo bar" + "\n\t") + ("hello" + ""))
        if v6 {
            if (65535 == 65535) {
                var v50 uint16 = 65535
                var v51 uint64 = 42
                _ = v50
                _ = v51
            }
        } else {
            if (true && ((v47 / ((v47) | 1)) < (^(v47)))) {
                v52 := GenS4{F0: 1000, F1: genFn0(1, 632, -1)}
                if (!(!true)) {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    var v53 int8 = 1
                    fmt.Println(v2, v49, v49)
                    _ = v53
                } else {
                    switch fi48 {
                    case 9223372036854775807:
                        for fi54 := 0; fi54 < 4; fi54++ {
                            v4 = append(v4, v49)
                            v55 := GenS4{F0: -1, F1: (v47 % ((1) | 1))}
                            _ = fi54
                            _ = v55
                        }
                    case -2147483648:
                        var v56 float64 = v2
                        _ = v56
                    case 0:
                        v57 := GenS2{F0: fi48, F1: 255, F2: -9223372036854775808}
                        _ = v57
                    default:
                        v58 := GenS1{F0: 171}
                        _ = v58
                    }
                    v59 := GenS1{F0: 73}
                    if (!v6) {
                        v49 = ((v49 + v49) + "a")
                    } else {
                        switch v49 {
                        case "\n\t":
                            v60 := GenS7{F0: -868, F1: 1, F2: 0}
                            _ = v60
                        }
                    }
                    _ = v59
                }
                _ = v52
            }
        }
        var v61 int16 = 42
        _ = fi48
        _ = v49
        _ = v61
    }
    v5 = append(v5, 0)
    if true {
        switch ("hello" + "foo bar") {
        case "hello":
            for fi62 := 0; fi62 < 9; fi62++ {
                var v63 uint64 = 2
                fmt.Println(v6)
                if (!v6) {
                    v64 := GenS7{F0: -48, F1: 0, F2: 42}
                    v47 = uint16(v63)
                    v47 = uint16(v63)
                    _ = v64
                }
                _ = fi62
                _ = v63
            }
        default:
            for fi65 := 0; fi65 < 16; fi65++ {
                var v66 uint32 = 921
                for fi67 := 0; fi67 < 8; fi67++ {
                    var v68 float64 = float64(0.1)
                    _ = fi67
                    _ = v68
                }
                var v69 bool = true
                _ = fi65
                _ = v66
                _ = v69
            }
        }
    }
    fmt.Println(v2, v2)
    var v70 string = "\n\t"
    v47 = (v47 + (v47 & (v47 - v47)))
    fmt.Println(v2, v6)
    v71 := v1[128]
    var v72 int16 = 42
    if v6 {
        for fi73 := 0; fi73 < 10; fi73++ {
            v72 = v72
            _ = fi73
        }
    }
    v72 -= 0
    v71 = v71
    var v74 int = (^(v71))
    v3 = append(v3, (v74 - int(v47)))
    var v75 uint64 = 18446744073709551615
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    fmt.Println("v72:", v72)
    fmt.Println("v74:", v74)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v6
    _ = v70
    _ = v71
    _ = v72
    _ = v74
    _ = v75
}
