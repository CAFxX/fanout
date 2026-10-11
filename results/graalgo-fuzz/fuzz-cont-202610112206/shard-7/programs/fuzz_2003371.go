package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string; F1 uint8; F2 int64 }
func genFn0(p0 uint64, p1 int) bool {
    var v4 bool = (!false)
    if ((100 < -1) || (v4 || true)) {
        p1 = (^(p1))
        for fi5 := 0; fi5 < 2; fi5++ {
            var v6 string = ""
            _ = fi5
            _ = v6
        }
        p1 = (p1 - p1)
    }
    _ = p0
    _ = p1
    _ = v4
    return (265 > 455)
}
func genFn1() bool {
    var v7 uint64 = 42
    v8 := GenS1{F0: 908}
    var v9 bool = true
    fmt.Println(v7, v9)
    switch 127 {
    case 4294967295:
        v7 += v7
    case -225:
        for fi10 := 0; fi10 < 12; fi10++ {
            v11 := GenS1{F0: 1}
            _ = fi10
            _ = v11
        }
    case 692:
        var v12 uint64 = 2
        _ = v12
    }
    _ = v7
    _ = v8
    _ = v9
    return v9
}

func main() {
    var v1 int = 32767
    var v2 uint = 4294967295
    var v3 int = v1
    fmt.Println(v1, v2, v3)
    fmt.Println(v1, v1, v3)
    v2 = (v2 % ((v2) | 1))
    for fi13 := 0; fi13 < 9; fi13++ {
        if false {
            v14 := []int{(v1 & (-(v3))), (fi13 | (fi13 % ((128) | 1))), fi13, v3}
            _ = v14
        }
        _ = fi13
    }
    v15 := GenS2{F0: ("a" + (("\n\t" + "") + "hello")), F1: 42, F2: 1}
    var v16 int8 = 42
    v17 := GenS1{F0: 1}
    var v18 uint = (v2 * 58)
    var v19 int16 = 32767
    v20 := []int{-2147483648, 127, v1, v3}
    for fi21 := 0; fi21 < 1; fi21++ {
        v22 := GenS1{F0: 531}
        _ = fi21
        _ = v22
    }
    v3 += -27
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
