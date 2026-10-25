package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int32; F1 uint64 }
type GenS5 struct { F0 int8; F1 int32; F2 int8 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint16; F1 int8 }
func genFn0(p0 bool, p1 uint16) uint32 {
    var v7 float32 = ((3.14159 / 0.1) / (0.1 + (1.5 + 1.5)))
    if (153 == -451) {
        if p0 {
            if p0 {
                if (!(1000 == 970)) {
                    p0 = (!(42 == 1))
                    if p0 {
                        v8 := GenS1{F0: (p1 << ((255) & 63))}
                        for fi9 := 0; fi9 < 17; fi9++ {
                            var v10 uint64 = 4294967295
                            _ = fi9
                            _ = v10
                        }
                        v11 := []int{}
                        _ = v11
                        _ = v8
                    }
                }
            } else {
                v12 := GenS1{F0: (p1 ^ 32767)}
                var v13 uint8 = 127
                v14 := GenS1{F0: p1}
                _ = v12
                _ = v13
                _ = v14
            }
            v7 += 0.0
        } else {
            var v15 float64 = ((0.1 + 1.5) / 3.14159)
            _ = v15
        }
        fmt.Println(v7)
    }
    v16 := GenS1{F0: 255}
    _ = p0
    _ = p1
    _ = v16
    _ = v7
    return 568
}
func genFn1() int32 {
    var v17 bool = (false || (("foo bar" + "foo bar") > ("\n\t" + "")))
    v18 := []int{9223372036854775807, -583, 127, 32768}
    var v19 uint32 = 4294967295
    for fi20 := 0; fi20 < 2; fi20++ {
        var v21 int32 = -497
        _ = fi20
        _ = v21
    }
    _ = v17
    _ = v18
    _ = v19
    return -128
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]string)
    v3 := []string{""}
    var v4 uint16 = 0
    var v5 uint32 = 0
    var v6 int64 = -75
    var v22 int = 2147483647
    if true {
        var v23 uint8 = 42
        _ = v23
    } else {
        delete(v1, (v22 | -144))
        fmt.Println(v22, v6)
    }
    fmt.Println(v4, v5)
    fmt.Println(v22)
    switch v22 {
    case 42:
        for fi24 := 0; fi24 < 7; fi24++ {
            v4 = (v4 % (((v4 % ((v4) | 1))) | 1))
            _ = fi24
        }
    default:
        if (!("a" > "hello")) {
            for fi25 := 0; fi25 < 11; fi25++ {
                if len(v3) > 0 {
                    v3[0] = ""
                }
                _ = fi25
            }
            v26 := GenS2{F0: 156, F1: 42}
            _ = v26
        } else {
            switch 2147483647 {
            case 2147483647:
                v3 = append(v3, (("" + "hello") + "foo bar"))
            }
        }
    }
    fmt.Println(v6)
    switch v22 {
    case -2147483648:
        for fi27 := 0; fi27 < 12; fi27++ {
            var v28 int16 = -118
            _ = fi27
            _ = v28
        }
    default:
        var v29 uint64 = 235
        _ = v29
    }
    fmt.Println(v4, v5, v22)
    for fi30 := 0; fi30 < 2; fi30++ {
        for fi31 := 0; fi31 < 13; fi31++ {
            var v32 string = "foo bar"
            v32 = "foo bar"
            _ = fi31
            _ = v32
        }
        _ = fi30
    }
    v33 := GenS2{F0: 42, F1: 127}
    for fi34 := 0; fi34 < 1; fi34++ {
        if false {
            fmt.Println(v5, v5)
            if (-1 == 42) {
                v4 += v4
            } else {
                var v35 float64 = 3.14159
                if len(v3) > 0 {
                    v3[0] = "foo bar"
                }
                v36 := GenS1{F0: 1}
                _ = v35
                _ = v36
            }
            var v37 float64 = float64(0.0)
            _ = v37
        } else {
            if (-32768 < -1) {
                var v38 bool = (-1 != -1)
                var v39 int8 = -25
                _ = v38
                _ = v39
            } else {
                if (true && true) {
                    v40 := GenS3{F0: 1}
                    _ = v40
                }
                _ = len(v1)
            }
            if false {
                v2[(v22 + (v22 << ((uint(v4)) & 63)))] = ((("" + "hello") + "hello") + (("\n\t" + "a") + ""))
                switch v22 {
                case 32767:
                    if true {
                        if false {
                            switch (^(fi34)) {
                            case 9223372036854775807:
                                fmt.Println(v22, v4, v5)
                            }
                            var v41 bool = ((true && false) && false)
                            for fi42 := 0; fi42 < 17; fi42++ {
                                v43 := GenS3{F0: 0}
                                _ = len(v2)
                                _ = fi42
                                _ = v43
                            }
                            _ = v41
                        } else {
                            v44 := GenS4{F0: 2, F1: 0}
                            if true {
                                delete(v1, (fi34 | (v22 & v22)))
                                var v45 uint8 = uint8(v5)
                                _ = v45
                            } else {
                                v22 = v22
                                fmt.Println(v22, v6)
                            }
                            _ = v44
                        }
                    } else {
                        if len(v3) > 0 {
                            v3[0] = ("\n\t" + "a")
                        }
                    }
                }
            } else {
                var v46 int8 = -16
                _ = v46
            }
        }
        fmt.Println(v4, fi34)
        v47 := GenS5{F0: 100, F1: -325, F2: 0}
        _ = fi34
        _ = v47
    }
    for fi48 := 0; fi48 < 5; fi48++ {
        v4 = v4
        v4 = (v4 & (v4 | (v4 << ((65535) & 63))))
        _ = fi48
    }
    switch ("a" + ("" + "foo bar")) {
    case "foo bar":
        v49 := GenS5{F0: 103, F1: -988, F2: 109}
        _ = v49
    case "":
        var v50 int16 = 1000
        _ = v50
    }
    if true {
        var v51 uint = 975
        _ = v51
    } else {
        _ = len(v3)
        v52 := v1[v22]
        for fi53 := 0; fi53 < 5; fi53++ {
            if true {
                _ = len(v3)
                if ((!v52) || v52) {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    if true {
                        var v54 uint16 = 42
                        for fi55 := 0; fi55 < 9; fi55++ {
                            v4 = 255
                            var v56 uint16 = 2
                            _ = fi55
                            _ = v56
                        }
                        switch "a" {
                        case "":
                            for fi57 := 0; fi57 < 10; fi57++ {
                                v3 = append(v3, "a")
                                fmt.Println(v52, fi57, v54)
                                _ = fi57
                            }
                        case "hello":
                            switch "a" {
                            case "a":
                                if false {
                                    fmt.Println(v5)
                                }
                            case "":
                                switch (v22 / (((fi53 + v22)) | 1)) {
                                case 42:
                                    v4 = v54
                                case 32767:
                                    v58 := GenS1{F0: v54}
                                    _ = v58
                                case 9223372036854775807:
                                    if v52 {
                                        switch "\n\t" {
                                        case "a":
                                            v5 += 4294967295
                                        default:
                                            v3 = append(v3, "a")
                                        }
                                        var v59 string = ""
                                        for fi60 := 0; fi60 < 11; fi60++ {
                                            v61 := GenS3{F0: 1}
                                            _ = len(v1)
                                            var v62 uint16 = 797
                                            _ = fi60
                                            _ = v61
                                            _ = v62
                                        }
                                        _ = v59
                                    } else {
                                        var v63 float64 = (((((1e10 + (0.1 + 0.1)) + (0.0 + float64(0.1))) - (-2.25 - 0.1)) * 1e10) + 1.5)
                                        _ = v63
                                    }
                                default:
                                    if false {
                                        for fi64 := 0; fi64 < 18; fi64++ {
                                            fmt.Println(v54)
                                            _ = fi64
                                        }
                                        switch (v22 | fi53) {
                                        case 9223372036854775807:
                                            var v65 uint16 = v54
                                            _ = v65
                                        }
                                        if v52 {
                                            _ = len(v3)
                                            if (1 < v5) {
                                                var v66 int64 = -1
                                                _ = v66
                                            } else {
                                                v54 = 1
                                                for fi67 := 0; fi67 < 10; fi67++ {
                                                    switch "foo bar" {
                                                    case "\n\t":
                                                        switch ((("a" + "foo bar") + (("a" + ("\n\t" + ("a" + ("\n\t" + "\n\t")))) + "")) + "") {
                                                        case "\n\t":
                                                            v5 -= v5
                                                        default:
                                                            var v68 int64 = 9223372036854775807
                                                            _ = v68
                                                        }
                                                    case "a":
                                                        var v69 uint16 = 0
                                                        _ = v69
                                                    }
                                                    var v70 uint16 = (v4 + 467)
                                                    var v71 uint32 = v5
                                                    _ = fi67
                                                    _ = v70
                                                    _ = v71
                                                }
                                            }
                                        }
                                    } else {
                                        var v72 int = 4294967295
                                        _ = v72
                                    }
                                }
                            case "\n\t":
                                var v73 float64 = float64(1.5)
                                _ = v73
                            }
                        default:
                            if true {
                                v6 = -9223372036854775808
                                v6 = (v6 | (v6 | v6))
                                fmt.Println(v4, v22, fi53)
                            }
                        }
                        _ = v54
                    }
                    var v74 int8 = 79
                    _ = v74
                } else {
                    v75 := GenS6{F0: v22}
                    _ = v75
                }
                var v76 int = (fi53 / (((fi53 ^ v22)) | 1))
                _ = v76
            } else {
                v5 += (v5 + v5)
            }
            v5 += (v5 - genFn0(v52, 1))
            switch (fi53 & v22) {
            case 128:
                var v77 int16 = 0
                _ = v77
            }
            _ = fi53
        }
        _ = v52
    }
    if true {
        fmt.Println(v5)
    } else {
        v78 := GenS7{F0: 42, F1: -70}
        _ = v78
    }
    v6 += (^(v6))
    var v79 int = v22
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v79
}
