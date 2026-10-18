package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 uint64 }
func genFn0() int32 {
    var v7 float32 = 1e10
    v8 := []int{127}
    v7 -= v7
    var v9 int16 = 62
    _ = v7
    _ = v8
    _ = v9
    return 0
}
func genFn1(p0 uint16, p1 int64, p2 string) int8 {
    p0 = 0
    for fi10 := 0; fi10 < 3; fi10++ {
        var v11 string = (((p2 + p2) + "a") + "a")
        if true {
            p0 = p0
        } else {
            p1 += p1
            switch (-(fi10)) {
            case -2147483648:
                p2 = "a"
            }
            var v12 uint16 = (^(p0))
            _ = v12
        }
        if false {
            for fi13 := 0; fi13 < 10; fi13++ {
                v14 := []int{198, fi13, fi10, 32768}
                v11 = p2
                fmt.Println(p0, p0)
                _ = fi13
                _ = v14
            }
            fmt.Println(p0)
        } else {
            v15 := GenS1{F0: (!true), F1: 42}
            _ = v15
        }
        _ = fi10
        _ = v11
    }
    var v16 uint8 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    return int8(v16)
}

func main() {
    var v1 uint8 = 255
    var v2 int32 = 0
    v3 := make(map[uint32]string)
    var v4 uint64 = 42
    var v5 uint8 = (v1 & (^(v1)))
    var v6 int32 = (v2 * int32(v4))
    var v17 float32 = -2.25
    v18 := v3[42]
    fmt.Println(v17)
    v1 = 1
    if false {
        v19 := v3[65535]
        var v20 uint16 = 0
        _ = v19
        _ = v20
    } else {
        var v21 int64 = -1
        _ = v21
    }
    for fi22 := 0; fi22 < 13; fi22++ {
        v18 = ""
        v23 := []int{fi22, 127, fi22, 2147483647}
        var v24 int = fi22
        _ = fi22
        _ = v23
        _ = v24
    }
    v25 := GenS2{F0: v4}
    fmt.Println(v6, v5)
    var v26 string = (v18 + "\n\t")
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
