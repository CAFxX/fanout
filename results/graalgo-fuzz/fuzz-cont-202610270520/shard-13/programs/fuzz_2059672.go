package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 uint }
type GenS2 struct { F0 uint32; F1 uint32 }
func genFn0(p0 uint16, p1 uint16) uint8 {
    switch 127 {
    case 127:
        p0 = 7
    case 621:
        var v6 uint64 = 4294967295
        _ = v6
    case 690:
        for fi7 := 0; fi7 < 20; fi7++ {
            p1 = p1
            v8 := GenS1{F0: 127, F1: ((("\n\t" + "foo bar") + "foo bar") + ("hello" + ("hello" + "a"))), F2: 2}
            for fi9 := 0; fi9 < 2; fi9++ {
                v10 := []int{fi7, -147, 722}
                p1 = p1
                _ = fi9
                _ = v10
            }
            _ = fi7
            _ = v8
        }
    default:
        for fi11 := 0; fi11 < 5; fi11++ {
            var v12 int8 = 0
            _ = fi11
            _ = v12
        }
    }
    p0 -= (p1 | (p0 / ((65535) | 1)))
    _ = p0
    _ = p1
    return 12
}

func main() {
    var v1 float32 = (float32(3.14159) - 1e10)
    var v2 int64 = -763
    var v3 uint16 = 2
    v4 := make(map[uint32]string)
    var v5 uint8 = 2
    if false {
        v1 += (3.14159 * (v1 - (3.14159 - 3.14159)))
        fmt.Println(v3)
        var v13 int = -659
        _ = v13
    }
    for fi14 := 0; fi14 < 5; fi14++ {
        v15 := GenS1{F0: genFn0((v3 - v3), (v3 ^ (^(v3)))), F1: "", F2: 1}
        v2 -= v2
        _ = fi14
        _ = v15
    }
    v2 = v2
    v16 := []int{32767, -2147483648}
    var v17 int8 = -128
    switch 2147483647 {
    case -644:
        if true {
            var v18 int8 = v17
            _ = v18
        }
    case -348:
        var v19 int = -1
        _ = v19
    }
    v20 := GenS1{F0: v5, F1: (("\n\t" + "\n\t") + ("" + "foo bar")), F2: 64}
    var v21 uint64 = 2
    delete(v4, 65535)
    for fi22 := 0; fi22 < 20; fi22++ {
        fmt.Println(v17, v3)
        _ = fi22
    }
    switch 127 {
    case 4294967295:
        fmt.Println(v17, v3, v3)
    case 42:
        v5 = 211
    default:
        if true {
            fmt.Println(v3, v17, v2)
        } else {
            var v23 int16 = -536
            for fi24 := 0; fi24 < 5; fi24++ {
                v21 -= v21
                _ = fi24
            }
            v25 := GenS2{F0: 660, F1: 880}
            _ = v23
            _ = v25
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
