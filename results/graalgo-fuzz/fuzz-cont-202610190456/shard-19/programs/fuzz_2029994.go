package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 uint8 }
type GenS2 struct { F0 uint; F1 uint32; F2 bool }
type GenS3 struct { F0 int8 }
func genFn0(p0 uint32) uint16 {
    fmt.Println(p0)
    p0 -= (^(p0))
    _ = p0
    return 42
}

func main() {
    var v1 bool = false
    var v2 string = "\n\t"
    v3 := []string{v2, (("hello" + v2) + v2), v2}
    v2 = (v2 + (v2 + v2))
    var v4 float64 = float64(-2.25)
    v2 = "foo bar"
    var v5 float32 = 0.0
    v6 := GenS1{F0: "", F1: 1000, F2: 34}
    for fi7 := 0; fi7 < 14; fi7++ {
        _ = len(v3)
        _ = fi7
    }
    var v8 uint = 0
    v9 := GenS2{F0: 2, F1: 456, F2: false}
    var v10 int16 = 0
    v5 = -2.25
    switch 1 {
    case 42:
        for fi11 := 0; fi11 < 16; fi11++ {
            _ = len(v3)
            for fi12 := 0; fi12 < 14; fi12++ {
                if len(v3) > 0 {
                    v3[0] = v2
                }
                fmt.Println(v10, v4, v2)
                if true {
                    var v13 int64 = -9223372036854775808
                    _ = v13
                } else {
                    if (!v1) {
                        v4 += ((v4 * ((((1.5 - v4) * v4) / v4) - v4)) * v4)
                    } else {
                        v14 := GenS3{F0: 0}
                        _ = v14
                    }
                    v8 = (^(v8))
                }
                _ = fi12
            }
            _ = fi11
        }
    case -816:
        v8 -= v8
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
