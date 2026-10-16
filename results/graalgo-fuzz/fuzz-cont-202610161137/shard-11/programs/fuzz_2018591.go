package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16 }
func genFn0(p0 int64, p1 int8, p2 int64) string {
    for fi5 := 0; fi5 < 13; fi5++ {
        v6 := []int{2147483647, fi5, (fi5 - 9223372036854775807), fi5}
        fmt.Println(p1, fi5)
        _ = fi5
        _ = v6
    }
    fmt.Println(p1, p1, p2)
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    return ("\n\t" + ("foo bar" + (("" + "hello") + "foo bar")))
}

func main() {
    var v1 string = ("" + "hello")
    v2 := []int{9223372036854775807, 128}
    var v3 int8 = 100
    v4 := make(map[uint32]string)
    if true {
        var v7 int16 = -1
        v8 := GenS1{F0: false}
        for fi9 := 0; fi9 < 11; fi9++ {
            v10 := v4[2]
            _ = fi9
            _ = v10
        }
        _ = v7
        _ = v8
    } else {
        v3 -= v3
    }
    v11 := v4[2147483647]
    v12 := GenS2{F0: 408}
    var v13 uint = 255
    v1 = ((v11 + (v11 + "\n\t")) + "foo bar")
    switch "foo bar" {
    case "foo bar":
        v14 := GenS1{F0: (!(true && false))}
        _ = v14
    }
    if (false || false) {
        var v15 int64 = -9223372036854775808
        var v16 int16 = 42
        fmt.Println(v13, v15, v11)
        _ = v15
        _ = v16
    } else {
        fmt.Println(v1)
    }
    var v17 float64 = (3.14159 - float64((0.0 * 1e10)))
    var v18 float64 = (-2.25 + (float64(0.1) / (v17 / v17)))
    var v19 int32 = 883
    var v20 int32 = (v19 % ((v19) | 1))
    if true {
        fmt.Println(v11, v17, v18)
    }
    var v21 float64 = v17
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
