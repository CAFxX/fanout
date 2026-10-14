package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 uint }
type GenS2 struct { F0 int32; F1 int8 }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]int)
    var v3 int64 = -1
    v4 := []string{(("foo bar" + "") + "foo bar"), "", ""}
    for fi5 := 0; fi5 < 13; fi5++ {
        v6 := GenS1{F0: 919, F1: "foo bar", F2: 255}
        var v7 float64 = float64((((float32((1e10 - (((0.1 / 1e10) + 0.1) - 1e10))) / (float32((1e10 * 0.1)) - (3.14159 - (-2.25 / 3.14159)))) / 3.14159) / 0.1))
        for fi8 := 0; fi8 < 12; fi8++ {
            var v9 string = (("foo bar" + "\n\t") + "a")
            var v10 string = "\n\t"
            v9 = ""
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    v3 = (v3 * v3)
    v3 += -9223372036854775808
    _ = len(v4)
    var v11 bool = (255 != 44)
    delete(v1, 1)
    if v11 {
        var v12 int8 = 127
        for fi13 := 0; fi13 < 9; fi13++ {
            var v14 bool = false
            var v15 float64 = 1e10
            var v16 uint = 66
            _ = fi13
            _ = v14
            _ = v15
            _ = v16
        }
        var v17 uint32 = 65535
        _ = v12
        _ = v17
    } else {
        fmt.Println(v3, v3)
    }
    var v18 int16 = 208
    var v19 string = ("" + "")
    var v20 int64 = v3
    v21 := GenS2{F0: -446, F1: 16}
    var v22 bool = false
    v23 := GenS1{F0: 4294967295, F1: v19, F2: 0}
    var v24 float64 = 1.5
    v25 := v2[65535]
    fmt.Println(v22)
    if false {
        v19 = "a"
        v11 = v11
        var v26 uint8 = 2
        _ = v26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
