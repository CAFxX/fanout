package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 int16, p1 int32, p2 uint8) int {
    v7 := []int{-661, int(p1)}
    v7 = append(v7, -9223372036854775808)
    p1 -= int32(p2)
    switch -9223372036854775808 {
    case -435:
        var v8 bool = false
        _ = v8
    }
    p0 -= 0
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 0
}

func main() {
    var v1 uint8 = 255
    var v2 int64 = -1
    v3 := make(map[uint32]int)
    v4 := make(map[uint32]string)
    var v5 int = 0
    var v6 uint8 = v1
    var v9 uint32 = 2147483648
    v6 = (v1 >> ((65535) & 63))
    v2 += -1
    if true {
        var v10 int = v5
        _ = v10
    }
    v5 = v5
    var v11 uint = 413
    if (-2147483648 == 0) {
        v6 -= v1
        if (!(!(1 >= -2147483648))) {
            v11 -= v11
            v9 = 2147483648
            fmt.Println(v1, v2, v5)
        } else {
            var v12 int = v5
            var v13 uint16 = 1
            var v14 int32 = -111
            _ = v12
            _ = v13
            _ = v14
        }
    } else {
        var v15 int32 = -2147483648
        var v16 float64 = 3.14159
        _ = v15
        _ = v16
    }
    v1 = (v6 >> (((v11 % ((v11) | 1))) & 63))
    var v17 uint64 = 42
    fmt.Println(v1, v11, v9)
    v18 := GenS1{F0: v11}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
