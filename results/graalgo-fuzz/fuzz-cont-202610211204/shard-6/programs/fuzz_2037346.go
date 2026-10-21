package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]uint32)
    var v3 string = "hello"
    v3 = ("" + "foo bar")
    var v4 uint32 = 2147483648
    if false {
        v5 := GenS1{F0: (1 > -10)}
        _ = v5
    }
    fmt.Println(v4, v4)
    var v6 uint32 = (v4 / ((2147483648) | 1))
    var v7 int16 = 604
    for fi8 := 0; fi8 < 14; fi8++ {
        fmt.Println(v4, v7)
        v9 := GenS1{F0: true}
        _ = fi8
        _ = v9
    }
    if false {
        delete(v2, 250)
        for fi10 := 0; fi10 < 5; fi10++ {
            var v11 string = v3
            if false {
                v3 = "\n\t"
                v12 := []int{(-(fi10)), fi10}
                _ = v12
            }
            fmt.Println(fi10, v11)
            _ = fi10
            _ = v11
        }
        fmt.Println(v4, v6)
    }
    var v13 int64 = -653
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
