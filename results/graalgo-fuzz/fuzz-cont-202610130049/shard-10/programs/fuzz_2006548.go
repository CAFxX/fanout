package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 string; F1 uint64; F2 uint }
type GenS3 struct { F0 uint; F1 int16; F2 uint64 }
type GenS4 struct { F0 int }
func genFn0(p0 uint32) int64 {
    var v4 int32 = 42
    for fi5 := 0; fi5 < 1; fi5++ {
        fmt.Println(fi5)
        var v6 int8 = 34
        _ = fi5
        _ = v6
    }
    v4 += v4
    for fi7 := 0; fi7 < 12; fi7++ {
        fmt.Println(fi7, fi7, v4)
        _ = fi7
    }
    _ = p0
    _ = v4
    return 625
}
func genFn1(p0 string, p1 uint, p2 uint32) bool {
    var v8 float64 = (0.1 * -2.25)
    p2 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return false
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]uint32)
    var v3 uint32 = 4294967295
    var v9 int32 = 42
    var v10 int32 = v9
    switch ("a" + ("hello" + "foo bar")) {
    case "foo bar":
        for fi11 := 0; fi11 < 15; fi11++ {
            v3 = (v3 & v3)
            var v12 float64 = 3.14159
            _ = fi11
            _ = v12
        }
    }
    if (!true) {
        v9 = (v9 % ((1) | 1))
        v9 = v10
        for fi13 := 0; fi13 < 13; fi13++ {
            var v14 uint64 = 0
            v15 := []int{(fi13 % ((2147483647) | 1))}
            if (0 > 0) {
                v3 -= (v3 + (v3 >> ((42) & 63)))
            } else {
                if true {
                    v16 := v1["foo bar"]
                    var v17 uint = 42
                    _ = v16
                    _ = v17
                }
                var v18 float32 = float32(3.14159)
                _ = v18
            }
            _ = fi13
            _ = v14
            _ = v15
        }
    }
    v3 = v3
    fmt.Println(v3, v10)
    var v19 uint = 18446744073709551615
    switch -723 {
    case 127:
        var v20 float32 = float32(float64(0.1))
        _ = v20
    default:
        if ((v19 | v19) <= (v19 % (((v19 * (v19 + v19))) | 1))) {
            fmt.Println(v19, v9, v3)
            for fi21 := 0; fi21 < 3; fi21++ {
                if ((false && false) || false) {
                    fmt.Println(v3, v3, v9)
                } else {
                    switch fi21 {
                    case 4294967295:
                        if (1.5 != float32(3.14159)) {
                            v22 := GenS1{F0: v19, F1: 260}
                            _ = v22
                        } else {
                            v23 := GenS1{F0: (v19 >> (((v19 + v19)) & 63)), F1: 64}
                            v24 := GenS2{F0: "a", F1: 581, F2: (v19 & v19)}
                            switch "hello" {
                            case "\n\t":
                                v25 := []int{}
                                _ = v25
                            }
                            _ = v23
                            _ = v24
                        }
                    default:
                        v3 -= (v3 >> ((v19) & 63))
                    }
                }
                for fi26 := 0; fi26 < 14; fi26++ {
                    fmt.Println(fi21, v3, fi26)
                    v27 := []int{fi26, (fi21 >> (((v19 - v19)) & 63))}
                    fmt.Println(v9)
                    fmt.Println(fi21, v10, v19)
                    fmt.Println(v3, fi21, v10)
                    _ = fi26
                    _ = v27
                }
                switch fi21 {
                case 0:
                    var v28 string = "a"
                    _ = v28
                case 32768:
                    if true {
                        if (!false) {
                            v29 := GenS2{F0: "", F1: uint64(v3), F2: v19}
                            _ = v29
                        }
                        fmt.Println(v19, fi21, v10)
                    }
                default:
                    var v30 int16 = 42
                    _ = v30
                }
                _ = fi21
            }
        } else {
            for fi31 := 0; fi31 < 17; fi31++ {
                switch fi31 {
                case 127:
                    v32 := GenS2{F0: "hello", F1: 4294967295, F2: v19}
                    _ = v32
                case 32768:
                    var v33 int16 = 419
                    _ = v33
                }
                v34 := []int{4294967295}
                var v35 int = 4294967295
                _ = fi31
                _ = v34
                _ = v35
            }
            if true {
                var v36 uint = v19
                for fi37 := 0; fi37 < 6; fi37++ {
                    v1["foo bar"] = (("foo bar" + "a") + ("\n\t" + "a"))
                    _ = fi37
                }
                _ = v36
            } else {
                switch -506 {
                case 9223372036854775807:
                    v38 := v1["a"]
                    _ = v38
                case -234:
                    var v39 int64 = -9223372036854775808
                    _ = v39
                case 42:
                    switch "a" {
                    case "foo bar":
                        v3 = (^(v3))
                    case "a":
                        v3 = (v3 - v3)
                    case "":
                        var v40 int = 489
                        _ = v40
                    }
                default:
                    if true {
                        var v41 int16 = -1
                        _ = v41
                    } else {
                        v19 -= 0
                    }
                }
                v1["hello"] = "hello"
                fmt.Println(v9)
            }
        }
    }
    switch 51 {
    case 42:
        switch -613 {
        case 2147483647:
            v3 -= (v3 % ((v3) | 1))
        default:
            v9 = v9
        }
    }
    v42 := GenS3{F0: (^(v19)), F1: 672, F2: 0}
    if true {
        if false {
            fmt.Println(v9)
            v1["\n\t"] = "hello"
            v43 := []int{4294967295}
            _ = v43
        }
        if true {
            fmt.Println(v3, v9, v3)
            if true {
                fmt.Println(v10)
                switch "hello" {
                case "\n\t":
                    fmt.Println(v10)
                case "a":
                    v44 := v1["foo bar"]
                    _ = v44
                default:
                    v45 := GenS3{F0: 558, F1: 786, F2: 1}
                    _ = v45
                }
            }
        } else {
            if false {
                var v46 uint = v19
                v10 -= v10
                _ = v46
            }
        }
        if true {
            v3 = 462
            var v47 string = ("foo bar" + "")
            _ = v47
        } else {
            for fi48 := 0; fi48 < 11; fi48++ {
                fmt.Println(v3, v3)
                fmt.Println(v19, v10, v10)
                if (!true) {
                    v49 := []int{-2147483648, fi48, 32768}
                    _ = len(v49)
                    for fi50 := 0; fi50 < 13; fi50++ {
                        v9 = -223
                        v51 := GenS4{F0: fi50}
                        v9 += v9
                        var v52 int32 = (^(v9))
                        _ = fi50
                        _ = v51
                        _ = v52
                    }
                    _ = v49
                } else {
                    for fi53 := 0; fi53 < 9; fi53++ {
                        var v54 string = "foo bar"
                        var v55 int16 = 1000
                        var v56 int = fi53
                        _ = fi53
                        _ = v54
                        _ = v55
                        _ = v56
                    }
                    var v57 int32 = (v10 & v10)
                    _ = v57
                }
                _ = fi48
            }
            for fi58 := 0; fi58 < 10; fi58++ {
                if true {
                    v9 = (v9 / ((-2147483648) | 1))
                }
                if (32767 > 0) {
                    v59 := v2["\n\t"]
                    _ = v59
                }
                var v60 int32 = (^(v9))
                _ = fi58
                _ = v60
            }
        }
    } else {
        switch "\n\t" {
        case "a":
            v61 := GenS3{F0: v19, F1: -95, F2: 2}
            _ = v61
        case "hello":
            var v62 uint16 = 188
            _ = v62
        }
        v10 = (v9 | v9)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v3
    _ = v42
    _ = v9
}
