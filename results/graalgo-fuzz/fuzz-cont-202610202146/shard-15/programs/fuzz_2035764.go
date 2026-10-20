package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 string }
func genFn0(p0 string, p1 int16) uint {
    var v4 string = ""
    p0 = "hello"
    _ = p0
    _ = p1
    _ = v4
    return 65535
}
func genFn1(p0 uint32) uint {
    p0 = p0
    var v5 uint8 = 255
    _ = p0
    _ = v5
    return genFn0("", -1)
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 int16 = -546
    var v3 int = -872
    fmt.Println(v3, v2, v1)
    if (!(v2 <= (v2 << ((1) & 63)))) {
        var v6 int32 = -2147483648
        switch 141 {
        case 42:
            v2 -= v2
        case 32768:
            var v7 uint32 = 480
            _ = v7
        default:
            var v8 float32 = -2.25
            _ = v8
        }
        v3 = (-(v3))
        _ = v6
    } else {
        v3 -= v3
        switch "" {
        case "\n\t":
            if (179 > 128) {
                var v9 float64 = ((0.0 - 1e10) + (0.1 / ((-2.25 - 0.0) * -2.25)))
                switch "foo bar" {
                case "\n\t":
                    v10 := []int{(^(v3))}
                    _ = v10
                }
                _ = v9
            } else {
                var v11 int8 = 0
                _ = v11
            }
        case "a":
            for fi12 := 0; fi12 < 8; fi12++ {
                v13 := []int{-948, (v3 | (fi12 - 32767)), (v3 ^ v3), (fi12 & (fi12 - v3))}
                _ = fi12
                _ = v13
            }
        default:
            for fi14 := 0; fi14 < 2; fi14++ {
                if false {
                    switch fi14 {
                    case 221:
                        v1 = v1
                    default:
                        v2 -= v2
                    }
                } else {
                    fmt.Println(fi14)
                    v15 := []int{fi14, v3, fi14}
                    var v16 bool = false
                    _ = v15
                    _ = v16
                }
                switch "\n\t" {
                case "hello":
                    v1 = (^(v1))
                case "":
                    v2 = v2
                default:
                    for fi17 := 0; fi17 < 17; fi17++ {
                        var v18 int = fi14
                        var v19 float32 = (float32((0.1 * ((0.1 + -2.25) + 1.5))) + 1.5)
                        _ = fi17
                        _ = v18
                        _ = v19
                    }
                }
                _ = fi14
            }
        }
    }
    var v20 bool = (!false)
    var v21 int64 = -808
    v22 := GenS1{F0: v2, F1: genFn0("foo bar", v2), F2: "a"}
    fmt.Println(v21, v20, v21)
    v23 := []int{(v3 / (((v3 >> ((1) & 63))) | 1))}
    v2 -= (-(v2))
    fmt.Println(v21, v1)
    if false {
        v23 = append(v23, v3)
        v20 = (v21 == 159)
    }
    v2 += v2
    if (!v20) {
        _ = len(v23)
        fmt.Println(v3)
    } else {
        if v20 {
            switch "\n\t" {
            case "hello":
                var v24 int32 = 441
                _ = v24
            case "a":
                var v25 int32 = int32(v2)
                _ = v25
            }
            if len(v23) > 0 {
                v23[0] = (v3 >> ((4294967295) & 63))
            }
        } else {
            fmt.Println(v2, v3)
        }
        for fi26 := 0; fi26 < 6; fi26++ {
            v27 := GenS1{F0: 42, F1: 2, F2: ("" + "a")}
            fmt.Println(v2, v21)
            _ = fi26
            _ = v27
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
}
