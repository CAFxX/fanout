package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 int8; F1 int16 }

func main() {
    var v1 float64 = (3.14159 / -2.25)
    v2 := make(map[int]int)
    var v3 uint = 18446744073709551615
    var v4 bool = (false || true)
    var v5 uint32 = 1
    v3 -= (v3 % (((v3 << ((uint(v5)) & 63))) | 1))
    var v6 uint64 = 4294967295
    v7 := GenS1{F0: (^(v6)), F1: -9223372036854775808}
    var v8 string = (("" + "hello") + "foo bar")
    v1 += v1
    var v9 float32 = ((-2.25 / 1e10) * -2.25)
    _ = len(v2)
    switch v8 {
    case "foo bar":
        delete(v2, 127)
    case "":
        v10 := []int{-2147483648, 127, 127}
        _ = v10
    }
    if (128 <= 0) {
        var v11 string = ""
        _ = v11
    }
    switch 1 {
    case -1:
        v6 = (v6 / ((v6) | 1))
    }
    fmt.Println(v5, v3, v4)
    if (v1 <= v1) {
        v12 := GenS2{F0: -44, F1: 1000}
        if true {
            delete(v2, 32768)
        } else {
            v3 = 2
            var v13 float32 = v9
            _ = v13
        }
        v14 := []int{32767}
        _ = v12
        _ = v14
    }
    var v15 uint64 = 952
    v16 := []int{-2147483648, 32767, 54}
    for fi17 := 0; fi17 < 15; fi17++ {
        fmt.Println(v3, fi17)
        _ = fi17
    }
    for fi18 := 0; fi18 < 12; fi18++ {
        var v19 int8 = -128
        _ = fi18
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
