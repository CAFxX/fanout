package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 uint }
func genFn0(p0 int8, p1 uint8) int16 {
    var v4 int8 = -128
    fmt.Println(p1, v4, v4)
    if (3.14159 <= -2.25) {
        p0 = 1
    }
    v4 += p0
    _ = p0
    _ = p1
    _ = v4
    return 261
}

func main() {
    var v1 int16 = 1
    var v2 string = "\n\t"
    var v3 uint8 = 42
    var v5 int16 = v1
    fmt.Println(v5, v3, v3)
    switch 901 {
    case 582:
        var v6 int32 = 0
        _ = v6
    case 378:
        v7 := []int{9223372036854775807, 127}
        _ = v7
    case 4294967295:
        var v8 int32 = 0
        _ = v8
    }
    fmt.Println(v5)
    v2 = v2
    v9 := []int{32767}
    fmt.Println(v3)
    if len(v9) > 0 {
        v9[0] = 4294967295
    }
    _ = len(v9)
    if (-613 == 722) {
        if false {
            v2 = v2
        } else {
            _ = len(v9)
        }
    }
    v10 := GenS1{F0: 42, F1: 761, F2: 2}
    for fi11 := 0; fi11 < 19; fi11++ {
        v12 := GenS1{F0: 42, F1: 1, F2: 2}
        for fi13 := 0; fi13 < 14; fi13++ {
            if len(v9) > 0 {
                v9[0] = 0
            }
            v3 -= 128
            switch "hello" {
            case "hello":
                v1 -= genFn0(1, (v3 << ((255) & 63)))
                if ("\n\t" <= "hello") {
                    if (genFn0(-83, 1) <= (v5 + (v5 << ((566) & 63)))) {
                        if true {
                            var v14 bool = true
                            v15 := GenS1{F0: 0, F1: 4294967295, F2: 4294967295}
                            _ = v14
                            _ = v15
                        }
                    } else {
                        if true {
                            var v16 int64 = 42
                            fmt.Println(v3)
                            if true {
                                switch v2 {
                                case "hello":
                                    var v17 int = (-(fi13))
                                    var v18 float64 = (3.14159 + (1.5 * 1e10))
                                    _ = v17
                                    _ = v18
                                }
                                var v19 int16 = (v5 >> ((997) & 63))
                                if len(v9) >= 2 {
                                    v9 = v9[1:len(v9)]
                                }
                                _ = v19
                            }
                            if (v16 <= (v16 + (v16 + v16))) {
                                if ("" <= "foo bar") {
                                    switch v2 {
                                    case "\n\t":
                                        v20 := GenS1{F0: 1, F1: 4294967295, F2: 2}
                                        _ = v20
                                    case "foo bar":
                                        if false {
                                            v21 := GenS1{F0: 549, F1: 2147483647, F2: 65535}
                                            _ = v21
                                        } else {
                                            v16 = v16
                                        }
                                    }
                                } else {
                                    v16 = v16
                                    var v22 int16 = v1
                                    _ = v22
                                }
                                v23 := GenS1{F0: 273, F1: 65535, F2: 71}
                                _ = v23
                            }
                            _ = v16
                        }
                    }
                }
                var v24 float32 = float32(float64(0.1))
                if len(v9) >= 2 {
                    v9 = v9[1:len(v9)]
                }
                _ = v24
            }
            _ = fi13
        }
        var v25 int32 = 2147483647
        _ = fi11
        _ = v12
        _ = v25
    }
    v9 = append(v9, 127)
    fmt.Println(v2, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v5
    _ = v9
}
