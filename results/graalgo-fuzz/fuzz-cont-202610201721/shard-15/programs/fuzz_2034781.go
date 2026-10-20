package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint32; F1 uint32; F2 string }
func genFn0(p0 uint8, p1 uint64, p2 uint64) int32 {
    switch -2147483648 {
    case 42:
        v4 := GenS1{F0: 948}
        _ = v4
    default:
        p0 = p0
    }
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    return 1
}
func genFn1(p0 uint32, p1 bool, p2 int8) uint16 {
    if false {
        fmt.Println(p0)
        for fi5 := 0; fi5 < 6; fi5++ {
            v6 := GenS2{F0: 1, F1: 236}
            if (4294967295 == 213) {
                var v7 uint16 = 65535
                v8 := []int{fi5, int(p0), fi5, fi5}
                _ = v7
                _ = v8
            } else {
                v9 := []int{-1, fi5, 1, fi5}
                var v10 int32 = -189
                _ = v10
                _ = v9
            }
            _ = fi5
            _ = v6
        }
        if ((p1 || p1) || p1) {
            for fi11 := 0; fi11 < 3; fi11++ {
                v12 := GenS3{F0: uint64(p0)}
                if false {
                    for fi13 := 0; fi13 < 17; fi13++ {
                        var v14 uint16 = 2
                        _ = fi13
                        _ = v14
                    }
                } else {
                    p2 -= 1
                    if ((-9223372036854775808 < -1) || (0 > 128)) {
                        v15 := []int{}
                        _ = v15
                    } else {
                        p1 = (-221 >= 42)
                    }
                }
                _ = fi11
                _ = v12
            }
            for fi16 := 0; fi16 < 18; fi16++ {
                var v17 int16 = 1
                p2 += -128
                _ = fi16
                _ = v17
            }
        }
    }
    p2 = p2
    var v18 uint32 = 4294967295
    var v19 int32 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v19
    return 63
}

func main() {
    var v1 uint16 = 694
    var v2 uint64 = 1
    var v3 int16 = 1000
    v20 := []int{0, 4294967295}
    switch ("a" + ("hello" + "hello")) {
    case "foo bar":
        v21 := GenS1{F0: 18446744073709551615}
        _ = v21
    case "\n\t":
        var v22 int16 = (^(v3))
        _ = v22
    default:
        if true {
            var v23 uint32 = 2
            if len(v20) > 0 {
                v20[0] = 128
            }
            _ = v23
        } else {
            switch "hello" {
            case "\n\t":
                var v24 float64 = (((0.0 + ((1.5 * float64(0.1)) - (0.0 + 0.0))) - 1e10) - 1e10)
                _ = v24
            }
            fmt.Println(v2, v3)
            var v25 int64 = -9223372036854775808
            _ = v25
        }
    }
    fmt.Println(v1, v1)
    fmt.Println(v3, v3, v1)
    for fi26 := 0; fi26 < 12; fi26++ {
        v3 = v3
        _ = fi26
    }
    v2 -= (v2 << ((831) & 63))
    v2 += (v2 ^ v2)
    for fi27 := 0; fi27 < 11; fi27++ {
        switch ("\n\t" + ("" + "")) {
        case "foo bar":
            var v28 uint64 = (v2 + (v2 >> ((980) & 63)))
            _ = v28
        default:
            v3 += -32768
        }
        _ = fi27
    }
    fmt.Println(v1, v2, v3)
    for fi29 := 0; fi29 < 6; fi29++ {
        var v30 bool = false
        var v31 int8 = 0
        fmt.Println(v31)
        _ = fi29
        _ = v30
        _ = v31
    }
    for fi32 := 0; fi32 < 12; fi32++ {
        var v33 bool = (!(-1 <= genFn0(122, v2, (v2 + (v2 / (((v2 % (((v2 << ((780) & 63))) | 1))) | 1))))))
        if len(v20) > 0 {
            v20[0] = (fi32 - fi32)
        }
        _ = len(v20)
        _ = fi32
        _ = v33
    }
    v34 := GenS4{F0: 100}
    v35 := GenS5{F0: 390}
    v2 -= 18446744073709551615
    v36 := GenS6{F0: 619, F1: 468, F2: ((((("hello" + "\n\t") + "a") + "\n\t") + "") + "foo bar")}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v34
    _ = v35
    _ = v36
}
