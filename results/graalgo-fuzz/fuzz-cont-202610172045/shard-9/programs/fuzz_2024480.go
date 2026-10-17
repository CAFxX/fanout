package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16; F1 int; F2 uint64 }
func genFn0(p0 uint8, p1 uint32) int {
    switch "" {
    case "":
        for fi4 := 0; fi4 < 7; fi4++ {
            p1 -= (p1 / (((p1 << ((0) & 63))) | 1))
            for fi5 := 0; fi5 < 5; fi5++ {
                v6 := []int{127, (fi4 << ((1) & 63)), fi4, (-(fi4))}
                if ((680 < 4294967295) || false) {
                    var v7 float32 = (0.1 / 1.5)
                    var v8 int = fi4
                    var v9 float64 = 0.0
                    _ = v7
                    _ = v8
                    _ = v9
                }
                fmt.Println(p1)
                switch ("hello" + "") {
                case "foo bar":
                    p1 += (p1 | (^(p1)))
                case "hello":
                    var v10 uint8 = (^(p0))
                    _ = v10
                }
                _ = fi5
                _ = v6
            }
            if true {
                var v11 string = "\n\t"
                v12 := GenS1{F0: fi4, F1: 42}
                var v13 int = (fi4 | fi4)
                _ = v11
                _ = v12
                _ = v13
            }
            _ = fi4
        }
    case "foo bar":
        var v14 uint = 380
        _ = v14
    default:
        v15 := []int{916, 32767}
        _ = v15
    }
    var v16 int64 = 1
    fmt.Println(p0, v16)
    _ = p0
    _ = p1
    _ = v16
    return 2147483647
}

func main() {
    v1 := []int{29, -75, -2147483648, 476}
    v2 := make(map[uint32]int)
    var v3 uint32 = 86
    switch "\n\t" {
    case "\n\t":
        var v17 uint16 = 32767
        _ = v17
    case "":
        v3 -= 682
    case "a":
        var v18 uint16 = 1
        _ = v18
    }
    var v19 uint16 = 42
    v20 := GenS1{F0: genFn0(1, (v3 - (v3 << ((0) & 63)))), F1: 1}
    var v21 float64 = 0.0
    var v22 int16 = 1
    v1 = append(v1, -2147483648)
    v23 := GenS1{F0: genFn0(133, v3), F1: (v22 & v22)}
    v1 = append(v1, 133)
    if false {
        v24 := GenS2{F0: 946}
        _ = v24
    }
    if len(v1) > 0 {
        v1[0] = -9223372036854775808
    }
    for fi25 := 0; fi25 < 20; fi25++ {
        switch ("a" + "") {
        case "":
            v22 -= 32767
        case "\n\t":
            if (false || true) {
                v26 := GenS3{F0: v19, F1: -132, F2: 844}
                _ = v26
            } else {
                for fi27 := 0; fi27 < 12; fi27++ {
                    if ((((0.0 / (1.5 / (1.5 - (((((1.5 + float32(float64(((float32(float64(float32(v21))) / 3.14159) * (1.5 * (-2.25 - 0.1)))))) * (0.0 * 1e10)) - (1e10 / -2.25)) * -2.25) * (float32(v21) - 1.5))))) + 0.1) + 1.5) <= (3.14159 * 0.0)) {
                        var v28 uint16 = (v19 << ((4) & 63))
                        v19 += 238
                        _ = v28
                    }
                    _ = fi27
                }
            }
        case "hello":
            v21 -= v21
        default:
            for fi29 := 0; fi29 < 17; fi29++ {
                switch (("" + "foo bar") + (("a" + ("hello" + "\n\t")) + ("hello" + (("\n\t" + "\n\t") + "foo bar")))) {
                case "a":
                    v21 = v21
                case "hello":
                    if len(v1) > 0 {
                        v1[0] = 1
                    }
                default:
                    v3 += 24
                }
                var v30 int16 = (v22 ^ (v22 >> ((877) & 63)))
                _ = fi29
                _ = v30
            }
        }
        v2[v3] = (fi25 * (fi25 >> ((2) & 63)))
        if len(v1) > 0 {
            v1[0] = fi25
        }
        _ = fi25
    }
    var v31 float64 = float64(1e10)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v31
}
