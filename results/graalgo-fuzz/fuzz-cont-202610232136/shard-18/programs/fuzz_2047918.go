package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 uint16; F1 uint8; F2 string }
func genFn0(p0 uint, p1 int, p2 bool) int64 {
    p1 -= 0
    fmt.Println(p0, p0, p0)
    v5 := GenS1{F0: 249, F1: 243}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 378
}
func genFn1(p0 uint, p1 int32, p2 uint16) uint {
    p2 = p2
    if (!false) {
        var v6 bool = false
        p0 -= (p0 & (p0 << ((255) & 63)))
        _ = v6
    }
    p2 = uint16(p0)
    _ = p0
    _ = p1
    _ = p2
    return p0
}

func main() {
    var v1 float32 = 3.14159
    var v2 int8 = 127
    var v3 uint = 4294967295
    var v4 int32 = 320
    v7 := []int{284, 4294967295, 0, 42}
    v8 := GenS1{F0: 1, F1: 0}
    var v9 int32 = (v4 ^ v4)
    for fi10 := 0; fi10 < 20; fi10++ {
        var v11 uint = v3
        switch ("hello" + "hello") {
        case "foo bar":
            v12 := GenS1{F0: 1, F1: 42}
            _ = v12
        case "":
            fmt.Println(v9)
        case "hello":
            fmt.Println(v9, v9, v9)
        }
        _ = fi10
        _ = v11
    }
    if (1 > -963) {
        fmt.Println(v9, v3)
        _ = len(v7)
    } else {
        fmt.Println(v1, v9, v9)
        var v13 int32 = v4
        v14 := GenS2{F0: 392, F1: 1, F2: "hello"}
        _ = v13
        _ = v14
    }
    var v15 float64 = 0.0
    _ = len(v7)
    v15 = -2.25
    _ = len(v7)
    for fi16 := 0; fi16 < 11; fi16++ {
        _ = len(v7)
        if false {
            fmt.Println(v9, v3, v3)
        } else {
            fmt.Println(v4)
            var v17 string = "\n\t"
            _ = v17
        }
        _ = len(v7)
        _ = fi16
    }
    var v18 int = -914
    var v19 string = "\n\t"
    v19 = v19
    v15 = 3.14159
    var v20 int32 = (v9 >> ((genFn1(v3, v9, 1)) & 63))
    for fi21 := 0; fi21 < 2; fi21++ {
        v3 -= v3
        v22 := GenS1{F0: 4294967295, F1: 17}
        v18 = 32768
        _ = fi21
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
