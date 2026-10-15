package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 uint }
func genFn0(p0 int16, p1 string, p2 bool) uint16 {
    p1 = p1
    for fi5 := 0; fi5 < 6; fi5++ {
        p2 = p2
        _ = fi5
    }
    switch 2147483647 {
    case 32767:
        var v6 bool = (0.1 != 1.5)
        _ = v6
    default:
        p0 = 839
    }
    _ = p0
    _ = p1
    _ = p2
    return 65535
}

func main() {
    var v1 int16 = -198
    var v2 float32 = -2.25
    v3 := make(map[uint32]uint32)
    var v4 uint8 = 255
    var v7 uint64 = 1
    delete(v3, 2147483648)
    v8 := GenS1{F0: 0, F1: true, F2: 866}
    var v9 int16 = int16(v7)
    v1 -= (v9 | v9)
    v4 = v4
    v7 = (v7 * (v7 ^ v7))
    if (127 >= 127) {
        if (false && (false && false)) {
            v10 := []int{}
            _ = v10
        } else {
            for fi11 := 0; fi11 < 2; fi11++ {
                v1 += (v9 & (^(v1)))
                var v12 bool = (true || false)
                v13 := GenS1{F0: 611, F1: (uint(v4) < 829), F2: 1}
                _ = fi11
                _ = v12
                _ = v13
            }
            v14 := []int{4294967295, 9223372036854775807, -1, 128}
            _ = v14
        }
        v15 := GenS1{F0: -1, F1: (663 == 340), F2: 0}
        fmt.Println(v9)
        _ = v15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
