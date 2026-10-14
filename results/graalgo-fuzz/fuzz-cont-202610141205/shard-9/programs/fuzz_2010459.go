package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 int; F1 bool }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int16; F1 uint }
type GenS7 struct { F0 uint8; F1 uint }
type GenS8 struct { F0 int64; F1 uint32; F2 string }
type GenS9 struct { F0 uint32 }
func genFn0(p0 bool) uint32 {
    for fi7 := 0; fi7 < 18; fi7++ {
        v8 := []int{36}
        p0 = true
        for fi9 := 0; fi9 < 2; fi9++ {
            p0 = p0
            v10 := GenS1{F0: 42, F1: 1}
            _ = fi9
            _ = v10
        }
        _ = fi7
        _ = v8
    }
    var v11 bool = ((p0 && true) && false)
    var v12 uint = 720
    v12 = (v12 | (v12 / ((1) | 1)))
    _ = p0
    _ = v11
    _ = v12
    return 786
}
func genFn1(p0 int64, p1 uint32) uint16 {
    if (((p0 <= p0) || true) || false) {
        var v13 int32 = 302
        v14 := GenS1{F0: 920, F1: 2}
        _ = v13
        _ = v14
    } else {
        fmt.Println(p1, p0)
    }
    v15 := []int{127}
    switch "hello" {
    case "":
        if len(v15) >= 2 {
            v15 = v15[1:len(v15)]
        }
    case "foo bar":
        for fi16 := 0; fi16 < 20; fi16++ {
            var v17 uint32 = (p1 / (((p1 << ((42) & 63))) | 1))
            v18 := GenS2{F0: genFn0(true), F1: 255}
            p1 = 1
            _ = fi16
            _ = v17
            _ = v18
        }
    case "hello":
        v15 = append(v15, 1)
    }
    fmt.Println(p0, p1)
    switch ("\n\t" + "") {
    case "foo bar":
        v19 := GenS3{F0: 42, F1: (!false)}
        _ = v19
    default:
        for fi20 := 0; fi20 < 13; fi20++ {
            v21 := GenS2{F0: (^(p1)), F1: 601}
            _ = fi20
            _ = v21
        }
    }
    _ = p0
    _ = p1
    _ = v15
    return 65535
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]bool)
    v3 := []uint32{}
    v4 := make(map[int]bool)
    var v5 int = 352
    var v6 uint32 = 298
    if false {
        v5 = v5
        var v22 int = v5
        switch 204 {
        case 32767:
            v22 += (v5 ^ v5)
        case 128:
            v6 -= (v6 % ((2147483648) | 1))
        case -861:
            if (1 >= 524) {
                for fi23 := 0; fi23 < 16; fi23++ {
                    v24 := v2[v6]
                    _ = fi23
                    _ = v24
                }
                var v25 uint64 = 2
                _ = v25
            } else {
                var v26 bool = true
                _ = v26
            }
        default:
            v3 = append(v3, v6)
        }
        _ = v22
    } else {
        v27 := GenS4{F0: 42}
        _ = v27
    }
    for fi28 := 0; fi28 < 15; fi28++ {
        switch v5 {
        case 938:
            var v29 uint16 = 42
            _ = v29
        case -1:
            for fi30 := 0; fi30 < 20; fi30++ {
                v6 -= 676
                var v31 uint64 = 419
                _ = fi30
                _ = v31
            }
        case 32768:
            v32 := GenS2{F0: v6, F1: 2}
            _ = v32
        }
        v6 -= 1
        _ = fi28
    }
    if true {
        if (false || true) {
            var v33 uint8 = 2
            if false {
                switch "" {
                case "hello":
                    v34 := GenS4{F0: 42}
                    _ = v34
                case "\n\t":
                    v6 -= (v6 >> ((1) & 63))
                default:
                    v6 = (^(v6))
                }
                v35 := GenS3{F0: (v5 >> ((255) & 63)), F1: false}
                var v36 uint16 = 1
                _ = v35
                _ = v36
            } else {
                if true {
                    var v37 int = v5
                    var v38 uint16 = genFn1(-764, v6)
                    for fi39 := 0; fi39 < 13; fi39++ {
                        v38 += (v38 * (^(v38)))
                        _ = fi39
                    }
                    _ = v37
                    _ = v38
                }
                v40 := GenS5{F0: genFn1(-402, genFn0((!(945 >= 2))))}
                var v41 int = (v5 * v5)
                _ = v40
                _ = v41
            }
            v6 += v6
            _ = v33
        } else {
            switch "a" {
            case "hello":
                for fi42 := 0; fi42 < 16; fi42++ {
                    v5 += (v5 - v5)
                    _ = fi42
                }
            case "":
                switch "" {
                case "":
                    v5 -= (v5 | v5)
                case "\n\t":
                    fmt.Println(v6)
                default:
                    if true {
                        switch v5 {
                        case 128:
                            v6 = v6
                        default:
                            switch (^(v5)) {
                            case 127:
                                v43 := GenS6{F0: 0, F1: 0}
                                _ = v43
                            case 2147483647:
                                var v44 uint32 = genFn0(((!false) || false))
                                _ = v44
                            }
                        }
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        v1["a"] = (("foo bar" + "\n\t") + (((("hello" + "\n\t") + "a") + "a") + "a"))
                    }
                }
            case "foo bar":
                v6 = genFn0((false && (2 > 93)))
            }
        }
        fmt.Println(v6, v6)
        v4[(v5 ^ v5)] = (1 == 142)
    } else {
        if false {
            _ = len(v1)
            for fi45 := 0; fi45 < 11; fi45++ {
                delete(v1, "")
                for fi46 := 0; fi46 < 9; fi46++ {
                    fmt.Println(v5, fi45)
                    var v47 int = 2147483647
                    _ = fi46
                    _ = v47
                }
                switch ("\n\t" + "foo bar") {
                case "\n\t":
                    v5 = (-(fi45))
                case "a":
                    if true {
                        var v48 int = (v5 | v5)
                        _ = v48
                    }
                default:
                    if false {
                        if (true || ("hello" < "\n\t")) {
                            fmt.Println(v6, v5)
                            fmt.Println(fi45, v6)
                        } else {
                            v6 -= genFn0(((-(v5)) < fi45))
                        }
                    }
                }
                _ = fi45
            }
            var v49 uint64 = 2
            _ = v49
        } else {
            for fi50 := 0; fi50 < 2; fi50++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi50
            }
            fmt.Println(v5)
            for fi51 := 0; fi51 < 11; fi51++ {
                var v52 float64 = (1.5 + (-2.25 + 3.14159))
                _ = fi51
                _ = v52
            }
        }
        v53 := GenS7{F0: 0, F1: 0}
        _ = v53
    }
    v54 := GenS8{F0: 0, F1: 65535, F2: ""}
    var v55 string = "foo bar"
    fmt.Println(v6, v6, v5)
    var v56 int64 = -753
    switch (v5 * v5) {
    case 42:
        var v57 uint = uint(v6)
        _ = v57
    default:
        v5 = (v5 * v5)
    }
    var v58 float64 = ((-2.25 / 3.14159) / 1.5)
    var v59 int32 = 42
    var v60 int32 = -346
    var v61 uint64 = 1
    var v62 int = 876
    v63 := GenS6{F0: -540, F1: 1}
    if true {
        for fi64 := 0; fi64 < 14; fi64++ {
            fmt.Println(v6, v6)
            var v65 bool = (v62 < fi64)
            v1["a"] = v55
            _ = fi64
            _ = v65
        }
    }
    if false {
        v66 := GenS5{F0: 255}
        var v67 int64 = -891
        if false {
            if (false && (!false)) {
                var v68 uint32 = v6
                v60 = -178
                if (true && true) {
                    var v69 float64 = float64((0.0 * 0.1))
                    _ = v69
                }
                _ = v68
            } else {
                v70 := GenS9{F0: 1}
                v67 = (v67 + 658)
                _ = v70
            }
        }
        _ = v66
        _ = v67
    }
    fmt.Println(v6)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v55
    _ = v56
    _ = v58
    _ = v59
    _ = v6
    _ = v60
    _ = v61
    _ = v62
    _ = v63
}
