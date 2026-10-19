package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 int8 }
type GenS2 struct { F0 uint }
func genFn0(p0 int, p1 bool, p2 int) bool {
    var v5 int = (p2 | -2147483648)
    for fi6 := 0; fi6 < 18; fi6++ {
        var v7 int16 = -1
        if (255 >= 347) {
            for fi8 := 0; fi8 < 9; fi8++ {
                var v9 int8 = 42
                v9 = v9
                v10 := GenS1{F0: 9223372036854775807, F1: -389, F2: (v9 + 1)}
                _ = fi8
                _ = v10
                _ = v9
            }
        }
        var v11 int32 = 0
        _ = fi6
        _ = v11
        _ = v7
    }
    p0 -= 4294967295
    v12 := []int{v5, (p0 - v5), -9223372036854775808, p0}
    fmt.Println(p1, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v5
    return (v5 <= p0)
}

func main() {
    var v1 float64 = float64(((3.14159 * -2.25) + 3.14159))
    v2 := make(map[int]uint32)
    var v3 int16 = 1000
    var v4 uint8 = 2
    v13 := GenS1{F0: 9223372036854775807, F1: 297, F2: 53}
    if (false || false) {
        fmt.Println(v1, v4)
        fmt.Println(v3, v4, v4)
    } else {
        v14 := []int{}
        _ = v14
    }
    var v15 uint8 = (v4 >> ((65535) & 63))
    for fi16 := 0; fi16 < 17; fi16++ {
        var v17 int16 = 0
        v18 := GenS2{F0: 111}
        for fi19 := 0; fi19 < 9; fi19++ {
            v4 = v15
            switch fi19 {
            case 42:
                var v20 uint8 = 183
                _ = v20
            case 2147483647:
                v21 := []int{(fi16 % ((-90) | 1)), (fi19 >> ((781) & 63)), -1, 142}
                _ = v21
            default:
                var v22 uint8 = v4
                _ = v22
            }
            _ = fi19
        }
        _ = fi16
        _ = v17
        _ = v18
    }
    if false {
        if true {
            v1 = v1
        }
        switch "hello" {
        case "a":
            v23 := GenS1{F0: 42, F1: -1, F2: 0}
            _ = v23
        case "foo bar":
            v1 = (v1 + v1)
        default:
            v1 -= v1
        }
    }
    for fi24 := 0; fi24 < 6; fi24++ {
        if false {
            if true {
                v15 -= 128
            }
        } else {
            switch "hello" {
            case "foo bar":
                var v25 int16 = (v3 >> ((4294967295) & 63))
                _ = v25
            }
        }
        for fi26 := 0; fi26 < 14; fi26++ {
            var v27 int64 = -1
            v28 := []int{fi26, 341, 0}
            v15 -= 2
            _ = fi26
            _ = v27
            _ = v28
        }
        _ = fi24
    }
    for fi29 := 0; fi29 < 14; fi29++ {
        var v30 uint32 = 0
        if true {
            var v31 uint8 = (v15 << ((201) & 63))
            _ = v31
        }
        _ = fi29
        _ = v30
    }
    fmt.Println(v4)
    if true {
        var v32 int8 = -106
        _ = v32
    } else {
        var v33 int32 = 454
        if true {
            switch 325 {
            case 2147483647:
                fmt.Println(v3, v15, v15)
            default:
                if (50 < 11) {
                    v33 = (v33 << ((88) & 63))
                    var v34 uint16 = 42
                    if false {
                        for fi35 := 0; fi35 < 10; fi35++ {
                            v4 -= v15
                            switch fi35 {
                            case 128:
                                v36 := []int{0}
                                _ = v36
                            default:
                                var v37 float64 = (v1 - v1)
                                _ = v37
                            }
                            var v38 uint64 = 404
                            _ = fi35
                            _ = v38
                        }
                    } else {
                        v39 := GenS1{F0: -717, F1: -1, F2: int8(v3)}
                        _ = v39
                    }
                    _ = v34
                }
            }
        } else {
            v33 = 36
            v40 := GenS2{F0: 437}
            _ = v40
        }
        v4 = v4
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
