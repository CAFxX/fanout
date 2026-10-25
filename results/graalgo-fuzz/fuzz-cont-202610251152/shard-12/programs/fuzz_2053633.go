package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16; F1 int8 }
func genFn0(p0 uint16, p1 uint) int8 {
    if false {
        fmt.Println(p0, p1)
        v4 := GenS1{F0: 32768}
        v5 := GenS1{F0: -691}
        _ = v4
        _ = v5
    } else {
        switch "a" {
        case "\n\t":
            if false {
                var v6 int32 = -689
                fmt.Println(p1, p0, v6)
                _ = v6
            }
        case "a":
            fmt.Println(p1)
        case "foo bar":
            fmt.Println(p1, p1)
        default:
            p1 -= 1
        }
    }
    p0 = (p0 + (p0 | uint16(p1)))
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    return 1
}
func genFn1(p0 uint8) uint {
    v7 := []int{118, 127, -9223372036854775808}
    for fi8 := 0; fi8 < 1; fi8++ {
        for fi9 := 0; fi9 < 7; fi9++ {
            if (103 <= 19) {
                var v10 int32 = 42
                p0 -= p0
                _ = v10
            }
            var v11 int = fi8
            _ = fi9
            _ = v11
        }
        for fi12 := 0; fi12 < 19; fi12++ {
            v13 := GenS1{F0: fi8}
            if false {
                var v14 uint32 = 2147483648
                _ = v14
            }
            _ = fi12
            _ = v13
        }
        p0 = (p0 << ((42) & 63))
        _ = fi8
    }
    fmt.Println(p0)
    if false {
        for fi15 := 0; fi15 < 12; fi15++ {
            v16 := GenS2{F0: 7}
            for fi17 := 0; fi17 < 20; fi17++ {
                if true {
                    var v18 int16 = 0
                    var v19 uint16 = 194
                    _ = v18
                    _ = v19
                }
                _ = fi17
            }
            if false {
                p0 = 1
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                if len(v7) > 0 {
                    v7[0] = fi15
                }
            } else {
                if false {
                    v7 = append(v7, (^(fi15)))
                    _ = len(v7)
                    var v20 int16 = -32768
                    _ = v20
                } else {
                    p0 = (^(p0))
                }
            }
            _ = fi15
            _ = v16
        }
        var v21 uint8 = p0
        var v22 float64 = 3.14159
        _ = v21
        _ = v22
    }
    for fi23 := 0; fi23 < 10; fi23++ {
        if true {
            v24 := GenS1{F0: fi23}
            _ = v24
        } else {
            p0 += p0
            fmt.Println(p0, p0)
            fmt.Println(p0, p0)
        }
        var v25 string = ("\n\t" + "")
        _ = fi23
        _ = v25
    }
    _ = p0
    _ = v7
    return 65535
}

func main() {
    var v1 bool = true
    var v2 float64 = ((1.5 + 3.14159) * 1.5)
    var v3 uint8 = 0
    v26 := []int{}
    var v27 int64 = -866
    v28 := GenS2{F0: 828}
    v29 := GenS2{F0: 355}
    var v30 uint64 = 0
    for fi31 := 0; fi31 < 14; fi31++ {
        v26 = append(v26, (fi31 << ((0) & 63)))
        for fi32 := 0; fi32 < 1; fi32++ {
            var v33 uint8 = v3
            _ = len(v26)
            _ = fi32
            _ = v33
        }
        var v34 uint16 = 0
        _ = fi31
        _ = v34
    }
    _ = len(v26)
    switch ("hello" + "foo bar") {
    case "":
        fmt.Println(v27, v1, v3)
    case "a":
        fmt.Println(v3)
    }
    v35 := GenS2{F0: 0}
    v3 = (v3 ^ v3)
    for fi36 := 0; fi36 < 1; fi36++ {
        if (true || v1) {
            v1 = v1
            fmt.Println(v30, fi36)
        }
        var v37 uint32 = uint32(v30)
        _ = fi36
        _ = v37
    }
    for fi38 := 0; fi38 < 4; fi38++ {
        for fi39 := 0; fi39 < 8; fi39++ {
            if v1 {
                v40 := GenS1{F0: 128}
                if true {
                    v2 += (((((((v2 - (float64(1e10) * 1e10)) + (v2 / v2)) * v2) - v2) * v2) / -2.25) / v2)
                    v41 := GenS2{F0: 63}
                    var v42 float64 = ((v2 + (3.14159 + -2.25)) * -2.25)
                    _ = v41
                    _ = v42
                } else {
                    v2 = v2
                    v30 = v30
                }
                v30 = 4294967295
                _ = v40
            }
            v43 := GenS3{F0: 2, F1: int8(v27)}
            var v44 uint32 = 2147483647
            var v45 uint8 = (v3 + v3)
            if len(v26) >= 2 {
                v26 = v26[1:len(v26)]
            }
            switch fi39 {
            case 870:
                switch fi38 {
                case -2147483648:
                    fmt.Println(v3)
                case 42:
                    fmt.Println(v1)
                default:
                    v45 = (v45 + v3)
                }
            }
            v46 := GenS2{F0: 1}
            _ = fi39
            _ = v43
            _ = v44
            _ = v45
            _ = v46
        }
        for fi47 := 0; fi47 < 14; fi47++ {
            if ((!false) || false) {
                v27 -= (v27 & 612)
            }
            var v48 int32 = 2147483647
            v49 := GenS2{F0: 255}
            v3 -= 1
            v50 := GenS1{F0: 2147483647}
            var v51 uint64 = uint64(v3)
            _ = fi47
            _ = v48
            _ = v49
            _ = v50
            _ = v51
        }
        v52 := GenS3{F0: 65535, F1: 0}
        _ = fi38
        _ = v52
    }
    var v53 uint32 = 2
    v3 = v3
    if len(v26) >= 2 {
        v26 = v26[1:len(v26)]
    }
    fmt.Println(v53, v27)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v35
    _ = v53
}
