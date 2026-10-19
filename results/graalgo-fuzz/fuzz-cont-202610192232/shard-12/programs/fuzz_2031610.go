package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 int16; F1 uint8; F2 uint64 }
type GenS3 struct { F0 int32; F1 uint; F2 uint32 }
func genFn0() int16 {
    for fi5 := 0; fi5 < 15; fi5++ {
        var v6 uint16 = 1
        _ = fi5
        _ = v6
    }
    var v7 int = 32768
    if false {
        var v8 int16 = 42
        v8 = v8
        _ = v8
    } else {
        v7 = v7
    }
    for fi9 := 0; fi9 < 12; fi9++ {
        if false {
            for fi10 := 0; fi10 < 4; fi10++ {
                fmt.Println(v7)
                if (!(719 >= 0)) {
                    v11 := []int{(v7 - fi9)}
                    v7 -= v7
                    _ = v11
                } else {
                    v12 := []int{127, fi10, (fi9 >> ((255) & 63))}
                    var v13 uint = 168
                    _ = v12
                    _ = v13
                }
                v14 := []int{(^(v7)), v7, (fi9 << ((18446744073709551615) & 63))}
                _ = fi10
                _ = v14
            }
            var v15 float32 = 1.5
            fmt.Println(fi9)
            _ = v15
        }
        var v16 string = ("foo bar" + ((("" + "foo bar") + ("a" + "a")) + (("\n\t" + "") + "a")))
        var v17 uint8 = 0
        _ = fi9
        _ = v16
        _ = v17
    }
    _ = v7
    return -682
}
func genFn1(p0 string) uint64 {
    for fi18 := 0; fi18 < 11; fi18++ {
        if true {
            p0 = ""
            v19 := []int{fi18, fi18}
            for fi20 := 0; fi20 < 14; fi20++ {
                p0 = (p0 + (((((p0 + (p0 + ("foo bar" + (p0 + p0)))) + (p0 + p0)) + p0) + "a") + p0))
                _ = fi20
            }
            _ = v19
        } else {
            p0 = (p0 + "")
            var v21 float32 = 1.5
            v21 = -2.25
            _ = v21
        }
        var v22 int8 = 1
        v22 = v22
        _ = fi18
        _ = v22
    }
    if (false || false) {
        for fi23 := 0; fi23 < 19; fi23++ {
            v24 := []int{}
            if len(v24) > 0 {
                v24[0] = (fi23 ^ 9223372036854775807)
            }
            _ = fi23
            _ = v24
        }
    } else {
        for fi25 := 0; fi25 < 17; fi25++ {
            var v26 uint64 = 18446744073709551615
            for fi27 := 0; fi27 < 12; fi27++ {
                switch "" {
                case "":
                    var v28 int16 = 965
                    _ = v28
                default:
                    if false {
                        v26 -= (v26 % ((834) | 1))
                    } else {
                        fmt.Println(fi27, fi25, fi27)
                    }
                }
                _ = fi27
            }
            _ = fi25
            _ = v26
        }
        p0 = p0
        var v29 uint16 = 32767
        _ = v29
    }
    for fi30 := 0; fi30 < 8; fi30++ {
        var v31 uint8 = 42
        v32 := GenS1{F0: -1, F1: 42}
        var v33 float64 = 0.0
        _ = fi30
        _ = v31
        _ = v32
        _ = v33
    }
    _ = p0
    return 895
}

