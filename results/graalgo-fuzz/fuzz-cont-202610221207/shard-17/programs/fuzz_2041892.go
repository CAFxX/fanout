package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 uint32 }

func main() {
    var v1 uint32 = 304
    var v2 float32 = 1e10
    var v3 string = "hello"
    var v4 uint64 = 42
    var v5 uint32 = (v1 | (v1 / ((v1) | 1)))
    v4 = (v4 + uint64(v5))
    if ((!(0 >= -9223372036854775808)) && ((false && (((true || ((1.5 / (((-2.25 + 0.1) * 1.5) / 3.14159)) == -2.25)) && false) && false)) && (1000 >= -32768))) {
        v5 = (v1 + uint32(v4))
        v6 := []int{437, 42}
        _ = v6
    }
    var v7 int64 = -9223372036854775808
    v5 += (v1 + 0)
    var v8 uint32 = (v5 ^ 290)
    var v9 float32 = v2
    var v10 float32 = float32(0.0)
    v11 := []int{4294967295, 127, -486, 2147483647}
    if true {
        if (0 < 42) {
            for fi12 := 0; fi12 < 5; fi12++ {
                var v13 float64 = (1.5 * (1e10 / float64((float32(float64((v10 * v10))) / v10))))
                _ = fi12
                _ = v13
            }
            v2 += ((v2 * 1.5) + v10)
        } else {
            for fi14 := 0; fi14 < 13; fi14++ {
                v15 := GenS1{F0: 1, F1: false, F2: (v8 << ((42) & 63))}
                if (false || (!true)) {
                    v9 += ((v2 / (1e10 * v9)) + 0.0)
                    var v16 uint16 = 932
                    _ = v16
                }
                _ = fi14
                _ = v15
            }
            v11 = append(v11, -2147483648)
            var v17 uint8 = 0
            _ = v17
        }
        for fi18 := 0; fi18 < 1; fi18++ {
            v9 -= 0.1
            _ = fi18
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
