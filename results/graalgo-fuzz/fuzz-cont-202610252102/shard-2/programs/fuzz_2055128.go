package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 uint32; F1 uint }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int8 }
func genFn0(p0 uint8, p1 uint64) uint8 {
    var v7 int16 = -963
    for fi8 := 0; fi8 < 3; fi8++ {
        switch (fi8 * 4294967295) {
        case 32768:
            var v9 float32 = ((1.5 + -2.25) + ((1e10 + 3.14159) / (-2.25 / float32(3.14159))))
            _ = v9
        }
        fmt.Println(p0)
        v10 := GenS1{F0: v7, F1: (v7 % ((1) | 1))}
        _ = fi8
        _ = v10
    }
    v11 := []int{}
    _ = p0
    _ = p1
    _ = v11
    _ = v7
    return 167
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]int)
    v3 := []uint32{649, 588}
    v4 := make(map[int]bool)
    var v5 float32 = -2.25
    var v6 uint64 = 0
    fmt.Println(v6)
    v12 := GenS2{F0: 65535, F1: 18446744073709551615}
    v6 += 18446744073709551615
    v5 = v5
    var v13 uint32 = 111
    var v14 uint = 255
    var v15 uint64 = (v6 / ((v6) | 1))
    if true {
        var v16 int = 1
        _ = v16
    } else {
        v14 = (v14 / ((v14) | 1))
        var v17 float64 = -2.25
        var v18 string = "a"
        _ = v17
        _ = v18
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v15 += (v15 + (v15 + 42))
    if false {
        for fi19 := 0; fi19 < 10; fi19++ {
            switch "\n\t" {
            case "foo bar":
                for fi20 := 0; fi20 < 16; fi20++ {
                    var v21 float64 = 1e10
                    fmt.Println(fi19, v15)
                    _ = fi20
                    _ = v21
                }
            case "\n\t":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            case "a":
                v22 := GenS1{F0: 0, F1: -1}
                _ = v22
            default:
                for fi23 := 0; fi23 < 12; fi23++ {
                    var v24 int = (fi23 & (^(fi19)))
                    if len(v3) > 0 {
                        v3[0] = v13
                    }
                    _ = fi23
                    _ = v24
                }
            }
            if ((("a" + "hello") + ("\n\t" + ("" + "hello"))) <= ("foo bar" + ("" + ("hello" + "a")))) {
                fmt.Println(v15, v6, v5)
                for fi25 := 0; fi25 < 9; fi25++ {
                    v6 = v15
                    fmt.Println(v13)
                    _ = fi25
                }
            } else {
                if (v6 > v6) {
                    var v26 uint8 = genFn0(68, 42)
                    var v27 uint8 = v26
                    _ = v26
                    _ = v27
                } else {
                    delete(v2, uint32(v14))
                    _ = len(v3)
                }
                if ((!false) && false) {
                    fmt.Println(v15, v5, fi19)
                }
                for fi28 := 0; fi28 < 11; fi28++ {
                    v4[fi28] = false
                    var v29 int16 = 32767
                    fmt.Println(v13, v5, fi19)
                    _ = fi28
                    _ = v29
                }
            }
            _ = fi19
        }
        if false {
            v30 := GenS2{F0: v13, F1: v14}
            v13 -= (v13 * (v13 & (v13 % ((521) | 1))))
            _ = v30
        } else {
            var v31 bool = (uint32(v6) != v13)
            _ = v31
        }
        if (196 <= 42) {
            v6 += (v6 & (v6 | (v6 << (((v14 % (((v14 << ((230) & 63))) | 1))) & 63))))
            fmt.Println(v15)
        }
    } else {
        v2[(v13 - 2)] = 0
        var v32 float32 = ((v5 * 1e10) / -2.25)
        _ = len(v3)
        _ = v32
    }
    var v33 int64 = 811
    for fi34 := 0; fi34 < 15; fi34++ {
        fmt.Println(v15, v15, fi34)
        _ = fi34
    }
    var v35 uint64 = (^(v15))
    if true {
        if (!(!(v33 != v33))) {
            if len(v3) > 0 {
                v3[0] = (^(v13))
            }
        } else {
            v36 := GenS3{F0: 1}
            switch 127 {
            case 37:
                var v37 uint8 = 42
                _ = v37
            case 610:
                for fi38 := 0; fi38 < 10; fi38++ {
                    switch (("hello" + "foo bar") + "foo bar") {
                    case "a":
                        switch "a" {
                        case "":
                            var v39 float32 = ((v5 - v5) * (float32(-2.25) / 1e10))
                            _ = v39
                        case "foo bar":
                            if (uint64(v13) == (v35 + (v35 >> (((^(v14))) & 63)))) {
                                v40 := GenS2{F0: (v13 >> (((^(v14))) & 63)), F1: v14}
                                _ = v40
                            }
                        }
                    }
                    v35 = 42
                    _ = fi38
                }
            }
            fmt.Println(v6)
            _ = v36
        }
        for fi41 := 0; fi41 < 18; fi41++ {
            if ((127 < 127) || false) {
                v42 := GenS3{F0: 0}
                _ = v42
            }
            v14 += 0
            _ = fi41
        }
    } else {
        var v43 int32 = 0
        v33 = v33
        var v44 uint8 = 32
        _ = v43
        _ = v44
    }
    v45 := GenS4{F0: -67}
    v46 := GenS1{F0: 42, F1: 0}
    v14 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v33
    _ = v35
    _ = v4
    _ = v45
    _ = v46
    _ = v5
    _ = v6
}
