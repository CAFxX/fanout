package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 int8 = 100
    var v2 float64 = 0.1
    var v3 float32 = 1.5
    v4 := make(map[string]bool)
    var v5 uint = 18446744073709551615
    var v6 int = 32767
    if true {
        if (v2 != v2) {
            if false {
                if (-32768 != 310) {
                    var v7 int16 = -838
                    fmt.Println(v3)
                    v1 += v1
                    _ = v7
                } else {
                    v8 := []int{(-(v6)), (v6 % ((v6) | 1)), v6, -596}
                    if false {
                        v9 := GenS1{F0: 973}
                        _ = v9
                    }
                    _ = v8
                }
            }
            if (1 == 472) {
                fmt.Println(v5)
                var v10 int64 = 110
                fmt.Println(v2, v5, v3)
                _ = v10
            } else {
                v11 := GenS1{F0: 984}
                _ = v11
            }
        }
    }
    v3 = 1.5
    var v12 int8 = v1
    var v13 string = "hello"
    v5 -= (v5 ^ (v5 % ((v5) | 1)))
    v14 := GenS1{F0: 2147483647}
    var v15 int32 = 835
    v1 -= (v12 | v12)
    var v16 int16 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
