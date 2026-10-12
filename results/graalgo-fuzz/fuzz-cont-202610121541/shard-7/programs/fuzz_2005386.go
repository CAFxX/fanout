package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 uint64 }
type GenS2 struct { F0 uint64; F1 bool }
func genFn0(p0 int, p1 int64) uint16 {
    p0 -= p0
    if true {
        var v5 uint64 = 195
        _ = v5
    }
    v6 := []int{-1, p0}
    var v7 uint16 = 0
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return v7
}

func main() {
    var v1 uint64 = 42
    var v2 uint = 42
    var v3 string = "foo bar"
    var v4 int8 = 85
    for fi8 := 0; fi8 < 1; fi8++ {
        v3 = ""
        if (((0.1 - (-2.25 / 1.5)) - 3.14159) > (-2.25 + (1.5 / 1.0))) {
            v3 = "foo bar"
            switch "a" {
            case "":
                if ((((v4 < (v4 % ((-1) | 1))) && ((-2.25 / ((0.0 + (((0.0 + 0.0) * ((float32(1.5) / 1.5) + 0.1)) + 1e10)) * float32(float64(-2.25)))) <= 3.14159)) && true) && true) {
                    if false {
                        fmt.Println(v2, fi8)
                    } else {
                        var v9 uint16 = 293
                        if ((v1 & (^(v1))) == v1) {
                            if true {
                                switch "a" {
                                case "a":
                                    v3 = ("hello" + ("\n\t" + v3))
                                default:
                                    v10 := []int{(fi8 & (fi8 & int(v9)))}
                                    _ = v10
                                }
                                for fi11 := 0; fi11 < 9; fi11++ {
                                    if false {
                                        v12 := GenS1{F0: fi8, F1: 183, F2: (v1 * 18446744073709551615)}
                                        _ = v12
                                    }
                                    v13 := GenS1{F0: (-(fi8)), F1: 2147483648, F2: 979}
                                    _ = fi11
                                    _ = v13
                                }
                            }
                            for fi14 := 0; fi14 < 7; fi14++ {
                                var v15 uint64 = v1
                                v16 := GenS1{F0: (-(fi8)), F1: 0, F2: 614}
                                v17 := GenS1{F0: (fi14 % ((0) | 1)), F1: 905, F2: 42}
                                _ = fi14
                                _ = v15
                                _ = v16
                                _ = v17
                            }
                        }
                        v3 = v3
                        _ = v9
                    }
                }
            }
            v18 := []int{}
            _ = v18
        }
        var v19 uint = (v2 << ((0) & 63))
        _ = fi8
        _ = v19
    }
    for fi20 := 0; fi20 < 17; fi20++ {
        fmt.Println(v4)
        _ = fi20
    }
    var v21 int = 0
    for fi22 := 0; fi22 < 16; fi22++ {
        var v23 uint8 = 0
        _ = fi22
        _ = v23
    }
    v2 += (v2 - v2)
    var v24 int8 = (v4 - v4)
    if (!(127 != 140)) {
        if false {
            v25 := GenS2{F0: v1, F1: ((v2 / ((v2) | 1)) <= v2)}
            var v26 uint = (v2 - v2)
            var v27 int32 = -1
            _ = v25
            _ = v26
            _ = v27
        } else {
            if ((!true) && (!true)) {
                fmt.Println(v21, v4)
                var v28 float64 = 3.14159
                v21 = (v21 + (^(v21)))
                _ = v28
            }
        }
    }
    v4 = (v24 * v24)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v24
    _ = v3
    _ = v4
}
