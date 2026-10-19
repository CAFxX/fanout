package main

import (
    "fmt"
)

func genFn0() int {
    for fi5 := 0; fi5 < 13; fi5++ {
        fmt.Println(fi5)
        var v6 uint64 = 42
        _ = fi5
        _ = v6
    }
    var v7 float32 = 3.14159
    _ = v7
    return 32767
}

func main() {
    var v1 string = ("a" + ("" + ("hello" + "foo bar")))
    v2 := make(map[int]string)
    v3 := make(map[uint32]bool)
    v4 := make(map[int]bool)
    v4[9223372036854775807] = false
    var v8 float64 = 0.1
    fmt.Println(v8, v8)
    v1 = v1
    var v9 uint16 = 388
    v10 := []int{genFn0(), 32767, -2147483648, -426}
    fmt.Println(v1)
    var v11 int16 = -32768
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
