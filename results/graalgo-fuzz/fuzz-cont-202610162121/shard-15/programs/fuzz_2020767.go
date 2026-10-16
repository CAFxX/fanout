package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    v1 := make(map[int]string)
    var v2 float64 = (1e10 + (1e10 * ((0.1 - ((((3.14159 + (0.1 * (3.14159 - 1e10))) * (-2.25 + ((1.5 - (1e10 + 3.14159)) - -2.25))) - -2.25) / 0.1)) / 1.0)))
    var v3 int8 = 1
    var v4 int32 = -2147483648
    _ = len(v1)
    var v5 string = ""
    delete(v1, 42)
    var v6 string = (v5 + v5)
    v6 = v6
    v7 := GenS1{F0: 94}
    var v8 float64 = v2
    fmt.Println(v5, v4, v3)
    var v9 int64 = 1
    v6 = (v6 + v6)
    fmt.Println(v6, v2)
    var v10 int8 = v3
    if (false && true) {
        v10 += 1
        v11 := GenS1{F0: 2}
        if (42 != 1) {
            v8 -= 0.1
        }
        _ = v11
    }
    var v12 float32 = -2.25
    v4 = 2147483647
    fmt.Println(v10)
    v12 += (-2.25 + (v12 - v12))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
