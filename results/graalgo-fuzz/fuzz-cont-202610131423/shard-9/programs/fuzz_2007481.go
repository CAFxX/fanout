package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 uint16; F1 uint16; F2 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint32; F1 int8; F2 int }
type GenS5 struct { F0 uint8 }
func genFn0(p0 uint16, p1 uint16, p2 string) uint8 {
    p2 = "hello"
    v5 := []int{873, 32768}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return uint8(p0)
}

func main() {
    v1 := make(map[string]int)
    v2 := []int{70, 32767, 9223372036854775807}
    var v3 string = "\n\t"
    v4 := make(map[uint32]bool)
    v3 = "hello"
    fmt.Println(v3)
    if len(v2) > 0 {
        v2[0] = -763
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v6 := GenS1{F0: 2, F1: 653}
    v3 = v3
    fmt.Println(v3)
    if true {
        v7 := GenS1{F0: 0, F1: 4294967295}
        var v8 float32 = (3.14159 / ((1.5 + -2.25) / (-2.25 + (0.1 * 3.14159))))
        _ = v7
        _ = v8
    } else {
        switch v3 {
        case "hello":
            _ = len(v1)
        case "foo bar":
            if true {
                if (true || false) {
                    fmt.Println(v3)
                    for fi9 := 0; fi9 < 11; fi9++ {
                        for fi10 := 0; fi10 < 13; fi10++ {
                            v11 := GenS1{F0: 1, F1: 65535}
                            _ = fi10
                            _ = v11
                        }
                        _ = fi9
                    }
                    if ("foo bar" > (v3 + "foo bar")) {
                        if true {
                            switch v3 {
                            case "\n\t":
                                var v12 int32 = 84
                                _ = v12
                            default:
                                switch (((v3 + "hello") + "\n\t") + "") {
                                case "a":
                                    v3 = ((v3 + v3) + ((((("foo bar" + v3) + v3) + "\n\t") + "\n\t") + (v3 + v3)))
                                }
                            }
                            v13 := GenS1{F0: 731, F1: 2147483647}
                            _ = v13
                        } else {
                            v3 = (("foo bar" + v3) + "")
                            for fi14 := 0; fi14 < 4; fi14++ {
                                v3 = v3
                                _ = fi14
                            }
                            for fi15 := 0; fi15 < 11; fi15++ {
                                fmt.Println(fi15)
                                var v16 bool = false
                                v17 := GenS2{F0: 36, F1: 45, F2: "\n\t"}
                                _ = fi15
                                _ = v16
                                _ = v17
                            }
                        }
                    }
                } else {
                    v18 := v1["\n\t"]
                    fmt.Println(v18, v3)
                    _ = v18
                }
                v3 = ("foo bar" + "hello")
            } else {
                if (false && (!true)) {
                    if true {
                        switch 32767 {
                        case -17:
                            var v19 int16 = 0
                            _ = v19
                        case 42:
                            v3 = ""
                        case 9223372036854775807:
                            v3 = v3
                        }
                        v3 = v3
                    } else {
                        var v20 string = ((v3 + ("" + ("foo bar" + ("hello" + (v3 + ("" + v3)))))) + ("\n\t" + ("" + v3)))
                        for fi21 := 0; fi21 < 13; fi21++ {
                            v22 := GenS3{F0: 111}
                            if ((true && false) && true) {
                                if false {
                                    v2 = append(v2, 127)
                                }
                                v20 = "foo bar"
                            }
                            _ = fi21
                            _ = v22
                        }
                        if (!(!false)) {
                            if (!false) {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                for fi23 := 0; fi23 < 11; fi23++ {
                                    v20 = "\n\t"
                                    _ = fi23
                                }
                                fmt.Println(v3, v3)
                            } else {
                                if (true && (0 > 32768)) {
                                    var v24 int = 1
                                    _ = v24
                                } else {
                                    delete(v4, 380)
                                    if (726 < 682) {
                                        _ = len(v1)
                                    }
                                    if (-541 <= -9223372036854775808) {
                                        if (!false) {
                                            if ((((3.14159 - 3.14159) / 1.0) - float32(0.1)) != -2.25) {
                                                if false {
                                                    var v25 bool = (true || (9223372036854775807 == 681))
                                                    _ = v25
                                                } else {
                                                    switch -884 {
                                                    case 0:
                                                        switch v3 {
                                                        case "":
                                                            var v26 uint = 799
                                                            _ = v26
                                                        }
                                                    case 9223372036854775807:
                                                        var v27 string = "a"
                                                        _ = v27
                                                    default:
                                                        var v28 int64 = -157
                                                        _ = v28
                                                    }
                                                }
                                            } else {
                                                var v29 string = v20
                                                fmt.Println(v20, v20)
                                                _ = v29
                                            }
                                            if true {
                                                v1[v3] = -387
                                                var v30 int32 = 1
                                                var v31 int64 = 0
                                                _ = v30
                                                _ = v31
                                            } else {
                                                if (-1 <= -9223372036854775808) {
                                                    v3 = v20
                                                }
                                            }
                                            delete(v4, 42)
                                        }
                                        v3 = v20
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    }
                                }
                                v20 = v3
                                switch 449 {
                                case 128:
                                    v32 := GenS4{F0: 4294967295, F1: 127, F2: 1}
                                    _ = v32
                                }
                            }
                        } else {
                            if len(v2) > 0 {
                                v2[0] = 0
                            }
                            for fi33 := 0; fi33 < 15; fi33++ {
                                v34 := GenS4{F0: 0, F1: 1, F2: (-(fi33))}
                                var v35 uint64 = 42
                                _ = fi33
                                _ = v34
                                _ = v35
                            }
                        }
                        _ = v20
                    }
                    var v36 float32 = 3.14159
                    v3 = ""
                    _ = v36
                } else {
                    v3 = ("" + ("\n\t" + "a"))
                }
                for fi37 := 0; fi37 < 20; fi37++ {
                    for fi38 := 0; fi38 < 7; fi38++ {
                        fmt.Println(fi37, fi37, fi37)
                        var v39 bool = false
                        _ = fi38
                        _ = v39
                    }
                    switch fi37 {
                    case 127:
                        fmt.Println(v3, fi37)
                    case 32767:
                        v3 = "foo bar"
                    }
                    if (!(true || true)) {
                        v40 := GenS5{F0: 2}
                        switch ((v3 + v3) + "foo bar") {
                        case "foo bar":
                            var v41 string = "\n\t"
                            _ = v41
                        }
                        _ = v40
                    } else {
                        var v42 int = fi37
                        switch (v3 + "hello") {
                        case "\n\t":
                            v3 = v3
                        case "hello":
                            fmt.Println(v42, fi37)
                        }
                        _ = v42
                    }
                    _ = fi37
                }
            }
        case "":
            var v43 uint8 = 0
            _ = v43
        }
        switch (("\n\t" + v3) + (v3 + (v3 + v3))) {
        case "\n\t":
            v3 = ("\n\t" + "hello")
        case "":
            v3 = ("a" + v3)
        case "a":
            for fi44 := 0; fi44 < 9; fi44++ {
                var v45 float64 = (3.14159 * 1.5)
                var v46 uint16 = 0
                if (true || true) {
                    v46 -= (^(v46))
                    v47 := GenS1{F0: uint64(v46), F1: 42}
                    _ = v47
                }
                _ = fi44
                _ = v45
                _ = v46
            }
        default:
            if true {
                v3 = "hello"
                v48 := GenS5{F0: 242}
                var v49 uint = 42
                _ = v48
                _ = v49
            }
        }
        for fi50 := 0; fi50 < 18; fi50++ {
            delete(v1, (v3 + v3))
            delete(v1, (v3 + v3))
            _ = fi50
        }
    }
    v3 = v3
    var v51 int64 = 0
    var v52 int = 928
    for fi53 := 0; fi53 < 4; fi53++ {
        v54 := GenS4{F0: 65535, F1: -30, F2: -9223372036854775808}
        if (!false) {
            var v55 float32 = 0.0
            _ = v55
        }
        var v56 float32 = 0.1
        _ = fi53
        _ = v54
        _ = v56
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v51
    _ = v52
    _ = v6
}
