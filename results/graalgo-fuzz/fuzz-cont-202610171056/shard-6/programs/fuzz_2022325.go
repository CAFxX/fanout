package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 uint32 }
type GenS2 struct { F0 bool; F1 bool; F2 int16 }

func main() {
    var v1 int64 = 42
    var v2 uint32 = 42
    v3 := make(map[int]bool)
    var v4 uint16 = 910
    v5 := make(map[uint32]string)
    var v6 string = "hello"
    v4 = v4
    switch int(v4) {
    case 4294967295:
        v7 := GenS1{F0: 42, F1: ((!false) && (false && false)), F2: v2}
        _ = v7
    case 2147483647:
        for fi8 := 0; fi8 < 6; fi8++ {
            for fi9 := 0; fi9 < 14; fi9++ {
                switch (fi9 << ((1) & 63)) {
                case 32767:
                    v10 := GenS1{F0: 2147483647, F1: (v1 > v1), F2: v2}
                    _ = v10
                case 9223372036854775807:
                    switch v6 {
                    case "\n\t":
                        v1 = 0
                    default:
                        switch "\n\t" {
                        case "a":
                            v11 := []int{(fi8 * fi9), (fi9 >> ((0) & 63)), 142, -605}
                            _ = v11
                        case "\n\t":
                            fmt.Println(v4, v2, fi9)
                        }
                    }
                case -2147483648:
                    if (52 > 458) {
                        var v12 uint = 255
                        v13 := GenS2{F0: ((v6 + "a") < "foo bar"), F1: false, F2: 612}
                        _ = v12
                        _ = v13
                    }
                }
                var v14 int = 508
                fmt.Println(v1, v14, v2)
                _ = fi9
                _ = v14
            }
            var v15 float32 = 0.1
            _ = fi8
            _ = v15
        }
    default:
        if false {
            for fi16 := 0; fi16 < 20; fi16++ {
                v17 := []int{fi16, fi16, fi16, int(v2)}
                if (false || (9223372036854775807 != (-(v1)))) {
                    v4 = (v4 * (^(v4)))
                    v1 += (^(v1))
                }
                _ = fi16
                _ = v17
            }
            for fi18 := 0; fi18 < 19; fi18++ {
                v19 := GenS2{F0: (!false), F1: false, F2: 843}
                _ = fi18
                _ = v19
            }
            for fi20 := 0; fi20 < 20; fi20++ {
                switch v6 {
                case "a":
                    var v21 float64 = 3.14159
                    _ = v21
                case "hello":
                    fmt.Println(v6)
                case "":
                    for fi22 := 0; fi22 < 1; fi22++ {
                        v5[(v2 | 65535)] = v6
                        fmt.Println(v1)
                        v23 := []int{4294967295, fi22, fi20, (^(fi22))}
                        if (v4 > v4) {
                            var v24 uint64 = 363
                            var v25 int8 = 83
                            v4 = v4
                            _ = v24
                            _ = v25
                        } else {
                            var v26 uint = 42
                            _ = v26
                        }
                        _ = fi22
                        _ = v23
                    }
                }
                v27 := []int{}
                _ = fi20
                _ = v27
            }
        }
    }
    fmt.Println(v1, v6)
    v28 := []int{-9223372036854775808, -1}
    _ = len(v3)
    v4 += v4
    v1 += v1
    var v29 uint32 = uint32(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
