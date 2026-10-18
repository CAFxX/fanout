package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint8; F1 int8 }
type GenS5 struct { F0 bool; F1 uint8; F2 uint }
type GenS6 struct { F0 uint64 }
func genFn0(p0 int, p1 string) uint8 {
    var v4 int = p0
    p1 = p1
    p1 = p1
    for fi5 := 0; fi5 < 1; fi5++ {
        var v6 uint32 = 65535
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = p1
    _ = v4
    return 1
}

func main() {
    var v1 int = 2147483647
    v2 := []string{("foo bar" + (("\n\t" + ("" + "a")) + "foo bar")), "", ("" + ("foo bar" + "\n\t")), "hello"}
    var v3 string = (("foo bar" + (((((("a" + "a") + "hello") + "\n\t") + "") + "a") + ("\n\t" + ("\n\t" + "hello")))) + "foo bar")
    var v7 int32 = -2147483648
    var v8 int32 = 1
    switch "foo bar" {
    case "foo bar":
        if true {
            for fi9 := 0; fi9 < 1; fi9++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                for fi10 := 0; fi10 < 17; fi10++ {
                    switch "hello" {
                    case "foo bar":
                        var v11 uint8 = 254
                        _ = v11
                    case "a":
                        v12 := GenS1{F0: 4294967295, F1: -1}
                        var v13 uint = 0
                        _ = v12
                        _ = v13
                    case "\n\t":
                        v1 += 1
                    default:
                        var v14 float64 = 3.14159
                        _ = v14
                    }
                    if true {
                        var v15 uint32 = 4294967295
                        _ = v15
                    }
                    v7 += 0
                    _ = fi10
                }
                _ = fi9
            }
            v16 := GenS2{F0: 255, F1: 0}
            switch "" {
            case "":
                for fi17 := 0; fi17 < 16; fi17++ {
                    v2 = append(v2, (v3 + ((v3 + v3) + "hello")))
                    v2 = append(v2, (v3 + (("hello" + v3) + v3)))
                    _ = fi17
                }
            case "hello":
                switch v3 {
                case "\n\t":
                    v18 := GenS2{F0: 1, F1: -128}
                    _ = v18
                default:
                    if (653 >= 2) {
                        v8 -= v8
                        if len(v2) > 0 {
                            v2[0] = v3
                        }
                    } else {
                        for fi19 := 0; fi19 < 19; fi19++ {
                            if (32767 > -1) {
                                v20 := GenS3{F0: v3}
                                _ = v20
                            }
                            fmt.Println(v7, v8)
                            _ = fi19
                        }
                        if (100 <= 100) {
                            v21 := GenS1{F0: 42, F1: 42}
                            _ = v21
                        } else {
                            for fi22 := 0; fi22 < 13; fi22++ {
                                if false {
                                    for fi23 := 0; fi23 < 2; fi23++ {
                                        if true {
                                            if len(v2) > 0 {
                                                v2[0] = v3
                                            }
                                        } else {
                                            v24 := GenS3{F0: "\n\t"}
                                            v3 = ("foo bar" + "")
                                            _ = v24
                                        }
                                        v1 -= -9223372036854775808
                                        _ = fi23
                                    }
                                    v8 = v7
                                }
                                v3 = v3
                                _ = fi22
                            }
                            var v25 uint = 734
                            _ = v25
                        }
                    }
                }
            case "\n\t":
                v2 = append(v2, v3)
            default:
                for fi26 := 0; fi26 < 3; fi26++ {
                    v3 = v3
                    _ = fi26
                }
            }
            _ = v16
        } else {
            v3 = v3
            if true {
                var v27 string = (v3 + ("hello" + v3))
                fmt.Println(v7)
                v2 = append(v2, v27)
                _ = v27
            } else {
                fmt.Println(v1, v3)
            }
        }
    case "":
        v8 = v7
    case "hello":
        for fi28 := 0; fi28 < 9; fi28++ {
            for fi29 := 0; fi29 < 11; fi29++ {
                var v30 int16 = -561
                var v31 string = "\n\t"
                if true {
                    switch "a" {
                    case "hello":
                        fmt.Println(v8)
                    case "a":
                        var v32 float32 = 0.0
                        _ = v32
                    case "foo bar":
                        var v33 int8 = -79
                        _ = v33
                    default:
                        var v34 int32 = v7
                        _ = v34
                    }
                } else {
                    v35 := GenS4{F0: 156, F1: 72}
                    _ = v35
                }
                v36 := GenS5{F0: (!true), F1: 0, F2: 1}
                var v37 uint64 = 0
                v38 := GenS6{F0: (v37 * (v37 ^ v37))}
                _ = fi29
                _ = v30
                _ = v31
                _ = v36
                _ = v37
                _ = v38
            }
            fmt.Println(v1)
            _ = fi28
        }
    default:
        v3 = v3
    }
    if len(v2) > 0 {
        v2[0] = ("foo bar" + v3)
    }
    v8 -= (v7 + v8)
    var v39 int64 = -1
    _ = len(v2)
    v1 += v1
    var v40 uint32 = 487
    if true {
        for fi41 := 0; fi41 < 12; fi41++ {
            fmt.Println(v40, v1)
            _ = fi41
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v40
    _ = v7
    _ = v8
}
