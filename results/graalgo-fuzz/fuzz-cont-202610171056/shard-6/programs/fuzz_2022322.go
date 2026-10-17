package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 string }
type GenS2 struct { F0 int64; F1 uint64; F2 int64 }
type GenS3 struct { F0 uint64; F1 int8 }
type GenS4 struct { F0 uint }
func genFn0() bool {
    v7 := []int{9223372036854775807, 117, 1}
    var v8 float64 = 0.0
    for fi9 := 0; fi9 < 13; fi9++ {
        if ((false && false) || (!(true || false))) {
            switch (fi9 - (fi9 - fi9)) {
            case 2147483647:
                if true {
                    v8 = 3.14159
                    v10 := GenS1{F0: -362, F1: false, F2: "a"}
                    v11 := GenS1{F0: -9223372036854775808, F1: false, F2: "hello"}
                    _ = v10
                    _ = v11
                } else {
                    var v12 float64 = ((1e10 * (1.5 - 1e10)) + -2.25)
                    var v13 uint32 = 912
                    var v14 int8 = -128
                    _ = v12
                    _ = v13
                    _ = v14
                }
            default:
                v15 := GenS1{F0: 380, F1: (491 >= 2), F2: "a"}
                _ = v15
            }
            var v16 bool = false
            var v17 int8 = int8(fi9)
            _ = v16
            _ = v17
        } else {
            v7 = append(v7, fi9)
        }
        for fi18 := 0; fi18 < 20; fi18++ {
            var v19 uint8 = 1
            _ = fi18
            _ = v19
        }
        if (((float64((float32(v8) / 0.1)) + float64((((0.1 * -2.25) / (-2.25 / 1.5)) / -2.25))) / v8) == 0.1) {
            if (!true) {
                for fi20 := 0; fi20 < 17; fi20++ {
                    v8 = (v8 - (0.1 - v8))
                    switch "" {
                    case "hello":
                        if (("hello" + ("" + "\n\t")) <= ("a" + ("a" + ""))) {
                            _ = len(v7)
                            v8 = (v8 + v8)
                            var v21 uint16 = 1
                            _ = v21
                        } else {
                            if (255 == 128) {
                                v22 := GenS1{F0: 437, F1: true, F2: "hello"}
                                var v23 uint64 = 966
                                _ = v22
                                _ = v23
                            } else {
                                if false {
                                    fmt.Println(v8, fi20)
                                } else {
                                    var v24 int8 = 92
                                    _ = v24
                                }
                            }
                            v8 = 3.14159
                            if (!(!((v8 * 3.14159) > (v8 - v8)))) {
                                fmt.Println(fi20, v8, fi20)
                            } else {
                                var v25 uint32 = 2147483648
                                fmt.Println(v25, v25)
                                _ = v25
                            }
                        }
                    default:
                        v26 := GenS2{F0: 0, F1: 1, F2: -1}
                        _ = v26
                    }
                    if false {
                        v8 += v8
                    }
                    _ = fi20
                }
                v8 += 0.0
                switch (("" + "\n\t") + ("\n\t" + "hello")) {
                case "\n\t":
                    var v27 uint16 = 255
                    _ = v27
                default:
                    for fi28 := 0; fi28 < 10; fi28++ {
                        if ((-1 > 1) || false) {
                            fmt.Println(v8, fi9)
                            if len(v7) > 0 {
                                v7[0] = fi9
                            }
                            if (1 > 4294967295) {
                                v8 += 0.1
                                fmt.Println(fi28)
                                fmt.Println(v8, v8, v8)
                            } else {
                                v8 -= v8
                            }
                        } else {
                            switch ("" + "foo bar") {
                            case "a":
                                var v29 int16 = 32767
                                _ = v29
                            case "":
                                _ = len(v7)
                            }
                            var v30 int = (fi28 / ((fi28) | 1))
                            _ = v30
                        }
                        _ = fi28
                    }
                }
            }
            switch fi9 {
            case -1:
                fmt.Println(v8, v8)
            case 9223372036854775807:
                if true {
                    var v31 uint = 4294967295
                    v32 := GenS1{F0: 1, F1: false, F2: "foo bar"}
                    if true {
                        var v33 int8 = 0
                        var v34 float32 = -2.25
                        _ = v33
                        _ = v34
                    } else {
                        fmt.Println(v8, v31)
                    }
                    _ = v31
                    _ = v32
                }
            }
            switch (("hello" + "hello") + ("\n\t" + "\n\t")) {
            case "hello":
                var v35 float64 = v8
                _ = v35
            case "a":
                var v36 uint8 = 87
                _ = v36
            case "\n\t":
                v8 += float64(3.14159)
            default:
                var v37 int = (fi9 - fi9)
                _ = v37
            }
        } else {
            v38 := GenS3{F0: 2, F1: -1}
            fmt.Println(fi9)
            _ = v38
        }
        _ = fi9
    }
    var v39 int32 = 42
    _ = v39
    _ = v7
    _ = v8
    return (!true)
}

