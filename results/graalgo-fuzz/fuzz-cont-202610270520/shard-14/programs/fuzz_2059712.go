package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 int8 }
func genFn0() string {
    var v7 float32 = -2.25
    v8 := GenS1{F0: 0, F1: 1, F2: 0}
    v9 := []int{4294967295, 32768, 128, 127}
    _ = v7
    _ = v8
    _ = v9
    return "hello"
}

func main() {
    var v1 float64 = -2.25
    var v2 int64 = -360
    v3 := make(map[int]int)
    var v4 string = ""
    v5 := make(map[int]string)
    v6 := []string{v4}
    var v10 int = -989
    switch ((v4 + "a") + ((((v4 + "foo bar") + ("" + (("" + "foo bar") + (("\n\t" + v4) + (v4 + (v4 + (v4 + "a"))))))) + v4) + v4)) {
    case "":
        if len(v6) > 0 {
            v6[0] = v4
        }
    }
    for fi11 := 0; fi11 < 17; fi11++ {
        for fi12 := 0; fi12 < 8; fi12++ {
            v5[v10] = v4
            switch v4 {
            case "\n\t":
                switch 128 {
                case -515:
                    var v13 uint32 = 534
                    _ = v13
                case 1:
                    fmt.Println(v10, fi11)
                }
            case "hello":
                v14 := v3[128]
                _ = v14
            default:
                fmt.Println(v2, fi12, v10)
            }
            _ = fi12
        }
        if (4294967295 < 42) {
            var v15 int64 = v2
            v10 = 127
            v2 -= v2
            _ = v15
        }
        var v16 float32 = ((((1.5 + 1e10) + 0.1) - 1e10) / (1e10 / 1.0))
        _ = fi11
        _ = v16
    }
    v17 := GenS1{F0: 1, F1: 42, F2: 127}
    if len(v6) > 0 {
        v6[0] = v4
    }
    v10 = -494
    delete(v5, v10)
    _ = len(v6)
    var v18 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
