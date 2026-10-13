package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 int }
type GenS2 struct { F0 uint16; F1 uint; F2 uint64 }
type GenS3 struct { F0 int16; F1 int16 }
func genFn0(p0 int16, p1 int, p2 int) uint8 {
    v6 := GenS1{F0: 255, F1: 820, F2: (p1 & p2)}
    if true {
        for fi7 := 0; fi7 < 11; fi7++ {
            switch p1 {
            case 32767:
                var v8 bool = true
                _ = v8
            default:
                v9 := []int{p1, (p1 * (p1 - fi7))}
                _ = v9
            }
            _ = fi7
        }
        var v10 int16 = (-(p0))
        v11 := GenS1{F0: 141, F1: 923, F2: p2}
        _ = v10
        _ = v11
    }
    fmt.Println(p0, p2, p2)
    for fi12 := 0; fi12 < 2; fi12++ {
        switch -2147483648 {
        case -9223372036854775808:
            var v13 uint8 = 168
            _ = v13
        case 2147483647:
            p2 -= (-(fi12))
        }
        _ = fi12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 0
}
func genFn1(p0 uint8, p1 uint16, p2 int) uint16 {
    switch "a" {
    case "":
        var v14 uint16 = 32767
        _ = v14
    case "hello":
        p2 += 210
    default:
        v15 := []int{86}
        _ = v15
    }
    for fi16 := 0; fi16 < 18; fi16++ {
        var v17 uint64 = 2
        _ = fi16
        _ = v17
    }
    var v18 int = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    return 353
}

func main() {
    var v1 int16 = 690
    v2 := make(map[string]int)
    var v3 uint64 = 19
    var v4 float64 = float64(0.1)
    v5 := make(map[int]uint32)
    v19 := []int{1}
    var v20 int8 = 127
    if len(v19) > 0 {
        v19[0] = -508
    }
    if (genFn1(202, 32767, -1) < 0) {
        var v21 int16 = (v1 & 623)
        var v22 uint32 = 1
        v20 = v20
        _ = v21
        _ = v22
    }
    v23 := GenS1{F0: 18446744073709551615, F1: 1, F2: 128}
    v4 = 3.14159
    fmt.Println(v20)
    var v24 float32 = (((1e10 - 0.1) / ((0.0 * 0.0) - (1.5 / 1.0))) - float32(0.1))
    for fi25 := 0; fi25 < 11; fi25++ {
        v24 = v24
        v26 := GenS2{F0: 1, F1: 0, F2: v3}
        var v27 float64 = v4
        _ = fi25
        _ = v26
        _ = v27
    }
    var v28 int32 = 467
    v29 := GenS3{F0: v1, F1: (^(v1))}
    v20 = -1
    var v30 string = ("\n\t" + ("\n\t" + ("hello" + "foo bar")))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
