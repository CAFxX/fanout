package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16; F1 string }
type GenS4 struct { F0 int32; F1 int64 }
type GenS5 struct { F0 int8 }
func genFn0(p0 int8, p1 int32, p2 int) int {
    for fi6 := 0; fi6 < 3; fi6++ {
        v7 := GenS1{F0: p2, F1: 1, F2: -273}
        var v8 uint64 = 1
        v9 := GenS1{F0: fi6, F1: 890, F2: 683}
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    fmt.Println(p0)
    if false {
        v10 := GenS2{F0: ("foo bar" + (("foo bar" + ("\n\t" + ("hello" + "\n\t"))) + ""))}
        _ = v10
    } else {
        if (1000 <= 1000) {
            switch ("" + ((("hello" + "foo bar") + ("a" + "foo bar")) + ("" + "foo bar"))) {
            case "hello":
                var v11 int = p2
                _ = v11
            case "foo bar":
                p1 -= p1
            default:
                v12 := []int{}
                _ = v12
            }
        }
        for fi13 := 0; fi13 < 11; fi13++ {
            switch fi13 {
            case 0:
                v14 := GenS3{F0: 2, F1: "hello"}
                _ = v14
            }
            _ = fi13
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]string)
    var v3 int16 = 1000
    var v4 uint8 = 2
    var v5 float64 = ((1e10 + 3.14159) + (3.14159 * ((1e10 * 0.1) + -2.25)))
    v3 += v3
    if (-1 >= 1) {
        if false {
            fmt.Println(v3, v4)
            var v15 string = "\n\t"
            _ = v15
        }
        fmt.Println(v3)
    }
    fmt.Println(v5)
    if (!true) {
        v16 := GenS4{F0: 2147483647, F1: 0}
        var v17 float32 = 0.1
        fmt.Println(v17, v4)
        _ = v16
        _ = v17
    }
    switch -873 {
    case 426:
        var v18 int16 = (v3 / (((^(v3))) | 1))
        _ = v18
    default:
        v19 := GenS5{F0: 127}
        _ = v19
    }
    var v20 string = "hello"
    if false {
        var v21 float32 = 0.1
        fmt.Println(v3, v3)
        _ = v21
    }
    var v22 uint64 = 2
    v1[v20] = 119
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
