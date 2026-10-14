package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int16; F1 bool }
func genFn0(p0 int32, p1 int64, p2 uint16) int8 {
    v7 := GenS1{F0: 255, F1: (p2 + p2)}
    v8 := []int{-1, -9223372036854775808, 128}
    var v9 float32 = 1e10
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 127
}
func genFn1(p0 int64, p1 uint16, p2 uint32) uint32 {
    var v10 string = ("\n\t" + "foo bar")
    v11 := GenS1{F0: 180, F1: p1}
    switch v10 {
    case "":
        if false {
            switch -1 {
            case 127:
                if (true || false) {
                    for fi12 := 0; fi12 < 7; fi12++ {
                        if (false && true) {
                            v13 := GenS2{F0: 42}
                            _ = v13
                        } else {
                            if true {
                                var v14 float32 = 1.5
                                _ = v14
                            } else {
                                var v15 uint16 = (p1 * 2)
                                _ = v15
                            }
                            var v16 int8 = 0
                            switch ("\n\t" + (v10 + ("hello" + (v10 + "hello")))) {
                            case "hello":
                                var v17 uint64 = 1
                                _ = v17
                            case "foo bar":
                                switch "" {
                                case "hello":
                                    v18 := GenS1{F0: 238, F1: (p1 >> ((2) & 63))}
                                    _ = v18
                                }
                            default:
                                if false {
                                    fmt.Println(p0, p2)
                                    p2 -= 42
                                    var v19 int64 = p0
                                    _ = v19
                                } else {
                                    v20 := GenS3{F0: 127}
                                    v10 = "hello"
                                    _ = v20
                                }
                            }
                            _ = v16
                        }
                        var v21 uint64 = 406
                        v22 := GenS1{F0: 1, F1: (p1 % (((p1 >> ((825) & 63))) | 1))}
                        _ = fi12
                        _ = v21
                        _ = v22
                    }
                    var v23 int = -233
                    var v24 string = "\n\t"
                    _ = v23
                    _ = v24
                }
            case -723:
                var v25 float64 = 0.1
                _ = v25
            case 0:
                v10 = ("" + "")
            default:
                if false {
                    for fi26 := 0; fi26 < 3; fi26++ {
                        v27 := GenS4{F0: 32767}
                        _ = fi26
                        _ = v27
                    }
                    v28 := GenS5{F0: -32768, F1: (!(!((false || (false && (true && (p0 == (p0 * (p0 - p0)))))) || true)))}
                    _ = v28
                } else {
                    v10 = (v10 + "hello")
                    p1 -= p1
                    fmt.Println(p1, p1, p0)
                }
            }
            fmt.Println(v10, v10, p1)
        } else {
            var v29 uint8 = 0
            var v30 int32 = -425
            _ = v29
            _ = v30
        }
    case "foo bar":
        v31 := GenS4{F0: 0}
        _ = v31
    case "\n\t":
        v32 := []int{}
        _ = v32
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return 42
}

func main() {
    var v1 int = -1
    var v2 uint = 255
    v3 := []string{}
    var v4 uint16 = 32767
    var v5 uint = (v2 ^ (v2 << (((v2 / ((v2) | 1))) & 63)))
    v6 := []int{}
    v1 += 1
    var v33 uint8 = 127
    switch "hello" {
    case "foo bar":
        fmt.Println(v1, v4, v5)
    case "a":
        v6 = append(v6, (v1 << (((^(v2))) & 63)))
    case "\n\t":
        var v34 uint64 = 295
        _ = v34
    }
    var v35 int = v1
    var v36 int64 = 42
    v35 += v35
    if false {
        fmt.Println(v4, v4)
        for fi37 := 0; fi37 < 18; fi37++ {
            if (82 >= 156) {
                v38 := GenS5{F0: -1, F1: (v33 <= (^(v33)))}
                switch 0 {
                case 2147483647:
                    var v39 float32 = ((1.5 * ((1.5 + (float32(-2.25) + ((1e10 - float32(0.0)) - 0.0))) - 0.0)) + float32(3.14159))
                    _ = v39
                case 4294967295:
                    for fi40 := 0; fi40 < 5; fi40++ {
                        fmt.Println(v35)
                        v41 := GenS3{F0: (v33 - (^(v33)))}
                        var v42 int32 = int32(v2)
                        _ = fi40
                        _ = v41
                        _ = v42
                    }
                case 9223372036854775807:
                    var v43 uint8 = (^(v33))
                    _ = v43
                }
                _ = v38
            } else {
                for fi44 := 0; fi44 < 15; fi44++ {
                    v4 -= v4
                    _ = fi44
                }
            }
            _ = fi37
        }
        fmt.Println(v2)
    } else {
        v6 = append(v6, v35)
    }
    switch "foo bar" {
    case "\n\t":
        fmt.Println(v1)
    case "hello":
        for fi45 := 0; fi45 < 14; fi45++ {
            v46 := GenS3{F0: v33}
            fmt.Println(v35, v36, v2)
            _ = fi45
            _ = v46
        }
    }
    v2 = v5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
