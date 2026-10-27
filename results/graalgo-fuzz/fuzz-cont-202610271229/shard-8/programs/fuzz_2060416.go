package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16; F1 uint64; F2 int64 }
type GenS4 struct { F0 bool; F1 int32 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 string; F1 int16 }
type GenS7 struct { F0 int; F1 uint }
type GenS8 struct { F0 uint; F1 int16; F2 uint }
type GenS9 struct { F0 bool; F1 int16 }
func genFn0() int16 {
    switch 0 {
    case 655:
        switch -2147483648 {
        case 42:
            if (!(-2.25 == 3.14159)) {
                var v4 int = 127
                for fi5 := 0; fi5 < 17; fi5++ {
                    v6 := GenS1{F0: 1, F1: 128}
                    _ = fi5
                    _ = v6
                }
                v7 := GenS1{F0: 127, F1: 1}
                _ = v4
                _ = v7
            } else {
                if false {
                    for fi8 := 0; fi8 < 1; fi8++ {
                        var v9 string = "foo bar"
                        v9 = v9
                        for fi10 := 0; fi10 < 11; fi10++ {
                            var v11 int16 = 0
                            switch (("hello" + v9) + "foo bar") {
                            case "hello":
                                var v12 bool = true
                                _ = v12
                            case "foo bar":
                                fmt.Println(v11, fi8, v11)
                            case "\n\t":
                                v13 := GenS2{F0: 1}
                                _ = v13
                            }
                            if (!false) {
                                v14 := []int{}
                                v11 -= (v11 << ((18446744073709551615) & 63))
                                fmt.Println(fi10)
                                if true {
                                    v9 = v9
                                }
                                _ = v14
                            } else {
                                if (1 <= 127) {
                                    var v15 string = ("a" + (v9 + v9))
                                    v9 = v9
                                    switch fi8 {
                                    case 1:
                                        var v16 float32 = 1e10
                                        _ = v16
                                    case 32767:
                                        switch (v9 + "hello") {
                                        case "\n\t":
                                            var v17 uint8 = 255
                                            _ = v17
                                        }
                                    }
                                    _ = v15
                                } else {
                                    fmt.Println(v9, v9, v9)
                                    if ((9223372036854775807 <= 0) && false) {
                                        v18 := GenS2{F0: (v11 << ((2) & 63))}
                                        if true {
                                            if true {
                                                v9 = v9
                                                fmt.Println(fi10, v11)
                                            }
                                            var v19 string = "a"
                                            switch (fi10 - (fi10 | fi8)) {
                                            case 2147483647:
                                                fmt.Println(fi10, fi10)
                                            case -474:
                                                v19 = ""
                                            }
                                            _ = v19
                                        }
                                        _ = v18
                                    }
                                }
                            }
                            _ = fi10
                            _ = v11
                        }
                        _ = fi8
                        _ = v9
                    }
                    var v20 int8 = -1
                    if false {
                        v21 := GenS3{F0: 42, F1: 581, F2: -1}
                        _ = v21
                    } else {
                        v22 := GenS2{F0: -1}
                        _ = v22
                    }
                    _ = v20
                } else {
                    var v23 int64 = 1
                    _ = v23
                }
                v24 := GenS2{F0: 32767}
                var v25 bool = (!((1e10 < 1e10) || ("foo bar" == "hello")))
                _ = v24
                _ = v25
            }
        case -578:
            var v26 uint8 = 1
            _ = v26
        default:
            v27 := []int{32767}
            _ = v27
        }
    }
    var v28 uint32 = 2
    var v29 int64 = 204
    if true {
        fmt.Println(v29, v29)
    } else {
        v28 -= v28
        var v30 int64 = (v29 & v29)
        _ = v30
    }
    v31 := []int{-2147483648, -2147483648, 320}
    _ = v28
    _ = v29
    _ = v31
    return -415
}

func main() {
    var v1 int = 42
    var v2 uint32 = 316
    v3 := []string{}
    v32 := GenS4{F0: true, F1: -2147483648}
    for fi33 := 0; fi33 < 2; fi33++ {
        fmt.Println(v1)
        for fi34 := 0; fi34 < 14; fi34++ {
            if (0 > 261) {
                fmt.Println(v1, v1, fi34)
                v1 -= 0
                if true {
                    v35 := GenS2{F0: 1}
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v36 := GenS4{F0: true, F1: 2147483647}
                    _ = v35
                    _ = v36
                } else {
                    v37 := GenS5{F0: (0 < 255)}
                    _ = v37
                }
            } else {
                var v38 int16 = -1
                var v39 int8 = 42
                switch ("" + ("" + "foo bar")) {
                case "":
                    v38 = int16(v39)
                default:
                    var v40 float64 = 3.14159
                    _ = v40
                }
                _ = v38
                _ = v39
            }
            _ = fi34
        }
        v1 += (fi33 + v1)
        _ = fi33
    }
    v41 := GenS1{F0: -128, F1: 255}
    if len(v3) > 0 {
        v3[0] = "foo bar"
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v2 = 4294967295
    if (100 >= 127) {
        v2 -= v2
        var v42 uint8 = 128
        _ = v42
    } else {
        v43 := GenS6{F0: ("hello" + (("hello" + "foo bar") + "hello")), F1: 0}
        if (!true) {
            v44 := GenS7{F0: (-(v1)), F1: 1}
            v1 -= v1
            _ = v44
        } else {
            var v45 int64 = 253
            _ = v45
        }
        _ = v43
    }
    if (false || false) {
        fmt.Println(v1, v2)
        v1 = v1
    }
    v46 := GenS8{F0: 42, F1: genFn0(), F2: 4294967295}
    v47 := GenS9{F0: ((1 > 7) || true), F1: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v41
    _ = v46
    _ = v47
}
