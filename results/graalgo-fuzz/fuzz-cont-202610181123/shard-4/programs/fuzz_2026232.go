package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 uint, p1 bool, p2 int) bool {
    v6 := GenS1{F0: 32767}
    p0 += 2
    var v7 uint = 255
    var v8 float32 = (1e10 * 0.1)
    v9 := []int{p2}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return (-2147483648 == 42)
}
func genFn1(p0 uint16, p1 int32) int {
    var v10 string = ""
    v10 = v10
    _ = p0
    _ = p1
    _ = v10
    return 32768
}

func main() {
    var v1 float64 = 3.14159
    var v2 uint32 = 191
    var v3 int = 42
    var v4 float64 = float64(((((1.5 + 1.5) / (((((3.14159 - 1.5) / 0.1) * 0.0) - (3.14159 - 1e10)) / 1e10)) + 0.1) + 0.1))
    var v5 bool = false
    if (42 >= -662) {
        v11 := GenS1{F0: 2}
        fmt.Println(v1, v4, v1)
        v3 = 4294967295
        _ = v11
    }
    v3 = 127
    v5 = (!v5)
    if false {
        v2 -= v2
        var v12 float32 = (1.5 * (1.5 + 3.14159))
        for fi13 := 0; fi13 < 17; fi13++ {
            var v14 float64 = (((v4 + 0.1) - 0.0) * 1e10)
            v14 = v4
            v15 := []int{(fi13 + fi13), (v3 / ((106) | 1))}
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = v12
    }
    var v16 string = ""
    v17 := []int{-156}
    v18 := GenS1{F0: 0}
    fmt.Println(v5, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
