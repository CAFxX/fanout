package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 uint8 }

func main() {
    var v1 uint16 = 42
    v2 := make(map[int]string)
    var v3 float64 = 0.0
    v4 := make(map[string]string)
    v5 := make(map[uint32]int)
    var v6 int = 32767
    var v7 int16 = 1
    var v8 int64 = 42
    var v9 uint8 = 127
    var v10 bool = true
    for fi11 := 0; fi11 < 1; fi11++ {
        switch ("a" + (((((("" + (("hello" + "") + ("" + ("a" + (("" + "hello") + (("\n\t" + ("\n\t" + "")) + (("a" + "hello") + ("foo bar" + "")))))))) + "foo bar") + "foo bar") + ("a" + "foo bar")) + "hello") + "foo bar")) {
        case "a":
            v7 = v7
        default:
            v1 += 600
        }
        _ = fi11
    }
    var v12 int64 = (-(v8))
    v13 := []int{(v6 / (((v6 & 32768)) | 1))}
    var v14 int8 = 126
    for fi15 := 0; fi15 < 13; fi15++ {
        fmt.Println(fi15, v7, v8)
        fmt.Println(v14, v14, v7)
        _ = fi15
    }
    v6 = v6
    v16 := GenS1{F0: 703, F1: (v6 << ((4294967295) & 63))}
    v17 := GenS2{F0: v9}
    var v18 int = v6
    for fi19 := 0; fi19 < 17; fi19++ {
        var v20 string = "a"
        v7 = 1
        var v21 uint = 255
        _ = fi19
        _ = v20
        _ = v21
    }
    for fi22 := 0; fi22 < 14; fi22++ {
        fmt.Println(v10, v1, v8)
        _ = fi22
    }
    switch v18 {
    case -2147483648:
        var v23 uint64 = uint64(v1)
        _ = v23
    }
    var v24 float64 = (v3 * v3)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
