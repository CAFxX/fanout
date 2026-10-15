package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 uint16 }
func genFn0(p0 int16) string {
    var v5 string = ""
    v6 := []int{32767, 591}
    for fi7 := 0; fi7 < 1; fi7++ {
        if len(v6) > 0 {
            v6[0] = -444
        }
        _ = fi7
    }
    var v8 int32 = -650
    _ = p0
    _ = v5
    _ = v6
    _ = v8
    return ((v5 + "foo bar") + "")
}
func genFn1(p0 int16, p1 bool) int {
    fmt.Println(p0, p0)
    for fi9 := 0; fi9 < 10; fi9++ {
        for fi10 := 0; fi10 < 16; fi10++ {
            v11 := []int{2147483647, fi10, (^(fi10)), 127}
            _ = fi10
            _ = v11
        }
        v12 := []int{(fi9 >> ((4294967295) & 63))}
        var v13 bool = (!p1)
        _ = fi9
        _ = v12
        _ = v13
    }
    if p1 {
        v14 := GenS1{F0: 42, F1: 109, F2: 42}
        _ = v14
    } else {
        p1 = (p1 || p1)
    }
    var v15 string = "hello"
    v16 := []int{-681, 32767}
    _ = p0
    _ = p1
    _ = v15
    _ = v16
    return 236
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]bool)
    var v3 uint8 = 42
    v4 := make(map[string]int)
    v17 := GenS1{F0: 1, F1: 127, F2: 106}
    switch ((("a" + ("foo bar" + (("a" + ("\n\t" + "foo bar")) + ""))) + "hello") + "") {
    case "a":
        delete(v2, 297)
    default:
        v18 := GenS1{F0: 9223372036854775807, F1: 0, F2: 2}
        _ = v18
    }
    for fi19 := 0; fi19 < 7; fi19++ {
        v20 := v4[("" + "hello")]
        v21 := GenS1{F0: (v20 / (((-(fi19))) | 1)), F1: -128, F2: 909}
        var v22 bool = false
        _ = fi19
        _ = v20
        _ = v21
        _ = v22
    }
    var v23 int16 = 42
    v24 := []int{164, -9223372036854775808, 686}
    v24 = append(v24, -277)
    switch "a" {
    case "\n\t":
        v1[42] = true
    case "foo bar":
        if false {
            var v25 float32 = (0.0 * 3.14159)
            fmt.Println(v25, v3, v25)
            _ = v25
        }
    case "hello":
        switch "foo bar" {
        case "":
            v26 := GenS1{F0: 61, F1: -1, F2: 526}
            _ = v26
        case "\n\t":
            if len(v24) > 0 {
                v24[0] = genFn1((v23 & -32768), (-9223372036854775808 >= 42))
            }
        }
    default:
        if false {
            fmt.Println(v23)
            switch ("a" + "a") {
            case "hello":
                fmt.Println(v3, v3)
            }
        } else {
            var v27 uint8 = 127
            if true {
                v23 += v23
            } else {
                v27 += 255
            }
            _ = v27
        }
    }
    var v28 uint16 = 245
    if ((1 <= 2) || ((false || (((float64(1e10) * 0.1) <= 0.1) || false)) && true)) {
        var v29 int32 = 1
        v2[898] = false
        if false {
            v28 = 65535
            v23 = 383
            switch "" {
            case "a":
                var v30 int8 = -1
                _ = v30
            case "\n\t":
                var v31 uint16 = 519
                _ = v31
            }
        } else {
            var v32 float64 = 0.0
            v33 := GenS1{F0: 9223372036854775807, F1: 1, F2: v28}
            _ = v32
            _ = v33
        }
        _ = v29
    }
    var v34 float32 = 1.5
    v23 -= v23
    v34 = v34
    v34 += v34
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v17
    _ = v2
    _ = v23
    _ = v24
    _ = v28
    _ = v3
    _ = v34
    _ = v4
}
