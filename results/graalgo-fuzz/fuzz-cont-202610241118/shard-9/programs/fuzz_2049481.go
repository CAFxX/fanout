package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 uint32 }
type GenS2 struct { F0 uint; F1 uint8; F2 int64 }
type GenS3 struct { F0 uint; F1 bool }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int }

func main() {
    var v1 uint16 = 65535
    var v2 string = ("\n\t" + "hello")
    v3 := []string{"hello", "foo bar"}
    var v4 uint16 = (v1 >> ((464) & 63))
    v5 := []string{}
    var v6 int16 = -32768
    var v7 int64 = -66
    var v8 uint64 = 399
    if (!(!(-2147483648 > 1))) {
        v2 = "\n\t"
        var v9 uint8 = 1
        _ = v9
    } else {
        v10 := GenS1{F0: 140, F1: uint64(v1), F2: 4294967295}
        var v11 string = ("a" + v2)
        if (!(!false)) {
            for fi12 := 0; fi12 < 9; fi12++ {
                v7 = v7
                fmt.Println(v1, v11, v8)
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = fi12
            }
            if (!true) {
                v1 += (v1 - 255)
            } else {
                var v13 float64 = (0.1 * 0.1)
                v14 := GenS1{F0: 2147483647, F1: v8, F2: 65535}
                for fi15 := 0; fi15 < 12; fi15++ {
                    v16 := GenS2{F0: 4294967295, F1: 2, F2: (v7 - v7)}
                    switch 32767 {
                    case 0:
                        for fi17 := 0; fi17 < 8; fi17++ {
                            var v18 int64 = v7
                            var v19 int16 = (v6 & v6)
                            v20 := GenS2{F0: 4294967295, F1: 1, F2: (v18 - v18)}
                            _ = fi17
                            _ = v18
                            _ = v19
                            _ = v20
                        }
                    case -9223372036854775808:
                        if ((float32(v13) >= ((0.0 / 1e10) - 0.0)) && (0 > 2147483647)) {
                            v1 -= 65535
                            for fi21 := 0; fi21 < 9; fi21++ {
                                if false {
                                    if true {
                                        v22 := GenS1{F0: fi21, F1: 18446744073709551615, F2: 2}
                                        switch v11 {
                                        case "foo bar":
                                            v8 += (v8 & 581)
                                        }
                                        _ = v22
                                    }
                                    v7 = 42
                                    if false {
                                        var v23 uint64 = (v8 | (^(v8)))
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        _ = v23
                                    }
                                    v11 = v2
                                    v24 := GenS2{F0: uint(v8), F1: 183, F2: (^(v7))}
                                    _ = v24
                                } else {
                                    var v25 uint32 = 0
                                    v4 += (v4 * (v1 - 32767))
                                    _ = v25
                                }
                                var v26 bool = (2 < (^(v1)))
                                var v27 int16 = 1000
                                if v26 {
                                    if v26 {
                                        var v28 int16 = (v27 - 32767)
                                        v26 = (2147483648 < 363)
                                        _ = v28
                                    }
                                }
                                var v29 float64 = 1e10
                                v7 = (v7 / (((v7 | (-(v7)))) | 1))
                                _ = fi21
                                _ = v26
                                _ = v27
                                _ = v29
                            }
                        }
                    case 32768:
                        for fi30 := 0; fi30 < 1; fi30++ {
                            fmt.Println(v11, v2)
                            _ = fi30
                        }
                    }
                    _ = fi15
                    _ = v16
                }
                _ = v13
                _ = v14
            }
        }
        _ = v10
        _ = v11
    }
    if true {
        v3 = append(v3, v2)
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    var v31 uint16 = 67
    var v32 uint = 4294967295
    for fi33 := 0; fi33 < 12; fi33++ {
        if ((v2 + v2) <= v2) {
            v1 = 65535
            var v34 string = ""
            v35 := GenS3{F0: (v32 ^ 936), F1: false}
            _ = v34
            _ = v35
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        for fi36 := 0; fi36 < 15; fi36++ {
            v37 := GenS2{F0: (v32 << ((v32) & 63)), F1: 128, F2: v7}
            v4 -= v4
            var v38 uint16 = v31
            _ = fi36
            _ = v37
            _ = v38
        }
        _ = fi33
    }
    for fi39 := 0; fi39 < 16; fi39++ {
        var v40 int32 = 1
        _ = fi39
        _ = v40
    }
    v5 = append(v5, (((v2 + (("hello" + v2) + v2)) + v2) + v2))
    fmt.Println(v4, v8)
    v41 := GenS1{F0: 1, F1: v8, F2: 2}
    for fi42 := 0; fi42 < 3; fi42++ {
        switch fi42 {
        case 42:
            if (v32 == v32) {
                v1 = 347
            }
        default:
            for fi43 := 0; fi43 < 6; fi43++ {
                var v44 uint32 = 65535
                _ = fi43
                _ = v44
            }
        }
        v45 := GenS3{F0: (v32 / (((^(v32))) | 1)), F1: (2 > 2147483648)}
        _ = fi42
        _ = v45
    }
    fmt.Println(v4, v31)
    var v46 uint64 = (v8 ^ 839)
    switch 1 {
    case 2147483647:
        switch 9223372036854775807 {
        case 32768:
            var v47 int64 = 42
            _ = v47
        case -1:
            var v48 int = 267
            _ = v48
        }
    case -9223372036854775808:
        v6 -= int16(v46)
    }
    if (!false) {
        for fi49 := 0; fi49 < 16; fi49++ {
            switch "foo bar" {
            case "":
                if true {
                    for fi50 := 0; fi50 < 3; fi50++ {
                        var v51 float32 = 1.5
                        _ = fi50
                        _ = v51
                    }
                    fmt.Println(v7, v1)
                }
            case "foo bar":
                fmt.Println(v31, v1, v46)
            case "hello":
                v6 += -917
            }
            var v52 float64 = 1.5
            var v53 uint = v32
            _ = fi49
            _ = v52
            _ = v53
        }
        v6 += (v6 / (((^(v6))) | 1))
    } else {
        switch v2 {
        case "hello":
            fmt.Println(v2, v8)
        case "foo bar":
            if false {
                _ = len(v5)
                if ((3.14159 / (((float32(float64(float32(3.14159))) - 0.1) - (-2.25 - 1e10)) * 0.1)) < 3.14159) {
                    var v54 uint32 = 65535
                    for fi55 := 0; fi55 < 15; fi55++ {
                        var v56 int8 = 100
                        _ = fi55
                        _ = v56
                    }
                    var v57 int = 4294967295
                    _ = v54
                    _ = v57
                }
                if len(v3) > 0 {
                    v3[0] = "a"
                }
            } else {
                _ = len(v3)
                v4 = (^(v1))
                var v58 float64 = 0.1
                _ = v58
            }
        case "\n\t":
            v59 := GenS4{F0: 42}
            _ = v59
        }
    }
    if (!true) {
        if true {
            v60 := GenS5{F0: -2147483648}
            _ = v60
        }
        var v61 int16 = (-(v6))
        _ = v61
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v41
    _ = v46
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
