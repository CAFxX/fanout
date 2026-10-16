package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64; F1 uint64; F2 uint32 }
type GenS3 struct { F0 int64; F1 uint32 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int8; F1 uint16 }
type GenS6 struct { F0 uint8; F1 int64; F2 uint16 }
func genFn0(p0 uint) uint16 {
    for fi5 := 0; fi5 < 8; fi5++ {
        p0 += p0
        var v6 uint16 = 0
        fmt.Println(fi5, fi5)
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: "a"}
    var v8 uint64 = 18446744073709551615
    v9 := []int{2147483647, 9223372036854775807}
    v8 = (v8 * v8)
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 208
}

func main() {
    v1 := []string{(("\n\t" + "") + "hello")}
    var v2 int16 = -501
    var v3 bool = false
    v4 := []string{"foo bar", "a", "hello", ("a" + (("" + ("" + "")) + "foo bar"))}
    switch 32768 {
    case 9223372036854775807:
        if v3 {
            var v10 bool = v3
            _ = v10
        } else {
            for fi11 := 0; fi11 < 17; fi11++ {
                v12 := GenS2{F0: 582, F1: 18446744073709551615, F2: 566}
                _ = fi11
                _ = v12
            }
        }
    }
    switch "" {
    case "hello":
        _ = len(v1)
    case "foo bar":
        if (!v3) {
            var v13 uint8 = 1
            var v14 string = "a"
            if true {
                v15 := GenS1{F0: (("" + v14) + v14)}
                v2 -= v2
                var v16 float64 = 3.14159
                _ = v15
                _ = v16
            } else {
                var v17 int32 = -1
                if (565 <= 0) {
                    if ((float32(-2.25) - (-2.25 / 3.14159)) > float32(3.14159)) {
                        for fi18 := 0; fi18 < 3; fi18++ {
                            for fi19 := 0; fi19 < 15; fi19++ {
                                switch int(v13) {
                                case 2147483647:
                                    if v3 {
                                        var v20 bool = (255 >= (v13 % (((v13 / (((v13 + v13)) | 1))) | 1)))
                                        if (0.1 == ((-2.25 / 1.0) / (((-2.25 + float64((0.0 / float32((1e10 * 1.5))))) * 3.14159) + ((float64(0.1) / (0.0 - 1.5)) + ((0.1 - 0.1) * 1e10))))) {
                                            var v21 float64 = float64(1e10)
                                            v4 = append(v4, v14)
                                            v22 := GenS3{F0: 0, F1: 2147483648}
                                            v4 = append(v4, "\n\t")
                                            _ = v21
                                            _ = v22
                                        }
                                        _ = v20
                                    } else {
                                        v23 := GenS1{F0: ("\n\t" + ((("a" + "foo bar") + v14) + (v14 + "foo bar")))}
                                        _ = v23
                                    }
                                }
                                switch v14 {
                                case "":
                                    var v24 bool = (1 < 65535)
                                    _ = v24
                                case "a":
                                    v13 += (v13 - (v13 << ((1) & 63)))
                                case "\n\t":
                                    v14 = "\n\t"
                                }
                                _ = fi19
                            }
                            v25 := GenS4{F0: 1}
                            fmt.Println(v14, v2, fi18)
                            _ = fi18
                            _ = v25
                        }
                        var v26 uint = 128
                        _ = v26
                    }
                    if ((!v3) && v3) {
                        v1 = append(v1, v14)
                        for fi27 := 0; fi27 < 9; fi27++ {
                            var v28 uint16 = 848
                            _ = fi27
                            _ = v28
                        }
                    } else {
                        var v29 uint8 = v13
                        for fi30 := 0; fi30 < 9; fi30++ {
                            var v31 int64 = -9223372036854775808
                            switch v14 {
                            case "\n\t":
                                switch fi30 {
                                case 979:
                                    for fi32 := 0; fi32 < 17; fi32++ {
                                        var v33 int16 = v2
                                        _ = fi32
                                        _ = v33
                                    }
                                case -194:
                                    switch v14 {
                                    case "a":
                                        if (-128 != -94) {
                                            for fi34 := 0; fi34 < 8; fi34++ {
                                                var v35 float64 = (0.0 * 3.14159)
                                                switch v14 {
                                                case "hello":
                                                    v36 := GenS5{F0: 127, F1: 805}
                                                    _ = v36
                                                }
                                                v17 = 0
                                                _ = fi34
                                                _ = v35
                                            }
                                            v37 := GenS5{F0: 0, F1: 2}
                                            for fi38 := 0; fi38 < 11; fi38++ {
                                                switch fi38 {
                                                case -946:
                                                    if v3 {
                                                        switch "hello" {
                                                        case "foo bar":
                                                            if len(v4) > 0 {
                                                                v4[0] = ""
                                                            }
                                                        case "hello":
                                                            v17 = -15
                                                        default:
                                                            fmt.Println(v17, fi38)
                                                        }
                                                        if (v3 || ((v29 + 0) >= (v13 | 42))) {
                                                            v3 = (!false)
                                                            v17 += 2147483647
                                                        } else {
                                                            v39 := GenS4{F0: 1}
                                                            v40 := GenS6{F0: 1, F1: v31, F2: 0}
                                                            _ = v39
                                                            _ = v40
                                                        }
                                                    } else {
                                                        v4 = append(v4, (v14 + v14))
                                                        v31 -= v31
                                                    }
                                                default:
                                                    v41 := GenS2{F0: 42, F1: 4294967295, F2: 976}
                                                    _ = v41
                                                }
                                                fmt.Println(v13, v29, v13)
                                                _ = fi38
                                            }
                                            _ = v37
                                        } else {
                                            if true {
                                                switch v14 {
                                                case "a":
                                                    v2 = v2
                                                case "":
                                                    fmt.Println(v3, v2, v3)
                                                case "hello":
                                                    var v42 float32 = 1e10
                                                    _ = v42
                                                }
                                            } else {
                                                _ = len(v4)
                                                var v43 int16 = v2
                                                fmt.Println(v31, v43)
                                                _ = v43
                                            }
                                        }
                                    case "hello":
                                        var v44 uint64 = 18446744073709551615
                                        _ = v44
                                    case "":
                                        var v45 int = fi30
                                        _ = v45
                                    }
                                case 4294967295:
                                    for fi46 := 0; fi46 < 12; fi46++ {
                                        var v47 uint = 18446744073709551615
                                        _ = fi46
                                        _ = v47
                                    }
                                default:
                                    v13 = v13
                                }
                            default:
                                var v48 float64 = float64(1e10)
                                _ = v48
                            }
                            _ = fi30
                            _ = v31
                        }
                        for fi49 := 0; fi49 < 7; fi49++ {
                            var v50 uint = 1
                            var v51 int = fi49
                            _ = fi49
                            _ = v50
                            _ = v51
                        }
                        _ = v29
                    }
                    if (390 > 18446744073709551615) {
                        switch -1 {
                        case 9223372036854775807:
                            var v52 uint = 238
                            _ = v52
                        }
                    }
                }
                _ = v17
            }
            _ = v13
            _ = v14
        }
    default:
        for fi53 := 0; fi53 < 4; fi53++ {
            v3 = (2147483647 <= -809)
            fmt.Println(fi53)
            var v54 int = fi53
            _ = fi53
            _ = v54
        }
    }
    v1 = append(v1, "foo bar")
    for fi55 := 0; fi55 < 18; fi55++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi55
    }
    var v56 int32 = 0
    v3 = true
    for fi57 := 0; fi57 < 3; fi57++ {
        _ = len(v4)
        _ = fi57
    }
    fmt.Println(v2, v56, v2)
    for fi58 := 0; fi58 < 7; fi58++ {
        v2 += v2
        _ = fi58
    }
    for fi59 := 0; fi59 < 16; fi59++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if (!v3) {
            var v60 uint16 = 477
            _ = v60
        }
        _ = fi59
    }
    for fi61 := 0; fi61 < 11; fi61++ {
        switch "hello" {
        case "":
            var v62 uint32 = 790
            _ = v62
        case "hello":
            for fi63 := 0; fi63 < 13; fi63++ {
                var v64 bool = (1 < 0)
                if false {
                    v65 := GenS1{F0: ("foo bar" + ("foo bar" + "\n\t"))}
                    _ = v65
                } else {
                    var v66 uint = 19
                    _ = v66
                }
                v67 := GenS3{F0: -9223372036854775808, F1: 42}
                if false {
                    v68 := GenS4{F0: 42}
                    v64 = true
                    _ = v68
                }
                var v69 uint16 = 628
                _ = fi63
                _ = v64
                _ = v67
                _ = v69
            }
        }
        _ = fi61
    }
    v56 -= v56
    for fi70 := 0; fi70 < 16; fi70++ {
        v2 = -1
        _ = fi70
    }
    v56 += (-(v56))
    _ = len(v1)
    switch ((("foo bar" + ("foo bar" + "a")) + "\n\t") + "") {
    case "a":
        var v71 uint64 = 1
        _ = v71
    default:
        if len(v1) > 0 {
            v1[0] = (("a" + "") + ("a" + "hello"))
        }
    }
    v2 += (v2 | v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v56
}
