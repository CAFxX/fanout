package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 bool; F1 int8 }
func genFn0(p0 uint, p1 int64) bool {
    var v7 bool = (1 < 1)
    p0 = p0
    fmt.Println(p0, v7, p0)
    _ = p0
    _ = p1
    _ = v7
    return v7
}

func main() {
    var v1 bool = true
    var v2 bool = false
    var v3 bool = (9223372036854775807 < -575)
    var v4 uint16 = 0
    var v5 int8 = 107
    v6 := []uint32{}
    var v8 int16 = 561
    var v9 uint8 = 131
    v1 = ((v4 >> ((255) & 63)) > 289)
    var v10 int32 = 2147483647
    var v11 string = "\n\t"
    if (v1 || (255 > 0)) {
        switch 32767 {
        case -495:
            for fi12 := 0; fi12 < 16; fi12++ {
                fmt.Println(v8)
                _ = fi12
            }
        default:
            v6 = append(v6, 599)
        }
        if v3 {
            var v13 int16 = v8
            for fi14 := 0; fi14 < 19; fi14++ {
                v5 = (v5 ^ 0)
                v13 += (-(v13))
                switch fi14 {
                case -1:
                    if len(v6) > 0 {
                        v6[0] = 0
                    }
                case -9223372036854775808:
                    v4 = v4
                default:
                    v15 := GenS1{F0: 167, F1: 255}
                    _ = v15
                }
                _ = fi14
            }
            v3 = true
            _ = v13
        } else {
            if (true || (!v3)) {
                if (212 >= (v9 & v9)) {
                    v16 := GenS1{F0: 384, F1: 243}
                    switch -2147483648 {
                    case 127:
                        var v17 string = "foo bar"
                        _ = v17
                    default:
                        var v18 uint8 = (v9 / ((v9) | 1))
                        _ = v18
                    }
                    _ = v16
                }
            } else {
                v2 = v2
                fmt.Println(v5, v1, v10)
            }
            v11 = v11
        }
        var v19 uint = 2
        _ = v19
    } else {
        if (!(!((127 >= 1) || v2))) {
            if (!v3) {
                for fi20 := 0; fi20 < 20; fi20++ {
                    var v21 uint32 = 2147483647
                    var v22 int32 = v10
                    _ = fi20
                    _ = v21
                    _ = v22
                }
                var v23 uint32 = 4294967295
                _ = v23
            } else {
                var v24 int32 = 2147483647
                fmt.Println(v9, v10)
                _ = v24
            }
            for fi25 := 0; fi25 < 12; fi25++ {
                if true {
                    v26 := GenS1{F0: 364, F1: 797}
                    _ = v26
                } else {
                    v27 := GenS2{F0: v2, F1: v5}
                    var v28 int32 = (v10 | 159)
                    _ = v27
                    _ = v28
                }
                fmt.Println(v5, v11, v8)
                _ = fi25
            }
            v5 -= v5
        }
        if v3 {
            var v29 uint32 = 65535
            v2 = ((!(true && (0 > 18446744073709551615))) || v3)
            _ = v29
        } else {
            fmt.Println(v1)
            var v30 bool = v1
            var v31 uint = 214
            _ = v30
            _ = v31
        }
        var v32 uint64 = 2
        _ = v32
    }
    v33 := GenS1{F0: 0, F1: 427}
    fmt.Println(v4, v1, v5)
    v3 = v2
    var v34 uint16 = v4
    fmt.Println(v4, v1, v2)
    var v35 int32 = v10
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
