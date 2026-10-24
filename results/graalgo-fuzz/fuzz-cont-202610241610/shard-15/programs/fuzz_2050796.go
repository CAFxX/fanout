package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 uint8 }
func genFn0() string {
    v5 := GenS1{F0: 2147483647, F1: 42, F2: 2}
    v6 := []int{2147483647, 128, 127, 4294967295}
    var v7 uint16 = 32767
    _ = v5
    _ = v6
    _ = v7
    return "\n\t"
}
func genFn1(p0 bool, p1 uint64, p2 int32) uint8 {
    if false {
        v8 := GenS1{F0: -731, F1: -1, F2: 0}
        _ = v8
    } else {
        for fi9 := 0; fi9 < 19; fi9++ {
            v10 := []int{127, 42, (fi9 << ((4294967295) & 63)), (fi9 ^ fi9)}
            var v11 uint64 = 825
            var v12 bool = p0
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
        }
    }
    var v13 uint64 = (p1 + (p1 / (((^(p1))) | 1)))
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 1
}

func main() {
    var v1 int8 = -1
    v2 := []int{42, 42}
    var v3 bool = (true && true)
    var v4 float64 = 1e10
    fmt.Println(v1)
    var v14 uint = 760
    v4 = (v4 + (((v4 * v4) / v4) * v4))
    if (255 <= 42) {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v15 float64 = ((v4 * v4) / v4)
        if (!true) {
            switch "" {
            case "foo bar":
                var v16 string = ("foo bar" + "")
                _ = v16
            case "":
                var v17 int16 = 387
                _ = v17
            case "\n\t":
                v3 = v3
            }
        }
        _ = v15
    }
    var v18 uint64 = 1
    switch ("\n\t" + "\n\t") {
    case "\n\t":
        var v19 uint16 = 912
        _ = v19
    case "a":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case "hello":
        var v20 int64 = -689
        _ = v20
    default:
        var v21 uint32 = 2147483647
        _ = v21
    }
    v3 = ("a" < "")
    switch 201 {
    case 42:
        v22 := GenS1{F0: -248, F1: v1, F2: 255}
        _ = v22
    case 2147483647:
        v18 += v18
    }
    if ((v3 || ((v14 - (v14 % ((v14) | 1))) <= v14)) && v3) {
        v2 = append(v2, 127)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
