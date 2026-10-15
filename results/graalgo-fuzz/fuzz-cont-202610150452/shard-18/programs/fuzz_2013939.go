package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 int8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint32; F1 bool }
type GenS4 struct { F0 int16; F1 uint32 }
func genFn0(p0 int64, p1 int16, p2 int) uint {
    if false {
        var v4 uint64 = 649
        p1 = p1
        for fi5 := 0; fi5 < 11; fi5++ {
            v6 := []int{-878}
            if true {
                if false {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                } else {
                    switch fi5 {
                    case -9223372036854775808:
                        switch 32768 {
                        case 42:
                            if (42 > -64) {
                                p1 = -32768
                                v7 := GenS1{F0: "\n\t", F1: -504, F2: -37}
                                _ = v7
                            }
                        case 2147483647:
                            p2 -= (^(fi5))
                        case 1:
                            p0 = -1
                        default:
                            v6 = append(v6, int(p0))
                        }
                    case 128:
                        _ = len(v6)
                    case -2147483648:
                        for fi8 := 0; fi8 < 8; fi8++ {
                            var v9 uint64 = v4
                            p2 = -1
                            _ = fi8
                            _ = v9
                        }
                    }
                    if true {
                        v10 := GenS2{F0: 65535}
                        if (false && true) {
                            switch fi5 {
                            case -1:
                                p1 += (p1 >> ((0) & 63))
                            case -2147483648:
                                var v11 string = ("" + "a")
                                _ = v11
                            }
                            v12 := GenS1{F0: "\n\t", F1: 0, F2: 127}
                            _ = v12
                        } else {
                            if (!(925 <= 2147483648)) {
                                var v13 float64 = 1.5
                                var v14 uint8 = 42
                                _ = v13
                                _ = v14
                            } else {
                                var v15 uint32 = 1
                                _ = v15
                            }
                        }
                        _ = v10
                    }
                }
                p1 = p1
                if len(v6) > 0 {
                    v6[0] = (p2 / (((p2 * 127)) | 1))
                }
            }
            _ = fi5
            _ = v6
        }
        _ = v4
    }
    p1 = -1
    _ = p0
    _ = p1
    _ = p2
    return 410
}

func main() {
    v1 := []uint32{2147483647, 272, 0}
    var v2 int16 = 655
    var v3 float32 = 1e10
    v2 += (v2 * (v2 << ((4294967295) & 63)))
    for fi16 := 0; fi16 < 18; fi16++ {
        if false {
            v17 := GenS3{F0: 2, F1: (100 > 1)}
            _ = v17
        }
        _ = fi16
    }
    v3 -= 0.1
    v3 -= ((v3 + float32(((0.1 - ((3.14159 * (1.5 / (1.5 / float64(v3)))) + 1.5)) * 1e10))) - v3)
    var v18 bool = (float32((3.14159 / 1e10)) < v3)
    var v19 uint32 = 988
    v20 := GenS3{F0: v19, F1: v18}
    v21 := GenS4{F0: 0, F1: 42}
    switch "\n\t" {
    case "a":
        fmt.Println(v2, v18, v18)
    case "":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
