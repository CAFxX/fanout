package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool }
func genFn0(p0 uint64, p1 int64, p2 int) int {
    if true {
        v4 := []int{}
        for fi5 := 0; fi5 < 7; fi5++ {
            if false {
                var v6 int8 = -21
                _ = v6
            }
            _ = fi5
        }
        _ = v4
    }
    v7 := []int{-1, p2}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (p2 / ((p2) | 1))
}
func genFn1(p0 int32, p1 bool, p2 int) uint8 {
    p1 = p1
    v8 := []int{p2, (p2 >> ((105) & 63))}
    if p1 {
        if len(v8) > 0 {
            v8[0] = -2147483648
        }
    } else {
        v9 := GenS1{F0: 42}
        var v10 int32 = (p0 & 959)
        if (1 > 0) {
            fmt.Println(p1, p2, p0)
            var v11 int16 = 1
            fmt.Println(v11)
            _ = v11
        }
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 1
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []uint32{}
    var v3 int8 = 1
    v3 = (v3 >> ((568) & 63))
    _ = len(v1)
    v12 := GenS2{F0: ((v3 + 56) > v3)}
    fmt.Println(v3)
    var v13 int = -608
    _ = len(v1)
    var v14 int64 = -424
    var v15 int8 = -1
    v16 := GenS2{F0: true}
    _ = len(v2)
    switch -9223372036854775808 {
    case 128:
        fmt.Println(v13)
    case 501:
        var v17 string = ("\n\t" + "hello")
        _ = v17
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
}
