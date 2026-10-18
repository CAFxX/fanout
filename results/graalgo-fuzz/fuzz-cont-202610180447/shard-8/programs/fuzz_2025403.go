package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 uint; F1 int8; F2 uint16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int; F1 int; F2 uint64 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int16; F1 bool }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 uint32; F1 uint64 }
type GenS10 struct { F0 bool; F1 int64; F2 uint32 }
func genFn0(p0 uint32, p1 int8) uint8 {
    v7 := GenS1{F0: -154, F1: 4294967295}
    v8 := GenS2{F0: 1, F1: (p1 + p1), F2: 65535}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 128
}
func genFn1(p0 bool, p1 int64) bool {
    p1 -= (p1 - (p1 + 42))
    for fi9 := 0; fi9 < 3; fi9++ {
        if p0 {
            p1 -= -534
            for fi10 := 0; fi10 < 4; fi10++ {
                v11 := []int{(fi10 | (fi10 + 1)), fi10}
                if ((p0 && (1 != 488)) || (65535 <= 308)) {
                    p0 = (993 <= 721)
                    if p0 {
                        var v12 string = ("foo bar" + "a")
                        if len(v11) >= 2 {
                            v11 = v11[1:len(v11)]
                        }
                        var v13 string = v12
                        _ = v12
                        _ = v13
                    }
                    p1 -= p1
                } else {
                    if (p0 || false) {
                        switch "foo bar" {
                        case "hello":
                            fmt.Println(p1, fi10)
                        case "\n\t":
                            var v14 uint16 = 697
                            _ = v14
                        case "":
                            switch "a" {
                            case "hello":
                                if len(v11) >= 2 {
                                    v11 = v11[1:len(v11)]
                                }
                                var v15 uint = 4294967295
                                if len(v11) > 0 {
                                    v11[0] = fi9
                                }
                                _ = v15
                            case "a":
                                v16 := GenS3{F0: 1}
                                _ = v16
                            }
                        }
                    } else {
                        var v17 uint8 = genFn0(4294967295, 42)
                        _ = v17
                    }
                }
                p0 = (4294967295 >= 572)
                v18 := GenS4{F0: fi10, F1: (fi9 & 0), F2: 42}
                p1 = p1
                _ = fi10
                _ = v11
                _ = v18
            }
        } else {
            fmt.Println(p1, fi9, fi9)
        }
        _ = fi9
    }
    p1 -= (-(p1))
    for fi19 := 0; fi19 < 1; fi19++ {
        switch "foo bar" {
        case "foo bar":
            v20 := []int{}
            _ = v20
        case "\n\t":
            p1 -= p1
        }
        _ = fi19
    }
    if p0 {
        var v21 uint = 0
        v22 := []int{42, -9223372036854775808, 128, 86}
        if (825 != 0) {
            for fi23 := 0; fi23 < 5; fi23++ {
                var v24 int = fi23
                var v25 string = "hello"
                var v26 int = v24
                _ = fi23
                _ = v24
                _ = v25
                _ = v26
            }
            switch 563 {
            case -1:
                v21 += 4294967295
            case -100:
                p0 = (p0 && p0)
            }
        } else {
            if false {
                if p0 {
                    switch "a" {
                    case "a":
                        _ = len(v22)
                    case "hello":
                        for fi27 := 0; fi27 < 2; fi27++ {
                            switch fi27 {
                            case 32767:
                                for fi28 := 0; fi28 < 11; fi28++ {
                                    var v29 uint16 = 486
                                    _ = fi28
                                    _ = v29
                                }
                            case -1:
                                p1 -= (p1 - (p1 + p1))
                            }
                            _ = fi27
                        }
                    }
                    p0 = p0
                }
                v30 := GenS4{F0: 127, F1: -1, F2: 1}
                for fi31 := 0; fi31 < 8; fi31++ {
                    v32 := GenS3{F0: -2147483648}
                    p1 += p1
                    var v33 uint16 = 32767
                    _ = fi31
                    _ = v32
                    _ = v33
                }
                _ = v30
            } else {
                v34 := GenS1{F0: p1, F1: 711}
                for fi35 := 0; fi35 < 14; fi35++ {
                    v36 := GenS5{F0: 42}
                    for fi37 := 0; fi37 < 11; fi37++ {
                        p1 -= (p1 << (((v21 ^ v21)) & 63))
                        fmt.Println(fi35)
                        var v38 int16 = -32768
                        if len(v22) >= 2 {
                            v22 = v22[1:len(v22)]
                        }
                        if len(v22) > 0 {
                            v22[0] = (fi35 * fi35)
                        }
                        v39 := GenS6{F0: 1}
                        _ = fi37
                        _ = v38
                        _ = v39
                    }
                    _ = fi35
                    _ = v36
                }
                v21 -= v21
                _ = v34
            }
            switch ("a" + "foo bar") {
            case "":
                var v40 float32 = ((1.5 + 1e10) / 0.1)
                _ = v40
            default:
                fmt.Println(v21, p0)
            }
        }
        _ = v21
        _ = v22
    } else {
        p1 -= p1
        v41 := []int{4294967295, int(p1), 127, 7}
        _ = v41
    }
    _ = p0
    _ = p1
    return true
}

