package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 uint; F1 int8; F2 int8 }
func genFn0(p0 int8, p1 int16, p2 uint8) int16 {
    switch "hello" {
    case "foo bar":
        switch 128 {
        case -9223372036854775808:
            var v5 uint16 = 255
            _ = v5
        }
    case "a":
        var v6 bool = ((!false) && false)
        _ = v6
    case "":
        v7 := []int{4294967295, 4294967295, 32767, 127}
        _ = v7
    }
    var v8 string = ("hello" + "foo bar")
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return p1
}

func main() {
    var v1 uint8 = 5
    v2 := []uint32{2147483647, 2, 4294967295, 2147483648}
    v3 := []int{42}
    v4 := []uint32{441, 2147483648, 2}
    for fi9 := 0; fi9 < 20; fi9++ {
        fmt.Println(v1, fi9)
        _ = fi9
    }
    v1 = v1
    v10 := GenS1{F0: 4294967295, F1: (^(v1))}
    if true {
        v1 = (v1 / (((^(v1))) | 1))
        v4 = append(v4, 42)
        for fi11 := 0; fi11 < 20; fi11++ {
            var v12 uint8 = v1
            _ = fi11
            _ = v12
        }
    }
    var v13 bool = (!true)
    var v14 float64 = (0.1 - float64(0.0))
    for fi15 := 0; fi15 < 3; fi15++ {
        var v16 bool = (v1 <= v1)
        var v17 uint8 = 128
        _ = fi15
        _ = v16
        _ = v17
    }
    var v18 int8 = 11
    v19 := GenS2{F0: 65535, F1: (v18 >> ((42) & 63)), F2: v18}
    for fi20 := 0; fi20 < 15; fi20++ {
        var v21 uint16 = 1
        _ = fi20
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
