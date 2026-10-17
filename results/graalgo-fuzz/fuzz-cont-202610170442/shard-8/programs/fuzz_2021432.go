package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 uint64; F1 uint16 }
type GenS4 struct { F0 uint16; F1 uint16; F2 uint64 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int16 }
func genFn0() string {
    for fi4 := 0; fi4 < 3; fi4++ {
        if (55 >= -25) {
            switch fi4 {
            case 520:
                var v5 int16 = -1
                _ = v5
            case 13:
                v6 := []int{(fi4 % ((fi4) | 1)), -642, 32767}
                _ = v6
            default:
                v7 := []int{fi4, fi4, (fi4 >> ((42) & 63))}
                _ = v7
            }
            var v8 uint8 = 1
            _ = v8
        }
        _ = fi4
    }
    var v9 bool = (746 < 18446744073709551615)
    var v10 uint = 42
    _ = v10
    _ = v9
    return "a"
}

func main() {
    v1 := []uint32{1, 1, 2147483647, 1}
    v2 := make(map[int]bool)
    var v3 int16 = -32768
    v11 := GenS1{F0: 2, F1: 612}
    v3 = (-(v3))
    var v12 uint8 = 127
    v3 = -903
    var v13 string = "\n\t"
    v14 := GenS2{F0: (v3 - 32767), F1: 899}
    v12 += (v12 % (((v12 + (v12 >> ((4294967295) & 63)))) | 1))
    var v15 int32 = 0
    for fi16 := 0; fi16 < 11; fi16++ {
        switch "" {
        case "":
            v3 = -32768
        case "hello":
            switch (^(fi16)) {
            case 42:
                var v17 uint64 = 18446744073709551615
                _ = v17
            default:
                for fi18 := 0; fi18 < 11; fi18++ {
                    if (2 != 0) {
                        v3 = v3
                        if true {
                            v19 := GenS3{F0: 4294967295, F1: 2}
                            _ = v19
                        } else {
                            var v20 uint8 = 123
                            fmt.Println(v3)
                            _ = v20
                        }
                    } else {
                        switch "a" {
                        case "":
                            var v21 uint64 = 4294967295
                            _ = v21
                        default:
                            var v22 float64 = 1.5
                            _ = v22
                        }
                        v2[fi18] = ((1 < 2147483648) || true)
                        v15 = v15
                        v3 += -32768
                        v2[fi18] = false
                    }
                    var v23 uint64 = 242
                    v24 := GenS2{F0: v3, F1: 936}
                    _ = fi18
                    _ = v23
                    _ = v24
                }
            }
        case "foo bar":
            for fi25 := 0; fi25 < 1; fi25++ {
                var v26 string = (v13 + "")
                if true {
                    v27 := v2[fi16]
                    var v28 float64 = float64(0.0)
                    _ = v27
                    _ = v28
                }
                _ = fi25
                _ = v26
            }
        default:
            var v29 uint8 = v12
            _ = v29
        }
        v30 := GenS4{F0: 100, F1: 1, F2: 42}
        _ = fi16
        _ = v30
    }
    fmt.Println(v15, v15, v15)
    if (!(((-128 == 127) && false) || false)) {
        if (780 != 1) {
            v12 -= v12
        }
    }
    for fi31 := 0; fi31 < 17; fi31++ {
        v32 := GenS5{F0: 42}
        _ = fi31
        _ = v32
    }
    v2[32768] = (true && (774 <= 42))
    if false {
        v12 = v12
        var v33 uint = 0
        _ = v33
    }
    switch "a" {
    case "":
        fmt.Println(v15)
    case "\n\t":
        if false {
            v34 := GenS6{F0: v3}
            for fi35 := 0; fi35 < 14; fi35++ {
                v15 = v15
                _ = len(v1)
                _ = fi35
            }
            _ = v34
        }
    case "foo bar":
        fmt.Println(v15)
    default:
        fmt.Println(v15)
    }
    v15 -= v15
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
}
