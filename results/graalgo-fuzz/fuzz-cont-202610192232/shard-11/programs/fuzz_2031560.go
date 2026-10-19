package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 uint16, p1 string, p2 uint64) int16 {
    v6 := []int{585}
    switch -846 {
    case 0:
        var v7 uint64 = 779
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 1000
}

func main() {
    var v1 int = -372
    var v2 int = (v1 << ((499) & 63))
    v3 := make(map[uint32]int)
    v4 := make(map[uint32]bool)
    var v5 uint = 1
    _ = len(v3)
    v1 = -1
    v2 -= 32768
    for fi8 := 0; fi8 < 18; fi8++ {
        switch (fi8 + 2147483647) {
        case 988:
            var v9 int = 9223372036854775807
            _ = v9
        default:
            var v10 float64 = 1e10
            _ = v10
        }
        if true {
            var v11 float32 = (((0.1 + 1e10) / 1.0) - 3.14159)
            var v12 int = v1
            v13 := GenS1{F0: 1}
            _ = v11
            _ = v12
            _ = v13
        } else {
            v5 = (v5 ^ v5)
            var v14 uint = (v5 - 919)
            _ = len(v4)
            _ = v14
        }
        _ = fi8
    }
    v15 := GenS1{F0: 1}
    fmt.Println(v5)
    v2 = v1
    var v16 uint8 = 127
    var v17 bool = (!(42 >= -9223372036854775808))
    for fi18 := 0; fi18 < 14; fi18++ {
        var v19 uint64 = 108
        var v20 float64 = (0.0 - 0.1)
        var v21 int = fi18
        _ = fi18
        _ = v19
        _ = v20
        _ = v21
    }
    v17 = false
    var v22 int64 = 9223372036854775807
    v23 := []int{42, v2, -327}
    v1 -= 566
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
