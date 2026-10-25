package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int8; F1 int8 }
type GenS5 struct { F0 bool; F1 bool }
func genFn0(p0 uint16, p1 int32, p2 int64) uint8 {
    p0 = p0
    var v6 int = 1
    fmt.Println(v6)
    for fi7 := 0; fi7 < 12; fi7++ {
        var v8 int8 = 68
        for fi9 := 0; fi9 < 3; fi9++ {
            if ((false && false) && false) {
                var v10 int32 = -344
                v11 := []int{}
                var v12 uint16 = p0
                _ = v10
                _ = v11
                _ = v12
            }
            _ = fi9
        }
        if (!((^(p1)) <= (^(p1)))) {
            v13 := GenS1{F0: 164, F1: 32767, F2: fi7}
            for fi14 := 0; fi14 < 19; fi14++ {
                if false {
                    fmt.Println(fi7, p2, v6)
                    v8 -= v8
                    p0 += (^(p0))
                } else {
                    if (32767 <= 32767) {
                        v8 = -1
                    }
                    var v15 int = -2147483648
                    _ = v15
                }
                v16 := []int{(v6 | 555), 128, (fi14 / (((-(v6))) | 1)), (fi14 * (v6 + (fi7 | (^(v6)))))}
                _ = fi14
                _ = v16
            }
            _ = v13
        }
        _ = fi7
        _ = v8
    }
    if (!(false && true)) {
        v17 := GenS1{F0: 2, F1: 1000, F2: 582}
        fmt.Println(p1)
        _ = v17
    } else {
        v18 := GenS2{F0: "hello"}
        p1 -= (p1 ^ (p1 | p1))
        for fi19 := 0; fi19 < 12; fi19++ {
            v6 = fi19
            if false {
                for fi20 := 0; fi20 < 4; fi20++ {
                    switch "a" {
                    case "":
                        v21 := []int{}
                        _ = v21
                    case "hello":
                        v6 = 0
                    case "\n\t":
                        v22 := GenS2{F0: "hello"}
                        _ = v22
                    }
                    _ = fi20
                }
                var v23 uint32 = 967
                for fi24 := 0; fi24 < 10; fi24++ {
                    v6 -= (fi19 * (^(fi19)))
                    v25 := GenS3{F0: -128}
                    _ = fi24
                    _ = v25
                }
                _ = v23
            }
            for fi26 := 0; fi26 < 5; fi26++ {
                if true {
                    switch ("hello" + "") {
                    case "foo bar":
                        var v27 float32 = float32(0.0)
                        _ = v27
                    case "":
                        v6 = 4294967295
                    case "hello":
                        v28 := []int{fi19, v6}
                        _ = v28
                    }
                    switch "hello" {
                    case "foo bar":
                        v29 := GenS1{F0: 4294967295, F1: 42, F2: (v6 / (((fi26 >> ((656) & 63))) | 1))}
                        _ = v29
                    case "\n\t":
                        p2 = p2
                    default:
                        switch fi26 {
                        case -1:
                            var v30 uint8 = 42
                            _ = v30
                        case 0:
                            var v31 int8 = 0
                            _ = v31
                        case -557:
                            switch "hello" {
                            case "\n\t":
                                v32 := []int{2147483647, (fi19 + (-(fi19))), fi26, -181}
                                _ = v32
                            case "hello":
                                p2 -= 9223372036854775807
                            case "":
                                p2 = p2
                                var v33 uint = 4294967295
                                if false {
                                    v34 := []int{}
                                    fmt.Println(p1)
                                    _ = v34
                                }
                                _ = v33
                            }
                        default:
                            p1 -= -707
                        }
                    }
                    var v35 float32 = (float32(0.1) / float32(0.1))
                    _ = v35
                }
                fmt.Println(fi26, p2, fi26)
                v36 := GenS1{F0: uint(p0), F1: 42, F2: fi26}
                _ = fi26
                _ = v36
            }
            _ = fi19
        }
        _ = v18
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 255
}
func genFn1() int8 {
    for fi37 := 0; fi37 < 14; fi37++ {
        v38 := []int{2147483647, 76}
        v39 := GenS4{F0: 127, F1: 84}
        var v40 uint32 = 123
        _ = fi37
        _ = v38
        _ = v39
        _ = v40
    }
    v41 := GenS3{F0: 31}
    var v42 float64 = 3.14159
    var v43 string = "hello"
    v43 = "foo bar"
    _ = v41
    _ = v42
    _ = v43
    return 127
}

func main() {
    var v1 string = ("" + "a")
    v2 := make(map[int]uint32)
    var v3 float32 = ((0.0 + 1.5) * 1e10)
    var v4 uint32 = 0
    var v5 int32 = 0
    var v44 uint32 = (v4 << ((989) & 63))
    for fi45 := 0; fi45 < 8; fi45++ {
        v5 = (v5 * (-(v5)))
        var v46 int = -9223372036854775808
        _ = fi45
        _ = v46
    }
    v47 := GenS5{F0: false, F1: (!false)}
    var v48 string = (v1 + "hello")
    if (1 >= 1) {
        var v49 bool = true
        if true {
            if v49 {
                v4 = v4
            } else {
                v50 := []int{128}
                var v51 int64 = -9223372036854775808
                var v52 int16 = 42
                _ = v50
                _ = v51
                _ = v52
            }
        }
        _ = v49
    }
    v4 = 2147483647
    var v53 int64 = -450
    var v54 uint16 = 780
    fmt.Println(v48, v1, v48)
    v55 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("slicelen v55:", len(v55), v55)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v47
    _ = v48
    _ = v5
    _ = v53
    _ = v54
    _ = v55
}
