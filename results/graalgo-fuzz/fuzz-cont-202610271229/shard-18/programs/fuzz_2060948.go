package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 uint64) uint {
    v5 := []int{370, -798, 315}
    if true {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        var v6 float32 = float32((((1.5 - float64(0.0)) * 3.14159) - 1.5))
        _ = v6
    } else {
        if false {
            if (4294967295 <= 1) {
                p0 = p0
            }
        } else {
            v7 := GenS1{F0: 2}
            _ = v7
        }
        var v8 uint32 = 4294967295
        var v9 bool = false
        _ = v8
        _ = v9
    }
    if (546 > -2147483648) {
        fmt.Println(p0)
        var v10 uint64 = 4294967295
        _ = v10
    } else {
        v5 = append(v5, 32767)
        v11 := GenS1{F0: 197}
        for fi12 := 0; fi12 < 6; fi12++ {
            p0 = (^(p0))
            p0 = (p0 & 2)
            _ = fi12
        }
        _ = v11
    }
    p0 -= p0
    _ = p0
    _ = v5
    return 255
}

func main() {
    var v1 uint64 = 42
    v2 := make(map[string]int)
    var v3 uint64 = v1
    v4 := []uint32{204, 917, 2}
    v1 = 1
    switch 188 {
    case -888:
        for fi13 := 0; fi13 < 19; fi13++ {
            if ((139 != 65535) && (42 > 127)) {
                _ = len(v4)
                for fi14 := 0; fi14 < 3; fi14++ {
                    fmt.Println(fi13, v3, v1)
                    _ = fi14
                }
            } else {
                v3 -= v1
                v3 = (v1 | (^(v3)))
                var v15 float64 = (float64((((0.0 / 3.14159) + float32(1.5)) / 1.0)) * 1.5)
                _ = v15
            }
            switch "hello" {
            case "hello":
                fmt.Println(v3, v3)
            }
            _ = fi13
        }
    }
    v1 = (v1 % ((366) | 1))
    delete(v2, "a")
    if (!false) {
        v16 := GenS1{F0: 2}
        v1 = (^(v3))
        fmt.Println(v1, v1)
        _ = v16
    }
    v1 += v1
    v17 := GenS1{F0: 2}
    v18 := GenS1{F0: 65535}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
