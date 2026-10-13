package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 uint8 }
type GenS2 struct { F0 int; F1 uint8 }
func genFn0(p0 uint32, p1 bool) uint64 {
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 int64 = -1
        _ = fi7
        _ = v8
    }
    v9 := []int{-321, 810, -538}
    p1 = (!p1)
    var v10 bool = (p1 || true)
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 18446744073709551615
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []uint32{284, 633, 42, 497}
    var v3 uint32 = 0
    var v4 int8 = 0
    var v5 int16 = -1
    var v6 bool = (!false)
    switch ("foo bar" + "") {
    case "foo bar":
        delete(v1, 32767)
    case "":
        fmt.Println(v5, v6, v5)
    case "\n\t":
        var v11 uint8 = uint8(v3)
        _ = v11
    default:
        if v6 {
            if (-9223372036854775808 == -9223372036854775808) {
                var v12 uint32 = v3
                for fi13 := 0; fi13 < 5; fi13++ {
                    for fi14 := 0; fi14 < 3; fi14++ {
                        var v15 bool = v6
                        if len(v2) > 0 {
                            v2[0] = (v3 * v3)
                        }
                        _ = fi14
                        _ = v15
                    }
                    _ = fi13
                }
                if (((((((2 != 1) && ((0.1 + 0.1) >= (1.5 + float32((-2.25 - 1e10))))) || false) || ((v5 / (((v5 | (v5 | 32767))) | 1)) == v5)) || (uint(v12) != 42)) || (!v6)) || true) {
                    v16 := GenS1{F0: v12, F1: ("foo bar" == "\n\t"), F2: 1}
                    v2 = append(v2, v12)
                    _ = v16
                } else {
                    v12 = 2
                    switch 127 {
                    case 32768:
                        var v17 uint32 = (^(v12))
                        _ = v17
                    }
                }
                _ = v12
            }
            v4 = (-(v4))
        } else {
            if (468 != 604) {
                for fi18 := 0; fi18 < 3; fi18++ {
                    var v19 bool = v6
                    v4 += (-(v4))
                    v5 += v5
                    _ = fi18
                    _ = v19
                }
                var v20 int = 127
                v3 += v3
                _ = v20
            } else {
                for fi21 := 0; fi21 < 18; fi21++ {
                    delete(v1, fi21)
                    fmt.Println(fi21, fi21, v4)
                    var v22 uint32 = (^(v3))
                    _ = fi21
                    _ = v22
                }
                var v23 uint16 = 65535
                _ = v23
            }
            var v24 int32 = 2147483647
            var v25 float32 = 0.0
            _ = v24
            _ = v25
        }
    }
    var v26 int64 = 654
    var v27 int64 = 1
    var v28 uint64 = 0
    var v29 int8 = (v4 + v4)
    for fi30 := 0; fi30 < 4; fi30++ {
        var v31 float64 = 3.14159
        v32 := GenS1{F0: 322, F1: (-200 <= -1), F2: 190}
        _ = fi30
        _ = v31
        _ = v32
    }
    _ = len(v1)
    v33 := GenS2{F0: 42, F1: uint8(v28)}
    if true {
        fmt.Println(v26, v26, v5)
        var v34 float32 = (1.5 - (-2.25 * 0.1))
        fmt.Println(v6, v3)
        _ = v34
    } else {
        if v6 {
            v6 = ((!true) && (("foo bar" > ("hello" + "\n\t")) || v6))
            for fi35 := 0; fi35 < 2; fi35++ {
                switch int(v28) {
                case 1:
                    v26 -= (v27 / (((v27 | int64(v5))) | 1))
                default:
                    var v36 int = fi35
                    _ = v36
                }
                _ = fi35
            }
        }
        var v37 int32 = -134
        _ = v37
    }
    var v38 int8 = -128
    v26 = (-(v27))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v33
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
