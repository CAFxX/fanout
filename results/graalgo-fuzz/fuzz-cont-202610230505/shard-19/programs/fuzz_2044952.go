package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32; F1 int; F2 uint64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint8; F1 int32 }
type GenS5 struct { F0 uint; F1 uint64 }
type GenS6 struct { F0 string; F1 int8 }
func genFn0(p0 uint64) int {
    var v7 string = ""
    for fi8 := 0; fi8 < 15; fi8++ {
        if (false && true) {
            switch "a" {
            case "":
                switch v7 {
                case "":
                    v7 = (v7 + (("hello" + (v7 + ("" + ("a" + "")))) + v7))
                default:
                    v9 := []int{}
                    _ = v9
                }
            default:
                var v10 uint = 18446744073709551615
                _ = v10
            }
        } else {
            v11 := GenS1{F0: p0}
            var v12 float64 = 1.5
            _ = v11
            _ = v12
        }
        _ = fi8
    }
    switch 127 {
    case 32767:
        switch 0 {
        case 0:
            v7 = ""
        }
    case 954:
        v13 := GenS2{F0: 42, F1: 42, F2: (p0 | p0)}
        _ = v13
    case 32768:
        if true {
            switch 128 {
            case 32767:
                v14 := []int{1}
                _ = v14
            case 32768:
                for fi15 := 0; fi15 < 20; fi15++ {
                    v16 := GenS1{F0: (p0 >> ((893) & 63))}
                    _ = fi15
                    _ = v16
                }
            }
        }
    }
    v17 := GenS2{F0: 0, F1: 1, F2: 280}
    _ = p0
    _ = v17
    _ = v7
    return -9223372036854775808
}
func genFn1(p0 uint, p1 int, p2 int32) bool {
    p1 -= genFn0(2)
    var v18 int = (p1 % ((9223372036854775807) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    return (!true)
}

func main() {
    var v1 float64 = (float64(0.0) / 1e10)
    var v2 int16 = 727
    var v3 uint16 = 329
    var v4 int64 = -524
    var v5 uint = 563
    v6 := make(map[string]bool)
    for fi19 := 0; fi19 < 20; fi19++ {
        v20 := GenS2{F0: 1, F1: (fi19 & fi19), F2: 589}
        if false {
            for fi21 := 0; fi21 < 11; fi21++ {
                if true {
                    var v22 bool = true
                    v23 := []int{fi19, fi19}
                    _ = len(v6)
                    fmt.Println(v3)
                    v24 := GenS1{F0: 2}
                    var v25 bool = (!v22)
                    if ((float32(v1) * ((1.5 * 1.5) + float32(0.0))) < (0.1 / -2.25)) {
                        v5 = (v5 << (((v5 % (((v5 | (v5 | v5))) | 1))) & 63))
                    }
                    _ = v22
                    _ = v23
                    _ = v24
                    _ = v25
                }
                _ = fi21
            }
            v2 += -468
        }
        switch ((("" + "foo bar") + "") + ("a" + "a")) {
        case "":
            switch (fi19 + (fi19 << ((734) & 63))) {
            case 32768:
                if false {
                    for fi26 := 0; fi26 < 9; fi26++ {
                        var v27 float64 = 0.1
                        if false {
                            switch "foo bar" {
                            case "\n\t":
                                v28 := GenS2{F0: 2147483647, F1: genFn0(573), F2: 0}
                                _ = v28
                            }
                        } else {
                            var v29 uint32 = 373
                            v30 := []int{fi26, fi26, 2147483647}
                            v31 := GenS3{F0: -1}
                            _ = v29
                            _ = v30
                            _ = v31
                        }
                        _ = fi26
                        _ = v27
                    }
                    if (!true) {
                        switch "foo bar" {
                        case "a":
                            var v32 int16 = (v2 ^ v2)
                            _ = v32
                        case "":
                            if (true && (v4 >= v4)) {
                                v33 := GenS1{F0: 4294967295}
                                fmt.Println(fi19, v5, v2)
                                for fi34 := 0; fi34 < 14; fi34++ {
                                    var v35 int = (fi34 << ((v5) & 63))
                                    v4 -= 1
                                    _ = fi34
                                    _ = v35
                                }
                                _ = v33
                            } else {
                                if (-79 >= 1) {
                                    var v36 float64 = float64(-2.25)
                                    _ = v36
                                } else {
                                    v6["hello"] = ((!true) && true)
                                }
                                v4 -= v4
                            }
                        }
                    } else {
                        v3 = 32767
                    }
                } else {
                    v3 += v3
                    v5 = (v5 + v5)
                }
            case 9223372036854775807:
                var v37 int32 = 2147483647
                _ = v37
            default:
                if false {
                    v38 := GenS4{F0: 255, F1: 0}
                    v39 := []int{32767, genFn0(0), (-(fi19)), (-(fi19))}
                    for fi40 := 0; fi40 < 9; fi40++ {
                        v41 := v6[""]
                        _ = fi40
                        _ = v41
                    }
                    _ = v38
                    _ = v39
                }
            }
        case "a":
            v2 = v2
        }
        _ = fi19
        _ = v20
    }
    var v42 int = -60
    if true {
        var v43 int32 = -2147483648
        v44 := []int{2147483647, 0, (v42 << ((v5) & 63)), v42}
        for fi45 := 0; fi45 < 11; fi45++ {
            v6["hello"] = true
            var v46 uint16 = uint16(v5)
            delete(v6, "hello")
            _ = fi45
            _ = v46
        }
        _ = v43
        _ = v44
    }
    v47 := GenS1{F0: 2}
    var v48 int8 = 1
    v49 := GenS3{F0: v48}
    v50 := GenS1{F0: 42}
    for fi51 := 0; fi51 < 19; fi51++ {
        var v52 uint64 = 254
        v53 := []int{}
        switch "hello" {
        case "foo bar":
            var v54 uint8 = 255
            _ = v54
        case "":
            delete(v6, "foo bar")
        }
        _ = fi51
        _ = v52
        _ = v53
    }
    v42 = (v42 >> ((v5) & 63))
    if true {
        var v55 float32 = 0.0
        if (!true) {
            switch int(v5) {
            case -1:
                for fi56 := 0; fi56 < 8; fi56++ {
                    if (v1 != (0.1 * ((v1 / (0.0 * (v1 - 0.1))) / v1))) {
                        v57 := GenS2{F0: 2147483647, F1: 127, F2: 18446744073709551615}
                        v58 := GenS1{F0: 18446744073709551615}
                        fmt.Println(fi56, v4, v48)
                        _ = v57
                        _ = v58
                    }
                    v3 -= (v3 >> (((v5 + v5)) & 63))
                    v42 = fi56
                    _ = fi56
                }
            case 32768:
                var v59 int64 = -9223372036854775808
                _ = v59
            default:
                v60 := []int{}
                _ = v60
            }
            fmt.Println(v4, v3)
            v42 -= (v42 * (v42 & (v42 | (v42 & -842))))
        }
        for fi61 := 0; fi61 < 9; fi61++ {
            v62 := GenS2{F0: 706, F1: (fi61 / ((v42) | 1)), F2: 376}
            _ = fi61
            _ = v62
        }
        _ = v55
    }
    v63 := GenS5{F0: (v5 * (v5 & (^(v5)))), F1: 0}
    var v64 uint = v5
    v65 := GenS6{F0: ("\n\t" + (("hello" + "\n\t") + "")), F1: int8(v42)}
    fmt.Println(v3, v1)
    v66 := []int{(v42 % (((v42 & (v42 % ((v42) | 1)))) | 1)), (v42 >> (((v5 % ((v5) | 1))) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v64:", v64)
    fmt.Println("slicelen v66:", len(v66), v66)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v6
    _ = v63
    _ = v64
    _ = v65
    _ = v66
}
