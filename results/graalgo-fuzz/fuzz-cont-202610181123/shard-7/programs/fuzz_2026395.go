package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int; F1 string; F2 int }
type GenS3 struct { F0 uint16; F1 uint; F2 int16 }

func main() {
    var v1 uint32 = 906
    v2 := make(map[string]bool)
    var v3 float32 = 0.1
    var v4 string = "hello"
    var v5 uint32 = (^(v1))
    var v6 uint32 = (v5 ^ v5)
    var v7 float32 = 1e10
    v8 := GenS1{F0: 9223372036854775807}
    v9 := GenS1{F0: 464}
    v10 := v2[("" + (((v4 + v4) + (v4 + ("" + v4))) + v4))]
    var v11 uint = 42
    for fi12 := 0; fi12 < 1; fi12++ {
        for fi13 := 0; fi13 < 10; fi13++ {
            v5 = (v6 + uint32(v11))
            v6 = v1
            _ = fi13
        }
        _ = fi12
    }
    for fi14 := 0; fi14 < 8; fi14++ {
        if (float32(float64((v7 / (v3 / v7)))) != v3) {
            for fi15 := 0; fi15 < 15; fi15++ {
                v1 = (v1 ^ v1)
                _ = fi15
            }
            for fi16 := 0; fi16 < 3; fi16++ {
                var v17 int8 = 95
                _ = fi16
                _ = v17
            }
        }
        _ = fi14
    }
    v5 = 4294967295
    var v18 int32 = -2147483648
    v11 = 15
    v19 := []int{}
    if v10 {
        fmt.Println(v1)
        if true {
            fmt.Println(v3, v6, v6)
            v20 := GenS2{F0: 32768, F1: "a", F2: 962}
            switch ((v4 + v4) + ((v4 + ((v4 + ("a" + "hello")) + (v4 + v4))) + v4)) {
            case "foo bar":
                v21 := GenS3{F0: 908, F1: (v11 * (v11 % ((v11) | 1))), F2: 0}
                _ = v21
            case "a":
                v22 := GenS2{F0: 127, F1: "\n\t", F2: 128}
                _ = v22
            case "":
                v7 = float32(0.0)
            }
            _ = v20
        }
    } else {
        for fi23 := 0; fi23 < 4; fi23++ {
            v24 := GenS1{F0: -343}
            if len(v19) > 0 {
                v19[0] = (fi23 / ((fi23) | 1))
            }
            _ = fi23
            _ = v24
        }
    }
    v1 = v6
    var v25 uint32 = v6
    v26 := GenS3{F0: 0, F1: (v11 / (((v11 << (((v11 << (((v11 % (((v11 ^ (v11 + (^(v11))))) | 1))) & 63))) & 63))) | 1)), F2: -1}
    v3 = v7
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v18
    _ = v19
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
