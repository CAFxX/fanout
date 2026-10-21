package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 bool }
func genFn0(p0 uint8) uint8 {
    if ((p0 & (p0 - (^(p0)))) > 197) {
        fmt.Println(p0)
    } else {
        if (!(true && true)) {
            for fi6 := 0; fi6 < 11; fi6++ {
                var v7 int8 = 127
                v8 := GenS1{F0: ""}
                _ = fi6
                _ = v7
                _ = v8
            }
        }
        var v9 uint8 = p0
        var v10 uint = 2
        _ = v10
        _ = v9
    }
    var v11 uint = 65535
    _ = p0
    _ = v11
    return 248
}

func main() {
    v1 := []uint32{}
    v2 := make(map[uint32]bool)
    v3 := make(map[int]bool)
    var v4 int = 2147483647
    var v5 float64 = 3.14159
    if true {
        var v12 int16 = 1000
        _ = v12
    }
    v13 := v2[4294967295]
    _ = len(v1)
    v5 += (0.0 / v5)
    v5 += (1.5 - (1e10 * (float64(0.0) + v5)))
    v14 := GenS1{F0: "foo bar"}
    v3[9223372036854775807] = ((18446744073709551615 <= 65535) && false)
    _ = len(v1)
    v15 := v2[1]
    if v13 {
        if (-475 != -2147483648) {
            v4 = v4
            var v16 uint32 = 952
            var v17 uint = 1
            _ = v16
            _ = v17
        } else {
            for fi18 := 0; fi18 < 1; fi18++ {
                if (int64(v4) > 9223372036854775807) {
                    v13 = v15
                }
                if true {
                    fmt.Println(fi18, v15)
                    for fi19 := 0; fi19 < 5; fi19++ {
                        v20 := GenS2{F0: 65535}
                        _ = fi19
                        _ = v20
                    }
                    v13 = v13
                } else {
                    var v21 uint32 = 0
                    _ = v21
                }
                v15 = (!v15)
                _ = fi18
            }
            for fi22 := 0; fi22 < 10; fi22++ {
                if v15 {
                    v4 = fi22
                    var v23 uint8 = 167
                    _ = v23
                } else {
                    v4 += v4
                    fmt.Println(v15, v5, v5)
                }
                v15 = v13
                var v24 float64 = (-2.25 / v5)
                _ = fi22
                _ = v24
            }
        }
        for fi25 := 0; fi25 < 15; fi25++ {
            v13 = v13
            _ = fi25
        }
        if (1.5 >= 0.1) {
            v15 = false
            for fi26 := 0; fi26 < 13; fi26++ {
                v27 := GenS3{F0: true}
                _ = fi26
                _ = v27
            }
        }
    }
    if true {
        v28 := GenS3{F0: ("hello" <= ("" + "a"))}
        switch "foo bar" {
        case "a":
            v13 = (1 <= 231)
        case "foo bar":
            fmt.Println(v13, v4, v5)
        default:
            v5 -= v5
        }
        _ = v28
    }
    v13 = v13
    fmt.Println(v5, v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
