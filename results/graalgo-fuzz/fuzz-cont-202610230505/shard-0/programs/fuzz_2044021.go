package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint16; F1 uint16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int32; F1 int64; F2 string }
type GenS9 struct { F0 uint; F1 bool }

func main() {
    v1 := make(map[uint32]string)
    var v2 int8 = 1
    var v3 float32 = -2.25
    v4 := []string{"a"}
    v5 := []string{"foo bar"}
    var v6 uint64 = 305
    switch 128 {
    case 2147483647:
        switch "" {
        case "a":
            v3 -= v3
        case "hello":
            if false {
                switch ("hello" + "") {
                case "hello":
                    for fi7 := 0; fi7 < 4; fi7++ {
                        var v8 int8 = -1
                        delete(v1, 318)
                        _ = fi7
                        _ = v8
                    }
                case "a":
                    fmt.Println(v2, v2, v6)
                default:
                    var v9 uint32 = 4294967295
                    _ = v9
                }
                switch "hello" {
                case "foo bar":
                    v10 := GenS1{F0: 24}
                    _ = v10
                case "a":
                    var v11 string = "foo bar"
                    _ = v11
                }
                v3 = v3
            } else {
                if false {
                    if true {
                        v1[4294967295] = ("" + ("\n\t" + ""))
                        fmt.Println(v6, v6, v3)
                        v2 = (v2 & (v2 % ((v2) | 1)))
                    }
                    v12 := GenS1{F0: 107}
                    _ = v12
                } else {
                    for fi13 := 0; fi13 < 11; fi13++ {
                        if false {
                            v14 := GenS1{F0: 1}
                            _ = v14
                        } else {
                            v2 = v2
                        }
                        _ = fi13
                    }
                    switch 4294967295 {
                    case -371:
                        v2 = (v2 + v2)
                    case -41:
                        v2 -= v2
                    case 0:
                        v15 := GenS1{F0: 2}
                        _ = v15
                    }
                    _ = len(v1)
                }
                _ = len(v1)
                var v16 string = "\n\t"
                _ = v16
            }
        }
    case -309:
        fmt.Println(v3, v6, v2)
    }
    v17 := GenS2{F0: v6}
    v18 := GenS3{F0: 42, F1: 2}
    if false {
        v19 := GenS1{F0: 1}
        v2 = (v2 - 100)
        _ = v19
    } else {
        v20 := GenS1{F0: 42}
        _ = v20
    }
    var v21 uint8 = 1
    v22 := v1[113]
    v6 += 2
    for fi23 := 0; fi23 < 13; fi23++ {
        fmt.Println(v6, v2)
        v24 := GenS4{F0: v22}
        _ = fi23
        _ = v24
    }
    if (v22 > "hello") {
        fmt.Println(v22, v6)
    } else {
        if true {
            var v25 int32 = 0
            _ = v25
        } else {
            for fi26 := 0; fi26 < 12; fi26++ {
                var v27 int32 = 42
                v6 += v6
                var v28 int16 = -272
                _ = fi26
                _ = v27
                _ = v28
            }
            for fi29 := 0; fi29 < 18; fi29++ {
                for fi30 := 0; fi30 < 8; fi30++ {
                    switch 626 {
                    case 32768:
                        switch ((v22 + v22) + (v22 + v22)) {
                        case "hello":
                            var v31 uint64 = v6
                            _ = v31
                        case "\n\t":
                            v32 := GenS5{F0: fi29}
                            _ = v32
                        case "":
                            if ((0 < 2) || (!true)) {
                                v22 = "foo bar"
                                _ = len(v1)
                                switch v22 {
                                case "foo bar":
                                    v33 := v1[65535]
                                    _ = v33
                                case "hello":
                                    v2 -= 1
                                case "\n\t":
                                    v34 := GenS4{F0: (((((v22 + "\n\t") + (v22 + "hello")) + v22) + v22) + v22)}
                                    _ = v34
                                }
                            } else {
                                var v35 int64 = 9223372036854775807
                                v35 = (v35 % ((v35) | 1))
                                _ = v35
                            }
                        default:
                            if (false || false) {
                                fmt.Println(fi30)
                                if false {
                                    if false {
                                        var v36 int64 = -1
                                        _ = v36
                                    } else {
                                        switch fi29 {
                                        case -703:
                                            v6 -= uint64(v21)
                                        default:
                                            delete(v1, 2)
                                        }
                                        var v37 string = v22
                                        v37 = "hello"
                                        v3 -= float32(float64(v3))
                                        _ = v37
                                    }
                                    var v38 uint = 1
                                    _ = v38
                                }
                                v39 := GenS6{F0: 1}
                                _ = len(v1)
                                v3 += 0.1
                                v22 = v22
                                var v40 int = (fi30 % ((int(v21)) | 1))
                                if ((!(42 < uint32(v6))) && (!true)) {
                                    var v41 string = "\n\t"
                                    _ = v41
                                } else {
                                    v42 := GenS6{F0: 1}
                                    _ = v42
                                }
                                switch v22 {
                                case "a":
                                    var v43 uint = 18446744073709551615
                                    _ = v43
                                case "hello":
                                    v44 := v1[42]
                                    _ = v44
                                }
                                if false {
                                    switch (^(v40)) {
                                    case 2147483647:
                                        v45 := GenS1{F0: 234}
                                        _ = v45
                                    default:
                                        _ = len(v4)
                                    }
                                    switch v22 {
                                    case "\n\t":
                                        switch v22 {
                                        case "hello":
                                            var v46 float64 = 0.0
                                            _ = v46
                                        case "foo bar":
                                            if len(v5) > 0 {
                                                v5[0] = ""
                                            }
                                        default:
                                            v47 := GenS6{F0: 1}
                                            v48 := GenS4{F0: v22}
                                            _ = v47
                                            _ = v48
                                        }
                                        var v49 uint = 0
                                        var v50 uint8 = (v21 % ((0) | 1))
                                        if (!(!(v3 > float32(3.14159)))) {
                                            delete(v1, 2147483647)
                                        } else {
                                            v4 = append(v4, v22)
                                            v2 -= -128
                                            v22 = ((v22 + v22) + "")
                                        }
                                        switch ("\n\t" + "hello") {
                                        case "":
                                            switch v40 {
                                            case 127:
                                                v40 = -2147483648
                                            default:
                                                if true {
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    v51 := GenS7{F0: v50}
                                                    v52 := GenS2{F0: (^(v6))}
                                                    _ = v51
                                                    _ = v52
                                                }
                                            }
                                        case "a":
                                            if (!(-9223372036854775808 < 0)) {
                                                delete(v1, 705)
                                                if (((v2 | 42) == v2) || ((v2 ^ v2) <= (v2 & (v2 + (v2 * -128))))) {
                                                    if (!false) {
                                                        if true {
                                                            if false {
                                                                if false {
                                                                    var v53 uint32 = 65535
                                                                    v53 += v53
                                                                    _ = v53
                                                                } else {
                                                                    _ = len(v4)
                                                                    v54 := GenS8{F0: 0, F1: 709, F2: v22}
                                                                    _ = v54
                                                                }
                                                                v21 = (v21 ^ v50)
                                                            }
                                                        }
                                                    } else {
                                                        v55 := v1[1]
                                                        var v56 float32 = (v3 + ((3.14159 / v3) / v3))
                                                        _ = v55
                                                        _ = v56
                                                    }
                                                    v22 = v22
                                                }
                                                v4 = append(v4, v22)
                                            }
                                        case "hello":
                                            _ = len(v1)
                                        }
                                        _ = v49
                                        _ = v50
                                    case "a":
                                        if true {
                                            if (1 <= 9223372036854775807) {
                                                _ = len(v1)
                                                var v57 bool = (4294967295 < 255)
                                                _ = v57
                                            } else {
                                                v58 := GenS1{F0: (v21 & (v21 % ((v21) | 1)))}
                                                _ = v58
                                            }
                                            if (true || ((v2 & (-(v2))) == (^(v2)))) {
                                                fmt.Println(fi29, v3, v21)
                                            } else {
                                                switch (fi29 >> ((65535) & 63)) {
                                                case -98:
                                                    v59 := GenS9{F0: 1, F1: (false || (427 == 510))}
                                                    _ = v59
                                                case 2147483647:
                                                    if true {
                                                        switch fi29 {
                                                        case 1:
                                                            var v60 uint16 = 0
                                                            switch v22 {
                                                            case "a":
                                                                v40 -= (fi30 % ((9223372036854775807) | 1))
                                                            case "foo bar":
                                                                v21 -= v21
                                                            }
                                                            fmt.Println(v40)
                                                            _ = v60
                                                        }
                                                        var v61 float64 = -2.25
                                                        _ = v61
                                                    }
                                                    v21 -= (v21 % ((v21) | 1))
                                                    var v62 int8 = (v2 + v2)
                                                    if (((0 <= uint(v6)) || false) || (-918 > -7)) {
                                                        v4 = append(v4, "\n\t")
                                                    }
                                                    var v63 int8 = -104
                                                    fmt.Println(fi29, v3, v62)
                                                    v5 = append(v5, (("foo bar" + v22) + v22))
                                                    _ = len(v1)
                                                    _ = v62
                                                    _ = v63
                                                default:
                                                    fmt.Println(v2, v2, v2)
                                                }
                                                switch "" {
                                                case "a":
                                                    v6 = (v6 | (v6 % (((v6 | v6)) | 1)))
                                                }
                                                if false {
                                                    v2 += (v2 % ((42) | 1))
                                                    _ = len(v5)
                                                }
                                            }
                                        } else {
                                            v40 = (v40 ^ -2147483648)
                                            switch v22 {
                                            case "a":
                                                _ = len(v4)
                                            default:
                                                v6 -= (v6 & v6)
                                            }
                                            fmt.Println(v3)
                                        }
                                    }
                                    switch fi29 {
                                    case 0:
                                        if (!(4294967295 <= 42)) {
                                            var v64 int16 = 1
                                            var v65 float64 = (((1e10 - float64(v3)) + 0.0) / 1.0)
                                            _ = v64
                                            _ = v65
                                        }
                                    case -1:
                                        if false {
                                            var v66 float32 = v3
                                            if true {
                                                v67 := GenS9{F0: 1, F1: (v40 == fi29)}
                                                _ = v67
                                            } else {
                                                var v68 float32 = 0.1
                                                var v69 string = v22
                                                _ = v68
                                                _ = v69
                                            }
                                            _ = v66
                                        }
                                    }
                                } else {
                                    fmt.Println(v21, v3)
                                    var v70 uint8 = (v21 & v21)
                                    _ = v70
                                }
                                switch ("" + v22) {
                                case "":
                                    var v71 uint32 = 4294967295
                                    _ = v71
                                case "foo bar":
                                    var v72 int32 = 42
                                    _ = v72
                                }
                                var v73 int16 = 209
                                _ = v39
                                _ = v40
                                _ = v73
                            } else {
                                switch (fi30 >> ((4294967295) & 63)) {
                                case 4294967295:
                                    fmt.Println(v21)
                                }
                                v74 := GenS7{F0: 219}
                                _ = v74
                            }
                        }
                        v6 -= 4294967295
                        var v75 int32 = -984
                        _ = v75
                    case -1:
                        v6 -= 0
                    }
                    _ = fi30
                }
                _ = fi29
            }
            var v76 uint16 = 2
            _ = v76
        }
        var v77 int8 = v2
        _ = v77
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