func main() {
    v1 := make(map[string]int)
    var v2 uint = 255
    var v3 string = "foo bar"
    var v4 uint32 = 198
    var v5 bool = (!(false && (!false)))
    v6 := make(map[int]string)
    switch "a" {
    case "a":
        var v40 float64 = 1.5
        _ = v40
    default:
        switch v3 {
        case "":
            if (false || v5) {
                v4 += v4
                _ = len(v6)
            } else {
                v4 -= (v4 | v4)
                v41 := GenS2{F0: 0, F1: 4294967295, F2: -9223372036854775808}
                fmt.Println(v5, v5)
                _ = v41
            }
        case "foo bar":
            if (v5 || (97 > -1)) {
                if (-179 != 728) {
                    v42 := GenS4{F0: 42}
                    _ = v42
                } else {
                    var v43 int32 = -2147483648
                    _ = v43
                }
                switch "" {
                case "":
                    var v44 int8 = -1
                    _ = v44
                case "hello":
                    fmt.Println(v2, v2, v2)
                case "foo bar":
                    for fi45 := 0; fi45 < 15; fi45++ {
                        v46 := []int{fi45, -1}
                        if (!v5) {
                            fmt.Println(v4, v5)
                        } else {
                            var v47 float64 = (-2.25 * 0.1)
                            _ = v47
                        }
                        _ = fi45
                        _ = v46
                    }
                }
            }
        }
    }
    for fi48 := 0; fi48 < 10; fi48++ {
        fmt.Println(fi48, v4)
        fmt.Println(v5, v2)
        _ = fi48
    }
    v49 := GenS4{F0: (v2 | v2)}
    var v50 int16 = 1
    switch -520 {
    case 127:
        v1[((v3 + "foo bar") + "\n\t")] = -647
    case 128:
        for fi51 := 0; fi51 < 15; fi51++ {
            v50 = (v50 << (((v2 >> ((255) & 63))) & 63))
            v52 := GenS2{F0: -456, F1: 25, F2: 0}
            v53 := v1[v3]
            _ = fi51
            _ = v52
            _ = v53
        }
    }
    for fi54 := 0; fi54 < 13; fi54++ {
        if ((3.14159 * float64(1e10)) < (1e10 - 1.5)) {
            for fi55 := 0; fi55 < 10; fi55++ {
                if v5 {
                    switch (^(fi54)) {
                    case -53:
                        var v56 uint8 = 1
                        _ = v56
                    case 0:
                        if true {
                            fmt.Println(v3)
                            var v57 int64 = -487
                            var v58 uint = (v2 ^ v2)
                            v59 := GenS3{F0: 42, F1: 83}
                            fmt.Println(v4, v50)
                            _ = v57
                            _ = v58
                            _ = v59
                        } else {
                            _ = len(v6)
                            if (true && v5) {
                                _ = len(v6)
                                if v5 {
                                    fmt.Println(v2, fi55, v2)
                                } else {
                                    fmt.Println(v3, fi54, v4)
                                    var v60 int = fi54
                                    v61 := GenS3{F0: 4294967295, F1: -1}
                                    _ = v60
                                    _ = v61
                                }
                            }
                        }
                    case -982:
                        v6[4294967295] = ""
                    default:
                        var v62 bool = (false || ((-2147483648 == 0) || v5))
                        _ = v62
                    }
                    var v63 uint = 1
                    _ = v63
                }
                if true {
                    fmt.Println(fi54, v3, fi54)
                } else {
                    fmt.Println(v2)
                    var v64 float64 = (1.5 / 0.1)
                    delete(v1, ((((((v3 + v3) + "\n\t") + (v3 + (v3 + "a"))) + v3) + ((v3 + v3) + "")) + v3))
                    _ = v64
                }
                _ = fi55
            }
            switch (fi54 + -9223372036854775808) {
            case -1:
                for fi65 := 0; fi65 < 17; fi65++ {
                    delete(v6, 32767)
                    if ((^(v50)) < (v50 << (((v2 & 42)) & 63))) {
                        var v66 bool = (v5 || v5)
                        var v67 float32 = -2.25
                        _ = v66
                        _ = v67
                    }
                    var v68 int = (fi65 | fi54)
                    _ = fi65
                    _ = v68
                }
            case 1:
                v3 = (v3 + v3)
            }
        }
        var v69 float32 = -2.25
        if v5 {
            var v70 uint16 = 255
            if v5 {
                for fi71 := 0; fi71 < 5; fi71++ {
                    v72 := []int{fi71, -463}
                    _ = fi71
                    _ = v72
                }
                v73 := GenS4{F0: (v2 % (((^(v2))) | 1))}
                var v74 uint = 13
                _ = v73
                _ = v74
            }
            var v75 int32 = -2147483648
            _ = v70
            _ = v75
        } else {
            var v76 int16 = 0
            _ = v76
        }
        _ = fi54
        _ = v69
    }
    var v77 string = v3
    var v78 string = (v77 + v3)
    fmt.Println(v2, v5, v50)
    var v79 bool = (-9223372036854775808 < 4294967295)
    _ = len(v1)
    v78 = (("hello" + "") + v77)
    v79 = v5
    v5 = v5
    var v80 int64 = 9223372036854775807
    v81 := []int{}
    var v82 uint64 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v77:", v77)
    fmt.Println("v78:", v78)
    fmt.Println("v79:", v79)
    fmt.Println("v80:", v80)
    fmt.Println("slicelen v81:", len(v81), v81)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v6
    _ = v77
    _ = v78
    _ = v79
    _ = v80
    _ = v81
    _ = v82
}
