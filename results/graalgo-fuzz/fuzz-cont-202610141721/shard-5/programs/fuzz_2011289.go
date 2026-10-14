package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 int16; F1 int; F2 uint64 }
type GenS3 struct { F0 uint8 }

func main() {
    var v1 int32 = -1
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]uint32)
    v4 := []string{"a", "foo bar", "a"}
    v4 = append(v4, "hello")
    var v5 int32 = v1
    fmt.Println(v5, v5)
    for fi6 := 0; fi6 < 5; fi6++ {
        var v7 bool = false
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 2; fi8++ {
        switch "a" {
        case "hello":
            if (!((0.1 / -2.25) < (float64((1.5 / 3.14159)) - float64(0.0)))) {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                if (("foo bar" <= "") || (2 == 599)) {
                    var v9 uint16 = 625
                    _ = v9
                }
                switch fi8 {
                case 4294967295:
                    var v10 int64 = 978
                    _ = v10
                case 42:
                    v11 := GenS1{F0: (v5 | (v1 & v1)), F1: 42}
                    _ = v11
                }
            }
        default:
            v12 := GenS1{F0: (v1 + (v1 % ((-486) | 1))), F1: -668}
            _ = v12
        }
        switch "\n\t" {
        case "\n\t":
            v5 += v1
        }
        _ = fi8
    }
    for fi13 := 0; fi13 < 2; fi13++ {
        v14 := GenS2{F0: 113, F1: fi13, F2: 0}
        var v15 bool = ("a" >= "\n\t")
        fmt.Println(v15, v1)
        _ = fi13
        _ = v14
        _ = v15
    }
    v5 = v5
    delete(v2, 209)
    delete(v2, 209)
    v1 += v5
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v1, v5)
    v16 := GenS3{F0: 127}
    v5 = v1
    var v17 int = 866
    if len(v4) > 0 {
        v4[0] = ""
    }
    v18 := GenS1{F0: (v1 >> ((255) & 63)), F1: 42}
    for fi19 := 0; fi19 < 15; fi19++ {
        var v20 int64 = 9223372036854775807
        fmt.Println(v5)
        v21 := GenS2{F0: -32768, F1: fi19, F2: 768}
        _ = fi19
        _ = v20
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
