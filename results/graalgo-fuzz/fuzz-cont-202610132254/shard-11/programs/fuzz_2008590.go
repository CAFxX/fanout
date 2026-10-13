package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 int64, p1 int16, p2 uint) uint32 {
    for fi6 := 0; fi6 < 15; fi6++ {
        for fi7 := 0; fi7 < 7; fi7++ {
            if (!(2 < (p2 | (p2 * p2)))) {
                var v8 bool = false
                _ = v8
            }
            var v9 float64 = 1.5
            _ = fi7
            _ = v9
        }
        v10 := []int{fi6}
        if len(v10) > 0 {
            v10[0] = (^(fi6))
        }
        _ = fi6
        _ = v10
    }
    p2 -= p2
    fmt.Println(p0, p2)
    var v11 int8 = 0
    if ((!false) && (!(688 > 42))) {
        fmt.Println(p2, p1, p0)
        p2 = (p2 >> (((^(p2))) & 63))
        for fi12 := 0; fi12 < 14; fi12++ {
            if (2 <= 42) {
                switch 308 {
                case 0:
                    fmt.Println(v11, p1)
                case 1:
                    for fi13 := 0; fi13 < 9; fi13++ {
                        switch ("hello" + "hello") {
                        case "a":
                            if false {
                                p1 = (p1 | (^(p1)))
                                v14 := []int{int(p0), 128, (fi12 / ((623) | 1))}
                                _ = v14
                            } else {
                                fmt.Println(p1, p2)
                                var v15 uint16 = 65535
                                _ = v15
                            }
                        case "foo bar":
                            var v16 uint16 = 1
                            _ = v16
                        case "\n\t":
                            v11 -= (v11 >> ((64) & 63))
                        default:
                            var v17 float32 = -2.25
                            _ = v17
                        }
                        _ = fi13
                    }
                case 469:
                    if (v11 != 127) {
                        fmt.Println(p1)
                        p1 -= p1
                    }
                default:
                    v18 := []int{(fi12 * (fi12 - 32768)), (fi12 - fi12), fi12}
                    _ = v18
                }
                v19 := []int{fi12, 32768, fi12}
                var v20 uint = p2
                _ = v19
                _ = v20
            }
            var v21 float64 = 1.5
            _ = fi12
            _ = v21
        }
    } else {
        v22 := []int{9223372036854775807, -889, 32768, -9223372036854775808}
        _ = v22
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 220
}

func main() {
    var v1 string = "hello"
    v2 := make(map[uint32]bool)
    var v3 float32 = 3.14159
    v4 := make(map[int]string)
    v5 := make(map[string]int)
    var v23 int16 = -1
    var v24 uint64 = 0
    if true {
        for fi25 := 0; fi25 < 4; fi25++ {
            for fi26 := 0; fi26 < 18; fi26++ {
                v27 := []int{-9223372036854775808, fi26, (fi25 / ((fi26) | 1))}
                v28 := v4[2147483647]
                v29 := GenS1{F0: v1}
                _ = fi26
                _ = v27
                _ = v28
                _ = v29
            }
            v23 = (v23 / (((v23 / ((1) | 1))) | 1))
            _ = fi25
        }
    } else {
        v30 := v2[674]
        var v31 int8 = 100
        _ = v30
        _ = v31
    }
    var v32 float64 = 3.14159
    v24 = 18446744073709551615
    var v33 int8 = 42
    v1 = ("hello" + v1)
    v2[1] = true
    var v34 uint16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
