package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 uint32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int32; F1 int64; F2 uint }
type GenS4 struct { F0 int8; F1 int16 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 string; F1 uint32; F2 int8 }

func main() {
    var v1 uint = 630
    var v2 uint64 = 42
    var v3 int32 = -2147483648
    var v4 int8 = 1
    var v5 float64 = 3.14159
    v6 := []int{}
    var v7 uint = v1
    v7 -= (v1 + (^(v7)))
    if (-735 >= 1) {
        var v8 int32 = 2147483647
        _ = v8
    }
    v3 = (v3 / ((0) | 1))
    if true {
        var v9 uint32 = 98
        _ = v9
    } else {
        var v10 uint32 = 0
        var v11 int64 = -499
        v12 := GenS1{F0: 32767, F1: ("hello" + ("\n\t" + "a")), F2: (v10 ^ v10)}
        _ = v10
        _ = v11
        _ = v12
    }
    v4 = (v4 - (v4 % ((v4) | 1)))
    for fi13 := 0; fi13 < 7; fi13++ {
        for fi14 := 0; fi14 < 9; fi14++ {
            v3 -= (v3 & (-(v3)))
            v15 := GenS1{F0: 2147483647, F1: "\n\t", F2: 4294967295}
            if (0 != 42) {
                var v16 int16 = 1
                _ = v16
            } else {
                v6 = append(v6, int(v2))
                var v17 string = "a"
                if (!(true && true)) {
                    var v18 int64 = 0
                    _ = len(v6)
                    if (1000 > 1000) {
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        switch (fi13 + 0) {
                        case 9223372036854775807:
                            var v19 int = fi13
                            _ = v19
                        }
                    } else {
                        v2 = v2
                    }
                    _ = v18
                } else {
                    switch -120 {
                    case 484:
                        switch fi13 {
                        case 1:
                            var v20 uint = (v7 & (v7 / (((v1 & 956)) | 1)))
                            _ = v20
                        }
                    case -9223372036854775808:
                        fmt.Println(v3, v3)
                    case 1:
                        switch "a" {
                        case "a":
                            v5 -= 1.5
                        default:
                            var v21 int16 = 32767
                            _ = v21
                        }
                    }
                    if true {
                        v5 = 0.0
                    } else {
                        if false {
                            if true {
                                var v22 float64 = ((float64((1.5 * 1.5)) - (0.0 - float64((1.5 * 1.5)))) * (1.5 * ((v5 * v5) / -2.25)))
                                var v23 float32 = 1.5
                                _ = v22
                                _ = v23
                            } else {
                                v24 := GenS2{F0: -9223372036854775808}
                                switch "a" {
                                case "\n\t":
                                    v5 = float64(float32(v5))
                                    if ((65535 == 213) || false) {
                                        v7 = 2
                                    }
                                    fmt.Println(v3)
                                }
                                v6 = append(v6, fi14)
                                v25 := GenS3{F0: int32(v7), F1: -9223372036854775808, F2: v1}
                                v26 := GenS2{F0: -753}
                                _ = v24
                                _ = v25
                                _ = v26
                            }
                            var v27 uint8 = 2
                            _ = v27
                        } else {
                            v7 -= 255
                        }
                    }
                    var v28 string = "\n\t"
                    _ = v28
                }
                if (499 < 42) {
                    v1 -= 2
                }
                _ = v17
            }
            _ = fi14
            _ = v15
        }
        v29 := GenS3{F0: (v3 ^ v3), F1: -793, F2: v1}
        _ = fi13
        _ = v29
    }
    for fi30 := 0; fi30 < 3; fi30++ {
        for fi31 := 0; fi31 < 13; fi31++ {
            v32 := GenS2{F0: -1}
            if false {
                v4 = (v4 | (v4 ^ v4))
                v33 := GenS4{F0: -23, F1: -119}
                _ = len(v6)
                _ = v33
            } else {
                fmt.Println(fi31, v1, v1)
                v34 := GenS2{F0: 0}
                _ = len(v6)
                _ = v34
            }
            _ = fi31
            _ = v32
        }
        _ = fi30
    }
    switch 40 {
    case 450:
        if false {
            if (!(false || (!true))) {
                var v35 int32 = -1
                _ = v35
            } else {
                if (!(((float64(1.5) + 1e10) + (v5 * ((3.14159 + v5) + (3.14159 - v5)))) != v5)) {
                    var v36 int = 9223372036854775807
                    _ = v36
                }
                switch int(v4) {
                case 42:
                    fmt.Println(v4, v1, v1)
                case 234:
                    switch int(v4) {
                    case 127:
                        fmt.Println(v1, v7, v1)
                    case -841:
                        if (false && (128 >= 2)) {
                            if true {
                                v37 := GenS5{F0: 59}
                                v38 := GenS4{F0: (v4 >> (((v7 & (v1 << (((v1 | 549)) & 63)))) & 63)), F1: 1}
                                _ = v37
                                _ = v38
                            } else {
                                v39 := GenS4{F0: (v4 & v4), F1: -32768}
                                _ = v39
                            }
                            var v40 int16 = 32767
                            var v41 bool = (42 <= 9223372036854775807)
                            _ = v40
                            _ = v41
                        }
                    case 4294967295:
                        v42 := GenS2{F0: 749}
                        _ = v42
                    }
                }
            }
        }
    case -2147483648:
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    case 4294967295:
        if false {
            v3 = v3
        }
    }
    for fi43 := 0; fi43 < 20; fi43++ {
        for fi44 := 0; fi44 < 8; fi44++ {
            var v45 bool = false
            fmt.Println(v45, fi44, fi44)
            var v46 uint64 = 4294967295
            if (((v5 + ((0.1 / v5) + v5)) + v5) < v5) {
                v45 = (127 < fi44)
                if ((!false) && (42 <= 1)) {
                    var v47 float32 = ((1.5 * -2.25) - (1e10 - 1.5))
                    v7 -= v7
                    _ = v47
                } else {
                    var v48 string = "foo bar"
                    var v49 float64 = (1.5 / (0.1 - 1.5))
                    _ = v48
                    _ = v49
                }
            } else {
                if true {
                    var v50 string = ("hello" + "")
                    _ = v50
                }
                switch "hello" {
                case "hello":
                    v45 = (fi43 >= fi43)
                case "a":
                    if v45 {
                        v3 += v3
                    }
                case "\n\t":
                    _ = len(v6)
                }
            }
            if len(v6) > 0 {
                v6[0] = (fi44 ^ fi44)
            }
            _ = fi44
            _ = v45
            _ = v46
        }
        v5 -= 3.14159
        _ = fi43
    }
    v6 = append(v6, 4294967295)
    v51 := GenS6{F0: "", F1: 0, F2: v4}
    fmt.Println(v2, v1)
    for fi52 := 0; fi52 < 14; fi52++ {
        fmt.Println(v5, v4, v4)
        _ = fi52
    }
    v5 -= 1e10
    v5 -= v5
    for fi53 := 0; fi53 < 5; fi53++ {
        var v54 float32 = (float32(float64((0.0 / (((float32(v5) * 3.14159) + (((3.14159 / (0.1 - 3.14159)) - 3.14159) - 0.1)) - (1.5 - 3.14159))))) + 1e10)
        for fi55 := 0; fi55 < 17; fi55++ {
            var v56 float64 = (-2.25 + (v5 / v5))
            _ = fi55
            _ = v56
        }
        _ = fi53
        _ = v54
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v6
    _ = v7
}
