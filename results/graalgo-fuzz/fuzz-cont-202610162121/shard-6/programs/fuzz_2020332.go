package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
func genFn0(p0 uint8, p1 string) int {
    v4 := GenS1{F0: p0, F1: 281}
    v5 := GenS1{F0: 0, F1: 0}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return -2147483648
}

func main() {
    var v1 float32 = 1.5
    v2 := make(map[uint32]string)
    v3 := make(map[int]bool)
    v6 := []int{484, 32767, 74}
    var v7 bool = true
    v8 := GenS1{F0: 255, F1: 1}
    v7 = (!(-720 >= -1))
    v9 := v2[28]
    switch "a" {
    case "\n\t":
        if (128 != 1) {
            var v10 int64 = 738
            _ = v10
        } else {
            v9 = v9
        }
    case "a":
        v11 := GenS1{F0: 149, F1: 273}
        _ = v11
    case "":
        fmt.Println(v7)
    default:
        v7 = v7
    }
    var v12 int8 = -19
    v13 := GenS1{F0: 2, F1: 0}
    fmt.Println(v9, v7)
    delete(v3, 32768)
    var v14 string = (("" + (v9 + ("\n\t" + (v9 + "a")))) + v9)
    v9 = ("\n\t" + "a")
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v15 uint8 = 111
    if len(v6) > 0 {
        v6[0] = 0
    }
    if (184 > 255) {
        var v16 int16 = -1
        _ = v16
    }
    v1 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
