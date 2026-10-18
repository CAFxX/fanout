package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 int16 }
type GenS3 struct { F0 int; F1 uint16; F2 uint32 }
func genFn0() bool {
    var v6 float64 = 0.0
    v7 := GenS1{F0: (true || (((!true) || (false && false)) && (2 < 327)))}
    v6 = v6
    fmt.Println(v6)
    _ = v6
    _ = v7
    return (false && false)
}
func genFn1(p0 int64, p1 bool, p2 uint8) int64 {
    fmt.Println(p0, p0, p1)
    var v8 float64 = -2.25
    for fi9 := 0; fi9 < 1; fi9++ {
        var v10 uint64 = 756
        v11 := GenS1{F0: (-9223372036854775808 <= (p0 - (-(p0))))}
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 int = 4294967295
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v8
    return p0
}

func main() {
    v1 := []int{127, -994}
    var v2 string = (("\n\t" + "\n\t") + "\n\t")
    v3 := make(map[int]int)
    v4 := []uint32{}
    var v5 uint16 = 65535
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v13 string = "a"
    var v14 float32 = (0.0 - 0.1)
    delete(v3, 1)
    _ = len(v4)
    var v15 string = "a"
    v16 := GenS1{F0: (522 == -2147483648)}
    if false {
        v5 = v5
        _ = len(v1)
        if (!false) {
            _ = len(v3)
            v17 := GenS2{F0: 0, F1: 830}
            v18 := GenS3{F0: 630, F1: (v5 - 0), F2: 318}
            _ = v17
            _ = v18
        }
    }
    fmt.Println(v13, v14)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
