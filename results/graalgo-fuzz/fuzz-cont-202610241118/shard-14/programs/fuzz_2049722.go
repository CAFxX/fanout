package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int }
func genFn0(p0 bool, p1 int32) string {
    var v7 int16 = 32767
    fmt.Println(p1, p0, p0)
    v8 := []int{42}
    var v9 int16 = -398
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return "hello"
}

func main() {
    v1 := make(map[string]int)
    var v2 int32 = 0
    var v3 float64 = (0.0 / (float64((3.14159 * 0.1)) * 0.1))
    var v4 uint64 = 1
    var v5 uint = 2
    v6 := make(map[string]int)
    var v10 uint = v5
    fmt.Println(v5, v2)
    if (!(1e10 <= (float32(v3) / (0.0 + (float32(3.14159) * float32(v3)))))) {
        fmt.Println(v10, v5, v4)
    } else {
        v2 = (^(v2))
        v10 = uint(v4)
    }
    v11 := GenS1{F0: 255}
    fmt.Println(v4, v10)
    var v12 uint = (v10 / (((^(v5))) | 1))
    v12 = (v12 << ((362) & 63))
    var v13 int32 = 2147483647
    delete(v6, "foo bar")
    v14 := GenS2{F0: 4294967295}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
