package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string; F1 uint16; F2 int32 }
type GenS3 struct { F0 uint32; F1 int8; F2 int8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint; F1 string; F2 uint8 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint; F1 uint32; F2 int8 }
func genFn0(p0 uint, p1 uint) uint16 {
    fmt.Println(p0, p1)
    switch "hello" {
    case "a":
        fmt.Println(p0, p0)
    case "\n\t":
        var v5 int32 = -1
        _ = v5
    case "hello":
        if (0.0 < (3.14159 * 1e10)) {
            p1 = (^(p0))
            v6 := GenS1{F0: 32767}
            v7 := GenS1{F0: 32767}
            _ = v6
            _ = v7
        } else {
            if (false || false) {
                var v8 uint = 1
                _ = v8
            }
            var v9 int16 = 1
            _ = v9
        }
    default:
        var v10 int16 = 0
        _ = v10
    }
    p1 = 18446744073709551615
    _ = p0
    _ = p1
    return 0
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]string)
    var v3 bool = true
    v4 := []string{"\n\t", "\n\t"}
    if true {
        var v11 bool = (0.1 != 3.14159)
        v12 := GenS1{F0: 217}
        _ = v11
        _ = v12
    }
    fmt.Println(v3)
    v3 = (42 > 2147483647)
    var v13 int = -9223372036854775808
    fmt.Println(v13, v13)
    fmt.Println(v13, v3)
    switch "\n\t" {
    case "":
        v3 = (1e10 >= 1.5)
    case "\n\t":
        switch ("foo bar" + ("a" + "a")) {
        case "":
            switch 1 {
            case 42:
                v14 := GenS1{F0: 0}
                _ = v14
            }
        case "hello":
            if v3 {
                var v15 uint8 = 127
                for fi16 := 0; fi16 < 3; fi16++ {
                    v13 = (fi16 * 2147483647)
                    var v17 uint64 = 75
                    var v18 float32 = 1e10
                    _ = fi16
                    _ = v17
                    _ = v18
                }
                _ = v15
            }
        case "a":
            v3 = (19 <= 42)
        }
    case "hello":
        v13 += (^(v13))
    }
    v3 = v3
    v13 = v13
    switch "\n\t" {
    case "\n\t":
        var v19 int32 = 790
        _ = v19
    case "a":
        v20 := GenS1{F0: 1000}
        _ = v20
    case "hello":
        if v3 {
            v13 = v13
            v13 -= v13
        } else {
            var v21 int16 = 32767
            if (!((v3 && (v3 || (-9223372036854775808 != 267))) || (!true))) {
                for fi22 := 0; fi22 < 10; fi22++ {
                    if false {
                        for fi23 := 0; fi23 < 13; fi23++ {
                            v21 = (v21 - (v21 | v21))
                            var v24 int64 = 9223372036854775807
                            _ = fi23
                            _ = v24
                        }
                    }
                    if (1 > 49) {
                        _ = len(v4)
                        v25 := GenS2{F0: ("hello" + ("foo bar" + "a")), F1: 42, F2: 905}
                        _ = v25
                    }
                    _ = fi22
                }
            } else {
                v3 = (0 == 42)
                if true {
                    for fi26 := 0; fi26 < 18; fi26++ {
                        var v27 uint16 = 42
                        switch ("hello" + "") {
                        case "hello":
                            fmt.Println(fi26)
                        default:
                            delete(v1, "foo bar")
                        }
                        _ = fi26
                        _ = v27
                    }
                    var v28 bool = v3
                    if (!true) {
                        if false {
                            var v29 uint32 = 0
                            v1["hello"] = v13
                            _ = v29
                        } else {
                            var v30 int64 = 419
                            switch (^(v13)) {
                            case 732:
                                var v31 int32 = -1
                                _ = v31
                            case 0:
                                if v3 {
                                    fmt.Println(v3, v3)
                                }
                            case 456:
                                v13 -= (v13 >> ((255) & 63))
                            default:
                                v13 += -1
                            }
                            var v32 int16 = (v21 % ((-663) | 1))
                            _ = v30
                            _ = v32
                        }
                        v2[536] = (("a" + "") + "foo bar")
                    }
                    _ = v28
                }
            }
            var v33 uint16 = 977
            _ = v21
            _ = v33
        }
    default:
        fmt.Println(v13, v3)
    }
    if v3 {
        v34 := GenS3{F0: 4294967295, F1: 0, F2: 127}
        var v35 uint64 = 4294967295
        _ = v34
        _ = v35
    }
    if len(v4) > 0 {
        v4[0] = ""
    }
    for fi36 := 0; fi36 < 8; fi36++ {
        for fi37 := 0; fi37 < 15; fi37++ {
            v3 = false
            fmt.Println(v3)
            _ = fi37
        }
        v3 = ((819 < 1) && (v3 || (!(!("\n\t" == "a")))))
        _ = fi36
    }
    for fi38 := 0; fi38 < 18; fi38++ {
        if len(v4) > 0 {
            v4[0] = ("hello" + ("hello" + (("" + "hello") + "")))
        }
        _ = fi38
    }
    v13 += (v13 % ((v13) | 1))
    if false {
        if (true || (2 > 255)) {
            switch (v13 << ((0) & 63)) {
            case -2147483648:
                if ((57 == 1) || true) {
                    fmt.Println(v3, v13)
                    switch "foo bar" {
                    case "foo bar":
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    default:
                        switch ((((("foo bar" + ("" + ("" + ("hello" + ("a" + ("" + ("hello" + ("foo bar" + "")))))))) + "hello") + "\n\t") + "a") + "a") {
                        case "hello":
                            for fi39 := 0; fi39 < 11; fi39++ {
                                var v40 int = 127
                                v41 := GenS4{F0: (!v3)}
                                _ = fi39
                                _ = v40
                                _ = v41
                            }
                        case "foo bar":
                            if (0.1 < 1e10) {
                                if v3 {
                                    var v42 uint8 = 42
                                    v3 = (1000 != -357)
                                    _ = v42
                                }
                                v43 := GenS5{F0: 2, F1: "foo bar", F2: 0}
                                if (127 > 127) {
                                    fmt.Println(v3, v13)
                                } else {
                                    for fi44 := 0; fi44 < 4; fi44++ {
                                        v45 := GenS5{F0: 357, F1: ("a" + "hello"), F2: 48}
                                        _ = fi44
                                        _ = v45
                                    }
                                    if (!v3) {
                                        v3 = ((((1e10 * ((float64((0.0 / 3.14159)) + -2.25) + 0.1)) - 0.1) - (1.5 / 1.5)) == 0.1)
                                        v3 = (v3 || true)
                                        var v46 bool = v3
                                        _ = v46
                                    }
                                }
                                _ = v43
                            } else {
                                var v47 string = ""
                                _ = len(v4)
                                _ = v47
                            }
                        case "":
                            v3 = (255 == 2)
                        default:
                            if v3 {
                                var v48 int64 = -445
                                fmt.Println(v48)
                                _ = v48
                            }
                        }
                    }
                    var v49 uint32 = 2
                    _ = v49
                } else {
                    v3 = v3
                }
            case 32768:
                v3 = (!(-731 <= 0))
            }
        }
    } else {
        switch "" {
        case "foo bar":
            var v50 string = "a"
            _ = v50
        case "hello":
            switch v13 {
            case 696:
                v51 := GenS6{F0: 255}
                _ = v51
            case 0:
                fmt.Println(v3, v3)
            }
        }
        v52 := GenS7{F0: 160, F1: 2, F2: 1}
        for fi53 := 0; fi53 < 2; fi53++ {
            v13 += v13
            var v54 int16 = 1
            _ = fi53
            _ = v54
        }
        _ = v52
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
