package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int }

func main() {
    var v1 int32 = -1
    var v2 float32 = 0.1
    v3 := []string{""}
    v4 := GenS1{F0: (v1 - (v1 % ((v1) | 1))), F1: 0}
    v5 := GenS2{F0: 32767}
    fmt.Println(v2)
    for fi6 := 0; fi6 < 15; fi6++ {
        for fi7 := 0; fi7 < 20; fi7++ {
            v3 = append(v3, "")
            if (v1 < -2147483648) {
                v8 := GenS1{F0: v1, F1: 2}
                v9 := GenS2{F0: 280}
                if len(v3) > 0 {
                    v3[0] = "hello"
                }
                _ = v8
                _ = v9
            }
            _ = fi7
        }
        _ = fi6
    }
    for fi10 := 0; fi10 < 17; fi10++ {
        v2 = v2
        for fi11 := 0; fi11 < 4; fi11++ {
            var v12 uint64 = 1
            _ = fi11
            _ = v12
        }
        var v13 uint16 = 992
        _ = fi10
        _ = v13
    }
    var v14 uint = 4294967295
    fmt.Println(v2, v1, v2)
    for fi15 := 0; fi15 < 6; fi15++ {
        v16 := GenS3{F0: -9223372036854775808}
        _ = fi15
        _ = v16
    }
    var v17 float64 = (0.1 + 0.1)
    var v18 int16 = 32767
    v17 += (1e10 * 1.5)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
