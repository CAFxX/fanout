package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]bool)
    var v3 float64 = 3.14159
    var v4 string = "a"
    v3 -= 1.5
    var v5 float32 = 1e10
    v6 := GenS1{F0: -9223372036854775808, F1: 4294967295}
    v5 = (v5 - v5)
    if (false && false) {
        v5 += (float32((v3 / 0.1)) / v5)
        if true {
            if (4294967295 >= 2147483647) {
                for fi7 := 0; fi7 < 1; fi7++ {
                    fmt.Println(v4, v4, fi7)
                    _ = fi7
                }
            } else {
                v8 := []int{9223372036854775807, 4294967295, 32768}
                var v9 int16 = 32767
                fmt.Println(v3)
                _ = v8
                _ = v9
            }
            for fi10 := 0; fi10 < 20; fi10++ {
                if false {
                    fmt.Println(fi10, v4, fi10)
                    var v11 uint16 = 42
                    _ = v11
                }
                fmt.Println(v4)
                _ = fi10
            }
            fmt.Println(v3)
        }
    }
    fmt.Println(v4)
    var v12 int = 127
    for fi13 := 0; fi13 < 19; fi13++ {
        fmt.Println(v5)
        for fi14 := 0; fi14 < 16; fi14++ {
            v3 += v3
            delete(v1, (v12 & v12))
            _ = fi14
        }
        v15 := []int{128}
        _ = fi13
        _ = v15
    }
    var v16 uint64 = 18446744073709551615
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
