package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 uint32; F1 bool }
type GenS3 struct { F0 uint16; F1 uint16; F2 uint16 }

func main() {
    var v1 uint32 = 0
    var v2 uint = 453
    var v3 uint32 = v1
    var v4 float32 = float32((((3.14159 / 0.1) / 0.1) / 1e10))
    v5 := make(map[uint32]bool)
    var v6 float32 = 3.14159
    for fi7 := 0; fi7 < 4; fi7++ {
        for fi8 := 0; fi8 < 20; fi8++ {
            var v9 int64 = 9223372036854775807
            v1 = (v1 | 1)
            _ = fi8
            _ = v9
        }
        switch 127 {
        case -780:
            v6 += (v4 - v4)
        case 1:
            v10 := GenS1{F0: 374, F1: 2147483647}
            _ = v10
        default:
            var v11 int16 = 0
            _ = v11
        }
        for fi12 := 0; fi12 < 20; fi12++ {
            var v13 int64 = 853
            _ = fi12
            _ = v13
        }
        _ = fi7
    }
    v14 := []int{128}
    switch "foo bar" {
    case "foo bar":
        for fi15 := 0; fi15 < 13; fi15++ {
            if false {
                var v16 int16 = 1000
                for fi17 := 0; fi17 < 19; fi17++ {
                    v6 += (0.1 - (((((v4 / ((((1.5 + v6) / v6) - (v4 / float32(-2.25))) / v6)) + ((((float32((1e10 / 3.14159)) * (v6 / v6)) * (v4 - v6)) / (3.14159 + ((float32(-2.25) / (v6 * v4)) + v6))) * v4)) / v6) * v4) - 1.5))
                    var v18 uint32 = (v3 % ((v1) | 1))
                    if true {
                        var v19 int64 = 1
                        fmt.Println(fi17)
                        _ = v19
                    } else {
                        if true {
                            v18 = 2147483648
                        } else {
                            fmt.Println(v16, v3)
                            var v20 int8 = 126
                            _ = v20
                        }
                        v1 -= v1
                    }
                    _ = fi17
                    _ = v18
                }
                var v21 uint = v2
                _ = v16
                _ = v21
            }
            _ = fi15
        }
    case "a":
        for fi22 := 0; fi22 < 6; fi22++ {
            for fi23 := 0; fi23 < 7; fi23++ {
                v6 = float32(1e10)
                switch fi23 {
                case 32768:
                    v24 := GenS1{F0: 42, F1: 2147483647}
                    _ = v24
                default:
                    v5[(v1 | v3)] = false
                }
                v25 := GenS1{F0: 15, F1: -995}
                _ = fi23
                _ = v25
            }
            fmt.Println(v6)
            v3 -= (v1 / (((v1 % ((v1) | 1))) | 1))
            _ = fi22
        }
    case "hello":
        fmt.Println(v1)
    }
    _ = len(v14)
    _ = len(v14)
    v4 += 1e10
    var v26 string = "hello"
    for fi27 := 0; fi27 < 18; fi27++ {
        v3 -= v3
        if (2 <= 1) {
            v28 := GenS2{F0: (v1 % ((v3) | 1)), F1: true}
            _ = v28
        } else {
            v2 = 255
            switch (fi27 | (fi27 ^ fi27)) {
            case 87:
                v29 := GenS2{F0: v3, F1: (uint64(v2) < 196)}
                _ = v29
            case 32768:
                v5[1] = (92 < -2147483648)
            case 121:
                for fi30 := 0; fi30 < 4; fi30++ {
                    v1 = (^(v1))
                    _ = fi30
                }
            }
        }
        _ = fi27
    }
    for fi31 := 0; fi31 < 7; fi31++ {
        var v32 int16 = 1
        var v33 int32 = 1
        if false {
            for fi34 := 0; fi34 < 1; fi34++ {
                v6 += (0.1 - v6)
                var v35 int8 = 1
                if false {
                    var v36 uint32 = v1
                    _ = v36
                } else {
                    v37 := GenS3{F0: 813, F1: 255, F2: 1}
                    if true {
                        v32 = v32
                    }
                    if ((v1 << ((v2) & 63)) > (v1 >> (((^(v2))) & 63))) {
                        var v38 uint8 = 127
                        _ = v38
                    } else {
                        var v39 int64 = -9223372036854775808
                        v40 := GenS3{F0: 720, F1: 672, F2: 1}
                        fmt.Println(v32, v4, v39)
                        v26 = v26
                        var v41 uint16 = 701
                        var v42 int8 = v35
                        _ = v39
                        _ = v40
                        _ = v41
                        _ = v42
                    }
                    _ = v37
                }
                v26 = v26
                fmt.Println(v1, fi31, v33)
                _ = fi34
                _ = v35
            }
            var v43 int16 = (v32 << (((^(v2))) & 63))
            var v44 uint = v2
            _ = v43
            _ = v44
        }
        _ = fi31
        _ = v32
        _ = v33
    }
    var v45 uint32 = (v1 + 42)
    switch 2147483647 {
    case -94:
        if len(v14) > 0 {
            v14[0] = 0
        }
    case 943:
        var v46 uint8 = 128
        _ = v46
    default:
        fmt.Println(v26)
    }
    _ = len(v5)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v6
}
