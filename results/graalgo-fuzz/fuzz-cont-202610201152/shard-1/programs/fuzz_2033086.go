package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 uint8 }
func genFn0(p0 uint32, p1 int8) int64 {
    var v6 string = ""
    p0 = 108
    _ = p0
    _ = p1
    _ = v6
    return 1
}

func main() {
    v1 := []string{("" + "hello"), ("hello" + ""), ("\n\t" + "")}
    v2 := []string{""}
    var v3 uint = 18446744073709551615
    var v4 uint16 = 32767
    v5 := []uint32{}
    var v7 float32 = 3.14159
    v7 = (((v7 * v7) * v7) * v7)
    v3 = (v3 % ((v3) | 1))
    v3 += 0
    switch "\n\t" {
    case "hello":
        v7 = (float32(float64(v7)) - (((3.14159 / (1e10 + -2.25)) - (v7 - (float32(float64(v7)) + v7))) / ((((3.14159 - 3.14159) * (v7 - v7)) * 1.5) + (-2.25 - -2.25))))
    case "":
        switch 127 {
        case 42:
            fmt.Println(v3, v7)
        case 128:
            var v8 int64 = 1
            _ = v8
        }
    default:
        v7 = 1e10
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if (!(!("foo bar" >= ("foo bar" + (((("hello" + "hello") + ((("hello" + "hello") + "hello") + "hello")) + "\n\t") + "a"))))) {
        var v9 float32 = float32(0.1)
        var v10 int16 = 0
        var v11 int32 = 2147483647
        _ = v10
        _ = v11
        _ = v9
    }
    v2 = append(v2, "\n\t")
    for fi12 := 0; fi12 < 7; fi12++ {
        v7 -= (v7 * (3.14159 + 1.5))
        _ = fi12
    }
    v13 := GenS1{F0: (v4 - v4), F1: 198, F2: 16}
    var v14 int16 = -1
    v3 += 763
    var v15 int = 32768
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
