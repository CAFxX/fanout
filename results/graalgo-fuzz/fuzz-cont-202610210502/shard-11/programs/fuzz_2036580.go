package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 uint32 }

func main() {
    v1 := make(map[int]string)
    var v2 int32 = 2147483647
    v3 := []int{127, 279, -9223372036854775808}
    var v4 int = 42
    v5 := []uint32{2147483647, 0, 2147483647}
    for fi6 := 0; fi6 < 3; fi6++ {
        for fi7 := 0; fi7 < 20; fi7++ {
            fmt.Println(v2, fi6, v4)
            _ = fi7
        }
        _ = fi6
    }
    var v8 int16 = -105
    fmt.Println(v2)
    v4 -= (v4 | (v4 | (-(v4))))
    var v9 uint32 = 880
    fmt.Println(v2, v2, v9)
    v2 = (v2 % ((-760) | 1))
    switch "hello" {
    case "hello":
        switch "a" {
        case "\n\t":
            fmt.Println(v4)
        case "hello":
            delete(v1, (v4 % (((-(v4))) | 1)))
        }
    case "a":
        switch (-(v4)) {
        case 42:
            if (true || (1 >= 127)) {
                _ = len(v3)
                var v10 int64 = 9223372036854775807
                v8 -= (v8 / ((1) | 1))
                _ = v10
            } else {
                for fi11 := 0; fi11 < 4; fi11++ {
                    for fi12 := 0; fi12 < 15; fi12++ {
                        v13 := GenS1{F0: 0, F1: 113, F2: 63}
                        v14 := GenS1{F0: -1, F1: 0, F2: v9}
                        v15 := GenS1{F0: -60, F1: 127, F2: 2}
                        _ = fi12
                        _ = v13
                        _ = v14
                        _ = v15
                    }
                    var v16 uint64 = 2
                    _ = fi11
                    _ = v16
                }
                var v17 uint8 = 12
                _ = v17
            }
        case -669:
            var v18 float32 = 1e10
            _ = v18
        case 4294967295:
            fmt.Println(v9, v9)
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
