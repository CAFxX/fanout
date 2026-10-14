package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16; F1 int64 }
func genFn0(p0 bool) bool {
    if (42 <= 630) {
        var v4 float32 = (0.0 / 1.0)
        _ = v4
    } else {
        var v5 uint = 42
        for fi6 := 0; fi6 < 8; fi6++ {
            for fi7 := 0; fi7 < 19; fi7++ {
                v8 := GenS1{F0: (fi7 << ((65535) & 63)), F1: -9223372036854775808, F2: 0}
                fmt.Println(fi6, fi6, v5)
                _ = fi7
                _ = v8
            }
            _ = fi6
        }
        var v9 int16 = 1000
        _ = v5
        _ = v9
    }
    v10 := GenS1{F0: 1, F1: 1, F2: 707}
    _ = p0
    _ = v10
    return (false || p0)
}

func main() {
    var v1 int32 = 2147483647
    var v2 float64 = 1.5
    v3 := []uint32{}
    var v11 int64 = 1
    if true {
        if false {
            var v12 float64 = (v2 * (0.1 + v2))
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = v12
        } else {
            if false {
                if false {
                    v13 := GenS1{F0: 891, F1: v11, F2: v11}
                    if ((!false) || true) {
                        v1 = v1
                    }
                    fmt.Println(v11, v1, v11)
                    _ = v13
                }
                _ = len(v3)
            } else {
                if len(v3) > 0 {
                    v3[0] = 2147483647
                }
                var v14 uint16 = 395
                _ = v14
            }
            v3 = append(v3, 2147483647)
            v1 += 16
        }
    } else {
        fmt.Println(v11, v2, v2)
        v15 := GenS2{F0: ((((("hello" + "foo bar") + "\n\t") + "") + "") + ("\n\t" + ("foo bar" + "foo bar")))}
        var v16 int16 = 1000
        _ = v15
        _ = v16
    }
    if true {
        v17 := GenS2{F0: ("" + "")}
        _ = v17
    } else {
        switch "" {
        case "":
            fmt.Println(v2, v2)
        case "a":
            var v18 string = ""
            _ = v18
        case "foo bar":
            for fi19 := 0; fi19 < 8; fi19++ {
                switch "foo bar" {
                case "a":
                    for fi20 := 0; fi20 < 7; fi20++ {
                        var v21 int = -9223372036854775808
                        _ = len(v3)
                        _ = fi20
                        _ = v21
                    }
                }
                if false {
                    fmt.Println(v2)
                    if false {
                        v1 += 1
                        v22 := GenS1{F0: 2147483647, F1: (v11 & 770), F2: v11}
                        var v23 int8 = -65
                        _ = v22
                        _ = v23
                    }
                } else {
                    switch fi19 {
                    case -9223372036854775808:
                        v24 := GenS3{F0: 552, F1: (^(v11))}
                        _ = v24
                    case -431:
                        v2 -= (v2 + (0.1 / (v2 / 1.5)))
                    }
                    v25 := GenS2{F0: ""}
                    for fi26 := 0; fi26 < 16; fi26++ {
                        v11 = (-(v11))
                        v11 = (v11 | (v11 * v11))
                        var v27 bool = false
                        _ = fi26
                        _ = v27
                    }
                    _ = v25
                }
                _ = fi19
            }
        default:
            var v28 uint16 = 0
            _ = v28
        }
    }
    var v29 int = 0
    var v30 uint64 = 18446744073709551615
    if len(v3) > 0 {
        v3[0] = 42
    }
    v1 = -2147483648
    var v31 uint16 = 542
    switch v29 {
    case 4294967295:
        v32 := GenS2{F0: "hello"}
        _ = v32
    case -969:
        v2 += float64(1e10)
    case 9223372036854775807:
        for fi33 := 0; fi33 < 19; fi33++ {
            var v34 int32 = v1
            if true {
                fmt.Println(v11)
                v35 := GenS2{F0: ("hello" + "hello")}
                var v36 uint16 = v31
                _ = v35
                _ = v36
            }
            _ = fi33
            _ = v34
        }
    }
    var v37 uint32 = uint32(v30)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v11
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v37
}
