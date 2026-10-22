package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 int16; F1 int64; F2 bool }
type GenS3 struct { F0 uint; F1 uint }

func main() {
    var v1 bool = true
    var v2 bool = (!(false || false))
    v3 := make(map[int]bool)
    var v4 uint = 1
    v5 := make(map[uint32]bool)
    _ = len(v5)
    for fi6 := 0; fi6 < 8; fi6++ {
        for fi7 := 0; fi7 < 1; fi7++ {
            v8 := []int{2147483647, fi6, (fi7 * (fi6 * (fi6 * 1))), fi7}
            _ = fi7
            _ = v8
        }
        fmt.Println(v4, v2, v2)
        _ = fi6
    }
    for fi9 := 0; fi9 < 14; fi9++ {
        v10 := GenS1{F0: 42, F1: true}
        for fi11 := 0; fi11 < 8; fi11++ {
            var v12 int = 0
            var v13 uint8 = 255
            if (!v2) {
                if (!(((((3.14159 + float64(-2.25)) * 0.0) + 0.1) + 3.14159) <= (-2.25 - float64((-2.25 / -2.25))))) {
                    var v14 uint32 = 65535
                    v2 = v2
                    if ((v14 | (v14 % ((2147483648) | 1))) < (v14 * 42)) {
                        v15 := GenS2{F0: 42, F1: 402, F2: (v1 || ((("hello" + "") + "\n\t") >= "a"))}
                        _ = v15
                    }
                    _ = v14
                }
                var v16 uint64 = 40
                v17 := v5[2]
                _ = v16
                _ = v17
            } else {
                v18 := GenS1{F0: -1, F1: v2}
                v19 := GenS2{F0: -782, F1: -1, F2: false}
                v20 := GenS2{F0: 837, F1: -9223372036854775808, F2: (("a" >= "a") || (v1 || false))}
                _ = v18
                _ = v19
                _ = v20
            }
            _ = fi11
            _ = v12
            _ = v13
        }
        v2 = v1
        _ = fi9
        _ = v10
    }
    for fi21 := 0; fi21 < 20; fi21++ {
        v4 += v4
        _ = fi21
    }
    for fi22 := 0; fi22 < 10; fi22++ {
        fmt.Println(v1, v4)
        var v23 int32 = 2147483647
        v23 = (v23 - 62)
        _ = fi22
        _ = v23
    }
    v24 := GenS3{F0: 461, F1: v4}
    var v25 bool = v1
    fmt.Println(v2, v1, v1)
    if v25 {
        v26 := []int{32767, 349, -9223372036854775808}
        var v27 uint = v4
        _ = v26
        _ = v27
    }
    fmt.Println(v1, v1, v4)
    v4 -= v4
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
