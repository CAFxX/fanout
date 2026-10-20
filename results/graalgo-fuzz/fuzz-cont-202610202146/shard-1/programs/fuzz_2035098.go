package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int64; F1 uint64; F2 int16 }
func genFn0(p0 int64, p1 int32) uint64 {
    if (!false) {
        v7 := GenS1{F0: 18446744073709551615}
        _ = v7
    } else {
        v8 := GenS1{F0: 291}
        _ = v8
    }
    switch -2147483648 {
    case 32768:
        var v9 float32 = 0.1
        _ = v9
    case -2147483648:
        v10 := GenS1{F0: 454}
        _ = v10
    case -9223372036854775808:
        for fi11 := 0; fi11 < 6; fi11++ {
            if (!(p1 >= (p1 % (((p1 | (p1 + int32(fi11)))) | 1)))) {
                v12 := GenS1{F0: 4294967295}
                _ = v12
            }
            _ = fi11
        }
    }
    fmt.Println(p1)
    v13 := []int{-2147483648, int(p0), 127, 9223372036854775807}
    _ = p0
    _ = p1
    _ = v13
    return 170
}
func genFn1(p0 int32, p1 uint8) int8 {
    v14 := GenS1{F0: 2}
    for fi15 := 0; fi15 < 5; fi15++ {
        p1 = p1
        if (!true) {
            var v16 uint64 = 1
            p1 -= (p1 / ((p1) | 1))
            v17 := []int{}
            _ = v16
            _ = v17
        } else {
            var v18 uint = 18446744073709551615
            fmt.Println(p0)
            _ = v18
        }
        var v19 string = "hello"
        _ = fi15
        _ = v19
    }
    if (!(1 < 400)) {
        v20 := GenS2{F0: 151}
        p0 += 0
        if (!true) {
            fmt.Println(p1, p0)
        } else {
            v21 := []int{0}
            if true {
                v22 := GenS2{F0: 272}
                var v23 uint8 = (p1 & 255)
                var v24 uint32 = 2147483648
                _ = v22
                _ = v23
                _ = v24
            } else {
                fmt.Println(p1, p0)
            }
            switch "foo bar" {
            case "a":
                var v25 int32 = p0
                _ = v25
            default:
                switch "\n\t" {
                case "foo bar":
                    for fi26 := 0; fi26 < 3; fi26++ {
                        p0 += (p0 & (^(p0)))
                        var v27 uint8 = p1
                        var v28 int16 = 1
                        _ = fi26
                        _ = v27
                        _ = v28
                    }
                default:
                    if true {
                        p0 -= p0
                        var v29 int = 0
                        _ = v29
                    }
                }
            }
            _ = v21
        }
        _ = v20
    } else {
        var v30 string = "hello"
        v31 := []int{}
        switch v30 {
        case "\n\t":
            var v32 float64 = -2.25
            _ = v32
        case "a":
            switch 9223372036854775807 {
            case 127:
                var v33 float32 = (3.14159 + 1e10)
                _ = v33
            case 1:
                if true {
                    v34 := GenS3{F0: -766}
                    _ = v34
                }
            case -348:
                p1 += p1
            }
        }
        _ = v30
        _ = v31
    }
    _ = p0
    _ = p1
    _ = v14
    return 0
}

func main() {
    var v1 uint64 = 227
    var v2 int32 = 2147483647
    v3 := make(map[int]bool)
    v4 := make(map[string]bool)
    v5 := []uint32{2147483647, 0}
    var v6 int64 = -538
    for fi35 := 0; fi35 < 3; fi35++ {
        v36 := GenS1{F0: 967}
        var v37 bool = ((-(fi35)) > -263)
        v2 = 42
        _ = fi35
        _ = v36
        _ = v37
    }
    fmt.Println(v1, v2)
    for fi38 := 0; fi38 < 13; fi38++ {
        v2 = -532
        _ = fi38
    }
    v1 = v1
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v6 = (-(v6))
    fmt.Println(v1, v6)
    if ((!true) && true) {
        var v39 int8 = 0
        switch 42 {
        case 0:
            for fi40 := 0; fi40 < 10; fi40++ {
                if len(v5) > 0 {
                    v5[0] = 0
                }
                _ = fi40
            }
        case 2147483647:
            var v41 int16 = -484
            _ = v41
        default:
            v6 -= v6
        }
        _ = v39
    }
    for fi42 := 0; fi42 < 10; fi42++ {
        if ("" != "foo bar") {
            var v43 int = 130
            for fi44 := 0; fi44 < 8; fi44++ {
                switch fi42 {
                case 42:
                    if len(v5) > 0 {
                        v5[0] = 2147483647
                    }
                case 567:
                    switch v43 {
                    case 2147483647:
                        switch ("a" + "a") {
                        case "foo bar":
                            v1 = v1
                        case "\n\t":
                            var v45 bool = true
                            _ = v45
                        }
                    case 573:
                        v46 := v3[(v43 - (fi44 << ((2) & 63)))]
                        _ = v46
                    default:
                        fmt.Println(fi44, fi42)
                    }
                case 127:
                    v47 := GenS4{F0: v6, F1: (^(v1)), F2: -1}
                    _ = v47
                default:
                    if (2 > 2) {
                        var v48 uint16 = 1
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        _ = v48
                    } else {
                        v1 = (^(v1))
                    }
                }
                _ = fi44
            }
            _ = v43
        }
        _ = fi42
    }
    var v49 int8 = 100
    var v50 int16 = 0
    v49 += (v49 | v49)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v49:", v49)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v6
}
