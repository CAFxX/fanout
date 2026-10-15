package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 string; F1 string; F2 int8 }
type GenS3 struct { F0 uint16; F1 int }
type GenS4 struct { F0 int }
func genFn0(p0 int, p1 string, p2 int8) int16 {
    var v6 string = "\n\t"
    v7 := GenS1{F0: 18446744073709551615, F1: 297}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 uint64, p1 bool) uint8 {
    if ((!p1) || (!(4294967295 < 4294967295))) {
        v8 := GenS2{F0: "foo bar", F1: "foo bar", F2: -128}
        p1 = (!p1)
        fmt.Println(p1, p1)
        _ = v8
    } else {
        if p1 {
            if (!p1) {
                p1 = p1
                for fi9 := 0; fi9 < 1; fi9++ {
                    switch (-(fi9)) {
                    case 4294967295:
                        p1 = p1
                    }
                    var v10 string = ("\n\t" + "hello")
                    _ = fi9
                    _ = v10
                }
            }
            fmt.Println(p1, p1)
            for fi11 := 0; fi11 < 12; fi11++ {
                p0 -= (p0 / ((p0) | 1))
                _ = fi11
            }
        } else {
            if p1 {
                p1 = p1
                v12 := GenS1{F0: 156, F1: 1}
                v13 := GenS2{F0: "", F1: "", F2: 103}
                _ = v12
                _ = v13
            } else {
                fmt.Println(p0, p0)
                var v14 uint32 = 387
                if false {
                    v15 := GenS1{F0: 4294967295, F1: (p0 >> ((0) & 63))}
                    _ = v15
                }
                _ = v14
            }
        }
    }
    var v16 uint32 = 0
    v17 := []int{-62}
    v18 := GenS1{F0: uint(v16), F1: (p0 - (^(p0)))}
    _ = p0
    _ = p1
    _ = v16
    _ = v17
    _ = v18
    return 127
}

func main() {
    v1 := []uint32{2147483648}
    var v2 int64 = 1
    v3 := make(map[uint32]bool)
    var v4 string = "hello"
    var v5 uint64 = 42
    var v19 int16 = -26
    v19 = (-(v19))
    var v20 int64 = v2
    var v21 int64 = v2
    if (-1 <= 42) {
        if true {
            v3[1] = true
            var v22 int32 = -7
            for fi23 := 0; fi23 < 3; fi23++ {
                v5 = (v5 * (^(v5)))
                var v24 uint16 = 1
                v2 -= (v2 - (v21 * v21))
                _ = fi23
                _ = v24
            }
            _ = v22
        }
        v20 -= 9223372036854775807
        if (v19 > -32768) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            fmt.Println(v19)
            switch "a" {
            case "foo bar":
                v4 = ((v4 + (v4 + v4)) + v4)
            default:
                v25 := GenS2{F0: ("" + v4), F1: v4, F2: 1}
                _ = v25
            }
        }
    }
    var v26 int16 = (^(v19))
    v27 := GenS2{F0: "", F1: (v4 + (v4 + "hello")), F2: -96}
    switch "a" {
    case "a":
        for fi28 := 0; fi28 < 14; fi28++ {
            v29 := GenS3{F0: 42, F1: (fi28 << ((102) & 63))}
            var v30 float64 = 0.1
            if (2147483647 != 734) {
                var v31 uint = 1
                switch (^(fi28)) {
                case 42:
                    for fi32 := 0; fi32 < 12; fi32++ {
                        if false {
                            v30 = (1.5 + v30)
                            _ = len(v1)
                            v20 = (v21 << ((1) & 63))
                        }
                        v2 = -220
                        _ = fi32
                    }
                case 9223372036854775807:
                    for fi33 := 0; fi33 < 15; fi33++ {
                        v21 -= (v20 % ((-1) | 1))
                        _ = fi33
                    }
                }
                v34 := GenS4{F0: (fi28 + (fi28 % (((fi28 >> ((v31) & 63))) | 1)))}
                _ = v31
                _ = v34
            } else {
                for fi35 := 0; fi35 < 14; fi35++ {
                    v30 -= (3.14159 / float64(0.1))
                    v19 = 959
                    v4 = (v4 + v4)
                    _ = fi35
                }
                fmt.Println(v21, fi28, v2)
            }
            _ = fi28
            _ = v29
            _ = v30
        }
    case "hello":
        for fi36 := 0; fi36 < 9; fi36++ {
            for fi37 := 0; fi37 < 12; fi37++ {
                switch (v4 + v4) {
                case "hello":
                    var v38 int32 = 42
                    _ = v38
                case "foo bar":
                    _ = len(v1)
                default:
                    var v39 string = v4
                    _ = v39
                }
                _ = fi37
            }
            if ((float32(3.14159) - 0.0) != 0.0) {
                v40 := GenS1{F0: 65535, F1: v5}
                _ = v40
            }
            _ = fi36
        }
    }
    var v41 int32 = 2147483647
    var v42 float32 = (0.1 / 3.14159)
    for fi43 := 0; fi43 < 5; fi43++ {
        var v44 int32 = (v41 << ((0) & 63))
        for fi45 := 0; fi45 < 3; fi45++ {
            _ = len(v3)
            fmt.Println(v20, v2, fi43)
            switch "a" {
            case "hello":
                var v46 int32 = v41
                _ = v46
            }
            _ = fi45
        }
        _ = fi43
        _ = v44
    }
    v2 -= (v21 | (^(v2)))
    for fi47 := 0; fi47 < 14; fi47++ {
        v26 = 1
        for fi48 := 0; fi48 < 3; fi48++ {
            fmt.Println(v20, fi48, v4)
            _ = fi48
        }
        _ = fi47
    }
    if ((0.0 * (0.0 - 1e10)) >= (0.1 / (1.5 * 1.5))) {
        var v49 int8 = -128
        v50 := v3[2]
        _ = v49
        _ = v50
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
}
