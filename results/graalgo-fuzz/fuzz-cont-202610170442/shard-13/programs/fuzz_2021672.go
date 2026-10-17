package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 uint8 }
type GenS2 struct { F0 int64 }

func main() {
    var v1 int16 = 983
    var v2 float64 = ((0.0 + 1.5) + 0.0)
    var v3 int32 = 0
    v4 := []int{}
    var v5 uint = 504
    v6 := GenS1{F0: 32767, F1: -1, F2: 0}
    v5 = (v5 * (^(v5)))
    var v7 float64 = v2
    v2 += ((float64(-2.25) - (v7 + (v7 / (v2 - 3.14159)))) * (v2 + (((v2 - (v2 - v7)) / ((((1.5 * v2) - v7) + ((0.1 / (v7 - float64((0.1 + 0.0)))) / v7)) + 1e10)) / (v7 / v2))))
    if true {
        v8 := GenS1{F0: 135, F1: 42, F2: 42}
        _ = v8
    }
    var v9 uint32 = 4294967295
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi10 := 0; fi10 < 20; fi10++ {
        if true {
            v7 = v7
        } else {
            fmt.Println(v3, fi10, fi10)
            switch fi10 {
            case 0:
                for fi11 := 0; fi11 < 10; fi11++ {
                    v12 := GenS1{F0: 0, F1: -9223372036854775808, F2: 0}
                    v13 := GenS2{F0: -1}
                    _ = fi11
                    _ = v12
                    _ = v13
                }
            case 440:
                for fi14 := 0; fi14 < 14; fi14++ {
                    if len(v4) > 0 {
                        v4[0] = (fi14 << ((821) & 63))
                    }
                    v5 += v5
                    var v15 float32 = 0.0
                    _ = fi14
                    _ = v15
                }
            case 4294967295:
                v1 += -1
            }
            v3 = (v3 % (((^(v3))) | 1))
        }
        _ = fi10
    }
    if true {
        var v16 uint64 = 582
        v17 := GenS2{F0: 1}
        _ = v16
        _ = v17
    } else {
        for fi18 := 0; fi18 < 4; fi18++ {
            v19 := GenS1{F0: 596, F1: 479, F2: 0}
            _ = fi18
            _ = v19
        }
        for fi20 := 0; fi20 < 17; fi20++ {
            var v21 uint = v5
            switch 4294967295 {
            case 182:
                fmt.Println(v3, v1, v7)
            case 0:
                v9 = (^(v9))
            }
            _ = fi20
            _ = v21
        }
    }
    if len(v4) > 0 {
        v4[0] = 290
    }
    if len(v4) > 0 {
        v4[0] = 32768
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
