package main

import (
    "fmt"
)

func genFn0(p0 uint, p1 string) uint8 {
    var v4 int16 = -528
    switch p1 {
    case "hello":
        v5 := []int{0}
        _ = v5
    }
    _ = p0
    _ = p1
    _ = v4
    return 94
}
func genFn1() uint64 {
    v6 := []int{-459, -1, 0, -201}
    var v7 bool = false
    _ = v6
    _ = v7
    return 18446744073709551615
}

func main() {
    var v1 uint8 = 212
    var v2 uint32 = 179
    v3 := make(map[uint32]int)
    var v8 uint32 = v2
    var v9 uint32 = (^(v8))
    v8 -= 819
    var v10 uint16 = 2
    var v11 uint8 = 42
    fmt.Println(v9, v1)
    var v12 float32 = (1.5 / 1.5)
    var v13 int8 = 127
    v3[(^(v2))] = -2147483648
    if (!((42 > 42) && true)) {
        v14 := []int{}
        _ = v14
    }
    var v15 string = ("foo bar" + ("hello" + ("" + "")))
    var v16 float32 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
