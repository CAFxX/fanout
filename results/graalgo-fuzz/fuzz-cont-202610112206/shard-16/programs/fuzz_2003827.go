package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 uint64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 bool; F1 bool; F2 bool }
type GenS4 struct { F0 int32; F1 bool; F2 int32 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int8; F1 int16; F2 uint16 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int16; F1 string; F2 uint }
type GenS9 struct { F0 string; F1 int32; F2 uint8 }
type GenS10 struct { F0 int8; F1 uint }
func genFn0(p0 uint, p1 uint, p2 int8) uint16 {
    p1 -= p0
    for fi4 := 0; fi4 < 10; fi4++ {
        for fi5 := 0; fi5 < 12; fi5++ {
            if (0 > 2) {
                v6 := GenS1{F0: 610, F1: 619, F2: 42}
                p0 -= p0
                var v7 uint = 1
                var v8 string = "foo bar"
                _ = v6
                _ = v7
                _ = v8
            }
            switch fi4 {
            case -2147483648:
                v9 := []int{}
                _ = v9
            case 4294967295:
                switch (fi5 - int(p2)) {
                case -2147483648:
                    var v10 int8 = (p2 % (((p2 << ((1) & 63))) | 1))
                    _ = v10
                case 20:
                    switch (fi5 << (((p0 | 65535)) & 63)) {
                    case 0:
                        p1 -= p0
                    case 1:
                        v11 := []int{(fi4 / (((fi5 >> ((p0) & 63))) | 1)), fi4, fi4, (fi4 ^ 127)}
                        _ = v11
                    default:
                        v12 := GenS1{F0: 2, F1: 4294967295, F2: 38}
                        _ = v12
                    }
                case 127:
                    p1 = p1
                default:
                    var v13 int8 = (p2 + 1)
                    _ = v13
                }
            default:
                fmt.Println(p0, p0, fi5)
            }
            v14 := []int{fi5, (fi4 - (fi5 + fi4)), fi5}
            p2 = p2
            if (false && true) {
                p1 -= p1
                var v15 uint8 = 127
                fmt.Println(v15)
                _ = v15
            } else {
                var v16 uint32 = 4294967295
                v17 := GenS1{F0: 42, F1: 176, F2: 206}
                switch "" {
                case "hello":
                    var v18 uint64 = 0
                    _ = v18
                case "a":
                    if ((v16 + v16) < v16) {
                        if true {
                            v16 += v16
                        } else {
                            var v19 float32 = 1.5
                            switch (-(fi5)) {
                            case 128:
                                v20 := GenS2{F0: 0}
                                _ = v20
                            case 32767:
                                fmt.Println(fi4)
                            }
                            if true {
                                var v21 int8 = p2
                                _ = v21
                            } else {
                                p0 = (p0 / (((^(p1))) | 1))
                                switch "a" {
                                case "\n\t":
                                    if false {
                                        var v22 string = ("foo bar" + (("\n\t" + ((("hello" + "hello") + "") + "")) + ""))
                                        fmt.Println(p1)
                                        _ = v22
                                    } else {
                                        v23 := GenS3{F0: (v16 <= v16), F1: true, F2: true}
                                        v24 := GenS4{F0: -639, F1: ((("foo bar" + "hello") + "a") >= ""), F2: 620}
                                        _ = v23
                                        _ = v24
                                    }
                                case "hello":
                                    var v25 int = (fi4 & fi4)
                                    _ = v25
                                default:
                                    v26 := GenS5{F0: 0}
                                    _ = len(v14)
                                    fmt.Println(p2)
                                    _ = v26
                                }
                                if (81 <= 2) {
                                    var v27 int = (fi5 | fi5)
                                    p0 += (p1 + (p1 | p0))
                                    _ = v27
                                } else {
                                    if (true || true) {
                                        if len(v14) > 0 {
                                            v14[0] = fi4
                                        }
                                        v14 = append(v14, int(p1))
                                    }
                                    var v28 int = (fi4 << ((42) & 63))
                                    var v29 float64 = 0.0
                                    if len(v14) >= 2 {
                                        v14 = v14[1:len(v14)]
                                    }
                                    v14 = append(v14, (fi5 >> ((p0) & 63)))
                                    p0 -= p0
                                    v19 = v19
                                    _ = v28
                                    _ = v29
                                }
                                v14 = append(v14, fi5)
                                var v30 string = "\n\t"
                                p1 = 450
                                _ = v30
                            }
                            _ = v19
                        }
                        v31 := GenS6{F0: -128, F1: 1000, F2: 0}
                        v14 = append(v14, fi4)
                        _ = v31
                    } else {
                        switch "" {
                        case "\n\t":
                            if ((1.5 / (0.1 * -2.25)) >= ((3.14159 * -2.25) - float32((0.0 - 1e10)))) {
                                var v32 uint32 = 1
                                switch -1 {
                                case 32768:
                                    fmt.Println(p1, fi4, p1)
                                case -1:
                                    var v33 float64 = 0.0
                                    _ = v33
                                }
                                v34 := GenS5{F0: 720}
                                _ = v32
                                _ = v34
                            } else {
                                var v35 int = (fi5 + (fi5 * fi4))
                                _ = v35
                            }
                        case "foo bar":
                            var v36 int = (fi5 ^ 32768)
                            _ = v36
                        default:
                            var v37 int8 = (p2 << ((p0) & 63))
                            _ = v37
                        }
                        v16 = v16
                    }
                    if len(v14) >= 2 {
                        v14 = v14[1:len(v14)]
                    }
                }
                _ = v16
                _ = v17
            }
            if true {
                p1 = p1
                var v38 int = (-(fi4))
                fmt.Println(p2, p1)
                _ = v38
            }
            _ = fi5
            _ = v14
        }
        _ = fi4
    }
    if false {
        p2 -= p2
        if (false || true) {
            if false {
                var v39 uint = p0
                _ = v39
            } else {
                fmt.Println(p2, p2, p1)
                var v40 string = "a"
                var v41 float32 = float32(1.5)
                _ = v40
                _ = v41
            }
            for fi42 := 0; fi42 < 12; fi42++ {
                var v43 int16 = 1000
                v44 := []int{-2147483648}
                _ = fi42
                _ = v43
                _ = v44
            }
        }
    }
    v45 := GenS3{F0: (128 <= 128), F1: (!(true && true)), F2: true}
    _ = p0
    _ = p1
    _ = p2
    _ = v45
    return 42
}

