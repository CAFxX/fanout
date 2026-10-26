package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 bool; F1 uint64; F2 int }
type GenS4 struct { F0 uint32; F1 uint32 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 string; F1 bool }
func genFn0() uint64 {
    var v7 bool = (!(2 != 65535))
    v8 := []int{830, 10, -64}
    var v9 float64 = (-2.25 / 3.14159)
    if v7 {
        switch "" {
        case "foo bar":
            var v10 int64 = 0
            _ = v10
        }
        var v11 uint = 1
        if (2 < 65535) {
            switch ("a" + "foo bar") {
            case "":
                if (v7 && false) {
                    var v12 string = "\n\t"
                    if (-32768 != 1) {
                        var v13 bool = (v7 || (!v7))
                        _ = v13
                    } else {
                        v14 := GenS1{F0: 18, F1: (((((v12 + v12) + v12) + (v12 + (v12 + v12))) + v12) + v12), F2: 541}
                        _ = v14
                    }
                    var v15 int16 = 195
                    _ = v12
                    _ = v15
                } else {
                    v9 = (((v9 - ((float64((1e10 + (0.0 - ((-2.25 - float32((v9 - v9))) * (((-2.25 * 3.14159) / (((3.14159 + float32(v9)) + 1.5) - (((((0.0 + ((0.1 + 0.1) / 1.0)) + (1.5 + 0.0)) * 1e10) - 0.1) / 1.5))) + 1e10))))) - v9) + v9)) * v9) * v9)
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                }
            case "foo bar":
                v7 = (42 > 191)
            case "hello":
                var v16 float64 = (1.5 * (v9 - -2.25))
                _ = v16
            }
            if len(v8) > 0 {
                v8[0] = 32767
            }
        } else {
            var v17 uint64 = 912
            _ = v17
        }
        _ = v11
    } else {
        if v7 {
            switch "a" {
            case "foo bar":
                var v18 int = -726
                _ = v18
            case "\n\t":
                var v19 float32 = 1e10
                _ = v19
            default:
                switch 42 {
                case 0:
                    if (v9 == (((0.1 - v9) - 0.0) * 0.1)) {
                        if (!true) {
                            _ = len(v8)
                            var v20 int64 = -1
                            _ = v20
                        }
                        if false {
                            var v21 float64 = v9
                            if (false && true) {
                                if (v7 && (38 > 0)) {
                                    _ = len(v8)
                                    if (-2.25 != 0.1) {
                                        var v22 uint32 = 0
                                        _ = v22
                                    }
                                    if (1e10 != (0.1 / 1e10)) {
                                        var v23 uint32 = 885
                                        if (309 <= 1) {
                                            for fi24 := 0; fi24 < 8; fi24++ {
                                                if ((fi24 % ((1) | 1)) > fi24) {
                                                    fmt.Println(v7)
                                                } else {
                                                    var v25 int8 = 22
                                                    var v26 uint16 = 898
                                                    for fi27 := 0; fi27 < 20; fi27++ {
                                                        fmt.Println(fi24)
                                                        v7 = true
                                                        var v28 uint32 = (v23 + v23)
                                                        if len(v8) >= 2 {
                                                            v8 = v8[1:len(v8)]
                                                        }
                                                        v29 := GenS1{F0: 1, F1: "hello", F2: 1}
                                                        _ = fi27
                                                        _ = v28
                                                        _ = v29
                                                    }
                                                    _ = v25
                                                    _ = v26
                                                }
                                                _ = fi24
                                            }
                                            for fi30 := 0; fi30 < 12; fi30++ {
                                                v23 += v23
                                                if v7 {
                                                    v31 := GenS1{F0: 54, F1: "foo bar", F2: 32767}
                                                    _ = v31
                                                }
                                                _ = fi30
                                            }
                                        }
                                        _ = v23
                                    }
                                } else {
                                    v32 := GenS1{F0: 42, F1: (("a" + (("hello" + "hello") + (("hello" + ("a" + ("" + ("" + ("" + "\n\t"))))) + "\n\t"))) + "\n\t"), F2: 32767}
                                    for fi33 := 0; fi33 < 6; fi33++ {
                                        var v34 uint = 0
                                        _ = fi33
                                        _ = v34
                                    }
                                    var v35 int8 = 42
                                    _ = v32
                                    _ = v35
                                }
                                for fi36 := 0; fi36 < 5; fi36++ {
                                    for fi37 := 0; fi37 < 19; fi37++ {
                                        fmt.Println(v9, fi36)
                                        v8 = append(v8, fi36)
                                        if len(v8) >= 2 {
                                            v8 = v8[1:len(v8)]
                                        }
                                        _ = fi37
                                    }
                                    var v38 bool = (2 == 255)
                                    v9 -= (v9 * (1e10 / v9))
                                    _ = fi36
                                    _ = v38
                                }
                            }
                            _ = v21
                        }
                    }
                case 808:
                    v39 := GenS2{F0: 917}
                    _ = v39
                }
            }
            v40 := GenS3{F0: v7, F1: 18446744073709551615, F2: -9223372036854775808}
            _ = v40
        } else {
            v41 := GenS3{F0: (2147483647 < 4294967295), F1: 1, F2: 893}
            _ = v41
        }
        v42 := GenS1{F0: 171, F1: ("" + "hello"), F2: -32768}
        for fi43 := 0; fi43 < 5; fi43++ {
            v9 = v9
            _ = fi43
        }
        _ = v42
    }
    _ = v7
    _ = v8
    _ = v9
    return 2
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]uint32)
    v3 := make(map[int]uint32)
    var v4 int16 = 276
    var v5 uint32 = 513
    var v6 uint32 = (v5 + v5)
    var v44 int8 = 0
    fmt.Println(v4)
    fmt.Println(v5, v6)
    v45 := v1[42]
    if false {
        v45 = ((-9223372036854775808 > -9223372036854775808) && (1.5 <= 1e10))
        var v46 int64 = 42
        for fi47 := 0; fi47 < 14; fi47++ {
            if v45 {
                var v48 bool = (!false)
                if (false && true) {
                    if true {
                        v49 := GenS4{F0: v6, F1: v5}
                        _ = v49
                    }
                    var v50 int16 = v4
                    v45 = v48
                    _ = v50
                } else {
                    for fi51 := 0; fi51 < 6; fi51++ {
                        var v52 int = (fi51 / (((fi47 % (((fi51 | fi47)) | 1))) | 1))
                        _ = fi51
                        _ = v52
                    }
                    v4 += v4
                }
                _ = v48
            }
            _ = fi47
        }
        _ = v46
    } else {
        fmt.Println(v45, v6)
        var v53 int64 = 1
        fmt.Println(v45)
        _ = v53
    }
    var v54 uint8 = 197
    var v55 float64 = 1e10
    _ = len(v3)
    v56 := GenS1{F0: 2, F1: ("\n\t" + ("hello" + "a")), F2: (v4 & v4)}
    for fi57 := 0; fi57 < 19; fi57++ {
        switch ("a" + "a") {
        case "\n\t":
            v58 := []int{}
            _ = v58
        default:
            fmt.Println(v44, fi57, v45)
        }
        v59 := GenS4{F0: v6, F1: (v6 / ((844) | 1))}
        switch fi57 {
        case -334:
            v60 := GenS5{F0: 1}
            _ = v60
        case 1:
            _ = len(v3)
        }
        _ = fi57
        _ = v59
    }
    v5 = 2
    if (65535 <= 0) {
        var v61 uint8 = v54
        _ = v61
    }
    v45 = (134 != 42)
    v62 := []int{}
    if len(v62) >= 2 {
        v62 = v62[1:len(v62)]
    }
    if v45 {
        v5 = 480
        var v63 uint64 = 42
        var v64 float64 = 1.5
        _ = v63
        _ = v64
    } else {
        var v65 float32 = float32((0.0 * v55))
        fmt.Println(v54, v65, v6)
        _ = v65
    }
    var v66 uint64 = 2
    if true {
        var v67 uint = 4294967295
        _ = v67
    } else {
        if (false && (true && (v45 || ((v5 | 382) <= v5)))) {
            var v68 string = "hello"
            var v69 float64 = (v55 + v55)
            if (1 > -1) {
                for fi70 := 0; fi70 < 3; fi70++ {
                    if v45 {
                        var v71 string = v68
                        fmt.Println(v69, v5)
                        v72 := GenS6{F0: (v71 + (v71 + "")), F1: v45}
                        _ = v71
                        _ = v72
                    } else {
                        v1[42] = (v45 || (v45 && v45))
                        v66 -= (^(v66))
                    }
                    v44 -= (v44 % (((v44 << ((0) & 63))) | 1))
                    _ = fi70
                }
            } else {
                var v73 string = "a"
                _ = v73
            }
            _ = v68
            _ = v69
        } else {
            var v74 uint16 = 255
            v55 = (v55 * v55)
            _ = v74
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v62:", len(v62), v62)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v5
    _ = v54
    _ = v55
    _ = v56
    _ = v6
    _ = v62
    _ = v66
}
