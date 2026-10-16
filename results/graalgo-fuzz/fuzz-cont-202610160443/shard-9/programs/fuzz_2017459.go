package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool; F1 int32 }
func genFn0(p0 uint64, p1 int32) uint32 {
    for fi7 := 0; fi7 < 7; fi7++ {
        var v8 uint64 = (^(p0))
        for fi9 := 0; fi9 < 16; fi9++ {
            var v10 int64 = -967
            p1 = 1
            if true {
                p1 += (p1 | p1)
                var v11 int32 = int32(v10)
                v12 := GenS1{F0: -558}
                _ = v11
                _ = v12
            }
            _ = fi9
            _ = v10
        }
        fmt.Println(p0, fi7)
        _ = fi7
        _ = v8
    }
    switch -225 {
    case 128:
        fmt.Println(p0, p1)
    }
    _ = p0
    _ = p1
    return 1
}
func genFn1() int8 {
    var v13 float64 = 1e10
    v14 := []int{-96, 128}
    v15 := GenS2{F0: ((!(!true)) && false), F1: 42}
    _ = v13
    _ = v14
    _ = v15
    return 42
}

func main() {
    var v1 int = -9223372036854775808
    var v2 int16 = 32767
    var v3 float32 = (-2.25 + float32(1e10))
    var v4 float64 = 1e10
    v5 := []string{"hello", "\n\t", "foo bar", ("foo bar" + "foo bar")}
    var v6 float64 = v4
    v3 += v3
    v3 -= (0.0 / v3)
    v4 = 1.5
    v16 := GenS1{F0: 42}
    v1 -= v1
    var v17 int16 = 1
    var v18 string = "a"
    if (4294967295 == 1) {
        var v19 uint32 = 24
        v17 -= v2
        _ = v19
    } else {
        var v20 string = "\n\t"
        _ = v20
    }
    v21 := GenS1{F0: v17}
    if true {
        if ((!true) || false) {
            var v22 uint32 = 74
            v18 = ("\n\t" + v18)
            _ = v22
        }
        var v23 int32 = 1
        var v24 int8 = 127
        _ = v23
        _ = v24
    } else {
        var v25 uint8 = 87
        for fi26 := 0; fi26 < 13; fi26++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi26
        }
        v2 -= -638
        _ = v25
    }
    switch v18 {
    case "foo bar":
        switch v1 {
        case -1:
            _ = len(v5)
        }
    case "a":
        for fi27 := 0; fi27 < 18; fi27++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            var v28 int16 = (v17 >> ((255) & 63))
            v29 := GenS1{F0: (v17 / (((v17 | v28)) | 1))}
            _ = fi27
            _ = v28
            _ = v29
        }
    }
    fmt.Println(v17)
    var v30 float64 = v6
    var v31 uint32 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
