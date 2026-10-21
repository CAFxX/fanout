package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int64; F1 int; F2 int }
type GenS6 struct { F0 string }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 uint; F1 uint }
func genFn0(p0 uint16, p1 int8) int8 {
    var v5 int64 = -9223372036854775808
    v6 := []int{0, 127, -1}
    var v7 int8 = (p1 + (p1 % ((int8(p0)) | 1)))
    var v8 uint8 = 128
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return (v7 + (p1 << ((255) & 63)))
}
func genFn1() uint16 {
    var v9 int32 = 944
    if (v9 == 139) {
        var v10 bool = true
        fmt.Println(v9)
        v10 = (0 >= 42)
        _ = v10
    } else {
        v11 := []int{-9223372036854775808, 0}
        if ((!false) && ("a" == "")) {
            var v12 string = "\n\t"
            _ = v12
        } else {
            var v13 int = 653
            if len(v11) >= 2 {
                v11 = v11[1:len(v11)]
            }
            if ((992 > 4294967295) || (!(517 != v9))) {
                if false {
                    v9 += (^(v9))
                }
                for fi14 := 0; fi14 < 8; fi14++ {
                    if ((127 < 2) && false) {
                        var v15 string = "\n\t"
                        _ = v15
                    }
                    var v16 uint64 = 529
                    if len(v11) > 0 {
                        v11[0] = 601
                    }
                    _ = fi14
                    _ = v16
                }
                fmt.Println(v9, v13)
            } else {
                fmt.Println(v9, v9)
            }
            _ = v13
        }
        _ = v11
    }
    v17 := GenS1{F0: -553}
    v9 = 0
    v18 := GenS2{F0: (v9 & v9), F1: 255}
    _ = v17
    _ = v18
    _ = v9
    return 65535
}

