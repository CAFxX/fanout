package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint64; F1 uint64 }
type GenS4 struct { F0 uint64; F1 uint }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint8; F1 uint16; F2 uint }
func genFn0() uint {
    var v6 int16 = -32768
    var v7 int16 = v6
    v8 := GenS1{F0: 2}
    var v9 int = 906
    var v10 string = ("a" + ("" + "\n\t"))
    _ = v10
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 255
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]int)
    v3 := []int{}
    v4 := make(map[int]int)
    var v5 uint = 4294967295
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v5 -= v5
    v11 := GenS2{F0: 128}
    _ = len(v4)
    v12 := GenS2{F0: 1}
    if true {
        if len(v3) > 0 {
            v3[0] = 1
        }
        v5 = (v5 | genFn0())
        v2[""] = 2147483647
    }
    for fi13 := 0; fi13 < 19; fi13++ {
        for fi14 := 0; fi14 < 5; fi14++ {
            if false {
                if false {
                    v15 := GenS3{F0: 4294967295, F1: 18446744073709551615}
                    _ = v15
                } else {
                    fmt.Println(v5)
                    v16 := GenS1{F0: 255}
                    fmt.Println(v5, fi14, fi14)
                    _ = v16
                }
            } else {
                var v17 bool = true
                var v18 float32 = float32(((1e10 - (3.14159 + 0.0)) / ((0.0 - 0.0) - -2.25)))
                fmt.Println(v18)
                _ = v17
                _ = v18
            }
            v19 := GenS4{F0: 0, F1: 124}
            switch fi14 {
            case 2147483647:
                v20 := v2["a"]
                _ = v20
            case 661:
                var v21 float64 = float64(float32((((0.0 + 0.0) + 1.5) * 3.14159)))
                _ = v21
            default:
                if (4294967295 != 1) {
                    if (3.14159 <= ((3.14159 * 1.5) * 1e10)) {
                        switch "" {
                        case "":
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        default:
                            v22 := GenS5{F0: 759}
                            _ = v22
                        }
                        v23 := GenS2{F0: 2}
                        _ = v23
                    } else {
                        if false {
                            var v24 uint64 = 0
                            _ = v24
                        }
                    }
                    _ = len(v1)
                }
            }
            _ = fi14
            _ = v19
        }
        var v25 string = ("" + "foo bar")
        _ = fi13
        _ = v25
    }
    var v26 uint32 = 248
    if (!false) {
        for fi27 := 0; fi27 < 3; fi27++ {
            var v28 int = fi27
            _ = fi27
            _ = v28
        }
        fmt.Println(v5)
        for fi29 := 0; fi29 < 3; fi29++ {
            fmt.Println(v26, v5)
            v30 := GenS6{F0: 255, F1: 0, F2: v5}
            _ = fi29
            _ = v30
        }
    }
    fmt.Println(v5)
    switch "a" {
    case "foo bar":
        for fi31 := 0; fi31 < 9; fi31++ {
            var v32 bool = false
            var v33 uint8 = 255
            _ = fi31
            _ = v32
            _ = v33
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
