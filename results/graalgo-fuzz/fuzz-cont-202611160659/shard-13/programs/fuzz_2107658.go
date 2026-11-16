package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 bool) int32 {
    for fi4 := 0; fi4 < 9; fi4++ {
        v5 := []int{32768, 0}
        _ = fi4
        _ = v5
    }
    p0 = (!p0)
    p0 = (!p0)
    _ = p0
    return -2147483648
}

func main() {
    var v1 int8 = -1
    v2 := []int{722, 9223372036854775807}
    v3 := []uint32{2, 2147483647, 4294967295, 0}
    var v6 int16 = 1
    var v7 string = "\n\t"
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v1 = -1
    v8 := GenS1{F0: 32767}
    v7 = "foo bar"
    var v9 uint32 = 2
    fmt.Println(v7)
    fmt.Println(v7, v6, v9)
    fmt.Println(v9, v9)
    var v10 uint = 1
    var v11 uint64 = 0
    if (v7 < (v7 + "hello")) {
        switch -780 {
        case 32768:
            v1 = v1
        case -130:
            v1 -= (-(v1))
        case 4294967295:
            v7 = v7
        default:
            var v12 bool = (false && (v7 >= v7))
            _ = v12
        }
        fmt.Println(v9, v6)
        v11 -= v11
    }
    for fi13 := 0; fi13 < 11; fi13++ {
        var v14 int16 = (v6 >> (((v10 ^ 344)) & 63))
        if ((^(v11)) == (v11 - 0)) {
            v14 = v6
            _ = len(v2)
        }
        _ = fi13
        _ = v14
    }
    if true {
        v15 := GenS1{F0: 0}
        _ = v15
    }
    if false {
        fmt.Println(v9, v10)
    } else {
        var v16 bool = true
        _ = v16
    }
    fmt.Println(v10, v6)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
