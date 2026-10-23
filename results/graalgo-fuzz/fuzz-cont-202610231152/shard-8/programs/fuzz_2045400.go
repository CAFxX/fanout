package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 int32 }
func genFn0(p0 string, p1 uint64) bool {
    v7 := []int{}
    v7 = append(v7, -1)
    _ = p0
    _ = p1
    _ = v7
    return false
}

func main() {
    var v1 uint8 = 189
    var v2 uint16 = 46
    var v3 int8 = -108
    v4 := make(map[string]bool)
    var v5 uint16 = (v2 >> ((711) & 63))
    var v6 float64 = 1e10
    var v8 int64 = -1
    for fi9 := 0; fi9 < 2; fi9++ {
        v10 := []int{fi9, (fi9 >> ((2) & 63))}
        _ = fi9
        _ = v10
    }
    var v11 bool = (1000 >= 0)
    fmt.Println(v1)
    v12 := v4["hello"]
    switch 1 {
    case 886:
        v13 := v4[""]
        _ = v13
    case 321:
        var v14 int8 = (v3 % ((v3) | 1))
        _ = v14
    }
    if ((v1 ^ 255) <= v1) {
        var v15 int = 0
        _ = len(v4)
        _ = v15
    }
    v16 := []int{21, 723, 127, -9223372036854775808}
    for fi17 := 0; fi17 < 8; fi17++ {
        var v18 bool = (false || (v3 != -128))
        _ = fi17
        _ = v18
    }
    v19 := GenS1{F0: (v3 >> ((825) & 63)), F1: 217, F2: 2147483647}
    var v20 int32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
