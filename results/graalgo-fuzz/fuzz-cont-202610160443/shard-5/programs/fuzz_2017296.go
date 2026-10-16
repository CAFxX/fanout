package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string }

func main() {
    var v1 int32 = -460
    var v2 uint16 = 227
    var v3 int8 = 42
    var v4 int32 = v1
    var v5 string = "\n\t"
    for fi6 := 0; fi6 < 16; fi6++ {
        v7 := []int{(fi6 << ((2) & 63)), fi6, fi6}
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: (^(v3))}
    v9 := []int{1, 128, -1, 1}
    if true {
        var v10 int32 = (v1 + (v4 - v1))
        v4 -= v4
        _ = v10
    } else {
        for fi11 := 0; fi11 < 20; fi11++ {
            fmt.Println(v3)
            _ = fi11
        }
    }
    v12 := GenS2{F0: (v5 + "\n\t")}
    if true {
        v13 := GenS1{F0: (^(v3))}
        v5 = (("foo bar" + v5) + ((("\n\t" + v5) + "foo bar") + v5))
        v9 = append(v9, 1)
        _ = v13
    } else {
        _ = len(v9)
    }
    v4 = v1
    if false {
        switch 0 {
        case 9223372036854775807:
            var v14 int = -668
            _ = v14
        default:
            if true {
                var v15 int8 = 29
                fmt.Println(v3)
                _ = v15
            }
        }
        v2 += 1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
