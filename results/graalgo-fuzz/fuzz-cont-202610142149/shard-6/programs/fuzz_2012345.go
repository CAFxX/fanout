package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 string; F1 int32 }

func main() {
    var v1 uint64 = 160
    var v2 int64 = 9223372036854775807
    var v3 bool = (42 == 2)
    v4 := GenS1{F0: 0, F1: 959}
    fmt.Println(v3, v2)
    v2 -= 9223372036854775807
    var v5 float32 = (-2.25 * 0.0)
    v3 = (-128 == 127)
    v2 = (v2 & -858)
    var v6 uint16 = 755
    if false {
        for fi7 := 0; fi7 < 13; fi7++ {
            v8 := GenS2{F0: "\n\t", F1: 115}
            if false {
                v5 = -2.25
            } else {
                var v9 uint8 = 2
                v9 = v9
                _ = v9
            }
            v3 = (true && (-2.25 < -2.25))
            _ = fi7
            _ = v8
        }
    } else {
        v6 = v6
        fmt.Println(v5)
    }
    v2 += v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
