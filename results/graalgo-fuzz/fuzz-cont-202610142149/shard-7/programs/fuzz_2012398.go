package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint8; F1 bool; F2 int16 }
func genFn0(p0 int16, p1 int8, p2 uint) int8 {
    var v7 int = 643
    var v8 uint8 = 255
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return -6
}

func main() {
    var v1 uint = 18446744073709551615
    var v2 int16 = -78
    v3 := make(map[uint32]string)
    v4 := []string{"\n\t", "a", "foo bar"}
    var v5 string = "\n\t"
    v6 := []uint32{362, 42, 65535}
    if (0.0 != (1e10 * (-2.25 - -2.25))) {
        v5 = "hello"
        v9 := GenS1{F0: (42 == -2147483648), F1: 1}
        var v10 int = 2147483647
        _ = v10
        _ = v9
    } else {
        switch "\n\t" {
        case "a":
            for fi11 := 0; fi11 < 17; fi11++ {
                var v12 string = ((v5 + (("" + v5) + ((v5 + v5) + "a"))) + "\n\t")
                var v13 uint32 = 42
                _ = fi11
                _ = v12
                _ = v13
            }
        case "foo bar":
            v2 = v2
        }
        var v14 uint = v1
        _ = v14
    }
    var v15 uint32 = 0
    switch -2147483648 {
    case 127:
        for fi16 := 0; fi16 < 4; fi16++ {
            v17 := GenS2{F0: (^(v1))}
            _ = fi16
            _ = v17
        }
    }
    var v18 int = 767
    var v19 int16 = v2
    v15 -= v15
    v5 = (v5 + "")
    var v20 uint32 = v15
    var v21 int = -9223372036854775808
    if (199 < 2) {
        _ = len(v3)
        var v22 int32 = -2147483648
        switch "foo bar" {
        case "a":
            v23 := GenS3{F0: 0, F1: (310 < 2), F2: (v19 * 116)}
            _ = v23
        case "\n\t":
            v3[(^(v15))] = v5
        case "":
            v24 := v3[v15]
            _ = v24
        }
        _ = v22
    } else {
        fmt.Println(v2)
        var v25 float64 = 0.1
        var v26 float32 = float32(3.14159)
        _ = v25
        _ = v26
    }
    v20 = v15
    delete(v3, v15)
    for fi27 := 0; fi27 < 17; fi27++ {
        switch v5 {
        case "a":
            var v28 bool = (42 > 255)
            _ = v28
        case "":
            fmt.Println(v21, v19, v1)
        }
        for fi29 := 0; fi29 < 12; fi29++ {
            v30 := GenS2{F0: v1}
            _ = fi29
            _ = v30
        }
        _ = fi27
    }
    v21 -= v21
    _ = len(v3)
    v21 += 60
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
