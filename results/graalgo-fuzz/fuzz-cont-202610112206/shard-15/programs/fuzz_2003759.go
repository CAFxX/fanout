package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 int32 }
type GenS2 struct { F0 bool; F1 string }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int; F1 uint64 }
type GenS5 struct { F0 bool }

func main() {
    var v1 int16 = 42
    v2 := []int{4294967295, -891}
    var v3 int32 = -112
    var v4 bool = ((!true) || (1 >= 260))
    var v5 uint = 18446744073709551615
    var v6 uint64 = 4294967295
    if ("hello" != "foo bar") {
        v7 := GenS1{F0: v1, F1: (v3 << ((uint(v6)) & 63)), F2: 42}
        var v8 int64 = 290
        var v9 int64 = v8
        _ = v7
        _ = v8
        _ = v9
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if true {
            v6 = 18446744073709551615
        } else {
            switch 127 {
            case 169:
                if v4 {
                    v5 = (v5 ^ v5)
                    v6 = (v6 & (v6 << (((v5 ^ (v5 / (((^(v5))) | 1)))) & 63)))
                }
            case 612:
                if true {
                    var v10 int = -944
                    _ = v10
                } else {
                    fmt.Println(v1, v4, v4)
                    _ = len(v2)
                    for fi11 := 0; fi11 < 2; fi11++ {
                        v5 = (v5 % (((^(v5))) | 1))
                        _ = fi11
                    }
                }
            }
            v6 = (^(v6))
            if v4 {
                switch -1 {
                case -1:
                    if (false || true) {
                        switch 32768 {
                        case 1:
                            for fi12 := 0; fi12 < 2; fi12++ {
                                v13 := GenS2{F0: ((!(!(!v4))) || true), F1: ("a" + ((("" + ("foo bar" + (("foo bar" + "hello") + ("a" + ("foo bar" + "hello"))))) + (("hello" + (("foo bar" + "foo bar") + "hello")) + "\n\t")) + "\n\t"))}
                                fmt.Println(v1, v3)
                                _ = fi12
                                _ = v13
                            }
                        case -1:
                            var v14 uint16 = 157
                            _ = v14
                        default:
                            switch "foo bar" {
                            case "foo bar":
                                var v15 int32 = v3
                                _ = v15
                            case "\n\t":
                                fmt.Println(v5)
                            case "a":
                                for fi16 := 0; fi16 < 14; fi16++ {
                                    v17 := GenS2{F0: ((("hello" + "") != (((((("\n\t" + "hello") + "\n\t") + (("" + ("\n\t" + "")) + "foo bar")) + "a") + ("hello" + "foo bar")) + "foo bar")) && v4), F1: ("foo bar" + ("hello" + "\n\t"))}
                                    if (42 == 9223372036854775807) {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        for fi18 := 0; fi18 < 10; fi18++ {
                                            v3 -= 2147483647
                                            v3 = (v3 | v3)
                                            _ = fi18
                                        }
                                        _ = len(v2)
                                    } else {
                                        switch ("a" + ("hello" + "")) {
                                        case "\n\t":
                                            for fi19 := 0; fi19 < 16; fi19++ {
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                if false {
                                                    v1 += 42
                                                } else {
                                                    if (82 != 1) {
                                                        var v20 bool = false
                                                        _ = v20
                                                    }
                                                }
                                                v21 := GenS3{F0: 127}
                                                _ = fi19
                                                _ = v21
                                            }
                                        case "a":
                                            fmt.Println(v5)
                                        default:
                                            if len(v2) > 0 {
                                                v2[0] = fi16
                                            }
                                        }
                                    }
                                    var v22 int = (fi16 ^ (fi16 | (fi16 * fi16)))
                                    _ = fi16
                                    _ = v17
                                    _ = v22
                                }
                            default:
                                fmt.Println(v5)
                            }
                        }
                        for fi23 := 0; fi23 < 4; fi23++ {
                            var v24 int = -9223372036854775808
                            for fi25 := 0; fi25 < 9; fi25++ {
                                switch "a" {
                                case "foo bar":
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                case "a":
                                    switch ("a" + ("hello" + "\n\t")) {
                                    case "a":
                                        var v26 string = "hello"
                                        _ = v26
                                    case "foo bar":
                                        if (v4 || v4) {
                                            fmt.Println(fi23, v6, v3)
                                            v24 = fi23
                                            v6 += (v6 << ((v5) & 63))
                                        } else {
                                            v1 += v1
                                            var v27 int32 = (v3 + (v3 * v3))
                                            _ = v27
                                        }
                                    case "":
                                        v28 := GenS3{F0: -53}
                                        _ = v28
                                    default:
                                        if (!(uint(v6) <= 377)) {
                                            if ((-2.25 / 1.5) <= ((-2.25 - float64(1e10)) * 3.14159)) {
                                                if (-1 > 42) {
                                                    var v29 uint64 = 2
                                                    switch (fi25 + v24) {
                                                    case 341:
                                                        fmt.Println(v24)
                                                    case -353:
                                                        if len(v2) > 0 {
                                                            v2[0] = fi23
                                                        }
                                                        if (766 <= 9223372036854775807) {
                                                            v1 += (v1 ^ (v1 & -1))
                                                            v3 = 2147483647
                                                            var v30 bool = v4
                                                            _ = v30
                                                        } else {
                                                            var v31 uint64 = v6
                                                            v32 := GenS4{F0: v24, F1: 0}
                                                            _ = v31
                                                            _ = v32
                                                        }
                                                    }
                                                    if false {
                                                        v4 = v4
                                                        var v33 bool = (-128 > -95)
                                                        v1 += v1
                                                        _ = len(v2)
                                                        _ = v33
                                                    } else {
                                                        var v34 uint8 = 42
                                                        if (4294967295 <= 481) {
                                                            if (2 < 42) {
                                                                v4 = (v3 <= (v3 - v3))
                                                            } else {
                                                                fmt.Println(fi23, v6, fi23)
                                                            }
                                                            switch ("hello" + (("a" + "hello") + ("foo bar" + ("hello" + (("hello" + ((("" + "hello") + (("foo bar" + "foo bar") + ((("\n\t" + ("\n\t" + "a")) + "a") + "a"))) + ("hello" + ((("hello" + "\n\t") + (("foo bar" + "foo bar") + "hello")) + ("foo bar" + ""))))) + "hello"))))) {
                                                            case "foo bar":
                                                                v2 = append(v2, v24)
                                                            }
                                                            fmt.Println(v5, v5, v1)
                                                        }
                                                        v29 = 2
                                                        _ = v34
                                                    }
                                                    var v35 int32 = (v3 & v3)
                                                    v3 += v3
                                                    _ = v29
                                                    _ = v35
                                                } else {
                                                    if true {
                                                        if false {
                                                            if false {
                                                                var v36 uint32 = 897
                                                                _ = v36
                                                            } else {
                                                                v6 = v6
                                                                v5 += uint(v6)
                                                            }
                                                            var v37 int32 = v3
                                                            var v38 int = (v24 << ((v5) & 63))
                                                            _ = v37
                                                            _ = v38
                                                        } else {
                                                            v24 = fi23
                                                            v39 := GenS1{F0: v1, F1: v3, F2: v3}
                                                            if v4 {
                                                                var v40 uint8 = 1
                                                                _ = v40
                                                            }
                                                            _ = v39
                                                        }
                                                    }
                                                    v41 := GenS3{F0: 114}
                                                    fmt.Println(fi23)
                                                    _ = v41
                                                }
                                                if len(v2) > 0 {
                                                    v2[0] = 544
                                                }
                                            }
                                        }
                                    }
                                case "\n\t":
                                    var v42 int = fi25
                                    _ = v42
                                default:
                                    if v4 {
                                        fmt.Println(fi25, v6)
                                    }
                                }
                                if (2147483647 != 42) {
                                    var v43 string = ("hello" + "a")
                                    _ = len(v2)
                                    _ = v43
                                }
                                v44 := GenS5{F0: ((!false) && (42 >= v1))}
                                _ = fi25
                                _ = v44
                            }
                            for fi45 := 0; fi45 < 15; fi45++ {
                                fmt.Println(v6, v5)
                                if len(v2) > 0 {
                                    v2[0] = (^(fi23))
                                }
                                _ = fi45
                            }
                            _ = fi23
                            _ = v24
                        }
                        fmt.Println(v3)
                    }
                case 434:
                    v46 := GenS4{F0: int(v3), F1: v6}
                    _ = v46
                case 128:
                    v47 := GenS2{F0: true, F1: "a"}
                    _ = v47
                }
                var v48 int64 = 1
                if false {
                    var v49 uint64 = 979
                    _ = v49
                }
                _ = v48
            }
        }
    }
    switch ("a" + "\n\t") {
    case "hello":
        for fi50 := 0; fi50 < 11; fi50++ {
            var v51 uint16 = 32767
            fmt.Println(v5, fi50)
            _ = fi50
            _ = v51
        }
    case "\n\t":
        v4 = true
    }
    switch "a" {
    case "hello":
        v52 := GenS2{F0: v4, F1: "foo bar"}
        _ = v52
    case "":
        switch ("" + ("hello" + "")) {
        case "":
            var v53 string = ("\n\t" + "")
            _ = v53
        case "a":
            for fi54 := 0; fi54 < 6; fi54++ {
                fmt.Println(v4, v3)
                var v55 uint = (v5 << ((v5) & 63))
                _ = fi54
                _ = v55
            }
        case "foo bar":
            v56 := GenS5{F0: true}
            _ = v56
        }
    case "\n\t":
        v2 = append(v2, -16)
    default:
        v57 := GenS5{F0: v4}
        _ = v57
    }
    v4 = (v1 < (v1 >> ((18446744073709551615) & 63)))
    _ = len(v2)
    var v58 uint8 = 1
    var v59 int64 = 1
    fmt.Println(v59, v58)
    var v60 int = 42
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v59
    _ = v6
    _ = v60
}
