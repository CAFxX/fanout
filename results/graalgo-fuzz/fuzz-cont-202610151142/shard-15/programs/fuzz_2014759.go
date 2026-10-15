package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 int64; F1 int8; F2 int16 }

func main() {
    var v1 bool = true
    var v2 float64 = 0.1
    var v3 int = 4294967295
    v4 := []string{"\n\t", ("a" + "a"), "a", "\n\t"}
    v1 = false
    if len(v4) > 0 {
        v4[0] = ("a" + "\n\t")
    }
    v5 := GenS1{F0: 0, F1: 1000}
    v6 := GenS1{F0: 747, F1: 574}
    v2 = v2
    var v7 uint8 = 255
    v1 = (!v1)
    var v8 int = (v3 + v3)
    var v9 int16 = 469
    for fi10 := 0; fi10 < 10; fi10++ {
        switch 32767 {
        case 52:
            switch "hello" {
            case "foo bar":
                if v1 {
                    var v11 int32 = -316
                    switch ((("\n\t" + "") + (("" + "\n\t") + "")) + "foo bar") {
                    case "hello":
                        if v1 {
                            v7 += (^(v7))
                            var v12 int32 = v11
                            _ = v12
                        }
                    case "\n\t":
                        fmt.Println(v3)
                    default:
                        if (v1 && (v9 == v9)) {
                            for fi13 := 0; fi13 < 2; fi13++ {
                                v8 -= -819
                                switch fi13 {
                                case -2147483648:
                                    v14 := GenS1{F0: 42, F1: v9}
                                    _ = v14
                                case -136:
                                    v3 = fi10
                                case 9223372036854775807:
                                    var v15 float64 = 0.0
                                    _ = v15
                                default:
                                    v16 := GenS1{F0: v11, F1: 32767}
                                    _ = v16
                                }
                                var v17 float64 = (1.5 + v2)
                                _ = fi13
                                _ = v17
                            }
                        } else {
                            v7 = 0
                        }
                    }
                    _ = v11
                } else {
                    var v18 float32 = 0.1
                    switch "foo bar" {
                    case "\n\t":
                        v19 := GenS2{F0: -164, F1: "a"}
                        _ = v19
                    }
                    var v20 uint16 = 2
                    _ = v18
                    _ = v20
                }
            }
        case -936:
            for fi21 := 0; fi21 < 7; fi21++ {
                fmt.Println(v2)
                var v22 float64 = (v2 / v2)
                _ = fi21
                _ = v22
            }
        case 234:
            var v23 uint16 = 596
            _ = v23
        }
        var v24 uint32 = 42
        _ = fi10
        _ = v24
    }
    if (1.5 >= (v2 + -2.25)) {
        for fi25 := 0; fi25 < 15; fi25++ {
            v26 := GenS3{F0: -9223372036854775808, F1: 0, F2: (v9 << ((42) & 63))}
            if v1 {
                v9 += (v9 - (v9 + -1))
            }
            v1 = ((((float64((((0.1 / ((-2.25 + float32((((v2 / 3.14159) - 0.1) - v2))) / (1e10 * 1e10))) - (((0.1 + ((-2.25 * float32((v2 + 3.14159))) / 1.5)) / (3.14159 - (-2.25 / (0.1 / float32((((v2 - v2) * float64(1.5)) * v2)))))) - (((1.5 * 1e10) * 1e10) / 0.1))) * (0.0 - 1.5))) * float64(0.0)) * (1e10 / v2)) + (((3.14159 - v2) + v2) + v2)) != 3.14159)
            _ = fi25
            _ = v26
        }
        var v27 int16 = v9
        _ = v27
    }
    _ = len(v4)
    var v28 int8 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
