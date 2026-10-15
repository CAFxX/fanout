package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 uint64 }
type GenS2 struct { F0 int8; F1 int32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint16; F1 int16 }
type GenS5 struct { F0 int32; F1 uint32; F2 bool }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int16; F1 uint64; F2 int16 }
func genFn0(p0 uint8, p1 uint64, p2 string) uint64 {
    if (!false) {
        for fi4 := 0; fi4 < 9; fi4++ {
            v5 := GenS1{F0: 42, F1: 100, F2: (p1 >> ((42) & 63))}
            v6 := GenS1{F0: 42, F1: 100, F2: p1}
            _ = fi4
            _ = v5
            _ = v6
        }
        v7 := GenS1{F0: -1, F1: -128, F2: p1}
        p2 = p2
        _ = v7
    }
    p1 += (p1 + 0)
    if (-128 < -25) {
        p0 += (p0 + 116)
        var v8 uint16 = 1
        fmt.Println(p1)
        _ = v8
    }
    var v9 int8 = 83
    p0 += (p0 << ((2) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return (p1 >> ((255) & 63))
}
func genFn1(p0 uint) int8 {
    v10 := GenS1{F0: 9223372036854775807, F1: 127, F2: 921}
    for fi11 := 0; fi11 < 6; fi11++ {
        if (32767 > 507) {
            v12 := GenS2{F0: 42, F1: -425}
            _ = v12
        }
        var v13 int8 = 0
        var v14 float64 = 0.0
        _ = fi11
        _ = v13
        _ = v14
    }
    _ = p0
    _ = v10
    return 13
}

func main() {
    var v1 uint8 = 255
    var v2 uint8 = v1
    var v3 int64 = 9223372036854775807
    if (-2.25 != (((-2.25 + 0.0) - (0.1 - 1.5)) - 1.5)) {
        v15 := GenS3{F0: -152}
        var v16 string = "a"
        var v17 int = -1
        _ = v15
        _ = v16
        _ = v17
    } else {
        fmt.Println(v1, v1, v2)
        var v18 float64 = 1e10
        _ = v18
    }
    var v19 uint16 = 1
    fmt.Println(v19, v3, v2)
    v20 := GenS4{F0: v19, F1: -473}
    v19 -= v19
    switch "a" {
    case "foo bar":
        for fi21 := 0; fi21 < 20; fi21++ {
            v22 := GenS3{F0: 1}
            v23 := []int{(fi21 << ((811) & 63))}
            _ = fi21
            _ = v22
            _ = v23
        }
    default:
        var v24 string = ""
        _ = v24
    }
    if true {
        if (!((1.5 / 0.1) < 3.14159)) {
            v3 = v3
        } else {
            var v25 uint32 = 1
            for fi26 := 0; fi26 < 20; fi26++ {
                fmt.Println(fi26, v25)
                _ = fi26
            }
            v27 := GenS1{F0: (v3 * (v3 | v3)), F1: 1, F2: 2}
            _ = v25
            _ = v27
        }
        for fi28 := 0; fi28 < 9; fi28++ {
            for fi29 := 0; fi29 < 14; fi29++ {
                v30 := GenS5{F0: 355, F1: 1, F2: (1 != (v3 | (-(v3))))}
                v31 := []int{-9223372036854775808, fi29, -508}
                _ = fi29
                _ = v30
                _ = v31
            }
            _ = fi28
        }
        v32 := []int{2147483647}
        _ = v32
    } else {
        var v33 bool = (v3 < (v3 - (-(v3))))
        fmt.Println(v1)
        var v34 int8 = -68
        _ = v33
        _ = v34
    }
    v2 -= v1
    fmt.Println(v1, v19, v19)
    fmt.Println(v1)
    if true {
        v35 := []int{-46, 275, 1}
        _ = v35
    } else {
        v19 += (^(v19))
        if (false || ((false && true) && (v19 >= v19))) {
            fmt.Println(v2, v3)
        }
    }
    if (true && true) {
        if (-722 <= 166) {
            for fi36 := 0; fi36 < 9; fi36++ {
                var v37 int = fi36
                v1 += v1
                _ = fi36
                _ = v37
            }
            v38 := []int{}
            _ = v38
        } else {
            v39 := GenS5{F0: -507, F1: 2147483647, F2: false}
            var v40 uint8 = 40
            _ = v39
            _ = v40
        }
        fmt.Println(v2)
        v2 = v2
    } else {
        if (!true) {
            v19 += (^(v19))
            for fi41 := 0; fi41 < 18; fi41++ {
                var v42 float32 = 3.14159
                fmt.Println(v1)
                _ = fi41
                _ = v42
            }
        }
        if (-1 <= 42) {
            var v43 int = -9223372036854775808
            fmt.Println(v1, v3, v43)
            v1 = v1
            _ = v43
        }
        fmt.Println(v1, v19, v1)
    }
    for fi44 := 0; fi44 < 5; fi44++ {
        if (!((!false) || true)) {
            for fi45 := 0; fi45 < 3; fi45++ {
                v46 := []int{fi44}
                _ = fi45
                _ = v46
            }
            v1 += (v2 | v1)
        }
        _ = fi44
    }
    v3 -= 42
    v47 := []int{2147483647, 0, 732}
    if (uint16(v2) >= (v19 / ((42) | 1))) {
        var v48 uint8 = (v1 ^ 2)
        for fi49 := 0; fi49 < 12; fi49++ {
            v50 := GenS6{F0: 0}
            for fi51 := 0; fi51 < 11; fi51++ {
                v3 = -1
                _ = fi51
            }
            _ = fi49
            _ = v50
        }
        for fi52 := 0; fi52 < 7; fi52++ {
            for fi53 := 0; fi53 < 14; fi53++ {
                if len(v47) > 0 {
                    v47[0] = 9223372036854775807
                }
                fmt.Println(v48)
                var v54 int64 = (v3 % ((v3) | 1))
                _ = fi53
                _ = v54
            }
            for fi55 := 0; fi55 < 5; fi55++ {
                switch ("foo bar" + ("hello" + (("hello" + ("a" + (((("a" + "") + ("hello" + "")) + "") + ("\n\t" + "\n\t")))) + ((("\n\t" + ((("a" + "") + "") + "a")) + "hello") + "\n\t")))) {
                case "hello":
                    v56 := GenS7{F0: 36, F1: 722, F2: -1}
                    _ = v56
                default:
                    v57 := GenS7{F0: 42, F1: 1, F2: -701}
                    _ = v57
                }
                switch fi52 {
                case -694:
                    switch "foo bar" {
                    case "\n\t":
                        v19 = 0
                    case "foo bar":
                        if false {
                            var v58 uint8 = v48
                            v59 := GenS1{F0: (v3 >> ((uint(v58)) & 63)), F1: 127, F2: 307}
                            var v60 uint8 = (v58 ^ (v2 ^ v2))
                            _ = v58
                            _ = v59
                            _ = v60
                        }
                        v2 = (v1 % ((183) | 1))
                    }
                case -2147483648:
                    if len(v47) > 0 {
                        v47[0] = (fi52 + (fi55 << ((4294967295) & 63)))
                    }
                default:
                    v3 = v3
                }
                fmt.Println(v2)
                _ = fi55
            }
            _ = fi52
        }
        _ = v48
    } else {
        fmt.Println(v1)
        _ = len(v47)
    }
    for fi61 := 0; fi61 < 15; fi61++ {
        v1 = (v1 / ((v1) | 1))
        var v62 int = fi61
        v3 -= int64(v19)
        _ = fi61
        _ = v62
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v47:", len(v47), v47)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v47
}
