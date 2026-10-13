package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16; F1 bool; F2 int32 }

func main() {
    v1 := make(map[string]string)
    var v2 float64 = (float64((0.1 - float32(-2.25))) * 1.5)
    v3 := make(map[uint32]uint32)
    v4 := make(map[string]string)
    v5 := GenS1{F0: -154}
    v2 -= (float64(3.14159) + (((0.1 - (((v2 + ((((v2 + v2) - v2) + v2) * 3.14159)) / v2) + (float64(float32(v2)) - 1.5))) / (0.0 / v2)) / (1.5 / 0.1)))
    var v6 float64 = v2
    v7 := GenS2{F0: 1, F1: (((-2.25 + (1e10 / -2.25)) * ((3.14159 - 0.1) / (float32((1.5 + (v6 - 3.14159))) * 3.14159))) == 3.14159), F2: 0}
    v8 := GenS1{F0: 0}
    v9 := GenS1{F0: 263}
    fmt.Println(v2)
    fmt.Println(v2)
    if true {
        if false {
            var v10 uint = 1
            _ = v10
        } else {
            v6 = float64(-2.25)
            switch -9223372036854775808 {
            case -2147483648:
                if true {
                    fmt.Println(v6)
                }
            }
            for fi11 := 0; fi11 < 4; fi11++ {
                for fi12 := 0; fi12 < 20; fi12++ {
                    v6 = (((v2 + 3.14159) / float64(0.1)) * (v2 - float64((0.0 * (float32(v6) / 1.0)))))
                    v2 -= 1.5
                    _ = fi12
                }
                switch "\n\t" {
                case "\n\t":
                    switch "hello" {
                    case "foo bar":
                        if true {
                            for fi13 := 0; fi13 < 19; fi13++ {
                                v14 := GenS2{F0: 420, F1: false, F2: 0}
                                _ = fi13
                                _ = v14
                            }
                        } else {
                            for fi15 := 0; fi15 < 17; fi15++ {
                                v6 += v2
                                v2 -= (((-2.25 * 0.0) / 1.5) + 3.14159)
                                v16 := GenS2{F0: -316, F1: true, F2: -1}
                                var v17 uint16 = 650
                                fmt.Println(v2)
                                v18 := GenS1{F0: (fi11 << ((1) & 63))}
                                v2 += -2.25
                                v19 := GenS2{F0: 42, F1: false, F2: 2147483647}
                                v20 := v3[2147483647]
                                if ((1 != 415) && (1 <= 2147483647)) {
                                    v2 = -2.25
                                }
                                _ = fi15
                                _ = v16
                                _ = v17
                                _ = v18
                                _ = v19
                                _ = v20
                            }
                        }
                    case "hello":
                        var v21 uint64 = 4294967295
                        _ = v21
                    }
                case "a":
                    fmt.Println(v6)
                }
                _ = fi11
            }
        }
        v22 := []int{4294967295, 127, -377}
        _ = v22
    } else {
        for fi23 := 0; fi23 < 17; fi23++ {
            var v24 int8 = 100
            var v25 int16 = 0
            _ = fi23
            _ = v24
            _ = v25
        }
        if false {
            if true {
                v2 += v6
                for fi26 := 0; fi26 < 20; fi26++ {
                    for fi27 := 0; fi27 < 2; fi27++ {
                        var v28 int64 = 42
                        fmt.Println(fi26, fi27, fi26)
                        _ = fi27
                        _ = v28
                    }
                    _ = fi26
                }
            } else {
                fmt.Println(v6, v6)
                var v29 float32 = -2.25
                _ = v29
            }
        } else {
            v6 += v6
        }
    }
    _ = len(v1)
    for fi30 := 0; fi30 < 2; fi30++ {
        v2 = (1.5 / 1e10)
        v31 := GenS1{F0: (fi30 ^ (fi30 % (((fi30 / ((32768) | 1))) | 1)))}
        _ = fi30
        _ = v31
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
