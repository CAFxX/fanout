package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32; F1 uint8; F2 uint }
func genFn0(p0 uint64) int {
    var v7 uint32 = uint32(p0)
    var v8 uint16 = 93
    for fi9 := 0; fi9 < 4; fi9++ {
        fmt.Println(p0, v8, v7)
        var v10 int = fi9
        _ = fi9
        _ = v10
    }
    var v11 string = ("a" + "\n\t")
    _ = p0
    _ = v11
    _ = v7
    _ = v8
    return 995
}
func genFn1(p0 uint8, p1 uint32) int {
    v12 := []int{-85}
    p1 = p1
    for fi13 := 0; fi13 < 15; fi13++ {
        for fi14 := 0; fi14 < 11; fi14++ {
            p0 += p0
            _ = fi14
        }
        switch "" {
        case "foo bar":
            for fi15 := 0; fi15 < 9; fi15++ {
                v16 := GenS1{F0: 4294967295}
                var v17 int = 1
                _ = fi15
                _ = v16
                _ = v17
            }
        case "":
            _ = len(v12)
        default:
            _ = len(v12)
        }
        v12 = append(v12, fi13)
        _ = fi13
    }
    if (true && true) {
        fmt.Println(p1, p0)
        p0 -= p0
    }
    _ = p0
    _ = p1
    _ = v12
    return -1
}

func main() {
    var v1 float64 = 1.5
    v2 := make(map[uint32]uint32)
    var v3 bool = true
    var v4 uint16 = 1
    var v5 uint16 = (v4 ^ v4)
    var v6 float64 = float64(1e10)
    v18 := []int{}
    if false {
        switch "" {
        case "hello":
            if v3 {
                v5 = (^(v5))
            } else {
                switch 32768 {
                case 127:
                    var v19 string = ("a" + "")
                    _ = v19
                case -9223372036854775808:
                    var v20 int = 819
                    _ = v20
                case -1:
                    v21 := GenS1{F0: 4294967295}
                    _ = v21
                }
                var v22 int64 = -406
                _ = v22
            }
        }
    } else {
        for fi23 := 0; fi23 < 19; fi23++ {
            var v24 uint = 18446744073709551615
            switch "hello" {
            case "a":
                fmt.Println(fi23)
            case "":
                var v25 uint64 = 178
                _ = v25
            case "\n\t":
                var v26 float32 = -2.25
                _ = v26
            }
            v5 -= 2
            _ = fi23
            _ = v24
        }
        switch 1 {
        case -2147483648:
            v27 := GenS1{F0: 817}
            _ = v27
        case 127:
            if false {
                v18 = append(v18, 32767)
            } else {
                v28 := GenS1{F0: 0}
                for fi29 := 0; fi29 < 9; fi29++ {
                    v5 += 441
                    _ = fi29
                }
                _ = v28
            }
        }
    }
    fmt.Println(v1, v1)
    v30 := v2[4294967295]
    var v31 int16 = 1000
    v30 += 850
    var v32 string = ("hello" + "foo bar")
    for fi33 := 0; fi33 < 13; fi33++ {
        for fi34 := 0; fi34 < 20; fi34++ {
            var v35 string = (v32 + "\n\t")
            v36 := GenS1{F0: 4294967295}
            _ = fi34
            _ = v35
            _ = v36
        }
        _ = fi33
    }
    if (9223372036854775807 > -9223372036854775808) {
        switch v32 {
        case "a":
            switch "a" {
            case "\n\t":
                for fi37 := 0; fi37 < 15; fi37++ {
                    v38 := GenS2{F0: v30, F1: 127, F2: 33}
                    fmt.Println(v6, v4, v31)
                    if v3 {
                        fmt.Println(v30)
                    }
                    _ = fi37
                    _ = v38
                }
            case "foo bar":
                _ = len(v2)
            case "hello":
                if v3 {
                    var v39 int32 = 1
                    v2[v30] = (^(v30))
                    v4 -= (v5 / ((v4) | 1))
                    _ = v39
                } else {
                    for fi40 := 0; fi40 < 16; fi40++ {
                        var v41 uint64 = 18446744073709551615
                        _ = fi40
                        _ = v41
                    }
                }
            default:
                var v42 int = -981
                _ = v42
            }
        case "\n\t":
            switch 386 {
            case 32768:
                for fi43 := 0; fi43 < 10; fi43++ {
                    var v44 int16 = (v31 / (((-(v31))) | 1))
                    var v45 float64 = 0.0
                    v46 := GenS1{F0: 4294967295}
                    _ = fi43
                    _ = v44
                    _ = v45
                    _ = v46
                }
            default:
                if len(v18) > 0 {
                    v18[0] = genFn0(uint64(v30))
                }
            }
        }
        if false {
            fmt.Println(v4, v30, v5)
        }
        if len(v18) > 0 {
            v18[0] = 2147483647
        }
    }
    if (!((v3 && ((127 > 95) || (977 == -1))) || (166 == uint8(v5)))) {
        delete(v2, 2147483647)
        for fi47 := 0; fi47 < 8; fi47++ {
            var v48 float64 = (3.14159 * (((-2.25 - (v1 / 0.1)) * float64(-2.25)) - 0.0))
            _ = fi47
            _ = v48
        }
    } else {
        var v49 uint8 = 127
        _ = v49
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