func main() {
    v1 := make(map[string]string)
    var v2 string = (("foo bar" + "") + "\n\t")
    v3 := []uint32{}
    var v4 bool = false
    var v5 int8 = 1
    v6 := make(map[int]string)
    if (true || v4) {
        _ = len(v6)
        var v42 uint16 = 472
        _ = v42
    } else {
        v4 = v4
    }
    switch -601 {
    case -1:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case 42:
        switch v2 {
        case "a":
            v43 := GenS7{F0: 0, F1: false}
            _ = v43
        case "":
            var v44 int16 = -1
            _ = v44
        }
    case -157:
        for fi45 := 0; fi45 < 20; fi45++ {
            for fi46 := 0; fi46 < 12; fi46++ {
                if ((v2 + (v2 + v2)) > v2) {
                    v6[(fi45 + fi45)] = ("hello" + ((v2 + "") + ""))
                    switch "a" {
                    case "\n\t":
                        var v47 bool = (0.1 == (0.1 * (1.5 + float64((float32(-2.25) / -2.25)))))
                        _ = v47
                    case "a":
                        if v4 {
                            v2 = v2
                            var v48 int32 = 0
                            _ = v48
                        } else {
                            var v49 uint = 884
                            if v4 {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                v5 = v5
                                v2 = ""
                            } else {
                                v49 = (v49 / ((v49) | 1))
                                v5 = v5
                            }
                            if len(v3) > 0 {
                                v3[0] = 0
                            }
                            _ = v49
                        }
                    }
                }
                v50 := v1["hello"]
                v4 = (32767 <= 32767)
                _ = fi46
                _ = v50
            }
            _ = fi45
        }
    default:
        var v51 float64 = (-2.25 - ((-2.25 - 1e10) * 3.14159))
        _ = v51
    }
    delete(v1, "hello")
    v52 := GenS8{F0: 4294967295}
    var v53 uint = 1
    var v54 int16 = 713
    if (1 >= 95) {
        for fi55 := 0; fi55 < 3; fi55++ {
            var v56 float32 = ((-2.25 * 0.0) - 0.0)
            for fi57 := 0; fi57 < 2; fi57++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                fmt.Println(fi57, fi57)
                _ = fi57
            }
            _ = fi55
            _ = v56
        }
        for fi58 := 0; fi58 < 9; fi58++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if v4 {
                delete(v6, 32768)
                if (!(129 != 2)) {
                    delete(v6, (fi58 % ((fi58) | 1)))
                    switch (fi58 ^ (fi58 << (((^(v53))) & 63))) {
                    case 32768:
                        var v59 int32 = 80
                        _ = v59
                    case 2147483647:
                        v60 := GenS4{F0: (fi58 >> ((v53) & 63)), F1: (fi58 % ((fi58) | 1)), F2: 42}
                        _ = v60
                    case 32767:
                        var v61 string = v2
                        _ = v61
                    }
                    switch (v2 + v2) {
                    case "":
                        var v62 int16 = (v54 % ((v54) | 1))
                        _ = v62
                    case "a":
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    case "hello":
                        for fi63 := 0; fi63 < 20; fi63++ {
                            v5 = (v5 ^ (v5 - (v5 - 127)))
                            switch (fi58 ^ 1) {
                            case 576:
                                fmt.Println(fi58, v54)
                            case 1:
                                v64 := GenS9{F0: 40, F1: 813}
                                _ = v64
                            case 127:
                                if false {
                                    if len(v3) > 0 {
                                        v3[0] = 4294967295
                                    }
                                    var v65 bool = v4
                                    var v66 string = v2
                                    _ = v65
                                    _ = v66
                                }
                            }
                            var v67 int8 = (-(v5))
                            _ = fi63
                            _ = v67
                        }
                    }
                }
                if false {
                    if v4 {
                        delete(v1, (v2 + "hello"))
                        var v68 int8 = 83
                        if (true || false) {
                            for fi69 := 0; fi69 < 5; fi69++ {
                                switch fi58 {
                                case 2147483647:
                                    var v70 float64 = (float64(1.5) + -2.25)
                                    _ = v70
                                case -586:
                                    fmt.Println(v2, v2, v5)
                                case 42:
                                    var v71 uint32 = 65535
                                    v54 = v54
                                    _ = v71
                                default:
                                    fmt.Println(v4, v54)
                                    var v72 int32 = 2147483647
                                    var v73 uint64 = 46
                                    var v74 int16 = -482
                                    _ = v72
                                    _ = v73
                                    _ = v74
                                }
                                v75 := GenS4{F0: fi58, F1: (fi69 + fi69), F2: 2}
                                _ = fi69
                                _ = v75
                            }
                            var v76 float32 = (1.5 + -2.25)
                            _ = v76
                        } else {
                            v5 -= (v68 % (((v5 ^ 0)) | 1))
                        }
                        _ = v68
                    } else {
                        if v4 {
                            for fi77 := 0; fi77 < 5; fi77++ {
                                fmt.Println(v4)
                                _ = fi77
                            }
                        }
                        v78 := GenS10{F0: v4, F1: 0, F2: 638}
                        for fi79 := 0; fi79 < 12; fi79++ {
                            if v4 {
                                var v80 bool = (!v4)
                                _ = v80
                            } else {
                                var v81 string = "foo bar"
                                var v82 bool = false
                                fmt.Println(v4)
                                fmt.Println(v2, v5)
                                _ = v81
                                _ = v82
                            }
                            _ = fi79
                        }
                        _ = v78
                    }
                    var v83 float32 = -2.25
                    switch "hello" {
                    case "hello":
                        var v84 uint32 = 2
                        _ = v84
                    }
                    _ = v83
                }
            }
            _ = fi58
        }
    }
    var v85 int = 32768
    v6[(v85 / (((v85 - v85)) | 1))] = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v85:", v85)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v53
    _ = v54
    _ = v6
    _ = v85
}
