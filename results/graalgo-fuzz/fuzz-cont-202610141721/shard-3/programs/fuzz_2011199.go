package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16; F1 int }

func main() {
    var v1 bool = (127 == -2147483648)
    var v2 uint64 = 2
    v3 := make(map[int]uint32)
    var v4 int = 42
    v5 := make(map[string]uint32)
    var v6 uint64 = v2
    if (0 == -574) {
        var v7 float64 = -2.25
        v2 += (^(v6))
        _ = v7
    } else {
        v2 -= v6
        v8 := v3[v4]
        _ = v8
    }
    _ = len(v5)
    if ((3.14159 / 3.14159) >= 3.14159) {
        var v9 uint64 = (v6 | (v6 >> ((4294967295) & 63)))
        if true {
            switch ("hello" + "a") {
            case "":
                var v10 uint8 = 93
                _ = v10
            case "a":
                var v11 uint32 = 704
                _ = v11
            }
        } else {
            v12 := GenS1{F0: 0, F1: 429}
            var v13 int32 = 2147483647
            _ = v12
            _ = v13
        }
        if (!false) {
            v14 := GenS1{F0: 0, F1: 954}
            fmt.Println(v6, v1, v6)
            v15 := GenS1{F0: 105, F1: 42}
            _ = v14
            _ = v15
        }
        _ = v9
    } else {
        v4 = v4
        var v16 int = v4
        _ = v16
    }
    switch 9223372036854775807 {
    case 997:
        var v17 uint16 = 298
        _ = v17
    case -2147483648:
        for fi18 := 0; fi18 < 13; fi18++ {
            v19 := []int{-362, fi18, v4}
            v19 = append(v19, 32767)
            _ = fi18
            _ = v19
        }
    default:
        for fi20 := 0; fi20 < 10; fi20++ {
            fmt.Println(v4, v1, v1)
            v6 += (v2 & v6)
            for fi21 := 0; fi21 < 14; fi21++ {
                v22 := GenS2{F0: 1}
                v1 = v1
                _ = fi21
                _ = v22
            }
            _ = fi20
        }
    }
    switch ("\n\t" + (("a" + "a") + ("\n\t" + "a"))) {
    case "a":
        fmt.Println(v2)
    }
    fmt.Println(v1)
    fmt.Println(v2, v6, v1)
    v23 := GenS3{F0: 419, F1: -9223372036854775808}
    for fi24 := 0; fi24 < 16; fi24++ {
        var v25 float32 = (0.0 + 1e10)
        _ = fi24
        _ = v25
    }
    switch 42 {
    case 2147483647:
        delete(v5, "")
    }
    v2 -= (v2 + 1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
