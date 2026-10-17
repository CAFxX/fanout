package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 int8) int16 {
    p0 -= p0
    v4 := []int{663}
    _ = len(v4)
    _ = p0
    _ = v4
    return 42
}
func genFn1(p0 int16) int16 {
    var v5 int = 42
    fmt.Println(v5, p0)
    for fi6 := 0; fi6 < 2; fi6++ {
        switch "hello" {
        case "a":
            fmt.Println(p0, v5)
        }
        if (1e10 <= 1.5) {
            v7 := GenS1{F0: (^(p0)), F1: 1}
            v8 := GenS2{F0: 0}
            v9 := GenS2{F0: 296}
            _ = v7
            _ = v8
            _ = v9
        }
        _ = fi6
    }
    v10 := GenS1{F0: (p0 << ((91) & 63)), F1: 108}
    switch v5 {
    case 32768:
        switch "hello" {
        case "foo bar":
            for fi11 := 0; fi11 < 6; fi11++ {
                p0 = p0
                _ = fi11
            }
        }
    default:
        v12 := []int{v5}
        _ = v12
    }
    _ = p0
    _ = v10
    _ = v5
    return (^(p0))
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint8 = 199
    v3 := []string{"\n\t"}
    var v13 float32 = 3.14159
    v14 := GenS1{F0: -32768, F1: 42}
    switch "" {
    case "foo bar":
        fmt.Println(v13)
    case "a":
        fmt.Println(v13, v2)
    default:
        v13 += v13
    }
    v15 := GenS1{F0: genFn1(genFn0(1)), F1: 1}
    v13 = float32(1.5)
    var v16 uint = 42
    for fi17 := 0; fi17 < 2; fi17++ {
        for fi18 := 0; fi18 < 9; fi18++ {
            v16 = 4294967295
            _ = fi18
        }
        for fi19 := 0; fi19 < 14; fi19++ {
            v13 += (v13 - v13)
            var v20 int16 = 313
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi19
            _ = v20
        }
        fmt.Println(v13, v13, fi17)
        _ = fi17
    }
    var v21 int = 1
    v22 := GenS1{F0: 0, F1: 127}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v3
}
