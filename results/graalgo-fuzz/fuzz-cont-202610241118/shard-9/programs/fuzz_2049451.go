package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64; F1 uint8; F2 int16 }
func genFn0(p0 bool, p1 uint16) int64 {
    if true {
        var v7 uint8 = 255
        v8 := GenS1{F0: true, F1: 1}
        var v9 uint = 1
        _ = v7
        _ = v8
        _ = v9
    }
    var v10 int8 = 100
    fmt.Println(p1, v10)
    var v11 float64 = (0.1 / 0.1)
    switch 32768 {
    case 127:
        v12 := GenS2{F0: 1}
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return -184
}
func genFn1() uint16 {
    for fi13 := 0; fi13 < 6; fi13++ {
        if ((("hello" + ((("\n\t" + ("a" + "\n\t")) + "\n\t") + "\n\t")) + "foo bar") < "") {
            fmt.Println(fi13)
        } else {
            var v14 float64 = 0.0
            v14 = (float64(0.0) + 1e10)
            v15 := GenS2{F0: -459}
            _ = v14
            _ = v15
        }
        _ = fi13
    }
    var v16 float32 = (0.0 + 0.0)
    _ = v16
    return 255
}

func main() {
    var v1 uint64 = 0
    var v2 uint64 = 1
    v3 := []int{42, -2147483648, -2147483648, 128}
    var v4 bool = (v1 == 136)
    var v5 uint = 687
    var v6 int16 = 42
    v1 += 2
    if (2 > 101) {
        fmt.Println(v2)
        v6 = v6
        v4 = (0.0 >= 1.5)
    } else {
        if v4 {
            for fi17 := 0; fi17 < 4; fi17++ {
                if len(v3) > 0 {
                    v3[0] = (fi17 / ((779) | 1))
                }
                _ = fi17
            }
            if v4 {
                switch 32767 {
                case 1:
                    var v18 int = 0
                    _ = v18
                case 0:
                    var v19 int16 = (-(v6))
                    _ = v19
                }
                switch 9223372036854775807 {
                case 127:
                    var v20 uint16 = genFn1()
                    _ = v20
                case 292:
                    v3 = append(v3, 566)
                case 871:
                    v6 += v6
                default:
                    switch "hello" {
                    case "hello":
                        fmt.Println(v1, v5, v6)
                    case "\n\t":
                        fmt.Println(v2, v2)
                    }
                }
                if len(v3) > 0 {
                    v3[0] = 128
                }
            }
            var v21 int64 = genFn0(false, 65535)
            _ = v21
        } else {
            if ((!(!v4)) && v4) {
                v2 = (v1 - (^(v1)))
                v5 -= 65535
                switch -243 {
                case -514:
                    v6 -= (-(v6))
                default:
                    for fi22 := 0; fi22 < 9; fi22++ {
                        v6 = (v6 * v6)
                        v4 = ((v1 % (((v2 * (^(v2)))) | 1)) >= (v2 / ((v2) | 1)))
                        _ = fi22
                    }
                }
            }
        }
    }
    v1 = 203
    var v23 uint16 = 344
    var v24 int8 = 42
    for fi25 := 0; fi25 < 17; fi25++ {
        v26 := GenS3{F0: (^(v1)), F1: 1, F2: 32767}
        _ = fi25
        _ = v26
    }
    v27 := GenS2{F0: 9223372036854775807}
    v28 := GenS2{F0: 9223372036854775807}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
