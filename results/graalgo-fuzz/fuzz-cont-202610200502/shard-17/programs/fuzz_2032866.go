package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 uint64; F1 string }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 string; F1 uint; F2 int }
func genFn0(p0 string, p1 bool) uint32 {
    if true {
        if p1 {
            v4 := GenS1{F0: 465}
            if p1 {
                var v5 int32 = 822
                var v6 int32 = (v5 + (v5 | v5))
                _ = v5
                _ = v6
            }
            switch p0 {
            case "a":
                v7 := GenS1{F0: 0}
                _ = v7
            case "hello":
                v8 := []int{-75, 127}
                _ = v8
            default:
                for fi9 := 0; fi9 < 14; fi9++ {
                    if (!true) {
                        if (128 <= (fi9 & fi9)) {
                            if true {
                                var v10 uint64 = 1
                                _ = v10
                            } else {
                                v11 := GenS1{F0: -2147483648}
                                p1 = (p1 || ((!false) && p1))
                                var v12 string = "hello"
                                _ = v11
                                _ = v12
                            }
                        } else {
                            v13 := GenS2{F0: 1000, F1: 141}
                            var v14 bool = (!true)
                            var v15 uint16 = 290
                            _ = v13
                            _ = v14
                            _ = v15
                        }
                        v16 := []int{2147483647, fi9, fi9, 32768}
                        _ = v16
                    } else {
                        fmt.Println(fi9, fi9, p0)
                        if p1 {
                            if true {
                                var v17 uint32 = 42
                                p1 = ((p1 && (467 > 814)) && (-128 == 42))
                                _ = v17
                            }
                            v18 := []int{fi9, fi9, (fi9 << ((65535) & 63)), (fi9 + (fi9 / ((fi9) | 1)))}
                            _ = v18
                        } else {
                            var v19 int64 = 1
                            _ = v19
                        }
                    }
                    if (!(p1 || false)) {
                        p1 = true
                        if false {
                            fmt.Println(p0, fi9, p1)
                        } else {
                            if true {
                                fmt.Println(p0, p0)
                                var v20 uint32 = 4294967295
                                _ = v20
                            } else {
                                var v21 int8 = 1
                                _ = v21
                            }
                        }
                    } else {
                        v22 := GenS1{F0: 42}
                        _ = v22
                    }
                    var v23 uint32 = 42
                    _ = fi9
                    _ = v23
                }
            }
            _ = v4
        } else {
            var v24 int8 = 127
            v25 := []int{-766, 2147483647}
            v26 := GenS3{F0: 83, F1: "hello"}
            _ = v24
            _ = v25
            _ = v26
        }
        v27 := GenS3{F0: 2, F1: ("hello" + p0)}
        _ = v27
    } else {
        p0 = (p0 + (p0 + p0))
        v28 := GenS1{F0: -2147483648}
        _ = v28
    }
    p1 = ("a" > p0)
    var v29 uint32 = 2
    v30 := GenS2{F0: 32767, F1: -1}
    _ = p0
    _ = p1
    _ = v29
    _ = v30
    return v29
}
func genFn1() int16 {
    var v31 uint32 = 4294967295
    v31 = (v31 + (v31 ^ 2147483648))
    switch ("foo bar" + "a") {
    case "hello":
        if false {
            for fi32 := 0; fi32 < 11; fi32++ {
                v31 -= (v31 * v31)
                _ = fi32
            }
            switch "\n\t" {
            case "hello":
                var v33 float64 = (float64((((1e10 / -2.25) - (((0.0 * 0.1) - (0.0 / -2.25)) + 0.0)) / 1e10)) / 0.1)
                _ = v33
            case "a":
                switch "hello" {
                case "\n\t":
                    switch 127 {
                    case -248:
                        v31 -= (v31 | v31)
                    }
                case "foo bar":
                    fmt.Println(v31)
                default:
                    switch -1 {
                    case 127:
                        for fi34 := 0; fi34 < 19; fi34++ {
                            v35 := []int{32768, fi34, (fi34 % (((fi34 / ((fi34) | 1))) | 1))}
                            _ = fi34
                            _ = v35
                        }
                    default:
                        switch ("" + "a") {
                        case "foo bar":
                            var v36 uint8 = 185
                            _ = v36
                        case "\n\t":
                            for fi37 := 0; fi37 < 19; fi37++ {
                                v38 := []int{fi37, 75}
                                _ = fi37
                                _ = v38
                            }
                        }
                    }
                }
            case "":
                v31 -= 609
            }
            v31 += 0
        }
    default:
        var v39 int64 = 0
        _ = v39
    }
    v40 := []int{-672, 2147483647, -2147483648}
    _ = v31
    _ = v40
    return 1000
}

func main() {
    v1 := make(map[string]string)
    var v2 float64 = ((0.1 - (0.1 * float64(1.5))) * float64(float32(float64(0.0))))
    var v3 bool = true
    fmt.Println(v3)
    var v41 uint16 = 186
    if (474 != -912) {
        if true {
            v3 = (!(false && false))
        } else {
            v42 := []int{-1, 2147483647, 1, 42}
            v2 += v2
            fmt.Println(v41, v2, v2)
            _ = v42
        }
    }
    switch "foo bar" {
    case "\n\t":
        switch 634 {
        case 9223372036854775807:
            for fi43 := 0; fi43 < 7; fi43++ {
                var v44 uint8 = 1
                v45 := GenS4{F0: -71}
                for fi46 := 0; fi46 < 1; fi46++ {
                    v3 = (3.14159 <= (v2 / v2))
                    _ = fi46
                }
                _ = fi43
                _ = v44
                _ = v45
            }
        case 32768:
            switch 1 {
            case 127:
                switch 42 {
                case 846:
                    v47 := GenS5{F0: (18446744073709551615 > 42)}
                    _ = v47
                case 988:
                    v2 = v2
                }
            case 32768:
                v41 = (v41 | 737)
            default:
                var v48 uint = 37
                _ = v48
            }
        default:
            for fi49 := 0; fi49 < 18; fi49++ {
                v50 := []int{fi49}
                _ = fi49
                _ = v50
            }
        }
    default:
        v2 = ((v2 + -2.25) / (0.1 - 1.5))
    }
    v51 := GenS2{F0: -1, F1: -1}
    for fi52 := 0; fi52 < 2; fi52++ {
        if true {
            var v53 float64 = 0.0
            fmt.Println(v41, fi52)
            v41 = 2
            _ = v53
        }
        fmt.Println(v3)
        _ = fi52
    }
    v3 = ((float64(3.14159) / v2) < v2)
    v54 := []int{-2147483648}
    v55 := GenS6{F0: "", F1: 439, F2: 246}
    fmt.Println(v41, v41, v41)
    v3 = (v2 > 3.14159)
    v2 = (1e10 / v2)
    v2 += (0.0 / v2)
    var v56 int = 32767
    if false {
        v56 -= v56
        var v57 int16 = -32768
        v58 := GenS2{F0: genFn1(), F1: -2147483648}
        _ = v57
        _ = v58
    }
    fmt.Println(v2, v3, v41)
    var v59 int32 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v54:", len(v54), v54)
    fmt.Println("v56:", v56)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v2
    _ = v3
    _ = v41
    _ = v51
    _ = v54
    _ = v55
    _ = v56
    _ = v59
}
