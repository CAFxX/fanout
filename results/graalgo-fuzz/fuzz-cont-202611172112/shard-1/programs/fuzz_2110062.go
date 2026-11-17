package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 int64 }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 uint8; F1 int64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint64 }
func genFn0(p0 uint64, p1 uint8) int {
    p1 = 127
    for fi4 := 0; fi4 < 3; fi4++ {
        var v5 uint = 5
        var v6 uint32 = 0
        v7 := GenS1{F0: -9223372036854775808, F1: 42, F2: -9223372036854775808}
        _ = fi4
        _ = v5
        _ = v6
        _ = v7
    }
    var v8 uint64 = p0
    switch "" {
    case "":
        var v9 int = 128
        _ = v9
    default:
        fmt.Println(p1, v8)
    }
    v10 := []int{-9223372036854775808}
    _ = p0
    _ = p1
    _ = v10
    _ = v8
    return 796
}
func genFn1(p0 int32, p1 uint16) int64 {
    var v11 uint = 1
    if true {
        switch ("a" + "\n\t") {
        case "foo bar":
            var v12 uint16 = p1
            _ = v12
        }
        for fi13 := 0; fi13 < 1; fi13++ {
            var v14 int16 = 1
            var v15 int64 = 0
            _ = fi13
            _ = v14
            _ = v15
        }
        var v16 bool = (386 != p1)
        _ = v16
    }
    _ = p0
    _ = p1
    _ = v11
    return 54
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int8 = -1
    v3 := make(map[uint32]int)
    var v17 uint64 = 18446744073709551615
    if (-643 >= 42) {
        v2 = (v2 % (((-(v2))) | 1))
        v18 := v1[497]
        delete(v1, v18)
        _ = v18
    }
    v19 := GenS1{F0: -674, F1: 2147483647, F2: 42}
    var v20 string = "hello"
    for fi21 := 0; fi21 < 11; fi21++ {
        if (-269 == 0) {
            if true {
                if true {
                    for fi22 := 0; fi22 < 6; fi22++ {
                        v17 -= (v17 >> ((18446744073709551615) & 63))
                        v2 -= (v2 | -45)
                        switch (fi22 >> ((1) & 63)) {
                        case 128:
                            switch v20 {
                            case "":
                                var v23 uint8 = 42
                                _ = v23
                            case "foo bar":
                                v17 = 4294967295
                            }
                        case 32768:
                            if (-2.25 >= (1e10 + 0.0)) {
                                if true {
                                    v2 = v2
                                } else {
                                    v24 := GenS1{F0: -312, F1: -2147483648, F2: -1}
                                    _ = v24
                                }
                                v2 -= 127
                                switch v20 {
                                case "hello":
                                    v20 = "hello"
                                default:
                                    if false {
                                        var v25 int32 = 1
                                        _ = v25
                                    }
                                }
                                if false {
                                    var v26 float32 = (float32((0.1 + 1e10)) / 3.14159)
                                    var v27 float32 = v26
                                    _ = v26
                                    _ = v27
                                }
                            }
                        case -558:
                            if (!(((-2.25 + (-2.25 - (float32(0.1) * 3.14159))) < ((1e10 - ((-2.25 + 3.14159) + 1e10)) - 1e10)) || (v17 >= v17))) {
                                v20 = "a"
                                v28 := GenS1{F0: 1, F1: 42, F2: 42}
                                switch ("foo bar" + ("a" + "\n\t")) {
                                case "":
                                    v2 += (v2 & 110)
                                case "\n\t":
                                    v2 = v2
                                case "foo bar":
                                    v29 := GenS1{F0: genFn1(-592, 399), F1: 0, F2: -9223372036854775808}
                                    _ = v29
                                }
                                fmt.Println(v2)
                                _ = v28
                            }
                        default:
                            if true {
                                var v30 int8 = v2
                                var v31 uint64 = 1
                                var v32 int64 = 1
                                _ = v30
                                _ = v31
                                _ = v32
                            } else {
                                v20 = ((("foo bar" + "a") + (v20 + "")) + "foo bar")
                            }
                        }
                        _ = fi22
                    }
                    v20 = "\n\t"
                } else {
                    for fi33 := 0; fi33 < 18; fi33++ {
                        v34 := GenS1{F0: 0, F1: -23, F2: genFn1(2147483647, 32767)}
                        var v35 uint = 596
                        _ = fi33
                        _ = v34
                        _ = v35
                    }
                    v36 := GenS1{F0: genFn1(0, 65535), F1: 506, F2: genFn1(-824, 0)}
                    if false {
                        var v37 bool = (false || true)
                        var v38 bool = (!true)
                        _ = v37
                        _ = v38
                    }
                    _ = v36
                }
            } else {
                var v39 float32 = 1e10
                _ = v39
            }
            for fi40 := 0; fi40 < 9; fi40++ {
                var v41 uint8 = 2
                if false {
                    if (!(true && true)) {
                        if true {
                            v17 += (v17 + (^(v17)))
                            var v42 bool = ((v41 >> ((623) & 63)) == v41)
                            v42 = (!(42 > -32768))
                            _ = v42
                        }
                        v2 = (v2 - 123)
                        var v43 int16 = 32767
                        _ = v43
                    }
                    v17 += uint64(v41)
                    v44 := []int{2147483647}
                    _ = v44
                }
                var v45 int = fi40
                var v46 int16 = 42
                _ = fi40
                _ = v41
                _ = v45
                _ = v46
            }
        } else {
            fmt.Println(fi21)
            switch 9223372036854775807 {
            case 42:
                var v47 float32 = float32((3.14159 + 1e10))
                _ = v47
            case 528:
                for fi48 := 0; fi48 < 11; fi48++ {
                    switch fi21 {
                    case -416:
                        switch 127 {
                        case 128:
                            v49 := GenS1{F0: 689, F1: 42, F2: -1}
                            _ = v49
                        case 320:
                            var v50 uint8 = 42
                            _ = v50
                        }
                    case -143:
                        var v51 float64 = 1e10
                        _ = v51
                    }
                    switch v20 {
                    case "\n\t":
                        var v52 int8 = (v2 | (v2 >> ((18446744073709551615) & 63)))
                        _ = v52
                    case "foo bar":
                        var v53 bool = (-9223372036854775808 < 9223372036854775807)
                        _ = v53
                    }
                    if (!true) {
                        v54 := GenS2{F0: v17, F1: -29}
                        v20 = v20
                        if false {
                            fmt.Println(fi48)
                            if false {
                                if (fi21 > 1) {
                                    var v55 uint = 491
                                    v17 = v17
                                    v20 = (("foo bar" + v20) + v20)
                                    fmt.Println(v20, v20, v55)
                                    fmt.Println(fi48)
                                    _ = v55
                                }
                                v56 := GenS1{F0: 9223372036854775807, F1: -592, F2: 42}
                                v57 := GenS3{F0: 42, F1: genFn1(-1, 255)}
                                var v58 uint16 = 0
                                v3[2147483647] = (fi48 * 1)
                                _ = v56
                                _ = v57
                                _ = v58
                            } else {
                                var v59 float32 = (((0.0 - float32((((0.0 * 1e10) + 0.1) * float64((1e10 * -2.25))))) * -2.25) + 0.0)
                                _ = v59
                            }
                            fmt.Println(fi48, fi21)
                        } else {
                            v17 = v17
                        }
                        _ = v54
                    } else {
                        v60 := []int{fi21, genFn0(v17, 0)}
                        _ = v60
                    }
                    if true {
                        v61 := []int{fi48, fi21, (-(fi21)), genFn0((v17 - v17), 78)}
                        v62 := GenS4{F0: 610}
                        _ = v61
                        _ = v62
                    } else {
                        v63 := GenS2{F0: v17, F1: 1}
                        _ = v63
                    }
                    var v64 uint8 = 110
                    _ = fi48
                    _ = v64
                }
            default:
                switch (v20 + (v20 + (("hello" + v20) + v20))) {
                case "\n\t":
                    v20 = ("\n\t" + v20)
                }
            }
        }
        switch fi21 {
        case 42:
            v20 = ((v20 + ((v20 + ((v20 + "foo bar") + v20)) + ("a" + v20))) + ((v20 + "hello") + v20))
        case -2147483648:
            v17 = v17
        default:
            var v65 int8 = v2
            _ = v65
        }
        _ = fi21
    }
    var v66 int32 = -2147483648
    switch v20 {
    case "hello":
        v67 := []int{-275, 645}
        _ = v67
    case "":
        v17 = (v17 + (^(v17)))
    case "a":
        var v68 uint = 42
        _ = v68
    default:
        v66 = (v66 + (^(v66)))
    }
    fmt.Println(v66, v66)
    fmt.Println(v2, v2)
    if true {
        var v69 uint = 2
        var v70 bool = false
        _ = v69
        _ = v70
    }
    for fi71 := 0; fi71 < 2; fi71++ {
        var v72 uint = 4294967295
        for fi73 := 0; fi73 < 14; fi73++ {
            v74 := GenS4{F0: 2}
            if ((v17 * (v17 << (((v72 ^ (^(v72)))) & 63))) >= 954) {
                if false {
                    v75 := []int{fi71, (fi71 << (((v72 ^ 4294967295)) & 63))}
                    var v76 int32 = -2147483648
                    _ = len(v3)
                    _ = v75
                    _ = v76
                } else {
                    var v77 float64 = 0.0
                    _ = v77
                }
                if true {
                    var v78 bool = (!true)
                    _ = v78
                }
                v79 := []int{-1, fi71}
                _ = v79
            } else {
                v80 := GenS5{F0: v17}
                v72 = (v72 >> (((v72 % (((^(v72))) | 1))) & 63))
                _ = v80
            }
            v20 = (v20 + ("hello" + (("\n\t" + (v20 + v20)) + v20)))
            _ = fi73
            _ = v74
        }
        _ = fi71
        _ = v72
    }
    var v81 bool = true
    v3[948] = 47
    var v82 string = v20
    fmt.Println(v81)
    delete(v1, 762)
    var v83 bool = false
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v66:", v66)
    fmt.Println("v81:", v81)
    fmt.Println("v82:", v82)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v66
    _ = v81
    _ = v82
    _ = v83
}
