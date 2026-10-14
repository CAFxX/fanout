package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 string }

func main() {
    var v1 uint8 = 51
    v2 := []uint32{12, 0}
    v3 := make(map[int]uint32)
    var v4 int32 = -332
    var v5 uint16 = 32767
    v5 += (v5 >> ((255) & 63))
    var v6 int16 = 825
    v5 -= (v5 + (^(v5)))
    fmt.Println(v1)
    var v7 uint = 0
    if false {
        var v8 bool = (-1 == int64(v1))
        for fi9 := 0; fi9 < 20; fi9++ {
            var v10 bool = ((0.1 > -2.25) && v8)
            v11 := GenS1{F0: v5, F1: 2147483648, F2: ("hello" + ("\n\t" + "foo bar"))}
            v4 += int32(v7)
            _ = fi9
            _ = v10
            _ = v11
        }
        if v8 {
            v12 := GenS1{F0: (^(v5)), F1: 101, F2: "hello"}
            _ = v12
        } else {
            for fi13 := 0; fi13 < 5; fi13++ {
                _ = len(v2)
                _ = fi13
            }
        }
        _ = v8
    }
    switch ("a" + ("a" + "\n\t")) {
    case "\n\t":
        var v14 uint32 = 2147483647
        _ = v14
    }
    v15 := GenS1{F0: v5, F1: 4294967295, F2: "foo bar"}
    _ = len(v2)
    v2 = append(v2, 634)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
