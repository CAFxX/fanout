package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 bool; F1 uint64; F2 uint16 }
func genFn0() string {
    for fi6 := 0; fi6 < 4; fi6++ {
        var v7 int16 = 1000
        v8 := []int{fi6, fi6, fi6}
        v9 := GenS1{F0: 2147483647, F1: 42}
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    var v10 float64 = 0.0
    v11 := []int{}
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    _ = v10
    _ = v11
    return "\n\t"
}
func genFn1(p0 uint16) uint8 {
    var v12 uint32 = 2
    v13 := GenS2{F0: (uint64(p0) != 59), F1: 2, F2: 2}
    _ = p0
    _ = v12
    _ = v13
    return 0
}

func main() {
    v1 := []int{-474}
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]uint32)
    var v4 string = ((("hello" + "foo bar") + ((("a" + "\n\t") + "foo bar") + "\n\t")) + "hello")
    v5 := []uint32{}
    var v14 uint16 = 255
    v4 = (v4 + ("hello" + v4))
    var v15 uint32 = 0
    for fi16 := 0; fi16 < 15; fi16++ {
        for fi17 := 0; fi17 < 14; fi17++ {
            if true {
                _ = len(v5)
            }
            v15 += (^(v15))
            _ = fi17
        }
        var v18 int16 = 32767
        v4 = v4
        _ = fi16
        _ = v18
    }
    v4 = "a"
    fmt.Println(v4)
    var v19 int32 = -2147483648
    v14 = uint16(v15)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
