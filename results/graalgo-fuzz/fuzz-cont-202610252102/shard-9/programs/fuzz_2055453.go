package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 uint64; F1 int }
type GenS3 struct { F0 int64; F1 int32 }
func genFn0(p0 uint64) int16 {
    var v7 uint64 = 217
    v7 = p0
    var v8 int8 = 42
    _ = p0
    _ = v7
    _ = v8
    return 42
}

func main() {
    v1 := make(map[uint32]int)
    var v2 float32 = (0.1 + 1.5)
    var v3 uint = 65535
    v4 := make(map[string]int)
    var v5 int32 = 1
    v6 := make(map[string]string)
    if (42 >= 0) {
        for fi9 := 0; fi9 < 5; fi9++ {
            switch ("a" + "\n\t") {
            case "":
                for fi10 := 0; fi10 < 8; fi10++ {
                    fmt.Println(v2, v5, fi9)
                    fmt.Println(v2, fi10)
                    _ = fi10
                }
            case "foo bar":
                v2 = 0.0
            case "hello":
                v5 -= (v5 & (v5 - v5))
            }
            var v11 int32 = 1
            fmt.Println(v3)
            _ = fi9
            _ = v11
        }
    } else {
        fmt.Println(v3, v5, v2)
        var v12 int8 = -128
        _ = v12
    }
    for fi13 := 0; fi13 < 19; fi13++ {
        if false {
            switch (("a" + "foo bar") + "\n\t") {
            case "foo bar":
                v2 += v2
            case "a":
                if ((v3 >= 0) && false) {
                    var v14 string = ""
                    v14 = (v14 + "foo bar")
                    _ = v14
                }
            case "hello":
                v2 -= (v2 - (v2 * (v2 - v2)))
            }
        }
        for fi15 := 0; fi15 < 8; fi15++ {
            if true {
                switch fi13 {
                case 507:
                    v16 := []int{fi13, 9223372036854775807, fi13}
                    _ = v16
                case 416:
                    v2 += 1e10
                case 945:
                    v1[2147483647] = fi13
                default:
                    v5 += (v5 % ((v5) | 1))
                }
            }
            v17 := GenS1{F0: -9223372036854775808, F1: 42}
            _ = fi15
            _ = v17
        }
        v5 -= 925
        _ = fi13
    }
    switch 127 {
    case 9223372036854775807:
        var v18 int16 = -542
        _ = v18
    default:
        for fi19 := 0; fi19 < 14; fi19++ {
            fmt.Println(v5, v5)
            for fi20 := 0; fi20 < 18; fi20++ {
                if (fi20 != (fi20 - (fi19 & fi20))) {
                    v21 := GenS2{F0: 2, F1: -1}
                    if true {
                        v22 := []int{}
                        _ = v22
                    } else {
                        v23 := []int{fi19, fi20, 128}
                        v5 += v5
                        switch "foo bar" {
                        case "hello":
                            if false {
                                _ = len(v4)
                            }
                        case "\n\t":
                            if true {
                                var v24 uint8 = 118
                                if len(v23) > 0 {
                                    v23[0] = fi20
                                }
                                _ = v24
                            }
                        }
                        v25 := GenS2{F0: 749, F1: (-(fi20))}
                        v5 = (^(v5))
                        v2 = v2
                        v3 += 890
                        _ = v23
                        _ = v25
                    }
                    var v26 uint32 = 1
                    _ = v21
                    _ = v26
                }
                var v27 int64 = -274
                _ = fi20
                _ = v27
            }
            _ = fi19
        }
    }
    fmt.Println(v2, v5, v2)
    v3 += (^(v3))
    var v28 int64 = 42
    v29 := GenS3{F0: v28, F1: (-(v5))}
    var v30 uint = (v3 & 0)
    if ((-(v5)) != v5) {
        var v31 uint32 = 42
        v2 = 1.5
        _ = v31
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
