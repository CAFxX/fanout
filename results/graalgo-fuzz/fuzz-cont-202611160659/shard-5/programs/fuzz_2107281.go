package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint; F1 int32; F2 int8 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint; F1 string }
type GenS7 struct { F0 int16; F1 uint }
type GenS8 struct { F0 string; F1 int8 }
type GenS9 struct { F0 int64; F1 uint64; F2 uint16 }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 uint; F1 int32; F2 uint8 }
type GenS12 struct { F0 string }
type GenS13 struct { F0 int32; F1 uint32 }
func genFn0(p0 int8, p1 uint64, p2 uint) bool {
    var v4 uint16 = 2
    p1 -= (p1 << (((p2 * (^(p2)))) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return (-2147483648 > -485)
}

func main() {
    var v1 uint8 = 0
    var v2 int = 0
    v3 := []string{"", "a", "foo bar"}
    fmt.Println(v2, v1)
    var v5 float64 = -2.25
    v3 = append(v3, "")
    v6 := GenS1{F0: 42}
    var v7 bool = false
    for fi8 := 0; fi8 < 15; fi8++ {
        if true {
            switch v2 {
            case 2147483647:
                v9 := GenS2{F0: -85}
                _ = v9
            case 42:
                switch fi8 {
                case 139:
                    v5 = v5
                }
            }
            v2 = (fi8 << ((42) & 63))
            var v10 uint = 972
            _ = v10
        } else {
            if v7 {
                v11 := GenS3{F0: "a"}
                _ = v11
            } else {
                v1 += v1
                if false {
                    for fi12 := 0; fi12 < 4; fi12++ {
                        if (v7 || v7) {
                            var v13 uint64 = 18446744073709551615
                            _ = v13
                        } else {
                            var v14 string = "\n\t"
                            switch fi12 {
                            case 128:
                                v2 = fi8
                            case 0:
                                fmt.Println(v1, v1)
                                if v7 {
                                    switch v14 {
                                    case "":
                                        fmt.Println(v2, v5)
                                    case "\n\t":
                                        if (!false) {
                                            var v15 float64 = 3.14159
                                            v5 = v5
                                            _ = v15
                                        }
                                    case "hello":
                                        fmt.Println(v2)
                                    default:
                                        v7 = v7
                                    }
                                    var v16 uint16 = 924
                                    _ = v16
                                }
                            case 370:
                                if true {
                                    var v17 uint64 = 42
                                    if v7 {
                                        v1 -= v1
                                        v2 += (^(v2))
                                    }
                                    var v18 int64 = 9223372036854775807
                                    v19 := GenS1{F0: 18446744073709551615}
                                    if (1 != 936) {
                                        var v20 int64 = int64(v2)
                                        if v7 {
                                            fmt.Println(v2, v20, v7)
                                            v14 = v14
                                        } else {
                                            fmt.Println(fi8, fi12, v20)
                                            if true {
                                                var v21 int8 = 127
                                                v1 += v1
                                                if len(v3) > 0 {
                                                    v3[0] = v14
                                                }
                                                var v22 uint32 = 22
                                                v5 = 1.5
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                _ = v21
                                                _ = v22
                                            }
                                            v20 = v20
                                        }
                                        var v23 bool = ((v7 && v7) || (true && v7))
                                        _ = v20
                                        _ = v23
                                    }
                                    var v24 float32 = (1.5 + (3.14159 / 1.5))
                                    _ = v17
                                    _ = v18
                                    _ = v19
                                    _ = v24
                                } else {
                                    switch v14 {
                                    case "foo bar":
                                        v1 += (v1 ^ 1)
                                    case "":
                                        var v25 int32 = -1
                                        _ = v25
                                    }
                                    v26 := GenS4{F0: 698, F1: 2147483647, F2: 100}
                                    _ = v26
                                }
                            }
                            var v27 int64 = -413
                            _ = v14
                            _ = v27
                        }
                        var v28 float64 = v5
                        v5 = v5
                        _ = fi12
                        _ = v28
                    }
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v29 := GenS5{F0: true}
                    _ = v29
                }
            }
            if v7 {
                v2 = fi8
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            } else {
                for fi30 := 0; fi30 < 14; fi30++ {
                    var v31 float32 = 0.1
                    _ = fi30
                    _ = v31
                }
                v7 = (128 > (-(v2)))
            }
            v32 := GenS6{F0: 411, F1: "\n\t"}
            _ = v32
        }
        v33 := GenS6{F0: 1, F1: ("" + "hello")}
        _ = fi8
        _ = v33
    }
    if len(v3) > 0 {
        v3[0] = ("foo bar" + ((("hello" + ("\n\t" + "foo bar")) + "") + "foo bar"))
    }
    if v7 {
        if (111 != -2147483648) {
            for fi34 := 0; fi34 < 13; fi34++ {
                var v35 uint = 18446744073709551615
                v1 -= v1
                for fi36 := 0; fi36 < 8; fi36++ {
                    v5 = ((0.1 / 1.5) - 0.0)
                    var v37 float64 = (v5 + (1.5 + v5))
                    _ = fi36
                    _ = v37
                }
                _ = fi34
                _ = v35
            }
            for fi38 := 0; fi38 < 4; fi38++ {
                for fi39 := 0; fi39 < 3; fi39++ {
                    if false {
                        var v40 int64 = 1
                        _ = v40
                    } else {
                        var v41 bool = true
                        if len(v3) > 0 {
                            v3[0] = "\n\t"
                        }
                        _ = v41
                    }
                    var v42 uint = 993
                    var v43 uint8 = v1
                    if (-1 > -1) {
                        switch "\n\t" {
                        case "foo bar":
                            if ((!(!v7)) || true) {
                                var v44 uint32 = 2
                                v2 += fi39
                                v44 -= 1
                                switch 128 {
                                case 32767:
                                    var v45 float64 = ((v5 / (float64(float32((v5 * float64(0.0)))) * v5)) - v5)
                                    _ = v45
                                case 1:
                                    v43 -= 127
                                case 128:
                                    fmt.Println(fi38, fi38)
                                }
                                _ = v44
                            }
                        case "a":
                            switch "a" {
                            case "foo bar":
                                v5 = (3.14159 - v5)
                            case "a":
                                v46 := GenS7{F0: -690, F1: v42}
                                _ = v46
                            }
                            v7 = v7
                            _ = len(v3)
                            var v47 float64 = (-2.25 - (v5 + v5))
                            if false {
                                v48 := GenS8{F0: "hello", F1: 103}
                                v3 = append(v3, "hello")
                                fmt.Println(v2, v47)
                                fmt.Println(v2, v5, v43)
                                v3 = append(v3, "hello")
                                v47 = 1.5
                                _ = v48
                            }
                            if v7 {
                                v49 := GenS9{F0: 0, F1: 42, F2: 2}
                                _ = v49
                            }
                            _ = v47
                        default:
                            v1 += v1
                            v50 := GenS10{F0: uint16(v43)}
                            v1 = v1
                            _ = v50
                        }
                        v43 += v1
                        if (false && (((float64(3.14159) * ((3.14159 * 1e10) - v5)) - float64((float32(((v5 / v5) - 1.5)) - ((3.14159 / ((1.5 / float32(v5)) + -2.25)) + 3.14159)))) == (0.1 - (v5 - ((v5 - v5) / ((0.1 * (v5 + v5)) * float64((((0.1 / 1.0) / -2.25) + 0.1)))))))) {
                            v51 := GenS11{F0: (v42 & v42), F1: -2147483648, F2: uint8(v42)}
                            _ = v51
                        }
                    }
                    _ = fi39
                    _ = v42
                    _ = v43
                }
                v2 = (fi38 & 4294967295)
                _ = fi38
            }
        }
    }
    v2 += v2
    var v52 uint = 33
    _ = len(v3)
    v7 = v7
    v53 := GenS4{F0: v52, F1: -2147483648, F2: -128}
    if v7 {
        var v54 int = (v2 % ((-933) | 1))
        _ = v54
    } else {
        for fi55 := 0; fi55 < 19; fi55++ {
            fmt.Println(v7, v1)
            switch "a" {
            case "\n\t":
                fmt.Println(v1, v5, v7)
            case "foo bar":
                switch "hello" {
                case "":
                    if (true && (32767 != int16(v1))) {
                        for fi56 := 0; fi56 < 7; fi56++ {
                            if len(v3) > 0 {
                                v3[0] = (((("hello" + "\n\t") + "") + ("hello" + ("\n\t" + "foo bar"))) + "")
                            }
                            _ = fi56
                        }
                        if v7 {
                            var v57 int16 = 322
                            fmt.Println(v57, v57)
                            _ = v57
                        } else {
                            var v58 float32 = 3.14159
                            for fi59 := 0; fi59 < 2; fi59++ {
                                var v60 float32 = float32(0.0)
                                _ = fi59
                                _ = v60
                            }
                            _ = v58
                        }
                        v52 = v52
                    } else {
                        var v61 uint32 = uint32(v52)
                        switch 0 {
                        case 32767:
                            v62 := GenS12{F0: "a"}
                            _ = v62
                        }
                        _ = v61
                    }
                case "\n\t":
                    v3 = append(v3, ("a" + "a"))
                default:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
            default:
                fmt.Println(v5, fi55, v2)
            }
            _ = fi55
        }
        fmt.Println(v1, v1, v2)
        var v63 int16 = -551
        _ = v63
    }
    var v64 uint8 = (v1 & v1)
    if v7 {
        v2 = 2147483647
        if false {
            if true {
                if false {
                    var v65 uint8 = v64
                    v2 = v2
                    _ = v65
                }
            }
        }
    } else {
        if v7 {
            v66 := GenS13{F0: 2147483647, F1: 42}
            _ = v66
        }
        for fi67 := 0; fi67 < 11; fi67++ {
            for fi68 := 0; fi68 < 13; fi68++ {
                v3 = append(v3, "a")
                fmt.Println(v64, v2, v52)
                _ = fi68
            }
            _ = fi67
        }
        v3 = append(v3, "")
    }
    v5 -= ((v5 / ((v5 - v5) - -2.25)) * 1e10)
    v3 = append(v3, ("" + ""))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v64:", v64)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v52
    _ = v53
    _ = v6
    _ = v64
    _ = v7
}
