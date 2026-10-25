package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 uint; F1 uint; F2 int8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 bool; F1 int64; F2 uint }
type GenS5 struct { F0 int8; F1 int; F2 bool }

func main() {
    var v1 int32 = -1
    var v2 string = "foo bar"
    v3 := []uint32{65535, 65535}
    for fi4 := 0; fi4 < 4; fi4++ {
        var v5 int64 = 0
        _ = fi4
        _ = v5
    }
    fmt.Println(v2, v1)
    if false {
        if true {
            v2 = (((v2 + v2) + v2) + v2)
        } else {
            switch "" {
            case "foo bar":
                for fi6 := 0; fi6 < 8; fi6++ {
                    fmt.Println(fi6, v1, v2)
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi6
                }
            case "hello":
                var v7 float32 = ((3.14159 * ((0.1 * 0.0) + 1e10)) - 1.5)
                _ = v7
            default:
                if (2 > 2) {
                    for fi8 := 0; fi8 < 3; fi8++ {
                        if (538 > 0) {
                            for fi9 := 0; fi9 < 4; fi9++ {
                                var v10 int32 = (v1 >> ((18446744073709551615) & 63))
                                _ = fi9
                                _ = v10
                            }
                        }
                        _ = fi8
                    }
                    if false {
                        v1 += -2147483648
                        v2 = v2
                    } else {
                        _ = len(v3)
                        if true {
                            switch 127 {
                            case 32768:
                                v1 = (v1 ^ (v1 >> ((65535) & 63)))
                            default:
                                var v11 uint8 = 0
                                _ = v11
                            }
                            v12 := GenS1{F0: v1, F1: 255}
                            v13 := GenS2{F0: 255, F1: 466, F2: 127}
                            _ = v12
                            _ = v13
                        } else {
                            v1 = 229
                        }
                    }
                    if (129 >= 590) {
                        v14 := GenS3{F0: -9223372036854775808}
                        _ = v14
                    }
                } else {
                    if true {
                        var v15 uint = 4294967295
                        _ = v15
                    } else {
                        v1 += v1
                        switch 32768 {
                        case 32768:
                            for fi16 := 0; fi16 < 10; fi16++ {
                                fmt.Println(v1)
                                for fi17 := 0; fi17 < 13; fi17++ {
                                    var v18 int8 = 36
                                    var v19 uint8 = 127
                                    _ = fi17
                                    _ = v18
                                    _ = v19
                                }
                                _ = fi16
                            }
                        case 0:
                            fmt.Println(v1, v2)
                        }
                        v2 = v2
                    }
                }
            }
            if (65535 != 2147483647) {
                v2 = ("foo bar" + v2)
                fmt.Println(v1)
            } else {
                if (false && (true || false)) {
                    switch 1 {
                    case 568:
                        var v20 uint32 = 2147483648
                        _ = v20
                    case -9223372036854775808:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    default:
                        v2 = v2
                    }
                    var v21 int16 = 42
                    _ = v21
                } else {
                    for fi22 := 0; fi22 < 13; fi22++ {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        var v23 int16 = 32767
                        v3 = append(v3, 2)
                        _ = fi22
                        _ = v23
                    }
                    fmt.Println(v2, v2)
                }
                if ((-251 == -28) || (-140 >= 299)) {
                    for fi24 := 0; fi24 < 11; fi24++ {
                        var v25 uint16 = 42
                        _ = fi24
                        _ = v25
                    }
                    for fi26 := 0; fi26 < 19; fi26++ {
                        v27 := GenS4{F0: true, F1: 318, F2: 1}
                        _ = fi26
                        _ = v27
                    }
                    switch "" {
                    case "foo bar":
                        for fi28 := 0; fi28 < 7; fi28++ {
                            switch ("a" + ("hello" + ((v2 + v2) + v2))) {
                            case "\n\t":
                                switch (v2 + v2) {
                                case "hello":
                                    for fi29 := 0; fi29 < 14; fi29++ {
                                        if len(v3) > 0 {
                                            v3[0] = 60
                                        }
                                        _ = fi29
                                    }
                                case "\n\t":
                                    var v30 uint = 255
                                    _ = v30
                                case "a":
                                    v1 = v1
                                }
                            case "a":
                                v2 = (v2 + v2)
                            }
                            for fi31 := 0; fi31 < 6; fi31++ {
                                v32 := GenS5{F0: 1, F1: fi28, F2: true}
                                _ = fi31
                                _ = v32
                            }
                            for fi33 := 0; fi33 < 3; fi33++ {
                                if (4294967295 < 0) {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                }
                                _ = fi33
                            }
                            _ = fi28
                        }
                    default:
                        var v34 float32 = ((0.1 * (0.1 / 0.1)) / 3.14159)
                        _ = v34
                    }
                }
            }
            var v35 bool = true
            _ = v35
        }
        for fi36 := 0; fi36 < 1; fi36++ {
            var v37 int = fi36
            if len(v3) > 0 {
                v3[0] = 56
            }
            _ = fi36
            _ = v37
        }
        v2 = ""
    } else {
        if (65535 == 2) {
            var v38 int8 = 0
            v39 := GenS2{F0: 18446744073709551615, F1: 429, F2: 1}
            _ = v38
            _ = v39
        }
    }
    var v40 float32 = -2.25
    for fi41 := 0; fi41 < 11; fi41++ {
        for fi42 := 0; fi42 < 3; fi42++ {
            var v43 uint16 = 740
            v43 = 1
            fmt.Println(v2, v43, v1)
            fmt.Println(v1)
            _ = fi42
            _ = v43
        }
        if (1 < -1) {
            v1 += (v1 | v1)
            for fi44 := 0; fi44 < 10; fi44++ {
                fmt.Println(v40, v40, v1)
                v2 = "foo bar"
                var v45 int8 = 42
                _ = fi44
                _ = v45
            }
            switch 2147483647 {
            case -9223372036854775808:
                v3 = append(v3, 2147483647)
            }
        } else {
            var v46 uint16 = 2
            var v47 bool = true
            if (!v47) {
                v48 := GenS1{F0: int32(fi41), F1: 2}
                _ = v48
            } else {
                v46 += v46
                var v49 int = (fi41 - fi41)
                _ = v49
            }
            _ = v46
            _ = v47
        }
        _ = fi41
    }
    fmt.Println(v2)
    if (-687 <= 854) {
        var v50 bool = (0 > 9223372036854775807)
        _ = v50
    }
    fmt.Println(v40, v2, v1)
    var v51 int64 = 9223372036854775807
    var v52 int16 = -529
    v40 = (0.0 * 3.14159)
    fmt.Println(v52)
    v3 = append(v3, 2147483647)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v40:", v40)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v40
    _ = v51
    _ = v52
}
