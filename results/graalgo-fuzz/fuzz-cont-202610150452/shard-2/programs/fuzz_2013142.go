package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8; F1 uint64 }
type GenS4 struct { F0 int64; F1 int32 }
type GenS5 struct { F0 uint64; F1 int32; F2 uint16 }
type GenS6 struct { F0 string; F1 int16 }
func genFn0(p0 uint64, p1 string, p2 uint) int64 {
    v5 := GenS1{F0: 65535, F1: -1}
    if true {
        p1 = "a"
        for fi6 := 0; fi6 < 14; fi6++ {
            if true {
                v7 := GenS2{F0: 42}
                var v8 int64 = -582
                _ = v7
                _ = v8
            }
            _ = fi6
        }
    }
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 366
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{"", "foo bar", "foo bar", ("foo bar" + ("hello" + "hello"))}
    v3 := make(map[string]bool)
    var v4 uint32 = 4294967295
    fmt.Println(v4)
    fmt.Println(v4)
    var v9 uint8 = 127
    v10 := GenS3{F0: 127, F1: 224}
    fmt.Println(v9, v4)
    v9 -= (v9 * v9)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v2 = append(v2, "a")
    switch "" {
    case "\n\t":
        if ((274 == 255) && false) {
            if ((false || (127 != 1)) || true) {
                v11 := GenS3{F0: 96, F1: 1}
                for fi12 := 0; fi12 < 10; fi12++ {
                    v4 = 2147483647
                    _ = fi12
                }
                _ = v11
            } else {
                v4 = (^(v4))
            }
        } else {
            fmt.Println(v9, v4)
            var v13 int32 = 42
            _ = v13
        }
    case "foo bar":
        v14 := GenS4{F0: 9223372036854775807, F1: 1}
        _ = v14
    }
    var v15 int64 = 42
    switch 128 {
    case 1:
        if (3.14159 > 0.1) {
            fmt.Println(v4, v9)
            v15 += 1
        } else {
            v3["\n\t"] = false
            if ((false && ((0 <= 286) && true)) && false) {
                var v16 uint32 = v4
                for fi17 := 0; fi17 < 14; fi17++ {
                    var v18 int16 = 0
                    if (!true) {
                        v3[""] = (false && ((v18 ^ (v18 >> ((4294967295) & 63))) > (^(v18))))
                    } else {
                        switch 32767 {
                        case 596:
                            for fi19 := 0; fi19 < 16; fi19++ {
                                if (fi19 == fi19) {
                                    if (18446744073709551615 > uint(v9)) {
                                        if true {
                                            switch (-(fi17)) {
                                            case -1:
                                                var v20 uint32 = (v4 + v16)
                                                _ = v20
                                            case 127:
                                                v21 := GenS5{F0: 228, F1: 782, F2: 32767}
                                                _ = v21
                                            case -391:
                                                v9 += v9
                                            default:
                                                _ = len(v2)
                                                switch fi17 {
                                                case 2147483647:
                                                    v16 += (v4 - 1)
                                                case -9223372036854775808:
                                                    v4 = (v4 % (((v16 + (v16 - (v16 >> ((4294967295) & 63))))) | 1))
                                                case 32768:
                                                    var v22 string = ""
                                                    _ = v22
                                                }
                                            }
                                            v3[("hello" + ("hello" + "hello"))] = false
                                        }
                                        var v23 float32 = -2.25
                                        _ = v23
                                    }
                                    _ = len(v1)
                                    fmt.Println(v9, fi17)
                                }
                                v24 := GenS5{F0: 2, F1: -1, F2: 684}
                                v25 := GenS2{F0: 0}
                                _ = fi19
                                _ = v24
                                _ = v25
                            }
                        case 902:
                            for fi26 := 0; fi26 < 8; fi26++ {
                                v9 = (v9 ^ v9)
                                if true {
                                    v9 += 56
                                    var v27 float32 = ((float32(-2.25) + float32(((1e10 + 3.14159) - -2.25))) / -2.25)
                                    v28 := GenS5{F0: 1, F1: 419, F2: 29}
                                    _ = v27
                                    _ = v28
                                }
                                _ = fi26
                            }
                        case 757:
                            var v29 int32 = int32(v9)
                            _ = v29
                        }
                    }
                    _ = fi17
                    _ = v18
                }
                _ = len(v2)
                _ = v16
            }
        }
    case 128:
        v4 = 65535
    }
    v15 = 42
    var v30 int64 = genFn0(18446744073709551615, "a", 982)
    v9 = v9
    v31 := GenS1{F0: 2, F1: (v15 + v30)}
    for fi32 := 0; fi32 < 16; fi32++ {
        if true {
            for fi33 := 0; fi33 < 6; fi33++ {
                var v34 int = fi33
                var v35 string = "a"
                _ = fi33
                _ = v34
                _ = v35
            }
            v36 := GenS1{F0: 0, F1: genFn0(2, "a", 1)}
            v30 -= v30
            _ = v36
        }
        v37 := GenS3{F0: -19, F1: 18446744073709551615}
        if ((^(fi32)) != fi32) {
            v9 = (v9 & v9)
            var v38 uint8 = v9
            _ = v38
        } else {
            var v39 uint32 = 2
            for fi40 := 0; fi40 < 19; fi40++ {
                switch "a" {
                case "\n\t":
                    delete(v1, v4)
                default:
                    var v41 int64 = (v30 & (v30 >> ((18446744073709551615) & 63)))
                    _ = v41
                }
                if (189 <= -2147483648) {
                    switch (("foo bar" + "hello") + ((("" + ("" + "")) + "") + "")) {
                    case "\n\t":
                        v42 := GenS2{F0: 10}
                        _ = v42
                    case "hello":
                        v4 = v4
                    default:
                        if false {
                            v9 -= (v9 ^ 0)
                            v43 := GenS2{F0: 296}
                            _ = v43
                        } else {
                            var v44 float32 = 1.5
                            if false {
                                var v45 int = (fi40 - 2147483647)
                                _ = v45
                            }
                            v44 = 0.1
                            _ = v44
                        }
                    }
                    var v46 float64 = (0.1 / float64(1.5))
                    _ = v46
                } else {
                    v47 := GenS6{F0: ("foo bar" + "hello"), F1: -337}
                    delete(v1, uint32(v9))
                    _ = v47
                }
                _ = fi40
            }
            v39 = (v4 >> ((18446744073709551615) & 63))
            _ = v39
        }
        _ = fi32
        _ = v37
    }
    _ = len(v3)
    v48 := GenS2{F0: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v48
    _ = v9
}
