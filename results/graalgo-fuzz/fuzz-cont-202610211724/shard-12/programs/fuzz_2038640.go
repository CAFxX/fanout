package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 int8 }
type GenS2 struct { F0 uint64; F1 bool; F2 int }
func genFn0(p0 int32, p1 int64, p2 uint16) uint16 {
    if (((uint32(p2) < 417) && (!(0 < 1))) || true) {
        switch 32768 {
        case -1:
            v6 := []int{}
            _ = v6
        case 4294967295:
            v7 := []int{}
            _ = v7
        }
        if (42 == 0) {
            fmt.Println(p1, p2)
            p0 -= (p0 >> ((4294967295) & 63))
            for fi8 := 0; fi8 < 2; fi8++ {
                fmt.Println(p0)
                fmt.Println(fi8, p1)
                _ = fi8
            }
        } else {
            if (!true) {
                var v9 int32 = 2147483647
                var v10 int64 = p1
                _ = v10
                _ = v9
            } else {
                v11 := []int{0, 2147483647, 32767}
                for fi12 := 0; fi12 < 16; fi12++ {
                    fmt.Println(p1, p1, p1)
                    p1 -= (p1 * p1)
                    fmt.Println(p2)
                    _ = fi12
                }
                if ((-103 <= 127) || (((!true) && false) || (false || false))) {
                    if len(v11) > 0 {
                        v11[0] = 4294967295
                    }
                }
                _ = v11
            }
            for fi13 := 0; fi13 < 19; fi13++ {
                v14 := GenS1{F0: (!(((-(p0)) >= (p0 >> ((799) & 63))) || (!(true && true)))), F1: (fi13 + (fi13 * fi13)), F2: -128}
                _ = fi13
                _ = v14
            }
            for fi15 := 0; fi15 < 6; fi15++ {
                p2 = (p2 ^ 2)
                v16 := GenS1{F0: (p0 > p0), F1: 42, F2: 1}
                _ = fi15
                _ = v16
            }
        }
    } else {
        fmt.Println(p2, p0)
        p0 = p0
    }
    fmt.Println(p1, p0)
    if (0 <= 255) {
        v17 := []int{}
        if len(v17) >= 2 {
            v17 = v17[1:len(v17)]
        }
        _ = v17
    }
    fmt.Println(p1, p1, p2)
    v18 := []int{971, 4294967295, 444, -923}
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    return (p2 >> ((309) & 63))
}

func main() {
    var v1 uint16 = 607
    var v2 uint16 = v1
    v3 := []int{-2147483648, 9223372036854775807, 32768, -378}
    var v4 uint64 = 42
    v5 := make(map[string]uint32)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v1 += v2
    for fi19 := 0; fi19 < 6; fi19++ {
        fmt.Println(v4, v4)
        _ = fi19
    }
    for fi20 := 0; fi20 < 12; fi20++ {
        fmt.Println(v4, fi20, v2)
        for fi21 := 0; fi21 < 12; fi21++ {
            switch fi21 {
            case 127:
                v2 += v1
            case -2147483648:
                var v22 uint16 = v1
                _ = v22
            case -9223372036854775808:
                switch (("foo bar" + "\n\t") + (("\n\t" + ("foo bar" + ("\n\t" + ("foo bar" + "")))) + "foo bar")) {
                case "":
                    fmt.Println(fi20)
                default:
                    v5["foo bar"] = 2147483647
                }
                if (!false) {
                    _ = len(v5)
                    var v23 bool = true
                    _ = v23
                } else {
                    if true {
                        v4 += (^(v4))
                        var v24 int = fi20
                        _ = v24
                    } else {
                        if true {
                            v25 := GenS1{F0: true, F1: (fi20 | fi21), F2: -1}
                            _ = v25
                        } else {
                            var v26 bool = false
                            if (42 < int8(v2)) {
                                var v27 int = 2147483647
                                v28 := GenS1{F0: (76 > 100), F1: (fi20 & fi20), F2: 127}
                                _ = v27
                                _ = v28
                            }
                            _ = v26
                        }
                        v4 = v4
                        v1 += (v1 ^ (^(v2)))
                    }
                    var v29 int64 = 42
                    var v30 int32 = 649
                    _ = v29
                    _ = v30
                }
            default:
                v1 = (^(v2))
            }
            if (((!true) || (!false)) || false) {
                var v31 uint8 = 42
                _ = v31
            } else {
                v32 := GenS2{F0: uint64(v1), F1: (!false), F2: fi21}
                _ = v32
            }
            _ = fi21
        }
        _ = fi20
    }
    for fi33 := 0; fi33 < 3; fi33++ {
        var v34 int32 = 42
        switch (("" + "foo bar") + ("\n\t" + "\n\t")) {
        case "hello":
            v35 := GenS2{F0: v4, F1: true, F2: -188}
            _ = v35
        default:
            var v36 int = (fi33 % (((^(fi33))) | 1))
            _ = v36
        }
        _ = fi33
        _ = v34
    }
    var v37 int32 = 0
    v1 = v1
    fmt.Println(v4, v37)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
}
