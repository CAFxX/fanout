package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 string) int32 {
    v6 := GenS1{F0: -1}
    var v7 int32 = 408
    v7 = (v7 & v7)
    _ = p0
    _ = v6
    _ = v7
    return v7
}

func main() {
    v1 := []int{128, 311, 2147483647}
    var v2 uint16 = 255
    var v3 float32 = (3.14159 + 1e10)
    var v4 bool = true
    var v5 uint32 = 4294967295
    v5 = 942
    fmt.Println(v4, v5, v5)
    if ((!false) || ("foo bar" != "\n\t")) {
        var v8 int16 = 1000
        _ = v8
    }
    v5 -= 517
    v5 -= v5
    v3 = float32(0.1)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v5, v4)
    v5 = (^(v5))
    if (v4 && (4294967295 > 18446744073709551615)) {
        v1 = append(v1, 4294967295)
        var v9 string = ("\n\t" + (((("foo bar" + "a") + "\n\t") + "a") + "\n\t"))
        _ = v9
    }
    v10 := GenS1{F0: -1}
    v5 -= v5
    var v11 uint16 = v2
    switch -2147483648 {
    case -538:
        v3 = v3
    default:
        fmt.Println(v11)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
