package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    var v1 uint64 = 775
    v2 := make(map[int]uint32)
    var v3 int64 = 9223372036854775807
    if true {
        v4 := GenS1{F0: 674}
        _ = v4
    } else {
        v2[-856] = 42
    }
    v5 := GenS1{F0: uint16(v1)}
    delete(v2, -2147483648)
    for fi6 := 0; fi6 < 7; fi6++ {
        var v7 int32 = 1
        fmt.Println(v7)
        _ = fi6
        _ = v7
    }
    fmt.Println(v3, v1)
    if (2147483647 != 42) {
        if (!("\n\t" > "a")) {
            if false {
                v3 = (v3 >> ((0) & 63))
            }
        } else {
            var v8 bool = true
            var v9 int8 = -17
            if true {
                v3 = (v3 >> ((255) & 63))
                v9 = int8(v1)
            } else {
                v3 = (-(v3))
            }
            _ = v8
            _ = v9
        }
        fmt.Println(v1, v3)
        var v10 int32 = 0
        _ = v10
    }
    if false {
        var v11 bool = (42 == -9223372036854775808)
        v12 := GenS1{F0: 1}
        if v11 {
            var v13 float32 = 3.14159
            _ = v13
        }
        _ = v11
        _ = v12
    } else {
        v1 -= v1
        for fi14 := 0; fi14 < 12; fi14++ {
            if ((!false) && false) {
                var v15 uint16 = 42
                _ = v15
            }
            _ = fi14
        }
    }
    v1 = (v1 % ((v1) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
}