func main() {
    v1 := []string{"foo bar", (("\n\t" + ("a" + ("a" + "hello"))) + ("\n\t" + ("foo bar" + ("a" + ("hello" + "a"))))), "hello", "hello"}
    v2 := make(map[int]int)
    v3 := make(map[uint32]string)
    var v4 uint64 = 2
    var v19 int64 = 9223372036854775807
    v19 = (v19 ^ (v19 + v19))
    v19 = -389
    var v20 float64 = 0.0
    var v21 uint = 1
    for fi22 := 0; fi22 < 10; fi22++ {
        var v23 uint8 = 0
        _ = fi22
        _ = v23
    }
    v24 := GenS3{F0: ("\n\t" + ("foo bar" + "a"))}
    switch (("hello" + "foo bar") + "a") {
    case "\n\t":
        var v25 uint16 = genFn1()
        _ = v25
    case "hello":
        if (214 > v19) {
            v26 := GenS4{F0: 0}
            for fi27 := 0; fi27 < 10; fi27++ {
                if true {
                    switch ("foo bar" + "\n\t") {
                    case "foo bar":
                        if (0 == 255) {
                            v28 := v2[fi27]
                            v1 = append(v1, ("" + "a"))
                            var v29 int64 = (v19 << ((304) & 63))
                            _ = v28
                            _ = v29
                        } else {
                            v19 = 9223372036854775807
                        }
                    case "":
                        var v30 bool = (v4 > (^(v4)))
                        _ = v30
                    case "hello":
                        for fi31 := 0; fi31 < 5; fi31++ {
                            var v32 uint8 = 2
                            v20 -= (0.0 + v20)
                            _ = fi31
                            _ = v32
                        }
                    default:
                        v33 := GenS5{F0: 0, F1: 780, F2: 32767}
                        _ = v33
                    }
                }
                switch "" {
                case "foo bar":
                    for fi34 := 0; fi34 < 14; fi34++ {
                        v21 -= (v21 * (v21 % ((v21) | 1)))
                        _ = fi34
                    }
                }
                var v35 float32 = 1e10
                _ = fi27
                _ = v35
            }
            _ = v26
        } else {
            if true {
                if true {
                    for fi36 := 0; fi36 < 3; fi36++ {
                        switch fi36 {
                        case 1:
                            if (false || false) {
                                var v37 int32 = -2147483648
                                _ = v37
                            }
                        case 0:
                            v21 -= v21
                        case 42:
                            var v38 uint32 = 4294967295
                            _ = v38
                        }
                        var v39 int64 = (v19 & 91)
                        _ = fi36
                        _ = v39
                    }
                }
                v40 := GenS6{F0: (((("a" + "hello") + "a") + "") + ("" + ""))}
                fmt.Println(v19, v4, v4)
                _ = v40
            }
            for fi41 := 0; fi41 < 1; fi41++ {
                v42 := GenS7{F0: "hello"}
                v21 = (v21 + (v21 ^ v21))
                var v43 bool = true
                _ = fi41
                _ = v42
                _ = v43
            }
            v21 = (v21 << ((v21) & 63))
        }
    case "foo bar":
        switch "" {
        case "":
            var v44 uint64 = (v4 ^ (v4 << ((255) & 63)))
            _ = v44
        }
    default:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    if false {
        for fi45 := 0; fi45 < 20; fi45++ {
            for fi46 := 0; fi46 < 2; fi46++ {
                if true {
                    var v47 float32 = (1.5 - float32(v20))
                    switch "\n\t" {
                    case "\n\t":
                        if (false || false) {
                            switch ("hello" + "a") {
                            case "a":
                                if true {
                                    fmt.Println(fi45, fi46)
                                    fmt.Println(v20)
                                } else {
                                    var v48 bool = true
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    switch (fi45 << ((v21) & 63)) {
                                    case 0:
                                        var v49 int = (fi46 % ((fi45) | 1))
                                        _ = v49
                                    default:
                                        v3[42] = (("" + "\n\t") + ("hello" + "hello"))
                                    }
                                    _ = v48
                                }
                            case "\n\t":
                                var v50 int8 = 100
                                _ = v50
                            }
                        } else {
                            if false {
                                var v51 string = "hello"
                                _ = v51
                            } else {
                                v52 := GenS6{F0: "\n\t"}
                                v53 := GenS7{F0: ("a" + "a")}
                                _ = v52
                                _ = v53
                            }
                            v20 = (v20 * (0.1 - (float64((-2.25 / -2.25)) / 3.14159)))
                        }
                    default:
                        var v54 string = ((((("" + "foo bar") + "hello") + "") + ("foo bar" + (("hello" + "foo bar") + "hello"))) + ((("hello" + (("" + "") + "a")) + (((("a" + ("\n\t" + "")) + "hello") + "hello") + "a")) + ""))
                        _ = v54
                    }
                    _ = len(v1)
                    fmt.Println(v20, v20)
                    v20 = v20
                    _ = v47
                }
                if (!(-25 > 42)) {
                    switch ("\n\t" + "a") {
                    case "":
                        if len(v1) > 0 {
                            v1[0] = ((((("a" + "\n\t") + (("hello" + "hello") + "\n\t")) + (("" + "hello") + (("\n\t" + "\n\t") + ""))) + ("\n\t" + ("\n\t" + "foo bar"))) + "hello")
                        }
                        fmt.Println(v19)
                    default:
                        _ = len(v1)
                    }
                    var v55 int32 = 290
                    fmt.Println(fi45, v21, v55)
                    var v56 uint32 = 2147483647
                    delete(v2, (fi46 % ((int(v21)) | 1)))
                    _ = v55
                    _ = v56
                }
                v19 = (v19 ^ 681)
                _ = fi46
            }
            _ = fi45
        }
    }
    _ = len(v1)
    for fi57 := 0; fi57 < 18; fi57++ {
        if true {
            switch "\n\t" {
            case "a":
                v20 -= (v20 * float64(float32(float64(1.5))))
            case "foo bar":
                if false {
                    v20 += float64(-2.25)
                    var v58 int = 128
                    _ = v58
                } else {
                    var v59 uint8 = 117
                    _ = v59
                }
            case "\n\t":
                switch "foo bar" {
                case "hello":
                    var v60 uint8 = 42
                    _ = v60
                default:
                    for fi61 := 0; fi61 < 9; fi61++ {
                        fmt.Println(fi61, fi61)
                        switch (("foo bar" + "hello") + "") {
                        case "":
                            var v62 int16 = -821
                            _ = v62
                        case "foo bar":
                            delete(v2, 32768)
                        }
                        var v63 uint = 206
                        v19 += (v19 ^ v19)
                        var v64 float64 = (float64(3.14159) * v20)
                        v19 = v19
                        var v65 uint64 = (v4 / (((^(v4))) | 1))
                        _ = fi61
                        _ = v63
                        _ = v64
                        _ = v65
                    }
                }
            }
            var v66 uint = v21
            _ = v66
        } else {
            if (true || (!false)) {
                v67 := GenS8{F0: 18446744073709551615}
                var v68 int = 518
                if (316 < 2) {
                    var v69 int32 = -1
                    _ = v69
                }
                _ = v67
                _ = v68
            } else {
                v21 = uint(v4)
            }
            v70 := GenS9{F0: v21, F1: 65535}
            _ = v70
        }
        switch 9223372036854775807 {
        case -9223372036854775808:
            var v71 int = (^(fi57))
            _ = v71
        }
        _ = fi57
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v3
    _ = v4
}
