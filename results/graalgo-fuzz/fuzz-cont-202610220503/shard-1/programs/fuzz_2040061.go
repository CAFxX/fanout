package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 uint }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64; F1 uint16 }
type GenS4 struct { F0 uint64; F1 int64; F2 int64 }
type GenS5 struct { F0 int32; F1 int16 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int; F1 uint64 }

func main() {
    var v1 string = "\n\t"
    v2 := []string{}
    var v3 string = ("foo bar" + (v1 + v1))
    var v4 bool = (32767 <= 118)
    v3 = "hello"
    fmt.Println(v3, v1, v4)
    v5 := GenS1{F0: 1, F1: 2147483648, F2: 226}
    if true {
        fmt.Println(v1, v3)
    } else {
        v4 = (((9223372036854775807 > 325) && (!v4)) && true)
    }
    if v4 {
        if true {
            v3 = ((v1 + v3) + v1)
            if (v4 || (255 >= 0)) {
                if v4 {
                    fmt.Println(v1, v3, v3)
                }
                if (!(2 < 2)) {
                    _ = len(v2)
                    v2 = append(v2, "a")
                }
                var v6 float64 = (-2.25 * 0.0)
                _ = v6
            }
        } else {
            switch 128 {
            case 0:
                for fi7 := 0; fi7 < 12; fi7++ {
                    var v8 bool = (!v4)
                    v9 := GenS2{F0: v4}
                    _ = fi7
                    _ = v8
                    _ = v9
                }
            case 42:
                switch v1 {
                case "a":
                    v10 := GenS1{F0: 42, F1: 0, F2: 40}
                    _ = v10
                case "hello":
                    var v11 uint32 = 65535
                    _ = v11
                default:
                    if v4 {
                        for fi12 := 0; fi12 < 18; fi12++ {
                            for fi13 := 0; fi13 < 6; fi13++ {
                                v3 = ""
                                fmt.Println(v3, v4)
                                v4 = v4
                                if ((-839 >= 327) && (true && (1 <= -1))) {
                                    v3 = ("hello" + v3)
                                } else {
                                    if (v4 && ((758 > 1) || true)) {
                                        var v14 int = (fi12 & -9223372036854775808)
                                        _ = v14
                                    }
                                    v15 := GenS2{F0: (1 == 460)}
                                    var v16 int32 = 0
                                    _ = v15
                                    _ = v16
                                }
                                if (v4 && (-128 <= 100)) {
                                    if v4 {
                                        var v17 uint = 2
                                        _ = v17
                                    } else {
                                        if v4 {
                                            var v18 uint64 = 2
                                            _ = v18
                                        } else {
                                            if len(v2) > 0 {
                                                v2[0] = "a"
                                            }
                                        }
                                        switch v1 {
                                        case "foo bar":
                                            v19 := GenS2{F0: false}
                                            _ = v19
                                        default:
                                            if false {
                                                if (42 != 18446744073709551615) {
                                                    if v4 {
                                                        v20 := GenS1{F0: -772, F1: 2147483647, F2: 1}
                                                        if v4 {
                                                            var v21 int32 = 62
                                                            var v22 uint = 18446744073709551615
                                                            _ = v21
                                                            _ = v22
                                                        } else {
                                                            v1 = "hello"
                                                        }
                                                        _ = v20
                                                    }
                                                } else {
                                                    _ = len(v2)
                                                    if (0.1 >= 1.5) {
                                                        switch 128 {
                                                        case 256:
                                                            switch fi12 {
                                                            case -9223372036854775808:
                                                                v2 = append(v2, (v3 + ""))
                                                            }
                                                        }
                                                    }
                                                    var v23 uint = 255
                                                    fmt.Println(v4, v4)
                                                    v4 = (v4 || (255 == 42))
                                                    _ = v23
                                                }
                                            }
                                        }
                                    }
                                    v1 = v3
                                    v4 = (1 < 1)
                                } else {
                                    var v24 int8 = 77
                                    var v25 int32 = -2147483648
                                    v2 = append(v2, v3)
                                    _ = v24
                                    _ = v25
                                }
                                _ = fi13
                            }
                            v1 = ""
                            _ = fi12
                        }
                        v2 = append(v2, ((("foo bar" + v1) + "\n\t") + "\n\t"))
                    }
                }
            case 127:
                switch "foo bar" {
                case "\n\t":
                    var v26 uint64 = 326
                    _ = v26
                case "foo bar":
                    v2 = append(v2, v1)
                case "a":
                    v1 = (v1 + ("a" + v1))
                }
            }
            if (0 == 1) {
                var v27 bool = true
                for fi28 := 0; fi28 < 18; fi28++ {
                    var v29 int64 = 406
                    v2 = append(v2, v1)
                    _ = len(v2)
                    _ = fi28
                    _ = v29
                }
                if (!false) {
                    var v30 int32 = 0
                    v31 := GenS2{F0: (42 == 2147483647)}
                    _ = v30
                    _ = v31
                } else {
                    fmt.Println(v27)
                }
                _ = v27
            } else {
                if (1 > 10) {
                    _ = len(v2)
                    if (255 == 259) {
                        for fi32 := 0; fi32 < 16; fi32++ {
                            for fi33 := 0; fi33 < 17; fi33++ {
                                if v4 {
                                    v3 = v3
                                    fmt.Println(fi32, fi33)
                                    if (0.0 > 1.5) {
                                        fmt.Println(fi33, fi33)
                                        var v34 int8 = 42
                                        _ = v34
                                    }
                                    var v35 uint64 = 1
                                    _ = v35
                                }
                                if false {
                                    if ((!true) && v4) {
                                        if len(v2) > 0 {
                                            v2[0] = "foo bar"
                                        }
                                        var v36 uint64 = 4294967295
                                        v37 := GenS3{F0: 969, F1: 65535}
                                        if ((v4 && false) && (!v4)) {
                                            var v38 bool = v4
                                            fmt.Println(v4)
                                            v39 := GenS4{F0: v36, F1: -494, F2: 1}
                                            _ = v38
                                            _ = v39
                                        } else {
                                            switch v3 {
                                            case "":
                                                var v40 int32 = -1
                                                _ = v40
                                            default:
                                                var v41 int8 = 66
                                                _ = v41
                                            }
                                            fmt.Println(v3)
                                            switch 787 {
                                            case -760:
                                                v42 := GenS3{F0: 634, F1: 2}
                                                _ = v42
                                            case 32767:
                                                if (0 < 226) {
                                                    if v4 {
                                                        v43 := GenS1{F0: -1, F1: 1, F2: 2}
                                                        _ = v43
                                                    }
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                }
                                            case -203:
                                                fmt.Println(v4, v4)
                                            }
                                        }
                                        _ = v36
                                        _ = v37
                                    }
                                }
                                _ = fi33
                            }
                            _ = fi32
                        }
                        v44 := GenS5{F0: 0, F1: 32767}
                        _ = v44
                    }
                }
                v3 = (v1 + "a")
            }
            var v45 float32 = -2.25
            _ = v45
        }
    } else {
        var v46 bool = (-1 > 100)
        v47 := GenS5{F0: 2147483647, F1: -1}
        _ = v46
        _ = v47
    }
    v3 = "foo bar"
    for fi48 := 0; fi48 < 13; fi48++ {
        fmt.Println(v1, v1, fi48)
        _ = fi48
    }
    var v49 bool = ((true || (-820 <= 0)) && (v4 && false))
    v4 = v4
    switch 32768 {
    case -1:
        for fi50 := 0; fi50 < 12; fi50++ {
            if true {
                _ = len(v2)
                var v51 uint = 851
                for fi52 := 0; fi52 < 19; fi52++ {
                    v51 -= (v51 * v51)
                    fmt.Println(fi52, v51, v49)
                    _ = fi52
                }
                _ = v51
            } else {
                if v4 {
                    if (127 == 1) {
                        var v53 int16 = -614
                        _ = v53
                    }
                }
                fmt.Println(fi50, v4)
            }
            _ = fi50
        }
    case 127:
        switch "a" {
        case "a":
            v54 := GenS1{F0: 1, F1: 576, F2: 65535}
            _ = v54
        }
    case -715:
        v55 := GenS6{F0: 127}
        _ = v55
    default:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    var v56 bool = (0.1 <= -2.25)
    if v4 {
        v57 := GenS5{F0: -1, F1: 1}
        v56 = v4
        _ = v57
    }
    v58 := GenS3{F0: 0, F1: 39}
    switch (("\n\t" + v1) + ("" + "hello")) {
    case "":
        v56 = (!v49)
    case "hello":
        fmt.Println(v49, v1)
    case "foo bar":
        for fi59 := 0; fi59 < 20; fi59++ {
            switch fi59 {
            case 0:
                for fi60 := 0; fi60 < 4; fi60++ {
                    if v49 {
                        v1 = ("\n\t" + ((("" + ((("" + v3) + v1) + ((v1 + v1) + v3))) + "") + v3))
                        fmt.Println(v3)
                    } else {
                        var v61 uint8 = 68
                        switch (fi59 - fi59) {
                        case 4294967295:
                            fmt.Println(v1)
                        default:
                            v49 = (0.1 >= 3.14159)
                        }
                        _ = v61
                    }
                    v62 := GenS7{F0: (fi60 ^ 32767), F1: 18446744073709551615}
                    _ = fi60
                    _ = v62
                }
            case 38:
                if false {
                    v3 = v3
                }
            default:
                var v63 bool = (0 == 42)
                _ = v63
            }
            _ = fi59
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v56
    _ = v58
}
