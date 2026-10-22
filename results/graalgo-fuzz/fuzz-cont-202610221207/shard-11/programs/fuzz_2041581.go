package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 int32 }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 uint16; F1 int16; F2 uint64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint; F1 uint; F2 bool }
type GenS6 struct { F0 uint32; F1 int }
func genFn0(p0 uint32, p1 uint16) uint16 {
    v7 := []int{0}
    var v8 uint8 = 77
    v9 := GenS1{F0: 1, F1: 855, F2: 0}
    var v10 uint16 = (p1 & p1)
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return p1
}
func genFn1() int64 {
    switch "a" {
    case "foo bar":
        var v11 bool = false
        _ = v11
    case "hello":
        var v12 int8 = -66
        _ = v12
    case "\n\t":
        if (((4294967295 != 255) && (1 > 42)) || (0 > -183)) {
            var v13 int16 = 0
            _ = v13
        }
    default:
        var v14 int8 = 127
        _ = v14
    }
    v15 := GenS2{F0: genFn0(4294967295, 65535), F1: ("\n\t" + ("" + ("foo bar" + "\n\t")))}
    v16 := []int{-9223372036854775808, 9223372036854775807, 127}
    if ("a" <= ("a" + "")) {
        var v17 float32 = 0.0
        fmt.Println(v17)
        _ = v17
    } else {
        var v18 uint8 = 42
        if (!(false || (false && (!true)))) {
            v19 := GenS3{F0: 1, F1: 1000, F2: 0}
            _ = v19
        }
        _ = v18
    }
    v20 := GenS4{F0: 2}
    _ = v15
    _ = v16
    _ = v20
    return 1
}

func main() {
    v1 := make(map[string]string)
    var v2 int = 821
    var v3 int16 = -32768
    var v4 uint8 = 127
    var v5 float32 = (((1.5 * -2.25) - 1e10) / ((((-2.25 + 0.0) - 3.14159) + 1e10) * 1.5))
    v6 := make(map[uint32]bool)
    v21 := GenS2{F0: genFn0(uint32(v4), genFn0(297, 0)), F1: ""}
    v2 += 32767
    fmt.Println(v2, v2, v3)
    if (false || true) {
        for fi22 := 0; fi22 < 16; fi22++ {
            switch (^(v2)) {
            case 32767:
                v23 := []int{}
                _ = v23
            case 1:
                var v24 int64 = 9223372036854775807
                _ = v24
            case 128:
                v25 := GenS2{F0: genFn0(712, 2), F1: ("a" + "hello")}
                _ = v25
            }
            for fi26 := 0; fi26 < 14; fi26++ {
                var v27 int8 = 0
                var v28 float64 = (3.14159 / 1e10)
                if true {
                    if (-2.25 >= float64(float32(0.0))) {
                        v27 = (v27 + (-(v27)))
                        v29 := []int{fi22, fi26}
                        _ = v29
                    } else {
                        v30 := GenS5{F0: 2, F1: 42, F2: (0 >= 91)}
                        v3 = (v3 | (-(v3)))
                        switch fi26 {
                        case -9223372036854775808:
                            v31 := []int{-552, 128, fi22}
                            _ = v31
                        case -916:
                            v2 += (v2 - fi22)
                        }
                        _ = v30
                    }
                } else {
                    v32 := []int{fi26}
                    v2 += (fi22 >> ((332) & 63))
                    v3 = v3
                    _ = v32
                }
                _ = fi26
                _ = v27
                _ = v28
            }
            _ = fi22
        }
    }
    v3 -= v3
    _ = len(v6)
    v33 := []int{-502, v2, v2}
    if true {
        v34 := GenS6{F0: 329, F1: v2}
        v3 = (v3 + v3)
        fmt.Println(v2, v5, v4)
        _ = v34
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
