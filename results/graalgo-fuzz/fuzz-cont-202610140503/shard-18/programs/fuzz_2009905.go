package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 int; F1 string; F2 uint8 }

func main() {
    var v1 float32 = 1.5
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]int)
    var v4 int8 = 74
    var v5 int16 = -1
    v6 := []int{}
    var v7 int = 127
    v2[65535] = 2147483647
    fmt.Println(v7, v5)
    if true {
        if (v5 < v5) {
            for fi8 := 0; fi8 < 19; fi8++ {
                v9 := GenS1{F0: 895, F1: -2147483648}
                switch (-(v7)) {
                case 42:
                    switch "\n\t" {
                    case "":
                        var v10 uint16 = 0
                        _ = v10
                    case "\n\t":
                        var v11 uint = 2
                        _ = v11
                    }
                case -2147483648:
                    if true {
                        v12 := GenS1{F0: 4294967295, F1: -136}
                        var v13 uint64 = 18446744073709551615
                        _ = v12
                        _ = v13
                    }
                }
                switch "foo bar" {
                case "\n\t":
                    var v14 float64 = (1.5 + 1e10)
                    _ = v14
                case "a":
                    var v15 uint8 = 0
                    _ = v15
                }
                _ = fi8
                _ = v9
            }
        }
        v16 := GenS1{F0: 4294967295, F1: 1}
        _ = v16
    }
    v7 = (v7 | -206)
    v4 -= (v4 - v4)
    var v17 uint16 = 565
    var v18 bool = ((v7 / ((32768) | 1)) <= v7)
    v19 := GenS1{F0: 18446744073709551615, F1: -25}
    var v20 bool = (0.1 != (1e10 + float64(v1)))
    for fi21 := 0; fi21 < 14; fi21++ {
        for fi22 := 0; fi22 < 12; fi22++ {
            fmt.Println(fi21)
            if true {
                v18 = v18
                v23 := GenS1{F0: 405, F1: -2147483648}
                v20 = false
                switch (v7 + (v7 / ((v7) | 1))) {
                case -1:
                    switch ("a" + "a") {
                    case "a":
                        v4 += (v4 + v4)
                    case "":
                        v17 += (^(v17))
                        var v24 int16 = 0
                        _ = v24
                    case "foo bar":
                        delete(v2, 2)
                        fmt.Println(v20)
                        switch fi22 {
                        case 42:
                            v5 += (v5 >> ((65535) & 63))
                        }
                    }
                }
                _ = v23
            }
            v25 := GenS2{F0: fi22, F1: "a", F2: 197}
            if (328 > 2) {
                v17 -= 65535
                switch 0 {
                case 128:
                    v20 = (!(2 >= 113))
                case -316:
                    switch "foo bar" {
                    case "foo bar":
                        var v26 uint64 = 0
                        _ = v26
                    default:
                        var v27 uint32 = 690
                        _ = v27
                    }
                    if (21 == 205) {
                        var v28 string = "hello"
                        v20 = (true || v20)
                        _ = v28
                    } else {
                        var v29 int64 = 1
                        switch "foo bar" {
                        case "a":
                            v4 -= (v4 | (v4 - (v4 << ((65535) & 63))))
                        case "foo bar":
                            v30 := v3["hello"]
                            _ = v30
                        case "":
                            v17 += 1
                        default:
                            v5 += v5
                        }
                        var v31 uint64 = 776
                        _ = v29
                        _ = v31
                    }
                case 9223372036854775807:
                    switch (fi21 ^ fi22) {
                    case 32768:
                        var v32 int16 = (^(v5))
                        _ = v32
                    case -9223372036854775808:
                        switch "foo bar" {
                        case "\n\t":
                            fmt.Println(v1, v5, v1)
                        case "":
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        case "a":
                            v6 = append(v6, fi22)
                        default:
                            fmt.Println(v20, v20, v5)
                        }
                    default:
                        var v33 int = (fi22 * v7)
                        var v34 bool = v18
                        _ = v33
                        _ = v34
                    }
                default:
                    v6 = append(v6, (fi22 - -1))
                }
                v35 := GenS1{F0: 645, F1: 2147483647}
                _ = v35
            }
            _ = fi22
            _ = v25
        }
        v17 += v17
        _ = fi21
    }
    if len(v6) > 0 {
        v6[0] = (-(v7))
    }
    for fi36 := 0; fi36 < 15; fi36++ {
        fmt.Println(v7)
        _ = fi36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
