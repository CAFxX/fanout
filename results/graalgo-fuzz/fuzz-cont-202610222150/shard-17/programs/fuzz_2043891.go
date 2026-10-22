package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 uint64 }
type GenS2 struct { F0 int64; F1 uint }
type GenS3 struct { F0 int16; F1 uint16; F2 uint64 }
type GenS4 struct { F0 int32 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint64 = 632
    var v3 uint64 = 4294967295
    if ("hello" <= "a") {
        for fi4 := 0; fi4 < 9; fi4++ {
            var v5 uint8 = 1
            _ = fi4
            _ = v5
        }
        v3 -= v2
    } else {
        for fi6 := 0; fi6 < 18; fi6++ {
            if (213 <= 100) {
                v7 := []int{fi6, (fi6 >> ((1) & 63)), (fi6 + (fi6 << ((uint(v2)) & 63)))}
                var v8 float64 = (((0.0 - 0.1) / 1e10) * 0.0)
                _ = v7
                _ = v8
            } else {
                switch (fi6 - fi6) {
                case -884:
                    v9 := GenS1{F0: 754, F1: fi6, F2: (^(v3))}
                    _ = v9
                default:
                    v2 += 18446744073709551615
                }
            }
            fmt.Println(v2, fi6)
            _ = fi6
        }
        if true {
            var v10 float64 = 0.0
            v3 = v2
            _ = v10
        }
        for fi11 := 0; fi11 < 11; fi11++ {
            v2 = v3
            _ = fi11
        }
    }
    delete(v1, 2147483648)
    v12 := GenS1{F0: 168, F1: -860, F2: (v3 & (v3 - 1))}
    if (0 > 0) {
        v3 -= 0
        var v13 bool = ((true || (((v2 / (((v3 | 42)) | 1)) <= (v3 / (((v3 % (((^(v3))) | 1))) | 1))) || true)) || true)
        v13 = true
        _ = v13
    } else {
        if (4294967295 == 250) {
            fmt.Println(v3, v2)
            v14 := GenS1{F0: 42, F1: -685, F2: 2}
            _ = v14
        } else {
            v3 -= v3
            if ((!true) || true) {
                switch "hello" {
                case "a":
                    v2 = (v3 >> ((42) & 63))
                case "hello":
                    v2 = v2
                case "":
                    v2 += (v3 - v2)
                default:
                    v15 := []int{-2147483648}
                    _ = v15
                }
                for fi16 := 0; fi16 < 2; fi16++ {
                    var v17 int8 = 125
                    fmt.Println(v17, v2, v17)
                    _ = fi16
                    _ = v17
                }
                fmt.Println(v2, v3)
            } else {
                v18 := GenS1{F0: 4294967295, F1: -2147483648, F2: 818}
                v2 = 2
                _ = v18
            }
        }
        var v19 uint16 = 1
        _ = v19
    }
    for fi20 := 0; fi20 < 8; fi20++ {
        v21 := []int{fi20}
        for fi22 := 0; fi22 < 11; fi22++ {
            v3 = v3
            if true {
                if true {
                    var v23 uint16 = 687
                    var v24 int16 = -198
                    _ = v23
                    _ = v24
                }
                if (!(-2.25 == -2.25)) {
                    v3 = v2
                    switch "a" {
                    case "\n\t":
                        v3 = (^(v2))
                    case "foo bar":
                        v25 := GenS1{F0: 2, F1: (fi22 % ((-9223372036854775808) | 1)), F2: v2}
                        _ = v25
                    default:
                        if (431 != 977) {
                            v26 := GenS2{F0: -686, F1: 42}
                            if ((357 > 782) || (true && ((255 == 655) || true))) {
                                switch fi22 {
                                case 569:
                                    var v27 uint32 = 4294967295
                                    _ = v27
                                case 0:
                                    delete(v1, 2)
                                case 326:
                                    v21 = append(v21, (fi20 - (^(fi22))))
                                    if len(v21) >= 2 {
                                        v21 = v21[1:len(v21)]
                                    }
                                    var v28 float32 = -2.25
                                    if len(v21) >= 2 {
                                        v21 = v21[1:len(v21)]
                                    }
                                    v3 -= (v2 & (v3 - 4294967295))
                                    _ = v28
                                }
                            }
                            _ = v26
                        }
                    }
                }
                if false {
                    v3 = v2
                }
                v29 := GenS3{F0: 923, F1: 2, F2: v2}
                _ = v29
            } else {
                var v30 float64 = 0.0
                _ = v30
            }
            _ = fi22
        }
        if len(v21) >= 2 {
            v21 = v21[1:len(v21)]
        }
        _ = fi20
        _ = v21
    }
    delete(v1, 2147483647)
    fmt.Println(v2, v3)
    v31 := GenS4{F0: -939}
    var v32 float64 = 0.1
    v2 -= (v3 << ((704) & 63))
    fmt.Println(v32, v2)
    v33 := GenS4{F0: -2147483648}
    var v34 uint64 = v2
    var v35 int = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
}
