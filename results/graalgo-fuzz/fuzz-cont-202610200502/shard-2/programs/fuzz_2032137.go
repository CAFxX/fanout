package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 int8 }
type GenS2 struct { F0 int8; F1 int8; F2 uint8 }

func main() {
    var v1 uint = 65535
    v2 := make(map[uint32]bool)
    var v3 uint16 = 65535
    var v4 int64 = -1
    v5 := []int{810}
    for fi6 := 0; fi6 < 5; fi6++ {
        v3 = 305
        for fi7 := 0; fi7 < 2; fi7++ {
            if (v3 >= (^(v3))) {
                var v8 uint32 = 856
                v9 := v2[(v8 ^ v8)]
                _ = v8
                _ = v9
            }
            var v10 int64 = (^(v4))
            _ = fi7
            _ = v10
        }
        _ = fi6
    }
    v2[42] = (((100 <= -115) && ((true && true) || ("hello" < "\n\t"))) && (-1 < int8(v4)))
    if (!true) {
        for fi11 := 0; fi11 < 11; fi11++ {
            v1 += 259
            fmt.Println(v1, v1)
            for fi12 := 0; fi12 < 2; fi12++ {
                v1 -= (v1 % (((v1 >> (((^(v1))) & 63))) | 1))
                _ = fi12
            }
            _ = fi11
        }
        if true {
            var v13 float32 = ((0.1 - (-2.25 + 0.1)) - -2.25)
            _ = v13
        } else {
            var v14 float32 = 1.5
            _ = v14
        }
    } else {
        for fi15 := 0; fi15 < 4; fi15++ {
            if true {
                fmt.Println(fi15)
                for fi16 := 0; fi16 < 13; fi16++ {
                    if (!true) {
                        var v17 int = -1
                        fmt.Println(v1, v1)
                        var v18 float64 = -2.25
                        _ = v17
                        _ = v18
                    } else {
                        var v19 uint16 = 250
                        fmt.Println(fi15)
                        _ = v19
                    }
                    var v20 uint8 = 207
                    v21 := GenS1{F0: v3, F1: -20, F2: 100}
                    _ = fi16
                    _ = v20
                    _ = v21
                }
                v3 = v3
            } else {
                v1 -= (^(v1))
                var v22 uint16 = (^(v3))
                v23 := GenS2{F0: 0, F1: 0, F2: 169}
                _ = v22
                _ = v23
            }
            var v24 uint8 = 240
            v25 := GenS1{F0: v3, F1: 1, F2: 76}
            _ = fi15
            _ = v24
            _ = v25
        }
        fmt.Println(v1, v4)
        var v26 int64 = (-(v4))
        _ = v26
    }
    v5 = append(v5, 42)
    var v27 bool = (!(65535 == 753))
    var v28 int16 = 42
    var v29 bool = (!false)
    var v30 uint = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
