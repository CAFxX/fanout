package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 int8, p1 string, p2 uint) bool {
    p0 = (-(p0))
    if (!true) {
        v4 := []int{198, -686}
        p1 = p1
        if len(v4) > 0 {
            v4[0] = 32767
        }
        _ = v4
    } else {
        switch 0 {
        case 9223372036854775807:
            fmt.Println(p1)
        case 32768:
            p0 = p0
        default:
            v5 := []int{330, 0, 943, 1}
            _ = v5
        }
        v6 := GenS1{F0: p2}
        p0 = (-(p0))
        _ = v6
    }
    var v7 uint = (^(p2))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return true
}
func genFn1(p0 int32) int32 {
    var v8 int = 0
    for fi9 := 0; fi9 < 14; fi9++ {
        p0 = 478
        fmt.Println(fi9)
        _ = fi9
    }
    v8 = (v8 - (v8 % ((-9223372036854775808) | 1)))
    var v10 int16 = -32768
    v11 := []int{(^(v8)), -302, v8}
    _ = p0
    _ = v10
    _ = v11
    _ = v8
    return 0
}

func main() {
    var v1 uint = 18
    v2 := make(map[uint32]int)
    var v3 uint64 = 101
    var v12 int32 = -2147483648
    v3 += 217
    v13 := []int{}
    var v14 string = "\n\t"
    switch v14 {
    case "a":
        fmt.Println(v1, v3, v1)
    case "":
        v2[0] = 127
    case "hello":
        var v15 uint64 = 684
        _ = v15
    default:
        v14 = v14
    }
    v16 := GenS1{F0: v1}
    v13 = append(v13, 1)
    var v17 uint16 = uint16(v3)
    v18 := GenS1{F0: (v1 << ((v1) & 63))}
    var v19 string = "a"
    var v20 uint = v1
    for fi21 := 0; fi21 < 14; fi21++ {
        var v22 bool = (1 <= int64(v1))
        _ = fi21
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
