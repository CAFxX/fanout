package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 int16; F1 uint64 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int16; F1 int32 }
type GenS6 struct { F0 int16; F1 string }

func main() {
    v1 := make(map[int]bool)
    v2 := []int{-844}
    var v3 uint8 = 103
    v4 := make(map[uint32]uint32)
    var v5 bool = false
    switch "foo bar" {
    case "\n\t":
        v6 := v4[42]
        _ = v6
    case "a":
        v7 := GenS1{F0: "\n\t"}
        _ = v7
    case "hello":
        var v8 uint64 = 42
        _ = v8
    }
    if v5 {
        var v9 uint16 = 2
        _ = v9
    } else {
        var v10 uint16 = 2
        _ = v10
    }
    if v5 {
        var v11 float32 = (3.14159 / (1.5 - float32(-2.25)))
        if v5 {
            fmt.Println(v3)
        } else {
            v11 = v11
        }
        _ = v11
    }
    v12 := GenS2{F0: 1, F1: 127}
    switch -1 {
    case 282:
        var v13 float64 = -2.25
        _ = v13
    case 4294967295:
        fmt.Println(v5, v3)
    case 42:
        for fi14 := 0; fi14 < 15; fi14++ {
            v15 := GenS3{F0: 32767, F1: 42}
            v5 = true
            _ = fi14
            _ = v15
        }
    default:
        v16 := GenS4{F0: 0}
        _ = v16
    }
    if true {
        v2 = append(v2, 128)
        var v17 int64 = 0
        _ = v17
    } else {
        v3 += (v3 ^ (v3 << ((42) & 63)))
    }
    var v18 uint = 255
    var v19 uint64 = 668
    for fi20 := 0; fi20 < 5; fi20++ {
        var v21 uint32 = 65535
        _ = fi20
        _ = v21
    }
    v3 -= (v3 * (^(v3)))
    fmt.Println(v19)
    switch "foo bar" {
    case "\n\t":
        switch -211 {
        case -9223372036854775808:
            if v5 {
                switch 127 {
                case 127:
                    for fi22 := 0; fi22 < 1; fi22++ {
                        if (0.0 > 0.1) {
                            for fi23 := 0; fi23 < 4; fi23++ {
                                var v24 uint64 = 97
                                var v25 int8 = 1
                                _ = fi23
                                _ = v24
                                _ = v25
                            }
                        }
                        var v26 int32 = 120
                        _ = fi22
                        _ = v26
                    }
                case 0:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                case 128:
                    for fi27 := 0; fi27 < 13; fi27++ {
                        fmt.Println(fi27, fi27, v3)
                        var v28 bool = (!(209 != v3))
                        _ = fi27
                        _ = v28
                    }
                }
                for fi29 := 0; fi29 < 9; fi29++ {
                    for fi30 := 0; fi30 < 17; fi30++ {
                        v31 := GenS5{F0: -476, F1: 1}
                        v19 = (v19 * (v19 % (((v19 & (v19 / ((v19) | 1)))) | 1)))
                        switch 0 {
                        case 127:
                            v32 := GenS2{F0: (v19 << ((v18) & 63)), F1: (v3 & (v3 / (((v3 ^ v3)) | 1)))}
                            _ = v32
                        case -413:
                            v4[uint32(v19)] = 813
                        default:
                            v33 := GenS6{F0: 244, F1: "a"}
                            _ = v33
                        }
                        v18 = v18
                        if v5 {
                            var v34 uint = v18
                            v35 := GenS1{F0: ("a" + "")}
                            var v36 int32 = 32
                            _ = v34
                            _ = v35
                            _ = v36
                        }
                        _ = fi30
                        _ = v31
                    }
                    _ = fi29
                }
            }
        case 962:
            v37 := GenS4{F0: -1}
            _ = v37
        }
    default:
        v19 += v19
    }
    var v38 uint8 = 127
    var v39 uint32 = 758
    v40 := GenS5{F0: 1000, F1: int32(v3)}
    for fi41 := 0; fi41 < 4; fi41++ {
        if (!(true && (!(!(0.1 < 0.1))))) {
            for fi42 := 0; fi42 < 19; fi42++ {
                var v43 string = "hello"
                _ = fi42
                _ = v43
            }
            if v5 {
                fmt.Println(v3, v19)
                switch "a" {
                case "foo bar":
                    v18 += v18
                case "a":
                    v38 += v38
                }
            } else {
                v5 = false
            }
        }
        v18 -= 255
        _ = fi41
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
}
