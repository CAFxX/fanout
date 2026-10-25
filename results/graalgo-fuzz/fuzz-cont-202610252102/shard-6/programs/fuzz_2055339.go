package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 bool; F1 string }
type GenS4 struct { F0 uint32; F1 uint }
func genFn0(p0 int, p1 int, p2 int16) int8 {
    if ((p1 > (p1 << ((18446744073709551615) & 63))) && false) {
        var v4 uint = 42
        _ = v4
    } else {
        v5 := []int{32768}
        v6 := GenS1{F0: 263, F1: 106}
        _ = v5
        _ = v6
    }
    for fi7 := 0; fi7 < 16; fi7++ {
        switch (-(p0)) {
        case 0:
            v8 := GenS1{F0: 9223372036854775807, F1: 127}
            _ = v8
        case 32768:
            v9 := GenS1{F0: 42, F1: 134}
            _ = v9
        case 127:
            v10 := []int{}
            _ = v10
        default:
            var v11 uint64 = 608
            _ = v11
        }
        var v12 int = (p0 | (p0 << ((645) & 63)))
        _ = fi7
        _ = v12
    }
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    var v1 uint32 = 2
    v2 := make(map[int]bool)
    var v3 int32 = 1
    v3 -= (v3 + v3)
    var v13 uint8 = 255
    for fi14 := 0; fi14 < 6; fi14++ {
        var v15 bool = (!(!(!((v13 >= v13) && false))))
        _ = fi14
        _ = v15
    }
    var v16 int16 = -556
    v17 := GenS1{F0: 1, F1: v13}
    v18 := []int{4294967295, -9223372036854775808}
    for fi19 := 0; fi19 < 9; fi19++ {
        var v20 int16 = (v16 - (-(v16)))
        _ = fi19
        _ = v20
    }
    var v21 uint8 = v13
    v13 = v21
    v13 = 42
    v22 := GenS2{F0: (v13 + 255), F1: v16}
    v3 = (v3 >> ((255) & 63))
    v23 := GenS3{F0: true, F1: "a"}
    v21 -= 2
    for fi24 := 0; fi24 < 13; fi24++ {
        v25 := GenS4{F0: v1, F1: 566}
        switch "hello" {
        case "foo bar":
            fmt.Println(v3)
        default:
            fmt.Println(fi24)
        }
        _ = fi24
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
}
