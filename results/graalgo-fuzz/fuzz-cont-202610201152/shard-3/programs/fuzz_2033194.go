package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 int8; F1 int32; F2 bool }
type GenS4 struct { F0 int8; F1 uint64; F2 uint }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint8 }
func genFn0(p0 uint16) uint {
    if (18446744073709551615 != 42) {
        fmt.Println(p0)
        switch "foo bar" {
        case "\n\t":
            var v4 float64 = -2.25
            _ = v4
        default:
            for fi5 := 0; fi5 < 1; fi5++ {
                if true {
                    var v6 float64 = (-2.25 + 0.0)
                    _ = v6
                } else {
                    v7 := []int{-607, fi5, fi5, fi5}
                    switch ("\n\t" + (("a" + "hello") + "hello")) {
                    case "a":
                        var v8 uint64 = 271
                        _ = v8
                    case "hello":
                        if (65535 < 65535) {
                            var v9 int32 = -16
                            _ = v9
                        }
                    case "":
                        p0 = 0
                    default:
                        var v10 float32 = float32(float64(1e10))
                        _ = v10
                    }
                    _ = v7
                }
                _ = fi5
            }
        }
    } else {
        p0 += (p0 << ((42) & 63))
        if (1 >= 11) {
            if ("hello" > ("foo bar" + "")) {
                p0 = 0
                for fi11 := 0; fi11 < 6; fi11++ {
                    var v12 int = (-(fi11))
                    _ = fi11
                    _ = v12
                }
                v13 := []int{-2147483648, -2147483648, -2147483648, 42}
                _ = v13
            } else {
                fmt.Println(p0)
                var v14 uint32 = 103
                _ = v14
            }
            if true {
                var v15 uint16 = (p0 << ((255) & 63))
                for fi16 := 0; fi16 < 16; fi16++ {
                    v17 := []int{fi16, fi16, (fi16 * 4294967295)}
                    _ = fi16
                    _ = v17
                }
                v18 := []int{-2147483648, -153, 404, 0}
                _ = v15
                _ = v18
            }
            v19 := GenS1{F0: -1}
            _ = v19
        } else {
            p0 -= (p0 >> ((2) & 63))
            var v20 int64 = -651
            _ = v20
        }
    }
    var v21 uint64 = 526
    switch 2147483647 {
    case 32767:
        switch 148 {
        case 32768:
            var v22 uint16 = p0
            _ = v22
        }
    case 399:
        fmt.Println(v21, v21)
    default:
        v23 := []int{-683, int(p0), 32767, 318}
        _ = v23
    }
    _ = p0
    _ = v21
    return 4294967295
}
func genFn1() bool {
    if (!true) {
        if (0 <= -1) {
            var v24 string = "a"
            _ = v24
        } else {
            var v25 uint64 = 2
            var v26 uint32 = 42
            fmt.Println(v25, v26)
            _ = v25
            _ = v26
        }
        v27 := GenS1{F0: 42}
        var v28 uint16 = 800
        _ = v27
        _ = v28
    } else {
        var v29 uint64 = 541
        _ = v29
    }
    for fi30 := 0; fi30 < 5; fi30++ {
        fmt.Println(fi30)
        var v31 int8 = 34
        v32 := GenS1{F0: 1000}
        _ = fi30
        _ = v31
        _ = v32
    }
    v33 := []int{-853, -252, 2147483647, 1}
    var v34 bool = false
    if (!v34) {
        for fi35 := 0; fi35 < 13; fi35++ {
            switch (fi35 | 978) {
            case -9223372036854775808:
                v34 = v34
            case 32768:
                for fi36 := 0; fi36 < 18; fi36++ {
                    v37 := GenS1{F0: -1}
                    v34 = (!(v34 || (127 > 42)))
                    fmt.Println(fi35, fi35, v34)
                    _ = fi36
                    _ = v37
                }
            case 9223372036854775807:
                switch "\n\t" {
                case "hello":
                    var v38 int32 = 940
                    _ = v38
                }
            default:
                v34 = ((fi35 > fi35) || (false && (2147483647 > 533)))
            }
            for fi39 := 0; fi39 < 4; fi39++ {
                var v40 uint16 = 2
                switch ("foo bar" + ("hello" + "")) {
                case "hello":
                    var v41 uint8 = 0
                    _ = v41
                case "":
                    fmt.Println(fi39, fi39)
                case "\n\t":
                    if len(v33) > 0 {
                        v33[0] = 127
                    }
                }
                _ = fi39
                _ = v40
            }
            _ = fi35
        }
        v34 = v34
    }
    _ = v33
    _ = v34
    return (!(128 > 196))
}

