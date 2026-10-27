package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 uint64; F1 int16; F2 uint }
type GenS3 struct { F0 int32; F1 uint }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 string; F1 int8 }
func genFn0() uint64 {
    var v4 int16 = 32767
    v5 := GenS1{F0: 9223372036854775807, F1: false}
    v6 := GenS1{F0: 9223372036854775807, F1: (0 == 1)}
    _ = v4
    _ = v5
    _ = v6
    return 4294967295
}
func genFn1(p0 uint64, p1 uint32, p2 bool) int {
    var v7 bool = (uint64(p1) == p0)
    for fi8 := 0; fi8 < 3; fi8++ {
        var v9 int16 = 79
        v10 := []int{}
        _ = fi8
        _ = v10
        _ = v9
    }
    switch 32767 {
    case 42:
        for fi11 := 0; fi11 < 7; fi11++ {
            var v12 uint32 = p1
            v13 := []int{fi11, 32768, fi11}
            _ = fi11
            _ = v12
            _ = v13
        }
    case -746:
        fmt.Println(p0, p2, p2)
    case -380:
        var v14 int8 = 0
        _ = v14
    }
    if (32767 > 1000) {
        var v15 uint16 = 255
        if v7 {
            var v16 uint64 = 435
            if p2 {
                v17 := []int{576, -2147483648}
                for fi18 := 0; fi18 < 16; fi18++ {
                    var v19 int32 = -961
                    _ = fi18
                    _ = v19
                }
                var v20 int8 = -25
                _ = v17
                _ = v20
            }
            _ = v16
        }
        switch ("" + "") {
        case "hello":
            p1 = (p1 ^ 2147483648)
        case "":
            v21 := GenS1{F0: 9223372036854775807, F1: (true && v7)}
            _ = v21
        case "\n\t":
            if v7 {
                for fi22 := 0; fi22 < 3; fi22++ {
                    if false {
                        var v23 int16 = -1
                        _ = v23
                    }
                    v24 := []int{(fi22 + 1), (fi22 & fi22), 127, fi22}
                    _ = fi22
                    _ = v24
                }
                v25 := GenS2{F0: (p0 >> ((942) & 63)), F1: 1000, F2: 999}
                _ = v25
            } else {
                v26 := GenS2{F0: p0, F1: -32768, F2: 2}
                v15 = (v15 ^ 2)
                _ = v26
            }
        default:
            fmt.Println(p0, v7, v15)
        }
        _ = v15
    } else {
        for fi27 := 0; fi27 < 17; fi27++ {
            p1 = p1
            v28 := GenS1{F0: 42, F1: false}
            _ = fi27
            _ = v28
        }
    }
    var v29 bool = (!((p0 % ((2) | 1)) == p0))
    _ = p0
    _ = p1
    _ = p2
    _ = v29
    _ = v7
    return 9223372036854775807
}

