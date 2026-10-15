package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
func genFn0(p0 uint64, p1 int64, p2 uint8) uint16 {
    p2 += p2
    v4 := []int{127}
    var v5 int8 = 127
    v6 := GenS1{F0: 0, F1: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return 2
}

func main() {
    var v1 int16 = -32768
    var v2 uint = 65535
    var v3 bool = false
    var v7 int8 = 1
    v8 := []int{128, 1, 2147483647, 4294967295}
    v9 := GenS1{F0: 1, F1: v2}
    if true {
        for fi10 := 0; fi10 < 15; fi10++ {
            var v11 uint64 = 0
            v3 = v3
            var v12 uint16 = 0
            _ = fi10
            _ = v11
            _ = v12
        }
        v2 = v2
        var v13 int = -1
        _ = v13
    }
    var v14 int = -1
    if (65535 < 42) {
        v2 += (v2 ^ v2)
    } else {
        switch v14 {
        case -370:
            switch -897 {
            case 127:
                var v15 float32 = (1e10 + (1.5 - (0.0 / 1e10)))
                _ = v15
            case 9223372036854775807:
                var v16 bool = v3
                _ = v16
            default:
                if ((0.0 + 3.14159) != (0.1 / (1.5 / (float64(-2.25) * ((-2.25 / 0.1) + 0.0))))) {
                    if (!v3) {
                        var v17 uint32 = 571
                        if v3 {
                            if len(v8) > 0 {
                                v8[0] = 9223372036854775807
                            }
                        } else {
                            v17 -= v17
                            v2 += (v2 * 255)
                        }
                        if len(v8) >= 2 {
                            v8 = v8[1:len(v8)]
                        }
                        _ = v17
                    }
                    for fi18 := 0; fi18 < 6; fi18++ {
                        v3 = true
                        var v19 int64 = -797
                        if len(v8) > 0 {
                            v8[0] = v14
                        }
                        _ = fi18
                        _ = v19
                    }
                }
            }
        case 127:
            var v20 int32 = -1
            _ = v20
        }
    }
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    var v21 string = ("" + ((("" + "foo bar") + ("" + ("a" + ""))) + "\n\t"))
    for fi22 := 0; fi22 < 8; fi22++ {
        var v23 bool = (1 > 463)
        _ = fi22
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v14
    _ = v2
    _ = v21
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
