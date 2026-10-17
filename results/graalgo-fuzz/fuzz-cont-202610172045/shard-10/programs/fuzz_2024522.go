package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 bool; F1 int32; F2 uint }

func main() {
    var v1 uint32 = 65535
    var v2 uint8 = 128
    v3 := make(map[string]int)
    var v4 float32 = float32(1.5)
    v5 := []int{362, 137}
    var v6 int = 32768
    v1 = (v1 / (((v1 * (v1 & (v1 * (v1 | (^(v1))))))) | 1))
    fmt.Println(v1, v4)
    switch 128 {
    case -2147483648:
        fmt.Println(v6)
    case 32768:
        fmt.Println(v4, v4, v1)
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v7 bool = true
    for fi8 := 0; fi8 < 15; fi8++ {
        var v9 uint32 = (v1 & (^(v1)))
        v7 = (v2 > v2)
        fmt.Println(v1, v4)
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: 0}
    if true {
        switch (v6 + (v6 - v6)) {
        case 0:
            v6 = -44
        case 2147483647:
            v11 := GenS1{F0: 2147483647}
            _ = v11
        default:
            if ("" > "a") {
                for fi12 := 0; fi12 < 9; fi12++ {
                    var v13 uint64 = 4294967295
                    fmt.Println(v4)
                    _ = fi12
                    _ = v13
                }
                v14 := v3[""]
                if (18446744073709551615 != 2) {
                    var v15 uint16 = 2
                    _ = v15
                }
                _ = v14
            } else {
                v5 = append(v5, (v6 | v6))
            }
        }
        var v16 uint = 198
        _ = v16
    }
    var v17 int32 = 1
    _ = len(v3)
    if v7 {
        v18 := GenS2{F0: (v6 ^ int(v2)), F1: ""}
        if (v7 || v7) {
            var v19 float64 = (-2.25 - -2.25)
            v3["hello"] = 128
            if v7 {
                var v20 int8 = 100
                switch "a" {
                case "hello":
                    _ = len(v3)
                }
                _ = v20
            } else {
                if true {
                    v7 = v7
                }
                v2 = 127
                v21 := GenS3{F0: (!(1 > 65535)), F1: v17, F2: 42}
                _ = v21
            }
            _ = v19
        }
        _ = v18
    }
    v22 := GenS1{F0: (v17 & (v17 - (^(v17))))}
    var v23 uint64 = 18446744073709551615
    var v24 float32 = v4
    v23 += v23
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
