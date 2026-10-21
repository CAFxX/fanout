package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string; F1 uint64; F2 uint16 }

func main() {
    v1 := make(map[string]uint32)
    var v2 bool = (!false)
    var v3 float32 = -2.25
    v4 := make(map[uint32]uint32)
    switch -924 {
    case -2147483648:
        fmt.Println(v2, v3)
    case -707:
        v5 := GenS1{F0: 10}
        _ = v5
    case -1:
        for fi6 := 0; fi6 < 3; fi6++ {
            var v7 int32 = -1
            _ = fi6
            _ = v7
        }
    default:
        var v8 uint = 243
        _ = v8
    }
    for fi9 := 0; fi9 < 1; fi9++ {
        fmt.Println(fi9, v2, v2)
        v2 = ("hello" > "hello")
        v10 := v4[65535]
        _ = fi9
        _ = v10
    }
    v1[""] = 2
    var v11 float64 = (float64(v3) * (((float64((-2.25 / 3.14159)) / 1.0) - (1.5 + -2.25)) * (-2.25 + 3.14159)))
    var v12 int16 = -523
    for fi13 := 0; fi13 < 19; fi13++ {
        if (v2 || v2) {
            v14 := GenS1{F0: 166}
            _ = v14
        } else {
            v15 := GenS2{F0: -335}
            v16 := GenS1{F0: 127}
            switch "foo bar" {
            case "a":
                if (v3 != v3) {
                    switch (fi13 | (fi13 | (fi13 >> ((935) & 63)))) {
                    case 1:
                        v11 -= v11
                    case -141:
                        fmt.Println(fi13, fi13)
                    case 32768:
                        var v17 string = "a"
                        _ = v17
                    }
                    v18 := GenS2{F0: -9223372036854775808}
                    v19 := []int{(^(fi13)), (fi13 / ((fi13) | 1)), -9223372036854775808, fi13}
                    _ = v18
                    _ = v19
                } else {
                    if v2 {
                        if (((127 >= 100) || (0 >= 9223372036854775807)) || false) {
                            v20 := GenS1{F0: 1}
                            for fi21 := 0; fi21 < 4; fi21++ {
                                v12 = (v12 | 173)
                                _ = fi21
                            }
                            v22 := []int{-179, -610, (fi13 ^ (fi13 / ((127) | 1))), fi13}
                            _ = v20
                            _ = v22
                        } else {
                            v4[65535] = 2147483648
                            v23 := GenS3{F0: "hello", F1: 833, F2: 65535}
                            _ = v23
                        }
                    }
                    var v24 string = "a"
                    _ = v24
                }
            default:
                v3 = 0.1
            }
            _ = v15
            _ = v16
        }
        _ = fi13
    }
    switch ("\n\t" + "\n\t") {
    case "a":
        fmt.Println(v3, v3, v3)
    case "\n\t":
        v25 := GenS3{F0: ("a" + "hello"), F1: 0, F2: 2}
        _ = v25
    case "foo bar":
        v12 = (v12 ^ (v12 + v12))
    default:
        var v26 uint16 = 255
        _ = v26
    }
    v27 := []int{9223372036854775807, 789, 9223372036854775807}
    if len(v27) >= 2 {
        v27 = v27[1:len(v27)]
    }
    var v28 uint64 = 791
    v28 += v28
    v2 = (v2 || (v3 == 0.1))
    var v29 uint16 = 464
    var v30 int32 = 42
    for fi31 := 0; fi31 < 16; fi31++ {
        v11 = float64((((v3 + 1e10) - 1e10) - v3))
        fmt.Println(v3, v11, v29)
        delete(v1, "hello")
        _ = fi31
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
