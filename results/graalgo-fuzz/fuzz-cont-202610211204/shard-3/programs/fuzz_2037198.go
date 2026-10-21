package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 uint16 }
type GenS2 struct { F0 uint64; F1 int64; F2 int32 }

func main() {
    var v1 float64 = float64(1.5)
    v2 := []string{("foo bar" + "foo bar"), "foo bar"}
    v3 := make(map[string]int)
    v4 := make(map[int]uint32)
    var v5 uint64 = 2
    fmt.Println(v1)
    var v6 string = "foo bar"
    v6 = v6
    fmt.Println(v5)
    var v7 int = 127
    switch "a" {
    case "":
        switch v6 {
        case "foo bar":
            var v8 uint = 752
            _ = v8
        case "\n\t":
            v9 := GenS1{F0: 42, F1: 42, F2: 0}
            _ = v9
        case "a":
            v10 := GenS1{F0: 0, F1: 32767, F2: 65535}
            _ = v10
        }
    case "\n\t":
        if (false && (-1 < 9223372036854775807)) {
            v11 := GenS2{F0: 0, F1: -892, F2: -2147483648}
            v2 = append(v2, v6)
            _ = v11
        }
    case "foo bar":
        v5 += (v5 % ((0) | 1))
    }
    var v12 float64 = 0.1
    for fi13 := 0; fi13 < 11; fi13++ {
        if false {
            v7 = int(v5)
        }
        _ = fi13
    }
    var v14 int16 = 1
    v6 = ("\n\t" + "")
    var v15 int = v7
    var v16 int64 = 9223372036854775807
    var v17 int32 = -1
    var v18 float64 = (v1 + (v1 / v12))
    for fi19 := 0; fi19 < 18; fi19++ {
        for fi20 := 0; fi20 < 6; fi20++ {
            fmt.Println(v5)
            _ = fi20
        }
        _ = fi19
    }
    delete(v3, v6)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
