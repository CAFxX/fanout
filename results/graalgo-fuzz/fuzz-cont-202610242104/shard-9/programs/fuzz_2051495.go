package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 bool }
type GenS2 struct { F0 int8; F1 uint64; F2 uint32 }
type GenS3 struct { F0 uint; F1 uint16; F2 int8 }

func main() {
    var v1 uint16 = 269
    var v2 float64 = (1.5 / 1.0)
    v3 := make(map[string]bool)
    v2 = v2
    _ = len(v3)
    switch 1 {
    case -1:
        var v4 bool = true
        _ = v4
    default:
        v1 += (v1 / (((v1 ^ 255)) | 1))
    }
    if (-1 > -128) {
        v1 -= 299
        for fi5 := 0; fi5 < 14; fi5++ {
            var v6 uint64 = 644
            v7 := []int{fi5}
            var v8 bool = true
            _ = fi5
            _ = v6
            _ = v7
            _ = v8
        }
    }
    v9 := GenS1{F0: 42, F1: -1, F2: (-2147483648 > 9223372036854775807)}
    fmt.Println(v1, v1)
    var v10 uint64 = 1
    switch 0 {
    case 0:
        fmt.Println(v2, v2, v1)
    default:
        v10 += v10
    }
    v11 := []int{42, 426, -980}
    if (!(-2147483648 > 0)) {
        for fi12 := 0; fi12 < 16; fi12++ {
            switch "foo bar" {
            case "":
                fmt.Println(v2)
            case "\n\t":
                var v13 int16 = 12
                _ = v13
            case "a":
                v14 := GenS1{F0: 42, F1: -2147483648, F2: false}
                _ = v14
            default:
                v10 -= 2
            }
            _ = fi12
        }
        for fi15 := 0; fi15 < 18; fi15++ {
            v16 := GenS1{F0: -34, F1: 0, F2: (-934 > -587)}
            if true {
                var v17 uint16 = (v1 + uint16(v10))
                for fi18 := 0; fi18 < 6; fi18++ {
                    v10 += (v10 >> ((18446744073709551615) & 63))
                    var v19 float64 = v2
                    _ = fi18
                    _ = v19
                }
                if (-1 < 9223372036854775807) {
                    fmt.Println(v17, v1)
                    switch fi15 {
                    case -174:
                        switch (-(fi15)) {
                        case 127:
                            v20 := GenS2{F0: int8(v10), F1: (v10 | (v10 << ((224) & 63))), F2: 42}
                            _ = v20
                        case 1:
                            v10 = (v10 & 2)
                        case -9223372036854775808:
                            v10 = (v10 % ((0) | 1))
                        }
                    default:
                        v2 -= (((0.0 - v2) / float64(1.5)) * v2)
                    }
                }
                _ = v17
            } else {
                v1 = (v1 >> ((65535) & 63))
            }
            v21 := v3[""]
            _ = fi15
            _ = v16
            _ = v21
        }
    }
    var v22 int8 = -128
    for fi23 := 0; fi23 < 18; fi23++ {
        var v24 bool = (!(true && (!false)))
        if len(v11) > 0 {
            v11[0] = fi23
        }
        _ = fi23
        _ = v24
    }
    delete(v3, ("" + "hello"))
    var v25 uint64 = (v10 & v10)
    v1 = 538
    if false {
        if (float64((float32((v2 - (v2 + v2))) / (float32((1e10 / ((-2.25 / v2) * float64(0.1)))) / float32(v2)))) == float64(-2.25)) {
            var v26 bool = true
            if false {
                if v26 {
                    v27 := GenS1{F0: (v22 + int8(v1)), F1: 0, F2: v26}
                    _ = v27
                }
                _ = len(v3)
                switch 2147483647 {
                case -94:
                    var v28 int32 = -2147483648
                    _ = v28
                case 0:
                    v2 -= v2
                case 2147483647:
                    if false {
                        v29 := v3[""]
                        v30 := GenS2{F0: (v22 * v22), F1: (v10 * 1), F2: 24}
                        fmt.Println(v29)
                        _ = v29
                        _ = v30
                    } else {
                        v10 = v10
                        if (!false) {
                            var v31 int = 127
                            for fi32 := 0; fi32 < 16; fi32++ {
                                var v33 int32 = -2147483648
                                _ = fi32
                                _ = v33
                            }
                            _ = v31
                        } else {
                            if len(v11) >= 2 {
                                v11 = v11[1:len(v11)]
                            }
                        }
                        for fi34 := 0; fi34 < 13; fi34++ {
                            v22 = (v22 + v22)
                            _ = fi34
                        }
                    }
                }
            } else {
                v1 = (^(v1))
            }
            switch "foo bar" {
            case "hello":
                var v35 bool = false
                _ = v35
            case "a":
                v3[("foo bar" + "a")] = v26
            case "foo bar":
                v26 = true
            }
            _ = v26
        }
        v22 = v22
    } else {
        var v36 string = "hello"
        v37 := GenS3{F0: 2, F1: 32767, F2: v22}
        _ = v36
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v22
    _ = v25
    _ = v3
    _ = v9
}
