package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 int64 }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 int32; F1 int32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int; F1 int }
type GenS6 struct { F0 int; F1 string }
type GenS7 struct { F0 uint64; F1 string; F2 int8 }
type GenS8 struct { F0 bool }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]string)
    var v3 bool = (90 > 42)
    v4 := []uint32{269, 4294967295, 928}
    var v5 float64 = 3.14159
    v6 := GenS1{F0: 572, F1: v3, F2: 1}
    for fi7 := 0; fi7 < 3; fi7++ {
        switch (("a" + "foo bar") + "") {
        case "foo bar":
            if false {
                var v8 uint64 = 0
                v8 = v8
                _ = v8
            } else {
                var v9 float32 = 0.0
                if (v9 >= (v9 - float32(0.0))) {
                    var v10 uint16 = 42
                    var v11 uint64 = 18446744073709551615
                    if v3 {
                        v3 = v3
                    }
                    _ = v10
                    _ = v11
                }
                switch 128 {
                case 32767:
                    for fi12 := 0; fi12 < 3; fi12++ {
                        var v13 int = (fi7 ^ 4294967295)
                        _ = fi12
                        _ = v13
                    }
                }
                _ = v9
            }
        default:
            if true {
                var v14 uint64 = 2
                var v15 uint16 = uint16(v14)
                fmt.Println(v5)
                _ = v14
                _ = v15
            }
        }
        _ = fi7
    }
    var v16 uint16 = 0
    var v17 int64 = 0
    var v18 uint16 = 580
    v19 := GenS1{F0: -2147483648, F1: true, F2: (v17 + v17)}
    var v20 int32 = 1
    for fi21 := 0; fi21 < 9; fi21++ {
        switch (fi21 + (fi21 ^ (fi21 + (fi21 | 4294967295)))) {
        case -2147483648:
            v22 := GenS1{F0: v20, F1: v3, F2: v17}
            _ = v22
        case 127:
            v23 := GenS1{F0: (^(v20)), F1: (v3 || true), F2: (v17 >> ((949) & 63))}
            _ = v23
        default:
            var v24 uint16 = v18
            _ = v24
        }
        switch "foo bar" {
        case "foo bar":
            var v25 string = (("a" + "\n\t") + "hello")
            _ = v25
        case "a":
            for fi26 := 0; fi26 < 13; fi26++ {
                v27 := GenS2{F0: 255, F1: 18446744073709551615}
                _ = fi26
                _ = v27
            }
        default:
            var v28 int8 = 42
            _ = v28
        }
        if (!v3) {
            v1["hello"] = fi21
        }
        _ = fi21
    }
    var v29 float32 = 1e10
    v30 := GenS1{F0: v20, F1: ((float32(0.0) / (v29 - -2.25)) < (float32(float64(float32((v5 / v5)))) + ((v29 / 3.14159) - (float32(((0.0 * float64(float32(v5))) - v5)) / v29)))), F2: -1}
    if v3 {
        if (484 > 42) {
            for fi31 := 0; fi31 < 13; fi31++ {
                if (v18 != (v16 / ((v18) | 1))) {
                    for fi32 := 0; fi32 < 6; fi32++ {
                        var v33 int16 = -32768
                        _ = fi32
                        _ = v33
                    }
                    v3 = v3
                }
                _ = fi31
            }
            if v3 {
                if v3 {
                    fmt.Println(v20)
                    for fi34 := 0; fi34 < 8; fi34++ {
                        var v35 int16 = -32768
                        _ = fi34
                        _ = v35
                    }
                    switch "" {
                    case "":
                        if v3 {
                            v18 = v16
                            if v3 {
                                if ((-194 < 42) && true) {
                                    if (!(342 <= 963)) {
                                        for fi36 := 0; fi36 < 16; fi36++ {
                                            for fi37 := 0; fi37 < 2; fi37++ {
                                                v5 -= (float64(0.0) + v5)
                                                switch "hello" {
                                                case "a":
                                                    var v38 uint8 = 1
                                                    _ = v38
                                                case "":
                                                    switch 128 {
                                                    case 128:
                                                        _ = len(v1)
                                                    }
                                                case "foo bar":
                                                    var v39 float64 = v5
                                                    _ = v39
                                                default:
                                                    v40 := v2["a"]
                                                    _ = v40
                                                }
                                                fmt.Println(v5, v3, fi36)
                                                _ = fi37
                                            }
                                            v41 := GenS3{F0: -1, F1: -2147483648}
                                            var v42 uint16 = 65535
                                            _ = fi36
                                            _ = v41
                                            _ = v42
                                        }
                                        v17 -= (v17 & (^(v17)))
                                        var v43 uint64 = 42
                                        _ = v43
                                    }
                                }
                            }
                        }
                    }
                }
                fmt.Println(v20, v17, v17)
                v4 = append(v4, 2147483648)
            } else {
                switch "hello" {
                case "hello":
                    for fi44 := 0; fi44 < 2; fi44++ {
                        var v45 int = (fi44 >> ((255) & 63))
                        var v46 int32 = v20
                        for fi47 := 0; fi47 < 7; fi47++ {
                            if v3 {
                                var v48 int16 = 634
                                if true {
                                    var v49 int8 = 1
                                    v2[""] = "foo bar"
                                    var v50 bool = true
                                    _ = v49
                                    _ = v50
                                }
                                v46 -= v46
                                _ = len(v4)
                                _ = v48
                            } else {
                                var v51 bool = ((true && (true || (v3 && (v3 && (!(!true)))))) || (!(-9223372036854775808 < (fi47 ^ 4294967295))))
                                v2["hello"] = "foo bar"
                                var v52 float64 = ((0.0 / (v5 * ((3.14159 + v5) * ((0.0 + (1.5 + ((1.5 + 0.1) + v5))) * ((1e10 / v5) / 1e10))))) + 3.14159)
                                v51 = v3
                                v53 := GenS2{F0: uint8(v18), F1: 2}
                                _ = v51
                                _ = v52
                                _ = v53
                            }
                            v54 := GenS4{F0: ("foo bar" + ("a" + "a"))}
                            if false {
                                delete(v1, ("foo bar" + "a"))
                            } else {
                                v55 := GenS1{F0: 0, F1: v3, F2: -526}
                                _ = v55
                            }
                            _ = fi47
                            _ = v54
                        }
                        _ = fi44
                        _ = v45
                        _ = v46
                    }
                case "foo bar":
                    for fi56 := 0; fi56 < 14; fi56++ {
                        fmt.Println(v29, v16, v18)
                        fmt.Println(v20, v20, v20)
                        _ = fi56
                    }
                default:
                    v57 := GenS5{F0: 4294967295, F1: 1}
                    _ = v57
                }
                for fi58 := 0; fi58 < 15; fi58++ {
                    for fi59 := 0; fi59 < 5; fi59++ {
                        v16 = 3
                        fmt.Println(v16)
                        if len(v4) > 0 {
                            v4[0] = 64
                        }
                        _ = fi59
                    }
                    _ = fi58
                }
            }
            v60 := GenS6{F0: 0, F1: "a"}
            _ = v60
        }
    }
    var v61 uint = 65535
    switch -9223372036854775808 {
    case 128:
        fmt.Println(v20)
    case 0:
        if false {
            if true {
                switch "a" {
                case "foo bar":
                    if len(v4) > 0 {
                        v4[0] = 64
                    }
                case "":
                    v62 := GenS6{F0: -2147483648, F1: ""}
                    _ = v62
                }
            } else {
                for fi63 := 0; fi63 < 10; fi63++ {
                    if (v3 || (true && (v3 || v3))) {
                        fmt.Println(v61)
                        v61 -= v61
                        v64 := GenS5{F0: (-(fi63)), F1: (fi63 % ((fi63) | 1))}
                        _ = v64
                    } else {
                        v65 := GenS4{F0: "foo bar"}
                        switch "a" {
                        case "":
                            fmt.Println(v20, v29, v3)
                        default:
                            if true {
                                var v66 int8 = 0
                                _ = v66
                            }
                        }
                        _ = v65
                    }
                    _ = fi63
                }
                switch ("a" + "") {
                case "foo bar":
                    var v67 uint16 = 32767
                    _ = v67
                case "":
                    v68 := GenS7{F0: 4294967295, F1: "\n\t", F2: 100}
                    _ = v68
                case "hello":
                    if (0 > 42) {
                        var v69 int = 32768
                        for fi70 := 0; fi70 < 8; fi70++ {
                            var v71 uint = v61
                            v71 = v71
                            _ = fi70
                            _ = v71
                        }
                        var v72 int = (v69 % ((v69) | 1))
                        _ = v69
                        _ = v72
                    }
                default:
                    for fi73 := 0; fi73 < 5; fi73++ {
                        for fi74 := 0; fi74 < 3; fi74++ {
                            v75 := GenS1{F0: v20, F1: (!v3), F2: -1}
                            var v76 uint = 4294967295
                            _ = fi74
                            _ = v75
                            _ = v76
                        }
                        _ = fi73
                    }
                }
                v18 = v16
            }
            fmt.Println(v61)
        }
    default:
        if (!v3) {
            var v77 float64 = v5
            var v78 string = (("hello" + ("hello" + "")) + "hello")
            fmt.Println(v17)
            _ = v77
            _ = v78
        }
    }
    var v79 uint = v61
    if v3 {
        var v80 int16 = -1
        v81 := GenS3{F0: v20, F1: (v20 ^ v20)}
        _ = v80
        _ = v81
    } else {
        v29 = -2.25
        v3 = (42 != 18446744073709551615)
        for fi82 := 0; fi82 < 13; fi82++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v83 := GenS8{F0: false}
            switch (fi82 % (((fi82 + (fi82 ^ (^(fi82))))) | 1)) {
            case 696:
                for fi84 := 0; fi84 < 20; fi84++ {
                    var v85 uint16 = 751
                    var v86 string = ""
                    _ = fi84
                    _ = v85
                    _ = v86
                }
            default:
                v18 = v18
            }
            _ = fi82
            _ = v83
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v61:", v61)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v61
    _ = v79
}
