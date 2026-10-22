package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 int8; F1 uint32; F2 int32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int32; F1 uint16 }

func main() {
    var v1 float32 = -2.25
    v2 := []uint32{65535, 0, 1}
    v3 := []uint32{2147483647}
    var v4 int64 = -457
    v1 = 0.1
    var v5 int64 = v4
    v4 = -1
    v6 := GenS1{F0: 401, F1: 100}
    v7 := GenS1{F0: 2147483648, F1: 61}
    for fi8 := 0; fi8 < 4; fi8++ {
        if true {
            var v9 int32 = -2147483648
            var v10 float64 = (1e10 - float64(3.14159))
            var v11 string = "hello"
            _ = v10
            _ = v11
            _ = v9
        } else {
            var v12 int16 = -1
            _ = v12
        }
        _ = fi8
    }
    v2 = append(v2, 2)
    v13 := GenS1{F0: 4294967295, F1: 100}
    if (true || false) {
        v1 = (v1 - float32(1.5))
    } else {
        for fi14 := 0; fi14 < 17; fi14++ {
            for fi15 := 0; fi15 < 9; fi15++ {
                var v16 string = "\n\t"
                fmt.Println(fi15)
                _ = fi15
                _ = v16
            }
            if (127 != fi14) {
                var v17 int = (fi14 + 9223372036854775807)
                switch (fi14 >> ((1) & 63)) {
                case 738:
                    v18 := GenS1{F0: 382, F1: -68}
                    _ = v18
                }
                _ = v17
            } else {
                var v19 float32 = (v1 * (1.5 * (v1 + v1)))
                v20 := GenS1{F0: 2147483647, F1: 1}
                v4 = 42
                _ = v19
                _ = v20
            }
            v5 = v4
            _ = fi14
        }
    }
    var v21 uint16 = 525
    for fi22 := 0; fi22 < 7; fi22++ {
        if false {
            if ("foo bar" == "foo bar") {
                if (((v4 / ((v4) | 1)) >= (v4 - 0)) || false) {
                    v21 += (^(v21))
                }
                v5 += v5
                var v23 int32 = 0
                _ = v23
            }
            v5 = (v4 & v5)
            switch fi22 {
            case -2147483648:
                v4 = 0
            case 32768:
                v4 += (-(v4))
            case 0:
                var v24 uint32 = 82
                _ = v24
            default:
                switch fi22 {
                case 32768:
                    for fi25 := 0; fi25 < 12; fi25++ {
                        if len(v2) > 0 {
                            v2[0] = 4294967295
                        }
                        v1 = (1e10 - v1)
                        _ = fi25
                    }
                case 916:
                    var v26 uint32 = 2
                    _ = v26
                case -704:
                    fmt.Println(fi22)
                default:
                    switch ("a" + ("foo bar" + ("hello" + ("hello" + "a")))) {
                    case "":
                        for fi27 := 0; fi27 < 14; fi27++ {
                            v28 := GenS1{F0: 42, F1: -64}
                            if (!false) {
                                if len(v3) > 0 {
                                    v3[0] = 65535
                                }
                            } else {
                                v1 -= 1e10
                                var v29 bool = (86 != 1)
                                v1 += 1.5
                                if (-2147483648 <= -791) {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    var v30 uint = 42
                                    switch (("hello" + "") + "foo bar") {
                                    case "hello":
                                        v30 = v30
                                    case "\n\t":
                                        v31 := GenS2{F0: -4, F1: 0, F2: -1}
                                        _ = v31
                                    case "":
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    }
                                    var v32 int32 = int32(v21)
                                    v33 := GenS2{F0: 0, F1: 4294967295, F2: v32}
                                    v21 += (v21 << ((791) & 63))
                                    _ = v30
                                    _ = v32
                                    _ = v33
                                } else {
                                    switch "\n\t" {
                                    case "":
                                        v2 = append(v2, 91)
                                    case "foo bar":
                                        v4 += (v4 + (^(v5)))
                                    }
                                    if (!true) {
                                        switch fi22 {
                                        case 128:
                                            v21 = 42
                                            var v34 uint16 = 335
                                            v35 := GenS2{F0: -1, F1: 2, F2: 0}
                                            _ = v34
                                            _ = v35
                                        case 1:
                                            v1 -= v1
                                        case -1:
                                            v21 = (^(v21))
                                        }
                                    } else {
                                        v36 := GenS3{F0: "foo bar"}
                                        _ = v36
                                    }
                                }
                                if true {
                                    v37 := GenS1{F0: 42, F1: 42}
                                    _ = v37
                                } else {
                                    var v38 string = ("hello" + "\n\t")
                                    if (!false) {
                                        v39 := GenS3{F0: v38}
                                        if len(v3) > 0 {
                                            v3[0] = 65535
                                        }
                                        var v40 uint = 65535
                                        _ = v39
                                        _ = v40
                                    }
                                    _ = v38
                                }
                                _ = v29
                            }
                            _ = fi27
                            _ = v28
                        }
                    case "a":
                        v41 := GenS3{F0: ""}
                        _ = v41
                    case "hello":
                        v1 += 3.14159
                    }
                }
            }
        } else {
            for fi42 := 0; fi42 < 10; fi42++ {
                var v43 float64 = 3.14159
                _ = fi42
                _ = v43
            }
            var v44 uint8 = 224
            fmt.Println(fi22)
            _ = v44
        }
        if len(v2) > 0 {
            v2[0] = 649
        }
        for fi45 := 0; fi45 < 17; fi45++ {
            v21 -= (v21 ^ v21)
            switch "\n\t" {
            case "":
                v21 += (^(v21))
            case "\n\t":
                v1 = 3.14159
            }
            if false {
                v21 += 65535
            } else {
                v5 -= (v5 >> ((255) & 63))
                if (false || (!(fi22 >= (fi22 * fi22)))) {
                    fmt.Println(v4)
                }
                var v46 int8 = -128
                _ = v46
            }
            _ = fi45
        }
        _ = fi22
    }
    var v47 bool = false
    if (18446744073709551615 >= 1) {
        v4 = (v4 / ((v4) | 1))
    }
    v1 = v1
    if false {
        for fi48 := 0; fi48 < 3; fi48++ {
            v49 := GenS3{F0: ("" + "hello")}
            if len(v2) > 0 {
                v2[0] = 0
            }
            var v50 int64 = v4
            _ = fi48
            _ = v49
            _ = v50
        }
        var v51 uint16 = v21
        _ = v51
    } else {
        if len(v3) > 0 {
            v3[0] = 2
        }
        switch (((("foo bar" + "hello") + "foo bar") + "") + "") {
        case "":
            v3 = append(v3, 2147483648)
        }
    }
    v52 := GenS4{F0: -1, F1: v21}
    for fi53 := 0; fi53 < 17; fi53++ {
        for fi54 := 0; fi54 < 5; fi54++ {
            var v55 uint8 = 1
            fmt.Println(v47)
            var v56 int8 = 11
            _ = fi54
            _ = v55
            _ = v56
        }
        var v57 int32 = 0
        v47 = (v21 > (v21 / ((v21) | 1)))
        _ = fi53
        _ = v57
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v52
    _ = v6
    _ = v7
}
