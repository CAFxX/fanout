package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 string; F1 uint8 }
type GenS4 struct { F0 int64; F1 uint64 }
type GenS5 struct { F0 int; F1 uint }
type GenS6 struct { F0 uint16; F1 int16 }
type GenS7 struct { F0 uint16; F1 string }
type GenS8 struct { F0 uint; F1 int16; F2 int8 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 uint32; F1 int; F2 string }
type GenS12 struct { F0 string; F1 int8; F2 uint16 }

func main() {
    var v1 bool = (-208 > 0)
    var v2 string = "\n\t"
    var v3 uint8 = 42
    v4 := make(map[int]bool)
    v5 := make(map[uint32]string)
    var v6 uint32 = 415
    var v7 int16 = 786
    _ = len(v4)
    for fi8 := 0; fi8 < 4; fi8++ {
        v9 := GenS1{F0: (v6 << ((0) & 63))}
        for fi10 := 0; fi10 < 20; fi10++ {
            v1 = false
            var v11 uint64 = 4294967295
            if (-957 > 2147483647) {
                v12 := GenS1{F0: 2}
                var v13 uint16 = 966
                v14 := []int{fi8, fi8, (^(fi10)), 86}
                _ = v12
                _ = v13
                _ = v14
            } else {
                switch v2 {
                case "foo bar":
                    v15 := GenS1{F0: 65535}
                    _ = v15
                case "a":
                    var v16 string = v2
                    _ = v16
                case "":
                    v17 := GenS2{F0: 0, F1: -1}
                    _ = v17
                default:
                    var v18 uint = 255
                    _ = v18
                }
                v7 = v7
            }
            v3 = 215
            if (!v1) {
                if false {
                    var v19 uint8 = (v3 / (((v3 >> ((18446744073709551615) & 63))) | 1))
                    _ = v19
                }
            } else {
                var v20 string = (v2 + ("\n\t" + v2))
                _ = v20
            }
            fmt.Println(fi10)
            var v21 uint32 = (v6 >> ((18446744073709551615) & 63))
            _ = fi10
            _ = v11
            _ = v21
        }
        _ = fi8
        _ = v9
    }
    if (!(v1 || ((3.14159 >= (0.0 / -2.25)) || v1))) {
        for fi22 := 0; fi22 < 6; fi22++ {
            v7 = (v7 % ((v7) | 1))
            v3 = v3
            _ = fi22
        }
        for fi23 := 0; fi23 < 1; fi23++ {
            if v1 {
                v2 = ""
                switch fi23 {
                case 529:
                    fmt.Println(v1, v1)
                case 2147483647:
                    v24 := GenS3{F0: (("" + v2) + v2), F1: 221}
                    _ = v24
                case 1:
                    fmt.Println(fi23, v7)
                }
            } else {
                var v25 float32 = (1e10 + 0.0)
                v25 -= v25
                _ = v25
            }
            for fi26 := 0; fi26 < 6; fi26++ {
                fmt.Println(fi26)
                var v27 int64 = 1
                if (!v1) {
                    var v28 float64 = (float64(float32(1e10)) * 1e10)
                    switch "" {
                    case "foo bar":
                        var v29 uint64 = 1
                        _ = v29
                    default:
                        v1 = (84 <= 42)
                        switch v2 {
                        case "a":
                            var v30 uint = 18446744073709551615
                            _ = v30
                        case "hello":
                            switch v2 {
                            case "a":
                                var v31 int32 = -684
                                var v32 float64 = v28
                                var v33 bool = v1
                                if (3.14159 < v32) {
                                    v7 = 615
                                    switch (-(fi26)) {
                                    case 1:
                                        switch -635 {
                                        case 128:
                                            v27 = v27
                                        case -1:
                                            v34 := GenS2{F0: v31, F1: (-(v27))}
                                            if false {
                                                if v1 {
                                                    var v35 uint8 = 42
                                                    fmt.Println(v6)
                                                    _ = v35
                                                } else {
                                                    var v36 int8 = -124
                                                    _ = v36
                                                }
                                                v37 := GenS2{F0: (v31 & v31), F1: v27}
                                                var v38 float32 = float32(v32)
                                                _ = v37
                                                _ = v38
                                            }
                                            _ = v34
                                        case 6:
                                            v39 := GenS4{F0: int64(v7), F1: 794}
                                            _ = v39
                                        default:
                                            var v40 int32 = (v31 | 1)
                                            _ = v40
                                        }
                                    }
                                }
                                _ = v31
                                _ = v32
                                _ = v33
                            case "hello":
                                if v1 {
                                    v41 := GenS5{F0: 128, F1: 4294967295}
                                    _ = v41
                                } else {
                                    v42 := GenS6{F0: 255, F1: (v7 + (^(v7)))}
                                    _ = v42
                                }
                            default:
                                v43 := GenS7{F0: 1, F1: ((v2 + v2) + ((((("foo bar" + (v2 + v2)) + "a") + v2) + ("hello" + (v2 + v2))) + "\n\t"))}
                                _ = v43
                            }
                        case "":
                            if v1 {
                                if false {
                                    v28 += v28
                                    if true {
                                        v7 = v7
                                        fmt.Println(v2)
                                        var v44 uint32 = v6
                                        v45 := GenS8{F0: 24, F1: v7, F2: 42}
                                        var v46 int64 = -700
                                        _ = v44
                                        _ = v45
                                        _ = v46
                                    }
                                } else {
                                    v47 := GenS9{F0: 0}
                                    fmt.Println(v6, fi23, v27)
                                    v48 := GenS10{F0: 111}
                                    _ = v47
                                    _ = v48
                                }
                                var v49 int = (fi23 >> ((18446744073709551615) & 63))
                                if ((^(v6)) <= (v6 << ((1) & 63))) {
                                    if v1 {
                                        v50 := []int{42, v49, 127}
                                        _ = v50
                                    }
                                } else {
                                    if v1 {
                                        if v1 {
                                            v51 := GenS11{F0: v6, F1: fi26, F2: v2}
                                            v2 = (v2 + (v2 + "\n\t"))
                                            switch (-(v49)) {
                                            case 128:
                                                var v52 uint32 = v6
                                                _ = v52
                                            default:
                                                var v53 float32 = -2.25
                                                _ = v53
                                            }
                                            _ = v51
                                        }
                                    }
                                }
                                _ = v49
                            }
                        default:
                            fmt.Println(fi23)
                        }
                        switch v2 {
                        case "a":
                            var v54 uint8 = 66
                            _ = v54
                        case "":
                            v27 = (v27 & (^(v27)))
                        }
                    }
                    _ = v28
                } else {
                    v55 := GenS8{F0: 560, F1: v7, F2: -30}
                    _ = v55
                }
                _ = fi26
                _ = v27
            }
            if (v1 && v1) {
                v2 = v2
                var v56 string = (v2 + "hello")
                for fi57 := 0; fi57 < 11; fi57++ {
                    if v1 {
                        fmt.Println(v7, fi23, fi23)
                    }
                    if v1 {
                        var v58 uint64 = 2
                        _ = v58
                    } else {
                        var v59 int64 = -1
                        if v1 {
                            v60 := GenS5{F0: fi57, F1: 0}
                            var v61 int = int(v6)
                            _ = v60
                            _ = v61
                        }
                        switch "foo bar" {
                        case "":
                            v1 = v1
                        case "a":
                            switch (fi23 & fi57) {
                            case -9223372036854775808:
                                if v1 {
                                    v6 = v6
                                    var v62 uint = 255
                                    _ = v62
                                }
                            case -908:
                                var v63 uint8 = (v3 % ((128) | 1))
                                _ = v63
                            case 4294967295:
                                v7 += v7
                            }
                        default:
                            if v1 {
                                v64 := v5[(v6 % (((v6 & 2147483647)) | 1))]
                                _ = v64
                            }
                        }
                        _ = v59
                    }
                    var v65 int16 = (v7 % (((v7 - int16(v3))) | 1))
                    _ = fi57
                    _ = v65
                }
                _ = v56
            } else {
                var v66 float64 = ((float64(1e10) * (float64(0.1) * 0.0)) / -2.25)
                _ = v66
            }
            _ = fi23
        }
    }
    var v67 int32 = 0
    if (0 >= 482) {
        v6 = v6
    }
    var v68 int64 = 1
    var v69 uint8 = v3
    v70 := GenS7{F0: 326, F1: v2}
    fmt.Println(v69)
    v71 := GenS12{F0: v2, F1: -27, F2: uint16(v6)}
    var v72 int = -1
    for fi73 := 0; fi73 < 7; fi73++ {
        v74 := GenS5{F0: fi73, F1: 18446744073709551615}
        _ = fi73
        _ = v74
    }
    v75 := []int{(^(v72))}
    switch ("foo bar" + v2) {
    case "":
        v7 += int16(v67)
    }
    for fi76 := 0; fi76 < 5; fi76++ {
        if (!(!(v1 || (2 < v6)))) {
            var v77 int8 = 42
            switch v72 {
            case -506:
                fmt.Println(v2)
            default:
                v2 = ""
            }
            _ = v77
        }
        _ = fi76
    }
    var v78 uint = 255
    var v79 uint32 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    fmt.Println("v69:", v69)
    fmt.Println("v7:", v7)
    fmt.Println("v72:", v72)
    fmt.Println("slicelen v75:", len(v75), v75)
    fmt.Println("v78:", v78)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v67
    _ = v68
    _ = v69
    _ = v7
    _ = v70
    _ = v71
    _ = v72
    _ = v75
    _ = v78
    _ = v79
}
