package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 int64; F1 uint16; F2 uint64 }
type GenS3 struct { F0 uint8; F1 bool }
type GenS4 struct { F0 bool; F1 int16 }
type GenS5 struct { F0 uint; F1 uint8; F2 uint8 }
func genFn0(p0 int64, p1 int16) bool {
    var v4 int = 128
    v5 := []int{}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return (!(!false))
}

func main() {
    var v1 string = "foo bar"
    v2 := make(map[uint32]int)
    v3 := []uint32{4294967295}
    v1 = v1
    var v6 int = 309
    v6 = (^(v6))
    v2[2147483648] = (v6 / ((v6) | 1))
    for fi7 := 0; fi7 < 17; fi7++ {
        if len(v3) > 0 {
            v3[0] = 2147483647
        }
        v6 -= (v6 << ((770) & 63))
        _ = fi7
    }
    var v8 uint16 = 0
    if (v6 != v6) {
        if (1000 != 1) {
            v8 += v8
            v9 := GenS1{F0: 1, F1: (4294967295 <= 145)}
            v10 := GenS1{F0: 42, F1: (uint64(v8) < 319)}
            _ = v10
            _ = v9
        }
        v1 = v1
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v11 uint = 4294967295
        var v12 int16 = -1
        _ = v11
        _ = v12
    }
    var v13 int = (v6 - v6)
    var v14 uint8 = 1
    if (-128 != 14) {
        var v15 uint16 = 466
        v15 = (v8 - uint16(v14))
        if false {
            var v16 bool = (2147483647 >= -42)
            v14 = (v14 & (v14 << ((880) & 63)))
            switch (v1 + ((v1 + ("" + v1)) + v1)) {
            case "hello":
                var v17 string = "hello"
                _ = v17
            default:
                switch (v6 * -9223372036854775808) {
                case 4294967295:
                    var v18 uint8 = (v14 << ((255) & 63))
                    _ = v18
                case 1:
                    v19 := GenS2{F0: 1, F1: v8, F2: 42}
                    _ = v19
                }
            }
            _ = v16
        } else {
            v20 := GenS3{F0: v14, F1: true}
            _ = v20
        }
        _ = v15
    } else {
        v14 = v14
    }
    v21 := GenS4{F0: (65535 >= v8), F1: 32767}
    if true {
        for fi22 := 0; fi22 < 8; fi22++ {
            switch fi22 {
            case 9223372036854775807:
                if ((42 >= -9223372036854775808) || (223 < uint32(v14))) {
                    var v23 float32 = 3.14159
                    if false {
                        var v24 int16 = -695
                        _ = v24
                    }
                    if (!(-1 > 2147483647)) {
                        fmt.Println(v1, v14)
                    }
                    _ = v23
                }
            case -991:
                if false {
                    v25 := GenS3{F0: v14, F1: true}
                    v2[2] = v13
                    switch (v13 % (((fi22 ^ fi22)) | 1)) {
                    case 4294967295:
                        for fi26 := 0; fi26 < 6; fi26++ {
                            var v27 uint16 = v8
                            switch "hello" {
                            case "a":
                                v28 := v2[209]
                                _ = v28
                            default:
                                v29 := GenS5{F0: 523, F1: (v14 - v14), F2: (^(v14))}
                                _ = v29
                            }
                            _ = fi26
                            _ = v27
                        }
                    case -407:
                        var v30 uint = 42
                        _ = v30
                    }
                    _ = v25
                }
            }
            _ = fi22
        }
        v6 += (^(v13))
    } else {
        delete(v2, 2)
    }
    for fi31 := 0; fi31 < 17; fi31++ {
        v6 += 97
        _ = fi31
    }
    if (true || (!true)) {
        if (2 >= v8) {
            switch "" {
            case "\n\t":
                v13 += int(v14)
            case "":
                var v32 int8 = 4
                _ = v32
            case "foo bar":
                v1 = ((v1 + "a") + (v1 + v1))
            }
            v3 = append(v3, 4294967295)
        }
    }
    if false {
        fmt.Println(v1, v8)
        var v33 float32 = (0.1 * 1e10)
        _ = v33
    } else {
        fmt.Println(v14, v14)
        switch (v6 % (((v6 - -9223372036854775808)) | 1)) {
        case 4294967295:
            for fi34 := 0; fi34 < 7; fi34++ {
                v13 = v6
                v35 := v2[821]
                _ = fi34
                _ = v35
            }
        case 9223372036854775807:
            fmt.Println(v1, v14)
        case 0:
            v36 := GenS3{F0: 128, F1: false}
            _ = v36
        default:
            var v37 bool = ((!false) || true)
            _ = v37
        }
        switch "a" {
        case "a":
            switch v6 {
            case 4294967295:
                var v38 int8 = 25
                _ = v38
            default:
                var v39 int32 = -1
                _ = v39
            }
        default:
            var v40 int8 = 42
            _ = v40
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v21
    _ = v3
    _ = v6
    _ = v8
}
