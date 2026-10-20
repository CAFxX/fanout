package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 string }
type GenS2 struct { F0 bool }

func main() {
    var v1 float32 = (-2.25 - 1e10)
    var v2 float64 = 1.5
    var v3 uint32 = 965
    var v4 string = ("" + "foo bar")
    var v5 uint32 = (v3 / (((v3 << ((4294967295) & 63))) | 1))
    v6 := []int{-1}
    for fi7 := 0; fi7 < 16; fi7++ {
        if (!(!true)) {
            var v8 int = fi7
            v4 = v4
            _ = v8
        } else {
            var v9 uint64 = 280
            switch 1 {
            case 0:
                if (-1 >= 42) {
                    if (74 != 710) {
                        var v10 int32 = 42
                        _ = v10
                    }
                } else {
                    var v11 int = fi7
                    if true {
                        var v12 int32 = 787
                        v6 = append(v6, fi7)
                        _ = v12
                    } else {
                        v13 := GenS1{F0: 1, F1: 2, F2: ((v4 + v4) + v4)}
                        _ = v13
                    }
                    _ = v11
                }
            }
            switch (fi7 - 0) {
            case 42:
                switch v4 {
                case "\n\t":
                    var v14 float64 = v2
                    _ = v14
                default:
                    v15 := GenS1{F0: 250, F1: 128, F2: (v4 + "")}
                    _ = v15
                }
            }
            _ = v9
        }
        var v16 float32 = v1
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = fi7
        _ = v16
    }
    v17 := GenS2{F0: false}
    var v18 uint32 = v5
    if len(v6) > 0 {
        v6[0] = 4294967295
    }
    var v19 int32 = 0
    var v20 int64 = 9223372036854775807
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    fmt.Println(v3, v2, v3)
    v4 = v4
    for fi21 := 0; fi21 < 4; fi21++ {
        var v22 string = v4
        v3 += (v3 * 119)
        _ = fi21
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
