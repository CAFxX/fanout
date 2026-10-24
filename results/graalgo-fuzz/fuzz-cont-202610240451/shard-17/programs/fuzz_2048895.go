package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 uint8; F1 uint16 }
func genFn0() int64 {
    var v5 int16 = -32768
    for fi6 := 0; fi6 < 19; fi6++ {
        var v7 int = 4294967295
        fmt.Println(fi6)
        v8 := GenS1{F0: "", F1: -10}
        _ = fi6
        _ = v7
        _ = v8
    }
    v9 := []int{-2147483648}
    _ = v5
    _ = v9
    return -1
}
func genFn1() int32 {
    for fi10 := 0; fi10 < 7; fi10++ {
        fmt.Println(fi10)
        for fi11 := 0; fi11 < 9; fi11++ {
            v12 := GenS1{F0: ("hello" + "foo bar"), F1: 127}
            _ = fi11
            _ = v12
        }
        _ = fi10
    }
    var v13 int8 = 127
    _ = v13
    return -322
}

func main() {
    var v1 string = "\n\t"
    var v2 int8 = -1
    v3 := make(map[string]int)
    v4 := []int{-740, 470, 236, 128}
    var v14 uint32 = 42
    var v15 float64 = (0.0 - -2.25)
    var v16 uint16 = 767
    var v17 int64 = genFn0()
    var v18 bool = false
    v2 = 57
    v19 := GenS1{F0: v1, F1: v2}
    v2 += (^(v2))
    fmt.Println(v18)
    _ = len(v3)
    switch ("foo bar" + "hello") {
    case "hello":
        fmt.Println(v16, v17)
    case "a":
        var v20 uint8 = 55
        _ = v20
    case "foo bar":
        v16 = v16
    default:
        v21 := GenS1{F0: v1, F1: (v2 | v2)}
        _ = v21
    }
    for fi22 := 0; fi22 < 8; fi22++ {
        if false {
            delete(v3, "")
            var v23 bool = v18
            v24 := GenS2{F0: 0, F1: v16}
            _ = v23
            _ = v24
        } else {
            v25 := GenS2{F0: 1, F1: 0}
            v26 := GenS1{F0: "", F1: int8(fi22)}
            var v27 int = fi22
            _ = v25
            _ = v26
            _ = v27
        }
        var v28 bool = v18
        _ = fi22
        _ = v28
    }
    fmt.Println(v1)
    if false {
        var v29 uint = 4294967295
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
