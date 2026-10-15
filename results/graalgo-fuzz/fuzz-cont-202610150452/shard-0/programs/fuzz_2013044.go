package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    v1 := make(map[string]uint32)
    var v2 int32 = 42
    var v3 uint32 = 42
    v4 := make(map[int]string)
    var v5 uint32 = (v3 | v3)
    if true {
        var v6 float32 = ((3.14159 / float32(0.1)) - 0.1)
        _ = v6
    }
    if false {
        v2 += v2
        if false {
            var v7 uint16 = 2
            _ = v7
        } else {
            if false {
                if (!(0.0 > 0.1)) {
                    var v8 uint16 = 2
                    var v9 float64 = 0.0
                    _ = v8
                    _ = v9
                }
                fmt.Println(v3, v5, v3)
                for fi10 := 0; fi10 < 14; fi10++ {
                    var v11 uint64 = 1
                    _ = fi10
                    _ = v11
                }
            }
            if (-2.25 <= (0.1 / 1.0)) {
                fmt.Println(v2, v3, v2)
                v5 = v5
            } else {
                v5 = v5
            }
            for fi12 := 0; fi12 < 11; fi12++ {
                var v13 int = (fi12 + (fi12 / ((fi12) | 1)))
                _ = fi12
                _ = v13
            }
        }
    }
    var v14 uint64 = 628
    fmt.Println(v2, v3)
    v15 := GenS1{F0: -1}
    v5 = (v5 & (^(v3)))
    v16 := GenS1{F0: 1}
    if ("a" >= "\n\t") {
        var v17 uint = 4294967295
        v4[-2147483648] = ("\n\t" + "")
        _ = v17
    } else {
        var v18 float64 = (1.5 / 1e10)
        _ = v18
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
