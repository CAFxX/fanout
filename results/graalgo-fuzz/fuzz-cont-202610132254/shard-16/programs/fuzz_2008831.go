package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
func genFn0(p0 uint, p1 int32, p2 int32) uint64 {
    v4 := []int{1, 9223372036854775807, 4294967295, 4294967295}
    v4 = append(v4, 926)
    v5 := GenS1{F0: p1, F1: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 0
}
func genFn1(p0 uint64) int8 {
    var v6 float32 = (0.0 + (3.14159 + 0.0))
    var v7 int64 = 0
    fmt.Println(p0, v6)
    for fi8 := 0; fi8 < 1; fi8++ {
        v9 := []int{}
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = v6
    _ = v7
    return 127
}

func main() {
    v1 := make(map[string]int)
    v2 := []int{639}
    v3 := make(map[string]string)
    switch -9223372036854775808 {
    case 32768:
        for fi10 := 0; fi10 < 8; fi10++ {
            var v11 uint16 = 2
            var v12 string = ((("hello" + "\n\t") + "foo bar") + "\n\t")
            _ = fi10
            _ = v11
            _ = v12
        }
    case -9223372036854775808:
        switch ("hello" + "a") {
        case "hello":
            var v13 bool = true
            _ = v13
        case "a":
            var v14 int = 186
            _ = v14
        case "foo bar":
            var v15 float64 = float64(0.1)
            _ = v15
        }
    }
    var v16 int32 = -1
    var v17 uint = 478
    switch 9223372036854775807 {
    case -350:
        v3["hello"] = ("foo bar" + "")
    case 128:
        fmt.Println(v16, v17)
    case 1:
        var v18 int32 = v16
        _ = v18
    default:
        v17 -= (v17 & (v17 * (v17 >> ((2) & 63))))
    }
    for fi19 := 0; fi19 < 13; fi19++ {
        _ = len(v1)
        _ = fi19
    }
    fmt.Println(v16, v16)
    v20 := v3[""]
    v17 = v17
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v3
}
