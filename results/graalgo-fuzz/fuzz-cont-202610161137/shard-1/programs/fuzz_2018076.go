package main

import (
    "fmt"
)


func main() {
    var v1 int = 1
    var v2 int32 = 2147483647
    var v3 bool = (((v1 + -1) > (v1 & -2147483648)) || (-2147483648 >= v2))
    v4 := make(map[uint32]int)
    delete(v4, 567)
    var v5 bool = ((!(-914 < -9223372036854775808)) || (42 != 2147483647))
    var v6 int8 = -1
    v6 = v6
    v6 = (^(v6))
    fmt.Println(v3)
    v6 = v6
    v2 = 0
    v7 := []int{127, 32768, int(v2), 9223372036854775807}
    var v8 string = ""
    fmt.Println(v5, v8, v3)
    var v9 float32 = (1.5 + 1.5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
