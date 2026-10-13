package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string }

func main() {
    var v1 string = "foo bar"
    var v2 string = "\n\t"
    var v3 uint16 = 887
    var v4 uint = 42
    v5 := make(map[uint32]bool)
    var v6 bool = true
    fmt.Println(v1)
    _ = len(v5)
    fmt.Println(v6, v6, v3)
    v2 = v2
    var v7 uint32 = 42
    fmt.Println(v4, v2, v2)
    fmt.Println(v2)
    if true {
        if (100 <= 0) {
            v8 := GenS1{F0: -2147483648}
            var v9 int32 = -2147483648
            v6 = v6
            _ = v8
            _ = v9
        } else {
            var v10 float64 = 0.1
            _ = v10
        }
    } else {
        v3 = 42
    }
    var v11 uint8 = 2
    var v12 int = 42
    v3 -= 352
    if v6 {
        v3 = 91
    } else {
        v13 := GenS2{F0: "a"}
        switch -122 {
        case 315:
            v14 := GenS1{F0: -1}
            _ = v14
        default:
            v6 = (false || false)
        }
        var v15 float32 = (-2.25 - 1.5)
        _ = v13
        _ = v15
    }
    v16 := []int{(v12 - int(v7)), (v12 << ((0) & 63)), (v12 << (((v4 % (((^(v4))) | 1))) & 63))}
    v7 += v7
    v17 := GenS1{F0: -139}
    fmt.Println(v3, v6, v6)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
