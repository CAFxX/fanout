package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint; F1 uint64 }
type GenS4 struct { F0 int8; F1 int64 }
func genFn0(p0 uint64, p1 int16) uint8 {
    switch "foo bar" {
    case "foo bar":
        v6 := []int{}
        _ = v6
    }
    if true {
        if (true || (9223372036854775807 != -249)) {
            p0 += p0
        } else {
            switch 2147483647 {
            case -423:
                v7 := GenS1{F0: -218}
                _ = v7
            default:
                fmt.Println(p0)
            }
            var v8 uint16 = 998
            v9 := GenS1{F0: -954}
            _ = v8
            _ = v9
        }
        fmt.Println(p0)
        fmt.Println(p0)
    } else {
        if ((false && true) && (false || (0 < 42))) {
            p0 += p0
            p1 = (-(p1))
        }
    }
    for fi10 := 0; fi10 < 1; fi10++ {
        v11 := []int{}
        switch (fi10 & 9223372036854775807) {
        case 32768:
            for fi12 := 0; fi12 < 20; fi12++ {
                var v13 int8 = -1
                switch "foo bar" {
                case "":
                    var v14 int8 = (v13 + -81)
                    _ = v14
                }
                fmt.Println(p0, p0)
                _ = fi12
                _ = v13
            }
        default:
            _ = len(v11)
        }
        _ = fi10
        _ = v11
    }
    v15 := GenS1{F0: int32(p1)}
    _ = p0
    _ = p1
    _ = v15
    return 1
}
func genFn1(p0 uint16) int64 {
    if false {
        p0 = p0
        v16 := GenS1{F0: -2147483648}
        _ = v16
    } else {
        p0 += 0
    }
    v17 := GenS2{F0: -956}
    if true {
        v18 := []int{-418, 648, 9223372036854775807}
        _ = v18
    } else {
        p0 = (^(p0))
        fmt.Println(p0)
    }
    _ = p0
    _ = v17
    return 9223372036854775807
}

func main() {
    var v1 uint = 65535
    var v2 float64 = 1.5
    var v3 uint32 = 2
    var v4 int8 = 12
    var v5 int8 = (v4 ^ v4)
    v1 = 65535
    v19 := GenS3{F0: 65535, F1: uint64(v1)}
    var v20 float64 = (v2 / 1e10)
    for fi21 := 0; fi21 < 19; fi21++ {
        fmt.Println(v1, fi21, v2)
        _ = fi21
    }
    v22 := []int{4294967295}
    var v23 int = -996
    for fi24 := 0; fi24 < 18; fi24++ {
        var v25 int64 = -9223372036854775808
        v26 := GenS2{F0: v23}
        var v27 uint = v1
        _ = fi24
        _ = v25
        _ = v26
        _ = v27
    }
    if (true || true) {
        v20 = (float64(0.0) * -2.25)
    }
    if (755 <= 0) {
        fmt.Println(v5)
        if (false && (float64(0.0) < (float64(1e10) + (v20 - float64(0.0))))) {
            v20 += v2
            var v28 int = v23
            v23 -= v28
            _ = v28
        } else {
            var v29 uint = v1
            for fi30 := 0; fi30 < 9; fi30++ {
                switch ("a" + "\n\t") {
                case "foo bar":
                    if true {
                        v31 := GenS4{F0: v4, F1: genFn1(255)}
                        if (((3.14159 + -2.25) - 0.0) == 1e10) {
                            if (-1 <= -32768) {
                                var v32 uint64 = 0
                                _ = v32
                            }
                        } else {
                            var v33 int = (-(v23))
                            var v34 uint64 = 97
                            if len(v22) > 0 {
                                v22[0] = fi30
                            }
                            _ = v33
                            _ = v34
                        }
                        _ = v31
                    }
                case "":
                    v3 = v3
                }
                _ = fi30
            }
            _ = v29
        }
        var v35 int64 = -9223372036854775808
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
