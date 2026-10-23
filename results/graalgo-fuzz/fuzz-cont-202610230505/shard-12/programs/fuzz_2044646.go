package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int; F1 uint32; F2 int8 }
func genFn0(p0 bool, p1 uint, p2 bool) uint64 {
    var v5 int32 = 2147483647
    v6 := []int{9223372036854775807, 281}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 42
}
func genFn1(p0 bool, p1 uint8, p2 int32) int {
    for fi7 := 0; fi7 < 3; fi7++ {
        var v8 uint = 723
        if (171 < 32767) {
            p2 -= 2147483647
            fmt.Println(p0, p0)
            p0 = p0
        } else {
            if p0 {
                fmt.Println(fi7, v8)
                for fi9 := 0; fi9 < 2; fi9++ {
                    p0 = false
                    var v10 uint16 = 65535
                    var v11 uint = (^(v8))
                    _ = fi9
                    _ = v10
                    _ = v11
                }
            }
            v12 := []int{}
            _ = v12
        }
        p0 = ((1 == 2) || p0)
        _ = fi7
        _ = v8
    }
    var v13 uint16 = 458
    if p0 {
        var v14 float64 = 0.0
        v15 := []int{643, 4294967295}
        switch ("" + "\n\t") {
        case "hello":
            v16 := GenS1{F0: p1}
            _ = v16
        default:
            fmt.Println(p0, p1)
        }
        _ = v14
        _ = v15
    }
    if p0 {
        var v17 string = ("foo bar" + "")
        var v18 uint = 0
        _ = v17
        _ = v18
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 0
}

func main() {
    var v1 float32 = float32(float64(1.5))
    v2 := make(map[uint32]string)
    v3 := []uint32{}
    var v4 uint8 = 0
    var v19 int = genFn1((2 <= 4294967295), v4, -1)
    _ = len(v3)
    fmt.Println(v19, v4)
    var v20 int64 = -1
    if true {
        for fi21 := 0; fi21 < 5; fi21++ {
            if false {
                v20 = v20
            }
            v2[2147483647] = "\n\t"
            _ = fi21
        }
    } else {
        v22 := GenS1{F0: (^(v4))}
        _ = v22
    }
    v23 := GenS1{F0: v4}
    var v24 uint8 = v4
    v25 := GenS2{F0: v19, F1: 1, F2: 0}
    var v26 uint32 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
