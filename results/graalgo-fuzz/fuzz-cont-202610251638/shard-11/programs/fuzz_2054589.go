package main

import (
    "fmt"
)


func main() {
    v1 := []int{-623, 321, 127}
    var v2 int64 = 0
    v3 := []int{127, 1}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v2 -= v2
    fmt.Println(v2)
    switch (("" + "") + "hello") {
    case "hello":
        fmt.Println(v2)
    default:
        var v4 uint32 = 347
        _ = v4
    }
    v2 -= 533
    if len(v1) > 0 {
        v1[0] = 189
    }
    if (42 <= 183) {
        var v5 uint8 = 42
        v2 -= (^(v2))
        for fi6 := 0; fi6 < 2; fi6++ {
            var v7 int = -949
            _ = fi6
            _ = v7
        }
        _ = v5
    }
    if (1 <= 4294967295) {
        fmt.Println(v2)
    } else {
        v2 += (v2 % ((-9223372036854775808) | 1))
        fmt.Println(v2)
        for fi8 := 0; fi8 < 15; fi8++ {
            _ = len(v3)
            _ = fi8
        }
    }
    var v9 uint16 = 32767
    switch 9223372036854775807 {
    case 32768:
        var v10 uint64 = 4294967295
        _ = v10
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v9
}
