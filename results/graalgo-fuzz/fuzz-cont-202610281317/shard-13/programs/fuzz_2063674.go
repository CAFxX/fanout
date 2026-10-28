package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 uint64; F1 int32; F2 uint32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 uint8 }

func main() {
    var v1 uint16 = 333
    v2 := make(map[uint32]uint32)
    var v3 int8 = 42
    v4 := []string{"", ("\n\t" + ("a" + ""))}
    v5 := make(map[string]bool)
    v6 := make(map[string]string)
    v1 = 363
    var v7 uint64 = 4294967295
    if (((v3 >> ((2) & 63)) > (v3 << ((18446744073709551615) & 63))) && false) {
        for fi8 := 0; fi8 < 8; fi8++ {
            if (true || ((!((true || false) && (-87 >= 9223372036854775807))) && (161 != -32768))) {
                v3 -= (v3 % ((0) | 1))
                var v9 int16 = 1
                for fi10 := 0; fi10 < 13; fi10++ {
                    if false {
                        switch (fi10 / (((fi8 * fi8)) | 1)) {
                        case -2147483648:
                            var v11 string = ("hello" + "foo bar")
                            _ = v11
                        case -448:
                            v4 = append(v4, ("foo bar" + ""))
                        case 9223372036854775807:
                            fmt.Println(fi10)
                        default:
                            switch "a" {
                            case "foo bar":
                                if (false || (((fi10 | (^(fi8))) < int(v1)) && (9223372036854775807 > 42))) {
                                    fmt.Println(v9)
                                    fmt.Println(fi8)
                                }
                            default:
                                fmt.Println(v1)
                            }
                        }
                        switch fi10 {
                        case 4294967295:
                            switch -9223372036854775808 {
                            case 32768:
                                v7 = 654
                            case -2147483648:
                                v4 = append(v4, ("" + ("hello" + "")))
                            }
                        case 0:
                            switch ("a" + "foo bar") {
                            case "foo bar":
                                var v12 int64 = 0
                                _ = v12
                            default:
                                fmt.Println(v9)
                                v9 = (^(v9))
                            }
                        case 9223372036854775807:
                            v13 := GenS1{F0: 2147483648, F1: 0}
                            _ = v13
                        }
                        fmt.Println(fi8, v3, fi10)
                    } else {
                        v14 := GenS2{F0: v7, F1: 124, F2: 0}
                        _ = len(v4)
                        var v15 float64 = ((1.5 + float64(((3.14159 + (0.1 + float32(float64((float32((1e10 * -2.25)) - -2.25))))) + 1.5))) / 3.14159)
                        _ = v14
                        _ = v15
                    }
                    _ = fi10
                }
                _ = v9
            } else {
                delete(v5, "foo bar")
                v2[869] = 357
            }
            var v16 bool = (true || true)
            if v16 {
                v17 := GenS3{F0: 142}
                v18 := GenS4{F0: 222}
                v16 = (v16 && ((v1 % ((uint16(v7)) | 1)) >= (^(v1))))
                _ = v17
                _ = v18
            } else {
                switch "a" {
                case "\n\t":
                    var v19 int32 = 2147483647
                    _ = v19
                case "hello":
                    switch "hello" {
                    case "foo bar":
                        v20 := GenS3{F0: 61}
                        _ = v20
                    default:
                        delete(v5, "foo bar")
                    }
                case "a":
                    fmt.Println(v1, fi8)
                default:
                    v21 := GenS2{F0: v7, F1: 2147483647, F2: 4294967295}
                    _ = v21
                }
                v4 = append(v4, "a")
                delete(v6, "a")
            }
            _ = fi8
            _ = v16
        }
        switch ("" + "a") {
        case "hello":
            _ = len(v4)
        default:
            var v22 int = 966
            _ = v22
        }
        if (128 > 2) {
            if true {
                v7 = 4294967295
            } else {
                var v23 bool = true
                var v24 uint16 = uint16(v7)
                for fi25 := 0; fi25 < 13; fi25++ {
                    switch ("\n\t" + "a") {
                    case "a":
                        v23 = v23
                    case "foo bar":
                        var v26 int64 = 9223372036854775807
                        _ = v26
                    default:
                        _ = len(v5)
                    }
                    if ((0.0 > 3.14159) || (v3 <= v3)) {
                        v27 := GenS2{F0: (v7 & (^(v7))), F1: 0, F2: 4294967295}
                        var v28 uint = uint(v24)
                        _ = v27
                        _ = v28
                    } else {
                        if ((1e10 - 1e10) != 0.1) {
                            var v29 int32 = 153
                            _ = v29
                        }
                        v30 := GenS1{F0: 65535, F1: 1}
                        var v31 bool = ((-2.25 != (0.0 + (0.1 - 0.0))) || v23)
                        _ = v30
                        _ = v31
                    }
                    if true {
                        v3 = (v3 >> ((65535) & 63))
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        v3 = v3
                    } else {
                        v3 += (v3 << ((0) & 63))
                    }
                    _ = fi25
                }
                _ = v23
                _ = v24
            }
            for fi32 := 0; fi32 < 2; fi32++ {
                var v33 uint = 183
                for fi34 := 0; fi34 < 5; fi34++ {
                    if false {
                        var v35 int = (fi32 * int(v1))
                        var v36 int = v35
                        _ = v35
                        _ = v36
                    } else {
                        if (!("\n\t" >= ((("a" + "") + "\n\t") + (("hello" + "\n\t") + "hello")))) {
                            v37 := GenS5{F0: true}
                            switch ("foo bar" + "") {
                            case "a":
                                var v38 uint64 = v7
                                _ = v38
                            case "foo bar":
                                var v39 uint8 = 2
                                _ = v39
                            }
                            _ = v37
                        }
                        v40 := GenS6{F0: 1}
                        if true {
                            fmt.Println(v1)
                        } else {
                            v7 += (v7 + v7)
                            fmt.Println(v33, v3, v1)
                        }
                        _ = v40
                    }
                    v41 := GenS7{F0: 2147483648}
                    var v42 string = ""
                    v1 = v1
                    if true {
                        v43 := v6["\n\t"]
                        if ((v3 + 42) >= (^(v3))) {
                            fmt.Println(fi32)
                            _ = len(v4)
                            v33 = v33
                        }
                        var v44 uint = 0
                        _ = v43
                        _ = v44
                    }
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = fi34
                    _ = v41
                    _ = v42
                }
                var v45 int = 2147483647
                _ = fi32
                _ = v33
                _ = v45
            }
            if len(v4) > 0 {
                v4[0] = "\n\t"
            }
        } else {
            for fi46 := 0; fi46 < 3; fi46++ {
                _ = len(v2)
                _ = fi46
            }
        }
    } else {
        for fi47 := 0; fi47 < 1; fi47++ {
            v7 = v7
            _ = fi47
        }
        if true {
            var v48 bool = (int8(v1) == v3)
            v49 := GenS8{F0: 0}
            _ = v48
            _ = v49
        } else {
            fmt.Println(v7)
            if true {
                v50 := v6["hello"]
                _ = v50
            }
        }
    }
    v51 := GenS8{F0: 2}
    if true {
        delete(v2, 2)
    }
    fmt.Println(v3, v3, v7)
    for fi52 := 0; fi52 < 9; fi52++ {
        v1 -= (v1 / ((v1) | 1))
        if true {
            var v53 bool = false
            switch fi52 {
            case -923:
                switch ("\n\t" + "\n\t") {
                case "":
                    var v54 int32 = 1
                    _ = v54
                case "\n\t":
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                case "foo bar":
                    if true {
                        if (v53 || v53) {
                            for fi55 := 0; fi55 < 20; fi55++ {
                                v56 := GenS9{F0: -1}
                                fmt.Println(v53)
                                switch 576 {
                                case 42:
                                    v57 := GenS7{F0: 2147483648}
                                    _ = v57
                                case 127:
                                    v53 = false
                                }
                                _ = fi55
                                _ = v56
                            }
                            fmt.Println(v1, v53)
                            if ((v1 - v1) == (v1 - v1)) {
                                fmt.Println(fi52)
                            } else {
                                v58 := GenS10{F0: 42}
                                _ = v58
                            }
                        }
                        for fi59 := 0; fi59 < 12; fi59++ {
                            if (v53 || ((42 == 127) || v53)) {
                                var v60 float32 = -2.25
                                v1 -= (v1 << ((4294967295) & 63))
                                _ = v60
                            }
                            if (v3 > v3) {
                                v7 += 42
                                switch 2147483647 {
                                case -382:
                                    v7 = v7
                                case -9223372036854775808:
                                    fmt.Println(v53, v1)
                                default:
                                    var v61 float64 = (float64(0.1) - ((0.1 / (3.14159 / 1.0)) + float64(1e10)))
                                    _ = v61
                                }
                            } else {
                                var v62 int32 = -118
                                v62 = v62
                                _ = v62
                            }
                            v7 = (^(v7))
                            _ = fi59
                        }
                        fmt.Println(v3, v3)
                    } else {
                        v63 := GenS4{F0: 255}
                        v7 = 998
                        _ = v63
                    }
                }
            case 337:
                fmt.Println(fi52)
            case -852:
                fmt.Println(v3, fi52)
            default:
                switch "\n\t" {
                case "hello":
                    var v64 uint8 = 42
                    _ = v64
                case "\n\t":
                    fmt.Println(v7, fi52, v3)
                default:
                    var v65 uint16 = (v1 / ((42) | 1))
                    _ = v65
                }
            }
            _ = v53
        }
        _ = fi52
    }
    fmt.Println(v1, v1, v1)
    var v66 uint64 = 18446744073709551615
    v3 += (-(v3))
    v67 := GenS3{F0: 255}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v66:", v66)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v6
    _ = v66
    _ = v67
    _ = v7
}
