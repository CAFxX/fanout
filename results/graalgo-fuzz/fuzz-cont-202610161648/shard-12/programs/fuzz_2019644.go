package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 int16 }
type GenS2 struct { F0 int; F1 int8 }

func main() {
    v1 := make(map[uint32]string)
    v2 := []string{"", "foo bar", ("foo bar" + "foo bar"), "a"}
    var v3 int = -1
    var v4 float64 = (1.5 * (-2.25 + 1.5))
    v5 := make(map[uint32]int)
    var v6 int32 = -987
    for fi7 := 0; fi7 < 15; fi7++ {
        v8 := GenS1{F0: 2, F1: 42, F2: 0}
        _ = fi7
        _ = v8
    }
    if (0 < 1) {
        v2 = append(v2, ("foo bar" + ""))
        for fi9 := 0; fi9 < 9; fi9++ {
            var v10 bool = true
            _ = fi9
            _ = v10
        }
        v11 := GenS1{F0: 4294967295, F1: 128, F2: 42}
        _ = v11
    }
    fmt.Println(v6, v6, v6)
    var v12 int16 = 638
    if true {
        v12 = v12
    } else {
        var v13 int = v3
        v14 := GenS1{F0: 208, F1: 2, F2: (v12 - -807)}
        v4 = v4
        _ = v13
        _ = v14
    }
    v15 := GenS2{F0: -1, F1: -1}
    for fi16 := 0; fi16 < 13; fi16++ {
        if true {
            delete(v1, 293)
            switch 42 {
            case 391:
                var v17 int = v3
                _ = v17
            case 982:
                var v18 uint = 0
                _ = v18
            case 325:
                if true {
                    v4 += ((v4 / 1.5) + 0.1)
                    var v19 int16 = (v12 >> ((42) & 63))
                    _ = v19
                } else {
                    v20 := v1[4294967295]
                    _ = v20
                }
            }
        } else {
            var v21 int64 = 289
            _ = v21
        }
        _ = fi16
    }
    for fi22 := 0; fi22 < 5; fi22++ {
        var v23 int32 = (v6 & 2147483647)
        _ = fi22
        _ = v23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
