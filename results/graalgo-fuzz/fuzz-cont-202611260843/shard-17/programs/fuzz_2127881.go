package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 string) int16 {
    fmt.Println(p0)
    v6 := GenS1{F0: -763}
    var v7 int16 = -1
    _ = p0
    _ = v6
    _ = v7
    return 1000
}

func main() {
    var v1 uint32 = 65535
    v2 := []int{678, 1, int(v1)}
    var v3 uint16 = 2
    var v4 uint = 80
    var v5 bool = (true && true)
    v4 = 0
    var v8 float64 = 1.5
    fmt.Println(v8, v4)
    _ = len(v2)
    var v9 float64 = (float64((3.14159 * (float32((float64(float32(float64(1.5))) / v8)) - 1.5))) - ((((0.1 - (-2.25 - v8)) - 3.14159) * 1e10) + 0.0))
    v3 += (v3 & 823)
    switch "hello" {
    case "":
        var v10 uint = (v4 & v4)
        _ = v10
    case "\n\t":
        var v11 uint = (v4 * v4)
        _ = v11
    default:
        v5 = (v5 || ((42 != 1) || (v5 || (2 == 240))))
    }
    if true {
        v1 += (^(v1))
    } else {
        if true {
            var v12 uint32 = (v1 | (^(v1)))
            v5 = false
            v1 -= (v1 % ((v1) | 1))
            _ = v12
        } else {
            _ = len(v2)
            v1 -= 0
        }
        fmt.Println(v4, v1)
    }
    fmt.Println(v8, v5, v3)
    var v13 float32 = -2.25
    if v5 {
        if false {
            var v14 uint8 = 105
            var v15 string = "hello"
            _ = v14
            _ = v15
        }
    } else {
        var v16 int8 = -65
        var v17 float32 = v13
        var v18 bool = (9223372036854775807 == 0)
        _ = v16
        _ = v17
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
