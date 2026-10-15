package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 uint16 }
type GenS2 struct { F0 bool; F1 int16; F2 int32 }

func main() {
    v1 := make(map[string]bool)
    var v2 float32 = float32(-2.25)
    var v3 int = -9223372036854775808
    var v4 float32 = (v2 * float32(float64((((((1e10 / 1.0) * (0.0 - v2)) / 1.5) - (v2 + v2)) + (((1e10 * (v2 - v2)) + 3.14159) + v2)))))
    v5 := make(map[int]int)
    var v6 uint32 = 0
    switch 879 {
    case -201:
        var v7 int32 = -723
        _ = v7
    default:
        v3 -= v3
    }
    if ((255 >= 42) && (-360 != 0)) {
        for fi8 := 0; fi8 < 5; fi8++ {
            fmt.Println(v3, v6)
            if (1.5 == v4) {
                v9 := v1["foo bar"]
                _ = v9
            }
            _ = fi8
        }
    }
    var v10 float32 = v4
    v4 = 1e10
    var v11 int32 = -725
    v12 := []int{(v3 + v3)}
    v6 += (v6 >> ((1) & 63))
    if false {
        var v13 uint = 2
        for fi14 := 0; fi14 < 3; fi14++ {
            for fi15 := 0; fi15 < 3; fi15++ {
                fmt.Println(v4)
                _ = fi15
            }
            var v16 string = ""
            _ = fi14
            _ = v16
        }
        _ = v13
    }
    switch v3 {
    case 42:
        v6 -= v6
    default:
        for fi17 := 0; fi17 < 19; fi17++ {
            v3 -= (^(fi17))
            for fi18 := 0; fi18 < 5; fi18++ {
                v19 := GenS1{F0: 65535, F1: 42, F2: 921}
                _ = fi18
                _ = v19
            }
            v20 := GenS1{F0: 32767, F1: 65535, F2: 2}
            _ = fi17
            _ = v20
        }
    }
    v5[v3] = (v3 | v3)
    v3 = (v3 >> ((42) & 63))
    switch (v3 - (-(v3))) {
    case 127:
        var v21 uint8 = 128
        _ = v21
    }
    _ = len(v1)
    var v22 int32 = (v11 | (v11 - (-(v11))))
    v23 := GenS2{F0: (-279 <= v11), F1: 709, F2: int32(v6)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
