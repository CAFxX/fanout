package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint32; F1 uint16; F2 uint64 }

func main() {
    var v1 string = "foo bar"
    v2 := make(map[uint32]int)
    var v3 int8 = -128
    var v4 uint16 = 0
    if (32767 > -1) {
        v2[42] = 4294967295
        v5 := []int{4294967295, 2147483647, 842, 32768}
        _ = v5
    } else {
        v3 = v3
        var v6 string = (((v1 + ((v1 + v1) + v1)) + v1) + (v1 + ("" + ("foo bar" + (v1 + v1)))))
        _ = v6
    }
    fmt.Println(v4, v4)
    for fi7 := 0; fi7 < 14; fi7++ {
        var v8 uint8 = 255
        v4 = (^(v4))
        _ = fi7
        _ = v8
    }
    v1 = (v1 + v1)
    fmt.Println(v4, v3)
    var v9 uint = 238
    var v10 bool = (2 >= 42)
    if true {
        var v11 int64 = -804
        v12 := []int{}
        v3 -= v3
        _ = v11
        _ = v12
    } else {
        var v13 uint32 = 0
        v3 += v3
        _ = v13
    }
    v1 = v1
    for fi14 := 0; fi14 < 8; fi14++ {
        v15 := []int{-1, 42}
        var v16 int = (fi14 ^ (fi14 * 9223372036854775807))
        _ = fi14
        _ = v15
        _ = v16
    }
    v3 = 98
    v17 := GenS1{F0: 0}
    v1 = v1
    if (-32768 >= -1) {
        if (false || (-1 > -670)) {
            for fi18 := 0; fi18 < 4; fi18++ {
                if (42 != -2147483648) {
                    fmt.Println(v9)
                    v19 := []int{127, fi18, (-(fi18))}
                    _ = v19
                } else {
                    if (v10 && (65535 > 2147483647)) {
                        v20 := GenS2{F0: -1, F1: 0}
                        switch v1 {
                        case "hello":
                            fmt.Println(v10, v3)
                        case "":
                            var v21 string = (v1 + ("\n\t" + ("hello" + ((v1 + v1) + "a"))))
                            _ = v21
                        default:
                            switch (fi18 + 1) {
                            case 32768:
                                v22 := []int{-645, fi18, (fi18 & (fi18 + (fi18 << (((v9 ^ 18446744073709551615)) & 63)))), fi18}
                                _ = v22
                            case 4294967295:
                                v23 := []int{}
                                _ = v23
                            default:
                                fmt.Println(v3)
                            }
                        }
                        _ = v20
                    } else {
                        var v24 int64 = 200
                        var v25 bool = v10
                        var v26 uint = 0
                        _ = v24
                        _ = v25
                        _ = v26
                    }
                    v27 := v2[2147483647]
                    _ = v27
                }
                var v28 float32 = ((1e10 + float32(((((1.5 - 1e10) / ((((3.14159 + 3.14159) / (float64(float32(0.0)) - 1e10)) / 3.14159) * 0.1)) * (-2.25 - 3.14159)) - (float64((-2.25 - -2.25)) + (0.1 / float64((((((((float32(-2.25) - (0.0 + (((1.5 / 1.5) / -2.25) + (0.1 - (1e10 + 3.14159))))) - 3.14159) + float32((-2.25 / 1e10))) - 1.5) * 1.5) - 1e10) / 1e10))))))) / 0.1)
                _ = fi18
                _ = v28
            }
            if v10 {
                if (0 != 603) {
                    var v29 string = ("\n\t" + v1)
                    _ = v29
                } else {
                    var v30 float64 = float64(1.5)
                    _ = v30
                }
            } else {
                var v31 int16 = -1
                v9 = v9
                _ = v31
            }
        }
        var v32 int32 = -2147483648
        v3 = v3
        _ = v32
    } else {
        if (-1 > -32768) {
            v1 = v1
            switch 4294967295 {
            case -841:
                _ = len(v2)
            case -2147483648:
                v33 := GenS3{F0: 4294967295}
                _ = v33
            case -763:
                for fi34 := 0; fi34 < 18; fi34++ {
                    v35 := v2[2]
                    for fi36 := 0; fi36 < 18; fi36++ {
                        var v37 int16 = 42
                        _ = fi36
                        _ = v37
                    }
                    _ = fi34
                    _ = v35
                }
            default:
                for fi38 := 0; fi38 < 4; fi38++ {
                    for fi39 := 0; fi39 < 7; fi39++ {
                        var v40 bool = (3.14159 > -2.25)
                        v41 := GenS4{F0: 922, F1: v4, F2: 18446744073709551615}
                        _ = fi39
                        _ = v40
                        _ = v41
                    }
                    for fi42 := 0; fi42 < 16; fi42++ {
                        v4 = v4
                        var v43 uint32 = 65535
                        if v10 {
                            v4 += (v4 & v4)
                            var v44 int16 = 1000
                            if (v1 < (v1 + v1)) {
                                v45 := GenS4{F0: v43, F1: v4, F2: 0}
                                v4 -= v4
                                _ = v45
                            } else {
                                v43 -= (v43 / ((2147483647) | 1))
                            }
                            _ = v44
                        }
                        _ = fi42
                        _ = v43
                    }
                    _ = fi38
                }
            }
        }
        fmt.Println(v3, v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
