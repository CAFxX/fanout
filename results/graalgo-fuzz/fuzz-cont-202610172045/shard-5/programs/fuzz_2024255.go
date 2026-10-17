package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 bool; F2 uint8 }
type GenS3 struct { F0 uint; F1 bool }
type GenS4 struct { F0 string; F1 uint32 }
func genFn0(p0 uint8, p1 uint16) int16 {
    var v7 int32 = 2147483647
    v8 := GenS1{F0: 1}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return -548
}
func genFn1(p0 bool, p1 uint16) uint {
    for fi9 := 0; fi9 < 17; fi9++ {
        v10 := []int{fi9, 32767, (fi9 | (fi9 | fi9))}
        if p0 {
            _ = len(v10)
            p1 += p1
            if len(v10) > 0 {
                v10[0] = 128
            }
        } else {
            for fi11 := 0; fi11 < 13; fi11++ {
                p0 = (((!(!("hello" == ""))) || false) && p0)
                switch "a" {
                case "hello":
                    if (840 != (p1 | (^(p1)))) {
                        fmt.Println(fi11)
                        switch (fi9 ^ fi11) {
                        case 2147483647:
                            v12 := GenS1{F0: 87}
                            _ = v12
                        default:
                            if p0 {
                                var v13 bool = false
                                _ = v13
                            }
                        }
                        p0 = true
                    } else {
                        if (128 < 230) {
                            if (!(!false)) {
                                v14 := GenS2{F0: -932, F1: true, F2: 2}
                                p0 = ((255 == (p1 ^ (^(p1)))) && ((!false) || (!((!false) && (65535 < 773)))))
                                var v15 float64 = 0.0
                                var v16 float32 = ((0.0 / 3.14159) + 0.0)
                                var v17 float64 = -2.25
                                fmt.Println(p0, p1, fi9)
                                if false {
                                    if true {
                                        var v18 uint8 = 0
                                        v10 = append(v10, (-(fi9)))
                                        fmt.Println(p1, fi9, p1)
                                        _ = v18
                                    } else {
                                        var v19 uint = 4294967295
                                        if len(v10) > 0 {
                                            v10[0] = (fi11 & fi9)
                                        }
                                        fmt.Println(fi9, v16, fi9)
                                        _ = v19
                                    }
                                    var v20 uint16 = p1
                                    _ = v20
                                } else {
                                    var v21 int = (fi9 << ((255) & 63))
                                    fmt.Println(p1, v16, v21)
                                    _ = v21
                                }
                                _ = v14
                                _ = v15
                                _ = v16
                                _ = v17
                            }
                        } else {
                            var v22 uint16 = (^(p1))
                            _ = v22
                        }
                    }
                case "\n\t":
                    p0 = (4294967295 <= 65535)
                case "foo bar":
                    switch "foo bar" {
                    case "a":
                        var v23 float64 = 0.0
                        _ = v23
                    case "\n\t":
                        p0 = p0
                    case "foo bar":
                        v24 := GenS2{F0: 0, F1: true, F2: 42}
                        _ = v24
                    }
                }
                _ = fi11
            }
            var v25 int8 = -128
            _ = v25
        }
        for fi26 := 0; fi26 < 2; fi26++ {
            switch (fi9 & fi26) {
            case 38:
                p1 = p1
            case 579:
                p1 += 255
            case 42:
                v27 := GenS1{F0: 127}
                var v28 int16 = 1
                var v29 float32 = 0.0
                _ = v27
                _ = v28
                _ = v29
            default:
                p0 = p0
            }
            _ = fi26
        }
        _ = fi9
        _ = v10
    }
    var v30 uint8 = 174
    v30 = 206
    v31 := []int{}
    v32 := GenS3{F0: 677, F1: true}
    _ = p0
    _ = p1
    _ = v30
    _ = v31
    _ = v32
    return 2
}

func main() {
    var v1 int8 = 127
    v2 := make(map[uint32]uint32)
    v3 := []int{42, -619}
    v4 := make(map[int]int)
    v5 := []string{"\n\t", "a", ""}
    var v6 float32 = float32(((1.5 * (3.14159 - -2.25)) / (-2.25 + float64(1.5))))
    var v33 float32 = (v6 * (v6 + 1.5))
    var v34 int32 = 0
    fmt.Println(v33)
    v34 -= v34
    _ = len(v5)
    v2[740] = 389
    v34 -= -774
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v33, v33, v1)
    if (true && (2 <= 2147483648)) {
        v1 = v1
        switch "\n\t" {
        case "foo bar":
            fmt.Println(v1, v1, v34)
        }
        if (false && (1 < 42)) {
            for fi35 := 0; fi35 < 9; fi35++ {
                for fi36 := 0; fi36 < 15; fi36++ {
                    switch 4294967295 {
                    case -9223372036854775808:
                        fmt.Println(v6, v1)
                    case 9223372036854775807:
                        if (false && (2 != 2147483648)) {
                            v6 = (v6 * (v6 - v33))
                            v33 -= (v33 - (v33 - (3.14159 * 0.1)))
                        } else {
                            v33 = 0.1
                            fmt.Println(v1)
                        }
                    }
                    _ = fi36
                }
                switch "hello" {
                case "hello":
                    if (0 >= v1) {
                        var v37 int8 = 42
                        _ = v37
                    } else {
                        fmt.Println(v6, v33)
                    }
                case "foo bar":
                    v2[42] = 513
                case "\n\t":
                    var v38 uint32 = 529
                    _ = v38
                }
                _ = fi35
            }
            v6 = (v33 + (v6 - (v33 - v6)))
        } else {
            for fi39 := 0; fi39 < 3; fi39++ {
                v40 := GenS4{F0: "", F1: 1}
                switch (^(fi39)) {
                case 4294967295:
                    switch "" {
                    case "hello":
                        v6 += float32(((-2.25 * 1e10) + (float64((v33 / v6)) * 1.5)))
                    case "foo bar":
                        v41 := GenS3{F0: 282, F1: ((!false) || false)}
                        _ = v41
                    case "a":
                        v33 += -2.25
                    }
                case -2147483648:
                    var v42 int32 = (^(v34))
                    _ = v42
                }
                _ = fi39
                _ = v40
            }
            if (-2147483648 == (v34 | v34)) {
                for fi43 := 0; fi43 < 19; fi43++ {
                    v44 := GenS2{F0: genFn0(0, 65535), F1: true, F2: 128}
                    v5 = append(v5, "foo bar")
                    _ = fi43
                    _ = v44
                }
                var v45 string = ""
                v6 += 0.1
                _ = v45
            }
            fmt.Println(v33)
        }
    }
    v33 = (v33 * -2.25)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
