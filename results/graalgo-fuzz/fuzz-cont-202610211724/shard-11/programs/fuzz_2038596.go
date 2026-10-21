package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 string; F1 int8; F2 uint64 }
type GenS3 struct { F0 uint16; F1 uint8 }
type GenS4 struct { F0 int; F1 uint16 }
func genFn0(p0 uint16, p1 bool, p2 int64) bool {
    var v7 int32 = -2147483648
    var v8 int64 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return (false && (((v7 | (v7 - v7)) != (v7 >> ((4294967295) & 63))) && ((920 < (v8 * -1)) && (!false))))
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]bool)
    v3 := make(map[int]int)
    var v4 uint32 = 2147483648
    v5 := make(map[string]int)
    var v6 int8 = 127
    if ((v4 << ((213) & 63)) != 2147483647) {
        v9 := []int{0, -2147483648, 0, 9223372036854775807}
        var v10 string = ("foo bar" + "foo bar")
        v6 = (-(v6))
        _ = v10
        _ = v9
    } else {
        v11 := GenS1{F0: 0, F1: true}
        var v12 int = -9223372036854775808
        _ = v11
        _ = v12
    }
    fmt.Println(v6)
    var v13 float64 = 1e10
    for fi14 := 0; fi14 < 2; fi14++ {
        switch "hello" {
        case "foo bar":
            v15 := GenS2{F0: "foo bar", F1: 1, F2: 0}
            _ = v15
        default:
            v16 := GenS2{F0: "a", F1: int8(v4), F2: 42}
            _ = v16
        }
        v17 := []int{(fi14 - int(v4)), (fi14 / (((fi14 / ((-110) | 1))) | 1)), 32768}
        _ = fi14
        _ = v17
    }
    v13 = float64(((float32(v13) * 0.0) * (3.14159 / 3.14159)))
    if false {
        fmt.Println(v13, v13)
        if (2 >= 65535) {
            v4 = (v4 % ((v4) | 1))
            v18 := v3[1]
            for fi19 := 0; fi19 < 20; fi19++ {
                var v20 string = (("a" + "a") + ("hello" + ""))
                _ = fi19
                _ = v20
            }
            _ = v18
        } else {
            v6 -= (v6 | v6)
        }
        switch "hello" {
        case "":
            switch "hello" {
            case "hello":
                for fi21 := 0; fi21 < 11; fi21++ {
                    v22 := GenS1{F0: 2, F1: (-1 >= 42)}
                    if false {
                        for fi23 := 0; fi23 < 16; fi23++ {
                            var v24 uint32 = 2147483647
                            v4 -= 832
                            _ = fi23
                            _ = v24
                        }
                    }
                    v25 := GenS3{F0: 0, F1: 30}
                    _ = fi21
                    _ = v22
                    _ = v25
                }
            case "foo bar":
                switch -68 {
                case 32767:
                    var v26 float64 = (-2.25 - v13)
                    _ = v26
                case 794:
                    v27 := GenS3{F0: 1, F1: 42}
                    _ = v27
                case 42:
                    var v28 uint = 927
                    _ = v28
                default:
                    for fi29 := 0; fi29 < 9; fi29++ {
                        v30 := []int{fi29, 128, 1}
                        _ = fi29
                        _ = v30
                    }
                }
            case "a":
                v4 = v4
            default:
                v31 := GenS3{F0: 255, F1: 0}
                _ = v31
            }
        default:
            v6 = (v6 & -99)
        }
    } else {
        var v32 uint8 = 184
        var v33 string = ("foo bar" + "")
        _ = v32
        _ = v33
    }
    switch "a" {
    case "\n\t":
        switch 1 {
        case -1:
            v34 := []int{989, 356, 630, 42}
            _ = v34
        default:
            if (true && false) {
                switch 1 {
                case -243:
                    for fi35 := 0; fi35 < 15; fi35++ {
                        v36 := []int{}
                        var v37 int16 = -1
                        _ = fi35
                        _ = v36
                        _ = v37
                    }
                case 2147483647:
                    v38 := GenS2{F0: ("" + ("a" + "foo bar")), F1: (v6 * (v6 ^ (v6 % ((-1) | 1)))), F2: 1}
                    _ = v38
                case -506:
                    for fi39 := 0; fi39 < 4; fi39++ {
                        for fi40 := 0; fi40 < 2; fi40++ {
                            v6 -= (^(v6))
                            _ = fi40
                        }
                        _ = fi39
                    }
                }
                v4 -= 2
            }
        }
    case "":
        switch 2147483647 {
        case 128:
            v41 := GenS1{F0: 18446744073709551615, F1: (2 < 32767)}
            _ = v41
        case -195:
            if ((((false && (!false)) || ((false && false) && ((!(1.5 >= float32(float64(0.1)))) && true))) && ("foo bar" != "a")) && false) {
                v42 := []int{}
                var v43 uint8 = 128
                _ = v42
                _ = v43
            } else {
                var v44 uint32 = v4
                _ = v44
            }
        case 0:
            v45 := []int{}
            _ = v45
        }
    }
    v6 = (v6 << ((2) & 63))
    var v46 int = -9223372036854775808
    v6 = v6
    v47 := v2[0]
    v4 += 688
    switch (((("foo bar" + ((("\n\t" + ((("a" + "hello") + "") + ("hello" + "hello"))) + ("" + "")) + "\n\t")) + "\n\t") + "hello") + ("" + "")) {
    case "":
        delete(v2, v4)
    case "a":
        if v47 {
            switch "\n\t" {
            case "foo bar":
                if v47 {
                    v6 = -117
                    if (640 != 224) {
                        v4 = v4
                    }
                }
            case "a":
                switch (v46 * -895) {
                case -247:
                    v48 := GenS3{F0: 35, F1: 1}
                    _ = v48
                case -943:
                    if v47 {
                        for fi49 := 0; fi49 < 8; fi49++ {
                            for fi50 := 0; fi50 < 14; fi50++ {
                                switch 32768 {
                                case 380:
                                    v6 -= (v6 + v6)
                                default:
                                    var v51 int64 = -9223372036854775808
                                    _ = v51
                                }
                                v52 := v1[(v4 * v4)]
                                var v53 string = "\n\t"
                                _ = fi50
                                _ = v52
                                _ = v53
                            }
                            _ = fi49
                        }
                        var v54 float64 = 1.5
                        _ = v54
                    }
                }
            case "hello":
                v55 := GenS1{F0: 205, F1: (!(v13 != 1e10))}
                _ = v55
            }
            fmt.Println(v13)
        }
    case "foo bar":
        if true {
            for fi56 := 0; fi56 < 13; fi56++ {
                v4 -= 551
                var v57 uint8 = 1
                var v58 uint = 255
                _ = fi56
                _ = v57
                _ = v58
            }
        }
    default:
        switch (^(v46)) {
        case 32768:
            v13 = (float64(3.14159) + 0.0)
        case 42:
            for fi59 := 0; fi59 < 2; fi59++ {
                var v60 uint = 44
                v61 := GenS1{F0: 522, F1: (2 != v4)}
                _ = fi59
                _ = v60
                _ = v61
            }
        case -256:
            var v62 int64 = -986
            _ = v62
        }
    }
    var v63 uint16 = 42
    var v64 uint16 = v63
    v46 -= 42
    v65 := GenS4{F0: 42, F1: v63}
    var v66 float32 = 1e10
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v5
    _ = v6
    _ = v63
    _ = v64
    _ = v65
    _ = v66
}
