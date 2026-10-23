package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 int32; F1 int32; F2 bool }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int; F1 int32; F2 uint32 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint64; F1 int64 }

func main() {
    var v1 int = 128
    var v2 uint8 = 255
    var v3 int32 = 0
    if true {
        var v4 uint64 = 1
        v2 -= (v2 ^ v2)
        _ = v4
    } else {
        switch "\n\t" {
        case "":
            var v5 int64 = -386
            _ = v5
        case "foo bar":
            v2 -= v2
        case "\n\t":
            var v6 float64 = (0.1 / 0.1)
            _ = v6
        default:
            switch "foo bar" {
            case "a":
                v7 := GenS1{F0: "\n\t", F1: 0}
                _ = v7
            case "":
                v1 = (v1 / (((v1 << ((2) & 63))) | 1))
            }
        }
        v2 = (^(v2))
        var v8 int16 = 1000
        _ = v8
    }
    var v9 uint = 202
    var v10 bool = (!(0 > 467))
    for fi11 := 0; fi11 < 15; fi11++ {
        v12 := GenS1{F0: "hello", F1: -1}
        if false {
            if (!true) {
                v13 := []int{(fi11 * 32767), (v1 + v1), v1}
                v14 := GenS2{F0: (v3 ^ v3), F1: 1, F2: true}
                switch -9223372036854775808 {
                case -226:
                    var v15 int8 = 0
                    _ = v15
                default:
                    v16 := GenS1{F0: ("" + ""), F1: 0}
                    _ = v16
                }
                _ = v13
                _ = v14
            } else {
                if (83 <= -128) {
                    var v17 uint64 = 930
                    var v18 int16 = -721
                    _ = v17
                    _ = v18
                }
                v10 = true
                v9 = v9
            }
            v19 := GenS3{F0: -287}
            v3 = (v3 ^ (v3 + (v3 - (v3 & v3))))
            _ = v19
        } else {
            fmt.Println(v9, fi11)
        }
        _ = fi11
        _ = v12
    }
    v20 := GenS1{F0: "", F1: -9223372036854775808}
    var v21 uint = v9
    v22 := GenS2{F0: -574, F1: (v3 - (v3 | -282)), F2: false}
    v1 += (v1 ^ (v1 * (v1 ^ (v1 - (v1 << ((835) & 63))))))
    for fi23 := 0; fi23 < 20; fi23++ {
        var v24 int = (-(v1))
        if false {
            v25 := []int{(v1 - v24)}
            for fi26 := 0; fi26 < 16; fi26++ {
                if (4294967295 > 42) {
                    fmt.Println(v2, v1)
                }
                v27 := GenS4{F0: 127, F1: (v3 * 1), F2: 0}
                switch ("a" + "") {
                case "\n\t":
                    v28 := GenS1{F0: "\n\t", F1: -1}
                    _ = v28
                }
                v3 -= (v3 ^ 0)
                fmt.Println(v3, v10)
                switch -1 {
                case -9223372036854775808:
                    v29 := GenS1{F0: ("hello" + "\n\t"), F1: -784}
                    _ = v29
                case 127:
                    var v30 int8 = 42
                    _ = v30
                case 128:
                    fmt.Println(v3, v10, fi23)
                }
                if v10 {
                    v31 := GenS3{F0: 1}
                    _ = v31
                } else {
                    v21 += v21
                    v24 = (^(v1))
                }
                _ = fi26
                _ = v27
            }
            _ = v25
        } else {
            v3 += (v3 << ((v21) & 63))
        }
        _ = fi23
        _ = v24
    }
    if v10 {
        v21 = 65535
        v32 := []int{(v1 & -9223372036854775808), 2147483647, v1, (v1 | v1)}
        var v33 int16 = 118
        _ = v32
        _ = v33
    }
    switch 1 {
    case 4294967295:
        if ((!v10) || (42 >= 2)) {
            fmt.Println(v3)
            v34 := []int{}
            for fi35 := 0; fi35 < 20; fi35++ {
                v36 := GenS1{F0: ("foo bar" + (("\n\t" + "") + "hello")), F1: 9223372036854775807}
                fmt.Println(v2, v2, v10)
                var v37 bool = v10
                _ = fi35
                _ = v36
                _ = v37
            }
            _ = v34
        } else {
            var v38 int8 = -1
            switch "" {
            case "a":
                if true {
                    if v10 {
                        var v39 uint = (v9 << ((1) & 63))
                        _ = v39
                    }
                    fmt.Println(v9)
                } else {
                    fmt.Println(v1)
                    v40 := []int{-2147483648, (v1 * v1)}
                    _ = v40
                }
            case "hello":
                v41 := GenS5{F0: 11}
                _ = v41
            default:
                switch "hello" {
                case "":
                    v38 += (v38 - (v38 - v38))
                case "\n\t":
                    v42 := []int{}
                    _ = v42
                case "foo bar":
                    v1 += v1
                default:
                    v3 += v3
                }
            }
            _ = v38
        }
    case -2147483648:
        for fi43 := 0; fi43 < 20; fi43++ {
            for fi44 := 0; fi44 < 3; fi44++ {
                if (42 > 0) {
                    v45 := GenS6{F0: v10}
                    v9 = (v21 & v9)
                    _ = v45
                }
                v46 := GenS7{F0: 4294967295, F1: 0}
                var v47 bool = (2 > (v9 & v9))
                _ = fi44
                _ = v46
                _ = v47
            }
            v1 = (^(fi43))
            v10 = (!((!v10) || v10))
            _ = fi43
        }
    case 128:
        v9 += (v9 - v9)
    }
    switch "foo bar" {
    case "":
        var v48 float64 = float64(-2.25)
        _ = v48
    }
    switch v1 {
    case 0:
        if (v1 >= v1) {
            v10 = false
            v3 += (v3 ^ (v3 + 437))
        }
    }
    if v10 {
        fmt.Println(v1)
    }
    fmt.Println(v21, v2)
    fmt.Println(v10, v3)
    for fi49 := 0; fi49 < 17; fi49++ {
        v3 = (v3 & v3)
        v50 := []int{}
        _ = fi49
        _ = v50
    }
    v51 := []int{v1}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v51:", len(v51), v51)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v51
    _ = v9
}