func main() {
    v1 := make(map[string]int)
    var v2 int32 = 1
    var v3 bool = false
    v42 := GenS2{F0: 2, F1: v3}
    switch 32767 {
    case 32768:
        fmt.Println(v3, v3)
    }
    v43 := GenS3{F0: 127, F1: v2, F2: v3}
    v44 := GenS4{F0: 127, F1: 4294967295, F2: 65535}
    var v45 uint64 = 18446744073709551615
    for fi46 := 0; fi46 < 10; fi46++ {
        for fi47 := 0; fi47 < 10; fi47++ {
            v48 := []int{(fi46 % ((fi47) | 1)), fi46, (fi46 - (fi46 / ((32768) | 1)))}
            _ = fi47
            _ = v48
        }
        var v49 string = ("a" + "")
        var v50 int8 = 79
        _ = fi46
        _ = v49
        _ = v50
    }
    var v51 int64 = 1
    var v52 int32 = 42
    if (!(false && v3)) {
        switch 2147483647 {
        case -187:
            if ((^(v45)) <= (v45 / (((v45 >> ((4294967295) & 63))) | 1))) {
                v3 = true
                switch "a" {
                case "hello":
                    if v3 {
                        if (-1 == -856) {
                            var v53 float64 = 1e10
                            v2 += (-(v52))
                            _ = v53
                        } else {
                            v1["foo bar"] = -310
                            v54 := GenS5{F0: 334}
                            _ = v54
                        }
                        var v55 int32 = -2147483648
                        _ = v55
                    }
                }
                if (4294967295 < genFn0(480)) {
                    v2 = (^(v2))
                } else {
                    v52 = v2
                }
            } else {
                v56 := GenS6{F0: 839}
                v1[("a" + "foo bar")] = 12
                v57 := v1["a"]
                _ = v56
                _ = v57
            }
        case 2147483647:
            if (42 == 1) {
                var v58 uint64 = (v45 & v45)
                _ = v58
            } else {
                fmt.Println(v2, v51)
            }
        }
        var v59 uint = 255
        _ = v59
    }
    switch 9223372036854775807 {
    case 680:
        for fi60 := 0; fi60 < 11; fi60++ {
            delete(v1, "a")
            _ = fi60
        }
    case -1:
        for fi61 := 0; fi61 < 10; fi61++ {
            v62 := GenS4{F0: 127, F1: (v45 >> ((42) & 63)), F2: 392}
            _ = fi61
            _ = v62
        }
    case 32768:
        delete(v1, "foo bar")
    }
    var v63 int64 = -609
    v52 += (-(v2))
    if true {
        if (uint(v45) > 596) {
            switch -9223372036854775808 {
            case 32768:
                var v64 int = 9223372036854775807
                _ = v64
            default:
                fmt.Println(v3, v51, v63)
            }
            var v65 uint16 = 42
            _ = v65
        }
        v52 += (v2 + v2)
        if v3 {
            for fi66 := 0; fi66 < 8; fi66++ {
                var v67 int64 = (v51 - (v51 | v63))
                if v3 {
                    v1[(("hello" + (("foo bar" + "") + (("\n\t" + "a") + ""))) + "hello")] = fi66
                    var v68 bool = (v51 != (-(v51)))
                    switch "\n\t" {
                    case "hello":
                        if (!((fi66 - (-(fi66))) >= (fi66 & (fi66 + (fi66 % ((fi66) | 1)))))) {
                            switch ("a" + "") {
                            case "":
                                v69 := GenS7{F0: 236}
                                _ = v69
                            case "hello":
                                if v3 {
                                    v52 = v2
                                    v70 := []int{(fi66 - fi66), -867, 127}
                                    for fi71 := 0; fi71 < 15; fi71++ {
                                        var v72 int64 = v51
                                        var v73 int8 = 1
                                        _ = fi71
                                        _ = v72
                                        _ = v73
                                    }
                                    _ = v70
                                }
                            }
                            var v74 uint32 = 2147483648
                            delete(v1, "")
                            _ = v74
                        }
                    default:
                        var v75 uint64 = v45
                        _ = v75
                    }
                    _ = v68
                } else {
                    fmt.Println(v52)
                }
                _ = fi66
                _ = v67
            }
            switch 127 {
            case -637:
                v76 := []int{}
                _ = v76
            default:
                v51 += v63
            }
            fmt.Println(v52, v52, v52)
        } else {
            var v77 bool = v3
            for fi78 := 0; fi78 < 20; fi78++ {
                var v79 string = ""
                for fi80 := 0; fi80 < 11; fi80++ {
                    v51 += (v51 % ((v63) | 1))
                    var v81 int8 = 119
                    _ = fi80
                    _ = v81
                }
                var v82 uint64 = (v45 - (v45 >> ((0) & 63)))
                _ = fi78
                _ = v79
                _ = v82
            }
            switch "hello" {
            case "":
                v63 -= v51
            default:
                v45 -= v45
            }
            _ = v77
        }
    } else {
        _ = len(v1)
    }
    fmt.Println(v45, v2, v45)
    for fi83 := 0; fi83 < 7; fi83++ {
        var v84 uint64 = (v45 << ((781) & 63))
        _ = fi83
        _ = v84
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v45:", v45)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v51
    _ = v52
    _ = v63
}
