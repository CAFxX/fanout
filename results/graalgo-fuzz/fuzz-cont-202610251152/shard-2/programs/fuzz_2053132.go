package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 int32 }
func genFn0(p0 int32, p1 string, p2 uint16) int32 {
    var v7 int64 = 1
    v8 := GenS1{F0: p1}
    p2 += (p2 / ((p2) | 1))
    var v9 uint = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return p0
}

func main() {
    v1 := []string{"foo bar"}
    var v2 int16 = 1000
    var v3 uint32 = 1
    v4 := make(map[uint32]uint32)
    var v5 bool = (false || (true || (426 == v2)))
    var v6 int8 = -75
    switch -772 {
    case 127:
        for fi10 := 0; fi10 < 13; fi10++ {
            if (v5 && v5) {
                for fi11 := 0; fi11 < 18; fi11++ {
                    fmt.Println(v3, fi11)
                    var v12 int = -9223372036854775808
                    _ = fi11
                    _ = v12
                }
                v5 = (4294967295 == 196)
            } else {
                switch (("foo bar" + "foo bar") + "\n\t") {
                case "hello":
                    v3 = v3
                default:
                    fmt.Println(fi10)
                }
            }
            fmt.Println(v2, fi10, v2)
            for fi13 := 0; fi13 < 14; fi13++ {
                var v14 int16 = v2
                _ = len(v1)
                _ = fi13
                _ = v14
            }
            _ = fi10
        }
    }
    var v15 uint = 65535
    if len(v1) > 0 {
        v1[0] = "\n\t"
    }
    if false {
        v16 := GenS2{F0: v6, F1: -865}
        var v17 string = ""
        _ = v16
        _ = v17
    }
    v15 -= (^(v15))
    _ = len(v1)
    var v18 uint32 = 736
    var v19 float64 = ((1.5 - (1.5 - (1e10 / (0.0 - 0.1)))) * 0.0)
    var v20 int16 = v2
    var v21 int32 = 0
    v5 = false
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
