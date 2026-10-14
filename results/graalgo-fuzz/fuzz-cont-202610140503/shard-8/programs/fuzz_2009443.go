package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int8; F1 uint; F2 uint16 }
type GenS4 struct { F0 int16; F1 int32; F2 string }
type GenS5 struct { F0 int8; F1 uint; F2 uint }
type GenS6 struct { F0 int; F1 int8 }
type GenS7 struct { F0 uint32; F1 uint8; F2 string }
type GenS8 struct { F0 int64; F1 int16; F2 uint32 }
func genFn0(p0 uint32, p1 uint, p2 uint16) string {
    switch "" {
    case "foo bar":
        if false {
            p0 += p0
            var v4 int8 = -1
            _ = v4
        }
    case "\n\t":
        p2 += uint16(p0)
    }
    for fi5 := 0; fi5 < 18; fi5++ {
        for fi6 := 0; fi6 < 18; fi6++ {
            var v7 int64 = 615
            var v8 int16 = 32767
            fmt.Println(p2)
            _ = fi6
            _ = v7
            _ = v8
        }
        if true {
            v9 := GenS1{F0: -9223372036854775808, F1: "foo bar", F2: -1}
            _ = v9
        } else {
            for fi10 := 0; fi10 < 10; fi10++ {
                switch (("hello" + ("" + "hello")) + "") {
                case "a":
                    if (1.5 <= float32(0.0)) {
                        var v11 float32 = 1e10
                        p0 -= 42
                        _ = v11
                    }
                    var v12 int = fi5
                    p1 += 262
                    _ = v12
                case "foo bar":
                    fmt.Println(p0)
                case "\n\t":
                    if (false && true) {
                        fmt.Println(p0, fi5)
                        var v13 float32 = 3.14159
                        _ = v13
                    }
                }
                v14 := GenS2{F0: (^(p2))}
                if ((!true) || false) {
                    fmt.Println(fi10)
                    v15 := []int{(fi5 << ((42) & 63)), 4294967295, -9223372036854775808, fi5}
                    _ = v15
                } else {
                    p1 = p1
                    v16 := GenS2{F0: 255}
                    _ = v16
                }
                p1 = p1
                if true {
                    p1 = (p1 >> ((p1) & 63))
                    switch ("foo bar" + "\n\t") {
                    case "\n\t":
                        fmt.Println(fi10, fi5, p2)
                    case "a":
                        p0 = p0
                    default:
                        var v17 int = fi5
                        if true {
                            p0 += 42
                            v18 := []int{351}
                            _ = v18
                        }
                        if (0.1 != 3.14159) {
                            if false {
                                fmt.Println(fi5, fi5, fi10)
                                var v19 uint16 = (p2 - p2)
                                _ = v19
                            }
                            var v20 int64 = 1
                            _ = v20
                        } else {
                            p2 = p2
                            v21 := GenS2{F0: (p2 + p2)}
                            var v22 int16 = 32767
                            _ = v21
                            _ = v22
                        }
                        _ = v17
                    }
                    p1 -= (p1 << ((60) & 63))
                } else {
                    if (33 != -32768) {
                        var v23 int16 = int16(p0)
                        _ = v23
                    } else {
                        p1 = p1
                        switch "" {
                        case "a":
                            v24 := []int{(-(fi10)), (fi10 & 2147483647), fi10}
                            _ = v24
                        case "":
                            var v25 int8 = 100
                            _ = v25
                        default:
                            v26 := GenS3{F0: -105, F1: 0, F2: (p2 ^ 0)}
                            _ = v26
                        }
                    }
                    fmt.Println(p1, p2, fi10)
                    switch fi10 {
                    case 42:
                        if true {
                            p0 = (p0 * (p0 << ((p1) & 63)))
                            fmt.Println(p1, p0)
                            p0 += (^(p0))
                            var v27 uint32 = 618
                            var v28 uint = (p1 & 65535)
                            _ = v27
                            _ = v28
                        } else {
                            switch -352 {
                            case 2147483647:
                                var v29 float32 = -2.25
                                _ = v29
                            default:
                                var v30 float64 = (float64(-2.25) / (((-2.25 / (0.1 - float64(float32(0.0)))) * -2.25) * 0.1))
                                _ = v30
                            }
                            v31 := GenS4{F0: 42, F1: -325, F2: (("a" + "") + "hello")}
                            _ = v31
                        }
                    case 0:
                        if (p0 == (p0 + uint32(p2))) {
                            var v32 float64 = ((float64(0.0) - (float64(3.14159) * 1e10)) - -2.25)
                            var v33 int8 = 100
                            var v34 uint8 = 128
                            v35 := GenS4{F0: 1000, F1: -1, F2: "hello"}
                            _ = v32
                            _ = v33
                            _ = v34
                            _ = v35
                        }
                    default:
                        switch fi5 {
                        case 42:
                            v36 := GenS5{F0: 0, F1: p1, F2: p1}
                            _ = v36
                        default:
                            if true {
                                var v37 int16 = -32768
                                switch "hello" {
                                case "foo bar":
                                    p2 -= p2
                                }
                                _ = v37
                            } else {
                                p1 = p1
                                fmt.Println(fi10, fi10, p1)
                                fmt.Println(fi10, fi10)
                            }
                        }
                    }
                }
                p1 += p1
                _ = fi10
                _ = v14
            }
            var v38 bool = (false || (42 > int64(p1)))
            fmt.Println(p1, fi5, v38)
            _ = v38
        }
        for fi39 := 0; fi39 < 20; fi39++ {
            fmt.Println(p0, p1)
            v40 := GenS4{F0: 1000, F1: -2147483648, F2: ((((("foo bar" + "hello") + "hello") + (("" + "\n\t") + "foo bar")) + "") + "")}
            var v41 int16 = -32768
            _ = fi39
            _ = v40
            _ = v41
        }
        _ = fi5
    }
    _ = p0
    _ = p1
    _ = p2
    return "foo bar"
}
func genFn1() uint8 {
    switch ("foo bar" + "\n\t") {
    case "\n\t":
        var v42 int16 = 1000
        _ = v42
    }
    v43 := []int{679}
    for fi44 := 0; fi44 < 13; fi44++ {
        for fi45 := 0; fi45 < 15; fi45++ {
            v46 := GenS6{F0: -2147483648, F1: 127}
            var v47 bool = (("" + ("a" + "foo bar")) <= ("\n\t" + ("hello" + "")))
            v43 = append(v43, fi45)
            _ = fi45
            _ = v46
            _ = v47
        }
        _ = fi44
    }
    if len(v43) > 0 {
        v43[0] = 9223372036854775807
    }
    _ = v43
    return 127
}

