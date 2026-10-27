package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 bool; F1 uint64 }
type GenS3 struct { F0 bool; F1 int16 }
type GenS4 struct { F0 int64; F1 int32; F2 int32 }
type GenS5 struct { F0 int; F1 uint; F2 uint8 }
type GenS6 struct { F0 uint32; F1 bool; F2 uint8 }
func genFn0(p0 uint8, p1 int16, p2 uint) uint32 {
    if false {
        v4 := GenS1{F0: false, F1: 2}
        _ = v4
    } else {
        var v5 int = 0
        var v6 int = -562
        for fi7 := 0; fi7 < 5; fi7++ {
            var v8 uint = p2
            var v9 int64 = 1
            fmt.Println(v8, v8)
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = v5
        _ = v6
    }
    v10 := GenS1{F0: false, F1: 757}
    p1 = (p1 & (^(p1)))
    v11 := []int{-2147483648}
    switch -208 {
    case 4294967295:
        fmt.Println(p1, p0)
    case -1:
        switch -44 {
        case 1:
            v12 := GenS2{F0: (-7 >= -128), F1: 2}
            _ = v12
        case -562:
            if (true || (false || true)) {
                var v13 uint = (p2 / ((p2) | 1))
                if len(v11) > 0 {
                    v11[0] = -23
                }
                _ = v13
            }
        }
    case 1:
        switch "a" {
        case "foo bar":
            var v14 uint16 = 42
            _ = v14
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return 974
}
func genFn1(p0 uint32, p1 uint32) int32 {
    switch (("" + "hello") + ("hello" + "\n\t")) {
    case "":
        if (1000 > 1000) {
            p0 = (p1 + (p1 & 65535))
        }
    case "\n\t":
        v15 := GenS1{F0: false, F1: uint64(p0)}
        _ = v15
    case "a":
        p0 = (^(p1))
    }
    p0 -= p1
    for fi16 := 0; fi16 < 6; fi16++ {
        var v17 int8 = 0
        _ = fi16
        _ = v17
    }
    for fi18 := 0; fi18 < 15; fi18++ {
        p1 += (p1 >> ((562) & 63))
        _ = fi18
    }
    _ = p0
    _ = p1
    return -774
}

func main() {
    var v1 bool = true
    var v2 uint64 = 42
    v3 := make(map[int]bool)
    fmt.Println(v1, v2)
    fmt.Println(v2)
    if (2147483648 <= 1) {
        switch "\n\t" {
        case "\n\t":
            for fi19 := 0; fi19 < 13; fi19++ {
                v20 := GenS1{F0: (!(v1 || (fi19 > fi19))), F1: v2}
                if (42 > 128) {
                    switch (-(fi19)) {
                    case 1:
                        for fi21 := 0; fi21 < 2; fi21++ {
                            v1 = v1
                            switch int(v2) {
                            case 1:
                                v22 := []int{(fi21 | (fi19 - fi19))}
                                _ = v22
                            case -9223372036854775808:
                                fmt.Println(fi19, fi21, fi21)
                            default:
                                fmt.Println(fi21, v1, v2)
                                var v23 bool = false
                                v2 -= v2
                                switch ("a" + ("\n\t" + "foo bar")) {
                                case "\n\t":
                                    if false {
                                        if (!((-666 > -32768) || (0 == -982))) {
                                            fmt.Println(fi21, fi21, v1)
                                            var v24 uint32 = 65535
                                            _ = v24
                                        } else {
                                            var v25 int64 = 9223372036854775807
                                            _ = v25
                                        }
                                        v26 := GenS1{F0: v1, F1: 2}
                                        v1 = v23
                                        _ = v26
                                    } else {
                                        v23 = (148 < 127)
                                        if true {
                                            var v27 int32 = 2147483647
                                            _ = v27
                                        } else {
                                            v28 := GenS1{F0: (v23 || (false && (!v1))), F1: (v2 & v2)}
                                            _ = v28
                                        }
                                    }
                                case "":
                                    v1 = v1
                                default:
                                    _ = len(v3)
                                }
                                _ = v23
                            }
                            v29 := GenS2{F0: ((v2 - v2) >= (v2 * (v2 * v2))), F1: (v2 % ((985) | 1))}
                            _ = fi21
                            _ = v29
                        }
                    case -692:
                        var v30 int8 = 100
                        _ = v30
                    case -1:
                        for fi31 := 0; fi31 < 13; fi31++ {
                            v32 := []int{(fi19 + fi19), (^(fi19)), (fi19 / ((fi19) | 1)), fi19}
                            _ = fi31
                            _ = v32
                        }
                    }
                }
                v33 := GenS3{F0: (!((v2 | v2) < (v2 << ((uint(v2)) & 63)))), F1: -32768}
                _ = fi19
                _ = v20
                _ = v33
            }
        case "":
            v2 -= v2
        }
        fmt.Println(v2, v2)
        for fi34 := 0; fi34 < 20; fi34++ {
            var v35 uint32 = 963
            if true {
                v36 := v3[fi34]
                v2 = v2
                var v37 float64 = (((3.14159 * -2.25) / -2.25) - (float64(((0.1 + (float32(((-2.25 * 1e10) - 1.5)) * 0.0)) + 0.1)) - 0.0))
                _ = v36
                _ = v37
            }
            v1 = false
            _ = fi34
            _ = v35
        }
    } else {
        for fi38 := 0; fi38 < 10; fi38++ {
            v39 := GenS2{F0: (!(669 <= 1)), F1: v2}
            var v40 uint = 587
            _ = fi38
            _ = v39
            _ = v40
        }
        v1 = (v1 && v1)
        var v41 int = -912
        _ = v41
    }
    var v42 int = 414
    for fi43 := 0; fi43 < 9; fi43++ {
        var v44 bool = v1
        v45 := []int{fi43, v42}
        _ = fi43
        _ = v44
        _ = v45
    }
    var v46 uint16 = 0
    switch -9223372036854775808 {
    case -9223372036854775808:
        v47 := GenS3{F0: true, F1: -32768}
        _ = v47
    case 32768:
        for fi48 := 0; fi48 < 12; fi48++ {
            v46 -= (v46 - v46)
            var v49 float32 = 3.14159
            _ = fi48
            _ = v49
        }
    case 2147483647:
        var v50 uint64 = v2
        _ = v50
    }
    for fi51 := 0; fi51 < 10; fi51++ {
        if (!false) {
            switch (v42 % ((fi51) | 1)) {
            case 42:
                if (!(2147483647 < 0)) {
                    v52 := GenS1{F0: ((-2.25 - (0.1 / (float64((((1e10 / -2.25) + -2.25) / 0.1)) - -2.25))) < 1e10), F1: v2}
                    var v53 uint64 = (v2 ^ 42)
                    if (false || v1) {
                        switch fi51 {
                        case -6:
                            delete(v3, (v42 - fi51))
                        case -1:
                            var v54 int32 = -947
                            _ = v54
                        default:
                            for fi55 := 0; fi55 < 17; fi55++ {
                                if (((float32((0.1 - 0.1)) - 0.0) + -2.25) <= (3.14159 + ((((1.5 / 0.1) - 3.14159) - 1.5) * -2.25))) {
                                    fmt.Println(v1, fi55)
                                    if v1 {
                                        var v56 uint64 = (v53 << ((0) & 63))
                                        switch (^(v42)) {
                                        case 32767:
                                            var v57 int64 = 0
                                            _ = v57
                                        default:
                                            v58 := []int{4294967295, (v42 / ((fi51) | 1)), -163, (v42 % ((90) | 1))}
                                            _ = v58
                                        }
                                        _ = v56
                                    } else {
                                        var v59 uint = 545
                                        var v60 int = fi55
                                        if (!(int8(v59) < 124)) {
                                            v46 = v46
                                            v61 := GenS2{F0: true, F1: v2}
                                            fmt.Println(v1, v53)
                                            _ = v61
                                        }
                                        _ = v59
                                        _ = v60
                                    }
                                    v1 = false
                                }
                                _ = fi55
                            }
                        }
                        _ = len(v3)
                    } else {
                        v42 += -665
                        for fi62 := 0; fi62 < 14; fi62++ {
                            v63 := GenS1{F0: false, F1: (v2 >> ((42) & 63))}
                            v64 := GenS4{F0: 1, F1: 2147483647, F2: 42}
                            v53 = 42
                            v65 := v3[v42]
                            _ = fi62
                            _ = v63
                            _ = v64
                            _ = v65
                        }
                    }
                    _ = v52
                    _ = v53
                } else {
                    var v66 string = "foo bar"
                    if ((v42 | (fi51 ^ (v42 >> ((255) & 63)))) != v42) {
                        v66 = (v66 + v66)
                    }
                    _ = v66
                }
            case 0:
                for fi67 := 0; fi67 < 3; fi67++ {
                    v46 -= 2
                    if (!(-128 != 0)) {
                        v42 = (-(v42))
                        var v68 int8 = 127
                        var v69 float64 = 0.0
                        _ = v68
                        _ = v69
                    }
                    _ = fi67
                }
            default:
                var v70 int64 = -1
                _ = v70
            }
            v71 := GenS2{F0: v1, F1: v2}
            _ = v71
        } else {
            v72 := GenS5{F0: v42, F1: 2, F2: 17}
            var v73 int64 = 1
            for fi74 := 0; fi74 < 7; fi74++ {
                var v75 uint = 255
                _ = fi74
                _ = v75
            }
            _ = v72
            _ = v73
        }
        v76 := GenS6{F0: 2147483647, F1: true, F2: 128}
        if (42 == 2147483647) {
            fmt.Println(v42, v1, v42)
        } else {
            fmt.Println(v2, v42, v1)
        }
        _ = fi51
        _ = v76
    }
    fmt.Println(v42, v1)
    v42 = -760
    var v77 bool = (0 == 0)
    for fi78 := 0; fi78 < 4; fi78++ {
        v46 -= (v46 * (v46 >> ((1) & 63)))
        fmt.Println(v1)
        _ = fi78
    }
    var v79 float64 = (1.5 - (-2.25 / 1.5))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v42:", v42)
    fmt.Println("v46:", v46)
    fmt.Println("v77:", v77)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v2
    _ = v3
    _ = v42
    _ = v46
    _ = v77
    _ = v79
}
