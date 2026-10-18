package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 int16; F1 uint; F2 uint }
type GenS3 struct { F0 bool; F1 uint16 }
func genFn0(p0 int64, p1 uint, p2 bool) bool {
    p0 += p0
    if false {
        var v7 uint16 = 2
        for fi8 := 0; fi8 < 20; fi8++ {
            v9 := GenS1{F0: 41, F1: fi8}
            for fi10 := 0; fi10 < 18; fi10++ {
                var v11 int = (-(fi8))
                var v12 int16 = -32768
                _ = fi10
                _ = v11
                _ = v12
            }
            for fi13 := 0; fi13 < 16; fi13++ {
                v14 := []int{(fi13 + fi13)}
                switch fi8 {
                case 1:
                    switch ("\n\t" + (("\n\t" + "hello") + ("hello" + "a"))) {
                    case "":
                        if (1 > 0) {
                            if false {
                                if len(v14) > 0 {
                                    v14[0] = fi13
                                }
                                var v15 int32 = 42
                                v15 = v15
                                _ = v15
                            }
                            var v16 string = "\n\t"
                            _ = v16
                        }
                    default:
                        var v17 int32 = 332
                        _ = v17
                    }
                    v18 := GenS1{F0: 65535, F1: 1}
                    p1 += (p1 & p1)
                    _ = v18
                case -1:
                    var v19 int32 = -2147483648
                    _ = v19
                case 127:
                    v14 = append(v14, (fi13 - fi8))
                default:
                    fmt.Println(fi8, v7, fi8)
                }
                _ = len(v14)
                _ = fi13
                _ = v14
            }
            _ = fi8
            _ = v9
        }
        _ = v7
    } else {
        for fi20 := 0; fi20 < 8; fi20++ {
            switch 820 {
            case -582:
                for fi21 := 0; fi21 < 15; fi21++ {
                    var v22 uint = (p1 + (p1 << ((p1) & 63)))
                    fmt.Println(fi21, v22, v22)
                    _ = fi21
                    _ = v22
                }
            case -2147483648:
                var v23 int32 = -1
                _ = v23
            default:
                var v24 uint16 = 2
                _ = v24
            }
            fmt.Println(p0)
            var v25 float32 = ((((1e10 + (0.1 + (((1.5 * ((1e10 * 1.5) / 1.0)) + (-2.25 + 1e10)) * 3.14159))) / -2.25) / 0.1) * ((0.0 * (0.0 * (3.14159 - ((-2.25 - (1e10 - (0.0 - (((0.0 + (-2.25 + float32((0.0 / (3.14159 * ((1.5 - ((0.0 + 0.0) * 1e10)) - 3.14159)))))) / 1e10) - (3.14159 + 0.0))))) * 1.5)))) * (((3.14159 / 1.5) - 1.5) / 1.5)))
            _ = fi20
            _ = v25
        }
        var v26 uint = p1
        _ = v26
    }
    _ = p0
    _ = p1
    _ = p2
    return p2
}

func main() {
    v1 := make(map[string]int)
    var v2 int8 = 100
    v3 := make(map[string]string)
    var v4 uint64 = 925
    var v5 uint64 = (v4 % (((^(v4))) | 1))
    var v6 bool = ((747 >= 0) && false)
    v5 = (v5 - v5)
    var v27 int16 = -420
    var v28 bool = (2 != 2147483648)
    v29 := GenS2{F0: (-(v27)), F1: 18446744073709551615, F2: 255}
    v6 = v28
    v30 := GenS3{F0: v28, F1: 42}
    var v31 bool = false
    if false {
        var v32 float32 = ((1e10 * 0.1) / 0.1)
        if v31 {
            switch "a" {
            case "\n\t":
                fmt.Println(v31, v31, v28)
            case "hello":
                _ = len(v3)
            case "a":
                v33 := []int{}
                _ = v33
            }
        }
        _ = v32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
