package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
func genFn0() uint8 {
    var v6 uint64 = 2
    v6 = v6
    var v7 float32 = 1.5
    _ = v6
    _ = v7
    return 128
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]uint32)
    var v3 float64 = 1e10
    var v4 bool = true
    var v5 int8 = 100
    fmt.Println(v5, v4)
    switch 9223372036854775807 {
    case 128:
        var v8 uint = 1
        _ = v8
    }
    for fi9 := 0; fi9 < 9; fi9++ {
        for fi10 := 0; fi10 < 7; fi10++ {
            if v4 {
                v4 = v4
                switch (fi10 ^ (fi9 / ((fi9) | 1))) {
                case 127:
                    v11 := v1[(fi9 & (fi10 & (fi10 ^ fi10)))]
                    _ = v11
                case 2147483647:
                    v12 := GenS1{F0: 4294967295, F1: 128}
                    _ = v12
                }
                var v13 uint16 = 65535
                _ = v13
            }
            v3 -= 3.14159
            var v14 bool = (v4 && (!(1000 > -78)))
            _ = fi10
            _ = v14
        }
        v3 += v3
        _ = fi9
    }
    v15 := GenS1{F0: 0, F1: 2}
    var v16 int = -1
    if (true && v4) {
        v5 = v5
        v3 = (0.0 / 1.0)
    } else {
        v4 = (!(1 <= 746))
        v17 := []int{-9223372036854775808, (v16 - -725), v16}
        _ = v17
    }
    v18 := GenS1{F0: 719, F1: 1}
    var v19 uint32 = 2147483647
    v20 := GenS1{F0: 18446744073709551615, F1: 0}
    v21 := GenS1{F0: 42, F1: 891}
    v22 := GenS1{F0: 255, F1: 609}
    for fi23 := 0; fi23 < 1; fi23++ {
        fmt.Println(v16, v16)
        _ = fi23
    }
    fmt.Println(v5, v19, v19)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
