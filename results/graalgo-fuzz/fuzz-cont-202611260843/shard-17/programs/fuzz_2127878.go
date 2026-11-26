package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string }
func genFn0(p0 int, p1 int) bool {
    v5 := []int{p1, p1}
    var v6 int = (p0 | p0)
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return false
}

func main() {
    var v1 uint16 = 32767
    var v2 int32 = 662
    var v3 bool = ((!(!true)) && false)
    v4 := make(map[uint32]uint32)
    v7 := []int{int(v2), -1}
    v8 := GenS1{F0: 18446744073709551615}
    v4[2147483648] = 1
    switch "" {
    case "":
        v9 := GenS1{F0: 4294967295}
        _ = v9
    case "hello":
        var v10 uint = 333
        _ = v10
    case "a":
        v3 = (v3 && v3)
    default:
        v11 := GenS2{F0: "hello"}
        _ = v11
    }
    var v12 uint16 = v1
    v3 = (-86 != -128)
    v3 = (128 == 255)
    v1 = v1
    var v13 uint32 = 614
    var v14 uint = 2
    v15 := v4[(v13 + (^(v13)))]
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
