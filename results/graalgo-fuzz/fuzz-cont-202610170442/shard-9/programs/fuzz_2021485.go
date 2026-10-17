package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }

func main() {
    var v1 bool = (-2147483648 >= -1)
    var v2 uint8 = 0
    var v3 uint8 = 42
    v3 = 0
    v4 := []int{127, 352}
    var v5 int32 = 0
    if false {
        v5 = v5
        switch "" {
        case "":
            if (1 > 0) {
                fmt.Println(v2, v2, v5)
                v1 = v1
                if v1 {
                    v6 := GenS1{F0: (0 < 2147483647), F1: 65535}
                    var v7 int32 = v5
                    _ = v6
                    _ = v7
                }
            }
        }
    }
    var v8 int8 = 53
    var v9 string = "a"
    v8 += (v8 << ((1) & 63))
    v4 = append(v4, -406)
    fmt.Println(v9, v2, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
