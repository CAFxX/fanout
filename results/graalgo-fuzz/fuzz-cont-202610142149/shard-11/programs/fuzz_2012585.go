package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 int32 }
type GenS2 struct { F0 int64; F1 uint }
type GenS3 struct { F0 uint64 }
func genFn0() int64 {
    var v7 uint16 = 0
    v7 = (^(v7))
    _ = v7
    return -1
}
func genFn1(p0 uint, p1 uint16, p2 int) uint {
    var v8 uint64 = 0
    switch "\n\t" {
    case "hello":
        fmt.Println(p0)
    case "foo bar":
        p2 = (-(p2))
    case "":
        var v9 uint8 = 128
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return p0
}

func main() {
    var v1 int16 = 0
    var v2 uint16 = 65535
    v3 := []uint32{2147483648, 112, 2147483647, 2}
    var v4 int8 = 42
    var v5 uint64 = 18446744073709551615
    v6 := make(map[string]uint32)
    fmt.Println(v5, v1, v4)
    var v10 int32 = 2147483647
    if (((false || true) && (-993 > -546)) && (0 < 255)) {
        _ = len(v3)
        v11 := GenS1{F0: v1, F1: "\n\t", F2: (v10 | 2147483647)}
        _ = v11
    } else {
        if len(v3) > 0 {
            v3[0] = 42
        }
        for fi12 := 0; fi12 < 17; fi12++ {
            var v13 float32 = 1e10
            for fi14 := 0; fi14 < 9; fi14++ {
                switch fi14 {
                case 127:
                    if (false && ((v4 / ((v4) | 1)) == (^(v4)))) {
                        var v15 uint16 = (v2 / (((^(v2))) | 1))
                        _ = len(v3)
                        _ = v15
                    } else {
                        v16 := GenS1{F0: -770, F1: "hello", F2: v10}
                        v17 := GenS1{F0: (v1 + (v1 + (v1 & (v1 + 710)))), F1: "foo bar", F2: v10}
                        _ = v16
                        _ = v17
                    }
                }
                var v18 int = fi12
                switch "hello" {
                case "\n\t":
                    v19 := GenS1{F0: -438, F1: "foo bar", F2: v10}
                    _ = v19
                }
                v20 := v6[((("a" + (("\n\t" + "foo bar") + "")) + "hello") + "a")]
                _ = fi14
                _ = v18
                _ = v20
            }
            _ = fi12
            _ = v13
        }
        for fi21 := 0; fi21 < 6; fi21++ {
            for fi22 := 0; fi22 < 12; fi22++ {
                v23 := GenS1{F0: v1, F1: ("a" + ("hello" + ("foo bar" + "a"))), F2: v10}
                _ = fi22
                _ = v23
            }
            _ = fi21
        }
    }
    var v24 uint16 = (v2 << ((42) & 63))
    v25 := GenS2{F0: genFn0(), F1: 42}
    var v26 uint = 1
    v5 = v5
    for fi27 := 0; fi27 < 7; fi27++ {
        if false {
            for fi28 := 0; fi28 < 18; fi28++ {
                fmt.Println(v24)
                switch fi27 {
                case 260:
                    if (v4 == v4) {
                        v29 := GenS3{F0: v5}
                        if (83 < -55) {
                            if false {
                                v30 := GenS3{F0: (v5 ^ 572)}
                                v2 = 32767
                                _ = v30
                            }
                        }
                        _ = v29
                    }
                case -510:
                    if (!(fi27 > (fi28 + (fi27 * (^(fi28)))))) {
                        v2 += 255
                        v24 -= 2
                    }
                case 42:
                    var v31 uint16 = v2
                    _ = v31
                }
                _ = fi28
            }
        } else {
            var v32 int = 4294967295
            _ = v32
        }
        var v33 int64 = 841
        _ = fi27
        _ = v33
    }
    var v34 int32 = v10
    var v35 uint16 = 1
    fmt.Println(v10, v34, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
