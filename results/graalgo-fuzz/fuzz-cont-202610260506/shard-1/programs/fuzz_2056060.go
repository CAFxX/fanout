package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 uint16 }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 int; F1 string }
type GenS4 struct { F0 uint64; F1 string }
type GenS5 struct { F0 bool; F1 string }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint16; F1 uint; F2 uint64 }
type GenS8 struct { F0 int16; F1 uint16; F2 int16 }
func genFn0(p0 uint32, p1 uint, p2 string) uint32 {
    var v4 int8 = -128
    fmt.Println(p1, p0)
    for fi5 := 0; fi5 < 15; fi5++ {
        if (0 > 42) {
            v6 := GenS1{F0: 784, F1: 255, F2: 32767}
            if (0 >= uint64(p0)) {
                if true {
                    p2 = "a"
                    var v7 float64 = 0.0
                    for fi8 := 0; fi8 < 9; fi8++ {
                        var v9 int8 = (^(v4))
                        var v10 uint64 = 18446744073709551615
                        v11 := []int{(fi8 + fi5), (fi8 - (-(fi5)))}
                        var v12 float64 = 1.5
                        _ = fi8
                        _ = v10
                        _ = v11
                        _ = v12
                        _ = v9
                    }
                    _ = v7
                }
            } else {
                if true {
                    v13 := []int{fi5, (fi5 / ((fi5) | 1)), -1}
                    fmt.Println(p0, p1)
                    _ = v13
                } else {
                    p1 -= (^(p1))
                    var v14 uint16 = 255
                    _ = v14
                }
            }
            p1 = 65535
            _ = v6
        } else {
            fmt.Println(p0)
            switch p2 {
            case "\n\t":
                p2 = "\n\t"
            default:
                if false {
                    var v15 int8 = v4
                    v4 += (^(v15))
                    var v16 bool = true
                    _ = v15
                    _ = v16
                }
            }
        }
        switch p2 {
        case "foo bar":
            for fi17 := 0; fi17 < 8; fi17++ {
                fmt.Println(p0, p2, fi5)
                var v18 uint64 = 0
                switch p2 {
                case "\n\t":
                    switch fi5 {
                    case -2147483648:
                        switch p2 {
                        case "a":
                            switch (((p2 + (p2 + (p2 + p2))) + ((p2 + ((((p2 + "foo bar") + (("hello" + (p2 + "foo bar")) + (p2 + (p2 + "foo bar")))) + p2) + p2)) + ((("" + "foo bar") + ("hello" + (p2 + p2))) + p2))) + (p2 + ((p2 + ("hello" + p2)) + "hello"))) {
                            case "\n\t":
                                v4 -= v4
                            }
                        default:
                            v19 := GenS2{F0: (fi5 ^ -639), F1: v4}
                            _ = v19
                        }
                    case 2147483647:
                        if (42 > 0) {
                            var v20 uint = 1
                            v20 += (v20 / ((65535) | 1))
                            _ = v20
                        }
                    default:
                        var v21 uint8 = uint8(p0)
                        _ = v21
                    }
                default:
                    v22 := []int{fi17, 476, fi17, 127}
                    _ = v22
                }
                _ = fi17
                _ = v18
            }
        case "a":
            v23 := []int{fi5}
            _ = v23
        case "\n\t":
            for fi24 := 0; fi24 < 11; fi24++ {
                var v25 int = fi24
                v26 := GenS1{F0: 4294967295, F1: 229, F2: 356}
                _ = fi24
                _ = v25
                _ = v26
            }
        default:
            switch fi5 {
            case 668:
                for fi27 := 0; fi27 < 11; fi27++ {
                    switch int(p1) {
                    case 581:
                        var v28 int16 = -779
                        _ = v28
                    case 127:
                        var v29 float32 = (float32(1e10) - float32((0.0 * 1e10)))
                        v30 := GenS3{F0: (fi5 / (((fi27 ^ (fi27 * (fi5 >> (((p1 << (((^(p1))) & 63))) & 63))))) | 1)), F1: (p2 + "foo bar")}
                        _ = v29
                        _ = v30
                    case 128:
                        v31 := GenS2{F0: fi5, F1: (v4 / ((-128) | 1))}
                        var v32 uint8 = 127
                        fmt.Println(p1, fi5)
                        _ = v31
                        _ = v32
                    }
                    _ = fi27
                }
            case -48:
                var v33 uint8 = 155
                _ = v33
            }
        }
        switch p2 {
        case "foo bar":
            v34 := GenS1{F0: 452, F1: 0, F2: 65535}
            _ = v34
        default:
            var v35 uint32 = (^(p0))
            _ = v35
        }
        _ = fi5
    }
    if (!(2147483647 > 78)) {
        var v36 uint32 = p0
        p1 += p1
        _ = v36
    } else {
        v37 := []int{2147483647}
        if (4294967295 <= 4294967295) {
            v38 := GenS4{F0: 1, F1: p2}
            v4 -= 95
            _ = v38
        }
        _ = v37
    }
    var v39 int64 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v39
    _ = v4
    return (^(p0))
}
func genFn1() int8 {
    v40 := GenS5{F0: false, F1: "\n\t"}
    var v41 uint32 = 1
    _ = v40
    _ = v41
    return 127
}

