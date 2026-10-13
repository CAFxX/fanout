package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 uint8, p1 bool, p2 bool) int {
    switch "\n\t" {
    case "":
        v6 := []int{1, 128, -13, -699}
        _ = v6
    }
    p0 -= 127
    v7 := []int{-2147483648, -9223372036854775808, 127, 9223372036854775807}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return -148
}

func main() {
    var v1 uint32 = 2147483647
    v2 := make(map[string]int)
    var v3 uint8 = 2
    v4 := make(map[int]string)
    var v5 uint = 255
    switch 2147483647 {
    case 32768:
        v5 = v5
    }
    v1 = v1
    v3 += (v3 << ((v5) & 63))
    var v8 int64 = -1
    v5 = (^(v5))
    var v9 int32 = 0
    fmt.Println(v3, v3)
    var v10 int64 = (-(v8))
    fmt.Println(v10, v10)
    var v11 bool = false
    for fi12 := 0; fi12 < 8; fi12++ {
        v8 += v10
        v13 := GenS1{F0: (fi12 << (((v5 << ((v5) & 63))) & 63))}
        _ = fi12
        _ = v13
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
