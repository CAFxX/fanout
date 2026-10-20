package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 bool }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 int; F1 uint8 }
type GenS4 struct { F0 string; F1 string; F2 bool }

func main() {
    var v1 int8 = 127
    v2 := []uint32{282, 4294967295, 736, 65535}
    var v3 bool = (!(v1 < -128))
    v4 := make(map[uint32]string)
    var v5 uint16 = 65535
    for fi6 := 0; fi6 < 18; fi6++ {
        v3 = v3
        _ = fi6
    }
    for fi7 := 0; fi7 < 14; fi7++ {
        v8 := GenS1{F0: -5, F1: 692, F2: v3}
        var v9 int64 = 9223372036854775807
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := GenS2{F0: 2147483647, F1: -128}
    v11 := GenS2{F0: 0, F1: v1}
    var v12 uint = 4294967295
    var v13 int64 = int64(v5)
    var v14 int = 484
    var v15 float64 = float64(1e10)
    switch "\n\t" {
    case "":
        for fi16 := 0; fi16 < 10; fi16++ {
            if (((3.14159 * float32(float64(-2.25))) + float32(((1.5 - (v15 * (v15 / 1.0))) + v15))) != (0.1 * 1.5)) {
                var v17 bool = ("" <= "hello")
                _ = v17
            } else {
                v1 = (v1 + v1)
                v2 = append(v2, 985)
                for fi18 := 0; fi18 < 8; fi18++ {
                    switch "a" {
                    case "a":
                        v19 := GenS2{F0: 42, F1: v1}
                        v20 := GenS1{F0: (v1 >> (((v12 ^ v12)) & 63)), F1: 2, F2: (v13 < v13)}
                        v13 += v13
                        _ = v19
                        _ = v20
                    default:
                        _ = len(v4)
                    }
                    v1 -= v1
                    _ = fi18
                }
            }
            if v3 {
                if (true && (4294967295 > 18446744073709551615)) {
                    var v21 int16 = int16(v12)
                    for fi22 := 0; fi22 < 12; fi22++ {
                        v23 := GenS2{F0: 2147483647, F1: (v1 | (v1 * (v1 & (^(v1)))))}
                        _ = fi22
                        _ = v23
                    }
                    v4[4294967295] = "\n\t"
                    _ = v21
                }
            } else {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = len(v4)
            }
            _ = fi16
        }
    }
    v14 += 367
    for fi24 := 0; fi24 < 10; fi24++ {
        for fi25 := 0; fi25 < 6; fi25++ {
            if (!(!((-(v13)) == (v13 & (v13 - v13))))) {
                var v26 string = (("a" + "foo bar") + (("foo bar" + ("hello" + "\n\t")) + ((("a" + ((("a" + "") + ("hello" + "foo bar")) + "a")) + "\n\t") + "a")))
                if len(v2) > 0 {
                    v2[0] = 4294967295
                }
                v27 := v4[1]
                _ = v26
                _ = v27
            } else {
                fmt.Println(v5)
                if (v3 || (!(0.0 == 0.1))) {
                    fmt.Println(v15, v5)
                    if (128 == 42) {
                        v28 := GenS3{F0: fi25, F1: 255}
                        v15 += v15
                        if (v3 && v3) {
                            fmt.Println(v3)
                        }
                        _ = v28
                    }
                } else {
                    fmt.Println(v5)
                    v29 := GenS4{F0: "hello", F1: "\n\t", F2: (!(2 == 42))}
                    var v30 string = "hello"
                    _ = v29
                    _ = v30
                }
            }
            _ = fi25
        }
        _ = fi24
    }
    var v31 int16 = -254
    var v32 bool = (!(!v3))
    v12 = v12
    v1 = 100
    v13 = 1
    var v33 int64 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
