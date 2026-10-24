package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 int64; F2 uint }
type GenS3 struct { F0 int; F1 int32; F2 uint16 }

func main() {
    var v1 uint = 65535
    v2 := []int{}
    var v3 string = "a"
    var v4 int16 = -669
    v5 := []int{}
    var v6 float64 = ((0.0 * (float64(1.5) + 3.14159)) + -2.25)
    v7 := GenS1{F0: 1}
    v1 -= v1
    var v8 float32 = ((((1e10 * 0.1) / 3.14159) - 3.14159) - 0.0)
    if len(v5) > 0 {
        v5[0] = 1
    }
    if ((^(v1)) <= (v1 & v1)) {
        var v9 string = "hello"
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = v9
    }
    var v10 int = -548
    if (100 < -1) {
        v11 := GenS1{F0: 42}
        _ = v11
    }
    _ = len(v2)
    var v12 uint32 = 4294967295
    v2 = append(v2, v10)
    for fi13 := 0; fi13 < 13; fi13++ {
        if len(v2) > 0 {
            v2[0] = (-(fi13))
        }
        var v14 uint = (v1 / ((v1) | 1))
        if false {
            if false {
                for fi15 := 0; fi15 < 20; fi15++ {
                    v12 = v12
                    var v16 uint16 = 2
                    _ = fi15
                    _ = v16
                }
            } else {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                var v17 uint8 = 222
                _ = v17
            }
            v8 += 3.14159
        }
        _ = fi13
        _ = v14
    }
    v18 := GenS1{F0: 2147483648}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v19 float32 = v8
    if (!true) {
        v8 = float32((3.14159 + -2.25))
        if (-2147483648 != -2147483648) {
            switch (v10 | v10) {
            case 1:
                var v20 float32 = v8
                _ = v20
            case 204:
                v5 = append(v5, v10)
            case 32768:
                switch (-(v10)) {
                case 4294967295:
                    var v21 int32 = -2147483648
                    _ = v21
                case 32768:
                    if (1000 != (-(v4))) {
                        v1 = (^(v1))
                        v1 = (v1 / ((v1) | 1))
                        v22 := GenS1{F0: (v12 * v12)}
                        _ = v22
                    } else {
                        for fi23 := 0; fi23 < 1; fi23++ {
                            v24 := GenS2{F0: -1, F1: -386, F2: (v1 | 1)}
                            fmt.Println(v12)
                            _ = fi23
                            _ = v24
                        }
                    }
                default:
                    var v25 uint16 = 264
                    _ = v25
                }
            default:
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
            var v26 int = (-(v10))
            _ = v26
        }
    }
    v27 := GenS3{F0: (-(v10)), F1: -614, F2: 255}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
