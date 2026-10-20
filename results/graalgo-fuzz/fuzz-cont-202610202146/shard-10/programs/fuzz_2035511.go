package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 int32 }

func main() {
    var v1 int16 = -1
    var v2 uint16 = 42
    var v3 int = 4294967295
    v4 := []string{"\n\t", ""}
    v5 := make(map[int]string)
    v6 := make(map[string]int)
    switch v3 {
    case 929:
        v7 := GenS1{F0: (-(v3)), F1: 1, F2: 42}
        _ = v7
    case 2147483647:
        v3 = v3
    case -9223372036854775808:
        v8 := GenS1{F0: -9223372036854775808, F1: 18446744073709551615, F2: 0}
        _ = v8
    default:
        _ = len(v6)
    }
    if ((v2 <= v2) && false) {
        var v9 uint16 = (^(v2))
        v9 = 42
        switch "" {
        case "foo bar":
            for fi10 := 0; fi10 < 7; fi10++ {
                var v11 float64 = (0.1 / 0.1)
                delete(v6, "")
                if (true && true) {
                    v2 += (v9 << ((255) & 63))
                } else {
                    if true {
                        switch 4294967295 {
                        case 127:
                            if ("a" != "") {
                                fmt.Println(v1)
                                var v12 int16 = (v1 * 1000)
                                _ = v12
                            } else {
                                if (v11 == v11) {
                                    var v13 int8 = int8(v3)
                                    _ = v13
                                }
                                if (false || (v2 >= v9)) {
                                    if false {
                                        v14 := v5[v3]
                                        var v15 int8 = 100
                                        var v16 float64 = (1e10 * v11)
                                        _ = v14
                                        _ = v15
                                        _ = v16
                                    }
                                    v1 = (v1 * (v1 ^ (v1 + v1)))
                                } else {
                                    v11 = v11
                                }
                                fmt.Println(v9)
                            }
                        case 42:
                            v9 += (v2 + v2)
                        case 2147483647:
                            v1 = v1
                        }
                        for fi17 := 0; fi17 < 17; fi17++ {
                            v4 = append(v4, "foo bar")
                            switch "hello" {
                            case "hello":
                                if (!false) {
                                    var v18 uint = 2
                                    _ = v18
                                } else {
                                    _ = len(v6)
                                    v3 = 9223372036854775807
                                    if false {
                                        var v19 int = fi17
                                        _ = v19
                                    } else {
                                        v11 -= v11
                                        v20 := GenS1{F0: (fi17 ^ 127), F1: 298, F2: 42}
                                        var v21 string = ("\n\t" + "\n\t")
                                        fmt.Println(v21)
                                        if false {
                                            v22 := GenS1{F0: -812, F1: 781, F2: 2147483647}
                                            _ = v22
                                        }
                                        _ = v20
                                        _ = v21
                                    }
                                }
                            case "":
                                var v23 uint = 18446744073709551615
                                _ = v23
                            default:
                                v3 += (v3 << ((18446744073709551615) & 63))
                            }
                            fmt.Println(v3, v3)
                            _ = fi17
                        }
                    }
                }
                _ = fi10
                _ = v11
            }
        case "":
            v2 = (v2 + v9)
        case "\n\t":
            v1 -= v1
        default:
            if len(v4) > 0 {
                v4[0] = "a"
            }
        }
        _ = v9
    }
    var v24 bool = (!(-2.25 >= 0.1))
    var v25 int64 = int64(v1)
    var v26 float32 = 3.14159
    var v27 float64 = 1e10
    v28 := GenS1{F0: (v3 & v3), F1: 65535, F2: 0}
    v25 = (v25 % (((^(v25))) | 1))
    v24 = false
    v2 -= v2
    var v29 int32 = 982
    switch "foo bar" {
    case "foo bar":
        v1 -= 32767
    case "\n\t":
        v30 := GenS1{F0: (v3 * (v3 & v3)), F1: 1, F2: (v29 * (^(v29)))}
        _ = v30
    case "a":
        fmt.Println(v24, v3, v2)
    default:
        var v31 int32 = (v29 * 1)
        _ = v31
    }
    v32 := GenS1{F0: v3, F1: 970, F2: (v29 - v29)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
