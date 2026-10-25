package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64; F1 uint64 }
type GenS4 struct { F0 bool; F1 int }

func main() {
    var v1 uint = 4294967295
    var v2 int16 = 0
    v3 := make(map[int]string)
    v4 := []int{338, 4294967295, 128}
    v5 := make(map[int]int)
    v6 := make(map[string]int)
    var v7 bool = ((42 != 127) || false)
    delete(v6, ("\n\t" + (("" + "\n\t") + "hello")))
    if true {
        for fi8 := 0; fi8 < 16; fi8++ {
            v7 = v7
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v9 := GenS1{F0: -510, F1: -2147483648, F2: -610}
            _ = fi8
            _ = v9
        }
        fmt.Println(v1)
        v10 := GenS1{F0: -9223372036854775808, F1: 514, F2: int64(v2)}
        _ = v10
    }
    _ = len(v4)
    v1 += 0
    for fi11 := 0; fi11 < 13; fi11++ {
        for fi12 := 0; fi12 < 10; fi12++ {
            v7 = false
            v2 = v2
            _ = fi12
        }
        fmt.Println(v2)
        _ = fi11
    }
    for fi13 := 0; fi13 < 19; fi13++ {
        var v14 int = (fi13 & 1)
        v15 := GenS1{F0: int64(v14), F1: 0, F2: -9223372036854775808}
        _ = fi13
        _ = v14
        _ = v15
    }
    if v7 {
        var v16 uint32 = 4294967295
        _ = v16
    } else {
        v1 += (^(v1))
        var v17 int64 = 944
        _ = v17
    }
    v1 -= v1
    var v18 uint = (v1 % ((2) | 1))
    var v19 int = -9223372036854775808
    var v20 float32 = float32(3.14159)
    if (v1 <= v1) {
        if (((v1 << ((66) & 63)) != 2) && v7) {
            v21 := GenS2{F0: 42}
            _ = v21
        }
    } else {
        delete(v5, (v19 << ((375) & 63)))
        v20 = v20
        v20 += (v20 - v20)
    }
    if false {
        for fi22 := 0; fi22 < 14; fi22++ {
            for fi23 := 0; fi23 < 1; fi23++ {
                v24 := GenS2{F0: 514}
                if v7 {
                    var v25 bool = (v19 > fi22)
                    var v26 uint = (v1 | v18)
                    if (v26 == 678) {
                        var v27 uint32 = 1
                        _ = v27
                    }
                    _ = v25
                    _ = v26
                }
                _ = fi23
                _ = v24
            }
            if (!(v7 || ((((0.1 + (3.14159 - -2.25)) * 1.5) * -2.25) != 0.0))) {
                fmt.Println(v2, fi22, v18)
            } else {
                fmt.Println(v2)
            }
            for fi28 := 0; fi28 < 12; fi28++ {
                var v29 string = "foo bar"
                fmt.Println(fi28, v18)
                _ = fi28
                _ = v29
            }
            _ = fi22
        }
        fmt.Println(v19, v19)
        var v30 int64 = -9223372036854775808
        _ = v30
    } else {
        v2 += v2
        for fi31 := 0; fi31 < 7; fi31++ {
            v32 := GenS2{F0: 32767}
            var v33 uint = 42
            _ = fi31
            _ = v32
            _ = v33
        }
        if v7 {
            v2 += v2
            for fi34 := 0; fi34 < 6; fi34++ {
                v35 := GenS1{F0: -1, F1: 710, F2: -9223372036854775808}
                if (1000 > -32768) {
                    if v7 {
                        switch "hello" {
                        case "foo bar":
                            var v36 uint32 = 0
                            _ = v36
                        default:
                            v2 = v2
                        }
                        var v37 uint16 = 857
                        var v38 uint8 = 255
                        _ = v37
                        _ = v38
                    }
                    var v39 uint16 = 0
                    _ = v39
                } else {
                    for fi40 := 0; fi40 < 5; fi40++ {
                        var v41 int64 = 1
                        var v42 float32 = -2.25
                        if true {
                            if ("hello" <= ("foo bar" + ((((("hello" + "foo bar") + "hello") + ((("foo bar" + "foo bar") + "hello") + "\n\t")) + "hello") + "hello"))) {
                                if true {
                                    var v43 uint64 = 0
                                    v1 -= 2
                                    if len(v4) > 0 {
                                        v4[0] = 42
                                    }
                                    _ = v43
                                }
                            }
                        } else {
                            v42 = 1.5
                            if (v7 || (true || (-2.25 < v20))) {
                                if (42 < 1) {
                                    v41 = 9223372036854775807
                                    if (163 == 1) {
                                        var v44 int32 = 452
                                        _ = v44
                                    } else {
                                        var v45 int32 = 42
                                        if v7 {
                                            var v46 uint16 = 32767
                                            delete(v6, "a")
                                            v42 -= v42
                                            _ = v46
                                        } else {
                                            var v47 uint16 = 32767
                                            var v48 int64 = (^(v41))
                                            v49 := GenS2{F0: (v47 / ((v47) | 1))}
                                            _ = v47
                                            _ = v48
                                            _ = v49
                                        }
                                        if (!true) {
                                            v41 = (v41 >> (((v1 * v1)) & 63))
                                            var v50 uint8 = 1
                                            _ = v50
                                        }
                                        _ = v45
                                    }
                                    var v51 float64 = -2.25
                                    v52 := GenS2{F0: 1}
                                    var v53 int = fi34
                                    v54 := GenS3{F0: (v41 << (((v1 * (v1 + 42))) & 63)), F1: 436}
                                    switch -9223372036854775808 {
                                    case -380:
                                        if (1 <= int16(v18)) {
                                            v20 += v20
                                            var v55 uint16 = 65535
                                            _ = v55
                                        } else {
                                            v56 := GenS3{F0: (v41 << ((v1) & 63)), F1: 1}
                                            if (127 == int8(v53)) {
                                                if v7 {
                                                    switch "\n\t" {
                                                    case "a":
                                                        if (!((fi40 & v19) > v53)) {
                                                            switch v19 {
                                                            case 628:
                                                                _ = len(v4)
                                                            case 4294967295:
                                                                v57 := GenS3{F0: -959, F1: 435}
                                                                _ = v57
                                                            default:
                                                                v42 = 1.5
                                                            }
                                                            var v58 bool = ((v51 * float64(1e10)) >= (v51 + (v51 * v51)))
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                            _ = v58
                                                        }
                                                    case "foo bar":
                                                        fmt.Println(v7, fi40, v20)
                                                    }
                                                    switch "" {
                                                    case "hello":
                                                        v20 -= 1e10
                                                        fmt.Println(v7, v7, v53)
                                                        var v59 int32 = 42
                                                        v41 -= 235
                                                        v6["foo bar"] = fi40
                                                        switch fi34 {
                                                        case 87:
                                                            fmt.Println(v19)
                                                        case 128:
                                                            var v60 string = "a"
                                                            _ = v60
                                                        case 32768:
                                                            var v61 int16 = (^(v2))
                                                            _ = v61
                                                        default:
                                                            _ = len(v6)
                                                        }
                                                        fmt.Println(v18, v51, fi40)
                                                        _ = v59
                                                    case "\n\t":
                                                        fmt.Println(v41)
                                                    case "":
                                                        var v62 int64 = v41
                                                        _ = v62
                                                    }
                                                    if (v7 || (v1 < v18)) {
                                                        v51 = v51
                                                        v18 = v18
                                                        v19 += fi34
                                                    }
                                                }
                                            } else {
                                                var v63 uint16 = 498
                                                _ = v63
                                            }
                                            _ = v56
                                        }
                                    }
                                    _ = v51
                                    _ = v52
                                    _ = v53
                                    _ = v54
                                } else {
                                    if (!(!true)) {
                                        if v7 {
                                            if (1 == 0) {
                                                var v64 int16 = v2
                                                var v65 int = (fi40 << (((v18 >> (((v1 << ((436) & 63))) & 63))) & 63))
                                                v66 := v3[(v19 % ((2147483647) | 1))]
                                                _ = v64
                                                _ = v65
                                                _ = v66
                                            } else {
                                                if (4294967295 != 353) {
                                                    fmt.Println(fi40, v19, v42)
                                                    v67 := GenS3{F0: v41, F1: 156}
                                                    _ = v67
                                                }
                                            }
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            if v7 {
                                                var v68 int8 = -16
                                                _ = v68
                                            } else {
                                                v69 := GenS4{F0: v7, F1: 4294967295}
                                                var v70 uint32 = 970
                                                switch int(v70) {
                                                case 32768:
                                                    v1 += (v1 | 2)
                                                case 83:
                                                    _ = len(v3)
                                                }
                                                _ = v69
                                                _ = v70
                                            }
                                        } else {
                                            if (v41 >= v41) {
                                                var v71 uint64 = 4294967295
                                                _ = v71
                                            } else {
                                                switch fi40 {
                                                case 654:
                                                    v2 += (v2 & -973)
                                                default:
                                                    v6["foo bar"] = (fi34 + (fi34 * fi40))
                                                }
                                                v41 = (v41 & (v41 << ((4294967295) & 63)))
                                            }
                                            v42 -= v42
                                        }
                                        fmt.Println(v18, v42)
                                        var v72 int = fi40
                                        _ = v72
                                    }
                                }
                                var v73 int = fi34
                                _ = v73
                            }
                            var v74 uint16 = 428
                            _ = v74
                        }
                        _ = fi40
                        _ = v41
                        _ = v42
                    }
                    var v75 int64 = 0
                    delete(v3, (v19 << ((255) & 63)))
                    _ = v75
                }
                _ = fi34
                _ = v35
            }
            if true {
                if v7 {
                    v1 = (v18 * 495)
                    v18 = (v18 + (^(v18)))
                    var v76 uint64 = 456
                    _ = v76
                }
                for fi77 := 0; fi77 < 3; fi77++ {
                    v18 += (v1 << ((2) & 63))
                    var v78 int = fi77
                    _ = fi77
                    _ = v78
                }
                if (!false) {
                    v18 -= (^(v1))
                    switch ("" + (("foo bar" + ("a" + "\n\t")) + (("hello" + "hello") + "\n\t"))) {
                    case "hello":
                        for fi79 := 0; fi79 < 17; fi79++ {
                            if (v7 && (!true)) {
                                fmt.Println(v2, v1)
                                v1 = v1
                                v18 = (v1 * v18)
                            } else {
                                var v80 string = ""
                                var v81 uint32 = 0
                                var v82 int = 42
                                _ = v80
                                _ = v81
                                _ = v82
                            }
                            var v83 int64 = 42
                            _ = fi79
                            _ = v83
                        }
                    case "\n\t":
                        var v84 int = (v19 | v19)
                        _ = v84
                    default:
                        v18 += (^(v1))
                    }
                }
            } else {
                var v85 uint64 = 0
                v86 := GenS3{F0: -492, F1: (v85 | (v85 % ((1) | 1)))}
                v18 = 0
                _ = v85
                _ = v86
            }
        }
    }
    for fi87 := 0; fi87 < 20; fi87++ {
        v88 := GenS4{F0: v7, F1: (v19 >> (((v18 ^ 363)) & 63))}
        v89 := v3[v19]
        _ = fi87
        _ = v88
        _ = v89
    }
    var v90 int32 = -1
    var v91 int = 32767
    v90 = int32(v91)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v90:", v90)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v90
    _ = v91
}
