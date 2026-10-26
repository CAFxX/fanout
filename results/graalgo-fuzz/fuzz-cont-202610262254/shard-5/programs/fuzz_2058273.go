package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 uint8 }
type GenS2 struct { F0 bool; F1 int8 }
type GenS3 struct { F0 bool; F1 uint; F2 uint32 }
type GenS4 struct { F0 uint32; F1 int16 }
func genFn0(p0 int8, p1 uint16) int16 {
    var v5 bool = false
    var v6 uint = 663
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 uint64 = 1
        for fi9 := 0; fi9 < 5; fi9++ {
            p1 += (^(p1))
            switch "a" {
            case "":
                if true {
                    var v10 uint = uint(v8)
                    v6 = (v10 & (v6 % ((0) | 1)))
                    _ = v10
                }
                if v5 {
                    if (!v5) {
                        p1 = (p1 / ((2) | 1))
                        var v11 bool = v5
                        var v12 uint64 = v8
                        var v13 int32 = -881
                        _ = v11
                        _ = v12
                        _ = v13
                    }
                    switch "hello" {
                    case "":
                        if (0 < 468) {
                            v8 = (v8 * (v8 ^ (v8 << ((v6) & 63))))
                        } else {
                            if (550 != 79) {
                                v8 -= (^(v8))
                                fmt.Println(v8)
                            } else {
                                v8 = (v8 >> ((v6) & 63))
                                var v14 uint64 = v8
                                fmt.Println(v14, v14, v6)
                                _ = v14
                            }
                        }
                    case "hello":
                        var v15 int = fi7
                        _ = v15
                    case "foo bar":
                        v8 = v8
                    }
                    switch fi7 {
                    case 0:
                        v16 := []int{fi9}
                        _ = v16
                    case 9223372036854775807:
                        p1 -= p1
                    case -198:
                        fmt.Println(fi9, p1)
                    default:
                        switch (fi7 - (fi7 >> (((v6 << ((v6) & 63))) & 63))) {
                        case 2147483647:
                            v17 := []int{}
                            _ = v17
                        case 127:
                            v18 := GenS1{F0: (^(fi7)), F1: v8, F2: 1}
                            _ = v18
                        }
                    }
                } else {
                    v8 -= v8
                    v19 := []int{}
                    v5 = (v5 && v5)
                    _ = v19
                }
            case "foo bar":
                v20 := GenS2{F0: v5, F1: (p0 - (^(p0)))}
                _ = v20
            default:
                v21 := GenS1{F0: (fi9 >> ((v6) & 63)), F1: (v8 % ((v8) | 1)), F2: 2}
                _ = v21
            }
            _ = fi9
        }
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 0
}

func main() {
    v1 := make(map[string]uint32)
    var v2 string = ""
    var v3 float32 = 1.5
    var v4 float32 = ((float32(1e10) + -2.25) * float32(1e10))
    if true {
        v2 = (v2 + ((v2 + "") + v2))
        for fi22 := 0; fi22 < 10; fi22++ {
            if (-730 == -574) {
                var v23 uint8 = 2
                delete(v1, v2)
                _ = v23
            }
            _ = fi22
        }
        switch (v2 + "") {
        case "":
            v24 := GenS2{F0: (v2 != ("hello" + "")), F1: -128}
            _ = v24
        default:
            v1["foo bar"] = 42
        }
    }
    switch (v2 + (v2 + v2)) {
    case "":
        v25 := GenS3{F0: (v2 <= "hello"), F1: 4294967295, F2: 786}
        _ = v25
    case "\n\t":
        v2 = v2
    case "foo bar":
        v26 := GenS2{F0: true, F1: 127}
        _ = v26
    }
    v1[(v2 + v2)] = 747
    var v27 int32 = -367
    fmt.Println(v4, v2)
    var v28 uint8 = 255
    v29 := GenS4{F0: 0, F1: 1}
    v30 := []int{127, 127, 9223372036854775807, -9223372036854775808}
    fmt.Println(v2)
    var v31 float64 = ((1.5 + 1.5) + -2.25)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
