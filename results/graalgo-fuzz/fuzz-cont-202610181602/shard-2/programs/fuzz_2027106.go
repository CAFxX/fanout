package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 bool; F1 int; F2 int16 }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint64 = 4294967295
    var v3 int64 = -876
    v4 := make(map[string]int)
    var v5 float64 = float64(1.5)
    var v6 int = 819
    v3 = 68
    var v7 int32 = 0
    v6 = v6
    for fi8 := 0; fi8 < 1; fi8++ {
        var v9 uint16 = 770
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: (v6 - 4294967295), F1: 219}
    v5 -= v5
    v4["foo bar"] = (v6 / ((v6) | 1))
    v11 := []int{v6}
    if len(v11) > 0 {
        v11[0] = (v6 >> ((1) & 63))
    }
    v5 = v5
    fmt.Println(v5)
    v2 = (v2 << ((4294967295) & 63))
    for fi12 := 0; fi12 < 4; fi12++ {
        if false {
            if ((v3 & v3) > (v3 ^ v3)) {
                v13 := GenS2{F0: false, F1: v6, F2: 567}
                for fi14 := 0; fi14 < 14; fi14++ {
                    if len(v11) > 0 {
                        v11[0] = (v6 * (fi12 - fi14))
                    }
                    _ = fi14
                }
                _ = v13
            } else {
                if (!(42 >= 458)) {
                    var v15 bool = (405 >= 982)
                    _ = v15
                } else {
                    v7 += 174
                    var v16 string = ""
                    _ = v16
                }
                v5 = v5
            }
        }
        var v17 uint = 2
        _ = fi12
        _ = v17
    }
    fmt.Println(v2)
    for fi18 := 0; fi18 < 9; fi18++ {
        if false {
            v6 = 32767
            if (((v5 * (v5 * float64(float32(v5)))) - 0.0) <= (float64(((1e10 * float32(3.14159)) + float32((float64((0.0 * 0.0)) - v5)))) + ((-2.25 - v5) / ((-2.25 + v5) + v5)))) {
                switch fi18 {
                case -1:
                    switch v6 {
                    case 2147483647:
                        var v19 float32 = (((0.1 * ((0.1 - 0.0) - float32((v5 * (v5 + float64(float32(-2.25))))))) / 1.5) * 0.0)
                        _ = v19
                    case -9223372036854775808:
                        _ = len(v11)
                    default:
                        v3 = (v3 & v3)
                    }
                case 0:
                    if true {
                        var v20 int = (^(v6))
                        _ = v20
                    }
                case -292:
                    delete(v4, "\n\t")
                }
                var v21 uint8 = 255
                v7 = v7
                _ = v21
            }
            fmt.Println(v6)
        }
        v5 -= 0.0
        _ = fi18
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