func main() {
    var v1 int16 = 0
    v2 := make(map[int]bool)
    v3 := make(map[uint32]uint32)
    var v30 bool = (0.1 != -2.25)
    v30 = v30
    if v30 {
        v30 = v30
        v31 := GenS2{F0: 709, F1: (v1 * v1), F2: 4294967295}
        fmt.Println(v30)
        _ = v31
    } else {
        var v32 int = -249
        v30 = (!(-601 < 2147483647))
        v33 := GenS2{F0: 18446744073709551615, F1: v1, F2: 4294967295}
        _ = v32
        _ = v33
    }
    v34 := []int{32768, 1}
    v35 := GenS2{F0: 912, F1: v1, F2: 722}
    switch -9223372036854775808 {
    case 2147483647:
        v1 += (^(v1))
    case -574:
        var v36 int = 32768
        _ = v36
    case -184:
        v30 = false
    }
    v37 := GenS1{F0: -1, F1: (303 <= 1)}
    switch ("" + "foo bar") {
    case "a":
        var v38 string = "a"
        _ = v38
    }
    for fi39 := 0; fi39 < 10; fi39++ {
        var v40 float32 = ((3.14159 + (1.5 * 1.5)) - 3.14159)
        if ((903 == 1) || v30) {
            for fi41 := 0; fi41 < 15; fi41++ {
                switch fi39 {
                case -1:
                    var v42 float64 = (((0.1 / (1.5 + (1e10 / 1e10))) + ((1.5 * (((-2.25 / (float64(1e10) / float64(float32(-2.25)))) * (float64(v40) - (1.5 - -2.25))) + 0.1)) - (-2.25 / float64(-2.25)))) * 3.14159)
                    _ = v42
                default:
                    if (!v30) {
                        switch genFn1(2, 985, v30) {
                        case 32768:
                            fmt.Println(v1, v40, fi39)
                        default:
                            v43 := GenS1{F0: 42, F1: (2147483647 > 162)}
                            _ = v43
                        }
                        if len(v34) > 0 {
                            v34[0] = (fi39 >> ((4294967295) & 63))
                        }
                        fmt.Println(v40, fi39, fi39)
                    } else {
                        fmt.Println(fi39)
                        var v44 uint16 = 773
                        _ = v44
                    }
                }
                _ = fi41
            }
            var v45 int = genFn1(42, 2, v30)
            v1 -= v1
            _ = v45
        }
        _ = fi39
        _ = v40
    }
    v46 := GenS2{F0: 42, F1: (^(v1)), F2: 2}
    var v47 float64 = 0.1
    switch "foo bar" {
    case "foo bar":
        var v48 bool = (!(2 == 0))
        _ = v48
    case "":
        v1 = 42
    case "\n\t":
        v49 := GenS3{F0: 2147483647, F1: 636}
        _ = v49
    default:
        for fi50 := 0; fi50 < 12; fi50++ {
            for fi51 := 0; fi51 < 7; fi51++ {
                v1 = v1
                v1 = v1
                fmt.Println(fi51, v30)
                _ = fi51
            }
            switch (("\n\t" + "foo bar") + "a") {
            case "foo bar":
                v30 = false
            case "\n\t":
                var v52 int32 = 2147483647
                _ = v52
            default:
                switch fi50 {
                case -508:
                    v1 += (v1 / ((797) | 1))
                case 555:
                    for fi53 := 0; fi53 < 17; fi53++ {
                        if v30 {
                            if true {
                                v54 := GenS4{F0: v1}
                                if false {
                                    if (0 != 771) {
                                        var v55 uint16 = 42
                                        v56 := GenS5{F0: ("foo bar" + ("\n\t" + "foo bar")), F1: -128}
                                        fmt.Println(fi50, v30, v47)
                                        _ = v55
                                        _ = v56
                                    }
                                    v57 := GenS1{F0: 9223372036854775807, F1: v30}
                                    if ((v1 / (((v1 | -1)) | 1)) > 1000) {
                                        v30 = v30
                                        v3[2147483648] = 755
                                    } else {
                                        v58 := v2[(fi50 ^ (fi50 / (((fi50 | (fi53 >> ((0) & 63)))) | 1)))]
                                        if len(v34) > 0 {
                                            v34[0] = 2147483647
                                        }
                                        v59 := GenS4{F0: (v1 % ((86) | 1))}
                                        _ = v58
                                        _ = v59
                                    }
                                    _ = v57
                                } else {
                                    v47 += 1.5
                                    switch ("" + "a") {
                                    case "hello":
                                        _ = len(v34)
                                    case "foo bar":
                                        var v60 string = ((("foo bar" + ("\n\t" + "")) + "foo bar") + ("hello" + "a"))
                                        _ = v60
                                    case "\n\t":
                                        switch (fi50 - genFn1(293, 2, (!(((fi53 != (fi53 % ((genFn1(0, 0, v30)) | 1))) && v30) && v30)))) {
                                        case -1:
                                            v30 = v30
                                        case 987:
                                            v1 -= 0
                                        default:
                                            v30 = (!(127 > -1))
                                        }
                                    }
                                }
                                v3[409] = 874
                                _ = v54
                            }
                            v3[42] = 4294967295
                        }
                        var v61 float64 = 1e10
                        var v62 uint = 2
                        _ = fi53
                        _ = v61
                        _ = v62
                    }
                case -379:
                    var v63 bool = (-1 < 42)
                    _ = v63
                default:
                    v47 += (0.0 / (v47 - v47))
                }
            }
            _ = fi50
        }
    }
    switch 4294967295 {
    case 2147483647:
        switch genFn1(4294967295, 315, v30) {
        case -2147483648:
            v64 := GenS1{F0: -449, F1: ((v30 || (2147483648 >= 944)) && true)}
            _ = v64
        default:
            v30 = v30
        }
    case 733:
        if v30 {
            for fi65 := 0; fi65 < 12; fi65++ {
                v30 = v30
                var v66 uint8 = 2
                if len(v34) > 0 {
                    v34[0] = fi65
                }
                _ = fi65
                _ = v66
            }
            _ = len(v34)
        }
    }
    var v67 uint8 = 127
    for fi68 := 0; fi68 < 12; fi68++ {
        for fi69 := 0; fi69 < 3; fi69++ {
            if v30 {
                v47 = (0.0 * v47)
                var v70 string = "hello"
                _ = v70
            } else {
                v47 = -2.25
                v1 = -32768
            }
            _ = fi69
        }
        switch (^(fi68)) {
        case -1:
            if true {
                v47 = ((1e10 - (v47 * v47)) - (float64(3.14159) + (v47 - float64(0.0))))
                v71 := GenS4{F0: (v1 + (v1 & int16(fi68)))}
                _ = v71
            } else {
                var v72 bool = (!false)
                v73 := GenS5{F0: "\n\t", F1: 127}
                _ = v72
                _ = v73
            }
        }
        _ = fi68
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v47:", v47)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v34
    _ = v35
    _ = v37
    _ = v46
    _ = v47
    _ = v67
}
