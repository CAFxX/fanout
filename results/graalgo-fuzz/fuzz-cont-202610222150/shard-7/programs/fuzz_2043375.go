package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
func genFn0(p0 uint16, p1 string) int {
    var v4 uint8 = 255
    var v5 float32 = -2.25
    fmt.Println(p1)
    p0 = p0
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 42
}

func main() {
    var v1 uint64 = 140
    var v2 int32 = 2147483647
    v3 := []int{2147483647, 128}
    if false {
        var v6 float64 = (-2.25 - (3.14159 - ((((1.5 - 0.0) * 1.5) / 1.5) / (float64(1.5) - float64((1.5 / 3.14159))))))
        v6 += v6
        _ = v6
    } else {
        _ = len(v3)
        if (!(false && true)) {
            if true {
                v2 += 42
                v2 = (v2 ^ v2)
            }
            if len(v3) > 0 {
                v3[0] = 2147483647
            }
            for fi7 := 0; fi7 < 6; fi7++ {
                var v8 string = "a"
                fmt.Println(v8, v1, v2)
                var v9 int16 = 0
                _ = fi7
                _ = v8
                _ = v9
            }
        }
        v2 = v2
    }
    switch genFn0(255, ("foo bar" + "\n\t")) {
    case 1:
        v1 = v1
    case -2147483648:
        fmt.Println(v1, v1)
    }
    fmt.Println(v1, v1)
    fmt.Println(v2)
    var v10 uint64 = (v1 / (((^(v1))) | 1))
    switch -9223372036854775808 {
    case -502:
        v1 -= 4294967295
    }
    v2 = (v2 * v2)
    _ = len(v3)
    switch "\n\t" {
    case "":
        if true {
            v10 = v1
            _ = len(v3)
        }
    case "a":
        if true {
            if false {
                if true {
                    if false {
                        v10 += v10
                        switch genFn0(0, "a") {
                        case -164:
                            v1 += (v10 | v1)
                        case 788:
                            v1 -= v10
                        case 128:
                            var v11 int = 545
                            _ = v11
                        default:
                            for fi12 := 0; fi12 < 6; fi12++ {
                                var v13 int8 = 98
                                _ = fi12
                                _ = v13
                            }
                        }
                        var v14 int64 = -1
                        _ = v14
                    }
                }
            } else {
                if len(v3) > 0 {
                    v3[0] = 878
                }
                var v15 uint32 = uint32(v10)
                var v16 uint8 = 0
                _ = v15
                _ = v16
            }
        }
    }
    v17 := GenS1{F0: 1, F1: 618}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v3
}
