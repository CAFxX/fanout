package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 uint8 }
type GenS2 struct { F0 int32; F1 int16 }

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 int16 = 32767
    v3 := []int{223}
    var v4 int16 = (v2 & v2)
    v4 = -804
    if len(v3) > 0 {
        v3[0] = -1
    }
    v4 -= -903
    v5 := GenS1{F0: 458, F1: "\n\t", F2: 89}
    v6 := GenS2{F0: 1, F1: (v2 * -32768)}
    v4 = (-(v2))
    v7 := GenS2{F0: 42, F1: (v2 * v2)}
    fmt.Println(v1)
    fmt.Println(v2)
    v8 := GenS2{F0: 591, F1: (v2 | (v4 % ((v2) | 1)))}
    if true {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        if (-1 != 42) {
            v1 = v1
            for fi9 := 0; fi9 < 10; fi9++ {
                var v10 bool = false
                var v11 int = 4294967295
                if v10 {
                    v1 = (^(v1))
                    if len(v3) > 0 {
                        v3[0] = -1
                    }
                }
                _ = fi9
                _ = v10
                _ = v11
            }
            v1 = v1
        }
        v1 += 2
    }
    v1 += v1
    v3 = append(v3, 128)
    if true {
        v12 := GenS1{F0: -1, F1: "hello", F2: 255}
        switch "hello" {
        case "\n\t":
            switch "hello" {
            case "":
                for fi13 := 0; fi13 < 4; fi13++ {
                    var v14 string = ""
                    fmt.Println(v14, fi13, v1)
                    _ = fi13
                    _ = v14
                }
            case "\n\t":
                var v15 uint32 = 287
                _ = v15
            default:
                var v16 bool = ((v4 ^ (^(v4))) <= (-(v2)))
                _ = v16
            }
        }
        v4 += (^(v4))
        _ = v12
    } else {
        v4 -= v2
    }
    fmt.Println(v4, v1)
    var v17 uint = uint(v1)
    for fi18 := 0; fi18 < 12; fi18++ {
        var v19 string = "a"
        _ = fi18
        _ = v19
    }
    var v20 int64 = int64(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
