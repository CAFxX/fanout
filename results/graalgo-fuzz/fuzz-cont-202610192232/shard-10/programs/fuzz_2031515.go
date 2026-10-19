package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 uint32 }
type GenS2 struct { F0 int; F1 bool }
type GenS3 struct { F0 int8; F1 int32 }
type GenS4 struct { F0 string; F1 bool }
type GenS5 struct { F0 bool; F1 uint64 }
func genFn0(p0 int8, p1 string) uint64 {
    fmt.Println(p1, p1)
    if false {
        var v7 uint16 = 806
        _ = v7
    } else {
        for fi8 := 0; fi8 < 3; fi8++ {
            for fi9 := 0; fi9 < 18; fi9++ {
                var v10 float32 = 0.0
                p1 = ("hello" + ("hello" + p1))
                if (true && true) {
                    v11 := GenS1{F0: 0, F1: fi9, F2: 65535}
                    _ = v11
                }
                v12 := GenS2{F0: (fi8 - (fi8 * (fi9 & (fi9 ^ fi8)))), F1: true}
                _ = fi9
                _ = v10
                _ = v12
            }
            _ = fi8
        }
        fmt.Println(p0)
        var v13 int8 = p0
        _ = v13
    }
    switch 9223372036854775807 {
    case 867:
        v14 := []int{}
        _ = v14
    case 32768:
        v15 := GenS2{F0: 4294967295, F1: true}
        _ = v15
    case -9223372036854775808:
        p1 = p1
    }
    p0 = p0
    _ = p0
    _ = p1
    return 269
}

func main() {
    var v1 float32 = 0.1
    var v2 int8 = -76
    v3 := []uint32{874, 2147483647, 1, 2147483648}
    var v4 int16 = 32767
    var v5 string = ("hello" + "\n\t")
    var v6 uint16 = 928
    for fi16 := 0; fi16 < 8; fi16++ {
        for fi17 := 0; fi17 < 13; fi17++ {
            if true {
                if false {
                    var v18 uint = 65535
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    if (0 > 2) {
                        v19 := GenS3{F0: (^(v2)), F1: 667}
                        _ = v19
                    }
                    if len(v3) > 0 {
                        v3[0] = 879
                    }
                    _ = v18
                }
                v2 = -30
                var v20 bool = (!(-9223372036854775808 >= 42))
                _ = v20
            }
            var v21 int8 = v2
            _ = fi17
            _ = v21
        }
        _ = fi16
    }
    v22 := GenS4{F0: v5, F1: (!true)}
    _ = len(v3)
    for fi23 := 0; fi23 < 6; fi23++ {
        if (198 != 1) {
            v1 -= ((v1 + (v1 - (1e10 * (v1 / 1.0)))) + (float32(float64(v1)) * (float32((3.14159 / 3.14159)) / v1)))
        } else {
            for fi24 := 0; fi24 < 20; fi24++ {
                v25 := GenS3{F0: -128, F1: 0}
                v26 := GenS5{F0: (false && false), F1: 18446744073709551615}
                var v27 uint = 1
                _ = fi24
                _ = v25
                _ = v26
                _ = v27
            }
        }
        _ = fi23
    }
    if len(v3) > 0 {
        v3[0] = 4294967295
    }
    for fi28 := 0; fi28 < 12; fi28++ {
        fmt.Println(v1, v1, fi28)
        fmt.Println(v6, fi28)
        var v29 uint = 1
        _ = fi28
        _ = v29
    }
    switch v5 {
    case "":
        for fi30 := 0; fi30 < 9; fi30++ {
            v2 += 127
            _ = fi30
        }
    case "\n\t":
        v5 = "a"
    case "foo bar":
        v6 += (v6 % (((v6 / (((^(v6))) | 1))) | 1))
    }
    v4 += v4
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
