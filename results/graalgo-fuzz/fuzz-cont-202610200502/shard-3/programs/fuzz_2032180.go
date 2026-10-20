package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 uint8 }

func main() {
    var v1 uint32 = 1
    var v2 int32 = 994
    v3 := []int{2147483647, -155}
    var v4 string = "foo bar"
    var v5 bool = (792 > 9223372036854775807)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v1, v1, v4)
    for fi6 := 0; fi6 < 8; fi6++ {
        v7 := GenS1{F0: 234, F1: 2, F2: 213}
        if len(v3) > 0 {
            v3[0] = (^(fi6))
        }
        v5 = false
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: -411, F1: 65535, F2: 2}
    switch 9223372036854775807 {
    case 298:
        for fi9 := 0; fi9 < 4; fi9++ {
            if (v5 || v5) {
                for fi10 := 0; fi10 < 4; fi10++ {
                    v11 := GenS1{F0: 486, F1: 385, F2: 0}
                    var v12 int32 = (v2 ^ (v2 % ((0) | 1)))
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                v13 := GenS1{F0: 42, F1: 0, F2: 127}
                _ = v13
            } else {
                if v5 {
                    for fi14 := 0; fi14 < 1; fi14++ {
                        v1 = 65535
                        var v15 bool = v5
                        _ = fi14
                        _ = v15
                    }
                }
            }
            v2 = (v2 % (((-(v2))) | 1))
            if (((!v5) && (true && ((v1 % ((v1) | 1)) != (v1 / ((v1) | 1))))) && true) {
                if len(v3) > 0 {
                    v3[0] = 128
                }
                var v16 uint16 = 402
                switch (^(fi9)) {
                case 1:
                    var v17 int16 = 1000
                    _ = v17
                case -1:
                    var v18 float64 = (0.0 / 3.14159)
                    _ = v18
                default:
                    var v19 uint8 = 211
                    _ = v19
                }
                _ = v16
            }
            _ = fi9
        }
    default:
        if (4294967295 < 247) {
            var v20 int64 = 0
            var v21 bool = (255 <= 42)
            if (true || (!true)) {
                v21 = v5
                fmt.Println(v20)
            } else {
                for fi22 := 0; fi22 < 20; fi22++ {
                    fmt.Println(v20, v5)
                    v2 = 45
                    var v23 uint64 = 1
                    _ = fi22
                    _ = v23
                }
            }
            _ = v20
            _ = v21
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if (1.5 > 3.14159) {
        var v24 uint32 = v1
        if len(v3) > 0 {
            v3[0] = 128
        }
        _ = v24
    } else {
        v5 = ((((0.0 * -2.25) / 1.0) + (0.0 - ((0.1 * (3.14159 * 0.0)) * float32(float64(1e10))))) > 1e10)
    }
    if (v4 <= v4) {
        fmt.Println(v5)
        for fi25 := 0; fi25 < 1; fi25++ {
            var v26 int = (fi25 % ((0) | 1))
            _ = fi25
            _ = v26
        }
    } else {
        for fi27 := 0; fi27 < 8; fi27++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi27
        }
    }
    if len(v3) > 0 {
        v3[0] = 127
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
