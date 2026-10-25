package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 uint64 }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 uint16; F1 int32 }
type GenS4 struct { F0 string }
func genFn0() uint {
    v6 := []int{}
    var v7 int16 = 648
    if (true || true) {
        v7 = (-(v7))
        for fi8 := 0; fi8 < 12; fi8++ {
            var v9 uint64 = 0
            _ = fi8
            _ = v9
        }
        switch 127 {
        case 42:
            switch 4294967295 {
            case -989:
                var v10 uint8 = 255
                _ = v10
            case -9223372036854775808:
                var v11 uint64 = 1
                _ = v11
            }
        case 0:
            if true {
                var v12 int32 = -1
                for fi13 := 0; fi13 < 17; fi13++ {
                    fmt.Println(fi13, v7, fi13)
                    switch 4294967295 {
                    case 4294967295:
                        fmt.Println(fi13, fi13)
                    case 117:
                        var v14 uint32 = 0
                        _ = v14
                    default:
                        var v15 string = ("a" + ("foo bar" + "a"))
                        _ = v15
                    }
                    _ = fi13
                }
                _ = v12
            }
        }
    } else {
        if len(v6) > 0 {
            v6[0] = 4294967295
        }
    }
    fmt.Println(v7)
    fmt.Println(v7)
    _ = v6
    _ = v7
    return 18446744073709551615
}
func genFn1(p0 int64, p1 uint16) int32 {
    for fi16 := 0; fi16 < 13; fi16++ {
        switch fi16 {
        case -400:
            p1 = p1
        case 777:
            switch ("\n\t" + "a") {
            case "":
                for fi17 := 0; fi17 < 7; fi17++ {
                    v18 := GenS1{F0: "foo bar", F1: 18446744073709551615, F2: 124}
                    _ = fi17
                    _ = v18
                }
            case "hello":
                v19 := []int{}
                _ = v19
            default:
                fmt.Println(p0, p1, fi16)
            }
        }
        for fi20 := 0; fi20 < 6; fi20++ {
            switch "a" {
            case "\n\t":
                p0 += 941
            case "foo bar":
                var v21 string = "\n\t"
                _ = v21
            case "a":
                if (true || false) {
                    switch "a" {
                    case "a":
                        v22 := []int{}
                        _ = v22
                    case "foo bar":
                        p1 += 383
                    case "\n\t":
                        switch fi16 {
                        case -1:
                            var v23 int16 = 1
                            _ = v23
                        case 4294967295:
                            var v24 uint16 = p1
                            _ = v24
                        }
                    }
                } else {
                    var v25 uint64 = 560
                    p0 -= (p0 << ((2) & 63))
                    _ = v25
                }
            }
            if false {
                if true {
                    switch fi20 {
                    case 2147483647:
                        var v26 string = ("a" + "\n\t")
                        _ = v26
                    case -9223372036854775808:
                        var v27 string = ""
                        _ = v27
                    case -844:
                        fmt.Println(fi20, p0)
                    default:
                        p0 -= (-(p0))
                        if true {
                            v28 := []int{}
                            v29 := GenS1{F0: (("foo bar" + "hello") + "a"), F1: 2, F2: 674}
                            v30 := GenS1{F0: ("" + "hello"), F1: 1, F2: 344}
                            _ = v28
                            _ = v29
                            _ = v30
                        } else {
                            var v31 uint16 = (p1 | (^(p1)))
                            v31 = p1
                            _ = v31
                        }
                    }
                } else {
                    v32 := GenS1{F0: "\n\t", F1: 18446744073709551615, F2: 42}
                    p0 = p0
                    _ = v32
                }
                p0 += (p0 % (((p0 >> ((2) & 63))) | 1))
            }
            fmt.Println(fi20, p1)
            _ = fi20
        }
        p0 += 48
        _ = fi16
    }
    var v33 uint64 = 2
    _ = p0
    _ = p1
    _ = v33
    return 2147483647
}

