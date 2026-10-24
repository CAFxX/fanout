package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string }

func main() {
    var v1 int8 = 0
    v2 := make(map[string]uint32)
    v3 := []string{("" + "a")}
    var v4 int64 = 51
    for fi5 := 0; fi5 < 20; fi5++ {
        var v6 int = fi5
        fmt.Println(v4, fi5, v1)
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: 127}
    for fi8 := 0; fi8 < 5; fi8++ {
        v9 := GenS2{F0: v4}
        var v10 int64 = v4
        _ = len(v2)
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := GenS3{F0: "foo bar"}
    var v12 float64 = ((1.5 - 3.14159) / (0.1 + (0.0 * (-2.25 + (((1e10 - 1e10) * ((1e10 - (1.5 * float64(0.1))) * 1e10)) - float64(1.5))))))
    for fi13 := 0; fi13 < 18; fi13++ {
        v2["\n\t"] = 534
        var v14 int64 = v4
        _ = fi13
        _ = v14
    }
    v15 := GenS1{F0: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
