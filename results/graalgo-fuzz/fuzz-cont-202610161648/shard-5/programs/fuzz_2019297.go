package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 string = "\n\t"
    var v2 int = 128
    v3 := []string{v1, "\n\t"}
    if len(v3) > 0 {
        v3[0] = v1
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v4 float64 = (1e10 + 0.0)
    var v5 string = ("\n\t" + ("\n\t" + "\n\t"))
    var v6 uint = 42
    _ = len(v3)
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 bool = (v4 != 0.0)
        v3 = append(v3, (v1 + "foo bar"))
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 8; fi9++ {
        var v10 uint64 = 978
        var v11 int32 = 162
        switch (v2 ^ fi9) {
        case 877:
            v12 := GenS1{F0: -82}
            _ = v12
        case -9223372036854775808:
            var v13 int8 = 47
            _ = v13
        case -2147483648:
            var v14 float64 = float64(3.14159)
            _ = v14
        default:
            var v15 string = v1
            _ = v15
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    for fi16 := 0; fi16 < 4; fi16++ {
        var v17 string = ("\n\t" + v5)
        _ = fi16
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
