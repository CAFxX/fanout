package main

import (
    "fmt"
)


func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]bool)
    v3 := make(map[string]bool)
    var v4 string = "hello"
    var v5 uint16 = 207
    v6 := make(map[uint32]uint32)
    switch 42 {
    case -2147483648:
        v4 = (v4 + (v4 + v4))
    case -645:
        _ = len(v2)
    }
    fmt.Println(v5)
    var v7 uint8 = 128
    fmt.Println(v7, v5)
    var v8 int16 = 1000
    switch (v4 + v4) {
    case "":
        v3[v4] = false
    }
    for fi9 := 0; fi9 < 2; fi9++ {
        for fi10 := 0; fi10 < 19; fi10++ {
            var v11 bool = true
            delete(v6, 2147483647)
            var v12 uint64 = 844
            v4 = v4
            v5 -= 418
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi9
    }
    var v13 int64 = 9223372036854775807
    v14 := []int{int(v13), -9223372036854775808, -1}
    var v15 uint32 = 65535
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
