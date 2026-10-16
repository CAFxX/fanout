package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int; F1 bool }
func genFn0() string {
    if (false || true) {
        var v4 uint8 = 2
        v5 := GenS1{F0: true}
        _ = v4
        _ = v5
    }
    var v6 bool = (float32(-2.25) >= 1e10)
    for fi7 := 0; fi7 < 5; fi7++ {
        var v8 uint16 = 898
        _ = fi7
        _ = v8
    }
    _ = v6
    return ("foo bar" + "\n\t")
}
func genFn1(p0 int, p1 uint32) string {
    p1 = 346
    v9 := []int{(-(p0))}
    var v10 int64 = 918
    switch (p0 << ((384) & 63)) {
    case 0:
        for fi11 := 0; fi11 < 12; fi11++ {
            p1 = p1
            v12 := GenS1{F0: (!(42 > 1))}
            _ = fi11
            _ = v12
        }
    case 32767:
        p0 = -900
    case 4294967295:
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return "\n\t"
}

func main() {
    v1 := make(map[int]int)
    var v2 uint8 = 127
    v3 := make(map[string]string)
    v2 -= (v2 / ((127) | 1))
    delete(v3, ("\n\t" + "\n\t"))
    switch -1 {
    case -556:
        for fi13 := 0; fi13 < 13; fi13++ {
            v2 -= v2
            _ = fi13
        }
    }
    var v14 uint8 = v2
    v15 := GenS1{F0: (true || (((3.14159 - 0.1) - -2.25) == 1e10))}
    switch ("" + (("\n\t" + "\n\t") + "")) {
    case "":
        v16 := GenS2{F0: 118, F1: (true && false)}
        _ = v16
    case "foo bar":
        var v17 float64 = ((0.1 - 1.5) + 1e10)
        _ = v17
    case "\n\t":
        if false {
            v14 += (v2 | 1)
        }
    }
    switch "foo bar" {
    case "":
        var v18 uint16 = 359
        _ = v18
    case "hello":
        switch "a" {
        case "":
            v19 := []int{127, -614}
            _ = v19
        case "hello":
            var v20 int32 = -2147483648
            _ = v20
        case "\n\t":
            v14 += 255
        default:
            for fi21 := 0; fi21 < 16; fi21++ {
                v2 -= 196
                _ = fi21
            }
        }
    case "\n\t":
        var v22 uint32 = 439
        _ = v22
    }
    v14 += (v2 / ((v14) | 1))
    v23 := []int{}
    v2 = v2
    _ = len(v1)
    v1[127] = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v23
    _ = v3
}
