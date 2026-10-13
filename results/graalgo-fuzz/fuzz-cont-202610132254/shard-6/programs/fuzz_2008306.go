package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 uint16 }
type GenS2 struct { F0 uint16; F1 string; F2 int }
func genFn0(p0 uint8, p1 uint8) uint16 {
    p0 += (^(p0))
    switch 127 {
    case 2147483647:
        for fi7 := 0; fi7 < 20; fi7++ {
            p0 = (p0 - p0)
            _ = fi7
        }
    case -980:
        var v8 uint = 18446744073709551615
        _ = v8
    case 127:
        var v9 int16 = 1
        _ = v9
    default:
        v10 := GenS1{F0: "foo bar", F1: 2, F2: 891}
        _ = v10
    }
    var v11 int = 2147483647
    _ = p0
    _ = p1
    _ = v11
    return 1
}

func main() {
    v1 := []string{("a" + ""), ("hello" + "")}
    var v2 uint = 238
    v3 := make(map[string]bool)
    var v4 float32 = ((0.0 + (float32(1e10) * float32(0.1))) + (float32((1.5 * 0.0)) * 0.0))
    var v5 int16 = -1
    v6 := make(map[string]int)
    fmt.Println(v2)
    if (127 != 42) {
        v2 = (^(v2))
        for fi12 := 0; fi12 < 2; fi12++ {
            v13 := GenS1{F0: "\n\t", F1: 1, F2: 957}
            v5 += (v5 * 0)
            v6["hello"] = (fi12 / ((fi12) | 1))
            _ = fi12
            _ = v13
        }
    } else {
        var v14 float32 = v4
        _ = v14
    }
    if (("" + "a") <= "") {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v15 uint64 = 2
        _ = v15
    } else {
        delete(v6, "foo bar")
        switch "" {
        case "a":
            v16 := GenS2{F0: 2, F1: "foo bar", F2: 32767}
            _ = v16
        }
    }
    v4 -= v4
    if true {
        v5 -= (v5 ^ -1)
        switch 42 {
        case 307:
            var v17 int8 = 42
            _ = v17
        case 128:
            v2 = v2
        case -746:
            for fi18 := 0; fi18 < 11; fi18++ {
                var v19 int = 127
                switch 128 {
                case 128:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case 127:
                    for fi20 := 0; fi20 < 8; fi20++ {
                        v1 = append(v1, "a")
                        v2 -= 255
                        switch "foo bar" {
                        case "foo bar":
                            var v21 uint8 = 124
                            _ = v21
                        case "hello":
                            v22 := GenS1{F0: ("foo bar" + "foo bar"), F1: 0, F2: genFn0(2, 145)}
                            _ = v22
                        case "":
                            delete(v3, "hello")
                        }
                        _ = fi20
                    }
                case 1:
                    v23 := GenS1{F0: ("" + ("a" + "\n\t")), F1: 709, F2: genFn0(1, 1)}
                    _ = v23
                default:
                    var v24 string = ""
                    _ = v24
                }
                _ = fi18
                _ = v19
            }
        }
    }
    switch 9223372036854775807 {
    case 4294967295:
        v5 -= 42
    }
    switch "\n\t" {
    case "hello":
        var v25 float64 = -2.25
        _ = v25
    }
    v26 := GenS2{F0: 559, F1: "foo bar", F2: 404}
    var v27 int8 = 75
    delete(v6, ("foo bar" + ("foo bar" + ("hello" + ""))))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
