package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 uint }
type GenS2 struct { F0 int64; F1 uint; F2 int64 }
type GenS3 struct { F0 uint8; F1 int64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint8; F1 string; F2 uint8 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 string; F1 int32; F2 int8 }
type GenS10 struct { F0 uint; F1 uint32 }
type GenS11 struct { F0 int16; F1 int32 }
func genFn0(p0 string, p1 uint32, p2 uint32) uint {
    v6 := GenS1{F0: 42, F1: 706, F2: 943}
    switch -2147483648 {
    case 378:
        var v7 uint32 = p2
        _ = v7
    case -785:
        fmt.Println(p2, p0, p2)
    case -431:
        var v8 int8 = 84
        _ = v8
    }
    v9 := []int{127}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v9
    return 1
}
func genFn1(p0 uint) int32 {
    if (1 == -630) {
        var v10 string = "a"
        _ = v10
    } else {
        v11 := GenS1{F0: 1, F1: 9223372036854775807, F2: 637}
        _ = v11
    }
    v12 := GenS1{F0: 782, F1: 2147483647, F2: genFn0("\n\t", uint32(p0), 0)}
    fmt.Println(p0)
    for fi13 := 0; fi13 < 20; fi13++ {
        var v14 uint16 = 42
        var v15 float32 = ((1.5 + (3.14159 * 3.14159)) - -2.25)
        _ = fi13
        _ = v14
        _ = v15
    }
    p0 = 255
    _ = p0
    _ = v12
    return 0
}

