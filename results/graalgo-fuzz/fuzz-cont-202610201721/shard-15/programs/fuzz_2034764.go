package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 bool; F1 uint8; F2 uint }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int; F1 int16 }
type GenS5 struct { F0 uint; F1 string }
type GenS6 struct { F0 string; F1 int; F2 int32 }
type GenS7 struct { F0 string; F1 uint; F2 bool }
type GenS8 struct { F0 int; F1 uint }
type GenS9 struct { F0 uint8 }
func genFn0() string {
    var v5 uint16 = 255
    v6 := []int{42, 42, -1, -388}
    var v7 int64 = -1
    _ = len(v6)
    _ = v5
    _ = v6
    _ = v7
    return "\n\t"
}
func genFn1(p0 uint, p1 uint8) string {
    p1 += (p1 | (p1 & p1))
    v8 := GenS1{F0: -2147483648, F1: 42}
    fmt.Println(p1, p0)
    p0 -= 270
    switch -666 {
    case -2147483648:
        p0 = p0
    default:
        switch -9223372036854775808 {
        case 128:
            var v9 uint16 = 287
            _ = v9
        case 1:
            var v10 int = 377
            _ = v10
        default:
            var v11 uint64 = 18446744073709551615
            _ = v11
        }
    }
    _ = p0
    _ = p1
    _ = v8
    return "\n\t"
}

