package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int64; F1 bool }
type GenS5 struct { F0 uint8; F1 uint }
type GenS6 struct { F0 int8; F1 uint; F2 bool }
type GenS7 struct { F0 int16; F1 int8 }
type GenS8 struct { F0 int; F1 bool }
type GenS9 struct { F0 uint32 }
func genFn0() int8 {
    if (false && true) {
        if (false || (2147483647 != -2147483648)) {
            v4 := GenS1{F0: -9223372036854775808}
            _ = v4
        } else {
            if (138 > 2147483648) {
                for fi5 := 0; fi5 < 14; fi5++ {
                    for fi6 := 0; fi6 < 2; fi6++ {
                        fmt.Println(fi6, fi5)
                        switch 744 {
                        case 32767:
                            v7 := GenS1{F0: 429}
                            _ = v7
                        case -985:
                            v8 := GenS2{F0: -1}
                            _ = v8
                        }
                        _ = fi6
                    }
                    v9 := []int{-842, fi5}
                    fmt.Println(fi5)
                    _ = fi5
                    _ = v9
                }
                v10 := GenS3{F0: 669}
                _ = v10
            }
        }
        for fi11 := 0; fi11 < 18; fi11++ {
            if false {
                var v12 bool = (!((0 == 0) && false))
                _ = v12
            }
            _ = fi11
        }
        var v13 uint32 = 0
        _ = v13
    }
    var v14 uint = 2
    v15 := []int{0}
    v14 = (v14 * 821)
    v16 := GenS4{F0: 0, F1: true}
    _ = v14
    _ = v15
    _ = v16
    return 100
}
func genFn1(p0 uint8, p1 int, p2 uint8) uint {
    v17 := GenS5{F0: p2, F1: 0}
    switch (p1 >> ((4294967295) & 63)) {
    case 127:
        p1 = 217
    case -40:
        for fi18 := 0; fi18 < 9; fi18++ {
            fmt.Println(p2, fi18, p0)
            switch (p1 / ((p1) | 1)) {
            case 962:
                var v19 float32 = -2.25
                _ = v19
            case 4294967295:
                p0 -= 0
            }
            _ = fi18
        }
    case -695:
        v20 := []int{p1}
        _ = v20
    default:
        p0 += p0
    }
    if true {
        var v21 float64 = float64(-2.25)
        _ = v21
    } else {
        v22 := GenS6{F0: 0, F1: 2, F2: (32767 <= 1)}
        var v23 uint8 = p2
        v24 := GenS7{F0: -249, F1: 30}
        _ = v22
        _ = v23
        _ = v24
    }
    switch (p1 % ((p1) | 1)) {
    case 42:
        if false {
            p1 += p1
            for fi25 := 0; fi25 < 11; fi25++ {
                fmt.Println(p2, p1)
                _ = fi25
            }
        } else {
            fmt.Println(p2)
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return 255
}

func main() {
    var v1 uint64 = 466
    var v2 uint32 = 684
    v3 := []uint32{v2, (v2 / (((v2 ^ v2)) | 1)), (^(v2)), v2}
    for fi26 := 0; fi26 < 10; fi26++ {
        var v27 string = "foo bar"
        var v28 uint64 = 2
        _ = fi26
        _ = v27
        _ = v28
    }
    var v29 int64 = 851
    v30 := GenS8{F0: 2147483647, F1: (!false)}
    var v31 uint16 = 255
    v32 := GenS9{F0: (v2 - (v2 + 909))}
    var v33 uint = 313
    v2 += 42
    v3 = append(v3, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
}
