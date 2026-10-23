package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int64; F1 uint16 }
type GenS4 struct { F0 string }

func main() {
    v1 := []int{544}
    v2 := make(map[uint32]uint32)
    var v3 float64 = (float64((((-2.25 / 0.1) * 1.5) + 0.1)) / (0.1 + 1.5))
    var v4 int8 = 127
    v5 := []int{32768, 1, 473}
    var v6 uint16 = 255
    for fi7 := 0; fi7 < 5; fi7++ {
        fmt.Println(v6)
        _ = fi7
    }
    v8 := GenS1{F0: -1, F1: (v4 | v4)}
    if (1000 >= 1000) {
        v4 -= (v4 * -45)
        if true {
            v9 := GenS1{F0: 1000, F1: v4}
            if (128 >= 42) {
                var v10 float64 = (float64(float32((1e10 * ((1.5 * 0.1) / (float64((0.0 - float32((1.5 - (v3 + v3))))) * ((((v3 - (1e10 - float64((0.1 * 0.0)))) / v3) - 0.1) - 0.0)))))) * 0.1)
                v3 += 1e10
                if (-2147483648 <= -911) {
                    if true {
                        v11 := v2[2147483648]
                        _ = v11
                    } else {
                        v3 = ((float64(0.0) - float64(1e10)) - (v10 + 3.14159))
                    }
                    for fi12 := 0; fi12 < 19; fi12++ {
                        var v13 float32 = 1e10
                        _ = fi12
                        _ = v13
                    }
                } else {
                    if false {
                        var v14 uint64 = 4294967295
                        var v15 float64 = v3
                        _ = v14
                        _ = v15
                    }
                    for fi16 := 0; fi16 < 16; fi16++ {
                        var v17 int32 = 2147483647
                        if ((0 < 2147483648) && (!true)) {
                            if true {
                                switch (fi16 | (-(fi16))) {
                                case 107:
                                    v4 -= (^(v4))
                                case 4294967295:
                                    fmt.Println(v6, v6, v10)
                                case 0:
                                    fmt.Println(v3, v6, v6)
                                default:
                                    var v18 float32 = 3.14159
                                    _ = v18
                                }
                            } else {
                                v3 = v3
                            }
                        } else {
                            if false {
                                for fi19 := 0; fi19 < 4; fi19++ {
                                    switch fi16 {
                                    case 2147483647:
                                        v4 += (-(v4))
                                    case 32768:
                                        _ = len(v2)
                                    case 0:
                                        var v20 uint8 = 2
                                        _ = v20
                                    }
                                    _ = fi19
                                }
                                fmt.Println(v17, v4)
                                v21 := GenS2{F0: 0}
                                _ = v21
                            }
                            if (!(!((1.5 + float32((3.14159 / 0.1))) < 1e10))) {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                if (-912 != 1) {
                                    v17 = v17
                                    fmt.Println(v3)
                                    if (171 < 2) {
                                        v1 = append(v1, (fi16 | (fi16 & (fi16 & 30))))
                                    }
                                }
                            }
                            fmt.Println(v6)
                        }
                        _ = fi16
                        _ = v17
                    }
                    v22 := GenS3{F0: 0, F1: 1}
                    _ = v22
                }
                _ = v10
            } else {
                fmt.Println(v6, v3)
            }
            _ = v9
        }
        v23 := GenS4{F0: "hello"}
        _ = v23
    } else {
        fmt.Println(v4)
    }
    delete(v2, 1)
    for fi24 := 0; fi24 < 17; fi24++ {
        v6 = 2
        var v25 int = 1
        fmt.Println(v6, fi24)
        _ = fi24
        _ = v25
    }
    v3 -= (1.5 * -2.25)
    v3 -= ((((v3 / v3) / (v3 * v3)) / (v3 * v3)) + (v3 / v3))
    var v26 string = ""
    for fi27 := 0; fi27 < 12; fi27++ {
        fmt.Println(v26, v4)
        if (!true) {
            var v28 float32 = float32(3.14159)
            if (true && true) {
                v29 := GenS3{F0: 850, F1: 65535}
                _ = v29
            }
            if (4294967295 > 65535) {
                v30 := GenS4{F0: v26}
                var v31 int8 = 102
                v26 = v26
                _ = v30
                _ = v31
            }
            _ = v28
        }
        _ = fi27
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
