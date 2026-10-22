package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 string }

func main() {
    v1 := []int{1, 0, 0}
    var v2 uint64 = 2
    v3 := []uint32{302, 1, 333}
    switch "a" {
    case "foo bar":
        var v4 uint64 = (v2 % (((v2 / ((v2) | 1))) | 1))
        _ = v4
    default:
        v5 := GenS1{F0: v2, F1: 777, F2: ""}
        _ = v5
    }
    switch 42 {
    case -2147483648:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case 128:
        var v6 float64 = -2.25
        _ = v6
    }
    fmt.Println(v2)
    var v7 int8 = 50
    var v8 uint = uint(v2)
    for fi9 := 0; fi9 < 19; fi9++ {
        v7 += v7
        fmt.Println(fi9, fi9, fi9)
        fmt.Println(v2, v2)
        _ = fi9
    }
    var v10 string = ((("hello" + ("\n\t" + "a")) + ((("a" + "") + ("foo bar" + "foo bar")) + ("" + "a"))) + "\n\t")
    v11 := GenS1{F0: uint64(v8), F1: 2147483647, F2: ("foo bar" + "")}
    var v12 float64 = (-2.25 * 0.1)
    for fi13 := 0; fi13 < 10; fi13++ {
        for fi14 := 0; fi14 < 7; fi14++ {
            v15 := GenS1{F0: (v2 % (((v2 - v2)) | 1)), F1: -1, F2: "foo bar"}
            var v16 uint64 = (v2 >> (((^(v8))) & 63))
            _ = fi14
            _ = v15
            _ = v16
        }
        switch 1 {
        case 128:
            _ = len(v1)
        }
        _ = fi13
    }
    fmt.Println(v7, v2)
    fmt.Println(v2)
    fmt.Println(v8)
    var v17 uint8 = 236
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
