package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int; F1 uint16; F2 int }
type GenS6 struct { F0 string; F1 uint; F2 int32 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 string; F1 int64; F2 uint16 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 bool; F1 int16 }

func main() {
    var v1 bool = (!true)
    v2 := []int{255}
    v3 := make(map[int]bool)
    var v4 bool = false
    var v5 int16 = 1000
    if v4 {
        switch 734 {
        case 396:
            v6 := GenS1{F0: true, F1: 544, F2: (v5 + 1000)}
            _ = v6
        }
    }
    if v1 {
        delete(v3, 4294967295)
    }
    var v7 float64 = 1e10
    v8 := GenS1{F0: v4, F1: 65535, F2: 463}
    if len(v2) > 0 {
        v2[0] = -253
    }
    if (v4 && (4294967295 == 2147483647)) {
        v7 = 0.0
        for fi9 := 0; fi9 < 11; fi9++ {
            var v10 uint16 = 1
            var v11 bool = (v10 < (v10 ^ v10))
            _ = fi9
            _ = v10
            _ = v11
        }
        v12 := v3[-2147483648]
        _ = v12
    } else {
        switch ("hello" + "a") {
        case "":
            v13 := GenS1{F0: v1, F1: 1, F2: (-(v5))}
            _ = v13
        case "\n\t":
            fmt.Println(v5, v7)
        case "foo bar":
            v4 = v1
        default:
            v7 -= v7
        }
        for fi14 := 0; fi14 < 5; fi14++ {
            if ((0 == -747) || (238 != 454)) {
                if (!false) {
                    v15 := v3[(fi14 | (fi14 & fi14))]
                    _ = v15
                }
                var v16 string = (("" + "foo bar") + ("hello" + "hello"))
                _ = v16
            } else {
                for fi17 := 0; fi17 < 15; fi17++ {
                    if v4 {
                        if v1 {
                            var v18 uint64 = 18446744073709551615
                            _ = v18
                        }
                    }
                    if (1 <= 1) {
                        var v19 uint8 = 255
                        v20 := GenS2{F0: 1}
                        if v4 {
                            v7 = 0.0
                        } else {
                            var v21 int8 = -128
                            v22 := GenS3{F0: -1}
                            _ = v21
                            _ = v22
                        }
                        _ = v19
                        _ = v20
                    } else {
                        v7 = v7
                        v1 = (v4 || false)
                        v7 -= float64(-2.25)
                    }
                    var v23 uint16 = 0
                    v5 += 318
                    var v24 uint8 = 0
                    _ = fi17
                    _ = v23
                    _ = v24
                }
                switch "hello" {
                case "a":
                    if v1 {
                        v25 := GenS1{F0: ("a" > "a"), F1: 2147483647, F2: (^(v5))}
                        switch fi14 {
                        case 127:
                            var v26 uint8 = 127
                            _ = v26
                        }
                        _ = v25
                    } else {
                        var v27 float64 = (3.14159 / v7)
                        fmt.Println(v7, v7)
                        _ = v27
                    }
                default:
                    var v28 int8 = -128
                    _ = v28
                }
            }
            v29 := GenS4{F0: 2}
            _ = fi14
            _ = v29
        }
    }
    v30 := GenS2{F0: v5}
    for fi31 := 0; fi31 < 17; fi31++ {
        v32 := GenS5{F0: 42, F1: 255, F2: 163}
        var v33 float32 = 1.5
        if v1 {
            for fi34 := 0; fi34 < 11; fi34++ {
                var v35 int16 = v5
                v1 = (-1 >= 1)
                if ((v7 / (float64(1e10) / 0.1)) >= (1.5 + -2.25)) {
                    if ((fi31 / ((fi31) | 1)) == fi34) {
                        v33 = float32(((0.0 / v7) + (v7 / v7)))
                        if v4 {
                            var v36 int8 = -109
                            _ = v36
                        } else {
                            v37 := GenS6{F0: "", F1: 745, F2: 0}
                            var v38 int16 = (v35 % ((v5) | 1))
                            if ((fi34 ^ 703) < -478) {
                                var v39 uint = 2
                                var v40 uint64 = 2
                                v39 = (v39 % ((v39) | 1))
                                _ = v39
                                _ = v40
                            } else {
                                var v41 bool = v1
                                v2 = append(v2, 0)
                                v42 := GenS6{F0: "foo bar", F1: 2, F2: 2147483647}
                                _ = v41
                                _ = v42
                            }
                            _ = v37
                            _ = v38
                        }
                        v35 = (v35 * v35)
                    }
                    var v43 bool = v4
                    switch "a" {
                    case "\n\t":
                        var v44 int64 = 0
                        _ = v44
                    case "foo bar":
                        v3[(fi34 % (((fi31 & fi31)) | 1))] = v43
                    default:
                        var v45 string = "foo bar"
                        _ = v45
                    }
                    _ = v43
                } else {
                    v46 := GenS7{F0: false}
                    _ = v46
                }
                _ = fi34
                _ = v35
            }
            var v47 bool = v4
            v48 := GenS8{F0: (("" + ("\n\t" + (("hello" + ("\n\t" + "foo bar")) + "foo bar"))) + "\n\t"), F1: -897, F2: 1}
            _ = v47
            _ = v48
        } else {
            var v49 float64 = (((((1.5 + (1.5 * (v7 - v7))) - 1e10) - v7) - ((float64((float32(float64((v33 / 3.14159))) + v33)) * v7) + (0.0 / v7))) * (v7 / float64(v33)))
            v50 := GenS9{F0: 42}
            _ = v49
            _ = v50
        }
        _ = fi31
        _ = v32
        _ = v33
    }
    if (!v4) {
        if v1 {
            var v51 string = "a"
            switch v51 {
            case "a":
                for fi52 := 0; fi52 < 4; fi52++ {
                    v53 := GenS3{F0: -691}
                    if ((42 != 18446744073709551615) && false) {
                        switch "a" {
                        case "":
                            var v54 uint16 = 32767
                            _ = v54
                        case "foo bar":
                            if (true && (1e10 <= float32(v7))) {
                                var v55 bool = v4
                                _ = v55
                            } else {
                                v2 = append(v2, (-(fi52)))
                                if (0.0 != (0.0 + (0.0 - float32(0.1)))) {
                                    if true {
                                        if (v1 && (!(-244 <= -9223372036854775808))) {
                                            v56 := GenS6{F0: "\n\t", F1: 2, F2: 42}
                                            _ = v56
                                        }
                                        var v57 bool = true
                                        _ = v57
                                    } else {
                                        v58 := GenS3{F0: -532}
                                        _ = v58
                                    }
                                }
                                var v59 uint64 = 1
                                _ = v59
                            }
                        case "\n\t":
                            v2 = append(v2, (fi52 >> ((2) & 63)))
                        }
                        fmt.Println(v51, v7)
                    }
                    _ = fi52
                    _ = v53
                }
            }
            _ = v51
        }
        v60 := GenS2{F0: v5}
        for fi61 := 0; fi61 < 18; fi61++ {
            var v62 float64 = ((float64(1e10) - (1e10 - ((v7 + (v7 * ((v7 - 0.0) * 1e10))) * v7))) * float64((-2.25 * float32(float64(float32((3.14159 * v7)))))))
            if v4 {
                v63 := GenS1{F0: (v1 && false), F1: 65535, F2: 1000}
                _ = v63
            } else {
                var v64 float64 = v62
                _ = len(v3)
                var v65 bool = v1
                _ = v64
                _ = v65
            }
            _ = fi61
            _ = v62
        }
        _ = v60
    } else {
        v66 := GenS10{F0: 0}
        switch ("\n\t" + (("hello" + "\n\t") + "\n\t")) {
        case "hello":
            var v67 int = -477
            _ = v67
        case "":
            if v1 {
                v4 = false
            } else {
                for fi68 := 0; fi68 < 6; fi68++ {
                    for fi69 := 0; fi69 < 2; fi69++ {
                        var v70 float64 = float64(1e10)
                        _ = fi69
                        _ = v70
                    }
                    for fi71 := 0; fi71 < 1; fi71++ {
                        v5 -= (v5 ^ v5)
                        switch fi68 {
                        case 128:
                            var v72 int = (fi71 ^ 0)
                            _ = v72
                        default:
                            var v73 int8 = 1
                            _ = v73
                        }
                        v7 = 0.0
                        var v74 int16 = (v5 & (v5 - v5))
                        _ = fi71
                        _ = v74
                    }
                    var v75 int32 = 0
                    _ = fi68
                    _ = v75
                }
                v4 = false
            }
        case "a":
            v4 = true
        default:
            fmt.Println(v1, v1, v1)
        }
        v76 := GenS11{F0: v4, F1: v5}
        _ = v66
        _ = v76
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
