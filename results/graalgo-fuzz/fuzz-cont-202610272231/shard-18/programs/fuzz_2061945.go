package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 uint16; F1 int16; F2 uint }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int64; F1 uint }
type GenS5 struct { F0 int16; F1 bool }
type GenS6 struct { F0 string; F1 int }
type GenS7 struct { F0 int32; F1 uint32; F2 int }
func genFn0(p0 uint64) string {
    v5 := []int{127, -9223372036854775808, 0, 9223372036854775807}
    fmt.Println(p0)
    if len(v5) > 0 {
        v5[0] = -1
    }
    fmt.Println(p0)
    _ = p0
    _ = v5
    return "a"
}
func genFn1(p0 bool) string {
    v6 := GenS1{F0: 65535, F1: 995}
    for fi7 := 0; fi7 < 5; fi7++ {
        v8 := []int{fi7, 1, 336}
        var v9 int16 = 32767
        var v10 bool = (511 > 0)
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    switch 19 {
    case -738:
        p0 = true
    case -1:
        p0 = (146 < 0)
    }
    _ = p0
    _ = v6
    return ("foo bar" + "hello")
}

func main() {
    var v1 int8 = 42
    v2 := make(map[string]int)
    var v3 float32 = 1.5
    var v4 uint8 = 140
    delete(v2, "a")
    switch 1 {
    case 2147483647:
        var v11 int64 = 9223372036854775807
        _ = v11
    default:
        for fi12 := 0; fi12 < 13; fi12++ {
            for fi13 := 0; fi13 < 4; fi13++ {
                fmt.Println(v4, fi12)
                v1 += (v1 >> ((2) & 63))
                v14 := GenS1{F0: 422, F1: 407}
                var v15 int8 = v1
                v2["a"] = (fi13 & fi13)
                _ = fi13
                _ = v14
                _ = v15
            }
            _ = fi12
        }
    }
    switch 4294967295 {
    case 0:
        var v16 float64 = (((float64(-2.25) * ((1.5 * (1.5 / -2.25)) / 1.5)) + 0.0) - -2.25)
        _ = v16
    case 42:
        fmt.Println(v1, v4, v3)
    case 32767:
        if (!true) {
            v17 := GenS1{F0: 870, F1: 954}
            _ = v17
        } else {
            fmt.Println(v4, v3)
            v18 := []int{}
            _ = v18
        }
    default:
        delete(v2, "hello")
    }
    fmt.Println(v1, v1)
    if (255 >= v4) {
        var v19 int32 = 1
        _ = v19
    }
    v20 := GenS2{F0: 42, F1: 42, F2: 4294967295}
    var v21 uint16 = 1
    var v22 int8 = -128
    switch -9223372036854775808 {
    case -97:
        v21 = v21
    }
    if false {
        fmt.Println(v4, v3)
    } else {
        var v23 uint32 = 4294967295
        if (-518 == 1) {
            for fi24 := 0; fi24 < 10; fi24++ {
                fmt.Println(v1)
                var v25 uint32 = v23
                _ = fi24
                _ = v25
            }
            switch ("" + "a") {
            case "foo bar":
                v26 := []int{2147483647}
                _ = v26
            case "":
                var v27 float32 = 3.14159
                _ = v27
            default:
                if false {
                    for fi28 := 0; fi28 < 15; fi28++ {
                        switch fi28 {
                        case -662:
                            v29 := GenS3{F0: 9223372036854775807}
                            _ = v29
                        case -1:
                            v30 := GenS2{F0: v21, F1: 11, F2: 4294967295}
                            _ = v30
                        default:
                            switch fi28 {
                            case -2147483648:
                                v31 := []int{-9223372036854775808}
                                _ = v31
                            }
                        }
                        _ = fi28
                    }
                    for fi32 := 0; fi32 < 3; fi32++ {
                        v21 += (v21 / (((v21 * v21)) | 1))
                        for fi33 := 0; fi33 < 18; fi33++ {
                            if false {
                                if false {
                                    if false {
                                        v2["\n\t"] = (-(fi32))
                                        v34 := []int{fi33, 9223372036854775807}
                                        if len(v34) > 0 {
                                            v34[0] = (fi32 % ((fi32) | 1))
                                        }
                                        _ = v34
                                    }
                                }
                                fmt.Println(v3)
                            }
                            v23 -= 2147483647
                            v35 := []int{fi32, -1, 42}
                            _ = fi33
                            _ = v35
                        }
                        for fi36 := 0; fi36 < 2; fi36++ {
                            v37 := GenS1{F0: 587, F1: 1}
                            var v38 uint8 = 94
                            v4 += (v4 + v4)
                            var v39 uint16 = (v21 | v21)
                            if ((-2147483648 > -254) || true) {
                                v40 := GenS2{F0: 190, F1: 32767, F2: 538}
                                v4 = v38
                                switch (fi36 & fi32) {
                                case 9223372036854775807:
                                    var v41 uint64 = 18446744073709551615
                                    _ = v41
                                case -2147483648:
                                    switch 885 {
                                    case -2147483648:
                                        fmt.Println(v4, v3)
                                    default:
                                        v23 = (v23 % ((v23) | 1))
                                    }
                                case 1:
                                    v42 := []int{}
                                    _ = v42
                                }
                                _ = v40
                            } else {
                                var v43 string = "foo bar"
                                fmt.Println(v23, fi32, v39)
                                _ = v43
                            }
                            _ = fi36
                            _ = v37
                            _ = v38
                            _ = v39
                        }
                        _ = fi32
                    }
                    v44 := GenS4{F0: -1, F1: uint(v4)}
                    _ = v44
                }
            }
            var v45 uint16 = 2
            _ = v45
        }
        _ = v23
    }
    var v46 uint16 = (v21 % ((200) | 1))
    if true {
        v47 := GenS5{F0: 1000, F1: false}
        for fi48 := 0; fi48 < 7; fi48++ {
            fmt.Println(fi48, v4, v1)
            v49 := []int{}
            _ = fi48
            _ = v49
        }
        for fi50 := 0; fi50 < 10; fi50++ {
            var v51 uint = 103
            for fi52 := 0; fi52 < 4; fi52++ {
                v22 = 4
                if false {
                    if true {
                        var v53 uint8 = (v4 >> ((v51) & 63))
                        var v54 int8 = (^(v1))
                        var v55 float64 = float64((v3 * v3))
                        v55 = v55
                        fmt.Println(v53, v53, v21)
                        _ = v53
                        _ = v54
                        _ = v55
                    } else {
                        fmt.Println(v21, v46)
                        v56 := GenS6{F0: "foo bar", F1: (-(fi52))}
                        _ = v56
                    }
                    delete(v2, "")
                } else {
                    switch "" {
                    case "hello":
                        if (!false) {
                            v4 = uint8(v51)
                            switch "\n\t" {
                            case "":
                                fmt.Println(v22)
                            case "a":
                                v57 := []int{fi52}
                                _ = v57
                            }
                            v4 = (v4 | v4)
                            v58 := []int{}
                            var v59 uint16 = v46
                            _ = v58
                            _ = v59
                        }
                    case "":
                        v1 += 42
                    default:
                        var v60 string = (("hello" + ("a" + "hello")) + "a")
                        _ = v60
                    }
                }
                v4 = 128
                v61 := GenS4{F0: 0, F1: v51}
                switch "\n\t" {
                case "\n\t":
                    _ = len(v2)
                case "a":
                    var v62 uint8 = (^(v4))
                    switch "\n\t" {
                    case "hello":
                        var v63 string = ("\n\t" + "\n\t")
                        _ = v63
                    case "foo bar":
                        v1 -= 42
                    }
                    v64 := []int{(fi50 * -74), fi50, (fi50 % (((fi52 + 310)) | 1))}
                    _ = v62
                    _ = v64
                case "":
                    var v65 int32 = -2147483648
                    _ = v65
                }
                _ = fi52
                _ = v61
            }
            for fi66 := 0; fi66 < 5; fi66++ {
                if (!(18446744073709551615 != uint64(v46))) {
                    v2[(("foo bar" + "") + "")] = int(v1)
                    v67 := []int{(-(fi50))}
                    var v68 uint = 1
                    _ = v67
                    _ = v68
                }
                switch ("hello" + "a") {
                case "foo bar":
                    var v69 int = (fi50 >> (((v51 * v51)) & 63))
                    _ = v69
                case "a":
                    v46 = v46
                    v70 := GenS7{F0: 2147483647, F1: 2147483647, F2: (fi50 / (((fi50 ^ 2147483647)) | 1))}
                    _ = v70
                case "\n\t":
                    v3 += 0.0
                    if (!true) {
                        fmt.Println(v3, fi66, fi50)
                        fmt.Println(v4, v51, v21)
                        var v71 int16 = -1
                        _ = v71
                    }
                }
                _ = fi66
            }
            _ = fi50
            _ = v51
        }
        _ = v47
    }
    var v72 uint64 = uint64(v46)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v46
    _ = v72
}
