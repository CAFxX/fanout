package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 int32; F1 uint16; F2 uint32 }
type GenS4 struct { F0 int; F1 int8; F2 int16 }
type GenS5 struct { F0 string; F1 uint32; F2 int16 }
type GenS6 struct { F0 int16; F1 uint; F2 string }
func genFn0() uint8 {
    if (!false) {
        var v5 int32 = 0
        if ((18446744073709551615 != 126) && (true && true)) {
            v6 := []int{-2147483648, -9223372036854775808, int(v5), 32767}
            fmt.Println(v5)
            _ = v6
        }
        switch 9223372036854775807 {
        case 9223372036854775807:
            v7 := GenS1{F0: (("hello" <= ("foo bar" + (("" + "") + ("foo bar" + "foo bar")))) && ("" < ("foo bar" + (("\n\t" + "") + "\n\t")))), F1: 4294967295}
            _ = v7
        case 300:
            v8 := GenS2{F0: ("foo bar" + ("a" + "")), F1: 32767}
            _ = v8
        case 794:
            if true {
                fmt.Println(v5)
                v5 = -353
            }
        }
        _ = v5
    } else {
        var v9 float32 = 0.1
        v9 -= ((v9 * v9) + (v9 / v9))
        for fi10 := 0; fi10 < 9; fi10++ {
            for fi11 := 0; fi11 < 7; fi11++ {
                v9 = float32(-2.25)
                v12 := GenS2{F0: ("" + (("hello" + "a") + "")), F1: (fi11 | (fi11 >> ((2) & 63)))}
                _ = fi11
                _ = v12
            }
            switch 0 {
            case 32767:
                for fi13 := 0; fi13 < 10; fi13++ {
                    var v14 uint8 = 223
                    var v15 uint8 = 220
                    v16 := []int{fi13}
                    _ = fi13
                    _ = v14
                    _ = v15
                    _ = v16
                }
            case 4294967295:
                for fi17 := 0; fi17 < 10; fi17++ {
                    var v18 bool = false
                    _ = fi17
                    _ = v18
                }
            case 42:
                fmt.Println(fi10, fi10)
            default:
                switch "\n\t" {
                case "foo bar":
                    for fi19 := 0; fi19 < 20; fi19++ {
                        var v20 int16 = -1
                        _ = fi19
                        _ = v20
                    }
                case "":
                    v21 := []int{}
                    _ = v21
                default:
                    v22 := []int{(fi10 & 42), 789, (fi10 & (fi10 | fi10))}
                    _ = v22
                }
            }
            _ = fi10
        }
        _ = v9
    }
    var v23 string = ""
    var v24 uint32 = 597
    _ = v23
    _ = v24
    return 2
}
func genFn1(p0 string, p1 int, p2 int32) int16 {
    p2 -= p2
    fmt.Println(p0, p2, p2)
    _ = p0
    _ = p1
    _ = p2
    return -1
}

