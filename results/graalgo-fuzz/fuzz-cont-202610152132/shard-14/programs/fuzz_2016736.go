package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int64; F1 bool }
type GenS5 struct { F0 uint8; F1 int16 }
type GenS6 struct { F0 uint32; F1 string; F2 bool }
type GenS7 struct { F0 uint32; F1 int16; F2 uint }
type GenS8 struct { F0 uint16; F1 uint64; F2 uint16 }
type GenS9 struct { F0 uint64; F1 uint16 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 uint16; F1 int; F2 string }
type GenS12 struct { F0 int16 }
func genFn0(p0 uint64) int8 {
    for fi5 := 0; fi5 < 4; fi5++ {
        var v6 float64 = 1.5
        _ = fi5
        _ = v6
    }
    var v7 uint64 = p0
    fmt.Println(p0, v7)
    p0 += v7
    p0 = p0
    _ = p0
    _ = v7
    return 0
}
func genFn1(p0 int16, p1 int32) int16 {
    switch 0 {
    case 857:
        switch "hello" {
        case "":
            if true {
                var v8 uint32 = 1
                _ = v8
            }
        case "a":
            if false {
                p0 = (p0 - (p0 & p0))
            }
        }
    default:
        if false {
            var v9 uint = 4294967295
            var v10 uint8 = 0
            _ = v10
            _ = v9
        } else {
            fmt.Println(p1, p0)
            switch -2147483648 {
            case 0:
                v11 := []int{809}
                _ = v11
            }
            if true {
                p1 = p1
                v12 := []int{32768, -2147483648, 616, 9223372036854775807}
                switch 854 {
                case 599:
                    fmt.Println(p0)
                case 42:
                    var v13 uint = 18446744073709551615
                    _ = v13
                }
                _ = v12
            }
        }
    }
    v14 := GenS1{F0: 2, F1: 42, F2: (p0 & (p0 * -834))}
    _ = p0
    _ = p1
    _ = v14
    return p0
}

