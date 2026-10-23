package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int; F1 uint64 }
type GenS5 struct { F0 uint; F1 int32 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int16; F1 uint64; F2 int32 }
func genFn0(p0 int64, p1 bool, p2 int8) string {
    v7 := []int{4294967295, 127, 9223372036854775807, 109}
    if (42 == 1) {
        v8 := GenS1{F0: 42, F1: (1 != 2147483647)}
        var v9 uint16 = 497
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 1; fi10++ {
        var v11 int = (fi10 << ((771) & 63))
        _ = fi10
        _ = v11
    }
    v7 = append(v7, 118)
    if len(v7) > 0 {
        v7[0] = -1
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return "foo bar"
}
func genFn1() int64 {
    v12 := []int{}
    var v13 int8 = -35
    var v14 float32 = 0.1
    var v15 float64 = float64((((v14 - v14) * float32(0.1)) + v14))
    v16 := GenS1{F0: 2, F1: (847 != 127)}
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    return 1
}

func main() {
    v1 := make(map[int]int)
    v2 := []int{}
    var v3 int8 = -128
    v4 := make(map[int]bool)
    var v5 int64 = 930
    var v6 float32 = (1e10 / (0.0 - (1.5 + 0.0)))
    fmt.Println(v3, v5, v6)
    switch ("a" + "hello") {
    case "hello":
        v17 := GenS1{F0: 214, F1: false}
        _ = v17
    case "":
        if (65535 < 255) {
            v18 := GenS2{F0: -600, F1: false}
            var v19 string = "foo bar"
            _ = v18
            _ = v19
        } else {
            _ = len(v1)
            var v20 int = -9223372036854775808
            _ = v20
        }
    case "a":
        v5 = (^(v5))
    default:
        if (false || true) {
            var v21 uint = 0
            switch int(v5) {
            case 0:
                v5 += v5
            case 825:
                v6 -= v6
            }
            for fi22 := 0; fi22 < 13; fi22++ {
                switch -2147483648 {
                case 695:
                    v21 = 833
                }
                _ = fi22
            }
            _ = v21
        }
    }
    if (!false) {
        for fi23 := 0; fi23 < 17; fi23++ {
            v24 := GenS3{F0: ((fi23 & (fi23 & (fi23 + -355))) != (fi23 & fi23))}
            for fi25 := 0; fi25 < 7; fi25++ {
                v26 := GenS1{F0: 2, F1: (("\n\t" + "hello") <= "\n\t")}
                v27 := GenS1{F0: 2, F1: (-15 >= (v3 % ((-1) | 1)))}
                fmt.Println(fi23)
                _ = fi25
                _ = v26
                _ = v27
            }
            switch "\n\t" {
            case "hello":
                for fi28 := 0; fi28 < 19; fi28++ {
                    v29 := GenS1{F0: 2, F1: (0.0 <= ((1e10 - ((3.14159 * 0.0) / 1.0)) / 3.14159))}
                    _ = fi28
                    _ = v29
                }
            case "":
                for fi30 := 0; fi30 < 15; fi30++ {
                    if (!(fi23 >= (fi30 % (((-(fi23))) | 1)))) {
                        var v31 int = (fi23 / ((fi30) | 1))
                        _ = v31
                    } else {
                        var v32 uint32 = 773
                        if (65535 >= 490) {
                            var v33 int8 = 42
                            var v34 uint16 = 32767
                            _ = v33
                            _ = v34
                        }
                        v32 += v32
                        _ = v32
                    }
                    var v35 uint = 0
                    fmt.Println(fi23, fi30, fi23)
                    var v36 uint64 = 721
                    v3 -= (-(v3))
                    v5 = 9223372036854775807
                    v37 := GenS4{F0: (fi30 & fi23), F1: (^(v36))}
                    _ = fi30
                    _ = v35
                    _ = v36
                    _ = v37
                }
            case "a":
                if false {
                    v5 -= 42
                    if false {
                        var v38 float32 = ((((v6 / 0.1) - v6) - (1.5 / 1.0)) - -2.25)
                        _ = v38
                    } else {
                        switch (fi23 / (((fi23 & (^(fi23)))) | 1)) {
                        case 42:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        case -1:
                            switch 0 {
                            case -9223372036854775808:
                                v39 := GenS1{F0: 65535, F1: false}
                                _ = v39
                            case -1:
                                var v40 string = "foo bar"
                                _ = v40
                            }
                        case 824:
                            v41 := v4[-656]
                            _ = v41
                        default:
                            var v42 uint8 = 0
                            _ = v42
                        }
                        v43 := GenS5{F0: 255, F1: -2147483648}
                        _ = v43
                    }
                    switch (fi23 ^ fi23) {
                    case 2147483647:
                        fmt.Println(fi23, v5, v3)
                    case 42:
                        v1[fi23] = 128
                    }
                }
            default:
                var v44 int8 = v3
                _ = v44
            }
            _ = fi23
            _ = v24
        }
        var v45 uint32 = 4294967295
        _ = v45
    }
    v4[-1] = true
    for fi46 := 0; fi46 < 15; fi46++ {
        v3 = v3
        _ = fi46
    }
    for fi47 := 0; fi47 < 16; fi47++ {
        var v48 float64 = ((-2.25 - 1.5) * (-2.25 * 3.14159))
        fmt.Println(v3, v48, v5)
        _ = fi47
        _ = v48
    }
    switch 4294967295 {
    case 9223372036854775807:
        v3 -= (v3 ^ -128)
    case 913:
        v3 -= (v3 | (v3 | v3))
    case 128:
        if false {
            delete(v4, 32767)
            switch 127 {
            case 127:
                fmt.Println(v5)
            case -9223372036854775808:
                if true {
                    if (1 >= v3) {
                        if (2147483647 == 32768) {
                            v3 = (v3 & v3)
                            v5 = v5
                        }
                        v49 := GenS5{F0: 873, F1: -2147483648}
                        _ = v49
                    }
                } else {
                    v3 += v3
                    var v50 uint8 = 255
                    _ = len(v4)
                    _ = v50
                }
            case -523:
                if true {
                    for fi51 := 0; fi51 < 4; fi51++ {
                        var v52 uint32 = 32
                        var v53 string = "foo bar"
                        _ = fi51
                        _ = v52
                        _ = v53
                    }
                }
            }
        }
    default:
        for fi54 := 0; fi54 < 5; fi54++ {
            for fi55 := 0; fi55 < 18; fi55++ {
                v3 = v3
                switch "" {
                case "a":
                    fmt.Println(v6, v5, fi55)
                case "\n\t":
                    if false {
                        var v56 uint32 = 65535
                        fmt.Println(v3)
                        v57 := GenS6{F0: false}
                        if (!true) {
                            v56 = 2147483647
                        } else {
                            var v58 uint = 4294967295
                            v59 := GenS7{F0: -223}
                            _ = v58
                            _ = v59
                        }
                        switch (fi54 % (((fi55 << ((0) & 63))) | 1)) {
                        case 0:
                            v5 += int64(v56)
                        }
                        v3 = (v3 / (((v3 << ((18446744073709551615) & 63))) | 1))
                        var v60 uint64 = 0
                        v61 := GenS8{F0: 42, F1: 4294967295, F2: 917}
                        _ = v56
                        _ = v57
                        _ = v60
                        _ = v61
                    }
                case "hello":
                    v3 += 0
                }
                var v62 bool = (0 >= 18446744073709551615)
                _ = fi55
                _ = v62
            }
            _ = fi54
        }
    }
    for fi63 := 0; fi63 < 16; fi63++ {
        v3 = 0
        _ = fi63
    }
    var v64 uint32 = 2147483647
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v64
}
