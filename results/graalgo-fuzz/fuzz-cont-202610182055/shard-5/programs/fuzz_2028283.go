package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 uint16; F1 uint; F2 int32 }
type GenS3 struct { F0 bool; F1 uint16 }
func genFn0(p0 uint64) int {
    if false {
        p0 -= p0
    }
    fmt.Println(p0)
    for fi7 := 0; fi7 < 11; fi7++ {
        fmt.Println(p0)
        _ = fi7
    }
    var v8 int64 = 0
    _ = p0
    _ = v8
    return -47
}

func main() {
    v1 := []string{}
    v2 := make(map[int]string)
    v3 := make(map[int]uint32)
    var v4 bool = (false && (-1 == 32768))
    var v5 float64 = 1.5
    var v6 int = 127
    fmt.Println(v4)
    for fi9 := 0; fi9 < 18; fi9++ {
        for fi10 := 0; fi10 < 18; fi10++ {
            v11 := GenS1{F0: -32768, F1: 42}
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    v4 = v4
    v12 := GenS2{F0: 255, F1: 4294967295, F2: -2147483648}
    fmt.Println(v6, v6, v6)
    for fi13 := 0; fi13 < 17; fi13++ {
        v5 += v5
        var v14 int64 = 42
        delete(v3, (fi13 | v6))
        _ = fi13
        _ = v14
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v1 = append(v1, ("a" + "a"))
    if (v4 || v4) {
        v15 := GenS2{F0: 42, F1: 4294967295, F2: -2147483648}
        v6 = 42
        v16 := GenS3{F0: true, F1: 0}
        _ = v15
        _ = v16
    }
    var v17 string = (("foo bar" + "hello") + ("a" + "foo bar"))
    switch (("a" + "hello") + v17) {
    case "a":
        var v18 uint32 = 0
        _ = v18
    case "\n\t":
        var v19 int8 = 0
        _ = v19
    default:
        switch (v6 & (v6 * (^(v6)))) {
        case 0:
            switch v6 {
            case 32768:
                v6 += 134
            case -2147483648:
                _ = len(v2)
            default:
                var v20 float32 = 0.0
                _ = v20
            }
        case -760:
            v5 -= v5
        }
    }
    fmt.Println(v6, v4, v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
