package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 int32) uint16 {
    p0 += -1
    v4 := []int{-184}
    var v5 float64 = float64((-2.25 - ((1e10 + -2.25) + float32((3.14159 / -2.25)))))
    var v6 uint = 437
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 255
}

func main() {
    var v1 uint16 = 65535
    var v2 int8 = -17
    var v3 bool = true
    var v7 float32 = (-2.25 * 1.5)
    for fi8 := 0; fi8 < 1; fi8++ {
        var v9 int32 = -2147483648
        _ = fi8
        _ = v9
    }
    fmt.Println(v2)
    var v10 bool = v3
    if v3 {
        if ((0 >= -1) && true) {
            fmt.Println(v7, v7, v2)
        }
        v3 = v10
    } else {
        var v11 string = ("foo bar" + "hello")
        _ = v11
    }
    fmt.Println(v1)
    if false {
        for fi12 := 0; fi12 < 17; fi12++ {
            v13 := GenS1{F0: genFn0(0)}
            var v14 uint16 = (v1 >> ((2) & 63))
            v1 += (v1 * v14)
            _ = fi12
            _ = v13
            _ = v14
        }
        var v15 float32 = (v7 + v7)
        v16 := []int{-820}
        _ = v15
        _ = v16
    }
    var v17 int = 32767
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v3
    _ = v7
}
