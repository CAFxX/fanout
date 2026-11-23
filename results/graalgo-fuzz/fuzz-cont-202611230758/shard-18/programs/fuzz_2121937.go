package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 int; F2 uint16 }
func genFn0(p0 string, p1 int32) string {
    fmt.Println(p0, p1)
    for fi4 := 0; fi4 < 13; fi4++ {
        v5 := []int{32768, fi4}
        _ = fi4
        _ = v5
    }
    switch (("" + (p0 + (p0 + p0))) + (p0 + p0)) {
    case "a":
        fmt.Println(p0, p1)
    case "hello":
        if ((4294967295 != 579) || (1 <= 484)) {
            var v6 float32 = float32(1e10)
            _ = v6
        } else {
            p1 += (p1 + (p1 % ((-2147483648) | 1)))
            var v7 uint32 = 229
            _ = v7
        }
    default:
        var v8 uint32 = 1
        _ = v8
    }
    var v9 int32 = p1
    v10 := []int{-93, 2147483647, 0, 128}
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return "hello"
}
func genFn1(p0 bool, p1 uint) uint32 {
    fmt.Println(p0, p1)
    if p0 {
        v11 := []int{-493}
        _ = v11
    }
    var v12 uint = 42
    switch 2147483647 {
    case 32767:
        v13 := []int{511, 623}
        _ = v13
    case 42:
        fmt.Println(v12)
    case 318:
        p0 = (p0 || ((v12 * (v12 << (((v12 << ((p1) & 63))) & 63))) == 255))
    }
    _ = p0
    _ = p1
    _ = v12
    return 0
}

func main() {
    v1 := []uint32{42}
    var v2 uint32 = 732
    v3 := make(map[uint32]bool)
    for fi14 := 0; fi14 < 13; fi14++ {
        switch fi14 {
        case -701:
            fmt.Println(fi14)
        }
        var v15 uint = 561
        _ = fi14
        _ = v15
    }
    switch "foo bar" {
    case "":
        if ((((65535 < 2) && ("" > ("foo bar" + "hello"))) || (1 < 0)) || (false || true)) {
            v16 := GenS1{F0: 1}
            fmt.Println(v2)
            _ = v16
        } else {
            var v17 float64 = float64((1e10 - 0.1))
            for fi18 := 0; fi18 < 15; fi18++ {
                v2 = v2
                v19 := v3[(v2 << ((310) & 63))]
                for fi20 := 0; fi20 < 14; fi20++ {
                    v19 = true
                    _ = fi20
                }
                _ = fi18
                _ = v19
            }
            _ = v17
        }
    case "\n\t":
        v2 -= v2
    case "a":
        var v21 bool = true
        _ = v21
    }
    for fi22 := 0; fi22 < 6; fi22++ {
        for fi23 := 0; fi23 < 6; fi23++ {
            switch (fi22 * (fi23 % ((fi23) | 1))) {
            case 128:
                v24 := GenS2{F0: -721, F1: (fi22 >> ((4294967295) & 63)), F2: 1}
                _ = v24
            case 176:
                var v25 int = (fi22 | (fi22 - 0))
                _ = v25
            case 1:
                fmt.Println(v2, fi22)
            }
            v2 -= (^(v2))
            _ = fi23
        }
        v26 := v3[(v2 ^ 469)]
        v27 := v3[(v2 ^ (v2 + v2))]
        _ = fi22
        _ = v26
        _ = v27
    }
    var v28 string = ("foo bar" + "")
    v29 := GenS1{F0: 127}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if (!false) {
        fmt.Println(v28, v2)
    }
    var v30 float32 = 0.1
    if true {
        v28 = "foo bar"
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
}
