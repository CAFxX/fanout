package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 uint16; F1 int8 }
func genFn0(p0 uint32, p1 uint8, p2 uint16) uint {
    v7 := GenS1{F0: (p2 << ((18446744073709551615) & 63)), F1: 1}
    var v8 int = 32768
    if false {
        v9 := []int{(v8 << ((916) & 63))}
        fmt.Println(v8, p2, p1)
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return uint(p1)
}

func main() {
    var v1 int64 = 42
    var v2 int16 = 1000
    v3 := []uint32{42, 65535, 495, 0}
    var v4 bool = true
    var v5 int = -618
    var v6 uint = 2
    var v10 uint8 = 103
    for fi11 := 0; fi11 < 8; fi11++ {
        var v12 uint = 255
        _ = fi11
        _ = v12
    }
    if v4 {
        v3 = append(v3, 0)
    }
    if false {
        var v13 int = (v5 ^ v5)
        if v4 {
            switch 4294967295 {
            case 32767:
                v10 -= v10
            case 128:
                if true {
                    if len(v3) > 0 {
                        v3[0] = uint32(v10)
                    }
                    if (42 != 0) {
                        switch "a" {
                        case "\n\t":
                            var v14 bool = false
                            _ = v14
                        default:
                            fmt.Println(v6, v4, v6)
                        }
                        switch "a" {
                        case "":
                            if len(v3) > 0 {
                                v3[0] = 2
                            }
                        default:
                            v5 -= (v13 & (v5 ^ 32768))
                        }
                        var v15 int = v13
                        _ = v15
                    }
                } else {
                    var v16 string = "a"
                    var v17 int16 = v2
                    var v18 uint32 = 4294967295
                    _ = v16
                    _ = v17
                    _ = v18
                }
            }
            v13 = (v13 + v5)
        }
        _ = v13
    } else {
        v1 = (-(v1))
    }
    fmt.Println(v6, v1, v4)
    v19 := GenS1{F0: 684, F1: 87}
    v6 += 65535
    fmt.Println(v2, v5)
    for fi20 := 0; fi20 < 8; fi20++ {
        fmt.Println(fi20, v5)
        if ("\n\t" >= "foo bar") {
            var v21 string = ((((((((((("a" + "\n\t") + (("a" + "") + "")) + (("\n\t" + "\n\t") + "hello")) + "a") + ("\n\t" + "a")) + "a") + (("foo bar" + "") + "foo bar")) + "a") + "a") + ("" + "\n\t")) + "")
            fmt.Println(fi20, fi20)
            _ = v21
        } else {
            _ = len(v3)
            for fi22 := 0; fi22 < 7; fi22++ {
                v23 := GenS1{F0: 971, F1: 115}
                v5 += (v5 & (fi20 / (((-(fi22))) | 1)))
                v5 += int(v1)
                _ = fi22
                _ = v23
            }
        }
        _ = fi20
    }
    v24 := GenS2{F0: (^(v6)), F1: (42 < v2)}
    _ = len(v3)
    var v25 float32 = float32(0.0)
    var v26 string = ""
    for fi27 := 0; fi27 < 10; fi27++ {
        v6 -= (v6 * (v6 << ((v6) & 63)))
        _ = fi27
    }
    v28 := GenS3{F0: 737, F1: 127}
    if (!(!v4)) {
        var v29 bool = (4294967295 <= 18446744073709551615)
        _ = v29
    } else {
        switch v5 {
        case 367:
            switch v26 {
            case "\n\t":
                v1 -= v1
            case "a":
                v30 := GenS2{F0: (v6 | genFn0(42, v10, 32767)), F1: v4}
                _ = v30
            case "hello":
                fmt.Println(v2, v25)
            }
        case -9223372036854775808:
            if true {
                if len(v3) > 0 {
                    v3[0] = 995
                }
                if len(v3) > 0 {
                    v3[0] = 289
                }
            }
        }
    }
    v31 := GenS2{F0: v6, F1: v4}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v28
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
