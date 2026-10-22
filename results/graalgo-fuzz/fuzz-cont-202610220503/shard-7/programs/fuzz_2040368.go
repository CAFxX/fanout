package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string; F1 int }

func main() {
    var v1 float64 = 3.14159
    var v2 uint8 = 128
    var v3 float32 = float32(float64(0.0))
    v4 := make(map[uint32]int)
    var v5 string = "foo bar"
    v1 = 1e10
    var v6 uint32 = 2147483647
    if (false || false) {
        var v7 string = v5
        _ = v7
    }
    _ = len(v4)
    v8 := GenS1{F0: (4294967295 != 579), F1: 2, F2: -2147483648}
    if true {
        if (false || (42 >= 1)) {
            fmt.Println(v2, v1)
            v4[v6] = 61
        } else {
            var v9 float32 = float32(v1)
            v9 = v9
            v9 = 1e10
            _ = v9
        }
        for fi10 := 0; fi10 < 20; fi10++ {
            v11 := GenS1{F0: (42 >= 709), F1: 0, F2: fi10}
            fmt.Println(fi10, v5, fi10)
            _ = fi10
            _ = v11
        }
    }
    v2 = (v2 * v2)
    v12 := GenS2{F0: 127}
    v1 = v1
    v2 -= 128
    for fi13 := 0; fi13 < 11; fi13++ {
        switch (v5 + (v5 + v5)) {
        case "a":
            if (false || false) {
                v14 := v4[819]
                _ = v14
            } else {
                v2 += (v2 / (((v2 % ((128) | 1))) | 1))
                var v15 string = v5
                for fi16 := 0; fi16 < 13; fi16++ {
                    v17 := []int{(fi16 % ((fi13) | 1)), (fi16 << ((454) & 63))}
                    _ = fi16
                    _ = v17
                }
                _ = v15
            }
        case "hello":
            v18 := v4[4294967295]
            _ = v18
        }
        v6 = 4294967295
        v1 += v1
        _ = fi13
    }
    var v19 int32 = -263
    if false {
        v20 := GenS1{F0: (v3 < float32((float64(v3) - (float64((1.5 + v3)) + (v1 - (v1 / v1)))))), F1: 255, F2: 1}
        if false {
            var v21 uint16 = 2
            v21 += uint16(v6)
            switch -1 {
            case 127:
                if false {
                    delete(v4, (^(v6)))
                    v4[v6] = 128
                    fmt.Println(v6, v6)
                } else {
                    v5 = ("\n\t" + "foo bar")
                }
            case 2147483647:
                v4[(v6 & v6)] = 0
            case -420:
                v6 += v6
            default:
                var v22 float32 = v3
                _ = v22
            }
            _ = v21
        } else {
            for fi23 := 0; fi23 < 1; fi23++ {
                delete(v4, 417)
                for fi24 := 0; fi24 < 9; fi24++ {
                    v1 = float64(((1e10 * ((v3 - ((float32(v1) * ((v3 - (v3 * (((float32(0.0) + (((0.0 / 0.1) / float32(v1)) / v3)) + 1e10) / 3.14159))) + 3.14159)) + v3)) * (v3 + v3))) * 0.1))
                    v25 := GenS1{F0: ((v2 ^ uint8(v6)) == 83), F1: 92, F2: 697}
                    fmt.Println(fi23)
                    if false {
                        if true {
                            v26 := GenS1{F0: ((v1 < v1) || (((((float64(v3) / (float64((v3 / (v3 - 1e10))) - (v1 - v1))) + 0.1) * (((float64(v3) + (0.0 - (3.14159 - v1))) * (v1 / float64(v3))) * (1.5 * 1e10))) > float64(((v3 + v3) - v3))) || (true || (1 == 1)))), F1: 0, F2: fi24}
                            if ((v19 - v19) < (v19 / (((v19 - (v19 & 42))) | 1))) {
                                v1 += -2.25
                                var v27 int8 = -56
                                fmt.Println(v27, v27)
                                _ = v27
                            }
                            if true {
                                v28 := GenS3{F0: (v5 + v5), F1: fi23}
                                delete(v4, v6)
                                _ = v28
                            } else {
                                fmt.Println(v6, fi23, fi23)
                            }
                            _ = v26
                        }
                    }
                    fmt.Println(v19, fi23, fi23)
                    _ = fi24
                    _ = v25
                }
                _ = fi23
            }
            v4[2147483648] = 628
            fmt.Println(v3)
        }
        _ = v20
    } else {
        v3 = 1.5
        var v29 int32 = v19
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
