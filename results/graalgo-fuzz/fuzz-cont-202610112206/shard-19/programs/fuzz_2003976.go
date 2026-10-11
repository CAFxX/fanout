package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 bool }
type GenS2 struct { F0 int; F1 int16 }
func genFn0() int64 {
    var v6 bool = true
    for fi7 := 0; fi7 < 14; fi7++ {
        var v8 uint64 = 2
        var v9 uint = 379
        _ = fi7
        _ = v8
        _ = v9
    }
    v6 = (0 >= 873)
    _ = v6
    return 42
}

func main() {
    v1 := make(map[int]string)
    var v2 int = 126
    v3 := make(map[string]uint32)
    v4 := []uint32{}
    v5 := make(map[int]string)
    var v10 int8 = 100
    v11 := GenS1{F0: 0, F1: "foo bar", F2: true}
    switch (v2 % (((v2 | -1)) | 1)) {
    case -906:
        v2 += 32767
    case 49:
        v12 := GenS2{F0: 32768, F1: 1000}
        _ = v12
    case 127:
        var v13 bool = ("\n\t" != ("hello" + "hello"))
        _ = v13
    }
    v2 += v2
    v10 -= v10
    delete(v1, (v2 >> ((18446744073709551615) & 63)))
    var v14 string = ""
    fmt.Println(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
