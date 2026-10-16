package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 string; F1 uint32; F2 uint64 }

func main() {
    v1 := []int{}
    v2 := []string{("foo bar" + ("" + "foo bar")), "a", "\n\t", "hello"}
    var v3 uint16 = 32767
    var v4 float64 = (0.0 * (1e10 * ((1e10 / (3.14159 / 1e10)) - -2.25)))
    v5 := []string{}
    var v6 int8 = 1
    switch -161 {
    case -9223372036854775808:
        var v7 uint = 18446744073709551615
        _ = v7
    case 2147483647:
        var v8 uint16 = 0
        _ = v8
    default:
        v9 := GenS1{F0: 0, F1: 1}
        _ = v9
    }
    v10 := GenS2{F0: "foo bar", F1: 65535, F2: 2}
    for fi11 := 0; fi11 < 11; fi11++ {
        for fi12 := 0; fi12 < 11; fi12++ {
            fmt.Println(v6, fi11)
            _ = fi12
        }
        var v13 int32 = -1
        var v14 uint64 = 890
        _ = fi11
        _ = v13
        _ = v14
    }
    var v15 bool = (v6 > (^(v6)))
    v15 = v15
    v5 = append(v5, ("a" + "hello"))
    for fi16 := 0; fi16 < 9; fi16++ {
        var v17 float32 = 3.14159
        var v18 bool = (1 > -2147483648)
        var v19 string = "foo bar"
        _ = fi16
        _ = v17
        _ = v18
        _ = v19
    }
    if (0 == 42) {
        var v20 uint32 = 65535
        _ = v20
    } else {
        v15 = (!(v15 || true))
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
