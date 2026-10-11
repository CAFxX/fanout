package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 int }
type GenS2 struct { F0 int; F1 int32; F2 uint }
type GenS3 struct { F0 int32 }

func main() {
    v1 := []int{526, 0, -218}
    v2 := make(map[string]int)
    v3 := []string{}
    var v4 string = ""
    v5 := []string{v4, (v4 + ("hello" + (("hello" + v4) + "")))}
    var v6 int8 = 42
    var v7 string = v4
    for fi8 := 0; fi8 < 11; fi8++ {
        if false {
            var v9 uint = 2
            _ = v9
        }
        _ = fi8
    }
    if ((0.0 / -2.25) < 1.5) {
        v6 -= v6
        v10 := GenS1{F0: (v6 | 57), F1: false, F2: 43}
        _ = v10
    } else {
        fmt.Println(v4, v6, v4)
    }
    v11 := GenS2{F0: 4294967295, F1: 2147483647, F2: 2}
    _ = len(v2)
    var v12 int8 = v6
    fmt.Println(v6)
    fmt.Println(v7, v6, v6)
    var v13 int8 = v12
    fmt.Println(v7, v6)
    for fi14 := 0; fi14 < 3; fi14++ {
        _ = len(v2)
        var v15 int8 = (v13 - v13)
        v6 -= v15
        _ = fi14
        _ = v15
    }
    if false {
        var v16 uint32 = 809
        _ = v16
    } else {
        v17 := GenS3{F0: 42}
        fmt.Println(v7, v6, v4)
        for fi18 := 0; fi18 < 18; fi18++ {
            v7 = (v7 + (("hello" + v4) + v7))
            _ = fi18
        }
        _ = v17
    }
    var v19 int32 = -1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
