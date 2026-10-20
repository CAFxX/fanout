package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 string; F1 int16 }
func genFn0(p0 int, p1 string) int8 {
    var v5 uint32 = 4294967295
    v6 := GenS1{F0: true, F1: (p0 | (p0 | 127)), F2: 0}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 72
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 uint64 = v1
    v3 := []string{(("a" + "") + ("foo bar" + "foo bar"))}
    var v4 float64 = 0.1
    for fi7 := 0; fi7 < 5; fi7++ {
        v8 := GenS1{F0: (!true), F1: (^(fi7)), F2: 1}
        for fi9 := 0; fi9 < 18; fi9++ {
            v1 -= 479
            fmt.Println(v1)
            if ((1.5 + (3.14159 + 0.0)) != 0.0) {
                switch ("a" + "foo bar") {
                case "":
                    v10 := GenS2{F0: -1}
                    _ = v10
                default:
                    switch (-(fi9)) {
                    case 465:
                        v2 += (v1 - v1)
                    }
                }
            } else {
                v11 := GenS3{F0: "", F1: 1000}
                _ = v11
            }
            v4 = v4
            fmt.Println(fi9, fi7)
            _ = fi9
        }
        var v12 int = 2147483647
        _ = fi7
        _ = v12
        _ = v8
    }
    v4 = 1e10
    fmt.Println(v2)
    if len(v3) > 0 {
        v3[0] = "hello"
    }
    fmt.Println(v2)
    fmt.Println(v1, v2, v4)
    v13 := GenS2{F0: -2147483648}
    switch 0 {
    case 42:
        v4 = -2.25
    default:
        v1 -= 2
    }
    var v14 uint8 = 127
    for fi15 := 0; fi15 < 14; fi15++ {
        v3 = append(v3, "\n\t")
        v2 += 1
        _ = fi15
    }
    fmt.Println(v2, v1)
    var v16 float32 = float32(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
