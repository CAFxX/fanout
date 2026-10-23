package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8; F1 uint; F2 uint16 }
type GenS3 struct { F0 uint8; F1 int16 }
type GenS4 struct { F0 int8; F1 int64 }
type GenS5 struct { F0 bool; F1 int32; F2 bool }
type GenS6 struct { F0 uint32 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []string{("" + ("" + ("hello" + (("a" + "\n\t") + ((("hello" + (("" + "foo bar") + "foo bar")) + ("" + "hello")) + "foo bar"))))), (("\n\t" + "a") + ""), "a", ""}
    v3 := make(map[uint32]bool)
    v4 := make(map[string]bool)
    v5 := []uint32{2}
    _ = len(v5)
    for fi6 := 0; fi6 < 19; fi6++ {
        _ = len(v1)
        v7 := GenS1{F0: "hello"}
        _ = fi6
        _ = v7
    }
    if (1 != 670) {
        if true {
            var v8 uint16 = 32767
            _ = v8
        } else {
            var v9 float64 = 0.1
            _ = v9
        }
        if (true && (378 > 2147483648)) {
            _ = len(v2)
        } else {
            var v10 int16 = -1
            var v11 uint32 = 370
            v11 += v11
            _ = v10
            _ = v11
        }
    }
    var v12 uint32 = 2147483648
    var v13 uint16 = 255
    if (-795 >= -12) {
        switch "a" {
        case "a":
            for fi14 := 0; fi14 < 8; fi14++ {
                fmt.Println(v13, v12)
                _ = fi14
            }
        case "\n\t":
            if true {
                v13 = 1
            } else {
                fmt.Println(v13, v13)
                if (!(1 == -1)) {
                    v15 := GenS2{F0: 18, F1: 42, F2: 42}
                    _ = v15
                }
            }
        }
        switch 228 {
        case -2147483648:
            var v16 string = ""
            _ = v16
        case 1:
            for fi17 := 0; fi17 < 15; fi17++ {
                fmt.Println(v12, v12)
                switch (fi17 - 4294967295) {
                case 1:
                    switch "" {
                    case "hello":
                        _ = len(v4)
                    }
                case 32767:
                    switch "" {
                    case "":
                        v12 += (v12 / ((v12) | 1))
                    case "\n\t":
                        var v18 uint16 = 42
                        _ = v18
                    }
                }
                switch fi17 {
                case 32768:
                    fmt.Println(v12, v12)
                case 2147483647:
                    v12 += (v12 + (v12 | (^(v12))))
                }
                _ = fi17
            }
        case 396:
            if true {
                delete(v3, (v12 % (((^(v12))) | 1)))
                v19 := GenS1{F0: ""}
                var v20 uint16 = (v13 % (((v13 << ((673) & 63))) | 1))
                _ = v19
                _ = v20
            }
        default:
            var v21 uint16 = 65535
            _ = v21
        }
    } else {
        for fi22 := 0; fi22 < 14; fi22++ {
            var v23 float64 = 0.0
            v13 = 65535
            var v24 string = "hello"
            _ = fi22
            _ = v23
            _ = v24
        }
        var v25 bool = ((3.14159 != 0.0) && false)
        _ = v25
    }
    if ((!true) && true) {
        v13 = (v13 << ((0) & 63))
        var v26 float64 = 3.14159
        _ = v26
    } else {
        switch ("" + "foo bar") {
        case "\n\t":
            var v27 float32 = (float32(3.14159) + (1.5 / 1.0))
            _ = v27
        case "foo bar":
            for fi28 := 0; fi28 < 18; fi28++ {
                if (9223372036854775807 == 9223372036854775807) {
                    var v29 int16 = -1
                    switch "hello" {
                    case "":
                        v12 = v12
                    case "\n\t":
                        v12 = (v12 << ((722) & 63))
                    }
                    _ = v29
                } else {
                    if true {
                        var v30 int32 = -353
                        _ = v30
                    }
                }
                _ = fi28
            }
        default:
            if false {
                v12 += v12
                if true {
                    fmt.Println(v12)
                    v13 -= (v13 ^ v13)
                    var v31 int16 = 42
                    _ = v31
                } else {
                    v13 = (v13 / ((v13) | 1))
                    v13 = v13
                }
                for fi32 := 0; fi32 < 11; fi32++ {
                    fmt.Println(v12, fi32)
                    _ = fi32
                }
            } else {
                switch "\n\t" {
                case "hello":
                    if len(v5) > 0 {
                        v5[0] = (v12 - 0)
                    }
                case "":
                    var v33 int = 0
                    _ = v33
                }
            }
        }
    }
    fmt.Println(v12)
    v34 := GenS3{F0: 128, F1: -760}
    if true {
        switch -2147483648 {
        case 4294967295:
            var v35 uint8 = 0
            _ = v35
        case -2147483648:
            var v36 int8 = 127
            _ = v36
        case 32767:
            fmt.Println(v12, v12)
        }
        v37 := GenS4{F0: -1, F1: -157}
        fmt.Println(v12, v13)
        _ = v37
    } else {
        for fi38 := 0; fi38 < 11; fi38++ {
            delete(v3, (v12 % (((v12 >> ((0) & 63))) | 1)))
            _ = fi38
        }
        v5 = append(v5, v12)
        if false {
            for fi39 := 0; fi39 < 20; fi39++ {
                for fi40 := 0; fi40 < 6; fi40++ {
                    var v41 uint32 = (v12 << ((0) & 63))
                    _ = fi40
                    _ = v41
                }
                switch "" {
                case "":
                    for fi42 := 0; fi42 < 15; fi42++ {
                        v13 -= (v13 % ((v13) | 1))
                        v43 := v4[("\n\t" + ("\n\t" + "foo bar"))]
                        if (v12 != 2147483648) {
                            var v44 float32 = (3.14159 / 1.5)
                            v45 := GenS2{F0: 127, F1: 785, F2: v13}
                            _ = v44
                            _ = v45
                        }
                        _ = fi42
                        _ = v43
                    }
                case "foo bar":
                    if (-1 >= -2147483648) {
                        v13 += 111
                        v46 := GenS5{F0: true, F1: 1, F2: (!true)}
                        fmt.Println(fi39)
                        _ = v46
                    } else {
                        var v47 float32 = 1.5
                        for fi48 := 0; fi48 < 19; fi48++ {
                            fmt.Println(fi39, fi39, v12)
                            v49 := GenS5{F0: (false && ((127 < uint8(v12)) || (true || (true || (!(-2147483648 != 0)))))), F1: -567, F2: true}
                            switch (("a" + "a") + ("foo bar" + "a")) {
                            case "foo bar":
                                if false {
                                    var v50 int64 = 42
                                    _ = v50
                                }
                            case "":
                                var v51 uint32 = 2147483647
                                _ = v51
                            case "a":
                                v47 = v47
                            }
                            if (42 >= 2) {
                                v52 := GenS5{F0: false, F1: 2147483647, F2: true}
                                v47 = v47
                                _ = v52
                            } else {
                                delete(v1, 0)
                            }
                            if true {
                                var v53 int16 = -32768
                                if (1 <= (^(v13))) {
                                    v47 = v47
                                    if false {
                                        v12 += (v12 / ((v12) | 1))
                                    }
                                    v54 := GenS5{F0: (!false), F1: -2147483648, F2: (!true)}
                                    _ = v54
                                }
                                _ = v53
                            }
                            _ = fi48
                            _ = v49
                        }
                        _ = len(v3)
                        _ = v47
                    }
                case "a":
                    v13 += v13
                default:
                    v55 := GenS6{F0: (v12 >> ((856) & 63))}
                    _ = v55
                }
                _ = fi39
            }
            v13 = 0
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
}
