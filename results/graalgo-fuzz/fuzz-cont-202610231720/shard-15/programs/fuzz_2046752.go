package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 uint8 }
type GenS2 struct { F0 uint; F1 string }

func main() {
    var v1 string = "\n\t"
    v2 := []uint32{}
    v3 := []int{4294967295, -36, 1}
    switch (v1 + ("" + "hello")) {
    case "a":
        var v4 int16 = -580
        _ = v4
    case "foo bar":
        fmt.Println(v1)
    case "hello":
        if len(v2) > 0 {
            v2[0] = 2147483648
        }
    default:
        if len(v3) > 0 {
            v3[0] = -629
        }
    }
    v1 = v1
    switch (v1 + "\n\t") {
    case "a":
        v5 := GenS1{F0: 806, F1: 128, F2: 128}
        _ = v5
    }
    fmt.Println(v1)
    for fi6 := 0; fi6 < 9; fi6++ {
        fmt.Println(fi6)
        _ = fi6
    }
    var v7 float32 = 1e10
    v1 = ((((v1 + v1) + "") + v1) + "foo bar")
    switch v1 {
    case "":
        for fi8 := 0; fi8 < 16; fi8++ {
            switch (fi8 ^ (fi8 + fi8)) {
            case -1:
                fmt.Println(v1, fi8, v7)
            case 32767:
                v2 = append(v2, 457)
            case 128:
                for fi9 := 0; fi9 < 4; fi9++ {
                    if len(v2) > 0 {
                        v2[0] = 42
                    }
                    v10 := GenS1{F0: -2147483648, F1: 255, F2: 127}
                    v7 -= float32(-2.25)
                    _ = fi9
                    _ = v10
                }
            default:
                for fi11 := 0; fi11 < 15; fi11++ {
                    var v12 string = (v1 + v1)
                    _ = fi11
                    _ = v12
                }
            }
            v13 := GenS2{F0: 4294967295, F1: (v1 + v1)}
            v7 = v7
            _ = fi8
            _ = v13
        }
    }
    for fi14 := 0; fi14 < 18; fi14++ {
        v15 := GenS1{F0: 2147483647, F1: 42, F2: 2}
        _ = fi14
        _ = v15
    }
    _ = len(v2)
    v1 = v1
    for fi16 := 0; fi16 < 16; fi16++ {
        v3 = append(v3, (^(fi16)))
        _ = len(v3)
        _ = fi16
    }
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v7
}
