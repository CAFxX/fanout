package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 uint8 }
type GenS2 struct { F0 int8; F1 int64; F2 uint32 }
func genFn0(p0 int32, p1 uint8) uint32 {
    var v6 uint = 18446744073709551615
    var v7 bool = false
    if false {
        if v7 {
            for fi8 := 0; fi8 < 16; fi8++ {
                switch ("hello" + "") {
                case "a":
                    v9 := []int{-1, 302, (fi8 << ((v6) & 63)), fi8}
                    _ = v9
                case "":
                    if (!(((!((342 != 2) && v7)) && false) && v7)) {
                        p1 = p1
                    } else {
                        for fi10 := 0; fi10 < 1; fi10++ {
                            v11 := GenS1{F0: p0, F1: (p0 << ((4294967295) & 63)), F2: 129}
                            v12 := []int{(fi8 - fi10), (fi8 % (((fi8 - fi8)) | 1)), fi8}
                            var v13 int64 = 0
                            if len(v12) >= 2 {
                                v12 = v12[1:len(v12)]
                            }
                            _ = fi10
                            _ = v11
                            _ = v12
                            _ = v13
                        }
                        var v14 bool = ((!true) && false)
                        _ = v14
                    }
                default:
                    v15 := []int{9223372036854775807, (fi8 & -9223372036854775808), fi8}
                    _ = v15
                }
                v16 := GenS1{F0: (p0 | 2147483647), F1: (p0 | (p0 % ((42) | 1))), F2: uint8(v6)}
                _ = fi8
                _ = v16
            }
        } else {
            if true {
                v17 := GenS1{F0: 42, F1: p0, F2: (p1 ^ (p1 & p1))}
                v6 = 841
                _ = v17
            }
        }
    } else {
        for fi18 := 0; fi18 < 5; fi18++ {
            v19 := GenS1{F0: p0, F1: p0, F2: uint8(v6)}
            p0 = p0
            _ = fi18
            _ = v19
        }
        fmt.Println(p1)
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 65535
}
func genFn1(p0 uint32) string {
    p0 = 2147483648
    v20 := []int{9223372036854775807, 1, 4294967295}
    _ = p0
    _ = v20
    return ""
}

func main() {
    var v1 uint32 = 0
    v2 := make(map[int]bool)
    var v3 uint64 = 2
    var v4 int16 = 14
    v5 := make(map[string]int)
    for fi21 := 0; fi21 < 11; fi21++ {
        if (372 != 65535) {
            if (18446744073709551615 > 39) {
                for fi22 := 0; fi22 < 8; fi22++ {
                    fmt.Println(fi22)
                    if true {
                        if true {
                            v23 := GenS1{F0: 1, F1: 1, F2: 255}
                            delete(v5, "foo bar")
                            _ = v23
                        }
                    } else {
                        _ = len(v2)
                    }
                    _ = fi22
                }
            } else {
                fmt.Println(v3, fi21)
                for fi24 := 0; fi24 < 14; fi24++ {
                    v4 = (v4 ^ (-(v4)))
                    _ = fi24
                }
            }
            for fi25 := 0; fi25 < 7; fi25++ {
                v26 := GenS1{F0: 292, F1: 2147483647, F2: 2}
                _ = fi25
                _ = v26
            }
        } else {
            v27 := GenS1{F0: 176, F1: 244, F2: 42}
            _ = v27
        }
        fmt.Println(fi21)
        _ = fi21
    }
    switch 128 {
    case 42:
        for fi28 := 0; fi28 < 18; fi28++ {
            switch (fi28 % (((fi28 - (fi28 | fi28))) | 1)) {
            case 4294967295:
                v1 = v1
            }
            v3 += v3
            _ = fi28
        }
    }
    var v29 int32 = -1
    v4 = (v4 & v4)
    v1 = 0
    if false {
        fmt.Println(v1, v3)
    }
    v3 = 18446744073709551615
    var v30 float64 = (float64(0.1) * 3.14159)
    v31 := GenS2{F0: 1, F1: -9223372036854775808, F2: 65535}
    switch -2147483648 {
    case 127:
        for fi32 := 0; fi32 < 11; fi32++ {
            fmt.Println(v3, v30)
            fmt.Println(v4)
            _ = fi32
        }
    default:
        v30 -= 0.1
    }
    for fi33 := 0; fi33 < 5; fi33++ {
        v34 := []int{(fi33 - fi33), (-(fi33)), fi33}
        _ = fi33
        _ = v34
    }
    if false {
        v3 += uint64(v1)
        for fi35 := 0; fi35 < 4; fi35++ {
            v29 = -903
            _ = fi35
        }
    } else {
        v29 = 42
        var v36 int16 = (v4 | v4)
        fmt.Println(v3, v3)
        _ = v36
    }
    v37 := []int{-9223372036854775808}
    delete(v5, "")
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v37
    _ = v4
    _ = v5
}
