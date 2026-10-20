package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool; F1 string; F2 uint16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 bool; F1 uint32; F2 uint16 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]int)
    var v3 uint64 = 4294967295
    v4 := []string{"", "", "a", "\n\t"}
    v5 := []int{-1, 128, -2147483648}
    var v6 uint = 340
    if len(v4) > 0 {
        v4[0] = "a"
    }
    var v7 int32 = int32(v3)
    v3 = 818
    var v8 string = (("\n\t" + "") + "")
    var v9 bool = (!(false && false))
    v6 -= (v6 ^ (v6 >> ((1) & 63)))
    switch 581 {
    case 32768:
        for fi10 := 0; fi10 < 2; fi10++ {
            var v11 float32 = -2.25
            switch fi10 {
            case 726:
                v5 = append(v5, (fi10 & 373))
            case 32767:
                if (uint(v3) == v6) {
                    v9 = ((true && (!(true && (v6 <= (v6 | 992))))) && ((((^(v6)) < (v6 | (v6 << ((888) & 63)))) || (false || (42 > 65535))) && (0 > 4294967295)))
                    v3 = (^(v3))
                    if false {
                        for fi12 := 0; fi12 < 4; fi12++ {
                            var v13 bool = (!(v9 && (!(1 != 0))))
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            if len(v5) > 0 {
                                v5[0] = (-(fi10))
                            }
                            switch ("\n\t" + "") {
                            case "\n\t":
                                var v14 int16 = 0
                                v8 = v8
                                switch (^(fi12)) {
                                case 4294967295:
                                    switch v8 {
                                    case "foo bar":
                                        var v15 string = ("hello" + (("\n\t" + v8) + ""))
                                        _ = v15
                                    case "a":
                                        var v16 int8 = 100
                                        _ = v16
                                    case "\n\t":
                                        v9 = ((0 != 255) && (true && (v14 > 42)))
                                    default:
                                        var v17 float32 = 3.14159
                                        _ = v17
                                    }
                                case 127:
                                    switch ((((((((v8 + "a") + v8) + v8) + v8) + (v8 + v8)) + "a") + (v8 + "foo bar")) + "") {
                                    case "":
                                        switch "foo bar" {
                                        case "hello":
                                            if false {
                                                _ = len(v5)
                                                v18 := v1[1]
                                                _ = v18
                                            }
                                        case "a":
                                            if (v14 != 1000) {
                                                switch fi12 {
                                                case 4294967295:
                                                    v19 := GenS1{F0: v3}
                                                    _ = v19
                                                case 12:
                                                    v11 -= v11
                                                default:
                                                    fmt.Println(v6)
                                                }
                                            } else {
                                                v8 = (v8 + v8)
                                            }
                                            fmt.Println(v6, v8, v6)
                                        case "foo bar":
                                            var v20 uint8 = 2
                                            _ = v20
                                        }
                                    case "hello":
                                        delete(v1, 2147483648)
                                    }
                                }
                                _ = v14
                            case "":
                                var v21 string = (("a" + v8) + v8)
                                _ = v21
                            case "hello":
                                if (v8 != v8) {
                                    if v13 {
                                        v6 = (v6 + v6)
                                        v22 := GenS1{F0: (v3 % ((916) | 1))}
                                        v23 := GenS1{F0: (v3 % ((v3) | 1))}
                                        _ = v22
                                        _ = v23
                                    } else {
                                        v11 = v11
                                        v24 := GenS1{F0: v3}
                                        if (!false) {
                                            v25 := GenS1{F0: v3}
                                            _ = len(v1)
                                            _ = v25
                                        }
                                        _ = v24
                                    }
                                    var v26 float64 = 1e10
                                    if ((65535 <= 42) && (0 != v3)) {
                                        var v27 uint32 = 426
                                        _ = v27
                                    }
                                    _ = v26
                                }
                            default:
                                fmt.Println(v11, fi10)
                            }
                            v28 := v1[65535]
                            switch "" {
                            case "hello":
                                var v29 int32 = v7
                                _ = v29
                            }
                            if (false && (65535 >= 65535)) {
                                var v30 int16 = 32767
                                _ = v30
                            } else {
                                if (!((v8 + ((v8 + ((v8 + v8) + (v8 + ""))) + "\n\t")) == v8)) {
                                    v6 = (v6 & uint(v3))
                                }
                                fmt.Println(v9, v11)
                            }
                            _ = fi12
                            _ = v13
                            _ = v28
                        }
                        v8 = ("foo bar" + "\n\t")
                        for fi31 := 0; fi31 < 6; fi31++ {
                            v32 := v1[2147483647]
                            switch (((("hello" + v8) + (v8 + "hello")) + v8) + v8) {
                            case "\n\t":
                                v33 := GenS2{F0: ((fi31 * fi10) != fi10)}
                                _ = v33
                            default:
                                var v34 int64 = -9223372036854775808
                                _ = v34
                            }
                            _ = fi31
                            _ = v32
                        }
                    }
                } else {
                    var v35 bool = (1 < 0)
                    _ = len(v4)
                    for fi36 := 0; fi36 < 11; fi36++ {
                        v37 := GenS1{F0: v3}
                        v3 = (v3 << ((v6) & 63))
                        _ = fi36
                        _ = v37
                    }
                    _ = v35
                }
            case 1:
                var v38 int64 = 0
                _ = v38
            }
            _ = fi10
            _ = v11
        }
    default:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    switch "hello" {
    case "\n\t":
        var v39 float32 = ((0.1 - (3.14159 + (((3.14159 / 1e10) / (0.1 * -2.25)) + (-2.25 - 3.14159)))) * float32(0.0))
        _ = v39
    case "a":
        var v40 int8 = 127
        _ = v40
    case "foo bar":
        switch "hello" {
        case "foo bar":
            fmt.Println(v6, v8)
        case "":
            if (!true) {
                if ((v3 << ((v6) & 63)) != 42) {
                    for fi41 := 0; fi41 < 6; fi41++ {
                        v42 := GenS3{F0: v9, F1: ("\n\t" + "foo bar"), F2: 1}
                        _ = fi41
                        _ = v42
                    }
                    v8 = (("hello" + (((v8 + v8) + v8) + v8)) + v8)
                    if v9 {
                        v43 := GenS1{F0: v3}
                        _ = v43
                    }
                }
                var v44 float32 = 1.5
                v45 := GenS4{F0: 32}
                _ = v44
                _ = v45
            } else {
                switch 4294967295 {
                case 518:
                    if false {
                        switch ("a" + v8) {
                        case "hello":
                            _ = len(v5)
                        case "a":
                            for fi46 := 0; fi46 < 19; fi46++ {
                                fmt.Println(v7, v8)
                                _ = fi46
                            }
                        }
                    }
                case 9223372036854775807:
                    v47 := GenS3{F0: true, F1: (("" + (v8 + v8)) + v8), F2: 65535}
                    _ = v47
                case 32768:
                    var v48 bool = (185 != 1)
                    _ = v48
                default:
                    v49 := GenS1{F0: (^(v3))}
                    _ = v49
                }
                v50 := GenS5{F0: ((v8 + "hello") != ""), F1: 632, F2: 32767}
                _ = v50
            }
        case "a":
            var v51 float64 = (float64(1e10) - 1e10)
            _ = v51
        default:
            var v52 float64 = (-2.25 - (1.5 / 1e10))
            _ = v52
        }
    }
    var v53 int = 1
    var v54 uint8 = 0
    for fi55 := 0; fi55 < 18; fi55++ {
        var v56 int = v53
        _ = fi55
        _ = v56
    }
    for fi57 := 0; fi57 < 7; fi57++ {
        switch "" {
        case "hello":
            v58 := GenS2{F0: v9}
            _ = v58
        case "foo bar":
            v6 += v6
        default:
            v59 := GenS1{F0: v3}
            _ = v59
        }
        for fi60 := 0; fi60 < 12; fi60++ {
            var v61 int64 = 9223372036854775807
            v62 := GenS3{F0: v9, F1: "a", F2: 42}
            v9 = v9
            _ = fi60
            _ = v61
            _ = v62
        }
        switch 1 {
        case -9223372036854775808:
            var v63 bool = false
            _ = v63
        default:
            for fi64 := 0; fi64 < 7; fi64++ {
                v65 := GenS1{F0: 18446744073709551615}
                _ = fi64
                _ = v65
            }
        }
        _ = fi57
    }
    switch v53 {
    case 32767:
        v3 = (v3 ^ (v3 - v3))
    case -1:
        if (!v9) {
            v6 -= (v6 | uint(v3))
        }
    default:
        if len(v5) > 0 {
            v5[0] = 128
        }
    }
    var v66 uint32 = 42
    _ = len(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v6:", v6)
    fmt.Println("v66:", v66)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v54
    _ = v6
    _ = v66
    _ = v7
    _ = v8
    _ = v9
}
