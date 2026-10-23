package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 uint64; F1 uint8; F2 uint16 }
type GenS3 struct { F0 int32 }

func main() {
    var v1 uint8 = 2
    var v2 uint64 = 823
    var v3 int8 = -1
    var v4 float32 = 0.0
    v5 := make(map[uint32]uint32)
    fmt.Println(v2)
    var v6 int = -922
    if (((3.14159 / 0.1) - (v4 * v4)) > 1e10) {
        v3 = v3
        v1 -= (v1 + 42)
    }
    v3 = (v3 * v3)
    var v7 int8 = (v3 >> ((4294967295) & 63))
    var v8 uint = 858
    for fi9 := 0; fi9 < 7; fi9++ {
        v7 = v7
        for fi10 := 0; fi10 < 8; fi10++ {
            switch v6 {
            case 128:
                v2 = v2
            }
            if (!true) {
                switch fi9 {
                case 556:
                    var v11 int8 = v7
                    _ = v11
                case -1:
                    if (-1 != 396) {
                        fmt.Println(v6)
                    }
                case -2147483648:
                    v12 := GenS1{F0: 42, F1: v1}
                    var v13 int64 = 664
                    _ = v12
                    _ = v13
                }
                v3 -= v7
                if false {
                    var v14 int64 = 9223372036854775807
                    v7 = (v7 + (^(v3)))
                    _ = v14
                } else {
                    v1 = 2
                    var v15 uint8 = 42
                    var v16 uint8 = uint8(v2)
                    _ = v15
                    _ = v16
                }
            }
            _ = fi10
        }
        _ = fi9
    }
    var v17 int = -9223372036854775808
    v18 := []int{4294967295, (v6 + -1), 42}
    _ = len(v18)
    v17 += 9223372036854775807
    if ((2147483647 > 128) && (!false)) {
        if true {
            if (v1 == 180) {
                if (!(317 == 1)) {
                    if false {
                        var v19 int16 = 693
                        var v20 uint8 = v1
                        fmt.Println(v6)
                        _ = v19
                        _ = v20
                    } else {
                        v18 = append(v18, 9223372036854775807)
                    }
                }
                for fi21 := 0; fi21 < 1; fi21++ {
                    v3 -= 42
                    v22 := GenS2{F0: (v2 ^ (^(v2))), F1: (v1 ^ (v1 - (v1 << (((v8 >> ((42) & 63))) & 63)))), F2: uint16(v8)}
                    v3 = (v7 ^ -80)
                    _ = fi21
                    _ = v22
                }
                v23 := v5[0]
                _ = v23
            }
        } else {
            var v24 bool = (!(468 != 87))
            _ = v24
        }
        _ = len(v18)
        if false {
            switch int(v3) {
            case 4294967295:
                var v25 float32 = v4
                _ = v25
            case 606:
                for fi26 := 0; fi26 < 4; fi26++ {
                    v27 := v5[546]
                    if true {
                        var v28 string = "hello"
                        if true {
                            for fi29 := 0; fi29 < 8; fi29++ {
                                switch v28 {
                                case "hello":
                                    if (v28 < "a") {
                                        v2 = v2
                                        _ = len(v5)
                                        v5[v27] = (v27 >> ((v8) & 63))
                                    }
                                case "\n\t":
                                    if len(v18) >= 2 {
                                        v18 = v18[1:len(v18)]
                                    }
                                }
                                _ = fi29
                            }
                        } else {
                            v7 -= v3
                            v17 -= (-(v17))
                        }
                        for fi30 := 0; fi30 < 14; fi30++ {
                            v31 := GenS3{F0: -1}
                            _ = fi30
                            _ = v31
                        }
                        _ = v28
                    } else {
                        for fi32 := 0; fi32 < 19; fi32++ {
                            v33 := GenS2{F0: v2, F1: v1, F2: 65535}
                            v2 = (^(v2))
                            _ = fi32
                            _ = v33
                        }
                        for fi34 := 0; fi34 < 12; fi34++ {
                            v27 = v27
                            _ = fi34
                        }
                    }
                    fmt.Println(v2, v3, v1)
                    _ = fi26
                    _ = v27
                }
            case 209:
                if false {
                    fmt.Println(v8)
                    var v35 int = v17
                    _ = v35
                }
            }
            if len(v18) > 0 {
                v18[0] = v17
            }
        }
    } else {
        _ = len(v5)
        v1 -= v1
        if len(v18) > 0 {
            v18[0] = (^(v17))
        }
    }
    for fi36 := 0; fi36 < 2; fi36++ {
        var v37 int16 = -1
        var v38 bool = false
        switch v17 {
        case 0:
            var v39 int = (fi36 & (^(v6)))
            _ = v39
        case 9223372036854775807:
            _ = len(v18)
        case 128:
            v40 := v5[2147483647]
            _ = v40
        default:
            for fi41 := 0; fi41 < 14; fi41++ {
                if v38 {
                    var v42 uint8 = (v1 - v1)
                    v5[777] = 42
                    var v43 float64 = float64(float32(0.0))
                    _ = v42
                    _ = v43
                } else {
                    fmt.Println(v3, v7, v17)
                }
                if (v17 < fi36) {
                    v37 = (v37 >> ((v8) & 63))
                }
                v2 -= v2
                _ = fi41
            }
        }
        _ = fi36
        _ = v37
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
