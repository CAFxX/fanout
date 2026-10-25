package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 int16, p1 uint32, p2 int8) int16 {
    fmt.Println(p1)
    var v7 string = "a"
    var v8 uint = 255
    var v9 int16 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return (v9 | 876)
}

func main() {
    var v1 float32 = 1e10
    v2 := make(map[string]int)
    var v3 uint8 = 127
    var v4 int = -2147483648
    var v5 uint16 = 0
    v6 := make(map[uint32]bool)
    v1 += 1.5
    for fi10 := 0; fi10 < 17; fi10++ {
        var v11 int16 = -211
        if (fi10 < fi10) {
            v5 = 65535
            var v12 string = "\n\t"
            _ = v12
        }
        v1 += float32(0.1)
        _ = fi10
        _ = v11
    }
    v3 = 255
    var v13 int64 = 364
    v14 := GenS1{F0: true}
    var v15 int64 = v13
    if ((2147483647 != v4) && true) {
        switch 2147483647 {
        case 987:
            v3 += (v3 << ((986) & 63))
        default:
            var v16 float64 = 0.1
            _ = v16
        }
        v17 := []int{v4, 512, (v4 >> ((763) & 63)), v4}
        v5 = v5
        _ = v17
    }
    fmt.Println(v4, v1)
    v3 += v3
    v18 := GenS1{F0: false}
    for fi19 := 0; fi19 < 12; fi19++ {
        fmt.Println(v13, v1)
        _ = fi19
    }
    v1 -= 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
