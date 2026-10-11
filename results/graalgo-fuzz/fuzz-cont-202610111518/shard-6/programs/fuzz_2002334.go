package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 int32 }
func genFn0(p0 int32, p1 uint64) uint64 {
    v4 := GenS1{F0: 42, F1: ("a" + "hello"), F2: p0}
    var v5 int = 1
    for fi6 := 0; fi6 < 18; fi6++ {
        if (622 <= 65535) {
            v7 := GenS1{F0: p1, F1: "foo bar", F2: (p0 * (p0 % (((p0 % ((p0) | 1))) | 1)))}
            v5 += (fi6 & v5)
            _ = v7
        } else {
            var v8 string = ""
            for fi9 := 0; fi9 < 18; fi9++ {
                var v10 float64 = 3.14159
                _ = fi9
                _ = v10
            }
            _ = v8
        }
        var v11 string = ("" + "\n\t")
        _ = fi6
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return (p1 * 4294967295)
}

func main() {
    var v1 uint64 = 0
    var v2 bool = (false && true)
    var v3 float32 = 0.1
    switch "a" {
    case "hello":
        if (0 != 0) {
            v12 := []int{}
            if true {
                v3 -= v3
                var v13 bool = (false && (v2 || v2))
                _ = v13
            }
            var v14 uint64 = 42
            _ = v12
            _ = v14
        }
    default:
        switch ("foo bar" + "hello") {
        case "foo bar":
            v2 = (float32(0.0) == (((float32(-2.25) + -2.25) + float32((float64(1.5) - -2.25))) * (float32((((1.5 * ((3.14159 / (1e10 * 0.1)) / 1.0)) / 1e10) / float64((v3 - (1.5 - 0.0))))) / 1.0)))
        case "a":
            v15 := GenS1{F0: (v1 / (((v1 ^ v1)) | 1)), F1: "a", F2: 42}
            _ = v15
        default:
            v2 = (-128 != 42)
        }
    }
    v3 = v3
    var v16 int8 = -71
    v2 = (4294967295 > 71)
    v16 = v16
    fmt.Println(v16, v16, v1)
    v3 = (0.1 - float32((3.14159 * -2.25)))
    fmt.Println(v16)
    var v17 int64 = 1
    if (v2 && v2) {
        fmt.Println(v3, v16)
    } else {
        var v18 uint = 584
        v19 := []int{}
        v20 := GenS1{F0: v1, F1: "foo bar", F2: -603}
        _ = v18
        _ = v19
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
