package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 uint16 }
type GenS2 struct { F0 uint64; F1 uint32 }
type GenS3 struct { F0 uint16; F1 string; F2 string }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int32; F1 uint8 }
type GenS6 struct { F0 int8; F1 uint16; F2 int64 }

func main() {
    var v1 bool = true
    var v2 float64 = (0.0 + (-2.25 - float64(1.5)))
    v3 := make(map[int]int)
    v4 := make(map[int]string)
    v2 += 3.14159
    v2 = ((float64(-2.25) * float64(float32(v2))) / 1.0)
    fmt.Println(v1, v1)
    for fi5 := 0; fi5 < 10; fi5++ {
        v6 := []int{(-(fi5)), (fi5 | fi5), (fi5 & (fi5 / (((-(fi5))) | 1))), fi5}
        if (!(!true)) {
            v7 := GenS1{F0: ("hello" + "hello"), F1: 0, F2: 65535}
            _ = v7
        }
        for fi8 := 0; fi8 < 12; fi8++ {
            switch fi5 {
            case -9223372036854775808:
                switch fi5 {
                case 42:
                    v6 = append(v6, (fi8 + -9223372036854775808))
                }
            case 268:
                v9 := GenS1{F0: "hello", F1: 255, F2: 255}
                _ = v9
            default:
                v2 -= v2
            }
            var v10 string = ("" + "a")
            _ = fi8
            _ = v10
        }
        _ = fi5
        _ = v6
    }
    if true {
        var v11 int64 = 9223372036854775807
        if (false && (32767 >= 1)) {
            v11 = -9223372036854775808
            v12 := GenS1{F0: (("foo bar" + "a") + "a"), F1: 451, F2: 263}
            _ = v12
        } else {
            var v13 bool = (!("" == "\n\t"))
            _ = len(v4)
            for fi14 := 0; fi14 < 8; fi14++ {
                for fi15 := 0; fi15 < 19; fi15++ {
                    v1 = v1
                    v16 := GenS2{F0: 42, F1: 2147483647}
                    _ = fi15
                    _ = v16
                }
                v17 := []int{(fi14 * (fi14 ^ fi14))}
                for fi18 := 0; fi18 < 18; fi18++ {
                    fmt.Println(v11, v11, v2)
                    if v13 {
                        fmt.Println(fi18, v13, fi18)
                    } else {
                        var v19 int = (fi14 >> ((2) & 63))
                        _ = v19
                    }
                    _ = fi18
                }
                _ = fi14
                _ = v17
            }
            _ = v13
        }
        switch 128 {
        case 42:
            for fi20 := 0; fi20 < 20; fi20++ {
                v2 -= ((v2 - 1.5) + v2)
                if v1 {
                    if (!(!(!v1))) {
                        switch fi20 {
                        case 0:
                            switch fi20 {
                            case 42:
                                fmt.Println(v1, v1)
                            default:
                                v11 += (^(v11))
                            }
                        case -1:
                            for fi21 := 0; fi21 < 17; fi21++ {
                                var v22 string = ("" + ("foo bar" + "foo bar"))
                                _ = fi21
                                _ = v22
                            }
                        }
                    } else {
                        v2 += (v2 + (((1.5 * ((float64(((0.0 / 0.1) * (1.5 / (1e10 + (-2.25 - float32(0.0)))))) - v2) - (v2 * float64((1.5 - -2.25))))) * (0.0 - v2)) + (v2 / v2)))
                        switch "\n\t" {
                        case "":
                            if false {
                                for fi23 := 0; fi23 < 2; fi23++ {
                                    v24 := v3[fi20]
                                    var v25 int = v24
                                    var v26 uint8 = 166
                                    _ = fi23
                                    _ = v24
                                    _ = v25
                                    _ = v26
                                }
                            }
                        case "a":
                            for fi27 := 0; fi27 < 18; fi27++ {
                                fmt.Println(fi20, v1)
                                v2 += (v2 + v2)
                                if false {
                                    var v28 float32 = (1.5 * 0.0)
                                    fmt.Println(fi20)
                                    _ = v28
                                } else {
                                    v2 = v2
                                    v2 = v2
                                    v11 = v11
                                }
                                _ = fi27
                            }
                        case "hello":
                            var v29 int64 = v11
                            _ = v29
                        }
                    }
                    switch "\n\t" {
                    case "a":
                        for fi30 := 0; fi30 < 15; fi30++ {
                            if v1 {
                                var v31 float64 = v2
                                if (!v1) {
                                    switch fi20 {
                                    case -204:
                                        fmt.Println(fi30, v31)
                                    case 42:
                                        v2 += 0.1
                                    }
                                    var v32 uint8 = 2
                                    _ = v32
                                } else {
                                    if v1 {
                                        var v33 uint8 = 0
                                        var v34 int = 127
                                        _ = v33
                                        _ = v34
                                    } else {
                                        fmt.Println(fi20)
                                        v2 += float64(1.5)
                                    }
                                    switch fi30 {
                                    case 42:
                                        var v35 int8 = -1
                                        _ = v35
                                    case -9223372036854775808:
                                        var v36 uint = 2
                                        _ = v36
                                    default:
                                        var v37 uint8 = 255
                                        _ = v37
                                    }
                                    switch (("" + (((("a" + "") + ("a" + ("hello" + "a"))) + "foo bar") + ("" + "\n\t"))) + ("\n\t" + "\n\t")) {
                                    case "foo bar":
                                        v2 += (v2 + (v31 + 0.0))
                                        v11 = (v11 ^ (v11 | v11))
                                        var v38 int8 = 42
                                        _ = v38
                                    case "\n\t":
                                        if (((v31 * v2) / v31) < (v2 + (1e10 + v31))) {
                                            v39 := GenS3{F0: 32767, F1: "\n\t", F2: ""}
                                            v40 := GenS4{F0: (fi20 - (fi20 >> ((255) & 63)))}
                                            v41 := GenS3{F0: 2, F1: "a", F2: "\n\t"}
                                            _ = v39
                                            _ = v40
                                            _ = v41
                                        }
                                    default:
                                        v31 -= v31
                                    }
                                }
                                var v42 float32 = float32(v2)
                                if (0 < 2) {
                                    var v43 float32 = 1.5
                                    v44 := []int{127}
                                    _ = v43
                                    _ = v44
                                } else {
                                    var v45 string = (("hello" + ("foo bar" + "a")) + "foo bar")
                                    _ = v45
                                }
                                _ = v31
                                _ = v42
                            }
                            v2 = (1e10 + v2)
                            v2 += (float64(1e10) + (v2 * v2))
                            _ = fi30
                        }
                    case "hello":
                        v11 += (^(v11))
                    }
                    if (false || v1) {
                        _ = len(v4)
                        for fi46 := 0; fi46 < 12; fi46++ {
                            if (424 != 792) {
                                v47 := GenS5{F0: -2147483648, F1: 2}
                                _ = v47
                            }
                            fmt.Println(fi46, fi20)
                            _ = fi46
                        }
                    }
                }
                if (!(32768 < (^(fi20)))) {
                    for fi48 := 0; fi48 < 18; fi48++ {
                        v11 -= 42
                        _ = fi48
                    }
                }
                _ = fi20
            }
        default:
            var v49 uint32 = 4294967295
            _ = v49
        }
        _ = v11
    }
    var v50 uint16 = 65535
    v51 := GenS4{F0: 279}
    switch "a" {
    case "a":
        switch 42 {
        case 2147483647:
            for fi52 := 0; fi52 < 8; fi52++ {
                v53 := GenS6{F0: 81, F1: 2, F2: -1}
                v50 = v50
                _ = fi52
                _ = v53
            }
        case 0:
            v54 := []int{9223372036854775807}
            _ = v54
        case 32767:
            v1 = (694 <= int(v50))
        }
    case "hello":
        v55 := []int{}
        _ = v55
    case "foo bar":
        v56 := []int{127, -714, 2147483647}
        _ = v56
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v50
    _ = v51
}
