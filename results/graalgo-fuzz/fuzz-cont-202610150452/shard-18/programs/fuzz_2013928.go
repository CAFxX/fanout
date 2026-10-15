package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32 }
func genFn0(p0 int) int {
    p0 = p0
    var v5 int64 = -107
    for fi6 := 0; fi6 < 11; fi6++ {
        p0 = (fi6 << ((4294967295) & 63))
        var v7 float64 = float64(1.5)
        var v8 uint16 = 0
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 int64 = (v5 / ((v5) | 1))
    if ((false || true) || (false || false)) {
        v9 -= (v5 ^ (v5 << ((2) & 63)))
    } else {
        for fi10 := 0; fi10 < 8; fi10++ {
            for fi11 := 0; fi11 < 3; fi11++ {
                if (0 <= 42) {
                    switch (fi10 >> ((255) & 63)) {
                    case 9223372036854775807:
                        var v12 int = (fi11 % (((p0 - (fi11 >> ((65535) & 63)))) | 1))
                        _ = v12
                    case -529:
                        v13 := []int{(fi10 / ((p0) | 1)), -360, 9223372036854775807, (fi10 % (((fi10 * fi10)) | 1))}
                        _ = v13
                    case 353:
                        if ((2147483648 == 4294967295) && (!(false && false))) {
                            var v14 uint32 = 967
                            p0 = fi11
                            v9 = (v5 ^ v5)
                            _ = v14
                        }
                    }
                }
                switch (-(fi11)) {
                case 963:
                    v15 := []int{}
                    _ = v15
                default:
                    p0 -= fi10
                }
                switch "foo bar" {
                case "\n\t":
                    v16 := GenS1{F0: 951}
                    fmt.Println(fi10, fi10, v9)
                    _ = v16
                default:
                    v5 += v9
                }
                _ = fi11
            }
            _ = fi10
        }
    }
    _ = p0
    _ = v5
    _ = v9
    return p0
}

func main() {
    var v1 uint64 = 1
    var v2 uint = 1
    v3 := make(map[uint32]bool)
    v4 := make(map[string]string)
    if (1 == 0) {
        v2 += (v2 * (v2 / ((4294967295) | 1)))
        var v17 int8 = 100
        if true {
            var v18 uint8 = 202
            for fi19 := 0; fi19 < 17; fi19++ {
                switch "a" {
                case "\n\t":
                    v20 := GenS1{F0: 65535}
                    _ = v20
                default:
                    v21 := GenS2{F0: (!true)}
                    _ = v21
                }
                _ = fi19
            }
            _ = v18
        } else {
            v22 := []int{9223372036854775807}
            if (1e10 != 1.5) {
                for fi23 := 0; fi23 < 15; fi23++ {
                    v17 = int8(fi23)
                    if (433 >= 2) {
                        v17 += (v17 + (^(v17)))
                    } else {
                        v24 := v3[65535]
                        fmt.Println(v2, v1)
                        switch fi23 {
                        case 32767:
                            v25 := GenS3{F0: 175}
                            _ = v25
                        default:
                            var v26 uint = (v2 - v2)
                            _ = v26
                        }
                        _ = v24
                    }
                    _ = fi23
                }
            }
            _ = v22
        }
        _ = v17
    }
    for fi27 := 0; fi27 < 8; fi27++ {
        var v28 int = (fi27 / (((fi27 - fi27)) | 1))
        _ = fi27
        _ = v28
    }
    v2 = (^(v2))
    v29 := GenS1{F0: 42}
    switch -2147483648 {
    case -395:
        v1 = v1
    case 361:
        v30 := []int{128}
        _ = v30
    case 2147483647:
        var v31 float64 = -2.25
        _ = v31
    default:
        for fi32 := 0; fi32 < 14; fi32++ {
            var v33 int32 = 2147483647
            _ = fi32
            _ = v33
        }
    }
    for fi34 := 0; fi34 < 17; fi34++ {
        v1 = v1
        _ = fi34
    }
    var v35 int64 = 9223372036854775807
    for fi36 := 0; fi36 < 2; fi36++ {
        v37 := GenS3{F0: 2147483647}
        _ = fi36
        _ = v37
    }
    for fi38 := 0; fi38 < 13; fi38++ {
        v39 := GenS1{F0: 0}
        v40 := []int{-9223372036854775808, 4294967295}
        v35 = v35
        _ = fi38
        _ = v39
        _ = v40
    }
    var v41 bool = true
    var v42 float32 = (-2.25 + 0.1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v35
    _ = v4
    _ = v41
    _ = v42
}
