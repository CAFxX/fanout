package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool; F1 string; F2 bool }
func genFn0() bool {
    v7 := GenS1{F0: 100}
    var v8 uint8 = 128
    _ = v7
    _ = v8
    return true
}
func genFn1(p0 uint, p1 int64, p2 uint64) bool {
    if (!false) {
        p0 = uint(p2)
    }
    var v9 uint = (^(p0))
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return false
}

func main() {
    var v1 string = "foo bar"
    v2 := make(map[string]bool)
    var v3 uint8 = 25
    v4 := []string{((v1 + v1) + v1), v1}
    v5 := []uint32{491, 65535, 420, 42}
    var v6 int8 = -128
    if true {
        v3 = v3
        var v10 float64 = 0.0
        _ = v10
    } else {
        if (42 < 798) {
            switch -268 {
            case 0:
                var v11 string = (("a" + (v1 + v1)) + (("\n\t" + v1) + "foo bar"))
                _ = v11
            case -9223372036854775808:
                for fi12 := 0; fi12 < 1; fi12++ {
                    switch 128 {
                    case 32768:
                        fmt.Println(v3)
                    default:
                        if false {
                            _ = len(v4)
                        } else {
                            fmt.Println(v3)
                        }
                    }
                    _ = fi12
                }
            case 916:
                v3 = v3
            default:
                var v13 uint64 = 4294967295
                _ = v13
            }
            fmt.Println(v6, v1)
        }
        var v14 bool = (!(true && false))
        _ = v14
    }
    var v15 int16 = -32768
    var v16 uint16 = 65535
    var v17 int64 = 9223372036854775807
    var v18 float32 = 0.0
    v19 := GenS1{F0: (v6 % ((v6) | 1))}
    var v20 int16 = v15
    if false {
        if (!(true || false)) {
            v2[v1] = true
        } else {
            _ = len(v2)
            var v21 uint32 = 0
            _ = v21
        }
    } else {
        v22 := v2[v1]
        _ = v22
    }
    v15 -= v20
    for fi23 := 0; fi23 < 9; fi23++ {
        if (18446744073709551615 > 0) {
            for fi24 := 0; fi24 < 17; fi24++ {
                v15 = int16(fi24)
                switch (fi24 ^ fi23) {
                case 603:
                    var v25 uint16 = v16
                    _ = v25
                case 0:
                    switch (-(fi23)) {
                    case 0:
                        v26 := GenS1{F0: v6}
                        _ = v26
                    case 32768:
                        v1 = ("foo bar" + v1)
                    default:
                        switch "hello" {
                        case "foo bar":
                            var v27 uint8 = v3
                            _ = v27
                        default:
                            v28 := GenS1{F0: 1}
                            v29 := GenS2{F0: true, F1: ((v1 + "foo bar") + "hello"), F2: (false && true)}
                            _ = len(v2)
                            _ = v28
                            _ = v29
                        }
                    }
                    fmt.Println(v18, v6, v3)
                    v3 += 42
                    v5 = append(v5, 0)
                    if (!true) {
                        if true {
                            v18 += v18
                            v6 -= (v6 - v6)
                            var v30 uint32 = 667
                            _ = v30
                        } else {
                            if len(v4) > 0 {
                                v4[0] = ("a" + "foo bar")
                            }
                            var v31 int16 = v15
                            v6 = v6
                            _ = v31
                        }
                        switch (fi24 >> ((993) & 63)) {
                        case 808:
                            var v32 float32 = ((v18 * v18) - (float32(3.14159) - v18))
                            _ = v32
                        case 2147483647:
                            var v33 int32 = 2147483647
                            _ = v33
                        case 880:
                            v18 -= float32((1e10 - (1e10 / 1e10)))
                        }
                    }
                case -2147483648:
                    if false {
                        v16 += 255
                        var v34 float64 = (-2.25 - (0.0 - 0.0))
                        _ = v34
                    } else {
                        var v35 uint64 = 18446744073709551615
                        v18 -= 1.5
                        v35 -= (v35 >> ((318) & 63))
                        _ = v35
                    }
                }
                _ = fi24
            }
            v36 := GenS2{F0: true, F1: ("hello" + "\n\t"), F2: false}
            _ = v36
        }
        _ = fi23
    }
    var v37 uint = 18446744073709551615
    fmt.Println(v17, v37)
    v38 := GenS1{F0: int8(v3)}
    var v39 uint32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
