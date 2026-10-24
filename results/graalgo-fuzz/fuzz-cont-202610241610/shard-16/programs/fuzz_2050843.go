package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 int; F1 int16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint; F1 int64 }
type GenS5 struct { F0 int; F1 int8; F2 bool }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int64; F1 int64; F2 bool }
type GenS8 struct { F0 bool; F1 bool; F2 int8 }
func genFn0(p0 int) uint32 {
    fmt.Println(p0)
    var v4 int64 = 9223372036854775807
    p0 = p0
    for fi5 := 0; fi5 < 7; fi5++ {
        p0 += (fi5 & fi5)
        p0 = 2147483647
        _ = fi5
    }
    var v6 int = 127
    _ = p0
    _ = v4
    _ = v6
    return 361
}
func genFn1() uint64 {
    var v7 float32 = 1.5
    for fi8 := 0; fi8 < 2; fi8++ {
        v7 = v7
        if false {
            var v9 string = ("a" + ("" + "hello"))
            switch (^(fi8)) {
            case 896:
                v7 += 3.14159
            case -781:
                if (false && ((255 >= 2) && (true && false))) {
                    var v10 uint32 = 2
                    var v11 int32 = 849
                    _ = v10
                    _ = v11
                } else {
                    var v12 int16 = 42
                    for fi13 := 0; fi13 < 13; fi13++ {
                        var v14 bool = (true || (((284 < 18446744073709551615) && false) || (1 == 891)))
                        if (42 <= 42) {
                            v9 = "a"
                        }
                        _ = fi13
                        _ = v14
                    }
                    fmt.Println(v7, v12)
                    _ = v12
                }
            case -309:
                var v15 bool = (117 > 127)
                _ = v15
            default:
                var v16 string = v9
                _ = v16
            }
            _ = v9
        }
        _ = fi8
    }
    _ = v7
    return 18446744073709551615
}

