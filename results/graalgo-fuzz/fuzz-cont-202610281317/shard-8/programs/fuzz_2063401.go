package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 uint64, p1 bool) uint8 {
    if (65535 < 255) {
        var v6 string = "foo bar"
        p1 = p1
        p1 = (-1 >= -9223372036854775808)
        _ = v6
    } else {
        for fi7 := 0; fi7 < 4; fi7++ {
            p0 += 18446744073709551615
            _ = fi7
        }
    }
    v8 := GenS1{F0: (true || (!p1))}
    p0 += (^(p0))
    var v9 uint16 = 982
    v10 := []int{}
    _ = p0
    _ = p1
    _ = v10
    _ = v8
    _ = v9
    return 141
}
func genFn1() int32 {
    var v11 int64 = 418
    v12 := []int{707, -2147483648, 2147483647}
    v11 = (v11 ^ v11)
    _ = v11
    _ = v12
    return -1
}

func main() {
    v1 := []int{128, 32767}
    var v2 int = 127
    var v3 uint16 = 42
    var v4 uint16 = (^(v3))
    var v5 int64 = 796
    v2 += 455
    var v13 uint64 = 751
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v14 int32 = -2147483648
    v2 += (v2 >> ((0) & 63))
    v2 -= -1
    v15 := GenS1{F0: false}
    v16 := GenS1{F0: false}
    fmt.Println(v4, v4, v2)
    v3 += v4
    fmt.Println(v4, v3)
    fmt.Println(v13, v4, v14)
    v1 = append(v1, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
