package main

import (
    "fmt"
)

func genFn0(p0 uint64) int {
    switch 4294967295 {
    case 127:
        p0 += (^(p0))
    case 1:
        var v5 float64 = 0.1
        _ = v5
    default:
        v6 := []int{1, -969}
        _ = v6
    }
    var v7 string = ("foo bar" + "hello")
    p0 -= 2
    _ = p0
    _ = v7
    return -804
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int64 = -9223372036854775808
    var v3 bool = (1 <= -1)
    var v4 int32 = -784
    v4 -= 2147483647
    v8 := []int{42, 462, 32767}
    fmt.Println(v2, v2, v2)
    if v3 {
        switch "" {
        case "\n\t":
            for fi9 := 0; fi9 < 11; fi9++ {
                v2 += v2
                var v10 int64 = v2
                fmt.Println(v3)
                _ = fi9
                _ = v10
            }
        case "hello":
            v2 = 902
        }
    }
    for fi11 := 0; fi11 < 2; fi11++ {
        v2 = v2
        _ = fi11
    }
    var v12 bool = v3
    v2 += -9223372036854775808
    fmt.Println(v3, v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
