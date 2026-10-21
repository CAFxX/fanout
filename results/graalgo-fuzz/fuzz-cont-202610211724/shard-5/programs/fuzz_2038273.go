package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32; F1 uint16; F2 int8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint64; F1 uint; F2 uint16 }

func main() {
    var v1 float32 = (-2.25 / (3.14159 + 3.14159))
    var v2 float64 = 0.0
    var v3 int64 = 0
    v4 := make(map[int]uint32)
    var v5 float32 = v1
    v6 := make(map[uint32]uint32)
    var v7 uint = 609
    if (507 >= 0) {
        v5 = v1
    }
    v8 := []int{128, -281, -975, 42}
    if false {
        v5 = v5
    }
    v9 := GenS1{F0: -65, F1: 0, F2: int(v3)}
    var v10 uint16 = 1
    v11 := GenS1{F0: -77, F1: 1, F2: -386}
    for fi12 := 0; fi12 < 4; fi12++ {
        v10 = v10
        _ = fi12
    }
    for fi13 := 0; fi13 < 2; fi13++ {
        for fi14 := 0; fi14 < 18; fi14++ {
            v5 -= (float32(v2) + 0.0)
            _ = fi14
        }
        _ = fi13
    }
    v2 = 0.0
    for fi15 := 0; fi15 < 5; fi15++ {
        v1 = v1
        switch ("hello" + ("" + "\n\t")) {
        case "hello":
            v2 = v2
        case "a":
            var v16 bool = false
            _ = v16
        case "foo bar":
            for fi17 := 0; fi17 < 18; fi17++ {
                v8 = append(v8, (fi17 % ((942) | 1)))
                _ = fi17
            }
        default:
            if len(v8) > 0 {
                v8[0] = (fi15 << ((255) & 63))
            }
        }
        v18 := GenS1{F0: 0, F1: -1, F2: (^(fi15))}
        _ = fi15
        _ = v18
    }
    v5 += v1
    switch -263 {
    case -1:
        if false {
            for fi19 := 0; fi19 < 1; fi19++ {
                if (false && (v7 == v7)) {
                    if ("" != "") {
                        delete(v6, 2)
                        v20 := GenS2{F0: 0}
                        _ = v20
                    }
                    v3 += v3
                    if len(v8) > 0 {
                        v8[0] = fi19
                    }
                } else {
                    v21 := GenS3{F0: 2147483647, F1: v10, F2: 62}
                    v5 = float32(1e10)
                    var v22 float32 = (v5 * v1)
                    _ = v21
                    _ = v22
                }
                var v23 int32 = -1
                var v24 float32 = v1
                _ = fi19
                _ = v23
                _ = v24
            }
            v5 = v1
        } else {
            delete(v4, 652)
            v5 -= float32(v2)
            if false {
                var v25 uint32 = 65535
                _ = v25
            }
        }
    }
    if (127 == 127) {
        v26 := GenS3{F0: 7, F1: (^(v10)), F2: 0}
        for fi27 := 0; fi27 < 4; fi27++ {
            for fi28 := 0; fi28 < 15; fi28++ {
                v6[54] = 2147483648
                switch ((("foo bar" + "foo bar") + "a") + "hello") {
                case "hello":
                    v29 := GenS1{F0: 127, F1: 42, F2: -9223372036854775808}
                    _ = v29
                case "foo bar":
                    if (true && (18446744073709551615 >= 4294967295)) {
                        v1 = 1e10
                        if ((^(v10)) <= v10) {
                            fmt.Println(v7)
                            switch (fi28 >> ((346) & 63)) {
                            case 9223372036854775807:
                                if true {
                                    switch (fi28 & (^(fi28))) {
                                    case -580:
                                        v1 += float32((v2 - (v2 + v2)))
                                    default:
                                        v30 := v4[fi28]
                                        _ = v30
                                    }
                                    var v31 int = (fi28 | 42)
                                    _ = v31
                                } else {
                                    var v32 bool = (!true)
                                    if true {
                                        v1 += -2.25
                                        v10 += v10
                                    }
                                    v7 = (v7 << (((^(v7))) & 63))
                                    if (!(0.0 > v5)) {
                                        v1 += (v1 - 0.1)
                                        switch 9223372036854775807 {
                                        case 770:
                                            if true {
                                                _ = len(v8)
                                                switch "a" {
                                                case "\n\t":
                                                    switch fi27 {
                                                    case 990:
                                                        v10 += (^(v10))
                                                    default:
                                                        if v32 {
                                                            switch "\n\t" {
                                                            case "foo bar":
                                                                delete(v6, 186)
                                                            case "a":
                                                                if len(v8) >= 2 {
                                                                    v8 = v8[1:len(v8)]
                                                                }
                                                                v33 := GenS1{F0: -128, F1: 127, F2: 0}
                                                                _ = v33
                                                            }
                                                            var v34 uint8 = 0
                                                            _ = v34
                                                        } else {
                                                            switch ("\n\t" + "foo bar") {
                                                            case "hello":
                                                                if v32 {
                                                                    if v32 {
                                                                        if ((v3 + v3) != (v3 * (^(v3)))) {
                                                                            v35 := GenS4{F0: -9223372036854775808}
                                                                            var v36 int32 = 42
                                                                            _ = v35
                                                                            _ = v36
                                                                        }
                                                                        v10 = (^(v10))
                                                                        var v37 float64 = (0.0 / 1.0)
                                                                        _ = v37
                                                                    } else {
                                                                        var v38 bool = (42 != 434)
                                                                        var v39 int64 = 42
                                                                        if true {
                                                                            if v32 {
                                                                                v5 = (float32((v2 / v2)) - 0.1)
                                                                            }
                                                                            var v40 bool = ((2 <= 42) && (335 > 1))
                                                                            if ((^(v10)) >= v10) {
                                                                                if (240 < 1) {
                                                                                    var v41 uint64 = 574
                                                                                    v42 := GenS3{F0: -361, F1: v10, F2: -128}
                                                                                    _ = v41
                                                                                    _ = v42
                                                                                }
                                                                                v4[fi28] = 532
                                                                            } else {
                                                                                _ = len(v8)
                                                                                v3 += v3
                                                                                _ = len(v8)
                                                                            }
                                                                            _ = v40
                                                                        } else {
                                                                            if ((!true) || (fi28 == 1)) {
                                                                                var v43 float32 = v5
                                                                                if len(v8) > 0 {
                                                                                    v8[0] = (fi28 >> ((2) & 63))
                                                                                }
                                                                                _ = v43
                                                                            } else {
                                                                                var v44 uint64 = 42
                                                                                if v38 {
                                                                                    fmt.Println(v38)
                                                                                    v7 += (v7 >> ((68) & 63))
                                                                                } else {
                                                                                    var v45 uint = v7
                                                                                    if v32 {
                                                                                        v46 := GenS3{F0: 1, F1: (^(v10)), F2: 0}
                                                                                        _ = v46
                                                                                    } else {
                                                                                        v2 = 1e10
                                                                                        var v47 string = (("foo bar" + (("\n\t" + ("hello" + "a")) + "hello")) + "hello")
                                                                                        fmt.Println(v32, v38)
                                                                                        _ = v47
                                                                                    }
                                                                                    fmt.Println(v1, v2, v38)
                                                                                    _ = v45
                                                                                }
                                                                                _ = v44
                                                                            }
                                                                            if false {
                                                                                fmt.Println(fi28, v2)
                                                                                if len(v8) >= 2 {
                                                                                    v8 = v8[1:len(v8)]
                                                                                }
                                                                            }
                                                                        }
                                                                        _ = v38
                                                                        _ = v39
                                                                    }
                                                                    v48 := v4[-9223372036854775808]
                                                                    _ = v48
                                                                }
                                                            case "foo bar":
                                                                v4[fi28] = 2147483648
                                                            default:
                                                                switch fi28 {
                                                                case 4294967295:
                                                                    v7 += uint(v10)
                                                                case 32768:
                                                                    v49 := GenS3{F0: 1, F1: v10, F2: 0}
                                                                    _ = v49
                                                                case 2147483647:
                                                                    switch (fi28 / (((fi27 | fi28)) | 1)) {
                                                                    case 1:
                                                                        if ((v7 + v7) > v7) {
                                                                            var v50 int64 = 1
                                                                            _ = v50
                                                                        }
                                                                    case -755:
                                                                        v2 = (v2 / (v2 + -2.25))
                                                                    case 376:
                                                                        v6[0] = 0
                                                                    default:
                                                                        if (1e10 > v1) {
                                                                            v51 := GenS4{F0: v3}
                                                                            _ = v51
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                            switch "" {
                                                            case "":
                                                                fmt.Println(v1, v32)
                                                            default:
                                                                v10 = uint16(v7)
                                                            }
                                                            v5 += v1
                                                        }
                                                    }
                                                case "a":
                                                    v52 := GenS5{F0: (fi27 | 32768)}
                                                    _ = v52
                                                case "hello":
                                                    fmt.Println(v3, v10, v2)
                                                }
                                                v3 -= v3
                                            }
                                        case 2147483647:
                                            v53 := GenS6{F0: uint64(v10), F1: v7, F2: v10}
                                            _ = v53
                                        }
                                    } else {
                                        switch "" {
                                        case "a":
                                            v6[65535] = 882
                                        case "hello":
                                            if true {
                                                if v32 {
                                                    v7 -= v7
                                                    var v54 string = "foo bar"
                                                    if len(v8) > 0 {
                                                        v8[0] = fi27
                                                    }
                                                    _ = v54
                                                }
                                                var v55 int8 = -1
                                                var v56 float64 = (v2 * (0.0 / v2))
                                                _ = v55
                                                _ = v56
                                            }
                                        default:
                                            var v57 int8 = -124
                                            _ = v57
                                        }
                                        v3 += (v3 & v3)
                                        var v58 int32 = 1
                                        _ = v58
                                    }
                                    v1 = v5
                                    _ = v32
                                }
                            }
                        } else {
                            v1 += (((((3.14159 - v5) / (v1 - 0.1)) - (((v5 - ((float32(1e10) * 0.0) + v5)) / v5) / (3.14159 / (0.0 / float32(v2))))) - (v1 * v1)) / v5)
                        }
                        if true {
                            var v59 uint8 = 2
                            v2 = v2
                            var v60 uint16 = 0
                            _ = v59
                            _ = v60
                        } else {
                            _ = len(v6)
                            fmt.Println(v5, v1, v5)
                            var v61 bool = (false && false)
                            _ = v61
                        }
                        v10 = (v10 + (v10 % ((897) | 1)))
                    } else {
                        if false {
                            v62 := GenS2{F0: 127}
                            var v63 uint16 = uint16(v7)
                            switch int(v7) {
                            case -2147483648:
                                var v64 float64 = (float64(float32(0.0)) - ((-2.25 - 1e10) + float64((v1 + (float32(float64(0.1)) * v5)))))
                                _ = v64
                            case -84:
                                fmt.Println(v3, v2)
                            case 42:
                                var v65 bool = false
                                _ = v65
                            }
                            _ = v62
                            _ = v63
                        } else {
                            switch "" {
                            case "hello":
                                v66 := GenS4{F0: -1}
                                _ = v66
                            case "\n\t":
                                var v67 uint64 = 1
                                _ = v67
                            case "foo bar":
                                var v68 int = (fi27 / (((fi28 - (fi27 * 32767))) | 1))
                                _ = v68
                            default:
                                v2 += v2
                            }
                        }
                        _ = len(v4)
                    }
                case "\n\t":
                    var v69 int16 = 69
                    _ = v69
                default:
                    var v70 float32 = v1
                    _ = v70
                }
                _ = fi28
            }
            var v71 int8 = 25
            v72 := v6[468]
            _ = fi27
            _ = v71
            _ = v72
        }
        _ = v26
    } else {
        if (("hello" >= (("a" + ((("a" + ("foo bar" + "hello")) + "a") + ("" + (("\n\t" + "foo bar") + "foo bar")))) + "a")) || true) {
            var v73 int16 = -1
            var v74 uint16 = v10
            _ = v73
            _ = v74
        } else {
            fmt.Println(v7)
        }
    }
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
