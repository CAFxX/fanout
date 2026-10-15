package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 uint16 }
func genFn0() uint8 {
    v5 := GenS1{F0: 18446744073709551615, F1: "a"}
    var v6 int16 = 0
    v6 = v6
    _ = v5
    _ = v6
    return 199
}
func genFn1(p0 int8) uint64 {
    v7 := GenS2{F0: 255}
    if (true && true) {
        var v8 uint8 = 188
        if false {
            v9 := []int{1, 4294967295, -2147483648, 32768}
            for fi10 := 0; fi10 < 17; fi10++ {
                p0 = p0
                for fi11 := 0; fi11 < 13; fi11++ {
                    switch 4294967295 {
                    case 4294967295:
                        switch fi10 {
                        case 127:
                            if true {
                                switch (^(fi11)) {
                                case 857:
                                    var v12 bool = false
                                    _ = v12
                                case 124:
                                    v8 = 127
                                case 0:
                                    var v13 float32 = 1.5
                                    _ = v13
                                default:
                                    v14 := GenS1{F0: 653, F1: ("\n\t" + ("hello" + "a"))}
                                    _ = v14
                                }
                            }
                        case 2147483647:
                            p0 -= -1
                        case 0:
                            if len(v9) > 0 {
                                v9[0] = (fi11 & 32768)
                            }
                        default:
                            p0 += -128
                        }
                    case 768:
                        fmt.Println(p0, p0)
                    default:
                        var v15 int = fi10
                        _ = v15
                    }
                    _ = len(v9)
                    _ = fi11
                }
                _ = fi10
            }
            var v16 string = "a"
            _ = v16
            _ = v9
        } else {
            v17 := []int{int(p0)}
            _ = v17
        }
        var v18 uint16 = 2
        _ = v18
        _ = v8
    }
    v19 := []int{1, -1, 0, 985}
    v20 := GenS1{F0: 0, F1: "a"}
    _ = p0
    _ = v19
    _ = v20
    _ = v7
    return 48
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 float64 = (0.1 * 0.0)
    v3 := make(map[string]int)
    v4 := make(map[uint32]uint32)
    for fi21 := 0; fi21 < 12; fi21++ {
        v22 := GenS1{F0: 660, F1: "\n\t"}
        v2 += (v2 * (1e10 / 1.5))
        _ = fi21
        _ = v22
    }
    switch "\n\t" {
    case "hello":
        var v23 uint16 = 30
        _ = v23
    default:
        v24 := v3["a"]
        _ = v24
    }
    if ((!true) && (2 <= 42)) {
        var v25 int16 = -32768
        v25 = v25
        for fi26 := 0; fi26 < 9; fi26++ {
            v25 = 1000
            _ = fi26
        }
        _ = v25
    }
    v27 := []int{}
    v28 := GenS2{F0: 255}
    var v29 uint32 = 679
    v27 = append(v27, 411)
    _ = len(v27)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
