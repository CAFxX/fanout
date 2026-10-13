package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 int32 }
type GenS2 struct { F0 uint64; F1 int8; F2 uint64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint32; F1 string; F2 uint16 }
type GenS5 struct { F0 int8; F1 int8 }
type GenS6 struct { F0 uint16; F1 uint32; F2 uint32 }
func genFn0(p0 int, p1 int16) uint {
    v4 := []int{(^(p0)), p0, 128}
    var v5 bool = false
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    _ = len(v4)
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 1
}
func genFn1(p0 int64, p1 int32) uint {
    fmt.Println(p1, p1)
    var v6 string = ("\n\t" + "foo bar")
    if true {
        p0 -= p0
        v7 := GenS1{F0: -202, F1: -123, F2: (p1 - p1)}
        v8 := GenS1{F0: 42, F1: 114, F2: (p1 & p1)}
        _ = v7
        _ = v8
    } else {
        for fi9 := 0; fi9 < 5; fi9++ {
            switch v6 {
            case "":
                if true {
                    fmt.Println(p0)
                    v6 = ((v6 + ("hello" + (v6 + v6))) + v6)
                    var v10 int64 = (^(p0))
                    _ = v10
                }
            case "\n\t":
                for fi11 := 0; fi11 < 14; fi11++ {
                    v12 := []int{fi9, (^(fi11)), (-(fi9))}
                    _ = len(v12)
                    var v13 uint = genFn0(-38, -729)
                    v14 := GenS2{F0: 696, F1: 100, F2: 18446744073709551615}
                    switch "a" {
                    case "hello":
                        switch fi9 {
                        case 42:
                            p1 -= 628
                        case -337:
                            v15 := GenS1{F0: 0, F1: 1, F2: p1}
                            _ = v15
                        case 207:
                            var v16 uint = 236
                            _ = v16
                        default:
                            var v17 int64 = -535
                            _ = v17
                        }
                    case "\n\t":
                        var v18 uint32 = 42
                        _ = v18
                    case "foo bar":
                        if (128 < 128) {
                            switch (v6 + "hello") {
                            case "\n\t":
                                var v19 uint32 = 2
                                _ = v19
                            case "hello":
                                var v20 float64 = (0.0 * 3.14159)
                                _ = v20
                            default:
                                v21 := GenS2{F0: 42, F1: -128, F2: 42}
                                _ = v21
                            }
                            v6 = v6
                            v12 = append(v12, (-(fi9)))
                        }
                    default:
                        v12 = append(v12, fi11)
                    }
                    _ = fi11
                    _ = v12
                    _ = v13
                    _ = v14
                }
            default:
                v22 := GenS1{F0: 0, F1: -57, F2: p1}
                _ = v22
            }
            v23 := []int{}
            _ = fi9
            _ = v23
        }
    }
    _ = p0
    _ = p1
    _ = v6
    return 2
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int8 = -1
    var v3 int16 = 1
    v24 := GenS1{F0: v3, F1: (^(v2)), F2: 1}
    var v25 int = 4294967295
    v26 := GenS3{F0: false}
    if true {
        var v27 float32 = 1.5
        v28 := GenS1{F0: 0, F1: v2, F2: -441}
        _ = v27
        _ = v28
    }
    v29 := GenS4{F0: 42, F1: "a", F2: 2}
    v30 := []int{(v25 * v25), v25, (v25 / ((v25) | 1)), v25}
    for fi31 := 0; fi31 < 11; fi31++ {
        var v32 float32 = (0.0 + 1.5)
        v33 := GenS5{F0: 41, F1: (-(v2))}
        _ = fi31
        _ = v32
        _ = v33
    }
    if (false || false) {
        switch 0 {
        case -9223372036854775808:
            switch ("" + "\n\t") {
            case "foo bar":
                for fi34 := 0; fi34 < 7; fi34++ {
                    v2 = v2
                    var v35 int = v25
                    var v36 uint16 = 255
                    _ = fi34
                    _ = v35
                    _ = v36
                }
            case "":
                var v37 int16 = (-(v3))
                _ = v37
            }
        case -688:
            var v38 uint64 = 722
            _ = v38
        case 2147483647:
            if false {
                switch int(v3) {
                case 0:
                    v25 = v25
                case 4294967295:
                    v39 := GenS4{F0: 62, F1: "a", F2: 716}
                    _ = v39
                case -751:
                    _ = len(v30)
                default:
                    var v40 uint32 = 2
                    _ = v40
                }
            } else {
                _ = len(v30)
            }
        default:
            var v41 float64 = 0.0
            _ = v41
        }
    }
    v25 = (v25 & (^(v25)))
    switch "\n\t" {
    case "foo bar":
        switch "\n\t" {
        case "hello":
            if len(v30) > 0 {
                v30[0] = (v25 & (-(v25)))
            }
        case "foo bar":
            for fi42 := 0; fi42 < 4; fi42++ {
                v2 -= (v2 >> ((4294967295) & 63))
                fmt.Println(v2)
                v25 -= (fi42 + fi42)
                _ = fi42
            }
        case "":
            var v43 int = 128
            _ = v43
        }
    case "\n\t":
        for fi44 := 0; fi44 < 7; fi44++ {
            fmt.Println(v2, v3, v2)
            v45 := GenS6{F0: 898, F1: 0, F2: 2}
            _ = fi44
            _ = v45
        }
    default:
        v3 = v3
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v30
}
