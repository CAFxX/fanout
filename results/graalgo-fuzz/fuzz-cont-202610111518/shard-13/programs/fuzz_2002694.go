package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }

func main() {
    v1 := []int{32767, 768}
    var v2 float32 = 0.1
    var v3 float64 = 1e10
    v4 := make(map[string]string)
    v5 := make(map[string]string)
    switch 128 {
    case 908:
        for fi6 := 0; fi6 < 7; fi6++ {
            if true {
                v3 -= (v3 / float64(v2))
            }
            _ = fi6
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v2 = (v2 - v2)
    var v7 uint = 994
    v2 += ((float32((((v3 * 0.0) + v3) - (v3 + (((v3 * (1.5 + v3)) - -2.25) - v3)))) / v2) - 0.1)
    v8 := GenS1{F0: 1, F1: -9223372036854775808}
    switch 127 {
    case 32768:
        var v9 float32 = v2
        _ = v9
    }
    _ = len(v1)
    v3 -= v3
    v1 = append(v1, 42)
    var v10 uint = 271
    for fi11 := 0; fi11 < 19; fi11++ {
        var v12 string = "\n\t"
        delete(v5, v12)
        if true {
            if (v2 != (v2 * float32(((v3 / 0.1) * ((v3 - (0.0 + v3)) - v3))))) {
                v13 := GenS1{F0: 0, F1: fi11}
                v12 = "foo bar"
                _ = v13
            }
        } else {
            v12 = ("a" + v12)
        }
        _ = fi11
        _ = v12
    }
    for fi14 := 0; fi14 < 4; fi14++ {
        var v15 int16 = 587
        if false {
            var v16 bool = (false || (2147483648 > 65535))
            v15 = (v15 >> ((v7) & 63))
            if (v15 != -502) {
                v2 -= v2
                for fi17 := 0; fi17 < 14; fi17++ {
                    if v16 {
                        var v18 uint16 = 65535
                        _ = v18
                    }
                    fmt.Println(v2, fi17, v15)
                    _ = fi17
                }
            }
            _ = v16
        }
        fmt.Println(v2, fi14, v7)
        _ = fi14
        _ = v15
    }
    var v19 float32 = v2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
