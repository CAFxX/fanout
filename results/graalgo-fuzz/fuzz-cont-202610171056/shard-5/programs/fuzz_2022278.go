package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 int8; F1 int8; F2 int8 }
type GenS3 struct { F0 uint64; F1 uint64 }
type GenS4 struct { F0 uint16; F1 uint8; F2 int8 }

func main() {
    v1 := []string{"foo bar", "a"}
    v2 := make(map[uint32]bool)
    v3 := []int{1, -688, 127}
    var v4 float32 = (1e10 / -2.25)
    var v5 string = ""
    if false {
        for fi6 := 0; fi6 < 18; fi6++ {
            switch 0 {
            case 2147483647:
                v5 = (v5 + v5)
            case 0:
                v7 := GenS1{F0: "", F1: 1}
                _ = v7
            }
            var v8 uint8 = 2
            _ = fi6
            _ = v8
        }
    } else {
        switch v5 {
        case "foo bar":
            var v9 uint8 = 2
            _ = v9
        default:
            v4 += v4
        }
        for fi10 := 0; fi10 < 15; fi10++ {
            v11 := v2[1]
            _ = fi10
            _ = v11
        }
    }
    v12 := GenS2{F0: -2, F1: 0, F2: 39}
    v13 := GenS1{F0: "foo bar", F1: 4294967295}
    switch 128 {
    case 4294967295:
        for fi14 := 0; fi14 < 2; fi14++ {
            switch (fi14 >> ((2) & 63)) {
            case 128:
                v5 = (v5 + "hello")
            case 127:
                if false {
                    fmt.Println(fi14, v4, fi14)
                    v5 = ""
                }
            }
            var v15 uint = 840
            _ = fi14
            _ = v15
        }
    case 42:
        for fi16 := 0; fi16 < 14; fi16++ {
            v4 += ((1e10 / v4) * 1e10)
            if len(v3) > 0 {
                v3[0] = (fi16 + (^(fi16)))
            }
            _ = fi16
        }
    default:
        switch v5 {
        case "":
            v17 := GenS3{F0: 2, F1: 732}
            _ = v17
        case "a":
            delete(v2, 193)
        default:
            v18 := GenS4{F0: 977, F1: 1, F2: -1}
            _ = v18
        }
    }
    var v19 bool = false
    v4 += (v4 / (v4 * (0.1 - v4)))
    var v20 uint64 = 18446744073709551615
    _ = len(v1)
    v21 := GenS4{F0: 506, F1: 2, F2: -1}
    v5 = "foo bar"
    fmt.Println(v5)
    fmt.Println(v4, v5, v5)
    v4 -= -2.25
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