func main() {
    v1 := []int{91, 391, -212}
    v2 := make(map[string]string)
    v3 := make(map[string]uint32)
    v4 := make(map[uint32]string)
    for fi34 := 0; fi34 < 16; fi34++ {
        for fi35 := 0; fi35 < 1; fi35++ {
            switch fi35 {
            case -808:
                var v36 uint8 = 1
                var v37 float32 = (1.5 + 1e10)
                _ = v36
                _ = v37
            default:
                if false {
                    var v38 bool = (185 > 0)
                    fmt.Println(v38)
                    v38 = v38
                    v38 = (-755 == 548)
                    var v39 uint8 = 175
                    if (!false) {
                        fmt.Println(fi34)
                        fmt.Println(v39)
                    } else {
                        switch "foo bar" {
                        case "\n\t":
                            v39 = v39
                        case "foo bar":
                            v40 := GenS1{F0: 9223372036854775807, F1: 42}
                            _ = v40
                        case "a":
                            switch (fi35 * -1) {
                            case 516:
                                var v41 string = (("hello" + "a") + "\n\t")
                                _ = v41
                            default:
                                var v42 uint32 = 0
                                _ = v42
                            }
                        }
                    }
                    _ = v38
                    _ = v39
                }
            }
            v43 := GenS1{F0: -1, F1: 743}
            var v44 uint16 = 263
            _ = len(v2)
            fmt.Println(v44, v44)
            _ = fi35
            _ = v43
            _ = v44
        }
        _ = fi34
    }
    v45 := GenS1{F0: 0, F1: 4294967295}
    for fi46 := 0; fi46 < 12; fi46++ {
        if false {
            var v47 int16 = 1
            if len(v1) > 0 {
                v1[0] = (fi46 ^ fi46)
            }
            _ = v47
        } else {
            var v48 uint = 858
            v48 = (v48 / ((v48) | 1))
            _ = v48
        }
        v49 := v4[65535]
        switch (v49 + (v49 + v49)) {
        case "a":
            if false {
                if true {
                    v50 := GenS1{F0: 0, F1: genFn1("\n\t")}
                    _ = len(v1)
                    v49 = v49
                    _ = v50
                } else {
                    delete(v3, v49)
                    if false {
                        v49 = v49
                    }
                }
                switch (((v49 + v49) + v49) + ("" + v49)) {
                case "hello":
                    v51 := v4[42]
                    _ = v51
                default:
                    var v52 uint32 = 670
                    _ = v52
                }
            } else {
                v53 := GenS2{F0: -32768, F1: 0, F2: 1}
                if (-1 < -9223372036854775808) {
                    v1 = append(v1, (fi46 % (((fi46 * fi46)) | 1)))
                    var v54 bool = ((1 == 0) && false)
                    _ = v54
                } else {
                    for fi55 := 0; fi55 < 8; fi55++ {
                        v49 = ("\n\t" + v49)
                        _ = fi55
                    }
                    delete(v2, "a")
                    var v56 int32 = -1
                    _ = v56
                }
                for fi57 := 0; fi57 < 10; fi57++ {
                    if (-2.25 <= 3.14159) {
                        if (v49 > (v49 + (v49 + v49))) {
                            var v58 bool = false
                            var v59 uint64 = 212
                            v60 := GenS3{F0: 42, F1: 162, F2: 619}
                            _ = v58
                            _ = v59
                            _ = v60
                        }
                        var v61 float32 = 1.5
                        _ = v61
                    } else {
                        v49 = (v49 + "\n\t")
                        var v62 uint64 = 1
                        _ = v62
                    }
                    v49 = ("" + ("" + ""))
                    v63 := GenS1{F0: 9223372036854775807, F1: 4294967295}
                    _ = fi57
                    _ = v63
                }
                _ = v53
            }
        case "hello":
            var v64 uint = 4294967295
            _ = v64
        case "\n\t":
            for fi65 := 0; fi65 < 13; fi65++ {
                _ = len(v2)
                _ = fi65
            }
        }
        _ = fi46
        _ = v49
    }
    var v66 uint16 = 156
    for fi67 := 0; fi67 < 12; fi67++ {
        var v68 int32 = 113
        v69 := GenS3{F0: 0, F1: 18446744073709551615, F2: 65535}
        v68 += (v68 - v68)
        _ = fi67
        _ = v68
        _ = v69
    }
    v1 = append(v1, 128)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v66 = 824
    var v70 uint32 = 1
    switch ("foo bar" + "\n\t") {
    case "":
        fmt.Println(v66)
    case "\n\t":
        var v71 int64 = 0
        _ = v71
    case "foo bar":
        fmt.Println(v66)
    default:
        v2["hello"] = "\n\t"
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v66:", v66)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v66
    _ = v70
}
