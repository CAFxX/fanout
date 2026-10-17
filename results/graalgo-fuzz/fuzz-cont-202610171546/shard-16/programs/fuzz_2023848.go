package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 bool; F1 string }

func main() {
    v1 := make(map[uint32]string)
    var v2 int64 = 0
    var v3 float32 = (1.5 * ((-2.25 - (1.5 * 1e10)) - 1.5))
    var v4 int32 = -2147483648
    v5 := []int{-649, 9223372036854775807, 0, 1}
    if len(v5) > 0 {
        v5[0] = 0
    }
    v6 := GenS1{F0: 0, F1: "\n\t"}
    for fi7 := 0; fi7 < 16; fi7++ {
        if (!(((v4 & v4) <= 1) && true)) {
            switch fi7 {
            case -1:
                for fi8 := 0; fi8 < 17; fi8++ {
                    var v9 int = (^(fi7))
                    v3 = v3
                    v2 -= (v2 >> ((255) & 63))
                    _ = fi8
                    _ = v9
                }
            case 128:
                if (v2 <= (v2 * v2)) {
                    for fi10 := 0; fi10 < 9; fi10++ {
                        v2 -= -9223372036854775808
                        var v11 uint = 663
                        switch 4294967295 {
                        case 2147483647:
                            v12 := GenS1{F0: v11, F1: ""}
                            _ = v12
                        case 1:
                            var v13 int = (fi10 & (fi7 - (^(fi10))))
                            _ = v13
                        default:
                            fmt.Println(fi7)
                        }
                        _ = fi10
                        _ = v11
                    }
                    v4 += v4
                }
            default:
                fmt.Println(fi7, fi7, fi7)
            }
            v14 := v1[0]
            _ = v14
        }
        var v15 string = ("hello" + "\n\t")
        for fi16 := 0; fi16 < 19; fi16++ {
            fmt.Println(v15)
            _ = fi16
        }
        _ = fi7
        _ = v15
    }
    v17 := GenS2{F0: false, F1: ((("\n\t" + ("" + "")) + "a") + "a")}
    v18 := v1[721]
    var v19 float32 = v3
    v19 -= 0.0
    if false {
        var v20 string = v18
        for fi21 := 0; fi21 < 17; fi21++ {
            switch -1 {
            case 42:
                v22 := v1[4294967295]
                _ = v22
            case -2147483648:
                for fi23 := 0; fi23 < 15; fi23++ {
                    var v24 int64 = (v2 ^ -9223372036854775808)
                    _ = fi23
                    _ = v24
                }
            case 1:
                var v25 string = "\n\t"
                _ = v25
            }
            _ = fi21
        }
        for fi26 := 0; fi26 < 18; fi26++ {
            v2 -= v2
            if (42 > v2) {
                v3 -= float32(1e10)
                v27 := GenS1{F0: 126, F1: v20}
                _ = v27
            }
            _ = fi26
        }
        _ = v20
    }
    if (335 == 1) {
        v3 = 0.0
        var v28 uint64 = 0
        _ = v28
    }
    v29 := GenS2{F0: (-1 != 32767), F1: ""}
    switch "a" {
    case "":
        v30 := GenS2{F0: (true && true), F1: "\n\t"}
        _ = v30
    case "\n\t":
        var v31 uint32 = 4294967295
        _ = v31
    }
    _ = len(v1)
    var v32 float64 = (0.0 - 1e10)
    var v33 int64 = 42
    v18 = "hello"
    if len(v5) > 0 {
        v5[0] = 516
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
