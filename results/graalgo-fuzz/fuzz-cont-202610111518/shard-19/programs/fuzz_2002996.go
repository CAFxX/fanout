package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }

func main() {
    v1 := []int{-2147483648, 510, -1, 32767}
    var v2 int = -1
    var v3 uint = 0
    v2 += (v2 | (^(v2)))
    v2 = v2
    var v4 uint64 = 0
    v3 += 18446744073709551615
    switch "foo bar" {
    case "\n\t":
        fmt.Println(v3, v4)
    case "hello":
        if (1 <= 754) {
            var v5 int16 = 32767
            _ = v5
        }
    }
    for fi6 := 0; fi6 < 10; fi6++ {
        for fi7 := 0; fi7 < 3; fi7++ {
            v4 += (v4 / (((v4 + 731)) | 1))
            switch fi7 {
            case 9223372036854775807:
                var v8 int32 = 1
                _ = v8
            case 836:
                v2 -= fi6
            case -744:
                v9 := GenS1{F0: 0, F1: 234}
                _ = v9
            }
            _ = fi7
        }
        _ = fi6
    }
    for fi10 := 0; fi10 < 16; fi10++ {
        for fi11 := 0; fi11 < 3; fi11++ {
            if false {
                var v12 string = (("foo bar" + "foo bar") + "hello")
                _ = v12
            }
            v3 -= v3
            _ = fi11
        }
        var v13 int8 = -1
        v1 = append(v1, 324)
        _ = fi10
        _ = v13
    }
    var v14 float64 = (float64(1.5) * 1e10)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
