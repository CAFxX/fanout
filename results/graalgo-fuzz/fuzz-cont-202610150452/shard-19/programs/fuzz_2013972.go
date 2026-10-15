package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 uint64 }
type GenS2 struct { F0 uint }

func main() {
    var v1 uint8 = 42
    var v2 int = -212
    var v3 string = "foo bar"
    v4 := make(map[string]uint32)
    var v5 uint = 18446744073709551615
    for fi6 := 0; fi6 < 8; fi6++ {
        v2 -= (fi6 | (v2 % ((v2) | 1)))
        if true {
            switch (((("a" + "\n\t") + (("foo bar" + v3) + ((("" + v3) + v3) + v3))) + ("\n\t" + v3)) + (v3 + v3)) {
            case "\n\t":
                v7 := GenS1{F0: 1, F1: 306, F2: 0}
                _ = v7
            default:
                if ((1 < 541) && (2 == 539)) {
                    fmt.Println(fi6)
                    var v8 int16 = -32768
                    var v9 int16 = v8
                    _ = v8
                    _ = v9
                } else {
                    var v10 string = v3
                    var v11 int16 = 32767
                    _ = v10
                    _ = v11
                }
            }
        } else {
            if (!(false || false)) {
                var v12 int16 = -786
                for fi13 := 0; fi13 < 20; fi13++ {
                    v2 += (-(v2))
                    _ = fi13
                }
                _ = v12
            }
            v14 := v4[(("\n\t" + "\n\t") + v3)]
            for fi15 := 0; fi15 < 16; fi15++ {
                switch fi15 {
                case 2147483647:
                    v16 := GenS1{F0: 0, F1: v14, F2: 42}
                    _ = v16
                default:
                    var v17 uint8 = (v1 & (^(v1)))
                    _ = v17
                }
                _ = fi15
            }
            _ = v14
        }
        if true {
            fmt.Println(v2)
            _ = len(v4)
            v5 = v5
        } else {
            v18 := GenS1{F0: 42, F1: 65535, F2: 4294967295}
            _ = v18
        }
        _ = fi6
    }
    var v19 bool = true
    v1 = 0
    v1 = (^(v1))
    v20 := GenS1{F0: uint32(v1), F1: 848, F2: 1}
    switch (("" + "foo bar") + (v3 + v3)) {
    case "a":
        var v21 uint64 = 561
        _ = v21
    }
    if (1e10 >= float64((1.5 + (0.0 - 1.5)))) {
        fmt.Println(v2, v1, v5)
        for fi22 := 0; fi22 < 4; fi22++ {
            for fi23 := 0; fi23 < 19; fi23++ {
                v24 := []int{fi22, 2147483647}
                _ = fi23
                _ = v24
            }
            _ = fi22
        }
    } else {
        v25 := GenS1{F0: 1, F1: 0, F2: 4294967295}
        var v26 uint = (^(v5))
        v2 = (v2 | -908)
        _ = v25
        _ = v26
    }
    switch v3 {
    case "foo bar":
        var v27 string = (v3 + v3)
        _ = v27
    case "\n\t":
        v28 := []int{}
        _ = v28
    case "":
        v3 = ((v3 + (v3 + "\n\t")) + v3)
    }
    v29 := GenS2{F0: v5}
    v2 -= (v2 + 9223372036854775807)
    v5 = 42
    var v30 int = (v2 / ((v2) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
