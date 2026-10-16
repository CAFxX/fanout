package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 uint8 }
type GenS2 struct { F0 uint64; F1 uint8; F2 string }
func genFn0() uint32 {
    var v4 int8 = -28
    var v5 uint32 = 42
    _ = v4
    _ = v5
    return 2147483648
}
func genFn1(p0 uint32, p1 int8) uint8 {
    for fi6 := 0; fi6 < 9; fi6++ {
        if (255 < 0) {
            v7 := GenS1{F0: -32768, F1: 280, F2: 1}
            _ = v7
        }
        _ = fi6
    }
    if false {
        switch 330 {
        case -1:
            v8 := []int{9223372036854775807, 32767}
            _ = v8
        default:
            if (((9223372036854775807 == -692) && false) || true) {
                var v9 int8 = p1
                if true {
                    var v10 uint8 = 255
                    _ = v10
                } else {
                    for fi11 := 0; fi11 < 17; fi11++ {
                        if (uint8(p0) != 1) {
                            fmt.Println(p1, v9)
                        } else {
                            if false {
                                var v12 string = "a"
                                _ = v12
                            } else {
                                var v13 int32 = 42
                                v14 := GenS1{F0: -1, F1: -1, F2: 128}
                                var v15 string = "hello"
                                _ = v13
                                _ = v14
                                _ = v15
                            }
                            var v16 int = 128
                            _ = v16
                        }
                        v9 = (v9 & 42)
                        _ = fi11
                    }
                }
                _ = v9
            } else {
                var v17 float64 = 1e10
                _ = v17
            }
        }
        var v18 uint = 65535
        if true {
            v18 = (v18 | (^(v18)))
        } else {
            if false {
                switch 1 {
                case 32767:
                    fmt.Println(v18)
                case 32768:
                    var v19 uint = v18
                    _ = v19
                }
                var v20 float32 = 1e10
                var v21 bool = ((!true) || false)
                _ = v20
                _ = v21
            }
            v22 := []int{-975}
            _ = v22
        }
        _ = v18
    }
    var v23 string = ("a" + "\n\t")
    var v24 uint16 = 467
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v23
    _ = v24
    return 255
}

func main() {
    var v1 uint16 = 42
    var v2 bool = (true && true)
    v3 := make(map[int]int)
    switch 9223372036854775807 {
    case 260:
        v3[0] = -631
    case -534:
        v25 := GenS1{F0: -32768, F1: -743, F2: 185}
        _ = v25
    }
    _ = len(v3)
    switch "a" {
    case "a":
        v1 = (v1 * v1)
    }
    v2 = v2
    fmt.Println(v1, v2)
    v26 := []int{849, 765, 9223372036854775807}
    _ = len(v26)
    v27 := GenS2{F0: 18446744073709551615, F1: 42, F2: "a"}
    v3[2147483647] = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
}
