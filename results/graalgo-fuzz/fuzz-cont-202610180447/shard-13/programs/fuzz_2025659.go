package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 uint64 }
type GenS2 struct { F0 uint }
func genFn0() uint16 {
    var v6 int16 = -1
    v7 := GenS1{F0: "a", F1: 9223372036854775807, F2: 4294967295}
    if (!true) {
        if true {
            v8 := []int{438}
            var v9 float32 = (1e10 - float32(1e10))
            for fi10 := 0; fi10 < 2; fi10++ {
                v9 = ((1.5 * v9) / v9)
                _ = len(v8)
                v9 -= ((v9 - -2.25) / 1.0)
                _ = fi10
            }
            _ = v8
            _ = v9
        }
        v6 = 42
    }
    fmt.Println(v6)
    _ = v6
    _ = v7
    return 1
}

func main() {
    v1 := make(map[string]string)
    var v2 uint64 = 42
    var v3 int = 32767
    v4 := make(map[int]uint32)
    var v5 uint8 = 2
    var v11 string = ("a" + "\n\t")
    for fi12 := 0; fi12 < 16; fi12++ {
        v5 -= 24
        var v13 string = v11
        _ = fi12
        _ = v13
    }
    fmt.Println(v2, v2)
    var v14 uint32 = 696
    for fi15 := 0; fi15 < 20; fi15++ {
        v14 = v14
        var v16 int = (fi15 ^ 987)
        _ = fi15
        _ = v16
    }
    var v17 bool = false
    v18 := GenS2{F0: 42}
    v11 = v11
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