func main() {
    var v1 uint16 = 303
    var v2 uint = 656
    var v3 bool = (48 <= 42)
    v4 := make(map[int]uint32)
    fmt.Println(v1)
    if (true && false) {
        v4[191] = 0
    } else {
        var v25 bool = ((-128 == -85) && true)
        _ = v25
    }
    v1 -= 0
    fmt.Println(v2)
    v26 := v4[32767]
    switch -548 {
    case -81:
        var v27 uint = (v2 % ((18446744073709551615) | 1))
        _ = v27
    default:
        if ((1.5 + -2.25) == (-2.25 + 3.14159)) {
            v2 += (v2 * v2)
        } else {
            var v28 float64 = (float64(1.5) + 0.0)
            v29 := []int{9223372036854775807, 128}
            v30 := GenS3{F0: 811, F1: v1, F2: uint32(v1)}
            _ = v28
            _ = v29
            _ = v30
        }
    }
    if (v3 && (float32(0.1) == -2.25)) {
        switch 32767 {
        case 9223372036854775807:
            fmt.Println(v26)
        default:
            fmt.Println(v3)
        }
        fmt.Println(v1, v26, v2)
    } else {
        var v31 int32 = 0
        _ = len(v4)
        var v32 int32 = v31
        _ = v31
        _ = v32
    }
    for fi33 := 0; fi33 < 5; fi33++ {
        if ((^(fi33)) > -2147483648) {
            if v3 {
                var v34 float32 = 0.0
                switch "\n\t" {
                case "hello":
                    for fi35 := 0; fi35 < 11; fi35++ {
                        fmt.Println(fi35)
                        v36 := GenS3{F0: -258, F1: (v1 - 489), F2: (^(v26))}
                        _ = fi35
                        _ = v36
                    }
                default:
                    if false {
                        v37 := GenS2{F0: "", F1: 4294967295}
                        fmt.Println(fi33, fi33)
                        fmt.Println(v26)
                        _ = v37
                    } else {
                        v38 := GenS2{F0: ("" + "foo bar"), F1: -868}
                        var v39 float32 = (-2.25 * (0.1 / (v34 - (v34 / (v34 * v34)))))
                        _ = v38
                        _ = v39
                    }
                }
                _ = len(v4)
                _ = v34
            } else {
                switch fi33 {
                case 42:
                    var v40 uint16 = (^(v1))
                    _ = v40
                case 127:
                    for fi41 := 0; fi41 < 13; fi41++ {
                        var v42 string = ("hello" + ("" + ("" + "hello")))
                        fmt.Println(v26, fi33, v3)
                        var v43 int32 = 42
                        _ = fi41
                        _ = v42
                        _ = v43
                    }
                }
                fmt.Println(v3)
                var v44 int = (fi33 >> (((v2 * v2)) & 63))
                _ = v44
            }
        } else {
            v45 := GenS4{F0: fi33, F1: 100, F2: -32768}
            v4[(fi33 | (fi33 - 127))] = v26
            for fi46 := 0; fi46 < 3; fi46++ {
                v2 = v2
                _ = fi46
            }
            _ = v45
        }
        _ = fi33
    }
    fmt.Println(v2, v1)
    for fi47 := 0; fi47 < 15; fi47++ {
        v26 = 45
        v48 := []int{1, fi47}
        for fi49 := 0; fi49 < 14; fi49++ {
            var v50 int64 = 9223372036854775807
            fmt.Println(v2, v2, fi47)
            _ = fi49
            _ = v50
        }
        _ = fi47
        _ = v48
    }
    fmt.Println(v3, v26)
    for fi51 := 0; fi51 < 6; fi51++ {
        v1 = (v1 - v1)
        if v3 {
            v2 -= (v2 ^ (v2 * v2))
        } else {
            var v52 float64 = 1.5
            if (9223372036854775807 <= 1) {
                v26 += 951
            } else {
                v3 = true
            }
            _ = v52
        }
        _ = fi51
    }
    var v53 float64 = 0.1
    for fi54 := 0; fi54 < 2; fi54++ {
        v55 := GenS5{F0: "\n\t", F1: 40, F2: 32767}
        if false {
            switch "a" {
            case "":
                v56 := GenS3{F0: 42, F1: 32767, F2: (v26 + v26)}
                _ = v56
            case "\n\t":
                v26 += v26
            }
        }
        v57 := GenS2{F0: "", F1: (fi54 % (((fi54 % (((fi54 & fi54)) | 1))) | 1))}
        _ = fi54
        _ = v55
        _ = v57
    }
    var v58 uint64 = 42
    for fi59 := 0; fi59 < 19; fi59++ {
        for fi60 := 0; fi60 < 7; fi60++ {
            fmt.Println(v58, v2)
            _ = fi60
        }
        v61 := GenS6{F0: 624, F1: v2, F2: ("foo bar" + (("foo bar" + "foo bar") + ""))}
        _ = fi59
        _ = v61
    }
    switch -2147483648 {
    case 128:
        v58 += 4294967295
    case 639:
        v2 = (v2 ^ 1)
    default:
        var v62 uint64 = v58
        _ = v62
    }
    var v63 float64 = v53
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v53:", v53)
    fmt.Println("v58:", v58)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v53
    _ = v58
    _ = v63
}
