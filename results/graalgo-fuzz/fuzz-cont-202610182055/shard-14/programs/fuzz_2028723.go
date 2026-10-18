package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int32; F1 int16 }
type GenS6 struct { F0 uint; F1 uint16 }

func main() {
    var v1 uint8 = 2
    v2 := make(map[int]string)
    var v3 bool = true
    var v4 float64 = (float64(-2.25) - 3.14159)
    if (!false) {
        v5 := GenS1{F0: 42}
        _ = v5
    }
    for fi6 := 0; fi6 < 18; fi6++ {
        var v7 uint = 18446744073709551615
        v8 := GenS1{F0: v7}
        v1 = (v1 << (((v7 | (v7 >> (((^(v7))) & 63)))) & 63))
        _ = fi6
        _ = v7
        _ = v8
    }
    fmt.Println(v3, v3, v1)
    for fi9 := 0; fi9 < 7; fi9++ {
        v10 := []int{fi9, -505}
        _ = fi9
        _ = v10
    }
    if v3 {
        var v11 bool = v3
        var v12 int8 = 100
        _ = v11
        _ = v12
    }
    v13 := v2[1]
    v14 := GenS2{F0: -567}
    for fi15 := 0; fi15 < 15; fi15++ {
        if true {
            v1 = (v1 & 1)
            for fi16 := 0; fi16 < 8; fi16++ {
                var v17 int = 832
                v18 := GenS3{F0: 334}
                var v19 int8 = -128
                v17 -= fi16
                switch (fi16 % ((v17) | 1)) {
                case 0:
                    v20 := []int{(fi16 | (v17 % (((fi16 * 951)) | 1)))}
                    fmt.Println(v19, v4, v13)
                    _ = v20
                case 128:
                    v4 += 0.0
                    v3 = (-760 <= 42)
                default:
                    v21 := GenS4{F0: 42}
                    _ = v21
                }
                _ = fi16
                _ = v17
                _ = v18
                _ = v19
            }
            for fi22 := 0; fi22 < 10; fi22++ {
                var v23 int = 0
                var v24 bool = ((0.0 + 0.0) >= 1.5)
                var v25 uint16 = 2
                _ = fi22
                _ = v23
                _ = v24
                _ = v25
            }
        } else {
            if false {
                v3 = (v3 && v3)
                var v26 float64 = v4
                v3 = (v26 <= -2.25)
                _ = v26
            } else {
                switch v13 {
                case "a":
                    if (2 > 65535) {
                        var v27 int = -9223372036854775808
                        fmt.Println(v1, v3, v13)
                        _ = v27
                    }
                case "":
                    var v28 int64 = 9223372036854775807
                    _ = v28
                case "hello":
                    v29 := GenS5{F0: -1, F1: 0}
                    _ = v29
                default:
                    if (954 == -9223372036854775808) {
                        v30 := []int{fi15, (fi15 | fi15), fi15}
                        _ = v30
                    } else {
                        fmt.Println(v4, v3, v1)
                        if v3 {
                            var v31 uint32 = 304
                            fmt.Println(fi15)
                            _ = v31
                        }
                    }
                }
                v32 := v2[fi15]
                if v3 {
                    var v33 int8 = 104
                    if (574 <= 0) {
                        v32 = "hello"
                    } else {
                        v3 = v3
                        var v34 int = (fi15 | fi15)
                        if v3 {
                            var v35 uint32 = 2
                            var v36 string = v32
                            _ = v35
                            _ = v36
                        }
                        _ = v34
                    }
                    v37 := []int{(fi15 / ((fi15) | 1)), fi15, fi15, 1}
                    _ = v33
                    _ = v37
                }
                _ = v32
            }
            if false {
                fmt.Println(v3)
                var v38 float64 = (-2.25 - (((0.0 + (v4 + 0.1)) - (float64(0.1) / 1.0)) - v4))
                _ = v38
            } else {
                var v39 string = v13
                _ = v39
            }
            var v40 int8 = -1
            _ = v40
        }
        var v41 int = 4294967295
        _ = fi15
        _ = v41
    }
    var v42 int64 = 1
    for fi43 := 0; fi43 < 11; fi43++ {
        v4 = 3.14159
        for fi44 := 0; fi44 < 13; fi44++ {
            switch fi44 {
            case 293:
                fmt.Println(v1, fi43, v13)
            case -129:
                var v45 int8 = int8(v1)
                _ = v45
            default:
                if v3 {
                    v42 -= v42
                    v46 := GenS2{F0: -32768}
                    _ = v46
                } else {
                    v1 = (^(v1))
                }
            }
            _ = fi44
        }
        v4 += 1e10
        _ = fi43
    }
    _ = len(v2)
    if (255 < (v1 * v1)) {
        for fi47 := 0; fi47 < 16; fi47++ {
            v48 := GenS2{F0: -427}
            _ = fi47
            _ = v48
        }
        fmt.Println(v42, v1)
    } else {
        if true {
            for fi49 := 0; fi49 < 1; fi49++ {
                if v3 {
                    var v50 bool = false
                    v50 = v50
                    var v51 string = v13
                    _ = v50
                    _ = v51
                } else {
                    var v52 string = (v13 + v13)
                    var v53 float32 = ((0.1 + float32(float64(3.14159))) * 0.1)
                    v3 = v3
                    _ = v52
                    _ = v53
                }
                for fi54 := 0; fi54 < 14; fi54++ {
                    if (v3 || true) {
                        v1 += v1
                        v55 := GenS6{F0: 255, F1: 255}
                        _ = v55
                    }
                    if (((1e10 / 1.5) - 0.0) >= 1e10) {
                        var v56 string = "hello"
                        _ = v56
                    }
                    _ = fi54
                }
                _ = fi49
            }
            v1 -= (v1 & 235)
            _ = len(v2)
        } else {
            v57 := GenS4{F0: 975}
            v2[127] = v13
            _ = v57
        }
    }
    switch -2147483648 {
    case -1:
        for fi58 := 0; fi58 < 4; fi58++ {
            var v59 int16 = -311
            fmt.Println(v59, v1)
            _ = fi58
            _ = v59
        }
    default:
        v42 = (v42 ^ (v42 * (^(v42))))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v42
}
