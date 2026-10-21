package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 string; F1 uint64 }
type GenS4 struct { F0 uint16 }

func main() {
    var v1 int16 = -513
    var v2 float64 = (-2.25 + 3.14159)
    var v3 int8 = 72
    var v4 string = "hello"
    if ((v1 * -32768) < -1) {
        v5 := []int{127, 32767, 32768}
        for fi6 := 0; fi6 < 3; fi6++ {
            var v7 int = (fi6 * fi6)
            fmt.Println(v2, fi6, fi6)
            _ = fi6
            _ = v7
        }
        _ = v5
    }
    v8 := GenS1{F0: 300, F1: 1}
    v2 = (v2 / v2)
    v2 -= float64(1e10)
    var v9 int32 = -475
    v10 := []int{32767, 32768}
    var v11 float32 = 0.0
    var v12 uint64 = 55
    var v13 float64 = 1.5
    var v14 bool = true
    v10 = append(v10, 1)
    for fi15 := 0; fi15 < 15; fi15++ {
        fmt.Println(v1, v11, v3)
        for fi16 := 0; fi16 < 8; fi16++ {
            var v17 bool = (2 <= 127)
            var v18 int64 = 9223372036854775807
            switch 522 {
            case -2147483648:
                var v19 int = fi15
                _ = v19
            default:
                var v20 float32 = 1.5
                _ = v20
            }
            var v21 uint = 1
            v22 := GenS2{F0: true, F1: 127}
            _ = fi16
            _ = v17
            _ = v18
            _ = v21
            _ = v22
        }
        switch (fi15 - fi15) {
        case 32768:
            if (v14 && (2 >= 265)) {
                var v23 int = -56
                if (float64(((((0.0 * v11) + v11) * (v11 / float32(v2))) * v11)) < 3.14159) {
                    for fi24 := 0; fi24 < 7; fi24++ {
                        v25 := GenS3{F0: v4, F1: (v12 & v12)}
                        v10 = append(v10, (v23 << ((1) & 63)))
                        _ = fi24
                        _ = v25
                    }
                    _ = len(v10)
                    if v14 {
                        for fi26 := 0; fi26 < 3; fi26++ {
                            if (!false) {
                                v23 = 2147483647
                            }
                            var v27 int32 = v9
                            var v28 uint8 = 72
                            if (18446744073709551615 < (^(v12))) {
                                var v29 int64 = 1
                                _ = v29
                            }
                            if v14 {
                                fmt.Println(v4, v27)
                            }
                            _ = fi26
                            _ = v27
                            _ = v28
                        }
                    }
                } else {
                    fmt.Println(v23, fi15, v23)
                    v23 = v23
                    v12 = v12
                }
                switch v4 {
                case "foo bar":
                    v30 := GenS4{F0: 2}
                    _ = v30
                case "a":
                    for fi31 := 0; fi31 < 7; fi31++ {
                        v32 := GenS2{F0: false, F1: 42}
                        _ = fi31
                        _ = v32
                    }
                }
                _ = v23
            }
        case 672:
            v12 = v12
        default:
            var v33 uint8 = 0
            _ = v33
        }
        _ = fi15
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
