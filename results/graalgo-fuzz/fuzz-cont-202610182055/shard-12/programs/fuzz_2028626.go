package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 int8 }
func genFn0(p0 uint, p1 int16) bool {
    if false {
        if true {
            fmt.Println(p0, p1)
            if (!(!false)) {
                p1 -= (p1 + (p1 % ((32767) | 1)))
                var v6 float64 = 1e10
                var v7 uint32 = 589
                _ = v6
                _ = v7
            } else {
                p0 += (p0 * 1)
                v8 := []int{}
                for fi9 := 0; fi9 < 17; fi9++ {
                    for fi10 := 0; fi10 < 13; fi10++ {
                        if len(v8) > 0 {
                            v8[0] = -2147483648
                        }
                        if true {
                            if (70 != uint8(p0)) {
                                var v11 int = fi9
                                p0 += (^(p0))
                                v12 := GenS1{F0: 1}
                                _ = v11
                                _ = v12
                            } else {
                                v13 := GenS1{F0: 128}
                                _ = v13
                            }
                            p0 += (p0 / (((^(p0))) | 1))
                            p0 -= (^(p0))
                            p0 = (p0 | p0)
                        }
                        v14 := GenS2{F0: p0, F1: (p0 & p0)}
                        _ = fi10
                        _ = v14
                    }
                    var v15 int8 = 127
                    v8 = append(v8, fi9)
                    _ = fi9
                    _ = v15
                }
                _ = v8
            }
        }
        switch 805 {
        case 32767:
            fmt.Println(p1, p0)
        default:
            v16 := GenS2{F0: p0, F1: (p0 >> ((p0) & 63))}
            _ = v16
        }
        var v17 uint64 = 1
        _ = v17
    }
    p0 += (p0 << (((p0 | p0)) & 63))
    var v18 float64 = (1.5 * 0.0)
    v19 := GenS3{F0: -1}
    _ = p0
    _ = p1
    _ = v18
    _ = v19
    return (true || true)
}

func main() {
    v1 := []int{}
    var v2 int32 = -2147483648
    var v3 int = -9223372036854775808
    var v4 uint = 2
    var v5 bool = false
    fmt.Println(v5, v5)
    for fi20 := 0; fi20 < 15; fi20++ {
        v5 = v5
        v5 = ("hello" <= "foo bar")
        _ = fi20
    }
    var v21 uint16 = 842
    for fi22 := 0; fi22 < 14; fi22++ {
        v5 = v5
        for fi23 := 0; fi23 < 2; fi23++ {
            if len(v1) > 0 {
                v1[0] = fi22
            }
            _ = fi23
        }
        _ = fi22
    }
    for fi24 := 0; fi24 < 16; fi24++ {
        for fi25 := 0; fi25 < 7; fi25++ {
            v3 = -296
            _ = fi25
        }
        for fi26 := 0; fi26 < 15; fi26++ {
            if len(v1) > 0 {
                v1[0] = v3
            }
            _ = fi26
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi24
    }
    v21 -= 65535
    var v27 int16 = -32768
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch "foo bar" {
    case "a":
        for fi28 := 0; fi28 < 5; fi28++ {
            for fi29 := 0; fi29 < 3; fi29++ {
                var v30 uint64 = uint64(v4)
                _ = fi29
                _ = v30
            }
            var v31 uint32 = 2
            _ = fi28
            _ = v31
        }
    case "\n\t":
        var v32 string = "\n\t"
        _ = v32
    case "hello":
        v27 += v27
    }
    var v33 float64 = -2.25
    fmt.Println(v27, v21, v4)
    var v34 int64 = -9223372036854775808
    v35 := GenS2{F0: (v4 >> ((4294967295) & 63)), F1: v4}
    v21 += (v21 + (v21 * (v21 - v21)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v27
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
