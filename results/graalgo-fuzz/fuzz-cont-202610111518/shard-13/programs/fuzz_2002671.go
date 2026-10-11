package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
func genFn0() uint32 {
    var v5 int8 = 0
    v6 := GenS1{F0: 0, F1: 127}
    _ = v5
    _ = v6
    return 2
}

func main() {
    var v1 uint = 65535
    v2 := []string{"", "", ("" + "hello")}
    var v3 float64 = -2.25
    var v4 string = ("\n\t" + "")
    var v7 float32 = (-2.25 - (1.5 + 0.1))
    v3 -= 1e10
    for fi8 := 0; fi8 < 11; fi8++ {
        v7 += float32((float64(v7) * (v3 + v3)))
        _ = fi8
    }
    v4 = v4
    fmt.Println(v7, v1)
    var v9 int8 = 72
    v10 := GenS1{F0: 0, F1: 0}
    var v11 uint8 = 128
    fmt.Println(v3, v11)
    fmt.Println(v7, v1)
    if (v4 >= "foo bar") {
        var v12 int = -334
        var v13 uint32 = 8
        if false {
            if true {
                v1 = (v1 + (v1 ^ (v1 - 65535)))
                var v14 uint64 = 2
                _ = v14
            }
        }
        _ = v12
        _ = v13
    }
    fmt.Println(v1)
    for fi15 := 0; fi15 < 1; fi15++ {
        v9 = -1
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi15
    }
    for fi16 := 0; fi16 < 3; fi16++ {
        var v17 uint64 = 903
        v4 = ("\n\t" + ((v4 + "hello") + v4))
        var v18 uint32 = 442
        _ = fi16
        _ = v17
        _ = v18
    }
    v7 = float32(v3)
    if len(v2) > 0 {
        v2[0] = "a"
    }
    v7 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v9
}