func main() {
    v1 := make(map[int]bool)
    var v2 int = 2147483647
    var v3 uint16 = 1
    v17 := GenS1{F0: v3, F1: 42}
    v2 -= v2
    var v18 string = "a"
    fmt.Println(v18)
    if (true || false) {
        fmt.Println(v18, v3)
    } else {
        if false {
            v3 = 255
            if false {
                for fi19 := 0; fi19 < 15; fi19++ {
                    v18 = v18
                    var v20 float32 = 0.1
                    _ = fi19
                    _ = v20
                }
            }
        }
        var v21 float32 = (float32((0.0 * (-2.25 * -2.25))) / (1.5 * 1.5))
        _ = v21
    }
    for fi22 := 0; fi22 < 2; fi22++ {
        v23 := GenS1{F0: v3, F1: 2}
        v24 := v1[v2]
        _ = fi22
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 7; fi25++ {
        var v26 int8 = -51
        switch v2 {
        case 1:
            v27 := []int{v2, v2}
            _ = v27
        case -1:
            var v28 bool = true
            _ = v28
        default:
            var v29 uint16 = (v3 / ((42) | 1))
            _ = v29
        }
        _ = fi25
        _ = v26
    }
    for fi30 := 0; fi30 < 1; fi30++ {
        switch "\n\t" {
        case "a":
            var v31 float64 = 0.0
            _ = v31
        case "\n\t":
            switch 9223372036854775807 {
            case 0:
                if (v18 <= v18) {
                    var v32 int16 = -606
                    var v33 float32 = 1.5
                    v32 = (v32 + v32)
                    _ = v32
                    _ = v33
                }
            default:
                var v34 uint = 65535
                _ = v34
            }
        default:
            v35 := GenS2{F0: fi30, F1: 1}
            _ = v35
        }
        _ = fi30
    }
    if true {
        for fi36 := 0; fi36 < 5; fi36++ {
            if false {
                for fi37 := 0; fi37 < 17; fi37++ {
                    var v38 int = fi36
                    switch "a" {
                    case "\n\t":
                        fmt.Println(v18, fi36)
                    case "":
                        if (v18 <= v18) {
                            v39 := GenS1{F0: (v3 - (v3 % ((255) | 1))), F1: 148}
                            var v40 int = fi37
                            _ = v39
                            _ = v40
                        } else {
                            var v41 float32 = ((((1.5 * 1e10) / ((((0.1 + float32(-2.25)) - -2.25) - -2.25) + (1.5 - (-2.25 - 0.1)))) * (0.0 + -2.25)) + 1.5)
                            _ = v41
                        }
                    default:
                        switch fi37 {
                        case -669:
                            var v42 int32 = 2147483647
                            _ = v42
                        case 9223372036854775807:
                            if (132 == 18446744073709551615) {
                                fmt.Println(v2, v38)
                                if true {
                                    v38 = 1
                                } else {
                                    if false {
                                        v1[v2] = (false || true)
                                    } else {
                                        var v43 string = v18
                                        if true {
                                            var v44 uint = 798
                                            _ = v44
                                        } else {
                                            var v45 float64 = (0.0 + (3.14159 + (1.5 + 1e10)))
                                            var v46 uint64 = 2
                                            v2 -= fi37
                                            _ = v45
                                            _ = v46
                                        }
                                        fmt.Println(v38, v43)
                                        if false {
                                            if true {
                                                if false {
                                                    v47 := GenS3{F0: 2147483647}
                                                    var v48 uint16 = v3
                                                    switch "hello" {
                                                    case "":
                                                        v2 = (^(fi37))
                                                    }
                                                    delete(v1, (fi36 >> ((18446744073709551615) & 63)))
                                                    v18 = "\n\t"
                                                    var v49 uint = 255
                                                    _ = v47
                                                    _ = v48
                                                    _ = v49
                                                } else {
                                                    fmt.Println(v18)
                                                }
                                            } else {
                                                var v50 int16 = -893
                                                _ = v50
                                            }
                                            var v51 int = 1
                                            _ = v51
                                        }
                                        _ = v43
                                    }
                                    v52 := GenS4{F0: 255, F1: 42}
                                    _ = v52
                                }
                            } else {
                                v3 += (v3 - (v3 + v3))
                                var v53 uint = 1
                                _ = v53
                            }
                        case 42:
                            var v54 int16 = -32768
                            _ = v54
                        }
                    }
                    switch (fi36 + fi37) {
                    case 128:
                        if (1 < 28) {
                            if ((float64(float32(1.5)) + (1.5 * (0.0 + (1e10 / 1.5)))) >= (1e10 + (0.1 / 1.5))) {
                                var v55 int16 = 42
                                var v56 int64 = 203
                                v55 += -1
                                _ = v55
                                _ = v56
                            }
                        }
                    case 164:
                        var v57 uint64 = 2
                        _ = v57
                    case 146:
                        if ((32767 > v3) && false) {
                            v58 := GenS2{F0: fi36, F1: 0}
                            _ = v58
                        }
                    }
                    _ = fi37
                    _ = v38
                }
            }
            for fi59 := 0; fi59 < 17; fi59++ {
                v3 = (v3 / ((v3) | 1))
                if false {
                    if true {
                        var v60 int64 = -1
                        v18 = "\n\t"
                        var v61 uint16 = v3
                        _ = v60
                        _ = v61
                    }
                } else {
                    if ((!false) || (false && false)) {
                        v62 := GenS5{F0: fi59, F1: -91, F2: false}
                        switch (fi36 >> ((2) & 63)) {
                        case 753:
                            v63 := GenS6{F0: 0}
                            _ = v63
                        case -2147483648:
                            v64 := []int{1, fi36}
                            _ = v64
                        default:
                            v3 = (v3 | (^(v3)))
                        }
                        if true {
                            var v65 int16 = -32768
                            switch (v18 + "") {
                            case "":
                                fmt.Println(fi59, v3)
                            default:
                                var v66 float32 = (float32(-2.25) / (0.1 + float32(0.0)))
                                _ = v66
                            }
                            if (-1 <= (-(v65))) {
                                fmt.Println(v2, fi59, v18)
                                switch v18 {
                                case "foo bar":
                                    if (true && false) {
                                        v67 := []int{1, 441, v2}
                                        v68 := GenS7{F0: 42, F1: 42, F2: (0.1 != float64((float32(1.5) - ((1e10 / (float32(0.1) - 1e10)) - (-2.25 * (3.14159 - -2.25))))))}
                                        v3 = 65535
                                        if true {
                                            switch "" {
                                            case "foo bar":
                                                var v69 string = v18
                                                _ = v69
                                            case "hello":
                                                v70 := GenS3{F0: -1}
                                                _ = v70
                                            case "\n\t":
                                                switch (v2 & -762) {
                                                case 208:
                                                    v67 = append(v67, (fi59 << ((255) & 63)))
                                                }
                                            }
                                            var v71 int8 = 0
                                            _ = v71
                                        }
                                        if true {
                                            v1[fi36] = false
                                        }
                                        _ = v67
                                        _ = v68
                                    }
                                case "a":
                                    v72 := []int{}
                                    _ = v72
                                default:
                                    switch (v18 + "foo bar") {
                                    case "\n\t":
                                        switch fi36 {
                                        case 202:
                                            v1[2147483647] = false
                                            v3 -= (v3 + 208)
                                            var v73 uint = 65535
                                            _ = v73
                                        }
                                    case "":
                                        var v74 uint16 = v3
                                        _ = v74
                                    }
                                }
                                v1[(^(fi59))] = (-125 >= 83)
                            } else {
                                v75 := GenS8{F0: (194 == 127), F1: false, F2: 1}
                                _ = v75
                            }
                            var v76 int8 = -1
                            var v77 uint16 = 0
                            _ = v65
                            _ = v76
                            _ = v77
                        }
                        var v78 uint64 = 1
                        _ = v62
                        _ = v78
                    }
                    fmt.Println(v3)
                    v3 = 42
                }
                var v79 float32 = (0.1 + 0.0)
                _ = fi59
                _ = v79
            }
            var v80 uint8 = 0
            _ = fi36
            _ = v80
        }
        v81 := []int{(v2 | (v2 << ((42) & 63)))}
        _ = v81
    } else {
        switch (v18 + (("\n\t" + v18) + ("\n\t" + v18))) {
        case "":
            if false {
                if true {
                    var v82 int64 = -9223372036854775808
                    _ = v82
                }
                delete(v1, (v2 >> ((65535) & 63)))
            }
        case "a":
            if true {
                var v83 bool = (true || (!(2 < 0)))
                _ = v83
            }
        case "hello":
            var v84 uint = 65535
            _ = v84
        default:
            v2 += (^(v2))
        }
    }
    v85 := []int{(v2 + (v2 + 4294967295)), (v2 | -599), v2}
    v3 = v3
    switch v2 {
    case -1:
        v3 += v3
    case -433:
        for fi86 := 0; fi86 < 2; fi86++ {
            var v87 float64 = 0.0
            for fi88 := 0; fi88 < 11; fi88++ {
                var v89 int32 = 0
                _ = fi88
                _ = v89
            }
            v18 = "hello"
            _ = fi86
            _ = v87
        }
    case -55:
        for fi90 := 0; fi90 < 19; fi90++ {
            v91 := GenS4{F0: 65535, F1: -416}
            _ = fi90
            _ = v91
        }
    default:
        for fi92 := 0; fi92 < 20; fi92++ {
            fmt.Println(fi92, v2, v18)
            var v93 float32 = float32(float64(0.0))
            _ = fi92
            _ = v93
        }
    }
    v85 = append(v85, (v2 | v2))
    v94 := GenS5{F0: v2, F1: 127, F2: (false && ((((1e10 + 0.0) / 1.0) / 1.0) != (0.1 + 3.14159)))}
    fmt.Println(v18, v2)
    var v95 int = (v2 * (v2 & 1))
    switch -1 {
    case 1:
        var v96 uint8 = 128
        _ = v96
    case 32768:
        v18 = v18
    default:
        v1[(-(v2))] = false
    }
    v3 -= (v3 % (((v3 ^ (^(v3)))) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v85:", len(v85), v85)
    fmt.Println("v95:", v95)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v85
    _ = v94
    _ = v95
}
