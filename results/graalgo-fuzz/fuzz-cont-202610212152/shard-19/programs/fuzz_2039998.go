package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }

func main() {
    v1 := []uint32{1}
    v2 := make(map[int]uint32)
    v3 := []string{("hello" + "hello")}
    delete(v2, -2147483648)
    for fi4 := 0; fi4 < 12; fi4++ {
        fmt.Println(fi4)
        fmt.Println(fi4)
        _ = fi4
    }
    var v5 int = 4294967295
    v6 := GenS1{F0: 2, F1: (false && true)}
    for fi7 := 0; fi7 < 4; fi7++ {
        v5 += fi7
        v2[fi7] = 42
        if (false || (!(true || (!true)))) {
            v5 -= v5
            var v8 float64 = (-2.25 - 0.1)
            _ = v8
        }
        _ = fi7
    }
    var v9 int64 = -984
    fmt.Println(v9)
    v9 -= -1
    switch "\n\t" {
    case "hello":
        if true {
            v9 += (v9 + v9)
        } else {
            v9 = (v9 & v9)
            var v10 int = (v5 - 771)
            _ = v10
        }
    case "":
        for fi11 := 0; fi11 < 11; fi11++ {
            v5 += 4294967295
            var v12 int32 = 42
            _ = fi11
            _ = v12
        }
    }
    if (0.1 > -2.25) {
        v9 += 9223372036854775807
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v9
}
