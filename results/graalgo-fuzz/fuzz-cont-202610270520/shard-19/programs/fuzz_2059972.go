package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 uint; F2 string }
type GenS3 struct { F0 int; F1 string; F2 uint16 }
type GenS4 struct { F0 uint64; F1 int8 }
type GenS5 struct { F0 int64; F1 uint }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]string)
    var v3 uint = 755
    v4 := GenS1{F0: ""}
    v3 += v3
    v5 := GenS1{F0: ""}
    if (!(-842 == 87)) {
        v6 := []int{932, -492, 128}
        v1[9223372036854775807] = true
        _ = v6
    } else {
        switch -2147483648 {
        case -2147483648:
            v3 -= 255
        case -79:
            fmt.Println(v3)
        case -9223372036854775808:
            var v7 int32 = 0
            _ = v7
        default:
            for fi8 := 0; fi8 < 14; fi8++ {
                switch (("a" + "a") + "foo bar") {
                case "a":
                    if false {
                        v3 = v3
                        for fi9 := 0; fi9 < 16; fi9++ {
                            fmt.Println(fi9)
                            _ = fi9
                        }
                        v10 := []int{32767, 127}
                        _ = v10
                    }
                case "\n\t":
                    for fi11 := 0; fi11 < 18; fi11++ {
                        var v12 float32 = -2.25
                        v13 := GenS2{F0: 2147483647, F1: v3, F2: (("" + ("\n\t" + ("" + ("a" + (("a" + "") + "\n\t"))))) + "hello")}
                        var v14 uint64 = 42
                        _ = fi11
                        _ = v12
                        _ = v13
                        _ = v14
                    }
                default:
                    v15 := GenS3{F0: fi8, F1: ("" + "\n\t"), F2: 1}
                    _ = v15
                }
                for fi16 := 0; fi16 < 6; fi16++ {
                    var v17 uint8 = 0
                    _ = fi16
                    _ = v17
                }
                _ = fi8
            }
        }
        var v18 uint8 = 2
        _ = v18
    }
    v19 := []int{42, -9223372036854775808, 70, 447}
    _ = len(v19)
    for fi20 := 0; fi20 < 4; fi20++ {
        var v21 int = fi20
        _ = fi20
        _ = v21
    }
    v22 := GenS4{F0: 511, F1: 0}
    var v23 uint16 = 65535
    var v24 string = "foo bar"
    switch 32768 {
    case 382:
        for fi25 := 0; fi25 < 12; fi25++ {
            switch fi25 {
            case 4294967295:
                var v26 uint = (^(v3))
                _ = v26
            case 128:
                v27 := GenS2{F0: 421, F1: (v3 + v3), F2: v24}
                _ = v27
            case 9223372036854775807:
                if (true && false) {
                    v23 -= 3
                    for fi28 := 0; fi28 < 4; fi28++ {
                        var v29 float32 = float32(float64(float32(float64((0.0 / 1e10)))))
                        fmt.Println(v24, fi25)
                        fmt.Println(v29, v29, fi25)
                        _ = fi28
                        _ = v29
                    }
                } else {
                    switch (v24 + v24) {
                    case "foo bar":
                        v30 := GenS1{F0: (v24 + "a")}
                        _ = v30
                    case "":
                        v23 = 0
                    case "hello":
                        var v31 uint = (^(v3))
                        _ = v31
                    default:
                        var v32 bool = (-821 >= 292)
                        _ = v32
                    }
                    var v33 uint16 = (v23 & 32767)
                    var v34 uint32 = 4294967295
                    _ = v33
                    _ = v34
                }
            default:
                v3 -= 148
            }
            _ = len(v19)
            _ = fi25
        }
    default:
        if (false && (1e10 != (3.14159 - (3.14159 - 0.0)))) {
            for fi35 := 0; fi35 < 8; fi35++ {
                for fi36 := 0; fi36 < 4; fi36++ {
                    v37 := GenS4{F0: 18446744073709551615, F1: 42}
                    _ = fi36
                    _ = v37
                }
                _ = fi35
            }
            v38 := GenS5{F0: 1, F1: (v3 - v3)}
            _ = v38
        } else {
            fmt.Println(v3, v23)
            fmt.Println(v23)
            v23 += v23
        }
    }
    v39 := GenS2{F0: 42, F1: (v3 + 431), F2: ("hello" + v24)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v39
    _ = v4
    _ = v5
}
