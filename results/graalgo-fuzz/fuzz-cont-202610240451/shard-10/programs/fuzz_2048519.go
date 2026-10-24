package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 uint64; F1 bool; F2 int8 }
func genFn0(p0 uint8, p1 int32, p2 int8) int16 {
    v7 := []int{2147483647, 548}
    for fi8 := 0; fi8 < 7; fi8++ {
        v9 := GenS1{F0: p2, F1: p1}
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 10; fi10++ {
        fmt.Println(fi10, p1)
        p0 += p0
        _ = fi10
    }
    if len(v7) > 0 {
        v7[0] = 127
    }
    for fi11 := 0; fi11 < 17; fi11++ {
        var v12 bool = true
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 1
}

func main() {
    v1 := []uint32{}
    v2 := make(map[string]string)
    v3 := make(map[string]bool)
    v4 := make(map[uint32]bool)
    v5 := []string{"\n\t", "\n\t", "\n\t", ""}
    v6 := []string{"", "\n\t", "\n\t"}
    v13 := GenS1{F0: 1, F1: 1}
    var v14 bool = (false || false)
    v15 := v2[""]
    v16 := GenS2{F0: 1, F1: (!v14), F2: 42}
    var v17 uint64 = 2
    v1 = append(v1, 111)
    for fi18 := 0; fi18 < 10; fi18++ {
        var v19 uint16 = 813
        var v20 uint = 1
        _ = fi18
        _ = v19
        _ = v20
    }
    if len(v5) > 0 {
        v5[0] = "hello"
    }
    v15 = (v15 + v15)
    for fi21 := 0; fi21 < 3; fi21++ {
        if (3.14159 > 3.14159) {
            for fi22 := 0; fi22 < 5; fi22++ {
                fmt.Println(v17, fi21)
                var v23 bool = v14
                _ = fi22
                _ = v23
            }
        } else {
            v15 = v15
        }
        if len(v5) > 0 {
            v5[0] = v15
        }
        var v24 int8 = 75
        _ = fi21
        _ = v24
    }
    var v25 float32 = 3.14159
    v15 = v15
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
