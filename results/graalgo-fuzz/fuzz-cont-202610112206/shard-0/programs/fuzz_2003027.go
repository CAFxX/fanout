package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }

func main() {
    v1 := make(map[int]string)
    var v2 uint = 18446744073709551615
    var v3 int16 = -204
    var v4 string = ""
    switch (("a" + "hello") + v4) {
    case "a":
        for fi5 := 0; fi5 < 9; fi5++ {
            v6 := GenS1{F0: (true || true), F1: 727}
            v3 += 0
            v4 = v4
            _ = fi5
            _ = v6
        }
    case "":
        _ = len(v1)
    }
    for fi7 := 0; fi7 < 11; fi7++ {
        v3 = 1
        var v8 int8 = 121
        v3 += (v3 & (v3 / (((v3 % ((v3) | 1))) | 1)))
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 18; fi9++ {
        v10 := GenS1{F0: false, F1: 65535}
        var v11 int8 = -1
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 float32 = float32(0.1)
    for fi13 := 0; fi13 < 7; fi13++ {
        var v14 uint16 = 2
        _ = fi13
        _ = v14
    }
    fmt.Println(v4, v3, v12)
    if (-2147483648 >= -2147483648) {
        v1[32768] = "a"
        for fi15 := 0; fi15 < 17; fi15++ {
            var v16 float32 = 0.1
            _ = fi15
            _ = v16
        }
    } else {
        fmt.Println(v2, v12)
        if false {
            var v17 string = (v4 + v4)
            _ = v17
        }
        fmt.Println(v3, v3)
    }
    v18 := []int{-1, -1, -130}
    var v19 uint16 = 1
    fmt.Println(v19, v4)
    var v20 string = v4
    delete(v1, 4294967295)
    for fi21 := 0; fi21 < 8; fi21++ {
        v22 := GenS1{F0: true, F1: 0}
        _ = fi21
        _ = v22
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
