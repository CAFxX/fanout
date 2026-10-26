package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 uint; F1 int32; F2 uint16 }
func genFn0(p0 uint16, p1 uint64) bool {
    for fi5 := 0; fi5 < 18; fi5++ {
        var v6 float32 = 3.14159
        var v7 int8 = int8(p0)
        _ = fi5
        _ = v6
        _ = v7
    }
    if (false || (2 == 1)) {
        fmt.Println(p1, p1)
        var v8 int8 = 127
        _ = v8
    }
    _ = p0
    _ = p1
    return true
}
func genFn1(p0 int64, p1 bool) int16 {
    switch 9223372036854775807 {
    case 32768:
        var v9 int8 = 127
        _ = v9
    case 127:
        p0 = p0
    }
    p0 = -9223372036854775808
    p1 = true
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    return 42
}

func main() {
    var v1 bool = (2 < 42)
    var v2 bool = (0 < 42)
    var v3 float64 = 0.0
    var v4 uint16 = 255
    if (42 < 42) {
        var v10 int = 9223372036854775807
        var v11 uint = 65535
        _ = v10
        _ = v11
    } else {
        fmt.Println(v1, v3)
    }
    if (float64((0.0 / 1.0)) != float64((0.0 * 1e10))) {
        switch 127 {
        case 9223372036854775807:
            v12 := []int{32768, -1}
            _ = v12
        case 1:
            v13 := []int{4294967295}
            _ = v13
        }
        v14 := GenS1{F0: 549, F1: ""}
        _ = v14
    }
    v2 = (!false)
    v15 := []int{-9223372036854775808, 834}
    switch -79 {
    case -1:
        var v16 int16 = 1
        _ = v16
    case 4294967295:
        if false {
            v3 = 1.5
            var v17 string = ("a" + "\n\t")
            for fi18 := 0; fi18 < 4; fi18++ {
                v19 := GenS2{F0: 116, F1: 1, F2: v4}
                if len(v15) >= 2 {
                    v15 = v15[1:len(v15)]
                }
                var v20 int64 = 1
                _ = fi18
                _ = v19
                _ = v20
            }
            _ = v17
        }
    case -135:
        v21 := GenS1{F0: 653, F1: "a"}
        _ = v21
    }
    fmt.Println(v3)
    switch 642 {
    case 42:
        v2 = (-114 < 127)
    }
    v4 = 2
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
