package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 uint8 }
func genFn0() int {
    var v7 uint32 = 102
    v8 := []int{-2147483648}
    v7 = (v7 | v7)
    if false {
        v7 += (v7 * (v7 / ((1) | 1)))
    } else {
        var v9 float32 = 1.5
        v9 -= 1.5
        v8 = append(v8, 1)
        _ = v9
    }
    v10 := GenS1{F0: 0, F1: false, F2: 2}
    _ = v10
    _ = v7
    _ = v8
    return 127
}
func genFn1(p0 int16) int32 {
    var v11 bool = (!(int32(p0) > 652))
    for fi12 := 0; fi12 < 14; fi12++ {
        var v13 uint16 = 255
        v11 = (!((false || v11) && ((false && v11) || (-1 != 42))))
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = v11
    return -104
}

func main() {
    v1 := []uint32{}
    v2 := []string{(("foo bar" + "\n\t") + "\n\t"), "", "a"}
    var v3 int32 = 1
    var v4 int32 = (v3 & v3)
    v5 := []string{"foo bar", ""}
    v6 := make(map[uint32]string)
    v14 := GenS1{F0: int(v3), F1: true, F2: 255}
    var v15 uint16 = 2
    var v16 int = genFn0()
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v4 = v4
    var v17 uint8 = 12
    var v18 bool = (true && (false && (!true)))
    var v19 uint64 = 25
    switch (v16 * (v16 | (v16 + (v16 & int(v17))))) {
    case 128:
        v20 := v6[828]
        _ = v20
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
