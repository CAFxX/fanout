package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int; F1 uint8; F2 uint64 }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]bool)
    var v3 float32 = (float32(0.1) * 1.5)
    v4 := []string{"hello", "\n\t", ((("a" + "foo bar") + "\n\t") + (((("hello" + "a") + "a") + "a") + ""))}
    v5 := []string{("" + ""), "hello", "hello", ""}
    var v6 float64 = (3.14159 / 1.0)
    if true {
        v1[""] = 2
        switch "\n\t" {
        case "a":
            if (false && (false || (115 != 128))) {
                var v7 float64 = v6
                _ = v7
            }
        default:
            v8 := GenS1{F0: 1}
            _ = v8
        }
    }
    v3 += (v3 / (v3 * ((float32(0.1) * v3) * 1e10)))
    fmt.Println(v3, v6)
    fmt.Println(v6)
    var v9 bool = true
    _ = len(v5)
    var v10 int64 = -207
    for fi11 := 0; fi11 < 11; fi11++ {
        fmt.Println(v3, v3, fi11)
        _ = len(v2)
        _ = fi11
    }
    switch 9223372036854775807 {
    case -669:
        if ((v3 + 1.5) < v3) {
            delete(v2, 42)
            v10 += (^(v10))
            fmt.Println(v6, v6)
        }
    case -9223372036854775808:
        for fi12 := 0; fi12 < 5; fi12++ {
            v9 = (v9 && (!(42 == 42)))
            var v13 uint32 = 2147483648
            var v14 uint64 = 543
            _ = fi12
            _ = v13
            _ = v14
        }
    case -2147483648:
        v3 = float32(((1.5 - v6) - v6))
    }
    var v15 uint32 = 785
    for fi16 := 0; fi16 < 20; fi16++ {
        for fi17 := 0; fi17 < 20; fi17++ {
            fmt.Println(v10, v9)
            _ = fi17
        }
        _ = fi16
    }
    v18 := GenS1{F0: 254}
    v19 := GenS1{F0: 1}
    var v20 int64 = 9223372036854775807
    v20 = -9223372036854775808
    for fi21 := 0; fi21 < 10; fi21++ {
        v20 = (v20 ^ (v10 % ((v10) | 1)))
        _ = fi21
    }
    for fi22 := 0; fi22 < 16; fi22++ {
        switch "hello" {
        case "":
            var v23 int = 32768
            _ = v23
        case "a":
            v20 = v10
        }
        if len(v4) > 0 {
            v4[0] = "\n\t"
        }
        _ = fi22
    }
    switch 128 {
    case 127:
        for fi24 := 0; fi24 < 8; fi24++ {
            v25 := GenS2{F0: 32767, F1: 97, F2: 0}
            _ = fi24
            _ = v25
        }
    case 213:
        var v26 bool = v9
        _ = v26
    default:
        var v27 string = ("a" + "foo bar")
        _ = v27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
