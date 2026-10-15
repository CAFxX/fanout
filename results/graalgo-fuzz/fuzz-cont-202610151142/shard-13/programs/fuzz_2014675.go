package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 string; F1 uint }

func main() {
    var v1 int = 9223372036854775807
    v2 := []string{"foo bar", ("" + ("hello" + "foo bar")), "hello", "\n\t"}
    var v3 string = "foo bar"
    v4 := make(map[string]string)
    var v5 string = v3
    var v6 uint16 = 449
    v7 := GenS1{F0: 231, F1: 0, F2: -57}
    for fi8 := 0; fi8 < 14; fi8++ {
        var v9 float64 = (float64(1e10) * (3.14159 / (0.1 / (1.5 - (1e10 + 1e10)))))
        _ = fi8
        _ = v9
    }
    var v10 uint32 = 2
    v6 -= v6
    if (43 == 65535) {
        v11 := GenS2{F0: 0}
        var v12 int64 = -9223372036854775808
        v13 := GenS1{F0: 255, F1: 1000, F2: -1}
        _ = v11
        _ = v12
        _ = v13
    } else {
        if true {
            switch (v1 / (((v1 >> ((844) & 63))) | 1)) {
            case 9223372036854775807:
                for fi14 := 0; fi14 < 13; fi14++ {
                    v3 = "\n\t"
                    v10 += v10
                    if (128 > 255) {
                        _ = len(v2)
                        v5 = v5
                        v15 := GenS3{F0: (v6 / ((v6) | 1))}
                        _ = v15
                    }
                    _ = fi14
                }
            default:
                if false {
                    var v16 int = (-(v1))
                    v1 -= (v1 % ((0) | 1))
                    _ = v16
                }
            }
        } else {
            for fi17 := 0; fi17 < 19; fi17++ {
                var v18 int8 = 0
                var v19 int32 = 415
                for fi20 := 0; fi20 < 4; fi20++ {
                    var v21 uint32 = 2147483648
                    v22 := GenS3{F0: v6}
                    v6 = v6
                    _ = fi20
                    _ = v21
                    _ = v22
                }
                _ = fi17
                _ = v18
                _ = v19
            }
            var v23 uint32 = (v10 ^ v10)
            _ = v23
        }
        v2 = append(v2, v5)
    }
    switch v5 {
    case "a":
        if (!false) {
            fmt.Println(v6, v1)
            v1 = v1
            fmt.Println(v5, v1, v3)
        }
    case "foo bar":
        if (uint8(v10) == 255) {
            v24 := GenS3{F0: v6}
            v10 = v10
            if (127 > -128) {
                if (float32(0.0) >= 1e10) {
                    delete(v4, v3)
                    for fi25 := 0; fi25 < 20; fi25++ {
                        var v26 int32 = 498
                        _ = fi25
                        _ = v26
                    }
                    v3 = "a"
                }
                if false {
                    var v27 uint64 = 535
                    fmt.Println(v10, v3)
                    _ = v27
                }
                if len(v2) > 0 {
                    v2[0] = (v5 + "")
                }
            } else {
                if (906 == 18446744073709551615) {
                    _ = len(v4)
                    switch -946 {
                    case 436:
                        var v28 int16 = 1000
                        _ = v28
                    case -914:
                        switch v3 {
                        case "a":
                            v10 += (v10 + 42)
                        case "":
                            var v29 int16 = -1
                            _ = v29
                        }
                    case 467:
                        v30 := GenS4{F0: 143}
                        _ = v30
                    default:
                        v1 += v1
                    }
                    for fi31 := 0; fi31 < 14; fi31++ {
                        var v32 uint = 628
                        var v33 uint = (v32 * v32)
                        var v34 uint16 = (v6 * (v6 - v6))
                        _ = fi31
                        _ = v32
                        _ = v33
                        _ = v34
                    }
                } else {
                    var v35 float64 = (3.14159 / 3.14159)
                    var v36 uint16 = 42
                    var v37 int = v1
                    _ = v35
                    _ = v36
                    _ = v37
                }
                v10 = v10
            }
            _ = v24
        } else {
            var v38 float64 = 1e10
            var v39 int = 0
            var v40 uint16 = v6
            _ = v38
            _ = v39
            _ = v40
        }
    case "hello":
        fmt.Println(v10)
    default:
        v5 = v3
    }
    if len(v2) > 0 {
        v2[0] = v3
    }
    var v41 bool = (((-2.25 * 1e10) * 0.0) >= (0.1 / float32(3.14159)))
    v41 = (0.1 < float64(0.0))
    for fi42 := 0; fi42 < 15; fi42++ {
        v10 -= (v10 + 2147483647)
        if (1.5 <= 0.0) {
            if (v6 >= v6) {
                switch fi42 {
                case -9223372036854775808:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                default:
                    fmt.Println(fi42, v1, v1)
                }
                var v43 uint16 = (v6 & v6)
                _ = v43
            }
            var v44 bool = v41
            v6 += v6
            _ = v44
        } else {
            v45 := GenS5{F0: (((v3 + "a") + v3) + v5), F1: 65535}
            v46 := GenS1{F0: 2, F1: -1, F2: 1}
            switch (v3 + "foo bar") {
            case "":
                v6 = v6
            }
            _ = v45
            _ = v46
        }
        v2 = append(v2, v3)
        _ = fi42
    }
    v4[v3] = v3
    if false {
        v10 += (v10 - (v10 / ((v10) | 1)))
        _ = len(v4)
        fmt.Println(v3)
    } else {
        switch (v1 & -703) {
        case -35:
            if (!v41) {
                v5 = v3
            } else {
                v6 = v6
                var v47 uint16 = (v6 * (v6 & 260))
                v41 = false
                _ = v47
            }
        }
        var v48 int64 = -9223372036854775808
        _ = v48
    }
    fmt.Println(v41)
    fmt.Println(v6, v1)
    var v49 uint32 = (v10 & (v10 + v10))
    var v50 float32 = (((1e10 / (0.0 - (1.5 - (float32(1e10) - 0.0)))) - ((0.0 + (0.0 / 1.5)) / 1e10)) / (0.1 + ((1.5 - 0.1) / float32(3.14159))))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v49
    _ = v5
    _ = v50
    _ = v6
    _ = v7
}
