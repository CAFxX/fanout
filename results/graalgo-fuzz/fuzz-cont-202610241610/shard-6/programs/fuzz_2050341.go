package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 uint8 }
type GenS2 struct { F0 int32; F1 int; F2 int }
type GenS3 struct { F0 int32; F1 int8 }
type GenS4 struct { F0 int32; F1 uint; F2 uint8 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 string; F1 int8; F2 int8 }
type GenS7 struct { F0 uint; F1 uint32; F2 uint }
type GenS8 struct { F0 uint64; F1 int; F2 string }
type GenS9 struct { F0 int16; F1 int64 }
type GenS10 struct { F0 uint; F1 uint32; F2 int16 }
func genFn0(p0 int64) uint64 {
    for fi4 := 0; fi4 < 7; fi4++ {
        if false {
            var v5 uint32 = 1
            _ = v5
        }
        if ((!false) && (42 == 0)) {
            v6 := []int{4294967295, 465}
            _ = len(v6)
            v7 := GenS1{F0: 255, F1: "a", F2: 221}
            _ = v6
            _ = v7
        }
        var v8 float32 = float32(3.14159)
        _ = fi4
        _ = v8
    }
    v9 := GenS1{F0: 255, F1: "hello", F2: 255}
    v10 := GenS2{F0: 0, F1: 32767, F2: 42}
    p0 = (p0 ^ 657)
    v11 := GenS1{F0: 2, F1: "", F2: 139}
    _ = p0
    _ = v10
    _ = v11
    _ = v9
    return 274
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]bool)
    var v3 uint8 = 2
    var v12 int16 = 1
    v12 += v12
    switch "a" {
    case "hello":
        if true {
            switch 0 {
            case 9223372036854775807:
                fmt.Println(v12, v3)
            case 0:
                v13 := []int{405, 32767}
                _ = v13
            default:
                v3 = v3
            }
        } else {
            for fi14 := 0; fi14 < 3; fi14++ {
                switch (-(fi14)) {
                case 32767:
                    v15 := GenS1{F0: 18446744073709551615, F1: "a", F2: v3}
                    _ = v15
                case -484:
                    v16 := []int{-1, fi14, fi14}
                    _ = v16
                }
                _ = fi14
            }
            v3 -= (v3 % ((255) | 1))
        }
    case "":
        if ((false && false) || true) {
            v12 -= v12
            switch "hello" {
            case "a":
                if false {
                    v17 := GenS3{F0: 1, F1: 28}
                    _ = v17
                } else {
                    v18 := v2[32767]
                    v19 := GenS4{F0: 0, F1: 42, F2: (v3 - 127)}
                    fmt.Println(v12)
                    _ = v18
                    _ = v19
                }
            case "foo bar":
                v3 += 128
            case "\n\t":
                var v20 int8 = -128
                _ = v20
            }
        } else {
            v21 := v1[0]
            _ = v21
        }
    }
    if (true && (false && (42 != -630))) {
        var v22 int32 = 2147483647
        _ = v22
    }
    if false {
        if (-2.25 <= 0.0) {
            var v23 string = "\n\t"
            _ = v23
        } else {
            v24 := GenS5{F0: ""}
            _ = v24
        }
        switch "" {
        case "":
            var v25 int8 = -95
            _ = v25
        case "hello":
            if true {
                v26 := []int{127}
                _ = v26
            } else {
                if (2 == 42) {
                    if (!true) {
                        if (3.14159 == -2.25) {
                            for fi27 := 0; fi27 < 14; fi27++ {
                                if false {
                                    v28 := GenS6{F0: "hello", F1: 42, F2: 100}
                                    _ = v28
                                }
                                _ = fi27
                            }
                            if true {
                                v29 := []int{}
                                var v30 uint8 = (v3 + 1)
                                _ = v29
                                _ = v30
                            }
                        }
                        v31 := []int{-1}
                        _ = v31
                    }
                    fmt.Println(v12, v3)
                    var v32 int64 = 0
                    _ = v32
                }
                v12 += (v12 + (v12 >> ((4294967295) & 63)))
                var v33 uint64 = 1
                _ = v33
            }
        case "a":
            switch 127 {
            case 9223372036854775807:
                for fi34 := 0; fi34 < 6; fi34++ {
                    var v35 float64 = (0.0 + 1e10)
                    var v36 int = (fi34 * fi34)
                    _ = fi34
                    _ = v35
                    _ = v36
                }
            default:
                if (!(-1 <= 42)) {
                    for fi37 := 0; fi37 < 5; fi37++ {
                        if (-1 >= -1) {
                            if true {
                                for fi38 := 0; fi38 < 9; fi38++ {
                                    switch (("" + "") + "\n\t") {
                                    case "\n\t":
                                        v39 := GenS3{F0: -1, F1: -79}
                                        _ = v39
                                    case "hello":
                                        var v40 uint32 = 2147483647
                                        if true {
                                            switch "" {
                                            case "":
                                                v41 := GenS3{F0: -824, F1: 0}
                                                _ = v41
                                            case "a":
                                                v42 := GenS7{F0: 4294967295, F1: v40, F2: 2}
                                                _ = v42
                                            case "foo bar":
                                                v43 := []int{(fi38 & (fi37 + (fi37 * (-(fi37)))))}
                                                _ = v43
                                            default:
                                                var v44 uint64 = 0
                                                _ = v44
                                            }
                                            v40 += 952
                                        }
                                        _ = v40
                                    }
                                    _ = fi38
                                }
                            }
                            var v45 int = 42
                            var v46 int = (v45 % (((^(fi37))) | 1))
                            _ = v45
                            _ = v46
                        }
                        switch "" {
                        case "hello":
                            var v47 string = ("\n\t" + ((("foo bar" + "\n\t") + ("" + "\n\t")) + "hello"))
                            _ = v47
                        default:
                            if true {
                                v48 := GenS8{F0: uint64(v3), F1: (fi37 << ((18446744073709551615) & 63)), F2: "a"}
                                delete(v2, fi37)
                                var v49 uint64 = 4294967295
                                _ = v48
                                _ = v49
                            }
                        }
                        if (!(((1.5 * (0.0 * (((0.0 + 1.5) - float32(float64((0.1 + 0.1)))) + 0.0))) == 0.0) && true)) {
                            switch (fi37 + (fi37 << ((2) & 63))) {
                            case 127:
                                fmt.Println(v3, fi37, v3)
                            default:
                                v50 := GenS9{F0: 1, F1: 976}
                                _ = v50
                            }
                            for fi51 := 0; fi51 < 3; fi51++ {
                                var v52 uint16 = 65535
                                v12 -= (v12 | v12)
                                var v53 int8 = 0
                                v52 -= (v52 - (^(v52)))
                                _ = fi51
                                _ = v52
                                _ = v53
                            }
                        }
                        _ = fi37
                    }
                } else {
                    var v54 string = "hello"
                    _ = v54
                }
            }
        }
    } else {
        fmt.Println(v3)
        if false {
            switch 42 {
            case 2147483647:
                switch ("a" + ("hello" + "foo bar")) {
                case "hello":
                    var v55 uint8 = 83
                    _ = v55
                default:
                    switch 128 {
                    case 4294967295:
                        for fi56 := 0; fi56 < 20; fi56++ {
                            v12 -= (v12 - (v12 / (((v12 % ((v12) | 1))) | 1)))
                            v57 := GenS8{F0: 0, F1: -1, F2: "a"}
                            switch ("\n\t" + "") {
                            case "\n\t":
                                v12 = v12
                            case "a":
                                switch ("\n\t" + (("hello" + "foo bar") + "")) {
                                case "hello":
                                    switch "hello" {
                                    case "a":
                                        v58 := []int{42, 128}
                                        _ = v58
                                    case "hello":
                                        v12 = v12
                                    case "foo bar":
                                        v59 := GenS3{F0: -1, F1: 1}
                                        _ = v59
                                    }
                                default:
                                    var v60 string = ((("" + ((((("a" + "hello") + "\n\t") + "\n\t") + "foo bar") + "foo bar")) + ("hello" + (("" + (("" + ("foo bar" + ("\n\t" + "a"))) + "a")) + "\n\t"))) + ("foo bar" + "hello"))
                                    _ = v60
                                }
                            case "foo bar":
                                for fi61 := 0; fi61 < 8; fi61++ {
                                    var v62 uint8 = (v3 * 239)
                                    _ = fi61
                                    _ = v62
                                }
                            }
                            _ = fi56
                            _ = v57
                        }
                    }
                }
            case 94:
                if false {
                    for fi63 := 0; fi63 < 9; fi63++ {
                        var v64 int8 = 100
                        var v65 uint = 255
                        _ = fi63
                        _ = v64
                        _ = v65
                    }
                    var v66 uint = 42
                    _ = v66
                }
            }
            switch 9223372036854775807 {
            case 1:
                v3 = (v3 & 127)
            case 128:
                v3 = (v3 & v3)
            default:
                v12 -= (v12 << ((967) & 63))
            }
            for fi67 := 0; fi67 < 18; fi67++ {
                v68 := GenS10{F0: 255, F1: 2, F2: v12}
                _ = fi67
                _ = v68
            }
        } else {
            var v69 uint = 42
            var v70 int16 = (v12 & (v12 / ((v12) | 1)))
            _ = v69
            _ = v70
        }
    }
    v71 := v1[-9223372036854775808]
    var v72 uint8 = (^(v3))
    var v73 int8 = 106
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v71:", v71)
    fmt.Println("v72:", v72)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v71
    _ = v72
    _ = v73
}
