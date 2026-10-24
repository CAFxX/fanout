package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 uint16 }
type GenS2 struct { F0 int32 }
func genFn0() uint {
    var v5 uint16 = 2
    var v6 int32 = -1
    v7 := GenS1{F0: 1, F1: 127, F2: (v5 + v5)}
    var v8 string = "foo bar"
    v9 := []int{0, -259}
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 0
}
func genFn1(p0 int64, p1 int16) int32 {
    v10 := []int{}
    var v11 string = "\n\t"
    for fi12 := 0; fi12 < 5; fi12++ {
        for fi13 := 0; fi13 < 20; fi13++ {
            v14 := GenS1{F0: 18446744073709551615, F1: 42, F2: 356}
            var v15 float64 = (-2.25 - 1e10)
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = len(v10)
        if len(v10) > 0 {
            v10[0] = fi12
        }
        _ = fi12
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return 0
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]bool)
    v3 := make(map[int]int)
    v4 := []int{32767, 0}
    var v16 uint16 = 32767
    v17 := GenS1{F0: 1, F1: 1, F2: 65535}
    v18 := GenS2{F0: -2147483648}
    var v19 int32 = 571
    var v20 float64 = 1e10
    if true {
        v16 += 42
    }
    delete(v2, -875)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v21 int32 = genFn1(-1, -789)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
