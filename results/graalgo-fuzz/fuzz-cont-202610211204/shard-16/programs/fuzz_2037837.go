package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0() uint {
    v6 := GenS1{F0: (18446744073709551615 == 18446744073709551615)}
    var v7 uint16 = 388
    _ = v6
    _ = v7
    return 1
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]string)
    var v3 int64 = 9223372036854775807
    v4 := make(map[string]uint32)
    var v5 bool = true
    switch "foo bar" {
    case "hello":
        fmt.Println(v3)
    }
    v5 = false
    switch "foo bar" {
    case "hello":
        if false {
            v8 := []int{}
            _ = v8
        }
    }
    var v9 uint64 = 0
    if v5 {
        fmt.Println(v5, v9, v3)
    }
    switch "\n\t" {
    case "a":
        var v10 int = 4294967295
        _ = v10
    case "foo bar":
        fmt.Println(v5, v9)
    default:
        switch "hello" {
        case "a":
            var v11 int16 = -32768
            _ = v11
        }
    }
    var v12 uint64 = v9
    fmt.Println(v5)
    v3 -= (v3 * 225)
    v3 += (v3 * (v3 / ((v3) | 1)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
