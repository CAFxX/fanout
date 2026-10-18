package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 uint }
type GenS2 struct { F0 int32; F1 uint8; F2 int32 }
type GenS3 struct { F0 uint64; F1 string }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int16; F1 bool }

func main() {
    v1 := []string{("\n\t" + "\n\t"), ("foo bar" + ""), "a", ("" + "\n\t")}
    v2 := make(map[int]int)
    var v3 int8 = -128
    var v4 uint32 = 2147483648
    if ((!false) || (-1 >= 421)) {
        v5 := GenS1{F0: 495, F1: 1, F2: 65535}
        v6 := GenS2{F0: 1, F1: 127, F2: -302}
        _ = v5
        _ = v6
    }
    switch "" {
    case "foo bar":
        if false {
            switch 32768 {
            case 128:
                v3 += (v3 << ((672) & 63))
            }
        } else {
            var v7 uint = 0
            if (9223372036854775807 >= 42) {
                if true {
                    for fi8 := 0; fi8 < 10; fi8++ {
                        if (true || false) {
                            fmt.Println(v7, v4, v7)
                            v3 += (v3 & 1)
                        }
                        _ = fi8
                    }
                    var v9 string = "foo bar"
                    if (-488 < 42) {
                        for fi10 := 0; fi10 < 4; fi10++ {
                            if len(v1) > 0 {
                                v1[0] = v9
                            }
                            _ = fi10
                        }
                        var v11 float32 = ((1e10 - (float32(((float64((((1e10 - (1e10 * 1e10)) / 1.0) * (0.1 / 1.5))) - 0.0) / 1.0)) - 1.5)) * ((0.0 / ((3.14159 + 3.14159) + 1e10)) - (float32((1.5 + 0.0)) / 1.5)))
                        _ = v11
                    }
                    _ = v9
                }
                if false {
                    var v12 uint8 = 255
                    switch "a" {
                    case "a":
                        var v13 float32 = float32(0.0)
                        _ = v13
                    case "":
                        if true {
                            if true {
                                if true {
                                    switch 127 {
                                    case 9223372036854775807:
                                        switch 42 {
                                        case 128:
                                            fmt.Println(v12, v7, v7)
                                        case 4294967295:
                                            v14 := GenS3{F0: 4294967295, F1: ""}
                                            _ = v14
                                        }
                                    case -811:
                                        if (false || true) {
                                            fmt.Println(v3)
                                        }
                                    }
                                }
                                if (3.14159 <= 0.1) {
                                    if false {
                                        fmt.Println(v3)
                                        switch 887 {
                                        case 957:
                                            fmt.Println(v12, v3)
                                        case 254:
                                            v15 := GenS2{F0: -1, F1: 1, F2: 42}
                                            _ = v15
                                        }
                                    }
                                }
                            } else {
                                var v16 uint64 = 35
                                v2[int(v3)] = 4294967295
                                if (0.1 <= 1e10) {
                                    for fi17 := 0; fi17 < 18; fi17++ {
                                        v18 := GenS4{F0: (^(v12))}
                                        fmt.Println(v16, v16, v12)
                                        _ = fi17
                                        _ = v18
                                    }
                                } else {
                                    v4 = (v4 & uint32(v7))
                                    for fi19 := 0; fi19 < 7; fi19++ {
                                        for fi20 := 0; fi20 < 9; fi20++ {
                                            v21 := GenS5{F0: -32768, F1: (uint64(v4) < 42)}
                                            var v22 int8 = (v3 * v3)
                                            var v23 uint16 = 247
                                            _ = fi20
                                            _ = v21
                                            _ = v22
                                            _ = v23
                                        }
                                        var v24 uint32 = (^(v4))
                                        _ = fi19
                                        _ = v24
                                    }
                                }
                                _ = v16
                            }
                            v25 := v2[127]
                            _ = v25
                        }
                    case "foo bar":
                        v2[-186] = 9223372036854775807
                    }
                    v12 += v12
                    _ = v12
                } else {
                    fmt.Println(v3)
                    for fi26 := 0; fi26 < 2; fi26++ {
                        v3 = v3
                        if len(v1) > 0 {
                            v1[0] = ("" + ((((("a" + "a") + "a") + ("a" + ("hello" + "a"))) + "") + "foo bar"))
                        }
                        var v27 int64 = 1
                        _ = fi26
                        _ = v27
                    }
                }
                var v28 int = -9223372036854775808
                _ = v28
            } else {
                var v29 uint = 501
                v30 := v2[-668]
                v2[v30] = v30
                _ = v29
                _ = v30
            }
            v2[-367] = -2147483648
            _ = v7
        }
    default:
        var v31 uint8 = 1
        _ = v31
    }
    fmt.Println(v4)
    v4 = (v4 * 947)
    var v32 uint = 18446744073709551615
    v32 -= (v32 & (v32 * v32))
    v3 += v3
    var v33 uint16 = 65535
    var v34 int64 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
}
