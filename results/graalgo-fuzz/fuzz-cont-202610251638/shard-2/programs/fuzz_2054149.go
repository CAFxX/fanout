package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 int8; F1 int }
type GenS3 struct { F0 string; F1 int32 }

func main() {
    var v1 string = ((("" + "") + "a") + ("\n\t" + "hello"))
    var v2 string = v1
    v3 := make(map[uint32]bool)
    v4 := make(map[string]uint32)
    if true {
        if true {
            for fi5 := 0; fi5 < 15; fi5++ {
                if false {
                    v2 = v1
                }
                var v6 uint = 0
                _ = fi5
                _ = v6
            }
            for fi7 := 0; fi7 < 6; fi7++ {
                v8 := GenS1{F0: 65535, F1: (4294967295 <= 42)}
                v3[0] = (float32(-2.25) <= 0.1)
                _ = fi7
                _ = v8
            }
        } else {
            for fi9 := 0; fi9 < 12; fi9++ {
                for fi10 := 0; fi10 < 14; fi10++ {
                    switch v1 {
                    case "foo bar":
                        var v11 int16 = 1000
                        _ = v11
                    case "\n\t":
                        v2 = ("hello" + v2)
                    case "a":
                        v2 = (v2 + v1)
                    }
                    if (-1 < 0) {
                        v4[((v2 + v1) + "a")] = 2147483647
                    } else {
                        switch 554 {
                        case 9223372036854775807:
                            var v12 bool = (1000 < -27)
                            _ = v12
                        default:
                            v13 := GenS1{F0: 2, F1: (!(0.1 > -2.25))}
                            _ = v13
                        }
                    }
                    if (!false) {
                        _ = len(v3)
                    } else {
                        v14 := []int{(^(fi10)), (fi9 - -2147483648), 0}
                        var v15 int16 = 1
                        _ = v14
                        _ = v15
                    }
                    v2 = v1
                    _ = fi10
                }
                _ = fi9
            }
        }
        v16 := []int{128}
        _ = v16
    }
    v2 = "foo bar"
    v17 := GenS2{F0: 127, F1: 574}
    v18 := GenS3{F0: (v1 + v1), F1: -2147483648}
    v2 = "foo bar"
    v1 = ("hello" + v2)
    v1 = "foo bar"
    var v19 float64 = 3.14159
    switch v2 {
    case "a":
        var v20 int64 = -983
        _ = v20
    case "\n\t":
        v1 = v1
    }
    var v21 uint16 = 270
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
