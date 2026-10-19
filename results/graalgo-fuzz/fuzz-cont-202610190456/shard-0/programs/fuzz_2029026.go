package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 int8; F1 uint64; F2 bool }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int8 }

func main() {
    var v1 uint32 = 753
    v2 := []string{"a", "hello", "hello"}
    v3 := make(map[int]bool)
    var v4 int32 = 6
    var v5 int32 = -2147483648
    v6 := GenS1{F0: 42, F1: v1}
    v7 := GenS2{F0: 49, F1: 724, F2: true}
    v5 = (v5 / (((v4 / (((v5 >> ((143) & 63))) | 1))) | 1))
    v8 := GenS3{F0: 92}
    v1 += (v1 & (^(v1)))
    v4 = v5
    switch "" {
    case "":
        var v9 int64 = -1
        _ = v9
    case "\n\t":
        for fi10 := 0; fi10 < 6; fi10++ {
            if (!false) {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                switch ("\n\t" + ("" + "")) {
                case "":
                    v11 := GenS4{F0: -96}
                    _ = v11
                default:
                    fmt.Println(v5)
                }
            }
            _ = len(v3)
            _ = fi10
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
