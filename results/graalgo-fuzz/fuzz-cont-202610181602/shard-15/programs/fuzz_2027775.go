package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 int }
type GenS2 struct { F0 bool; F1 uint8; F2 int8 }
type GenS3 struct { F0 bool; F1 int }

func main() {
    v1 := make(map[string]uint32)
    var v2 int64 = -915
    v3 := []string{}
    v4 := make(map[int]int)
    var v5 uint8 = 2
    var v6 string = "foo bar"
    _ = len(v3)
    var v7 int32 = -394
    v8 := GenS1{F0: v6, F1: v5, F2: 1}
    for fi9 := 0; fi9 < 13; fi9++ {
        v2 -= (v2 >> ((18446744073709551615) & 63))
        _ = fi9
    }
    v10 := v1[((v6 + (v6 + v6)) + (("foo bar" + ("foo bar" + (((v6 + "hello") + v6) + v6))) + v6))]
    v2 += 0
    v10 = v10
    v11 := GenS1{F0: "foo bar", F1: 1, F2: 4294967295}
    v12 := GenS1{F0: v6, F1: (v5 + v5), F2: 127}
    switch 4294967295 {
    case -2147483648:
        v7 = 967
    case 2147483647:
        switch -143 {
        case 1:
            var v13 bool = (false || true)
            _ = v13
        case 127:
            switch 9223372036854775807 {
            case 71:
                v14 := GenS2{F0: (0 == 1), F1: v5, F2: 42}
                _ = v14
            case 4294967295:
                if (v2 == (v2 >> ((255) & 63))) {
                    v15 := GenS1{F0: v6, F1: v5, F2: 32768}
                    v16 := GenS3{F0: false, F1: 1}
                    var v17 uint = 255
                    _ = v15
                    _ = v16
                    _ = v17
                }
            }
        case 730:
            v18 := GenS3{F0: (false || false), F1: 9223372036854775807}
            _ = v18
        }
    case -533:
        var v19 float64 = 1.5
        _ = v19
    default:
        v2 = (^(v2))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
