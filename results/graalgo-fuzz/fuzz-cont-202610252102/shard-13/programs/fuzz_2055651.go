package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 uint64; F1 uint32 }
type GenS3 struct { F0 uint16; F1 string }
type GenS4 struct { F0 uint; F1 int8; F2 int16 }
type GenS5 struct { F0 uint16; F1 int; F2 int }
type GenS6 struct { F0 int32; F1 int64; F2 uint }
type GenS7 struct { F0 int }
func genFn0(p0 string, p1 uint64, p2 uint) int32 {
    v7 := []int{127, 128}
    var v8 uint64 = (p1 + 18446744073709551615)
    for fi9 := 0; fi9 < 14; fi9++ {
        for fi10 := 0; fi10 < 14; fi10++ {
            var v11 string = (p0 + (p0 + "a"))
            v12 := GenS1{F0: 873, F1: 801}
            _ = fi10
            _ = v11
            _ = v12
        }
        fmt.Println(p1, p1, p0)
        v8 = (v8 >> (((^(p2))) & 63))
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return -1
}

func main() {
    v1 := []string{("a" + "\n\t"), ("a" + "")}
    var v2 int8 = -1
    var v3 int16 = 0
    var v4 int = 32767
    var v5 int = (^(v4))
    var v6 bool = (!true)
    for fi13 := 0; fi13 < 9; fi13++ {
        switch 32768 {
        case -9223372036854775808:
            switch v5 {
            case 4294967295:
                var v14 uint = 18446744073709551615
                _ = v14
            case -1:
                _ = len(v1)
            default:
                var v15 float32 = ((((0.1 - (0.1 + 1e10)) / 1.0) * 3.14159) * 0.0)
                _ = v15
            }
        }
        v16 := GenS1{F0: 2, F1: 0}
        _ = fi13
        _ = v16
    }
    fmt.Println(v2, v4)
    switch (v5 + (v5 % ((v4) | 1))) {
    case 32768:
        v2 += (v2 | v2)
    }
    switch (v5 >> ((4294967295) & 63)) {
    case 1:
        v2 -= (v2 ^ (v2 / ((v2) | 1)))
    case 2147483647:
        for fi17 := 0; fi17 < 12; fi17++ {
            for fi18 := 0; fi18 < 7; fi18++ {
                if ((!v6) && v6) {
                    var v19 uint32 = 4294967295
                    v3 += (v3 | v3)
                    if (42 >= -1) {
                        _ = len(v1)
                    }
                    v20 := GenS1{F0: 61, F1: (v19 - (v19 | 748))}
                    if len(v1) > 0 {
                        v1[0] = "a"
                    }
                    _ = v19
                    _ = v20
                } else {
                    v21 := GenS1{F0: 42, F1: 42}
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = v21
                }
                v22 := GenS1{F0: 133, F1: 2147483648}
                if v6 {
                    var v23 int32 = 906
                    _ = len(v1)
                    v6 = (2 == 65535)
                    v23 -= v23
                    _ = v23
                }
                _ = fi18
                _ = v22
            }
            for fi24 := 0; fi24 < 20; fi24++ {
                if true {
                    switch "" {
                    case "a":
                        if (1 < 128) {
                            if len(v1) > 0 {
                                v1[0] = "\n\t"
                            }
                        }
                    case "":
                        v6 = (!v6)
                    default:
                        if v6 {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                    }
                    v25 := GenS2{F0: 0, F1: 4294967295}
                    switch fi24 {
                    case 972:
                        var v26 uint32 = 752
                        _ = v26
                    default:
                        switch "hello" {
                        case "foo bar":
                            v1 = append(v1, "")
                        case "":
                            fmt.Println(v3)
                        default:
                            _ = len(v1)
                        }
                    }
                    _ = v25
                }
                v3 = 32767
                if v6 {
                    if (v5 > (fi17 << ((595) & 63))) {
                        v4 += fi17
                        v6 = (!v6)
                        fmt.Println(v4)
                        v27 := GenS2{F0: 501, F1: 2147483648}
                        _ = v27
                    } else {
                        v4 += fi17
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        var v28 uint32 = 2147483647
                        var v29 uint64 = 4294967295
                        var v30 float32 = 1e10
                        var v31 int8 = (^(v2))
                        var v32 int16 = (v3 * (v3 + (v3 >> ((65535) & 63))))
                        _ = v28
                        _ = v29
                        _ = v30
                        _ = v31
                        _ = v32
                    }
                    if (197 == -900) {
                        if true {
                            if (-375 != int64(v4)) {
                                var v33 int8 = -128
                                _ = v33
                            } else {
                                if len(v1) > 0 {
                                    v1[0] = "foo bar"
                                }
                                var v34 int16 = (v3 / ((v3) | 1))
                                if true {
                                    switch -215 {
                                    case -2147483648:
                                        var v35 int64 = 470
                                        _ = v35
                                    case 147:
                                        var v36 uint16 = 0
                                        _ = v36
                                    case 128:
                                        var v37 uint8 = 255
                                        _ = v37
                                    default:
                                        var v38 int32 = 42
                                        _ = v38
                                    }
                                }
                                _ = v34
                            }
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            fmt.Println(v2, v4)
                            if (2147483647 > (v4 % (((-(v5))) | 1))) {
                                fmt.Println(v4, v4)
                                if true {
                                    var v39 int64 = -9223372036854775808
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    var v40 int8 = v2
                                    _ = v39
                                    _ = v40
                                } else {
                                    var v41 float32 = 3.14159
                                    if (!true) {
                                        if ((v6 && (!(v6 || v6))) || (false && v6)) {
                                            switch 127 {
                                            case 42:
                                                switch ("\n\t" + ("\n\t" + ("hello" + ""))) {
                                                case "foo bar":
                                                    v42 := GenS1{F0: 291, F1: 469}
                                                    _ = v42
                                                }
                                            case -1:
                                                fmt.Println(v4, v3)
                                            }
                                            v1 = append(v1, ("hello" + "foo bar"))
                                        }
                                    } else {
                                        if len(v1) > 0 {
                                            v1[0] = ("" + ("a" + "\n\t"))
                                        }
                                    }
                                    _ = v41
                                }
                                if true {
                                    switch (v4 + (v4 % ((fi17) | 1))) {
                                    case 32768:
                                        var v43 int16 = 42
                                        _ = v43
                                    default:
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        if v6 {
                                            v5 = fi24
                                            v44 := GenS3{F0: 0, F1: ("\n\t" + "\n\t")}
                                            var v45 float64 = 1.5
                                            v46 := GenS4{F0: 42, F1: (v2 - v2), F2: -32768}
                                            _ = v44
                                            _ = v45
                                            _ = v46
                                        } else {
                                            v47 := GenS1{F0: 255, F1: 2147483647}
                                            var v48 uint32 = 2147483648
                                            v5 = v5
                                            if len(v1) > 0 {
                                                v1[0] = "foo bar"
                                            }
                                            if false {
                                                switch "foo bar" {
                                                case "hello":
                                                    switch "a" {
                                                    case "\n\t":
                                                        v48 = 2
                                                    case "foo bar":
                                                        fmt.Println(v6)
                                                        fmt.Println(v6, v4)
                                                        if true {
                                                            var v49 int32 = 42
                                                            _ = v49
                                                        } else {
                                                            var v50 bool = ((v48 & v48) <= (v48 >> ((2) & 63)))
                                                            v1 = append(v1, "hello")
                                                            _ = v50
                                                        }
                                                        v4 -= v4
                                                        fmt.Println(v2, v6)
                                                    case "a":
                                                        fmt.Println(fi17, fi24, fi17)
                                                    default:
                                                        var v51 int8 = (v2 % (((v2 ^ (v2 + v2))) | 1))
                                                        _ = v51
                                                    }
                                                case "":
                                                    var v52 int16 = v3
                                                    _ = v52
                                                default:
                                                    var v53 int = (fi17 + -669)
                                                    _ = v53
                                                }
                                                v3 = (v3 + (v3 & v3))
                                                var v54 string = "foo bar"
                                                _ = v54
                                            }
                                            _ = v47
                                            _ = v48
                                        }
                                        if len(v1) > 0 {
                                            v1[0] = ("foo bar" + "a")
                                        }
                                        v55 := GenS5{F0: 942, F1: (fi17 % (((fi24 & fi17)) | 1)), F2: (fi17 - (v4 % ((-9223372036854775808) | 1)))}
                                        _ = v55
                                    }
                                    var v56 int = (fi17 % ((int(v2)) | 1))
                                    _ = v56
                                }
                            }
                        }
                        switch (fi17 >> ((2) & 63)) {
                        case -1:
                            v6 = true
                        case -2147483648:
                            v1 = append(v1, "\n\t")
                        }
                        v1 = append(v1, "")
                    } else {
                        v3 = v3
                    }
                    switch "\n\t" {
                    case "foo bar":
                        if len(v1) > 0 {
                            v1[0] = "hello"
                        }
                    }
                    fmt.Println(v5, fi17, v6)
                    v4 = (fi17 << ((42) & 63))
                    if v6 {
                        fmt.Println(v2, v3, fi17)
                        v57 := GenS4{F0: 1, F1: -128, F2: v3}
                        var v58 uint64 = 159
                        fmt.Println(v6, v3)
                        _ = v57
                        _ = v58
                    }
                } else {
                    v59 := GenS6{F0: -364, F1: 42, F2: 18446744073709551615}
                    v5 += 112
                    _ = v59
                }
                _ = fi24
            }
            switch 2147483647 {
            case 42:
                if len(v1) > 0 {
                    v1[0] = "a"
                }
            case 401:
                var v60 string = "hello"
                _ = v60
            case 2147483647:
                var v61 int32 = -2147483648
                _ = v61
            }
            _ = fi17
        }
    default:
        if ((3.14159 / 1.0) >= ((1.5 - 1.5) - float64(0.1))) {
            switch (v4 * (v4 ^ 1)) {
            case 0:
                v2 += (v2 / (((^(v2))) | 1))
            case -2147483648:
                if false {
                    var v62 uint16 = 65535
                    for fi63 := 0; fi63 < 15; fi63++ {
                        v1 = append(v1, "hello")
                        _ = fi63
                    }
                    v6 = true
                    _ = v62
                } else {
                    v64 := GenS1{F0: 32767, F1: 2}
                    _ = v64
                }
            case -596:
                if v6 {
                    var v65 uint16 = 42
                    _ = v65
                }
            }
            switch "" {
            case "foo bar":
                v3 = v3
            case "a":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case "":
                fmt.Println(v3)
            default:
                fmt.Println(v4, v6)
            }
            _ = len(v1)
        }
    }
    v66 := GenS7{F0: (v5 & v5)}
    v67 := GenS1{F0: 2, F1: 1}
    var v68 bool = v6
    v69 := GenS6{F0: -2147483648, F1: -763, F2: 167}
    fmt.Println(v4)
    fmt.Println(v4, v5)
    fmt.Println(v2, v4, v2)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v66
    _ = v67
    _ = v68
    _ = v69
}
