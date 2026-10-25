package main

import (
    "fmt"
)


func main() {
    var v1 uint8 = 2
    v2 := make(map[int]string)
    var v3 float64 = float64(1.5)
    var v4 uint = 2
    var v5 float64 = 3.14159
    var v6 uint = (v4 << (((^(v4))) & 63))
    var v7 uint32 = 2147483647
    v1 = (v1 * 141)
    switch ("\n\t" + "") {
    case "a":
        v5 = 1e10
    case "":
        fmt.Println(v6, v5, v6)
    }
    for fi8 := 0; fi8 < 13; fi8++ {
        v9 := []int{}
        _ = fi8
        _ = v9
    }
    v10 := []int{}
    fmt.Println(v5, v3, v5)
    v3 -= v3
    for fi11 := 0; fi11 < 1; fi11++ {
        for fi12 := 0; fi12 < 5; fi12++ {
            var v13 int32 = -2147483648
            delete(v2, (^(fi11)))
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
