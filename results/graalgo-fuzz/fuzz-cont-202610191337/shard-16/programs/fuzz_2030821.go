package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 int8 }
type GenS2 struct { F0 uint16; F1 bool; F2 uint }
type GenS3 struct { F0 uint32; F1 int64 }
type GenS4 struct { F0 uint }
func genFn0() string {
    v6 := []int{32767, -985, 0}
    switch 4 {
    case 2147483647:
        v7 := GenS1{F0: 127, F1: "a", F2: 120}
        _ = v7
    case 1:
        for fi8 := 0; fi8 < 10; fi8++ {
            var v9 uint16 = 2
            if len(v6) > 0 {
                v6[0] = -109
            }
            _ = fi8
            _ = v9
        }
    }
    switch ("a" + (("\n\t" + ("a" + "\n\t")) + ("a" + "\n\t"))) {
    case "a":
        var v10 int16 = 438
        _ = v10
    case "hello":
        for fi11 := 0; fi11 < 8; fi11++ {
            fmt.Println(fi11)
            var v12 uint = 852
            if true {
                v13 := GenS1{F0: -128, F1: (("" + "") + "\n\t"), F2: 1}
                var v14 int16 = -32768
                _ = v13
                _ = v14
            }
            _ = fi11
            _ = v12
        }
    case "\n\t":
        var v15 bool = true
        _ = v15
    }
    _ = v6
    return ""
}
func genFn1(p0 uint8, p1 int64) uint32 {
    p0 = p0
    if ((1.5 < (3.14159 * 0.0)) || (((1e10 * (((1e10 + float32(0.1)) + 1.5) + 1.5)) < float32((((((0.0 * 3.14159) + -2.25) / -2.25) * 3.14159) - float64(((0.1 * 0.1) * ((0.0 * 1.5) * 0.1)))))) || false)) {
        var v16 uint32 = 508
        _ = v16
    } else {
        p0 += (^(p0))
    }
    p0 = 29
    v17 := []int{}
    _ = p0
    _ = p1
    _ = v17
    return 1
}

func main() {
    v1 := []int{32767, 9223372036854775807}
    var v2 float32 = -2.25
    var v3 float64 = float64(v2)
    v4 := make(map[uint32]bool)
    var v5 uint64 = 2
    for fi18 := 0; fi18 < 17; fi18++ {
        v5 = 0
        v19 := GenS2{F0: 391, F1: (("" + "hello") >= "a"), F2: 65535}
        switch "a" {
        case "a":
            fmt.Println(fi18, v5, fi18)
        case "":
            v5 += (v5 | v5)
        }
        _ = fi18
        _ = v19
    }
    v2 = v2
    var v20 float64 = (v3 - v3)
    switch "hello" {
    case "\n\t":
        fmt.Println(v2, v20)
    case "a":
        for fi21 := 0; fi21 < 19; fi21++ {
            v22 := v4[0]
            _ = fi21
            _ = v22
        }
    }
    var v23 uint8 = 127
    v2 = -2.25
    if true {
        for fi24 := 0; fi24 < 9; fi24++ {
            if (!true) {
                if (0 == 42) {
                    v25 := v4[65535]
                    var v26 int64 = -896
                    if v25 {
                        v4[68] = v25
                        v27 := GenS3{F0: genFn1(v23, -79), F1: (v26 + int64(fi24))}
                        fmt.Println(v3)
                        _ = v27
                    }
                    _ = v25
                    _ = v26
                }
                var v28 string = (("\n\t" + "a") + (("\n\t" + ("hello" + (("" + "foo bar") + "\n\t"))) + "hello"))
                for fi29 := 0; fi29 < 14; fi29++ {
                    if (!true) {
                        v30 := GenS3{F0: 42, F1: 1}
                        v20 = v20
                        _ = v30
                    }
                    _ = fi29
                }
                _ = v28
            }
            _ = fi24
        }
        v20 = 0.1
    } else {
        v5 -= v5
        if ("" < ("" + "foo bar")) {
            v5 -= (v5 | (v5 ^ 1))
        }
    }
    v31 := GenS1{F0: 48, F1: ("a" + "\n\t"), F2: 100}
    v2 = v2
    var v32 int8 = 87
    v33 := GenS4{F0: 2}
    var v34 uint = 4294967295
    for fi35 := 0; fi35 < 19; fi35++ {
        v2 += v2
        _ = fi35
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
