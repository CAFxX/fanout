package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 uint16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int }

func main() {
    var v1 int16 = 1000
    v2 := make(map[string]uint32)
    var v3 int8 = 57
    switch 32767 {
    case 4294967295:
        var v4 int32 = -996
        _ = v4
    case 882:
        var v5 string = (("a" + "\n\t") + ("foo bar" + "a"))
        _ = v5
    default:
        var v6 int16 = (v1 & (v1 | v1))
        _ = v6
    }
    if (4294967295 == -9223372036854775808) {
        var v7 float64 = 3.14159
        v8 := GenS1{F0: 32767, F1: v1, F2: 110}
        _ = v7
        _ = v8
    } else {
        v1 = v1
        v3 = (v3 >> ((854) & 63))
    }
    var v9 uint64 = 710
    var v10 float64 = (float64(0.1) + -2.25)
    v1 = (v1 / ((int16(v3)) | 1))
    fmt.Println(v9, v10)
    v3 = v3
    fmt.Println(v9)
    v11 := []int{4294967295, -9223372036854775808}
    if (2147483647 != 32767) {
        for fi12 := 0; fi12 < 4; fi12++ {
            switch (fi12 + -147) {
            case -395:
                v10 = ((v10 + v10) + (0.1 * v10))
            case 9223372036854775807:
                switch 128 {
                case -9223372036854775808:
                    v3 -= (v3 * v3)
                default:
                    v11 = append(v11, (fi12 << ((65535) & 63)))
                }
            case -2147483648:
                v9 = (v9 ^ (v9 / ((0) | 1)))
            }
            _ = fi12
        }
        if false {
            var v13 uint64 = (v9 >> ((488) & 63))
            v14 := GenS1{F0: 32767, F1: v1, F2: 255}
            if (-32768 != v1) {
                for fi15 := 0; fi15 < 5; fi15++ {
                    switch fi15 {
                    case 0:
                        v13 += 42
                    default:
                        if true {
                            for fi16 := 0; fi16 < 12; fi16++ {
                                v1 -= 1
                                fmt.Println(fi16, v9, v9)
                                _ = fi16
                            }
                            var v17 uint8 = 230
                            _ = v17
                        }
                    }
                    v18 := GenS1{F0: fi15, F1: v1, F2: 32767}
                    _ = fi15
                    _ = v18
                }
                var v19 uint8 = 2
                v19 -= v19
                _ = v19
            } else {
                for fi20 := 0; fi20 < 10; fi20++ {
                    v1 -= v1
                    v21 := GenS2{F0: -492}
                    _ = fi20
                    _ = v21
                }
                _ = len(v2)
                v22 := GenS3{F0: 42}
                _ = v22
            }
            _ = v13
            _ = v14
        }
        for fi23 := 0; fi23 < 5; fi23++ {
            v24 := GenS4{F0: (fi23 + fi23)}
            v3 -= v3
            _ = fi23
            _ = v24
        }
    } else {
        var v25 bool = false
        _ = len(v11)
        _ = v25
    }
    var v26 int16 = v1
    v27 := GenS2{F0: 0}
    var v28 float64 = v10
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v9
}