func main() {
    var v1 float64 = 1e10
    var v2 int = 360
    var v3 bool = false
    var v4 float64 = v1
    v12 := []int{9223372036854775807, v2, 32767}
    v2 = v2
    for fi13 := 0; fi13 < 18; fi13++ {
        v1 += v1
        var v14 int32 = -33
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 19; fi15++ {
        switch fi15 {
        case 128:
            var v16 int = -1
            _ = v16
        case 1:
            if (2147483647 == 2147483647) {
                for fi17 := 0; fi17 < 20; fi17++ {
                    if ((-(fi15)) == 264) {
                        v18 := GenS2{F0: false, F1: 215, F2: 255}
                        v19 := GenS2{F0: false, F1: 0, F2: 99}
                        var v20 int32 = 606
                        v21 := GenS2{F0: v3, F1: 254, F2: 179}
                        _ = v18
                        _ = v19
                        _ = v20
                        _ = v21
                    } else {
                        v22 := GenS1{F0: (^(fi17)), F1: 440}
                        _ = v22
                    }
                    v1 -= v1
                    if len(v12) >= 2 {
                        v12 = v12[1:len(v12)]
                    }
                    if false {
                        _ = len(v12)
                    } else {
                        if (((-1 >= 42) || ("hello" >= (("" + "hello") + ""))) || (65535 < 255)) {
                            fmt.Println(fi15, fi17)
                            var v23 string = ("hello" + "a")
                            if true {
                                v24 := GenS2{F0: ((1 >= 2) && false), F1: 116, F2: 677}
                                if (!v3) {
                                    switch (("\n\t" + v23) + "\n\t") {
                                    case "foo bar":
                                        var v25 uint16 = 65535
                                        _ = v25
                                    case "":
                                        if (v23 < "foo bar") {
                                            var v26 int32 = -2147483648
                                            v27 := GenS3{F0: (v26 >> ((4294967295) & 63))}
                                            _ = v26
                                            _ = v27
                                        } else {
                                            var v28 string = (((("hello" + (v23 + v23)) + v23) + "hello") + (("a" + (v23 + v23)) + (v23 + v23)))
                                            switch "hello" {
                                            case "hello":
                                                var v29 bool = v3
                                                _ = v29
                                            case "a":
                                                v2 = (fi15 / ((fi17) | 1))
                                            }
                                            _ = v28
                                        }
                                    }
                                    switch "hello" {
                                    case "a":
                                        var v30 int = v2
                                        _ = v30
                                    default:
                                        v31 := GenS2{F0: v3, F1: 255, F2: 4294967295}
                                        _ = v31
                                    }
                                    var v32 uint8 = 127
                                    _ = v32
                                }
                                _ = v24
                            }
                            _ = v23
                        }
                    }
                    _ = fi17
                }
            } else {
                if true {
                    if ((("" + "\n\t") + ("a" + ("hello" + "foo bar"))) > "\n\t") {
                        var v33 uint32 = 4294967295
                        var v34 uint16 = 2
                        _ = v33
                        _ = v34
                    }
                }
            }
        }
        var v35 uint = 42
        if v3 {
            switch "hello" {
            case "":
                v1 = 3.14159
            case "hello":
                for fi36 := 0; fi36 < 1; fi36++ {
                    var v37 uint16 = 32767
                    _ = fi36
                    _ = v37
                }
            default:
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
            }
        }
        _ = fi15
        _ = v35
    }
    switch "" {
    case "hello":
        if v3 {
            var v38 uint8 = 42
            switch "a" {
            case "":
                if (0.0 >= 0.1) {
                    switch (("\n\t" + "foo bar") + "hello") {
                    case "a":
                        if true {
                            v39 := GenS4{F0: (-(v2)), F1: -32768}
                            if v3 {
                                var v40 uint = 0
                                v3 = (592 != 1)
                                _ = v40
                            } else {
                                if ((-2.25 - (0.1 - 1e10)) > 1e10) {
                                    v12 = append(v12, 1)
                                    v41 := GenS5{F0: 364, F1: "foo bar"}
                                    _ = v41
                                }
                                v12 = append(v12, 127)
                            }
                            _ = v39
                        } else {
                            v42 := GenS3{F0: 0}
                            var v43 string = "a"
                            var v44 int = v2
                            _ = v42
                            _ = v43
                            _ = v44
                        }
                    case "\n\t":
                        switch v2 {
                        case -909:
                            var v45 string = "\n\t"
                            _ = v45
                        case 127:
                            switch "foo bar" {
                            case "a":
                                for fi46 := 0; fi46 < 18; fi46++ {
                                    var v47 uint = 42
                                    v48 := GenS6{F0: "hello", F1: fi46, F2: -2147483648}
                                    _ = fi46
                                    _ = v47
                                    _ = v48
                                }
                            default:
                                var v49 int32 = 42
                                _ = v49
                            }
                        }
                    default:
                        for fi50 := 0; fi50 < 10; fi50++ {
                            var v51 uint16 = 0
                            _ = fi50
                            _ = v51
                        }
                    }
                }
            default:
                var v52 int16 = -32768
                _ = v52
            }
            _ = v38
        } else {
            fmt.Println(v2)
            v53 := GenS7{F0: ("\n\t" + ("" + "foo bar")), F1: 255, F2: (-128 != 0)}
            v54 := GenS8{F0: v2, F1: 284}
            _ = v53
            _ = v54
        }
    case "\n\t":
        v2 = -9223372036854775808
    case "foo bar":
        v3 = (!(v3 && (-1 <= v2)))
    }
    var v55 int8 = 1
    var v56 int8 = -118
    var v57 int = v2
    for fi58 := 0; fi58 < 17; fi58++ {
        if (1.5 > 3.14159) {
            v2 += fi58
            v1 += v4
        } else {
            for fi59 := 0; fi59 < 7; fi59++ {
                v2 += v2
                _ = fi59
            }
            v3 = ("\n\t" != "\n\t")
        }
        for fi60 := 0; fi60 < 1; fi60++ {
            v12 = append(v12, fi60)
            _ = fi60
        }
        _ = fi58
    }
    if false {
        if false {
            for fi61 := 0; fi61 < 18; fi61++ {
                if len(v12) > 0 {
                    v12[0] = (v2 - 4294967295)
                }
                _ = fi61
            }
        } else {
            switch (v57 >> ((14) & 63)) {
            case 9223372036854775807:
                if len(v12) > 0 {
                    v12[0] = (v2 / (((v2 / ((v2) | 1))) | 1))
                }
            case 420:
                v57 = -907
            case 4294967295:
                var v62 int = (v2 >> ((1) & 63))
                _ = v62
            default:
                v63 := GenS9{F0: 255}
                _ = v63
            }
        }
        for fi64 := 0; fi64 < 2; fi64++ {
            var v65 int64 = -9223372036854775808
            _ = fi64
            _ = v65
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v55
    _ = v56
    _ = v57
}
