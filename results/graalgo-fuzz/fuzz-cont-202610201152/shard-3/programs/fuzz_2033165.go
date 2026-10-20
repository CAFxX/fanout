package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 int8 }
func genFn0(p0 string) uint {
    fmt.Println(p0)
    fmt.Println(p0)
    for fi6 := 0; fi6 < 15; fi6++ {
        if (true || (!false)) {
            fmt.Println(fi6, p0)
        }
        _ = fi6
    }
    _ = p0
    return 437
}

func main() {
    var v1 uint32 = 982
    v2 := []string{"\n\t"}
    var v3 uint = 255
    v4 := make(map[int]int)
    var v5 uint8 = 128
    var v7 uint8 = 255
    var v8 bool = (-1 == 242)
    v7 += v7
    if (false && ((((902 == -779) && (1 != 555)) && v8) && (!v8))) {
        var v9 float64 = float64(-2.25)
        _ = v9
    }
    v3 += (v3 % ((573) | 1))
    v10 := GenS1{F0: 92, F1: 1, F2: -12}
    v8 = ((32767 != 255) || true)
    var v11 uint8 = v7
    var v12 bool = v8
    for fi13 := 0; fi13 < 11; fi13++ {
        var v14 float64 = 3.14159
        var v15 int = (fi13 << ((803) & 63))
        _ = fi13
        _ = v14
        _ = v15
    }
    if len(v2) > 0 {
        v2[0] = "hello"
    }
    var v16 int32 = -2147483648
    var v17 uint32 = 569
    var v18 uint8 = (^(v7))
    var v19 int32 = -514
    var v20 uint64 = 42
    var v21 uint = (v3 << (((v3 & (v3 >> ((v3) & 63)))) & 63))
    var v22 uint32 = 821
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
