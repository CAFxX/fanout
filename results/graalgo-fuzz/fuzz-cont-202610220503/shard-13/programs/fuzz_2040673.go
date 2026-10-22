package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 int8 }
type GenS2 struct { F0 int }
func genFn0() int32 {
    var v4 float64 = 1.5
    var v5 string = "hello"
    _ = v4
    _ = v5
    return 42
}

func main() {
    var v1 int32 = -852
    v2 := make(map[string]bool)
    var v3 bool = ("foo bar" != "a")
    if false {
        if v3 {
            v3 = (!(!v3))
            v6 := []int{-932}
            var v7 float64 = 1.5
            _ = v6
            _ = v7
        } else {
            var v8 int = 42
            for fi9 := 0; fi9 < 10; fi9++ {
                var v10 string = (("a" + "") + "foo bar")
                _ = fi9
                _ = v10
            }
            v3 = v3
            _ = v8
        }
        var v11 uint = 42
        v3 = false
        _ = v11
    }
    var v12 int32 = v1
    for fi13 := 0; fi13 < 5; fi13++ {
        v14 := v2[("hello" + "foo bar")]
        v1 = v12
        _ = fi13
        _ = v14
    }
    v2["a"] = (!(!(!false)))
    v15 := v2[("" + "\n\t")]
    switch 187 {
    case 2147483647:
        v16 := GenS1{F0: -1, F1: -1, F2: -128}
        _ = v16
    default:
        v17 := GenS2{F0: 224}
        _ = v17
    }
    fmt.Println(v3, v15, v1)
    v18 := []int{120, -158, 7}
    if v3 {
        v15 = v3
        v19 := GenS1{F0: 1, F1: genFn0(), F2: 127}
        var v20 uint = 227
        _ = v19
        _ = v20
    } else {
        v21 := GenS2{F0: 256}
        fmt.Println(v3, v3)
        _ = v21
    }
    for fi22 := 0; fi22 < 7; fi22++ {
        v18 = append(v18, fi22)
        v23 := GenS2{F0: (fi22 << ((18446744073709551615) & 63))}
        _ = fi22
        _ = v23
    }
    v3 = v15
    var v24 uint = 65535
    v25 := GenS2{F0: -2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v15
    _ = v18
    _ = v2
    _ = v24
    _ = v25
    _ = v3
}
