package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16; F1 int; F2 int16 }
type GenS3 struct { F0 int8 }

func main() {
    var v1 int8 = 42
    var v2 uint16 = 255
    var v3 float32 = ((float32((float64(((((float32((0.1 + (float64(3.14159) / 1.0))) / (0.1 - (float32(((0.1 + (1e10 * (1.5 / ((0.0 / 0.1) + 0.1)))) * float64((1.5 * 0.1)))) + 0.1))) + float32(float64(-2.25))) - (0.1 * float32(float64(float32(1e10))))) / 0.1)) * 1.5)) / -2.25) / 1e10)
    var v4 string = ("\n\t" + "hello")
    v5 := []int{-116, 32768, 9223372036854775807}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v5 = append(v5, 42)
    for fi6 := 0; fi6 < 13; fi6++ {
        v3 = (((v3 + float32(0.0)) - 3.14159) + float32(3.14159))
        switch v4 {
        case "a":
            var v7 uint16 = (v2 ^ (v2 >> ((1) & 63)))
            _ = v7
        case "hello":
            v3 = v3
        }
        _ = fi6
    }
    v5 = append(v5, -771)
    if false {
        if true {
            if true {
                var v8 bool = (false && true)
                v9 := GenS1{F0: 32767}
                _ = v8
                _ = v9
            } else {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = len(v5)
                var v10 int16 = 42
                _ = v10
            }
            v2 += (^(v2))
            if ((v4 + v4) <= "foo bar") {
                v4 = (v4 + (v4 + v4))
                v11 := GenS1{F0: 1000}
                _ = v11
            }
        } else {
            for fi12 := 0; fi12 < 11; fi12++ {
                for fi13 := 0; fi13 < 14; fi13++ {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = fi13
                }
                fmt.Println(v4, v1, v3)
                _ = fi12
            }
            fmt.Println(v3, v3, v4)
        }
    }
    var v14 uint32 = 65535
    if false {
        switch (v4 + "hello") {
        case "foo bar":
            var v15 uint32 = v14
            _ = v15
        case "hello":
            fmt.Println(v2)
        default:
            if true {
                switch -543 {
                case 127:
                    fmt.Println(v4, v14, v3)
                case 9223372036854775807:
                    if true {
                        switch v4 {
                        case "hello":
                            switch v4 {
                            case "hello":
                                var v16 float64 = (1e10 / ((float64((0.1 * 1e10)) * 1e10) / 3.14159))
                                _ = v16
                            default:
                                v4 = v4
                            }
                        default:
                            if (false || (0 <= 659)) {
                                if false {
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    var v17 float64 = -2.25
                                    _ = v17
                                } else {
                                    fmt.Println(v2)
                                    v18 := GenS2{F0: 401, F1: 1, F2: 42}
                                    _ = v18
                                }
                                if true {
                                    v2 += (^(v2))
                                    var v19 int8 = (v1 & v1)
                                    fmt.Println(v1, v4)
                                    _ = v19
                                } else {
                                    v20 := GenS3{F0: -69}
                                    _ = v20
                                }
                            } else {
                                fmt.Println(v4)
                            }
                        }
                    } else {
                        v3 += (float32(0.0) / (float32(1e10) + v3))
                        _ = len(v5)
                    }
                case 525:
                    v21 := GenS1{F0: 1}
                    _ = v21
                default:
                    v1 = v1
                }
            }
        }
        if false {
            v14 = v14
            v5 = append(v5, 32768)
            v2 -= (v2 >> ((920) & 63))
        }
        v3 = v3
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
