package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 int16 }
type GenS2 struct { F0 uint16; F1 uint; F2 uint }
type GenS3 struct { F0 uint8; F1 string; F2 int16 }
type GenS4 struct { F0 bool }

func main() {
    v1 := []uint32{794, 42, 210, 2147483648}
    v2 := make(map[string]uint32)
    var v3 bool = false
    switch ("hello" + "foo bar") {
    case "\n\t":
        switch -246 {
        case 32768:
            v4 := GenS1{F0: 856, F1: 100, F2: -32768}
            _ = v4
        default:
            delete(v2, "a")
        }
    case "foo bar":
        fmt.Println(v3)
    case "hello":
        var v5 int32 = -319
        _ = v5
    default:
        v3 = v3
    }
    if (0 != -32768) {
        var v6 int = 42
        _ = v6
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if (555 == 32) {
            v7 := GenS1{F0: 754, F1: 100, F2: 1}
            var v8 string = "a"
            _ = v7
            _ = v8
        }
        v3 = (((3.14159 + 1.5) * (((3.14159 * 0.0) - 1.5) - (3.14159 - (-2.25 + 3.14159)))) == 0.1)
    }
    v9 := GenS1{F0: -2147483648, F1: 1, F2: 1}
    v10 := GenS2{F0: 0, F1: 2, F2: 18446744073709551615}
    v3 = false
    fmt.Println(v3)
    if (!(!v3)) {
        if true {
            v3 = (!(-1 < 42))
            v3 = (-2.25 > (3.14159 + 3.14159))
            switch "" {
            case "foo bar":
                var v11 int64 = 42
                _ = v11
            case "a":
                v3 = (2147483647 <= 128)
            case "":
                if v3 {
                    if len(v1) > 0 {
                        v1[0] = 242
                    }
                    if v3 {
                        var v12 uint32 = 1
                        _ = v12
                    } else {
                        var v13 uint32 = 42
                        _ = v13
                    }
                    var v14 uint = 18446744073709551615
                    _ = v14
                } else {
                    v3 = true
                    delete(v2, "a")
                }
            }
        } else {
            v3 = v3
        }
        var v15 int8 = -1
        _ = v15
    } else {
        v3 = (255 >= 2)
        if false {
            for fi16 := 0; fi16 < 19; fi16++ {
                if v3 {
                    v3 = (false && (667 >= 88))
                    var v17 int16 = 42
                    _ = v17
                } else {
                    switch ("a" + "a") {
                    case "":
                        v3 = (3.14159 >= float64(1e10))
                    default:
                        var v18 uint64 = 4294967295
                        _ = v18
                    }
                }
                v2[(("\n\t" + "") + ("foo bar" + "a"))] = 4294967295
                _ = fi16
            }
        }
    }
    for fi19 := 0; fi19 < 13; fi19++ {
        fmt.Println(fi19, v3)
        var v20 int8 = 1
        _ = fi19
        _ = v20
    }
    v1 = append(v1, 4294967295)
    v21 := GenS2{F0: 359, F1: 255, F2: 42}
    var v22 uint64 = 2
    if false {
        if false {
            if v3 {
                v22 = v22
                fmt.Println(v3, v3)
                v23 := v2["foo bar"]
                _ = v23
            } else {
                var v24 string = "foo bar"
                for fi25 := 0; fi25 < 20; fi25++ {
                    v3 = false
                    for fi26 := 0; fi26 < 18; fi26++ {
                        switch "hello" {
                        case "a":
                            var v27 int8 = 0
                            _ = v27
                        default:
                            v2[v24] = 1
                            switch "foo bar" {
                            case "":
                                switch 2147483647 {
                                case 522:
                                    var v28 int = fi25
                                    _ = v28
                                default:
                                    var v29 float32 = float32(float64(1e10))
                                    if (fi25 == 128) {
                                        fmt.Println(v24, v24, fi25)
                                        fmt.Println(v24, v29, v24)
                                        var v30 int32 = -878
                                        _ = v30
                                    } else {
                                        var v31 int = 9223372036854775807
                                        _ = v31
                                    }
                                    var v32 int32 = 0
                                    _ = v29
                                    _ = v32
                                }
                            case "hello":
                                v22 -= (^(v22))
                            default:
                                switch -824 {
                                case 32767:
                                    if ((!(42 >= 1000)) && v3) {
                                        if (fi25 < (fi26 % (((fi26 ^ 4294967295)) | 1))) {
                                            fmt.Println(fi26)
                                        } else {
                                            v33 := GenS1{F0: 2147483647, F1: 80, F2: 1000}
                                            v22 = 0
                                            _ = v33
                                        }
                                    } else {
                                        switch (fi25 & (fi25 ^ fi26)) {
                                        case -2147483648:
                                            if (848 >= 255) {
                                                var v34 uint16 = 982
                                                v24 = v24
                                                var v35 int = -462
                                                _ = v34
                                                _ = v35
                                            } else {
                                                switch ("" + "hello") {
                                                case "\n\t":
                                                    v22 = (v22 - v22)
                                                case "":
                                                    fmt.Println(v22, fi25)
                                                case "hello":
                                                    fmt.Println(fi26)
                                                }
                                                v22 += v22
                                                var v36 string = v24
                                                _ = v36
                                            }
                                        case -718:
                                            if false {
                                                v3 = false
                                            } else {
                                                if false {
                                                    fmt.Println(fi26, v22)
                                                    fmt.Println(v24, v24)
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                }
                                                var v37 uint64 = v22
                                                _ = v37
                                            }
                                        default:
                                            v3 = false
                                        }
                                        switch v24 {
                                        case "":
                                            v1 = append(v1, 0)
                                        case "a":
                                            v22 += v22
                                        case "hello":
                                            v3 = v3
                                        default:
                                            var v38 uint16 = 42
                                            _ = v38
                                        }
                                    }
                                }
                            }
                        }
                        _ = fi26
                    }
                    _ = fi25
                }
                _ = v24
            }
            v22 += v22
            for fi39 := 0; fi39 < 14; fi39++ {
                if ((-2.25 + (1e10 - 3.14159)) >= ((1.5 / ((0.0 * (1e10 + 1.5)) - (-2.25 + (0.0 / (-2.25 - (0.1 - 0.1)))))) / 1.0)) {
                    var v40 float64 = 3.14159
                    for fi41 := 0; fi41 < 16; fi41++ {
                        v42 := GenS1{F0: -258, F1: -87, F2: -32768}
                        v22 = 756
                        _ = fi41
                        _ = v42
                    }
                    fmt.Println(fi39, v22)
                    _ = v40
                } else {
                    v3 = (fi39 < (-(fi39)))
                    v43 := GenS2{F0: 878, F1: 4294967295, F2: 2}
                    for fi44 := 0; fi44 < 16; fi44++ {
                        v45 := GenS1{F0: -1, F1: 42, F2: 1}
                        var v46 float32 = (1.5 + (1.5 - 1e10))
                        v1 = append(v1, 42)
                        _ = fi44
                        _ = v45
                        _ = v46
                    }
                    _ = v43
                }
                v3 = false
                _ = fi39
            }
        } else {
            v2["a"] = 1
            if (1 > 128) {
                var v47 uint16 = 0
                _ = v47
            }
            var v48 bool = (-855 != -1)
            _ = v48
        }
        for fi49 := 0; fi49 < 13; fi49++ {
            for fi50 := 0; fi50 < 9; fi50++ {
                v3 = (42 == -3)
                _ = fi50
            }
            _ = fi49
        }
    } else {
        delete(v2, ("\n\t" + ("a" + ("a" + "a"))))
    }
    if true {
        var v51 uint = 0
        _ = v51
    }
    if (!v3) {
        v22 = 617
        v3 = true
    } else {
        if false {
            if v3 {
                v52 := GenS3{F0: 255, F1: "a", F2: -461}
                if v3 {
                    v53 := GenS4{F0: ((0.0 * float32(-2.25)) >= ((((float32(0.0) - (-2.25 / float32(0.1))) + 3.14159) * 3.14159) - 1e10))}
                    v1 = append(v1, 2147483648)
                    var v54 int = 2147483647
                    _ = v53
                    _ = v54
                }
                _ = v52
            } else {
                var v55 int16 = 1000
                for fi56 := 0; fi56 < 12; fi56++ {
                    fmt.Println(v55, v55, v22)
                    var v57 uint = 65535
                    var v58 int = (fi56 / (((fi56 & fi56)) | 1))
                    _ = fi56
                    _ = v57
                    _ = v58
                }
                _ = v55
            }
            for fi59 := 0; fi59 < 16; fi59++ {
                fmt.Println(v3, v3, fi59)
                _ = fi59
            }
            v22 = (v22 / (((v22 - v22)) | 1))
        } else {
            if (4294967295 < (v22 / (((^(v22))) | 1))) {
                v22 = v22
            }
        }
        fmt.Println(v22, v22)
        switch 9223372036854775807 {
        case 126:
            switch "foo bar" {
            case "foo bar":
                var v60 uint8 = 1
                _ = v60
            default:
                var v61 float32 = float32(3.14159)
                _ = v61
            }
        default:
            var v62 bool = v3
            _ = v62
        }
    }
    var v63 int8 = -128
    for fi64 := 0; fi64 < 3; fi64++ {
        v3 = v3
        var v65 int16 = 0
        if (!(v22 < 18446744073709551615)) {
            v66 := GenS4{F0: false}
            delete(v2, "")
            v67 := GenS1{F0: -536, F1: v63, F2: (v65 >> ((1) & 63))}
            _ = v66
            _ = v67
        }
        _ = fi64
        _ = v65
    }
    var v68 int8 = v63
    v22 -= 535
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v63:", v63)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v63
    _ = v68
    _ = v9
}
