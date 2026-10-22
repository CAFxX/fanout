package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
func genFn0() uint {
    v6 := GenS1{F0: 36, F1: 4294967295}
    switch "hello" {
    case "":
        var v7 float32 = (-2.25 / 1.5)
        _ = v7
    }
    var v8 string = ("foo bar" + "")
    v9 := []int{4294967295, -256, 243}
    _ = v6
    _ = v8
    _ = v9
    return 255
}

func main() {
    var v1 uint32 = 4294967295
    v2 := make(map[int]string)
    v3 := make(map[uint32]int)
    v4 := []int{0}
    var v5 uint64 = 3
    var v10 uint64 = v5
    for fi11 := 0; fi11 < 5; fi11++ {
        v10 -= 16
        var v12 float32 = 3.14159
        _ = fi11
        _ = v12
    }
    v2[42] = "hello"
    fmt.Println(v1, v10, v1)
    v13 := GenS1{F0: v1, F1: -9223372036854775808}
    var v14 int32 = -103
    v14 += (v14 & (v14 - (v14 * -2147483648)))
    var v15 uint8 = 127
    switch -613 {
    case -1:
        v1 += 65535
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
