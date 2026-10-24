package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 int; F1 bool; F2 bool }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint32; F1 uint }
func genFn0(p0 int64, p1 uint8) uint8 {
    var v4 uint64 = 0
    var v5 int16 = 1
    fmt.Println(v4, v4, v4)
    fmt.Println(v4)
    var v6 int16 = v5
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 0
}
func genFn1(p0 int16, p1 bool, p2 int64) int {
    var v7 int64 = p2
    switch 4294967295 {
    case 1:
        p1 = (-2147483648 >= 32767)
    }
    if p1 {
        if p1 {
            p0 = (-(p0))
        } else {
            v8 := GenS1{F0: v7, F1: 34}
            p2 += p2
            _ = v8
        }
        fmt.Println(p2, p2, p2)
        for fi9 := 0; fi9 < 1; fi9++ {
            if ((((((("\n\t" + "hello") + "foo bar") + "foo bar") + "") + "a") + ("a" + ("foo bar" + "hello"))) >= ("foo bar" + "")) {
                v10 := []int{fi9, (^(fi9))}
                if len(v10) > 0 {
                    v10[0] = (fi9 * fi9)
                }
                _ = v10
            } else {
                var v11 int8 = -26
                p0 += (p0 | (p0 + (-(p0))))
                _ = v11
            }
            p1 = (!(32767 <= 1))
            if false {
                v12 := []int{}
                _ = v12
            } else {
                p0 = p0
                for fi13 := 0; fi13 < 17; fi13++ {
                    fmt.Println(fi9)
                    _ = fi13
                }
            }
            _ = fi9
        }
    }
    var v14 uint64 = 18446744073709551615
    if ((0.1 - 0.0) >= 1e10) {
        fmt.Println(p1, p0, v14)
        for fi15 := 0; fi15 < 2; fi15++ {
            if (!(p1 && true)) {
                var v16 uint64 = v14
                v17 := GenS1{F0: (v7 | v7), F1: -1}
                _ = v16
                _ = v17
            }
            v18 := GenS1{F0: 0, F1: 127}
            _ = fi15
            _ = v18
        }
        v19 := GenS1{F0: (p2 ^ (-(p2))), F1: -80}
        _ = v19
    } else {
        var v20 uint = 1
        _ = v20
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v7
    return 643
}

func main() {
    var v1 float64 = 0.0
    var v2 int32 = 2147483647
    var v3 uint32 = 2147483648
    v2 = v2
    switch "hello" {
    case "":
        v21 := GenS2{F0: 127, F1: (!false), F2: true}
        _ = v21
    case "hello":
        v2 += 1
    }
    v1 -= v1
    v22 := GenS1{F0: 42, F1: 127}
    for fi23 := 0; fi23 < 12; fi23++ {
        if (true || (173 >= 255)) {
            for fi24 := 0; fi24 < 14; fi24++ {
                var v25 int64 = 165
                var v26 int32 = (v2 % (((^(v2))) | 1))
                v27 := []int{}
                v28 := GenS3{F0: -118}
                var v29 int64 = 42
                if true {
                    var v30 int32 = (^(v26))
                    _ = v30
                } else {
                    var v31 string = "foo bar"
                    v31 = ("foo bar" + ("\n\t" + v31))
                    var v32 int = (fi24 >> ((42) & 63))
                    _ = v31
                    _ = v32
                }
                _ = fi24
                _ = v25
                _ = v26
                _ = v27
                _ = v28
                _ = v29
            }
            v1 += (float64((1.5 + (0.1 - 0.1))) / v1)
            v33 := GenS4{F0: (v3 / ((v3) | 1)), F1: 851}
            _ = v33
        } else {
            v2 = (v2 - (v2 & (v2 & (v2 | v2))))
        }
        v34 := []int{}
        _ = fi23
        _ = v34
    }
    v1 += v1
    var v35 int64 = 42
    v3 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v35
}
