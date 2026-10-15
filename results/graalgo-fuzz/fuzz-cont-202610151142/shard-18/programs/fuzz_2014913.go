package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int; F1 int16 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int; F1 string }
type GenS8 struct { F0 string; F1 int32; F2 int8 }
type GenS9 struct { F0 uint16; F1 int; F2 int64 }
type GenS10 struct { F0 int; F1 uint64 }
type GenS11 struct { F0 int8; F1 uint; F2 uint }
type GenS12 struct { F0 int; F1 bool; F2 int64 }
type GenS13 struct { F0 int16; F1 uint32 }

func main() {
    var v1 float32 = ((float32(1.5) / (float32(float64(((0.1 / 3.14159) * (-2.25 - (3.14159 * -2.25))))) - (0.0 + (1e10 + float32(((0.1 / 1.5) + 1.5)))))) / (((0.0 - (1e10 / float32(1.5))) + (1e10 - (((1.5 + -2.25) + (float32((1.5 * ((float64(0.0) + 3.14159) + 1.5))) + (((3.14159 * 3.14159) + (((float32((1.5 + (-2.25 / 3.14159))) * -2.25) + (3.14159 + ((1.5 * ((1e10 / 1e10) - 0.1)) - float32(float64((0.0 + float32(0.0))))))) + 0.0)) + (1.5 / (1.5 * 3.14159))))) / float32(1.5)))) * 3.14159))
    var v2 string = ("\n\t" + "a")
    var v3 bool = false
    var v4 string = ("" + v2)
    v5 := []int{9223372036854775807, 9223372036854775807, -544}
    var v6 int16 = -32768
    for fi7 := 0; fi7 < 13; fi7++ {
        if len(v5) > 0 {
            v5[0] = fi7
        }
        v1 -= (v1 - (1.5 / v1))
        _ = fi7
    }
    v8 := GenS1{F0: 1, F1: 127}
    var v9 int16 = (v6 << ((255) & 63))
    v10 := GenS2{F0: ((v6 + -1) >= (^(v6)))}
    v9 -= 32767
    var v11 bool = false
    switch "hello" {
    case "foo bar":
        if v3 {
            if (!v3) {
                for fi12 := 0; fi12 < 4; fi12++ {
                    if (!v3) {
                        v9 -= v6
                        v13 := GenS2{F0: ((v4 + "a") < v2)}
                        var v14 int8 = 42
                        _ = v13
                        _ = v14
                    }
                    switch ((v4 + (v4 + ("\n\t" + v4))) + "a") {
                    case "foo bar":
                        for fi15 := 0; fi15 < 4; fi15++ {
                            var v16 float32 = (v1 / float32(1.5))
                            switch (((((v4 + v4) + "hello") + "foo bar") + v4) + (v4 + "foo bar")) {
                            case "foo bar":
                                var v17 int = fi12
                                _ = v17
                            case "":
                                switch -2147483648 {
                                case -2147483648:
                                    v18 := GenS3{F0: v11}
                                    _ = v18
                                case 1:
                                    var v19 int64 = int64(v9)
                                    _ = v19
                                case 32768:
                                    v16 += 1.5
                                }
                            case "hello":
                                var v20 float32 = v16
                                _ = v20
                            }
                            _ = fi15
                            _ = v16
                        }
                    }
                    _ = fi12
                }
                for fi21 := 0; fi21 < 14; fi21++ {
                    v22 := GenS1{F0: 1, F1: 255}
                    v4 = v2
                    _ = fi21
                    _ = v22
                }
            }
            fmt.Println(v11)
        } else {
            var v23 int16 = v6
            v9 -= (^(v9))
            _ = v23
        }
    case "":
        switch v2 {
        case "\n\t":
            if false {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                fmt.Println(v6, v3, v11)
                v4 = v2
            }
        default:
            v24 := GenS4{F0: 1}
            _ = v24
        }
    case "hello":
        v6 = (v9 * (v9 * (v6 | (v6 | 1000))))
    }
    var v25 uint8 = 127
    var v26 bool = (float64((1e10 + (v1 + (((v1 - 3.14159) * v1) - v1)))) < ((0.0 + float64(3.14159)) - 3.14159))
    switch v4 {
    case "":
        for fi27 := 0; fi27 < 17; fi27++ {
            v28 := GenS5{F0: fi27, F1: v6}
            if (!(-2147483648 == -1)) {
                switch v4 {
                case "":
                    switch int(v6) {
                    case 984:
                        v6 -= 32767
                    case -1:
                        var v29 string = "a"
                        _ = v29
                    }
                case "\n\t":
                    if len(v5) > 0 {
                        v5[0] = (fi27 >> ((2) & 63))
                    }
                case "foo bar":
                    var v30 float64 = float64(v1)
                    _ = v30
                default:
                    if false {
                        if false {
                            v31 := GenS6{F0: 4294967295}
                            switch v4 {
                            case "a":
                                if (!true) {
                                    v32 := GenS6{F0: 2}
                                    _ = v32
                                } else {
                                    switch ("foo bar" + "\n\t") {
                                    case "hello":
                                        _ = len(v5)
                                    case "":
                                        var v33 int32 = -1
                                        _ = v33
                                    case "a":
                                        if (-211 == -9223372036854775808) {
                                            var v34 uint16 = 317
                                            switch fi27 {
                                            case -474:
                                                for fi35 := 0; fi35 < 11; fi35++ {
                                                    switch (("foo bar" + v4) + v4) {
                                                    case "foo bar":
                                                        v2 = ("hello" + ("\n\t" + v2))
                                                    case "hello":
                                                        v2 = ("foo bar" + ((v2 + v2) + v2))
                                                    case "a":
                                                        var v36 int8 = -128
                                                        _ = v36
                                                    default:
                                                        switch fi27 {
                                                        case 32768:
                                                            switch (fi27 >> ((2) & 63)) {
                                                            case 32768:
                                                                if v11 {
                                                                    var v37 string = ("" + ("hello" + ("" + v2)))
                                                                    switch (fi35 + (^(fi27))) {
                                                                    case -53:
                                                                        var v38 int64 = -9223372036854775808
                                                                        _ = v38
                                                                    case 4294967295:
                                                                        var v39 int = fi35
                                                                        _ = v39
                                                                    case -592:
                                                                        v25 = 62
                                                                    }
                                                                    v40 := GenS7{F0: (fi35 << ((943) & 63)), F1: "foo bar"}
                                                                    _ = v37
                                                                    _ = v40
                                                                } else {
                                                                    v41 := GenS8{F0: v2, F1: 42, F2: -65}
                                                                    if (v3 && ((v25 ^ (v25 + (v25 / (((v25 % ((0) | 1))) | 1)))) > v25)) {
                                                                        _ = len(v5)
                                                                        var v42 bool = false
                                                                        var v43 string = (v4 + v4)
                                                                        _ = v42
                                                                        _ = v43
                                                                    }
                                                                    var v44 string = ((v4 + v2) + v2)
                                                                    _ = v41
                                                                    _ = v44
                                                                }
                                                            case 999:
                                                                switch 0 {
                                                                case 266:
                                                                    v45 := GenS3{F0: (2 <= 4294967295)}
                                                                    switch 0 {
                                                                    case 0:
                                                                        var v46 bool = (184 <= 65535)
                                                                        _ = v46
                                                                    default:
                                                                        if (v25 > v25) {
                                                                            var v47 uint32 = 808
                                                                            var v48 uint8 = v25
                                                                            _ = v47
                                                                            _ = v48
                                                                        }
                                                                    }
                                                                    if (((fi27 * fi27) >= (fi35 & -9223372036854775808)) && (!true)) {
                                                                        switch (fi35 | 2147483647) {
                                                                        case 442:
                                                                            var v49 uint32 = 42
                                                                            _ = v49
                                                                        }
                                                                    } else {
                                                                        var v50 int64 = 9223372036854775807
                                                                        var v51 bool = (2 <= v25)
                                                                        _ = v50
                                                                        _ = v51
                                                                    }
                                                                    _ = v45
                                                                case -2147483648:
                                                                    fmt.Println(v2, v11)
                                                                    var v52 bool = v11
                                                                    if len(v5) > 0 {
                                                                        v5[0] = (fi35 | (fi35 | (fi27 + (fi27 & 2147483647))))
                                                                    }
                                                                    fmt.Println(v9, v26, v6)
                                                                    _ = v52
                                                                case 4294967295:
                                                                    var v53 uint16 = (v34 % ((v34) | 1))
                                                                    _ = v53
                                                                default:
                                                                    v54 := GenS9{F0: 65535, F1: (fi27 + fi27), F2: 233}
                                                                    _ = v54
                                                                }
                                                                var v55 int64 = 42
                                                                v9 += 0
                                                                _ = v55
                                                            default:
                                                                var v56 uint8 = 179
                                                                _ = v56
                                                            }
                                                        case 128:
                                                            var v57 int32 = 0
                                                            v11 = (-9223372036854775808 > -1)
                                                            v4 = v2
                                                            var v58 int16 = 1
                                                            _ = v57
                                                            _ = v58
                                                        case -9223372036854775808:
                                                            v11 = false
                                                        }
                                                    }
                                                    v59 := GenS10{F0: -9223372036854775808, F1: 4294967295}
                                                    _ = fi35
                                                    _ = v59
                                                }
                                            case 0:
                                                v1 -= 1.5
                                            case 4294967295:
                                                for fi60 := 0; fi60 < 14; fi60++ {
                                                    var v61 int32 = 1
                                                    fmt.Println(v1)
                                                    _ = len(v5)
                                                    if true {
                                                        var v62 float32 = v1
                                                        var v63 int32 = (v61 & 2147483647)
                                                        _ = v62
                                                        _ = v63
                                                    } else {
                                                        var v64 uint8 = 255
                                                        var v65 float64 = float64(v1)
                                                        var v66 uint64 = 4294967295
                                                        _ = v64
                                                        _ = v65
                                                        _ = v66
                                                    }
                                                    _ = fi60
                                                    _ = v61
                                                }
                                            }
                                            fmt.Println(fi27, v34)
                                            _ = v34
                                        } else {
                                            fmt.Println(v3, v25, v9)
                                            if ((!v11) || true) {
                                                v5 = append(v5, fi27)
                                            }
                                        }
                                    default:
                                        var v67 uint16 = 255
                                        _ = v67
                                    }
                                    fmt.Println(v1, fi27, v9)
                                }
                            case "":
                                v1 -= v1
                            default:
                                fmt.Println(v3, v1, v3)
                            }
                            _ = len(v5)
                            _ = v31
                        }
                        if (839 <= 42) {
                            var v68 int32 = 1
                            var v69 bool = (42 > 0)
                            _ = v68
                            _ = v69
                        } else {
                            var v70 int32 = 1
                            switch fi27 {
                            case -383:
                                v71 := GenS11{F0: 42, F1: 906, F2: 20}
                                _ = v71
                            }
                            switch v2 {
                            case "a":
                                if len(v5) > 0 {
                                    v5[0] = (fi27 << ((398) & 63))
                                }
                            case "\n\t":
                                var v72 float32 = (((v1 * v1) / 0.1) - v1)
                                _ = v72
                            }
                            _ = v70
                        }
                    }
                }
            }
            v73 := GenS12{F0: (-(fi27)), F1: (true && v3), F2: 1}
            _ = fi27
            _ = v28
            _ = v73
        }
    case "a":
        fmt.Println(v25, v25, v2)
    case "hello":
        var v74 uint8 = v25
        _ = v74
    }
    var v75 int64 = 9223372036854775807
    if (float64(v1) == float64(1.5)) {
        v4 = (v2 + "hello")
    }
    v76 := GenS1{F0: -128, F1: (^(v25))}
    if (2147483647 == 2147483647) {
        var v77 int32 = 875
        fmt.Println(v25)
        if (2 >= 2) {
            var v78 int32 = int32(v6)
            var v79 int32 = v78
            _ = v78
            _ = v79
        } else {
            var v80 int32 = v77
            var v81 string = v4
            _ = v80
            _ = v81
        }
        _ = v77
    } else {
        switch 42 {
        case 447:
            v1 -= ((v1 - (v1 + v1)) - float32(((((float64(0.0) / 0.1) * -2.25) - 0.1) * (1.5 / 3.14159))))
        default:
            v82 := GenS3{F0: v26}
            _ = v82
        }
        switch v4 {
        case "hello":
            v83 := GenS13{F0: v6, F1: 386}
            _ = v83
        case "":
            fmt.Println(v75, v6, v3)
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    var v84 uint8 = (v25 & 0)
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v75:", v75)
    fmt.Println("v84:", v84)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v75
    _ = v76
    _ = v8
    _ = v84
    _ = v9
}
