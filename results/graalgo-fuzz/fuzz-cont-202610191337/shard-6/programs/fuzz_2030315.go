package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int; F1 string; F2 int8 }
func genFn0(p0 int32, p1 int32, p2 int32) uint8 {
    v7 := []int{}
    var v8 int8 = 42
    v9 := GenS1{F0: 65535}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 0
}
func genFn1() uint64 {
    v10 := GenS1{F0: 255}
    var v11 float32 = (3.14159 / (1.5 / ((0.1 - float32(1e10)) / 0.1)))
    _ = v10
    _ = v11
    return 4294967295
}

func main() {
    var v1 uint = 424
    v2 := make(map[int]string)
    var v3 int64 = 0
    var v4 float64 = (1.5 / 1.0)
    v5 := []string{"a", "hello", "foo bar", (("a" + "a") + ("" + "hello"))}
    v6 := []int{}
    var v12 float64 = (-2.25 + (v4 + 3.14159))
    for fi13 := 0; fi13 < 7; fi13++ {
        fmt.Println(v1)
        v12 = v12
        _ = fi13
    }
    if true {
        var v14 int = 128
        _ = v14
    } else {
        for fi15 := 0; fi15 < 9; fi15++ {
            switch "foo bar" {
            case "a":
                if false {
                    delete(v2, fi15)
                    v1 -= (v1 - 255)
                    fmt.Println(v1, v12)
                }
            case "":
                var v16 int32 = 2147483647
                _ = v16
            case "hello":
                v4 = (v12 / 1.0)
            default:
                var v17 float32 = 1.5
                _ = v17
            }
            if len(v5) > 0 {
                v5[0] = ""
            }
            v4 = v12
            _ = fi15
        }
    }
    var v18 uint16 = 45
    fmt.Println(v1, v3, v12)
    v12 -= v12
    for fi19 := 0; fi19 < 17; fi19++ {
        v2[(fi19 | (fi19 | fi19))] = "foo bar"
        _ = fi19
    }
    v2[42] = ""
    v12 += 3.14159
    v20 := GenS1{F0: 3}
    switch "foo bar" {
    case "":
        for fi21 := 0; fi21 < 6; fi21++ {
            for fi22 := 0; fi22 < 6; fi22++ {
                var v23 string = "foo bar"
                _ = fi22
                _ = v23
            }
            v12 -= v4
            var v24 int16 = 465
            _ = fi21
            _ = v24
        }
    case "hello":
        delete(v2, -909)
    case "a":
        var v25 uint8 = 42
        _ = v25
    }
    for fi26 := 0; fi26 < 1; fi26++ {
        var v27 uint = v1
        v28 := GenS2{F0: (fi26 % ((fi26) | 1)), F1: ("hello" + "\n\t"), F2: 42}
        _ = fi26
        _ = v27
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
