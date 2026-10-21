package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 int32; F1 uint16 }
type GenS3 struct { F0 int64; F1 int32 }
type GenS4 struct { F0 uint64; F1 int8 }
type GenS5 struct { F0 uint; F1 uint8 }

func main() {
    var v1 uint32 = 2147483648
    v2 := make(map[int]string)
    var v3 int16 = 0
    v4 := make(map[int]bool)
    var v5 uint8 = 80
    if (!(true && false)) {
        v6 := GenS1{F0: 1, F1: 4294967295}
        _ = v6
    } else {
        for fi7 := 0; fi7 < 19; fi7++ {
            var v8 int = fi7
            switch 127 {
            case 128:
                var v9 int32 = 42
                _ = v9
            case 786:
                var v10 uint16 = 32767
                _ = v10
            default:
                var v11 float64 = (float64(1e10) + (((3.14159 + float64(0.1)) + (-2.25 - -2.25)) / 0.1))
                _ = v11
            }
            fmt.Println(v8)
            _ = fi7
            _ = v8
        }
        v5 += v5
    }
    var v12 float32 = 0.0
    v12 += 1e10
    var v13 int = 42
    var v14 float64 = (float64(float32(((0.1 / 0.1) + 0.0))) / 3.14159)
    for fi15 := 0; fi15 < 10; fi15++ {
        for fi16 := 0; fi16 < 6; fi16++ {
            if (!false) {
                v17 := GenS2{F0: 2147483647, F1: 65535}
                v18 := GenS2{F0: 0, F1: 42}
                if (((1.5 * -2.25) + (float32((v14 - (((float64((v12 - v12)) / 0.1) - ((v14 - 0.1) / (((float64(0.1) / v14) * v14) / (((0.0 * float64(v12)) + (v14 * (1.5 * v14))) * v14)))) - (v14 - ((v14 - ((v14 / (v14 * v14)) / 1.0)) * v14))))) - (1.5 * 1e10))) >= float32(((((v14 * 3.14159) / float64(float32(float64((v12 + (float32(v14) + v12)))))) + (float64((v12 * 1.5)) + 0.0)) / v14))) {
                    switch (fi15 * fi16) {
                    case 506:
                        fmt.Println(v14, v12, v14)
                    }
                    var v19 int32 = 1
                    if true {
                        fmt.Println(fi15)
                    } else {
                        if ((("a" < "a") && ((v1 | (^(v1))) != (^(v1)))) && ((v1 / ((uint32(v5)) | 1)) >= (v1 ^ v1))) {
                            var v20 uint8 = v5
                            v5 += v20
                            _ = v20
                        } else {
                            var v21 uint8 = v5
                            v22 := []int{fi16, (^(fi16)), fi16, (fi15 >> ((4294967295) & 63))}
                            var v23 uint16 = 32767
                            _ = v21
                            _ = v22
                            _ = v23
                        }
                    }
                    _ = v19
                } else {
                    v24 := GenS1{F0: 18446744073709551615, F1: 1}
                    v14 = v14
                    v25 := GenS2{F0: -379, F1: uint16(v1)}
                    _ = v24
                    _ = v25
                }
                _ = v17
                _ = v18
            } else {
                v2[fi16] = "foo bar"
                var v26 uint = 386
                var v27 int16 = (v3 >> ((235) & 63))
                var v28 uint8 = (v5 & (v5 << (((v26 - (v26 | 0))) & 63)))
                v29 := GenS1{F0: v26, F1: 42}
                var v30 uint16 = 408
                if true {
                    var v31 int64 = 9223372036854775807
                    v32 := []int{(fi15 | (fi16 >> ((389) & 63)))}
                    fmt.Println(v5)
                    _ = v31
                    _ = v32
                }
                _ = v26
                _ = v27
                _ = v28
                _ = v29
                _ = v30
            }
            var v33 int8 = -128
            var v34 uint32 = 4294967295
            v5 = 163
            var v35 int8 = (v33 | 42)
            _ = fi16
            _ = v33
            _ = v34
            _ = v35
        }
        fmt.Println(v5)
        _ = fi15
    }
    var v36 uint16 = 311
    v5 += v5
    if (!true) {
        v12 += (((v12 + v12) * v12) + float32(float64(0.0)))
        v5 = v5
    }
    var v37 int = (v13 - (-(v13)))
    v38 := GenS2{F0: 2147483647, F1: (^(v36))}
    fmt.Println(v12, v3)
    if (!false) {
        var v39 uint32 = (v1 & v1)
        v12 = 3.14159
        _ = v39
    }
    v40 := GenS3{F0: -9223372036854775808, F1: 0}
    switch int(v36) {
    case 32767:
        v41 := []int{v37, 32768, 9223372036854775807}
        _ = v41
    case 0:
        v42 := GenS4{F0: 106, F1: 127}
        _ = v42
    case -1:
        for fi43 := 0; fi43 < 7; fi43++ {
            v44 := GenS3{F0: 1, F1: 2147483647}
            v45 := GenS5{F0: 477, F1: (v5 + 254)}
            _ = fi43
            _ = v44
            _ = v45
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v40
    _ = v5
}
