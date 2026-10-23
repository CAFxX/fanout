package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 uint64; F1 int64; F2 int }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 string; F1 bool }
func genFn0(p0 int, p1 int64) int8 {
    var v6 string = ("\n\t" + "")
    if (v6 > v6) {
        fmt.Println(p1, v6)
        for fi7 := 0; fi7 < 2; fi7++ {
            switch 128 {
            case 2147483647:
                var v8 int = p0
                _ = v8
            }
            for fi9 := 0; fi9 < 17; fi9++ {
                var v10 uint64 = 0
                switch 128 {
                case -686:
                    if ((((v6 + (v6 + v6)) + v6) + v6) <= "foo bar") {
                        v6 = v6
                        v11 := []int{fi7}
                        _ = v11
                    }
                }
                v12 := GenS1{F0: 635, F1: 65535}
                v13 := []int{(fi7 * fi9)}
                _ = fi9
                _ = v10
                _ = v12
                _ = v13
            }
            _ = fi7
        }
    } else {
        v14 := GenS1{F0: 1, F1: 76}
        for fi15 := 0; fi15 < 19; fi15++ {
            v6 = ((v6 + v6) + v6)
            v16 := GenS1{F0: 255, F1: 4294967295}
            if true {
                v17 := GenS1{F0: 608, F1: 0}
                v18 := []int{p0, (fi15 << ((1) & 63)), (fi15 | (^(fi15))), 990}
                _ = v17
                _ = v18
            }
            _ = fi15
            _ = v16
        }
        _ = v14
    }
    v19 := GenS1{F0: 237, F1: 42}
    for fi20 := 0; fi20 < 11; fi20++ {
        v21 := []int{fi20, (^(p0))}
        _ = fi20
        _ = v21
    }
    _ = p0
    _ = p1
    _ = v19
    _ = v6
    return -1
}

func main() {
    var v1 uint32 = 0
    var v2 int16 = 32767
    var v3 uint32 = (^(v1))
    var v4 float32 = 1.5
    v5 := make(map[uint32]uint32)
    var v22 float32 = (v4 - (v4 + (v4 + (1.5 * v4))))
    switch "foo bar" {
    case "":
        switch 0 {
        case 42:
            var v23 int32 = -885
            _ = v23
        }
    }
    switch "a" {
    case "":
        v2 = 725
    }
    var v24 int8 = -128
    v25 := GenS2{F0: 2, F1: -619, F2: 28}
    _ = len(v5)
    var v26 int64 = 0
    if true {
        fmt.Println(v24, v22, v3)
        switch 32768 {
        case 32767:
            v27 := GenS1{F0: uint(v1), F1: 4294967295}
            _ = v27
        case 4294967295:
            switch 867 {
            case 909:
                for fi28 := 0; fi28 < 12; fi28++ {
                    var v29 int32 = 412
                    var v30 uint32 = (v1 + 2147483648)
                    _ = fi28
                    _ = v29
                    _ = v30
                }
            }
        case -89:
            switch 9223372036854775807 {
            case 1:
                var v31 int16 = v2
                _ = v31
            case 408:
                if (v4 != 0.0) {
                    for fi32 := 0; fi32 < 5; fi32++ {
                        v33 := []int{(fi32 | (fi32 % ((fi32) | 1)))}
                        v34 := GenS2{F0: 0, F1: (-(v26)), F2: 417}
                        v35 := GenS3{F0: 888}
                        _ = fi32
                        _ = v33
                        _ = v34
                        _ = v35
                    }
                    v36 := v5[1]
                    v37 := GenS1{F0: 42, F1: 18446744073709551615}
                    _ = v36
                    _ = v37
                }
            }
        default:
            if true {
                v1 -= v3
            }
        }
    } else {
        v3 += v1
        var v38 float64 = 1e10
        if (false || false) {
            v38 -= 1e10
            v5[v3] = (^(v3))
        } else {
            if ((v26 | 0) <= (v26 * (v26 % (((v26 >> ((455) & 63))) | 1)))) {
                if true {
                    v24 -= v24
                    v39 := v5[v1]
                    _ = v39
                }
            }
            if (!true) {
                if false {
                    v40 := GenS2{F0: 18446744073709551615, F1: (v26 * (v26 ^ (^(v26)))), F2: -492}
                    _ = v40
                }
                v38 = v38
            }
        }
        _ = v38
    }
    fmt.Println(v1, v2, v24)
    _ = len(v5)
    fmt.Println(v24, v2)
    for fi41 := 0; fi41 < 11; fi41++ {
        v26 += (^(v26))
        v2 += (v2 | (v2 & 0))
        for fi42 := 0; fi42 < 19; fi42++ {
            v43 := GenS4{F0: "hello", F1: false}
            fmt.Println(v3, fi42, v22)
            var v44 float32 = ((((v4 - v22) / ((v4 - v4) / v22)) * v22) / (1e10 + ((v4 - -2.25) / -2.25)))
            _ = fi42
            _ = v43
            _ = v44
        }
        _ = fi41
    }
    v45 := []int{42, 32768, -848}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v45:", len(v45), v45)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v22
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v45
    _ = v5
}
