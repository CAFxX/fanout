package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 int = 375
    var v2 float64 = 1e10
    var v3 float32 = 0.1
    v4 := []int{}
    v2 = float64(3.14159)
    _ = len(v4)
    v5 := GenS1{F0: "\n\t"}
    switch "" {
    case "hello":
        var v6 int = (v1 % (((v1 + (v1 ^ v1))) | 1))
        _ = v6
    case "foo bar":
        v7 := GenS1{F0: ("\n\t" + "foo bar")}
        _ = v7
    }
    v1 += (v1 | 127)
    if len(v4) > 0 {
        v4[0] = 32768
    }
    for fi8 := 0; fi8 < 3; fi8++ {
        _ = len(v4)
        _ = fi8
    }
    for fi9 := 0; fi9 < 19; fi9++ {
        var v10 int64 = 106
        if false {
            fmt.Println(v3, v1)
        }
        _ = fi9
        _ = v10
    }
    v11 := GenS1{F0: ("foo bar" + "")}
    for fi12 := 0; fi12 < 12; fi12++ {
        for fi13 := 0; fi13 < 4; fi13++ {
            v14 := GenS1{F0: "\n\t"}
            var v15 int8 = 42
            v3 = 0.1
            _ = fi13
            _ = v14
            _ = v15
        }
        v2 = 0.1
        _ = fi12
    }
    if len(v4) > 0 {
        v4[0] = v1
    }
    v3 -= float32((v2 - (v2 - (v2 * v2))))
    for fi16 := 0; fi16 < 5; fi16++ {
        var v17 int = -1
        v4 = append(v4, v17)
        _ = fi16
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
