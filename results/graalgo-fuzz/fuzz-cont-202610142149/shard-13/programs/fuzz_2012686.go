package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 int; F2 uint64 }
type GenS3 struct { F0 bool; F1 bool; F2 int }
type GenS4 struct { F0 uint8; F1 uint8 }

func main() {
    var v1 uint32 = 2
    v2 := []string{"hello", ("\n\t" + "a"), "a", "foo bar"}
    var v3 int64 = -9223372036854775808
    var v4 int64 = v3
    v5 := make(map[int]int)
    var v6 uint = 2
    var v7 string = "\n\t"
    switch v7 {
    case "foo bar":
        for fi8 := 0; fi8 < 5; fi8++ {
            v2 = append(v2, "foo bar")
            switch 110 {
            case 575:
                v3 = v3
            case 344:
                delete(v5, (fi8 + fi8))
            default:
                v9 := GenS1{F0: -23}
                _ = v9
            }
            for fi10 := 0; fi10 < 18; fi10++ {
                var v11 uint64 = 4294967295
                v4 = v3
                if ((!true) && true) {
                    v12 := GenS1{F0: -1}
                    var v13 int32 = 42
                    v14 := GenS2{F0: 100, F1: fi10, F2: (v11 - v11)}
                    _ = v12
                    _ = v13
                    _ = v14
                }
                _ = fi10
                _ = v11
            }
            _ = fi8
        }
    default:
        v4 = 42
    }
    var v15 float64 = (((float64(0.1) + 0.0) * 0.0) + (0.1 + 0.0))
    for fi16 := 0; fi16 < 12; fi16++ {
        var v17 int = -705
        if true {
            v18 := v5[v17]
            v19 := GenS3{F0: false, F1: false, F2: 1}
            var v20 int64 = (v3 >> ((255) & 63))
            _ = v18
            _ = v19
            _ = v20
        }
        _ = fi16
        _ = v17
    }
    v21 := GenS4{F0: 42, F1: 238}
    v7 = v7
    v3 = -9223372036854775808
    v4 = v4
    switch 18 {
    case -53:
        for fi22 := 0; fi22 < 2; fi22++ {
            var v23 uint8 = 0
            var v24 string = "a"
            _ = fi22
            _ = v23
            _ = v24
        }
    case 128:
        _ = len(v2)
    default:
        v25 := GenS1{F0: 91}
        _ = v25
    }
    var v26 uint16 = 187
    var v27 int64 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v2
    _ = v21
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
