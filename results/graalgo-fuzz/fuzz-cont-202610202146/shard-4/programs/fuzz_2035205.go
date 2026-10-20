package main

import (
    "fmt"
)

func genFn0(p0 uint8, p1 int16) uint8 {
    for fi4 := 0; fi4 < 2; fi4++ {
        p1 -= p1
        _ = fi4
    }
    v5 := []int{}
    _ = len(v5)
    _ = p0
    _ = p1
    _ = v5
    return 28
}

func main() {
    var v1 uint16 = 2
    var v2 bool = false
    var v3 int32 = 1
    for fi6 := 0; fi6 < 11; fi6++ {
        if (v1 <= (v1 * (v1 >> ((1) & 63)))) {
            v7 := []int{fi6}
            v2 = v2
            _ = v7
        } else {
            v3 += (v3 / (((v3 + (-(v3)))) | 1))
        }
        v8 := []int{}
        _ = fi6
        _ = v8
    }
    var v9 int32 = (v3 & v3)
    v9 = (v3 * (v9 & (v9 + (v9 & (-(v3))))))
    var v10 float32 = 0.0
    v11 := []int{128, -289}
    if len(v11) > 0 {
        v11[0] = -9223372036854775808
    }
    fmt.Println(v10)
    v11 = append(v11, 42)
    fmt.Println(v10, v3)
    v1 -= v1
    if len(v11) > 0 {
        v11[0] = 755
    }
    var v12 uint16 = (v1 << ((218) & 63))
    var v13 uint16 = 665
    var v14 uint = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v9
}
