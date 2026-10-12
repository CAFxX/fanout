package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 string; F1 bool; F2 int8 }
type GenS4 struct { F0 bool; F1 int64; F2 string }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int64; F1 uint32; F2 uint8 }
type GenS7 struct { F0 int16; F1 bool; F2 uint8 }
type GenS8 struct { F0 uint8; F1 int16; F2 uint64 }
func genFn0() uint {
    var v6 string = ""
    var v7 int = 32767
    switch (^(v7)) {
    case -601:
        var v8 string = ((v6 + v6) + (v6 + (v6 + v6)))
        _ = v8
    case 9223372036854775807:
        if ("" == ("" + "\n\t")) {
            v9 := []int{(v7 * (-(v7))), 128, v7}
            var v10 bool = (127 >= 42)
            _ = v10
            _ = v9
        } else {
            v6 = v6
            v11 := GenS1{F0: false}
            _ = v11
        }
    case 298:
        if (1 < 1) {
            var v12 int = v7
            if true {
                if false {
                    fmt.Println(v6, v12, v7)
                    switch v6 {
                    case "hello":
                        if (0 <= 0) {
                            if false {
                                switch (v7 ^ (v12 % ((v7) | 1))) {
                                case 127:
                                    fmt.Println(v7, v7)
                                }
                            } else {
                                var v13 float32 = 3.14159
                                var v14 int16 = 32767
                                _ = v13
                                _ = v14
                            }
                        } else {
                            fmt.Println(v6, v7)
                            v15 := []int{(v12 % ((v12) | 1)), 9223372036854775807, 42}
                            switch v12 {
                            case 128:
                                v15 = append(v15, v12)
                            case 42:
                                if len(v15) > 0 {
                                    v15[0] = (v7 - v7)
                                }
                            default:
                                v16 := GenS2{F0: ((false && false) || (!false)), F1: 128}
                                _ = v16
                            }
                            _ = v15
                        }
                    case "a":
                        if true {
                            for fi17 := 0; fi17 < 6; fi17++ {
                                v18 := GenS3{F0: v6, F1: (42 > 100), F2: -1}
                                _ = fi17
                                _ = v18
                            }
                        }
                    default:
                        var v19 float32 = (((1e10 + 1e10) - 3.14159) / 1e10)
                        _ = v19
                    }
                }
            } else {
                var v20 uint8 = 42
                v21 := []int{9223372036854775807, v7}
                var v22 int8 = 73
                _ = v20
                _ = v21
                _ = v22
            }
            v6 = (v6 + (v6 + "\n\t"))
            _ = v12
        } else {
            if false {
                switch (-(v7)) {
                case 128:
                    v23 := GenS4{F0: (true || (!false)), F1: 9223372036854775807, F2: v6}
                    _ = v23
                default:
                    v7 = v7
                }
            } else {
                var v24 int64 = 1
                _ = v24
            }
            fmt.Println(v6, v6)
            v6 = "hello"
        }
    }
    for fi25 := 0; fi25 < 2; fi25++ {
        v26 := GenS5{F0: -32768}
        if ((-24 == 1) || false) {
            var v27 uint = 255
            _ = v27
        } else {
            v6 = v6
            for fi28 := 0; fi28 < 16; fi28++ {
                v29 := GenS6{F0: 1, F1: 2147483648, F2: 0}
                switch (fi25 * (fi28 >> ((792) & 63))) {
                case 127:
                    v30 := GenS7{F0: 0, F1: ((-128 != 127) && false), F2: 0}
                    _ = v30
                }
                _ = fi28
                _ = v29
            }
        }
        for fi31 := 0; fi31 < 18; fi31++ {
            switch "hello" {
            case "\n\t":
                v6 = v6
            case "a":
                var v32 uint32 = 45
                _ = v32
            case "":
                v33 := GenS3{F0: v6, F1: true, F2: -128}
                var v34 uint16 = 32767
                _ = v33
                _ = v34
            }
            var v35 uint32 = 1
            var v36 uint16 = 166
            _ = fi31
            _ = v35
            _ = v36
        }
        _ = fi25
        _ = v26
    }
    var v37 int64 = 42
    _ = v37
    _ = v6
    _ = v7
    return 2
}
func genFn1() int {
    if (739 != 42) {
        var v38 int64 = 313
        switch "" {
        case "\n\t":
            v39 := []int{42, 0}
            _ = v39
        }
        v40 := []int{-1}
        _ = v38
        _ = v40
    } else {
        var v41 uint8 = 153
        _ = v41
    }
    var v42 float32 = ((0.1 - ((((float32(0.0) / 1.0) / 3.14159) * (((0.0 - 0.1) - (0.1 + 0.1)) / (float32((1.5 + 1.5)) / ((0.1 / (3.14159 * 1.5)) / -2.25)))) / 3.14159)) / (3.14159 - 0.0))
    if true {
        v43 := GenS1{F0: ("foo bar" < "a")}
        v42 = (((float32(0.0) / (v42 / 1e10)) - (v42 / v42)) - v42)
        if ((2147483647 == 2147483647) && (false || false)) {
            var v44 bool = (v42 == v42)
            _ = v44
        }
        _ = v43
    }
    _ = v42
    return -889
}

func main() {
    v1 := []int{}
    v2 := []string{"\n\t", ("hello" + "hello")}
    v3 := make(map[int]uint32)
    v4 := make(map[int]int)
    var v5 string = ("hello" + ("hello" + "a"))
    v5 = "foo bar"
    fmt.Println(v5)
    if false {
        v5 = (v5 + v5)
        for fi45 := 0; fi45 < 4; fi45++ {
            var v46 int64 = -1
            switch v5 {
            case "hello":
                v3[fi45] = 1
            case "\n\t":
                switch (-(fi45)) {
                case -697:
                    var v47 int8 = 79
                    _ = v47
                default:
                    v5 = "a"
                }
            case "":
                if true {
                    if false {
                        v46 += 0
                        var v48 bool = ((255 > 33) && (176 < 1))
                        _ = v48
                    }
                } else {
                    var v49 string = v5
                    for fi50 := 0; fi50 < 13; fi50++ {
                        var v51 uint = 18446744073709551615
                        _ = fi50
                        _ = v51
                    }
                    var v52 uint8 = 127
                    _ = v49
                    _ = v52
                }
            }
            var v53 int16 = 1
            _ = fi45
            _ = v46
            _ = v53
        }
    } else {
        fmt.Println(v5)
        _ = len(v2)
    }
    v54 := GenS4{F0: true, F1: 9223372036854775807, F2: "a"}
    v5 = "\n\t"
    var v55 int32 = 387
    fmt.Println(v55, v55)
    if len(v2) > 0 {
        v2[0] = (v5 + v5)
    }
    var v56 bool = false
    var v57 int8 = 1
    if (!true) {
        v55 += v55
    }
    v4[914] = 0
    delete(v3, 127)
    delete(v4, -639)
    v58 := GenS1{F0: (!((3.14159 != (((0.0 - (1e10 + 0.0)) * 3.14159) + 1e10)) || v56))}
    for fi59 := 0; fi59 < 18; fi59++ {
        v60 := GenS8{F0: 128, F1: -816, F2: 1}
        _ = fi59
        _ = v60
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v55
    _ = v56
    _ = v57
    _ = v58
}
