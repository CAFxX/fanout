package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string; F1 uint64 }
type GenS4 struct { F0 string; F1 uint32; F2 int16 }
type GenS5 struct { F0 uint8; F1 int; F2 uint }
type GenS6 struct { F0 uint8; F1 int; F2 uint64 }
type GenS7 struct { F0 int32; F1 int64 }
func genFn0() int8 {
    v5 := []int{-1, 0}
    v6 := GenS1{F0: 0, F1: (-1 == -1), F2: 738}
    var v7 int16 = 32767
    if (true || false) {
        if true {
            v7 = v7
            var v8 int8 = -99
            v8 += (^(v8))
            _ = v8
        } else {
            switch "hello" {
            case "\n\t":
                if ((((1 > 892) || (123 != -128)) && false) || false) {
                    var v9 uint64 = 18446744073709551615
                    var v10 string = "foo bar"
                    _ = v10
                    _ = v9
                }
            default:
                v7 += v7
            }
        }
        var v11 float64 = 3.14159
        if true {
            var v12 uint16 = 78
            var v13 float64 = v11
            for fi14 := 0; fi14 < 14; fi14++ {
                fmt.Println(fi14, v13, fi14)
                _ = fi14
            }
            _ = v12
            _ = v13
        }
        _ = v11
    } else {
        if len(v5) > 0 {
            v5[0] = 32768
        }
    }
    for fi15 := 0; fi15 < 4; fi15++ {
        v5 = append(v5, fi15)
        if true {
            if (!true) {
                for fi16 := 0; fi16 < 9; fi16++ {
                    if (v7 == (v7 % (((v7 / ((412) | 1))) | 1))) {
                        switch "foo bar" {
                        case "a":
                            v7 += v7
                        case "\n\t":
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        }
                        v17 := GenS2{F0: 2}
                        _ = v17
                    } else {
                        fmt.Println(fi15, v7, fi16)
                    }
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = fi16
                }
                for fi18 := 0; fi18 < 17; fi18++ {
                    v7 = v7
                    v7 = (-(v7))
                    v7 -= v7
                    _ = fi18
                }
            } else {
                v7 = (v7 - 0)
                fmt.Println(v7)
                v19 := GenS1{F0: 128, F1: (((0 != -781) || false) || (true || (-758 <= 0))), F2: 0}
                _ = v19
            }
        } else {
            v7 = v7
            if (-2.25 <= 0.1) {
                fmt.Println(v7)
                if len(v5) > 0 {
                    v5[0] = fi15
                }
                fmt.Println(v7)
            } else {
                v7 += 0
                switch fi15 {
                case 127:
                    _ = len(v5)
                }
                var v20 uint64 = 1
                _ = v20
            }
            var v21 uint = 43
            _ = v21
        }
        v22 := GenS2{F0: 255}
        _ = fi15
        _ = v22
    }
    _ = v5
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 int64, p1 int, p2 int) string {
    p0 -= (p0 >> ((42) & 63))
    if false {
        if (2 > 4294967295) {
            v23 := GenS3{F0: "a", F1: 2}
            v24 := GenS2{F0: 116}
            _ = v23
            _ = v24
        } else {
            v25 := []int{}
            if true {
                for fi26 := 0; fi26 < 17; fi26++ {
                    p1 += -2147483648
                    _ = len(v25)
                    _ = fi26
                }
                if (32767 <= 65535) {
                    if false {
                        v27 := GenS4{F0: ((("a" + "") + ("hello" + "\n\t")) + ("foo bar" + "hello")), F1: 0, F2: 42}
                        p0 += 9223372036854775807
                        v28 := GenS5{F0: 42, F1: 9223372036854775807, F2: 2}
                        _ = v27
                        _ = v28
                    }
                    if len(v25) >= 2 {
                        v25 = v25[1:len(v25)]
                    }
                } else {
                    fmt.Println(p1, p1)
                    var v29 uint64 = 524
                    for fi30 := 0; fi30 < 4; fi30++ {
                        if true {
                            var v31 uint64 = v29
                            var v32 float64 = 1.5
                            _ = v31
                            _ = v32
                        } else {
                            v33 := GenS5{F0: 2, F1: -637, F2: 797}
                            v34 := GenS4{F0: "", F1: 246, F2: -554}
                            var v35 uint8 = 60
                            _ = v33
                            _ = v34
                            _ = v35
                        }
                        _ = fi30
                    }
                    _ = v29
                }
            }
            _ = v25
        }
        if true {
            if (-485 != -32768) {
                fmt.Println(p0, p0, p2)
            } else {
                v36 := GenS1{F0: 42, F1: false, F2: 0}
                _ = v36
            }
            var v37 string = "foo bar"
            fmt.Println(p2)
            _ = v37
        }
    } else {
        p1 -= p2
        p0 = (p0 % ((p0) | 1))
    }
    if false {
        switch ("a" + "hello") {
        case "foo bar":
            p1 = int(p0)
        case "\n\t":
            v38 := []int{p1, p1, p2, p1}
            _ = v38
        default:
            v39 := GenS5{F0: 255, F1: p1, F2: 255}
            _ = v39
        }
        if (true && true) {
            v40 := []int{p1, (p2 & (p1 - (-(p1)))), (p1 >> ((4294967295) & 63))}
            fmt.Println(p0)
            var v41 uint32 = 4294967295
            _ = v40
            _ = v41
        }
        p2 -= (p1 & (p1 - -65))
    } else {
        v42 := GenS6{F0: 0, F1: p2, F2: 0}
        var v43 string = ""
        for fi44 := 0; fi44 < 9; fi44++ {
            switch v43 {
            case "hello":
                var v45 uint16 = 65535
                _ = v45
            case "\n\t":
                fmt.Println(v43, fi44)
            }
            v46 := []int{p1}
            for fi47 := 0; fi47 < 11; fi47++ {
                var v48 float64 = 3.14159
                v43 = v43
                p0 += (p0 | p0)
                var v49 int8 = 0
                v50 := GenS7{F0: 2147483647, F1: (p0 % (((p0 - p0)) | 1))}
                _ = fi47
                _ = v48
                _ = v49
                _ = v50
            }
            _ = fi44
            _ = v46
        }
        _ = v42
        _ = v43
    }
    v51 := []int{p1}
    _ = p0
    _ = p1
    _ = p2
    _ = v51
    return (("a" + "a") + "hello")
}

func main() {
    v1 := make(map[uint32]int)
    var v2 string = ("a" + "")
    var v3 uint16 = 42
    v4 := []uint32{}
    var v52 uint32 = 65535
    var v53 int32 = 2147483647
    for fi54 := 0; fi54 < 13; fi54++ {
        var v55 string = "a"
        _ = len(v1)
        delete(v1, (v52 | v52))
        _ = fi54
        _ = v55
    }
    v56 := v1[v52]
    v4 = append(v4, (v52 + v52))
    fmt.Println(v2, v53, v2)
    v3 = 1
    v57 := v1[(v52 - v52)]
    v58 := GenS7{F0: v53, F1: 674}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v52
    _ = v53
    _ = v56
    _ = v57
    _ = v58
}
