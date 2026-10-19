package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 int16 }
type GenS2 struct { F0 int; F1 int8; F2 uint64 }
type GenS3 struct { F0 uint64 }
func genFn0() int32 {
    var v7 float32 = 0.1
    v8 := []int{4294967295, 0, -2147483648}
    if (840 > 42) {
        v7 = v7
        var v9 uint16 = 425
        fmt.Println(v7)
        _ = v9
    } else {
        v7 = 0.1
        v7 = (v7 / v7)
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
    }
    v10 := GenS1{F0: 1, F1: 255, F2: -1}
    _ = v10
    _ = v7
    _ = v8
    return 481
}

func main() {
    v1 := []string{"hello", ("hello" + ("\n\t" + (("" + "hello") + "foo bar"))), ""}
    var v2 uint16 = 255
    var v3 int16 = -1
    v4 := make(map[int]int)
    v5 := []int{}
    var v6 float32 = float32(((0.0 * float64(((1e10 * (float32(0.1) / 1.0)) * float32(1.5)))) - (1e10 - 3.14159)))
    v6 = v6
    switch ("" + "a") {
    case "hello":
        _ = len(v1)
    case "":
        v3 += (v3 * 42)
    }
    v6 += (3.14159 + 1.5)
    fmt.Println(v2)
    if true {
        for fi11 := 0; fi11 < 2; fi11++ {
            v12 := GenS1{F0: 132, F1: (v2 ^ v2), F2: (-(v3))}
            var v13 int = (fi11 / (((^(fi11))) | 1))
            if ((!true) || (!(127 != 193))) {
                switch "foo bar" {
                case "foo bar":
                    v6 += v6
                case "hello":
                    var v14 uint64 = 1
                    _ = v14
                case "":
                    if (!false) {
                        v2 -= 0
                        fmt.Println(fi11, v6)
                    }
                default:
                    v6 += (3.14159 + v6)
                }
                fmt.Println(v13, v3, v3)
                if (float32(1.5) != 0.1) {
                    var v15 int = (v13 + (fi11 | fi11))
                    _ = v15
                } else {
                    v13 -= fi11
                    if false {
                        v4[(fi11 >> ((4294967295) & 63))] = (fi11 - fi11)
                        for fi16 := 0; fi16 < 6; fi16++ {
                            v3 = 380
                            v17 := GenS2{F0: 127, F1: int8(v13), F2: 42}
                            fmt.Println(fi11, v2, v2)
                            _ = fi16
                            _ = v17
                        }
                        fmt.Println(fi11, v6)
                    } else {
                        for fi18 := 0; fi18 < 1; fi18++ {
                            v13 = (v13 << ((18446744073709551615) & 63))
                            v4[fi11] = 32767
                            _ = fi18
                        }
                    }
                    v3 = 963
                }
            } else {
                var v19 uint16 = (^(v2))
                if false {
                    v20 := GenS2{F0: (fi11 & 128), F1: 100, F2: 666}
                    _ = v20
                } else {
                    fmt.Println(v19, v6, v19)
                }
                _ = v19
            }
            _ = fi11
            _ = v12
            _ = v13
        }
        for fi21 := 0; fi21 < 4; fi21++ {
            fmt.Println(v6)
            v22 := GenS2{F0: fi21, F1: 1, F2: 4294967295}
            var v23 bool = false
            _ = fi21
            _ = v22
            _ = v23
        }
        v24 := GenS3{F0: 42}
        _ = v24
    }
    for fi25 := 0; fi25 < 6; fi25++ {
        switch ("hello" + "a") {
        case "\n\t":
            fmt.Println(v2, v2, v2)
        default:
            v5 = append(v5, (-(fi25)))
        }
        _ = fi25
    }
    for fi26 := 0; fi26 < 12; fi26++ {
        v27 := GenS2{F0: 42, F1: -62, F2: 868}
        var v28 float32 = ((1e10 - float32((((3.14159 - 0.0) + float64(v6)) * 3.14159))) / v6)
        v2 = 53
        _ = fi26
        _ = v27
        _ = v28
    }
    fmt.Println(v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
