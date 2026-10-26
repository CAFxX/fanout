package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 uint64 }

func main() {
    var v1 bool = (222 < 352)
    v2 := make(map[string]int)
    v3 := make(map[string]string)
    v4 := []uint32{62, 42}
    v1 = true
    var v5 int8 = 42
    v1 = (false && (-2.25 != 0.1))
    var v6 int64 = 42
    v1 = (-914 < 42)
    var v7 bool = (42 != 31)
    for fi8 := 0; fi8 < 9; fi8++ {
        var v9 bool = (42 != 487)
        var v10 int16 = -716
        for fi11 := 0; fi11 < 20; fi11++ {
            switch 934 {
            case 471:
                if v9 {
                    switch (fi8 * -9223372036854775808) {
                    case -1:
                        if v7 {
                            var v12 bool = (((!(0.1 <= float32(((float64(float32(-2.25)) + (1.5 - (-2.25 / 0.1))) * (1e10 * 3.14159))))) && v7) || ((v10 * 1) > (v10 | (v10 * (v10 ^ 1)))))
                            _ = v12
                        } else {
                            fmt.Println(fi8)
                            var v13 float64 = (0.0 / 1.0)
                            _ = v13
                        }
                    case 2147483647:
                        switch fi11 {
                        case 1:
                            if v9 {
                                var v14 uint32 = 107
                                _ = v14
                            }
                        case 128:
                            v15 := v2["a"]
                            _ = v15
                        default:
                            var v16 string = "\n\t"
                            _ = v16
                        }
                    }
                    v5 -= int8(v6)
                    var v17 int = fi8
                    _ = v17
                } else {
                    delete(v2, (("" + "\n\t") + "hello"))
                    _ = len(v3)
                    if (v1 || (!false)) {
                        v10 += (v10 << ((18446744073709551615) & 63))
                    }
                    if v1 {
                        if (1 >= 2) {
                            v7 = (7 >= (v6 * (v6 / ((-1) | 1))))
                            var v18 float64 = -2.25
                            _ = v18
                        }
                        v19 := GenS1{F0: 4294967295, F1: 42, F2: 532}
                        v6 -= v6
                        _ = v19
                    }
                    v9 = true
                }
            }
            var v20 uint64 = 970
            if (2 == 4294967295) {
                var v21 uint16 = 2
                var v22 int16 = (v10 << ((468) & 63))
                _ = v21
                _ = v22
            } else {
                v4 = append(v4, 2147483647)
            }
            _ = fi11
            _ = v20
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    v1 = v1
    v7 = false
    var v23 int32 = -628
    v5 = 62
    v24 := GenS1{F0: 617, F1: 255, F2: 42}
    var v25 uint = 65535
    var v26 int = -2147483648
    v27 := v2["a"]
    fmt.Println(v27, v7, v23)
    var v28 uint16 = 255
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
