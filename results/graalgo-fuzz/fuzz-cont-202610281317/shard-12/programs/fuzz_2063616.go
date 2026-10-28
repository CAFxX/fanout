package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint; F1 uint8; F2 uint }
func genFn0(p0 int16, p1 int, p2 int32) uint32 {
    v6 := []int{9223372036854775807, 0}
    var v7 int = 42
    var v8 int32 = (p2 >> ((4294967295) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 4294967295
}
func genFn1(p0 string, p1 bool, p2 uint16) uint {
    v9 := GenS1{F0: p2}
    var v10 int8 = 1
    for fi11 := 0; fi11 < 11; fi11++ {
        for fi12 := 0; fi12 < 15; fi12++ {
            if (!false) {
                var v13 string = "hello"
                if p1 {
                    var v14 uint = 42
                    p1 = false
                    if false {
                        v15 := GenS1{F0: p2}
                        _ = v15
                    } else {
                        v13 = (("\n\t" + p0) + ("" + "a"))
                        switch ("\n\t" + (((v13 + ("foo bar" + "\n\t")) + v13) + v13)) {
                        case "foo bar":
                            v16 := []int{fi12, (^(fi12))}
                            _ = v16
                        case "a":
                            p2 = p2
                        case "":
                            switch p0 {
                            case "":
                                if true {
                                    fmt.Println(v10)
                                    v17 := []int{0, fi11}
                                    _ = v17
                                } else {
                                    var v18 float64 = 1e10
                                    var v19 uint64 = 18446744073709551615
                                    _ = v18
                                    _ = v19
                                }
                            case "\n\t":
                                var v20 int8 = (v10 % (((v10 & (-(v10)))) | 1))
                                _ = v20
                            }
                            fmt.Println(p2, p2, fi11)
                            if false {
                                var v21 uint32 = 65535
                                _ = v21
                            } else {
                                var v22 float32 = float32(3.14159)
                                var v23 float64 = -2.25
                                p0 = p0
                                _ = v22
                                _ = v23
                            }
                        }
                        v24 := []int{}
                        _ = v24
                    }
                    _ = v14
                } else {
                    fmt.Println(p0, p1, v10)
                    v10 = v10
                }
                _ = v13
            }
            if (p0 < ("foo bar" + p0)) {
                p0 = p0
                fmt.Println(fi11)
                var v25 uint16 = 134
                p0 = (p0 + p0)
                _ = v25
            }
            var v26 int = fi12
            _ = fi12
            _ = v26
        }
        _ = fi11
    }
    var v27 float32 = 0.1
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v27
    _ = v9
    return 42
}

func main() {
    v1 := make(map[int]bool)
    var v2 int = 4294967295
    v3 := make(map[uint32]string)
    var v4 uint8 = 194
    var v5 uint32 = 382
    fmt.Println(v5)
    for fi28 := 0; fi28 < 1; fi28++ {
        var v29 bool = true
        if v29 {
            v29 = (v29 || false)
        }
        _ = fi28
        _ = v29
    }
    v4 -= (v4 * 2)
    switch "a" {
    case "foo bar":
        for fi30 := 0; fi30 < 14; fi30++ {
            for fi31 := 0; fi31 < 9; fi31++ {
                v2 += int(v4)
                var v32 bool = (332 <= 2)
                v33 := GenS1{F0: 12}
                if (!v32) {
                    v2 += fi30
                    var v34 float32 = 3.14159
                    _ = len(v1)
                    _ = v34
                } else {
                    v4 += v4
                }
                _ = fi31
                _ = v32
                _ = v33
            }
            v5 = (v5 / (((^(v5))) | 1))
            _ = fi30
        }
    case "a":
        var v35 uint32 = (v5 & 466)
        _ = v35
    }
    v36 := []int{-1, v2, (v2 - (v2 << ((uint(v5)) & 63)))}
    v36 = append(v36, v2)
    for fi37 := 0; fi37 < 10; fi37++ {
        v36 = append(v36, int(v5))
        _ = fi37
    }
    var v38 int64 = -1
    switch "a" {
    case "":
        v2 += 32768
    }
    v39 := GenS2{F0: genFn1("", true, 105), F1: true}
    for fi40 := 0; fi40 < 4; fi40++ {
        fmt.Println(v2, v2, v5)
        _ = fi40
    }
    v38 -= v38
    _ = len(v36)
    var v41 uint8 = 170
    switch "foo bar" {
    case "\n\t":
        for fi42 := 0; fi42 < 18; fi42++ {
            var v43 float64 = float64(1e10)
            _ = fi42
            _ = v43
        }
    case "foo bar":
        v41 = 2
    case "":
        fmt.Println(v5, v4, v4)
    default:
        switch "\n\t" {
        case "hello":
            for fi44 := 0; fi44 < 7; fi44++ {
                switch "foo bar" {
                case "a":
                    v45 := GenS2{F0: 18446744073709551615, F1: ((-128 <= -128) && (false || true))}
                    _ = v45
                case "hello":
                    _ = len(v3)
                }
                switch ("hello" + ("\n\t" + ("foo bar" + "a"))) {
                case "hello":
                    var v46 float32 = 3.14159
                    _ = v46
                case "foo bar":
                    v41 -= (v41 + (v41 / ((v4) | 1)))
                case "a":
                    switch v2 {
                    case 370:
                        v2 -= (fi44 / (((-(v2))) | 1))
                    }
                default:
                    for fi47 := 0; fi47 < 20; fi47++ {
                        fmt.Println(fi47, fi44)
                        _ = fi47
                    }
                }
                _ = fi44
            }
        case "a":
            _ = len(v3)
        case "\n\t":
            for fi48 := 0; fi48 < 14; fi48++ {
                var v49 uint16 = 65535
                for fi50 := 0; fi50 < 7; fi50++ {
                    switch -1 {
                    case 32768:
                        var v51 uint16 = v49
                        _ = v51
                    }
                    fmt.Println(fi48)
                    _ = fi50
                }
                _ = fi48
                _ = v49
            }
        }
    }
    if ((genFn0(1, v2, 42) == v5) || (-1 < 0)) {
        if false {
            for fi52 := 0; fi52 < 14; fi52++ {
                v53 := GenS2{F0: 4294967295, F1: false}
                _ = fi52
                _ = v53
            }
        }
        var v54 int16 = 820
        var v55 int = (v2 % ((127) | 1))
        _ = v54
        _ = v55
    } else {
        v56 := GenS2{F0: 0, F1: ((42 >= 42) && ("a" == "foo bar"))}
        for fi57 := 0; fi57 < 12; fi57++ {
            for fi58 := 0; fi58 < 17; fi58++ {
                var v59 uint8 = (v4 | 2)
                if true {
                    fmt.Println(v38)
                    switch "" {
                    case "a":
                        fmt.Println(v4, v41)
                    case "\n\t":
                        v60 := GenS3{F0: 1}
                        _ = v60
                    default:
                        _ = len(v36)
                    }
                    if false {
                        var v61 float32 = 1.5
                        _ = v61
                    } else {
                        var v62 uint16 = 1
                        if (("hello" >= "foo bar") && ((v59 * 5) == v41)) {
                            var v63 bool = (!false)
                            if v63 {
                                var v64 uint32 = 42
                                var v65 int8 = 72
                                _ = v64
                                _ = v65
                            } else {
                                switch "\n\t" {
                                case "a":
                                    v66 := GenS3{F0: genFn1("", false, v62)}
                                    _ = v66
                                case "hello":
                                    v63 = v63
                                case "":
                                    v67 := GenS4{F0: 116}
                                    _ = v67
                                }
                                v68 := GenS5{F0: 65535, F1: (v59 * (^(v4))), F2: 1}
                                _ = len(v36)
                                _ = v68
                            }
                            var v69 int8 = -128
                            _ = v63
                            _ = v69
                        }
                        delete(v3, (v5 >> ((uint(v5)) & 63)))
                        _ = v62
                    }
                }
                switch "hello" {
                case "hello":
                    v38 += v38
                }
                _ = fi58
                _ = v59
            }
            _ = fi57
        }
        var v70 int32 = -584
        _ = v56
        _ = v70
    }
    fmt.Println(v4)
    v36 = append(v36, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v38
    _ = v39
    _ = v4
    _ = v41
    _ = v5
}
