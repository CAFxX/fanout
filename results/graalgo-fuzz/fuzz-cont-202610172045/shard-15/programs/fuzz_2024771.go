package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 int8; F1 uint32 }
func genFn0() uint8 {
    var v4 int32 = 1
    v4 += v4
    _ = v4
    return 255
}
func genFn1() bool {
    var v5 uint32 = 0
    var v6 int32 = -430
    _ = v5
    _ = v6
    return true
}

func main() {
    var v1 uint = 255
    var v2 bool = false
    var v3 int16 = -1
    var v7 int16 = 32767
    var v8 string = ""
    for fi9 := 0; fi9 < 20; fi9++ {
        if true {
            v10 := GenS1{F0: v1, F1: 870}
            _ = v10
        }
        _ = fi9
    }
    if (v8 == (v8 + (v8 + v8))) {
        v8 = v8
    }
    for fi11 := 0; fi11 < 1; fi11++ {
        v12 := []int{(fi11 ^ 2147483647)}
        v1 -= v1
        _ = fi11
        _ = v12
    }
    v13 := GenS2{F0: 127, F1: 335}
    v2 = (!false)
    v8 = v8
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
