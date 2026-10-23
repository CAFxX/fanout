package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 bool; F1 uint32; F2 bool }
func genFn0(p0 string) uint8 {
    v4 := []int{9223372036854775807, 32767}
    for fi5 := 0; fi5 < 6; fi5++ {
        var v6 string = "hello"
        var v7 int32 = 880
        _ = fi5
        _ = v6
        _ = v7
    }
    _ = p0
    _ = v4
    return 6
}

func main() {
    var v1 bool = (0 == 0)
    v2 := make(map[int]uint32)
    var v3 uint = 217
    var v8 uint = (v3 / ((v3) | 1))
    v9 := GenS1{F0: "a", F1: 63}
    for fi10 := 0; fi10 < 8; fi10++ {
        v11 := []int{}
        if false {
            v8 -= v3
            for fi12 := 0; fi12 < 16; fi12++ {
                var v13 int32 = 1
                _ = fi12
                _ = v13
            }
            var v14 uint32 = 507
            _ = v14
        } else {
            v15 := GenS1{F0: "\n\t", F1: 196}
            fmt.Println(v3, v8)
            if false {
                for fi16 := 0; fi16 < 10; fi16++ {
                    fmt.Println(fi16, v1, fi10)
                    v17 := v2[9223372036854775807]
                    _ = fi16
                    _ = v17
                }
                v8 = v8
                v8 = v8
            }
            _ = v15
        }
        _ = fi10
        _ = v11
    }
    fmt.Println(v8, v8, v8)
    v8 = 65535
    var v18 uint32 = 534
    v18 += v18
    v3 -= 4294967295
    _ = len(v2)
    v19 := GenS1{F0: "\n\t", F1: 2}
    fmt.Println(v1, v18, v3)
    v20 := GenS2{F0: v1, F1: 567, F2: (!(v1 && (uint16(v3) > uint16(v8))))}
    var v21 float64 = ((1.5 + (0.1 * ((1.5 - 3.14159) * 0.1))) / ((1e10 / 3.14159) * float64((3.14159 + 1.5))))
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v8
    _ = v9
}
