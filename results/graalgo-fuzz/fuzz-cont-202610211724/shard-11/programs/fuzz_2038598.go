package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32; F1 bool; F2 uint64 }
func genFn0(p0 string, p1 uint64, p2 uint8) uint32 {
    var v4 uint = 42
    v5 := []int{}
    var v6 int32 = 42
    var v7 uint8 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 2147483648
}

func main() {
    v1 := make(map[uint32]int)
    var v2 float64 = (1.5 + -2.25)
    var v3 bool = (1 > 42)
    for fi8 := 0; fi8 < 4; fi8++ {
        v9 := []int{fi8, -386}
        v10 := GenS1{F0: 941}
        var v11 uint64 = 1
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    var v12 uint64 = 0
    if ((false && v3) || v3) {
        v1[42] = -705
    }
    switch "hello" {
    case "":
        var v13 int = 32768
        _ = v13
    case "\n\t":
        v14 := GenS2{F0: 9223372036854775807}
        _ = v14
    default:
        v15 := v1[2]
        _ = v15
    }
    v16 := GenS2{F0: 0}
    _ = len(v1)
    var v17 int32 = 0
    for fi18 := 0; fi18 < 13; fi18++ {
        var v19 float64 = (1.5 - (v2 - (v2 - float64((0.1 - 0.0)))))
        _ = fi18
        _ = v19
    }
    switch ("" + ("a" + "a")) {
    case "hello":
        for fi20 := 0; fi20 < 13; fi20++ {
            v12 += v12
            v12 = 2
            _ = fi20
        }
    default:
        v17 -= v17
    }
    var v21 float64 = v2
    var v22 int64 = 0
    switch (("hello" + "hello") + "hello") {
    case "a":
        var v23 uint64 = 658
        _ = v23
    default:
        v24 := GenS2{F0: (v22 >> ((4294967295) & 63))}
        _ = v24
    }
    if (v3 || (2147483648 <= 2147483647)) {
        v25 := GenS3{F0: 1, F1: false, F2: (v12 + 415)}
        fmt.Println(v2)
        if false {
            fmt.Println(v17, v12, v17)
            delete(v1, 623)
        }
        _ = v25
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v3
}
