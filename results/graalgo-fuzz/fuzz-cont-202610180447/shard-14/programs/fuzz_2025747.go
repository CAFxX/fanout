package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 int64 }

func main() {
    v1 := []uint32{1, 2147483648, 4294967295}
    var v2 uint64 = 18446744073709551615
    v3 := []string{"", "a"}
    v2 += 2
    var v4 float64 = 1e10
    var v5 uint32 = 1
    var v6 uint64 = (v2 >> ((530) & 63))
    for fi7 := 0; fi7 < 13; fi7++ {
        v4 -= -2.25
        _ = fi7
    }
    switch "\n\t" {
    case "\n\t":
        if ((-32768 > 1000) && ((v2 << ((65535) & 63)) >= (v2 / (((v6 + v6)) | 1)))) {
            _ = len(v3)
            v4 += -2.25
            fmt.Println(v4, v4)
        }
    case "a":
        v4 = (0.0 + 1.5)
    default:
        v4 -= ((0.0 + 1.5) + v4)
    }
    var v8 uint16 = 65535
    fmt.Println(v6)
    v3 = append(v3, "")
    for fi9 := 0; fi9 < 5; fi9++ {
        for fi10 := 0; fi10 < 10; fi10++ {
            v5 += 2
            fmt.Println(v6, v2)
            _ = fi10
        }
        v11 := GenS1{F0: 3, F1: -1, F2: -108}
        _ = fi9
        _ = v11
    }
    v2 += (v2 / (((^(v2))) | 1))
    if len(v1) > 0 {
        v1[0] = v5
    }
    if false {
        var v12 bool = false
        var v13 bool = v12
        v5 = 1
        _ = v12
        _ = v13
    } else {
        for fi14 := 0; fi14 < 1; fi14++ {
            v15 := GenS1{F0: 127, F1: 9223372036854775807, F2: -653}
            fmt.Println(v6)
            switch (fi14 + fi14) {
            case -1:
                _ = len(v1)
            default:
                fmt.Println(v5, v2, v2)
            }
            _ = fi14
            _ = v15
        }
        var v16 int8 = -1
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = v16
    }
    var v17 int = -409
    v17 += (v17 | v17)
    if false {
        for fi18 := 0; fi18 < 14; fi18++ {
            var v19 uint64 = (^(v6))
            _ = fi18
            _ = v19
        }
    }
    v8 = v8
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
