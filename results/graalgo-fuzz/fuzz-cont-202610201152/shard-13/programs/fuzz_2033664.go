package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8; F1 uint8; F2 int16 }
type GenS4 struct { F0 int64; F1 int8; F2 bool }
type GenS5 struct { F0 int16; F1 bool }
type GenS6 struct { F0 int16 }
func genFn0() uint64 {
    v7 := GenS1{F0: -9223372036854775808, F1: "\n\t"}
    v8 := []int{}
    var v9 uint32 = 2147483647
    _ = v7
    _ = v8
    _ = v9
    return 146
}
func genFn1(p0 int16, p1 int) uint {
    p0 = p0
    var v10 uint8 = 127
    _ = p0
    _ = p1
    _ = v10
    return 42
}

func main() {
    v1 := []int{32768, -39, 9223372036854775807, 2147483647}
    v2 := make(map[uint32]bool)
    var v3 int32 = 0
    var v4 int8 = -1
    var v5 string = "hello"
    var v6 int64 = 42
    if (128 > 255) {
        for fi11 := 0; fi11 < 7; fi11++ {
            v1 = append(v1, -505)
            _ = fi11
        }
    } else {
        v12 := GenS1{F0: 15, F1: (v5 + ((("foo bar" + v5) + (v5 + "\n\t")) + v5))}
        _ = v12
    }
    switch 879 {
    case -792:
        v6 += v6
    case 128:
        var v13 string = "a"
        _ = v13
    case 32767:
        v5 = "hello"
    }
    var v14 uint16 = 0
    var v15 uint8 = 0
    v3 -= v3
    if ((v3 * (v3 / ((v3) | 1))) > v3) {
        var v16 int8 = v4
        var v17 uint32 = 1
        if (true || (false || true)) {
            for fi18 := 0; fi18 < 9; fi18++ {
                if true {
                    if (1 > 356) {
                        if (v6 > 9223372036854775807) {
                            if (!((!(v3 > v3)) || false)) {
                                v6 -= (v6 | (^(v6)))
                            }
                        } else {
                            v6 = int64(fi18)
                            for fi19 := 0; fi19 < 20; fi19++ {
                                v5 = "\n\t"
                                switch fi18 {
                                case 32768:
                                    v3 = -1
                                case 42:
                                    if (42 == -32768) {
                                        v16 -= 127
                                        v15 += 128
                                    }
                                case 4294967295:
                                    fmt.Println(v5, v3)
                                }
                                switch (fi19 * 1) {
                                case 9223372036854775807:
                                    fmt.Println(fi19)
                                case -880:
                                    fmt.Println(v6)
                                default:
                                    if len(v1) > 0 {
                                        v1[0] = (-(fi18))
                                    }
                                }
                                _ = fi19
                            }
                            if (9223372036854775807 >= (-(fi18))) {
                                var v20 int8 = v16
                                v21 := GenS2{F0: ("foo bar" + v5)}
                                _ = v20
                                _ = v21
                            } else {
                                v22 := GenS2{F0: (v5 + v5)}
                                _ = v22
                            }
                        }
                    } else {
                        v23 := GenS3{F0: (v15 & v15), F1: v15, F2: 1000}
                        var v24 int32 = -883
                        _ = v23
                        _ = v24
                    }
                } else {
                    switch (fi18 + 32767) {
                    case -760:
                        v25 := GenS4{F0: (v6 >> ((902) & 63)), F1: int8(fi18), F2: (42 != (v17 - (^(v17))))}
                        _ = v25
                    default:
                        v26 := GenS2{F0: "foo bar"}
                        _ = v26
                    }
                }
                _ = fi18
            }
        }
        _ = v16
        _ = v17
    }
    for fi27 := 0; fi27 < 15; fi27++ {
        for fi28 := 0; fi28 < 7; fi28++ {
            switch "foo bar" {
            case "a":
                v29 := GenS1{F0: int(v6), F1: "foo bar"}
                _ = v29
            }
            switch "hello" {
            case "hello":
                fmt.Println(fi28)
            default:
                delete(v2, 4294967295)
            }
            v14 += v14
            if false {
                v6 += -892
            }
            _ = fi28
        }
        _ = fi27
    }
    if ((float64(0.1) / (1e10 / ((1.5 * 1e10) - ((0.0 * ((1.5 + -2.25) + (((1.5 * 1e10) - float64(1e10)) + 3.14159))) / 1.5)))) >= ((float64(3.14159) * 1e10) / 3.14159)) {
        switch v5 {
        case "":
            v30 := v2[971]
            _ = v30
        }
        if ((v3 + -1) < (v3 >> ((4294967295) & 63))) {
            fmt.Println(v15, v3, v14)
        } else {
            if false {
                if false {
                    for fi31 := 0; fi31 < 20; fi31++ {
                        var v32 int = (fi31 - fi31)
                        if true {
                            var v33 float64 = (1e10 - (1e10 - 1e10))
                            for fi34 := 0; fi34 < 16; fi34++ {
                                var v35 uint8 = 0
                                _ = fi34
                                _ = v35
                            }
                            _ = v33
                        }
                        _ = fi31
                        _ = v32
                    }
                }
                var v36 float64 = -2.25
                _ = v36
            }
        }
    }
    fmt.Println(v3)
    for fi37 := 0; fi37 < 8; fi37++ {
        if true {
            var v38 int8 = (v4 | int8(v14))
            if (!true) {
                var v39 int8 = (v4 & (^(v4)))
                for fi40 := 0; fi40 < 16; fi40++ {
                    v41 := v2[4294967295]
                    v41 = true
                    v6 += 372
                    _ = fi40
                    _ = v41
                }
                var v42 int64 = (v6 >> ((255) & 63))
                _ = v39
                _ = v42
            }
            _ = v38
        }
        v43 := GenS5{F0: -1, F1: true}
        for fi44 := 0; fi44 < 13; fi44++ {
            v15 -= (v15 + v15)
            _ = fi44
        }
        _ = fi37
        _ = v43
    }
    for fi45 := 0; fi45 < 2; fi45++ {
        if (4294967295 > 425) {
            for fi46 := 0; fi46 < 14; fi46++ {
                if (3.14159 != ((-2.25 * (3.14159 / ((0.1 / -2.25) + 0.1))) + (1.5 * (-2.25 + 1.5)))) {
                    switch "a" {
                    case "a":
                        var v47 int = (-(fi46))
                        _ = v47
                    case "foo bar":
                        v4 += 0
                    case "\n\t":
                        if false {
                            if true {
                                switch ("foo bar" + "a") {
                                case "":
                                    fmt.Println(v6, v6, fi46)
                                }
                                var v48 int8 = (v4 ^ (v4 * (v4 / ((-27) | 1))))
                                _ = v48
                            } else {
                                switch fi46 {
                                case 16:
                                    var v49 float32 = 0.1
                                    _ = v49
                                case -9223372036854775808:
                                    fmt.Println(v15, v3)
                                default:
                                    switch v5 {
                                    case "hello":
                                        fmt.Println(v6, v15, v4)
                                    }
                                }
                                v2[1] = false
                            }
                            fmt.Println(v3, v3)
                            v50 := GenS6{F0: 42}
                            _ = v50
                        }
                    }
                    v4 = (v4 & (v4 % ((42) | 1)))
                    switch (fi45 + fi45) {
                    case 42:
                        switch v5 {
                        case "a":
                            fmt.Println(v14)
                        }
                    case 441:
                        v3 += (v3 ^ (v3 * (^(v3))))
                    case 4294967295:
                        if (0 < -1) {
                            v3 += v3
                        }
                    }
                }
                fmt.Println(v5, fi45, v6)
                switch (fi45 << ((364) & 63)) {
                case -652:
                    _ = len(v1)
                default:
                    var v51 int = (fi45 | int(v14))
                    _ = v51
                }
                _ = fi46
            }
        }
        var v52 string = v5
        fmt.Println(v3)
        _ = fi45
        _ = v52
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
