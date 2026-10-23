package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int }
func genFn0() uint32 {
    var v4 float32 = 0.0
    var v5 int16 = 205
    for fi6 := 0; fi6 < 8; fi6++ {
        fmt.Println(v5)
        _ = fi6
    }
    v7 := []int{}
    _ = v4
    _ = v5
    _ = v7
    return 2
}
func genFn1(p0 uint16, p1 int8) uint16 {
    fmt.Println(p1, p0)
    var v8 uint8 = 252
    if true {
        var v9 int = 9223372036854775807
        var v10 int32 = 42
        switch ("foo bar" + "hello") {
        case "hello":
            v11 := []int{v9, (v9 * (v9 & v9))}
            _ = v11
        case "\n\t":
            switch 32767 {
            case 647:
                v10 -= -1
            case -236:
                v12 := GenS1{F0: v9}
                _ = v12
            default:
                fmt.Println(v10)
            }
        case "":
            v13 := []int{-2147483648}
            _ = v13
        }
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v8
    return (p0 << ((42) & 63))
}

func main() {
    var v1 float32 = 1.5
    v2 := []int{}
    v3 := make(map[int]int)
    switch "" {
    case "hello":
        for fi14 := 0; fi14 < 17; fi14++ {
            v1 += v1
            _ = fi14
        }
    }
    v1 = v1
    var v15 string = "foo bar"
    var v16 int32 = 32
    fmt.Println(v15, v1, v16)
    if false {
        v17 := GenS2{F0: 9223372036854775807}
        for fi18 := 0; fi18 < 6; fi18++ {
            v16 = v16
            _ = fi18
        }
        v16 = v16
        _ = v17
    }
    var v19 uint8 = 255
    var v20 uint64 = 826
    var v21 uint32 = genFn0()
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