func main() {
    v1 := make(map[int]int)
    var v2 string = (("a" + "") + "\n\t")
    var v3 int32 = 0
    var v42 int32 = (v3 | v3)
    v43 := []int{-406, -1, 9223372036854775807}
    fmt.Println(v3, v42, v3)
    v2 = v2
    switch v2 {
    case "\n\t":
        v2 = v2
    default:
        fmt.Println(v3)
    }
    switch (v2 + "a") {
    case "hello":
        v2 = ""
    case "":
        if (!(!true)) {
            if (!false) {
                v42 = v3
            } else {
                v42 += (v42 - (v3 | v42))
                if true {
                    v44 := GenS3{F0: 128, F1: "\n\t"}
                    v2 = "a"
                    _ = v44
                }
                fmt.Println(v2, v2, v2)
            }
            for fi45 := 0; fi45 < 1; fi45++ {
                var v46 uint = 2
                _ = fi45
                _ = v46
            }
        } else {
            if true {
                switch v2 {
                case "hello":
                    for fi47 := 0; fi47 < 6; fi47++ {
                        v3 -= v42
                        _ = fi47
                    }
                case "foo bar":
                    var v48 int8 = 1
                    _ = v48
                default:
                    var v49 float64 = 3.14159
                    _ = v49
                }
                var v50 float64 = (((1e10 + (0.0 * 1.5)) * 3.14159) / ((float64(0.1) / -2.25) + 3.14159))
                if (v50 == v50) {
                    v51 := GenS3{F0: 2147483647, F1: "\n\t"}
                    _ = v51
                } else {
                    v52 := GenS6{F0: 42}
                    v1[-2147483648] = -614
                    v53 := GenS4{F0: 1, F1: (v2 + "\n\t")}
                    _ = v52
                    _ = v53
                }
                _ = v50
            } else {
                var v54 string = v2
                _ = v54
            }
            for fi55 := 0; fi55 < 13; fi55++ {
                var v56 string = ((("a" + (v2 + "")) + (v2 + "foo bar")) + v2)
                var v57 uint64 = 18446744073709551615
                _ = fi55
                _ = v56
                _ = v57
            }
            var v58 float64 = float64((1e10 * 1.5))
            _ = v58
        }
    default:
        v2 = (((((v2 + ((v2 + v2) + ("hello" + ""))) + v2) + v2) + ("hello" + "hello")) + ("a" + v2))
    }
    _ = len(v43)
    switch ("foo bar" + (v2 + ((v2 + v2) + ""))) {
    case "a":
        switch 127 {
        case 32767:
            v3 = v3
        }
    case "foo bar":
        v59 := GenS4{F0: 764, F1: ((((v2 + (v2 + (v2 + "hello"))) + v2) + "\n\t") + v2)}
        _ = v59
    }
    v43 = append(v43, 46)
    for fi60 := 0; fi60 < 4; fi60++ {
        if len(v43) > 0 {
            v43[0] = 32768
        }
        v42 = 1
        v42 += (v3 ^ v42)
        _ = fi60
    }
    switch ((v2 + ((v2 + v2) + v2)) + v2) {
    case "":
        v1[52] = 792
    case "foo bar":
        v2 = "\n\t"
    default:
        switch (v2 + v2) {
        case "hello":
            if true {
                var v61 float64 = 1.5
                for fi62 := 0; fi62 < 10; fi62++ {
                    for fi63 := 0; fi63 < 20; fi63++ {
                        switch (fi63 & -185) {
                        case 2147483647:
                            fmt.Println(fi63)
                        case 42:
                            v61 -= ((float64((1e10 * (0.1 / -2.25))) + (v61 * v61)) + v61)
                        default:
                            var v64 string = v2
                            fmt.Println(fi63, v61, v42)
                            _ = v64
                        }
                        var v65 int8 = -97
                        v43 = append(v43, fi62)
                        _ = fi63
                        _ = v65
                    }
                    _ = fi62
                }
                _ = v61
            }
        case "foo bar":
            if (65535 <= 42) {
                if (0 < 0) {
                    if (1.5 < (0.0 - 1e10)) {
                        if true {
                            v3 += v42
                            switch (((v2 + (("" + v2) + (v2 + (v2 + "")))) + (("hello" + v2) + ((v2 + v2) + ""))) + v2) {
                            case "":
                                v42 = v42
                            case "\n\t":
                                if false {
                                    v2 = ""
                                    if (1 > genFn1()) {
                                        switch 42 {
                                        case -707:
                                            var v66 int16 = 949
                                            _ = v66
                                        case 4294967295:
                                            var v67 uint32 = 1
                                            _ = v67
                                        case 2147483647:
                                            switch 1 {
                                            case 170:
                                                for fi68 := 0; fi68 < 4; fi68++ {
                                                    var v69 float32 = ((-2.25 / 1.0) + (0.0 / -2.25))
                                                    _ = fi68
                                                    _ = v69
                                                }
                                            case 4294967295:
                                                _ = len(v1)
                                            }
                                        default:
                                            var v70 uint8 = 42
                                            _ = v70
                                        }
                                    }
                                } else {
                                    var v71 float64 = (((((-2.25 * (float64(-2.25) - 0.1)) - (1e10 * 0.0)) - (0.0 / 1e10)) + (float64((float32(0.0) * 0.1)) / -2.25)) + 1.5)
                                    for fi72 := 0; fi72 < 11; fi72++ {
                                        fmt.Println(v3, v71, v2)
                                        var v73 int64 = 9223372036854775807
                                        _ = fi72
                                        _ = v73
                                    }
                                    _ = v71
                                }
                            }
                        } else {
                            v42 = -2147483648
                        }
                        if len(v43) >= 2 {
                            v43 = v43[1:len(v43)]
                        }
                        v2 = "\n\t"
                    } else {
                        for fi74 := 0; fi74 < 11; fi74++ {
                            delete(v1, (fi74 ^ (fi74 - 127)))
                            var v75 uint = 42
                            _ = fi74
                            _ = v75
                        }
                        var v76 int32 = (v42 & v42)
                        for fi77 := 0; fi77 < 15; fi77++ {
                            var v78 uint32 = 393
                            v2 = ((("\n\t" + "hello") + v2) + (v2 + (v2 + "foo bar")))
                            v76 = (v42 * (v76 | -2147483648))
                            _ = fi77
                            _ = v78
                        }
                        _ = v76
                    }
                    v79 := GenS1{F0: 18446744073709551615, F1: 137, F2: 593}
                    var v80 float64 = 3.14159
                    _ = v79
                    _ = v80
                }
            }
        case "":
            _ = len(v43)
        default:
            v3 += -1
        }
    }
    for fi81 := 0; fi81 < 13; fi81++ {
        v42 = (v3 >> ((42) & 63))
        fmt.Println(fi81)
        switch fi81 {
        case 127:
            if true {
                v2 = (v2 + "")
                var v82 uint8 = 255
                _ = len(v1)
                _ = v82
            } else {
                fmt.Println(v3, fi81, v2)
                if false {
                    fmt.Println(v3, v42)
                    if len(v43) > 0 {
                        v43[0] = (fi81 + 9223372036854775807)
                    }
                    v2 = v2
                } else {
                    var v83 uint8 = 97
                    v83 = (v83 / ((42) | 1))
                    v84 := GenS7{F0: 0, F1: 255, F2: 300}
                    _ = v83
                    _ = v84
                }
                if true {
                    switch fi81 {
                    case 32768:
                        v85 := GenS8{F0: -32768, F1: 1, F2: 1}
                        _ = v85
                    case 42:
                        for fi86 := 0; fi86 < 7; fi86++ {
                            switch fi86 {
                            case 32767:
                                var v87 string = "a"
                                _ = v87
                            }
                            _ = fi86
                        }
                    }
                    var v88 uint = 4294967295
                    _ = v88
                }
            }
        case -194:
            for fi89 := 0; fi89 < 8; fi89++ {
                if (18446744073709551615 != 123) {
                    fmt.Println(fi81, v2)
                    fmt.Println(v2, v42, fi89)
                    var v90 float32 = 0.1
                    _ = v90
                } else {
                    var v91 uint = 2
                    if len(v43) >= 2 {
                        v43 = v43[1:len(v43)]
                    }
                    v42 -= v42
                    fmt.Println(fi89, fi89, v42)
                    v92 := GenS8{F0: 1000, F1: 32767, F2: -32768}
                    _ = v91
                    _ = v92
                }
                v1[fi81] = 127
                v2 = ("" + v2)
                _ = fi89
            }
        }
        _ = fi81
    }
    v43 = append(v43, -316)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v43:", len(v43), v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v42
    _ = v43
}
