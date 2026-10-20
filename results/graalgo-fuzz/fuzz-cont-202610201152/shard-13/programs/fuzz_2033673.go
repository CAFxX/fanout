package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 uint32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int32; F1 uint16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint16; F1 bool }
type GenS6 struct { F0 string; F1 int }
type GenS7 struct { F0 uint16; F1 int16 }
type GenS8 struct { F0 uint64; F1 int32; F2 int }

func main() {
    var v1 string = ("foo bar" + "a")
    var v2 int16 = 1000
    var v3 int8 = 100
    var v4 uint32 = 4294967295
    v5 := make(map[int]bool)
    var v6 uint16 = 32767
    var v7 uint16 = v6
    for fi8 := 0; fi8 < 13; fi8++ {
        v1 = ("foo bar" + ((((("\n\t" + "foo bar") + v1) + (v1 + v1)) + v1) + (v1 + v1)))
        for fi9 := 0; fi9 < 5; fi9++ {
            if true {
                var v10 string = v1
                v11 := GenS1{F0: v7, F1: 255, F2: (v4 | v4)}
                _ = v10
                _ = v11
            } else {
                if true {
                    v12 := GenS1{F0: 1, F1: 2, F2: 0}
                    _ = v12
                } else {
                    if false {
                        v1 = "a"
                        if ("a" != ((v1 + v1) + ((v1 + "hello") + v1))) {
                            v7 = (v7 - v6)
                        }
                    } else {
                        if false {
                            var v13 string = (v1 + (v1 + ("" + v1)))
                            _ = v13
                        } else {
                            var v14 uint8 = 96
                            if (-2147483648 != 2147483647) {
                                var v15 uint8 = (^(v14))
                                switch fi9 {
                                case 4294967295:
                                    v1 = ("" + v1)
                                    fmt.Println(v2, v3, v3)
                                default:
                                    var v16 uint64 = 4294967295
                                    _ = v16
                                }
                                if ((v2 >> ((42) & 63)) < v2) {
                                    fmt.Println(v14)
                                    fmt.Println(v15, v3, v2)
                                } else {
                                    v14 += (v15 * (v14 >> ((0) & 63)))
                                    if (!(!false)) {
                                        var v17 uint32 = (v4 / ((v4) | 1))
                                        v18 := GenS1{F0: (v6 / ((488) | 1)), F1: 4294967295, F2: (v4 | v17)}
                                        var v19 uint = 42
                                        _ = v17
                                        _ = v18
                                        _ = v19
                                    }
                                }
                                if (uint(v6) == 2) {
                                    v4 = (v4 * uint32(v14))
                                    v20 := GenS2{F0: 458}
                                    if (-2147483648 != 0) {
                                        v21 := GenS3{F0: -2147483648, F1: 255}
                                        v22 := []int{fi8, fi8, fi8, -1}
                                        v6 = v7
                                        var v23 int = 644
                                        switch v1 {
                                        case "hello":
                                            v15 -= (v15 >> ((18446744073709551615) & 63))
                                        case "foo bar":
                                            v14 = uint8(v4)
                                        case "":
                                            var v24 bool = true
                                            _ = v24
                                        }
                                        _ = v21
                                        _ = v22
                                        _ = v23
                                    } else {
                                        v25 := GenS4{F0: 255}
                                        _ = v25
                                    }
                                    _ = v20
                                } else {
                                    var v26 uint32 = v4
                                    _ = v26
                                }
                                fmt.Println(v2, v6)
                                _ = v15
                            } else {
                                if false {
                                    v27 := []int{fi9, fi8}
                                    _ = v27
                                } else {
                                    var v28 uint16 = v7
                                    var v29 int = fi8
                                    _ = v28
                                    _ = v29
                                }
                                if true {
                                    var v30 bool = false
                                    v3 += v3
                                    _ = v30
                                } else {
                                    if false {
                                        switch (fi9 ^ (fi9 ^ fi8)) {
                                        case 128:
                                            v7 = v7
                                        case 32768:
                                            var v31 uint = 18446744073709551615
                                            _ = v31
                                        case 548:
                                            var v32 string = (("" + ((v1 + "a") + v1)) + v1)
                                            v32 = v32
                                            fmt.Println(v4, v32)
                                            _ = v32
                                        }
                                    } else {
                                        fmt.Println(v6, v1, v6)
                                    }
                                    v33 := []int{fi8, 32768, fi9, -85}
                                    _ = v33
                                }
                                var v34 bool = ((^(v4)) <= v4)
                                _ = v34
                            }
                            v35 := GenS2{F0: 9223372036854775807}
                            _ = v14
                            _ = v35
                        }
                        switch ((v1 + "\n\t") + v1) {
                        case "":
                            fmt.Println(v2, v2, v7)
                        default:
                            switch v1 {
                            case "foo bar":
                                v2 -= v2
                            case "a":
                                v2 = v2
                                delete(v5, fi9)
                                var v36 uint16 = v6
                                _ = v36
                            }
                            v37 := GenS2{F0: -9223372036854775808}
                            v38 := []int{}
                            _ = v37
                            _ = v38
                        }
                    }
                    v4 += (v4 >> ((4294967295) & 63))
                }
            }
            var v39 uint8 = 1
            _ = fi9
            _ = v39
        }
        fmt.Println(v3, v4, v7)
        _ = fi8
    }
    if false {
        if (612 > 18446744073709551615) {
            if false {
                for fi40 := 0; fi40 < 18; fi40++ {
                    fmt.Println(v3, v7, v4)
                    v41 := []int{(-(fi40)), int(v4), (fi40 / ((-947) | 1))}
                    if false {
                        v3 -= (v3 - v3)
                        if len(v41) >= 2 {
                            v41 = v41[1:len(v41)]
                        }
                        var v42 uint32 = 967
                        _ = v42
                    } else {
                        v7 = v6
                    }
                    _ = fi40
                    _ = v41
                }
            } else {
                v43 := GenS4{F0: v6}
                v7 = (v6 / ((v7) | 1))
                v2 -= (v2 ^ -935)
                _ = v43
            }
        } else {
            for fi44 := 0; fi44 < 14; fi44++ {
                var v45 int32 = 0
                for fi46 := 0; fi46 < 8; fi46++ {
                    var v47 uint64 = 147
                    _ = fi46
                    _ = v47
                }
                switch v1 {
                case "a":
                    switch (-(fi44)) {
                    case -550:
                        switch -1 {
                        case 127:
                            if true {
                                if (!(!((!false) && true))) {
                                    if true {
                                        v6 = 2
                                        var v48 string = (("\n\t" + (("a" + v1) + ("a" + v1))) + "foo bar")
                                        _ = v48
                                    }
                                    var v49 uint8 = 255
                                    v7 += v6
                                    _ = v49
                                } else {
                                    v50 := []int{fi44}
                                    v7 = (v7 & v6)
                                    _ = v50
                                }
                            } else {
                                if false {
                                    for fi51 := 0; fi51 < 14; fi51++ {
                                        v1 = v1
                                        _ = fi51
                                    }
                                    v52 := v5[fi44]
                                    var v53 string = v1
                                    _ = v52
                                    _ = v53
                                }
                                if true {
                                    var v54 uint8 = 213
                                    _ = v54
                                }
                            }
                        }
                    case 1:
                        var v55 string = v1
                        _ = v55
                    default:
                        v56 := GenS4{F0: v6}
                        _ = v56
                    }
                case "\n\t":
                    switch (fi44 | -612) {
                    case -2147483648:
                        if true {
                            _ = len(v5)
                        } else {
                            for fi57 := 0; fi57 < 7; fi57++ {
                                fmt.Println(v1, v45, v6)
                                var v58 uint64 = 915
                                var v59 uint8 = 127
                                _ = fi57
                                _ = v58
                                _ = v59
                            }
                            fmt.Println(v7, fi44)
                            for fi60 := 0; fi60 < 20; fi60++ {
                                v61 := GenS3{F0: (v45 >> ((378) & 63)), F1: (v7 / ((v7) | 1))}
                                switch "hello" {
                                case "foo bar":
                                    v6 = 255
                                case "a":
                                    _ = len(v5)
                                case "\n\t":
                                    var v62 uint16 = v6
                                    var v63 uint64 = 4294967295
                                    v64 := GenS5{F0: v7, F1: true}
                                    _ = v62
                                    _ = v63
                                    _ = v64
                                default:
                                    var v65 float32 = ((1.5 - 1e10) * 0.1)
                                    _ = v65
                                }
                                v7 -= (v7 | (v6 - v6))
                                _ = fi60
                                _ = v61
                            }
                        }
                    case 127:
                        v2 = v2
                    case 2147483647:
                        v66 := GenS5{F0: (^(v6)), F1: true}
                        _ = v66
                    default:
                        fmt.Println(fi44, v1, v7)
                    }
                default:
                    v67 := GenS6{F0: "\n\t", F1: (fi44 << ((534) & 63))}
                    _ = v67
                }
                _ = fi44
                _ = v45
            }
            v68 := GenS1{F0: 902, F1: uint(v4), F2: v4}
            var v69 int32 = -82
            _ = v68
            _ = v69
        }
        fmt.Println(v7, v3)
        v70 := GenS7{F0: v7, F1: (^(v2))}
        _ = v70
    }
    v71 := GenS2{F0: 24}
    for fi72 := 0; fi72 < 9; fi72++ {
        var v73 int8 = (v3 + (v3 & -128))
        v74 := GenS8{F0: uint64(v4), F1: -2147483648, F2: fi72}
        _ = fi72
        _ = v73
        _ = v74
    }
    v75 := []int{-840, -23, 127}
    var v76 uint32 = (v4 - (v4 * (v4 | v4)))
    var v77 int64 = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v75:", len(v75), v75)
    fmt.Println("v76:", v76)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v71
    _ = v75
    _ = v76
    _ = v77
}
