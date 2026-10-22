package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 uint32 }

func main() {
    var v1 int64 = 181
    var v2 string = ""
    var v3 uint32 = 1
    v2 = "\n\t"
    v1 = v1
    var v4 bool = (0 >= -1)
    v5 := []int{4294967295}
    if len(v5) > 0 {
        v5[0] = 32767
    }
    v6 := GenS1{F0: (v1 / (((v1 % (((v1 - v1)) | 1))) | 1)), F1: 0, F2: v3}
    v4 = (54 > 18446744073709551615)
    v4 = false
    var v7 string = v2
    v8 := GenS1{F0: (v1 * v1), F1: -2147483648, F2: 764}
    for fi9 := 0; fi9 < 10; fi9++ {
        var v10 int64 = 1
        switch v2 {
        case "a":
            v1 -= v10
        default:
            v4 = ((-2147483648 >= -267) && (42 < 580))
        }
        var v11 uint = 18446744073709551615
        _ = fi9
        _ = v10
        _ = v11
    }
    v7 = v7
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
