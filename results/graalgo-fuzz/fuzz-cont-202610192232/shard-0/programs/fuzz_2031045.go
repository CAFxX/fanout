package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 int64, p1 bool, p2 int64) bool {
    v5 := GenS1{F0: 128}
    fmt.Println(p0, p2, p1)
    var v6 float64 = 1e10
    if p1 {
        v7 := []int{}
        _ = v7
    }
    p0 -= p0
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return (-2147483648 <= 208)
}
func genFn1(p0 uint8, p1 uint32, p2 int16) int8 {
    p0 = 255
    var v8 int64 = int64(p0)
    p2 += -99
    switch -558 {
    case -786:
        if (308 < 986) {
            fmt.Println(v8, v8, v8)
        }
    case -2147483648:
        var v9 float64 = 1e10
        _ = v9
    case 2147483647:
        if (!(true && false)) {
            v10 := []int{32767, -1}
            p0 = (p0 ^ 127)
            _ = len(v10)
            _ = v10
        } else {
            for fi11 := 0; fi11 < 3; fi11++ {
                if (!true) {
                    v12 := []int{-9223372036854775808, 779, (fi11 + -605)}
                    fmt.Println(fi11)
                    _ = v12
                } else {
                    p1 = (^(p1))
                    var v13 int = (fi11 ^ 2147483647)
                    p0 -= 255
                    _ = v13
                }
                p2 += 42
                _ = fi11
            }
            for fi14 := 0; fi14 < 10; fi14++ {
                if (16 != 102) {
                    v15 := []int{-2147483648, (fi14 & -975), fi14, (fi14 << ((18446744073709551615) & 63))}
                    _ = len(v15)
                    _ = v15
                } else {
                    fmt.Println(p2)
                }
                var v16 int64 = (v8 >> ((1) & 63))
                _ = fi14
                _ = v16
            }
            v17 := GenS1{F0: 0}
            _ = v17
        }
    default:
        var v18 float64 = 0.1
        _ = v18
    }
    fmt.Println(p2, p0, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 0
}

func main() {
    v1 := []uint32{2147483648, 2147483647, 2147483648, 1}
    v2 := make(map[int]uint32)
    v3 := []string{"hello"}
    v4 := make(map[uint32]uint32)
    delete(v2, 32768)
    var v19 uint8 = 96
    fmt.Println(v19)
    switch 9223372036854775807 {
    case -1:
        if (false || (88 <= 42)) {
            v19 = (v19 >> ((18446744073709551615) & 63))
            fmt.Println(v19)
            for fi20 := 0; fi20 < 4; fi20++ {
                fmt.Println(v19)
                _ = fi20
            }
        } else {
            var v21 float32 = -2.25
            _ = v21
        }
    case 825:
        var v22 float64 = (1.5 / ((((float64(-2.25) - float64((0.0 - -2.25))) / 1.0) + 1e10) + 1e10))
        _ = v22
    default:
        v2[32768] = 1
    }
    for fi23 := 0; fi23 < 19; fi23++ {
        var v24 string = ("" + (("a" + "\n\t") + "a"))
        _ = fi23
        _ = v24
    }
    v19 = 255
    fmt.Println(v19)
    if ("\n\t" == "foo bar") {
        v19 = v19
        v19 = (v19 - (v19 % ((v19) | 1)))
    } else {
        var v25 float32 = 1e10
        if (!(1 >= 389)) {
            for fi26 := 0; fi26 < 1; fi26++ {
                for fi27 := 0; fi27 < 18; fi27++ {
                    var v28 uint64 = 0
                    _ = fi27
                    _ = v28
                }
                switch fi26 {
                case 127:
                    var v29 uint32 = 42
                    _ = v29
                default:
                    var v30 int32 = 42
                    _ = v30
                }
                _ = len(v1)
                _ = fi26
            }
        } else {
            if len(v1) > 0 {
                v1[0] = 2147483647
            }
        }
        switch 4294967295 {
        case 32768:
            var v31 bool = true
            _ = v31
        default:
            if (-2.25 < (0.0 / float64((v25 / v25)))) {
                _ = len(v1)
                v19 -= (^(v19))
            } else {
                v1 = append(v1, 23)
            }
        }
        _ = v25
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
