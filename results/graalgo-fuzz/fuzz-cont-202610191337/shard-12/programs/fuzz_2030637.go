package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 int16 }
func genFn0(p0 uint8, p1 string, p2 uint8) uint8 {
    v4 := []int{-9223372036854775808, 42}
    var v5 string = ("" + "\n\t")
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return p2
}
func genFn1(p0 bool, p1 uint, p2 int64) bool {
    var v6 bool = false
    var v7 uint64 = 793
    if true {
        p1 -= 4294967295
    }
    p2 += 42
    p1 -= p1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return (!v6)
}

func main() {
    var v1 uint = 0
    var v2 int32 = -1
    v3 := []uint32{2147483648}
    for fi8 := 0; fi8 < 3; fi8++ {
        v2 += (v2 ^ (v2 - v2))
        v2 -= 42
        _ = fi8
    }
    if len(v3) > 0 {
        v3[0] = 0
    }
    v9 := GenS1{F0: (v1 % ((v1) | 1)), F1: 42, F2: 0}
    var v10 uint8 = 0
    switch "foo bar" {
    case "":
        for fi11 := 0; fi11 < 7; fi11++ {
            var v12 int32 = int32(fi11)
            _ = fi11
            _ = v12
        }
    case "foo bar":
        var v13 bool = (!false)
        _ = v13
    }
    v2 = (v2 & -1)
    v3 = append(v3, 1)
    switch -498 {
    case 9223372036854775807:
        v10 = (v10 << ((2) & 63))
    case 128:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v9
}
