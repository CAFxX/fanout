package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 uint16; F1 uint; F2 uint32 }
func genFn0(p0 int64) bool {
    for fi6 := 0; fi6 < 5; fi6++ {
        for fi7 := 0; fi7 < 20; fi7++ {
            var v8 int = fi6
            var v9 uint = 198
            fmt.Println(v8, p0)
            _ = fi7
            _ = v8
            _ = v9
        }
        if (true || false) {
            v10 := []int{fi6, 127}
            if (0 > 129) {
                p0 += p0
            } else {
                p0 += 465
                var v11 int8 = -128
                for fi12 := 0; fi12 < 7; fi12++ {
                    if ((p0 << ((255) & 63)) < 1) {
                        if false {
                            fmt.Println(fi12)
                            p0 += -9223372036854775808
                        } else {
                            _ = len(v10)
                            fmt.Println(fi12)
                            if (1.5 != 3.14159) {
                                var v13 uint8 = 128
                                if len(v10) > 0 {
                                    v10[0] = (fi6 >> ((2) & 63))
                                }
                                _ = len(v10)
                                var v14 int64 = int64(v11)
                                fmt.Println(v14)
                                _ = v13
                                _ = v14
                            }
                        }
                        fmt.Println(fi6)
                    } else {
                        var v15 float64 = (0.1 * -2.25)
                        var v16 float32 = ((-2.25 / 1.5) / ((((float32((v15 / (v15 / (v15 / 3.14159)))) * (float32(v15) - 0.1)) - 1.5) - -2.25) / 1.0))
                        _ = v15
                        _ = v16
                    }
                    _ = fi12
                }
                _ = v11
            }
            _ = v10
        }
        _ = fi6
    }
    v17 := GenS1{F0: 32768, F1: 406}
    if ((p0 | (-(p0))) != (^(p0))) {
        for fi18 := 0; fi18 < 4; fi18++ {
            p0 = p0
            _ = fi18
        }
        v19 := GenS2{F0: 4294967295, F1: -128}
        var v20 int32 = 1
        _ = v19
        _ = v20
    }
    _ = p0
    _ = v17
    return (0.1 < 0.1)
}

func main() {
    var v1 uint64 = 1
    v2 := make(map[string]string)
    var v3 int32 = 42
    v4 := make(map[uint32]int)
    v5 := make(map[int]int)
    _ = len(v5)
    v21 := GenS3{F0: 255, F1: 483, F2: 197}
    v22 := GenS3{F0: 648, F1: 77, F2: 2}
    var v23 uint = 890
    v24 := GenS3{F0: 255, F1: 0, F2: uint32(v23)}
    var v25 uint8 = 0
    var v26 float32 = 0.0
    var v27 uint64 = 0
    v23 -= uint(v27)
    var v28 float32 = ((3.14159 * (1e10 * v26)) * (1.5 + ((0.1 + (0.1 + 0.1)) - float32(float64((3.14159 / v26))))))
    for fi29 := 0; fi29 < 1; fi29++ {
        if true {
            switch fi29 {
            case -320:
                var v30 uint64 = 0
                _ = v30
            default:
                if true {
                    var v31 bool = true
                    _ = v31
                }
            }
            switch fi29 {
            case 886:
                var v32 uint8 = v25
                _ = v32
            default:
                v33 := GenS2{F0: v27, F1: 42}
                _ = v33
            }
        }
        _ = fi29
    }
    v1 = v1
    for fi34 := 0; fi34 < 8; fi34++ {
        v35 := GenS3{F0: 0, F1: (v23 & v23), F2: 422}
        v3 = (v3 >> (((^(v23))) & 63))
        _ = fi34
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
