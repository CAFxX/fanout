package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 uint64 }
func genFn0(p0 uint8, p1 int8, p2 int8) int8 {
    var v4 uint64 = 1
    switch ("a" + "hello") {
    case "\n\t":
        v5 := []int{42, 146, 42}
        _ = v5
    default:
        v4 += v4
    }
    for fi6 := 0; fi6 < 18; fi6++ {
        var v7 uint16 = 255
        p0 += p0
        _ = fi6
        _ = v7
    }
    fmt.Println(v4)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return (p2 % ((p1) | 1))
}

func main() {
    var v1 int64 = 0
    var v2 int = -692
    v3 := make(map[string]bool)
    for fi8 := 0; fi8 < 6; fi8++ {
        v9 := v3[""]
        for fi10 := 0; fi10 < 10; fi10++ {
            v1 = (v1 << ((42) & 63))
            v11 := []int{32767}
            _ = fi10
            _ = v11
        }
        _ = fi8
        _ = v9
    }
    v12 := GenS1{F0: 4294967295, F1: ("\n\t" + ("hello" + "foo bar"))}
    for fi13 := 0; fi13 < 8; fi13++ {
        switch v2 {
        case 0:
            for fi14 := 0; fi14 < 1; fi14++ {
                v15 := GenS2{F0: 0}
                fmt.Println(v2, v2, fi14)
                _ = fi14
                _ = v15
            }
        default:
            v16 := []int{(v2 | v2), (fi13 & 127), (^(fi13))}
            _ = v16
        }
        _ = fi13
    }
    switch v2 {
    case 4294967295:
        delete(v3, "hello")
    default:
        for fi17 := 0; fi17 < 9; fi17++ {
            var v18 int32 = 2147483647
            _ = fi17
            _ = v18
        }
    }
    fmt.Println(v2, v2)
    v19 := []int{v2, (v2 ^ 42), (-(v2))}
    v2 = 278
    v20 := GenS2{F0: 42}
    var v21 uint64 = 383
    var v22 int16 = -32768
    fmt.Println(v22)
    v19 = append(v19, v2)
    fmt.Println(v22, v21)
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
}
