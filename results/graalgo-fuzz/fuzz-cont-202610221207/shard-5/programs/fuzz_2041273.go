package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32; F1 uint16 }
type GenS4 struct { F0 uint64; F1 uint8; F2 int }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int64 }
func genFn0(p0 uint8, p1 int, p2 uint32) int8 {
    p0 = 255
    var v7 uint = 2
    for fi8 := 0; fi8 < 18; fi8++ {
        v9 := []int{562, (-(fi8))}
        if false {
            p1 = -713
            v10 := GenS1{F0: p2, F1: (p2 / ((p2) | 1)), F2: true}
            _ = v10
        }
        _ = fi8
        _ = v9
    }
    fmt.Println(p1, p2)
    v7 -= v7
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return -18
}
func genFn1(p0 uint32) uint32 {
    var v11 bool = (!(965 >= 42))
    var v12 int32 = 42
    p0 = p0
    fmt.Println(v12, p0)
    if v11 {
        if true {
            var v13 string = ("\n\t" + ("foo bar" + "a"))
            _ = v13
        }
        if v11 {
            if (42 != -1) {
                switch "\n\t" {
                case "a":
                    if ((0.1 * (1e10 * 3.14159)) <= 3.14159) {
                        v14 := []int{496}
                        v11 = true
                        _ = v14
                    }
                case "":
                    for fi15 := 0; fi15 < 19; fi15++ {
                        for fi16 := 0; fi16 < 13; fi16++ {
                            switch "a" {
                            case "":
                                v11 = false
                            case "hello":
                                var v17 bool = v11
                                _ = v17
                            }
                            var v18 float32 = 1.5
                            fmt.Println(v12, p0)
                            _ = fi16
                            _ = v18
                        }
                        if v11 {
                            v19 := []int{2147483647, (fi15 % ((0) | 1))}
                            _ = v19
                        } else {
                            for fi20 := 0; fi20 < 11; fi20++ {
                                v12 = v12
                                var v21 uint16 = 32767
                                v22 := []int{}
                                _ = fi20
                                _ = v21
                                _ = v22
                            }
                            if true {
                                v23 := []int{(fi15 / ((111) | 1)), 2147483647}
                                var v24 uint8 = 255
                                v25 := GenS2{F0: 465}
                                _ = v23
                                _ = v24
                                _ = v25
                            } else {
                                v12 -= -2147483648
                            }
                        }
                        fmt.Println(v11, fi15)
                        _ = fi15
                    }
                }
            } else {
                var v26 uint16 = 42
                var v27 uint64 = 42
                v28 := []int{}
                _ = v26
                _ = v27
                _ = v28
            }
        } else {
            var v29 float64 = 1e10
            for fi30 := 0; fi30 < 6; fi30++ {
                if v11 {
                    v31 := GenS3{F0: (p0 / ((p0) | 1)), F1: 98}
                    _ = v31
                } else {
                    switch fi30 {
                    case 32768:
                        var v32 uint16 = 42
                        _ = v32
                    case -1:
                        var v33 float64 = 1e10
                        _ = v33
                    }
                }
                v11 = v11
                var v34 bool = true
                _ = fi30
                _ = v34
            }
            _ = v29
        }
    } else {
        v12 += (v12 << ((255) & 63))
    }
    _ = p0
    _ = v11
    _ = v12
    return 2
}

func main() {
    var v1 int8 = -128
    var v2 uint = 18446744073709551615
    var v3 float32 = 0.1
    var v4 int64 = -9223372036854775808
    v5 := []string{"\n\t", "", ("foo bar" + "hello")}
    var v6 uint64 = 0
    v2 -= 18446744073709551615
    for fi35 := 0; fi35 < 1; fi35++ {
        v2 = (v2 << ((v2) & 63))
        _ = fi35
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v4 += 42
    v36 := GenS1{F0: genFn1(genFn1(2147483647)), F1: 65535, F2: true}
    v37 := GenS3{F0: 619, F1: 65535}
    fmt.Println(v2, v6, v4)
    v6 = v6
    for fi38 := 0; fi38 < 11; fi38++ {
        switch "\n\t" {
        case "hello":
            if (v4 >= (v4 % ((v4) | 1))) {
                if false {
                    var v39 uint16 = 42
                    var v40 uint32 = 2
                    if true {
                        switch ("" + "foo bar") {
                        case "\n\t":
                            v1 += (v1 * (v1 & (v1 + v1)))
                        case "a":
                            for fi41 := 0; fi41 < 16; fi41++ {
                                var v42 uint16 = v39
                                if len(v5) > 0 {
                                    v5[0] = ("a" + "\n\t")
                                }
                                _ = fi41
                                _ = v42
                            }
                        case "":
                            v40 = 160
                        }
                    }
                    _ = v39
                    _ = v40
                } else {
                    if true {
                        var v43 uint32 = 0
                        _ = v43
                    } else {
                        for fi44 := 0; fi44 < 12; fi44++ {
                            v45 := GenS1{F0: 4294967295, F1: genFn1(0), F2: true}
                            fmt.Println(v6, v3)
                            var v46 string = (("" + "foo bar") + ("foo bar" + "hello"))
                            _ = fi44
                            _ = v45
                            _ = v46
                        }
                        for fi47 := 0; fi47 < 1; fi47++ {
                            v2 = (v2 + uint(v6))
                            v4 = 0
                            fmt.Println(v4, fi47, v2)
                            _ = fi47
                        }
                    }
                }
                for fi48 := 0; fi48 < 7; fi48++ {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = fi48
                }
                var v49 uint32 = 2147483647
                _ = v49
            } else {
                v4 = (v4 - v4)
                fmt.Println(v6)
                for fi50 := 0; fi50 < 12; fi50++ {
                    switch "\n\t" {
                    case "\n\t":
                        v51 := GenS4{F0: (^(v6)), F1: 128, F2: (-(fi50))}
                        _ = v51
                    case "hello":
                        v6 += (v6 & v6)
                    case "":
                        v52 := GenS5{F0: true}
                        _ = v52
                    }
                    v53 := GenS4{F0: v6, F1: 2, F2: fi50}
                    fmt.Println(fi38, fi50)
                    _ = fi50
                    _ = v53
                }
            }
        }
        _ = fi38
    }
    fmt.Println(v2)
    if (248 <= v6) {
        var v54 int32 = 106
        _ = v54
    } else {
        var v55 uint64 = v6
        _ = len(v5)
        _ = v55
    }
    for fi56 := 0; fi56 < 3; fi56++ {
        if (42 >= 127) {
            v57 := GenS5{F0: ((821 <= (fi56 % (((fi56 / (((fi56 & (-(fi56)))) | 1))) | 1))) && true)}
            _ = v57
        } else {
            for fi58 := 0; fi58 < 20; fi58++ {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = fi58
            }
            v59 := GenS6{F0: v4}
            _ = v59
        }
        _ = fi56
    }
    v60 := GenS5{F0: ((!(genFn1(42) == 612)) && true)}
    v1 += (v1 + (-(v1)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v60
}
