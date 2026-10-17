package main

import (
    "fmt"
)


func main() {
    v1 := make(map[string]string)
    var v2 uint8 = 1
    v3 := make(map[uint32]int)
    var v4 uint16 = uint16(v2)
    var v5 float64 = (float64(1e10) - 1e10)
    var v6 uint32 = 0
    fmt.Println(v4, v2, v5)
    var v7 float64 = (-2.25 - (v5 - v5))
    var v8 uint16 = (v4 / (((v4 | v4)) | 1))
    if (v2 >= v2) {
        if (-9223372036854775808 != 894) {
            _ = len(v1)
            var v9 float32 = 0.1
            _ = v9
        } else {
            for fi10 := 0; fi10 < 16; fi10++ {
                v11 := []int{(-(fi10)), 687, fi10}
                var v12 uint16 = (^(v4))
                _ = fi10
                _ = v11
                _ = v12
            }
            var v13 uint32 = v6
            _ = v13
        }
        fmt.Println(v7, v4)
    }
    v6 -= (^(v6))
    fmt.Println(v5)
    for fi14 := 0; fi14 < 11; fi14++ {
        v5 += v5
        v15 := []int{(fi14 << ((255) & 63)), fi14}
        _ = fi14
        _ = v15
    }
    if (230 > 1) {
        for fi16 := 0; fi16 < 10; fi16++ {
            if (!false) {
                v4 = (^(v4))
            }
            fmt.Println(v7, v8)
            fmt.Println(v2, v2)
            _ = fi16
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
