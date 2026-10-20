package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool }
func genFn0() int16 {
    v6 := []int{128, 0}
    var v7 float32 = ((-2.25 * float32(-2.25)) / 1.0)
    var v8 bool = (false && (312 != -1))
    var v9 float32 = v7
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 32767
}

func main() {
    v1 := []uint32{147}
    v2 := []int{9223372036854775807, 995}
    v3 := []string{(("foo bar" + "hello") + "hello"), "", "a"}
    var v4 uint8 = 1
    var v5 int32 = 109
    if false {
        v10 := GenS1{F0: 4294967295}
        v4 -= (v4 + (v4 >> ((4294967295) & 63)))
        switch "foo bar" {
        case "a":
            var v11 uint32 = 0
            _ = v11
        case "":
            v5 = (v5 + (v5 << ((255) & 63)))
        }
        _ = v10
    }
    v12 := GenS2{F0: (0 < 1)}
    fmt.Println(v4, v5)
    v13 := GenS1{F0: -693}
    for fi14 := 0; fi14 < 4; fi14++ {
        _ = len(v2)
        switch fi14 {
        case -2147483648:
            v4 += (^(v4))
        case 4294967295:
            if (-11 < 42) {
                v5 = (v5 ^ (v5 | (v5 | v5)))
                switch fi14 {
                case -556:
                    v4 -= (v4 >> ((1) & 63))
                case 32767:
                    if false {
                        var v15 uint16 = 33
                        if len(v2) > 0 {
                            v2[0] = (^(fi14))
                        }
                        _ = v15
                    }
                case -681:
                    for fi16 := 0; fi16 < 16; fi16++ {
                        var v17 float64 = (0.1 * -2.25)
                        var v18 string = "a"
                        var v19 int16 = -32768
                        _ = fi16
                        _ = v17
                        _ = v18
                        _ = v19
                    }
                }
            } else {
                if ((uint64(v4) != 4294967295) && false) {
                    var v20 uint64 = 659
                    switch (fi14 + fi14) {
                    case 9223372036854775807:
                        if (42 <= 1) {
                            var v21 int = (fi14 ^ fi14)
                            v22 := GenS1{F0: 32767}
                            _ = v21
                            _ = v22
                        } else {
                            fmt.Println(v5, fi14)
                            v3 = append(v3, "")
                            for fi23 := 0; fi23 < 9; fi23++ {
                                var v24 float32 = -2.25
                                var v25 uint16 = 32767
                                _ = fi23
                                _ = v24
                                _ = v25
                            }
                        }
                    case 0:
                        v20 = 439
                    default:
                        switch "\n\t" {
                        case "hello":
                            if false {
                                v2 = append(v2, (fi14 >> ((433) & 63)))
                            }
                        default:
                            var v26 uint = 42
                            _ = v26
                        }
                    }
                    _ = v20
                } else {
                    var v27 uint64 = 1
                    _ = v27
                }
            }
        default:
            var v28 float32 = 1.5
            _ = v28
        }
        _ = fi14
    }
    for fi29 := 0; fi29 < 2; fi29++ {
        if (244 > 2) {
            v4 = (^(v4))
            var v30 uint64 = 2
            _ = v30
        } else {
            v31 := GenS2{F0: (true || true)}
            v32 := GenS2{F0: false}
            _ = v31
            _ = v32
        }
        v33 := GenS2{F0: true}
        _ = fi29
        _ = v33
    }
    fmt.Println(v4)
    v5 += (^(v5))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
