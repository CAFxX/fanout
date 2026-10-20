package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool; F1 uint64; F2 int16 }
type GenS3 struct { F0 uint16; F1 string }
type GenS4 struct { F0 int32; F1 uint8; F2 int32 }
type GenS5 struct { F0 int32 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint16 = 626
    v3 := []string{("hello" + "foo bar")}
    v4 := make(map[string]uint32)
    var v5 uint16 = 0
    var v6 int16 = 0
    fmt.Println(v6, v6)
    switch 4294967295 {
    case 378:
        fmt.Println(v5, v5)
    }
    if true {
        switch 128 {
        case 242:
            v3 = append(v3, "foo bar")
        case 42:
            if len(v3) > 0 {
                v3[0] = ""
            }
        }
    }
    if (!true) {
        v7 := GenS1{F0: 1}
        v5 = (v2 << ((65535) & 63))
        _ = v7
    } else {
        fmt.Println(v6)
        var v8 float64 = 1.5
        _ = v8
    }
    v2 = 2
    for fi9 := 0; fi9 < 14; fi9++ {
        v10 := GenS1{F0: 2147483647}
        if (fi9 > 4294967295) {
            for fi11 := 0; fi11 < 18; fi11++ {
                var v12 int = fi9
                switch ("\n\t" + ("" + "a")) {
                case "foo bar":
                    switch "foo bar" {
                    case "hello":
                        if false {
                            var v13 bool = true
                            _ = v13
                        }
                    default:
                        var v14 string = "\n\t"
                        _ = v14
                    }
                }
                var v15 int = (v12 * v12)
                _ = fi11
                _ = v12
                _ = v15
            }
        }
        _ = fi9
        _ = v10
    }
    switch "\n\t" {
    case "\n\t":
        for fi16 := 0; fi16 < 2; fi16++ {
            v6 = (v6 - -677)
            v5 = v5
            _ = fi16
        }
    }
    switch 127 {
    case 128:
        switch "a" {
        case "":
            v5 = 32767
        case "a":
            if true {
                var v17 int8 = -128
                if true {
                    var v18 int32 = -29
                    if (0 <= 2147483647) {
                        var v19 int16 = -1
                        v17 = 87
                        var v20 string = "foo bar"
                        _ = v19
                        _ = v20
                    }
                    var v21 uint8 = 2
                    _ = v18
                    _ = v21
                } else {
                    for fi22 := 0; fi22 < 14; fi22++ {
                        fmt.Println(v6)
                        v23 := GenS2{F0: true, F1: 0, F2: (v6 << ((4294967295) & 63))}
                        _ = fi22
                        _ = v23
                    }
                    switch 32767 {
                    case -838:
                        v4["a"] = 42
                    }
                }
                _ = v17
            } else {
                switch -737 {
                case -187:
                    v24 := GenS1{F0: -987}
                    _ = v24
                }
            }
        default:
            if false {
                switch "a" {
                case "a":
                    v25 := GenS3{F0: (v5 + (v2 * (^(v2)))), F1: ("foo bar" + "foo bar")}
                    _ = v25
                case "\n\t":
                    var v26 float32 = 1e10
                    _ = v26
                }
            } else {
                v2 = v2
                for fi27 := 0; fi27 < 8; fi27++ {
                    for fi28 := 0; fi28 < 12; fi28++ {
                        switch fi28 {
                        case 1:
                            switch "" {
                            case "a":
                                if false {
                                    var v29 uint8 = 104
                                    _ = v29
                                } else {
                                    v6 = v6
                                    var v30 float64 = (((0.1 - -2.25) - -2.25) + 0.1)
                                    if len(v3) > 0 {
                                        v3[0] = ""
                                    }
                                    _ = v30
                                }
                            case "\n\t":
                                var v31 string = "\n\t"
                                _ = v31
                            case "":
                                v5 = v5
                            }
                        case 2147483647:
                            if false {
                                var v32 bool = (-2.25 <= (float64(((1e10 + 3.14159) - 3.14159)) / float64((1.5 * 1.5))))
                                _ = v32
                            } else {
                                var v33 uint32 = 612
                                fmt.Println(v6, v6, fi27)
                                v5 -= v2
                                _ = v33
                            }
                        case 9223372036854775807:
                            if false {
                                v34 := GenS4{F0: 2147483647, F1: 255, F2: -1}
                                v2 += v2
                                _ = v34
                            }
                            if ((^(fi27)) >= fi27) {
                                var v35 string = "\n\t"
                                var v36 bool = (42 == 42)
                                _ = v35
                                _ = v36
                            }
                            var v37 int8 = 0
                            _ = v37
                        }
                        if len(v3) > 0 {
                            v3[0] = ""
                        }
                        _ = fi28
                    }
                    var v38 int16 = v6
                    _ = fi27
                    _ = v38
                }
            }
        }
    case 4294967295:
        v2 += v5
    case -9223372036854775808:
        v5 += (v5 ^ 65535)
    default:
        var v39 int64 = -422
        _ = v39
    }
    v3 = append(v3, "foo bar")
    var v40 float32 = float32(0.0)
    v41 := GenS4{F0: 2147483647, F1: 128, F2: 812}
    v6 = (v6 ^ (v6 / ((v6) | 1)))
    if false {
        if (false || (-2147483648 >= -2147483648)) {
            v42 := GenS3{F0: (v5 * v5), F1: "hello"}
            v1[324] = ((float32(3.14159) / 1.0) <= -2.25)
            fmt.Println(v6, v40, v6)
            _ = v42
        } else {
            if (!(!(!(((((3.14159 / 0.1) * 1.5) / v40) / 1.0) > 0.0)))) {
                switch "foo bar" {
                case "foo bar":
                    var v43 int = -558
                    _ = v43
                }
                var v44 int16 = v6
                var v45 bool = false
                _ = v44
                _ = v45
            } else {
                var v46 uint16 = (v2 - v2)
                _ = v46
            }
        }
        v5 = 0
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    switch 127 {
    case 127:
        for fi47 := 0; fi47 < 2; fi47++ {
            var v48 uint16 = 815
            v6 = (v6 % ((v6) | 1))
            if (0 <= 1) {
                v5 += v2
            }
            _ = fi47
            _ = v48
        }
    default:
        var v49 int32 = 1
        _ = v49
    }
    switch 9223372036854775807 {
    case 655:
        delete(v1, 65535)
    case 127:
        if (-990 != -974) {
            var v50 uint8 = uint8(v2)
            v6 = (v6 + int16(v5))
            switch "a" {
            case "hello":
                v2 = (v5 % ((v2) | 1))
            case "foo bar":
                var v51 bool = true
                _ = v51
            case "":
                var v52 uint32 = 2147483647
                _ = v52
            default:
                var v53 int16 = 32767
                _ = v53
            }
            _ = v50
        } else {
            for fi54 := 0; fi54 < 9; fi54++ {
                var v55 int64 = 9223372036854775807
                v56 := GenS5{F0: 0}
                v40 = 3.14159
                _ = fi54
                _ = v55
                _ = v56
            }
            v5 = v5
        }
    case 1:
        var v57 uint32 = 42
        _ = v57
    default:
        v3 = append(v3, ("a" + ""))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
