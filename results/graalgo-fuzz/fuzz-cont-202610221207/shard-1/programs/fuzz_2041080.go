package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 int64 }
type GenS2 struct { F0 int; F1 uint; F2 int16 }
type GenS3 struct { F0 int64; F1 int32; F2 bool }
func genFn0(p0 string) int64 {
    p0 = "\n\t"
    if false {
        p0 = "a"
        var v7 float64 = 0.1
        _ = v7
    } else {
        v8 := GenS1{F0: 746, F1: 2, F2: 0}
        p0 = (p0 + ("hello" + p0))
        if true {
            fmt.Println(p0)
        }
        _ = v8
    }
    _ = p0
    return -1
}
func genFn1(p0 int16, p1 uint64) uint32 {
    for fi9 := 0; fi9 < 5; fi9++ {
        var v10 float64 = 3.14159
        _ = fi9
        _ = v10
    }
    p1 = (p1 % ((p1) | 1))
    if false {
        v11 := GenS2{F0: -1, F1: 2, F2: p0}
        var v12 int16 = p0
        var v13 bool = ("" < ("\n\t" + ""))
        _ = v11
        _ = v12
        _ = v13
    } else {
        v14 := []int{127}
        var v15 uint = 51
        if len(v14) > 0 {
            v14[0] = 9223372036854775807
        }
        _ = v14
        _ = v15
    }
    p0 = (p0 & (p0 / ((p0) | 1)))
    _ = p0
    _ = p1
    return 62
}

func main() {
    var v1 int16 = 32767
    var v2 bool = false
    v3 := []int{-322, 9223372036854775807}
    var v4 string = "\n\t"
    v5 := []int{-663, -1, -1}
    v6 := []uint32{4294967295, 0, 1}
    var v16 bool = false
    var v17 uint16 = 623
    if true {
        var v18 float64 = 0.1
        v1 = (v1 << ((0) & 63))
        _ = v18
    }
    if (2 == 42) {
        v4 = (v4 + ("" + (v4 + "hello")))
        v6 = append(v6, 40)
    } else {
        switch 1 {
        case 127:
            v19 := GenS2{F0: -2147483648, F1: 727, F2: (v1 + (v1 << ((65535) & 63)))}
            _ = v19
        case -2147483648:
            var v20 string = v4
            _ = v20
        case 128:
            _ = len(v5)
        }
        if (!((v17 / ((984) | 1)) == 786)) {
            if v16 {
                v17 += (v17 << ((0) & 63))
            } else {
                for fi21 := 0; fi21 < 10; fi21++ {
                    v4 = v4
                    for fi22 := 0; fi22 < 16; fi22++ {
                        switch -288 {
                        case -364:
                            var v23 int16 = -473
                            _ = v23
                        case -238:
                            v16 = v16
                        case 9223372036854775807:
                            v2 = v2
                        }
                        if len(v6) > 0 {
                            v6[0] = 42
                        }
                        _ = fi22
                    }
                    _ = fi21
                }
                fmt.Println(v1)
                if len(v3) > 0 {
                    v3[0] = 42
                }
            }
            switch 1 {
            case 0:
                for fi24 := 0; fi24 < 6; fi24++ {
                    fmt.Println(v1, v16)
                    fmt.Println(v2)
                    _ = fi24
                }
            case 9223372036854775807:
                fmt.Println(v4, v17)
            case 32768:
                v4 = "foo bar"
            default:
                v16 = (!false)
            }
        } else {
            v25 := GenS2{F0: -1, F1: 0, F2: v1}
            for fi26 := 0; fi26 < 19; fi26++ {
                if false {
                    v1 += -32768
                }
                v27 := GenS3{F0: 42, F1: 0, F2: ((255 != 4294967295) || false)}
                _ = fi26
                _ = v27
            }
            _ = v25
        }
        fmt.Println(v1)
    }
    if v16 {
        v16 = ((^(v17)) != (v17 ^ 1))
        for fi28 := 0; fi28 < 12; fi28++ {
            fmt.Println(v1)
            v29 := GenS1{F0: 1, F1: 42, F2: -1}
            fmt.Println(v2, v1, v2)
            _ = fi28
            _ = v29
        }
    } else {
        for fi30 := 0; fi30 < 14; fi30++ {
            fmt.Println(v1, v16)
            _ = fi30
        }
    }
    v31 := GenS1{F0: 0, F1: 1, F2: 488}
    v16 = false
    if (255 == 2) {
        v32 := GenS2{F0: -1, F1: 18446744073709551615, F2: v1}
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
