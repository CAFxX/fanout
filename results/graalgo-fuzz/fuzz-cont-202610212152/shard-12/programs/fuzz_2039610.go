package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 string; F1 uint32; F2 uint }

func main() {
    var v1 int64 = 310
    var v2 bool = false
    var v3 uint64 = 42
    for fi4 := 0; fi4 < 6; fi4++ {
        if v2 {
            fmt.Println(v1, v2, v3)
            if ("\n\t" != "\n\t") {
                var v5 uint32 = 638
                _ = v5
            }
            if false {
                if (1.5 == 3.14159) {
                    v6 := []int{(fi4 * (fi4 >> ((623) & 63)))}
                    _ = len(v6)
                    _ = v6
                } else {
                    v7 := []int{251, 32767, int(v1)}
                    switch (fi4 & (fi4 - (fi4 ^ fi4))) {
                    case 9223372036854775807:
                        var v8 int16 = 32767
                        _ = v8
                    }
                    _ = v7
                }
            } else {
                var v9 uint8 = 190
                for fi10 := 0; fi10 < 17; fi10++ {
                    switch fi10 {
                    case 0:
                        v11 := GenS1{F0: -708, F1: 42}
                        _ = v11
                    case -745:
                        var v12 uint16 = 717
                        _ = v12
                    case 1:
                        var v13 float32 = 0.0
                        _ = v13
                    }
                    var v14 uint32 = 42
                    _ = fi10
                    _ = v14
                }
                _ = v9
            }
        } else {
            v15 := []int{(fi4 / (((fi4 + 32767)) | 1)), 47}
            switch ("a" + "foo bar") {
            case "":
                v1 = v1
            default:
                if (!v2) {
                    v1 += (v1 & v1)
                }
            }
            _ = len(v15)
            _ = v15
        }
        _ = fi4
    }
    switch "\n\t" {
    case "foo bar":
        v2 = (!(238 > 255))
    case "\n\t":
        var v16 float32 = 0.0
        _ = v16
    }
    v17 := GenS1{F0: -332, F1: 2}
    v2 = (v2 || v2)
    for fi18 := 0; fi18 < 2; fi18++ {
        if false {
            fmt.Println(fi18, v3, v3)
        }
        var v19 int8 = -128
        switch "a" {
        case "\n\t":
            v20 := []int{(-(fi18)), (^(fi18)), fi18}
            _ = v20
        default:
            var v21 float64 = 3.14159
            _ = v21
        }
        _ = fi18
        _ = v19
    }
    for fi22 := 0; fi22 < 18; fi22++ {
        v23 := []int{(fi22 ^ (fi22 * (fi22 * (fi22 - (fi22 >> ((399) & 63)))))), (fi22 + (fi22 >> ((399) & 63))), (^(fi22)), int(v1)}
        v2 = v2
        _ = fi22
        _ = v23
    }
    v24 := []int{}
    if v2 {
        var v25 float64 = (-2.25 + (0.0 + (3.14159 - 0.0)))
        var v26 uint16 = 2
        _ = len(v24)
        _ = v25
        _ = v26
    } else {
        var v27 int8 = 10
        switch 9223372036854775807 {
        case 4294967295:
            fmt.Println(v2)
        case 573:
            v2 = (!false)
        }
        _ = v27
    }
    if (0.1 >= (float64(3.14159) * -2.25)) {
        fmt.Println(v1, v2)
        v28 := GenS2{F0: (("" + "foo bar") + "\n\t"), F1: 2147483648, F2: 255}
        _ = v28
    }
    v29 := GenS2{F0: "foo bar", F1: 0, F2: 2}
    var v30 int16 = 937
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v17
    _ = v2
    _ = v24
    _ = v29
    _ = v3
    _ = v30
}
