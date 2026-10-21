package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 int32; F1 uint16 }

func main() {
    var v1 int32 = 328
    v2 := []uint32{42, 278, 1}
    v3 := make(map[uint32]int)
    var v4 uint64 = 4294967295
    var v5 int8 = 9
    var v6 uint16 = 42
    v5 = -51
    var v7 bool = true
    v1 += (v1 % ((-1) | 1))
    var v8 uint64 = v4
    v9 := v3[2147483647]
    _ = len(v3)
    for fi10 := 0; fi10 < 5; fi10++ {
        v11 := GenS1{F0: 32767}
        if true {
            fmt.Println(fi10)
            if v7 {
                v12 := GenS2{F0: 128, F1: 1}
                switch "" {
                case "a":
                    if (65535 <= 686) {
                        switch ("\n\t" + "foo bar") {
                        case "":
                            v8 = v8
                        }
                    } else {
                        v4 -= (v8 / ((v8) | 1))
                    }
                case "foo bar":
                    v6 -= v6
                case "hello":
                    if len(v2) > 0 {
                        v2[0] = 200
                    }
                default:
                    v13 := GenS3{F0: -1, F1: (v6 >> ((0) & 63))}
                    _ = v13
                }
                v7 = true
                _ = v12
            }
            if (v7 && (v5 < (v5 * v5))) {
                var v14 float64 = 0.1
                _ = v14
            } else {
                v7 = (v7 || true)
            }
        }
        v15 := GenS1{F0: int(v5)}
        _ = fi10
        _ = v11
        _ = v15
    }
    for fi16 := 0; fi16 < 14; fi16++ {
        v17 := v3[4294967295]
        switch "" {
        case "hello":
            v8 = (v4 - (v8 - 1))
        default:
            if v7 {
                if (!v7) {
                    var v18 float64 = (-2.25 / 0.1)
                    for fi19 := 0; fi19 < 5; fi19++ {
                        _ = len(v2)
                        _ = fi19
                    }
                    for fi20 := 0; fi20 < 12; fi20++ {
                        var v21 uint16 = 786
                        _ = fi20
                        _ = v21
                    }
                    _ = v18
                }
                var v22 int16 = 507
                switch v9 {
                case 360:
                    v4 = 1
                case 9223372036854775807:
                    v9 += 989
                }
                _ = v22
            }
        }
        _ = fi16
        _ = v17
    }
    var v23 int32 = -506
    v9 = 42
    v8 = v8
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
