package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool }
func genFn0() uint {
    v7 := []int{454}
    var v8 uint64 = 42
    v9 := GenS1{F0: 122}
    _ = v7
    _ = v8
    _ = v9
    return 255
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int32 = -2147483648
    var v3 bool = (42 > 42)
    var v4 uint32 = 496
    var v5 int = 181
    var v6 int8 = -66
    var v10 uint16 = 921
    v11 := v1[v4]
    v12 := GenS2{F0: (1.5 < (3.14159 - 3.14159))}
    v11 = "a"
    v2 += (-(v2))
    var v13 int = -52
    var v14 int = v13
    var v15 uint64 = 1
    fmt.Println(v14, v15, v13)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
