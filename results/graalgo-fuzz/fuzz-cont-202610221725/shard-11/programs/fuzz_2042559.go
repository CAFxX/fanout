package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
func genFn0(p0 int64) uint16 {
    var v6 int = 0
    for fi7 := 0; fi7 < 14; fi7++ {
        v8 := GenS1{F0: true, F1: 1}
        _ = fi7
        _ = v8
    }
    if true {
        p0 += p0
    }
    _ = p0
    _ = v6
    return 1
}

func main() {
    var v1 int32 = 244
    var v2 uint16 = 32767
    v3 := make(map[int]uint32)
    v4 := make(map[uint32]uint32)
    var v5 int32 = v1
    var v9 float64 = (3.14159 * 0.0)
    v1 += -2147483648
    var v10 int16 = 0
    fmt.Println(v2, v5)
    if (!true) {
        v2 = v2
    } else {
        v11 := GenS1{F0: false, F1: 2}
        switch "hello" {
        case "a":
            for fi12 := 0; fi12 < 5; fi12++ {
                v13 := []int{0, (-(fi12)), fi12}
                _ = fi12
                _ = v13
            }
        case "foo bar":
            var v14 int8 = 42
            _ = v14
        }
        _ = v11
    }
    var v15 uint64 = 1
    v10 = -32768
    var v16 int16 = 0
    v16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
