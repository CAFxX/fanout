package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 uint64) uint16 {
    for fi6 := 0; fi6 < 11; fi6++ {
        v7 := []int{fi6, -389}
        _ = fi6
        _ = v7
    }
    var v8 float64 = float64(-2.25)
    var v9 uint64 = (^(p0))
    _ = p0
    _ = v8
    _ = v9
    return 32767
}
func genFn1() int32 {
    var v10 float64 = 1.5
    var v11 bool = true
    v12 := []int{32768, -262, 0, 985}
    _ = v10
    _ = v11
    _ = v12
    return -2147483648
}

func main() {
    v1 := []uint32{631}
    v2 := make(map[int]bool)
    v3 := make(map[uint32]int)
    var v4 int = -721
    v5 := make(map[string]int)
    _ = len(v1)
    var v13 int8 = 1
    for fi14 := 0; fi14 < 11; fi14++ {
        v1 = append(v1, 2147483648)
        _ = fi14
    }
    v4 = (v4 ^ v4)
    if false {
        fmt.Println(v13, v13)
        v2[1] = ((1e10 + 3.14159) <= (3.14159 + float64(float32(0.0))))
        var v15 uint64 = 0
        _ = v15
    } else {
        var v16 uint8 = 1
        v16 += v16
        v17 := GenS1{F0: 4294967295}
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 14; fi18++ {
        v19 := GenS1{F0: 0}
        var v20 int64 = 503
        for fi21 := 0; fi21 < 13; fi21++ {
            v22 := GenS1{F0: 255}
            delete(v2, (fi18 / (((fi18 * 4294967295)) | 1)))
            if true {
                v20 = 1
            }
            _ = fi21
            _ = v22
        }
        _ = fi18
        _ = v19
        _ = v20
    }
    fmt.Println(v4)
    var v23 uint64 = 237
    fmt.Println(v4)
    for fi24 := 0; fi24 < 7; fi24++ {
        _ = len(v5)
        var v25 int16 = 1
        _ = fi24
        _ = v25
    }
    for fi26 := 0; fi26 < 6; fi26++ {
        v13 += (v13 * v13)
        v27 := v2[(v4 >> ((1) & 63))]
        _ = fi26
        _ = v27
    }
    switch (v4 & (v4 & 42)) {
    case -9223372036854775808:
        v4 += 1
    case -635:
        v28 := GenS1{F0: 18446744073709551615}
        _ = v28
    default:
        for fi29 := 0; fi29 < 9; fi29++ {
            if false {
                if false {
                    v13 += (v13 ^ 0)
                    for fi30 := 0; fi30 < 10; fi30++ {
                        if false {
                            v31 := GenS1{F0: 4294967295}
                            var v32 int64 = -9223372036854775808
                            _ = v31
                            _ = v32
                        }
                        var v33 uint8 = 0
                        _ = fi30
                        _ = v33
                    }
                    if ((!(-274 >= 1)) || false) {
                        v1 = append(v1, 42)
                    }
                }
            }
            for fi34 := 0; fi34 < 3; fi34++ {
                v13 += v13
                _ = fi34
            }
            _ = fi29
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
