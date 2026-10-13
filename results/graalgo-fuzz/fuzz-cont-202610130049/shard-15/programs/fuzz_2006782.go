package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32; F1 uint8 }

func main() {
    var v1 float64 = 1.5
    var v2 string = ("" + "hello")
    v3 := make(map[uint32]string)
    var v4 bool = (false || false)
    var v5 int64 = 1
    if ((255 > 255) || (606 > 1)) {
        if false {
            fmt.Println(v2, v4)
            v6 := GenS1{F0: 588}
            _ = v6
        } else {
            if v4 {
                var v7 float32 = (1.5 - ((1e10 - 1e10) + 3.14159))
                _ = v7
            }
        }
        var v8 uint8 = 42
        fmt.Println(v5)
        _ = v8
    } else {
        var v9 int = 4294967295
        if (4294967295 != 18446744073709551615) {
            v10 := []int{}
            for fi11 := 0; fi11 < 4; fi11++ {
                var v12 uint64 = 1
                v2 = "\n\t"
                v2 = (("\n\t" + ("\n\t" + ("" + v2))) + v2)
                _ = fi11
                _ = v12
            }
            _ = v10
        } else {
            fmt.Println(v5, v4)
            delete(v3, 1)
        }
        var v13 int64 = (v5 + v5)
        _ = v13
        _ = v9
    }
    if v4 {
        v14 := []int{9223372036854775807, 2147483647}
        var v15 int64 = v5
        _ = v14
        _ = v15
    } else {
        for fi16 := 0; fi16 < 5; fi16++ {
            if v4 {
                var v17 uint16 = 255
                _ = v17
            }
            _ = fi16
        }
        var v18 float32 = float32(0.1)
        var v19 int64 = v5
        _ = v18
        _ = v19
    }
    v20 := []int{9223372036854775807}
    _ = len(v20)
    var v21 uint64 = 4294967295
    switch "a" {
    case "hello":
        v22 := GenS2{F0: 42, F1: 151}
        _ = v22
    }
    v23 := GenS2{F0: 808, F1: 14}
    v4 = ((4294967295 >= 4294967295) && (1 < -1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
