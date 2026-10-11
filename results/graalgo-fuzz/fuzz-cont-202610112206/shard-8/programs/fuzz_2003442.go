package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 int32; F1 uint; F2 uint16 }
type GenS3 struct { F0 string; F1 int16; F2 bool }
type GenS4 struct { F0 int64 }
func genFn0() bool {
    var v6 bool = false
    var v7 uint32 = 514
    var v8 string = (("foo bar" + "hello") + "hello")
    for fi9 := 0; fi9 < 2; fi9++ {
        var v10 int32 = 42
        _ = fi9
        _ = v10
    }
    _ = v6
    _ = v7
    _ = v8
    return v6
}

func main() {
    var v1 uint16 = 32767
    v2 := make(map[uint32]bool)
    var v3 uint16 = (v1 / (((v1 << ((999) & 63))) | 1))
    v4 := []uint32{65535, 496, 768, 65535}
    var v5 bool = (true || true)
    v11 := v2[428]
    for fi12 := 0; fi12 < 8; fi12++ {
        for fi13 := 0; fi13 < 17; fi13++ {
            var v14 int32 = 805
            _ = fi13
            _ = v14
        }
        _ = fi12
    }
    v15 := v2[42]
    for fi16 := 0; fi16 < 19; fi16++ {
        switch "\n\t" {
        case "a":
            for fi17 := 0; fi17 < 9; fi17++ {
                switch "a" {
                case "hello":
                    if true {
                        fmt.Println(v11)
                    }
                case "foo bar":
                    v11 = (4294967295 != 200)
                    fmt.Println(v15, v5)
                case "\n\t":
                    v1 = 42
                default:
                    v1 = (v3 % ((711) | 1))
                }
                _ = fi17
            }
        case "\n\t":
            if true {
                if v5 {
                    if true {
                        if false {
                            var v18 uint64 = 260
                            var v19 int32 = -2147483648
                            _ = v18
                            _ = v19
                        } else {
                            var v20 bool = (970 == 423)
                            for fi21 := 0; fi21 < 7; fi21++ {
                                v4 = append(v4, 65535)
                                _ = fi21
                            }
                            _ = v20
                        }
                        for fi22 := 0; fi22 < 9; fi22++ {
                            if len(v4) > 0 {
                                v4[0] = 360
                            }
                            v23 := GenS1{F0: (fi16 << ((411) & 63)), F1: 42}
                            v24 := GenS2{F0: -643, F1: 42, F2: 836}
                            _ = fi22
                            _ = v23
                            _ = v24
                        }
                        v25 := GenS3{F0: "hello", F1: -1, F2: (1 < v3)}
                        _ = v25
                    } else {
                        v26 := v2[0]
                        _ = v26
                    }
                    fmt.Println(v1, v3)
                }
            } else {
                var v27 float64 = float64((0.0 * 0.0))
                var v28 int = (fi16 - fi16)
                _ = v27
                _ = v28
            }
        }
        v15 = v15
        delete(v2, 65535)
        _ = fi16
    }
    switch 128 {
    case 820:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case 1:
        var v29 int = -424
        _ = v29
    }
    switch "" {
    case "a":
        fmt.Println(v1)
    case "\n\t":
        var v30 bool = true
        _ = v30
    case "hello":
        v31 := GenS4{F0: -9223372036854775808}
        _ = v31
    default:
        delete(v2, 2147483647)
    }
    delete(v2, 2147483647)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v32 string = "hello"
    var v33 int16 = -1
    v3 = (v3 % (((v3 << ((18446744073709551615) & 63))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
