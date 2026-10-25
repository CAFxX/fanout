package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 int16; F1 uint }
func genFn0() int32 {
    v7 := GenS1{F0: (1e10 >= (3.14159 / (-2.25 + (0.1 / 1e10)))), F1: 4294967295}
    var v8 float32 = 3.14159
    _ = v7
    _ = v8
    return 2147483647
}

func main() {
    v1 := []int{-2147483648, -9223372036854775808, 128}
    v2 := []uint32{42, 479, 42, 42}
    v3 := make(map[string]bool)
    var v4 float64 = (3.14159 * 1.5)
    var v5 int64 = 9223372036854775807
    var v6 int16 = 1
    for fi9 := 0; fi9 < 3; fi9++ {
        switch fi9 {
        case 32767:
            fmt.Println(v4)
        case 570:
            v5 = (-(v5))
        default:
            for fi10 := 0; fi10 < 14; fi10++ {
                fmt.Println(fi10, v6)
                if (!(true || true)) {
                    switch (("\n\t" + "foo bar") + "hello") {
                    case "foo bar":
                        v6 += (^(v6))
                    case "hello":
                        v6 = v6
                    }
                } else {
                    v6 = -32768
                }
                var v11 bool = ((v6 << ((2) & 63)) < v6)
                _ = fi10
                _ = v11
            }
        }
        fmt.Println(v4, v4)
        _ = fi9
    }
    v4 = 1.5
    v6 = v6
    for fi12 := 0; fi12 < 2; fi12++ {
        if (!false) {
            v5 += (-(v5))
            switch "a" {
            case "":
                fmt.Println(v5, v4, v4)
            }
        }
        _ = fi12
    }
    var v13 float64 = v4
    v14 := GenS2{F0: 1, F1: 65535}
    switch 1 {
    case -2147483648:
        _ = len(v1)
    case -359:
        var v15 int16 = (-(v6))
        _ = v15
    case 4294967295:
        switch "foo bar" {
        case "foo bar":
            v16 := GenS2{F0: (v6 + -32768), F1: 874}
            _ = v16
        case "hello":
            var v17 uint16 = 32767
            _ = v17
        }
    default:
        switch "\n\t" {
        case "":
            v5 -= v5
        case "hello":
            for fi18 := 0; fi18 < 2; fi18++ {
                switch (("\n\t" + "a") + "\n\t") {
                case "":
                    for fi19 := 0; fi19 < 16; fi19++ {
                        v1 = append(v1, -867)
                        var v20 bool = (genFn0() < -221)
                        v13 -= (1e10 + v4)
                        _ = fi19
                        _ = v20
                    }
                default:
                    v5 -= (v5 * 539)
                }
                _ = fi18
            }
        default:
            var v21 uint32 = 2147483648
            _ = v21
        }
    }
    _ = len(v3)
    var v22 int = 32768
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
