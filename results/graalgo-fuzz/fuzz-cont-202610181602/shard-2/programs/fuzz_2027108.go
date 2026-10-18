package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 int32, p1 int64) bool {
    switch "\n\t" {
    case "\n\t":
        var v5 int16 = 42
        _ = v5
    case "a":
        fmt.Println(p0)
    default:
        v6 := []int{2147483647, 9223372036854775807}
        _ = v6
    }
    if (false || (18446744073709551615 == 880)) {
        if (1e10 > -2.25) {
            var v7 float64 = 1.5
            _ = v7
        } else {
            var v8 bool = false
            _ = v8
        }
    } else {
        for fi9 := 0; fi9 < 13; fi9++ {
            var v10 string = "a"
            p0 -= p0
            _ = fi9
            _ = v10
        }
    }
    _ = p0
    _ = p1
    return true
}

func main() {
    v1 := make(map[int]string)
    var v2 bool = (!false)
    v3 := []string{("a" + "foo bar"), "hello", "hello"}
    v4 := []string{("a" + "hello")}
    v11 := GenS1{F0: 166}
    v2 = (42 > -1)
    for fi12 := 0; fi12 < 5; fi12++ {
        v2 = (!(v2 || (v2 || v2)))
        _ = fi12
    }
    for fi13 := 0; fi13 < 10; fi13++ {
        var v14 int8 = -24
        _ = fi13
        _ = v14
    }
    var v15 float64 = (((1.5 - 3.14159) * 1.5) / 3.14159)
    var v16 float32 = 1.5
    var v17 float64 = v15
    v16 = (v16 / 0.1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
