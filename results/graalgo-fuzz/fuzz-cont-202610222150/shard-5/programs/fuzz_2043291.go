package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 uint; F1 uint8; F2 bool }

func main() {
    var v1 bool = (!false)
    v2 := make(map[uint32]uint32)
    var v3 uint32 = 65535
    var v4 float32 = -2.25
    var v5 int16 = -310
    v6 := GenS1{F0: 194, F1: 255}
    v7 := []int{-2147483648}
    fmt.Println(v1)
    v8 := GenS1{F0: 167, F1: 255}
    if v1 {
        v4 -= v4
        if false {
            if (100 == -1) {
                for fi9 := 0; fi9 < 19; fi9++ {
                    var v10 bool = (((v4 * v4) * v4) == (v4 * v4))
                    switch (fi9 >> ((18446744073709551615) & 63)) {
                    case 4294967295:
                        v11 := GenS2{F0: v5, F1: (1 == 9223372036854775807)}
                        _ = v11
                    case 1:
                        v10 = v10
                    default:
                        var v12 uint16 = 723
                        _ = v12
                    }
                    if v1 {
                        if (4294967295 < (v3 - (v3 >> ((30) & 63)))) {
                            v13 := GenS1{F0: 2, F1: 42}
                            for fi14 := 0; fi14 < 16; fi14++ {
                                if v10 {
                                    var v15 uint16 = 2
                                    var v16 uint16 = 42
                                    var v17 string = "\n\t"
                                    _ = v15
                                    _ = v16
                                    _ = v17
                                } else {
                                    v3 = v3
                                    var v18 uint64 = 907
                                    v4 += (v4 - (v4 + (0.0 + 0.0)))
                                    _ = v18
                                }
                                _ = fi14
                            }
                            if false {
                                if ((v3 & (v3 & v3)) != v3) {
                                    v19 := v2[(^(v3))]
                                    var v20 int = 127
                                    _ = v19
                                    _ = v20
                                }
                            }
                            _ = v13
                        } else {
                            v21 := GenS2{F0: -1, F1: v1}
                            if v10 {
                                v22 := GenS2{F0: (v5 << ((4294967295) & 63)), F1: v1}
                                _ = v22
                            } else {
                                for fi23 := 0; fi23 < 10; fi23++ {
                                    var v24 uint = 402
                                    v3 = v3
                                    _ = fi23
                                    _ = v24
                                }
                                v10 = false
                            }
                            _ = v21
                        }
                        v1 = ((1.5 - -2.25) < 1e10)
                    }
                    _ = fi9
                    _ = v10
                }
                if (1 >= 2) {
                    var v25 uint16 = 884
                    _ = v25
                }
                var v26 int64 = 1
                _ = v26
            } else {
                var v27 float32 = 1.5
                delete(v2, (v3 ^ 2))
                _ = v27
            }
        } else {
            v5 = 42
            v28 := GenS2{F0: (v5 * v5), F1: (127 != -55)}
            _ = v28
        }
    }
    switch 1 {
    case 128:
        for fi29 := 0; fi29 < 14; fi29++ {
            for fi30 := 0; fi30 < 1; fi30++ {
                var v31 uint8 = 0
                if len(v7) > 0 {
                    v7[0] = (fi29 << ((1) & 63))
                }
                _ = fi30
                _ = v31
            }
            v1 = v1
            _ = fi29
        }
    case -9223372036854775808:
        v32 := GenS1{F0: 1, F1: 42}
        _ = v32
    }
    fmt.Println(v1)
    var v33 int = -380
    v1 = (((v3 ^ (v3 * 989)) < (v3 * v3)) && v1)
    var v34 uint = 65535
    for fi35 := 0; fi35 < 3; fi35++ {
        if v1 {
            v33 += 9223372036854775807
        }
        _ = fi35
    }
    v36 := v2[v3]
    fmt.Println(v36, v33, v33)
    var v37 uint8 = 1
    v38 := GenS3{F0: 655, F1: (v37 + (v37 | (v37 - (v37 / ((v37) | 1))))), F2: v1}
    var v39 int32 = -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
