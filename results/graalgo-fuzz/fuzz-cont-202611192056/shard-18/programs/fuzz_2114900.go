package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 uint }
type GenS2 struct { F0 int32; F1 uint64 }
type GenS3 struct { F0 uint16; F1 uint8 }
type GenS4 struct { F0 uint32; F1 string; F2 int64 }
func genFn0(p0 bool, p1 int64, p2 int) uint16 {
    v5 := []int{p2, p2}
    var v6 uint = 65535
    fmt.Println(p1, p2, v6)
    if p0 {
        for fi7 := 0; fi7 < 12; fi7++ {
            if false {
                p2 += fi7
            } else {
                var v8 float64 = -2.25
                _ = v8
            }
            var v9 uint64 = 490
            _ = fi7
            _ = v9
        }
        var v10 uint16 = 1
        _ = v10
    }
    if ("hello" == (("\n\t" + ("a" + "")) + ("" + ((("" + "a") + ("foo bar" + (("foo bar" + "hello") + "foo bar"))) + ("a" + ""))))) {
        fmt.Println(p2, p2, p1)
    } else {
        p1 = -236
        var v11 uint16 = 0
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 32767
}

func main() {
    v1 := []int{128, 32768, -9223372036854775808, -9223372036854775808}
    var v2 string = "hello"
    var v3 int32 = 42
    var v4 int64 = 9223372036854775807
    fmt.Println(v4, v2, v2)
    var v12 uint16 = 32767
    fmt.Println(v2)
    if true {
        var v13 int = -921
        _ = v13
    } else {
        v3 += 1
        for fi14 := 0; fi14 < 20; fi14++ {
            var v15 int32 = 825
            _ = fi14
            _ = v15
        }
    }
    v16 := GenS1{F0: 800, F1: 239, F2: 4294967295}
    var v17 string = v2
    var v18 int = -560
    v4 = (-(v4))
    for fi19 := 0; fi19 < 5; fi19++ {
        var v20 int32 = (v3 | (v3 * 534))
        v4 -= (v4 / ((v4) | 1))
        _ = fi19
        _ = v20
    }
    fmt.Println(v17, v3, v12)
    switch (^(v18)) {
    case 9223372036854775807:
        switch (("\n\t" + (v2 + v17)) + v17) {
        case "\n\t":
            if true {
                var v21 uint = 2
                v2 = v2
                _ = v21
            } else {
                if true {
                    for fi22 := 0; fi22 < 2; fi22++ {
                        v4 = v4
                        if (1 <= v18) {
                            if true {
                                v23 := GenS1{F0: fi22, F1: 0, F2: 1}
                                _ = v23
                            } else {
                                if false {
                                    v18 += (fi22 + 32768)
                                } else {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    for fi24 := 0; fi24 < 14; fi24++ {
                                        v4 -= (v4 >> ((501) & 63))
                                        if (!true) {
                                            _ = len(v1)
                                        } else {
                                            if true {
                                                var v25 uint = 0
                                                v4 += (v4 * (v4 * v4))
                                                _ = v25
                                            } else {
                                                var v26 float32 = ((1.5 * 1e10) * 0.0)
                                                if (!false) {
                                                    var v27 int8 = 127
                                                    _ = len(v1)
                                                    _ = v27
                                                }
                                                _ = v26
                                            }
                                            _ = len(v1)
                                            v3 += v3
                                        }
                                        _ = fi24
                                    }
                                }
                            }
                        } else {
                            switch v17 {
                            case "\n\t":
                                v3 += (v3 / (((v3 | (v3 | (v3 ^ (v3 & (v3 >> ((2) & 63))))))) | 1))
                            case "hello":
                                switch v18 {
                                case 1:
                                    for fi28 := 0; fi28 < 10; fi28++ {
                                        fmt.Println(v17, v17)
                                        v29 := GenS2{F0: (v3 | v3), F1: 18446744073709551615}
                                        _ = fi28
                                        _ = v29
                                    }
                                default:
                                    v18 = int(v3)
                                }
                            default:
                                v18 -= 127
                            }
                            switch fi22 {
                            case 4294967295:
                                var v30 bool = ((-30 < 42) || (32767 != -32768))
                                _ = v30
                            }
                        }
                        v31 := GenS1{F0: (fi22 ^ fi22), F1: 1, F2: 65535}
                        _ = fi22
                        _ = v31
                    }
                    if false {
                        fmt.Println(v3)
                        v4 -= (v4 & -9223372036854775808)
                    } else {
                        if true {
                            var v32 uint = 65535
                            switch (v18 + (v18 % ((v18) | 1))) {
                            case 127:
                                var v33 int32 = v3
                                _ = v33
                            case 9223372036854775807:
                                var v34 uint64 = uint64(v12)
                                _ = v34
                            default:
                                var v35 uint8 = 108
                                _ = v35
                            }
                            var v36 int = 161
                            _ = v32
                            _ = v36
                        }
                        fmt.Println(v17, v2, v12)
                        v4 = (v4 ^ -975)
                    }
                } else {
                    if len(v1) > 0 {
                        v1[0] = 128
                    }
                    v37 := GenS2{F0: (v3 >> ((1) & 63)), F1: 42}
                    v12 += (^(v12))
                    _ = v37
                }
                v2 = (v17 + v2)
                fmt.Println(v4, v17)
            }
        case "foo bar":
            v1 = append(v1, 9223372036854775807)
        default:
            switch (v18 ^ (-(v18))) {
            case 127:
                if true {
                    for fi38 := 0; fi38 < 5; fi38++ {
                        if len(v1) > 0 {
                            v1[0] = (v18 ^ v18)
                        }
                        _ = fi38
                    }
                    _ = len(v1)
                } else {
                    v18 -= 128
                }
            }
        }
    case 32767:
        for fi39 := 0; fi39 < 17; fi39++ {
            v2 = "\n\t"
            _ = fi39
        }
    case 1:
        fmt.Println(v3)
    default:
        if (!false) {
            var v40 int32 = 42
            var v41 int8 = 127
            var v42 int16 = 42
            _ = v40
            _ = v41
            _ = v42
        } else {
            var v43 uint16 = (v12 + v12)
            switch v18 {
            case -9223372036854775808:
                v12 += v43
            }
            for fi44 := 0; fi44 < 12; fi44++ {
                var v45 float64 = float64(3.14159)
                _ = fi44
                _ = v45
            }
            _ = v43
        }
    }
    if len(v1) > 0 {
        v1[0] = 32767
    }
    switch (v18 % ((v18) | 1)) {
    case -2147483648:
        v46 := GenS3{F0: v12, F1: 42}
        _ = v46
    case 42:
        var v47 float64 = 1e10
        _ = v47
    default:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    switch v17 {
    case "\n\t":
        v48 := GenS4{F0: 4294967295, F1: v2, F2: (v4 | 0)}
        _ = v48
    default:
        v49 := GenS3{F0: (v12 / ((genFn0(true, (v4 % ((v4) | 1)), v18)) | 1)), F1: 42}
        _ = v49
    }
    if (float32(0.0) < 0.1) {
        var v50 uint64 = 0
        _ = v50
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
