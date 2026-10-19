package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 int32 }
func genFn0(p0 int16, p1 uint, p2 uint) bool {
    var v7 uint64 = 463
    for fi8 := 0; fi8 < 16; fi8++ {
        for fi9 := 0; fi9 < 8; fi9++ {
            v10 := []int{}
            _ = fi9
            _ = v10
        }
        var v11 bool = false
        _ = fi8
        _ = v11
    }
    switch -961 {
    case 942:
        if (-588 <= 955) {
            if false {
                var v12 string = ("hello" + ("a" + "a"))
                _ = v12
            }
            v13 := []int{1, 42, -9223372036854775808, 279}
            if len(v13) > 0 {
                v13[0] = 128
            }
            _ = v13
        }
    case -881:
        fmt.Println(p1, p1, p1)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return true
}

func main() {
    var v1 uint64 = 55
    var v2 int32 = 1
    v3 := make(map[int]uint32)
    v4 := make(map[string]string)
    v5 := []string{"foo bar", "foo bar", ("hello" + "")}
    var v6 int64 = 0
    if true {
        v1 -= (^(v1))
    } else {
        if (!true) {
            v14 := GenS1{F0: 65535, F1: 0}
            _ = v14
        }
    }
    v5 = append(v5, "\n\t")
    var v15 int32 = v2
    v6 = (v6 % ((-384) | 1))
    v16 := GenS1{F0: 292, F1: -367}
    v17 := GenS1{F0: 861, F1: 42}
    if len(v5) > 0 {
        v5[0] = "foo bar"
    }
    fmt.Println(v1, v1)
    for fi18 := 0; fi18 < 8; fi18++ {
        v6 = v6
        switch "" {
        case "":
            switch "\n\t" {
            case "hello":
                switch -1 {
                case 4294967295:
                    for fi19 := 0; fi19 < 20; fi19++ {
                        v20 := GenS1{F0: 1, F1: -819}
                        switch "a" {
                        case "a":
                            var v21 uint16 = 0
                            _ = v21
                        default:
                            v15 = (-(v2))
                        }
                        if (false || (false || false)) {
                            var v22 bool = (1 >= 128)
                            v3[fi19] = 333
                            var v23 uint = 1
                            _ = v22
                            _ = v23
                        } else {
                            var v24 uint64 = (v1 | v1)
                            _ = v24
                        }
                        v15 -= (v15 - v15)
                        _ = fi19
                        _ = v20
                    }
                case 42:
                    fmt.Println(v6, v1)
                case 968:
                    if false {
                        if (128 > 26) {
                            var v25 float32 = 1.5
                            _ = v25
                        }
                        for fi26 := 0; fi26 < 11; fi26++ {
                            v27 := GenS2{F0: (v2 << ((823) & 63))}
                            if len(v5) > 0 {
                                v5[0] = "hello"
                            }
                            _ = fi26
                            _ = v27
                        }
                    }
                default:
                    var v28 float32 = (-2.25 * 0.1)
                    _ = v28
                }
            case "":
                var v29 uint = 255
                _ = v29
            default:
                var v30 uint8 = 255
                _ = v30
            }
        }
        _ = fi18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
