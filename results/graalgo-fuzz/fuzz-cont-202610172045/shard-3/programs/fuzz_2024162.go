package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 uint }
type GenS2 struct { F0 uint8; F1 uint16 }
func genFn0(p0 uint8, p1 int64, p2 int) int32 {
    p2 = (p2 << ((157) & 63))
    var v6 int = p2
    switch "\n\t" {
    case "":
        p2 += 128
    case "foo bar":
        var v7 float32 = 1.5
        _ = v7
    }
    switch (p2 >> ((18446744073709551615) & 63)) {
    case 0:
        v8 := GenS1{F0: 4294967295, F1: 665, F2: 42}
        _ = v8
    case -2147483648:
        for fi9 := 0; fi9 < 4; fi9++ {
            for fi10 := 0; fi10 < 2; fi10++ {
                v11 := []int{}
                switch -335 {
                case 0:
                    if true {
                        _ = len(v11)
                        if len(v11) > 0 {
                            v11[0] = int(p0)
                        }
                    } else {
                        fmt.Println(v6)
                        var v12 uint8 = p0
                        _ = v12
                    }
                case 9223372036854775807:
                    var v13 float64 = 1e10
                    _ = v13
                default:
                    if len(v11) >= 2 {
                        v11 = v11[1:len(v11)]
                    }
                }
                v14 := GenS1{F0: 1, F1: 2147483648, F2: 1}
                _ = fi10
                _ = v11
                _ = v14
            }
            v15 := []int{fi9, (p2 + (p2 << ((255) & 63))), (-(v6)), -2147483648}
            p0 -= (p0 + (p0 * 98))
            _ = fi9
            _ = v15
        }
    case 32767:
        var v16 int32 = -2147483648
        _ = v16
    default:
        var v17 float64 = 1e10
        _ = v17
    }
    v18 := []int{(v6 >> ((754) & 63)), -2147483648, -9223372036854775808, p2}
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v6
    return 0
}

func main() {
    var v1 string = (("a" + ("" + ("hello" + "foo bar"))) + "foo bar")
    var v2 uint = 82
    var v3 float32 = (0.0 - (3.14159 - ((-2.25 + float32(1.5)) + (0.1 * (float32((0.1 - (0.0 * (3.14159 + (1e10 + -2.25))))) + 1.5)))))
    v4 := []int{}
    var v5 bool = false
    v19 := GenS2{F0: 255, F1: 0}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if (42 > 255) {
        for fi20 := 0; fi20 < 19; fi20++ {
            if v5 {
                fmt.Println(v1)
            }
            for fi21 := 0; fi21 < 9; fi21++ {
                var v22 uint16 = 0
                var v23 uint16 = v22
                var v24 float64 = float64(v3)
                _ = fi21
                _ = v22
                _ = v23
                _ = v24
            }
            _ = fi20
        }
    }
    _ = len(v4)
    v1 = ""
    var v25 uint8 = 2
    v2 += (v2 ^ v2)
    v5 = (!(!(uint32(v25) < 485)))
    for fi26 := 0; fi26 < 1; fi26++ {
        v27 := GenS1{F0: 195, F1: 338, F2: v2}
        _ = fi26
        _ = v27
    }
    for fi28 := 0; fi28 < 19; fi28++ {
        for fi29 := 0; fi29 < 9; fi29++ {
            fmt.Println(v1)
            fmt.Println(v1, v25, v25)
            _ = fi29
        }
        var v30 float32 = (v3 * v3)
        _ = fi28
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
