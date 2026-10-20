package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint64; F1 int }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint }
func genFn0(p0 uint64) string {
    var v5 float64 = 3.14159
    switch "hello" {
    case "foo bar":
        for fi6 := 0; fi6 < 6; fi6++ {
            if (1e10 != v5) {
                fmt.Println(v5, p0, p0)
                var v7 uint64 = (p0 - (p0 - (^(p0))))
                var v8 float32 = 3.14159
                _ = v7
                _ = v8
            } else {
                for fi9 := 0; fi9 < 7; fi9++ {
                    var v10 int16 = 0
                    var v11 uint64 = p0
                    var v12 bool = false
                    p0 += (p0 << ((2) & 63))
                    _ = fi9
                    _ = v10
                    _ = v11
                    _ = v12
                }
                p0 = (^(p0))
                v5 -= 0.1
            }
            var v13 int16 = 42
            fmt.Println(v13)
            _ = fi6
            _ = v13
        }
    case "\n\t":
        p0 = 1
    default:
        for fi14 := 0; fi14 < 6; fi14++ {
            if true {
                for fi15 := 0; fi15 < 6; fi15++ {
                    if false {
                        if (637 >= 1) {
                            v16 := []int{fi14}
                            _ = v16
                        }
                    }
                    v5 -= v5
                    var v17 string = "hello"
                    _ = fi15
                    _ = v17
                }
                v18 := []int{fi14}
                fmt.Println(fi14, p0)
                _ = v18
            }
            _ = fi14
        }
    }
    var v19 float32 = float32((1e10 * float64(float32(v5))))
    _ = p0
    _ = v19
    _ = v5
    return "foo bar"
}
func genFn1(p0 uint16) uint8 {
    p0 -= 0
    p0 -= (p0 ^ p0)
    p0 = p0
    v20 := []int{-430}
    v21 := GenS1{F0: (^(p0))}
    _ = p0
    _ = v20
    _ = v21
    return 169
}

func main() {
    var v1 int = 206
    var v2 uint32 = 2147483648
    var v3 uint32 = (^(v2))
    var v4 string = ("foo bar" + "")
    v4 = v4
    switch -721 {
    case 128:
        if (107 != 115) {
            var v22 string = ("hello" + (("" + (v4 + ("foo bar" + ""))) + ("a" + v4)))
            for fi23 := 0; fi23 < 9; fi23++ {
                v24 := GenS2{F0: v4}
                _ = fi23
                _ = v24
            }
            _ = v22
        } else {
            if true {
                v2 = 1
            }
        }
    case 1:
        v3 = (^(v2))
    }
    v25 := []int{2147483647, -9223372036854775808}
    if false {
        v4 = v4
        for fi26 := 0; fi26 < 17; fi26++ {
            if (-70 != -185) {
                v2 = 2147483647
            } else {
                var v27 bool = (42 > 0)
                var v28 uint8 = genFn1(42)
                v29 := GenS3{F0: uint64(v28), F1: -2147483648}
                _ = v27
                _ = v28
                _ = v29
            }
            _ = fi26
        }
    } else {
        switch v1 {
        case 2147483647:
            v30 := GenS3{F0: 2, F1: v1}
            _ = v30
        case 42:
            if (!("foo bar" == v4)) {
                for fi31 := 0; fi31 < 6; fi31++ {
                    if false {
                        fmt.Println(v1, v4, v2)
                        v2 -= (v2 << ((807) & 63))
                    } else {
                        for fi32 := 0; fi32 < 1; fi32++ {
                            v33 := GenS3{F0: 42, F1: (fi32 & fi32)}
                            v34 := GenS4{F0: 18446744073709551615}
                            _ = fi32
                            _ = v33
                            _ = v34
                        }
                        v4 = v4
                        fmt.Println(v2)
                    }
                    _ = fi31
                }
            }
        default:
            switch "\n\t" {
            case "":
                v2 = (v3 ^ 593)
            }
        }
        var v35 int32 = 42
        fmt.Println(v2)
        _ = v35
    }
    v2 = 42
    fmt.Println(v3)
    var v36 int16 = 42
    for fi37 := 0; fi37 < 14; fi37++ {
        var v38 uint64 = 2
        for fi39 := 0; fi39 < 13; fi39++ {
            v2 = v3
            v25 = append(v25, 4294967295)
            if len(v25) > 0 {
                v25[0] = 127
            }
            _ = fi39
        }
        _ = fi37
        _ = v38
    }
    v4 = v4
    if true {
        v25 = append(v25, v1)
        v25 = append(v25, v1)
        v40 := GenS5{F0: 255}
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v36
    _ = v4
}
