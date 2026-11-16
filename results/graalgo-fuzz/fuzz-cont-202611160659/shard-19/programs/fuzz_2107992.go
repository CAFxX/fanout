package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 int16; F1 uint16 }
type GenS4 struct { F0 uint8; F1 int32 }
type GenS5 struct { F0 int16 }
func genFn0(p0 uint) int16 {
    p0 = (^(p0))
    v6 := []int{42, -1, -634, 42}
    if ("\n\t" == "a") {
        switch "hello" {
        case "foo bar":
            v6 = append(v6, -414)
        case "":
            for fi7 := 0; fi7 < 7; fi7++ {
                if true {
                    if (!((943 > 1000) || (float64((1e10 + float32(0.1))) != 1.5))) {
                        p0 += (^(p0))
                        v8 := GenS1{F0: 835}
                        var v9 int32 = 1
                        _ = v8
                        _ = v9
                    } else {
                        var v10 int = (fi7 % ((fi7) | 1))
                        var v11 uint32 = 914
                        _ = v10
                        _ = v11
                    }
                    var v12 uint16 = 1
                    _ = v12
                }
                for fi13 := 0; fi13 < 18; fi13++ {
                    var v14 int8 = -128
                    switch "a" {
                    case "hello":
                        var v15 int16 = 84
                        _ = v15
                    case "":
                        p0 = p0
                    case "foo bar":
                        var v16 uint16 = 636
                        _ = v16
                    }
                    _ = fi13
                    _ = v14
                }
                _ = fi7
            }
        default:
            switch 32767 {
            case -9223372036854775808:
                var v17 int = 32767
                _ = v17
            case 2147483647:
                switch "hello" {
                case "hello":
                    switch -23 {
                    case 32767:
                        for fi18 := 0; fi18 < 11; fi18++ {
                            var v19 int64 = 9223372036854775807
                            if true {
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                v19 = (v19 + int64(fi18))
                                var v20 int64 = v19
                                _ = v20
                            }
                            _ = fi18
                            _ = v19
                        }
                    case -9223372036854775808:
                        switch ("" + "hello") {
                        case "foo bar":
                            var v21 bool = (255 > 2)
                            _ = v21
                        default:
                            var v22 int32 = 42
                            _ = v22
                        }
                    }
                case "\n\t":
                    for fi23 := 0; fi23 < 3; fi23++ {
                        if ((false && false) && false) {
                            _ = len(v6)
                            p0 = 4294967295
                        }
                        _ = fi23
                    }
                default:
                    var v24 float64 = (3.14159 - -2.25)
                    _ = v24
                }
            case -376:
                v25 := GenS1{F0: 605}
                _ = v25
            default:
                for fi26 := 0; fi26 < 10; fi26++ {
                    p0 -= (p0 | p0)
                    _ = fi26
                }
            }
        }
    } else {
        if (2 <= 55) {
            v27 := GenS2{F0: 1, F1: 2}
            var v28 uint16 = 664
            for fi29 := 0; fi29 < 9; fi29++ {
                for fi30 := 0; fi30 < 17; fi30++ {
                    v28 = v28
                    if len(v6) > 0 {
                        v6[0] = -991
                    }
                    _ = fi30
                }
                _ = fi29
            }
            _ = v27
            _ = v28
        } else {
            if true {
                var v31 int8 = 15
                p0 = p0
                _ = v31
            }
        }
        if true {
            switch 9223372036854775807 {
            case -933:
                p0 = p0
            case 127:
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            }
            fmt.Println(p0)
        }
    }
    _ = p0
    _ = v6
    return -1
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 int16 = 753
    v3 := []string{}
    v4 := make(map[string]int)
    v5 := make(map[uint32]int)
    v32 := v4[("hello" + "foo bar")]
    v33 := GenS3{F0: (^(v2)), F1: 255}
    for fi34 := 0; fi34 < 6; fi34++ {
        var v35 string = (("hello" + "\n\t") + ("" + "a"))
        _ = fi34
        _ = v35
    }
    v36 := v5[42]
    for fi37 := 0; fi37 < 19; fi37++ {
        var v38 uint8 = 42
        switch "\n\t" {
        case "\n\t":
            if false {
                switch v36 {
                case 603:
                    for fi39 := 0; fi39 < 14; fi39++ {
                        if (!(1 != 42)) {
                            v40 := GenS4{F0: 255, F1: 1}
                            v3 = append(v3, "")
                            _ = v40
                        }
                        _ = len(v5)
                        var v41 int64 = 0
                        _ = fi39
                        _ = v41
                    }
                case -872:
                    v38 = 0
                case 42:
                    switch ((((((("" + (("" + ("\n\t" + "\n\t")) + "a")) + "a") + ("" + (("foo bar" + "\n\t") + ((("a" + "") + ("\n\t" + ("foo bar" + ("hello" + "\n\t")))) + "\n\t")))) + "a") + (("" + ("\n\t" + "hello")) + ("\n\t" + ""))) + ("a" + "a")) + ((("" + "\n\t") + (((("" + "") + "\n\t") + ("a" + "\n\t")) + "hello")) + "a")) {
                    case "a":
                        fmt.Println(v2, fi37, v38)
                    case "\n\t":
                        if false {
                            var v42 int16 = (v2 << ((2) & 63))
                            if true {
                                v43 := GenS5{F0: (v42 >> ((0) & 63))}
                                for fi44 := 0; fi44 < 13; fi44++ {
                                    switch fi44 {
                                    case 42:
                                        switch "" {
                                        case "a":
                                            v3 = append(v3, "\n\t")
                                        case "\n\t":
                                            fmt.Println(v2, v38, v2)
                                        default:
                                            var v45 int64 = (v1 - (v1 ^ -490))
                                            _ = v45
                                        }
                                    default:
                                        switch v32 {
                                        case 32767:
                                            var v46 int32 = 990
                                            _ = v46
                                        case 42:
                                            switch (fi37 << ((2) & 63)) {
                                            case 4294967295:
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                            case 127:
                                                var v47 uint32 = 65535
                                                _ = v47
                                            case -850:
                                                fmt.Println(fi44, v38, v1)
                                            }
                                        }
                                    }
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    v5[42] = 128
                                    v36 += v32
                                    _ = len(v5)
                                    v48 := v4[("foo bar" + "")]
                                    v38 = (v38 * v38)
                                    _ = fi44
                                    _ = v48
                                }
                                v42 = int16(v36)
                                _ = v43
                            } else {
                                var v49 uint8 = v38
                                switch (v36 | 720) {
                                case 48:
                                    v5[0] = fi37
                                case 2147483647:
                                    var v50 string = ("" + "")
                                    _ = v50
                                default:
                                    var v51 int32 = -2147483648
                                    _ = v51
                                }
                                v32 = int(v2)
                                _ = v49
                            }
                            var v52 uint32 = 2147483647
                            _ = v42
                            _ = v52
                        } else {
                            if true {
                                var v53 uint8 = 37
                                v38 += v38
                                if false {
                                    v38 += 50
                                }
                                _ = v53
                            }
                        }
                    }
                default:
                    switch "a" {
                    case "foo bar":
                        if false {
                            for fi54 := 0; fi54 < 4; fi54++ {
                                var v55 float64 = float64(-2.25)
                                fmt.Println(v1)
                                var v56 bool = (v1 == (v1 | v1))
                                _ = fi54
                                _ = v55
                                _ = v56
                            }
                            _ = len(v3)
                            if false {
                                var v57 uint8 = v38
                                _ = v57
                            }
                        }
                    case "":
                        for fi58 := 0; fi58 < 3; fi58++ {
                            v59 := GenS2{F0: 778, F1: 1}
                            _ = fi58
                            _ = v59
                        }
                    default:
                        v38 += 1
                    }
                }
            }
        case "":
            var v60 uint64 = 783
            _ = v60
        case "hello":
            if false {
                _ = len(v3)
                switch "hello" {
                case "foo bar":
                    v1 -= -587
                case "\n\t":
                    var v61 uint32 = 490
                    _ = v61
                default:
                    fmt.Println(v2)
                }
            } else {
                fmt.Println(v1, fi37, v38)
                switch "\n\t" {
                case "hello":
                    v3 = append(v3, "")
                }
            }
        default:
            var v62 uint = 18446744073709551615
            _ = v62
        }
        if false {
            var v63 int = (v36 & 4294967295)
            v63 -= v63
            _ = v63
        }
        _ = fi37
        _ = v38
    }
    var v64 int32 = 2147483647
    _ = len(v5)
    v65 := GenS4{F0: 128, F1: -160}
    var v66 int8 = 42
    if (true || (false || (110 != (-(v66))))) {
        delete(v5, 42)
    } else {
        v2 -= v2
        for fi67 := 0; fi67 < 9; fi67++ {
            switch (-(fi67)) {
            case 32768:
                switch "a" {
                case "hello":
                    v68 := v5[0]
                    _ = v68
                case "\n\t":
                    for fi69 := 0; fi69 < 8; fi69++ {
                        v70 := GenS3{F0: 0, F1: 32767}
                        var v71 int = v36
                        _ = fi69
                        _ = v70
                        _ = v71
                    }
                }
            case 1:
                fmt.Println(v66)
            }
            _ = fi67
        }
        v66 = -16
    }
    for fi72 := 0; fi72 < 5; fi72++ {
        v4["a"] = (-(fi72))
        _ = fi72
    }
    switch (("\n\t" + "") + "\n\t") {
    case "a":
        var v73 float32 = 1.5
        _ = v73
    case "foo bar":
        fmt.Println(v2, v66)
    case "hello":
        switch "\n\t" {
        case "foo bar":
            var v74 int64 = 1
            _ = v74
        case "a":
            var v75 int16 = -32768
            _ = v75
        default:
            v66 += -1
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v64:", v64)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v36
    _ = v4
    _ = v5
    _ = v64
    _ = v65
    _ = v66
}
