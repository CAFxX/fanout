package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16; F1 bool; F2 uint32 }
type GenS4 struct { F0 int64; F1 uint16; F2 int16 }
type GenS5 struct { F0 uint16; F1 bool; F2 int8 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint; F1 uint8 }
func genFn0(p0 uint32) int64 {
    if false {
        switch -2147483648 {
        case 127:
            var v7 uint = 1
            _ = v7
        case 4294967295:
            v8 := []int{32767, -772, 1, 968}
            _ = v8
        default:
            var v9 uint = 692
            _ = v9
        }
        if true {
            fmt.Println(p0)
            for fi10 := 0; fi10 < 4; fi10++ {
                var v11 int32 = 42
                p0 = (^(p0))
                _ = fi10
                _ = v11
            }
        } else {
            if (701 > 4294967295) {
                v12 := GenS1{F0: -734, F1: "a"}
                v13 := GenS1{F0: int16(p0), F1: ("" + "\n\t")}
                _ = v12
                _ = v13
            } else {
                switch -46 {
                case 32767:
                    fmt.Println(p0)
                case -498:
                    p0 += p0
                default:
                    var v14 uint32 = p0
                    _ = v14
                }
                p0 = p0
            }
            var v15 int32 = -1
            var v16 int = -9223372036854775808
            _ = v15
            _ = v16
        }
        var v17 uint64 = 337
        _ = v17
    } else {
        p0 = (p0 & p0)
        switch "hello" {
        case "":
            if false {
                for fi18 := 0; fi18 < 12; fi18++ {
                    v19 := GenS2{F0: 2}
                    var v20 int32 = -1
                    _ = fi18
                    _ = v19
                    _ = v20
                }
            } else {
                p0 = p0
                if false {
                    if false {
                        v21 := GenS3{F0: 0, F1: (627 >= 42), F2: (p0 ^ p0)}
                        for fi22 := 0; fi22 < 10; fi22++ {
                            var v23 float32 = 1.5
                            var v24 float64 = 1e10
                            for fi25 := 0; fi25 < 13; fi25++ {
                                v23 = (v23 / (v23 - (v23 / ((v23 / (3.14159 * 0.1)) - v23))))
                                p0 = p0
                                _ = fi25
                            }
                            _ = fi22
                            _ = v23
                            _ = v24
                        }
                        _ = v21
                    } else {
                        for fi26 := 0; fi26 < 16; fi26++ {
                            p0 += p0
                            v27 := GenS2{F0: 97}
                            _ = fi26
                            _ = v27
                        }
                        v28 := []int{516, 2147483647, -1, 153}
                        _ = v28
                    }
                    v29 := []int{1}
                    _ = v29
                }
                p0 = (p0 % ((p0) | 1))
            }
        case "foo bar":
            v30 := GenS1{F0: -1, F1: ("foo bar" + (("foo bar" + "") + ("foo bar" + "foo bar")))}
            _ = v30
        case "hello":
            p0 -= p0
        }
    }
    for fi31 := 0; fi31 < 7; fi31++ {
        fmt.Println(fi31, fi31)
        v32 := GenS3{F0: 255, F1: false, F2: p0}
        _ = fi31
        _ = v32
    }
    var v33 int32 = 18
    _ = p0
    _ = v33
    return 42
}

func main() {
    var v1 int8 = -87
    var v2 uint32 = 65535
    v3 := make(map[string]int)
    var v4 int64 = 42
    v5 := make(map[string]uint32)
    var v6 float64 = (((1.5 * 1e10) / (0.1 + 3.14159)) / 1.0)
    for fi34 := 0; fi34 < 19; fi34++ {
        var v35 uint32 = 42
        var v36 uint64 = 2
        var v37 string = "a"
        _ = fi34
        _ = v35
        _ = v36
        _ = v37
    }
    v38 := GenS4{F0: genFn0(v2), F1: 42, F2: 1000}
    switch "\n\t" {
    case "hello":
        v39 := GenS5{F0: 65535, F1: false, F2: v1}
        _ = v39
    }
    if ((!false) && true) {
        for fi40 := 0; fi40 < 18; fi40++ {
            fmt.Println(v4, v4)
            var v41 int64 = (v4 & (v4 & int64(fi40)))
            v42 := GenS1{F0: 1, F1: "a"}
            _ = fi40
            _ = v41
            _ = v42
        }
        if (false || true) {
            if ("" >= (((("" + "foo bar") + "\n\t") + "hello") + "foo bar")) {
                for fi43 := 0; fi43 < 3; fi43++ {
                    v6 += (((0.1 - 1.5) * 3.14159) + (0.1 + v6))
                    v44 := GenS6{F0: (v2 + v2)}
                    _ = fi43
                    _ = v44
                }
                switch "\n\t" {
                case "a":
                    v4 -= v4
                case "hello":
                    var v45 int32 = 498
                    _ = v45
                default:
                    v4 -= v4
                }
            }
        }
        v2 = v2
    }
    if true {
        var v46 int64 = v4
        var v47 int32 = 42
        _ = v46
        _ = v47
    }
    v48 := GenS3{F0: 3, F1: true, F2: v2}
    v2 = (v2 - v2)
    if false {
        v3[("foo bar" + ("a" + ("foo bar" + "\n\t")))] = -9223372036854775808
        for fi49 := 0; fi49 < 5; fi49++ {
            fmt.Println(fi49)
            var v50 int16 = -305
            v3[""] = fi49
            _ = fi49
            _ = v50
        }
        v51 := GenS7{F0: 65535, F1: 127}
        _ = v51
    }
    switch -888 {
    case 971:
        var v52 int = 508
        _ = v52
    }
    for fi53 := 0; fi53 < 6; fi53++ {
        delete(v5, "foo bar")
        var v54 float32 = 0.0
        v6 += v6
        _ = fi53
        _ = v54
    }
    v1 = (-(v1))
    switch "foo bar" {
    case "\n\t":
        v6 += 0.1
    case "a":
        var v55 int = 785
        _ = v55
    }
    delete(v5, "")
    switch (("foo bar" + ("\n\t" + "\n\t")) + ("foo bar" + "a")) {
    case "\n\t":
        for fi56 := 0; fi56 < 7; fi56++ {
            v57 := []int{(fi56 - (fi56 * fi56)), 0}
            var v58 float32 = (-2.25 - -2.25)
            for fi59 := 0; fi59 < 8; fi59++ {
                v6 = 3.14159
                _ = fi59
            }
            _ = fi56
            _ = v57
            _ = v58
        }
    case "foo bar":
        var v60 int = 2147483647
        _ = v60
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v48
    _ = v5
    _ = v6
}
