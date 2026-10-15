package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8; F1 bool }

func main() {
    var v1 float32 = 0.0
    v2 := make(map[string]string)
    var v3 int64 = -9223372036854775808
    var v4 string = ""
    v1 = (1e10 * (v1 / (v1 + v1)))
    v5 := GenS1{F0: (127 <= 0)}
    v6 := []int{1, 127, 128}
    for fi7 := 0; fi7 < 2; fi7++ {
        if (false || true) {
            v8 := GenS1{F0: false}
            var v9 int32 = 2147483647
            _ = v8
            _ = v9
        } else {
            var v10 uint64 = 4294967295
            if len(v6) > 0 {
                v6[0] = fi7
            }
            _ = v10
        }
        v1 += float32(-2.25)
        _ = fi7
    }
    switch (v4 + v4) {
    case "foo bar":
        v11 := v2["hello"]
        _ = v11
    case "":
        var v12 uint16 = 47
        _ = v12
    default:
        var v13 uint32 = 698
        _ = v13
    }
    if (0.1 == 0.0) {
        v14 := GenS1{F0: ((3.14159 >= 0.1) && (398 != -2147483648))}
        _ = v14
    }
    for fi15 := 0; fi15 < 18; fi15++ {
        v16 := GenS2{F0: 2, F1: (float64((v1 * v1)) > -2.25)}
        _ = fi15
        _ = v16
    }
    v17 := GenS2{F0: 2, F1: (42 == 4294967295)}
    switch -2147483648 {
    case -2147483648:
        fmt.Println(v3, v1)
    case 42:
        switch (v4 + v4) {
        case "\n\t":
            if len(v6) > 0 {
                v6[0] = -933
            }
        case "foo bar":
            var v18 float32 = v1
            _ = v18
        default:
            v3 += -70
        }
    }
    var v19 int = 32768
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
