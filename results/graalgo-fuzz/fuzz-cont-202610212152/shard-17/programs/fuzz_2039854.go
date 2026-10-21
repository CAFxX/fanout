package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64; F1 uint; F2 bool }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int; F1 uint8; F2 uint64 }
type GenS5 struct { F0 string; F1 int16 }
func genFn0(p0 uint64) string {
    p0 += p0
    p0 = p0
    var v4 uint8 = 127
    for fi5 := 0; fi5 < 14; fi5++ {
        fmt.Println(p0, v4)
        v6 := GenS1{F0: false}
        for fi7 := 0; fi7 < 5; fi7++ {
            v8 := []int{}
            var v9 float32 = 0.1
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = v4
    return "\n\t"
}
func genFn1(p0 int32, p1 uint, p2 int32) int32 {
    if false {
        if true {
            v10 := GenS1{F0: (!false)}
            var v11 string = "foo bar"
            p2 -= p0
            _ = v10
            _ = v11
        } else {
            v12 := []int{510, 363, 1, -1}
            switch "hello" {
            case "hello":
                var v13 float32 = 1e10
                _ = v13
            case "\n\t":
                p0 = p0
            case "a":
                v12 = append(v12, 9223372036854775807)
            }
            switch "\n\t" {
            case "hello":
                _ = len(v12)
            case "a":
                switch "" {
                case "":
                    var v14 float64 = 3.14159
                    _ = v14
                }
            }
            _ = v12
        }
    }
    p1 -= (p1 / ((p1) | 1))
    if (!((!(((9223372036854775807 < 790) && (-1 > 9223372036854775807)) || ((p0 / ((2147483647) | 1)) <= p2))) && false)) {
        v15 := GenS2{F0: 804, F1: (p1 + (p1 | p1)), F2: (false || true)}
        var v16 bool = true
        _ = v15
        _ = v16
    } else {
        p2 -= 686
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    var v1 uint8 = 127
    var v2 int32 = -2147483648
    var v3 bool = false
    v17 := GenS2{F0: 375, F1: 65535, F2: v3}
    switch ("\n\t" + ("hello" + "a")) {
    case "":
        fmt.Println(v3)
    default:
        if (-862 <= 42) {
            for fi18 := 0; fi18 < 19; fi18++ {
                switch "hello" {
                case "":
                    for fi19 := 0; fi19 < 18; fi19++ {
                        v1 += v1
                        fmt.Println(v1, v3, fi18)
                        v20 := []int{fi19, fi18, int(v1)}
                        if v3 {
                            if (4294967295 == 255) {
                                v3 = v3
                                v21 := GenS2{F0: 1, F1: 255, F2: (!(((1.5 - 1e10) / (1.5 + (-2.25 + ((((1.5 * ((0.1 - ((0.1 / (3.14159 - 1e10)) / -2.25)) / (1e10 / 0.1))) - (3.14159 + float32(3.14159))) - (0.1 + ((1.5 - 0.1) / (((3.14159 * (-2.25 / -2.25)) / -2.25) + 1.5)))) + 3.14159)))) == (1.5 - float32((1.5 * -2.25)))))}
                                _ = v21
                            }
                        } else {
                            switch "a" {
                            case "":
                                fmt.Println(v2, v3)
                            }
                        }
                        _ = fi19
                        _ = v20
                    }
                default:
                    v2 = (v2 / ((v2) | 1))
                }
                v22 := []int{(fi18 - fi18)}
                for fi23 := 0; fi23 < 3; fi23++ {
                    var v24 float32 = ((0.0 + (3.14159 + float32(1e10))) / (3.14159 / 1e10))
                    fmt.Println(fi23, fi23)
                    _ = fi23
                    _ = v24
                }
                _ = fi18
                _ = v22
            }
            switch ("foo bar" + "\n\t") {
            case "\n\t":
                fmt.Println(v3)
            }
            for fi25 := 0; fi25 < 9; fi25++ {
                v1 -= 42
                var v26 int32 = 327
                _ = fi25
                _ = v26
            }
        } else {
            fmt.Println(v3, v3)
            if v3 {
                v3 = true
                if v3 {
                    v27 := GenS3{F0: 556}
                    var v28 int8 = 127
                    _ = v27
                    _ = v28
                }
                var v29 float64 = (1.5 * 3.14159)
                _ = v29
            }
        }
    }
    if ((1 >= (v2 | (v2 - genFn1((^(v2)), 255, 1)))) && v3) {
        v30 := []int{0, -812, 420, 265}
        switch -2147483648 {
        case 32768:
            var v31 int = 2147483647
            _ = v31
        case 2147483647:
            v32 := GenS1{F0: (986 > 101)}
            _ = v32
        default:
            fmt.Println(v1)
        }
        _ = v30
    } else {
        v2 += genFn1((v2 / (((-(v2))) | 1)), 0, -488)
    }
    v33 := []int{-9223372036854775808}
    for fi34 := 0; fi34 < 6; fi34++ {
        v35 := GenS1{F0: ((v3 || (v1 <= v1)) && v3)}
        _ = fi34
        _ = v35
    }
    switch -1 {
    case 1:
        v3 = v3
    case -2147483648:
        for fi36 := 0; fi36 < 2; fi36++ {
            if false {
                var v37 int16 = -32768
                _ = v37
            }
            _ = fi36
        }
    case 32767:
        v2 = v2
    default:
        fmt.Println(v3)
    }
    v38 := GenS2{F0: 9223372036854775807, F1: 0, F2: false}
    for fi39 := 0; fi39 < 17; fi39++ {
        for fi40 := 0; fi40 < 17; fi40++ {
            v41 := GenS4{F0: fi39, F1: 255, F2: 334}
            v42 := GenS5{F0: "\n\t", F1: 0}
            v43 := GenS1{F0: (-253 <= 1)}
            _ = fi40
            _ = v41
            _ = v42
            _ = v43
        }
        fmt.Println(v2, fi39, v1)
        v2 += (v2 / ((genFn1(v2, 1, -2147483648)) | 1))
        _ = fi39
    }
    _ = len(v33)
    if ((82 != 735) || v3) {
        v44 := GenS3{F0: -9223372036854775808}
        v45 := GenS2{F0: -1, F1: 0, F2: v3}
        _ = v44
        _ = v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v33
    _ = v38
}
