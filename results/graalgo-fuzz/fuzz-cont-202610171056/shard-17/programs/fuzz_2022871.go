package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 uint; F1 uint32 }

func main() {
    v1 := make(map[string]string)
    var v2 int64 = 502
    var v3 float64 = 0.1
    v4 := make(map[int]int)
    v5 := make(map[string]int)
    v6 := []string{("\n\t" + "")}
    v3 = (float64(float32((v3 + float64(1e10)))) / 1e10)
    for fi7 := 0; fi7 < 5; fi7++ {
        fmt.Println(fi7, v2, fi7)
        var v8 float32 = -2.25
        _ = fi7
        _ = v8
    }
    v3 = v3
    var v9 uint32 = 1
    if false {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        var v10 float64 = v3
        _ = v10
    } else {
        var v11 int8 = 121
        var v12 uint = 18446744073709551615
        _ = v11
        _ = v12
    }
    v13 := GenS1{F0: 358, F1: "a"}
    var v14 float32 = 1.5
    if ((!(!(-393 == 0))) && (true && true)) {
        if (true || true) {
            fmt.Println(v2)
        }
    }
    v3 += float64(1.5)
    v15 := GenS2{F0: 65535, F1: (v9 - (v9 - 42))}
    var v16 float64 = 0.1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
