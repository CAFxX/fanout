package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int8; F1 int32 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint8; F1 bool }
func genFn0() uint64 {
    var v4 bool = true
    var v5 float64 = float64(1.5)
    if (!v4) {
        fmt.Println(v4, v4)
        var v6 uint16 = 1
        v5 += v5
        _ = v6
    } else {
        var v7 string = ""
        _ = v7
    }
    var v8 uint64 = 4294967295
    _ = v4
    _ = v5
    _ = v8
    return (v8 * 970)
}

func main() {
    v1 := make(map[string]bool)
    var v2 float64 = 3.14159
    v3 := make(map[int]string)
    var v9 float64 = ((1.5 - v2) + v2)
    v10 := []int{32768, 32768, -873, 165}
    var v11 int64 = -1
    v12 := GenS1{F0: false}
    for fi13 := 0; fi13 < 9; fi13++ {
        for fi14 := 0; fi14 < 18; fi14++ {
            if (!(0 != 42)) {
                if false {
                    if true {
                        var v15 int8 = 1
                        _ = len(v10)
                        var v16 float32 = 1.5
                        v17 := GenS2{F0: 18446744073709551615, F1: (("foo bar" + "\n\t") + "foo bar")}
                        _ = len(v3)
                        var v18 float32 = v16
                        _ = v15
                        _ = v16
                        _ = v17
                        _ = v18
                    }
                    v2 = 1.5
                    var v19 int8 = -109
                    v20 := GenS3{F0: 42}
                    switch "" {
                    case "a":
                        var v21 string = "a"
                        _ = v21
                    case "hello":
                        delete(v3, fi13)
                    case "foo bar":
                        v11 = 166
                    }
                    _ = v19
                    _ = v20
                }
                if (366 <= 42) {
                    if ((2 < 1) && false) {
                        if true {
                            var v22 float32 = 1e10
                            _ = v22
                        }
                        var v23 uint8 = 255
                        _ = v23
                    }
                    fmt.Println(fi13)
                    var v24 int32 = -2147483648
                    var v25 int64 = (v11 | (v11 % ((-383) | 1)))
                    var v26 uint = 0
                    _ = v24
                    _ = v25
                    _ = v26
                } else {
                    if ("a" == ("hello" + "foo bar")) {
                        if (!(!(1 <= 2))) {
                            var v27 int = fi13
                            _ = v27
                        } else {
                            if (0 > 42) {
                                var v28 uint16 = 1
                                var v29 int64 = v11
                                _ = v28
                                _ = v29
                            } else {
                                var v30 float32 = 0.0
                                var v31 int8 = 68
                                _ = v30
                                _ = v31
                            }
                            v2 -= (v2 - 1.5)
                        }
                        if false {
                            if ((!false) && false) {
                                v32 := GenS2{F0: 65535, F1: "\n\t"}
                                delete(v3, 4294967295)
                                var v33 float64 = -2.25
                                _ = v32
                                _ = v33
                            }
                            if ((!true) && (!true)) {
                                if true {
                                    var v34 int32 = 2147483647
                                    if len(v10) >= 2 {
                                        v10 = v10[1:len(v10)]
                                    }
                                    v35 := v3[fi13]
                                    _ = v34
                                    _ = v35
                                }
                                fmt.Println(v9, v2)
                                var v36 int32 = 229
                                v36 = 1
                                v36 -= (v36 & (^(v36)))
                                _ = v36
                            }
                        }
                        v9 -= v9
                        v37 := GenS3{F0: 980}
                        fmt.Println(v11)
                        _ = v37
                    }
                }
                v2 = float64((0.0 + 1.5))
            }
            var v38 bool = true
            _ = fi14
            _ = v38
        }
        for fi39 := 0; fi39 < 4; fi39++ {
            if (3.14159 <= ((1.5 + ((-2.25 - 0.0) * (float32(float64(1.5)) - (3.14159 + 3.14159)))) * (1e10 * 1e10))) {
                _ = len(v10)
            } else {
                var v40 uint = 65535
                fmt.Println(v9, v2, v2)
                var v41 int8 = 15
                _ = v40
                _ = v41
            }
            _ = fi39
        }
        _ = fi13
    }
    var v42 uint8 = 76
    v43 := GenS1{F0: (("" + "a") > "hello")}
    for fi44 := 0; fi44 < 3; fi44++ {
        switch "\n\t" {
        case "hello":
            v11 -= (^(v11))
        case "foo bar":
            switch "" {
            case "hello":
                switch "" {
                case "foo bar":
                    v9 += v9
                case "\n\t":
                    v10 = append(v10, (^(fi44)))
                case "":
                    v45 := GenS4{F0: -58, F1: 2147483647}
                    _ = v45
                }
            default:
                switch (fi44 / ((fi44) | 1)) {
                case 1:
                    for fi46 := 0; fi46 < 15; fi46++ {
                        v47 := GenS3{F0: 255}
                        _ = fi46
                        _ = v47
                    }
                default:
                    if true {
                        if (!false) {
                            v2 = 1e10
                        } else {
                            var v48 uint64 = genFn0()
                            _ = v48
                        }
                        v2 -= v9
                        delete(v3, 529)
                    }
                }
            }
        case "":
            var v49 int32 = 432
            _ = v49
        default:
            var v50 float32 = 0.1
            _ = v50
        }
        _ = fi44
    }
    for fi51 := 0; fi51 < 6; fi51++ {
        _ = len(v10)
        _ = fi51
    }
    switch (("hello" + "a") + "") {
    case "foo bar":
        var v52 float64 = v9
        _ = v52
    default:
        if false {
            switch 42 {
            case -9223372036854775808:
                if true {
                    if (461 >= 2) {
                        switch 42 {
                        case 2147483647:
                            for fi53 := 0; fi53 < 7; fi53++ {
                                v11 += (v11 << ((4294967295) & 63))
                                var v54 int32 = 42
                                _ = fi53
                                _ = v54
                            }
                        case 380:
                            fmt.Println(v11)
                        }
                        var v55 int32 = -2147483648
                        _ = v55
                    } else {
                        var v56 int64 = (v11 << ((1) & 63))
                        if true {
                            v42 -= (v42 >> ((1) & 63))
                            v57 := GenS5{F0: 635}
                            var v58 string = ((("a" + "") + "foo bar") + ("\n\t" + ("hello" + ("foo bar" + "hello"))))
                            _ = v57
                            _ = v58
                        }
                        _ = v56
                    }
                    if false {
                        for fi59 := 0; fi59 < 17; fi59++ {
                            fmt.Println(v42)
                            v2 = v2
                            _ = fi59
                        }
                        if len(v10) >= 2 {
                            v10 = v10[1:len(v10)]
                        }
                        v60 := GenS5{F0: 4294967295}
                        _ = v60
                    } else {
                        for fi61 := 0; fi61 < 19; fi61++ {
                            if len(v10) >= 2 {
                                v10 = v10[1:len(v10)]
                            }
                            _ = fi61
                        }
                        var v62 uint16 = 384
                        _ = v62
                    }
                    v63 := GenS6{F0: "foo bar"}
                    _ = v63
                }
            case 713:
                for fi64 := 0; fi64 < 7; fi64++ {
                    v42 = (v42 >> ((42) & 63))
                    var v65 int32 = 2147483647
                    var v66 int64 = (-(v11))
                    _ = fi64
                    _ = v65
                    _ = v66
                }
            }
        } else {
            switch 1 {
            case -2147483648:
                v11 -= (^(v11))
            default:
                v67 := GenS7{F0: 128, F1: (true || (true && false))}
                _ = v67
            }
        }
    }
    for fi68 := 0; fi68 < 8; fi68++ {
        for fi69 := 0; fi69 < 13; fi69++ {
            fmt.Println(fi69, v42, v2)
            _ = fi69
        }
        v70 := GenS3{F0: 1}
        _ = fi68
        _ = v70
    }
    var v71 uint32 = 65535
    switch "hello" {
    case "\n\t":
        v72 := GenS2{F0: 65535, F1: "a"}
        _ = v72
    case "":
        delete(v3, -2147483648)
    }
    var v73 uint = 255
    if (!false) {
        for fi74 := 0; fi74 < 16; fi74++ {
            var v75 uint16 = 0
            _ = fi74
            _ = v75
        }
        var v76 uint64 = genFn0()
        _ = v76
    } else {
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
        v77 := GenS2{F0: uint(v42), F1: "foo bar"}
        v11 -= (v11 * 9223372036854775807)
        _ = v77
    }
    _ = len(v3)
    if true {
        var v78 int8 = 127
        _ = v78
    }
    v2 -= v9
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v42:", v42)
    fmt.Println("v71:", v71)
    fmt.Println("v73:", v73)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v42
    _ = v43
    _ = v71
    _ = v73
    _ = v9
}
