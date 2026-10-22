package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64 }
func genFn0(p0 uint8, p1 int) uint16 {
    var v6 float32 = 0.0
    v7 := []int{}
    if true {
        switch "a" {
        case "\n\t":
            p0 += p0
        }
        v8 := GenS1{F0: -1, F1: 0}
        fmt.Println(p0, p0, p0)
        _ = v8
    } else {
        fmt.Println(v6, v6)
        v9 := GenS1{F0: -144, F1: 0}
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 255
}
func genFn1(p0 int, p1 string) bool {
    for fi10 := 0; fi10 < 2; fi10++ {
        var v11 uint = 65535
        _ = fi10
        _ = v11
    }
    var v12 uint16 = 42
    v13 := GenS1{F0: 0, F1: 1}
    var v14 int = -1
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    _ = v14
    return (!((true || (!false)) && (1000 >= 42)))
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []uint32{1}
    var v3 bool = true
    v4 := make(map[int]string)
    var v5 int64 = 1
    v5 += -9223372036854775808
    var v15 uint8 = 0
    v16 := GenS2{F0: 333}
    for fi17 := 0; fi17 < 4; fi17++ {
        v15 -= v15
        _ = fi17
    }
    var v18 uint8 = v15
    v19 := v1[4294967295]
    v15 -= 128
    v1[-9223372036854775808] = v19
    var v20 string = ("hello" + ("\n\t" + "hello"))
    v21 := GenS3{F0: 0}
    v22 := GenS2{F0: int(v5)}
    for fi23 := 0; fi23 < 6; fi23++ {
        v3 = ((^(fi23)) == 42)
        _ = fi23
    }
    v15 -= (v18 % ((v18) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