func main() {
    v1 := []string{(("hello" + ("" + "foo bar")) + ("\n\t" + "foo bar")), "hello", ("" + "hello")}
    var v2 int64 = -67
    var v3 string = "\n\t"
    var v4 int8 = -1
    fmt.Println(v2)
    v15 := GenS1{F0: 255, F1: (v2 * v2), F2: 32767}
    var v16 int32 = -440
    for fi17 := 0; fi17 < 16; fi17++ {
        if len(v1) > 0 {
            v1[0] = "a"
        }
        if (!true) {
            v18 := GenS1{F0: 33, F1: v2, F2: -1}
            _ = v18
        } else {
            var v19 int = fi17
            _ = v19
        }
        var v20 float64 = ((float64(0.0) - 1.5) - 0.1)
        _ = fi17
        _ = v20
    }
    var v21 string = v3
    var v22 uint32 = 2
    for fi23 := 0; fi23 < 15; fi23++ {
        v24 := GenS2{F0: v16}
        if (32767 != 0) {
            v4 -= v4
            v25 := GenS3{F0: uint64(v22)}
            var v26 int16 = -1
            _ = v25
            _ = v26
        } else {
            var v27 float32 = 0.1
            _ = v27
        }
        if true {
            v1 = append(v1, ("" + "hello"))
            fmt.Println(v3, v3)
            var v28 uint64 = 616
            _ = v28
        } else {
            switch "hello" {
            case "foo bar":
                if true {
                    switch (v3 + v21) {
                    case "a":
                        v21 = (((v3 + v3) + v21) + "")
                    }
                    var v29 float32 = (0.1 * (1e10 + 1.5))
                    for fi30 := 0; fi30 < 16; fi30++ {
                        var v31 string = v21
                        if true {
                            switch v21 {
                            case "\n\t":
                                fmt.Println(v29, v4, v4)
                            case "a":
                                var v32 uint32 = (v22 * v22)
                                _ = v32
                            case "foo bar":
                                var v33 string = "\n\t"
                                _ = v33
                            default:
                                if len(v1) > 0 {
                                    v1[0] = v21
                                }
                            }
                            switch v3 {
                            case "foo bar":
                                v16 = v16
                            case "\n\t":
                                v22 += 460
                            }
                            v34 := GenS4{F0: v2, F1: true}
                            _ = v34
                        } else {
                            switch (v3 + "") {
                            case "foo bar":
                                var v35 float32 = v29
                                _ = v35
                            case "":
                                v36 := GenS3{F0: 0}
                                _ = v36
                            }
                            v21 = v3
                            v29 += (v29 + (((1e10 / -2.25) + v29) + (1e10 * (v29 - 3.14159))))
                        }
                        switch (-(fi30)) {
                        case 911:
                            switch v31 {
                            case "\n\t":
                                switch ("hello" + (v21 + "foo bar")) {
                                case "hello":
                                    switch fi30 {
                                    case 127:
                                        v1 = append(v1, "\n\t")
                                    default:
                                        var v37 int64 = v2
                                        var v38 string = v3
                                        v39 := GenS5{F0: 0, F1: -1}
                                        v37 -= v2
                                        var v40 int8 = genFn0(42)
                                        fmt.Println(v21)
                                        _ = v37
                                        _ = v38
                                        _ = v39
                                        _ = v40
                                    }
                                case "\n\t":
                                    var v41 uint8 = 9
                                    _ = v41
                                case "a":
                                    var v42 float64 = float64(v29)
                                    _ = v42
                                }
                                var v43 float64 = 0.1
                                _ = v43
                            }
                        case 32768:
                            fmt.Println(v22, v4, v31)
                        case -224:
                            switch v31 {
                            case "hello":
                                v4 += (v4 >> ((65535) & 63))
                            }
                        default:
                            var v44 int32 = (^(v16))
                            switch ("hello" + "\n\t") {
                            case "hello":
                                var v45 float64 = float64((v29 + 0.1))
                                _ = v45
                            case "\n\t":
                                if false {
                                    var v46 int16 = 42
                                    v1 = append(v1, v21)
                                    if (!(((v2 >> ((571) & 63)) <= -9223372036854775808) || (!((((v29 * float32(0.0)) - float32(((-2.25 / 1.0) + ((1e10 + (1.5 - -2.25)) - 1e10)))) == -2.25) || (((true && true) && true) && true))))) {
                                        v46 = (v46 + (v46 & v46))
                                        v21 = v31
                                        if (65535 > v22) {
                                            v16 += (v44 ^ (v44 << ((42) & 63)))
                                            var v47 uint32 = (v22 / ((v22) | 1))
                                            v2 -= 9223372036854775807
                                            _ = v47
                                        } else {
                                            if false {
                                                switch v31 {
                                                case "":
                                                    v31 = v31
                                                case "a":
                                                    if true {
                                                        var v48 uint = 18446744073709551615
                                                        switch v3 {
                                                        case "foo bar":
                                                            v49 := GenS5{F0: 0, F1: (v46 + (v46 & (v46 % ((genFn1(v46, (v44 - (-(v16))))) | 1))))}
                                                            _ = v49
                                                        }
                                                        _ = v48
                                                    }
                                                default:
                                                    v29 = v29
                                                }
                                                v50 := GenS6{F0: 42, F1: v3, F2: true}
                                                _ = v50
                                            } else {
                                                fmt.Println(v2, v2)
                                                v44 += v16
                                                v51 := GenS5{F0: 128, F1: v46}
                                                _ = v51
                                            }
                                            v52 := GenS7{F0: 42, F1: v46, F2: 584}
                                            v53 := GenS5{F0: 128, F1: (v46 ^ (v46 - -32768))}
                                            _ = v52
                                            _ = v53
                                        }
                                    }
                                    _ = v46
                                } else {
                                    v1 = append(v1, v21)
                                    if (635 != 2) {
                                        fmt.Println(v29)
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        var v54 uint8 = 1
                                        _ = v54
                                    } else {
                                        v55 := GenS8{F0: 42, F1: 1, F2: 175}
                                        v1 = append(v1, (v3 + v31))
                                        v31 = (v31 + (((v3 + ("\n\t" + "foo bar")) + v31) + ("hello" + v21)))
                                        fmt.Println(fi30)
                                        var v56 uint64 = 4294967295
                                        _ = v55
                                        _ = v56
                                    }
                                }
                            case "foo bar":
                                v44 -= (v44 << ((2) & 63))
                            }
                            _ = v44
                        }
                        _ = fi30
                        _ = v31
                    }
                    _ = v29
                }
            case "":
                if (v3 >= "") {
                    fmt.Println(v3, v21, v21)
                    v57 := GenS9{F0: 4294967295, F1: 255}
                    switch v3 {
                    case "\n\t":
                        var v58 uint32 = v22
                        _ = v58
                    default:
                        v16 -= 42
                    }
                    _ = v57
                }
            }
            var v59 bool = (4294967295 > 42)
            v59 = v59
            _ = v59
        }
        _ = fi23
        _ = v24
    }
    switch 4294967295 {
    case 32767:
        switch 2147483647 {
        case 127:
            var v60 int64 = -771
            _ = v60
        case 975:
            var v61 uint16 = 84
            _ = v61
        case 197:
            fmt.Println(v21, v16)
        }
    case -2147483648:
        if true {
            v22 -= v22
            if false {
                switch v21 {
                case "a":
                    switch ((v3 + v3) + ("\n\t" + (v3 + "a"))) {
                    case "hello":
                        v3 = v3
                    case "\n\t":
                        switch 32768 {
                        case 42:
                            v2 -= (v2 - v2)
                        }
                    case "":
                        if false {
                            var v62 uint64 = 0
                            for fi63 := 0; fi63 < 5; fi63++ {
                                for fi64 := 0; fi64 < 20; fi64++ {
                                    v1 = append(v1, v3)
                                    v1 = append(v1, "hello")
                                    if true {
                                        fmt.Println(v21, v4, v3)
                                        var v65 float64 = (0.1 / (3.14159 * float64(((((((3.14159 * 1e10) / (0.1 + 0.1)) - 1.5) - (float32(1.5) - (1.5 * (float32((1e10 / 3.14159)) + 1e10)))) - 0.0) - 0.0))))
                                        _ = v65
                                    } else {
                                        if false {
                                            v66 := GenS4{F0: (v2 >> ((866) & 63)), F1: ((^(v62)) == (v62 | 0))}
                                            switch fi64 {
                                            case 32767:
                                                var v67 uint = 1
                                                v68 := GenS10{F0: fi63}
                                                _ = v67
                                                _ = v68
                                            }
                                            var v69 uint = 255
                                            _ = v66
                                            _ = v69
                                        }
                                    }
                                    _ = fi64
                                }
                                _ = fi63
                            }
                            _ = v62
                        } else {
                            v4 -= (v4 % (((v4 / ((genFn0(1)) | 1))) | 1))
                            v3 = (("\n\t" + v3) + "\n\t")
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                    }
                }
            }
            var v70 uint = 4294967295
            _ = v70
        }
    case 310:
        v22 += v22
    default:
        if false {
            v21 = v21
            var v71 float32 = 1e10
            _ = v71
        } else {
            v72 := GenS8{F0: 708, F1: 18446744073709551615, F2: 654}
            v2 = v2
            for fi73 := 0; fi73 < 7; fi73++ {
                _ = len(v1)
                for fi74 := 0; fi74 < 10; fi74++ {
                    if false {
                        v75 := GenS11{F0: 305, F1: (^(fi73)), F2: v21}
                        switch fi74 {
                        case -9223372036854775808:
                            v76 := GenS12{F0: -32768}
                            _ = v76
                        case 662:
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        default:
                            v2 = (v2 & (v2 * (v2 % (((v2 % ((v2) | 1))) | 1))))
                            switch v3 {
                            case "foo bar":
                                var v77 int32 = v16
                                _ = v77
                            case "hello":
                                v2 = 55
                            case "\n\t":
                                _ = len(v1)
                            default:
                                v21 = ("a" + ("hello" + ((("" + (v21 + v3)) + ((v21 + ("a" + "\n\t")) + "hello")) + v3)))
                            }
                            var v78 float64 = (0.1 * 3.14159)
                            var v79 uint8 = 154
                            _ = v78
                            _ = v79
                        }
                        _ = v75
                    } else {
                        var v80 float64 = 1e10
                        if (2 <= 354) {
                            v3 = v21
                            v1 = append(v1, v3)
                            v2 -= -456
                        }
                        _ = v80
                    }
                    _ = fi74
                }
                v21 = v21
                _ = fi73
            }
            _ = v72
        }
    }
    switch (v3 + (v3 + "")) {
    case "":
        if false {
            v21 = ((((v21 + (v21 + (v21 + v3))) + ("a" + "a")) + "") + ("a" + v21))
            if len(v1) > 0 {
                v1[0] = ("" + ("hello" + v21))
            }
            fmt.Println(v21, v21, v4)
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
