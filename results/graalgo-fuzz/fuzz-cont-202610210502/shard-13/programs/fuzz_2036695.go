package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32; F1 string; F2 uint }
func genFn0(p0 int32, p1 int, p2 uint32) int16 {
    v7 := GenS1{F0: 1}
    p0 += p0
    v8 := GenS1{F0: 0}
    fmt.Println(p2, p2)
    v9 := GenS1{F0: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 1
}
func genFn1(p0 uint8, p1 int) string {
    var v10 int32 = 1
    v11 := []int{}
    switch "foo bar" {
    case "foo bar":
        v12 := GenS1{F0: 1}
        _ = v12
    case "\n\t":
        if ((0.1 / (float32((1.5 - float64(0.1))) / 3.14159)) <= (0.1 * ((3.14159 / 1.5) * float32((1.5 + (1e10 * ((-2.25 + 1e10) * -2.25))))))) {
            var v13 uint = 18446744073709551615
            v13 = (v13 & 1)
            _ = v13
        }
    }
    if len(v11) > 0 {
        v11[0] = p1
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return "\n\t"
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint32 = 42
    var v3 int32 = 2147483647
    var v4 float32 = 1.5
    v5 := []string{"foo bar", ("foo bar" + ("foo bar" + ("\n\t" + ("hello" + "a"))))}
    var v6 string = ("" + "foo bar")
    v6 = ("" + (((("hello" + "hello") + "a") + ((((v6 + v6) + v6) + v6) + "\n\t")) + v6))
    _ = len(v1)
    v14 := GenS2{F0: 0, F1: (v6 + v6), F2: 0}
    var v15 float64 = (((3.14159 * (3.14159 + (1.5 / (((1e10 * ((3.14159 * 0.0) + 1.5)) - 1e10) + float64(v4))))) + 1e10) * 1.5)
    var v16 int64 = 9223372036854775807
    v2 = v2
    var v17 float32 = v4
    for fi18 := 0; fi18 < 17; fi18++ {
        var v19 string = ("\n\t" + (v6 + (v6 + ("a" + (v6 + (("a" + v6) + (v6 + v6)))))))
        _ = fi18
        _ = v19
    }
    switch "foo bar" {
    case "foo bar":
        v3 = v3
    case "hello":
        fmt.Println(v2)
    case "":
        _ = len(v5)
    default:
        for fi20 := 0; fi20 < 2; fi20++ {
            v21 := GenS2{F0: v3, F1: v6, F2: 32}
            _ = fi20
            _ = v21
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
