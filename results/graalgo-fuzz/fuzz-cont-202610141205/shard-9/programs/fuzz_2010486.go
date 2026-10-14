package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16; F1 int64 }
func genFn0(p0 uint64) bool {
    v7 := []int{}
    if (83 < 42) {
        for fi8 := 0; fi8 < 13; fi8++ {
            p0 += (p0 ^ (p0 << ((42) & 63)))
            _ = fi8
        }
    } else {
        p0 = (^(p0))
        if (1 >= -9223372036854775808) {
            _ = len(v7)
        }
        var v9 float64 = 0.0
        _ = v9
    }
    v10 := GenS1{F0: -53}
    var v11 uint16 = 0
    _ = p0
    _ = v10
    _ = v11
    _ = v7
    return false
}

func main() {
    v1 := []int{128, 559, -207, -1}
    var v2 int32 = 0
    var v3 uint32 = 315
    var v4 uint8 = 127
    var v5 int8 = -72
    var v6 bool = true
    var v12 int32 = v2
    v13 := GenS2{F0: 730, F1: 54}
    v5 -= (v5 - v5)
    fmt.Println(v5, v6, v2)
    if len(v1) > 0 {
        v1[0] = 566
    }
    if v6 {
        if (v6 || (!true)) {
            for fi14 := 0; fi14 < 9; fi14++ {
                v3 = (v3 | 2)
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi14
            }
            v12 += -1
        }
        var v15 string = "a"
        var v16 uint = 65535
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 15; fi17++ {
        var v18 uint16 = 974
        switch "" {
        case "\n\t":
            switch (("foo bar" + "hello") + "a") {
            case "hello":
                for fi19 := 0; fi19 < 1; fi19++ {
                    if v6 {
                        v12 += v2
                    } else {
                        v20 := GenS1{F0: int8(v18)}
                        if false {
                            v6 = v6
                            v1 = append(v1, (fi19 % (((-(fi19))) | 1)))
                        }
                        var v21 float32 = 3.14159
                        var v22 uint64 = 4294967295
                        var v23 uint8 = v4
                        _ = v20
                        _ = v21
                        _ = v22
                        _ = v23
                    }
                    var v24 int32 = (v12 >> ((18446744073709551615) & 63))
                    v3 -= 529
                    _ = fi19
                    _ = v24
                }
            case "a":
                for fi25 := 0; fi25 < 6; fi25++ {
                    fmt.Println(v2)
                    _ = fi25
                }
            case "\n\t":
                fmt.Println(v12, v6)
            }
        default:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        _ = fi17
        _ = v18
    }
    v5 -= -1
    var v26 int16 = 1
    switch 9223372036854775807 {
    case 1:
        if false {
            switch 32768 {
            case 128:
                v2 -= (v12 << ((1) & 63))
            case -9223372036854775808:
                v27 := GenS1{F0: (^(v5))}
                _ = v27
            case 1:
                if (!v6) {
                    v2 += int32(v5)
                    var v28 float32 = (1.5 * (-2.25 - (1e10 - -2.25)))
                    switch 316 {
                    case -101:
                        if false {
                            v2 -= v2
                        }
                    case 32768:
                        v29 := GenS2{F0: 1, F1: 0}
                        _ = v29
                    default:
                        fmt.Println(v6, v5)
                    }
                    _ = v28
                }
            }
            var v30 float64 = 3.14159
            fmt.Println(v6, v12)
            _ = v30
        }
    default:
        if len(v1) > 0 {
            v1[0] = 32768
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi31 := 0; fi31 < 17; fi31++ {
        if false {
            var v32 int = fi31
            for fi33 := 0; fi33 < 17; fi33++ {
                v4 = v4
                var v34 uint8 = (v4 * v4)
                _ = fi33
                _ = v34
            }
            v6 = true
            _ = v32
        }
        switch "" {
        case "":
            for fi35 := 0; fi35 < 6; fi35++ {
                v26 += v26
                _ = len(v1)
                fmt.Println(v3, fi31, v5)
                _ = fi35
            }
        }
        _ = fi31
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
