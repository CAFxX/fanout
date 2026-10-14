package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 uint8 }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint }

func main() {
    var v1 int16 = 1
    var v2 string = "\n\t"
    var v3 int8 = 20
    v4 := make(map[string]int)
    v5 := []uint32{}
    if (v2 == ("foo bar" + ("\n\t" + "\n\t"))) {
        v2 = v2
    } else {
        v1 += (^(v1))
    }
    if len(v5) > 0 {
        v5[0] = 252
    }
    fmt.Println(v2, v3, v3)
    if (-261 < -1) {
        v6 := v4["foo bar"]
        var v7 int8 = v3
        v8 := v4[(("a" + v2) + v2)]
        _ = v6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 5; fi9++ {
        var v10 uint8 = 216
        v11 := GenS1{F0: 4294967295, F1: (v3 % ((-1) | 1)), F2: (v10 + v10)}
        if (!false) {
            v12 := GenS1{F0: 192, F1: v3, F2: v10}
            v4["\n\t"] = fi9
            _ = v12
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    if true {
        v2 = v2
        v2 = v2
    }
    v13 := GenS1{F0: 4294967295, F1: (-(v3)), F2: 0}
    if true {
        if true {
            var v14 uint = 42
            _ = v14
        }
        switch v2 {
        case "a":
            for fi15 := 0; fi15 < 3; fi15++ {
                var v16 uint8 = 0
                var v17 float32 = 3.14159
                var v18 uint64 = 42
                _ = fi15
                _ = v16
                _ = v17
                _ = v18
            }
        case "\n\t":
            var v19 uint8 = 22
            _ = v19
        case "hello":
            if false {
                for fi20 := 0; fi20 < 18; fi20++ {
                    for fi21 := 0; fi21 < 15; fi21++ {
                        v2 = ("" + v2)
                        fmt.Println(v1)
                        _ = fi21
                    }
                    v1 = (v1 * -1)
                    for fi22 := 0; fi22 < 20; fi22++ {
                        v23 := GenS2{F0: -1, F1: 128}
                        fmt.Println(v2, fi20, v1)
                        _ = fi22
                        _ = v23
                    }
                    _ = fi20
                }
                for fi24 := 0; fi24 < 5; fi24++ {
                    if true {
                        if ((fi24 | fi24) == fi24) {
                            for fi25 := 0; fi25 < 5; fi25++ {
                                v3 += int8(v1)
                                var v26 int = fi24
                                v27 := GenS2{F0: -879, F1: 255}
                                _ = fi25
                                _ = v26
                                _ = v27
                            }
                        } else {
                            fmt.Println(fi24, v2, v3)
                        }
                        v28 := GenS3{F0: ((v1 / ((v1) | 1)) >= (v1 * v1))}
                        v29 := GenS4{F0: 65535}
                        _ = v28
                        _ = v29
                    } else {
                        _ = len(v4)
                    }
                    var v30 uint32 = 835
                    v3 = (v3 - v3)
                    _ = fi24
                    _ = v30
                }
                for fi31 := 0; fi31 < 7; fi31++ {
                    fmt.Println(v3, v2)
                    _ = fi31
                }
            }
        }
    }
    var v32 uint = 255
    fmt.Println(v1, v32)
    v3 -= (v3 >> ((42) & 63))
    v2 = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
