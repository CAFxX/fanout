package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64; F1 uint32; F2 int16 }
func genFn0(p0 uint8, p1 bool) int64 {
    for fi7 := 0; fi7 < 6; fi7++ {
        for fi8 := 0; fi8 < 6; fi8++ {
            v9 := []int{-1, fi7}
            fmt.Println(fi8, fi7, p0)
            v9 = append(v9, (-(fi7)))
            _ = fi8
            _ = v9
        }
        if p1 {
            for fi10 := 0; fi10 < 5; fi10++ {
                var v11 int8 = -20
                p1 = (p1 && false)
                _ = fi10
                _ = v11
            }
            switch (^(fi7)) {
            case -9223372036854775808:
                var v12 uint16 = 171
                _ = v12
            case -704:
                var v13 string = (("foo bar" + ("\n\t" + "\n\t")) + "a")
                _ = v13
            case -478:
                fmt.Println(p0)
            }
            var v14 float32 = ((1e10 * 0.0) * 0.1)
            _ = v14
        }
        var v15 int64 = -397
        _ = fi7
        _ = v15
    }
    v16 := []int{128, 0, 2147483647}
    _ = p0
    _ = p1
    _ = v16
    return 0
}
func genFn1(p0 int16, p1 uint16) int32 {
    if true {
        fmt.Println(p1, p1)
    }
    v17 := []int{}
    v18 := GenS1{F0: 2147483648, F1: true, F2: 170}
    _ = len(v17)
    var v19 float64 = -2.25
    _ = p0
    _ = p1
    _ = v17
    _ = v18
    _ = v19
    return 0
}

func main() {
    var v1 uint64 = 233
    var v2 uint = 0
    v3 := []uint32{}
    v4 := make(map[string]bool)
    var v5 int8 = -105
    v6 := make(map[int]bool)
    if ((351 == uint16(v2)) && (42 != v5)) {
        v20 := GenS1{F0: 65535, F1: (-549 != 128), F2: 0}
        fmt.Println(v1)
        _ = v20
    }
    fmt.Println(v1, v5)
    switch ("foo bar" + "") {
    case "foo bar":
        v21 := GenS1{F0: 279, F1: true, F2: 255}
        _ = v21
    case "hello":
        v22 := GenS1{F0: 834, F1: ((1 != 292) && false), F2: 0}
        _ = v22
    }
    for fi23 := 0; fi23 < 20; fi23++ {
        v5 = v5
        var v24 float64 = 0.1
        _ = fi23
        _ = v24
    }
    var v25 uint32 = 0
    for fi26 := 0; fi26 < 12; fi26++ {
        v5 = (v5 >> (((v2 % ((v2) | 1))) & 63))
        _ = fi26
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v25, v25)
    switch "a" {
    case "foo bar":
        v1 = (v1 >> (((v2 & v2)) & 63))
    case "":
        v27 := GenS1{F0: (v25 % ((v25) | 1)), F1: (!((v2 << ((v2) & 63)) >= (^(v2)))), F2: 128}
        _ = v27
    default:
        if (true && ((^(v1)) <= 2)) {
            for fi28 := 0; fi28 < 16; fi28++ {
                var v29 uint = v2
                var v30 uint64 = v1
                _ = fi28
                _ = v29
                _ = v30
            }
            switch "hello" {
            case "foo bar":
                v25 += (v25 & 691)
            case "a":
                if (1 == (v1 * 1)) {
                    for fi31 := 0; fi31 < 4; fi31++ {
                        var v32 int16 = 1000
                        _ = fi31
                        _ = v32
                    }
                    if (false || (true && true)) {
                        if len(v3) > 0 {
                            v3[0] = uint32(v2)
                        }
                    } else {
                        var v33 uint = (v2 / ((v2) | 1))
                        _ = v33
                    }
                } else {
                    if (!true) {
                        for fi34 := 0; fi34 < 4; fi34++ {
                            _ = len(v6)
                            var v35 uint64 = v1
                            fmt.Println(v35)
                            _ = fi34
                            _ = v35
                        }
                        var v36 int32 = 1
                        v5 = (v5 % (((v5 / ((v5) | 1))) | 1))
                        _ = v36
                    }
                    var v37 uint16 = 2
                    _ = v37
                }
            case "":
                for fi38 := 0; fi38 < 17; fi38++ {
                    v5 = v5
                    fmt.Println(v2, fi38, v25)
                    _ = fi38
                }
            default:
                fmt.Println(v25)
            }
            switch 0 {
            case 127:
                for fi39 := 0; fi39 < 2; fi39++ {
                    v2 += v2
                    var v40 int8 = (v5 % ((v5) | 1))
                    _ = fi39
                    _ = v40
                }
            case 4294967295:
                for fi41 := 0; fi41 < 3; fi41++ {
                    v25 -= (^(v25))
                    v42 := GenS2{F0: genFn1(32767, 927)}
                    v5 += v5
                    _ = fi41
                    _ = v42
                }
            }
        } else {
            var v43 bool = false
            _ = v43
        }
    }
    var v44 uint = (v2 + (v2 % (((v2 % ((487) | 1))) | 1)))
    v25 = (v25 >> ((1) & 63))
    if ((!(255 == 2)) && ((^(v25)) <= v25)) {
        v25 = v25
        for fi45 := 0; fi45 < 14; fi45++ {
            if (v5 != v5) {
                v3 = append(v3, v25)
                v2 = uint(v25)
            } else {
                delete(v6, 0)
                var v46 int32 = 42
                _ = v46
            }
            for fi47 := 0; fi47 < 13; fi47++ {
                v44 = v2
                _ = fi47
            }
            _ = fi45
        }
    } else {
        v44 += v44
        fmt.Println(v2)
        var v48 uint8 = 10
        _ = v48
    }
    v44 += (v2 / ((0) | 1))
    fmt.Println(v2, v44, v44)
    v49 := GenS3{F0: 9223372036854775807, F1: (^(v25)), F2: 1000}
    v5 = (v5 + (v5 << ((86) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v44
    _ = v49
    _ = v5
    _ = v6
}
