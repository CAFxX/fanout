package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 bool; F1 uint64; F2 uint32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int8; F1 string }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint64; F1 int32 }

func main() {
    var v1 float64 = (float64(1.5) / ((-2.25 * 1.5) / 1.0))
    v2 := []int{368, 474, 431}
    v3 := []int{32768, 4294967295, -395, 947}
    var v4 uint = 0
    switch "a" {
    case "hello":
        fmt.Println(v4)
    case "foo bar":
        fmt.Println(v1, v1)
    case "a":
        var v5 uint32 = 4294967295
        _ = v5
    default:
        fmt.Println(v4)
    }
    v6 := GenS1{F0: -9223372036854775808, F1: 1}
    var v7 uint32 = 2
    if (!(false || false)) {
        v8 := GenS1{F0: 9223372036854775807, F1: 739}
        v7 = v7
        if (v7 > 2) {
            for fi9 := 0; fi9 < 7; fi9++ {
                fmt.Println(v4, v4, v7)
                var v10 int32 = -359
                v11 := GenS2{F0: false, F1: 18446744073709551615, F2: (v7 & v7)}
                _ = fi9
                _ = v10
                _ = v11
            }
        } else {
            v12 := GenS2{F0: (1.5 >= (3.14159 / (((0.1 - 0.0) / -2.25) - 3.14159))), F1: 591, F2: (v7 | v7)}
            v13 := GenS3{F0: 1}
            _ = len(v2)
            _ = v12
            _ = v13
        }
        _ = v8
    } else {
        v2 = append(v2, 4294967295)
    }
    switch 814 {
    case 1:
        v7 = 801
    case 32768:
        var v14 bool = false
        _ = v14
    }
    if ((false && ((v7 / (((v7 * 4294967295)) | 1)) < (v7 << ((138) & 63)))) && (v4 != 92)) {
        if false {
            if false {
                if false {
                    v1 = 3.14159
                } else {
                    for fi15 := 0; fi15 < 20; fi15++ {
                        fmt.Println(v7)
                        if false {
                            switch fi15 {
                            case -555:
                                _ = len(v2)
                            case 4294967295:
                                v16 := GenS4{F0: -1}
                                _ = v16
                            default:
                                v2 = append(v2, fi15)
                            }
                            var v17 float32 = -2.25
                            v18 := GenS2{F0: false, F1: uint64(v4), F2: (v7 + (v7 << (((v4 / ((v4) | 1))) & 63)))}
                            _ = v17
                            _ = v18
                        }
                        _ = fi15
                    }
                    v7 = (v7 ^ (^(v7)))
                    if ((false || (9223372036854775807 != 1)) && ("" >= (("hello" + "foo bar") + "\n\t"))) {
                        v2 = append(v2, -1)
                        fmt.Println(v4)
                    }
                }
                v19 := GenS4{F0: -9223372036854775808}
                if len(v2) > 0 {
                    v2[0] = 32767
                }
                _ = v19
            }
        } else {
            switch 127 {
            case 9223372036854775807:
                for fi20 := 0; fi20 < 6; fi20++ {
                    var v21 float64 = v1
                    _ = fi20
                    _ = v21
                }
            case 42:
                v7 -= 2147483647
            }
            if true {
                for fi22 := 0; fi22 < 5; fi22++ {
                    var v23 int = (fi22 << ((v4) & 63))
                    switch "" {
                    case "a":
                        var v24 uint64 = 2
                        _ = v24
                    case "hello":
                        if ((false && ((v1 - (-2.25 + v1)) == v1)) && (1 > 1)) {
                            var v25 int32 = 0
                            if true {
                                var v26 float64 = (v1 * (v1 + v1))
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                _ = v26
                            } else {
                                for fi27 := 0; fi27 < 19; fi27++ {
                                    v4 = v4
                                    v23 = fi27
                                    _ = fi27
                                }
                                v28 := GenS1{F0: 9223372036854775807, F1: 496}
                                var v29 float64 = float64(((0.1 * (((3.14159 * 1.5) - 0.1) / 3.14159)) + 0.0))
                                _ = v28
                                _ = v29
                            }
                            _ = v25
                        } else {
                            v30 := GenS5{F0: -128, F1: ("hello" + "hello")}
                            _ = v30
                        }
                    }
                    if true {
                        v3 = append(v3, fi22)
                        v31 := GenS1{F0: 1, F1: 0}
                        for fi32 := 0; fi32 < 10; fi32++ {
                            v33 := GenS6{F0: -1}
                            v34 := GenS7{F0: 1, F1: -1}
                            _ = fi32
                            _ = v33
                            _ = v34
                        }
                        _ = v31
                    } else {
                        switch (v23 % (((v23 % ((fi22) | 1))) | 1)) {
                        case 32767:
                            var v35 bool = false
                            _ = v35
                        }
                    }
                    _ = fi22
                    _ = v23
                }
                var v36 uint = (^(v4))
                fmt.Println(v7, v36, v7)
                _ = v36
            }
        }
        v37 := GenS3{F0: uint16(v7)}
        _ = v37
    }
    var v38 uint = uint(v7)
    fmt.Println(v7, v4, v4)
    for fi39 := 0; fi39 < 18; fi39++ {
        if (0 >= 0) {
            fmt.Println(fi39)
        }
        v38 += v38
        v40 := GenS6{F0: 42}
        _ = fi39
        _ = v40
    }
    v4 += (v4 + (v4 + uint(v7)))
    for fi41 := 0; fi41 < 14; fi41++ {
        v42 := GenS7{F0: 68, F1: -2147483648}
        _ = fi41
        _ = v42
    }
    var v43 uint32 = (v7 >> ((v38) & 63))
    fmt.Println(v43, v4)
    fmt.Println(v1, v38, v7)
    var v44 string = "a"
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v43
    _ = v44
    _ = v6
    _ = v7
}
