package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 uint64; F2 int32 }
type GenS3 struct { F0 uint }
func genFn0() uint64 {
    var v6 bool = (1 >= 1)
    for fi7 := 0; fi7 < 19; fi7++ {
        v8 := GenS1{F0: 823}
        v9 := GenS1{F0: 891}
        switch "foo bar" {
        case "foo bar":
            v10 := GenS2{F0: fi7, F1: 4294967295, F2: 2147483647}
            _ = v10
        case "a":
            v11 := []int{fi7, (fi7 ^ fi7), (^(fi7)), (fi7 ^ (fi7 + fi7))}
            _ = v11
        case "":
            v6 = (!v6)
        default:
            fmt.Println(fi7)
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    v12 := GenS1{F0: 4294967295}
    fmt.Println(v6)
    v6 = ((0.0 + float64((((1.5 / 0.1) * float32(1.5)) * 0.0))) != -2.25)
    _ = v12
    _ = v6
    return 18446744073709551615
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int32 = 1
    v3 := []uint32{0}
    var v4 string = "\n\t"
    var v5 float64 = 1.5
    fmt.Println(v2, v4, v4)
    for fi13 := 0; fi13 < 6; fi13++ {
        for fi14 := 0; fi14 < 4; fi14++ {
            v5 -= v5
            delete(v1, 0)
            switch (-(fi13)) {
            case 9223372036854775807:
                if (("" != "") || (37 != 255)) {
                    v15 := GenS2{F0: fi13, F1: 969, F2: (v2 | (v2 * (v2 ^ v2)))}
                    _ = len(v3)
                    var v16 uint64 = 42
                    _ = v15
                    _ = v16
                } else {
                    var v17 uint32 = 42
                    v17 = (^(v17))
                    switch -936 {
                    case -979:
                        if false {
                            v18 := GenS1{F0: genFn0()}
                            fmt.Println(v5, v5)
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = v18
                        }
                    case 96:
                        if (14 == 1) {
                            delete(v1, v17)
                        }
                    default:
                        fmt.Println(fi14, fi13)
                    }
                    _ = v17
                }
            case -274:
                fmt.Println(fi14, v4, fi14)
            default:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            var v19 uint8 = 245
            var v20 int32 = v2
            v21 := GenS2{F0: (^(fi14)), F1: genFn0(), F2: (v20 & v20)}
            v2 = 2147483647
            _ = fi14
            _ = v19
            _ = v20
            _ = v21
        }
        switch -2147483648 {
        case 42:
            fmt.Println(v4, v5, v5)
        case 0:
            v5 -= 0.1
        case 573:
            var v22 float32 = 1e10
            _ = v22
        }
        _ = fi13
    }
    fmt.Println(v4, v5, v5)
    v23 := GenS2{F0: 2147483647, F1: 4294967295, F2: -1}
    v5 = 0.1
    if false {
        v2 -= 792
        v24 := GenS1{F0: genFn0()}
        _ = v24
    } else {
        if ((v2 == 2147483647) || (!(!false))) {
            fmt.Println(v5)
        } else {
            for fi25 := 0; fi25 < 18; fi25++ {
                v26 := v1[2147483647]
                var v27 int8 = int8(v26)
                v28 := v1[232]
                _ = fi25
                _ = v26
                _ = v27
                _ = v28
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        fmt.Println(v4)
        v29 := GenS3{F0: 202}
        _ = v29
    }
    var v30 int8 = -128
    switch 32768 {
    case 4294967295:
        v3 = append(v3, 1)
    case 821:
        var v31 uint64 = 2
        _ = v31
    default:
        fmt.Println(v5, v5, v2)
    }
    fmt.Println(v4, v4)
    var v32 uint16 = 2
    var v33 int16 = 0
    var v34 int16 = (v33 % ((412) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v30
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
