package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 int16 }
func genFn0(p0 int32, p1 int8, p2 string) int {
    p0 = (p0 << ((65535) & 63))
    p1 = (p1 | p1)
    if ((true || false) || true) {
        v5 := GenS1{F0: 32768, F1: false, F2: 1000}
        _ = v5
    } else {
        v6 := GenS1{F0: 9223372036854775807, F1: (false || ((255 != 42) && (!(false && (true && (!false)))))), F2: 798}
        p1 -= p1
        for fi7 := 0; fi7 < 10; fi7++ {
            switch ("hello" + p2) {
            case "":
                v8 := []int{fi7, (fi7 * fi7), fi7}
                _ = v8
            default:
                if (1 == 9223372036854775807) {
                    v9 := []int{-9223372036854775808, 4294967295, (fi7 << ((678) & 63)), (fi7 ^ 211)}
                    v9 = append(v9, fi7)
                    if len(v9) > 0 {
                        v9[0] = -1
                    }
                    _ = v9
                } else {
                    var v10 bool = (fi7 == fi7)
                    _ = v10
                }
            }
            switch p2 {
            case "foo bar":
                if (!(0 > -164)) {
                    p1 -= p1
                } else {
                    v11 := []int{-136}
                    if (2147483647 < fi7) {
                        for fi12 := 0; fi12 < 19; fi12++ {
                            var v13 float32 = 1e10
                            _ = fi12
                            _ = v13
                        }
                    }
                    for fi14 := 0; fi14 < 13; fi14++ {
                        p0 = 978
                        var v15 uint = 18446744073709551615
                        _ = fi14
                        _ = v15
                    }
                    _ = v11
                }
            default:
                if true {
                    if false {
                        if (true || (false && (!true))) {
                            var v16 bool = false
                            v17 := []int{9223372036854775807, (fi7 + (fi7 + fi7)), fi7}
                            p0 -= p0
                            _ = v16
                            _ = v17
                        } else {
                            var v18 uint = 255
                            p0 += (p0 << ((18446744073709551615) & 63))
                            _ = v18
                        }
                        var v19 int = fi7
                        var v20 uint32 = 148
                        _ = v19
                        _ = v20
                    } else {
                        var v21 float64 = 1e10
                        v22 := GenS1{F0: (fi7 & fi7), F1: (((p0 | p0) < p0) && (1 == 1)), F2: 0}
                        _ = v21
                        _ = v22
                    }
                    if (760 <= p0) {
                        p0 -= p0
                    }
                    var v23 uint32 = 24
                    _ = v23
                }
            }
            p2 = p2
            _ = fi7
        }
        _ = v6
    }
    var v24 float64 = 1e10
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    return 128
}

func main() {
    v1 := []string{"hello", "hello", ("a" + ("foo bar" + (("hello" + "\n\t") + "foo bar")))}
    v2 := make(map[int]string)
    var v3 int = -1
    v4 := make(map[uint32]bool)
    if false {
        fmt.Println(v3)
    }
    var v25 int64 = 953
    fmt.Println(v3)
    var v26 uint16 = 255
    var v27 uint16 = 0
    var v28 bool = true
    switch v3 {
    case 641:
        fmt.Println(v3, v28, v25)
    case 133:
        fmt.Println(v3)
    case -845:
        var v29 bool = v28
        _ = v29
    default:
        var v30 uint8 = 253
        _ = v30
    }
    fmt.Println(v25)
    v31 := GenS1{F0: v3, F1: (1 != 1000), F2: 1}
    fmt.Println(v3, v27, v3)
    v25 -= 1
    var v32 uint64 = 42
    v27 = (v27 * (v26 * (v26 / ((v27) | 1))))
    var v33 uint = 319
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
}
