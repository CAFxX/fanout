package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 int16 }
type GenS2 struct { F0 bool; F1 bool; F2 int8 }
type GenS3 struct { F0 uint; F1 int16 }

func main() {
    var v1 int16 = -868
    v2 := make(map[string]int)
    v3 := make(map[uint32]uint32)
    v4 := []int{0, 140, 32768, 2147483647}
    for fi5 := 0; fi5 < 9; fi5++ {
        for fi6 := 0; fi6 < 4; fi6++ {
            var v7 int16 = (v1 | (v1 & (-(v1))))
            v8 := GenS1{F0: ("hello" + "a"), F1: 42, F2: 934}
            var v9 uint32 = 2147483648
            _ = fi6
            _ = v7
            _ = v8
            _ = v9
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi5
    }
    var v10 float32 = (3.14159 * 0.1)
    v11 := GenS2{F0: false, F1: false, F2: -55}
    for fi12 := 0; fi12 < 9; fi12++ {
        var v13 uint32 = 577
        _ = fi12
        _ = v13
    }
    fmt.Println(v10)
    v14 := GenS2{F0: false, F1: true, F2: -7}
    delete(v3, 2147483648)
    fmt.Println(v10)
    if (1 >= 465) {
        v10 -= v10
    } else {
        var v15 int64 = -16
        if false {
            fmt.Println(v15, v1, v1)
            var v16 uint32 = 4294967295
            for fi17 := 0; fi17 < 8; fi17++ {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi17
            }
            _ = v16
        } else {
            if true {
                var v18 int32 = -644
                switch 32768 {
                case 205:
                    if false {
                        v19 := GenS1{F0: "foo bar", F1: -1, F2: (v1 / ((v1) | 1))}
                        _ = v19
                    }
                case 0:
                    var v20 uint = 1
                    _ = v20
                case -223:
                    v1 = 0
                default:
                    v21 := GenS2{F0: ((-(v18)) <= v18), F1: (!false), F2: 1}
                    _ = v21
                }
                v22 := GenS3{F0: 18446744073709551615, F1: v1}
                _ = v18
                _ = v22
            } else {
                var v23 uint16 = 867
                switch 2147483647 {
                case -9223372036854775808:
                    v10 = v10
                default:
                    for fi24 := 0; fi24 < 10; fi24++ {
                        var v25 float64 = float64(v10)
                        _ = fi24
                        _ = v25
                    }
                }
                v10 += ((float32(3.14159) - v10) * (v10 + ((3.14159 + v10) * 0.0)))
                _ = v23
            }
            fmt.Println(v15, v1)
        }
        _ = v15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
