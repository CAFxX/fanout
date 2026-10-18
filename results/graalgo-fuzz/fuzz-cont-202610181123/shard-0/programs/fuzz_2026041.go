package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 int }
type GenS2 struct { F0 int32; F1 int8; F2 uint16 }
type GenS3 struct { F0 uint8; F1 int16 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint16; F1 int16 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint32 }

func main() {
    v1 := []uint32{822}
    var v2 bool = false
    v3 := make(map[uint32]uint32)
    var v4 uint = 147
    if v2 {
        v5 := GenS1{F0: v2, F1: true, F2: -1}
        if true {
            if v2 {
                for fi6 := 0; fi6 < 15; fi6++ {
                    var v7 float64 = ((3.14159 / (0.0 - 1.5)) / -2.25)
                    v8 := GenS1{F0: v2, F1: v2, F2: fi6}
                    for fi9 := 0; fi9 < 12; fi9++ {
                        v4 = (v4 | 42)
                        if false {
                            var v10 int8 = 51
                            if v2 {
                                _ = len(v1)
                            } else {
                                var v11 uint16 = 543
                                v10 = (v10 ^ v10)
                                v12 := v3[2147483648]
                                _ = v11
                                _ = v12
                            }
                            var v13 int8 = (v10 * int8(fi6))
                            _ = v10
                            _ = v13
                        }
                        if v2 {
                            v14 := GenS2{F0: -1, F1: -128, F2: 0}
                            _ = v14
                        }
                        _ = fi9
                    }
                    _ = fi6
                    _ = v7
                    _ = v8
                }
                if false {
                    delete(v3, 2147483648)
                    for fi15 := 0; fi15 < 13; fi15++ {
                        fmt.Println(fi15)
                        _ = fi15
                    }
                } else {
                    var v16 float64 = 0.1
                    _ = v16
                }
            }
            var v17 bool = v2
            delete(v3, 658)
            _ = v17
        } else {
            var v18 string = "foo bar"
            fmt.Println(v4)
            _ = v18
        }
        _ = v5
    }
    for fi19 := 0; fi19 < 5; fi19++ {
        v2 = v2
        var v20 int16 = 0
        switch "" {
        case "a":
            fmt.Println(v20, v20)
        case "hello":
            if len(v1) > 0 {
                v1[0] = 4294967295
            }
        case "":
            for fi21 := 0; fi21 < 8; fi21++ {
                v22 := GenS3{F0: 237, F1: (v20 << ((v4) & 63))}
                var v23 uint32 = 4294967295
                var v24 uint64 = 2
                var v25 int64 = 776
                v26 := GenS1{F0: true, F1: ((!(!v2)) || v2), F2: fi19}
                _ = fi21
                _ = v22
                _ = v23
                _ = v24
                _ = v25
                _ = v26
            }
        default:
            v20 = 1
        }
        _ = fi19
        _ = v20
    }
    v4 -= (v4 ^ 281)
    var v27 int8 = 22
    v3[2] = 2147483647
    if (!(!(9223372036854775807 == 0))) {
        fmt.Println(v4, v2, v2)
        if ((32767 == 255) || v2) {
            fmt.Println(v2, v2, v4)
            v28 := GenS1{F0: true, F1: (32767 <= 1000), F2: 32768}
            var v29 float32 = 0.1
            _ = v28
            _ = v29
        } else {
            switch (("hello" + "\n\t") + ("" + ("\n\t" + ""))) {
            case "a":
                if false {
                    var v30 int16 = -1
                    for fi31 := 0; fi31 < 11; fi31++ {
                        v32 := GenS3{F0: 0, F1: 1}
                        v30 = -1
                        _ = fi31
                        _ = v32
                    }
                    _ = v30
                }
            case "hello":
                v2 = ((!false) || false)
            case "\n\t":
                for fi33 := 0; fi33 < 14; fi33++ {
                    if v2 {
                        fmt.Println(v2, v2)
                        fmt.Println(v2, fi33, v4)
                        switch fi33 {
                        case 32768:
                            var v34 uint64 = 42
                            _ = v34
                        case 2147483647:
                            v1 = append(v1, 2)
                        default:
                            if (3.14159 != 0.1) {
                                v35 := GenS2{F0: 42, F1: v27, F2: 834}
                                _ = v35
                            }
                        }
                    }
                    _ = fi33
                }
            default:
                switch 32768 {
                case 32768:
                    var v36 uint = v4
                    _ = v36
                case -616:
                    var v37 float64 = (float64(float32(1e10)) - (3.14159 + 0.0))
                    _ = v37
                default:
                    if (!(9223372036854775807 != -9223372036854775808)) {
                        for fi38 := 0; fi38 < 7; fi38++ {
                            v2 = v2
                            if v2 {
                                v2 = true
                                v39 := GenS4{F0: (!false)}
                                if false {
                                    var v40 uint64 = 4294967295
                                    v41 := GenS5{F0: 2, F1: 1}
                                    v42 := GenS6{F0: 255}
                                    _ = v40
                                    _ = v41
                                    _ = v42
                                }
                                _ = v39
                            } else {
                                var v43 uint = v4
                                _ = v43
                            }
                            v44 := GenS2{F0: 0, F1: v27, F2: 2}
                            _ = fi38
                            _ = v44
                        }
                    } else {
                        v2 = (-257 > 1)
                        v4 = v4
                    }
                }
            }
            v2 = (v4 == (^(v4)))
        }
    }
    v4 -= v4
    var v45 uint32 = 2147483647
    var v46 string = ("hello" + "\n\t")
    var v47 uint = 4294967295
    fmt.Println(v2, v47)
    switch v46 {
    case "foo bar":
        switch "a" {
        case "\n\t":
            v4 = (v47 + v47)
        case "a":
            v47 -= (v4 % (((v47 + uint(v45))) | 1))
        default:
            fmt.Println(v4, v47, v45)
        }
    }
    if (0.1 <= 3.14159) {
        if v2 {
            if v2 {
                v46 = ((v46 + (("a" + "foo bar") + v46)) + (v46 + v46))
                v48 := GenS3{F0: 180, F1: -51}
                _ = v48
            } else {
                v46 = "\n\t"
                var v49 uint = 192
                _ = v49
            }
        } else {
            v50 := GenS7{F0: 42}
            v51 := GenS4{F0: ((v2 && (32767 != 795)) && true)}
            v2 = (v2 || v2)
            _ = v50
            _ = v51
        }
        switch 9223372036854775807 {
        case 2147483647:
            v47 = v4
        case -317:
            var v52 uint32 = (v45 | (^(v45)))
            _ = v52
        case -1:
            for fi53 := 0; fi53 < 15; fi53++ {
                delete(v3, v45)
                if (v2 || ((v4 + v4) >= v4)) {
                    v45 = 0
                } else {
                    v45 -= v45
                    var v54 uint64 = 4294967295
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = v54
                }
                _ = fi53
            }
        }
        fmt.Println(v46, v45)
    }
    switch 1 {
    case -795:
        if (!(!(!(false || (2 < 0))))) {
            switch (v46 + "hello") {
            case "hello":
                var v55 int16 = 1000
                _ = v55
            }
        } else {
            for fi56 := 0; fi56 < 14; fi56++ {
                if true {
                    _ = len(v3)
                    v46 = (((("a" + "foo bar") + "foo bar") + v46) + "a")
                } else {
                    v57 := v3[65535]
                    v58 := GenS8{F0: v45}
                    var v59 float64 = (0.0 * 1.5)
                    _ = v57
                    _ = v58
                    _ = v59
                }
                _ = fi56
            }
            if v2 {
                v60 := GenS5{F0: 2, F1: 42}
                fmt.Println(v46, v27, v4)
                fmt.Println(v45, v2)
                _ = v60
            } else {
                if v2 {
                    if (0 >= -1) {
                        if ((((1.5 - float64(-2.25)) - (float64(0.1) + (1.5 * (((0.0 * (1e10 + 0.0)) / 0.1) / 1.5)))) >= (0.1 / 1.0)) || v2) {
                            v2 = false
                            v61 := GenS8{F0: (v45 * (v45 * v45))}
                            v47 = v4
                            _ = v61
                        }
                        switch v46 {
                        case "a":
                            v46 = v46
                        case "\n\t":
                            _ = len(v3)
                        }
                    }
                }
            }
        }
    case 0:
        _ = len(v1)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
}
