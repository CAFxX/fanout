package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int; F1 int32 }
func genFn0(p0 int8, p1 int, p2 uint32) int {
    p0 -= (p0 << ((42) & 63))
    p1 = (^(p1))
    if ((p1 < p1) && false) {
        var v7 uint32 = p2
        v8 := GenS1{F0: 1}
        var v9 uint16 = 1
        _ = v7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    return (p1 + 127)
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int16 = -895
    v3 := make(map[uint32]string)
    v4 := make(map[int]int)
    var v5 string = ("\n\t" + ("a" + "foo bar"))
    var v6 uint64 = 2
    fmt.Println(v6, v2)
    v10 := GenS2{F0: 281, F1: 955}
    if false {
        var v11 uint8 = 180
        v12 := GenS3{F0: true}
        var v13 uint32 = 42
        _ = v11
        _ = v12
        _ = v13
    }
    v14 := GenS2{F0: -339, F1: 42}
    v15 := GenS1{F0: 42}
    fmt.Println(v5, v6)
    var v16 float32 = float32(1e10)
    var v17 float64 = 3.14159
    fmt.Println(v16)
    v18 := GenS4{F0: genFn0(42, 127, 857), F1: 588}
    for fi19 := 0; fi19 < 9; fi19++ {
        var v20 string = v5
        v21 := GenS4{F0: (fi19 & (fi19 % ((fi19) | 1))), F1: 1}
        var v22 int = fi19
        _ = fi19
        _ = v20
        _ = v21
        _ = v22
    }
    v23 := GenS2{F0: -1, F1: genFn0(-128, 4294967295, 2147483648)}
    for fi24 := 0; fi24 < 14; fi24++ {
        var v25 uint16 = 2
        _ = fi24
        _ = v25
    }
    var v26 string = v5
    var v27 uint16 = 375
    var v28 int8 = 100
    v29 := []int{0, genFn0((v28 * (v28 & v28)), genFn0((-(v28)), 2147483647, 4294967295), 2), -61}
    var v30 bool = true
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
