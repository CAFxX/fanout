package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 int16 }
type GenS2 struct { F0 uint; F1 bool }

func main() {
    v1 := make(map[string]uint32)
    var v2 int8 = 38
    v3 := []uint32{439, 505}
    v4 := []uint32{65535, 2147483647}
    fmt.Println(v2)
    var v5 int8 = v2
    var v6 uint64 = 18446744073709551615
    if true {
        var v7 int32 = int32(v6)
        _ = v7
    }
    for fi8 := 0; fi8 < 15; fi8++ {
        var v9 bool = (-2147483648 <= 675)
        for fi10 := 0; fi10 < 19; fi10++ {
            v11 := GenS1{F0: fi8, F1: 42, F2: 178}
            _ = len(v1)
            _ = fi10
            _ = v11
        }
        var v12 bool = ((v5 >> ((2) & 63)) == v2)
        _ = fi8
        _ = v12
        _ = v9
    }
    switch int(v6) {
    case 1:
        v5 -= v5
    case 0:
        switch "hello" {
        case "a":
            v13 := GenS2{F0: 18446744073709551615, F1: false}
            _ = v13
        case "foo bar":
            if true {
                fmt.Println(v5)
                _ = len(v3)
                if false {
                    if (2147483647 == 466) {
                        v5 += (v5 + (-(v5)))
                        v2 -= v2
                        var v14 string = ""
                        _ = v14
                    }
                } else {
                    if false {
                        switch ("\n\t" + "a") {
                        case "":
                            fmt.Println(v5, v2)
                        }
                        switch 32767 {
                        case -2147483648:
                            v15 := v1["\n\t"]
                            _ = v15
                        case 2147483647:
                            var v16 uint = 4294967295
                            _ = v16
                        case 0:
                            var v17 uint32 = 4294967295
                            _ = v17
                        }
                        fmt.Println(v2, v6, v6)
                    }
                    _ = len(v3)
                }
            } else {
                v6 = v6
                if true {
                    v6 = 0
                } else {
                    var v18 float64 = float64((float32(((1.5 - 1e10) + 1e10)) + -2.25))
                    _ = v18
                }
            }
        case "\n\t":
            if (65535 > 518) {
                for fi19 := 0; fi19 < 3; fi19++ {
                    v20 := v1["foo bar"]
                    v21 := GenS2{F0: 0, F1: (false || (1 != 789))}
                    if (0 > 50) {
                        v22 := v1[("hello" + "\n\t")]
                        if true {
                            var v23 int16 = -463
                            if ("\n\t" != "foo bar") {
                                if (834 == 4294967295) {
                                    v2 -= 59
                                    fmt.Println(fi19, v2)
                                } else {
                                    v2 = (v2 | v5)
                                    v23 = v23
                                }
                            }
                            _ = v23
                        } else {
                            switch ("foo bar" + "hello") {
                            case "hello":
                                v2 += v5
                            case "foo bar":
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                        }
                        for fi24 := 0; fi24 < 11; fi24++ {
                            var v25 uint16 = 1
                            var v26 int8 = (v5 * (v2 * v2))
                            _ = fi24
                            _ = v25
                            _ = v26
                        }
                        _ = v22
                    } else {
                        v6 += v6
                    }
                    _ = fi19
                    _ = v20
                    _ = v21
                }
            }
        default:
            if false {
                v3 = append(v3, 661)
            }
        }
    case 127:
        var v27 uint = 852
        _ = v27
    }
    v28 := v1[""]
    switch 32767 {
    case 0:
        fmt.Println(v6)
    default:
        v29 := GenS1{F0: 128, F1: -1, F2: 42}
        _ = v29
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
