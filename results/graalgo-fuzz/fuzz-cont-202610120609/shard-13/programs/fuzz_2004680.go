package main

import (
    "fmt"
)

func genFn0() int {
    var v5 uint16 = 27
    v6 := []int{-192}
    v5 -= v5
    for fi7 := 0; fi7 < 4; fi7++ {
        v5 = v5
        _ = fi7
    }
    _ = v5
    _ = v6
    return -1
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]int)
    v3 := make(map[int]uint32)
    var v4 int8 = -71
    fmt.Println(v4)
    v4 += (^(v4))
    switch "foo bar" {
    case "":
        if (v4 > (^(v4))) {
            v8 := []int{}
            _ = v8
        }
    case "a":
        v4 = v4
    }
    if false {
        v4 = v4
        v9 := []int{}
        _ = v9
    }
    v4 -= 100
    var v10 int64 = -913
    for fi11 := 0; fi11 < 10; fi11++ {
        if false {
            var v12 float64 = 0.1
            var v13 int16 = 0
            v10 += -842
            _ = v12
            _ = v13
        }
        v4 = v4
        v4 += 1
        _ = fi11
    }
    var v14 bool = false
    var v15 uint64 = 2
    v16 := []int{-9223372036854775808, genFn0(), 128, 2147483647}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
