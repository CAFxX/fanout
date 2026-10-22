package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int16 }
func genFn0() int16 {
    switch "foo bar" {
    case "foo bar":
        var v7 int8 = 0
        _ = v7
    }
    var v8 int16 = 1
    var v9 int8 = 100
    _ = v8
    _ = v9
    return (^(v8))
}
func genFn1(p0 string, p1 int32, p2 uint32) int64 {
    p1 = (p1 + (p1 / ((p1) | 1)))
    v10 := GenS1{F0: 2, F1: 4294967295, F2: (false && (!(!false)))}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 1
}

func main() {
    var v1 uint32 = 2
    var v2 int32 = 818
    var v3 uint64 = 4294967295
    v4 := make(map[uint32]bool)
    var v5 uint32 = 717
    v6 := []string{}
    for fi11 := 0; fi11 < 20; fi11++ {
        var v12 int16 = 32767
        if (uint16(v1) <= 894) {
            var v13 int32 = (v2 / (((-(v2))) | 1))
            _ = v13
        } else {
            v14 := GenS1{F0: 847, F1: fi11, F2: false}
            for fi15 := 0; fi15 < 19; fi15++ {
                v12 += genFn0()
                _ = fi15
            }
            v16 := GenS2{F0: 2}
            _ = v14
            _ = v16
        }
        v17 := v4[v5]
        _ = fi11
        _ = v12
        _ = v17
    }
    var v18 int = -629
    fmt.Println(v18, v3, v2)
    if false {
        v3 -= 18446744073709551615
        var v19 uint64 = (v3 ^ (v3 - (v3 | v3)))
        _ = v19
    }
    v20 := GenS1{F0: 65535, F1: v18, F2: false}
    fmt.Println(v18)
    if (false || (127 >= -128)) {
        for fi21 := 0; fi21 < 9; fi21++ {
            for fi22 := 0; fi22 < 17; fi22++ {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                v3 = (^(v3))
                _ = fi22
            }
            v3 += (v3 / ((v3) | 1))
            _ = fi21
        }
        v23 := GenS2{F0: 1}
        for fi24 := 0; fi24 < 13; fi24++ {
            v5 = v1
            v3 += v3
            _ = fi24
        }
        _ = v23
    } else {
        var v25 uint32 = (v5 >> ((4294967295) & 63))
        switch "foo bar" {
        case "a":
            var v26 uint8 = 128
            _ = v26
        case "":
            var v27 int16 = 1
            _ = v27
        case "\n\t":
            if (255 < 888) {
                switch "\n\t" {
                case "hello":
                    for fi28 := 0; fi28 < 17; fi28++ {
                        v29 := GenS1{F0: 65535, F1: (fi28 / (((v18 ^ (v18 << ((2) & 63)))) | 1)), F2: (!false)}
                        var v30 int = (fi28 & (fi28 << ((878) & 63)))
                        _ = fi28
                        _ = v29
                        _ = v30
                    }
                case "":
                    if (258 == uint(v5)) {
                        if false {
                            var v31 int16 = -388
                            _ = v31
                        }
                    } else {
                        v2 -= v2
                    }
                case "\n\t":
                    v5 = 42
                default:
                    fmt.Println(v25, v1)
                }
                var v32 uint16 = 718
                _ = v32
            }
        default:
            for fi33 := 0; fi33 < 9; fi33++ {
                _ = len(v4)
                v34 := v4[v25]
                _ = fi33
                _ = v34
            }
        }
        _ = v25
    }
    v35 := GenS2{F0: 0}
    v2 = (v2 % ((v2) | 1))
    v2 = v2
    v2 = 1
    v5 = v1
    fmt.Println(v3, v3, v2)
    if ((42 == 65535) || (0 == (v18 + v18))) {
        var v36 uint64 = (v3 >> ((0) & 63))
        var v37 int16 = 32767
        v5 = (v5 << ((744) & 63))
        _ = v36
        _ = v37
    } else {
        var v38 int = v18
        delete(v4, 799)
        _ = v38
    }
    v39 := GenS3{F0: 42}
    for fi40 := 0; fi40 < 17; fi40++ {
        v1 = (^(v1))
        v18 += -1
        if (false || true) {
            v41 := GenS1{F0: 42, F1: -120, F2: false}
            var v42 int32 = (v2 / ((1) | 1))
            _ = v41
            _ = v42
        } else {
            if (!true) {
                fmt.Println(v18)
            }
        }
        _ = fi40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v35
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
