package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 int64 }
type GenS2 struct { F0 int; F1 string }
func genFn0(p0 uint16, p1 uint32) int8 {
    p1 -= (^(p1))
    v4 := []int{}
    v5 := GenS1{F0: "hello", F1: 2, F2: 42}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 42
}
func genFn1(p0 int64, p1 uint8, p2 int8) uint64 {
    v6 := GenS1{F0: ((("\n\t" + "foo bar") + ("" + "hello")) + ("foo bar" + ((("a" + "") + "") + ("hello" + (("" + ((("a" + ("a" + "\n\t")) + "\n\t") + "a")) + "foo bar"))))), F1: 98, F2: -1}
    v7 := GenS2{F0: int(p1), F1: ("foo bar" + (("hello" + ("" + "a")) + "\n\t"))}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 560
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[string]int)
    v3 := make(map[string]int)
    v8 := GenS1{F0: "\n\t", F1: 625, F2: 42}
    var v9 float32 = 3.14159
    v10 := []int{}
    for fi11 := 0; fi11 < 15; fi11++ {
        var v12 int8 = 127
        var v13 int8 = (v12 ^ (v12 ^ -1))
        _ = fi11
        _ = v12
        _ = v13
    }
    if true {
        fmt.Println(v9)
        var v14 uint = 2
        _ = v14
    } else {
        var v15 int8 = 42
        var v16 uint32 = 2
        if true {
            v16 -= 42
            v17 := v3["\n\t"]
            var v18 float32 = v9
            _ = v17
            _ = v18
        }
        _ = v15
        _ = v16
    }
    v9 -= ((v9 * v9) * float32(0.1))
    for fi19 := 0; fi19 < 8; fi19++ {
        switch (fi19 / (((fi19 + (fi19 - fi19))) | 1)) {
        case 32767:
            for fi20 := 0; fi20 < 6; fi20++ {
                if (373 >= 42) {
                    v9 = v9
                    fmt.Println(v9, fi20)
                }
                if ((fi20 * (^(fi20))) != 790) {
                    var v21 float32 = v9
                    _ = v21
                }
                var v22 int64 = -1
                var v23 float32 = (v9 / 1.0)
                _ = fi20
                _ = v22
                _ = v23
            }
        default:
            fmt.Println(fi19, v9)
        }
        _ = fi19
    }
    v24 := v2[""]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v24
    _ = v3
    _ = v8
    _ = v9
}
