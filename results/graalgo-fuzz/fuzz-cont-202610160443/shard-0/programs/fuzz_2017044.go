package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 uint16; F2 int8 }
type GenS3 struct { F0 uint; F1 int; F2 uint }

func main() {
    var v1 uint32 = 1
    v2 := make(map[uint32]int)
    v3 := []string{("" + "hello"), "foo bar", ""}
    var v4 float32 = 0.1
    if len(v3) > 0 {
        v3[0] = ""
    }
    if false {
        v4 -= (v4 - 0.1)
        fmt.Println(v1)
        var v5 float64 = -2.25
        _ = v5
    } else {
        var v6 int64 = 410
        var v7 int16 = -32768
        _ = v6
        _ = v7
    }
    fmt.Println(v4)
    switch -2147483648 {
    case -808:
        v4 += v4
    }
    v1 = 2
    var v8 int32 = -1
    v1 = (v1 >> ((42) & 63))
    var v9 int64 = 1
    var v10 int16 = 0
    v1 -= v1
    fmt.Println(v9)
    v8 = 42
    v11 := v2[v1]
    if true {
        var v12 uint8 = 42
        fmt.Println(v10, v1, v11)
        if len(v3) > 0 {
            v3[0] = "hello"
        }
        _ = v12
    } else {
        v10 -= (v10 >> ((255) & 63))
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        for fi13 := 0; fi13 < 1; fi13++ {
            for fi14 := 0; fi14 < 9; fi14++ {
                v15 := GenS1{F0: 4294967295}
                v16 := GenS2{F0: 100, F1: 0, F2: 100}
                var v17 uint = 172
                _ = fi14
                _ = v15
                _ = v16
                _ = v17
            }
            _ = fi13
        }
    }
    var v18 uint = 378
    var v19 float64 = float64(v4)
    if (686 >= 740) {
        v19 = v19
        var v20 string = ""
        _ = v20
    }
    v21 := GenS3{F0: v18, F1: 128, F2: 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
