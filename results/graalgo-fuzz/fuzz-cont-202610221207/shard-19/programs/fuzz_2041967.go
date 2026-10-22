package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int16; F1 uint32 }
type GenS5 struct { F0 uint32; F1 uint64 }
type GenS6 struct { F0 int32; F1 int8; F2 string }
type GenS7 struct { F0 int16; F1 uint16; F2 uint8 }
func genFn0(p0 int8, p1 int64) uint {
    if (!((p1 * 42) == (p1 % (((p1 << ((18446744073709551615) & 63))) | 1)))) {
        if false {
            v6 := GenS1{F0: 0, F1: 1, F2: 65535}
            for fi7 := 0; fi7 < 19; fi7++ {
                var v8 uint8 = 48
                if (0.0 < -2.25) {
                    fmt.Println(p0, fi7, p1)
                    v8 -= v8
                }
                _ = fi7
                _ = v8
            }
            _ = v6
        } else {
            var v9 uint64 = 18446744073709551615
            p0 = p0
            var v10 int8 = -20
            _ = v10
            _ = v9
        }
    } else {
        for fi11 := 0; fi11 < 7; fi11++ {
            if true {
                var v12 int = 4294967295
                fmt.Println(p1, p0)
                _ = v12
            } else {
                var v13 float64 = float64(((1e10 + 1.5) * 1.5))
                _ = v13
            }
            _ = fi11
        }
        v14 := []int{891}
        fmt.Println(p0, p1)
        _ = v14
    }
    var v15 uint = 267
    for fi16 := 0; fi16 < 4; fi16++ {
        v17 := GenS1{F0: 255, F1: 926, F2: 32767}
        v18 := []int{-638}
        _ = fi16
        _ = v17
        _ = v18
    }
    v19 := GenS2{F0: 588}
    switch "\n\t" {
    case "foo bar":
        if false {
            for fi20 := 0; fi20 < 15; fi20++ {
                if (0 != 877) {
                    var v21 float32 = float32(1.5)
                    _ = v21
                }
                var v22 int64 = (-(p1))
                _ = fi20
                _ = v22
            }
            var v23 int = int(v15)
            _ = v23
        } else {
            var v24 bool = false
            var v25 float32 = (0.1 / 1.5)
            _ = v24
            _ = v25
        }
    case "a":
        for fi26 := 0; fi26 < 1; fi26++ {
            if (true && true) {
                var v27 float32 = 0.0
                _ = v27
            } else {
                if false {
                    v28 := []int{fi26}
                    switch (fi26 << ((v15) & 63)) {
                    case 588:
                        for fi29 := 0; fi29 < 17; fi29++ {
                            var v30 int32 = 0
                            p1 += (p1 % ((-9223372036854775808) | 1))
                            _ = fi29
                            _ = v30
                        }
                    default:
                        for fi31 := 0; fi31 < 11; fi31++ {
                            _ = len(v28)
                            var v32 int8 = 100
                            v15 += (^(v15))
                            p1 -= p1
                            _ = fi31
                            _ = v32
                        }
                    }
                    if true {
                        var v33 int32 = int32(p0)
                        _ = v33
                    }
                    _ = v28
                } else {
                    switch ("foo bar" + (("\n\t" + "") + "a")) {
                    case "hello":
                        fmt.Println(fi26, p1, fi26)
                    case "foo bar":
                        for fi34 := 0; fi34 < 8; fi34++ {
                            p1 = (p1 - 42)
                            v15 += (v15 & (v15 + (^(v15))))
                            p0 -= p0
                            _ = fi34
                        }
                    default:
                        switch "\n\t" {
                        case "hello":
                            v35 := GenS3{F0: (-(p0))}
                            _ = v35
                        }
                    }
                    for fi36 := 0; fi36 < 20; fi36++ {
                        v37 := GenS2{F0: 32767}
                        _ = fi36
                        _ = v37
                    }
                }
                for fi38 := 0; fi38 < 19; fi38++ {
                    v39 := GenS1{F0: 255, F1: 1, F2: 32767}
                    var v40 int32 = 60
                    p1 = p1
                    _ = fi38
                    _ = v39
                    _ = v40
                }
                v41 := GenS4{F0: -32768, F1: 2147483648}
                _ = v41
            }
            v42 := GenS3{F0: -22}
            fmt.Println(fi26, fi26, v15)
            _ = fi26
            _ = v42
        }
    default:
        for fi43 := 0; fi43 < 4; fi43++ {
            if true {
                if false {
                    v44 := []int{}
                    switch "hello" {
                    case "a":
                        var v45 bool = (true && false)
                        _ = v45
                    case "foo bar":
                        if true {
                            _ = len(v44)
                        }
                    default:
                        for fi46 := 0; fi46 < 1; fi46++ {
                            var v47 string = "hello"
                            var v48 int8 = (p0 - (p0 * 42))
                            _ = fi46
                            _ = v47
                            _ = v48
                        }
                    }
                    _ = v44
                } else {
                    v49 := []int{fi43, (fi43 | (fi43 + (-(fi43)))), fi43}
                    _ = v49
                }
                v50 := []int{}
                switch ("hello" + "\n\t") {
                case "a":
                    var v51 bool = (!true)
                    _ = v51
                }
                _ = v50
            }
            p0 = p0
            _ = fi43
        }
    }
    _ = p0
    _ = p1
    _ = v15
    _ = v19
    return (v15 / ((314) | 1))
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int32 = -59
    var v3 uint64 = 0
    var v4 uint16 = 42
    var v5 uint64 = v3
    var v52 uint8 = 248
    switch "" {
    case "foo bar":
        v53 := GenS3{F0: 42}
        _ = v53
    case "":
        v54 := []int{-9223372036854775808, 803, 4294967295}
        _ = v54
    default:
        v4 -= v4
    }
    v55 := GenS5{F0: 2147483648, F1: v5}
    if (0 >= 654) {
        switch 127 {
        case -1:
            fmt.Println(v5)
        case 128:
            for fi56 := 0; fi56 < 4; fi56++ {
                var v57 int8 = -1
                v58 := []int{(fi56 ^ fi56), (fi56 & -289), (fi56 * (fi56 & (fi56 - (fi56 - (fi56 * (fi56 >> ((65535) & 63)))))))}
                var v59 string = ("foo bar" + "a")
                _ = fi56
                _ = v57
                _ = v58
                _ = v59
            }
        case 32767:
            if (672 >= 42) {
                fmt.Println(v4)
            }
        default:
            switch "hello" {
            case "foo bar":
                switch "foo bar" {
                case "":
                    v60 := GenS6{F0: 0, F1: -42, F2: ("foo bar" + "")}
                    _ = v60
                }
            case "a":
                for fi61 := 0; fi61 < 17; fi61++ {
                    v2 += (v2 | (v2 >> ((2) & 63)))
                    _ = fi61
                }
            }
        }
        v62 := GenS1{F0: (v52 * (v52 | (v52 + v52))), F1: 4294967295, F2: (v4 + v4)}
        _ = v62
    } else {
        switch (("foo bar" + "a") + "a") {
        case "foo bar":
            v5 = v5
        }
        fmt.Println(v52, v2, v4)
        var v63 float64 = 0.0
        _ = v63
    }
    fmt.Println(v5, v5)
    if (("" + ("foo bar" + "")) > "a") {
        var v64 int16 = 32767
        v65 := []int{}
        _ = v64
        _ = v65
    }
    if false {
        v66 := GenS4{F0: -477, F1: 363}
        _ = v66
    } else {
        v5 -= v3
        fmt.Println(v4)
        v67 := GenS7{F0: 1000, F1: (v4 | 1), F2: (^(v52))}
        _ = v67
    }
    for fi68 := 0; fi68 < 19; fi68++ {
        v5 = 18446744073709551615
        _ = fi68
    }
    var v69 uint64 = v3
    v4 = (v4 / (((^(v4))) | 1))
    fmt.Println(v5)
    fmt.Println(v69, v5, v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v55
    _ = v69
}
