package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 uint; F1 int32 }
type GenS3 struct { F0 bool; F1 int8 }
func genFn0() uint {
    var v4 uint16 = 1
    v4 += (^(v4))
    _ = v4
    return 583
}

func main() {
    var v1 int8 = 127
    v2 := []string{"foo bar", (((("" + "\n\t") + "hello") + "hello") + "foo bar")}
    var v3 int8 = (v1 ^ 42)
    switch 565 {
    case 32767:
        var v5 uint16 = 0
        _ = v5
    case -812:
        for fi6 := 0; fi6 < 12; fi6++ {
            v2 = append(v2, ("\n\t" + ("hello" + "foo bar")))
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v3 += (v1 / (((v3 | 42)) | 1))
            _ = fi6
        }
    default:
        if (18446744073709551615 > 1) {
            if false {
                var v7 uint16 = 219
                v3 = (v3 - -1)
                v8 := GenS1{F0: true, F1: 992}
                _ = v7
                _ = v8
            } else {
                v2 = append(v2, "hello")
            }
            for fi9 := 0; fi9 < 12; fi9++ {
                for fi10 := 0; fi10 < 5; fi10++ {
                    switch "foo bar" {
                    case "foo bar":
                        fmt.Println(v3, v3, v1)
                    case "hello":
                        var v11 uint = 42
                        _ = v11
                    default:
                        v12 := GenS2{F0: genFn0(), F1: -542}
                        _ = v12
                    }
                    _ = fi10
                }
                _ = fi9
            }
            var v13 int16 = -32768
            _ = v13
        }
    }
    var v14 float32 = 0.1
    v14 += (v14 - (-2.25 * -2.25))
    var v15 bool = true
    var v16 int64 = 1
    v17 := GenS3{F0: ((v1 ^ (v1 & v3)) <= v1), F1: -128}
    for fi18 := 0; fi18 < 18; fi18++ {
        v15 = v15
        _ = len(v2)
        _ = fi18
    }
    v16 -= v16
    switch 970 {
    case 475:
        for fi19 := 0; fi19 < 14; fi19++ {
            for fi20 := 0; fi20 < 12; fi20++ {
                v14 += (v14 + (-2.25 / 1e10))
                v15 = true
                var v21 uint32 = 2147483647
                _ = fi20
                _ = v21
            }
            for fi22 := 0; fi22 < 10; fi22++ {
                _ = len(v2)
                var v23 float32 = 3.14159
                _ = fi22
                _ = v23
            }
            var v24 uint8 = 255
            _ = fi19
            _ = v24
        }
    default:
        var v25 float64 = float64(v14)
        _ = v25
    }
    var v26 int64 = (v16 >> ((65535) & 63))
    var v27 int = 128
    var v28 uint32 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
}
