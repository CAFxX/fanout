package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 int32 }

func main() {
    v1 := make(map[string]string)
    var v2 float64 = -2.25
    var v3 bool = true
    v1["hello"] = ("foo bar" + "\n\t")
    fmt.Println(v3)
    fmt.Println(v3, v2)
    v2 += v2
    fmt.Println(v2, v2)
    v4 := []int{384, 4294967295, -1, 42}
    v5 := GenS1{F0: 767, F1: -27, F2: 1}
    for fi6 := 0; fi6 < 14; fi6++ {
        fmt.Println(fi6, fi6, v2)
        if len(v4) > 0 {
            v4[0] = fi6
        }
        _ = fi6
    }
    for fi7 := 0; fi7 < 10; fi7++ {
        v8 := GenS1{F0: fi7, F1: 42, F2: 799}
        for fi9 := 0; fi9 < 18; fi9++ {
            var v10 float64 = v2
            if (1 >= 18446744073709551615) {
                switch fi7 {
                case -866:
                    var v11 int32 = 1
                    v12 := GenS1{F0: (fi7 >> ((1) & 63)), F1: 1, F2: v11}
                    _ = v11
                    _ = v12
                default:
                    var v13 int64 = -1
                    _ = v13
                }
            }
            _ = fi9
            _ = v10
        }
        v4 = append(v4, fi7)
        _ = fi7
        _ = v8
    }
    v1["a"] = "a"
    v3 = (v3 || (v3 || (88 == 255)))
    var v14 int16 = 1000
    var v15 uint = 42
    var v16 uint64 = 0
    if v3 {
        fmt.Println(v15, v15)
        v14 = (v14 & v14)
        var v17 int32 = 2147483647
        _ = v17
    } else {
        var v18 int64 = -1
        _ = v18
    }
    var v19 string = ("" + "a")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