func main() {
    v1 := []uint32{}
    v2 := make(map[string]uint32)
    var v3 int16 = 754
    var v4 int8 = 100
    var v5 float64 = (1.5 + 0.0)
    var v34 int64 = 539
    switch -600 {
    case -9223372036854775808:
        _ = len(v1)
    }
    v3 -= (v3 + v3)
    for fi35 := 0; fi35 < 19; fi35++ {
        v1 = append(v1, 2147483647)
        v4 = 1
        for fi36 := 0; fi36 < 20; fi36++ {
            if true {
                if (v5 >= 1.5) {
                    v37 := GenS1{F0: "hello", F1: 990, F2: 18446744073709551615}
                    var v38 float64 = ((v5 - 3.14159) - ((v5 - (1e10 * v5)) - v5))
                    v39 := GenS1{F0: "", F1: 4294967295, F2: 42}
                    _ = v37
                    _ = v38
                    _ = v39
                }
                if (65535 > 255) {
                    v34 += (v34 | v34)
                } else {
                    var v40 int32 = 42
                    var v41 uint = 517
                    _ = v40
                    _ = v41
                }
            } else {
                var v42 uint16 = 0
                v3 -= v3
                _ = v42
            }
            var v43 uint = genFn0()
            _ = fi36
            _ = v43
        }
        _ = fi35
    }
    delete(v2, ("hello" + "foo bar"))
    fmt.Println(v4)
    if true {
        v2["foo bar"] = 42
        v44 := GenS2{F0: (-(v3)), F1: ("hello" + ("hello" + "hello"))}
        _ = v44
    } else {
        var v45 int32 = -1
        v1 = append(v1, 819)
        _ = v45
    }
    _ = len(v2)
    switch "a" {
    case "hello":
        switch "foo bar" {
        case "":
            var v46 int64 = v34
            _ = v46
        case "hello":
            switch -547 {
            case -9223372036854775808:
                v47 := GenS2{F0: v3, F1: ""}
                _ = v47
            case 4294967295:
                var v48 float64 = (v5 + -2.25)
                _ = v48
            case 554:
                fmt.Println(v4, v5)
            }
        default:
            v49 := GenS1{F0: "\n\t", F1: 42, F2: 4294967295}
            _ = v49
        }
    case "foo bar":
        switch 4294967295 {
        case 0:
            fmt.Println(v5)
        }
    }
    for fi50 := 0; fi50 < 18; fi50++ {
        if (false && (4294967295 <= 2147483648)) {
            v51 := GenS2{F0: -1, F1: ""}
            delete(v2, ("foo bar" + ""))
            var v52 uint8 = 1
            _ = v51
            _ = v52
        } else {
            v53 := GenS3{F0: 528, F1: 2147483647}
            v34 = v34
            _ = v53
        }
        _ = fi50
    }
    _ = len(v1)
    switch ((("hello" + ((("foo bar" + ("" + "")) + "a") + ((("\n\t" + "a") + ((("" + "a") + "foo bar") + "")) + ("" + (("hello" + "hello") + "\n\t"))))) + "\n\t") + ("" + ((("hello" + "hello") + "foo bar") + "foo bar"))) {
    case "\n\t":
        var v54 int16 = v3
        _ = v54
    case "a":
        v34 = v34
    default:
        fmt.Println(v34, v34, v34)
    }
    var v55 int32 = 2147483647
    v4 -= (v4 - (v4 >> ((1) & 63)))
    if true {
        for fi56 := 0; fi56 < 13; fi56++ {
            fmt.Println(v3)
            switch "a" {
            case "\n\t":
                var v57 uint8 = 19
                _ = v57
            }
            var v58 uint16 = 0
            _ = fi56
            _ = v58
        }
        var v59 int8 = int8(v3)
        var v60 int32 = (v55 >> ((255) & 63))
        _ = v59
        _ = v60
    } else {
        var v61 bool = (122 == 127)
        _ = v61
    }
    v62 := GenS4{F0: (("\n\t" + "") + "a")}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v55
    _ = v62
}
