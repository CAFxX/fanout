package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 int8 }

func main() {
    var v1 int16 = 1000
    var v2 float64 = 1e10
    v3 := []string{}
    v1 += (v1 - (-(v1)))
    if false {
        for fi4 := 0; fi4 < 3; fi4++ {
            if false {
                if true {
                    v5 := GenS1{F0: -1, F1: -2147483648, F2: 127}
                    if true {
                        v2 += (v2 * (((v2 / 3.14159) + (v2 - v2)) - (3.14159 / v2)))
                        var v6 int8 = 1
                        _ = v6
                    }
                    _ = len(v3)
                    _ = v5
                }
                v2 -= ((v2 + (1e10 / 0.1)) / v2)
            } else {
                switch "foo bar" {
                case "a":
                    v7 := GenS1{F0: -1, F1: 597, F2: 42}
                    _ = v7
                case "":
                    for fi8 := 0; fi8 < 19; fi8++ {
                        if (false || (!(false && true))) {
                            var v9 float32 = 0.0
                            var v10 int64 = -9223372036854775808
                            var v11 uint8 = 94
                            var v12 uint16 = 2
                            _ = v10
                            _ = v11
                            _ = v12
                            _ = v9
                        }
                        v13 := GenS1{F0: -1, F1: -1, F2: 1}
                        _ = fi8
                        _ = v13
                    }
                default:
                    for fi14 := 0; fi14 < 19; fi14++ {
                        v2 -= v2
                        _ = fi14
                    }
                }
                v3 = append(v3, (("" + "") + ("" + "a")))
                var v15 uint32 = 0
                _ = v15
            }
            fmt.Println(v2)
            _ = fi4
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    } else {
        fmt.Println(v1)
        v16 := GenS1{F0: 593, F1: -1, F2: -106}
        v2 = 0.1
        _ = v16
    }
    fmt.Println(v1)
    var v17 int32 = 257
    _ = len(v3)
    var v18 int16 = v1
    if len(v3) > 0 {
        v3[0] = ("\n\t" + "")
    }
    v17 = (^(v17))
    var v19 int = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
