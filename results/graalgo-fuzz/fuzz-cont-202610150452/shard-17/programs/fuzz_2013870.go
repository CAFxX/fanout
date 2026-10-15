package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 int16; F2 uint64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint16; F1 int16; F2 uint64 }
type GenS5 struct { F0 int32; F1 uint; F2 uint8 }
func genFn0(p0 int, p1 int, p2 int8) int {
    v6 := []int{617, p1, (p1 * (p1 & (p1 % (((p0 << ((42) & 63))) | 1))))}
    for fi7 := 0; fi7 < 18; fi7++ {
        for fi8 := 0; fi8 < 14; fi8++ {
            var v9 uint8 = 128
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    var v10 string = "a"
    p0 -= 32767
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v6
    return p0
}

func main() {
    var v1 int64 = 594
    var v2 uint16 = 42
    var v3 uint16 = 1
    var v4 int = -9223372036854775808
    v5 := []string{((("hello" + ("" + "hello")) + ((("\n\t" + "\n\t") + "foo bar") + "foo bar")) + "hello"), "hello", (("" + ("hello" + "hello")) + ("foo bar" + "hello")), (("" + "\n\t") + "hello")}
    var v11 uint8 = 157
    for fi12 := 0; fi12 < 7; fi12++ {
        var v13 uint16 = (v3 % ((440) | 1))
        _ = fi12
        _ = v13
    }
    v1 += v1
    if false {
        switch (-(v4)) {
        case 32768:
            v14 := GenS1{F0: uint32(v11)}
            _ = v14
        case 1:
            v15 := GenS2{F0: "", F1: 42, F2: 0}
            _ = v15
        case 9223372036854775807:
            var v16 float64 = (0.0 * 0.0)
            _ = v16
        default:
            switch (-(v4)) {
            case 1:
                if true {
                    if false {
                        var v17 uint8 = v11
                        var v18 string = ("\n\t" + "\n\t")
                        if len(v5) > 0 {
                            v5[0] = (v18 + "a")
                        }
                        _ = v17
                        _ = v18
                    }
                } else {
                    v19 := GenS3{F0: (v11 | 128)}
                    v3 += uint16(v11)
                    fmt.Println(v4)
                    _ = v19
                }
            case -9223372036854775808:
                var v20 int64 = v1
                _ = v20
            case 4294967295:
                if (true && false) {
                    var v21 uint8 = (v11 & v11)
                    _ = v21
                } else {
                    v1 = (v1 - (v1 - (v1 - v1)))
                    for fi22 := 0; fi22 < 17; fi22++ {
                        v3 = (v2 | (v3 ^ (v2 >> ((440) & 63))))
                        v23 := GenS4{F0: 32767, F1: 32767, F2: 42}
                        _ = fi22
                        _ = v23
                    }
                    var v24 uint32 = 42
                    _ = v24
                }
            default:
                var v25 int32 = 1
                _ = v25
            }
        }
        v11 = v11
        fmt.Println(v1)
    } else {
        v2 = (^(v3))
        if true {
            switch (("a" + "foo bar") + "a") {
            case "\n\t":
                v26 := GenS5{F0: 971, F1: 328, F2: 0}
                _ = v26
            case "a":
                fmt.Println(v3, v1)
            case "foo bar":
                v4 -= v4
            }
            var v27 bool = ((1.5 / -2.25) > (3.14159 + 0.0))
            _ = v27
        }
    }
    v28 := GenS5{F0: 0, F1: 635, F2: v11}
    if true {
        var v29 float64 = (3.14159 - 0.1)
        _ = v29
    }
    var v30 uint64 = 0
    for fi31 := 0; fi31 < 15; fi31++ {
        switch (fi31 | v4) {
        case 0:
            v5 = append(v5, "a")
        case -783:
            v32 := GenS1{F0: 2}
            _ = v32
        default:
            v30 = v30
        }
        var v33 uint16 = v2
        _ = fi31
        _ = v33
    }
    v11 += (v11 & v11)
    _ = len(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v28
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
