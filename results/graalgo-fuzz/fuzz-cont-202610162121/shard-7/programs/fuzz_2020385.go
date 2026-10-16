package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int }
func genFn0(p0 int64, p1 uint16) int64 {
    p1 -= p1
    var v4 uint16 = p1
    _ = p0
    _ = p1
    _ = v4
    return (p0 + -9223372036854775808)
}

func main() {
    var v1 float32 = float32((0.0 + float64(1e10)))
    var v2 int32 = -2147483648
    var v3 uint32 = 65535
    v5 := GenS1{F0: 42}
    v6 := GenS1{F0: 698}
    var v7 bool = ((!(!((0 >= 80) || ((true || true) || true)))) && (!(42 < 1000)))
    v8 := GenS1{F0: 0}
    for fi9 := 0; fi9 < 10; fi9++ {
        if true {
            if false {
                v3 = 0
            }
            v10 := []int{}
            _ = v10
        }
        var v11 uint32 = 237
        _ = fi9
        _ = v11
    }
    v12 := []int{-2147483648, -504}
    var v13 int = -471
    v14 := GenS2{F0: v13}
    if len(v12) > 0 {
        v12[0] = v13
    }
    v12 = append(v12, v13)
    switch "a" {
    case "a":
        if ((v2 * int32(v3)) != v2) {
            fmt.Println(v2, v1)
        }
    case "\n\t":
        var v15 string = "foo bar"
        _ = v15
    }
    v7 = (469 <= -330)
    var v16 int = 87
    var v17 bool = v7
    v2 -= (v2 & 683)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
