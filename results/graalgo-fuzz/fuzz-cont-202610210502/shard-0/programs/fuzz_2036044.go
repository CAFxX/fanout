package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 uint }

func main() {
    v1 := []int{-692, 830}
    v2 := make(map[int]uint32)
    var v3 int32 = -1
    v4 := v2[-171]
    for fi5 := 0; fi5 < 19; fi5++ {
        fmt.Println(v3)
        var v6 uint = 0
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 3; fi7++ {
        v4 += 2
        var v8 int = fi7
        _ = fi7
        _ = v8
    }
    v4 = v4
    v9 := GenS1{F0: "\n\t", F1: 234, F2: 18446744073709551615}
    var v10 int16 = -32768
    var v11 float64 = -2.25
    switch 42 {
    case 32768:
        for fi12 := 0; fi12 < 16; fi12++ {
            if true {
                var v13 uint64 = 4294967295
                v11 -= 1e10
                if (4294967295 > 18446744073709551615) {
                    _ = len(v2)
                    _ = len(v1)
                }
                _ = v13
            }
            var v14 int16 = (v10 % (((v10 ^ v10)) | 1))
            var v15 int32 = (v3 + v3)
            _ = fi12
            _ = v14
            _ = v15
        }
    case 42:
        _ = len(v1)
    }
    for fi16 := 0; fi16 < 17; fi16++ {
        if (2147483648 <= v4) {
            var v17 string = ("" + "a")
            _ = v17
        } else {
            v11 -= 0.0
        }
        v11 = v11
        _ = fi16
    }
    var v18 int32 = v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
