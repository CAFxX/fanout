package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 string; F2 uint }
func genFn0(p0 string, p1 string) uint32 {
    p1 = ((p1 + (((p1 + "\n\t") + p1) + "foo bar")) + "a")
    fmt.Println(p1, p0)
    v7 := []int{-842, -403, -695, 0}
    v7 = append(v7, -90)
    var v8 string = "hello"
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 4294967295
}

func main() {
    var v1 int16 = 42
    v2 := make(map[uint32]uint32)
    var v3 int32 = -1
    var v4 bool = false
    var v5 uint8 = 127
    v6 := make(map[int]bool)
    fmt.Println(v4, v1)
    var v9 int32 = 515
    var v10 uint8 = (v5 | v5)
    var v11 uint16 = 65535
    v3 = v3
    fmt.Println(v5, v3)
    v12 := v6[-2147483648]
    v13 := GenS1{F0: 4294967295}
    fmt.Println(v11, v12, v4)
    v4 = (true || (v12 || true))
    v14 := GenS2{F0: v11, F1: "hello", F2: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
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
    _ = v9
}
