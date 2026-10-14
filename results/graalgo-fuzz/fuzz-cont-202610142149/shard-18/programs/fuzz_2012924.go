package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int; F1 uint }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]string)
    v3 := []string{"\n\t", ("foo bar" + "foo bar")}
    v4 := make(map[uint32]uint32)
    var v5 int32 = 0
    v6 := []string{"hello", "a"}
    v7 := GenS1{F0: -924}
    if false {
        if ((v5 != -496) || true) {
            var v8 float64 = (1e10 / 1.0)
            if len(v6) > 0 {
                v6[0] = ((("\n\t" + ("\n\t" + "hello")) + "\n\t") + ("" + "a"))
            }
            var v9 int = 0
            _ = v8
            _ = v9
        } else {
            v10 := GenS1{F0: 1}
            for fi11 := 0; fi11 < 19; fi11++ {
                v5 += v5
                _ = fi11
            }
            _ = v10
        }
    }
    v5 = v5
    v4[0] = 973
    v12 := GenS2{F0: 770, F1: 18446744073709551615}
    var v13 uint = 18446744073709551615
    v13 = v13
    fmt.Println(v5, v13)
    if len(v3) > 0 {
        v3[0] = "a"
    }
    for fi14 := 0; fi14 < 19; fi14++ {
        v13 = (v13 / (((v13 >> ((357) & 63))) | 1))
        _ = fi14
    }
    for fi15 := 0; fi15 < 17; fi15++ {
        var v16 string = "\n\t"
        var v17 float32 = 3.14159
        _ = fi15
        _ = v16
        _ = v17
    }
    var v18 uint = v13
    var v19 string = "\n\t"
    v19 = (v19 + v19)
    if (2 <= 42) {
        var v20 bool = (!true)
        _ = v20
    }
    for fi21 := 0; fi21 < 4; fi21++ {
        var v22 string = "hello"
        fmt.Println(v22)
        fmt.Println(v18)
        _ = fi21
        _ = v22
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
