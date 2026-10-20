package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 uint8 }
func genFn0(p0 int8) string {
    var v4 bool = false
    if (v4 && ((9223372036854775807 <= -9223372036854775808) && v4)) {
        for fi5 := 0; fi5 < 8; fi5++ {
            var v6 int16 = 47
            _ = fi5
            _ = v6
        }
    }
    var v7 int8 = -128
    _ = p0
    _ = v4
    _ = v7
    return (("" + "a") + "foo bar")
}

func main() {
    var v1 int8 = -123
    var v2 int16 = -853
    var v3 float64 = 1.5
    var v8 uint32 = 358
    switch "foo bar" {
    case "foo bar":
        v9 := []int{127}
        _ = v9
    }
    v10 := []int{199, int(v1), 9223372036854775807, 1}
    var v11 int16 = 0
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    if (127 >= 128) {
        fmt.Println(v3, v8, v3)
        if false {
            v11 += (v11 - 0)
            var v12 int = -235
            _ = v12
        } else {
            var v13 float32 = 3.14159
            var v14 int16 = (v2 * v2)
            _ = v13
            _ = v14
        }
        if (181 <= 0) {
            v1 = (v1 % (((v1 << ((73) & 63))) | 1))
            v15 := GenS1{F0: 371, F1: 60, F2: 149}
            var v16 uint = 0
            _ = v15
            _ = v16
        }
    } else {
        v11 = (-(v2))
        v2 = 1
    }
    v11 -= v11
    v11 -= v11
    var v17 string = "hello"
    v10 = append(v10, -280)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v2
    _ = v3
    _ = v8
}
