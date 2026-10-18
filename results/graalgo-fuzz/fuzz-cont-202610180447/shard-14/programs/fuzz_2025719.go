package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 uint64 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 string, p1 uint8, p2 uint32) int8 {
    var v7 int = 127
    var v8 int64 = 0
    for fi9 := 0; fi9 < 3; fi9++ {
        var v10 int8 = 127
        v10 = (v10 * (v10 % ((v10) | 1)))
        if true {
            var v11 int16 = 42
            _ = v11
        }
        _ = fi9
        _ = v10
    }
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 100
}

func main() {
    var v1 float64 = 3.14159
    var v2 uint8 = 0
    var v3 float64 = (v1 / (v1 / v1))
    v4 := []int{}
    var v5 int64 = 0
    var v6 bool = (127 > 42)
    v5 = (v5 ^ (v5 % ((v5) | 1)))
    if true {
        v3 += (v3 - -2.25)
        _ = len(v4)
    }
    switch 1 {
    case 4294967295:
        var v12 string = (("foo bar" + "") + "")
        _ = v12
    default:
        if (v6 || v6) {
            switch "a" {
            case "\n\t":
                var v13 int64 = -9223372036854775808
                _ = v13
            case "hello":
                var v14 int64 = v5
                _ = v14
            case "foo bar":
                if v6 {
                    v3 += (v3 * 0.1)
                }
            default:
                var v15 uint64 = 18446744073709551615
                _ = v15
            }
        }
    }
    v16 := GenS1{F0: (-(v5)), F1: "", F2: 1}
    v17 := GenS2{F0: v2}
    v1 -= v1
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v18 uint = 65535
    for fi19 := 0; fi19 < 19; fi19++ {
        v18 = v18
        v1 = (0.1 * v1)
        _ = len(v4)
        _ = fi19
    }
    var v20 uint64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
