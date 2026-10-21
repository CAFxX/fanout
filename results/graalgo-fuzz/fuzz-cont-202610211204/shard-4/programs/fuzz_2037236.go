package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int16; F1 int16 }
type GenS4 struct { F0 int16 }
func genFn0(p0 int8, p1 bool, p2 uint16) string {
    if p1 {
        p1 = true
        fmt.Println(p2)
        v7 := GenS1{F0: "a"}
        _ = v7
    }
    p1 = p1
    fmt.Println(p2, p0, p2)
    v8 := []int{9223372036854775807, 2147483647, -1, 9223372036854775807}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return "\n\t"
}

func main() {
    var v1 int = 4294967295
    var v2 uint = 95
    v3 := []string{(("hello" + ("foo bar" + (("a" + "\n\t") + "a"))) + ""), "foo bar"}
    var v4 uint16 = uint16(v2)
    v5 := []int{-1, (v1 & (v1 >> ((v2) & 63))), v1}
    v6 := make(map[int]int)
    _ = len(v5)
    if true {
        if false {
            v9 := GenS1{F0: "hello"}
            _ = v9
        } else {
            for fi10 := 0; fi10 < 5; fi10++ {
                if false {
                    var v11 int = v1
                    for fi12 := 0; fi12 < 18; fi12++ {
                        var v13 int64 = -688
                        var v14 int = 910
                        v1 = 32767
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                    fmt.Println(v1, fi10, v2)
                    _ = v11
                }
                _ = fi10
            }
            v1 = 1
            switch v1 {
            case 32767:
                var v15 uint64 = 2
                _ = v15
            case 0:
                v16 := GenS2{F0: -2147483648}
                _ = v16
            default:
                v2 = 0
            }
        }
    }
    fmt.Println(v4)
    v17 := GenS2{F0: -2147483648}
    v1 = v1
    var v18 uint64 = 1
    v18 = 869
    switch "foo bar" {
    case "foo bar":
        var v19 float64 = -2.25
        _ = v19
    case "\n\t":
        v4 = (^(v4))
    }
    var v20 int16 = -1
    v4 -= 0
    v21 := GenS3{F0: v20, F1: (-(v20))}
    var v22 uint = v2
    v23 := v6[v1]
    for fi24 := 0; fi24 < 2; fi24++ {
        v25 := GenS4{F0: v20}
        var v26 float64 = 1.5
        _ = fi24
        _ = v25
        _ = v26
    }
    var v27 uint = v22
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
