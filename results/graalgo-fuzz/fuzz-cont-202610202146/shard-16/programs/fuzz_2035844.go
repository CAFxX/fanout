package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 uint64 }

func main() {
    var v1 float32 = (3.14159 + float32(((0.1 / 1.0) - (-2.25 / (((((1.5 * (0.0 / 0.1)) + (float64(1e10) + ((1.5 / 1.0) * 1.5))) - 1.5) - 1.5) - 0.0)))))
    v2 := []int{-616, -477, 0}
    var v3 string = "hello"
    var v4 int64 = 0
    fmt.Println(v3, v3)
    fmt.Println(v3, v3, v4)
    v1 = v1
    v1 += ((((float32(1e10) * v1) * 0.0) + v1) - v1)
    var v5 int = 42
    for fi6 := 0; fi6 < 11; fi6++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi6
    }
    v2 = append(v2, v5)
    if ((42 > -1) || (-128 > -128)) {
        for fi7 := 0; fi7 < 11; fi7++ {
            v4 -= (v4 + 9223372036854775807)
            _ = fi7
        }
        var v8 uint32 = 584
        _ = v8
    }
    for fi9 := 0; fi9 < 11; fi9++ {
        for fi10 := 0; fi10 < 12; fi10++ {
            v11 := GenS1{F0: v3, F1: 2, F2: 42}
            var v12 uint = 843
            fmt.Println(v12)
            var v13 int8 = 114
            if len(v2) > 0 {
                v2[0] = v5
            }
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
        }
        for fi14 := 0; fi14 < 2; fi14++ {
            v3 = "hello"
            if false {
                if (191 <= -32768) {
                    var v15 int16 = 32767
                    fmt.Println(v3)
                    v1 -= float32(float64(v1))
                    v15 -= (v15 + 710)
                    _ = v15
                }
                v3 = v3
                v4 = 42
            }
            var v16 string = v3
            _ = fi14
            _ = v16
        }
        _ = fi9
    }
    v2 = append(v2, (v5 % ((32768) | 1)))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if (0 >= 127) {
        var v17 int16 = 42
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v17
    }
    var v18 uint8 = 0
    if len(v2) > 0 {
        v2[0] = 4294967295
    }
    v18 = 1
    switch (-(v5)) {
    case 0:
        var v19 uint8 = (v18 + v18)
        _ = v19
    case 786:
        if (v3 != v3) {
            if true {
                v18 += 1
                v2 = append(v2, v5)
                v2 = append(v2, v5)
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
