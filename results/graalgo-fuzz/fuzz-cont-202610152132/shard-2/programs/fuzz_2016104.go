package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 uint64; F1 uint32 }
type GenS3 struct { F0 bool; F1 uint64; F2 uint }
type GenS4 struct { F0 string; F1 int32 }
type GenS5 struct { F0 int64; F1 uint16; F2 int }
func genFn0() string {
    var v6 int16 = 1000
    v6 -= (v6 / (((v6 | (-(v6)))) | 1))
    v7 := GenS1{F0: 1, F1: 643}
    _ = v6
    _ = v7
    return "hello"
}
func genFn1(p0 int32, p1 int) int8 {
    var v8 float64 = ((-2.25 + float64(1.5)) / (3.14159 - (-2.25 + 0.1)))
    switch "foo bar" {
    case "\n\t":
        if true {
            v8 = (v8 + ((1e10 * -2.25) * -2.25))
        } else {
            switch 520 {
            case 9223372036854775807:
                v9 := []int{p1, (p1 | p1), p1, (^(p1))}
                _ = v9
            default:
                v10 := []int{(p1 & p1)}
                _ = v10
            }
            v8 = float64(1e10)
        }
    case "":
        fmt.Println(p1, p1, p0)
    default:
        v11 := []int{p1, 250}
        _ = v11
    }
    var v12 uint64 = 4294967295
    v13 := GenS1{F0: 2, F1: 42}
    if false {
        v14 := GenS2{F0: v12, F1: 65535}
        v15 := []int{(p1 % (((p1 / (((p1 / ((101) | 1))) | 1))) | 1)), p1}
        var v16 bool = true
        _ = v14
        _ = v15
        _ = v16
    } else {
        for fi17 := 0; fi17 < 18; fi17++ {
            var v18 int16 = 240
            if false {
                switch p1 {
                case 10:
                    p1 = fi17
                case 32767:
                    switch p1 {
                    case -1:
                        switch "\n\t" {
                        case "":
                            for fi19 := 0; fi19 < 2; fi19++ {
                                var v20 float32 = float32(-2.25)
                                if (0 < -9223372036854775808) {
                                    v8 = (((v8 * 1.5) * v8) - v8)
                                    var v21 uint16 = 255
                                    _ = v21
                                }
                                _ = fi19
                                _ = v20
                            }
                        case "\n\t":
                            switch "a" {
                            case "hello":
                                p0 = (p0 ^ (p0 * 42))
                            case "foo bar":
                                v8 -= -2.25
                            case "a":
                                v22 := []int{}
                                _ = v22
                            default:
                                v23 := []int{p1, -9223372036854775808}
                                _ = v23
                            }
                        }
                    case 0:
                        for fi24 := 0; fi24 < 11; fi24++ {
                            fmt.Println(fi24, fi17, v8)
                            p1 -= fi24
                            _ = fi24
                        }
                    case -866:
                        v25 := GenS2{F0: v12, F1: 1}
                        _ = v25
                    default:
                        var v26 int8 = -28
                        _ = v26
                    }
                case 0:
                    v27 := GenS3{F0: (p0 <= p0), F1: (v12 | (^(v12))), F2: 0}
                    _ = v27
                }
                for fi28 := 0; fi28 < 2; fi28++ {
                    if false {
                        var v29 int32 = p0
                        v30 := GenS4{F0: ((("hello" + "foo bar") + "") + "foo bar"), F1: p0}
                        _ = v29
                        _ = v30
                    } else {
                        p1 = (fi17 << ((42) & 63))
                        var v31 uint32 = 2147483648
                        _ = v31
                    }
                    var v32 int = 32768
                    _ = fi28
                    _ = v32
                }
            } else {
                var v33 int32 = p0
                _ = v33
            }
            p0 = (p0 | p0)
            _ = fi17
            _ = v18
        }
        v34 := GenS5{F0: -1, F1: 158, F2: 32768}
        _ = v34
    }
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    _ = v8
    return -128
}

