package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 string; F1 uint64 }
type GenS3 struct { F0 uint16; F1 int16; F2 uint16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint64; F1 uint; F2 uint32 }
type GenS6 struct { F0 uint8; F1 uint16 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint; F1 uint32 }
type GenS9 struct { F0 uint8; F1 bool }
func genFn0(p0 uint, p1 int32) uint8 {
    p1 -= (p1 & p1)
    var v4 bool = (!false)
    for fi5 := 0; fi5 < 19; fi5++ {
        var v6 int8 = 127
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 18; fi7++ {
        if v4 {
            v8 := []int{(-(fi7)), fi7, fi7, (fi7 | -1)}
            p1 -= -2147483648
            v4 = v4
            _ = v8
        } else {
            v9 := []int{(-(fi7)), 524, 96, fi7}
            p1 -= (^(p1))
            for fi10 := 0; fi10 < 11; fi10++ {
                if ((1 == 100) || (42 < 931)) {
                    var v11 float64 = float64(0.1)
                    var v12 string = "hello"
                    _ = v11
                    _ = v12
                } else {
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                    var v13 int16 = -343
                    var v14 uint8 = 0
                    _ = v13
                    _ = v14
                }
                _ = fi10
            }
            _ = v9
        }
        _ = fi7
    }
    if v4 {
        var v15 uint64 = 685
        switch -2147483648 {
        case -613:
            v15 = (v15 * 724)
        case 1:
            var v16 uint = p0
            _ = v16
        default:
            switch 1 {
            case 4294967295:
                p1 += (p1 * 722)
            case 1:
                v17 := []int{-277, 4294967295}
                _ = v17
            default:
                p0 += p0
            }
        }
        _ = v15
    }
    _ = p0
    _ = p1
    _ = v4
    return 83
}
func genFn1(p0 int8, p1 int8) int {
    if false {
        v18 := GenS1{F0: 277, F1: p1}
        _ = v18
    }
    if false {
        fmt.Println(p0, p0)
    } else {
        if (0 > -1) {
            if ((("\n\t" + "\n\t") + (("" + "hello") + "\n\t")) >= (("" + "foo bar") + "hello")) {
                fmt.Println(p1)
                p0 = p1
            } else {
                p0 += (p1 | -128)
                for fi19 := 0; fi19 < 11; fi19++ {
                    var v20 int8 = -90
                    v21 := GenS2{F0: ("a" + ((("" + "a") + "") + "")), F1: 2}
                    _ = fi19
                    _ = v20
                    _ = v21
                }
            }
        }
        if false {
            var v22 int8 = (-(p0))
            v22 = (p0 / (((v22 - (p1 - (p0 + 100)))) | 1))
            fmt.Println(p1, v22, p1)
            _ = v22
        }
        fmt.Println(p0, p1)
    }
    v23 := []int{-558, -1}
    _ = p0
    _ = p1
    _ = v23
    return 311
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint = 1
    var v3 int = 643
    v24 := GenS3{F0: 32767, F1: 42, F2: 177}
    v3 -= (v3 * v3)
    for fi25 := 0; fi25 < 20; fi25++ {
        v3 = 451
        _ = fi25
    }
    v26 := GenS2{F0: "a", F1: 18446744073709551615}
    v3 = (v3 % ((-324) | 1))
    fmt.Println(v3, v2)
    var v27 float64 = (1.5 * 3.14159)
    switch "" {
    case "a":
        for fi28 := 0; fi28 < 10; fi28++ {
            fmt.Println(v2, v2, fi28)
            v27 = 3.14159
            var v29 bool = (16 > 416)
            _ = fi28
            _ = v29
        }
    }
    v3 = v3
    for fi30 := 0; fi30 < 5; fi30++ {
        for fi31 := 0; fi31 < 8; fi31++ {
            var v32 bool = (-9223372036854775808 < -7)
            if true {
                var v33 int = (v3 * (-(fi30)))
                var v34 string = ("" + "\n\t")
                if v32 {
                    v3 -= 2147483647
                    v27 = v27
                    v2 = 4294967295
                    var v35 uint8 = 2
                    fmt.Println(v27, v27, v2)
                    v34 = v34
                    _ = v35
                }
                _ = v33
                _ = v34
            } else {
                v3 = fi31
                fmt.Println(v2, fi30)
                v36 := GenS1{F0: 2, F1: 127}
                _ = v36
            }
            var v37 uint8 = 70
            var v38 uint64 = 0
            _ = fi31
            _ = v32
            _ = v37
            _ = v38
        }
        _ = fi30
    }
    for fi39 := 0; fi39 < 9; fi39++ {
        var v40 int64 = 42
        fmt.Println(v2, v27, v2)
        v41 := GenS4{F0: 2147483647}
        _ = fi39
        _ = v40
        _ = v41
    }
    v42 := GenS5{F0: 42, F1: v2, F2: 2147483647}
    var v43 uint32 = 2147483647
    var v44 uint16 = 42
    if ((!false) || (float32(1.5) < 3.14159)) {
        if (2147483647 > (v43 >> ((v2) & 63))) {
            v45 := GenS6{F0: 255, F1: (v44 + v44)}
            if ((!false) || false) {
                for fi46 := 0; fi46 < 18; fi46++ {
                    v47 := GenS7{F0: (v44 | v44)}
                    _ = fi46
                    _ = v47
                }
                v27 -= 0.1
                var v48 int32 = int32(v43)
                _ = v48
            } else {
                var v49 float64 = (v27 - (float64((1.5 - 0.0)) / -2.25))
                v44 += (^(v44))
                var v50 uint8 = 255
                _ = v49
                _ = v50
            }
            for fi51 := 0; fi51 < 16; fi51++ {
                if ((v44 + (v44 >> ((v2) & 63))) != (v44 / ((65535) | 1))) {
                    fmt.Println(v2, v43, v27)
                    fmt.Println(v44, v3)
                }
                _ = fi51
            }
            _ = v45
        } else {
            v52 := GenS8{F0: (v2 ^ 2), F1: (v43 + v43)}
            if (!false) {
                v53 := GenS9{F0: 128, F1: true}
                _ = v53
            } else {
                v43 -= v43
                for fi54 := 0; fi54 < 19; fi54++ {
                    var v55 int8 = 1
                    _ = fi54
                    _ = v55
                }
            }
            _ = v52
        }
    } else {
        var v56 uint = (v2 + v2)
        _ = v56
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v24
    _ = v26
    _ = v27
    _ = v3
    _ = v42
    _ = v43
    _ = v44
}
