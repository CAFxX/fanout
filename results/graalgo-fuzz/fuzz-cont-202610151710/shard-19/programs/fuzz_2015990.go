package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 int16 }
type GenS2 struct { F0 int64 }

func main() {
    var v1 float32 = (1.5 - -2.25)
    var v2 int16 = 390
    v3 := []string{("" + ("\n\t" + "hello")), "a"}
    var v4 float32 = v1
    v5 := GenS1{F0: v2, F1: 42, F2: v2}
    _ = len(v3)
    var v6 uint32 = 915
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v7 := GenS1{F0: v2, F1: 42, F2: 1}
    v8 := GenS1{F0: (-(v2)), F1: 333, F2: (v2 ^ (v2 / ((v2) | 1)))}
    if (1 < 2147483647) {
        v9 := GenS1{F0: (v2 / ((1000) | 1)), F1: 0, F2: (v2 >> ((354) & 63))}
        var v10 int = -2147483648
        _ = v10
        _ = v9
    }
    v3 = append(v3, "hello")
    v2 += v2
    for fi11 := 0; fi11 < 7; fi11++ {
        v12 := GenS1{F0: (v2 / ((v2) | 1)), F1: 255, F2: 0}
        v1 -= 1.5
        _ = fi11
        _ = v12
    }
    var v13 string = "a"
    var v14 uint = 65535
    fmt.Println(v13)
    v2 = (v2 * 1000)
    if true {
        v1 = (1e10 * (float32(float64(v4)) / v1))
        v2 -= v2
        var v15 int32 = 1
        _ = v15
    } else {
        if (v14 >= (^(v14))) {
            v4 += 1.5
            v16 := GenS2{F0: -252}
            _ = v16
        } else {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        v14 -= (^(v14))
    }
    var v17 uint16 = 42
    v2 -= v2
    v2 = 16
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
