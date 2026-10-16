package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 bool }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 int32 }

func main() {
    v1 := make(map[string]int)
    var v2 float64 = 3.14159
    var v3 uint16 = 168
    var v4 int32 = 1
    switch "a" {
    case "\n\t":
        var v5 int64 = 1
        _ = v5
    case "hello":
        var v6 string = ("\n\t" + "hello")
        _ = v6
    }
    v3 += (^(v3))
    v7 := GenS1{F0: -9223372036854775808, F1: 1, F2: (1 > 43)}
    if (9223372036854775807 != 1) {
        v8 := GenS2{F0: 2, F1: -1}
        _ = v8
    }
    fmt.Println(v4, v3, v2)
    delete(v1, "\n\t")
    var v9 int64 = 9223372036854775807
    for fi10 := 0; fi10 < 15; fi10++ {
        for fi11 := 0; fi11 < 18; fi11++ {
            v12 := []int{fi10}
            v12 = append(v12, fi11)
            switch (fi10 << ((42) & 63)) {
            case 9223372036854775807:
                if len(v12) > 0 {
                    v12[0] = (fi11 ^ 636)
                }
            case 1:
                v4 = v4
            case 32767:
                if (("a" != "hello") && true) {
                    var v13 float32 = -2.25
                    _ = v13
                } else {
                    v3 = v3
                }
            }
            _ = fi11
            _ = v12
        }
        _ = fi10
    }
    v2 = v2
    fmt.Println(v2, v3, v3)
    for fi14 := 0; fi14 < 20; fi14++ {
        for fi15 := 0; fi15 < 13; fi15++ {
            fmt.Println(v4, v9)
            _ = fi15
        }
        var v16 int32 = (v4 ^ 0)
        _ = fi14
        _ = v16
    }
    v17 := []int{}
    v18 := GenS3{F0: v4}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v9
}
