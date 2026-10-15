package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 uint16; F1 int64; F2 uint32 }
type GenS3 struct { F0 int; F1 bool }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]int)
    v3 := make(map[int]string)
    v4 := GenS1{F0: 0, F1: 1}
    var v5 int = -2147483648
    var v6 uint64 = 2
    var v7 int16 = 1000
    var v8 string = (((("\n\t" + "hello") + ("\n\t" + "foo bar")) + "a") + ((("a" + "foo bar") + ("" + "")) + "\n\t"))
    if false {
        v9 := GenS1{F0: 1, F1: 231}
        _ = v9
    } else {
        switch v5 {
        case 0:
            for fi10 := 0; fi10 < 18; fi10++ {
                switch v8 {
                case "hello":
                    for fi11 := 0; fi11 < 3; fi11++ {
                        v12 := v3[v5]
                        _ = fi11
                        _ = v12
                    }
                case "a":
                    v5 -= v5
                case "":
                    switch (-(v5)) {
                    case 32768:
                        v7 -= (v7 / ((v7) | 1))
                    case -708:
                        switch "a" {
                        case "a":
                            if ((0.0 / 1.0) >= (1e10 + ((((0.1 * (-2.25 + 1e10)) / (1.5 * float64(1.5))) + 1.5) - (1e10 + 0.1)))) {
                                v13 := GenS2{F0: 795, F1: -953, F2: 2147483648}
                                for fi14 := 0; fi14 < 4; fi14++ {
                                    v15 := GenS3{F0: fi14, F1: true}
                                    if true {
                                        var v16 int64 = -9223372036854775808
                                        var v17 int8 = -128
                                        _ = v16
                                        _ = v17
                                    } else {
                                        v8 = "a"
                                        var v18 bool = false
                                        _ = v18
                                    }
                                    if (3.14159 <= (-2.25 - 0.0)) {
                                        if (0.1 < 0.1) {
                                            v6 += v6
                                        }
                                        v5 = (fi10 | -242)
                                    } else {
                                        fmt.Println(v8, fi10, v6)
                                        var v19 float32 = 3.14159
                                        v20 := GenS3{F0: v5, F1: (!((100 <= 127) || false))}
                                        _ = v19
                                        _ = v20
                                    }
                                    _ = fi14
                                    _ = v15
                                }
                                v21 := []int{32768, -2147483648, fi10, -779}
                                _ = v13
                                _ = v21
                            }
                        default:
                            var v22 int32 = 1
                            _ = v22
                        }
                    }
                }
                _ = fi10
            }
        case 42:
            v7 = (v7 / ((v7) | 1))
        default:
            var v23 float64 = 1.5
            _ = v23
        }
    }
    v24 := []int{1}
    var v25 uint8 = 115
    v3[v5] = (v8 + ("foo bar" + v8))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
