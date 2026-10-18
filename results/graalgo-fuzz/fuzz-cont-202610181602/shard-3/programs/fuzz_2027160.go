package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    var v1 uint = 758
    var v2 int = 1
    var v3 float64 = ((((float64(0.1) + 0.1) * float64((((0.1 + (0.0 + float32(1.5))) - 0.1) / ((3.14159 * -2.25) / 0.1)))) / 1e10) - float64(0.1))
    v4 := []string{}
    var v5 uint16 = 32767
    var v6 int64 = -9223372036854775808
    var v7 int8 = -38
    if (192 >= 2147483648) {
        v8 := GenS1{F0: (!true)}
        _ = v8
    }
    v5 = v5
    v7 = v7
    fmt.Println(v5, v6, v7)
    if true {
        for fi9 := 0; fi9 < 13; fi9++ {
            var v10 string = "\n\t"
            var v11 uint32 = 4294967295
            _ = fi9
            _ = v10
            _ = v11
        }
        if len(v4) > 0 {
            v4[0] = "foo bar"
        }
        v1 += 42
    } else {
        var v12 float64 = 0.1
        _ = v12
    }
    for fi13 := 0; fi13 < 10; fi13++ {
        fmt.Println(v3, v1, v1)
        var v14 float32 = 3.14159
        _ = fi13
        _ = v14
    }
    switch "\n\t" {
    case "":
        v2 = 128
    default:
        v5 = 0
    }
    if (true && (0 > 42)) {
        v15 := GenS1{F0: false}
        var v16 int16 = -691
        v5 = v5
        _ = v15
        _ = v16
    }
    v3 = float64(((((0.1 / 3.14159) * 3.14159) + 0.1) - 0.1))
    _ = len(v4)
    for fi17 := 0; fi17 < 19; fi17++ {
        v5 += 426
        v18 := GenS1{F0: true}
        _ = fi17
        _ = v18
    }
    fmt.Println(v6, v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
