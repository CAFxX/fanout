package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool; F1 uint32; F2 uint8 }
func genFn0(p0 int16, p1 int64, p2 uint32) bool {
    var v5 uint32 = 65535
    if ((p0 >= (p0 * 1000)) && ((4294967295 != 2) && (127 != 42))) {
        for fi6 := 0; fi6 < 20; fi6++ {
            for fi7 := 0; fi7 < 17; fi7++ {
                var v8 uint8 = 128
                var v9 int64 = (^(p1))
                var v10 uint32 = (p2 >> ((4294967295) & 63))
                _ = fi7
                _ = v10
                _ = v8
                _ = v9
            }
            switch (fi6 * (-(fi6))) {
            case -225:
                switch (("a" + ("a" + (("\n\t" + ("" + "hello")) + "a"))) + "a") {
                case "hello":
                    v11 := []int{fi6}
                    _ = v11
                }
            case -616:
                v12 := GenS1{F0: -897}
                _ = v12
            default:
                if (float64(0.0) < 1.5) {
                    v13 := GenS1{F0: 1}
                    _ = v13
                } else {
                    var v14 uint8 = 2
                    _ = v14
                }
            }
            _ = fi6
        }
        v15 := []int{-1}
        v16 := GenS1{F0: 2147483647}
        _ = v15
        _ = v16
    } else {
        p0 -= -1
        p0 = p0
        if true {
            p0 = (p0 & (p0 % ((142) | 1)))
            if ("a" < "foo bar") {
                if (-128 <= 26) {
                    v17 := GenS1{F0: 343}
                    _ = v17
                }
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return ((!(true || (p1 >= p1))) && (33 <= 0))
}
func genFn1(p0 uint64, p1 bool) string {
    for fi18 := 0; fi18 < 9; fi18++ {
        var v19 int8 = 18
        if ((p0 / (((p0 * p0)) | 1)) < (p0 + p0)) {
            if p1 {
                var v20 bool = false
                _ = v20
            }
            if true {
                switch 1 {
                case -206:
                    fmt.Println(p1)
                case 42:
                    v21 := GenS1{F0: 42}
                    _ = v21
                case 9223372036854775807:
                    var v22 int16 = -1
                    _ = v22
                default:
                    var v23 uint8 = 128
                    _ = v23
                }
                if (1e10 >= 1e10) {
                    v24 := []int{(fi18 & fi18), int(v19), -363}
                    fmt.Println(fi18, v19)
                    _ = v24
                } else {
                    var v25 uint32 = 2
                    v19 -= v19
                    v19 += 42
                    _ = v25
                }
            }
            v26 := GenS2{F0: ((0 >= -968) && ((127 >= (v19 | v19)) || p1)), F1: 42, F2: 248}
            _ = v26
        }
        _ = fi18
        _ = v19
    }
    var v27 int32 = 704
    var v28 int64 = 9223372036854775807
    v27 = (-(v27))
    _ = p0
    _ = p1
    _ = v27
    _ = v28
    return "\n\t"
}

func main() {
    v1 := []int{-2147483648}
    var v2 int8 = 54
    var v3 uint32 = 2147483647
    var v4 uint = 2
    var v29 uint32 = v3
    v1 = append(v1, -878)
    v29 = 0
    v2 -= (v2 ^ v2)
    var v30 int8 = (v2 + -115)
    fmt.Println(v4)
    switch 1 {
    case 128:
        v3 += (v3 % ((2147483648) | 1))
    }
    for fi31 := 0; fi31 < 1; fi31++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi31
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
