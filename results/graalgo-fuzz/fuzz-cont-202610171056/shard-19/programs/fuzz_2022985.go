package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 int16 }
type GenS2 struct { F0 uint8; F1 string; F2 int32 }
type GenS3 struct { F0 uint16; F1 int64 }
type GenS4 struct { F0 uint16; F1 int64; F2 string }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 bool; F1 int8; F2 string }
type GenS7 struct { F0 uint8 }

func main() {
    var v1 uint = 1
    v2 := make(map[string]uint32)
    var v3 int16 = 1000
    v4 := make(map[uint32]string)
    if (false && (!true)) {
        for fi5 := 0; fi5 < 17; fi5++ {
            var v6 uint = (^(v1))
            if (!(0 != -764)) {
                v7 := GenS1{F0: (v6 ^ 2), F1: (v3 | (v3 + 1)), F2: (v3 - (^(v3)))}
                if false {
                    if false {
                        v6 += (^(v1))
                    }
                    v6 += 733
                } else {
                    var v8 uint32 = 2
                    fmt.Println(v6)
                    _ = v8
                }
                _ = v7
            }
            v9 := v4[42]
            _ = fi5
            _ = v6
            _ = v9
        }
    }
    v3 = (-(v3))
    fmt.Println(v3, v3)
    if false {
        v10 := GenS1{F0: (v1 - (v1 - 42)), F1: (v3 * v3), F2: 42}
        var v11 int32 = -990
        _ = v10
        _ = v11
    } else {
        switch 127 {
        case 32767:
            var v12 bool = (-253 > -2147483648)
            _ = v12
        case 0:
            var v13 uint8 = 2
            _ = v13
        case -1:
            if true {
                v14 := []int{4294967295, 127}
                _ = v14
            } else {
                v3 += (v3 * int16(v1))
            }
        default:
            switch 42 {
            case 427:
                v15 := GenS2{F0: 204, F1: "\n\t", F2: 1}
                _ = v15
            case -9223372036854775808:
                for fi16 := 0; fi16 < 12; fi16++ {
                    v3 += v3
                    var v17 float32 = ((3.14159 / 0.1) + (1.5 * (float32((0.1 + 0.1)) - 1.5)))
                    _ = fi16
                    _ = v17
                }
            case -1:
                fmt.Println(v1, v1)
            }
        }
        var v18 string = ("" + "foo bar")
        _ = v18
    }
    if (3.14159 != (0.1 + (float64((3.14159 * (float32((float64(0.1) / 1e10)) + 1e10))) / (1e10 + float64((1e10 / 0.1)))))) {
        fmt.Println(v1)
    }
    for fi19 := 0; fi19 < 16; fi19++ {
        for fi20 := 0; fi20 < 15; fi20++ {
            var v21 string = ""
            v21 = ""
            _ = fi20
            _ = v21
        }
        fmt.Println(fi19)
        _ = fi19
    }
    var v22 uint16 = 255
    v23 := GenS3{F0: v22, F1: -425}
    if false {
        var v24 int64 = 383
        v25 := []int{32767, 32768, 42}
        v1 = (v1 | v1)
        _ = v24
        _ = v25
    } else {
        switch "a" {
        case "":
            v22 -= v22
        case "foo bar":
            v26 := []int{-9223372036854775808, 1, -465}
            _ = v26
        }
    }
    var v27 uint16 = (v22 * 1)
    var v28 string = "\n\t"
    if false {
        for fi29 := 0; fi29 < 6; fi29++ {
            _ = len(v2)
            var v30 uint8 = 2
            _ = fi29
            _ = v30
        }
        switch -568 {
        case 32768:
            if (!false) {
                for fi31 := 0; fi31 < 16; fi31++ {
                    v32 := GenS4{F0: (v27 + v27), F1: -713, F2: "a"}
                    _ = fi31
                    _ = v32
                }
                v33 := GenS2{F0: 127, F1: (v28 + v28), F2: 0}
                for fi34 := 0; fi34 < 2; fi34++ {
                    switch "" {
                    case "foo bar":
                        if true {
                            var v35 int8 = 100
                            var v36 uint8 = 42
                            _ = v35
                            _ = v36
                        } else {
                            fmt.Println(v27, v22, v22)
                            fmt.Println(v22)
                            fmt.Println(v3, fi34)
                        }
                    case "a":
                        var v37 string = v28
                        _ = v37
                    }
                    if true {
                        var v38 float32 = 0.1
                        fmt.Println(v22)
                        _ = v38
                    } else {
                        v2["a"] = 241
                    }
                    for fi39 := 0; fi39 < 12; fi39++ {
                        switch ((((v28 + v28) + ("hello" + v28)) + v28) + ("foo bar" + v28)) {
                        case "":
                            v1 = (v1 >> ((v1) & 63))
                        case "foo bar":
                            var v40 float64 = float64(-2.25)
                            fmt.Println(v27, v1, fi34)
                            _ = v40
                        }
                        _ = fi39
                    }
                    _ = fi34
                }
                _ = v33
            } else {
                var v41 int8 = 1
                if false {
                    var v42 float32 = -2.25
                    v43 := GenS5{F0: 0}
                    _ = v42
                    _ = v43
                }
                switch 207 {
                case -26:
                    v44 := []int{0, 128, -9223372036854775808, -9223372036854775808}
                    _ = v44
                case -9223372036854775808:
                    v2[v28] = 4294967295
                case 128:
                    for fi45 := 0; fi45 < 19; fi45++ {
                        _ = len(v4)
                        fmt.Println(v22)
                        v22 += 1
                        _ = fi45
                    }
                }
                _ = v41
            }
        case 710:
            v27 -= 32767
        }
        for fi46 := 0; fi46 < 6; fi46++ {
            v28 = ""
            v47 := GenS6{F0: false, F1: 0, F2: "\n\t"}
            _ = fi46
            _ = v47
        }
    } else {
        for fi48 := 0; fi48 < 16; fi48++ {
            v22 -= (v27 | 142)
            var v49 int16 = (v3 + (v3 + (v3 + v3)))
            if (!(42 >= 0)) {
                switch 0 {
                case 4294967295:
                    v50 := GenS5{F0: 9223372036854775807}
                    _ = v50
                case -1:
                    v51 := GenS7{F0: 255}
                    _ = v51
                case 1:
                    if false {
                        v52 := []int{(fi48 % ((fi48) | 1)), 32767, (fi48 % (((fi48 | (fi48 + (fi48 + (fi48 ^ (fi48 ^ (fi48 % ((fi48) | 1)))))))) | 1))}
                        _ = v52
                    }
                }
            } else {
                v4[1] = (((v28 + v28) + v28) + ((((v28 + "") + "a") + ((v28 + v28) + v28)) + (((v28 + (v28 + v28)) + (v28 + v28)) + v28)))
                if (!true) {
                    if true {
                        var v53 int32 = 1
                        _ = v53
                    } else {
                        v28 = "hello"
                    }
                }
            }
            _ = fi48
            _ = v49
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
