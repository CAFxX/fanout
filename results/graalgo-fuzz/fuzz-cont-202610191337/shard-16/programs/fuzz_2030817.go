package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 uint8; F1 string; F2 int8 }
type GenS3 struct { F0 string; F1 int8 }
type GenS4 struct { F0 int16; F1 int }
func genFn0() uint {
    for fi4 := 0; fi4 < 13; fi4++ {
        var v5 bool = true
        if (true && v5) {
            fmt.Println(fi4, fi4)
            switch fi4 {
            case 4294967295:
                v6 := []int{128, 32768, (fi4 / ((fi4) | 1))}
                _ = v6
            case 32767:
                var v7 int8 = 1
                _ = v7
            }
        }
        _ = fi4
        _ = v5
    }
    for fi8 := 0; fi8 < 7; fi8++ {
        for fi9 := 0; fi9 < 6; fi9++ {
            var v10 uint8 = 42
            var v11 uint16 = 255
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = fi8
    }
    for fi12 := 0; fi12 < 16; fi12++ {
        v13 := GenS1{F0: 2147483647, F1: 503}
        _ = fi12
        _ = v13
    }
    v14 := []int{32768, 0, 128, -2147483648}
    _ = v14
    return 73
}
func genFn1(p0 bool, p1 uint, p2 int8) int8 {
    var v15 uint64 = 2
    var v16 bool = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v16
    return 0
}

func main() {
    var v1 string = "\n\t"
    v2 := []string{}
    var v3 uint32 = 618
    fmt.Println(v3, v3)
    v17 := GenS2{F0: 2, F1: (v1 + v1), F2: 0}
    for fi18 := 0; fi18 < 7; fi18++ {
        v3 = 4294967295
        var v19 float32 = 1.5
        var v20 int8 = 100
        _ = fi18
        _ = v19
        _ = v20
    }
    var v21 uint = 42
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if ((0.1 - (0.0 - float64(1e10))) >= (float64(3.14159) * float64(-2.25))) {
        v22 := GenS3{F0: ("foo bar" + ("hello" + v1)), F1: -128}
        v2 = append(v2, "\n\t")
        v21 = v21
        _ = v22
    } else {
        var v23 int32 = 0
        _ = v23
    }
    var v24 int = -1
    if (((0.1 + 0.0) + float64(0.0)) <= 3.14159) {
        v25 := GenS4{F0: 42, F1: (v24 + v24)}
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v24
    _ = v3
}
