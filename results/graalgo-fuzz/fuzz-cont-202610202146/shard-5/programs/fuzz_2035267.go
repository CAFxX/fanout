package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 uint8 }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 int8; F1 uint16; F2 int16 }

func main() {
    var v1 uint16 = 379
    var v2 uint8 = 255
    v3 := make(map[string]int)
    v4 := []int{}
    fmt.Println(v1)
    var v5 int16 = -430
    if false {
        v5 = v5
        v4 = append(v4, 42)
    } else {
        v6 := GenS1{F0: 42, F1: (v1 | v1), F2: (^(v2))}
        fmt.Println(v1)
        var v7 int32 = 0
        _ = v6
        _ = v7
    }
    if ("foo bar" > "foo bar") {
        switch 2147483647 {
        case 4294967295:
            v8 := GenS1{F0: 1, F1: v1, F2: 2}
            _ = v8
        }
    } else {
        v9 := GenS2{F0: -2147483648, F1: 722}
        _ = v9
    }
    var v10 string = "hello"
    for fi11 := 0; fi11 < 11; fi11++ {
        var v12 uint = 167
        for fi13 := 0; fi13 < 8; fi13++ {
            switch -40 {
            case 128:
                v3[("a" + "")] = (fi11 - (fi13 - 1))
            case -961:
                switch "\n\t" {
                case "\n\t":
                    v5 = (v5 & int16(v12))
                case "":
                    v14 := GenS1{F0: (v12 * v12), F1: (v1 + uint16(v2)), F2: 255}
                    _ = v14
                case "foo bar":
                    if false {
                        switch "hello" {
                        case "foo bar":
                            v10 = v10
                        case "":
                            v5 -= v5
                        case "a":
                            if len(v4) > 0 {
                                v4[0] = (fi13 ^ fi11)
                            }
                        }
                        v3[v10] = (fi13 & (-(fi13)))
                        var v15 uint64 = 447
                        _ = v15
                    } else {
                        if (!(!true)) {
                            v2 -= v2
                        } else {
                            var v16 uint16 = 32767
                            _ = v16
                        }
                        v1 = (^(v1))
                        v1 -= (v1 | 255)
                        if false {
                            fmt.Println(v12, v5)
                        } else {
                            fmt.Println(v1, v10)
                        }
                    }
                }
            }
            _ = len(v4)
            _ = fi13
        }
        _ = fi11
        _ = v12
    }
    for fi17 := 0; fi17 < 17; fi17++ {
        if (9223372036854775807 == 1) {
            if true {
                var v18 uint64 = 514
                delete(v3, "\n\t")
                if false {
                    if ((!false) || (((uint32(v2) <= 65535) && ((32767 == v5) && (float32(0.1) <= 1.5))) && false)) {
                        var v19 int = fi17
                        fmt.Println(v19, v18, fi17)
                        _ = v19
                    }
                    fmt.Println(fi17)
                }
                _ = v18
            } else {
                if (-116 == 119) {
                    if false {
                        fmt.Println(v1, v2, v1)
                    }
                    v10 = v10
                    switch v10 {
                    case "hello":
                        v5 -= v5
                    case "a":
                        v20 := GenS3{F0: -13, F1: (v1 / ((v1) | 1)), F2: -392}
                        _ = v20
                    default:
                        v2 = (^(v2))
                    }
                }
                var v21 float32 = (3.14159 * (1e10 + 3.14159))
                var v22 uint64 = 989
                _ = v21
                _ = v22
            }
        } else {
            for fi23 := 0; fi23 < 11; fi23++ {
                var v24 bool = (((v1 << ((4294967295) & 63)) <= (v1 % (((v1 & (v1 % ((255) | 1)))) | 1))) && (((true && false) && (255 == 0)) && (-656 == 1)))
                v25 := GenS1{F0: 1, F1: v1, F2: v2}
                var v26 string = v10
                _ = fi23
                _ = v24
                _ = v25
                _ = v26
            }
        }
        for fi27 := 0; fi27 < 2; fi27++ {
            var v28 uint8 = (v2 | 237)
            _ = len(v4)
            _ = fi27
            _ = v28
        }
        v1 -= (^(v1))
        _ = fi17
    }
    delete(v3, ("\n\t" + v10))
    v10 = v10
    v2 += v2
    v5 += 0
    switch ("hello" + "hello") {
    case "":
        _ = len(v3)
    default:
        v10 = (("" + (v10 + v10)) + "")
    }
    _ = len(v3)
    if true {
        v29 := GenS1{F0: 2, F1: v1, F2: (v2 & v2)}
        var v30 uint = 407
        _ = v29
        _ = v30
    }
    var v31 uint32 = 815
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
