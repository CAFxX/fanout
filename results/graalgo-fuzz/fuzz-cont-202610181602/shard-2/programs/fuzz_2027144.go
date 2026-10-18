package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 int }
func genFn0(p0 uint32, p1 uint16, p2 bool) uint {
    v6 := GenS1{F0: 40, F1: 305, F2: 1}
    var v7 int = -2147483648
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 int, p1 uint32, p2 int32) bool {
    var v8 float64 = 3.14159
    p0 = 879
    if false {
        if ((1.5 + 0.0) >= (0.1 + ((float32(1e10) * 1.5) + 3.14159))) {
            var v9 uint16 = 255
            v10 := []int{}
            _ = v10
            _ = v9
        } else {
            for fi11 := 0; fi11 < 11; fi11++ {
                var v12 uint32 = 391
                fmt.Println(p2, v12)
                for fi13 := 0; fi13 < 17; fi13++ {
                    fmt.Println(fi13, p0, fi11)
                    var v14 int8 = 0
                    _ = fi13
                    _ = v14
                }
                _ = fi11
                _ = v12
            }
        }
        p2 = (p2 ^ int32(p0))
        fmt.Println(p1)
    } else {
        fmt.Println(p0)
    }
    var v15 bool = true
    var v16 float32 = 1e10
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v16
    _ = v8
    return true
}

func main() {
    var v1 int8 = 1
    var v2 float32 = (3.14159 / 3.14159)
    var v3 int64 = 9223372036854775807
    v4 := make(map[int]string)
    v5 := make(map[string]int)
    v3 = 946
    var v17 int32 = -2147483648
    var v18 string = "hello"
    v19 := []int{0, 905, 127, 87}
    v20 := v5[v18]
    v21 := GenS1{F0: (v1 * 0), F1: 983, F2: -2147483648}
    delete(v5, v18)
    if true {
        v22 := GenS1{F0: v1, F1: 42, F2: v20}
        fmt.Println(v2)
        _ = v22
    }
    v3 -= -83
    var v23 int64 = (^(v3))
    fmt.Println(v2, v2, v20)
    v19 = append(v19, int(v23))
    v23 += 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
