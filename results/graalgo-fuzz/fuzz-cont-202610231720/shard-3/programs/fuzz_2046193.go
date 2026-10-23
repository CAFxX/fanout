package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 int }

func main() {
    var v1 bool = (-894 <= 16)
    var v2 int16 = 0
    var v3 int16 = v2
    var v4 int32 = 1
    v1 = v1
    v3 -= v2
    for fi5 := 0; fi5 < 2; fi5++ {
        if (-128 > -1) {
            v3 += v2
            var v6 bool = v1
            _ = v6
        }
        if ((v4 * v4) < (v4 + v4)) {
            v4 = (v4 | (^(v4)))
            v2 = v2
        }
        v3 -= 1000
        _ = fi5
    }
    v7 := []int{641, 467, -429}
    var v8 string = "\n\t"
    switch 2147483647 {
    case -1:
        v9 := GenS1{F0: 42, F1: 409, F2: 42}
        _ = v9
    case 128:
        for fi10 := 0; fi10 < 14; fi10++ {
            v1 = false
            v3 += (v2 << ((2) & 63))
            _ = fi10
        }
    default:
        var v11 int = 0
        _ = v11
    }
    v12 := GenS1{F0: 212, F1: 1, F2: 32768}
    if (!false) {
        v13 := GenS1{F0: 127, F1: 18446744073709551615, F2: 4294967295}
        var v14 int64 = 0
        _ = v13
        _ = v14
    } else {
        if (-1 <= -745) {
            fmt.Println(v8, v2)
            var v15 int32 = (v4 - v4)
            _ = v15
        } else {
            if (65535 <= 65535) {
                v8 = v8
            } else {
                for fi16 := 0; fi16 < 14; fi16++ {
                    v17 := GenS1{F0: 2, F1: 2, F2: (fi16 - (fi16 >> ((2) & 63)))}
                    switch "foo bar" {
                    case "a":
                        if len(v7) > 0 {
                            v7[0] = fi16
                        }
                    case "\n\t":
                        v4 = 1
                    default:
                        v18 := GenS1{F0: 127, F1: 0, F2: 127}
                        _ = v18
                    }
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                    _ = fi16
                    _ = v17
                }
                for fi19 := 0; fi19 < 15; fi19++ {
                    for fi20 := 0; fi20 < 19; fi20++ {
                        v1 = ((fi19 % ((-2147483648) | 1)) >= (-(fi19)))
                        v1 = v1
                        v21 := GenS1{F0: 134, F1: 255, F2: fi19}
                        _ = fi20
                        _ = v21
                    }
                    var v22 int = (fi19 / ((128) | 1))
                    for fi23 := 0; fi23 < 18; fi23++ {
                        v22 -= fi19
                        v7 = append(v7, v22)
                        _ = fi23
                    }
                    _ = fi19
                    _ = v22
                }
            }
            var v24 int64 = 485
            if ((127 != 197) && true) {
                _ = len(v7)
                var v25 uint8 = 127
                if len(v7) > 0 {
                    v7[0] = 32767
                }
                _ = v25
            }
            _ = v24
        }
        var v26 int64 = -213
        v1 = (637 == 2147483647)
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
