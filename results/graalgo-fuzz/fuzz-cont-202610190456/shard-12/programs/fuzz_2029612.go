package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 int16 }
type GenS2 struct { F0 uint; F1 int32 }
type GenS3 struct { F0 uint; F1 uint16 }
type GenS4 struct { F0 uint32; F1 int32 }
type GenS5 struct { F0 uint8; F1 uint }
type GenS6 struct { F0 uint; F1 int16 }
func genFn0(p0 uint16, p1 int, p2 int32) int16 {
    var v6 int32 = p2
    var v7 string = "a"
    v8 := GenS1{F0: -1, F1: v6, F2: 0}
    for fi9 := 0; fi9 < 18; fi9++ {
        p0 -= p0
        v10 := GenS1{F0: p1, F1: p2, F2: -1}
        v7 = "a"
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 98
}
func genFn1(p0 uint) bool {
    p0 -= p0
    var v11 uint8 = 2
    _ = p0
    _ = v11
    return ((v11 / (((v11 ^ (v11 << ((18446744073709551615) & 63)))) | 1)) >= (^(v11)))
}

func main() {
    v1 := []string{}
    var v2 int8 = 123
    v3 := make(map[int]int)
    var v4 int64 = 1
    var v5 int64 = 1
    var v12 int64 = v5
    v4 = v5
    var v13 string = ((("" + ("\n\t" + "foo bar")) + ("hello" + "hello")) + "foo bar")
    for fi14 := 0; fi14 < 3; fi14++ {
        v15 := GenS1{F0: (fi14 | 4294967295), F1: -803, F2: -717}
        v16 := GenS2{F0: 255, F1: -337}
        _ = fi14
        _ = v15
        _ = v16
    }
    var v17 int16 = 0
    for fi18 := 0; fi18 < 14; fi18++ {
        switch fi18 {
        case -219:
            switch ("a" + v13) {
            case "a":
                fmt.Println(v12, v12)
            case "":
                if true {
                    if (!false) {
                        v12 -= v4
                        for fi19 := 0; fi19 < 18; fi19++ {
                            var v20 uint16 = 42
                            _ = fi19
                            _ = v20
                        }
                        if (127 < 0) {
                            v21 := GenS3{F0: 65535, F1: 720}
                            v3[42] = fi18
                            _ = v21
                        } else {
                            switch "foo bar" {
                            case "":
                                _ = len(v1)
                            case "foo bar":
                                v22 := GenS3{F0: 4294967295, F1: 1}
                                _ = v22
                            case "a":
                                switch fi18 {
                                case 32767:
                                    switch fi18 {
                                    case -2147483648:
                                        for fi23 := 0; fi23 < 8; fi23++ {
                                            v24 := GenS3{F0: 599, F1: 42}
                                            v17 += (^(v17))
                                            var v25 int8 = -1
                                            _ = fi23
                                            _ = v24
                                            _ = v25
                                        }
                                    }
                                case 42:
                                    v12 = (^(v12))
                                case -2147483648:
                                    var v26 int8 = v2
                                    _ = v26
                                }
                            }
                            var v27 uint8 = 1
                            _ = v27
                        }
                    } else {
                        switch 36 {
                        case -9223372036854775808:
                            v2 = v2
                        }
                        for fi28 := 0; fi28 < 5; fi28++ {
                            switch (fi18 * (fi18 - (fi18 / ((32767) | 1)))) {
                            case 0:
                                var v29 int = (fi18 << ((65535) & 63))
                                _ = v29
                            case 670:
                                if true {
                                    delete(v3, fi28)
                                    if (!(false || (true && false))) {
                                        _ = len(v1)
                                        v2 -= (-(v2))
                                        var v30 int = fi18
                                        var v31 uint32 = 2147483647
                                        var v32 float32 = 1.5
                                        var v33 uint64 = 42
                                        var v34 string = v13
                                        v35 := GenS1{F0: fi18, F1: -249, F2: int16(v33)}
                                        _ = v30
                                        _ = v31
                                        _ = v32
                                        _ = v33
                                        _ = v34
                                        _ = v35
                                    }
                                } else {
                                    v36 := GenS4{F0: 2147483647, F1: -2147483648}
                                    _ = v36
                                }
                                var v37 uint32 = 1
                                if (!true) {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    fmt.Println(fi28, v13)
                                    switch (v13 + "\n\t") {
                                    case "\n\t":
                                        delete(v3, fi18)
                                    case "foo bar":
                                        fmt.Println(v2)
                                    case "a":
                                        if (999 <= 4294967295) {
                                            var v38 int16 = genFn0(2, (fi18 | (fi18 & fi18)), 1)
                                            fmt.Println(v4, v17, v5)
                                            v39 := GenS5{F0: 0, F1: 42}
                                            _ = v38
                                            _ = v39
                                        } else {
                                            if false {
                                                switch v13 {
                                                case "\n\t":
                                                    delete(v3, (^(fi28)))
                                                default:
                                                    v17 -= int16(v37)
                                                }
                                            } else {
                                                _ = len(v3)
                                                var v40 float32 = 0.0
                                                _ = v40
                                            }
                                            v5 -= (v5 + v12)
                                        }
                                    default:
                                        v37 -= (v37 & 24)
                                    }
                                    v41 := GenS6{F0: 97, F1: v17}
                                    _ = v41
                                } else {
                                    v4 = v5
                                    if len(v1) > 0 {
                                        v1[0] = "foo bar"
                                    }
                                }
                                _ = v37
                            }
                            v17 += (v17 % (((v17 ^ int16(v4))) | 1))
                            if (int32(v17) < 0) {
                                v17 -= genFn0(0, fi28, -1)
                                var v42 string = ("\n\t" + "\n\t")
                                _ = v42
                            }
                            _ = fi28
                        }
                    }
                    delete(v3, (fi18 % ((int(v2)) | 1)))
                    v5 += v5
                }
            case "foo bar":
                if true {
                    v43 := GenS4{F0: 4294967295, F1: 1}
                    var v44 int = 42
                    var v45 bool = (691 > -692)
                    _ = v43
                    _ = v44
                    _ = v45
                }
            }
        case 9223372036854775807:
            var v46 int = fi18
            _ = v46
        default:
            if len(v1) > 0 {
                v1[0] = ((v13 + ((v13 + v13) + v13)) + v13)
            }
        }
        switch fi18 {
        case 32768:
            if false {
                v47 := GenS5{F0: 6, F1: 255}
                _ = v47
            } else {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v48 int64 = v4
                switch v13 {
                case "\n\t":
                    switch -2147483648 {
                    case 9223372036854775807:
                        var v49 bool = ((1.5 - (1.5 * (((1.5 / 1.5) + (-2.25 + ((1e10 * 1.5) + 1e10))) * 3.14159))) >= -2.25)
                        _ = v49
                    default:
                        var v50 int = (-(fi18))
                        _ = v50
                    }
                case "a":
                    v17 = (v17 >> ((1) & 63))
                default:
                    for fi51 := 0; fi51 < 4; fi51++ {
                        v52 := GenS3{F0: 2, F1: 678}
                        var v53 int64 = (-(v12))
                        _ = fi51
                        _ = v52
                        _ = v53
                    }
                }
                _ = v48
            }
        default:
            switch ("" + "foo bar") {
            case "a":
                if (18446744073709551615 != 443) {
                    _ = len(v3)
                }
            case "":
                v4 -= (v12 * v4)
            }
        }
        var v54 int32 = -297
        _ = fi18
        _ = v54
    }
    v1 = append(v1, (v13 + ""))
    var v55 uint = 2
    if (!(!true)) {
        v13 = ("\n\t" + (v13 + v13))
    }
    for fi56 := 0; fi56 < 6; fi56++ {
        v57 := GenS2{F0: v55, F1: 0}
        var v58 int = 42
        _ = fi56
        _ = v57
        _ = v58
    }
    var v59 int8 = 69
    if (0 < 650) {
        var v60 string = v13
        var v61 int32 = 0
        _ = v60
        _ = v61
    }
    var v62 int32 = 276
    if (!true) {
        if (!(!(false && false))) {
            var v63 int8 = (v59 >> ((v55) & 63))
            switch "hello" {
            case "":
                fmt.Println(v59)
            case "hello":
                for fi64 := 0; fi64 < 18; fi64++ {
                    v65 := GenS4{F0: 2147483647, F1: -1}
                    switch int(v55) {
                    case 32767:
                        if (!false) {
                            fmt.Println(v63)
                            if true {
                                var v66 int16 = v17
                                var v67 uint16 = 42
                                _ = v66
                                _ = v67
                            }
                        } else {
                            var v68 int64 = int64(v62)
                            _ = v68
                        }
                    case -2147483648:
                        var v69 string = v13
                        _ = v69
                    }
                    switch "foo bar" {
                    case "foo bar":
                        var v70 int64 = v5
                        _ = v70
                    default:
                        var v71 float32 = 1.5
                        _ = v71
                    }
                    _ = fi64
                    _ = v65
                }
            case "a":
                fmt.Println(v63)
            default:
                v5 += (v5 >> (((v55 % (((v55 << ((v55) & 63))) | 1))) & 63))
            }
            _ = v63
        } else {
            v72 := GenS6{F0: (^(v55)), F1: v17}
            _ = v72
        }
        var v73 float32 = (((float32(3.14159) / 1.5) - 3.14159) - (3.14159 - -2.25))
        for fi74 := 0; fi74 < 3; fi74++ {
            for fi75 := 0; fi75 < 19; fi75++ {
                var v76 int16 = (v17 * (v17 & (v17 >> ((v55) & 63))))
                _ = fi75
                _ = v76
            }
            _ = len(v1)
            fmt.Println(fi74)
            _ = fi74
        }
        _ = v73
    }
    var v77 int = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("v59:", v59)
    fmt.Println("v62:", v62)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v12
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v59
    _ = v62
    _ = v77
}
