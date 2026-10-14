package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint64; F1 int64; F2 int16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint64; F1 uint64 }
type GenS7 struct { F0 uint32; F1 bool }
type GenS8 struct { F0 bool; F1 uint16; F2 bool }
func genFn0(p0 uint64, p1 uint64) uint16 {
    if true {
        switch "foo bar" {
        case "hello":
            for fi4 := 0; fi4 < 3; fi4++ {
                fmt.Println(fi4, fi4)
                _ = fi4
            }
        case "":
            switch ("foo bar" + "\n\t") {
            case "a":
                fmt.Println(p1, p0)
            case "hello":
                fmt.Println(p1)
            default:
                var v5 string = ((("foo bar" + (("a" + "\n\t") + (("hello" + "foo bar") + ("foo bar" + "a")))) + "hello") + (("foo bar" + "foo bar") + "a"))
                _ = v5
            }
        case "a":
            if false {
                for fi6 := 0; fi6 < 14; fi6++ {
                    fmt.Println(fi6, p0)
                    if false {
                        v7 := GenS1{F0: 1, F1: (^(p1)), F2: ""}
                        _ = v7
                    }
                    _ = fi6
                }
                fmt.Println(p0)
                var v8 int8 = -1
                _ = v8
            } else {
                for fi9 := 0; fi9 < 19; fi9++ {
                    if ((723 == 459) || (3.14159 == 1e10)) {
                        var v10 uint64 = (p0 + p0)
                        _ = v10
                    } else {
                        for fi11 := 0; fi11 < 1; fi11++ {
                            v12 := GenS1{F0: 18446744073709551615, F1: p0, F2: (("foo bar" + "\n\t") + (("\n\t" + ("foo bar" + (("a" + (("a" + ("foo bar" + ("foo bar" + ""))) + ("a" + ("\n\t" + "\n\t")))) + "\n\t"))) + "\n\t"))}
                            var v13 uint8 = 42
                            if true {
                                if true {
                                    var v14 int = (fi9 % ((32767) | 1))
                                    _ = v14
                                }
                                p1 -= p1
                                var v15 bool = true
                                _ = v15
                            }
                            _ = fi11
                            _ = v12
                            _ = v13
                        }
                        v16 := []int{(fi9 << ((0) & 63)), (fi9 ^ (fi9 >> ((960) & 63)))}
                        v17 := GenS1{F0: 42, F1: (^(p0)), F2: "\n\t"}
                        _ = v16
                        _ = v17
                    }
                    _ = fi9
                }
                if true {
                    if false {
                        p1 = (^(p0))
                    } else {
                        if (false && true) {
                            v18 := []int{403, 127, 9223372036854775807}
                            _ = v18
                        }
                        p0 = (p0 | p1)
                    }
                    v19 := GenS1{F0: 671, F1: (^(p1)), F2: "foo bar"}
                    p1 -= (p1 ^ p0)
                    _ = v19
                } else {
                    v20 := GenS2{F0: ("foo bar" + "\n\t")}
                    _ = v20
                }
            }
        default:
            if (1 >= 255) {
                v21 := []int{}
                _ = v21
            }
        }
        if true {
            if (2 < 127) {
                if true {
                    p1 -= p0
                } else {
                    for fi22 := 0; fi22 < 6; fi22++ {
                        v23 := GenS1{F0: 211, F1: p1, F2: "foo bar"}
                        p1 = p0
                        _ = fi22
                        _ = v23
                    }
                    var v24 float32 = 1.5
                    _ = v24
                }
                var v25 int = 1
                for fi26 := 0; fi26 < 18; fi26++ {
                    switch fi26 {
                    case 4294967295:
                        switch v25 {
                        case -1:
                            v27 := GenS3{F0: (p0 - p1), F1: -1, F2: 32767}
                            _ = v27
                        case 32768:
                            var v28 int16 = 1
                            _ = v28
                        case 42:
                            if ((662 > 0) && ((v25 ^ fi26) < v25)) {
                                p0 = 2
                            }
                        default:
                            var v29 string = "hello"
                            _ = v29
                        }
                    default:
                        var v30 int16 = 9
                        _ = v30
                    }
                    var v31 float32 = ((1.5 - (-2.25 * float32(0.1))) * float32((3.14159 - (0.1 + 1.5))))
                    v32 := GenS2{F0: ("foo bar" + "")}
                    _ = fi26
                    _ = v31
                    _ = v32
                }
                _ = v25
            } else {
                v33 := GenS4{F0: 1}
                _ = v33
            }
            fmt.Println(p0, p1)
        }
    }
    v34 := []int{-1, 300, 1}
    var v35 uint = 124
    if len(v34) > 0 {
        v34[0] = -612
    }
    _ = p0
    _ = p1
    _ = v34
    _ = v35
    return 65535
}
func genFn1() string {
    var v36 int8 = -1
    fmt.Println(v36)
    v36 = (v36 >> ((18446744073709551615) & 63))
    v37 := GenS5{F0: 0}
    _ = v36
    _ = v37
    return "a"
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]int)
    var v3 int64 = 230
    if false {
        for fi38 := 0; fi38 < 6; fi38++ {
            v39 := GenS2{F0: ("a" + "foo bar")}
            _ = fi38
            _ = v39
        }
        v3 = 1
    } else {
        v40 := GenS2{F0: "hello"}
        _ = v40
    }
    var v41 int8 = 18
    v41 = v41
    if true {
        if (false && (1 < 65535)) {
            switch -600 {
            case -1:
                delete(v1, "hello")
            default:
                for fi42 := 0; fi42 < 20; fi42++ {
                    v3 = -1
                    for fi43 := 0; fi43 < 13; fi43++ {
                        var v44 uint32 = 748
                        _ = fi43
                        _ = v44
                    }
                    if (!((-32768 > 384) && ((1e10 - 3.14159) == 3.14159))) {
                        fmt.Println(v3)
                    }
                    _ = fi42
                }
            }
            v1[("hello" + "\n\t")] = ("hello" + ("hello" + ""))
            v3 = v3
        }
    } else {
        for fi45 := 0; fi45 < 7; fi45++ {
            v3 = (^(v3))
            v3 += v3
            v46 := v2["hello"]
            _ = fi45
            _ = v46
        }
        v47 := []int{-495, 2147483647}
        fmt.Println(v41)
        _ = v47
    }
    var v48 bool = false
    for fi49 := 0; fi49 < 7; fi49++ {
        fmt.Println(v3)
        if false {
            v50 := GenS6{F0: 1, F1: 18446744073709551615}
            fmt.Println(v48)
            fmt.Println(v41, fi49)
            _ = v50
        } else {
            v51 := []int{(fi49 & -1), fi49, (fi49 ^ fi49), (^(fi49))}
            switch "" {
            case "foo bar":
                var v52 int64 = -1
                _ = v52
            default:
                v3 -= 42
            }
            _ = v51
        }
        _ = fi49
    }
    for fi53 := 0; fi53 < 3; fi53++ {
        v48 = false
        _ = fi53
    }
    fmt.Println(v41, v3, v41)
    delete(v1, "a")
    if v48 {
        switch -686 {
        case 0:
            if ((4294967295 > 0) && (!v48)) {
                delete(v1, ("foo bar" + ((("a" + "hello") + "\n\t") + "hello")))
                delete(v1, "\n\t")
                for fi54 := 0; fi54 < 16; fi54++ {
                    if v48 {
                        for fi55 := 0; fi55 < 13; fi55++ {
                            switch fi55 {
                            case -190:
                                v48 = (((!v48) || v48) && v48)
                            default:
                                var v56 uint = 2
                                _ = v56
                            }
                            _ = fi55
                        }
                    }
                    var v57 uint16 = 255
                    v3 -= v3
                    _ = fi54
                    _ = v57
                }
            }
        case -43:
            fmt.Println(v48)
        }
    } else {
        v58 := GenS2{F0: "a"}
        _ = v58
    }
    delete(v1, "\n\t")
    v1["foo bar"] = "foo bar"
    if (1 > 0) {
        v59 := GenS7{F0: 4294967295, F1: v48}
        var v60 uint8 = 42
        _ = v59
        _ = v60
    } else {
        _ = len(v2)
        var v61 float32 = (1.5 - 0.1)
        _ = v61
    }
    v62 := GenS8{F0: (!v48), F1: 255, F2: (!v48)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v41:", v41)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v3
    _ = v41
    _ = v48
    _ = v62
}
