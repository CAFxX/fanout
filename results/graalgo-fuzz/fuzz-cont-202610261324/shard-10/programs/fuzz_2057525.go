package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 uint = 2
    var v2 float32 = 1.5
    var v3 uint64 = 0
    v4 := make(map[int]string)
    var v5 uint8 = 51
    var v6 int64 = 42
    v7 := GenS1{F0: (v5 * 0)}
    for fi8 := 0; fi8 < 6; fi8++ {
        v5 = v5
        _ = fi8
    }
    v1 -= v1
    var v9 uint16 = 662
    var v10 int = 1
    v3 -= 737
    for fi11 := 0; fi11 < 17; fi11++ {
        for fi12 := 0; fi12 < 13; fi12++ {
            fmt.Println(v1, v10)
            var v13 uint64 = (v3 | (^(v3)))
            _ = fi12
            _ = v13
        }
        for fi14 := 0; fi14 < 5; fi14++ {
            if (!(!true)) {
                v15 := []int{v10, (fi11 ^ (v10 / ((v10) | 1)))}
                var v16 uint8 = 42
                _ = v15
                _ = v16
            } else {
                var v17 uint32 = 4294967295
                _ = v17
            }
            _ = fi14
        }
        fmt.Println(v9, v5, fi11)
        _ = fi11
    }
    v1 += 65535
    fmt.Println(v5)
    var v18 uint32 = 1
    if true {
        v18 = (^(v18))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
