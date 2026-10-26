package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 int }
type GenS2 struct { F0 uint8; F1 uint32; F2 int8 }
type GenS3 struct { F0 string }

func main() {
    var v1 int64 = 42
    var v2 int16 = -1
    var v3 uint8 = 42
    v4 := []uint32{76, 65535, 2, 1}
    v5 := make(map[uint32]int)
    for fi6 := 0; fi6 < 4; fi6++ {
        v3 = 2
        var v7 uint8 = (v3 + 0)
        _ = fi6
        _ = v7
    }
    var v8 uint32 = 2147483648
    fmt.Println(v3, v3, v3)
    if false {
        v1 -= (-(v1))
        v9 := v5[(v8 ^ 2147483648)]
        _ = len(v5)
        _ = v9
    } else {
        if (3.14159 != 3.14159) {
            v3 += (v3 ^ 255)
        } else {
            if false {
                for fi10 := 0; fi10 < 4; fi10++ {
                    v11 := GenS1{F0: 594, F1: "hello", F2: fi10}
                    _ = fi10
                    _ = v11
                }
                var v12 int16 = v2
                v12 = (-(v2))
                _ = v12
            }
            for fi13 := 0; fi13 < 16; fi13++ {
                v14 := GenS2{F0: (^(v3)), F1: (v8 % ((v8) | 1)), F2: 0}
                _ = fi13
                _ = v14
            }
        }
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v4 = append(v4, (v8 * 309))
    var v15 float64 = 1e10
    if (-12 >= -50) {
        var v16 uint = 2
        if true {
            v3 = v3
            if (!(-2147483648 < -48)) {
                v16 += (v16 / (((v16 - v16)) | 1))
                if ((v1 ^ v1) > v1) {
                    v5[2147483647] = 128
                    v17 := v5[v8]
                    _ = v17
                }
            } else {
                var v18 int64 = v1
                v2 -= -160
                _ = v18
            }
        } else {
            var v19 uint32 = (v8 ^ (v8 - v8))
            _ = v19
        }
        switch "" {
        case "hello":
            var v20 float64 = (v15 + 3.14159)
            _ = v20
        case "a":
            v16 = uint(v3)
        case "foo bar":
            switch -759 {
            case 352:
                fmt.Println(v8)
            default:
                for fi21 := 0; fi21 < 8; fi21++ {
                    if false {
                        v22 := GenS1{F0: int32(v2), F1: "hello", F2: fi21}
                        fmt.Println(v1, v8)
                        for fi23 := 0; fi23 < 10; fi23++ {
                            v2 -= v2
                            switch ("" + "a") {
                            case "":
                                v24 := GenS3{F0: "foo bar"}
                                _ = v24
                            case "a":
                                v4 = append(v4, 832)
                            }
                            v25 := GenS1{F0: 0, F1: (("" + "foo bar") + "foo bar"), F2: 552}
                            _ = fi23
                            _ = v25
                        }
                        _ = v22
                    }
                    _ = fi21
                }
            }
        }
        _ = v16
    } else {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v26 uint8 = (^(v3))
        _ = v26
    }
    v8 = (v8 >> ((4294967295) & 63))
    v27 := v5[v8]
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