func main() {
    var v1 bool = (65535 > 42)
    v2 := make(map[uint32]int)
    var v3 uint32 = 108
    v4 := make(map[uint32]int)
    var v5 int32 = 143
    v35 := v4[v3]
    fmt.Println(v3)
    for fi36 := 0; fi36 < 4; fi36++ {
        switch v35 {
        case -465:
            var v37 bool = ("foo bar" < "foo bar")
            _ = v37
        default:
            if ((1 <= 0) || false) {
                v4[v3] = fi36
                v38 := GenS4{F0: ("foo bar" + "hello"), F1: v5}
                var v39 int8 = -1
                _ = v38
                _ = v39
            }
        }
        v1 = ((v5 / (((v5 >> ((0) & 63))) | 1)) >= v5)
        _ = fi36
    }
    v40 := GenS3{F0: v1, F1: 981, F2: 251}
    v41 := GenS4{F0: ("hello" + ("a" + "foo bar")), F1: 42}
    v3 = 0
    var v42 uint64 = 672
    for fi43 := 0; fi43 < 15; fi43++ {
        var v44 string = "foo bar"
        v44 = "\n\t"
        v45 := []int{(fi43 + v35), (fi43 >> ((2) & 63)), (fi43 ^ 42)}
        _ = fi43
        _ = v44
        _ = v45
    }
    switch (v35 + v35) {
    case 2147483647:
        switch v35 {
        case 0:
            v46 := GenS4{F0: "", F1: (v5 % ((-688) | 1))}
            _ = v46
        }
    case -409:
        v5 += int32(v42)
    case -2147483648:
        if v1 {
            fmt.Println(v42)
        }
    default:
        fmt.Println(v5, v42)
    }
    v35 += v35
    var v47 int32 = 2147483647
    switch (-(v35)) {
    case 1:
        v48 := []int{v35, v35}
        _ = v48
    }
    switch (v35 & 9223372036854775807) {
    case 9223372036854775807:
        if (-851 != 9223372036854775807) {
            fmt.Println(v35, v42)
            if (uint64(v3) <= 42) {
                if ((-2.25 / (0.1 / 1e10)) == (3.14159 * 0.0)) {
                    switch "a" {
                    case "\n\t":
                        for fi49 := 0; fi49 < 16; fi49++ {
                            fmt.Println(v47, v42, v5)
                            for fi50 := 0; fi50 < 18; fi50++ {
                                v47 = v47
                                _ = fi50
                            }
                            switch ((("hello" + "a") + "foo bar") + "a") {
                            case "a":
                                for fi51 := 0; fi51 < 1; fi51++ {
                                    v3 = v3
                                    v35 = fi51
                                    _ = fi51
                                }
                            case "hello":
                                for fi52 := 0; fi52 < 2; fi52++ {
                                    var v53 int64 = int64(v42)
                                    v42 = (v42 | (v42 * v42))
                                    _ = fi52
                                    _ = v53
                                }
                            case "foo bar":
                                var v54 int8 = -112
                                _ = v54
                            }
                            _ = fi49
                        }
                    default:
                        switch ("hello" + "\n\t") {
                        case "foo bar":
                            switch (v35 >> ((2) & 63)) {
                            case 1:
                                var v55 uint = 482
                                _ = v55
                            case 32767:
                                v47 -= v47
                            }
                        }
                    }
                    var v56 int8 = 42
                    _ = v56
                }
                var v57 int = -1
                v57 = v57
                _ = v57
            }
        }
    case 679:
        v5 -= -1
    case -9223372036854775808:
        for fi58 := 0; fi58 < 20; fi58++ {
            v3 = v3
            var v59 int64 = -900
            _ = fi58
            _ = v59
        }
    }
    for fi60 := 0; fi60 < 7; fi60++ {
        var v61 bool = ((-(v5)) == (-(v47)))
        var v62 int16 = 1000
        _ = fi60
        _ = v61
        _ = v62
    }
    for fi63 := 0; fi63 < 1; fi63++ {
        switch "a" {
        case "\n\t":
            v3 = (v3 ^ (v3 | v3))
        case "a":
            for fi64 := 0; fi64 < 12; fi64++ {
                v65 := []int{fi63}
                _ = fi64
                _ = v65
            }
        }
        var v66 string = "foo bar"
        var v67 int = -1
        _ = fi63
        _ = v66
        _ = v67
    }
    for fi68 := 0; fi68 < 13; fi68++ {
        v69 := v4[v3]
        _ = fi68
        _ = v69
    }
    v1 = (-128 < -35)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v47
    _ = v5
}
