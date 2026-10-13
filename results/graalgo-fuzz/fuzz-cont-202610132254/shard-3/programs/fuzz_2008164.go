package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 uint8 }
type GenS2 struct { F0 uint16; F1 int16 }
func genFn0() int {
    var v5 string = "a"
    fmt.Println(v5)
    _ = v5
    return 42
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int64 = -1
    v3 := []uint32{42, 2}
    var v4 bool = false
    if len(v3) > 0 {
        v3[0] = 376
    }
    var v6 uint32 = 1
    var v7 string = "\n\t"
    var v8 uint32 = v6
    v1[v8] = -187
    for fi9 := 0; fi9 < 13; fi9++ {
        v10 := GenS1{F0: 42, F1: 2, F2: 0}
        v6 = (^(v8))
        _ = fi9
        _ = v10
    }
    var v11 uint64 = 18446744073709551615
    var v12 string = ""
    if true {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v6 = 247
    } else {
        v3 = append(v3, v8)
    }
    v13 := GenS2{F0: 42, F1: int16(v11)}
    var v14 uint16 = 32767
    switch 127 {
    case -2147483648:
        var v15 int8 = 92
        _ = v15
    case 128:
        v8 = v6
    default:
        var v16 uint = 42
        _ = v16
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
}
