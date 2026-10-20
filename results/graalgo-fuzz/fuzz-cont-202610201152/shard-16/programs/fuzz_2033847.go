package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int16; F1 uint8; F2 int16 }
type GenS6 struct { F0 uint32; F1 uint64 }
type GenS7 struct { F0 uint; F1 int16 }
type GenS8 struct { F0 string }
func genFn0(p0 uint32) uint32 {
    p0 = p0
    var v5 uint32 = (p0 & 977)
    _ = p0
    _ = v5
    return v5
}

func main() {
    var v1 uint64 = 1
    var v2 int8 = -128
    v3 := []int{-1}
    var v4 uint32 = 4294967295
    var v6 int16 = -1
    fmt.Println(v4)
    fmt.Println(v6)
    v6 = (-(v6))
    var v7 int64 = 992
    var v8 uint16 = uint16(v4)
    switch 9223372036854775807 {
    case 2147483647:
        if false {
            fmt.Println(v2, v6)
            v8 = 1
            switch "hello" {
            case "a":
                var v9 bool = (2 == 65535)
                _ = v9
            }
        }
    case -541:
        if (false || true) {
            var v10 int64 = v7
            _ = v10
        }
    case 452:
        if true {
            for fi11 := 0; fi11 < 8; fi11++ {
                if true {
                    if true {
                        switch (fi11 | fi11) {
                        case 4294967295:
                            if true {
                                var v12 bool = (-106 < 0)
                                _ = v12
                            } else {
                                switch "\n\t" {
                                case "hello":
                                    _ = len(v3)
                                }
                                if (true || (("\n\t" + "hello") > ("a" + ("foo bar" + "foo bar")))) {
                                    var v13 int64 = v7
                                    var v14 string = "hello"
                                    v3 = append(v3, (fi11 & 0))
                                    _ = v13
                                    _ = v14
                                } else {
                                    var v15 int16 = (v6 >> ((2) & 63))
                                    _ = v15
                                }
                            }
                        case 1:
                            switch "foo bar" {
                            case "foo bar":
                                v16 := GenS1{F0: (2 <= 1)}
                                _ = v16
                            case "hello":
                                switch (("hello" + "a") + ("foo bar" + ("" + ("hello" + "\n\t")))) {
                                case "hello":
                                    v17 := GenS2{F0: 42, F1: (v4 - (v4 >> ((545) & 63)))}
                                    _ = v17
                                case "":
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                case "a":
                                    if (true || (false && ((1e10 == -2.25) && ((float64((1.5 + (float32(1.5) / 1e10))) / (0.1 + 1e10)) <= (0.0 * ((0.1 * float64(-2.25)) + -2.25)))))) {
                                        v3 = append(v3, fi11)
                                        switch "foo bar" {
                                        case "foo bar":
                                            var v18 bool = true
                                            _ = v18
                                        case "\n\t":
                                            for fi19 := 0; fi19 < 5; fi19++ {
                                                switch (("a" + "foo bar") + ((("foo bar" + "foo bar") + "") + (("a" + "a") + "\n\t"))) {
                                                case "":
                                                    var v20 int8 = (v2 ^ int8(fi19))
                                                    _ = v20
                                                case "\n\t":
                                                    if (42 <= (^(v4))) {
                                                        _ = len(v3)
                                                        var v21 bool = true
                                                        switch fi19 {
                                                        case 32768:
                                                            var v22 uint64 = (^(v1))
                                                            _ = v22
                                                        case 1:
                                                            switch fi19 {
                                                            case 756:
                                                                var v23 uint64 = 1
                                                                if v21 {
                                                                    _ = len(v3)
                                                                    switch (fi19 * fi19) {
                                                                    case -592:
                                                                        var v24 int64 = v7
                                                                        _ = v24
                                                                    }
                                                                    v25 := GenS3{F0: 883}
                                                                    fmt.Println(v6, v8, fi11)
                                                                    _ = v25
                                                                }
                                                                _ = v23
                                                            case -977:
                                                                v7 -= int64(fi19)
                                                            case 569:
                                                                fmt.Println(v7, v21)
                                                            }
                                                        default:
                                                            v2 -= 1
                                                        }
                                                        var v26 uint = 18446744073709551615
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        v3 = append(v3, fi19)
                                                        _ = v21
                                                        _ = v26
                                                    } else {
                                                        v6 -= (v6 ^ (v6 * v6))
                                                        v27 := GenS4{F0: 255}
                                                        if len(v3) > 0 {
                                                            v3[0] = -2147483648
                                                        }
                                                        _ = v27
                                                    }
                                                case "a":
                                                    v28 := GenS5{F0: v6, F1: 127, F2: -804}
                                                    _ = v28
                                                }
                                                if false {
                                                    v1 = v1
                                                } else {
                                                    v29 := GenS6{F0: v4, F1: (v1 & v1)}
                                                    var v30 int = (fi19 ^ -566)
                                                    switch (fi11 % ((fi11) | 1)) {
                                                    case -729:
                                                        var v31 float64 = 0.0
                                                        _ = v31
                                                    case 0:
                                                        v2 -= int8(v30)
                                                    case -2147483648:
                                                        v3 = append(v3, v30)
                                                    default:
                                                        fmt.Println(v7)
                                                    }
                                                    _ = v29
                                                    _ = v30
                                                }
                                                if len(v3) > 0 {
                                                    v3[0] = (-(fi19))
                                                }
                                                _ = fi19
                                            }
                                        default:
                                            var v32 string = ("a" + "")
                                            _ = v32
                                        }
                                    } else {
                                        v8 += v8
                                        var v33 bool = (78 != uint8(v1))
                                        _ = v33
                                    }
                                }
                            case "":
                                var v34 int8 = (v2 & (^(v2)))
                                _ = v34
                            default:
                                v35 := GenS4{F0: 219}
                                _ = v35
                            }
                        case -805:
                            v2 -= (v2 - v2)
                        }
                        for fi36 := 0; fi36 < 11; fi36++ {
                            var v37 float32 = 1.5
                            fmt.Println(v8)
                            _ = fi36
                            _ = v37
                        }
                        v38 := GenS1{F0: (1 == -32768)}
                        _ = v38
                    } else {
                        for fi39 := 0; fi39 < 20; fi39++ {
                            if len(v3) > 0 {
                                v3[0] = fi39
                            }
                            switch (("foo bar" + "foo bar") + "hello") {
                            case "a":
                                switch "a" {
                                case "hello":
                                    fmt.Println(v8, v8, fi11)
                                default:
                                    if false {
                                        v40 := GenS3{F0: (fi39 / ((959) | 1))}
                                        switch "a" {
                                        case "hello":
                                            v2 += 1
                                        case "a":
                                            v3 = append(v3, -2147483648)
                                        }
                                        var v41 uint8 = 42
                                        _ = v40
                                        _ = v41
                                    }
                                }
                            case "foo bar":
                                var v42 uint8 = 0
                                _ = v42
                            case "":
                                fmt.Println(v7, fi11)
                            default:
                                v3 = append(v3, (fi39 - (fi11 * int(v6))))
                                switch ("" + "a") {
                                case "a":
                                    if true {
                                        if (false && (9223372036854775807 >= v7)) {
                                            if ((v7 & (-(v7))) >= v7) {
                                                switch "\n\t" {
                                                case "":
                                                    fmt.Println(fi11)
                                                }
                                                switch (fi11 & -1) {
                                                case 370:
                                                    v43 := GenS4{F0: 255}
                                                    _ = v43
                                                case -1:
                                                    switch fi39 {
                                                    case 0:
                                                        v44 := GenS7{F0: 0, F1: v6}
                                                        _ = v44
                                                    }
                                                default:
                                                    fmt.Println(fi39, v6)
                                                }
                                            } else {
                                                var v45 uint = 12
                                                v46 := GenS2{F0: -881, F1: (v4 >> ((11) & 63))}
                                                _ = v45
                                                _ = v46
                                            }
                                            var v47 uint32 = v4
                                            _ = v47
                                        } else {
                                            var v48 bool = (v7 != v7)
                                            var v49 uint32 = v4
                                            _ = v48
                                            _ = v49
                                        }
                                        v50 := GenS3{F0: fi39}
                                        if false {
                                            var v51 int = fi11
                                            var v52 uint8 = 42
                                            v52 += 1
                                            _ = v51
                                            _ = v52
                                        }
                                        _ = v50
                                    }
                                case "foo bar":
                                    _ = len(v3)
                                default:
                                    v8 = 215
                                }
                                if len(v3) > 0 {
                                    v3[0] = (fi39 / (((fi39 >> ((18446744073709551615) & 63))) | 1))
                                }
                            }
                            v53 := GenS3{F0: (fi11 << ((17) & 63))}
                            _ = fi39
                            _ = v53
                        }
                    }
                    for fi54 := 0; fi54 < 1; fi54++ {
                        _ = len(v3)
                        v7 = v7
                        var v55 uint = 2
                        _ = fi54
                        _ = v55
                    }
                }
                _ = fi11
            }
        }
    }
    for fi56 := 0; fi56 < 5; fi56++ {
        var v57 bool = true
        var v58 uint = 18446744073709551615
        var v59 int64 = (v7 << ((v58) & 63))
        _ = fi56
        _ = v57
        _ = v58
        _ = v59
    }
    if true {
        fmt.Println(v6, v1)
    } else {
        v4 = v4
        fmt.Println(v7)
        var v60 uint = 255
        _ = v60
    }
    for fi61 := 0; fi61 < 9; fi61++ {
        v6 = (v6 >> ((uint(v8)) & 63))
        _ = fi61
    }
    v62 := GenS8{F0: ""}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v62
    _ = v7
    _ = v8
}
