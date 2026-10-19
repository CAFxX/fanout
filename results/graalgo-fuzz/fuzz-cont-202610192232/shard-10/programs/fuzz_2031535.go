package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 bool; F1 uint32; F2 int64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint; F1 int32 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint32; F1 int }
type GenS8 struct { F0 int64; F1 bool }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 uint8; F1 uint16 }
func genFn0(p0 uint64, p1 uint32, p2 string) int16 {
    v5 := GenS1{F0: 0, F1: 42}
    var v6 int64 = -170
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 1
}
func genFn1(p0 string, p1 int8) string {
    v7 := []int{127}
    v8 := GenS1{F0: 582, F1: 2}
    var v9 string = (("\n\t" + (p0 + (p0 + ("foo bar" + p0)))) + p0)
    var v10 int32 = -2147483648
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return p0
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]bool)
    v3 := make(map[int]string)
    v4 := make(map[string]string)
    var v11 int64 = 1
    var v12 string = ""
    for fi13 := 0; fi13 < 11; fi13++ {
        v12 = v12
        if (true || true) {
            v12 = v12
        } else {
            var v14 string = v12
            var v15 int64 = (v11 * -318)
            _ = v14
            _ = v15
        }
        _ = fi13
    }
    for fi16 := 0; fi16 < 13; fi16++ {
        if false {
            v11 = (v11 << ((4294967295) & 63))
            if true {
                v17 := GenS1{F0: 255, F1: 42}
                if true {
                    switch (fi16 / ((32768) | 1)) {
                    case 9223372036854775807:
                        v12 = v12
                    case 2147483647:
                        if true {
                            for fi18 := 0; fi18 < 12; fi18++ {
                                v11 -= (v11 & -745)
                                v19 := v2[""]
                                v19 = false
                                v20 := GenS1{F0: 1, F1: 0}
                                _ = fi18
                                _ = v19
                                _ = v20
                            }
                            if false {
                                v21 := GenS2{F0: ((627 > v11) || false), F1: 2147483648, F2: v11}
                                _ = len(v3)
                                for fi22 := 0; fi22 < 13; fi22++ {
                                    var v23 uint8 = 2
                                    _ = fi22
                                    _ = v23
                                }
                                _ = v21
                            } else {
                                if (0 != 1) {
                                    v24 := []int{(^(fi16)), 2147483647}
                                    v1[fi16] = -1
                                    delete(v4, (v12 + (((v12 + "") + "foo bar") + v12)))
                                    _ = v24
                                } else {
                                    for fi25 := 0; fi25 < 8; fi25++ {
                                        v12 = "foo bar"
                                        fmt.Println(fi25, v12, fi16)
                                        _ = fi25
                                    }
                                }
                                var v26 bool = true
                                _ = v26
                            }
                        }
                    }
                    if false {
                        var v27 uint64 = 18446744073709551615
                        switch fi16 {
                        case 4294967295:
                            v28 := GenS3{F0: 95}
                            _ = v28
                        default:
                            v29 := GenS3{F0: 127}
                            _ = v29
                        }
                        if (fi16 != 127) {
                            v2[(v12 + v12)] = (!(91 < 255))
                            var v30 float64 = -2.25
                            v31 := []int{fi16, (^(fi16)), int(v11), (fi16 | fi16)}
                            _ = v30
                            _ = v31
                        }
                        _ = v27
                    } else {
                        var v32 float64 = ((1.5 + 1.5) - (float64(3.14159) * 1.5))
                        _ = v32
                    }
                    v11 -= -9223372036854775808
                }
                _ = v17
            } else {
                fmt.Println(fi16, v12, v11)
                fmt.Println(fi16)
            }
        }
        var v33 uint = 42
        _ = fi16
        _ = v33
    }
    v12 = v12
    v34 := GenS4{F0: 609}
    v12 = (v12 + (v12 + "foo bar"))
    if false {
        v12 = v12
        for fi35 := 0; fi35 < 7; fi35++ {
            v12 = v12
            switch (fi35 << ((255) & 63)) {
            case -38:
                switch (v12 + ((((((v12 + "a") + v12) + v12) + (v12 + (((v12 + (v12 + "foo bar")) + v12) + ("" + v12)))) + v12) + "a")) {
                case "a":
                    for fi36 := 0; fi36 < 8; fi36++ {
                        if true {
                            var v37 int16 = 42
                            v12 = v12
                            v12 = ((("foo bar" + v12) + v12) + v12)
                            _ = v37
                        } else {
                            delete(v4, v12)
                            var v38 uint64 = 4294967295
                            _ = v38
                        }
                        v11 += v11
                        _ = fi36
                    }
                case "":
                    v39 := []int{(^(fi35))}
                    _ = v39
                }
            case 42:
                if (false || false) {
                    v12 = (("foo bar" + v12) + "a")
                    switch "\n\t" {
                    case "":
                        if false {
                            var v40 uint = 18446744073709551615
                            _ = v40
                        } else {
                            var v41 float32 = -2.25
                            switch fi35 {
                            case -367:
                                var v42 int16 = 398
                                _ = v42
                            case 42:
                                fmt.Println(v11, v41, v41)
                            default:
                                v43 := GenS5{F0: 179, F1: 42}
                                _ = v43
                            }
                            var v44 float64 = -2.25
                            _ = v41
                            _ = v44
                        }
                    }
                    if ("" < (v12 + v12)) {
                        for fi45 := 0; fi45 < 14; fi45++ {
                            var v46 uint = 65535
                            _ = fi45
                            _ = v46
                        }
                        if (false && false) {
                            var v47 uint64 = 2
                            _ = v47
                        } else {
                            switch (v12 + ("" + v12)) {
                            case "hello":
                                switch fi35 {
                                case 871:
                                    var v48 float32 = 0.0
                                    _ = v48
                                case 2147483647:
                                    delete(v1, (fi35 - -897))
                                }
                            case "\n\t":
                                fmt.Println(fi35, v11)
                            }
                            if false {
                                var v49 bool = false
                                fmt.Println(v12, v49, fi35)
                                _ = v49
                            }
                            if (1 == 905) {
                                v50 := GenS6{F0: 2}
                                _ = v50
                            }
                        }
                    }
                } else {
                    fmt.Println(v12, v12)
                    switch "foo bar" {
                    case "foo bar":
                        for fi51 := 0; fi51 < 5; fi51++ {
                            if false {
                                switch fi35 {
                                case -1:
                                    v52 := GenS1{F0: 1, F1: 521}
                                    _ = v52
                                case 381:
                                    if true {
                                        v53 := GenS7{F0: 1, F1: (fi51 * 1)}
                                        if (658 > 65535) {
                                            var v54 uint32 = 42
                                            if false {
                                                var v55 int8 = -125
                                                v56 := GenS8{F0: 1, F1: (42 < 2147483647)}
                                                v57 := v2[v12]
                                                _ = v55
                                                _ = v56
                                                _ = v57
                                            } else {
                                                v11 = (v11 >> ((18446744073709551615) & 63))
                                            }
                                            _ = v54
                                        } else {
                                            fmt.Println(v12, v11)
                                            v11 = (v11 << ((18446744073709551615) & 63))
                                        }
                                        var v58 int8 = 42
                                        var v59 string = "foo bar"
                                        _ = len(v1)
                                        _ = v53
                                        _ = v58
                                        _ = v59
                                    }
                                default:
                                    v12 = v12
                                }
                            } else {
                                v60 := GenS9{F0: 2147483648}
                                var v61 string = "\n\t"
                                if (-460 == -236) {
                                    v62 := GenS10{F0: 131, F1: 502}
                                    _ = v62
                                }
                                _ = v60
                                _ = v61
                            }
                            _ = fi51
                        }
                    }
                }
            default:
                fmt.Println(v11)
            }
            var v63 int = int(v11)
            _ = fi35
            _ = v63
        }
        var v64 uint = 42
        _ = v64
    } else {
        if (false || true) {
            v12 = "\n\t"
            delete(v2, v12)
            var v65 string = v12
            _ = v65
        }
        if true {
            _ = len(v3)
        }
    }
    var v66 float64 = 0.0
    for fi67 := 0; fi67 < 2; fi67++ {
        var v68 string = (v12 + (v12 + ("foo bar" + "a")))
        _ = fi67
        _ = v68
    }
    var v69 float64 = v66
    v70 := GenS4{F0: 1}
    if true {
        for fi71 := 0; fi71 < 11; fi71++ {
            var v72 int32 = -2147483648
            v73 := v4[v12]
            _ = fi71
            _ = v72
            _ = v73
        }
        v74 := GenS10{F0: 128, F1: 922}
        if (!false) {
            fmt.Println(v12)
        } else {
            if true {
                var v75 uint8 = 127
                _ = v75
            } else {
                var v76 uint = 587
                _ = v76
            }
        }
        _ = v74
    }
    var v77 string = "hello"
    v77 = (v77 + "hello")
    v4[((("" + v77) + v77) + v77)] = ("" + (v77 + ((v77 + v77) + v77)))
    for fi78 := 0; fi78 < 6; fi78++ {
        fmt.Println(v12, fi78, v11)
        _ = fi78
    }
    if true {
        var v79 string = (v12 + v77)
        _ = v79
    } else {
        var v80 uint8 = 1
        v81 := []int{0}
        for fi82 := 0; fi82 < 16; fi82++ {
            v80 -= 2
            switch "foo bar" {
            case "a":
                if true {
                    fmt.Println(v69)
                } else {
                    fmt.Println(v66, v69)
                    _ = len(v81)
                    v66 = float64(((3.14159 - -2.25) + 1.5))
                }
            }
            _ = fi82
        }
        _ = v80
        _ = v81
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v66:", v66)
    fmt.Println("v69:", v69)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v66
    _ = v69
    _ = v70
    _ = v77
}
