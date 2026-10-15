package main

import (
    "fmt"
)


func main() {
    var v1 int = 9223372036854775807
    var v2 int8 = -63
    var v3 string = ("a" + (("\n\t" + "foo bar") + ""))
    var v4 uint64 = 1
    v5 := []int{v1, (v1 % ((1) | 1))}
    v4 = v4
    for fi6 := 0; fi6 < 1; fi6++ {
        v2 = (v2 >> ((18446744073709551615) & 63))
        var v7 int8 = v2
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 2; fi8++ {
        v4 -= 711
        _ = fi8
    }
    v2 = 1
    var v9 int32 = 0
    fmt.Println(v9)
    v2 -= v2
    fmt.Println(v3, v9, v9)
    v9 = v9
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
