package main

import (
    "fmt"
)


func main() {
    var v1 uint32 = 42
    v2 := []string{((("a" + "") + "a") + "foo bar"), "foo bar", "", "hello"}
    var v3 int8 = 1
    var v4 bool = true
    v1 = (v1 & v1)
    v1 -= (v1 << ((18446744073709551615) & 63))
    switch 757 {
    case -9223372036854775808:
        for fi5 := 0; fi5 < 20; fi5++ {
            v1 = v1
            _ = fi5
        }
    }
    v4 = (false || v4)
    v2 = append(v2, "hello")
    var v6 int64 = -1
    var v7 int8 = v3
    for fi8 := 0; fi8 < 9; fi8++ {
        var v9 int8 = (^(v7))
        _ = fi8
        _ = v9
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
