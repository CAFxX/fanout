package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0() string {
    var v4 int = 32768
    v5 := []int{(v4 + -266), v4}
    fmt.Println(v4)
    _ = len(v5)
    _ = v4
    _ = v5
    return "a"
}

func main() {
    var v1 float32 = 1.5
    var v2 int16 = -32768
    var v3 uint = 886
    var v6 float64 = ((0.0 + 1.5) + float64((1e10 - 1.5)))
    v2 += 0
    var v7 float32 = v1
    var v8 int64 = -914
    var v9 uint32 = 0
    v8 -= 1
    v10 := GenS1{F0: -2147483648}
    v11 := GenS1{F0: 0}
    var v12 uint8 = 204
    v6 = float64(float32((1e10 + ((float64(v7) - -2.25) * 1.5))))
    var v13 int16 = 84
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
