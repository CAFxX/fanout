package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 int32; F1 bool; F2 int64 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 bool; F1 int16 }
type GenS6 struct { F0 int; F1 uint32 }
type GenS7 struct { F0 string; F1 int8; F2 string }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 uint; F1 int8; F2 int32 }
func genFn0(p0 bool, p1 bool, p2 uint64) uint16 {
    for fi4 := 0; fi4 < 12; fi4++ {
        for fi5 := 0; fi5 < 11; fi5++ {
            if ("foo bar" == ("\n\t" + "a")) {
                v6 := []int{}
                var v7 uint32 = 931
                _ = len(v6)
                _ = v6
                _ = v7
            }
            if p1 {
                p2 = 18446744073709551615
            } else {
                v8 := GenS1{F0: p1}
                v9 := []int{102, (fi5 >> ((4294967295) & 63))}
                _ = v8
                _ = v9
            }
            _ = fi5
        }
        _ = fi4
    }
    for fi10 := 0; fi10 < 14; fi10++ {
        var v11 string = ("hello" + ("hello" + "hello"))
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}
func genFn1(p0 int, p1 int64) int16 {
    switch (("\n\t" + "hello") + "a") {
    case "hello":
        var v12 float64 = 0.0
        _ = v12
    default:
        var v13 int64 = p1
        _ = v13
    }
    v14 := []int{(p0 + p0)}
    _ = p0
    _ = p1
    _ = v14
    return -660
}

func main() {
    var v1 uint8 = 2
    v2 := make(map[string]string)
    var v3 int16 = -414
    switch "a" {
    case "foo bar":
        v15 := GenS2{F0: 0, F1: 1}
        _ = v15
    case "hello":
        fmt.Println(v1, v3)
    default:
        v3 += (v3 ^ v3)
    }
    fmt.Println(v1, v3)
    fmt.Println(v3, v1)
    if (v3 <= (-(v3))) {
        var v16 string = (("a" + "\n\t") + "foo bar")
        for fi17 := 0; fi17 < 9; fi17++ {
            var v18 int16 = v3
            _ = fi17
            _ = v18
        }
        _ = v16
    } else {
        var v19 float32 = 0.0
        v1 = v1
        v19 = v19
        _ = v19
    }
    for fi20 := 0; fi20 < 9; fi20++ {
        switch ("a" + "") {
        case "\n\t":
            v21 := GenS3{F0: 558, F1: false, F2: -1}
            _ = v21
        case "foo bar":
            v3 = (^(v3))
        }
        v22 := GenS3{F0: -2147483648, F1: true, F2: 537}
        var v23 int16 = (v3 >> ((0) & 63))
        _ = fi20
        _ = v22
        _ = v23
    }
    v1 += (v1 >> ((2) & 63))
    for fi24 := 0; fi24 < 16; fi24++ {
        v25 := []int{(fi24 * (fi24 - fi24))}
        v1 = (v1 | (v1 >> ((255) & 63)))
        _ = fi24
        _ = v25
    }
    switch "" {
    case "foo bar":
        var v26 float32 = 1.5
        _ = v26
    case "":
        switch "hello" {
        case "\n\t":
            v2["foo bar"] = (("" + "hello") + "foo bar")
        case "a":
            switch "foo bar" {
            case "":
                var v27 uint16 = genFn0((!((v1 != v1) || (!true))), true, 735)
                _ = v27
            default:
                switch 122 {
                case -1:
                    var v28 uint16 = 575
                    _ = v28
                }
            }
        case "":
            for fi29 := 0; fi29 < 13; fi29++ {
                if true {
                    switch (fi29 + fi29) {
                    case 128:
                        var v30 uint64 = 42
                        _ = v30
                    case 42:
                        for fi31 := 0; fi31 < 12; fi31++ {
                            if true {
                                var v32 float32 = float32(1.5)
                                _ = v32
                            }
                            _ = fi31
                        }
                    }
                    v33 := GenS2{F0: 42, F1: (v3 / (((-(v3))) | 1))}
                    var v34 uint64 = uint64(v1)
                    _ = v33
                    _ = v34
                } else {
                    v3 = v3
                    if false {
                        v35 := GenS4{F0: "foo bar"}
                        switch ("foo bar" + "hello") {
                        case "a":
                            var v36 int16 = (v3 % ((v3) | 1))
                            _ = v36
                        case "foo bar":
                            var v37 uint32 = 2147483647
                            _ = v37
                        }
                        switch -165 {
                        case 2147483647:
                            var v38 string = ("" + (("a" + "\n\t") + ""))
                            _ = v38
                        }
                        _ = v35
                    } else {
                        v39 := GenS5{F0: true, F1: genFn1(fi29, 1)}
                        fmt.Println(v1)
                        v3 = v3
                        _ = v39
                    }
                    v1 -= v1
                }
                fmt.Println(v1, v1)
                _ = fi29
            }
        }
    }
    if (!false) {
        if (1.5 <= (3.14159 / (1.5 + 0.1))) {
            for fi40 := 0; fi40 < 17; fi40++ {
                var v41 int64 = 0
                v42 := []int{-2147483648, (fi40 >> ((42) & 63))}
                v41 += v41
                _ = fi40
                _ = v41
                _ = v42
            }
        } else {
            var v43 int32 = -1
            _ = v43
        }
        v2[""] = ("foo bar" + ((("\n\t" + ("" + ("hello" + ("a" + ("\n\t" + "hello"))))) + "") + (((("hello" + ("\n\t" + ("" + "hello"))) + ("hello" + "\n\t")) + "hello") + (("" + "\n\t") + ("" + "")))))
        var v44 bool = true
        _ = v44
    } else {
        v45 := GenS6{F0: 2147483647, F1: 2}
        if false {
            var v46 int8 = 120
            v47 := GenS7{F0: ("hello" + "a"), F1: (v46 - int8(v1)), F2: (("\n\t" + "\n\t") + (("hello" + ("foo bar" + ("hello" + ""))) + ""))}
            _ = v46
            _ = v47
        }
        for fi48 := 0; fi48 < 1; fi48++ {
            if true {
                v49 := GenS2{F0: -87, F1: v3}
                v1 += v1
                _ = v49
            } else {
                v3 += genFn1((fi48 * fi48), 0)
                for fi50 := 0; fi50 < 2; fi50++ {
                    v51 := GenS8{F0: 65535}
                    _ = fi50
                    _ = v51
                }
            }
            v3 += -98
            _ = fi48
        }
        _ = v45
    }
    var v52 string = (("hello" + "\n\t") + "")
    switch v52 {
    case "a":
        var v53 uint8 = 42
        _ = v53
    case "hello":
        v3 += genFn1(1, 0)
    case "":
        var v54 uint = 42
        _ = v54
    default:
        for fi55 := 0; fi55 < 20; fi55++ {
            if true {
                var v56 int64 = 135
                if false {
                    for fi57 := 0; fi57 < 10; fi57++ {
                        var v58 uint16 = 0
                        fmt.Println(fi55, v56)
                        var v59 uint16 = v58
                        v56 = (^(v56))
                        var v60 uint16 = v58
                        _ = fi57
                        _ = v58
                        _ = v59
                        _ = v60
                    }
                    switch (fi55 + fi55) {
                    case 0:
                        v61 := []int{}
                        _ = v61
                    case -406:
                        var v62 int8 = -128
                        _ = v62
                    case -1:
                        v63 := GenS9{F0: 65535, F1: -1, F2: 2147483647}
                        _ = v63
                    }
                }
                if true {
                    for fi64 := 0; fi64 < 12; fi64++ {
                        fmt.Println(fi64, v52)
                        v2["\n\t"] = v52
                        _ = fi64
                    }
                    v65 := []int{809, fi55}
                    _ = v65
                }
                _ = v56
            }
            _ = fi55
        }
    }
    v1 += v1
    delete(v2, v52)
    switch -1 {
    case 9223372036854775807:
        for fi66 := 0; fi66 < 10; fi66++ {
            v67 := GenS1{F0: (false && true)}
            _ = fi66
            _ = v67
        }
    default:
        var v68 uint16 = 255
        _ = v68
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v52
}
