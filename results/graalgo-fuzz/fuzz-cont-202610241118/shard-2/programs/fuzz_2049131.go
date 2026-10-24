package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 int32; F1 bool }
type GenS4 struct { F0 string; F1 int; F2 int32 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int8; F1 int8 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint32; F1 uint16; F2 uint64 }
type GenS9 struct { F0 int32; F1 uint32 }
type GenS10 struct { F0 int16; F1 uint16 }
type GenS11 struct { F0 string; F1 uint; F2 string }
type GenS12 struct { F0 int32; F1 int8 }
type GenS13 struct { F0 uint32; F1 uint32 }
type GenS14 struct { F0 uint }
type GenS15 struct { F0 int64; F1 int8 }
type GenS16 struct { F0 int8; F1 uint32; F2 uint16 }
func genFn0() int8 {
    var v6 uint = 1
    if (!true) {
        v6 = (v6 | (^(v6)))
    } else {
        fmt.Println(v6)
    }
    if false {
        var v7 bool = (0 != 622)
        _ = v7
    }
    _ = v6
    return 127
}

func main() {
    var v1 uint8 = 226
    var v2 float32 = 3.14159
    v3 := []uint32{2147483648, 4294967295, 2, 1}
    v4 := []string{"\n\t", "a", "\n\t", ("foo bar" + "foo bar")}
    var v5 int16 = 42
    var v8 uint16 = 218
    fmt.Println(v8)
    switch "foo bar" {
    case "hello":
        v8 -= (^(v8))
    case "foo bar":
        v1 = (v1 % (((v1 & 1)) | 1))
    case "\n\t":
        v9 := GenS1{F0: 42}
        _ = v9
    default:
        for fi10 := 0; fi10 < 5; fi10++ {
            v11 := GenS2{F0: v8, F1: 2147483647}
            switch fi10 {
            case 1:
                if len(v4) > 0 {
                    v4[0] = ("foo bar" + "hello")
                }
            case 790:
                v12 := GenS1{F0: -93}
                _ = v12
            default:
                v5 -= (v5 | (v5 - -1))
            }
            for fi13 := 0; fi13 < 8; fi13++ {
                if (float64(v2) == float64((v2 * 0.1))) {
                    v14 := GenS3{F0: 2147483647, F1: true}
                    var v15 bool = true
                    if v15 {
                        v15 = v15
                    } else {
                        if ((0 <= 0) && v15) {
                            v8 = (v8 | (v8 & (v8 / (((v8 + (^(v8)))) | 1))))
                            switch "" {
                            case "foo bar":
                                v5 += 761
                            case "\n\t":
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            default:
                                var v16 int16 = v5
                                _ = v16
                            }
                            if true {
                                fmt.Println(fi13, v1, fi13)
                                var v17 uint32 = 0
                                fmt.Println(v8)
                                fmt.Println(fi13, v5)
                                fmt.Println(fi13)
                                v18 := GenS2{F0: v8, F1: -1}
                                _ = len(v4)
                                v19 := GenS4{F0: "hello", F1: (fi13 * fi10), F2: 42}
                                _ = v17
                                _ = v18
                                _ = v19
                            }
                        } else {
                            switch ("a" + "hello") {
                            case "a":
                                var v20 int64 = 42
                                _ = v20
                            case "\n\t":
                                if (390 > 23) {
                                    var v21 float32 = ((v2 / v2) - float32(0.0))
                                    v15 = v15
                                    v8 -= (v8 | uint16(v1))
                                    _ = v21
                                } else {
                                    v2 = float32(-2.25)
                                }
                            }
                            v1 = v1
                            _ = len(v3)
                            switch ("\n\t" + "foo bar") {
                            case "a":
                                v22 := GenS5{F0: 9}
                                _ = v22
                            case "":
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            case "foo bar":
                                fmt.Println(v1, fi13)
                            default:
                                var v23 int16 = v5
                                _ = v23
                            }
                            var v24 uint8 = (v1 - uint8(v8))
                            _ = v24
                        }
                    }
                    switch ("hello" + "hello") {
                    case "foo bar":
                        switch ("a" + "a") {
                        case "a":
                            if true {
                                v15 = v15
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                        }
                    default:
                        v25 := GenS6{F0: 43, F1: 1}
                        _ = v25
                    }
                    var v26 int32 = -1
                    _ = v14
                    _ = v15
                    _ = v26
                }
                var v27 string = (("\n\t" + "") + (((("\n\t" + "\n\t") + "hello") + ((("hello" + "\n\t") + "") + "")) + ""))
                var v28 bool = true
                _ = fi13
                _ = v27
                _ = v28
            }
            _ = fi10
            _ = v11
        }
    }
    v3 = append(v3, 65535)
    var v29 bool = false
    var v30 int32 = 0
    var v31 uint16 = uint16(v1)
    v32 := GenS7{F0: v8}
    v30 = (v30 << ((2) & 63))
    switch -855 {
    case -769:
        switch "\n\t" {
        case "a":
            var v33 int16 = (v5 / ((-928) | 1))
            _ = v33
        case "":
            v34 := GenS8{F0: 694, F1: 1, F2: 178}
            _ = v34
        case "hello":
            v35 := GenS9{F0: v30, F1: 65535}
            _ = v35
        }
    case 0:
        for fi36 := 0; fi36 < 10; fi36++ {
            switch (("foo bar" + ((("a" + "") + "a") + "\n\t")) + "") {
            case "foo bar":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            case "\n\t":
                switch ("\n\t" + ("hello" + ("a" + "hello"))) {
                case "\n\t":
                    switch 2147483647 {
                    case 140:
                        v31 = (v8 >> ((42) & 63))
                    case 728:
                        fmt.Println(v2)
                    case 32768:
                        switch "hello" {
                        case "":
                            v37 := GenS5{F0: 0}
                            _ = v37
                        case "a":
                            fmt.Println(v5)
                        case "\n\t":
                            if (v30 != int32(v5)) {
                                var v38 int64 = 9223372036854775807
                                _ = v38
                            } else {
                                fmt.Println(v8, v2)
                                var v39 string = "a"
                                var v40 int16 = v5
                                _ = v39
                                _ = v40
                            }
                        default:
                            var v41 float64 = 3.14159
                            _ = v41
                        }
                    default:
                        var v42 uint8 = (v1 ^ v1)
                        _ = v42
                    }
                case "a":
                    v8 += v8
                case "hello":
                    var v43 float64 = float64(((v2 / v2) * float32(1.5)))
                    _ = v43
                default:
                    for fi44 := 0; fi44 < 19; fi44++ {
                        if (!v29) {
                            if false {
                                var v45 int64 = -9223372036854775808
                                v5 = (v5 | (v5 / ((v5) | 1)))
                                if ((v45 | v45) <= int64(v30)) {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                } else {
                                    var v46 uint32 = 274
                                    var v47 uint8 = 255
                                    fmt.Println(fi44, fi36, fi44)
                                    if (!false) {
                                        if (839 != 35) {
                                            var v48 uint8 = v47
                                            if len(v4) > 0 {
                                                v4[0] = "\n\t"
                                            }
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            v49 := GenS10{F0: -950, F1: v31}
                                            v50 := GenS11{F0: "a", F1: 255, F2: "hello"}
                                            _ = v48
                                            _ = v49
                                            _ = v50
                                        }
                                    }
                                    if (("hello" + "\n\t") > "foo bar") {
                                        if true {
                                            v4 = append(v4, "\n\t")
                                            if (v29 && false) {
                                                v2 -= 0.1
                                                v3 = append(v3, (v46 / (((^(v46))) | 1)))
                                                var v51 uint16 = v31
                                                _ = v51
                                            } else {
                                                if len(v3) > 0 {
                                                    v3[0] = (v46 & v46)
                                                }
                                            }
                                        } else {
                                            if true {
                                                var v52 int = fi36
                                                if v29 {
                                                    _ = len(v4)
                                                }
                                                _ = v52
                                            }
                                            fmt.Println(fi36, v29)
                                            switch 9223372036854775807 {
                                            case -948:
                                                var v53 uint16 = 2
                                                _ = v53
                                            case 1:
                                                v31 = (^(v8))
                                            case 9223372036854775807:
                                                v54 := GenS12{F0: v30, F1: -8}
                                                _ = v54
                                            }
                                            fmt.Println(v30, v2)
                                        }
                                    } else {
                                        v55 := GenS13{F0: 42, F1: 2147483647}
                                        if ((!(65535 != 4294967295)) || (!(uint8(v46) < 128))) {
                                            v31 -= 65535
                                            v56 := GenS10{F0: (v5 << ((4294967295) & 63)), F1: v8}
                                            if v29 {
                                                v45 += v45
                                                switch "a" {
                                                case "\n\t":
                                                    fmt.Println(v45)
                                                default:
                                                    var v57 uint = 42
                                                    v46 = (v46 / (((v46 << (((v57 % (((v57 << ((42) & 63))) | 1))) & 63))) | 1))
                                                    v5 -= (-(v5))
                                                    _ = v57
                                                }
                                            }
                                            _ = v56
                                        }
                                        if v29 {
                                            var v58 uint32 = (v46 - v46)
                                            v5 += 0
                                            if (uint16(v58) <= v8) {
                                                var v59 uint64 = 18446744073709551615
                                                _ = v59
                                            } else {
                                                v58 -= v46
                                                v60 := GenS8{F0: v46, F1: 65535, F2: 1}
                                                var v61 uint64 = 4294967295
                                                _ = v60
                                                _ = v61
                                            }
                                            _ = v58
                                        }
                                        v62 := GenS12{F0: v30, F1: genFn0()}
                                        var v63 int8 = 100
                                        _ = v55
                                        _ = v62
                                        _ = v63
                                    }
                                    v64 := GenS3{F0: int32(v45), F1: (((fi44 - fi36) >= fi44) && v29)}
                                    _ = v46
                                    _ = v47
                                    _ = v64
                                }
                                _ = v45
                            } else {
                                var v65 int32 = (v30 % (((v30 + (^(v30)))) | 1))
                                var v66 string = "a"
                                _ = v65
                                _ = v66
                            }
                            switch ((("foo bar" + "\n\t") + "\n\t") + "hello") {
                            case "":
                                var v67 uint = 726
                                _ = v67
                            case "foo bar":
                                var v68 uint64 = 42
                                _ = v68
                            case "\n\t":
                                v30 += -580
                            }
                        }
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        v31 -= v31
                        _ = fi44
                    }
                }
            default:
                v1 -= 0
            }
            if false {
                if len(v3) > 0 {
                    v3[0] = 739
                }
                switch 9223372036854775807 {
                case -942:
                    var v69 int = (fi36 >> ((1) & 63))
                    _ = v69
                case -502:
                    var v70 int8 = 15
                    _ = v70
                default:
                    v2 = (float32(1.5) - (v2 / v2))
                }
            } else {
                v71 := GenS14{F0: 4294967295}
                v72 := GenS11{F0: "a", F1: 0, F2: ("a" + "foo bar")}
                v73 := GenS1{F0: -35}
                _ = v71
                _ = v72
                _ = v73
            }
            _ = fi36
        }
    }
    var v74 int8 = -128
    if false {
        v75 := GenS15{F0: 31, F1: genFn0()}
        _ = v75
    } else {
        v29 = (!v29)
        for fi76 := 0; fi76 < 17; fi76++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v77 := GenS16{F0: v74, F1: 851, F2: 255}
            _ = fi76
            _ = v77
        }
    }
    var v78 int = 0
    v1 += (v1 << ((4294967295) & 63))
    _ = len(v3)
    switch (^(v78)) {
    case -500:
        v74 += v74
    case 128:
        fmt.Println(v2)
    case -589:
        v30 += (v30 | 208)
    }
    if v29 {
        var v79 float32 = 3.14159
        _ = v79
    } else {
        v30 += -2147483648
        fmt.Println(v30, v1)
        var v80 string = ""
        _ = v80
    }
    v81 := GenS9{F0: -2147483648, F1: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v74:", v74)
    fmt.Println("v78:", v78)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v74
    _ = v78
    _ = v8
    _ = v81
}
