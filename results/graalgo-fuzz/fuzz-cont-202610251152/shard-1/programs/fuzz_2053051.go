package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 bool; F1 string }
type GenS3 struct { F0 uint; F1 int64; F2 int64 }
func genFn0(p0 string, p1 uint16, p2 uint64) string {
    var v7 int32 = -42
    var v8 int64 = -9223372036854775808
    fmt.Println(p1, v8, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return "foo bar"
}

func main() {
    v1 := []string{}
    var v2 int64 = 42
    var v3 int8 = 42
    v4 := make(map[int]int)
    v5 := []string{(("\n\t" + "hello") + "\n\t"), "a", (("hello" + ("hello" + "a")) + "\n\t")}
    var v6 int = 9223372036854775807
    for fi9 := 0; fi9 < 12; fi9++ {
        if true {
            v3 -= (v3 % ((v3) | 1))
        } else {
            v2 = v2
            for fi10 := 0; fi10 < 20; fi10++ {
                v6 = (^(v6))
                switch fi10 {
                case 32767:
                    v11 := GenS1{F0: (fi9 / ((fi10) | 1)), F1: 798}
                    _ = v11
                case 1:
                    v12 := v4[v6]
                    v3 = 1
                    v13 := GenS2{F0: false, F1: ("\n\t" + "foo bar")}
                    var v14 bool = true
                    _ = v12
                    _ = v13
                    _ = v14
                case 571:
                    if (false && (!((((((((((-2.25 + (1e10 / 1e10)) + 3.14159) / 1.5) - 3.14159) / float32(1e10)) - 0.0) * float32(3.14159)) > (0.1 - 1e10)) && false) || false))) {
                        v15 := GenS3{F0: 1, F1: -67, F2: v2}
                        fmt.Println(v3, fi9, v2)
                        if len(v1) > 0 {
                            v1[0] = ("a" + "foo bar")
                        }
                        _ = v15
                    }
                }
                _ = fi10
            }
        }
        v2 += v2
        _ = fi9
    }
    v4[880] = (^(v6))
    if false {
        v16 := GenS1{F0: (-(v6)), F1: 0}
        v17 := v4[v6]
        var v18 uint16 = 255
        _ = v16
        _ = v17
        _ = v18
    } else {
        v2 = 1
        var v19 int8 = 42
        var v20 float32 = float32(0.0)
        _ = v19
        _ = v20
    }
    fmt.Println(v3, v3)
    switch (("hello" + ("" + ("hello" + ("foo bar" + "a")))) + "\n\t") {
    case "foo bar":
        _ = len(v5)
    default:
        v6 -= v6
    }
    v6 += v6
    v3 += v3
    fmt.Println(v6)
    v3 = -128
    if true {
        v21 := GenS2{F0: true, F1: "\n\t"}
        for fi22 := 0; fi22 < 9; fi22++ {
            v2 = -9223372036854775808
            _ = fi22
        }
        for fi23 := 0; fi23 < 1; fi23++ {
            var v24 uint16 = 2
            _ = fi23
            _ = v24
        }
        _ = v21
    } else {
        fmt.Println(v2, v6, v6)
    }
    v25 := GenS3{F0: 1, F1: v2, F2: (v2 >> ((1) & 63))}
    v5 = append(v5, "foo bar")
    v26 := GenS3{F0: 488, F1: (-(v2)), F2: (v2 | v2)}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
