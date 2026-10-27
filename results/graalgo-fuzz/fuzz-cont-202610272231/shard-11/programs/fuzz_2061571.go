package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 uint64; F1 uint32; F2 uint }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint8 }

func main() {
    var v1 int = 1
    var v2 int = v1
    var v3 float32 = -2.25
    var v4 float64 = (float64(v3) + (float64(-2.25) * (1e10 / ((-2.25 - 1e10) + -2.25))))
    v5 := make(map[int]int)
    var v6 int = 518
    v7 := GenS1{F0: 4294967295, F1: 82}
    v4 = 0.1
    if true {
        var v8 uint32 = 975
        v9 := GenS2{F0: 1, F1: (v8 | (v8 * (v8 % ((v8) | 1)))), F2: 1}
        _ = v8
        _ = v9
    }
    var v10 float32 = 0.1
    v6 = 826
    switch "a" {
    case "foo bar":
        var v11 float32 = (-2.25 + float32(0.1))
        _ = v11
    case "a":
        switch ("hello" + ((("foo bar" + ("hello" + ("foo bar" + "a"))) + ("a" + "hello")) + "a")) {
        case "\n\t":
            for fi12 := 0; fi12 < 12; fi12++ {
                v13 := GenS1{F0: 65535, F1: 65535}
                if false {
                    delete(v5, 4294967295)
                } else {
                    var v14 int64 = 259
                    switch "" {
                    case "":
                        if (true || false) {
                            if (42 <= -128) {
                                v3 += (-2.25 - float32((v4 - v4)))
                            }
                            var v15 int = v1
                            for fi16 := 0; fi16 < 20; fi16++ {
                                v4 -= 3.14159
                                _ = fi16
                            }
                            _ = v15
                        }
                    case "\n\t":
                        var v17 uint64 = 813
                        _ = v17
                    case "hello":
                        v5[v1] = (v2 >> ((18446744073709551615) & 63))
                    }
                    _ = v14
                }
                _ = fi12
                _ = v13
            }
        case "foo bar":
            var v18 float32 = (v10 / v10)
            _ = v18
        case "hello":
            if (!(1 != 18446744073709551615)) {
                var v19 bool = false
                _ = v19
            }
        default:
            var v20 float32 = v3
            _ = v20
        }
    case "":
        v6 = (v2 / ((v6) | 1))
    default:
        for fi21 := 0; fi21 < 10; fi21++ {
            for fi22 := 0; fi22 < 17; fi22++ {
                if (42 < 127) {
                    var v23 string = "\n\t"
                    var v24 int = (fi21 | v2)
                    _ = v23
                    _ = v24
                } else {
                    v25 := GenS3{F0: 4294967295}
                    v6 = (v2 << ((0) & 63))
                    if ((v4 <= 1.5) && true) {
                        var v26 uint64 = 266
                        v6 += 128
                        _ = v26
                    } else {
                        var v27 bool = false
                        v28 := GenS2{F0: 85, F1: 4294967295, F2: 0}
                        v29 := []int{v1}
                        _ = v27
                        _ = v28
                        _ = v29
                    }
                    v3 = ((v3 + ((0.0 + (((v10 + 1e10) * 0.0) + (v10 + 0.0))) / -2.25)) * -2.25)
                    var v30 int16 = 733
                    _ = v25
                    _ = v30
                }
                if true {
                    v10 = float32((v4 / 1e10))
                    v3 = (v3 / v3)
                    if true {
                        if ((4294967295 >= 18446744073709551615) || (v6 == fi21)) {
                            v1 -= 32767
                            var v31 string = "a"
                            var v32 uint32 = 2147483647
                            _ = v31
                            _ = v32
                        } else {
                            v3 += v10
                            fmt.Println(v10, v1, v6)
                        }
                        fmt.Println(v1, fi22)
                    }
                } else {
                    if true {
                        if (!(false && false)) {
                            v1 = (fi21 / ((v1) | 1))
                        } else {
                            v33 := GenS3{F0: 1}
                            v34 := GenS1{F0: 0, F1: 202}
                            var v35 uint32 = 65535
                            _ = v33
                            _ = v34
                            _ = v35
                        }
                        v3 = float32(1.5)
                        fmt.Println(v3, v10, v1)
                        v3 = float32((v4 / 1.0))
                        v2 -= fi22
                        var v36 uint16 = 42
                        v37 := v5[-1]
                        _ = v36
                        _ = v37
                    }
                }
                fmt.Println(v2, v4)
                _ = fi22
            }
            switch "foo bar" {
            case "":
                if ((!(true || (-1 == -1))) || ((4294967295 != 738) && (2147483647 < 42))) {
                    if (878 == 4294967295) {
                        var v38 uint = 65535
                        _ = v38
                    }
                }
            case "a":
                _ = len(v5)
            case "foo bar":
                for fi39 := 0; fi39 < 12; fi39++ {
                    _ = len(v5)
                    _ = fi39
                }
            }
            _ = fi21
        }
    }
    for fi40 := 0; fi40 < 8; fi40++ {
        if true {
            var v41 uint8 = 107
            _ = v41
        } else {
            v42 := []int{(v2 % (((fi40 * (v2 >> ((203) & 63)))) | 1)), 128, (fi40 * (v6 | -675))}
            for fi43 := 0; fi43 < 13; fi43++ {
                v1 = (v1 << ((2) & 63))
                _ = fi43
            }
            v1 += -1
            _ = v42
        }
        for fi44 := 0; fi44 < 3; fi44++ {
            var v45 uint64 = 1
            _ = fi44
            _ = v45
        }
        v6 -= v6
        _ = fi40
    }
    for fi46 := 0; fi46 < 20; fi46++ {
        for fi47 := 0; fi47 < 11; fi47++ {
            if false {
                delete(v5, -9223372036854775808)
                var v48 uint32 = 878
                switch "" {
                case "foo bar":
                    var v49 int16 = 42
                    _ = v49
                }
                _ = v48
            } else {
                v50 := GenS3{F0: 0}
                var v51 string = ""
                _ = v50
                _ = v51
            }
            v52 := GenS4{F0: 319}
            _ = fi47
            _ = v52
        }
        v53 := []int{v2, v6, v2, -65}
        v3 = (float32(v4) - 1e10)
        _ = fi46
        _ = v53
    }
    v54 := GenS5{F0: 0}
    v4 -= -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v6
    _ = v7
}
