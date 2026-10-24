package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int32; F1 uint }

func main() {
    v1 := []uint32{288, 867}
    var v2 uint16 = 448
    var v3 float32 = (float32((1e10 * -2.25)) + -2.25)
    var v4 float64 = -2.25
    var v5 uint8 = 1
    var v6 string = ("hello" + ("\n\t" + ("a" + "foo bar")))
    fmt.Println(v2)
    v7 := GenS1{F0: (v2 + (v2 * 65535)), F1: 0}
    v2 = 1
    for fi8 := 0; fi8 < 14; fi8++ {
        switch (-(fi8)) {
        case 1:
            var v9 uint16 = v2
            _ = v9
        case 2147483647:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case -418:
            var v10 uint8 = 2
            _ = v10
        default:
            if false {
                var v11 int16 = -1
                var v12 int16 = -1
                _ = v11
                _ = v12
            }
        }
        _ = fi8
    }
    switch 42 {
    case -9223372036854775808:
        v13 := GenS2{F0: 0}
        _ = v13
    case 1:
        switch 128 {
        case 42:
            if true {
                v1 = append(v1, 2147483648)
                if len(v1) > 0 {
                    v1[0] = 575
                }
            } else {
                for fi14 := 0; fi14 < 16; fi14++ {
                    v2 = (v2 - v2)
                    _ = fi14
                }
                v2 = v2
                for fi15 := 0; fi15 < 17; fi15++ {
                    v16 := GenS3{F0: 42}
                    switch "foo bar" {
                    case "hello":
                        v17 := GenS3{F0: -128}
                        _ = v17
                    case "foo bar":
                        if (1.5 == v4) {
                            v4 += (float64((v3 * float32(float64(3.14159)))) * v4)
                            v4 = v4
                            v18 := GenS3{F0: -128}
                            _ = v18
                        } else {
                            v19 := GenS4{F0: 0, F1: 18446744073709551615}
                            v20 := GenS1{F0: 837, F1: 4294967295}
                            _ = v19
                            _ = v20
                        }
                    default:
                        v1 = append(v1, 4294967295)
                    }
                    _ = fi15
                    _ = v16
                }
            }
        }
    }
    v4 -= (1.5 - (-2.25 / 1.0))
    var v21 int = -1
    v22 := GenS2{F0: -32768}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
