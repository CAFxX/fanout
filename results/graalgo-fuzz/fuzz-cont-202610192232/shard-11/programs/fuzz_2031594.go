package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 bool }
func genFn0(p0 int16, p1 int) bool {
    if true {
        p1 = (^(p1))
        p1 += (p1 / ((p1) | 1))
    }
    var v7 int16 = p0
    fmt.Println(v7, p0)
    _ = p0
    _ = p1
    _ = v7
    return true
}
func genFn1() int {
    v8 := []int{127, 32767, 4294967295}
    _ = len(v8)
    if (2147483647 >= 0) {
        if (511 > 383) {
            for fi9 := 0; fi9 < 17; fi9++ {
                var v10 string = (("\n\t" + "foo bar") + "a")
                _ = fi9
                _ = v10
            }
        } else {
            if ("\n\t" == "foo bar") {
                var v11 int16 = 1
                v11 += (v11 | 0)
                v8 = append(v8, 128)
                _ = v11
            } else {
                var v12 uint = 4294967295
                switch ("" + "foo bar") {
                case "hello":
                    if true {
                        var v13 uint8 = 192
                        _ = v13
                    } else {
                        var v14 uint16 = 1
                        switch 9223372036854775807 {
                        case 74:
                            var v15 float32 = ((1.5 / -2.25) - 0.0)
                            _ = v15
                        case 127:
                            v14 -= 1
                        case 0:
                            var v16 int16 = -32768
                            _ = v16
                        default:
                            var v17 int8 = -128
                            _ = v17
                        }
                        _ = v14
                    }
                case "\n\t":
                    fmt.Println(v12)
                default:
                    for fi18 := 0; fi18 < 13; fi18++ {
                        if false {
                            var v19 int16 = 32767
                            var v20 uint = 255
                            v20 = 25
                            _ = v19
                            _ = v20
                        }
                        v12 += v12
                        _ = fi18
                    }
                }
                var v21 int = -1
                _ = v12
                _ = v21
            }
        }
    }
    _ = v8
    return 32767
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]int)
    v3 := make(map[string]int)
    v4 := make(map[string]int)
    var v5 float64 = 0.0
    v6 := make(map[uint32]int)
    v5 -= 0.1
    v22 := []int{-576, 32768}
    v5 += ((v5 / ((v5 / v5) - 0.1)) * 3.14159)
    var v23 uint8 = 246
    v24 := GenS1{F0: 1, F1: -128, F2: false}
    v25 := GenS1{F0: -1, F1: -1, F2: (!(42 > 32767))}
    if true {
        for fi26 := 0; fi26 < 4; fi26++ {
            _ = len(v22)
            delete(v3, ("" + "foo bar"))
            for fi27 := 0; fi27 < 19; fi27++ {
                var v28 bool = true
                var v29 int32 = 2147483647
                if v28 {
                    var v30 uint = 1
                    if (!v28) {
                        v5 += (v5 + (0.1 * v5))
                    }
                    if len(v22) > 0 {
                        v22[0] = fi26
                    }
                    _ = v30
                } else {
                    v22 = append(v22, (fi26 << ((2) & 63)))
                    var v31 int = (fi26 << ((514) & 63))
                    _ = v31
                }
                fmt.Println(fi26, fi27)
                _ = fi27
                _ = v28
                _ = v29
            }
            _ = fi26
        }
        var v32 int16 = 1
        fmt.Println(v32, v5, v32)
        _ = v32
    }
    fmt.Println(v23, v23)
    v23 = (v23 + v23)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
