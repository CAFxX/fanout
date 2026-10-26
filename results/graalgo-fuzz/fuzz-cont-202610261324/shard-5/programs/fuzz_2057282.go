package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 string; F1 uint8; F2 uint32 }
type GenS5 struct { F0 bool; F1 int8 }
type GenS6 struct { F0 int8; F1 uint; F2 uint8 }
type GenS7 struct { F0 uint; F1 int64; F2 uint16 }
type GenS8 struct { F0 int16; F1 uint; F2 uint16 }
type GenS9 struct { F0 bool; F1 int; F2 uint8 }
type GenS10 struct { F0 int32; F1 int }

func main() {
    var v1 int32 = 1
    v2 := make(map[uint32]uint32)
    v3 := []string{("" + "hello"), "foo bar"}
    var v4 bool = (!true)
    v5 := []int{}
    v6 := make(map[string]int)
    v7 := GenS1{F0: (!v4)}
    var v8 string = ""
    var v9 uint16 = 2
    v8 = (v8 + v8)
    if false {
        if false {
            for fi10 := 0; fi10 < 5; fi10++ {
                v11 := GenS1{F0: (!v4)}
                if ((fi10 | (fi10 ^ 32767)) >= (fi10 >> ((255) & 63))) {
                    for fi12 := 0; fi12 < 19; fi12++ {
                        var v13 int32 = -2147483648
                        if (-7 <= -128) {
                            fmt.Println(v1, v1, v1)
                        } else {
                            var v14 uint16 = (v9 - (v9 + (v9 + v9)))
                            switch ("foo bar" + (("" + v8) + ("" + v8))) {
                            case "foo bar":
                                if (true && (!(v4 || v4))) {
                                    v6[v8] = fi12
                                } else {
                                    if (0.1 < 1.5) {
                                        if false {
                                            var v15 uint64 = uint64(v9)
                                            v14 = (v14 + v9)
                                            _ = v15
                                        } else {
                                            switch 2147483647 {
                                            case -1:
                                                if ((!false) || (v1 >= (v1 + v13))) {
                                                    var v16 string = (("" + ("\n\t" + (v8 + (("\n\t" + v8) + v8)))) + "foo bar")
                                                    switch (v8 + v8) {
                                                    case "\n\t":
                                                        fmt.Println(v8)
                                                    case "foo bar":
                                                        v17 := GenS2{F0: 158}
                                                        _ = v17
                                                    case "hello":
                                                        v1 = -69
                                                    }
                                                    v2[0] = 1
                                                    _ = v16
                                                }
                                            case -2147483648:
                                                fmt.Println(v8, v1)
                                            case 0:
                                                v18 := GenS3{F0: 255}
                                                _ = v18
                                            default:
                                                var v19 uint = 0
                                                _ = v19
                                            }
                                            var v20 int8 = 127
                                            var v21 uint = 42
                                            _ = v20
                                            _ = v21
                                        }
                                    }
                                }
                            case "a":
                                var v22 uint16 = v14
                                v23 := GenS4{F0: (v8 + v8), F1: 1, F2: 2}
                                if (v4 || true) {
                                    var v24 uint8 = 2
                                    v25 := GenS2{F0: v24}
                                    _ = v24
                                    _ = v25
                                } else {
                                    v26 := GenS4{F0: v8, F1: 0, F2: 65535}
                                    _ = len(v3)
                                    var v27 int32 = (v13 - v1)
                                    if (!true) {
                                        fmt.Println(v9)
                                        var v28 int32 = -2147483648
                                        _ = v28
                                    } else {
                                        v29 := GenS5{F0: v4, F1: 127}
                                        switch (v8 + ("a" + (((v8 + "hello") + (v8 + v8)) + (v8 + (v8 + v8))))) {
                                        case "hello":
                                            v9 = v9
                                        case "\n\t":
                                            if (!false) {
                                                fmt.Println(v14)
                                                switch fi10 {
                                                case 4294967295:
                                                    v9 -= v9
                                                    v1 = 657
                                                    if (1 <= 4294967295) {
                                                        v4 = (true || v4)
                                                        v4 = false
                                                        v30 := GenS2{F0: 2}
                                                        v2[160] = 0
                                                        v1 = v1
                                                        _ = v30
                                                    }
                                                default:
                                                    switch (fi12 >> ((uint(v14)) & 63)) {
                                                    case 128:
                                                        v4 = (796 != 0)
                                                    }
                                                }
                                            } else {
                                                var v31 int16 = -32768
                                                delete(v6, v8)
                                                _ = v31
                                            }
                                        case "foo bar":
                                            v9 = (v22 % ((v14) | 1))
                                            var v32 int8 = -128
                                            v33 := v6[v8]
                                            _ = v32
                                            _ = v33
                                        default:
                                            var v34 int16 = -32768
                                            _ = v34
                                        }
                                        _ = v29
                                    }
                                    var v35 string = (v8 + "hello")
                                    _ = v26
                                    _ = v27
                                    _ = v35
                                }
                                _ = v22
                                _ = v23
                            case "":
                                var v36 uint16 = (v9 ^ (v14 / (((v14 >> ((651) & 63))) | 1)))
                                _ = v36
                            }
                            v8 = "hello"
                            _ = v14
                        }
                        _ = fi12
                        _ = v13
                    }
                    fmt.Println(v1, v8)
                    if v4 {
                        v4 = (!v4)
                        _ = len(v6)
                        v4 = (1000 > 0)
                    }
                } else {
                    for fi37 := 0; fi37 < 11; fi37++ {
                        switch v8 {
                        case "\n\t":
                            var v38 float64 = float64(((0.0 / (0.1 + 0.0)) / 0.1))
                            _ = v38
                        default:
                            switch ((((v8 + "a") + "\n\t") + (((v8 + "a") + "a") + (((v8 + "foo bar") + v8) + v8))) + (v8 + v8)) {
                            case "hello":
                                var v39 float64 = ((3.14159 / 1e10) + 1e10)
                                _ = v39
                            case "a":
                                var v40 uint = 255
                                _ = v40
                            case "\n\t":
                                var v41 int32 = v1
                                _ = v41
                            default:
                                if ((uint32(v9) == 56) && (109 > 1000)) {
                                    v4 = v4
                                } else {
                                    fmt.Println(v8, v4)
                                    delete(v6, (("" + (v8 + (("foo bar" + ((v8 + v8) + v8)) + v8))) + "hello"))
                                    var v42 uint16 = 65535
                                    _ = v42
                                }
                            }
                        }
                        _ = fi37
                    }
                    for fi43 := 0; fi43 < 19; fi43++ {
                        _ = len(v3)
                        var v44 bool = ((((0.1 * (1.5 * 0.1)) * 3.14159) / (1.5 * -2.25)) <= 1e10)
                        v45 := GenS1{F0: v44}
                        _ = fi43
                        _ = v44
                        _ = v45
                    }
                    v46 := GenS3{F0: 477}
                    _ = v46
                }
                _ = fi10
                _ = v11
            }
            if true {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                for fi47 := 0; fi47 < 9; fi47++ {
                    if true {
                        v48 := v2[42]
                        if (2 != 255) {
                            var v49 string = v8
                            for fi50 := 0; fi50 < 5; fi50++ {
                                var v51 string = v8
                                v1 = 2147483647
                                var v52 int32 = (v1 * 0)
                                v3 = append(v3, ("a" + v49))
                                _ = fi50
                                _ = v51
                                _ = v52
                            }
                            v4 = (("foo bar" + ((("\n\t" + "") + v49) + "foo bar")) >= v8)
                            _ = v49
                        }
                        _ = v48
                    } else {
                        if ((fi47 * 584) > fi47) {
                            if v4 {
                                var v53 int64 = 9223372036854775807
                                v54 := v6[v8]
                                _ = v53
                                _ = v54
                            } else {
                                for fi55 := 0; fi55 < 15; fi55++ {
                                    fmt.Println(fi47, v1, fi55)
                                    _ = fi55
                                }
                            }
                            var v56 int64 = -1
                            _ = v56
                        } else {
                            var v57 bool = v4
                            v9 = 0
                            _ = v57
                        }
                        fmt.Println(v9, v1)
                    }
                    delete(v6, "a")
                    if len(v3) > 0 {
                        v3[0] = ""
                    }
                    _ = fi47
                }
            }
            for fi58 := 0; fi58 < 8; fi58++ {
                v59 := GenS6{F0: 117, F1: 547, F2: 127}
                for fi60 := 0; fi60 < 16; fi60++ {
                    v61 := GenS6{F0: -30, F1: 4294967295, F2: 255}
                    _ = fi60
                    _ = v61
                }
                _ = fi58
                _ = v59
            }
        }
        v62 := GenS5{F0: v4, F1: -109}
        _ = v62
    }
    if (5 < uint(v9)) {
        if (v9 != (v9 | (v9 * 42))) {
            switch v8 {
            case "a":
                fmt.Println(v9, v1)
            case "foo bar":
                v8 = (v8 + v8)
            }
            fmt.Println(v1)
            fmt.Println(v9, v4)
        }
        for fi63 := 0; fi63 < 16; fi63++ {
            v4 = (219 <= 1)
            switch (fi63 * (^(fi63))) {
            case 280:
                v64 := GenS6{F0: 42, F1: 42, F2: 42}
                _ = v64
            case -1:
                var v65 float32 = (3.14159 / -2.25)
                _ = v65
            default:
                v66 := v2[4294967295]
                _ = v66
            }
            _ = fi63
        }
    } else {
        v9 = (v9 << ((65535) & 63))
        for fi67 := 0; fi67 < 5; fi67++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            fmt.Println(fi67, v4)
            v68 := v2[65535]
            _ = fi67
            _ = v68
        }
        for fi69 := 0; fi69 < 9; fi69++ {
            delete(v6, v8)
            switch v8 {
            case "\n\t":
                v1 += v1
            case "":
                v70 := GenS1{F0: (true || (int64(v1) >= -632))}
                _ = v70
            default:
                for fi71 := 0; fi71 < 11; fi71++ {
                    if v4 {
                        if true {
                            v72 := v2[42]
                            if v4 {
                                var v73 uint64 = 0
                                if ("a" <= "foo bar") {
                                    if false {
                                        v74 := GenS3{F0: 791}
                                        if ((!(-1 == 1000)) || v4) {
                                            if v4 {
                                                v3 = append(v3, v8)
                                                switch (fi69 * 128) {
                                                case 9223372036854775807:
                                                    fmt.Println(fi69)
                                                }
                                                var v75 uint = 926
                                                _ = v75
                                            }
                                            fmt.Println(v73, v9, fi69)
                                            switch (fi69 >> ((2) & 63)) {
                                            case 4294967295:
                                                var v76 uint64 = (v73 % (((^(v73))) | 1))
                                                _ = v76
                                            case 128:
                                                if len(v3) > 0 {
                                                    v3[0] = v8
                                                }
                                            case 1:
                                                if len(v3) > 0 {
                                                    v3[0] = ""
                                                }
                                            }
                                        } else {
                                            if (1.5 >= 1e10) {
                                                v4 = (false || ("\n\t" < v8))
                                                if v4 {
                                                    v8 = "\n\t"
                                                }
                                            }
                                            var v77 int64 = 42
                                            _ = v77
                                        }
                                        _ = v74
                                    } else {
                                        var v78 uint64 = 0
                                        _ = v78
                                    }
                                    var v79 uint = 255
                                    _ = v79
                                }
                                _ = v73
                            } else {
                                v8 = v8
                                v4 = (!(87 < -9223372036854775808))
                                v80 := v2[(v72 - v72)]
                                _ = v80
                            }
                            _ = v72
                        } else {
                            var v81 int = (fi71 >> ((65535) & 63))
                            _ = v81
                        }
                        if (float32(float64(3.14159)) < float32(0.1)) {
                            fmt.Println(v8, v8, fi71)
                            v82 := GenS6{F0: -116, F1: 42, F2: 73}
                            switch v8 {
                            case "":
                                v83 := v2[333]
                                _ = v83
                            case "foo bar":
                                if false {
                                    v3 = append(v3, (("a" + (("a" + v8) + v8)) + "a"))
                                    if (-128 < 16) {
                                        v84 := v6[(v8 + ("" + ((v8 + v8) + (v8 + "foo bar"))))]
                                        v85 := GenS7{F0: 65535, F1: -649, F2: (^(v9))}
                                        _ = v84
                                        _ = v85
                                    } else {
                                        fmt.Println(fi71, v9, fi69)
                                        v3 = append(v3, ((v8 + v8) + v8))
                                    }
                                    v86 := GenS6{F0: 38, F1: 394, F2: 42}
                                    var v87 float64 = (1e10 + (((1e10 - ((1e10 + 1e10) + 0.0)) / -2.25) - 3.14159))
                                    _ = v86
                                    _ = v87
                                }
                            case "a":
                                delete(v2, 1)
                            }
                            _ = v82
                        }
                        switch (fi71 / ((-605) | 1)) {
                        case 128:
                            if (((1.5 + 3.14159) * (0.1 - 3.14159)) != 0.1) {
                                v8 = ("a" + v8)
                                switch (fi71 | (-(fi71))) {
                                case -1:
                                    v1 -= (v1 / (((v1 % ((v1) | 1))) | 1))
                                default:
                                    var v88 int64 = -1
                                    _ = v88
                                }
                                v89 := GenS1{F0: ("hello" >= "foo bar")}
                                _ = v89
                            }
                        case 127:
                            v90 := GenS8{F0: -32768, F1: 18446744073709551615, F2: v9}
                            _ = v90
                        }
                    } else {
                        fmt.Println(fi69, fi71, v4)
                        v8 = ((("a" + v8) + "hello") + (((v8 + v8) + v8) + "a"))
                        switch fi71 {
                        case -312:
                            var v91 int32 = v1
                            _ = v91
                        }
                    }
                    v92 := GenS9{F0: ((1e10 < 1e10) && (-2.25 >= 3.14159)), F1: 32768, F2: 128}
                    if (2147483647 != 42) {
                        if len(v3) > 0 {
                            v3[0] = v8
                        }
                        v93 := GenS1{F0: false}
                        _ = v93
                    } else {
                        v94 := GenS2{F0: 148}
                        _ = v94
                    }
                    if (!true) {
                        v95 := GenS3{F0: 4294967295}
                        if (!false) {
                            if v4 {
                                v96 := GenS10{F0: 1, F1: fi69}
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = v96
                            }
                        }
                        _ = v95
                    }
                    var v97 int8 = 1
                    _ = fi71
                    _ = v92
                    _ = v97
                }
            }
            v8 = "foo bar"
            _ = fi69
        }
    }
    if len(v5) > 0 {
        v5[0] = 4294967295
    }
    var v98 uint16 = v9
    v2[1] = 1
    var v99 int8 = -1
    delete(v2, 42)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    fmt.Println("v98:", v98)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    _ = v98
    _ = v99
}
