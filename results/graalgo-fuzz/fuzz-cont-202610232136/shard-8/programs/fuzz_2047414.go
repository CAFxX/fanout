package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 bool; F1 int8; F2 bool }
type GenS3 struct { F0 uint64; F1 bool }
func genFn0() int {
    for fi6 := 0; fi6 < 7; fi6++ {
        fmt.Println(fi6)
        v7 := GenS1{F0: 42, F1: -1}
        v8 := []int{}
        _ = fi6
        _ = v7
        _ = v8
    }
    if true {
        if true {
            for fi9 := 0; fi9 < 6; fi9++ {
                if (fi9 == fi9) {
                    var v10 float64 = (3.14159 / ((3.14159 / 0.1) / float64(1e10)))
                    v11 := []int{184, (fi9 / (((^(fi9))) | 1))}
                    _ = v10
                    _ = v11
                }
                fmt.Println(fi9)
                _ = fi9
            }
            for fi12 := 0; fi12 < 19; fi12++ {
                var v13 float32 = ((0.0 / 1.5) / 1.0)
                _ = fi12
                _ = v13
            }
        }
        for fi14 := 0; fi14 < 19; fi14++ {
            fmt.Println(fi14)
            if (32767 > 0) {
                var v15 int = fi14
                _ = v15
            }
            v16 := GenS1{F0: 2, F1: -741}
            _ = fi14
            _ = v16
        }
    }
    v17 := GenS1{F0: 762, F1: -1}
    for fi18 := 0; fi18 < 19; fi18++ {
        v19 := GenS1{F0: 654, F1: 2147483647}
        _ = fi18
        _ = v19
    }
    _ = v17
    return 282
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []uint32{65535, 113, 42, 2147483647}
    var v3 int64 = 189
    v4 := make(map[uint32]bool)
    var v5 int64 = (v3 >> ((18446744073709551615) & 63))
    v20 := GenS2{F0: (428 != -266), F1: 59, F2: (true || false)}
    switch ("\n\t" + ("" + ("\n\t" + "hello"))) {
    case "foo bar":
        v5 -= (v3 % (((v3 * 46)) | 1))
    }
    v5 = v5
    v5 += v5
    var v21 uint = 4294967295
    v5 -= v3
    var v22 int32 = -1
    fmt.Println(v21, v3)
    v5 -= (v3 | (v3 - v3))
    v21 += (v21 * (v21 / (((^(v21))) | 1)))
    fmt.Println(v21, v5)
    v21 -= v21
    if ("a" < "foo bar") {
        switch -192 {
        case 2147483647:
            for fi23 := 0; fi23 < 4; fi23++ {
                var v24 bool = true
                if len(v2) > 0 {
                    v2[0] = 42
                }
                _ = fi23
                _ = v24
            }
        case 954:
            v5 = (-(v3))
        default:
            if (!true) {
                var v25 uint8 = 2
                for fi26 := 0; fi26 < 3; fi26++ {
                    var v27 string = ((("foo bar" + "hello") + "\n\t") + (("foo bar" + "foo bar") + "a"))
                    v27 = v27
                    var v28 uint = (v21 << ((v21) & 63))
                    _ = fi26
                    _ = v27
                    _ = v28
                }
                _ = v25
            } else {
                v1["hello"] = 2147483647
                _ = len(v1)
                v21 += v21
            }
        }
        v29 := GenS3{F0: 519, F1: false}
        _ = v29
    }
    var v30 uint32 = 585
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
