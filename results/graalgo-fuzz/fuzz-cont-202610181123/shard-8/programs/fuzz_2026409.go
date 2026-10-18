package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 uint; F1 int64; F2 bool }
type GenS3 struct { F0 uint32; F1 uint16 }
type GenS4 struct { F0 uint32; F1 uint8 }
type GenS5 struct { F0 uint32 }
func genFn0() int64 {
    if (0.0 < 3.14159) {
        var v5 float64 = -2.25
        var v6 uint = 473
        switch "foo bar" {
        case "foo bar":
            v7 := []int{}
            _ = v7
        case "hello":
            var v8 uint8 = 2
            _ = v8
        }
        _ = v5
        _ = v6
    }
    var v9 int8 = 1
    _ = v9
    return 42
}
func genFn1(p0 uint8) uint16 {
    v10 := []int{32767, -655}
    var v11 float32 = (3.14159 * ((3.14159 / (1.5 * -2.25)) + -2.25))
    var v12 int64 = 9223372036854775807
    _ = p0
    _ = v10
    _ = v11
    _ = v12
    return 439
}

func main() {
    var v1 uint8 = 127
    v2 := []uint32{678, 65535, 2147483648, 449}
    var v3 int = 42
    var v4 uint16 = 65535
    var v13 int16 = 32767
    switch v3 {
    case 128:
        if len(v2) > 0 {
            v2[0] = 4294967295
        }
    }
    v2 = append(v2, 1)
    var v14 int32 = -310
    v13 += v13
    v13 = v13
    var v15 int32 = v14
    v16 := GenS1{F0: 4294967295, F1: v3}
    var v17 uint64 = 0
    fmt.Println(v14, v13)
    for fi18 := 0; fi18 < 18; fi18++ {
        _ = len(v2)
        v2 = append(v2, 181)
        var v19 int64 = 0
        _ = fi18
        _ = v19
    }
    switch "foo bar" {
    case "hello":
        if len(v2) > 0 {
            v2[0] = 271
        }
    case "a":
        if ((^(v13)) < 42) {
            v14 += v14
            if (-32768 == -32768) {
                fmt.Println(v1, v13, v13)
                var v20 uint32 = 365
                var v21 float32 = (float32(float64((3.14159 / (0.0 - (0.0 + ((1e10 - 0.0) * 3.14159)))))) / 1.0)
                _ = v20
                _ = v21
            }
        }
    case "":
        _ = len(v2)
    default:
        v22 := GenS2{F0: 4294967295, F1: -1, F2: (1.5 == -2.25)}
        _ = v22
    }
    fmt.Println(v1, v14, v1)
    var v23 int8 = -128
    v15 = v15
    if (-32768 <= v13) {
        var v24 uint8 = v1
        switch v3 {
        case 32768:
            if (42 <= uint32(v4)) {
                v25 := GenS3{F0: 14, F1: v4}
                fmt.Println(v14, v4, v15)
                v26 := GenS3{F0: uint32(v4), F1: genFn1(1)}
                _ = v25
                _ = v26
            } else {
                v27 := GenS1{F0: (v17 | 155), F1: v3}
                _ = v27
            }
        default:
            v28 := GenS4{F0: 583, F1: v24}
            _ = v28
        }
        _ = v24
    } else {
        for fi29 := 0; fi29 < 17; fi29++ {
            switch 0 {
            case 32768:
                var v30 uint32 = 418
                _ = v30
            case 127:
                v1 = (v1 + (^(v1)))
            }
            _ = fi29
        }
        switch "foo bar" {
        case "a":
            for fi31 := 0; fi31 < 18; fi31++ {
                if true {
                    var v32 float32 = -2.25
                    switch (-(v3)) {
                    case 32767:
                        v14 += int32(v23)
                    case 2147483647:
                        switch v3 {
                        case 32767:
                            if (!(0.1 > (-2.25 - 0.0))) {
                                if false {
                                    var v33 uint16 = (v4 & 724)
                                    _ = v33
                                }
                                fmt.Println(v1)
                                v23 = int8(v4)
                            } else {
                                if (-9223372036854775808 == -9223372036854775808) {
                                    fmt.Println(v17, v17, v14)
                                    v34 := GenS5{F0: 4294967295}
                                    v13 = (v13 / ((195) | 1))
                                    _ = v34
                                }
                                _ = len(v2)
                                if ((v23 >> ((4294967295) & 63)) == (v23 ^ (v23 - (v23 + v23)))) {
                                    switch "hello" {
                                    case "foo bar":
                                        v13 += v13
                                    case "a":
                                        for fi35 := 0; fi35 < 3; fi35++ {
                                            if len(v2) > 0 {
                                                v2[0] = 4294967295
                                            }
                                            if false {
                                                v3 = -336
                                                v32 = (((0.1 - v32) / (float32((float64(v32) * 0.0)) + v32)) * (0.1 / v32))
                                                v36 := GenS1{F0: 18446744073709551615, F1: (fi35 << ((540) & 63))}
                                                _ = v36
                                            } else {
                                                var v37 bool = true
                                                v15 -= v14
                                                var v38 int8 = 39
                                                _ = v37
                                                _ = v38
                                            }
                                            v3 = fi35
                                            _ = fi35
                                        }
                                    default:
                                        if (("hello" + "foo bar") < (("\n\t" + "foo bar") + ("foo bar" + ""))) {
                                            if (!((((v1 < 255) || ((float32(1e10) - (0.1 + v32)) >= 0.1)) || (42 != (v17 >> ((65535) & 63)))) && false)) {
                                                v39 := GenS1{F0: 42, F1: (fi31 << ((255) & 63))}
                                                _ = v39
                                            }
                                            v15 = -19
                                            switch (fi31 >> ((4294967295) & 63)) {
                                            case 0:
                                                if true {
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                } else {
                                                    fmt.Println(fi31, v4)
                                                }
                                            case -728:
                                                if (!((!true) || true)) {
                                                    if (true || false) {
                                                        var v40 uint = 255
                                                        _ = v40
                                                    }
                                                } else {
                                                    fmt.Println(v1, v13)
                                                    v15 += (v14 & v14)
                                                    var v41 bool = false
                                                    _ = v41
                                                }
                                            default:
                                                fmt.Println(v32, v4)
                                            }
                                        }
                                    }
                                    for fi42 := 0; fi42 < 5; fi42++ {
                                        var v43 int = fi31
                                        var v44 bool = (284 < fi31)
                                        _ = fi42
                                        _ = v43
                                        _ = v44
                                    }
                                }
                            }
                        case -9223372036854775808:
                            v1 -= v1
                        case 1:
                            fmt.Println(v15)
                        }
                    }
                    _ = v32
                } else {
                    var v45 int32 = v14
                    var v46 uint64 = (^(v17))
                    for fi47 := 0; fi47 < 1; fi47++ {
                        fmt.Println(v23, v1)
                        _ = fi47
                    }
                    _ = v45
                    _ = v46
                }
                if false {
                    var v48 uint = 2
                    for fi49 := 0; fi49 < 19; fi49++ {
                        switch (("\n\t" + ("hello" + "")) + "foo bar") {
                        case "":
                            v50 := GenS3{F0: 2147483648, F1: 32767}
                            _ = v50
                        case "\n\t":
                            if (float32(1e10) < -2.25) {
                                if true {
                                    var v51 uint8 = 127
                                    v23 = (^(v23))
                                    var v52 bool = ((^(v48)) <= (^(v48)))
                                    var v53 float64 = 0.1
                                    var v54 uint16 = v4
                                    _ = v51
                                    _ = v52
                                    _ = v53
                                    _ = v54
                                }
                            } else {
                                var v55 uint32 = 4294967295
                                if len(v2) > 0 {
                                    v2[0] = v55
                                }
                                _ = v55
                            }
                        }
                        _ = fi49
                    }
                    _ = v48
                }
                v1 -= (v1 ^ 127)
                _ = fi31
            }
        }
        v56 := GenS3{F0: 2147483648, F1: (^(v4))}
        _ = v56
    }
    v3 += v3
    var v57 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v57
}
