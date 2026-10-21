package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 string }
type GenS2 struct { F0 uint16; F1 uint32; F2 uint64 }
type GenS3 struct { F0 uint64; F1 int16 }
type GenS4 struct { F0 string; F1 uint }
func genFn0(p0 uint32) int32 {
    p0 += 2
    var v4 int64 = 1
    fmt.Println(p0)
    for fi5 := 0; fi5 < 7; fi5++ {
        v6 := GenS1{F0: (v4 * (v4 % (((^(v4))) | 1))), F1: -190, F2: ("" + "foo bar")}
        p0 = (p0 >> ((1) & 63))
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = v4
    return 1
}
func genFn1(p0 string, p1 uint8) bool {
    var v7 uint8 = (p1 >> ((1) & 63))
    var v8 int = 4294967295
    var v9 int = v8
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return true
}

func main() {
    var v1 uint16 = 17
    var v2 uint32 = 577
    v3 := []uint32{4294967295}
    for fi10 := 0; fi10 < 16; fi10++ {
        v2 += (v2 | v2)
        if true {
            var v11 int8 = 100
            _ = v11
        } else {
            v12 := GenS2{F0: v1, F1: v2, F2: 42}
            v13 := GenS1{F0: 42, F1: -185, F2: ("\n\t" + "\n\t")}
            _ = v12
            _ = v13
        }
        if true {
            switch (fi10 >> ((0) & 63)) {
            case -1:
                var v14 uint16 = 32767
                _ = v14
            case 2147483647:
                if (!false) {
                    switch -1 {
                    case -64:
                        var v15 float64 = 3.14159
                        _ = v15
                    default:
                        v1 += v1
                    }
                    for fi16 := 0; fi16 < 14; fi16++ {
                        var v17 uint16 = (v1 % (((v1 << ((1) & 63))) | 1))
                        _ = fi16
                        _ = v17
                    }
                } else {
                    v2 -= v2
                    if false {
                        v1 = v1
                        switch (fi10 - fi10) {
                        case 2147483647:
                            switch fi10 {
                            case 128:
                                switch ("foo bar" + ("a" + "hello")) {
                                case "\n\t":
                                    v1 = (v1 * v1)
                                case "":
                                    for fi18 := 0; fi18 < 18; fi18++ {
                                        if false {
                                            var v19 uint32 = (v2 + 3)
                                            switch "\n\t" {
                                            case "hello":
                                                var v20 bool = true
                                                _ = v20
                                            case "a":
                                                fmt.Println(fi10, fi10, fi18)
                                            default:
                                                if false {
                                                    if len(v3) > 0 {
                                                        v3[0] = 65535
                                                    }
                                                    var v21 int = (-(fi10))
                                                    fmt.Println(v2, v1, v19)
                                                    if false {
                                                        if (660 == -889) {
                                                            if true {
                                                                switch "foo bar" {
                                                                case "foo bar":
                                                                    v19 -= (^(v2))
                                                                }
                                                                if len(v3) > 0 {
                                                                    v3[0] = (v19 + (v19 + (v19 << ((18446744073709551615) & 63))))
                                                                }
                                                                fmt.Println(v21, v19, v19)
                                                            } else {
                                                                v22 := GenS2{F0: v1, F1: (v2 << ((0) & 63)), F2: 1}
                                                                var v23 uint8 = 127
                                                                _ = v22
                                                                _ = v23
                                                            }
                                                            v3 = append(v3, (v2 + v2))
                                                            switch ("hello" + ("\n\t" + "a")) {
                                                            case "hello":
                                                                v1 += v1
                                                            default:
                                                                v19 = (^(v19))
                                                            }
                                                        }
                                                        var v24 float64 = 1.5
                                                        v19 = 1
                                                        _ = v24
                                                    }
                                                    _ = len(v3)
                                                    _ = v21
                                                } else {
                                                    v25 := GenS2{F0: v1, F1: v19, F2: uint64(v19)}
                                                    v26 := GenS3{F0: 929, F1: -32768}
                                                    _ = v25
                                                    _ = v26
                                                }
                                            }
                                            var v27 int16 = 0
                                            var v28 uint64 = 18446744073709551615
                                            v27 += int16(v28)
                                            _ = v19
                                            _ = v27
                                            _ = v28
                                        }
                                        if len(v3) > 0 {
                                            v3[0] = (v2 * 0)
                                        }
                                        _ = fi18
                                    }
                                case "foo bar":
                                    v1 -= v1
                                default:
                                    for fi29 := 0; fi29 < 5; fi29++ {
                                        if len(v3) > 0 {
                                            v3[0] = (^(v2))
                                        }
                                        _ = fi29
                                    }
                                }
                            case -729:
                                _ = len(v3)
                            case 2147483647:
                                v1 = (v1 << ((uint(v1)) & 63))
                            }
                        case 127:
                            if len(v3) > 0 {
                                v3[0] = v2
                            }
                        default:
                            switch ("" + "hello") {
                            case "hello":
                                if (127 == fi10) {
                                    v1 -= (v1 % ((209) | 1))
                                }
                            }
                        }
                    }
                    v1 += v1
                }
            case 127:
                var v30 int8 = int8(v1)
                _ = v30
            }
        } else {
            v31 := GenS3{F0: 119, F1: -1}
            _ = v31
        }
        _ = fi10
    }
    v32 := GenS4{F0: "a", F1: 1}
    var v33 int64 = -1
    var v34 uint32 = 65535
    switch 127 {
    case 1:
        fmt.Println(v1)
    default:
        v34 -= 65535
    }
    var v35 int32 = 1
    for fi36 := 0; fi36 < 7; fi36++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v37 uint8 = 159
        _ = fi36
        _ = v37
    }
    v35 = (v35 | 42)
    if len(v3) > 0 {
        v3[0] = 4294967295
    }
    var v38 int32 = genFn0((v2 ^ (v34 * (v34 | (v2 ^ (v2 << ((255) & 63)))))))
    var v39 uint8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v38
    _ = v39
}
