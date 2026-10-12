package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 bool }

func main() {
    var v1 uint64 = 0
    var v2 int16 = 72
    var v3 int64 = -1
    var v4 float64 = 1.5
    var v5 int32 = -1
    v6 := []int{32768, -9223372036854775808, -447, 1}
    if (34 <= 1) {
        fmt.Println(v2, v3, v4)
    }
    v7 := GenS1{F0: ("" + "\n\t"), F1: (v2 ^ (v2 * v2)), F2: (v3 == -230)}
    switch -586 {
    case 9223372036854775807:
        if len(v6) > 0 {
            v6[0] = 1
        }
    }
    var v8 float32 = float32((v4 * v4))
    var v9 int16 = v2
    v9 = v2
    if len(v6) > 0 {
        v6[0] = -1
    }
    var v10 int = -892
    if (false || (v5 <= -928)) {
        _ = len(v6)
    } else {
        for fi11 := 0; fi11 < 20; fi11++ {
            if len(v6) > 0 {
                v6[0] = (fi11 * 1)
            }
            var v12 int8 = -1
            _ = fi11
            _ = v12
        }
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        v4 -= 3.14159
    }
    if false {
        _ = len(v6)
        v10 += (v10 + (v10 * (v10 / ((42) | 1))))
        var v13 uint = 65535
        _ = v13
    } else {
        fmt.Println(v8, v2)
        for fi14 := 0; fi14 < 2; fi14++ {
            var v15 int16 = (^(v2))
            if (32767 > v2) {
                var v16 uint8 = 249
                _ = v16
            }
            v5 -= (^(v5))
            _ = fi14
            _ = v15
        }
    }
    v2 += -799
    fmt.Println(v5, v3)
    fmt.Println(v5, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
