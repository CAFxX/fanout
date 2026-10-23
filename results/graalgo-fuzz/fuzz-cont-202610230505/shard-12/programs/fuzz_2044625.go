package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 int }
func genFn0(p0 int16, p1 bool, p2 int8) int32 {
    var v5 int = -1
    for fi6 := 0; fi6 < 3; fi6++ {
        var v7 uint = 843
        v8 := GenS1{F0: "a", F1: 65535, F2: 9223372036854775807}
        _ = fi6
        _ = v7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 2147483647
}
func genFn1() int32 {
    var v9 uint64 = 42
    fmt.Println(v9)
    _ = v9
    return genFn0(-91, (!true), 1)
}

func main() {
    v1 := make(map[int]int)
    v2 := []int{}
    v3 := []int{508, -46}
    v4 := []uint32{65535}
    _ = len(v4)
    var v10 int8 = 100
    var v11 uint16 = 255
    v10 -= 127
    var v12 uint = 403
    if false {
        var v13 float32 = (3.14159 / float32(0.1))
        switch -2147483648 {
        case 878:
            var v14 int8 = v10
            _ = v14
        case 4294967295:
            var v15 float32 = (v13 / -2.25)
            _ = v15
        default:
            v1[-1] = 128
        }
        _ = v13
    } else {
        fmt.Println(v10)
        if false {
            switch ((("foo bar" + ("" + "a")) + ("foo bar" + "a")) + "foo bar") {
            case "\n\t":
                for fi16 := 0; fi16 < 12; fi16++ {
                    v11 = 42
                    v12 = 18446744073709551615
                    for fi17 := 0; fi17 < 3; fi17++ {
                        v12 = 1
                        _ = fi17
                    }
                    _ = fi16
                }
            case "":
                switch -390 {
                case 421:
                    switch "foo bar" {
                    case "a":
                        for fi18 := 0; fi18 < 3; fi18++ {
                            if (false || false) {
                                var v19 float32 = 1e10
                                switch fi18 {
                                case 128:
                                    v20 := GenS1{F0: "\n\t", F1: 152, F2: 9223372036854775807}
                                    _ = v20
                                case -903:
                                    v21 := GenS1{F0: ("hello" + "a"), F1: 904, F2: (fi18 * fi18)}
                                    _ = v21
                                case 1:
                                    var v22 float64 = float64(1e10)
                                    _ = v22
                                }
                                v4 = append(v4, 2147483648)
                                _ = v19
                            } else {
                                for fi23 := 0; fi23 < 6; fi23++ {
                                    if false {
                                        v24 := GenS1{F0: "hello", F1: 0, F2: (^(fi23))}
                                        var v25 uint16 = v11
                                        fmt.Println(v25)
                                        if false {
                                            if len(v2) > 0 {
                                                v2[0] = (^(fi23))
                                            }
                                            v1[fi18] = (-(fi18))
                                        } else {
                                            v25 = (^(v25))
                                            v26 := GenS1{F0: "hello", F1: 2147483647, F2: (fi23 * fi18)}
                                            if false {
                                                v3 = append(v3, (fi23 + fi23))
                                                var v27 uint8 = 8
                                                _ = v27
                                            }
                                            _ = v26
                                        }
                                        var v28 bool = ((v25 ^ 996) > v25)
                                        _ = v24
                                        _ = v25
                                        _ = v28
                                    } else {
                                        var v29 int32 = -1
                                        v4 = append(v4, 847)
                                        _ = v29
                                    }
                                    _ = fi23
                                }
                                var v30 float32 = 0.0
                                _ = v30
                            }
                            _ = fi18
                        }
                    case "hello":
                        fmt.Println(v12, v12, v10)
                    }
                case 2147483647:
                    if (!false) {
                        if len(v4) > 0 {
                            v4[0] = 1
                        }
                    }
                default:
                    v31 := v1[42]
                    _ = v31
                }
            case "hello":
                var v32 uint8 = 128
                _ = v32
            default:
                v11 = v11
            }
            if true {
                var v33 uint16 = v11
                _ = v33
            }
        }
        v34 := GenS1{F0: "foo bar", F1: 1, F2: 559}
        _ = v34
    }
    v11 += (^(v11))
    var v35 int32 = -1
    if (895 < uint16(v12)) {
        for fi36 := 0; fi36 < 6; fi36++ {
            switch ("foo bar" + "foo bar") {
            case "hello":
                v37 := GenS1{F0: "a", F1: 219, F2: int(v11)}
                _ = v37
            case "":
                v12 = (v12 * (^(v12)))
            case "foo bar":
                switch (^(fi36)) {
                case -461:
                    v10 += int8(fi36)
                case 127:
                    var v38 int8 = -28
                    _ = v38
                case 0:
                    v35 += (v35 << ((1) & 63))
                }
            }
            _ = fi36
        }
        v10 = v10
        v12 = 2
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v35
    _ = v4
}
