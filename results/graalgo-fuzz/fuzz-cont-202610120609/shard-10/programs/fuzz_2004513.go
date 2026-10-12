package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint8 }
func genFn0(p0 uint8, p1 int16, p2 uint) uint8 {
    v6 := []int{2147483647}
    fmt.Println(p0, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return p0
}

func main() {
    var v1 uint = 4294967295
    var v2 string = ("a" + "a")
    var v3 uint64 = 18446744073709551615
    var v4 int = -272
    var v5 int32 = 786
    for fi7 := 0; fi7 < 16; fi7++ {
        switch fi7 {
        case 9223372036854775807:
            v8 := []int{}
            _ = v8
        default:
            v9 := []int{int(v1), fi7}
            _ = v9
        }
        v3 = 0
        _ = fi7
    }
    v3 = 0
    var v10 uint16 = 32767
    for fi11 := 0; fi11 < 16; fi11++ {
        v12 := []int{(fi11 - 2147483647)}
        if ((v1 / ((v1) | 1)) == 4294967295) {
            for fi13 := 0; fi13 < 18; fi13++ {
                v2 = ((v2 + (("a" + v2) + v2)) + ((((("\n\t" + v2) + ((v2 + "hello") + v2)) + v2) + v2) + ("\n\t" + (v2 + v2))))
                _ = fi13
            }
            var v14 int16 = -32768
            _ = v14
        } else {
            if (1e10 > -2.25) {
                fmt.Println(fi11, v1, v3)
                var v15 uint32 = 951
                var v16 uint16 = (v10 ^ v10)
                _ = v15
                _ = v16
            } else {
                fmt.Println(v10, v4, v1)
                var v17 float64 = ((((0.0 + ((1.5 + -2.25) + (0.0 * (0.1 + 1e10)))) + float64((0.0 + 0.1))) / 3.14159) - (1e10 + 0.1))
                v2 = ""
                _ = v17
            }
            if false {
                var v18 int = fi11
                var v19 int = (fi11 * -640)
                var v20 uint = v1
                _ = v18
                _ = v19
                _ = v20
            } else {
                _ = len(v12)
            }
        }
        v21 := GenS1{F0: 2, F1: v5, F2: -128}
        _ = fi11
        _ = v12
        _ = v21
    }
    v22 := []int{v4, v4}
    var v23 int16 = 1000
    var v24 float64 = 0.1
    v25 := GenS1{F0: 128, F1: (v5 + 1), F2: int8(v4)}
    if (v2 > (v2 + (v2 + "\n\t"))) {
        fmt.Println(v24, v24, v23)
        switch v4 {
        case -890:
            v22 = append(v22, (^(v4)))
        case 42:
            v26 := GenS1{F0: 255, F1: -2147483648, F2: 120}
            _ = v26
        case 485:
            fmt.Println(v2)
        default:
            v27 := GenS2{F0: 1}
            _ = v27
        }
    } else {
        for fi28 := 0; fi28 < 3; fi28++ {
            if len(v22) > 0 {
                v22[0] = (fi28 << ((v1) & 63))
            }
            _ = fi28
        }
    }
    if (!true) {
        var v29 uint = (v1 % ((0) | 1))
        v29 += (v29 & (v29 << ((2) & 63)))
        v1 = uint(v10)
        _ = v29
    } else {
        v30 := GenS1{F0: genFn0(0, v23, (v1 >> (((v1 << ((v1) & 63))) & 63))), F1: v5, F2: -54}
        fmt.Println(v4, v1)
        _ = v30
    }
    for fi31 := 0; fi31 < 10; fi31++ {
        v32 := GenS2{F0: (^(v3))}
        v1 = v1
        _ = fi31
        _ = v32
    }
    fmt.Println(v2, v10, v2)
    var v33 uint64 = uint64(v10)
    fmt.Println(v5)
    for fi34 := 0; fi34 < 7; fi34++ {
        fmt.Println(fi34, fi34)
        if false {
            switch 127 {
            case 1:
                v23 -= v23
            case -9223372036854775808:
                var v35 int16 = (v23 - 1)
                _ = v35
            default:
                if (255 == (v1 - (v1 | v1))) {
                    if true {
                        var v36 string = v2
                        v5 += -1
                        v37 := GenS2{F0: v33}
                        _ = v36
                        _ = v37
                    }
                    if true {
                        var v38 int8 = -128
                        _ = v38
                    }
                } else {
                    var v39 int32 = -1
                    _ = v39
                }
            }
        } else {
            fmt.Println(v23, v10, v2)
            var v40 float64 = float64((-2.25 / 3.14159))
            v41 := GenS3{F0: v10}
            _ = v40
            _ = v41
        }
        if len(v22) >= 2 {
            v22 = v22[1:len(v22)]
        }
        _ = fi34
    }
    v3 = (v33 + v33)
    var v42 float32 = (-2.25 / (-2.25 / float32(0.1)))
    v43 := GenS4{F0: 31}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v33
    _ = v4
    _ = v42
    _ = v43
    _ = v5
}
