package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 uint32 = 42
    v2 := make(map[uint32]bool)
    v3 := make(map[int]bool)
    for fi4 := 0; fi4 < 11; fi4++ {
        _ = len(v2)
        v5 := GenS1{F0: 4294967295}
        _ = fi4
        _ = v5
    }
    if (2147483647 != 1) {
        v1 = v1
        v1 = 2147483648
    }
    switch ("\n\t" + "\n\t") {
    case "a":
        v6 := v3[128]
        _ = v6
    case "":
        v7 := GenS1{F0: 799}
        _ = v7
    }
    v8 := []int{-745, 9223372036854775807, 32767}
    switch "" {
    case "a":
        var v9 int32 = -1
        _ = v9
    default:
        switch -599 {
        case 0:
            fmt.Println(v1)
        default:
            var v10 int8 = 100
            _ = v10
        }
    }
    var v11 int64 = 9223372036854775807
    for fi12 := 0; fi12 < 16; fi12++ {
        switch ("" + "hello") {
        case "\n\t":
            v1 += v1
        case "foo bar":
            v1 += v1
        }
        _ = fi12
    }
    v13 := GenS1{F0: 1}
    if len(v8) > 0 {
        v8[0] = 42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v8
}
