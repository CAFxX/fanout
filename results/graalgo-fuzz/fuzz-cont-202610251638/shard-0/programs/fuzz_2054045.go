package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 int }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32 }

func main() {
    var v1 string = (("foo bar" + ("foo bar" + "")) + "hello")
    var v2 int16 = -32768
    var v3 uint = 0
    v4 := make(map[string]string)
    var v5 float32 = -2.25
    v6 := []string{v1}
    if (!false) {
        var v7 int64 = 380
        _ = v7
    } else {
        fmt.Println(v2, v2)
    }
    for fi8 := 0; fi8 < 12; fi8++ {
        v9 := GenS1{F0: (("\n\t" + "\n\t") < "a"), F1: 1, F2: fi8}
        _ = fi8
        _ = v9
    }
    v4[""] = v1
    var v10 string = v1
    var v11 float64 = (1.5 / 0.1)
    var v12 int16 = v2
    var v13 int16 = v12
    if false {
        fmt.Println(v1)
        v13 = (v13 | v13)
    } else {
        if ((v13 - -319) != v12) {
            v14 := GenS1{F0: false, F1: 255, F2: 4294967295}
            _ = v14
        }
    }
    var v15 float64 = v11
    var v16 float32 = v5
    fmt.Println(v10, v13)
    if (-2.25 < (1e10 / v16)) {
        fmt.Println(v11, v10)
    } else {
        v17 := GenS2{F0: 852}
        var v18 float64 = -2.25
        if false {
            v19 := GenS1{F0: false, F1: 11, F2: 32767}
            var v20 string = "foo bar"
            var v21 float64 = v18
            _ = v19
            _ = v20
            _ = v21
        } else {
            v22 := GenS3{F0: -2147483648}
            var v23 string = v10
            var v24 string = "hello"
            _ = v22
            _ = v23
            _ = v24
        }
        _ = v17
        _ = v18
    }
    v11 = v11
    v1 = v10
    v25 := GenS1{F0: (true || false), F1: 255, F2: 352}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
