package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32 }
func genFn0(p0 uint16, p1 uint8) int16 {
    p0 = (p0 & (p0 / ((p0) | 1)))
    p0 = p0
    for fi6 := 0; fi6 < 5; fi6++ {
        var v7 uint64 = 4294967295
        _ = fi6
        _ = v7
    }
    v8 := []int{458, -2147483648, -533}
    _ = p0
    _ = p1
    _ = v8
    return -260
}

func main() {
    var v1 int64 = 0
    v2 := []int{-13}
    var v3 int = 127
    v4 := []int{(v3 / ((-661) | 1)), 128}
    var v5 uint64 = 1
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi9 := 0; fi9 < 4; fi9++ {
        var v10 uint8 = 121
        var v11 bool = ((!((^(v5)) == v5)) && ("\n\t" != ("a" + ("hello" + ((("foo bar" + "foo bar") + ("a" + "\n\t")) + ("hello" + "\n\t"))))))
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 int64 = (v1 & (v1 % ((42) | 1)))
    var v13 bool = false
    v14 := GenS1{F0: 1, F1: "\n\t", F2: 0}
    if len(v2) > 0 {
        v2[0] = v3
    }
    v5 = v5
    var v15 float64 = 3.14159
    v16 := GenS1{F0: 2147483647, F1: "foo bar", F2: 137}
    if v13 {
        switch v3 {
        case -709:
            _ = len(v4)
        default:
            if len(v4) > 0 {
                v4[0] = v3
            }
        }
    }
    switch ("foo bar" + ("\n\t" + (("foo bar" + "hello") + ("foo bar" + ("" + "\n\t"))))) {
    case "foo bar":
        var v17 uint16 = 883
        _ = v17
    default:
        v18 := GenS2{F0: -203}
        _ = v18
    }
    v19 := GenS3{F0: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
