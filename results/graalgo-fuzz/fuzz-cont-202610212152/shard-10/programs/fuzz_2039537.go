package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 int16; F1 int64; F2 uint32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int; F1 uint32 }
type GenS5 struct { F0 int64; F1 uint; F2 string }
type GenS6 struct { F0 string; F1 uint8; F2 uint32 }
type GenS7 struct { F0 uint64; F1 uint8 }
type GenS8 struct { F0 uint; F1 uint64 }
type GenS9 struct { F0 uint8 }

func main() {
    var v1 float64 = (3.14159 - float64(3.14159))
    v2 := make(map[string]int)
    var v3 float32 = (1e10 - 0.0)
    var v4 string = "hello"
    var v5 float32 = 3.14159
    fmt.Println(v1, v3)
    for fi6 := 0; fi6 < 14; fi6++ {
        var v7 int16 = 42
        var v8 int8 = 0
        var v9 uint64 = 0
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    switch v4 {
    case "foo bar":
        var v10 int16 = 1000
        _ = v10
    }
    v3 = float32(v1)
    for fi11 := 0; fi11 < 14; fi11++ {
        switch fi11 {
        case 32768:
            fmt.Println(fi11, v4)
        }
        v5 = (v3 - (v5 + (1.5 / v3)))
        _ = fi11
    }
    v12 := GenS1{F0: 255, F1: -32768}
    var v13 float32 = v3
    if (229 >= 935) {
        switch v4 {
        case "\n\t":
            for fi14 := 0; fi14 < 18; fi14++ {
                fmt.Println(fi14)
                _ = fi14
            }
        case "a":
            for fi15 := 0; fi15 < 1; fi15++ {
                if (956 >= 730) {
                    fmt.Println(v3, v13, v13)
                    fmt.Println(v5)
                    v16 := []int{(fi15 << ((110) & 63)), (fi15 << ((1) & 63)), 825}
                    _ = v16
                } else {
                    if (1 < int16(fi15)) {
                        switch 32768 {
                        case -2147483648:
                            fmt.Println(fi15, v4, v1)
                        case -228:
                            v3 = ((v3 / (v5 - float32(v1))) / ((0.0 + v5) / float32(1.5)))
                        default:
                            fmt.Println(fi15, v1)
                        }
                        fmt.Println(v5, v1, v13)
                    } else {
                        v17 := GenS2{F0: 166, F1: -9223372036854775808, F2: 65535}
                        for fi18 := 0; fi18 < 14; fi18++ {
                            v5 -= v13
                            if true {
                                v13 -= float32(v1)
                            }
                            _ = fi18
                        }
                        _ = v17
                    }
                    if (!true) {
                        var v19 bool = (false || false)
                        v20 := GenS3{F0: -1}
                        _ = v19
                        _ = v20
                    } else {
                        v3 = v3
                        v21 := GenS1{F0: 0, F1: 0}
                        _ = v21
                    }
                    switch -404 {
                    case -1:
                        v5 += 0.1
                    case -2147483648:
                        var v22 uint = 2
                        _ = v22
                    case 32768:
                        for fi23 := 0; fi23 < 1; fi23++ {
                            var v24 bool = true
                            v25 := GenS4{F0: (fi15 >> ((1) & 63)), F1: 2147483648}
                            fmt.Println(fi23, v24, v5)
                            _ = fi23
                            _ = v24
                            _ = v25
                        }
                    }
                }
                _ = fi15
            }
        case "hello":
            switch v4 {
            case "foo bar":
                var v26 uint16 = 0
                _ = v26
            case "\n\t":
                if false {
                    fmt.Println(v3, v1, v3)
                    v3 = float32(v1)
                    v27 := GenS5{F0: 152, F1: 368, F2: v4}
                    _ = v27
                } else {
                    v13 -= v3
                    var v28 int16 = -833
                    switch -710 {
                    case -2147483648:
                        v13 -= 1.5
                    case -1:
                        delete(v2, ("a" + v4))
                    case 127:
                        v1 = (((0.0 * float64(0.1)) + 0.1) + (((v1 - v1) * ((v1 - v1) / (float64(v3) * v1))) + ((3.14159 - (((v1 / (3.14159 / float64(((v5 + v3) * (v3 / 0.1))))) - v1) - v1)) + 3.14159)))
                    default:
                        v1 = float64(1.5)
                    }
                    _ = v28
                }
            }
        }
        v3 = v5
        if (0.0 > (v13 / v3)) {
            v4 = "foo bar"
        } else {
            if true {
                v29 := []int{-363, -653, 469, 4294967295}
                _ = v29
            } else {
                v5 -= v13
                v30 := GenS2{F0: 1, F1: 1, F2: 645}
                v3 = 1e10
                _ = v30
            }
            v1 = v1
        }
    }
    v31 := []int{-370, 1, 42, -9223372036854775808}
    switch -326 {
    case 4294967295:
        v32 := GenS6{F0: (v4 + (v4 + ("hello" + "a"))), F1: 0, F2: 1}
        _ = v32
    default:
        v33 := v2[(v4 + v4)]
        _ = v33
    }
    fmt.Println(v1, v3, v3)
    for fi34 := 0; fi34 < 3; fi34++ {
        v35 := GenS1{F0: 733, F1: -32768}
        var v36 uint = 0
        if false {
            var v37 uint32 = 4294967295
            if (v13 != 1.5) {
                v4 = v4
            } else {
                switch 148 {
                case 9223372036854775807:
                    v5 -= (float32((v1 - (v1 / (((v1 * ((v1 + 1e10) + (v1 / (v1 * (v1 + v1))))) - (v1 + float64(v3))) + (float64(v13) * (3.14159 + (v1 * ((v1 - 1e10) * float64(v3))))))))) - (3.14159 - 1e10))
                default:
                    v3 = (v3 * v13)
                }
                if ((v1 - 3.14159) > (float64(v3) + ((0.1 * float64(v3)) - (3.14159 / v1)))) {
                    switch v4 {
                    case "hello":
                        var v38 bool = true
                        _ = v38
                    case "foo bar":
                        v39 := GenS7{F0: 18446744073709551615, F1: 2}
                        _ = v39
                    case "\n\t":
                        _ = len(v31)
                    default:
                        if true {
                            v40 := GenS2{F0: -32768, F1: -309, F2: (v37 & v37)}
                            _ = v40
                        }
                    }
                    var v41 float64 = v1
                    _ = v41
                }
                var v42 float64 = 0.1
                _ = v42
            }
            var v43 uint8 = 127
            _ = v37
            _ = v43
        }
        _ = fi34
        _ = v35
        _ = v36
    }
    v3 += float32((v1 - (0.0 - v1)))
    if false {
        v1 = v1
    } else {
        if ((1 != 65535) && true) {
            v44 := GenS8{F0: 86, F1: 447}
            for fi45 := 0; fi45 < 11; fi45++ {
                if len(v31) > 0 {
                    v31[0] = (^(fi45))
                }
                var v46 int8 = -86
                switch -1 {
                case 32767:
                    fmt.Println(v3)
                case 42:
                    for fi47 := 0; fi47 < 7; fi47++ {
                        switch (fi47 + fi45) {
                        case 654:
                            var v48 uint16 = 168
                            _ = v48
                        case -303:
                            fmt.Println(v46)
                        case -1:
                            if len(v31) >= 2 {
                                v31 = v31[1:len(v31)]
                            }
                        }
                        v31 = append(v31, (fi45 & 32767))
                        v49 := GenS4{F0: (fi45 & fi45), F1: 646}
                        _ = fi47
                        _ = v49
                    }
                }
                _ = fi45
                _ = v46
            }
            _ = v44
        } else {
            var v50 string = ""
            v51 := v2["\n\t"]
            _ = v50
            _ = v51
        }
    }
    v1 = (1e10 - 1.5)
    var v52 bool = false
    fmt.Println(v4)
    if v52 {
        if (1 <= 2) {
            v5 = v5
        } else {
            for fi53 := 0; fi53 < 8; fi53++ {
                fmt.Println(fi53, v52)
                v1 = ((((v1 - float64(v13)) / v1) / ((0.1 * (v1 * 1.5)) * v1)) * (float64((1.5 - v3)) / v1))
                fmt.Println(v13, v13, v5)
                _ = fi53
            }
            for fi54 := 0; fi54 < 11; fi54++ {
                if (false || false) {
                    for fi55 := 0; fi55 < 1; fi55++ {
                        v4 = (((v4 + v4) + "foo bar") + (v4 + "\n\t"))
                        var v56 uint8 = 42
                        _ = fi55
                        _ = v56
                    }
                    v2[(v4 + v4)] = (fi54 - fi54)
                    v13 = (v5 - 1e10)
                } else {
                    switch fi54 {
                    case 127:
                        v4 = ((v4 + ("\n\t" + v4)) + "foo bar")
                    case 9223372036854775807:
                        if v52 {
                            switch ((((((v4 + "") + v4) + v4) + ("" + "a")) + (v4 + (((v4 + v4) + (v4 + ("a" + v4))) + ""))) + "hello") {
                            case "foo bar":
                                var v57 uint16 = 65535
                                _ = v57
                            case "":
                                fmt.Println(fi54, v52, fi54)
                            case "a":
                                v13 = v13
                            }
                            v58 := v2[((v4 + v4) + ((v4 + v4) + v4))]
                            var v59 uint8 = 2
                            _ = v58
                            _ = v59
                        } else {
                            for fi60 := 0; fi60 < 3; fi60++ {
                                var v61 uint16 = 255
                                _ = fi60
                                _ = v61
                            }
                            var v62 bool = true
                            _ = v62
                        }
                    case 1:
                        v13 += (((3.14159 - v13) / 1.5) / (0.0 * v3))
                    default:
                        switch fi54 {
                        case 9223372036854775807:
                            fmt.Println(v13, fi54, v52)
                        case 2147483647:
                            fmt.Println(v4)
                        case 127:
                            v5 += ((((((1.5 / v5) / v13) / 3.14159) + v13) / 1.5) * 1.5)
                        }
                    }
                    switch (fi54 >> ((0) & 63)) {
                    case 858:
                        v3 = v13
                    case 1:
                        for fi63 := 0; fi63 < 3; fi63++ {
                            v5 = (v3 * (v13 / (v3 - v3)))
                            fmt.Println(v13, fi54, v1)
                            var v64 string = v4
                            _ = fi63
                            _ = v64
                        }
                    }
                }
                v3 += v13
                _ = fi54
            }
            var v65 int64 = -501
            _ = v65
        }
        v3 += float32(float64(v13))
        switch 85 {
        case 32767:
            for fi66 := 0; fi66 < 8; fi66++ {
                switch v4 {
                case "":
                    if (814 < 65535) {
                        if ((916 >= 58) || v52) {
                            if (v52 || ((42 < 1) && v52)) {
                                v2[v4] = (fi66 >> ((42) & 63))
                                var v67 uint = 255
                                _ = v67
                            } else {
                                var v68 uint = 40
                                v68 = v68
                                _ = v68
                            }
                            if (0 >= -1) {
                                v69 := GenS9{F0: 1}
                                v70 := GenS2{F0: -1, F1: 1, F2: 274}
                                _ = v69
                                _ = v70
                            }
                            switch v4 {
                            case "\n\t":
                                var v71 int = fi66
                                _ = v71
                            case "a":
                                v72 := GenS2{F0: 32767, F1: 1, F2: 570}
                                _ = v72
                            case "hello":
                                var v73 uint = 1
                                _ = v73
                            default:
                                fmt.Println(fi66, v3, v4)
                            }
                        } else {
                            v74 := GenS1{F0: 65535, F1: 42}
                            if (2147483648 == 155) {
                                v3 = v13
                                if v52 {
                                    switch 9223372036854775807 {
                                    case 0:
                                        var v75 int = (fi66 << ((296) & 63))
                                        _ = v75
                                    case 4294967295:
                                        for fi76 := 0; fi76 < 6; fi76++ {
                                            v52 = (1 >= 610)
                                            var v77 uint32 = 249
                                            var v78 uint = 4294967295
                                            _ = fi76
                                            _ = v77
                                            _ = v78
                                        }
                                    case 820:
                                        var v79 int8 = 100
                                        _ = v79
                                    }
                                    v1 -= v1
                                }
                                for fi80 := 0; fi80 < 17; fi80++ {
                                    v4 = ((("" + v4) + v4) + ("hello" + "foo bar"))
                                    _ = fi80
                                }
                            }
                            _ = v74
                        }
                        var v81 uint64 = 317
                        v82 := GenS4{F0: (fi66 * fi66), F1: 42}
                        _ = v81
                        _ = v82
                    } else {
                        v5 = float32(3.14159)
                    }
                }
                if (553 >= 4294967295) {
                    var v83 int16 = 32767
                    var v84 int8 = 127
                    _ = v83
                    _ = v84
                } else {
                    switch fi66 {
                    case 9223372036854775807:
                        v85 := GenS1{F0: 1, F1: -1}
                        _ = v85
                    case 128:
                        var v86 int32 = 42
                        _ = v86
                    default:
                        switch v4 {
                        case "\n\t":
                            v2[v4] = 128
                        default:
                            var v87 string = v4
                            _ = v87
                        }
                    }
                    fmt.Println(v52)
                    fmt.Println(v52)
                }
                _ = fi66
            }
        }
    } else {
        var v88 uint32 = 2
        for fi89 := 0; fi89 < 9; fi89++ {
            var v90 string = v4
            var v91 string = ((v4 + (v4 + (((((v90 + v90) + ("" + (((v4 + (v90 + "\n\t")) + (v4 + (v90 + "hello"))) + v4))) + "\n\t") + v4) + "hello"))) + v90)
            _ = fi89
            _ = v90
            _ = v91
        }
        _ = v88
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v52
}
