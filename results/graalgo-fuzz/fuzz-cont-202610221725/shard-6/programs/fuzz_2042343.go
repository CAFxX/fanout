package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 bool }
type GenS2 struct { F0 uint; F1 string; F2 uint32 }

func main() {
    var v1 int16 = -32768
    var v2 int8 = 100
    var v3 uint16 = 255
    fmt.Println(v3, v1, v1)
    var v4 int64 = 1
    if true {
        v5 := GenS1{F0: "", F1: -1, F2: (-1 == 42)}
        var v6 uint8 = 2
        _ = v5
        _ = v6
    } else {
        var v7 int64 = v4
        v8 := GenS2{F0: 18446744073709551615, F1: "\n\t", F2: 2}
        _ = v7
        _ = v8
    }
    switch 839 {
    case 1:
        for fi9 := 0; fi9 < 12; fi9++ {
            var v10 uint8 = 190
            var v11 float64 = float64(-2.25)
            _ = fi9
            _ = v10
            _ = v11
        }
    case 32768:
        v12 := []int{370, 573, -9}
        _ = v12
    case 2147483647:
        var v13 string = ((("hello" + "\n\t") + "\n\t") + "hello")
        _ = v13
    default:
        var v14 float32 = 0.0
        _ = v14
    }
    v15 := []int{4294967295, 597}
    switch -540 {
    case 2147483647:
        for fi16 := 0; fi16 < 4; fi16++ {
            var v17 uint = 42
            fmt.Println(v1)
            _ = fi16
            _ = v17
        }
    }
    fmt.Println(v4, v1)
    v1 = v1
    _ = len(v15)
    v2 += (v2 * v2)
    if true {
        var v18 string = "a"
        _ = v18
    }
    var v19 string = (("\n\t" + "\n\t") + "a")
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
