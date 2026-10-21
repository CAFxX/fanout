package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 int64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint; F1 uint16 }
type GenS7 struct { F0 int64; F1 int; F2 uint }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 string; F1 bool; F2 uint16 }

func main() {
    v1 := make(map[int]string)
    var v2 bool = true
    v3 := []int{42, 347, -170}
    v4 := make(map[int]uint32)
    if v2 {
        v5 := GenS1{F0: 2147483647, F1: (v2 || (128 != 32767)), F2: 0}
        if (127 > -1) {
            if false {
                var v6 float64 = (1.5 / 1.5)
                _ = v6
            }
            v7 := GenS2{F0: 42}
            var v8 int64 = 42
            _ = v7
            _ = v8
        } else {
            v2 = ((1 >= 0) || ((2 < 1) || v2))
            for fi9 := 0; fi9 < 12; fi9++ {
                v2 = (!(-1 == -459))
                var v10 uint16 = 65535
                fmt.Println(v2, v2, v10)
                _ = fi9
                _ = v10
            }
        }
        for fi11 := 0; fi11 < 3; fi11++ {
            var v12 int32 = -2147483648
            var v13 bool = v2
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = v5
    } else {
        v14 := GenS1{F0: 1, F1: (42 < 9223372036854775807), F2: 0}
        fmt.Println(v2)
        _ = v14
    }
    var v15 bool = (false && (!false))
    for fi16 := 0; fi16 < 18; fi16++ {
        v17 := GenS2{F0: 127}
        _ = fi16
        _ = v17
    }
    switch "" {
    case "\n\t":
        var v18 int32 = 2147483647
        _ = v18
    case "a":
        switch -264 {
        case 42:
            v15 = (-128 < 66)
        default:
            if v15 {
                v15 = (0 > -1)
                for fi19 := 0; fi19 < 2; fi19++ {
                    var v20 uint64 = 0
                    v15 = (0.0 <= float64(1.5))
                    _ = fi19
                    _ = v20
                }
                var v21 float64 = 1.5
                _ = v21
            }
        }
    }
    for fi22 := 0; fi22 < 6; fi22++ {
        if len(v3) > 0 {
            v3[0] = -487
        }
        switch fi22 {
        case 408:
            v23 := GenS3{F0: 2}
            _ = v23
        case -9223372036854775808:
            if (!v2) {
                v15 = (42 == 356)
            }
        default:
            var v24 uint32 = 42
            _ = v24
        }
        switch fi22 {
        case -9223372036854775808:
            for fi25 := 0; fi25 < 2; fi25++ {
                v2 = v15
                if ((((1e10 / (float32(1e10) + (3.14159 + -2.25))) * ((3.14159 * (((((((1.5 - 0.0) + 1.5) - -2.25) - 0.1) / 1.5) - 3.14159) * 3.14159)) + (3.14159 - 0.1))) + (float32(0.1) * (float32(0.0) - 0.0))) == 0.0) {
                    if v2 {
                        fmt.Println(fi25, fi25)
                        var v26 float64 = (-2.25 * -2.25)
                        _ = v26
                    } else {
                        if true {
                            var v27 bool = (4294967295 >= 4294967295)
                            _ = v27
                        } else {
                            v28 := GenS1{F0: -1, F1: false, F2: -773}
                            _ = v28
                        }
                        var v29 string = ("foo bar" + "foo bar")
                        v2 = (0 != -1)
                        _ = v29
                    }
                    var v30 uint8 = 95
                    fmt.Println(v2, fi22)
                    var v31 bool = (!false)
                    fmt.Println(v15, v15, v30)
                    _ = v30
                    _ = v31
                } else {
                    switch (fi25 ^ 839) {
                    case 1:
                        if v2 {
                            v32 := GenS1{F0: -2147483648, F1: v2, F2: 42}
                            _ = v32
                        } else {
                            var v33 int32 = -699
                            _ = v33
                        }
                    default:
                        _ = len(v3)
                    }
                }
                var v34 uint8 = 1
                _ = fi25
                _ = v34
            }
        }
        _ = fi22
    }
    v35 := GenS4{F0: ""}
    if (18446744073709551615 != 552) {
        if len(v3) > 0 {
            v3[0] = -2147483648
        }
        if (2 == 42) {
            _ = len(v4)
            var v36 int8 = 44
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = v36
        }
        var v37 int16 = 32767
        _ = v37
    } else {
        if v15 {
            var v38 uint = 99
            v38 = v38
            if (!(255 > 219)) {
                v2 = v2
                v15 = (0 <= -2147483648)
            }
            _ = v38
        } else {
            for fi39 := 0; fi39 < 5; fi39++ {
                if false {
                    var v40 uint8 = 2
                    var v41 string = ("\n\t" + "\n\t")
                    _ = v40
                    _ = v41
                }
                _ = fi39
            }
        }
        _ = len(v3)
        switch "\n\t" {
        case "":
            v42 := GenS5{F0: 849}
            _ = v42
        }
    }
    delete(v4, 128)
    _ = len(v3)
    if (!v2) {
        v15 = v15
        fmt.Println(v2, v15)
    } else {
        for fi43 := 0; fi43 < 11; fi43++ {
            var v44 int = fi43
            _ = len(v4)
            _ = fi43
            _ = v44
        }
        v45 := GenS3{F0: 1}
        fmt.Println(v15)
        _ = v45
    }
    var v46 float32 = 0.0
    for fi47 := 0; fi47 < 15; fi47++ {
        if true {
            switch "hello" {
            case "":
                v46 += 0.1
            case "foo bar":
                v48 := GenS1{F0: -453, F1: false, F2: -1}
                _ = v48
            case "\n\t":
                if v2 {
                    v49 := GenS2{F0: 110}
                    _ = v49
                }
            default:
                v50 := GenS4{F0: "\n\t"}
                _ = v50
            }
            v51 := GenS4{F0: "foo bar"}
            if true {
                if v2 {
                    for fi52 := 0; fi52 < 3; fi52++ {
                        v53 := GenS6{F0: 42, F1: 0}
                        var v54 uint64 = 1
                        if v2 {
                            switch (fi52 % ((fi47) | 1)) {
                            case -1:
                                var v55 uint16 = 42
                                _ = v55
                            case 1:
                                if len(v3) > 0 {
                                    v3[0] = fi47
                                }
                            case -794:
                                fmt.Println(v15, v46)
                                v56 := GenS1{F0: 0, F1: (0 == 1), F2: 9223372036854775807}
                                _ = v56
                            default:
                                var v57 uint8 = 202
                                _ = v57
                            }
                            var v58 uint64 = (v54 & (v54 + (v54 / (((^(v54))) | 1))))
                            _ = v58
                        } else {
                            v2 = (fi52 == (fi52 ^ fi47))
                            var v59 int8 = 42
                            var v60 uint8 = 42
                            _ = v59
                            _ = v60
                        }
                        if (42 <= uint8(v54)) {
                            var v61 uint16 = 1
                            v62 := GenS7{F0: -856, F1: (fi47 * (fi52 % ((2147483647) | 1))), F2: 42}
                            _ = v61
                            _ = v62
                        } else {
                            var v63 float64 = ((0.0 - 0.1) + -2.25)
                            v64 := GenS8{F0: 1}
                            _ = v63
                            _ = v64
                        }
                        _ = fi52
                        _ = v53
                        _ = v54
                    }
                } else {
                    fmt.Println(v15)
                    fmt.Println(v2, fi47)
                }
            } else {
                if (!(v15 && false)) {
                    v46 += v46
                    v65 := GenS7{F0: -1, F1: (fi47 >> ((0) & 63)), F2: 4294967295}
                    var v66 float64 = float64((v46 / ((v46 * 0.1) - (v46 * 0.0))))
                    _ = v65
                    _ = v66
                }
                var v67 int = fi47
                switch 330 {
                case 4294967295:
                    if false {
                        for fi68 := 0; fi68 < 1; fi68++ {
                            if v2 {
                                v2 = true
                                var v69 float32 = (0.0 / (v46 + v46))
                                _ = v69
                            }
                            _ = fi68
                        }
                        for fi70 := 0; fi70 < 19; fi70++ {
                            var v71 int8 = 100
                            if (0.0 == 1e10) {
                                fmt.Println(v67)
                            } else {
                                var v72 uint = 18446744073709551615
                                _ = v72
                            }
                            var v73 float32 = ((((v46 + v46) * -2.25) - (v46 + 1.5)) / (v46 - 0.0))
                            _ = fi70
                            _ = v71
                            _ = v73
                        }
                        v2 = v15
                    }
                case 32767:
                    fmt.Println(v2)
                case 2147483647:
                    for fi74 := 0; fi74 < 20; fi74++ {
                        if (!v2) {
                            var v75 int32 = 42
                            v76 := GenS2{F0: 255}
                            _ = v75
                            _ = v76
                        } else {
                            _ = len(v3)
                        }
                        _ = fi74
                    }
                }
                _ = v67
            }
            _ = v51
        }
        switch "a" {
        case "\n\t":
            if (!v15) {
                v46 += (float32((((-2.25 + 1.5) * (-2.25 / 0.1)) - 0.1)) * (v46 * float32(3.14159)))
                v46 = 0.1
                if (934 <= 255) {
                    var v77 uint16 = 1
                    var v78 int32 = -1
                    _ = v77
                    _ = v78
                } else {
                    v15 = v15
                }
            }
        case "hello":
            var v79 uint64 = 2
            _ = v79
        case "foo bar":
            if (0 == -974) {
                if (!v15) {
                    for fi80 := 0; fi80 < 15; fi80++ {
                        v2 = false
                        _ = fi80
                    }
                } else {
                    v46 -= (3.14159 + 1e10)
                }
            }
        }
        for fi81 := 0; fi81 < 6; fi81++ {
            v82 := GenS6{F0: 4294967295, F1: 2}
            if false {
                var v83 int = (^(fi81))
                fmt.Println(v15, v2)
                _ = v83
            }
            v84 := GenS8{F0: 128}
            _ = fi81
            _ = v82
            _ = v84
        }
        _ = fi47
    }
    fmt.Println(v2, v46)
    _ = len(v4)
    v85 := GenS9{F0: "a", F1: (0 != 577), F2: 42}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v46
    _ = v85
}
