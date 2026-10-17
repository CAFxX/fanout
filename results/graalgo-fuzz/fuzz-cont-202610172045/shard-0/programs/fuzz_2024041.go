package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
func genFn0(p0 uint16) bool {
    v5 := []int{-299, 127, -2147483648}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v6 int = 2147483647
    if (false || false) {
        var v7 uint64 = 18446744073709551615
        var v8 float64 = 1.5
        _ = v7
        _ = v8
    }
    _ = p0
    _ = v5
    _ = v6
    return (false || false)
}

func main() {
    var v1 int32 = -2147483648
    v2 := []string{"hello", "foo bar", "foo bar"}
    var v3 int = 32767
    var v4 string = ""
    var v9 string = ((v4 + "hello") + "foo bar")
    switch v3 {
    case 127:
        if len(v2) > 0 {
            v2[0] = v4
        }
    default:
        v10 := GenS1{F0: 255, F1: 255}
        _ = v10
    }
    var v11 uint16 = 255
    var v12 string = v9
    var v13 int8 = -1
    var v14 int32 = v1
    if true {
        var v15 int32 = 2147483647
        var v16 bool = true
        v15 = v1
        _ = v15
        _ = v16
    }
    if true {
        v3 = (v3 << ((141) & 63))
        var v17 uint = 65535
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
