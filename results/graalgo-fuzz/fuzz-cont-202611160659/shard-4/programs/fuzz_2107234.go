package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 string }
type GenS2 struct { F0 uint8; F1 int; F2 int32 }
type GenS3 struct { F0 uint64; F1 int64; F2 uint32 }
func genFn0(p0 int8) int32 {
    p0 -= p0
    for fi6 := 0; fi6 < 19; fi6++ {
        var v7 uint64 = 271
        _ = fi6
        _ = v7
    }
    var v8 uint8 = 42
    _ = p0
    _ = v8
    return -63
}
func genFn1(p0 int64, p1 string, p2 uint64) uint8 {
    v9 := GenS1{F0: (p0 / ((p0) | 1)), F1: 127, F2: p1}
    for fi10 := 0; fi10 < 16; fi10++ {
        switch (p1 + p1) {
        case "a":
            v11 := GenS1{F0: p0, F1: 254, F2: ""}
            _ = v11
        case "hello":
            for fi12 := 0; fi12 < 11; fi12++ {
                p2 = p2
                if (0 < (-(p0))) {
                    if ((false && true) || true) {
                        var v13 uint16 = 32767
                        _ = v13
                    } else {
                        v14 := []int{-263, 32768, (fi10 ^ (fi10 >> ((2) & 63)))}
                        _ = v14
                    }
                    if true {
                        var v15 int64 = (p0 * (p0 ^ 0))
                        v15 += v15
                        fmt.Println(p0, fi12)
                        p1 = ""
                        fmt.Println(v15, p0)
                        v15 -= 9223372036854775807
                        _ = v15
                    }
                } else {
                    v16 := GenS2{F0: 127, F1: -1, F2: -129}
                    _ = v16
                }
                _ = fi12
            }
        case "\n\t":
            v17 := GenS1{F0: (p0 | p0), F1: 128, F2: ""}
            _ = v17
        default:
            v18 := []int{fi10, (fi10 | (fi10 + fi10)), 0}
            _ = v18
        }
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 128
}

func main() {
    var v1 int16 = 0
    var v2 uint64 = 43
    var v3 int8 = 42
    var v4 float32 = (1.5 - (1e10 - (-2.25 * (((((3.14159 / 1.0) + float32(0.1)) / float32((3.14159 - 0.0))) + 0.0) - 1.5))))
    v5 := []int{32767, -9223372036854775808, -715, -233}
    fmt.Println(v1)
    var v19 uint = 0
    v20 := GenS1{F0: -149, F1: 2, F2: "hello"}
    for fi21 := 0; fi21 < 20; fi21++ {
        v22 := GenS2{F0: 42, F1: (-(fi21)), F2: 44}
        for fi23 := 0; fi23 < 5; fi23++ {
            v19 += (v19 ^ (v19 + (v19 * (^(v19)))))
            v24 := GenS3{F0: v2, F1: 0, F2: 42}
            var v25 uint16 = 36
            _ = fi23
            _ = v24
            _ = v25
        }
        _ = fi21
        _ = v22
    }
    v1 = v1
    var v26 string = (((("a" + ("" + "hello")) + ((("a" + ("hello" + "a")) + (("a" + "\n\t") + "")) + "hello")) + "") + "\n\t")
    if false {
        v26 = v26
        var v27 int8 = v3
        _ = v27
    }
    for fi28 := 0; fi28 < 12; fi28++ {
        v1 -= (v1 * -1)
        v29 := GenS1{F0: -455, F1: 2, F2: ("" + "\n\t")}
        _ = fi28
        _ = v29
    }
    var v30 uint = v19
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v26
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
