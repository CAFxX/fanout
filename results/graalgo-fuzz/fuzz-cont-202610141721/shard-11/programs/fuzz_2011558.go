package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint32; F1 bool }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint8; F1 int8; F2 int16 }
type GenS8 struct { F0 uint; F1 int16; F2 int8 }
type GenS9 struct { F0 uint8; F1 uint16 }
type GenS10 struct { F0 int32; F1 int; F2 uint64 }
type GenS11 struct { F0 int; F1 string }

func main() {
    var v1 int64 = -9223372036854775808
    var v2 bool = false
    v3 := []string{}
    v4 := make(map[string]uint32)
    v5 := make(map[string]string)
    v6 := make(map[int]string)
    if (v2 && ("foo bar" < ("" + "\n\t"))) {
        switch "" {
        case "foo bar":
            v1 -= (^(v1))
        case "a":
            for fi7 := 0; fi7 < 3; fi7++ {
                v1 = v1
                v2 = false
                switch "foo bar" {
                case "\n\t":
                    var v8 uint32 = 4294967295
                    _ = v8
                default:
                    if len(v3) > 0 {
                        v3[0] = "hello"
                    }
                }
                _ = fi7
            }
        case "hello":
            var v9 float64 = float64(0.1)
            _ = v9
        }
        for fi10 := 0; fi10 < 13; fi10++ {
            if v2 {
                if v2 {
                    fmt.Println(fi10, v1)
                    var v11 string = "hello"
                    _ = v11
                }
                v3 = append(v3, "a")
                if (-154 <= 32767) {
                    if v2 {
                        switch fi10 {
                        case 802:
                            v2 = v2
                        default:
                            v1 = 0
                        }
                        switch (fi10 % ((-56) | 1)) {
                        case -503:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        case 2147483647:
                            if (false || (!true)) {
                                fmt.Println(fi10, v1)
                            }
                        case 151:
                            switch fi10 {
                            case 986:
                                v12 := GenS1{F0: -32768, F1: -2147483648}
                                _ = v12
                            case 4294967295:
                                var v13 uint16 = 1
                                _ = v13
                            default:
                                switch (fi10 - (-(fi10))) {
                                case -1:
                                    v14 := v6[0]
                                    _ = v14
                                case 32768:
                                    var v15 uint8 = 36
                                    _ = v15
                                case 17:
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                default:
                                    for fi16 := 0; fi16 < 9; fi16++ {
                                        delete(v6, (^(fi16)))
                                        if true {
                                            v2 = (2 != 1)
                                        } else {
                                            fmt.Println(v1, v2, v2)
                                            v1 = v1
                                            fmt.Println(v2, fi16, v1)
                                        }
                                        _ = len(v6)
                                        _ = fi16
                                    }
                                }
                            }
                        default:
                            v2 = ((true || true) || v2)
                        }
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    } else {
                        for fi17 := 0; fi17 < 15; fi17++ {
                            fmt.Println(v2, fi17)
                            if (9 >= 65535) {
                                if false {
                                    fmt.Println(v2)
                                    if len(v3) > 0 {
                                        v3[0] = (("hello" + ("foo bar" + "foo bar")) + "hello")
                                    }
                                    v18 := GenS1{F0: 0, F1: 2147483647}
                                    _ = v18
                                }
                                fmt.Println(v2, fi17, fi10)
                                var v19 int64 = v1
                                _ = v19
                            }
                            if v2 {
                                v20 := GenS2{F0: "a"}
                                v21 := GenS1{F0: -338, F1: -489}
                                if true {
                                    var v22 uint32 = 599
                                    v23 := GenS1{F0: 813, F1: -740}
                                    _ = v22
                                    _ = v23
                                }
                                _ = v20
                                _ = v21
                            } else {
                                v2 = v2
                                var v24 float32 = 1e10
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = v24
                            }
                            v2 = (1 != 255)
                            if ((!v2) && ((fi17 << ((0) & 63)) == fi10)) {
                                if v2 {
                                    _ = len(v3)
                                    v25 := GenS3{F0: (v1 % ((v1) | 1))}
                                    switch (fi10 + (fi17 / (((-(fi17))) | 1))) {
                                    case 128:
                                        fmt.Println(v1, v1, v2)
                                    case 2147483647:
                                        v3 = append(v3, "a")
                                    }
                                    _ = v25
                                } else {
                                    _ = len(v3)
                                }
                                _ = len(v5)
                            } else {
                                v2 = (v2 || v2)
                            }
                            if true {
                                if false {
                                    v26 := v4["hello"]
                                    v2 = (!v2)
                                    _ = v26
                                } else {
                                    var v27 uint16 = 65535
                                    var v28 int64 = v1
                                    _ = v27
                                    _ = v28
                                }
                                v1 = 9223372036854775807
                            }
                            if (false || v2) {
                                v2 = true
                                switch "foo bar" {
                                case "\n\t":
                                    var v29 uint16 = 0
                                    _ = v29
                                }
                                if true {
                                    fmt.Println(v1, v1)
                                } else {
                                    switch (-(fi10)) {
                                    case 32767:
                                        var v30 bool = true
                                        _ = v30
                                    }
                                    var v31 bool = (2 > 4294967295)
                                    _ = v31
                                }
                            } else {
                                v2 = (32767 == 237)
                                switch (fi10 % ((128) | 1)) {
                                case 4294967295:
                                    var v32 uint8 = 233
                                    _ = v32
                                case 993:
                                    v33 := v4["\n\t"]
                                    _ = v33
                                default:
                                    switch (fi17 + fi17) {
                                    case 1:
                                        var v34 float32 = (1.5 / 1.0)
                                        _ = v34
                                    case 0:
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        v2 = (65535 < 42)
                                    case -1:
                                        v35 := GenS4{F0: 65535, F1: (42 != -499)}
                                        _ = v35
                                    }
                                }
                            }
                            _ = fi17
                        }
                        v36 := GenS1{F0: 1, F1: -295}
                        v37 := GenS1{F0: 375, F1: 1}
                        _ = v36
                        _ = v37
                    }
                    v38 := GenS5{F0: (v1 << ((2) & 63))}
                    v39 := GenS6{F0: false}
                    _ = v38
                    _ = v39
                } else {
                    v1 = v1
                    var v40 uint64 = 0
                    switch "hello" {
                    case "\n\t":
                        var v41 uint = 1
                        _ = v41
                    case "a":
                        v42 := GenS1{F0: 690, F1: -2147483648}
                        _ = v42
                    case "foo bar":
                        for fi43 := 0; fi43 < 6; fi43++ {
                            var v44 float64 = -2.25
                            switch fi10 {
                            case -1:
                                var v45 float64 = v44
                                _ = v45
                            }
                            _ = fi43
                            _ = v44
                        }
                    default:
                        var v46 float32 = 0.1
                        _ = v46
                    }
                    _ = v40
                }
            }
            _ = fi10
        }
    }
    if false {
        var v47 int16 = 1000
        v48 := GenS7{F0: 1, F1: -128, F2: (-(v47))}
        var v49 int64 = 42
        _ = v47
        _ = v48
        _ = v49
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        if v2 {
            for fi50 := 0; fi50 < 6; fi50++ {
                var v51 uint16 = 65535
                _ = fi50
                _ = v51
            }
            switch 9223372036854775807 {
            case 4294967295:
                if (634 > 4294967295) {
                    v2 = (v2 && false)
                }
            case -23:
                v1 += v1
            }
            v2 = (1 > 42)
        } else {
            if ((-(v1)) != v1) {
                v1 = (v1 & v1)
            }
            v4["a"] = 65535
        }
    }
    if (1 <= 127) {
        v1 += v1
    } else {
        var v52 uint64 = 921
        v1 -= (v1 % (((v1 << ((2) & 63))) | 1))
        v52 -= (v52 & (v52 << ((2) & 63)))
        _ = v52
    }
    switch 662 {
    case 369:
        if v2 {
            fmt.Println(v2, v1)
            v1 -= 1
            v3 = append(v3, "\n\t")
        } else {
            if true {
                var v53 string = "a"
                switch 42 {
                case 850:
                    if false {
                        var v54 int8 = -103
                        _ = v54
                    }
                case -406:
                    v55 := GenS8{F0: 255, F1: -1, F2: int8(v1)}
                    _ = v55
                default:
                    var v56 bool = ((v1 * (v1 - 0)) != v1)
                    _ = v56
                }
                _ = v53
            } else {
                if ((!v2) || ((42 == 18446744073709551615) && v2)) {
                    if false {
                        var v57 int64 = (v1 / (((-(v1))) | 1))
                        var v58 uint16 = 1
                        v1 += (v57 / (((v57 * v57)) | 1))
                        _ = v57
                        _ = v58
                    } else {
                        fmt.Println(v1)
                        v2 = (!(2147483648 < 42))
                        var v59 uint32 = 2147483648
                        _ = v59
                    }
                    for fi60 := 0; fi60 < 18; fi60++ {
                        switch "\n\t" {
                        case "":
                            if v2 {
                                switch "" {
                                case "a":
                                    var v61 int16 = 1000
                                    _ = v61
                                case "hello":
                                    if (0.0 <= float32(1e10)) {
                                        switch fi60 {
                                        case -149:
                                            switch (fi60 | 651) {
                                            case 2147483647:
                                                switch fi60 {
                                                case -2147483648:
                                                    v62 := GenS9{F0: 148, F1: 1}
                                                    _ = v62
                                                case 32767:
                                                    v63 := v5["a"]
                                                    _ = v63
                                                default:
                                                    v64 := GenS7{F0: 42, F1: -106, F2: 1}
                                                    _ = v64
                                                }
                                            case 42:
                                                switch "\n\t" {
                                                case "hello":
                                                    for fi65 := 0; fi65 < 11; fi65++ {
                                                        switch "foo bar" {
                                                        case "foo bar":
                                                            v2 = v2
                                                            if len(v3) > 0 {
                                                                v3[0] = "a"
                                                            }
                                                        }
                                                        v2 = v2
                                                        _ = fi65
                                                    }
                                                case "foo bar":
                                                    fmt.Println(v1)
                                                case "a":
                                                    switch (fi60 | fi60) {
                                                    case 183:
                                                        v66 := GenS3{F0: -1}
                                                        _ = v66
                                                    }
                                                }
                                            case 4294967295:
                                                v67 := GenS10{F0: -3, F1: (fi60 + int(v1)), F2: 519}
                                                _ = v67
                                            default:
                                                v2 = v2
                                            }
                                        case 32767:
                                            v1 -= v1
                                        default:
                                            if len(v3) > 0 {
                                                v3[0] = ("foo bar" + "hello")
                                            }
                                        }
                                    } else {
                                        var v68 uint8 = 42
                                        _ = len(v3)
                                        _ = v68
                                    }
                                default:
                                    var v69 float32 = (0.0 + -2.25)
                                    _ = v69
                                }
                            }
                        case "foo bar":
                            fmt.Println(v1, v2, fi60)
                        case "\n\t":
                            _ = len(v5)
                        }
                        if (0.1 >= -2.25) {
                            v70 := GenS9{F0: 127, F1: 32767}
                            _ = v70
                        }
                        _ = fi60
                    }
                }
            }
        }
    default:
        v71 := GenS11{F0: 9223372036854775807, F1: ""}
        _ = v71
    }
    if len(v3) > 0 {
        v3[0] = ""
    }
    if v2 {
        v3 = append(v3, "a")
        v72 := GenS6{F0: (!true)}
        if (2147483647 == -2147483648) {
            v1 = -9223372036854775808
            var v73 int64 = (-(v1))
            var v74 float64 = 0.1
            _ = v73
            _ = v74
        } else {
            if (true && v2) {
                var v75 bool = (true || false)
                if v2 {
                    _ = len(v3)
                    v2 = (160 >= 438)
                }
                fmt.Println(v2, v1)
                _ = v75
            }
            if (2147483648 == 2147483648) {
                for fi76 := 0; fi76 < 12; fi76++ {
                    if (255 >= 187) {
                        v77 := v6[fi76]
                        _ = v77
                    }
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi76
                }
                delete(v4, "\n\t")
            } else {
                if (-128 != 100) {
                    if v2 {
                        v2 = v2
                    } else {
                        v78 := v6[127]
                        _ = v78
                    }
                    if v2 {
                        v2 = v2
                    } else {
                        var v79 float64 = ((float64(0.0) / 1e10) + (float64(-2.25) * float64((1.5 - ((3.14159 * (1.5 - 0.1)) * 0.0)))))
                        for fi80 := 0; fi80 < 10; fi80++ {
                            v79 = 1.5
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            v2 = (1000 <= 42)
                            _ = fi80
                        }
                        var v81 string = ("foo bar" + "a")
                        _ = v79
                        _ = v81
                    }
                    var v82 int = 49
                    _ = v82
                }
            }
            v2 = v2
        }
        _ = v72
    }
    var v83 int16 = 163
    var v84 int16 = v83
    if true {
        v2 = false
        var v85 int16 = v83
        _ = v85
    }
    v86 := GenS10{F0: 1, F1: 0, F2: 42}
    fmt.Println(v2, v83, v84)
    var v87 int8 = 0
    v1 -= (v1 << ((133) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v83:", v83)
    fmt.Println("v84:", v84)
    fmt.Println("v87:", v87)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v83
    _ = v84
    _ = v86
    _ = v87
}
