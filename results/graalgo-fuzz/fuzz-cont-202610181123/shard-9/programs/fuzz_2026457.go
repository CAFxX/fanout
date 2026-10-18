package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 uint32 }
type GenS2 struct { F0 uint8; F1 bool; F2 uint16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint32; F1 int64; F2 int }
type GenS5 struct { F0 int64; F1 string; F2 uint8 }
func genFn0(p0 int32, p1 uint) uint16 {
    v7 := GenS1{F0: -67, F1: 2, F2: 65535}
    if false {
        fmt.Println(p1)
    }
    _ = p0
    _ = p1
    _ = v7
    return 65535
}

func main() {
    var v1 bool = false
    var v2 float64 = (-2.25 / -2.25)
    v3 := []uint32{2, 65535}
    v4 := make(map[int]string)
    v5 := []string{("\n\t" + "a")}
    v6 := []string{(("\n\t" + "foo bar") + ("a" + ((("\n\t" + ("hello" + ((("" + "") + ("hello" + "foo bar")) + "a"))) + "hello") + "foo bar"))), ("hello" + "a"), "hello"}
    var v8 uint8 = 2
    switch 0 {
    case 1:
        if len(v3) > 0 {
            v3[0] = 1
        }
    case -5:
        switch "a" {
        case "":
            switch "foo bar" {
            case "":
                if false {
                    v9 := GenS1{F0: 1, F1: 656, F2: 65535}
                    switch "" {
                    case "foo bar":
                        for fi10 := 0; fi10 < 13; fi10++ {
                            v11 := GenS1{F0: -123, F1: 911, F2: 4294967295}
                            for fi12 := 0; fi12 < 3; fi12++ {
                                fmt.Println(v8, v1, v2)
                                v6 = append(v6, ("" + ("\n\t" + "hello")))
                                _ = fi12
                            }
                            _ = fi10
                            _ = v11
                        }
                    case "":
                        fmt.Println(v8, v8, v1)
                    default:
                        for fi13 := 0; fi13 < 13; fi13++ {
                            var v14 int16 = -834
                            for fi15 := 0; fi15 < 17; fi15++ {
                                fmt.Println(v1)
                                var v16 int16 = (v14 / ((v14) | 1))
                                _ = fi15
                                _ = v16
                            }
                            _ = fi13
                            _ = v14
                        }
                    }
                    _ = v9
                } else {
                    for fi17 := 0; fi17 < 2; fi17++ {
                        if (!v1) {
                            fmt.Println(v1, fi17, v8)
                        } else {
                            _ = len(v4)
                            v18 := GenS2{F0: 1, F1: (v1 || false), F2: 42}
                            for fi19 := 0; fi19 < 16; fi19++ {
                                var v20 uint32 = 145
                                v3 = append(v3, (v20 + (v20 * v20)))
                                _ = fi19
                                _ = v20
                            }
                            _ = v18
                        }
                        _ = fi17
                    }
                }
            case "hello":
                delete(v4, 42)
            default:
                v8 = (v8 + (v8 / ((v8) | 1)))
            }
        case "a":
            v2 = v2
        case "\n\t":
            v2 = -2.25
        default:
            for fi21 := 0; fi21 < 17; fi21++ {
                if (v1 || ((0.0 >= -2.25) || true)) {
                    var v22 int64 = 102
                    var v23 int16 = int16(fi21)
                    _ = v22
                    _ = v23
                }
                fmt.Println(v2)
                _ = fi21
            }
        }
    }
    v24 := GenS2{F0: (^(v8)), F1: v1, F2: 2}
    fmt.Println(v1)
    v5 = append(v5, (("\n\t" + "") + "\n\t"))
    for fi25 := 0; fi25 < 5; fi25++ {
        delete(v4, 42)
        _ = fi25
    }
    fmt.Println(v2)
    fmt.Println(v1, v2, v2)
    fmt.Println(v8, v8)
    v26 := GenS3{F0: 0}
    for fi27 := 0; fi27 < 1; fi27++ {
        v28 := GenS4{F0: 42, F1: -9223372036854775808, F2: (fi27 & fi27)}
        _ = fi27
        _ = v28
    }
    v8 -= v8
    fmt.Println(v1, v1, v1)
    var v29 uint64 = 226
    v30 := v4[127]
    v31 := GenS1{F0: -1, F1: 765, F2: 20}
    v32 := GenS5{F0: -781, F1: (v30 + v30), F2: (v8 >> ((273) & 63))}
    var v33 int = int(v8)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v24
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
