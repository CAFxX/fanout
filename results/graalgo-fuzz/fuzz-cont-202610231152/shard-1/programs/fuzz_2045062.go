package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 uint64 }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 string; F1 uint8 }

func main() {
    var v1 uint64 = 2
    v2 := make(map[uint32]int)
    v3 := make(map[int]int)
    v4 := []uint32{0}
    if (true && (1e10 < 0.1)) {
        for fi5 := 0; fi5 < 1; fi5++ {
            delete(v3, fi5)
            for fi6 := 0; fi6 < 10; fi6++ {
                var v7 uint8 = 77
                if (!true) {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
                v1 += 841
                _ = fi6
                _ = v7
            }
            v8 := v2[2147483647]
            _ = fi5
            _ = v8
        }
        var v9 float32 = 0.0
        var v10 uint = 65535
        _ = v10
        _ = v9
    }
    v1 -= v1
    v11 := GenS1{F0: v1, F1: 668, F2: (v1 + v1)}
    switch 59 {
    case -2147483648:
        switch 128 {
        case -2147483648:
            var v12 float64 = float64((1.5 * 1.5))
            _ = v12
        case 127:
            v13 := GenS1{F0: v1, F1: 958, F2: (v1 + v1)}
            _ = v13
        case 32768:
            v1 += v1
        }
    default:
        fmt.Println(v1)
    }
    fmt.Println(v1)
    var v14 int16 = 1
    var v15 int32 = 1
    var v16 int64 = -1
    v14 = -735
    v1 -= v1
    var v17 int64 = (v16 << ((4294967295) & 63))
    if false {
        for fi18 := 0; fi18 < 18; fi18++ {
            v4 = append(v4, 861)
            for fi19 := 0; fi19 < 13; fi19++ {
                if (v15 <= (v15 / (((v15 + v15)) | 1))) {
                    switch "foo bar" {
                    case "a":
                        fmt.Println(v17, v1, v14)
                    }
                    var v20 bool = ((!(fi19 > 1)) && false)
                    var v21 int16 = v14
                    _ = v20
                    _ = v21
                }
                _ = fi19
            }
            _ = fi18
        }
        if len(v4) > 0 {
            v4[0] = 884
        }
        v17 += (v16 | v16)
    } else {
        _ = len(v4)
        var v22 float64 = 1.5
        fmt.Println(v15)
        _ = v22
    }
    fmt.Println(v1, v15, v14)
    for fi23 := 0; fi23 < 8; fi23++ {
        v24 := GenS1{F0: v1, F1: 663, F2: v1}
        v14 += (v14 | v14)
        _ = fi23
        _ = v24
    }
    var v25 int32 = (v15 ^ v15)
    v26 := GenS2{F0: 1, F1: 958}
    v27 := GenS3{F0: "\n\t", F1: 1}
    if (2147483647 != 65535) {
        if false {
            switch 327 {
            case 127:
                var v28 uint = 255
                _ = v28
            default:
                if (808 < 4294967295) {
                    for fi29 := 0; fi29 < 15; fi29++ {
                        var v30 int16 = (v14 - (v14 << ((54) & 63)))
                        if len(v4) > 0 {
                            v4[0] = uint32(v1)
                        }
                        _ = fi29
                        _ = v30
                    }
                } else {
                    v14 += v14
                    v1 += v1
                }
            }
            _ = len(v3)
            var v31 int16 = v14
            _ = v31
        }
        switch 32767 {
        case 265:
            v3[-1] = 4294967295
        }
        if true {
            var v32 int32 = (^(v25))
            _ = v32
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
