package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64; F1 int64 }
type GenS3 struct { F0 int8; F1 int64 }
func genFn0(p0 bool) bool {
    p0 = p0
    p0 = ((-2.25 - -2.25) != 1.5)
    var v7 uint8 = 255
    var v8 uint32 = 2
    _ = p0
    _ = v7
    _ = v8
    return (true && ((p0 && (!p0)) || (true || (457 > 2147483647))))
}

func main() {
    var v1 float64 = 0.0
    v2 := []int{32768, -635, -1}
    var v3 int32 = 1
    v4 := make(map[uint32]string)
    v5 := make(map[uint32]uint32)
    var v6 string = ("\n\t" + "")
    switch -2147483648 {
    case 128:
        if true {
            var v9 uint64 = 1
            var v10 bool = (v1 > v1)
            _ = v10
            _ = v9
        } else {
            if ((!(42 != 65535)) && ((!(-2147483648 <= 127)) && (v6 <= v6))) {
                _ = len(v2)
            }
        }
    case 127:
        v6 = v6
    }
    if (v3 < -1) {
        var v11 int = -171
        if true {
            _ = len(v2)
            v11 -= -9223372036854775808
            v1 += v1
        }
        _ = v11
    } else {
        v6 = ("\n\t" + v6)
        v12 := GenS1{F0: -624}
        v3 += (v3 % (((v3 ^ v3)) | 1))
        _ = v12
    }
    if (0 < 0) {
        if true {
            if (1e10 < float32((v1 - float64((float32(v1) * 1e10))))) {
                var v13 uint32 = 4294967295
                var v14 int64 = 42
                _ = v13
                _ = v14
            }
            var v15 int64 = 0
            if (100 <= 127) {
                var v16 uint8 = 145
                _ = v16
            } else {
                var v17 uint16 = 0
                v4[2147483647] = ""
                if true {
                    if ((2 != 42) || (float32(((1e10 / 0.1) * float64(-2.25))) != 0.0)) {
                        for fi18 := 0; fi18 < 6; fi18++ {
                            v19 := GenS1{F0: (v15 + (v15 * (v15 - (v15 | v15))))}
                            switch fi18 {
                            case 2147483647:
                                v20 := GenS2{F0: v15, F1: (^(v15))}
                                _ = v20
                            }
                            for fi21 := 0; fi21 < 1; fi21++ {
                                delete(v5, 65535)
                                fmt.Println(v1, v1, fi21)
                                v17 = v17
                                _ = fi21
                            }
                            _ = fi18
                            _ = v19
                        }
                    }
                }
                _ = v17
            }
            _ = v15
        } else {
            v2 = append(v2, 4294967295)
            v3 = (v3 & v3)
            for fi22 := 0; fi22 < 8; fi22++ {
                for fi23 := 0; fi23 < 20; fi23++ {
                    v24 := GenS2{F0: int64(v3), F1: 797}
                    v25 := GenS3{F0: 115, F1: 9223372036854775807}
                    _ = fi23
                    _ = v24
                    _ = v25
                }
                _ = len(v2)
                var v26 int = fi22
                _ = fi22
                _ = v26
            }
        }
        var v27 uint16 = 42
        v1 += ((0.1 * float64(-2.25)) / (float64(1e10) - v1))
        _ = v27
    } else {
        var v28 int16 = -211
        v28 = v28
        _ = v28
    }
    var v29 int32 = (-(v3))
    for fi30 := 0; fi30 < 6; fi30++ {
        var v31 uint32 = 51
        _ = fi30
        _ = v31
    }
    v32 := v5[65535]
    for fi33 := 0; fi33 < 18; fi33++ {
        if true {
            v34 := GenS3{F0: -115, F1: -1}
            _ = v34
        }
        v35 := GenS1{F0: 42}
        _ = fi33
        _ = v35
    }
    v36 := GenS3{F0: 127, F1: 1}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v32
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