func main() {
    v1 := make(map[int]string)
    var v2 int8 = 30
    var v3 int16 = 32767
    var v4 string = "a"
    v5 := make(map[uint32]bool)
    var v16 uint = 255
    for fi17 := 0; fi17 < 18; fi17++ {
        for fi18 := 0; fi18 < 16; fi18++ {
            if (false && (v4 < ("hello" + ("a" + "\n\t")))) {
                if false {
                    fmt.Println(v3, v3, v4)
                    var v19 uint32 = 42
                    var v20 bool = true
                    _ = v19
                    _ = v20
                } else {
                    switch (fi17 - fi17) {
                    case 42:
                        v3 -= (v3 % ((v3) | 1))
                    case 127:
                        var v21 string = (v4 + (v4 + v4))
                        _ = v21
                    default:
                        v22 := []int{(fi17 ^ (fi17 - (^(fi18))))}
                        _ = v22
                    }
                    v1[fi18] = v4
                    v3 = v3
                    var v23 int16 = v3
                    v23 += -106
                    _ = v23
                }
                v24 := GenS1{F0: 560, F1: (fi18 >> ((255) & 63)), F2: genFn0(("hello" + "foo bar"), 967, 42)}
                v25 := []int{}
                _ = v24
                _ = v25
            } else {
                v3 = (v3 * (v3 + (-(v3))))
                if (false && (genFn1(v16) > 2147483647)) {
                    v3 += v3
                    _ = len(v1)
                } else {
                    v2 = (v2 * (v2 - v2))
                }
                var v26 float32 = -2.25
                v27 := GenS1{F0: 2, F1: (fi18 << ((65535) & 63)), F2: genFn0(("" + (v4 + v4)), 108, 65535)}
                _ = len(v5)
                v28 := GenS2{F0: 9223372036854775807, F1: (v16 & 18446744073709551615), F2: -1}
                v29 := GenS1{F0: 663, F1: fi18, F2: (v16 + (^(v16)))}
                _ = v26
                _ = v27
                _ = v28
                _ = v29
            }
            if ((v2 / (((v2 >> ((genFn0("hello", 1, 2)) & 63))) | 1)) < (v2 / (((v2 % (((v2 | -1)) | 1))) | 1))) {
                v30 := []int{(-(fi18)), (fi18 / ((9223372036854775807) | 1)), (fi17 - -901), -192}
                _ = v30
            } else {
                v31 := GenS2{F0: 9223372036854775807, F1: v16, F2: -870}
                var v32 int = (-(fi17))
                if true {
                    v4 = v4
                } else {
                    var v33 int = (fi17 | 32768)
                    v34 := GenS3{F0: 255, F1: -118}
                    _ = v33
                    _ = v34
                }
                _ = v31
                _ = v32
            }
            switch fi17 {
            case 127:
                fmt.Println(v4)
            case -336:
                switch (fi18 - 535) {
                case -2147483648:
                    switch "foo bar" {
                    case "a":
                        v35 := GenS2{F0: 1, F1: v16, F2: -9223372036854775808}
                        _ = v35
                    }
                    switch "" {
                    case "":
                        v36 := v1[fi18]
                        _ = v36
                    case "a":
                        var v37 int = fi17
                        _ = v37
                    case "\n\t":
                        v4 = v4
                    }
                case 2147483647:
                    var v38 uint8 = 134
                    _ = v38
                case 32768:
                    var v39 int32 = genFn1(genFn0(v4, 1, uint32(v16)))
                    _ = v39
                default:
                    v40 := GenS4{F0: true}
                    _ = v40
                }
                var v41 int = (fi18 - fi17)
                var v42 int16 = -32768
                _ = v41
                _ = v42
            case 32767:
                var v43 int8 = (v2 + (v2 << (((v16 >> ((genFn0((("a" + v4) + v4), 42, 2147483647)) & 63))) & 63)))
                _ = v43
            }
            _ = fi18
        }
        if false {
            if ((!false) || (v16 <= (v16 & 4294967295))) {
                switch ((v4 + (((v4 + ("" + (v4 + v4))) + "a") + ("a" + v4))) + ("" + ((((v4 + ("hello" + "foo bar")) + "foo bar") + ("a" + "foo bar")) + v4))) {
                case "foo bar":
                    if (1e10 > 0.1) {
                        var v44 int = (fi17 + 573)
                        var v45 int32 = 0
                        v3 += (v3 + 1000)
                        _ = v44
                        _ = v45
                    } else {
                        var v46 int32 = 2147483647
                        var v47 float32 = 1e10
                        _ = len(v1)
                        _ = v46
                        _ = v47
                    }
                }
                var v48 float64 = 1e10
                v49 := []int{fi17, fi17, fi17}
                _ = v48
                _ = v49
            } else {
                var v50 uint = (v16 % ((1) | 1))
                var v51 int32 = 2147483647
                if false {
                    v52 := GenS4{F0: (1e10 == 3.14159)}
                    _ = v52
                }
                _ = v50
                _ = v51
            }
        }
        _ = fi17
    }
    v3 = v3
    v2 -= 42
    v53 := GenS5{F0: 2}
    fmt.Println(v3)
    v54 := GenS6{F0: 189, F1: "hello", F2: 255}
    if (-9223372036854775808 != 0) {
        switch int(v16) {
        case 128:
            var v55 int8 = (v2 * (v2 ^ (v2 >> (((v16 * 94)) & 63))))
            _ = v55
        default:
            v56 := []int{2147483647, 32767, -1, 0}
            _ = v56
        }
    }
    for fi57 := 0; fi57 < 19; fi57++ {
        if (true || true) {
            if true {
                for fi58 := 0; fi58 < 13; fi58++ {
                    if (false && false) {
                        var v59 uint16 = 470
                        var v60 float64 = 1e10
                        var v61 float32 = 0.0
                        _ = v59
                        _ = v60
                        _ = v61
                    }
                    var v62 uint8 = 128
                    v63 := GenS7{F0: v4}
                    var v64 int16 = v3
                    var v65 uint32 = 4294967295
                    if false {
                        var v66 float64 = float64(0.0)
                        var v67 string = (v4 + "hello")
                        var v68 uint = (v16 * v16)
                        _ = v66
                        _ = v67
                        _ = v68
                    } else {
                        if (65535 < (v65 >> ((607) & 63))) {
                            var v69 int32 = genFn1(0)
                            _ = v69
                        } else {
                            if ((-2.25 + (0.1 / float32(0.1))) >= (1e10 - 1.5)) {
                                v2 += (v2 % (((v2 ^ 16)) | 1))
                            }
                            var v70 uint64 = 243
                            switch 32768 {
                            case 128:
                                var v71 string = v4
                                _ = v71
                            case -831:
                                v72 := []int{fi57, 128, fi57, 150}
                                _ = v72
                            case 4294967295:
                                v73 := GenS5{F0: v65}
                                _ = v73
                            default:
                                if (127 != (v62 - (v62 ^ v62))) {
                                    v74 := []int{(fi58 / (((fi58 + fi58)) | 1)), (fi57 * (fi58 << ((uint(v70)) & 63))), fi58, (fi57 + fi57)}
                                    _ = v74
                                } else {
                                    v70 -= (v70 << ((0) & 63))
                                }
                            }
                            _ = v70
                        }
                        switch fi58 {
                        case 32767:
                            delete(v5, v65)
                        }
                        var v75 uint8 = 255
                        _ = v75
                    }
                    fmt.Println(v2, v2, v2)
                    var v76 bool = (!true)
                    _ = fi58
                    _ = v62
                    _ = v63
                    _ = v64
                    _ = v65
                    _ = v76
                }
                var v77 uint8 = 1
                v78 := GenS8{F0: -1}
                _ = v77
                _ = v78
            }
            if (1 != 0) {
                v79 := GenS4{F0: (!false)}
                v3 = (v3 | (v3 & v3))
                v2 = -70
                _ = v79
            }
        } else {
            for fi80 := 0; fi80 < 16; fi80++ {
                if false {
                    v4 = "foo bar"
                    var v81 uint8 = 253
                    v2 += v2
                    switch (fi80 >> (((^(v16))) & 63)) {
                    case 1:
                        v82 := GenS6{F0: v81, F1: (v4 + v4), F2: v81}
                        _ = v82
                    case 32768:
                        if (!false) {
                            fmt.Println(v4, v2)
                            if (v2 >= 86) {
                                v3 = 42
                                v2 = 86
                            } else {
                                var v83 float32 = 0.1
                                _ = v83
                            }
                            fmt.Println(fi80, v81)
                        }
                    default:
                        v4 = v4
                    }
                    fmt.Println(v2, v3)
                    _ = v81
                } else {
                    switch v4 {
                    case "hello":
                        var v84 uint64 = 0
                        _ = v84
                    case "a":
                        v85 := GenS9{F0: (v4 + v4), F1: 2147483647, F2: v2}
                        _ = v85
                    }
                    switch (fi57 >> (((v16 & (^(v16)))) & 63)) {
                    case 128:
                        var v86 int16 = 32767
                        _ = v86
                    case -997:
                        var v87 uint8 = 191
                        _ = v87
                    }
                    switch ((v4 + ("hello" + v4)) + "foo bar") {
                    case "\n\t":
                        var v88 int8 = 127
                        v89 := GenS10{F0: 2, F1: 2147483647}
                        v90 := GenS11{F0: v3, F1: -52}
                        _ = v88
                        _ = v89
                        _ = v90
                    case "foo bar":
                        var v91 int8 = v2
                        _ = v91
                    }
                }
                _ = fi80
            }
        }
        fmt.Println(v3)
        _ = fi57
    }
    if (!(9223372036854775807 <= 1)) {
        v92 := GenS9{F0: "", F1: 2147483647, F2: (v2 ^ (v2 & v2))}
        for fi93 := 0; fi93 < 15; fi93++ {
            var v94 uint8 = 2
            v16 += 144
            _ = fi93
            _ = v94
        }
        _ = v92
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v54
}
