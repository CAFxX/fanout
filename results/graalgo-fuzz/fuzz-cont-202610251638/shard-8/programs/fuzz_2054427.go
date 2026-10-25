package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8; F1 int; F2 int32 }
func genFn0(p0 uint8, p1 string, p2 int16) uint32 {
    v6 := []int{664, 2147483647, -245, 588}
    switch p1 {
    case "foo bar":
        p2 = 42
    case "":
        for fi7 := 0; fi7 < 18; fi7++ {
            v8 := GenS1{F0: (fi7 | (fi7 / ((fi7) | 1)))}
            v6 = append(v6, (fi7 ^ -382))
            _ = fi7
            _ = v8
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 958
}
func genFn1(p0 string) int32 {
    for fi9 := 0; fi9 < 17; fi9++ {
        v10 := GenS2{F0: 0}
        v11 := GenS1{F0: fi9}
        _ = fi9
        _ = v10
        _ = v11
    }
    v12 := GenS2{F0: 1}
    fmt.Println(p0)
    p0 = p0
    _ = p0
    _ = v12
    return 464
}

func main() {
    var v1 bool = ((909 != 289) && false)
    v2 := make(map[uint32]string)
    var v3 int16 = 0
    var v4 int64 = -619
    v5 := make(map[uint32]string)
    if true {
        v1 = (!v1)
        var v13 float32 = 0.1
        v4 += v4
        _ = v13
    } else {
        for fi14 := 0; fi14 < 1; fi14++ {
            switch fi14 {
            case 32767:
                v4 += v4
            case 4294967295:
                fmt.Println(fi14, v3)
            case 9223372036854775807:
                _ = len(v5)
            }
            var v15 uint32 = 85
            _ = fi14
            _ = v15
        }
    }
    fmt.Println(v3)
    v1 = true
    if (-1 == genFn1(("" + ("hello" + "a")))) {
        fmt.Println(v3, v3, v3)
        var v16 float32 = (0.1 * float32(((-2.25 * (0.1 + -2.25)) + float64(1e10))))
        _ = v16
    }
    for fi17 := 0; fi17 < 3; fi17++ {
        v18 := GenS1{F0: 9223372036854775807}
        v19 := GenS1{F0: fi17}
        _ = fi17
        _ = v18
        _ = v19
    }
    _ = len(v5)
    delete(v5, genFn0(175, ("foo bar" + "\n\t"), 1))
    for fi20 := 0; fi20 < 6; fi20++ {
        for fi21 := 0; fi21 < 13; fi21++ {
            fmt.Println(v4, v4)
            _ = fi21
        }
        var v22 uint32 = 2
        _ = fi20
        _ = v22
    }
    v23 := GenS1{F0: int(v4)}
    var v24 bool = (!false)
    v25 := GenS3{F0: 42, F1: -579, F2: genFn1("")}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
