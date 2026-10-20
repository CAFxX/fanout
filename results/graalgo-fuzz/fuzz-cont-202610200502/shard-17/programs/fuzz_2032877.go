package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32 }
func genFn0() uint8 {
    if false {
        var v4 uint = 42
        if true {
            var v5 uint8 = 42
            var v6 int = 0
            _ = v5
            _ = v6
        }
        fmt.Println(v4)
        _ = v4
    }
    switch ("a" + "foo bar") {
    case "":
        var v7 uint8 = 1
        _ = v7
    case "foo bar":
        var v8 uint = 42
        _ = v8
    default:
        var v9 int = 996
        _ = v9
    }
    var v10 string = ""
    var v11 uint64 = 42
    _ = v10
    _ = v11
    return 170
}
func genFn1(p0 uint32) string {
    switch 0 {
    case -1:
        p0 = p0
    case 128:
        var v12 string = "foo bar"
        _ = v12
    case 0:
        fmt.Println(p0)
    default:
        v13 := []int{49}
        _ = v13
    }
    var v14 uint = 42
    _ = p0
    _ = v14
    return "a"
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []string{"\n\t", ("a" + "\n\t"), "a"}
    v3 := []uint32{42, 2147483648, 0}
    if len(v3) > 0 {
        v3[0] = 884
    }
    var v15 uint64 = 0
    v16 := GenS1{F0: 255}
    var v17 uint = 631
    var v18 string = ("" + "a")
    v2 = append(v2, (v18 + v18))
    var v19 uint16 = 109
    for fi20 := 0; fi20 < 20; fi20++ {
        v17 -= v17
        v21 := GenS2{F0: -1}
        _ = fi20
        _ = v21
    }
    v3 = append(v3, 4294967295)
    if true {
        var v22 float64 = 3.14159
        v17 = (v17 / ((v17) | 1))
        _ = v22
    } else {
        fmt.Println(v19)
    }
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
