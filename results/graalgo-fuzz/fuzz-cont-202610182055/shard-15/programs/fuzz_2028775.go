package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64; F1 uint16; F2 int8 }
type GenS4 struct { F0 bool; F1 uint; F2 bool }

func main() {
    var v1 int16 = 952
    var v2 float64 = 0.0
    v3 := make(map[uint32]uint32)
    var v4 bool = false
    v5 := make(map[int]uint32)
    for fi6 := 0; fi6 < 14; fi6++ {
        if false {
            v7 := GenS1{F0: 18446744073709551615, F1: 1}
            v8 := []int{(fi6 | fi6), fi6}
            _ = v7
            _ = v8
        }
        fmt.Println(v4, v1)
        var v9 int16 = (v1 & (v1 >> ((255) & 63)))
        _ = fi6
        _ = v9
    }
    v10 := GenS1{F0: 18446744073709551615, F1: -5}
    v11 := []int{127}
    v12 := GenS2{F0: 761}
    switch "" {
    case "\n\t":
        v13 := GenS3{F0: 42, F1: 32767, F2: 0}
        _ = v13
    }
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    v14 := GenS4{F0: (-2147483648 == -2147483648), F1: 387, F2: (("\n\t" + (("" + "\n\t") + ("\n\t" + "hello"))) == ("a" + ("" + (("\n\t" + "hello") + (("a" + ((("\n\t" + ("foo bar" + ("hello" + "hello"))) + "foo bar") + ("" + "\n\t"))) + "hello")))))}
    var v15 string = "hello"
    if v4 {
        v2 = 1.5
        for fi16 := 0; fi16 < 6; fi16++ {
            var v17 uint64 = 18446744073709551615
            v17 = (v17 ^ (v17 ^ 0))
            _ = fi16
            _ = v17
        }
    }
    var v18 int16 = v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
