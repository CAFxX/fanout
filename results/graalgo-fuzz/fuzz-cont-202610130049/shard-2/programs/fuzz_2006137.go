package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int; F1 uint8; F2 int }
type GenS5 struct { F0 int64; F1 uint16 }
type GenS6 struct { F0 uint8; F1 int; F2 int64 }
type GenS7 struct { F0 string; F1 int16; F2 uint16 }

func main() {
    var v1 uint16 = 255
    var v2 int16 = -609
    var v3 int8 = -20
    var v4 int32 = 2147483647
    var v5 int8 = v3
    switch -64 {
    case 42:
        var v6 bool = (!false)
        _ = v6
    case -2147483648:
        if (!true) {
            fmt.Println(v3, v1, v5)
        }
    }
    switch -193 {
    case -617:
        if (-95 < v5) {
            var v7 string = "foo bar"
            v8 := GenS1{F0: (-(v4))}
            if ((0 > 2) && false) {
                var v9 int8 = 1
                v10 := []int{4294967295, 1}
                if false {
                    var v11 int = -9223372036854775808
                    if len(v10) >= 2 {
                        v10 = v10[1:len(v10)]
                    }
                    _ = v11
                } else {
                    v12 := GenS2{F0: (^(v1))}
                    if ((v2 >= (v2 * (v2 << ((255) & 63)))) || (true && true)) {
                        fmt.Println(v1, v3, v5)
                        v7 = "foo bar"
                    } else {
                        if (v5 >= (v9 + v9)) {
                            v13 := GenS2{F0: (v1 << ((4294967295) & 63))}
                            _ = v13
                        } else {
                            v7 = v7
                            if false {
                                v9 = -128
                                switch (v7 + (v7 + "\n\t")) {
                                case "hello":
                                    switch v7 {
                                    case "a":
                                        if ((v4 - (v4 >> ((18446744073709551615) & 63))) > (v4 << ((255) & 63))) {
                                            fmt.Println(v3, v2, v3)
                                        }
                                    case "hello":
                                        fmt.Println(v7, v2, v3)
                                    case "":
                                        if len(v10) > 0 {
                                            v10[0] = 9223372036854775807
                                        }
                                    }
                                case "foo bar":
                                    var v14 uint32 = 2147483648
                                    _ = v14
                                case "":
                                    fmt.Println(v7)
                                default:
                                    v15 := GenS3{F0: 18446744073709551615}
                                    _ = v15
                                }
                                v16 := GenS1{F0: (v4 | -2147483648)}
                                _ = v16
                            } else {
                                v7 = v7
                                switch v7 {
                                case "a":
                                    fmt.Println(v3, v3, v3)
                                case "hello":
                                    v5 = v3
                                default:
                                    for fi17 := 0; fi17 < 16; fi17++ {
                                        var v18 bool = false
                                        v2 -= v2
                                        _ = fi17
                                        _ = v18
                                    }
                                }
                            }
                        }
                        v19 := GenS2{F0: v1}
                        for fi20 := 0; fi20 < 17; fi20++ {
                            v21 := GenS2{F0: v1}
                            for fi22 := 0; fi22 < 20; fi22++ {
                                if (!(!true)) {
                                    v2 += (v2 & (^(v2)))
                                }
                                v23 := GenS4{F0: (fi22 + fi22), F1: 2, F2: fi22}
                                var v24 int64 = 9223372036854775807
                                _ = fi22
                                _ = v23
                                _ = v24
                            }
                            _ = fi20
                            _ = v21
                        }
                        _ = v19
                    }
                    switch ("" + "foo bar") {
                    case "foo bar":
                        var v25 float32 = ((float32(((float64((0.1 - 3.14159)) + -2.25) / ((1e10 * -2.25) / -2.25))) + (((0.1 / 3.14159) / 0.1) - 0.1)) * (0.1 + 3.14159))
                        _ = v25
                    case "":
                        for fi26 := 0; fi26 < 14; fi26++ {
                            for fi27 := 0; fi27 < 6; fi27++ {
                                switch "a" {
                                case "\n\t":
                                    if (42 == 371) {
                                        v1 += (v1 & v1)
                                    }
                                case "foo bar":
                                    v5 += (v3 % (((v3 / (((v9 << ((18446744073709551615) & 63))) | 1))) | 1))
                                case "a":
                                    if ((0 >= v2) || (false || (true || (!(!((v1 != v1) && (255 <= 18446744073709551615))))))) {
                                        v1 -= (v1 * (v1 * 764))
                                        var v28 int8 = (v3 | (v5 * v3))
                                        _ = v28
                                    }
                                default:
                                    if false {
                                        v29 := GenS1{F0: v4}
                                        if false {
                                            v5 -= v9
                                            if true {
                                                var v30 uint16 = 160
                                                _ = v30
                                            } else {
                                                fmt.Println(v9, v2, v9)
                                                v4 = v4
                                                v31 := GenS5{F0: -9223372036854775808, F1: (^(v1))}
                                                _ = v31
                                            }
                                            v32 := GenS4{F0: -2147483648, F1: 240, F2: 994}
                                            _ = v32
                                        }
                                        _ = len(v10)
                                        _ = v29
                                    }
                                }
                                if false {
                                    fmt.Println(v3)
                                    v33 := GenS3{F0: 18446744073709551615}
                                    _ = v33
                                } else {
                                    v1 -= v1
                                    v3 = 0
                                }
                                _ = len(v10)
                                _ = fi27
                            }
                            _ = fi26
                        }
                    }
                    _ = v12
                }
                _ = v10
                _ = v9
            } else {
                v1 = (^(v1))
                v34 := GenS6{F0: 255, F1: -257, F2: int64(v4)}
                _ = v34
            }
            _ = v7
            _ = v8
        }
    default:
        switch "\n\t" {
        case "\n\t":
            v5 = 0
        case "a":
            if false {
                v35 := GenS4{F0: -2147483648, F1: 1, F2: 32767}
                _ = v35
            } else {
                fmt.Println(v2)
                v2 = v2
            }
        case "":
            for fi36 := 0; fi36 < 14; fi36++ {
                if true {
                    var v37 int8 = -102
                    switch fi36 {
                    case -9223372036854775808:
                        var v38 uint32 = 607
                        _ = v38
                    }
                    v39 := GenS7{F0: "hello", F1: v2, F2: 255}
                    _ = v37
                    _ = v39
                }
                _ = fi36
            }
        }
    }
    switch "hello" {
    case "":
        var v40 int32 = v4
        _ = v40
    case "a":
        var v41 bool = ((v5 << ((65535) & 63)) > v5)
        _ = v41
    case "hello":
        var v42 uint = 1
        _ = v42
    default:
        v43 := []int{127}
        _ = v43
    }
    switch -9223372036854775808 {
    case 9223372036854775807:
        var v44 uint = 472
        _ = v44
    case -2147483648:
        v3 = v3
    case 1:
        var v45 uint16 = (^(v1))
        _ = v45
    }
    v2 -= v2
    var v46 int8 = 109
    v46 = (v46 + 0)
    if true {
        v4 -= v4
        var v47 float64 = 3.14159
        _ = v47
    } else {
        var v48 float64 = (float64(0.0) - (((0.1 / ((1e10 * 1.5) - 1.5)) / -2.25) + ((1.5 * 1e10) / -2.25)))
        _ = v48
    }
    var v49 int = 127
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v49
    _ = v5
}
