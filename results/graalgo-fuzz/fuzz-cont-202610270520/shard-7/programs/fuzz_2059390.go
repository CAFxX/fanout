package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
func genFn0(p0 int8) int64 {
    v7 := []int{2147483647, 2147483647, 127, 458}
    p0 = (p0 * (p0 % (((p0 - (p0 + 1))) | 1)))
    _ = p0
    _ = v7
    return 132
}

func main() {
    v1 := make(map[string]bool)
    v2 := []string{"foo bar", "hello", (((("" + ("" + ("foo bar" + ""))) + ("hello" + "")) + "\n\t") + "foo bar"), "\n\t"}
    var v3 float32 = (float32(1.5) + 0.0)
    var v4 uint64 = 2
    var v5 uint = 65535
    v6 := make(map[uint32]bool)
    for fi8 := 0; fi8 < 13; fi8++ {
        v1["\n\t"] = false
        v9 := v1[("" + ("foo bar" + (("foo bar" + "hello") + ("foo bar" + "foo bar"))))]
        if (219 < 176) {
            switch "\n\t" {
            case "hello":
                fmt.Println(v5, v5)
            case "":
                fmt.Println(fi8)
            case "foo bar":
                if false {
                    if (2147483647 != fi8) {
                        for fi10 := 0; fi10 < 16; fi10++ {
                            var v11 uint = 399
                            delete(v1, "")
                            var v12 bool = (0.1 > 0.0)
                            _ = fi10
                            _ = v11
                            _ = v12
                        }
                        v5 = v5
                    }
                    v5 += (^(v5))
                } else {
                    for fi13 := 0; fi13 < 5; fi13++ {
                        var v14 int16 = 32767
                        switch (-(fi13)) {
                        case 42:
                            var v15 string = "\n\t"
                            _ = v15
                        case 32768:
                            v16 := GenS1{F0: 42, F1: 18446744073709551615}
                            delete(v6, 65535)
                            v14 = -1
                            _ = v16
                        default:
                            v17 := v6[65535]
                            _ = v17
                        }
                        fmt.Println(v4, fi8)
                        _ = fi13
                        _ = v14
                    }
                    var v18 float32 = 1.5
                    _ = v18
                }
            default:
                for fi19 := 0; fi19 < 20; fi19++ {
                    v4 = (v4 & (v4 >> ((255) & 63)))
                    var v20 uint64 = v4
                    _ = fi19
                    _ = v20
                }
            }
            var v21 uint64 = 4294967295
            _ = v21
        }
        _ = fi8
        _ = v9
    }
    v3 -= -2.25
    if false {
        var v22 int16 = 1000
        delete(v6, 1)
        _ = v22
    }
    v6[2] = false
    switch -893 {
    case -921:
        var v23 int64 = -562
        _ = v23
    case 42:
        var v24 uint8 = 255
        _ = v24
    case 2147483647:
        v25 := GenS1{F0: 42, F1: (v5 | (v5 & 439))}
        _ = v25
    default:
        v3 = float32((float64(3.14159) / (3.14159 * ((1e10 / (1e10 / -2.25)) + (0.0 + 0.0)))))
    }
    for fi26 := 0; fi26 < 9; fi26++ {
        v5 = (v5 / ((v5) | 1))
        v27 := GenS1{F0: 255, F1: 42}
        _ = fi26
        _ = v27
    }
    v28 := GenS1{F0: 2, F1: v5}
    var v29 uint32 = 4294967295
    v5 = (v5 / (((v5 % ((v5) | 1))) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
