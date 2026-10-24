package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 int16 }
type GenS2 struct { F0 int64; F1 uint64 }
type GenS3 struct { F0 uint64; F1 uint; F2 int16 }

func main() {
    var v1 bool = true
    v2 := make(map[string]bool)
    var v3 int32 = 1
    v4 := GenS1{F0: -146, F1: v1, F2: 0}
    var v5 uint32 = 491
    v6 := GenS1{F0: 620, F1: (v5 > v5), F2: -170}
    fmt.Println(v5, v5)
    var v7 float64 = float64(0.0)
    v8 := []int{308, -2147483648}
    _ = len(v8)
    v7 = (((v7 * ((v7 - -2.25) - (0.0 * 1.5))) - v7) / (v7 - 3.14159))
    v1 = ((v7 / 1e10) > 3.14159)
    var v9 uint32 = (v5 << ((0) & 63))
    v10 := GenS2{F0: 1, F1: 4294967295}
    switch 32767 {
    case 9223372036854775807:
        _ = len(v8)
    case 0:
        for fi11 := 0; fi11 < 18; fi11++ {
            if (!v1) {
                v12 := v2[("hello" + ("a" + "\n\t"))]
                _ = v12
            } else {
                for fi13 := 0; fi13 < 6; fi13++ {
                    if v1 {
                        v14 := v2[(("a" + "hello") + "foo bar")]
                        v1 = ((v3 << ((314) & 63)) < (v3 >> ((42) & 63)))
                        v2["a"] = v1
                        fmt.Println(v14, fi13, fi11)
                        _ = v14
                    }
                    _ = fi13
                }
            }
            v15 := GenS3{F0: 0, F1: 1, F2: 1000}
            if (2 <= 86) {
                var v16 float32 = (0.1 * 0.1)
                fmt.Println(v16, fi11, v3)
                var v17 string = "a"
                _ = v16
                _ = v17
            }
            _ = fi11
            _ = v15
        }
    }
    for fi18 := 0; fi18 < 15; fi18++ {
        switch "foo bar" {
        case "":
            v3 = (^(v3))
        case "foo bar":
            for fi19 := 0; fi19 < 1; fi19++ {
                var v20 float64 = 0.0
                _ = len(v8)
                var v21 int8 = -119
                _ = fi19
                _ = v20
                _ = v21
            }
        default:
            switch (fi18 + 411) {
            case 999:
                v8 = append(v8, -9223372036854775808)
            case -9223372036854775808:
                v8 = append(v8, (fi18 * fi18))
            }
        }
        _ = fi18
    }
    switch 128 {
    case -495:
        fmt.Println(v9)
    case -9:
        v8 = append(v8, -1)
    case 758:
        var v22 bool = v1
        _ = v22
    default:
        v3 = v3
    }
    var v23 float32 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
