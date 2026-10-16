package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 int8; F1 int16; F2 int8 }
type GenS3 struct { F0 bool; F1 uint; F2 int }
func genFn0() int32 {
    v4 := GenS1{F0: 1, F1: 2147483648}
    if (!false) {
        v5 := GenS1{F0: 42, F1: 0}
        if true {
            v6 := []int{}
            v7 := GenS1{F0: -2147483648, F1: 2147483647}
            for fi8 := 0; fi8 < 4; fi8++ {
                fmt.Println(fi8)
                _ = fi8
            }
            _ = v6
            _ = v7
        }
        _ = v5
    } else {
        var v9 float32 = 1.5
        switch (("foo bar" + (("a" + "foo bar") + "")) + "") {
        case "a":
            if true {
                v9 = v9
                v10 := GenS2{F0: -1, F1: 18, F2: 0}
                _ = v10
            }
        }
        v9 = v9
        _ = v9
    }
    switch ("\n\t" + "") {
    case "":
        var v11 uint32 = 4294967295
        _ = v11
    default:
        var v12 uint64 = 457
        _ = v12
    }
    var v13 float64 = 1e10
    _ = v13
    _ = v4
    return -2147483648
}

func main() {
    var v1 uint32 = 2
    var v2 bool = false
    var v3 int8 = 0
    v14 := GenS1{F0: 42, F1: 0}
    var v15 uint16 = 42
    var v16 uint16 = (^(v15))
    switch ("\n\t" + "foo bar") {
    case "a":
        var v17 uint64 = uint64(v16)
        _ = v17
    case "hello":
        v18 := GenS1{F0: -637, F1: (v1 >> ((697) & 63))}
        _ = v18
    default:
        fmt.Println(v16, v15, v15)
    }
    v19 := []int{-2147483648, 4294967295}
    v20 := GenS3{F0: (-495 <= 777), F1: 42, F2: 344}
    var v21 float32 = 0.1
    if (-1 != -933) {
        fmt.Println(v1)
        var v22 string = "a"
        var v23 bool = (-360 <= 755)
        _ = v22
        _ = v23
    } else {
        fmt.Println(v21, v3, v2)
        v16 = (^(v15))
    }
    fmt.Println(v15)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
