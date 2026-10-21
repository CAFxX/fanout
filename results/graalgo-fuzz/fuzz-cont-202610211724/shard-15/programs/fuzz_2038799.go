package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 uint16 }
func genFn0(p0 uint, p1 int64) uint32 {
    v7 := []int{-2147483648, 357}
    p1 = (p1 * (p1 + -516))
    var v8 int = -2147483648
    if (42 <= 0) {
        for fi9 := 0; fi9 < 17; fi9++ {
            _ = len(v7)
            _ = len(v7)
            for fi10 := 0; fi10 < 13; fi10++ {
                fmt.Println(fi9, fi9)
                switch "\n\t" {
                case "\n\t":
                    if ((fi9 ^ (v8 * (v8 + (fi10 & -2147483648)))) != (^(fi9))) {
                        p0 = 42
                        p1 = p1
                    }
                case "a":
                    if (127 == 11) {
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                    } else {
                        v8 = (v8 ^ 32767)
                        var v11 uint = 946
                        _ = v11
                    }
                }
                v8 = 127
                _ = fi10
            }
            _ = fi9
        }
        var v12 int16 = -453
        _ = v12
    } else {
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        var v13 bool = true
        _ = v13
    }
    var v14 int64 = (p1 | (p1 * 643))
    _ = p0
    _ = p1
    _ = v14
    _ = v7
    _ = v8
    return 2
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []int{}
    var v3 bool = true
    v4 := make(map[uint32]uint32)
    v5 := []string{}
    v6 := []string{"hello", ("foo bar" + "foo bar")}
    var v15 uint64 = 0
    var v16 string = "hello"
    for fi17 := 0; fi17 < 20; fi17++ {
        fmt.Println(v3, v3)
        v15 = v15
        v4[65535] = 0
        _ = fi17
    }
    for fi18 := 0; fi18 < 17; fi18++ {
        v3 = (!v3)
        _ = fi18
    }
    switch ("a" + v16) {
    case "a":
        switch 42 {
        case 2147483647:
            v3 = v3
        case 4294967295:
            v3 = (v3 && (42 == -128))
        default:
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
    case "foo bar":
        delete(v4, 918)
    case "":
        var v19 uint64 = (v15 - 613)
        _ = v19
    default:
        if v3 {
            if (!v3) {
                v15 = v15
                v16 = (v16 + (v16 + ("foo bar" + v16)))
                switch "hello" {
                case "foo bar":
                    v3 = (!v3)
                case "hello":
                    if len(v5) > 0 {
                        v5[0] = "hello"
                    }
                case "\n\t":
                    var v20 int32 = 907
                    _ = v20
                default:
                    switch v16 {
                    case "foo bar":
                        v15 += v15
                    case "\n\t":
                        if len(v2) > 0 {
                            v2[0] = -2147483648
                        }
                    case "":
                        v21 := GenS1{F0: 128, F1: 255, F2: 245}
                        _ = v21
                    }
                }
            } else {
                for fi22 := 0; fi22 < 13; fi22++ {
                    var v23 int8 = 118
                    for fi24 := 0; fi24 < 12; fi24++ {
                        switch ("a" + (v16 + (v16 + v16))) {
                        case "a":
                            var v25 uint8 = 128
                            _ = v25
                        }
                        var v26 int64 = -9223372036854775808
                        fmt.Println(v16)
                        _ = fi24
                        _ = v26
                    }
                    _ = fi22
                    _ = v23
                }
                for fi27 := 0; fi27 < 8; fi27++ {
                    var v28 uint32 = 65535
                    _ = fi27
                    _ = v28
                }
                var v29 int = -2147483648
                _ = v29
            }
        }
    }
    var v30 bool = (v16 < v16)
    v31 := GenS1{F0: 128, F1: 143, F2: 32767}
    var v32 uint = 193
    delete(v4, 2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
