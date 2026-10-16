package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 int16 }
func genFn0(p0 uint16, p1 uint8, p2 uint) int16 {
    v4 := []int{128}
    v4 = append(v4, -104)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 1000
}
func genFn1(p0 string, p1 string) uint {
    fmt.Println(p0)
    v5 := GenS1{F0: -72, F1: 42, F2: -227}
    _ = p0
    _ = p1
    _ = v5
    return 1
}

func main() {
    var v1 uint32 = 65535
    v2 := make(map[int]uint32)
    var v3 int16 = 32767
    for fi6 := 0; fi6 < 14; fi6++ {
        fmt.Println(v1, v1, v1)
        _ = fi6
    }
    if (((!(2 != 0)) || ("\n\t" < "\n\t")) || (v3 < (v3 % (((v3 ^ v3)) | 1)))) {
        v7 := GenS1{F0: 1, F1: -91, F2: 1}
        _ = v7
    } else {
        switch -1 {
        case 127:
            for fi8 := 0; fi8 < 2; fi8++ {
                switch fi8 {
                case 32767:
                    switch ("hello" + "a") {
                    case "a":
                        var v9 string = "foo bar"
                        _ = v9
                    }
                case 2147483647:
                    v10 := []int{-1, (-(fi8))}
                    _ = v10
                case 42:
                    v1 = v1
                default:
                    if (!(0 == uint64(v1))) {
                        v3 = (v3 & (-(v3)))
                        if true {
                            v2[-943] = (v1 * v1)
                            v1 = 0
                            var v11 uint64 = 4294967295
                            _ = v11
                        }
                        v1 -= v1
                    }
                }
                v12 := GenS1{F0: 649, F1: 0, F2: 1000}
                v1 = (v1 * 2147483647)
                _ = fi8
                _ = v12
            }
        default:
            if false {
                if false {
                    if false {
                        v3 -= -32768
                        var v13 uint8 = 84
                        _ = v13
                    } else {
                        var v14 int64 = 275
                        var v15 uint64 = 0
                        _ = v14
                        _ = v15
                    }
                } else {
                    var v16 uint = 960
                    v1 += 2
                    _ = v16
                }
                switch -121 {
                case -502:
                    if (42 < 1) {
                        v17 := GenS1{F0: -2147483648, F1: 42, F2: -30}
                        if false {
                            var v18 bool = (91 <= 117)
                            fmt.Println(v3)
                            var v19 uint16 = 536
                            _ = v18
                            _ = v19
                        }
                        _ = v17
                    }
                case 0:
                    for fi20 := 0; fi20 < 2; fi20++ {
                        v21 := GenS1{F0: -2147483648, F1: -1, F2: 0}
                        for fi22 := 0; fi22 < 20; fi22++ {
                            v1 += (v1 ^ v1)
                            _ = len(v2)
                            _ = fi22
                        }
                        fmt.Println(v3)
                        _ = fi20
                        _ = v21
                    }
                }
                v23 := []int{-75}
                _ = v23
            }
        }
    }
    var v24 float32 = 0.0
    var v25 bool = false
    switch "a" {
    case "foo bar":
        v25 = (false && (42 <= -9223372036854775808))
    }
    v3 = v3
    v26 := v2[-9223372036854775808]
    for fi27 := 0; fi27 < 11; fi27++ {
        var v28 int = (-(fi27))
        v24 = (0.0 - (float32(1.5) + (3.14159 - ((v24 - 3.14159) * v24))))
        switch fi27 {
        case 4294967295:
            fmt.Println(v1)
        case 32768:
            var v29 int8 = 0
            _ = v29
        default:
            v25 = false
        }
        _ = fi27
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
}
