package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 uint16; F1 uint64; F2 uint16 }
func genFn0(p0 bool, p1 uint, p2 int16) string {
    v5 := []int{-578, 2147483647}
    var v6 int8 = int8(p1)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v7 uint32 = 610
    fmt.Println(v6, v6, v7)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return "hello"
}

func main() {
    v1 := []string{"hello", "hello", ("" + "hello")}
    var v2 int16 = -1
    var v3 string = ("\n\t" + "foo bar")
    v4 := make(map[uint32]uint32)
    if true {
        _ = len(v1)
        fmt.Println(v2, v2)
    } else {
        var v8 uint = 42
        var v9 bool = (0 != 0)
        _ = v8
        _ = v9
    }
    fmt.Println(v2)
    var v10 int32 = 402
    if ((1.5 == (((0.1 - float64(0.0)) / float64((0.1 + (-2.25 - ((0.0 * 0.0) - 0.0))))) / ((1e10 / 0.1) + ((3.14159 / ((0.1 + -2.25) * (0.1 * -2.25))) - 0.1)))) && (!(true && (-426 == -178)))) {
        var v11 uint8 = 42
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v11
    }
    switch -951 {
    case 127:
        for fi12 := 0; fi12 < 10; fi12++ {
            v13 := GenS1{F0: v10, F1: 0}
            _ = fi12
            _ = v13
        }
    case 32768:
        fmt.Println(v2)
    case -440:
        v4[42] = 362
    }
    for fi14 := 0; fi14 < 17; fi14++ {
        var v15 int16 = (v2 << ((170) & 63))
        _ = fi14
        _ = v15
    }
    v16 := GenS1{F0: v10, F1: 18446744073709551615}
    var v17 uint32 = 1
    var v18 float32 = (float32(1e10) - float32((0.0 + 3.14159)))
    if (1 > -9223372036854775808) {
        for fi19 := 0; fi19 < 10; fi19++ {
            switch "foo bar" {
            case "a":
                if ((191 > 127) && (!(!true))) {
                    v17 -= v17
                    v17 = v17
                    for fi20 := 0; fi20 < 13; fi20++ {
                        fmt.Println(fi20)
                        _ = fi20
                    }
                } else {
                    var v21 int = (^(fi19))
                    v21 -= v21
                    v18 = ((v18 * (v18 / -2.25)) + (v18 + -2.25))
                    _ = v21
                }
            case "\n\t":
                var v22 uint8 = 84
                _ = v22
            default:
                switch (fi19 ^ -9223372036854775808) {
                case 2147483647:
                    switch (fi19 ^ 128) {
                    case 42:
                        v1 = append(v1, (v3 + v3))
                    case 128:
                        v23 := GenS1{F0: (v10 << ((0) & 63)), F1: 458}
                        _ = v23
                    case -13:
                        var v24 int8 = -106
                        _ = v24
                    }
                case 1:
                    for fi25 := 0; fi25 < 6; fi25++ {
                        v3 = v3
                        var v26 uint32 = (^(v17))
                        var v27 int16 = v2
                        _ = fi25
                        _ = v26
                        _ = v27
                    }
                default:
                    if false {
                        v2 = 0
                        fmt.Println(v17, fi19, fi19)
                    } else {
                        delete(v4, v17)
                    }
                }
            }
            var v28 uint8 = 0
            _ = fi19
            _ = v28
        }
        for fi29 := 0; fi29 < 17; fi29++ {
            switch fi29 {
            case -1:
                var v30 float32 = (-2.25 * v18)
                _ = v30
            case 139:
                v2 = v2
            default:
                v31 := GenS2{F0: 228, F1: 809, F2: 49}
                _ = v31
            }
            _ = fi29
        }
        fmt.Println(v2, v17)
    }
    var v32 int16 = (v2 | (^(v2)))
    if false {
        fmt.Println(v17, v10, v2)
        for fi33 := 0; fi33 < 5; fi33++ {
            if false {
                if (2 > 4294967295) {
                    var v34 int = 4294967295
                    var v35 uint64 = uint64(v17)
                    var v36 int64 = 9223372036854775807
                    _ = v34
                    _ = v35
                    _ = v36
                }
            } else {
                switch -1 {
                case 32768:
                    for fi37 := 0; fi37 < 15; fi37++ {
                        v32 += v32
                        _ = fi37
                    }
                default:
                    if true {
                        v2 = 532
                    }
                }
            }
            if true {
                fmt.Println(fi33)
            }
            _ = fi33
        }
    } else {
        fmt.Println(v32, v2)
    }
    v38 := GenS1{F0: (v10 | 802), F1: 2}
    for fi39 := 0; fi39 < 13; fi39++ {
        switch (v3 + (v3 + v3)) {
        case "":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case "foo bar":
            v10 = (v10 >> ((255) & 63))
        default:
            fmt.Println(v18, v3, v2)
        }
        v40 := GenS1{F0: (v10 - v10), F1: 282}
        _ = fi39
        _ = v40
    }
    v17 = (v17 & (^(v17)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v32
    _ = v38
    _ = v4
}
