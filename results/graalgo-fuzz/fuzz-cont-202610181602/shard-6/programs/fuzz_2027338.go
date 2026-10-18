package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 string }
type GenS3 struct { F0 int64; F1 uint16 }

func main() {
    var v1 uint = 0
    v2 := make(map[string]string)
    var v3 uint8 = uint8(v1)
    v4 := []uint32{2147483648, 768}
    var v5 string = "foo bar"
    fmt.Println(v5, v5)
    fmt.Println(v3)
    if true {
        v6 := GenS1{F0: 2}
        _ = v6
    }
    fmt.Println(v3, v1)
    for fi7 := 0; fi7 < 10; fi7++ {
        if false {
            v2[v5] = v5
            v1 += (v1 >> ((v1) & 63))
            if true {
                fmt.Println(v1, v3)
            } else {
                if ("\n\t" > ("hello" + v5)) {
                    v3 += v3
                    if true {
                        switch (fi7 - 2147483647) {
                        case 128:
                            fmt.Println(v3, v3)
                        }
                    }
                }
            }
        } else {
            for fi8 := 0; fi8 < 5; fi8++ {
                if (v1 == 0) {
                    if ((9223372036854775807 > -9223372036854775808) || false) {
                        v3 -= v3
                    } else {
                        var v9 string = v5
                        if false {
                            v10 := v2["a"]
                            var v11 float64 = (float64(0.1) + 1.5)
                            var v12 int32 = -378
                            var v13 uint64 = 18446744073709551615
                            v11 += (v11 * (1.5 + v11))
                            _ = v10
                            _ = v11
                            _ = v12
                            _ = v13
                        } else {
                            var v14 bool = false
                            v15 := GenS2{F0: v5, F1: (v9 + v5)}
                            v16 := GenS1{F0: 2}
                            _ = v14
                            _ = v15
                            _ = v16
                        }
                        _ = v9
                    }
                } else {
                    var v17 float32 = (1.5 - 1e10)
                    var v18 bool = false
                    _ = v17
                    _ = v18
                }
                v1 += (v1 ^ (v1 & (v1 / ((2) | 1))))
                switch (fi8 * (fi8 ^ fi8)) {
                case -1:
                    _ = len(v2)
                    var v19 uint = v1
                    _ = v19
                case -435:
                    v1 += (^(v1))
                default:
                    v20 := GenS1{F0: 2147483647}
                    _ = v20
                }
                _ = fi8
            }
            switch (fi7 * fi7) {
            case -218:
                v21 := GenS3{F0: 161, F1: 504}
                _ = v21
            case 0:
                _ = len(v4)
            default:
                var v22 int16 = -87
                _ = v22
            }
        }
        _ = len(v4)
        _ = fi7
    }
    var v23 float32 = ((-2.25 / (float32(3.14159) - (float32(1.5) + (((3.14159 * 1.5) - -2.25) * ((0.0 + (1.5 + 0.1)) * -2.25))))) * 0.0)
    for fi24 := 0; fi24 < 2; fi24++ {
        if false {
            v25 := GenS1{F0: 864}
            if len(v4) > 0 {
                v4[0] = 717
            }
            if len(v4) > 0 {
                v4[0] = 197
            }
            _ = v25
        }
        _ = fi24
    }
    v23 = (float32((0.0 + 3.14159)) * (v23 + v23))
    var v26 uint64 = 728
    var v27 string = v5
    _ = len(v2)
    var v28 int64 = 1
    _ = len(v2)
    var v29 uint64 = v26
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
