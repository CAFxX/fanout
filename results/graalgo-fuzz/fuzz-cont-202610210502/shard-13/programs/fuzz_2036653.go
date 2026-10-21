package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 uint64; F1 uint; F2 int32 }
type GenS3 struct { F0 uint; F1 int16 }
type GenS4 struct { F0 uint32; F1 uint64; F2 uint16 }
type GenS5 struct { F0 uint64; F1 int64; F2 uint8 }
type GenS6 struct { F0 int8 }
func genFn0(p0 int32, p1 uint32) uint16 {
    fmt.Println(p1, p1)
    switch -9223372036854775808 {
    case 1:
        v6 := GenS1{F0: 104, F1: 65535}
        _ = v6
    case 42:
        p1 += (^(p1))
    case 32768:
        p1 += (p1 << ((4294967295) & 63))
    }
    _ = p0
    _ = p1
    return 42
}
func genFn1(p0 int64) int {
    var v7 bool = ((0 > 1) && true)
    var v8 int = -1
    var v9 uint64 = 0
    v10 := GenS2{F0: (v9 % ((1) | 1)), F1: 422, F2: 42}
    _ = p0
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return v8
}

func main() {
    var v1 int32 = 1
    v2 := make(map[string]int)
    var v3 int64 = -498
    var v4 int8 = 42
    var v5 uint64 = 18446744073709551615
    switch genFn1(v3) {
    case 2147483647:
        v4 += (v4 % ((v4) | 1))
    default:
        v11 := []int{-593, 474, -1}
        _ = v11
    }
    var v12 uint = 18446744073709551615
    v13 := []int{1, genFn1((v3 & v3)), 1, genFn1((v3 / ((int64(v1)) | 1)))}
    fmt.Println(v12, v12, v4)
    if false {
        var v14 int16 = -32768
        _ = v14
    } else {
        v5 = v5
        var v15 uint = v12
        _ = v15
    }
    if len(v13) > 0 {
        v13[0] = -540
    }
    switch "" {
    case "hello":
        for fi16 := 0; fi16 < 9; fi16++ {
            if ((32767 <= -920) && (false || (!(!(55 <= v4))))) {
                v1 = 0
                var v17 int = (fi16 ^ (^(fi16)))
                fmt.Println(fi16, v3)
                _ = v17
            }
            if true {
                switch "a" {
                case "\n\t":
                    v12 += (v12 >> (((v12 - (v12 | 1))) & 63))
                }
                if len(v13) > 0 {
                    v13[0] = (fi16 / (((fi16 >> (((v12 * 65535)) & 63))) | 1))
                }
            }
            v4 = v4
            _ = fi16
        }
    case "\n\t":
        fmt.Println(v12)
    case "a":
        switch "hello" {
        case "a":
            fmt.Println(v1, v5)
        default:
            v4 = v4
        }
    }
    v13 = append(v13, int(v1))
    v2["foo bar"] = -9223372036854775808
    var v18 int = 32768
    switch ("" + (((("a" + "foo bar") + ("a" + "\n\t")) + ("hello" + ("" + ""))) + ("\n\t" + "foo bar"))) {
    case "foo bar":
        for fi19 := 0; fi19 < 6; fi19++ {
            if true {
                switch v18 {
                case -2147483648:
                    if true {
                        v3 = (^(v3))
                    } else {
                        v20 := GenS2{F0: 4294967295, F1: v12, F2: (^(v1))}
                        switch -523 {
                        case 262:
                            v5 = (v5 >> ((uint(v5)) & 63))
                        case -243:
                            v21 := v2["a"]
                            _ = v21
                        case 215:
                            switch v18 {
                            case -380:
                                v1 = v1
                            case -9223372036854775808:
                                var v22 uint = 656
                                _ = v22
                            default:
                                _ = len(v13)
                            }
                        default:
                            v23 := GenS3{F0: (v12 & 255), F1: 1}
                            _ = v23
                        }
                        _ = v20
                    }
                default:
                    switch (-(fi19)) {
                    case -58:
                        v1 -= (^(v1))
                    case 42:
                        if (!false) {
                            v2["foo bar"] = (v18 ^ v18)
                            v5 = 4294967295
                            switch fi19 {
                            case 2147483647:
                                switch "foo bar" {
                                case "a":
                                    var v24 uint8 = 255
                                    _ = v24
                                case "":
                                    for fi25 := 0; fi25 < 7; fi25++ {
                                        if (int(v1) >= (^(fi25))) {
                                            var v26 int16 = int16(v1)
                                            v12 = (v12 >> ((0) & 63))
                                            _ = v26
                                        } else {
                                            v27 := GenS4{F0: 65535, F1: (^(v5)), F2: 880}
                                            if false {
                                                v28 := GenS4{F0: 395, F1: (v5 & (^(v5))), F2: 32767}
                                                var v29 uint32 = 974
                                                fmt.Println(v1, fi19, v5)
                                                _ = v28
                                                _ = v29
                                            } else {
                                                fmt.Println(fi19, fi25, fi25)
                                            }
                                            _ = v27
                                        }
                                        v4 -= (v4 * v4)
                                        v3 += -800
                                        _ = fi25
                                    }
                                case "foo bar":
                                    var v30 int32 = (v1 - v1)
                                    _ = v30
                                }
                            case -473:
                                v31 := GenS5{F0: 18446744073709551615, F1: v3, F2: 59}
                                _ = v31
                            default:
                                fmt.Println(v5)
                            }
                        }
                    }
                }
                v1 -= (^(v1))
            } else {
                v32 := GenS6{F0: -128}
                _ = v32
            }
            v33 := v2["foo bar"]
            _ = fi19
            _ = v33
        }
    case "\n\t":
        var v34 uint8 = 127
        _ = v34
    }
    v5 = v5
    var v35 uint32 = 42
    if (42 == 127) {
        v2["hello"] = (v18 / (((v18 << (((v12 >> ((v12) & 63))) & 63))) | 1))
        v36 := GenS2{F0: (v5 << ((v12) & 63)), F1: (v12 % ((0) | 1)), F2: v1}
        if len(v13) >= 2 {
            v13 = v13[1:len(v13)]
        }
        _ = v36
    } else {
        var v37 uint16 = genFn0(1, 2147483647)
        var v38 int8 = (v4 << ((uint(v37)) & 63))
        _ = v37
        _ = v38
    }
    var v39 uint32 = (v35 + (^(v35)))
    fmt.Println(v39, v3, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v35
    _ = v39
    _ = v4
    _ = v5
}
