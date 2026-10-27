package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 string; F1 int32; F2 int16 }
func genFn0(p0 uint32, p1 string) uint64 {
    var v7 int = 4294967295
    switch ("" + "foo bar") {
    case "foo bar":
        switch 32768 {
        case 128:
            for fi8 := 0; fi8 < 9; fi8++ {
                p0 = (p0 ^ (p0 - (p0 - (p0 / ((594) | 1)))))
                for fi9 := 0; fi9 < 5; fi9++ {
                    if true {
                        p0 -= (p0 * p0)
                        v10 := []int{v7, (fi8 % (((^(fi9))) | 1))}
                        var v11 int32 = -1
                        _ = v10
                        _ = v11
                    } else {
                        v7 = (fi9 << ((65535) & 63))
                    }
                    v12 := GenS1{F0: int8(v7), F1: 65535}
                    var v13 uint16 = 390
                    p1 = (((p1 + "foo bar") + ("" + "a")) + ("\n\t" + (("a" + p1) + ("hello" + p1))))
                    v14 := GenS1{F0: 0, F1: v13}
                    v15 := GenS1{F0: -128, F1: v13}
                    v16 := GenS1{F0: 1, F1: v13}
                    _ = fi9
                    _ = v12
                    _ = v13
                    _ = v14
                    _ = v15
                    _ = v16
                }
                fmt.Println(p0, fi8, p1)
                _ = fi8
            }
        }
    case "":
        v17 := []int{(v7 & (v7 << ((255) & 63)))}
        _ = v17
    default:
        v7 -= v7
    }
    for fi18 := 0; fi18 < 18; fi18++ {
        v7 -= v7
        if true {
            if false {
                p0 -= (p0 & (p0 >> ((4294967295) & 63)))
            }
            v19 := []int{}
            v20 := GenS1{F0: int8(p0), F1: 453}
            _ = v19
            _ = v20
        } else {
            var v21 int32 = 42
            _ = v21
        }
        _ = fi18
    }
    _ = p0
    _ = p1
    _ = v7
    return 42
}
func genFn1(p0 int32, p1 int32, p2 bool) uint8 {
    p2 = p2
    switch ("\n\t" + "") {
    case "a":
        p1 -= (p0 - (^(p1)))
    case "hello":
        var v22 uint32 = 0
        _ = v22
    case "\n\t":
        v23 := []int{-1}
        _ = v23
    default:
        v24 := GenS2{F0: "", F1: p1, F2: 42}
        _ = v24
    }
    if (!true) {
        p1 += -1
        var v25 int8 = -1
        _ = v25
    } else {
        v26 := []int{127, -940}
        switch 349 {
        case -405:
            v26 = append(v26, 173)
        case 376:
            v26 = append(v26, 2147483647)
        default:
            p1 = p0
        }
        if ((1.5 / (float64(0.0) + 0.1)) < 1.5) {
            for fi27 := 0; fi27 < 5; fi27++ {
                var v28 uint = 18446744073709551615
                fmt.Println(fi27, fi27)
                _ = fi27
                _ = v28
            }
        }
        _ = v26
    }
    _ = p0
    _ = p1
    _ = p2
    return 127
}

func main() {
    var v1 int16 = 362
    var v2 string = "foo bar"
    v3 := make(map[int]string)
    v4 := make(map[uint32]string)
    var v5 uint16 = 65535
    var v6 bool = (255 >= 127)
    var v29 int32 = -1
    v1 -= (v1 ^ v1)
    v30 := []int{32768, 9223372036854775807, 4294967295, 42}
    var v31 uint32 = 2147483648
    v31 = v31
    v2 = v2
    switch 9223372036854775807 {
    case 4294967295:
        v5 -= (v5 | (^(v5)))
    case 42:
        v31 = v31
    default:
        v3[2147483647] = "\n\t"
    }
    if len(v30) > 0 {
        v30[0] = 0
    }
    v6 = false
    var v32 float32 = 1e10
    if v6 {
        var v33 uint16 = v5
        delete(v4, 234)
        if (100 < 44) {
            if len(v30) > 0 {
                v30[0] = 32768
            }
        }
        _ = v33
    }
    if len(v30) >= 2 {
        v30 = v30[1:len(v30)]
    }
    fmt.Println(v29, v2)
    var v34 uint64 = genFn0(v31, (v2 + v2))
    v2 = ""
    v35 := GenS1{F0: 109, F1: (v5 % ((255) | 1))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
