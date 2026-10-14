package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }

func main() {
    v1 := make(map[int]uint32)
    var v2 int32 = -2147483648
    v3 := make(map[int]string)
    v4 := make(map[string]int)
    var v5 int16 = 69
    var v6 string = ""
    var v7 uint = 1
    var v8 string = v6
    if ((!false) || true) {
        v2 += (v2 / (((-(v2))) | 1))
        switch v8 {
        case "hello":
            for fi9 := 0; fi9 < 16; fi9++ {
                v10 := GenS1{F0: v5, F1: 0}
                _ = fi9
                _ = v10
            }
        case "\n\t":
            var v11 bool = false
            _ = v11
        }
    }
    delete(v4, (v8 + v8))
    for fi12 := 0; fi12 < 10; fi12++ {
        v13 := GenS1{F0: 32767, F1: 9223372036854775807}
        for fi14 := 0; fi14 < 17; fi14++ {
            v1[(^(fi14))] = 713
            var v15 bool = true
            var v16 uint64 = 1
            _ = fi14
            _ = v15
            _ = v16
        }
        v5 += (v5 << (((^(v7))) & 63))
        _ = fi12
        _ = v13
    }
    v8 = (v8 + "a")
    var v17 uint = 255
    var v18 int16 = 1
    v17 += (v7 / ((v7) | 1))
    var v19 float64 = 3.14159
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