func main() {
    v1 := []int{9223372036854775807}
    v2 := make(map[int]uint32)
    var v3 int32 = -1
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v48 int8 = 1
    v48 += (v48 & (v48 << ((256) & 63)))
    v48 -= v48
    var v49 uint16 = 42
    var v50 int = 2147483647
    _ = len(v2)
    if true {
        switch (v50 ^ 32768) {
        case 332:
            v3 -= -2147483648
        case -641:
            for fi51 := 0; fi51 < 4; fi51++ {
                if (int8(v49) < (v48 & v48)) {
                    v52 := GenS1{F0: 0, F1: "", F2: 32767}
                    if true {
                        if (1 >= (v49 << ((510) & 63))) {
                            v53 := GenS7{F0: 580, F1: 130, F2: "foo bar"}
                            if true {
                                if len(v1) > 0 {
                                    v1[0] = (fi51 & (v50 - fi51))
                                }
                                switch "foo bar" {
                                case "":
                                    v49 = v49
                                case "hello":
                                    v54 := GenS8{F0: 9223372036854775807, F1: -1, F2: 65535}
                                    _ = v54
                                case "foo bar":
                                    v2[(-(v50))] = 65535
                                default:
                                    fmt.Println(v48, v49)
                                }
                                v48 = (v48 % (((v48 + v48)) | 1))
                            } else {
                                var v55 uint32 = 662
                                _ = v55
                            }
                            v3 -= v3
                            _ = v53
                        }
                    } else {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        if false {
                            fmt.Println(v50, v3)
                            v56 := GenS8{F0: -1, F1: -1, F2: 1}
                            _ = v56
                        } else {
                            var v57 int32 = v3
                            for fi58 := 0; fi58 < 9; fi58++ {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                fmt.Println(v3, fi51, v48)
                                if (-1 != 0) {
                                    v3 = v3
                                    v3 -= 42
                                } else {
                                    fmt.Println(v57)
                                }
                                _ = fi58
                            }
                            var v59 uint32 = 0
                            _ = v57
                            _ = v59
                        }
                        var v60 uint32 = 65535
                        _ = v60
                    }
                    var v61 int16 = 0
                    _ = v52
                    _ = v61
                }
                _ = fi51
            }
        case 128:
            fmt.Println(v3)
        default:
            for fi62 := 0; fi62 < 13; fi62++ {
                if false {
                    v63 := GenS2{F0: v49}
                    _ = v63
                } else {
                    var v64 int8 = (^(v48))
                    _ = v64
                }
                _ = fi62
            }
        }
        v65 := GenS7{F0: 281, F1: 42, F2: "a"}
        _ = v65
    }
    var v66 float64 = 3.14159
    v67 := v2[1]
    switch (v50 >> ((4294967295) & 63)) {
    case 679:
        v68 := v2[-515]
        _ = v68
    case 32767:
        switch ("\n\t" + ("" + "\n\t")) {
        case "a":
            if len(v1) > 0 {
                v1[0] = (v50 << ((534) & 63))
            }
        case "hello":
            v69 := GenS3{F0: 50, F1: 853, F2: v49}
            _ = v69
        default:
            for fi70 := 0; fi70 < 4; fi70++ {
                fmt.Println(v50)
                var v71 float32 = 0.0
                _ = fi70
                _ = v71
            }
        }
    }
    if false {
        for fi72 := 0; fi72 < 16; fi72++ {
            v2[(fi72 % ((fi72) | 1))] = 2147483648
            if true {
                v66 += 1e10
                var v73 uint32 = 1
                v2[v50] = v73
                _ = v73
            }
            _ = fi72
        }
        v48 = (v48 - (v48 >> ((1) & 63)))
        if (9223372036854775807 <= 0) {
            fmt.Println(v49, v67)
        } else {
            v66 += v66
            var v74 uint32 = (v67 - 65535)
            _ = v74
        }
    }
    fmt.Println(v66, v67)
    var v75 float32 = float32((3.14159 + float64((0.1 - (1.5 - (float32(0.0) - float32((v66 / v66))))))))
    fmt.Println(v49)
    if true {
        switch ((("foo bar" + ("\n\t" + "hello")) + (("foo bar" + "\n\t") + "foo bar")) + (("\n\t" + "foo bar") + "a")) {
        case "":
            var v76 float32 = (-2.25 - v75)
            _ = v76
        case "\n\t":
            v66 -= 0.1
        case "hello":
            switch v50 {
            case 292:
                for fi77 := 0; fi77 < 12; fi77++ {
                    var v78 int32 = -1
                    var v79 uint32 = (v67 / (((v67 & (v67 * (^(v67))))) | 1))
                    _ = fi77
                    _ = v78
                    _ = v79
                }
            case 9223372036854775807:
                v48 = v48
            case 4294967295:
                for fi80 := 0; fi80 < 17; fi80++ {
                    if ((0.1 - v75) == v75) {
                        for fi81 := 0; fi81 < 9; fi81++ {
                            fmt.Println(v49, v75, v48)
                            _ = fi81
                        }
                    }
                    v48 += v48
                    v82 := v2[32767]
                    _ = fi80
                    _ = v82
                }
            }
        }
        v49 -= 0
    } else {
        v75 += v75
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v50:", v50)
    fmt.Println("v66:", v66)
    fmt.Println("v67:", v67)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v2
    _ = v3
    _ = v48
    _ = v49
    _ = v50
    _ = v66
    _ = v67
    _ = v75
}
