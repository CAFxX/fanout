package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
func genFn0(p0 uint16) int32 {
    if (143 <= 4294967295) {
        v6 := []int{128, 322}
        for fi7 := 0; fi7 < 15; fi7++ {
            for fi8 := 0; fi8 < 4; fi8++ {
                p0 = (^(p0))
                _ = fi8
            }
            _ = fi7
        }
        v9 := GenS1{F0: 65535, F1: 0}
        _ = v6
        _ = v9
    } else {
        p0 = 255
    }
    v10 := []int{-9223372036854775808, 127, -9223372036854775808}
    p0 -= p0
    var v11 uint = 255
    _ = p0
    _ = v10
    _ = v11
    return 2147483647
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int8 = 0
    v3 := make(map[string]bool)
    var v4 uint16 = 568
    v5 := make(map[string]int)
    v4 = (^(v4))
    if (!(!true)) {
        v12 := GenS1{F0: 42, F1: 0}
        _ = v12
    } else {
        v13 := []int{-104, 537, 0, 9223372036854775807}
        _ = v13
    }
    var v14 int16 = 712
    delete(v3, "foo bar")
    if (1 <= 1) {
        v2 = 42
        for fi15 := 0; fi15 < 10; fi15++ {
            if (!true) {
                fmt.Println(v14)
                var v16 uint16 = (v4 + (v4 / ((65535) | 1)))
                fmt.Println(v16, v14, v14)
                _ = v16
            }
            var v17 int8 = v2
            _ = fi15
            _ = v17
        }
    } else {
        v18 := []int{2147483647, -9223372036854775808}
        _ = len(v1)
        _ = v18
    }
    var v19 float32 = ((1.5 + 1.5) * (1.5 * (-2.25 * 1e10)))
    var v20 uint8 = 0
    v21 := GenS1{F0: 2147483647, F1: (v20 % ((98) | 1))}
    fmt.Println(v20, v20)
    _ = len(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
