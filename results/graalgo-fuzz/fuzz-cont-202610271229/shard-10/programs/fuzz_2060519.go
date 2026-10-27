package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int }
type GenS5 struct { F0 string; F1 string }
type GenS6 struct { F0 uint16; F1 uint16 }
type GenS7 struct { F0 string; F1 int64 }
type GenS8 struct { F0 int16 }

func main() {
    v1 := []uint32{1}
    var v2 uint32 = 245
    v3 := make(map[uint32]int)
    var v4 uint8 = 0
    v5 := make(map[uint32]uint32)
    fmt.Println(v4, v4)
    switch "hello" {
    case "":
        switch 557 {
        case 9223372036854775807:
            v6 := v3[v2]
            _ = v6
        }
    case "\n\t":
        delete(v3, (v2 % ((1) | 1)))
    default:
        switch "foo bar" {
        case "":
            if (false || ((65535 <= uint(v4)) || false)) {
                var v7 int16 = 42
                fmt.Println(v4)
                fmt.Println(v2)
                _ = v7
            } else {
                v8 := GenS1{F0: v4, F1: v4, F2: "hello"}
                _ = v8
            }
        case "\n\t":
            var v9 uint8 = v4
            _ = v9
        }
    }
    v10 := GenS1{F0: v4, F1: (v4 % ((1) | 1)), F2: "\n\t"}
    switch "" {
    case "\n\t":
        if (42 < 0) {
            fmt.Println(v2, v4)
            if (!false) {
                for fi11 := 0; fi11 < 14; fi11++ {
                    v12 := GenS2{F0: v4}
                    var v13 uint8 = (^(v4))
                    for fi14 := 0; fi14 < 14; fi14++ {
                        v2 = v2
                        if (v2 == uint32(v13)) {
                            v1 = append(v1, v2)
                            v13 = v13
                        }
                        var v15 int = 1
                        var v16 uint = 697
                        if ((float64(1.5) - 1.5) >= 1.5) {
                            var v17 uint = v16
                            v17 = (v17 + 49)
                            if (v16 >= v16) {
                                v18 := GenS1{F0: (v4 >> (((v17 + v17)) & 63)), F1: (v4 | v4), F2: "a"}
                                fmt.Println(fi14, v17, fi14)
                                _ = v18
                            }
                            var v19 int8 = -126
                            var v20 float64 = 0.1
                            _ = v17
                            _ = v19
                            _ = v20
                        }
                        v16 = 114
                        if (!true) {
                            fmt.Println(v2, fi11)
                        }
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                    _ = fi11
                    _ = v12
                    _ = v13
                }
            }
        }
    case "":
        v21 := GenS2{F0: v4}
        _ = v21
    case "foo bar":
        if true {
            if false {
                var v22 string = ("a" + ("hello" + ("a" + "a")))
                var v23 bool = true
                v23 = (-2.25 != 1e10)
                _ = v22
                _ = v23
            } else {
                var v24 int64 = 1
                if true {
                    v25 := GenS3{F0: 1}
                    v26 := GenS4{F0: int(v4)}
                    _ = v25
                    _ = v26
                }
                if false {
                    if false {
                        for fi27 := 0; fi27 < 13; fi27++ {
                            if false {
                                var v28 uint32 = v2
                                if true {
                                    v2 = (v28 * (v28 & v28))
                                    if (1 <= -414) {
                                        v29 := GenS5{F0: "hello", F1: "foo bar"}
                                        fmt.Println(fi27)
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        _ = v29
                                    }
                                    var v30 uint64 = 139
                                    _ = v30
                                } else {
                                    v31 := GenS1{F0: v4, F1: (^(v4)), F2: ("hello" + "foo bar")}
                                    _ = v31
                                }
                                for fi32 := 0; fi32 < 2; fi32++ {
                                    fmt.Println(v24, v4, v2)
                                    _ = fi32
                                }
                                _ = v28
                            }
                            _ = fi27
                        }
                        _ = len(v5)
                        var v33 string = ("\n\t" + "foo bar")
                        _ = v33
                    } else {
                        if len(v1) > 0 {
                            v1[0] = (v2 * (^(v2)))
                        }
                        var v34 int8 = -81
                        v2 = v2
                        _ = v34
                    }
                } else {
                    if false {
                        v35 := GenS1{F0: 42, F1: v4, F2: "a"}
                        var v36 int16 = 0
                        _ = v35
                        _ = v36
                    }
                }
                _ = v24
            }
            switch 9223372036854775807 {
            case 42:
                var v37 bool = (-474 == -1)
                _ = v37
            default:
                var v38 uint16 = 42
                _ = v38
            }
            var v39 int32 = 698
            _ = v39
        }
    default:
        if ((255 != 725) || false) {
            v1 = append(v1, (v2 | (v2 * (v2 << ((4294967295) & 63)))))
            _ = len(v1)
        } else {
            v40 := v3[(^(v2))]
            v40 += v40
            v4 += v4
            _ = v40
        }
    }
    var v41 uint8 = 2
    var v42 string = ""
    var v43 uint16 = 255
    v44 := GenS6{F0: 599, F1: uint16(v41)}
    v4 = v41
    for fi45 := 0; fi45 < 20; fi45++ {
        if false {
            for fi46 := 0; fi46 < 14; fi46++ {
                var v47 uint16 = uint16(v4)
                v2 = 2147483648
                v48 := GenS7{F0: "", F1: 1}
                _ = fi46
                _ = v47
                _ = v48
            }
            var v49 uint64 = 2
            _ = v49
        }
        v42 = (v42 + ((("foo bar" + (v42 + "\n\t")) + v42) + "foo bar"))
        var v50 int64 = -201
        _ = fi45
        _ = v50
    }
    var v51 int32 = -2147483648
    v52 := GenS8{F0: -32768}
    var v53 float32 = float32((3.14159 - (((0.0 * (1.5 + (((1e10 - (1e10 / ((0.1 + 3.14159) / (3.14159 + float64((0.1 * (((-2.25 * (3.14159 / 3.14159)) * 3.14159) + 0.1))))))) * ((float64(3.14159) / ((3.14159 / (-2.25 * -2.25)) + 0.1)) - 1.5)) * 0.0))) / 1.0) / 1.0)))
    _ = len(v5)
    var v54 uint32 = uint32(v41)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v51:", v51)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v5
    _ = v51
    _ = v52
    _ = v53
    _ = v54
}
