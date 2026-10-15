package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 string) uint {
    if (!false) {
        p0 = p0
    }
    var v6 uint = 42
    _ = p0
    _ = v6
    return v6
}

func main() {
    var v1 float64 = ((3.14159 - 0.1) - -2.25)
    var v2 int64 = -664
    var v3 bool = false
    v4 := make(map[int]string)
    v5 := []string{"\n\t", "a"}
    v3 = v3
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v7 float32 = ((-2.25 * 3.14159) - (1e10 / float32(3.14159)))
    if (2147483647 == 824) {
        var v8 bool = v3
        for fi9 := 0; fi9 < 4; fi9++ {
            v7 += (((v7 + (v7 / float32(v1))) / 1.0) + v7)
            var v10 int64 = 743
            v7 += float32(float64(float32(((1e10 - 3.14159) - (v1 + float64(1e10))))))
            _ = fi9
            _ = v10
        }
        switch 4294967295 {
        case -1:
            var v11 float32 = 1e10
            _ = v11
        case 127:
            fmt.Println(v8, v7)
        case 4294967295:
            for fi12 := 0; fi12 < 4; fi12++ {
                v13 := GenS1{F0: -128}
                _ = fi12
                _ = v13
            }
        }
        _ = v8
    }
    if ((0 == 123) && true) {
        var v14 float64 = float64(0.1)
        var v15 float64 = v1
        _ = v14
        _ = v15
    }
    fmt.Println(v7, v2, v2)
    if (v3 && v3) {
        var v16 int64 = -9223372036854775808
        v7 -= 1.5
        v5 = append(v5, "a")
        _ = v16
    }
    var v17 uint16 = 32767
    fmt.Println(v1, v3)
    v18 := GenS1{F0: -104}
    if true {
        v7 = v7
        var v19 int16 = 315
        _ = v19
    }
    v7 = (v7 / float32(v1))
    for fi20 := 0; fi20 < 16; fi20++ {
        v1 = v1
        _ = fi20
    }
    var v21 uint64 = 18446744073709551615
    var v22 uint8 = 42
    switch ((((("" + "foo bar") + "\n\t") + "hello") + "") + (("\n\t" + ("\n\t" + ("" + ("a" + ("\n\t" + ""))))) + "hello")) {
    case "hello":
        for fi23 := 0; fi23 < 16; fi23++ {
            fmt.Println(v17, v2, v1)
            _ = fi23
        }
    default:
        fmt.Println(v17, v3, v17)
    }
    v3 = true
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
