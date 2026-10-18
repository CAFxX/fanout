package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 int; F1 int; F2 int64 }
type GenS3 struct { F0 uint; F1 int64; F2 int32 }
type GenS4 struct { F0 int16; F1 uint32; F2 uint32 }
type GenS5 struct { F0 int32; F1 string; F2 int8 }

func main() {
    v1 := make(map[int]string)
    var v2 uint8 = 127
    v3 := []int{32767, 62}
    var v4 float64 = 3.14159
    var v5 int16 = 1000
    var v6 uint16 = 32767
    v7 := GenS1{F0: "", F1: 2}
    switch "" {
    case "":
        v8 := GenS2{F0: 1, F1: 986, F2: -1}
        _ = v8
    case "hello":
        fmt.Println(v4, v6, v6)
    default:
        var v9 int = 67
        _ = v9
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v3 = append(v3, 1)
    fmt.Println(v6, v6)
    var v10 float32 = 0.0
    switch ("hello" + "\n\t") {
    case "hello":
        v3 = append(v3, -2147483648)
    case "\n\t":
        var v11 uint64 = 18446744073709551615
        _ = v11
    }
    fmt.Println(v2)
    var v12 int8 = 34
    fmt.Println(v6)
    for fi13 := 0; fi13 < 20; fi13++ {
        if (false || ((v12 < (v12 >> ((2) & 63))) || (2 != 449))) {
            v2 = 2
            v5 = v5
        } else {
            v6 = v6
            if (57 != 1) {
                switch "hello" {
                case "":
                    var v14 string = (("\n\t" + "a") + ("" + ("hello" + "")))
                    _ = v14
                case "\n\t":
                    v5 = (v5 ^ v5)
                }
                v4 -= ((((0.1 * v4) / v4) * v4) + (v4 - v4))
                v15 := GenS1{F0: "foo bar", F1: 18446744073709551615}
                _ = v15
            }
        }
        v16 := GenS3{F0: 4294967295, F1: 42, F2: -2147483648}
        var v17 float32 = float32(((1.5 / ((0.0 * 1e10) - v4)) * 1e10))
        _ = fi13
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 20; fi18++ {
        v19 := GenS4{F0: 42, F1: 2, F2: 2147483648}
        var v20 string = ("foo bar" + ("foo bar" + ("\n\t" + "a")))
        v21 := GenS4{F0: v5, F1: 2, F2: 42}
        _ = fi18
        _ = v19
        _ = v20
        _ = v21
    }
    if (769 > 1) {
        for fi22 := 0; fi22 < 16; fi22++ {
            for fi23 := 0; fi23 < 10; fi23++ {
                switch "hello" {
                case "hello":
                    switch "\n\t" {
                    case "":
                        fmt.Println(v5, fi22, v5)
                    }
                }
                if (v4 < 1e10) {
                    var v24 bool = true
                    v25 := GenS5{F0: 42, F1: ("" + ""), F2: (v12 << ((726) & 63))}
                    v12 -= (v12 - 0)
                    _ = v24
                    _ = v25
                } else {
                    _ = len(v1)
                }
                _ = fi23
            }
            v26 := GenS3{F0: 0, F1: 9223372036854775807, F2: 42}
            _ = fi22
            _ = v26
        }
        fmt.Println(v4)
        v27 := GenS2{F0: 0, F1: 32768, F2: 9223372036854775807}
        _ = v27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
