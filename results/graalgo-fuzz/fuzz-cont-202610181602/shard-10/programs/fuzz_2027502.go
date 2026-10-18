package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 int8; F2 int }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint32 }

func main() {
    var v1 float64 = 1.5
    var v2 bool = true
    var v3 uint32 = 373
    var v4 int8 = -109
    v5 := GenS1{F0: 128}
    var v6 float64 = v1
    v4 = (v4 + -128)
    var v7 uint64 = 273
    if (("hello" + "\n\t") == "") {
        fmt.Println(v1)
    }
    fmt.Println(v7, v6)
    var v8 int = int(v3)
    v9 := []int{}
    v2 = false
    v10 := GenS2{F0: uint(v3), F1: 59, F2: v8}
    for fi11 := 0; fi11 < 15; fi11++ {
        for fi12 := 0; fi12 < 12; fi12++ {
            switch -104 {
            case -176:
                v13 := GenS3{F0: (v4 >> ((42) & 63))}
                _ = v13
            case 127:
                v7 = 4294967295
            }
            var v14 int64 = 42
            if v2 {
                v6 = v6
                if (0 <= 32767) {
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                    var v15 bool = (true || true)
                    v16 := GenS4{F0: (^(v3))}
                    v14 = v14
                    _ = v15
                    _ = v16
                }
                _ = len(v9)
            }
            _ = fi12
            _ = v14
        }
        _ = fi11
    }
    v7 = (v7 / (((^(v7))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
