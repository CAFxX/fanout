package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 uint64 }

func main() {
    v1 := make(map[string]string)
    var v2 string = ("hello" + "\n\t")
    var v3 uint = 1
    v4 := GenS1{F0: 396, F1: 2147483647}
    v3 = v3
    v5 := GenS1{F0: 2, F1: 2}
    v1[""] = "hello"
    v3 = v3
    var v6 int = -1
    if (!(true && false)) {
        fmt.Println(v2, v3, v3)
    }
    if false {
        var v7 string = "hello"
        var v8 uint32 = 2147483647
        v7 = v2
        _ = v7
        _ = v8
    } else {
        var v9 int = v6
        _ = v9
    }
    var v10 uint8 = 209
    v10 += v10
    var v11 float64 = -2.25
    fmt.Println(v3, v11, v10)
    v12 := v1[v2]
    for fi13 := 0; fi13 < 4; fi13++ {
        v14 := GenS2{F0: uint64(v10)}
        if (v6 > v6) {
            v2 = v2
        } else {
            delete(v1, (((("" + ((v2 + ("" + (((v12 + v12) + v2) + v2))) + "")) + v12) + "hello") + (("\n\t" + ("\n\t" + v2)) + v12)))
        }
        v15 := GenS1{F0: 32767, F1: 2147483648}
        _ = fi13
        _ = v14
        _ = v15
    }
    var v16 int32 = 383
    v17 := GenS2{F0: 2}
    v18 := []int{(^(v6)), 554}
    v19 := v1[v12]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
