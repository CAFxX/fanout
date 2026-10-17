package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 string, p1 bool, p2 uint32) int {
    var v6 uint16 = 65535
    v7 := GenS1{F0: 127}
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return -2147483648
}

func main() {
    var v1 uint64 = 708
    var v2 string = "\n\t"
    var v3 uint64 = 1
    var v4 float32 = (-2.25 - 0.0)
    var v5 uint = 255
    v8 := []int{}
    if false {
        v5 = v5
        if true {
            v3 = v1
            v1 = 4294967295
            var v9 bool = false
            _ = v9
        }
        v1 -= 2
    }
    var v10 uint = v5
    if false {
        var v11 float64 = float64(float32((0.1 - (1.5 * 1.5))))
        _ = v11
    }
    if false {
        v3 -= (v3 >> ((591) & 63))
        v1 -= v1
        var v12 int8 = 11
        _ = v12
    }
    var v13 uint32 = 2147483648
    var v14 uint8 = 124
    var v15 uint = 0
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
