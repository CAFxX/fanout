package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint16; F1 uint16; F2 uint }
func genFn0(p0 uint32, p1 int16, p2 uint32) int32 {
    var v4 uint32 = (p0 << ((4294967295) & 63))
    v5 := []int{32768, 127, -9223372036854775808, 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 42
}

func main() {
    var v1 float64 = (0.1 + 0.0)
    v2 := []int{-2147483648, -273}
    v3 := make(map[uint32]int)
    var v6 float32 = 0.1
    if false {
        for fi7 := 0; fi7 < 11; fi7++ {
            v6 = float32(v1)
            switch 127 {
            case -519:
                if true {
                    for fi8 := 0; fi8 < 3; fi8++ {
                        if (true && true) {
                            v6 = ((float32(0.0) + float32(float64(v6))) / (float32((1.5 * float64(-2.25))) * 3.14159))
                            var v9 int16 = 0
                            var v10 float64 = (0.0 * v1)
                            _ = v10
                            _ = v9
                        } else {
                            switch "" {
                            case "":
                                var v11 float32 = v6
                                _ = v11
                            default:
                                delete(v3, 65535)
                            }
                            var v12 float64 = (v1 - float64(v6))
                            switch fi7 {
                            case -242:
                                if false {
                                    v3[768] = (fi8 & (fi7 ^ fi8))
                                    var v13 int16 = -76
                                    var v14 int = fi7
                                    _ = v13
                                    _ = v14
                                }
                            default:
                                var v15 float64 = v1
                                _ = v15
                            }
                            if false {
                                v12 += (3.14159 + ((v1 + (v12 - 1e10)) - (float64((float32(v1) * (v6 * 1.5))) - 3.14159)))
                                v16 := GenS1{F0: (fi7 % ((fi8) | 1))}
                                v6 = v6
                                _ = v16
                            }
                            _ = v12
                        }
                        _ = fi8
                    }
                    v17 := GenS1{F0: (-(fi7))}
                    if (!false) {
                        var v18 int64 = -1
                        for fi19 := 0; fi19 < 19; fi19++ {
                            var v20 int32 = int32(v18)
                            switch (fi19 - int(v18)) {
                            case 42:
                                var v21 uint8 = 128
                                _ = v21
                            default:
                                var v22 uint16 = 695
                                _ = v22
                            }
                            var v23 int64 = 56
                            v6 += v6
                            _ = fi19
                            _ = v20
                            _ = v23
                        }
                        _ = v18
                    }
                    _ = v17
                } else {
                    var v24 uint32 = 201
                    _ = v24
                }
            }
            fmt.Println(v1, v1, fi7)
            _ = fi7
        }
        v25 := GenS2{F0: 100}
        for fi26 := 0; fi26 < 20; fi26++ {
            _ = len(v2)
            if false {
                var v27 uint32 = 2147483648
                var v28 float64 = 1e10
                _ = v27
                _ = v28
            } else {
                v1 += 0.1
                if true {
                    v29 := GenS3{F0: (true || (true || true))}
                    v1 -= v1
                    _ = len(v2)
                    _ = v29
                }
            }
            _ = fi26
        }
        _ = v25
    } else {
        var v30 int32 = -1
        var v31 int64 = -9223372036854775808
        v6 += 0.0
        _ = v30
        _ = v31
    }
    v1 += v1
    if ((v1 == (v1 - (float64((v6 + -2.25)) + 0.1))) || true) {
        var v32 int64 = -1
        _ = v32
    } else {
        v33 := GenS4{F0: 255, F1: 584, F2: 846}
        var v34 int32 = 174
        switch "\n\t" {
        case "":
            switch 0 {
            case 32767:
                v34 += (v34 % (((v34 - v34)) | 1))
            }
        default:
            var v35 uint16 = 42
            _ = v35
        }
        _ = v33
        _ = v34
    }
    var v36 uint = 1
    v6 -= -2.25
    var v37 string = ""
    var v38 string = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v6
}
