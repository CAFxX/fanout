package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 uint64, p1 uint16) int16 {
    var v5 int8 = int8(p1)
    var v6 uint64 = (p0 + (^(p0)))
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 1000
}

func main() {
    var v1 uint16 = 0
    var v2 uint = 65535
    var v3 int8 = -128
    var v4 float64 = 1e10
    if false {
        var v7 int64 = 1
        _ = v7
    }
    for fi8 := 0; fi8 < 19; fi8++ {
        v9 := GenS1{F0: (v2 << ((v2) & 63))}
        var v10 string = "foo bar"
        v11 := GenS1{F0: (^(v2))}
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    var v12 int16 = genFn0(393, (^(v1)))
    if false {
        var v13 int = 4294967295
        _ = v13
    } else {
        v14 := []int{-159, 2147483647}
        _ = v14
    }
    switch "hello" {
    case "\n\t":
        var v15 uint16 = (v1 << (((v2 + (v2 | (v2 >> ((v2) & 63))))) & 63))
        _ = v15
    default:
        v16 := []int{0}
        _ = v16
    }
    var v17 bool = (v3 < (v3 ^ (v3 + (v3 >> ((42) & 63)))))
    var v18 uint64 = 18446744073709551615
    var v19 bool = ((3.14159 / v4) == v4)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
