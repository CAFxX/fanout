package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 bool }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint; F1 bool }
type GenS5 struct { F0 uint32; F1 int64 }
type GenS6 struct { F0 int; F1 string }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint16; F1 uint64 }

func main() {
    v1 := make(map[uint32]int)
    var v2 float64 = 1e10
    var v3 uint16 = 255
    var v4 int16 = 997
    var v5 int32 = -1
    v6 := []uint32{2147483647, 72, 65535}
    for fi7 := 0; fi7 < 15; fi7++ {
        switch (fi7 & fi7) {
        case 128:
            if false {
                v5 -= -360
            } else {
                v5 += (v5 - v5)
                for fi8 := 0; fi8 < 15; fi8++ {
                    v9 := GenS1{F0: -1, F1: 562, F2: (42 <= v3)}
                    fmt.Println(v5)
                    v5 = v5
                    v10 := GenS2{F0: 743, F1: "hello"}
                    v11 := GenS3{F0: -18}
                    _ = fi8
                    _ = v10
                    _ = v11
                    _ = v9
                }
            }
        default:
            if (2 == 127) {
                v4 = -32768
                fmt.Println(v3, v4, v2)
            }
        }
        v12 := GenS4{F0: 1, F1: (!false)}
        v5 = v5
        _ = fi7
        _ = v12
    }
    var v13 int8 = 42
    if (false || false) {
        fmt.Println(v3)
    } else {
        for fi14 := 0; fi14 < 1; fi14++ {
            for fi15 := 0; fi15 < 10; fi15++ {
                delete(v1, uint32(v3))
                if len(v6) > 0 {
                    v6[0] = 1
                }
                v2 += (v2 / (v2 * v2))
                switch (fi15 ^ fi15) {
                case -525:
                    v13 += (v13 & (v13 ^ (v13 & -7)))
                case -9223372036854775808:
                    v16 := GenS5{F0: 157, F1: -197}
                    _ = v16
                }
                _ = fi15
            }
            v13 = v13
            if false {
                v5 = v5
                var v17 int = (-(fi14))
                _ = v17
            } else {
                var v18 float64 = (((v2 / v2) - 3.14159) * v2)
                _ = v18
            }
            _ = fi14
        }
        for fi19 := 0; fi19 < 10; fi19++ {
            v5 = (v5 | 0)
            _ = fi19
        }
    }
    fmt.Println(v5)
    for fi20 := 0; fi20 < 2; fi20++ {
        var v21 float32 = float32(0.0)
        for fi22 := 0; fi22 < 17; fi22++ {
            v13 += v13
            var v23 uint = 315
            _ = fi22
            _ = v23
        }
        var v24 float64 = v2
        _ = fi20
        _ = v21
        _ = v24
    }
    if (false || true) {
        fmt.Println(v5, v5)
        for fi25 := 0; fi25 < 9; fi25++ {
            for fi26 := 0; fi26 < 4; fi26++ {
                v1[42] = (fi26 * 4294967295)
                var v27 int32 = 2147483647
                _ = fi26
                _ = v27
            }
            var v28 int32 = int32(v4)
            v13 = (v13 << ((61) & 63))
            _ = fi25
            _ = v28
        }
        if false {
            if false {
                for fi29 := 0; fi29 < 15; fi29++ {
                    if true {
                        delete(v1, 0)
                    }
                    _ = fi29
                }
                v5 = (v5 / (((v5 << ((641) & 63))) | 1))
            }
            v30 := GenS6{F0: 97, F1: ("" + ("hello" + ("hello" + ((("a" + "\n\t") + "\n\t") + "a"))))}
            v31 := GenS6{F0: 762, F1: "hello"}
            _ = v30
            _ = v31
        } else {
            if false {
                v4 = v4
                v32 := GenS2{F0: 2147483648, F1: ("hello" + "\n\t")}
                var v33 int16 = v4
                _ = v32
                _ = v33
            } else {
                v34 := GenS7{F0: 0}
                _ = v34
            }
            v35 := v1[42]
            _ = v35
        }
    }
    for fi36 := 0; fi36 < 3; fi36++ {
        fmt.Println(v13)
        if (((!true) && true) || true) {
            v37 := GenS8{F0: v3, F1: 42}
            var v38 uint16 = v3
            _ = v37
            _ = v38
        } else {
            if false {
                switch "\n\t" {
                case "":
                    switch (fi36 + -9223372036854775808) {
                    case 285:
                        var v39 string = "\n\t"
                        _ = v39
                    default:
                        v4 = 0
                    }
                }
                for fi40 := 0; fi40 < 3; fi40++ {
                    fmt.Println(v2, fi40)
                    _ = fi40
                }
            }
            var v41 float64 = ((0.1 * 1.5) / 1e10)
            _ = v41
        }
        _ = fi36
    }
    var v42 uint16 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