func main() {
    var v1 float32 = -2.25
    var v2 string = "foo bar"
    var v3 float64 = (0.1 * (-2.25 * (1e10 * (3.14159 * (1.5 / ((3.14159 * 0.0) - (-2.25 - float64((v1 - v1)))))))))
    var v46 int8 = 42
    fmt.Println(v46, v46)
    if true {
        fmt.Println(v3)
        for fi47 := 0; fi47 < 4; fi47++ {
            switch ((v2 + v2) + "a") {
            case "hello":
                if false {
                    for fi48 := 0; fi48 < 8; fi48++ {
                        v49 := []int{}
                        _ = fi48
                        _ = v49
                    }
                    v50 := GenS7{F0: 127}
                    _ = v50
                }
            case "":
                v2 = v2
            case "a":
                var v51 float64 = 3.14159
                _ = v51
            }
            _ = fi47
        }
        for fi52 := 0; fi52 < 10; fi52++ {
            var v53 uint64 = 2
            v53 = v53
            _ = fi52
            _ = v53
        }
    }
    v1 -= (-2.25 + v1)
    var v54 uint = 0
    for fi55 := 0; fi55 < 14; fi55++ {
        for fi56 := 0; fi56 < 3; fi56++ {
            v3 = v3
            fmt.Println(v1, fi56, v3)
            var v57 float32 = 1e10
            _ = fi56
            _ = v57
        }
        if true {
            switch v2 {
            case "hello":
                v58 := []int{-709, (fi55 | (fi55 | fi55)), fi55}
                _ = v58
            }
            var v59 uint64 = 1
            _ = v59
        } else {
            v1 -= 1.5
            v60 := []int{fi55, (fi55 << ((v54) & 63))}
            _ = v60
        }
        var v61 int = fi55
        _ = fi55
        _ = v61
    }
    if (371 > -1) {
        v62 := GenS8{F0: 1, F1: v2, F2: v54}
        var v63 int16 = 1
        var v64 int16 = -1
        _ = v62
        _ = v63
        _ = v64
    }
    var v65 bool = true
    for fi66 := 0; fi66 < 1; fi66++ {
        v67 := GenS8{F0: 32767, F1: v2, F2: v54}
        if (128 != 0) {
            switch v2 {
            case "foo bar":
                v68 := GenS9{F0: (v2 + v2), F1: 42, F2: 1}
                _ = v68
            case "hello":
                if true {
                    if (4294967295 >= 2) {
                        var v69 uint8 = 51
                        _ = v69
                    } else {
                        switch v2 {
                        case "a":
                            v46 += v46
                        case "":
                            var v70 float64 = 0.1
                            _ = v70
                        case "hello":
                            for fi71 := 0; fi71 < 2; fi71++ {
                                fmt.Println(v65, v46)
                                _ = fi71
                            }
                        }
                    }
                } else {
                    var v72 int16 = -1
                    var v73 int32 = 0
                    _ = v72
                    _ = v73
                }
            default:
                for fi74 := 0; fi74 < 12; fi74++ {
                    if true {
                        var v75 int16 = 696
                        _ = v75
                    }
                    _ = fi74
                }
            }
            fmt.Println(v2)
            var v76 int = -764
            _ = v76
        } else {
            switch fi66 {
            case 128:
                var v77 string = "hello"
                _ = v77
            case -2147483648:
                var v78 uint32 = 1
                _ = v78
            case -1:
                if v65 {
                    v3 += ((v3 / (v3 + ((((1e10 - 0.1) / (v3 + v3)) + (-2.25 + ((v3 * (float64(0.1) - 0.1)) - 0.0))) + (3.14159 / v3)))) - v3)
                    var v79 int = fi66
                    _ = v79
                } else {
                    if true {
                        switch "\n\t" {
                        case "foo bar":
                            v2 = "a"
                        case "\n\t":
                            v54 = (v54 ^ (v54 << (((v54 + v54)) & 63)))
                        }
                    }
                }
            default:
                fmt.Println(v54, v65)
            }
            for fi80 := 0; fi80 < 7; fi80++ {
                v81 := []int{}
                v2 = "\n\t"
                v46 -= 127
                _ = fi80
                _ = v81
            }
        }
        _ = fi66
        _ = v67
    }
    fmt.Println(v65)
    v82 := []int{1}
    var v83 uint16 = 586
    v3 = (0.0 / v3)
    v82 = append(v82, 670)
    for fi84 := 0; fi84 < 3; fi84++ {
        fmt.Println(v65)
        v85 := GenS10{F0: (^(v46)), F1: (v54 | v54)}
        _ = fi84
        _ = v85
    }
    if (!true) {
        var v86 float32 = float32(v3)
        v65 = v65
        _ = v86
    }
    for fi87 := 0; fi87 < 4; fi87++ {
        fmt.Println(fi87, v1, v3)
        v1 = v1
        v88 := GenS3{F0: (-832 > 42), F1: false, F2: v65}
        _ = fi87
        _ = v88
    }
    var v89 string = "a"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v46:", v46)
    fmt.Println("v54:", v54)
    fmt.Println("v65:", v65)
    fmt.Println("slicelen v82:", len(v82), v82)
    fmt.Println("v83:", v83)
    fmt.Println("v89:", v89)
    _ = v1
    _ = v2
    _ = v3
    _ = v46
    _ = v54
    _ = v65
    _ = v82
    _ = v83
    _ = v89
}
