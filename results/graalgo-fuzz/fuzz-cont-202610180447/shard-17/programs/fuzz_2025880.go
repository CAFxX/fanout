package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 bool; F1 uint16 }
type GenS5 struct { F0 int16; F1 int16 }

func main() {
    var v1 uint16 = 1
    v2 := []int{1, 32768, 4294967295}
    var v3 int64 = 1
    var v4 float32 = (0.1 - 3.14159)
    var v5 int64 = -374
    var v6 uint64 = 207
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 int32 = 2147483647
        v9 := GenS1{F0: 0, F1: (true && true)}
        var v10 uint16 = 854
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    if false {
        v4 -= ((0.1 + v4) - (1.5 + float32(-2.25)))
        v2 = append(v2, 42)
    } else {
        if false {
            fmt.Println(v4)
            for fi11 := 0; fi11 < 11; fi11++ {
                if false {
                    var v12 uint = 255
                    var v13 uint64 = v6
                    var v14 int64 = (v5 >> (((v12 * (v12 % ((uint(v1)) | 1)))) & 63))
                    _ = v12
                    _ = v13
                    _ = v14
                }
                var v15 float64 = float64(v4)
                _ = fi11
                _ = v15
            }
        }
        var v16 float32 = 0.1
        _ = v16
    }
    fmt.Println(v5)
    v5 = -1
    for fi17 := 0; fi17 < 3; fi17++ {
        if (127 < 42) {
            v18 := GenS2{F0: 363, F1: (^(v1))}
            _ = v18
        }
        v5 = (v5 - v5)
        v6 = 2
        _ = fi17
    }
    v3 = (v3 - 9223372036854775807)
    if false {
        for fi19 := 0; fi19 < 13; fi19++ {
            for fi20 := 0; fi20 < 20; fi20++ {
                v1 = 1
                _ = fi20
            }
            _ = len(v2)
            _ = fi19
        }
        if true {
            var v21 int = -2147483648
            v5 = int64(v6)
            _ = v21
        }
    } else {
        fmt.Println(v4, v6, v6)
        v22 := GenS2{F0: 65535, F1: v1}
        var v23 int64 = v5
        _ = v22
        _ = v23
    }
    v2 = append(v2, 2147483647)
    fmt.Println(v6, v6)
    if false {
        v24 := GenS3{F0: 1}
        var v25 int8 = -115
        _ = v24
        _ = v25
    } else {
        for fi26 := 0; fi26 < 7; fi26++ {
            if true {
                v27 := GenS1{F0: 0, F1: (false || (v4 >= (float32((3.14159 - ((0.1 / float64(v4)) - (float64(v4) + 1e10)))) / 1e10)))}
                if (1 == 32767) {
                    switch fi26 {
                    case 785:
                        v1 -= 255
                    case 306:
                        v4 += v4
                    }
                    var v28 uint = 440
                    var v29 uint32 = 2
                    _ = v28
                    _ = v29
                } else {
                    fmt.Println(v1, v6, v4)
                    v5 = (v5 ^ v5)
                    v1 = (v1 >> ((65535) & 63))
                }
                _ = v27
            }
            _ = fi26
        }
    }
    v2 = append(v2, -2147483648)
    for fi30 := 0; fi30 < 2; fi30++ {
        if (!false) {
            if (-2147483648 <= -440) {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                var v31 uint32 = 4294967295
                _ = v31
            } else {
                v4 = (v4 / 1.5)
                for fi32 := 0; fi32 < 6; fi32++ {
                    switch "a" {
                    case "hello":
                        if true {
                            var v33 bool = ((v4 == (0.0 - (v4 - v4))) && (!false))
                            var v34 int16 = -600
                            fmt.Println(v34, v33)
                            _ = v33
                            _ = v34
                        }
                    case "foo bar":
                        var v35 int = fi30
                        _ = v35
                    case "\n\t":
                        v36 := GenS3{F0: 2147483647}
                        _ = v36
                    default:
                        v37 := GenS4{F0: (v6 != (v6 >> ((255) & 63))), F1: v1}
                        _ = v37
                    }
                    switch fi30 {
                    case 2147483647:
                        v1 = 255
                        v5 += (v3 | -9223372036854775808)
                    case 786:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    case 9223372036854775807:
                        fmt.Println(fi32)
                    }
                    fmt.Println(fi30, v5, v1)
                    _ = fi32
                }
            }
            var v38 float32 = float32(3.14159)
            _ = v38
        }
        _ = fi30
    }
    v39 := GenS5{F0: -1, F1: 522}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v6 = v6
    fmt.Println(v6, v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
