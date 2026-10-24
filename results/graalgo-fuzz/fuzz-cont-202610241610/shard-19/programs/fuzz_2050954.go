package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
func genFn0(p0 int32, p1 uint64) int {
    var v6 float32 = 1.5
    v7 := []int{9223372036854775807}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 32767
}

func main() {
    var v1 uint16 = 442
    var v2 int64 = 42
    v3 := make(map[int]bool)
    var v4 uint64 = 4294967295
    var v5 uint8 = 128
    if true {
        var v8 uint32 = 903
        if (!false) {
            v2 += (-(v2))
        }
        _ = v8
    }
    var v9 uint64 = (v4 >> ((18446744073709551615) & 63))
    var v10 uint16 = v1
    switch 2147483647 {
    case 9223372036854775807:
        v11 := GenS1{F0: v1, F1: uint32(v5)}
        _ = v11
    case -989:
        var v12 uint64 = (v9 & (v4 / ((v9) | 1)))
        _ = v12
    case 32768:
        if (!true) {
            switch "" {
            case "\n\t":
                v1 = v10
            default:
                v13 := []int{4294967295, -9223372036854775808}
                _ = v13
            }
            var v14 int64 = v2
            _ = v14
        }
    default:
        v15 := []int{genFn0(0, v4), 1000, -507, 127}
        _ = v15
    }
    v10 -= v10
    for fi16 := 0; fi16 < 4; fi16++ {
        if true {
            var v17 uint16 = (v10 / ((v1) | 1))
            _ = v17
        } else {
            if ((v9 & (v9 >> ((0) & 63))) < (v4 | (v4 + v4))) {
                for fi18 := 0; fi18 < 7; fi18++ {
                    if ((v2 ^ (v2 ^ v2)) > -375) {
                        if ((((v5 < (v5 * (v5 | 128))) && false) && (!(735 > 32767))) && (1.5 <= 3.14159)) {
                            if (65535 < (v1 + (v10 / ((383) | 1)))) {
                                if true {
                                    var v19 int = genFn0(-1, 407)
                                    _ = v19
                                }
                            }
                            if (true || (0 <= 4294967295)) {
                                switch "hello" {
                                case "a":
                                    var v20 uint = 1
                                    _ = v20
                                case "foo bar":
                                    var v21 int64 = (v2 & v2)
                                    _ = v21
                                case "":
                                    v3[fi16] = (!(!false))
                                }
                                delete(v3, fi16)
                            } else {
                                fmt.Println(v9)
                                fmt.Println(v2)
                            }
                            var v22 int = genFn0(42, (v4 - (v9 >> ((18446744073709551615) & 63))))
                            _ = v22
                        }
                    }
                    switch (-(fi18)) {
                    case 128:
                        var v23 bool = false
                        _ = v23
                    }
                    fmt.Println(v1, fi16)
                    v24 := []int{(-(fi16))}
                    _ = fi18
                    _ = v24
                }
            } else {
                for fi25 := 0; fi25 < 7; fi25++ {
                    fmt.Println(fi16, fi25)
                    delete(v3, 238)
                    _ = fi25
                }
            }
        }
        _ = fi16
    }
    v1 -= 1
    v10 = (v10 >> ((2) & 63))
    var v26 int32 = -734
    for fi27 := 0; fi27 < 1; fi27++ {
        var v28 bool = true
        fmt.Println(v1, v26)
        _ = fi27
        _ = v28
    }
    v4 -= uint64(v10)
    var v29 uint32 = 2
    var v30 int64 = (v2 >> ((130) & 63))
    if false {
        fmt.Println(v30, v9)
    }
    var v31 float64 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v9
}
