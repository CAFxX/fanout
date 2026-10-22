package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 int64; F1 int8; F2 uint32 }
type GenS3 struct { F0 int64; F1 int32 }
func genFn0(p0 uint, p1 uint32) int {
    switch 32768 {
    case 203:
        var v5 uint64 = 550
        _ = v5
    case 32768:
        var v6 string = ("foo bar" + ("a" + ""))
        _ = v6
    case 4294967295:
        if true {
            v7 := []int{2147483647, 1}
            p1 -= p1
            if true {
                if len(v7) > 0 {
                    v7[0] = 128
                }
                for fi8 := 0; fi8 < 13; fi8++ {
                    for fi9 := 0; fi9 < 6; fi9++ {
                        switch (fi8 * 9223372036854775807) {
                        case 1:
                            v10 := GenS1{F0: p0, F1: p1}
                            _ = v10
                        case 32767:
                            var v11 uint16 = 183
                            _ = v11
                        case 9223372036854775807:
                            p0 -= 696
                        }
                        _ = fi9
                    }
                    _ = fi8
                }
                p0 = (p0 * p0)
            }
            _ = v7
        }
    default:
        p0 += (p0 << ((uint(p1)) & 63))
    }
    var v12 int64 = 1
    for fi13 := 0; fi13 < 17; fi13++ {
        var v14 int16 = 0
        p1 -= p1
        v15 := []int{}
        _ = fi13
        _ = v14
        _ = v15
    }
    v16 := []int{827, -9223372036854775808, -2147483648}
    _ = p0
    _ = p1
    _ = v12
    _ = v16
    return -1
}
func genFn1(p0 int8, p1 int32) uint16 {
    p1 += -1
    var v17 string = "foo bar"
    _ = p0
    _ = p1
    _ = v17
    return 2
}

func main() {
    v1 := make(map[string]bool)
    var v2 int = 4294967295
    var v3 uint8 = 160
    v4 := []uint32{1, 1}
    if (100 <= -106) {
        v3 -= 0
        if true {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            var v18 float64 = 0.0
            fmt.Println(v3, v3)
            _ = v18
        } else {
            if (!false) {
                for fi19 := 0; fi19 < 1; fi19++ {
                    v20 := GenS1{F0: 541, F1: 65535}
                    switch (^(v2)) {
                    case 2147483647:
                        var v21 uint64 = 4294967295
                        _ = v21
                    case 9223372036854775807:
                        v1["foo bar"] = (!(!((100 < 100) && ((v3 | v3) < (v3 / ((118) | 1))))))
                    }
                    _ = fi19
                    _ = v20
                }
                if (!(true && true)) {
                    v2 += (-(v2))
                    v3 -= v3
                    var v22 uint32 = 42
                    _ = v22
                }
            } else {
                switch "" {
                case "a":
                    v2 -= v2
                }
                v3 = 0
                fmt.Println(v3)
            }
        }
        var v23 uint8 = (v3 / (((^(v3))) | 1))
        _ = v23
    } else {
        var v24 int32 = 250
        v24 += v24
        _ = v24
    }
    for fi25 := 0; fi25 < 18; fi25++ {
        v26 := GenS2{F0: -9223372036854775808, F1: 13, F2: uint32(v3)}
        fmt.Println(v2)
        _ = fi25
        _ = v26
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v27 float64 = (0.0 + 1e10)
    if false {
        for fi28 := 0; fi28 < 8; fi28++ {
            v2 = -785
            v29 := GenS1{F0: 277, F1: 0}
            var v30 string = (("a" + ("\n\t" + "hello")) + ("hello" + "a"))
            _ = fi28
            _ = v29
            _ = v30
        }
        v31 := GenS2{F0: 9223372036854775807, F1: -45, F2: 963}
        switch "\n\t" {
        case "hello":
            for fi32 := 0; fi32 < 14; fi32++ {
                fmt.Println(v2, fi32)
                _ = fi32
            }
        }
        _ = v31
    } else {
        fmt.Println(v2, v2)
        v1["hello"] = false
        v27 = v27
    }
    v3 = v3
    var v33 int64 = -1
    v34 := GenS3{F0: 1, F1: -598}
    v27 = v27
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v33
    _ = v34
    _ = v4
}
