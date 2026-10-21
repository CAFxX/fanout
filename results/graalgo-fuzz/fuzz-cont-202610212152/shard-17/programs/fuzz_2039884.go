package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 int64 }

func main() {
    var v1 uint16 = 255
    var v2 int64 = 608
    v3 := make(map[string]string)
    v4 := []int{42}
    v5 := GenS1{F0: v2, F1: 1, F2: (v2 | (v2 % ((0) | 1)))}
    for fi6 := 0; fi6 < 2; fi6++ {
        v1 += (^(v1))
        for fi7 := 0; fi7 < 3; fi7++ {
            if false {
                v3[("a" + (("a" + "hello") + "\n\t"))] = "hello"
                var v8 uint32 = uint32(v1)
                _ = v8
            } else {
                v9 := GenS1{F0: (v2 * v2), F1: -1, F2: (v2 % ((-659) | 1))}
                _ = v9
            }
            v2 = -9223372036854775808
            _ = fi7
        }
        _ = fi6
    }
    var v10 uint = 65535
    if len(v4) > 0 {
        v4[0] = 32768
    }
    fmt.Println(v1, v2)
    var v11 bool = false
    fmt.Println(v1, v10)
    for fi12 := 0; fi12 < 5; fi12++ {
        switch -9223372036854775808 {
        case 32767:
            var v13 uint64 = 2
            _ = v13
        case 173:
            for fi14 := 0; fi14 < 5; fi14++ {
                _ = len(v4)
                _ = fi14
            }
        case 128:
            if (((((0.0 * 3.14159) + 1.5) * -2.25) - ((float32(float64(0.1)) * 1.5) / -2.25)) == -2.25) {
                v11 = v11
                v4 = append(v4, fi12)
                if (9223372036854775807 > -9223372036854775808) {
                    if v11 {
                        _ = len(v4)
                    }
                    v11 = v11
                } else {
                    v3[("" + (("hello" + "foo bar") + "foo bar"))] = "\n\t"
                    var v15 int8 = 42
                    _ = v15
                }
            }
        }
        for fi16 := 0; fi16 < 11; fi16++ {
            var v17 int8 = -128
            _ = fi16
            _ = v17
        }
        _ = fi12
    }
    var v18 float32 = (1e10 / (0.1 - 1.5))
    v18 -= v18
    v10 = v10
    v19 := GenS1{F0: (v2 / ((v2) | 1)), F1: 42, F2: v2}
    v18 -= ((float32(-2.25) * v18) / v18)
    var v20 int16 = 317
    v11 = (v11 || (!(!v11)))
    v11 = (v11 || (false || v11))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
