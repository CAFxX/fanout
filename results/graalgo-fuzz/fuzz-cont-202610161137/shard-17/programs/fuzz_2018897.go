package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 string }
func genFn0() uint32 {
    var v5 float64 = (1e10 - (-2.25 / ((0.1 + 0.0) + 1e10)))
    v6 := []int{2147483647}
    v5 += ((float64(3.14159) / v5) / v5)
    fmt.Println(v5)
    _ = v5
    _ = v6
    return 398
}

func main() {
    v1 := make(map[string]string)
    var v2 int64 = -9223372036854775808
    v3 := make(map[uint32]string)
    var v4 uint64 = 780
    v7 := GenS1{F0: -1, F1: (v2 / (((v2 % (((v2 + (-(v2)))) | 1))) | 1)), F2: "hello"}
    if true {
        var v8 uint8 = 2
        _ = v8
    }
    fmt.Println(v4, v4)
    fmt.Println(v2, v2)
    for fi9 := 0; fi9 < 19; fi9++ {
        v4 = 42
        fmt.Println(v2, fi9, fi9)
        _ = fi9
    }
    var v10 float32 = float32(1.5)
    fmt.Println(v2)
    for fi11 := 0; fi11 < 18; fi11++ {
        switch (("hello" + "hello") + "hello") {
        case "":
            for fi12 := 0; fi12 < 5; fi12++ {
                v2 -= v2
                v13 := v3[680]
                if (!false) {
                    var v14 uint32 = genFn0()
                    v15 := GenS1{F0: 1000, F1: v2, F2: ("a" + v13)}
                    _ = v14
                    _ = v15
                }
                _ = fi12
                _ = v13
            }
        case "hello":
            var v16 float32 = 0.1
            _ = v16
        }
        _ = fi11
    }
    var v17 float64 = -2.25
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
