package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8; F1 bool }
type GenS3 struct { F0 string; F1 int }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int; F1 int64 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint; F1 int8; F2 int32 }
type GenS9 struct { F0 uint16 }

func main() {
    var v1 float64 = 3.14159
    v2 := make(map[int]uint32)
    v3 := []string{("a" + ((("foo bar" + "") + (("foo bar" + ("a" + "hello")) + "hello")) + "hello")), "foo bar", "\n\t", "a"}
    var v4 bool = false
    v5 := []uint32{969, 2147483647, 2147483648}
    v6 := make(map[int]bool)
    switch (("foo bar" + "") + "foo bar") {
    case "hello":
        if false {
            for fi7 := 0; fi7 < 2; fi7++ {
                v8 := GenS1{F0: -2147483648}
                _ = fi7
                _ = v8
            }
            v9 := GenS1{F0: -2147483648}
            _ = v9
        } else {
            fmt.Println(v1, v1)
        }
    case "foo bar":
        _ = len(v3)
    case "a":
        if true {
            for fi10 := 0; fi10 < 11; fi10++ {
                var v11 uint16 = 32767
                v12 := GenS2{F0: -57, F1: v4}
                v4 = (32767 > fi10)
                _ = fi10
                _ = v11
                _ = v12
            }
        } else {
            v4 = (1e10 >= ((0.0 / 1.0) + 1.5))
            v13 := GenS2{F0: -62, F1: (127 >= -1)}
            _ = v13
        }
    }
    var v14 int16 = -1
    var v15 uint32 = 1
    var v16 int16 = (v14 << ((65535) & 63))
    switch "a" {
    case "foo bar":
        v17 := v6[32768]
        _ = v17
    case "hello":
        v14 += 0
    case "\n\t":
        for fi18 := 0; fi18 < 17; fi18++ {
            if (2 <= 128) {
                if false {
                    switch 9223372036854775807 {
                    case 293:
                        for fi19 := 0; fi19 < 11; fi19++ {
                            v20 := GenS3{F0: "\n\t", F1: (fi19 | -1)}
                            v15 = v15
                            v21 := GenS3{F0: "hello", F1: 127}
                            _ = fi19
                            _ = v20
                            _ = v21
                        }
                    case -2147483648:
                        v22 := GenS3{F0: "foo bar", F1: fi18}
                        _ = v22
                    case 9223372036854775807:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    default:
                        var v23 int64 = -334
                        _ = v23
                    }
                } else {
                    for fi24 := 0; fi24 < 5; fi24++ {
                        v25 := GenS3{F0: "a", F1: -9223372036854775808}
                        _ = fi24
                        _ = v25
                    }
                    for fi26 := 0; fi26 < 3; fi26++ {
                        switch "foo bar" {
                        case "hello":
                            if v4 {
                                var v27 uint16 = 2
                                var v28 bool = v4
                                _ = len(v3)
                                _ = v27
                                _ = v28
                            }
                        case "\n\t":
                            if v4 {
                                var v29 uint = 234
                                _ = v29
                            }
                        }
                        _ = fi26
                    }
                }
                if v4 {
                    for fi30 := 0; fi30 < 12; fi30++ {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        if true {
                            fmt.Println(v15, v14, v1)
                            if (!v4) {
                                var v31 int16 = -977
                                _ = v31
                            } else {
                                var v32 int16 = v16
                                v33 := GenS4{F0: (v14 + -495)}
                                _ = v32
                                _ = v33
                            }
                        } else {
                            v4 = v4
                            _ = len(v6)
                        }
                        _ = fi30
                    }
                    v34 := GenS5{F0: 42}
                    _ = v34
                }
                var v35 uint32 = (v15 - (^(v15)))
                _ = v35
            } else {
                if (255 != 604) {
                    if (!(1e10 > ((v1 + (0.1 + ((3.14159 - (v1 - (v1 / -2.25))) * v1))) + v1))) {
                        v16 -= 1
                    } else {
                        for fi36 := 0; fi36 < 2; fi36++ {
                            v37 := GenS6{F0: -2147483648, F1: 9223372036854775807}
                            _ = fi36
                            _ = v37
                        }
                        v1 -= (v1 * ((v1 * v1) * (v1 - ((1e10 - (v1 - (v1 / v1))) - (v1 * v1)))))
                    }
                }
                v38 := GenS6{F0: (fi18 + (fi18 << ((65535) & 63))), F1: 103}
                _ = v38
            }
            v16 = v16
            if v4 {
                v1 -= (0.1 + 0.1)
                v2[(fi18 << ((398) & 63))] = v15
            } else {
                v3 = append(v3, "foo bar")
                if v4 {
                    fmt.Println(v15, v1)
                    v39 := GenS6{F0: 32768, F1: 1}
                    _ = v39
                } else {
                    if ((9223372036854775807 < 326) && v4) {
                        fmt.Println(v16, v16)
                        v4 = v4
                        switch ("hello" + (("foo bar" + ("\n\t" + ("hello" + "hello"))) + "a")) {
                        case "\n\t":
                            switch int(v16) {
                            case 32767:
                                delete(v6, (-(fi18)))
                            case 9223372036854775807:
                                var v40 uint = 657
                                _ = v40
                            case 27:
                                switch "foo bar" {
                                case "foo bar":
                                    fmt.Println(v4, v14)
                                case "\n\t":
                                    var v41 uint16 = 46
                                    _ = v41
                                case "a":
                                    for fi42 := 0; fi42 < 1; fi42++ {
                                        v43 := GenS6{F0: fi42, F1: -9223372036854775808}
                                        switch (fi18 << ((93) & 63)) {
                                        case 9223372036854775807:
                                            var v44 int32 = -1
                                            _ = v44
                                        case -1:
                                            if (v4 && (v4 || v4)) {
                                                if v4 {
                                                    v14 = (v14 & 436)
                                                    v16 += v16
                                                } else {
                                                    v16 -= -974
                                                }
                                            } else {
                                                switch 4294967295 {
                                                case -971:
                                                    v14 = (v14 - 46)
                                                case -493:
                                                    var v45 float64 = 0.0
                                                    _ = v45
                                                }
                                                var v46 uint = 1
                                                fmt.Println(v4, v16)
                                                _ = v46
                                            }
                                        }
                                        var v47 int64 = 42
                                        var v48 float64 = v1
                                        var v49 string = ("foo bar" + "\n\t")
                                        _ = fi42
                                        _ = v43
                                        _ = v47
                                        _ = v48
                                        _ = v49
                                    }
                                default:
                                    fmt.Println(v16, v14, fi18)
                                }
                            }
                        case "a":
                            for fi50 := 0; fi50 < 19; fi50++ {
                                fmt.Println(v14)
                                v51 := GenS5{F0: v15}
                                _ = fi50
                                _ = v51
                            }
                        case "":
                            v52 := GenS3{F0: "\n\t", F1: 4294967295}
                            _ = v52
                        }
                    }
                    v4 = (!((255 < 526) || v4))
                    var v53 uint64 = 1
                    _ = v53
                }
                var v54 uint = 1
                _ = v54
            }
            _ = fi18
        }
    }
    v4 = (true && v4)
    switch 42 {
    case -503:
        var v55 int = -555
        _ = v55
    case -2147483648:
        var v56 float32 = 0.1
        _ = v56
    case 42:
        v5 = append(v5, 4294967295)
    default:
        if false {
            var v57 uint = 816
            v15 += (^(v15))
            _ = v57
        } else {
            v14 = (^(v16))
            v14 += -1
            v58 := GenS7{F0: 42}
            _ = v58
        }
    }
    v59 := v6[2147483647]
    var v60 string = "hello"
    v3 = append(v3, v60)
    for fi61 := 0; fi61 < 5; fi61++ {
        if v4 {
            v3 = append(v3, (v60 + v60))
            if true {
                v59 = (!true)
                if ((2 == 2) && false) {
                    var v62 uint32 = 65535
                    if (!v4) {
                        v4 = true
                    } else {
                        v63 := v6[fi61]
                        v15 = (^(v15))
                        fmt.Println(v62, v1, v14)
                        _ = v63
                    }
                    v64 := v6[(fi61 | (fi61 ^ 4294967295))]
                    _ = v62
                    _ = v64
                }
            } else {
                delete(v2, (fi61 & (fi61 >> ((4294967295) & 63))))
                switch -827 {
                case 852:
                    v65 := GenS5{F0: v15}
                    _ = v65
                case -9223372036854775808:
                    var v66 int32 = 2147483647
                    _ = v66
                case -796:
                    switch "" {
                    case "foo bar":
                        if (!true) {
                            switch fi61 {
                            case 9223372036854775807:
                                if true {
                                    v60 = ((v60 + v60) + v60)
                                }
                            default:
                                v67 := GenS4{F0: (v16 | (v16 << ((4294967295) & 63)))}
                                _ = v67
                            }
                            var v68 int64 = 42
                            v69 := GenS3{F0: v60, F1: fi61}
                            _ = v68
                            _ = v69
                        } else {
                            for fi70 := 0; fi70 < 14; fi70++ {
                                v14 -= (v14 << ((65535) & 63))
                                _ = fi70
                            }
                            if true {
                                var v71 uint64 = 42
                                _ = v71
                            }
                            v14 = 833
                        }
                    }
                }
            }
        }
        if (true && (-33 < 0)) {
            if (!(!v59)) {
                switch fi61 {
                case 498:
                    var v72 string = "\n\t"
                    _ = v72
                case 4294967295:
                    if ((v15 / ((2147483647) | 1)) <= v15) {
                        v14 -= 663
                    }
                }
                if ((((((0.1 * 1e10) / (((0.0 + float32(float64(-2.25))) - ((1e10 + -2.25) / (1e10 - (((3.14159 / (float32(0.1) + (0.0 * -2.25))) * 1e10) - ((3.14159 / 1.0) - 0.1))))) * 0.1)) * (3.14159 - ((1.5 / 1e10) + ((0.1 + ((-2.25 - (-2.25 + 3.14159)) / 1.0)) + (0.0 * 0.0))))) - (-2.25 - 1e10)) / ((0.1 - ((1.5 - 1.5) - -2.25)) / 3.14159)) < 3.14159) {
                    v73 := v6[21]
                    v4 = ((!v73) || (!((v1 / 1.0) != (v1 + v1))))
                    if false {
                        v3 = append(v3, v60)
                    }
                    _ = v73
                }
                v59 = true
            } else {
                v4 = (!((v4 || true) && (!("a" < v60))))
                v74 := GenS8{F0: 261, F1: -128, F2: 372}
                _ = v74
            }
            v75 := GenS9{F0: 32767}
            _ = v75
        }
        v5 = append(v5, v15)
        _ = fi61
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v59:", v59)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v60:", v60)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v59
    _ = v6
    _ = v60
}
