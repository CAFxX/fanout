package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 uint16 }
type GenS2 struct { F0 uint }
func genFn0(p0 int64, p1 int32) uint8 {
    v6 := GenS1{F0: 2147483648, F1: 1, F2: 180}
    v7 := []int{1, 0, -2147483648, 0}
    for fi8 := 0; fi8 < 8; fi8++ {
        var v9 int8 = 100
        _ = fi8
        _ = v9
    }
    switch "hello" {
    case "foo bar":
        fmt.Println(p0, p1)
    case "\n\t":
        v7 = append(v7, -311)
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 255
}

func main() {
    v1 := []string{"\n\t"}
    v2 := make(map[int]uint32)
    var v3 uint32 = 2
    var v4 bool = true
    var v5 string = "a"
    v2[-1] = 381
    v10 := GenS1{F0: (^(v3)), F1: 0, F2: 65535}
    if len(v1) > 0 {
        v1[0] = (v5 + v5)
    }
    var v11 bool = true
    if (0 > -472) {
        for fi12 := 0; fi12 < 10; fi12++ {
            if false {
                var v13 uint8 = 84
                var v14 float32 = (1e10 * -2.25)
                _ = v13
                _ = v14
            }
            _ = fi12
        }
        v3 -= (v3 - (v3 + (v3 % ((v3) | 1))))
    }
    v4 = true
    if v11 {
        switch (v5 + v5) {
        case "hello":
            v1 = append(v1, v5)
        case "\n\t":
            if len(v1) > 0 {
                v1[0] = v5
            }
        case "":
            v5 = v5
        default:
            var v15 uint64 = 966
            _ = v15
        }
        if v4 {
            fmt.Println(v3, v11, v11)
        } else {
            var v16 int16 = -924
            var v17 int16 = v16
            if false {
                if (757 != 2) {
                    switch (v5 + v5) {
                    case "foo bar":
                        v11 = (!false)
                    case "hello":
                        fmt.Println(v16)
                    case "\n\t":
                        delete(v2, 2147483647)
                    }
                    var v18 int64 = -9223372036854775808
                    var v19 string = "\n\t"
                    _ = v18
                    _ = v19
                }
                delete(v2, -1)
            }
            _ = v16
            _ = v17
        }
    }
    switch 32767 {
    case 2147483647:
        for fi20 := 0; fi20 < 13; fi20++ {
            v21 := GenS1{F0: v3, F1: -705, F2: 1}
            _ = fi20
            _ = v21
        }
    case -725:
        var v22 uint64 = 2
        _ = v22
    case 1:
        v3 = v3
    default:
        v23 := GenS1{F0: 1, F1: 0, F2: 65535}
        _ = v23
    }
    v4 = false
    v24 := GenS1{F0: (v3 * 65535), F1: -2147483648, F2: 32767}
    v5 = "a"
    switch 4294967295 {
    case 32767:
        if true {
            var v25 float32 = (3.14159 - (3.14159 - 1.5))
            _ = v25
        } else {
            for fi26 := 0; fi26 < 9; fi26++ {
                v3 = (v3 & v3)
                for fi27 := 0; fi27 < 11; fi27++ {
                    switch "hello" {
                    case "\n\t":
                        v5 = "hello"
                    default:
                        if len(v1) > 0 {
                            v1[0] = "a"
                        }
                    }
                    _ = len(v2)
                    if v4 {
                        v4 = (((v5 + (("a" + v5) + "")) + v5) <= "")
                        fmt.Println(v11)
                    } else {
                        v28 := GenS1{F0: (v3 * (v3 - (v3 << ((466) & 63)))), F1: -1, F2: 198}
                        var v29 uint64 = 4294967295
                        _ = v28
                        _ = v29
                    }
                    v30 := GenS2{F0: 516}
                    _ = fi27
                    _ = v30
                }
                _ = fi26
            }
        }
    case 9223372036854775807:
        v31 := v2[188]
        _ = v31
    case 4294967295:
        var v32 bool = (v11 || (65535 >= 1))
        _ = v32
    default:
        fmt.Println(v4, v11, v3)
    }
    var v33 int8 = 42
    _ = len(v1)
    v34 := GenS2{F0: 4294967295}
    v4 = v11
    switch 0 {
    case -164:
        for fi35 := 0; fi35 < 8; fi35++ {
            v3 = 4294967295
            _ = fi35
        }
    case 42:
        v36 := v2[739]
        _ = v36
    default:
        var v37 int = -29
        _ = v37
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v24
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
