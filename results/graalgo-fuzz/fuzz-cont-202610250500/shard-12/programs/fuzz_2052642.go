package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 bool; F1 uint16; F2 uint }
type GenS5 struct { F0 uint; F1 string; F2 uint64 }
type GenS6 struct { F0 bool }
func genFn0(p0 int16, p1 uint) uint16 {
    fmt.Println(p1, p1)
    v5 := []int{32768, 4294967295, 920}
    v6 := GenS1{F0: (1 > -128), F1: 629}
    if (!(717 == 2147483648)) {
        fmt.Println(p0)
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 1
}
func genFn1(p0 uint16, p1 uint8) uint64 {
    p0 = (^(p0))
    var v7 uint16 = p0
    fmt.Println(p0, v7)
    _ = p0
    _ = p1
    _ = v7
    return 4294967295
}

func main() {
    var v1 int = 32767
    var v2 uint16 = 326
    var v3 float64 = 0.1
    var v4 uint64 = 797
    v8 := GenS1{F0: (189 >= 0), F1: v2}
    v2 -= (v2 + 532)
    var v9 int32 = -520
    if ((true && (false && (false || (("foo bar" + "\n\t") == ("\n\t" + ("" + "a")))))) && true) {
        if true {
            v10 := GenS1{F0: (255 != 9), F1: v2}
            if (!true) {
                v9 = (v9 - -23)
                if true {
                    v2 = v2
                    switch "\n\t" {
                    case "a":
                        for fi11 := 0; fi11 < 8; fi11++ {
                            var v12 int8 = 1
                            var v13 int8 = v12
                            var v14 int64 = -1
                            _ = fi11
                            _ = v12
                            _ = v13
                            _ = v14
                        }
                    default:
                        fmt.Println(v1, v1, v2)
                    }
                    v15 := []int{}
                    _ = v15
                }
                var v16 bool = ((v9 | v9) == v9)
                _ = v16
            }
            if false {
                for fi17 := 0; fi17 < 10; fi17++ {
                    switch "\n\t" {
                    case "":
                        switch "hello" {
                        case "a":
                            v18 := []int{}
                            _ = v18
                        case "hello":
                            v19 := GenS2{F0: 1}
                            _ = v19
                        default:
                            v20 := GenS2{F0: 42}
                            _ = v20
                        }
                    case "hello":
                        v21 := []int{-1, v1, -9223372036854775808, 1}
                        _ = v21
                    default:
                        var v22 uint64 = (v4 / (((v4 >> ((311) & 63))) | 1))
                        _ = v22
                    }
                    v3 -= (float64(-2.25) + float64(((((float32(v3) + ((0.1 / 3.14159) - ((1.5 / 3.14159) + (1e10 / 1.0)))) / -2.25) + (float32(v3) * 1e10)) + 0.0)))
                    var v23 bool = ("foo bar" != ("\n\t" + (((("\n\t" + (("" + "") + ("\n\t" + "\n\t"))) + "\n\t") + "hello") + "hello")))
                    _ = fi17
                    _ = v23
                }
            }
            _ = v10
        } else {
            var v24 float64 = (0.1 / (v3 - ((1.5 * -2.25) - (0.0 - (v3 * 1.5)))))
            for fi25 := 0; fi25 < 10; fi25++ {
                switch "hello" {
                case "hello":
                    v24 = (v24 * 1e10)
                case "\n\t":
                    switch (fi25 | 0) {
                    case 2147483647:
                        var v26 bool = true
                        _ = v26
                    case 42:
                        for fi27 := 0; fi27 < 5; fi27++ {
                            v28 := GenS3{F0: fi25}
                            _ = fi27
                            _ = v28
                        }
                    case 4294967295:
                        v29 := []int{(-(fi25)), v1}
                        _ = v29
                    }
                default:
                    if true {
                        v4 += genFn1((v2 / ((genFn0(-32768, 522)) | 1)), 49)
                        v9 -= (v9 ^ v9)
                        v30 := []int{412}
                        _ = v30
                    }
                }
                var v31 int8 = 100
                _ = fi25
                _ = v31
            }
            _ = v24
        }
        v32 := []int{v1, (v1 % (((v1 ^ (v1 + v1))) | 1)), (-(v1))}
        var v33 uint = 65535
        _ = v32
        _ = v33
    } else {
        v34 := []int{(-(v1)), (v1 ^ 303)}
        _ = v34
    }
    v35 := GenS3{F0: v1}
    v9 += (v9 + (v9 >> ((uint(v4)) & 63)))
    v9 += (v9 << ((2) & 63))
    if ((3.14159 >= -2.25) && false) {
        var v36 float64 = float64((0.0 * 3.14159))
        _ = v36
    }
    if true {
        v3 = v3
        for fi37 := 0; fi37 < 9; fi37++ {
            if ((false && (350 > -1)) || true) {
                v38 := GenS4{F0: false, F1: (v2 + 42), F2: uint(v2)}
                v39 := []int{}
                _ = v38
                _ = v39
            } else {
                v40 := []int{}
                v41 := GenS3{F0: 4294967295}
                _ = v40
                _ = v41
            }
            _ = fi37
        }
        switch (v1 ^ (v1 << ((4294967295) & 63))) {
        case 0:
            var v42 int16 = -124
            _ = v42
        case 235:
            var v43 bool = false
            _ = v43
        case 1:
            if ((-35 > 1) && true) {
                var v44 int8 = 42
                fmt.Println(v4, v4, v4)
                v45 := []int{}
                _ = v44
                _ = v45
            } else {
                switch "\n\t" {
                case "hello":
                    fmt.Println(v4, v3, v1)
                case "foo bar":
                    for fi46 := 0; fi46 < 7; fi46++ {
                        var v47 uint64 = (v4 * 249)
                        var v48 uint = 18446744073709551615
                        v9 = (v9 | v9)
                        _ = fi46
                        _ = v47
                        _ = v48
                    }
                default:
                    v1 = v1
                }
                v4 += v4
            }
        default:
            if ((v1 / ((0) | 1)) != -1) {
                v49 := GenS5{F0: 0, F1: "hello", F2: genFn1(v2, 87)}
                v50 := GenS1{F0: false, F1: (v2 ^ 42)}
                if false {
                    v51 := []int{v1, (^(v1)), (v1 >> ((255) & 63)), (v1 * (v1 ^ (-(v1))))}
                    _ = v51
                }
                _ = v49
                _ = v50
            } else {
                fmt.Println(v9, v3)
                v52 := GenS4{F0: true, F1: 65535, F2: 2}
                v1 += (^(v1))
                _ = v52
            }
        }
    } else {
        for fi53 := 0; fi53 < 6; fi53++ {
            v1 = (v1 << ((4294967295) & 63))
            _ = fi53
        }
    }
    switch "\n\t" {
    case "\n\t":
        for fi54 := 0; fi54 < 18; fi54++ {
            fmt.Println(v9, v2, v1)
            var v55 bool = true
            _ = fi54
            _ = v55
        }
    }
    if (-118 > 42) {
        v56 := GenS6{F0: true}
        _ = v56
    } else {
        for fi57 := 0; fi57 < 14; fi57++ {
            switch (("\n\t" + ("a" + "a")) + "foo bar") {
            case "hello":
                for fi58 := 0; fi58 < 14; fi58++ {
                    var v59 uint8 = 144
                    _ = fi58
                    _ = v59
                }
            case "foo bar":
                switch v1 {
                case -9223372036854775808:
                    var v60 string = ""
                    _ = v60
                case 32767:
                    v2 = v2
                case -831:
                    var v61 uint16 = (v2 + 32767)
                    _ = v61
                }
            }
            for fi62 := 0; fi62 < 12; fi62++ {
                v1 += fi57
                v63 := []int{fi62, fi57}
                _ = fi62
                _ = v63
            }
            _ = fi57
        }
        v64 := GenS4{F0: false, F1: v2, F2: 65535}
        _ = v64
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v8
    _ = v9
}
