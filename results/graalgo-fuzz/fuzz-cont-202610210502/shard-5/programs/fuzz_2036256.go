package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32; F1 int32; F2 uint64 }
type GenS3 struct { F0 uint16; F1 int }
func genFn0(p0 int16) int16 {
    if (!false) {
        v7 := []int{0, 4294967295, 455, 127}
        for fi8 := 0; fi8 < 13; fi8++ {
            if (1.5 != 3.14159) {
                if (101 > 65535) {
                    switch 127 {
                    case -9223372036854775808:
                        if (!(752 > 0)) {
                            p0 -= p0
                            p0 += p0
                            p0 = -32768
                        }
                    case 32768:
                        var v9 int16 = (^(p0))
                        _ = v9
                    case 32767:
                        for fi10 := 0; fi10 < 9; fi10++ {
                            v11 := GenS1{F0: 18}
                            _ = fi10
                            _ = v11
                        }
                    default:
                        v12 := GenS2{F0: 0, F1: 42, F2: 2}
                        _ = v12
                    }
                    if true {
                        var v13 int16 = (p0 << ((2) & 63))
                        var v14 uint64 = 559
                        _ = v13
                        _ = v14
                    }
                    var v15 bool = (255 < 65535)
                    _ = v15
                }
            } else {
                v7 = append(v7, 32767)
                p0 += (p0 * p0)
            }
            if true {
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                switch "" {
                case "a":
                    fmt.Println(fi8, fi8)
                default:
                    var v16 uint = 18446744073709551615
                    _ = v16
                }
            } else {
                _ = len(v7)
                p0 = (p0 << ((437) & 63))
            }
            v17 := GenS3{F0: 32767, F1: fi8}
            _ = fi8
            _ = v17
        }
        _ = v7
    }
    var v18 int32 = -256
    v19 := []int{128, 128}
    v18 -= -1
    _ = p0
    _ = v18
    _ = v19
    return p0
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []string{}
    v3 := make(map[string]int)
    var v4 float32 = 3.14159
    v5 := []uint32{457, 0, 742, 42}
    v6 := []int{32768, 128}
    for fi20 := 0; fi20 < 4; fi20++ {
        switch (fi20 / ((fi20) | 1)) {
        case 32767:
            for fi21 := 0; fi21 < 1; fi21++ {
                if true {
                    switch "hello" {
                    case "\n\t":
                        v4 = v4
                    default:
                        fmt.Println(v4, v4)
                    }
                    var v22 uint = 65535
                    _ = v22
                } else {
                    if (4294967295 < 255) {
                        var v23 int8 = 0
                        _ = v23
                    }
                    fmt.Println(v4)
                }
                var v24 float64 = (-2.25 / 1.0)
                _ = fi21
                _ = v24
            }
        case 723:
            for fi25 := 0; fi25 < 14; fi25++ {
                v26 := GenS3{F0: 2, F1: 0}
                var v27 uint64 = 348
                var v28 float32 = v4
                _ = fi25
                _ = v26
                _ = v27
                _ = v28
            }
        default:
            v4 += ((v4 / v4) + ((v4 * v4) * float32(1.5)))
        }
        v4 -= v4
        _ = fi20
    }
    var v29 string = ""
    var v30 bool = false
    var v31 uint16 = 367
    var v32 float32 = (((v4 - v4) + ((((v4 * (v4 * (3.14159 + (v4 / v4)))) / 1.5) + v4) + float32(3.14159))) * 0.1)
    if (-2.25 != (-2.25 - ((1e10 + 1e10) / -2.25))) {
        v29 = v29
        var v33 uint16 = 65535
        v32 = v4
        _ = v33
    }
    v34 := v1[2147483648]
    var v35 string = "a"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
