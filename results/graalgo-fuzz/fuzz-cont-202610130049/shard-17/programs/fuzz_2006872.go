package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 uint; F1 int8 }
func genFn0(p0 uint8, p1 int16, p2 uint) bool {
    for fi7 := 0; fi7 < 15; fi7++ {
        switch ("hello" + "\n\t") {
        case "\n\t":
            p2 -= 4294967295
        case "a":
            var v8 uint = (p2 << ((p2) & 63))
            _ = v8
        case "":
            for fi9 := 0; fi9 < 7; fi9++ {
                var v10 uint16 = 917
                _ = fi9
                _ = v10
            }
        }
        _ = fi7
    }
    v11 := []int{-538}
    p2 = p2
    var v12 uint16 = 42
    for fi13 := 0; fi13 < 14; fi13++ {
        var v14 int8 = -128
        _ = fi13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return false
}
func genFn1(p0 uint16) uint {
    var v15 uint8 = 255
    p0 += p0
    switch (("\n\t" + "") + "") {
    case "foo bar":
        switch "" {
        case "hello":
            var v16 float64 = 0.0
            _ = v16
        }
    case "":
        v17 := GenS1{F0: 888, F1: 42}
        _ = v17
    case "hello":
        v15 = 127
    default:
        v18 := GenS1{F0: -849, F1: 472}
        _ = v18
    }
    _ = p0
    _ = v15
    return 18446744073709551615
}

func main() {
    var v1 uint16 = 977
    var v2 uint8 = 0
    var v3 uint8 = 2
    v4 := make(map[string]string)
    var v5 int16 = -882
    v6 := []string{"hello", "hello", "hello"}
    var v19 int = 128
    _ = len(v6)
    _ = len(v6)
    var v20 int = int(v5)
    var v21 bool = true
    var v22 int = v19
    for fi23 := 0; fi23 < 3; fi23++ {
        var v24 int8 = 127
        for fi25 := 0; fi25 < 2; fi25++ {
            v22 = (fi25 / (((v20 * v22)) | 1))
            _ = fi25
        }
        v4["hello"] = "foo bar"
        _ = fi23
        _ = v24
    }
    if true {
        v26 := GenS1{F0: 2147483647, F1: 1}
        fmt.Println(v21, v20)
        _ = v26
    }
    for fi27 := 0; fi27 < 5; fi27++ {
        v21 = (0.1 >= 3.14159)
        _ = fi27
    }
    v5 += (v5 / ((v5) | 1))
    v28 := GenS2{F0: 1, F1: 87}
    var v29 uint64 = 131
    v20 = (^(v19))
    fmt.Println(v1)
    var v30 int8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
