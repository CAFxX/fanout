package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 uint16 }
type GenS2 struct { F0 uint32; F1 uint16; F2 uint16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint64; F1 string; F2 uint }
type GenS6 struct { F0 uint64; F1 uint32; F2 int32 }

func main() {
    var v1 float32 = ((((1e10 + 1e10) * -2.25) + float32((3.14159 + 1e10))) + 3.14159)
    v2 := []uint32{}
    var v3 bool = ((!false) || false)
    var v4 int = -97
    fmt.Println(v4)
    switch (v4 & (-(v4))) {
    case -525:
        var v5 string = "a"
        _ = v5
    default:
        for fi6 := 0; fi6 < 7; fi6++ {
            v7 := GenS1{F0: (("foo bar" + ("a" + (("foo bar" + "\n\t") + ("a" + "")))) + ("foo bar" + ("hello" + ("a" + "")))), F1: 749, F2: 65535}
            v8 := GenS1{F0: "foo bar", F1: 1, F2: 42}
            _ = fi6
            _ = v7
            _ = v8
        }
    }
    v9 := GenS1{F0: "\n\t", F1: 40, F2: 255}
    v10 := GenS2{F0: 327, F1: 42, F2: 42}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if (!true) {
        v1 = v1
    }
    for fi11 := 0; fi11 < 3; fi11++ {
        if (0 <= 1) {
            fmt.Println(v3)
            var v12 int16 = int16(v4)
            _ = v12
        }
        _ = fi11
    }
    if len(v2) > 0 {
        v2[0] = 1
    }
    var v13 float64 = 0.0
    for fi14 := 0; fi14 < 8; fi14++ {
        switch 9223372036854775807 {
        case 32768:
            v2 = append(v2, 4294967295)
        default:
            v15 := GenS1{F0: "a", F1: 42, F2: 255}
            _ = v15
        }
        var v16 uint64 = 495
        _ = fi14
        _ = v16
    }
    var v17 int8 = 42
    for fi18 := 0; fi18 < 8; fi18++ {
        switch (fi18 ^ -32) {
        case 32768:
            v13 += 1.5
        case -922:
            if v3 {
                if (2 > 65535) {
                    if len(v2) > 0 {
                        v2[0] = 321
                    }
                } else {
                    if len(v2) > 0 {
                        v2[0] = 42
                    }
                    if v3 {
                        v4 += (-(fi18))
                        v19 := GenS3{F0: (42 == 18446744073709551615)}
                        _ = v19
                    }
                }
                switch -236 {
                case 32768:
                    v17 -= v17
                case 626:
                    fmt.Println(v17, v3, v3)
                }
            }
        case 2147483647:
            if v3 {
                fmt.Println(fi18, v17, fi18)
                fmt.Println(v3)
            }
        default:
            v20 := GenS4{F0: 255}
            _ = v20
        }
        if false {
            v1 -= v1
            v2 = append(v2, 471)
            var v21 int8 = v17
            _ = v21
        } else {
            _ = len(v2)
        }
        switch fi18 {
        case 128:
            fmt.Println(v17, v3)
        }
        _ = fi18
    }
    if v3 {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    } else {
        v22 := GenS5{F0: 4294967295, F1: "a", F2: 4294967295}
        if len(v2) > 0 {
            v2[0] = 4294967295
        }
        _ = v22
    }
    v23 := GenS4{F0: 396}
    fmt.Println(v1, v4, v1)
    v1 -= 1.5
    if ((v4 >> ((0) & 63)) <= int(v17)) {
        v17 = -1
    } else {
        v24 := GenS6{F0: 586, F1: 2147483648, F2: 1}
        var v25 int = (v4 & (-(v4)))
        switch v25 {
        case 32767:
            fmt.Println(v17, v4, v17)
        default:
            for fi26 := 0; fi26 < 7; fi26++ {
                for fi27 := 0; fi27 < 6; fi27++ {
                    switch v25 {
                    case 978:
                        if v3 {
                            v28 := GenS6{F0: 0, F1: 2147483648, F2: 42}
                            switch (fi27 % ((4294967295) | 1)) {
                            case -130:
                                var v29 uint16 = 977
                                _ = v29
                            }
                            if (!v3) {
                                if len(v2) > 0 {
                                    v2[0] = 1
                                }
                                fmt.Println(v1)
                                v30 := GenS6{F0: 320, F1: 632, F2: 2147483647}
                                _ = v30
                            }
                            _ = v28
                        }
                        switch "hello" {
                        case "\n\t":
                            if ((v17 % ((v17) | 1)) >= 1) {
                                fmt.Println(v3)
                                v13 = 1e10
                                var v31 uint64 = 42
                                if v3 {
                                    v13 = 0.1
                                }
                                _ = v31
                            }
                        }
                        v13 -= (-2.25 - v13)
                    case -9223372036854775808:
                        v1 += v1
                    case 1:
                        var v32 uint64 = 18446744073709551615
                        _ = v32
                    }
                    _ = fi27
                }
                for fi33 := 0; fi33 < 5; fi33++ {
                    _ = len(v2)
                    if len(v2) > 0 {
                        v2[0] = 42
                    }
                    if ("hello" != "") {
                        var v34 uint = 42
                        var v35 uint16 = 42
                        _ = v34
                        _ = v35
                    } else {
                        switch fi33 {
                        case 0:
                            var v36 uint64 = 1
                            _ = v36
                        default:
                            v2 = append(v2, 1)
                        }
                        switch v25 {
                        case -2147483648:
                            v4 = v25
                        case 4294967295:
                            switch (fi26 ^ fi33) {
                            case -442:
                                v3 = (!(true && (v3 && true)))
                            case 2147483647:
                                v2 = append(v2, 65535)
                            case -671:
                                v4 = (^(fi33))
                            }
                        case 9223372036854775807:
                            var v37 uint32 = 2147483648
                            _ = v37
                        default:
                            var v38 int64 = -1
                            _ = v38
                        }
                    }
                    _ = fi33
                }
                _ = fi26
            }
        }
        _ = v24
        _ = v25
    }
    fmt.Println(v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v13
    _ = v17
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v9
}
