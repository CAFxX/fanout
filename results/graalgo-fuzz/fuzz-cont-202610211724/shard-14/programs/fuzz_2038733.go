package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
func genFn0() bool {
    for fi5 := 0; fi5 < 7; fi5++ {
        var v6 uint32 = 2
        _ = fi5
        _ = v6
    }
    var v7 int32 = 1
    v8 := []int{-9, -880}
    v9 := GenS1{F0: 119, F1: "foo bar"}
    v7 += 0
    _ = v7
    _ = v8
    _ = v9
    return (true || ((float64((1.5 - float32(1.5))) > 3.14159) || (!(!true))))
}
func genFn1(p0 uint16) uint64 {
    p0 = (p0 * 1)
    v10 := GenS1{F0: 1, F1: (("\n\t" + ("foo bar" + (("a" + "") + ""))) + ("foo bar" + "foo bar"))}
    _ = p0
    _ = v10
    return 18446744073709551615
}

func main() {
    v1 := []int{}
    var v2 float32 = 0.0
    v3 := make(map[uint32]uint32)
    v4 := make(map[string]bool)
    var v11 int = -829
    if false {
        var v12 uint32 = 0
        _ = v12
    } else {
        var v13 float32 = 3.14159
        if (false || (146 > 128)) {
            v13 -= (-2.25 + v2)
        } else {
            v2 = v2
            v13 = (((1e10 - 1e10) - v13) + float32(3.14159))
        }
        _ = v13
    }
    var v14 float64 = (1e10 / 0.1)
    fmt.Println(v11)
    for fi15 := 0; fi15 < 8; fi15++ {
        for fi16 := 0; fi16 < 4; fi16++ {
            fmt.Println(fi15, v2)
            fmt.Println(v2, v2)
            if len(v1) > 0 {
                v1[0] = (v11 + (fi16 / ((-542) | 1)))
            }
            _ = fi16
        }
        _ = fi15
    }
    var v17 float64 = v14
    switch ("" + "\n\t") {
    case "":
        _ = len(v4)
    }
    fmt.Println(v14, v17, v17)
    var v18 float64 = v14
    v14 -= (v17 - (-2.25 + v14))
    var v19 bool = false
    var v20 int16 = 1
    var v21 uint32 = 0
    if v19 {
        v21 -= (^(v21))
    } else {
        if v19 {
            for fi22 := 0; fi22 < 4; fi22++ {
                for fi23 := 0; fi23 < 4; fi23++ {
                    v24 := GenS1{F0: 100, F1: ("\n\t" + "\n\t")}
                    _ = fi23
                    _ = v24
                }
                _ = fi22
            }
        }
        fmt.Println(v20)
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch "" {
    case "foo bar":
        v25 := GenS1{F0: 96, F1: "a"}
        _ = v25
    default:
        v3[v21] = v21
    }
    var v26 string = "foo bar"
    if len(v1) > 0 {
        v1[0] = (v11 | v11)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v26
    _ = v3
    _ = v4
}
