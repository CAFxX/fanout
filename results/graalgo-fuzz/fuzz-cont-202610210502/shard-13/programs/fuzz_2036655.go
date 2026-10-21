package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16; F1 bool }
func genFn0() int32 {
    if (0.0 > ((((0.1 / 1.0) / 1.5) * 0.0) - (3.14159 / ((1e10 - (3.14159 / 1.0)) / (1e10 - (1e10 - 0.1)))))) {
        var v6 bool = true
        _ = v6
    }
    v7 := GenS1{F0: (2 == 42)}
    var v8 string = "foo bar"
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 uint) uint {
    fmt.Println(p0)
    if (true || (!true)) {
        fmt.Println(p0)
    }
    var v9 uint8 = 1
    _ = p0
    _ = v9
    return 2
}

func main() {
    var v1 uint = 636
    var v2 float32 = ((0.0 + float32(0.0)) + 1.5)
    v3 := make(map[uint32]uint32)
    v4 := []string{"\n\t", ("" + ""), ""}
    var v5 bool = false
    if v5 {
        if (!v5) {
            switch "hello" {
            case "\n\t":
                switch -910 {
                case 127:
                    var v10 float64 = 1.5
                    _ = v10
                case -1:
                    var v11 int = 725
                    _ = v11
                case 9223372036854775807:
                    _ = len(v3)
                default:
                    var v12 uint32 = 2147483647
                    _ = v12
                }
            }
        }
        v13 := v3[2147483648]
        _ = v13
    }
    v14 := GenS1{F0: true}
    var v15 int64 = 42
    _ = len(v4)
    fmt.Println(v15, v2, v5)
    var v16 uint8 = 42
    v17 := v3[0]
    fmt.Println(v5, v2, v1)
    var v18 uint8 = (^(v16))
    switch -1 {
    case 128:
        v19 := GenS1{F0: ("foo bar" != "a")}
        _ = v19
    case -835:
        v20 := GenS1{F0: (v1 > v1)}
        _ = v20
    default:
        var v21 int64 = v15
        _ = v21
    }
    if v5 {
        switch "a" {
        case "hello":
            if (v5 || (("" + "\n\t") <= (("" + "a") + "\n\t"))) {
                var v22 float64 = 3.14159
                _ = v22
            } else {
                fmt.Println(v5, v16, v2)
                if false {
                    if v5 {
                        v17 -= 481
                        v4 = append(v4, "hello")
                        var v23 int = 0
                        _ = v23
                    } else {
                        v1 += v1
                        for fi24 := 0; fi24 < 12; fi24++ {
                            for fi25 := 0; fi25 < 16; fi25++ {
                                v26 := GenS1{F0: (!(v2 < v2))}
                                _ = fi25
                                _ = v26
                            }
                            fmt.Println(v2, v15, v5)
                            v5 = false
                            _ = fi24
                        }
                        v27 := GenS2{F0: 411, F1: ((!(v17 <= (v17 + uint32(v1)))) && v5)}
                        _ = v27
                    }
                    fmt.Println(v17, v15)
                }
                v28 := GenS1{F0: v5}
                _ = v28
            }
        case "foo bar":
            switch -156 {
            case 42:
                _ = len(v3)
            case 127:
                switch ("hello" + ("a" + "a")) {
                case "a":
                    if len(v4) > 0 {
                        v4[0] = ("a" + "")
                    }
                }
            case 4294967295:
                _ = len(v3)
            }
        }
        fmt.Println(v1, v16)
    } else {
        if len(v4) > 0 {
            v4[0] = "\n\t"
        }
        v29 := GenS2{F0: 255, F1: v5}
        v16 -= v18
        _ = v29
    }
    var v30 float32 = v2
    v17 = (v17 / ((v17) | 1))
    fmt.Println(v16)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
