package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 uint32; F1 uint16 }
type GenS4 struct { F0 int16; F1 int8; F2 int8 }
type GenS5 struct { F0 int64; F1 int8; F2 uint8 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int64; F1 int64; F2 uint32 }
type GenS8 struct { F0 int64; F1 uint16; F2 uint8 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 bool }

func main() {
    var v1 float64 = ((1e10 / 1.0) * float64(0.1))
    var v2 float64 = float64(1e10)
    v3 := make(map[uint32]bool)
    var v4 uint = 18446744073709551615
    v5 := []string{("foo bar" + "a")}
    v6 := make(map[int]int)
    if false {
        fmt.Println(v1, v4, v4)
        v1 -= v2
    }
    v7 := GenS1{F0: (^(v4))}
    var v8 bool = (0.1 != 0.0)
    var v9 string = "foo bar"
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v10 := GenS2{F0: 135, F1: 18446744073709551615}
    var v11 int64 = 0
    v1 += 1e10
    if (32767 >= 2) {
        v12 := GenS1{F0: (^(v4))}
        v11 = 0
        var v13 float64 = v2
        _ = v12
        _ = v13
    } else {
        var v14 float32 = (float32(3.14159) + -2.25)
        _ = v14
    }
    v2 -= 0.1
    if v8 {
        var v15 int64 = v11
        switch (((v9 + "a") + "a") + v9) {
        case "":
            if (!(int32(v15) >= 1)) {
                if (!true) {
                    switch 706 {
                    case -9223372036854775808:
                        v16 := GenS1{F0: (v4 & 2)}
                        _ = v16
                    case 32768:
                        for fi17 := 0; fi17 < 19; fi17++ {
                            for fi18 := 0; fi18 < 14; fi18++ {
                                _ = len(v5)
                                if true {
                                    v2 = v2
                                    var v19 uint64 = 1
                                    var v20 string = "\n\t"
                                    _ = v19
                                    _ = v20
                                }
                                _ = fi18
                            }
                            if (((355 != 65535) || v8) || (v8 && true)) {
                                v21 := GenS3{F0: 4294967295, F1: 255}
                                if false {
                                    switch "a" {
                                    case "hello":
                                        v4 = (v4 & (^(v4)))
                                    case "":
                                        if true {
                                            if false {
                                                v22 := GenS4{F0: -661, F1: 100, F2: 42}
                                                var v23 uint8 = 1
                                                _ = v22
                                                _ = v23
                                            }
                                            var v24 bool = (109 < 128)
                                            _ = v24
                                        } else {
                                            switch fi17 {
                                            case -9223372036854775808:
                                                var v25 int32 = 2147483647
                                                _ = v25
                                            case 43:
                                                if v8 {
                                                    v9 = v9
                                                    if (0 == 0) {
                                                        v26 := GenS2{F0: 234, F1: 119}
                                                        v27 := GenS2{F0: 42, F1: 77}
                                                        if (v15 != v11) {
                                                            v8 = v8
                                                            if v8 {
                                                                v28 := GenS5{F0: v15, F1: 103, F2: 2}
                                                                v5 = append(v5, ("a" + (v9 + ((v9 + ("foo bar" + v9)) + ((v9 + "\n\t") + v9)))))
                                                                var v29 uint32 = 2
                                                                _ = v28
                                                                _ = v29
                                                            }
                                                        }
                                                        _ = v26
                                                        _ = v27
                                                    }
                                                }
                                            case 128:
                                                for fi30 := 0; fi30 < 1; fi30++ {
                                                    if (v8 || v8) {
                                                        switch (fi17 * fi17) {
                                                        case 127:
                                                            v31 := GenS6{F0: 128}
                                                            _ = v31
                                                        case -500:
                                                            v9 = v9
                                                        }
                                                    } else {
                                                        v9 = "a"
                                                        var v32 bool = (!true)
                                                        _ = v32
                                                    }
                                                    if v8 {
                                                        var v33 string = v9
                                                        v1 -= (v2 / v2)
                                                        var v34 int16 = 1000
                                                        var v35 uint32 = 2147483648
                                                        if true {
                                                            switch "" {
                                                            case "foo bar":
                                                                fmt.Println(v15, v11)
                                                            case "":
                                                                v9 = v33
                                                            case "\n\t":
                                                                fmt.Println(v35, v34)
                                                            default:
                                                                switch (fi17 << (((v4 / ((v4) | 1))) & 63)) {
                                                                case 32767:
                                                                    v36 := GenS4{F0: v34, F1: 1, F2: 127}
                                                                    _ = v36
                                                                default:
                                                                    v4 += 65535
                                                                }
                                                            }
                                                            v15 -= (v11 * v15)
                                                        } else {
                                                            delete(v3, v35)
                                                            switch v9 {
                                                            case "a":
                                                                var v37 int32 = -2147483648
                                                                _ = v37
                                                            case "\n\t":
                                                                v5 = append(v5, "")
                                                            case "":
                                                                if (!(!v8)) {
                                                                    v34 = (v34 | (v34 + (-(v34))))
                                                                } else {
                                                                    switch v33 {
                                                                    case "hello":
                                                                        var v38 bool = (((v2 / float64(3.14159)) + 0.0) > v2)
                                                                        _ = v38
                                                                    case "foo bar":
                                                                        v4 = (v4 & 1)
                                                                    case "":
                                                                        var v39 int = 2147483647
                                                                        _ = v39
                                                                    }
                                                                    v40 := GenS4{F0: (v34 >> ((v4) & 63)), F1: 0, F2: 100}
                                                                    _ = v40
                                                                }
                                                            }
                                                            v6[(fi30 / ((fi17) | 1))] = (fi30 ^ 9223372036854775807)
                                                        }
                                                        var v41 uint16 = 774
                                                        var v42 int8 = -128
                                                        v2 -= v1
                                                        _ = v33
                                                        _ = v34
                                                        _ = v35
                                                        _ = v41
                                                        _ = v42
                                                    }
                                                    _ = fi30
                                                }
                                            }
                                        }
                                    default:
                                        v43 := GenS7{F0: -1, F1: v11, F2: 42}
                                        _ = v43
                                    }
                                    for fi44 := 0; fi44 < 17; fi44++ {
                                        v45 := GenS4{F0: 924, F1: 127, F2: -23}
                                        switch fi44 {
                                        case 0:
                                            v11 += (v11 * (^(v11)))
                                        case 9223372036854775807:
                                            _ = len(v5)
                                        default:
                                            fmt.Println(v4, fi17, v2)
                                        }
                                        _ = fi44
                                        _ = v45
                                    }
                                }
                                _ = v21
                            }
                            for fi46 := 0; fi46 < 3; fi46++ {
                                var v47 uint64 = 18446744073709551615
                                fmt.Println(fi46)
                                _ = fi46
                                _ = v47
                            }
                            _ = fi17
                        }
                    case 894:
                        v5 = append(v5, v9)
                    default:
                        v5 = append(v5, v9)
                    }
                } else {
                    v48 := GenS7{F0: v11, F1: v15, F2: 65535}
                    for fi49 := 0; fi49 < 3; fi49++ {
                        v15 = int64(v4)
                        _ = fi49
                    }
                    switch 122 {
                    case -1:
                        var v50 uint64 = 0
                        _ = v50
                    case 9223372036854775807:
                        v11 = v11
                    case 255:
                        var v51 uint16 = 70
                        _ = v51
                    default:
                        for fi52 := 0; fi52 < 13; fi52++ {
                            for fi53 := 0; fi53 < 20; fi53++ {
                                if (false || (false && v8)) {
                                    var v54 bool = v8
                                    switch v9 {
                                    case "hello":
                                        var v55 int8 = -128
                                        _ = v55
                                    case "a":
                                        v54 = v54
                                        delete(v3, 2147483648)
                                    case "":
                                        var v56 float32 = float32(float64((-2.25 + 1e10)))
                                        var v57 uint8 = 2
                                        _ = v56
                                        _ = v57
                                    default:
                                        v54 = true
                                    }
                                    _ = v54
                                }
                                var v58 uint8 = 127
                                _ = fi53
                                _ = v58
                            }
                            _ = fi52
                        }
                    }
                    _ = v48
                }
                v11 = 42
            } else {
                v15 = (v11 * 0)
                for fi59 := 0; fi59 < 4; fi59++ {
                    v60 := GenS8{F0: v15, F1: 42, F2: 0}
                    v61 := GenS5{F0: (v11 % (((v15 >> (((v4 % (((v4 | 4294967295)) | 1))) & 63))) | 1)), F1: -128, F2: 42}
                    var v62 bool = v8
                    _ = fi59
                    _ = v60
                    _ = v61
                    _ = v62
                }
                v15 += v11
            }
        case "foo bar":
            for fi63 := 0; fi63 < 3; fi63++ {
                for fi64 := 0; fi64 < 14; fi64++ {
                    v65 := GenS1{F0: 18446744073709551615}
                    v4 -= (^(v4))
                    var v66 int32 = 0
                    v3[2147483647] = false
                    _ = fi64
                    _ = v65
                    _ = v66
                }
                switch "a" {
                case "a":
                    if v8 {
                        v4 = v4
                        var v67 int32 = -2147483648
                        _ = v67
                    }
                case "\n\t":
                    for fi68 := 0; fi68 < 4; fi68++ {
                        v5 = append(v5, "\n\t")
                        var v69 int16 = 1000
                        _ = fi68
                        _ = v69
                    }
                default:
                    if v8 {
                        switch 32768 {
                        case 1:
                            fmt.Println(v9, v11, v2)
                        }
                        for fi70 := 0; fi70 < 14; fi70++ {
                            v1 = v1
                            _ = fi70
                        }
                    } else {
                        var v71 uint32 = 349
                        var v72 uint16 = 32767
                        fmt.Println(v11, v9)
                        _ = v71
                        _ = v72
                    }
                }
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = fi63
            }
        case "a":
            v9 = (("\n\t" + "hello") + "a")
        }
        v4 += 1
        _ = v15
    } else {
        v2 -= v2
        var v73 bool = true
        v74 := GenS9{F0: -194}
        _ = v73
        _ = v74
    }
    v2 = v2
    for fi75 := 0; fi75 < 2; fi75++ {
        for fi76 := 0; fi76 < 2; fi76++ {
            var v77 uint16 = 413
            v78 := v6[(fi75 % (((fi76 / ((1) | 1))) | 1))]
            v78 = v78
            _ = fi76
            _ = v77
            _ = v78
        }
        v1 -= ((float64((1e10 / -2.25)) + v2) * v2)
        _ = fi75
    }
    if ((v11 > 42) || true) {
        v79 := v3[15]
        if (208 == 255) {
            switch ("\n\t" + ("" + "a")) {
            case "a":
                v80 := GenS3{F0: 1, F1: 255}
                _ = v80
            case "hello":
                if (32767 != 32767) {
                    v5 = append(v5, ("\n\t" + (v9 + "a")))
                    for fi81 := 0; fi81 < 12; fi81++ {
                        v82 := GenS10{F0: 1}
                        if (0 == 502) {
                            fmt.Println(v11)
                            for fi83 := 0; fi83 < 8; fi83++ {
                                var v84 int16 = -1
                                var v85 int8 = 100
                                _ = fi83
                                _ = v84
                                _ = v85
                            }
                            for fi86 := 0; fi86 < 9; fi86++ {
                                v87 := GenS11{F0: true}
                                v2 = (-2.25 * 1.5)
                                _ = fi86
                                _ = v87
                            }
                        }
                        _ = fi81
                        _ = v82
                    }
                    v79 = true
                }
            }
        }
        var v88 int = 32768
        _ = v79
        _ = v88
    } else {
        if v8 {
            fmt.Println(v1)
            v89 := GenS11{F0: (-90 == 1)}
            fmt.Println(v11, v2)
            _ = v89
        } else {
            fmt.Println(v4, v1, v11)
            v1 += float64(3.14159)
        }
        v90 := GenS5{F0: v11, F1: 100, F2: 128}
        _ = v90
    }
    var v91 int16 = 42
    v4 = 2
    var v92 string = (("a" + (v9 + "\n\t")) + ((v9 + (v9 + "\n\t")) + (("foo bar" + "hello") + v9)))
    for fi93 := 0; fi93 < 14; fi93++ {
        if true {
            v94 := GenS4{F0: v91, F1: 0, F2: 97}
            var v95 uint = v4
            _ = v94
            _ = v95
        }
        v92 = v92
        _ = fi93
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    fmt.Println("v91:", v91)
    fmt.Println("v92:", v92)
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
    _ = v91
    _ = v92
}
