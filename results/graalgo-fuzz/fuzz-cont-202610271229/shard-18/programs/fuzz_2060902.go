package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int }
func genFn0(p0 uint16, p1 string) int {
    var v7 uint8 = 1
    var v8 int8 = 127
    v9 := GenS1{F0: 127}
    p1 = "foo bar"
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return -1
}
func genFn1(p0 uint8, p1 int) string {
    v10 := GenS2{F0: (p1 + p1)}
    fmt.Println(p1, p1)
    switch genFn0(2, "hello") {
    case -659:
        v11 := GenS1{F0: p1}
        _ = v11
    default:
        p0 = p0
    }
    if true {
        var v12 bool = false
        v13 := GenS2{F0: p1}
        v12 = false
        _ = v12
        _ = v13
    }
    switch "" {
    case "hello":
        p0 = 201
    case "foo bar":
        v14 := []int{p1}
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v10
    return "\n\t"
}

func main() {
    v1 := []uint32{2147483648, 42, 65535, 615}
    v2 := make(map[string]uint32)
    v3 := []int{-9223372036854775808, 9223372036854775807}
    v4 := make(map[uint32]uint32)
    v5 := make(map[int]int)
    v6 := make(map[uint32]uint32)
    var v15 bool = (679 == 365)
    v15 = false
    var v16 uint8 = 82
    v17 := v5[42]
    v16 -= 2
    v18 := GenS2{F0: genFn0(0, ("foo bar" + ("" + "")))}
    for fi19 := 0; fi19 < 17; fi19++ {
        v17 += fi19
        v17 += (v17 << ((342) & 63))
        switch fi19 {
        case 913:
            v20 := GenS2{F0: (v17 * v17)}
            _ = v20
        }
        _ = fi19
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if (-128 != 100) {
        var v21 uint8 = 127
        v15 = v15
        _ = v21
    }
    for fi22 := 0; fi22 < 9; fi22++ {
        v1 = append(v1, 4294967295)
        var v23 int = v17
        _ = fi22
        _ = v23
    }
    for fi24 := 0; fi24 < 3; fi24++ {
        var v25 uint = 0
        _ = fi24
        _ = v25
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
