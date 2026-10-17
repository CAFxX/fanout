package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 int64; F2 uint8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int16; F1 int32; F2 int16 }
type GenS5 struct { F0 int8; F1 uint; F2 int }
type GenS6 struct { F0 int; F1 uint64; F2 int64 }
type GenS7 struct { F0 int32; F1 bool; F2 string }
type GenS8 struct { F0 uint; F1 int32; F2 int32 }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 int16 }
func genFn0(p0 bool, p1 uint16) int16 {
    v4 := GenS1{F0: 1}
    p0 = (0.0 == 0.0)
    if (3.14159 <= ((1e10 + 3.14159) + 3.14159)) {
        fmt.Println(p0, p0)
        fmt.Println(p0)
    }
    switch "a" {
    case "hello":
        if (!(true && p0)) {
            var v5 int16 = 1000
            var v6 bool = p0
            v7 := []int{1, -404, 1}
            _ = v5
            _ = v6
            _ = v7
        } else {
            if p0 {
                v8 := []int{42, 32767, 32767, -961}
                var v9 uint = 468
                p0 = (p0 || (174 != -224))
                _ = v8
                _ = v9
            } else {
                if true {
                    v10 := GenS1{F0: 42}
                    if ("foo bar" < "") {
                        v11 := GenS2{F0: 18446744073709551615, F1: -848, F2: 255}
                        _ = v11
                    } else {
                        fmt.Println(p0)
                    }
                    p0 = (-653 > 9223372036854775807)
                    _ = v10
                } else {
                    if true {
                        for fi12 := 0; fi12 < 19; fi12++ {
                            if false {
                                v13 := GenS3{F0: 1}
                                for fi14 := 0; fi14 < 8; fi14++ {
                                    if p0 {
                                        var v15 uint16 = 42
                                        var v16 uint64 = 124
                                        if p0 {
                                            var v17 float32 = float32(float64(((1e10 - 3.14159) / (0.1 - 0.0))))
                                            if p0 {
                                                p0 = p0
                                                p1 += (^(v15))
                                            }
                                            _ = v17
                                        } else {
                                            var v18 uint64 = (v16 << ((255) & 63))
                                            _ = v18
                                        }
                                        var v19 int32 = 42
                                        _ = v15
                                        _ = v16
                                        _ = v19
                                    }
                                    var v20 uint16 = (p1 & 799)
                                    _ = fi14
                                    _ = v20
                                }
                                _ = v13
                            } else {
                                v21 := []int{(^(fi12)), (fi12 / ((-2147483648) | 1))}
                                _ = len(v21)
                                _ = v21
                            }
                            _ = fi12
                        }
                        fmt.Println(p0)
                        switch ("a" + ("foo bar" + "")) {
                        case "a":
                            var v22 uint8 = 0
                            _ = v22
                        }
                    } else {
                        switch -87 {
                        case -55:
                            for fi23 := 0; fi23 < 3; fi23++ {
                                var v24 uint = 1
                                _ = fi23
                                _ = v24
                            }
                        }
                    }
                    v25 := GenS4{F0: 42, F1: 2147483647, F2: -1}
                    _ = v25
                }
            }
            v26 := GenS1{F0: 42}
            _ = v26
        }
    case "\n\t":
        var v27 uint32 = 1
        _ = v27
    case "":
        v28 := []int{2147483647, 236}
        _ = v28
    }
    _ = p0
    _ = p1
    _ = v4
    return 1000
}
func genFn1(p0 uint16) uint16 {
    if (583 <= 0) {
        var v29 int32 = 2147483647
        v30 := GenS5{F0: -128, F1: 388, F2: 4294967295}
        v29 += (v29 ^ v29)
        _ = v29
        _ = v30
    }
    v31 := GenS6{F0: 2147483647, F1: uint64(p0), F2: -931}
    fmt.Println(p0)
    _ = p0
    _ = v31
    return p0
}

func main() {
    var v1 float64 = 0.0
    v2 := []string{("hello" + (("foo bar" + "a") + "foo bar")), ("" + (("foo bar" + "") + "foo bar")), "foo bar"}
    var v3 float64 = ((((v1 - float64((3.14159 * ((-2.25 + ((0.1 / (3.14159 * 3.14159)) * 3.14159)) * 0.0)))) * ((1e10 / v1) * 1.5)) / ((3.14159 / v1) * -2.25)) - (((0.0 * (v1 * (v1 - v1))) - (((v1 / v1) * 0.0) / v1)) - 1e10))
    v2 = append(v2, "hello")
    v1 = 0.1
    var v32 uint64 = 202
    v1 -= v3
    var v33 uint = 255
    v34 := GenS7{F0: -793, F1: ((32767 >= uint16(v33)) || (false || false)), F2: ""}
    var v35 int64 = 0
    fmt.Println(v1, v32)
    switch ((("\n\t" + "foo bar") + ("a" + "hello")) + "") {
    case "foo bar":
        v32 = v32
    case "":
        v36 := GenS7{F0: -2147483648, F1: true, F2: "a"}
        _ = v36
    case "hello":
        v37 := GenS7{F0: 1, F1: false, F2: "\n\t"}
        _ = v37
    }
    for fi38 := 0; fi38 < 17; fi38++ {
        v33 -= 2
        var v39 int32 = 891
        var v40 float64 = ((v1 * v1) / ((v3 * v3) + 0.0))
        _ = fi38
        _ = v39
        _ = v40
    }
    var v41 bool = (-680 > 1)
    v42 := GenS6{F0: -1, F1: (^(v32)), F2: -9223372036854775808}
    var v43 int = 2147483647
    v43 = v43
    v35 += (v35 + v35)
    if false {
        var v44 int8 = -128
        switch (v43 >> ((42) & 63)) {
        case 629:
            var v45 string = ""
            _ = v45
        case 631:
            v46 := GenS4{F0: -1, F1: 1, F2: genFn0((23 > 1), 835)}
            _ = v46
        case -2147483648:
            var v47 int16 = 0
            _ = v47
        default:
            switch 856 {
            case 32768:
                var v48 float32 = (-2.25 + (1e10 * (0.0 / 1.0)))
                _ = v48
            case -564:
                v35 += v35
            case -946:
                var v49 int = v43
                _ = v49
            }
        }
        fmt.Println(v41, v1, v43)
        _ = v44
    } else {
        for fi50 := 0; fi50 < 13; fi50++ {
            for fi51 := 0; fi51 < 6; fi51++ {
                v52 := GenS8{F0: 4294967295, F1: -1, F2: 844}
                _ = fi51
                _ = v52
            }
            var v53 float64 = 3.14159
            _ = fi50
            _ = v53
        }
        switch "\n\t" {
        case "\n\t":
            if (0 != 1) {
                v54 := GenS9{F0: 0}
                fmt.Println(v3, v33)
                var v55 bool = v41
                _ = v54
                _ = v55
            } else {
                v56 := GenS10{F0: 1000}
                for fi57 := 0; fi57 < 15; fi57++ {
                    var v58 string = "foo bar"
                    _ = fi57
                    _ = v58
                }
                _ = v56
            }
        case "foo bar":
            fmt.Println(v41)
        default:
            v41 = false
        }
    }
    for fi59 := 0; fi59 < 15; fi59++ {
        v33 -= v33
        _ = fi59
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v41
    _ = v42
    _ = v43
}
