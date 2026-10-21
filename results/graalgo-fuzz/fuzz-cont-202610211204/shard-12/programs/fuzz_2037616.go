package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 int64 }
func genFn0(p0 int8, p1 uint32) int {
    p1 = p1
    var v7 int8 = 100
    var v8 uint = 255
    fmt.Println(p1, v7, p0)
    p0 -= (v7 << ((v8) & 63))
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 128
}

func main() {
    var v1 int64 = -1
    v2 := []uint32{}
    var v3 int32 = 42
    v4 := make(map[string]bool)
    var v5 int64 = v1
    v6 := []string{}
    for fi9 := 0; fi9 < 8; fi9++ {
        if false {
            var v10 float64 = 1e10
            for fi11 := 0; fi11 < 2; fi11++ {
                if (255 != 156) {
                    v5 -= 9223372036854775807
                }
                v3 = (v3 | 96)
                _ = fi11
            }
            _ = v10
        } else {
            for fi12 := 0; fi12 < 16; fi12++ {
                v1 += -9223372036854775808
                if true {
                    v5 += (v5 | v1)
                    var v13 uint16 = 873
                    var v14 int32 = (v3 + 1)
                    _ = v13
                    _ = v14
                } else {
                    v4["hello"] = true
                }
                var v15 int8 = -128
                _ = fi12
                _ = v15
            }
            if (false || true) {
                var v16 uint32 = 658
                _ = v16
            }
        }
        if ((v3 - v3) < 42) {
            for fi17 := 0; fi17 < 12; fi17++ {
                v1 = (v1 + v5)
                var v18 int8 = 42
                var v19 float64 = -2.25
                _ = fi17
                _ = v18
                _ = v19
            }
        } else {
            v3 = (^(v3))
            var v20 int = (fi9 >> ((255) & 63))
            _ = v20
        }
        v1 += (v5 >> ((2) & 63))
        _ = fi9
    }
    var v21 uint8 = 128
    for fi22 := 0; fi22 < 11; fi22++ {
        fmt.Println(v21, v21, v21)
        _ = fi22
    }
    var v23 string = "\n\t"
    if true {
        v3 = v3
        var v24 int8 = 127
        _ = v24
    }
    fmt.Println(v21)
    v5 += v5
    fmt.Println(v21, v21)
    var v25 int = 415
    v26 := GenS1{F0: (v25 << ((444) & 63)), F1: 1, F2: 42}
    for fi27 := 0; fi27 < 7; fi27++ {
        for fi28 := 0; fi28 < 19; fi28++ {
            var v29 float32 = 1.5
            _ = len(v2)
            if true {
                switch genFn0(127, 4294967295) {
                case -9223372036854775808:
                    _ = len(v4)
                case 32767:
                    switch (v23 + (v23 + "hello")) {
                    case "hello":
                        v3 = -545
                    }
                default:
                    var v30 bool = (v3 < (-(v3)))
                    _ = v30
                }
            } else {
                fmt.Println(v25, v3)
                v4[(v23 + "hello")] = true
                fmt.Println(v5, v23)
                var v31 float32 = v29
                _ = v31
            }
            _ = fi28
            _ = v29
        }
        _ = fi27
    }
    for fi32 := 0; fi32 < 14; fi32++ {
        var v33 int64 = -44
        _ = fi32
        _ = v33
    }
    switch v25 {
    case 1:
        v4[v23] = false
    case -2147483648:
        v34 := GenS1{F0: v25, F1: 18446744073709551615, F2: v5}
        _ = v34
    case 0:
        var v35 uint8 = (v21 / (((v21 & (v21 % ((v21) | 1)))) | 1))
        _ = v35
    }
    if false {
        v25 -= genFn0(0, 2147483647)
    }
    v4[("\n\t" + ("hello" + ("hello" + v23)))] = ((-(v3)) > v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v23
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
